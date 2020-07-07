### socket
```
#include <sys/socket.h>
int socket(int domain, int type, int protocol)

//domain(域)
AF_INET     IPv4
AF_INET6    IPv6
AF_UNIX     UNIX域

//type
SOCK_DGRAM  UDP(固定长度、无连接、不可靠报文传递)
SOCK_RAM    ip协议数据报接口，用于直接访问网络层，绕过传输层(tcp、udp)，需要超级用户特权
SOCK_STREAM TCP(有序、可靠、双向、面向连接字节流)

//protocol
0           表示为给定的域和套接字选择默认协议
```

### 字节序
```
小端/大端         大端
处理器字节序 ---> 网络字节序

#include <arpa/inet.h>
uint32_t htonl(uint32_t hostint32)  返回 网络字节序32位整数
uint16_t htons(uint16_t hostint16)  返回 网络字节序16位整数
uint32_t ntonl(uint32_t netint32)   返回 主机字节序32位整数
uint16_t ntons(uint16_t netint16)   返回 主机字节序16位整数
```

### 地址
```
1、通用socket地址
    不同的地址格式必须转换为此格式
    #include <sys/socket.h>
    struct sockaddr {
        sa_family_t  sa_family;     //地址族 unsigned short, AF_xxx
        char         sa_data[14];   //14字节 包含套接字中的目标地址和端口信息     
    }
    // 大小16个字节
2、专用socket地址
    (1) IPv4
        #include<netinet/in.h>
        typedef uint16_t in_port_t;
        typedef uint32_t in_addr_t;
        struct sockaddr_in {    
            sa_family_t     sin_family;     // 地址族 AF_INET
            in_port_t       sin_port;       // 16位端口号
            struct in_addr  sin_addr;       // 32位IP地址
            unsigned char   sin_zero[8];    // Same size as struct sockaddr，补齐剩余的字符
        }
        struct in_addr {
            in_addr_t       s_addr           // 32位IPv4地址；A.B.C.D 
        }

    (2) IPv6
        struct sockaddr_in6 { 
            sa_family_t     sin6_family;    // AF_INET6
            in_port_t       sin6_port;      // 16位端口号
            uint32_t        sin6_flowinfo;  // IPv6 flow information
            struct in6_addr sin6_addr;      // IPv6 address
            uint32_t        sin6_scope_id;  // 
        }
        struct in6_addr { 
            unsigned char   s6_addr[16];    // 128位IPv6地址长度；XXXX:XXXX:XXXX:XXXX:XXXX:XXXX:XXXX:XXXX
        }

    (3) unix域套接字地址
        #include <sys/un.h>
        struct sockaddr_un {
			sa_family_t     sun_family;     // AF_UNIX
            char            sun_path[108];  // pathname
        }

3、addr转换
    (1) tcp套接字转换
		struct sockaddr_in my_addr;
		my_addr.sin_family      = AF_INET;
		my_addr.sin_port        = htons(80);                 //uint16转换成网络字节序
		my_addr.sin_addr.s_addr = inet_addr("192.168.2.201") //inet_addr将字符串转换为网络字节序，inet_ntoa则将网络字节序转换为字符串
		bzero(&(my_addr.sin_zero), 8);                       //sin_zero置0
		struct sockaddr* myaddr = (struct sockaddr*)&my_addr //转换成sockaddr

	(2) unix域套接字转换
		struct sockaddr_un un;
		memset(&un, 0, sizeof(un));
		un.sun_family = AF_UNIX;
		strcpy(un.sun_path, "foo.socket");
		struct sockaddr *myaddr = (struct sockaddr*)&un;
	
	// socket绑定addr
	if((fd = socket(AF_UNIX, SOCK_STREAM, 0)) < 0)
		err_sys("socket failed");
	if(bind(fd, myaddr, sizeof(myaddr)) < 0)
		ERR_EXIT("bind");
```

