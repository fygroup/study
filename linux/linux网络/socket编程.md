# socket
```
#include <sys/socket.h>
int socket(int domain, int type, int protocol)
//domain(域)
AF_INET     IPv4
AF_INET6    IPv6
AF_UNIX     UNIX域
//type
SOCK_DGRAM  UDP（固定长度、无连接、不可靠报文传递）
SOCK_RAM    ip协议数据报接口，用于直接访问网络层，绕过传输层（tcp、udp），需要超级用户特权
SOCK_STREAM TCP(有序、可靠、双向、面向连接字节流)
//protocol
0           表示为给定的域和套接字选择默认协议
```

# 字节序
```
  小端/大端         大端
处理器字节序 ---> 网络字节序
#include <arpa/inet.h>
uint32_t htonl(uint32_t hostint32)  返回 网络字节序32位整数
uint16_t htons(uint16_t hostint16)  返回 网络字节序16位整数
uint32_t ntonl(uint32_t netint32)   返回 主机字节序32位整数
uint16_t ntons(uint16_t netint16)   返回 主机字节序16位整数
```

# 地址
1、通用socket地址
```
不同的地址格式必须转换为此格式
#include <sys/socket.h>
struct sockaddr {
    sa_family_t  sa_family;     //地址族,unsigned short,AF_xxx
    char         sa_data[14];   //14字节，包含套接字中的目标地址和端口信息     
}
```
2、专用socket地址
```
//IPv4
#include<netinet/in.h>
typedef uint16_t in_port_t;
typedef uint32_t in_addr_t;
struct sockaddr_in {    
    sa_family_t    sin_family;    //地址族
    in_port_t      sin_port;      //16位端口号
    struct in_addr sin_addr;      //32位IP地址
    unsigned  char  sin_zero[8];         /* Same size as struct sockaddr */
}
struct in_addr {
    in_addr_t      s_addr         //32位IPv4地址
}

//IPv6
...

//unix域套接字地址
#include <sys/un.h>
struct sockaddr_un {
    sa_family    sun_family;    /* AF_UNIX */
    char         sun_path[108];    /* pathname */
};


```
3、addr转换
```
//tcp套接字转换
struct sockaddr_in my_addr;
my_addr.sin_family      = AF_INET;
my_addr.sin_port        = htons(80);                 //uint16转换成网络字节序
my_addr.sin_addr.s_addr = inet_addr("192.168.2.201") //inet_addr将字符串转换为网络addr字节， inet_ntoa相反
bzero(&(my_addr.sin_zero), 8);                       //sin_zero置0
struct sockaddr* myaddr = (struct sockaddr*)&my_addr //转换成sockaddr

//unix域套接字转换
struct sockaddr_un un;
memset(&un, 0, sizeof(un));
un.sun_family = AF_UNIX;
strcpy(un.sun_path, "foo.socket");
if((fd = socket(AF_UNIX, SOCK_STREAM, 0)) < 0)
        err_sys("socket failed");
if(bind(fd, (struct sockaddr*)&servaddr, sizeof(servaddr)) < 0)
        ERR_EXIT("bind");


```
4、hostent
```

```



### accept
```
https://www.cnblogs.com/wangcq/p/3520400.html

三次握手发生在这一步

TCP服务器端依次调用socket()、bind()、listen()之后，就会监听指定的socket地址了。TCP客户端依次调用socket()、connect()之后就想TCP服务器发送了一个连接请求。TCP服务器监听到这个请求之后，就会调用accept()函数取接收请求，这样连接就建立好了。之后就可以开始网络I/O操作了，即类同于普通文件的读写I/O操作。

int accept(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
参数1：服务器的socket描述字
参数2：客户端的协议地址
参数3：第三个参数为协议地址的长度
返回值：由内核自动生成的一个全新的描述字，代表与返回客户的TCP连接。

注意：内核为每个由服务器进程接受的客户连接创建了一个已连接socket描述字，当服务器完成了对某个客户的服务，相应的已连接socket描述字就被关闭。
```

### read/write的返回
```
(1) 对于阻塞socket
    能read时，读缓冲区没有数据，或者write时，写缓冲区满了。这是就发生阻塞，如果返回-1代表网络出错了
(2) 对于非阻塞socket
    不能read或write时，就会返回-1，同时errno设置为EAGAIN（再试一次）。
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
    1) 句柄从不可读变成可读，或者句柄写缓冲区有新的数据进来且超过SO_RCVLOWAT。
    2) 产生事件的情况
        > socket有一个未清除的错误。如非阻塞的connect连接错误会使socket变成可读写状态。
        > 非阻塞accept有新的连接进来。
        > socket写对端关闭，read返回0。
        > socket读缓冲区有新的数据进来且超过SO_RCVLOWAT

(2) 写事件
    1) 句柄从不可写变成可写，或者句柄写缓冲区有新的数据进来而且缓冲区水位高于SO_SNDLOWAT。
    2) 产生事件的情况
        > socket有一个未清除的错误。例如非阻塞connect连接出错会导致socket变成可读可写状态。
        > 非阻塞connect连接成功后端口状态会变成可写。
        > socket读对端关闭，socket变成可写状态，产生SIGPIPE信号。
        > socket写缓冲区有新的数据进来且超过SO_SNDLOWAT
    
在epoll中，读事件对应EPOLLIN，写事件对应EPOLLOUT。
```

### epoll
1、原理
```
基于事件驱动的I/O方式

(1) 重要概念
    1) socket对象
        由文件系统管理的，包含了发送缓冲区、接收缓冲区、等待队列
    2) 等待队列
        创建socket时，会创建socket文件对象，里面包含一个非常重要的结构（等待队列），它指向所有需要等待该socket事件的进程。
    2) epoll对象
        调用epoll_create方法时，内核会创建一个eventpoll对象，eventpoll对象也是文件系统中的一员，和socket一样，它也会有等待队列。epoll对象时进程与socket之间的中介，包含了监视队列和就绪列表
    4) epitem对象
        监视队列和就绪列表的基本结构，每一个要监听的fd对应一个此结构体    
    5) 监视队列(rbr 红黑树)
        用epoll_ctl添加或删除所要监听的socket，内核会将eventpoll添加到目标socket的等待队列中。
    6) 就绪列表(rdlist 双向链表)
        当socket收到数据后（由不可读变成可或由不可写变成可写，见上述），中断函数将添加socket的引用到就绪列表，另一方面唤醒eventpoll等待队列中的进程
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

(2) 流程图

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
    EPOLLONESHOT：只监听一次事件，当监听完这次事件之后，如果还需要继续监听这个socket的话，需要再次把这个socket加入到EPOLL队列里

(3) int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
    等待事件的产生，类似于select()调用。参数events用来从内核得到事件的集合，maxevents告之内核这个events有多大，这个 maxevents的值不能大于创建epoll_create()时的size，参数timeout是超时时间（毫秒，0会立即返回，-1将不确定，也有说法说是永久阻塞）。该函数返回需要处理的事件数目，如返回0表示已超时。

```
3、用法
```
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