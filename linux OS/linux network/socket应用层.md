

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
基于事件驱动的I/O方式

https://blog.csdn.net/u010155023/article/details/53507788 [select #(select实例)]
https://blog.csdn.net/davidsguo008/article/details/73556811 [epoll #(epoll实例)]
https://yq.aliyun.com/articles/683282 [关于epoll的IO模型是同步异步的一次纠结过程]
https://segmentfault.com/a/1190000003063859
https://www.cnblogs.com/lojunren/p/3856290.html❗
http://blog.chinaunix.net/uid-28541347-id-4273856.html❗
https://blog.csdn.net/dog250/article/details/80837278❗

(1) 重要概念
    1) socket对象
        由文件系统管理的，包含了发送缓冲区、接收缓冲区、等待队列
    2) epoll对象
        调用epoll_create方法时，内核会创建一个eventpoll对象，eventpoll对象也是文件系统中的一员，和socket一样，它也会有等待队列
        epoll对象是进程与socket之间的中介，包含了监视队列和就绪列表

        struct eventpoll{
            ....
            // 红黑树的根节点，这颗树中存储着所有添加到epoll中的需要'监控的事件'
            struct rb_root  rbr;
            // 双链表中则存放着将要通过epoll_wait返回给用户的'满足条件的事件'
            struct list_head rdlist;
            ....
        };

        每一个epoll对象都有一个独立的eventpoll结构体，用于存放通过epoll_ctl方法向epoll对象中添加进来的事件
        
        这些'事件'都会挂载在红黑树中(重复添加的事件就可以通过红黑树而高效的识别出来，红黑树的插入时间效率是O(lgn)，其中n为元素个数)

    3) 等待队列
        创建socket时，会创建socket文件对象，里面包含一个非常重要的结构(等待队列)，它指向所有需要等待该socket事件的进程
    4) 监视队列(rbr 红黑树)
        用epoll_ctl添加或删除所要监听的socket，内核会将eventpoll添加到目标socket的等待队列中
    5) epitem对象
        监视队列和就绪列表的基本结构，每一个要监听的fd对应一个此结构体   
        
        所有添加到epoll中的事件都会与设备(网卡)驱动程序建立回调关系，也就是说当相应的事件发生时会调用这个回调方法
        
        此回调方法在内核中叫ep_poll_callback，它会将发生的事件添加到rdlist双链表中

        在epoll中，对于每一个事件，都会建立一个epitem结构体，如下所示
        struct epitem{
            struct rb_node rbn;         //红黑树节点
            struct list_head rdllink;   //双向链表节点
            struct epoll_filefd ffd;    //事件句柄信息
            struct eventpoll *ep;       //指向其所属的eventpoll对象
            struct epoll_event event;   //期待发生的事件类型
        } 
        当调用epoll_wait检查是否有事件发生时(所以epoll不是异步)，只需要检查eventpoll对象中的rdlist双链表中是否有epitem元素即可
        如果rdlist不为空，则把发生的事件复制到用户态，同时将事件数量返回给用户

    6) 就绪列表(rdlist 双向链表)
        当socket收到数据后(由不可读变成可读或由不可写变成可写)，中断函数将添加socket的引用到就绪列表
        另一方面唤醒eventpoll等待队列中的进程(可能会发生惊群)

    7) ET和FT
        https://www.jianshu.com/p/d3442ff24ba6
        http://kimi.it/515.html
        
        使用ET的例子 -> nginx
        使用LT的例子 -> redis
        
        > ET
            socket的接收缓冲区状态变化时触发读事件，即空的接收缓冲区刚接收到数据时触发读事件
            socket的发送缓冲区状态变化时触发写事件，即满的缓冲区刚空出空间时触发读事件
            边沿触发仅触发一次
            > 特点
                句柄在发生读写事件时只会通知用户一次
                ET模式主要关注fd从不可用到可用或者可用到不可用的情况
                ET只支持非阻塞模式
                ET模式下读写操作要时用while循环，直到读/写够足够多的数据，或者读/写到返回EAGAIN
                写大块数据时，一次write操作不足以写完全部数据，或者在读大块数据时，应用层缓冲区数据太小，一次read操作不足以读完全部数据，应用层要么一直调用while循环一直IO到EGAIN，或者自己调用epoll_ctl手动触发ET响应

        > LT(epoll默认模式)
            socket接收缓冲区不为空 有数据可读 读事件一直触发
            socket发送缓冲区不满 可以继续写入数据 写事件一直触发
            水平触发会一直触发
            > 特点
                只要句柄一直处于可用状态，就会一直通知用户
                LT模式下，句柄读缓冲区被读空后，句柄会从可用转变未不可以用，这个时候不会通知用户。写缓冲区只要还没写满，就会一直通知用户
                LT模式支持阻塞和非阻塞两种方式
                LT下，应用层的业务逻辑比较简单，更不容易遗漏事件，更不容易出错。通常，在将数据写完后，我们会关闭句柄的写事件
        > ET vs LT
            LT模式下，只要一个句柄上的事件一次没有处理完，会在以后调用epoll_wait时次次返回这个句柄，而ET模式仅在第一次返回

            当一个socket句柄上有事件时，内核会把该句柄插入准备就绪list链表
            我们调用epoll_wait，会把准备就绪的socket拷贝到用户态内存，然后清空准备就绪list链表，然后epoll_wait检查这些socket的模式
            如果是LT模式，并且这些socket上确实有未处理的事件时，又把该句柄放回到刚刚清空的准备就绪链表了。所以，LT的句柄，只要它上面还有事件，epoll_wait每次都会返回
            如果是ET模式的句柄，除非有新中断到，即使socket上的事件没有处理完，也是不会次次从epoll_wait返回的
        > 使用建议
            对于监听的sockfd，最好使用水平触发模式，边缘触发模式会导致高并发情况下，有的客户端会连接不上。如果非要使用边缘触发，可以用 while 来循环 accept()。
            对于读写的 connfd，水平触发模式下，阻塞和非阻塞效果都一样，建议设置非阻塞。
            对于读写的 connfd，边缘触发模式下，必须使用非阻塞 IO，并要求一次性地完整读写全部数据。



