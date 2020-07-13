

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

### epoll的相关概念
```
基于事件驱动的IO方式

https://blog.csdn.net/u010155023/article/details/53507788 [select #(select实例)]
https://blog.csdn.net/davidsguo008/article/details/73556811 [epoll #(epoll实例)]
https://yq.aliyun.com/articles/683282 [关于epoll的IO模型是同步异步的一次纠结过程]
https://segmentfault.com/a/1190000003063859
https://www.cnblogs.com/lojunren/p/3856290.html❗
http://blog.chinaunix.net/uid-28541347-id-4273856.html❗
https://blog.csdn.net/dog250/article/details/80837278❗

1、底层实现
    (1) 两个数据结构
        epoll底层实现最重要的两个数据结构 epitem和eventpoll
        epitem是和每个用户态监控IO的fd对应的
        eventpoll是用户态创建的管理所有被监控fd的结构
        1) eventpoll
            调用epoll_create方法时，内核会创建一个eventpoll对象，eventpoll对象也是文件系统中的一员，和socket一样，它也会有等待队列
            epoll对象是进程与socket之间的中介，包含了监视队列和就绪列表

            struct eventpoll {
                spinlock_t lock;
                struct mutex mtx;

                wait_queue_head_t wq;           // sys_epoll_wait()使用的等待队列
                wait_queue_head_t poll_wait;    // file->poll()使用的等待队列

                struct list_head rdllist;       // 所有'准备就绪'的文件描述符列表
                struct rb_root rbr;             // 用于储存'已监控'fd的红黑树根节点
                
                struct epitem *ovflist;         // 当正在向用户空间传递事件，则就绪事件会临时放到该队列，否则直接放到rdllist
                ...
            };

            每一个epoll对象都有一个独立的eventpoll结构体，用于存放通过epoll_ctl方法向epoll对象中添加进来的事件
            
            这些'事件'都会挂载在红黑树中(重复添加的事件就可以通过红黑树而高效的识别出来，红黑树的插入时间效率是O(lgn)，其中n为元素个数)

        2) epitem
            监视队列和就绪列表的基本结构，每一个要监听的fd对应一个此结构体   
            
            所有添加到epoll中的事件都会与设备(网卡)驱动程序建立回调关系，也就是说当相应的事件发生时会调用这个回调方法
            
            此回调方法在内核中叫ep_poll_callback，它会将发生的事件添加到rdlist双链表中

            在epoll中，对于每一个事件，都会建立一个epitem结构体，如下所示
            struct epitem{
                struct rb_node rbn;         //红黑树节点
                struct list_head rdllink;   //双向链表节点
                struct epitem *next; 
                struct epoll_filefd ffd;    //此条目引用的文件描述符信息
                struct eventpoll *ep;       //指向其所属的eventpoll对象
                struct epoll_event event;   //期待发生的事件类型
            } 
            当调用epoll_wait检查是否有事件发生时(所以epoll不是异步)，只需要检查eventpoll对象中的rdlist双链表中是否有epitem元素即可
            如果rdlist不为空，则把发生的事件复制到用户态，同时将事件数量返回给用户

    (3) 等待队列
        创建socket时，会创建socket文件对象，里面包含一个非常重要的结构(等待队列)，它指向所有需要等待该socket事件的进程
    (4) 监视队列(rbr 红黑树)
        用epoll_ctl添加或删除所要监听的socket，内核会将eventpoll添加到目标socket的等待队列中
    (5) 就绪列表(rdlist 双向链表)
        当socket收到数据后(由不可读变成可读或由不可写变成可写)，中断函数将添加socket的引用到就绪列表
        另一方面唤醒eventpoll等待队列中的进程(可能会发生惊群)

2、工作模式
    https://www.jianshu.com/p/d3442ff24ba6
    http://kimi.it/515.html

    ET模式和FT模式
    使用ET的例子 -> nginx
    使用LT的例子 -> redis

    (1) ET(边缘触发)
        socket的接收缓冲区状态变化时触发读事件，即空的接收缓冲区刚接收到数据时触发读事件
        socket的发送缓冲区状态变化时触发写事件，即满的缓冲区刚空出空间时触发读事件
        边沿触发仅触发一次
        1) 特点
            > 句柄在发生读写事件时只会通知用户一次
            > ET模式主要关注fd从不可用到可用或者可用到不可用的情况
            > ET只支持非阻塞模式
        2) 一般用法
            > ET模式下读写操作要时用while循环，直到读/写够足够多的数据，或者读/写到返回EAGAIN
            > 写大块数据时，一次write操作不足以写完全部数据，或者在读大块数据时，应用层缓冲区数据太小，一次read操作不足以读完全部数据，应用层要么一直调用while循环一直IO到EGAIN，或者自己调用epoll_ctl手动触发ET响应

    (2) LT(水平触发，默认)
        socket接收缓冲区不为空，有数据可读，读事件一直触发
        socket发送缓冲区不满，可以继续写入数据，写事件一直触发
        水平触发会一直触发
        1) 特点
            > 只要句柄一直处于可用状态，就会一直通知用户
            > LT模式下，句柄读缓冲区被读空后，句柄会从可用转变未不可以用，这个时候不会通知用户。写缓冲区只要还没写满，就会一直通知用户
            > LT模式支持阻塞和非阻塞两种方式
            > LT下，应用层的业务逻辑比较简单，更不容易遗漏事件，更不容易出错。通常，在将数据写完后，我们会关闭句柄的写事件

    (3) ET vs LT
        LT模式下，只要一个句柄上的事件一次没有处理完，会在以后调用epoll_wait时次次返回这个句柄，而ET模式仅在第一次返回

        > 当一个socket句柄上有事件时，内核会把该句柄插入准备就绪list链表
        > 调用epoll_wait，会把准备就绪的socket拷贝到用户态内存，然后清空准备就绪list链表
        > 然后epoll_wait检查这些socket的模式
            > 如果是LT模式，并且这些socket上确实有未处理的事件时，又把该句柄放回到刚刚清空的准备就绪链表了。所以，LT的句柄，只要它上面还有事件，epoll_wait每次都会返回
            > 如果是ET模式的句柄，除非有新中断到，即使socket上的事件没有处理完，也是不会次次从epoll_wait返回的
    
    (4) 使用建议
        对于监听的sockfd，最好使用水平触发模式，边缘触发模式会导致高并发情况下，有的客户端会连接不上。如果非要使用边缘触发，可以用while来循环accept()
        
        对于读写的 connfd，水平触发模式下，阻塞和非阻塞效果都一样，建议设置非阻塞
        
        对于读写的 connfd，边缘触发模式下，必须使用非阻塞 IO，并要求一次性地完整读写全部数据
```

