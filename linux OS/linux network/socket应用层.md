

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