### accept
```
https://www.cnblogs.com/wangcq/p/3520400.html

TCP服务器端依次调用socket()、bind()、listen()之后，就会监听指定的socket地址了
TCP客户端依次调用socket()、connect()之后就向服务器发送了一个连接请求
TCP服务器监听到这个请求之后，就会调用accept()函数取接收请求，这样连接就建立好了。之后就可以开始网络I/O操作了，即类同于普通文件的读写I/O操作

int accept(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
// sockfd 	服务器的socket描述字
// addr		客户端的协议地址
// addrlen	第三个参数为协议地址的长度
// 返回		生成一个全新的描述字，代表与返回客户的TCP连接

// 内核为每个由服务器进程接受的客户连接创建了一个已连接socket描述字，当服务器完成了对某个客户的服务，相应的已连接socket描述字就被关闭

// 三次握手发生在这一步
```

### read/write的返回
```
(1) 对于阻塞socket
	read:	读缓冲区没有数据，发生阻塞
	write:	写缓冲区满了，发生阻塞
    如果返回-1代表网络出错了
(2) 对于非阻塞socket
    不能read或write时，就会返回-1，同时errno设置为EAGAIN(再试一次)
```

### tcp输出
```
用户空间        内核
应用   ----->   TCP  ------------------------------>    IP  ----------->  输出
缓冲区          套接字发送缓冲区（SO_SNDBUF）             MTU大小分组
                MSS大小的TCP分节
                通常MSS<=MTU-40(IPv4) 或 MTU-60(IPv6)
```

### socket IO
```
(1) 读事件
    1) 句柄从不可读变成可读，或者句柄写缓冲区有新的数据进来且超过SO_RCVLOWAT
    2) 产生事件的情况
        > socket有一个未清除的错误。如非阻塞的connect连接错误会使socket变成可读写状态
        > 非阻塞accept有新的连接进来
        > socket写对端关闭，read返回0
        > socket读缓冲区有新的数据进来且超过SO_RCVLOWAT

(2) 写事件
    1) 句柄从不可写变成可写，或者句柄写缓冲区有新的数据进来而且缓冲区水位高于SO_SNDLOWAT
    2) 产生事件的情况
        > socket有一个未清除的错误。例如非阻塞connect连接出错会导致socket变成可读可写状态
        > 非阻塞connect连接成功后端口状态会变成可写
        > socket读对端关闭，socket变成可写状态，产生SIGPIPE信号
        > socket写缓冲区有新的数据进来且超过SO_SNDLOWAT
    
在epoll中，读事件对应EPOLLIN，写事件对应EPOLLOUT
```

### 套接字关联的选项
```
#include <sys/socket.h>

int setsockopt(int socket, int level, int option_name,const void *option_value, socklen_t option_len);
int getsockopt(int socket, int level, int option_name, void *option_value, socklen_t *option_len);

// socket		套接字
// level		所在的协议层，一般设置SOL_SOCKET
// option_name	设置的选项，选项如下
				SO_DEBUG 		打开或关闭排错模式
				SO_REUSEADDR 	允许在bind()过程中本地地址可重复使用
				SO_TYPE 		返回socket形态
				SO_ERROR 		返回socket已发生的错误原因
				SO_DONTROUTE 	送出的数据包不要利用路由设备来传输
				SO_BROADCAST 	使用广播方式传送
				SO_SNDBUF 		设置送出的暂存区大小
				SO_RCVBUF 		设置接收的暂存区大小
				SO_KEEPALIVE 	定期确定连线是否已终止
				SO_OOBINLINE	当接收到OOB数据时会马上送至标准输入设备
				SO_LINGER		确保数据安全且可靠的传送出去
// option_value	代表欲设置的值
// option_len	则为option_value的长度
// 返回值		成功则返回0, 错误返回-1, 错误原因存于errno


```