### epoll相关函数
```
https://blog.csdn.net/ljx0305/article/details/4065058

(1) int epoll_create(int size);
    创建一个epoll的句柄
    // size     用来告诉内核这个监听的数目一共有多大，这个参数不同于select()中的第一个参数，给出最大监听的fd+1的值(不过目前size并不起作用了)
    // 返回     创建的epoll句柄
    // 注意
        当创建好epoll句柄后，它就是会占用一个fd值，在linux下如果查看/proc/进程id/fd/，是能够看到这个fd的，所以在使用完epoll后，必须调用close()关闭，否则可能导致fd被耗尽

(2) int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
    epoll的事件注册函数，它不同与select()是在监听事件时告诉内核要监听什么类型的事件，而是在这里先注册要监听的事件类型
    // epfd     epoll_create()的返回值
    // op       表示动作，用三个宏来表示
        EPOLL_CTL_ADD   注册新的fd到epfd中
        EPOLL_CTL_MOD   修改已经注册的fd的监听事件
        EPOLL_CTL_DEL   从epfd中删除一个fd
    // fd       需要监听的fd
    // event    告知内核需要监听的事件
        // event结构
            struct epoll_event {
                __uint32_t events;      // epoll events
                epoll_data_t data;      // user data variable
            };
            
            typedef union epoll_data {  //用户数据载体
                void *ptr;              // 传进我们想要的内容，wait出来的时候，用我们自己的函数进行处理
                int fd;                 // 上面的fd
                __uint32_t u32;
                __uint64_t u64;
            } epoll_data_t;
            
        // epoll events操作
            EPOLLIN         表示对应的文件描述符可以读(包括对端SOCKET正常关闭)
            EPOLLOUT        表示对应的文件描述符可以写
            EPOLLPRI        表示对应的文件描述符有紧急的数据可读(表示有带外数据到来)
            EPOLLERR        表示对应的文件描述符发生错误
            EPOLLHUP        表示对应的文件描述符被挂断
            EPOLLET         将此描述符设为边缘触发(ET)模式
            EPOLLONESHOT    防止不同进程处理同一个socket事件
                A线程读完某socket上数据后开始处理这些数据，此时该socket上又有新数据可读，B线程被唤醒读新的数据，造成2个线程同时操作一个socket的局面 ，EPOLLONESHOT保证一个socket连接在任一时刻只被一个线程处理
            EPOLLEXCLUSIVE  linux内核4.5新加的，保证一个事件发生时候只有部分线程会被唤醒，缓解惊群问题

(3) int epoll_wait(int epfd, struct epoll_event *events, int maxevents, int timeout);
    等待epoll句柄上的I/O事件，最多返回maxevents个事件
    // epfd         同上
    // events       返回就绪事件的集合
    // maxevents    告之内核这个events有多大，这个 maxevents的值不能大于创建epoll_create()时的size
    // timeout      超时时间(毫秒)，0会立即返回，-1将不确定(永久阻塞)
    // 返回         大于0返回就绪事件数目，等于0表示超时

```

