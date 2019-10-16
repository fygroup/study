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
4、hostent地址查询
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

### SO_REUSEADDR和SO_REUSEPORT
```
https://zhuanlan.zhihu.com/p/35367402

(1) SO_REUSEADDR
    1) 使用场景
        server端在调用bind函数时
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
