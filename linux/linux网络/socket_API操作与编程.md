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