### epoll demo
```
这块在IO多线程模型中是reactor设计模型
没有考虑多线程的应用

#define MAX_EVENTS 10
struct epoll_event ev, events[MAX_EVENTS];
int listen_sock, conn_sock, nfds, epollfd;

/* Set up listening socket, 'listen_sock' (socket(),
  bind(), listen()) */

//创建一个描述符
epollfd = epoll_create(10);
if(epollfd == -1) {
    perror("epoll_create");
    exit(EXIT_FAILURE);
}

//添加监听描述符事件
ev.events = EPOLLIN;
ev.data.fd = listen_sock;
if(epoll_ctl(epollfd, EPOLL_CTL_ADD, listen_sock, &ev) == -1) {
    perror("epoll_ctl: listen_sock");
    exit(EXIT_FAILURE);
}

for(;;) {
    // 返回已经准备好的描述符事件数目
    nfds = epoll_wait(epollfd, events, MAX_EVENTS, -1);
    if (nfds == -1) {
        perror("epoll_pwait");
        exit(EXIT_FAILURE);
    }

    // 遍历已经准备好的io事件
    for (n = 0; n < nfds; ++n) {
        if (events[n].data.fd == listen_sock) {
            // 主监听socket有新连接
            conn_sock = accept(listen_sock, (struct sockaddr *) &local, &addrlen);
            if (conn_sock == -1) {
                perror("accept");
                exit(EXIT_FAILURE);
            }
            // 设置非阻塞IO
            setnonblocking(conn_sock);
            ev.events = EPOLLIN | EPOLLET;
            ev.data.fd = conn_sock;
            // 添加新连接到epoll中
            if (epoll_ctl(epollfd, EPOLL_CTL_ADD, conn_sock, &ev) == -1) {
                perror("epoll_ctl: conn_sock");
                exit(EXIT_FAILURE);
            }
        } else {
            //已建立连接的可读写句柄
            do_use_fd(events[n].data.fd);
        }
    }
}

```

### 惊群
```
惊群效应(thundering herd)是指多进/线程在同时阻塞等待同一个事件的时候(休眠状态)，如果等待的这个事件发生，那么他就会唤醒等待的所有进/线程，但是最终却只能有一个进/线程获得这个事件的"控制权"，对该事件进行处理，而其他进程(线程)获取"控制权"失败，只能重新进入休眠状态，这种现象和性能浪费就叫做惊群效应

多个进程因为一个连接请求而被同时唤醒，称为惊群效应，在高并发情况下，大部分进程会无效地被唤醒然后因为抢占不到连接请求又重新进入睡眠，是会造成系统极大的性能损耗

1、accept惊群
    多个进程同时监听一个sockfd，每个进程都阻塞在accept，当一个新的连接到来时候，所有的进程都会被唤醒，但是其中只有一个进程会接受成功，其余皆失败，重新休眠

    在linux2.6版本以后，linux内核已经解决了accept()函数的"惊群"现象
    大概的处理方式是，当内核接收到一个客户连接后，只会唤醒等待队列上的第一个进程(线程),所以如果服务器采用accept阻塞调用方式，在最新的linux系统中已经没有'惊群效应'

2、epoll惊群
https://www.cnblogs.com/sduzh/p/6810469.html
https://www.zhihu.com/question/24169490/answers/updated
https://zhuanlan.zhihu.com/p/87843750

(1) epoll存在惊群
    调用epoll_create创建一个epfd，然后多进/线程同时epoll_wait这个epfd事件
    当新的连接到来时，这些子进程全部(也可能是部分)被唤醒并处理事件
    但是只有一个子进程是可以处理连接，剩余的进程会获得一个EAGAIN信号，存在EAGAIN的进程不会继续读写，而错误退出

(2) LT模式存在惊群
    epoll惊群只会在LT模式下发生
    在epoll的LT模式下，每次epoll_wait将readylist的event返回用户空间后，会将epi立刻再加入到ready_list里，而不像ET模式下是清空当前的ready_list。
    由于ready_list不为空了，当连接再次到来时，会再次调用waitqueue_active来唤醒epoll本身的等待队列，即其他线程的epoll_wait会返回, 造成惊群的现象

(3) 解决办法
    https://www.jianshu.com/p/21c3e5b99f4a
    https://simpleyyt.com/2017/06/25/how-ngnix-solve-thundering-herd/ [Ngnix 是如何解决 epoll 惊群的]

    1) accept_mutex锁
        每个worker都会先去抢自旋锁，只有抢占成功了，才把socket加入到epoll中，accept请求，然后释放锁。accept_mutex锁也有负载均衡的作用
        // 伪代码如下
        lock()
        epoll_wait(...);
        accept(...);
        unlock()

        效率低下，特别是在长连接时，一个进程长时间占用accept_mutex锁，使得其它进程得不到 accept 的机会
        accept_mutex的分配不均可能会导致进程饥饿

    2) EPOLLEXCLUSIVE(Nginx > 1.11.3)
        在epoll层面上，EPOLLEXCLUSIVE标识会保证一个事件发生时候只有部分(one or more)线程会被唤醒，一定程度上缓解了惊群效应
        不过任一时候只能有一个工作线程调用accept，限制了真正并行的吞吐量

    3) SO_REUSEPORT(Nginx > 1.9.1)
        SO_REUSEPORT是惊群最推荐的解决方法，每个worker都有自己的socket，这些socket都bind同一个端口。当新请求到来时，内核根据四元组信息进行负载均衡，非常高效
```