### SO_REUSEADDR和SO_REUSEPORT
```
https://zhuanlan.zhihu.com/p/35367402

(1) SO_REUSEADDR
    1) 设置套接字属性
        setsockopt(listenfd, SOL_SOCKET, SO_REUSEADDR,(const void *)&reuse , sizeof(int));
    2) 目的
        当服务端出现timewait状态的链接时，确保server能够重启成功

(2) SO_REUSEPORT
    1) 使用场景
        linux kernel 3.9 引入了最新的SO_REUSEPORT选项，使得多进程或者多线程创建多个绑定同一个ip:port的监听socket，提高服务器的接收链接的并发能力,程序的扩展性更好；此时需要设置SO_REUSEPORT（注意所有进程都要设置才生效）
        setsockopt(listenfd, SOL_SOCKET, SO_REUSEPORT,(const void *)&reuse , sizeof(int));
    2) 目的
        每一个进程有一个独立的监听socket，并且bind相同的ip:port，独立的listen()和accept()；提高接收连接的能力。
        nginx新版本是多进程同时监听同一个ip:port（每个进程bind同一个ip:port，但是只有一个进程会得到响应
    3) 解决的问题
        > 避免了应用层多线程或者进程监听同一ip:port的“惊群效应”。
        > 内核层面实现负载均衡，保证每个进程或者线程接收均衡的连接数。
        > 只有effective-user-id相同的服务器进程才能监听同一ip:port （安全性考虑）

(3) 示例
    #include <stdlib.h>
    #include <string.h>
    #include <netinet/in.h>
    #include <sys/socket.h>
    #include <arpa/inet.h>
    #include <sys/types.h>
    #include <errno.h>
    #include <time.h>
    #include <unistd.h>
    #include <sys/wait.h>
    void work () {
        int listenfd = socket(AF_INET, SOCK_STREAM, 0);
        if (listenfd < 0) {
            perror("listen socket");
            _exit(-1);
        }
        int ret = 0;
        int reuse = 1;
        ret = setsockopt(listenfd, SOL_SOCKET, SO_REUSEADDR,(const void *)&reuse , sizeof(int));
        if (ret < 0) {
            perror("setsockopt");
            _exit(-1);
        }
        ret = setsockopt(listenfd, SOL_SOCKET, SO_REUSEPORT,(const void *)&reuse , sizeof(int));
        if (ret < 0) {
            perror("setsockopt");
            _exit(-1);
        }
        struct sockaddr_in addr;
        memset(&addr, 0, sizeof(addr));
        addr.sin_family = AF_INET;
        //addr.sin_addr.s_addr = inet_addr("10.95.118.221");
        addr.sin_addr.s_addr = inet_addr("0.0.0.0");                                                                             
        addr.sin_port = htons(9980);
        ret = bind(listenfd, (struct sockaddr *)&addr, sizeof(addr));
        if (ret < 0) {
            perror("bind addr");
            _exit(-1);
        }
        printf("bind success\n");
        ret = listen(listenfd,10);
        if (ret < 0) {
            perror("listen");
            _exit(-1);
        }
        printf("listen success\n");
        struct sockaddr clientaddr;
        int len = 0;
        while(1) {
            printf("process:%d accept...\n", getpid());
            int clientfd = accept(listenfd, (struct sockaddr*)&clientaddr, &len);
            if (clientfd < 0) {
                printf("accept:%d %s", getpid(),strerror(errno));
                _exit(-1);
            }
            close(clientfd);
            printf("process:%d close socket\n", getpid());
        }
    }
    int main(){
        printf("uid:%d euid:%d\n", getuid(),geteuid());
        int i = 0;
        for (i = 0; i< 6; i++) {
            pid_t pid = fork();
            if (pid == 0) {
                work();
            }
            if(pid < 0) {
                perror("fork");
                continue;
            }
        }
        int status,id;
        while((id=waitpid(-1, &status, 0)) > 0) {
            printf("%d exit\n", id);
        }
        if(errno == ECHILD) {
            printf("all child exit\n");
        }
        return 0;
    } 

```