```

### epoll相关函数
```
https://blog.csdn.net/ljx0305/article/details/4065058

(1) int epoll_create(int size);
    创建一个epoll的句柄
    // size 用来告诉内核这个监听的数目一共有多大，这个参数不同于select()中的第一个参数，给出最大监听的fd+1的值
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
    // event    监听的事件
        struct epoll_event {
            __uint32_t events;  // Epoll events
            epoll_data_t data;  // User data variable
        };

        typedef union epoll_data {
            void *ptr;
            int fd;
            __uint32_t u32;
            __uint64_t u64;
        } epoll_data_t;
        
        // events   表示对应的文件描述符的操作
            EPOLLIN         表示对应的文件描述符可以读(包括对端SOCKET正常关闭)
            EPOLLOUT        表示对应的文件描述符可以写
            EPOLLPRI        表示对应的文件描述符有紧急的数据可读(表示有带外数据到来)
            EPOLLERR        表示对应的文件描述符发生错误
            EPOLLHUP        表示对应的文件描述符被挂断
            EPOLLET         将此描述符设为边缘触发(ET)模式
            EPOLLONESHOT    多线程模式下下仅触发一个线程执行事件，且这次执行完，下次其他线程执行
                对于注册了EPOLLONESHOT事件的文件描述符，操作系统最多其上注册的一个可读，可写或异常事件，且只能触发一次
                注册了EPOLLONESHOT事件的socket一旦被某个线程处理完毕，该线程就应该立即重置这个socket上的EPOLLONESHOT事件，以确保这个socket下一次可读时，其EPOLLIN事件能被触发，进而让其他工作线程有机会继续处理这个socket

(3) int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
    等待事件的产生，类似于select()调用
    // epfd         同上
    // events       返回就绪事件的集合
    // maxevents    告之内核这个events有多大，这个 maxevents的值不能大于创建epoll_create()时的size
    // timeout      超时时间(毫秒)，0会立即返回，-1将不确定(永久阻塞)
    // 返回         返回就绪事件数目，如返回0表示已超时

```

### epoll示例
```
#define IPADDRESS   "127.0.0.1"
#define PORT        8787
#define MAXSIZE     1024
#define LISTENQ     5
#define FDSIZE      1000
#define EPOLLEVENTS 100

int main(){
    listenfd = socket_bind(IPADDRESS,PORT);

    struct epoll_event events[EPOLLEVENTS];

    //创建一个描述符
    epollfd = epoll_create(FDSIZE);

    //添加监听描述符事件
    add_event(epollfd, listenfd, EPOLLIN);

    //循环等待
    for ( ; ; ){
        //该函数返回已经准备好的描述符事件数目
        ret = epoll_wait(epollfd,events,EPOLLEVENTS,-1);
        //处理接收到的连接
        handle_events(epollfd,events,ret,listenfd,buf);
    }
}