## IO多路复用
### epoll
1、原理
```
基于事件驱动的I/O方式
select #(select实例)[https://blog.csdn.net/u010155023/article/details/53507788]
epoll #(epoll实例)[https://blog.csdn.net/davidsguo008/article/details/73556811]
https://yq.aliyun.com/articles/683282[关于epoll的IO模型是同步异步的一次纠结过程]
https://segmentfault.com/a/1190000003063859
https://www.cnblogs.com/lojunren/p/3856290.html
https://www.cnblogs.com/lojunren/p/3856290.html❗
http://blog.chinaunix.net/uid-28541347-id-4273856.html❗
https://blog.csdn.net/dog250/article/details/80837278❗

(1) 重要概念
    1) socket对象
        由文件系统管理的，包含了发送缓冲区、接收缓冲区、等待队列
    2) epoll对象
        调用epoll_create方法时，内核会创建一个eventpoll对象，eventpoll对象也是文件系统中的一员，和socket一样，它也会有等待队列。epoll对象时进程与socket之间的中介，包含了监视队列和就绪列表
        struct eventpoll{
            ....
            /*红黑树的根节点，这颗树中存储着所有添加到epoll中的需要监控的事件*/
            struct rb_root  rbr;
            /*双链表中则存放着将要通过epoll_wait返回给用户的满足条件的事件*/
            struct list_head rdlist;
            ....
        };
        每一个epoll对象都有一个独立的eventpoll结构体，用于存放通过epoll_ctl方法向epoll对象中添加进来的事件。这些事件都会挂载在红黑树中,如此，重复添加的事件就可以通过红黑树而高效的识别出来(红黑树的插入时间效率是lgn，其中n为树的高度)。
    3) 等待队列
        创建socket时，会创建socket文件对象，里面包含一个非常重要的结构（等待队列），它指向所有需要等待该socket事件的进程。
    4) 监视队列(rbr 红黑树)
        用epoll_ctl添加或删除所要监听的socket，内核会将eventpoll添加到目标socket的等待队列中
    5) epitem对象
        监视队列和就绪列表的基本结构，每一个要监听的fd对应一个此结构体   
        所有添加到epoll中的事件都会与设备(网卡)驱动程序建立回调关系，也就是说，当相应的事件发生时会调用这个回调方法。这个回调方法在内核中叫ep_poll_callback,它会将发生的事件添加到rdlist双链表中。
        在epoll中，对于每一个事件，都会建立一个epitem结构体，如下所示：
        struct epitem{
            struct rb_node  rbn;//红黑树节点
            struct list_head    rdllink;//双向链表节点
            struct epoll_filefd  ffd;  //事件句柄信息
            struct eventpoll *ep;    //指向其所属的eventpoll对象
            struct epoll_event event; //期待发生的事件类型
        } 
        当调用epoll_wait检查是否有事件发生时(所以epoll不是异步)，只需要检查eventpoll对象中的rdlist双链表中是否有epitem元素即可。如果rdlist不为空，则把发生的事件复制到用户态，同时将事件数量返回给用户。
    6) 就绪列表(rdlist 双向链表)
        当socket收到数据后（由不可读变成可或由不可写变成可写，见上述），中断函数将添加socket的引用到就绪列表，另一方面唤醒eventpoll等待队列中的进程
        //注意：这是epoll不会引起惊群的关键步骤，中断除了将socket引用添加到就绪列表后，还会
    7) ET和FT
        https://www.jianshu.com/p/d3442ff24ba6
        http://kimi.it/515.html
        使用ET的例子 nginx; 使用LT的例子 redis
        > ET(不到边缘情况，是死都不会触发的)
            > 特点
                句柄在发生读写事件时只会通知用户一次,
                ET模式主要关注fd从不可用到可用或者可用到不可用的情况。
                ET只支持非阻塞模式。
                ET模式下读写操作要时用while循环，直到读/写够足够多的数据，或者读/写到返回EAGAIN。尤其时在写大块数据时，一次write操作不足以写完全部数据，或者在读大块数据时，应用层缓冲区数据太小，一次read操作不足以读完全部数据，应用层要么一直调用while循环一直IO到EGAIN,或者自己调用epoll_ctl手动触发ET响应。
        > LT
            > 特点
                只要句柄一直处于可用状态，就会一直通知用户。
                LT模式下，句柄读缓冲区被读空后，句柄会从可用转变未不可以用，这个时候不会通知用户。写缓冲区只要还没写满，就会一直通知用户。
                LT模式支持阻塞和非阻塞两种方式。epoll默认的模式是LT。
                LT下，应用层的业务逻辑比较简单，更不容易遗漏事件，更不容易出错。通常，在将数据写完后，我们会关闭句柄的写事件。
        > ET vs LT
            LT模式下，只要一个句柄上的事件一次没有处理完，会在以后调用epoll_wait时次次返回这个句柄，而ET模式仅在第一次返回。
            这件事怎么做到的呢？当一个socket句柄上有事件时，内核会把该句柄插入上面所说的准备就绪list链表，这时我们调用epoll_wait，会把准备就绪的socket拷贝到用户态内存，然后清空准备就绪list链表，最后，epoll_wait干了件事，就是检查这些socket，如果不是ET模式（就是LT模式的句柄了），并且这些socket上确实有未处理的事件时，又把该句柄放回到刚刚清空的准备就绪链表了。所以，非ET的句柄，只要它上面还有事件，epoll_wait每次都会返回。而ET模式的句柄，除非有新中断到，即使socket上的事件没有处理完，也是不会次次从epoll_wait返回的。
        > 使用建议
            对于监听的sockfd，最好使用水平触发模式，边缘触发模式会导致高并发情况下，有的客户端会连接不上。如果非要使用边缘触发，可以用 while 来循环 accept()。
            对于读写的 connfd，水平触发模式下，阻塞和非阻塞效果都一样，建议设置非阻塞。
            对于读写的 connfd，边缘触发模式下，必须使用非阻塞 IO，并要求一次性地完整读写全部数据。



```
2、函数
```
https://blog.csdn.net/ljx0305/article/details/4065058

(1) int epoll_create(int size);
    创建一个epoll的句柄，size用来告诉内核这个监听的数目一共有多大。这个参数不同于select()中的第一个参数，给出最大监听的fd+1的值。需要注意的是，当创建好epoll句柄后，它就是会占用一个fd值，在linux下如果查看/proc/进程id/fd/，是能够看到这个fd的，所以在使用完epoll后，必须调用close()关闭，否则可能导致fd被耗尽。


(2) int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
    epoll的事件注册函数，它不同与select()是在监听事件时告诉内核要监听什么类型的事件，而是在这里先注册要监听的事件类型。第一个参数是epoll_create()的返回值，第二个参数表示动作，用三个宏来表示：
    EPOLL_CTL_ADD：注册新的fd到epfd中；
    EPOLL_CTL_MOD：修改已经注册的fd的监听事件；
    EPOLL_CTL_DEL：从epfd中删除一个fd；
    第三个参数是需要监听的fd，第四个参数是告诉内核需要监听什么事，struct epoll_event结构如下：

    typedef union epoll_data {
        void *ptr;
        int fd;
        __uint32_t u32;
        __uint64_t u64;
    } epoll_data_t;

    struct epoll_event {
        __uint32_t events; /* Epoll events */
        epoll_data_t data; /* User data variable */
    };

    events可以是以下几个宏的集合：
    EPOLLIN ：表示对应的文件描述符可以读（包括对端SOCKET正常关闭）；
    EPOLLOUT：表示对应的文件描述符可以写；
    EPOLLPRI：表示对应的文件描述符有紧急的数据可读（这里应该表示有带外数据到来）；
    EPOLLERR：表示对应的文件描述符发生错误；
    EPOLLHUP：表示对应的文件描述符被挂断；
    EPOLLET： 将EPOLL设为边缘触发(Edge Triggered)模式，这是相对于水平触发(Level Triggered)来说的。
    EPOLLONESHOT：
        作用：对于注册了EPOLLONESHOT事件的文件描述符，操作系统最多出发其上注册的一个可读，可写或异常事件，且只能触发一次
        使用：注册了EPOLLONESHOT事件的socket一旦被某个线程处理完毕，该线程就应该立即重置这个socket上的EPOLLONESHOT事件，以确保这个socket下一次可读时，其EPOLLIN事件能被触发，进而让其他工作线程有机会继续处理这个sockt。

(3) int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
    等待事件的产生，类似于select()调用。参数events用来从内核得到事件的集合，maxevents告之内核这个events有多大，这个 maxevents的值不能大于创建epoll_create()时的size，参数timeout是超时时间（毫秒，0会立即返回，-1将不确定，也有说法说是永久阻塞）。该函数返回需要处理的事件数目，如返回0表示已超时。

```
3、示例
```
#define IPADDRESS   "127.0.0.1"
#define PORT        8787
#define MAXSIZE     1024
#define LISTENQ     5
#define FDSIZE      1000
#define EPOLLEVENTS 100

listenfd = socket_bind(IPADDRESS,PORT);

struct epoll_event events[EPOLLEVENTS];

//创建一个描述符
epollfd = epoll_create(FDSIZE);

//添加监听描述符事件
add_event(epollfd,listenfd,EPOLLIN);

//循环等待
for ( ; ; ){
    //该函数返回已经准备好的描述符事件数目
    ret = epoll_wait(epollfd,events,EPOLLEVENTS,-1);
    //处理接收到的连接
    handle_events(epollfd,events,ret,listenfd,buf);
}

//事件处理函数
static void handle_events(int epollfd,struct epoll_event *events,int num,int listenfd,char *buf)
{
     int i;
     int fd;
     //进行遍历;这里只要遍历已经准备好的io事件。num并不是当初epoll_create时的FDSIZE。
     for (i = 0;i < num;i++)
     {
         fd = events[i].data.fd;
        //根据描述符的类型和事件类型进行处理
         if ((fd == listenfd) &&(events[i].events & EPOLLIN))
            handle_accpet(epollfd,listenfd);
         else if (events[i].events & EPOLLIN)
            do_read(epollfd,fd,buf);
         else if (events[i].events & EPOLLOUT)
            do_write(epollfd,fd,buf);
     }
}

//添加事件
static void add_event(int epollfd,int fd,int state){
    struct epoll_event ev;
    ev.events = state;
    ev.data.fd = fd;
    epoll_ctl(epollfd,EPOLL_CTL_ADD,fd,&ev);
}

//处理接收到的连接
static void handle_accpet(int epollfd,int listenfd){
     int clifd;     
     struct sockaddr_in cliaddr;     
     socklen_t  cliaddrlen;     
     clifd = accept(listenfd,(struct sockaddr*)&cliaddr,&cliaddrlen);     
     if (clifd == -1)         
     perror("accpet error:");     
     else {         
         printf("accept a new client: %s:%d\n",inet_ntoa(cliaddr.sin_addr),cliaddr.sin_port);                       //添加一个客户描述符和事件         
         add_event(epollfd,clifd,EPOLLIN);     
     } 
}

//读处理
static void do_read(int epollfd,int fd,char *buf){
    int nread;
    nread = read(fd,buf,MAXSIZE);
    if (nread == -1)     {         
        perror("read error:");         
        close(fd); //记住close fd        
        delete_event(epollfd,fd,EPOLLIN); //删除监听 
    }
    else if (nread == 0)     {         
        fprintf(stderr,"client close.\n");
        close(fd); //记住close fd       
        delete_event(epollfd,fd,EPOLLIN); //删除监听 
    }     
    else {         
        printf("read message is : %s",buf);        
        //修改描述符对应的事件，由读改为写         
        modify_event(epollfd,fd,EPOLLOUT);     
    } 
}

//写处理
static void do_write(int epollfd,int fd,char *buf) {     
    int nwrite;     
    nwrite = write(fd,buf,strlen(buf));     
    if (nwrite == -1){         
        perror("write error:");        
        close(fd);   //记住close fd       
        delete_event(epollfd,fd,EPOLLOUT);  //删除监听    
    }else{
        modify_event(epollfd,fd,EPOLLIN); 
    }    
    memset(buf,0,MAXSIZE); 
}

//删除事件
static void delete_event(int epollfd,int fd,int state) {
    struct epoll_event ev;
    ev.events = state;
    ev.data.fd = fd;
    epoll_ctl(epollfd,EPOLL_CTL_DEL,fd,&ev);
}

//修改事件
static void modify_event(int epollfd,int fd,int state){     
    struct epoll_event ev;
    ev.events = state;
    ev.data.fd = fd;
    epoll_ctl(epollfd,EPOLL_CTL_MOD,fd,&ev);
}

几乎所有的epoll程序都使用下面的框架：
for( ; ; )
    {
        nfds = epoll_wait(epfd,events,20,500);
        for(i=0;i<nfds;++i)
        {
            if(events[i].data.fd==listenfd)                                 //有新的连接
            {
                connfd = accept(listenfd,(sockaddr *)&clientaddr, &clilen); //accept这个连接
                ev.data.fd=connfd;
                ev.events=EPOLLIN|EPOLLET;
                epoll_ctl(epfd,EPOLL_CTL_ADD,connfd,&ev);                   //将新的fd添加到epoll的监听队列中
            }
            else if( events[i].events&EPOLLIN )                             //接收到数据，读socket
            {
                n = read(sockfd, line, MAXLINE)) < 0                        //读
                ev.data.ptr = md;                                           //md为自定义类型，添加数据
                ev.events=EPOLLOUT|EPOLLET;
                epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev);                   //修改标识符，等待下一个循环时发送数据，异步处理的精髓
            }
            else if(events[i].events&EPOLLOUT)                              //有数据待发送，写socket
            {
                struct myepoll_data* md = (myepoll_data*)events[i].data.ptr;//取数据
                sockfd = md->fd;
                send( sockfd, md->ptr, strlen((char*)md->ptr), 0 );         //发送数据
                ev.data.fd=sockfd;
                ev.events=EPOLLIN|EPOLLET;
                epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev); //修改标识符，等待下一个循环时接收数据
            }
            else
            {
                //其他的处理
            }
        }
    }
```

### 惊群
```
惊群效应（thundering herd）是指多进程(多线程)在同时阻塞等待同一个事件的时候(休眠状态)，如果等待的这个事件发生，那么他就会唤醒等待的所有进程(或者线程)，但是最终却只能有一个进程(线程)获得这个时间的"控制权"，对该事件进行处理，而其他进程(线程)获取"控制权"失败，只能重新进入休眠状态，这种现象和性能浪费就叫做惊群效应。

1、accept惊群
    (1) 场景
        主进程创建了socket、bind、listen之后，fork()出来多个进程，每个子进程都开始循环处理accept这个listen_fd。每个进程都阻塞在accept上，当一个新的连接到来时候，所有的进程都会被唤醒，但是其中只有一个进程会接受成功，其余皆失败，重新休眠
    (2) 存在
        历史上，Linux的accpet确实存在惊群问题
        在linux2.6版本以后，linux内核已经解决了accept()函数的“惊群”现象，大概的处理方式就是，当内核接收到一个客户连接后，只会唤醒等待队列上的第一个进程（线程）,所以如果服务器采用accept阻塞调用方式，在最新的linux系统中已经没有“惊群效应”了


2、epoll惊群
https://www.cnblogs.com/sduzh/p/6810469.html
https://www.zhihu.com/question/24169490/answers/updated

(1) epoll存在惊群
    当有一个新的连接到来时，唤醒的进程可能只有一个，也可能有部分，也可能是全部。
    但是只有一个进程是可以读写的，剩余的进程会获得一个EAGAIN信号，存在EAGAIN的进程不会继续读写，而错误退出

(2) 水平模式下才会发生
    epoll惊群只会在水平模式下才会发生
    在epoll的LT模式下，每次epoll_wait 将 readylist的event返回用户空间后，会将epi立刻再加入到ready_list里, 而不像ET模式下是清空当前的ready_list, 由于ready_list不为空了，会再次调用 waitqueue_active来唤醒epoll本身的等待队列，即其他线程的 epoll_wait会返回, 造成惊群的现象

(3) 解决办法
    https://www.jianshu.com/p/21c3e5b99f4a
    void ngx_process_events_and_timers(ngx_cycle_t *cycle)
    {
        ...
        //这里面会对监听socket处理
        //1、获得锁则加入wait集合
        //2、没有获得则去除
        if (ngx_trylock_accept_mutex(cycle) == NGX_ERROR) {
            return;
        }

        ...

        //设置网络读写事件延迟处理标志，即在释放锁后处理
        if (ngx_accept_mutex_held) {
            flags |= NGX_POST_EVENTS;
        } 

        ...

        //这里面epollwait等待网络事件
        //网络连接事件，放入ngx_posted_accept_events队列
        //网络读写事件，放入ngx_posted_events队列
        (void) ngx_process_events(cycle, timer, flags);

        ...

        //先处理网络连接事件，只有获取到锁，这里才会有连接事件
        ngx_event_process_posted(cycle, &ngx_posted_accept_events);

        //释放锁，让其他进程也能够拿到
        if (ngx_accept_mutex_held) {
            ngx_shmtx_unlock(&ngx_accept_mutex);
        }

        //处理网络读写事件
        ngx_event_process_posted(cycle, &ngx_posted_events);
    }

```