//事件处理函数
static void handle_events(int epollfd,struct epoll_event *events,int num,int listenfd,char *buf) {
    int i;
    int fd;
    //进行遍历，这里只要遍历已经准备好的io事件
    // num是事件就绪的数目，小于等于当初epoll_create时的FDSIZE
    for (i = 0;i < num;i++) {
        fd = events[i].data.fd;
        //根据描述符的类型和事件类型进行处理
        if ((fd == listenfd) && (events[i].events & EPOLLIN))
            handle_accpet(epollfd,listenfd);
        else if (events[i].events & EPOLLIN)
            do_read(epollfd,fd,buf);
         else if (events[i].events & EPOLLOUT)
            do_write(epollfd,fd,buf);
     }
}

//添加事件
static void add_event(int epollfd, int fd, int state){
    struct epoll_event ev;
    ev.events = state;
    ev.data.fd = fd;
    epoll_ctl(epollfd, EPOLL_CTL_ADD, fd, &ev);
}

//处理接收到的新连接
static void handle_accpet(int epollfd,int listenfd){
    int clifd;     
    struct sockaddr_in cliaddr;     
    socklen_t  cliaddrlen;     
    clifd = accept(listenfd, (struct sockaddr*)&cliaddr, &cliaddrlen);     
    if (clifd == -1) perror("accpet error:");     
    else {         
        printf("accept a new client: %s:%d\n",inet_ntoa(cliaddr.sin_addr),cliaddr.sin_port);                       
        //添加一个客户描述符和事件
        add_event(epollfd,clifd,EPOLLIN);     
    } 
}

//读处理
static void do_read(int epollfd, int fd, char *buf){
    int nread;
    nread = read(fd,buf,MAXSIZE);
    if (nread == -1) {         
        perror("read error:");         
        close(fd);                          // 记住close fd        
        delete_event(epollfd,fd,EPOLLIN);   // 删除监听 
    } else if (nread == 0) {         
        fprintf(stderr,"client close.\n");
        close(fd);                          //记住close fd       
        delete_event(epollfd,fd,EPOLLIN);   //删除监听 
    } else {         
        printf("read message is : %s",buf);        
        // 修改描述符对应的事件，由读改为写         
        modify_event(epollfd,fd,EPOLLOUT);     
    } 
}

//写处理
static void do_write(int epollfd,int fd,char *buf) {     
    int nwrite;     
    nwrite = write(fd,buf,strlen(buf));     
    if (nwrite == -1){         
        perror("write error:");        
        close(fd);                          //记住close fd       
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
        for(i=0; i<nfds; ++i)
        {
            if(events[i].data.fd == listenfd) {                                 //有新的连接
                connfd = accept(listenfd,(sockaddr *)&clientaddr, &clilen);     //accept这个连接
                ev.data.fd = connfd;
                ev.events = EPOLLIN|EPOLLET;
                epoll_ctl(epfd,EPOLL_CTL_ADD,connfd,&ev);                       //将新的fd添加到epoll的监听队列中
            } else if(events[i].events & EPOLLIN) {                             //接收到数据，读socket
                n = read(sockfd, line, MAXLINE);
                ev.data.ptr = md;                                               //md为自定义类型，添加数据
                ev.events = EPOLLOUT|EPOLLET;
                epoll_ctl(epfd, EPOLL_CTL_MOD, sockfd, &ev);                    //修改标识符，等待下一个循环时发送数据(异步处理的精髓)
            } else if(events[i].events & EPOLLOUT) {                            //有数据待发送，写socket
                struct myepoll_data* md = (myepoll_data*)events[i].data.ptr;    //取数据
                sockfd = md->fd;
                send(sockfd, md->ptr, strlen((char*)md->ptr), 0);               //发送数据
                ev.data.fd = sockfd;
                ev.events = EPOLLIN|EPOLLET;
                epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev);                       //修改标识符，等待下一个循环时接收数据
            } else {
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
        在linux2.6版本以后，linux内核已经解决了accept()函数的“惊群”现象，大概的处理方式就是，当内核接收到一个客户连接后，只会唤醒等待队列上的第一个进程（线程）,所以如果服务器采用accept阻塞调用方式，在最新的linux系统中已经没有'惊群效应'


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
