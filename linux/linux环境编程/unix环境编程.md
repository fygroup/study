### linux常用头文件
```
POSIX标准定义的头文件
<dirent.h>        目录项
<fcntl.h>         文件控制
<fnmatch.h>    文件名匹配类型
<glob.h>    路径名模式匹配类型
<grp.h>        组文件
<netdb.h>    网络数据库操作
<pwd.h>        口令文件
<regex.h>    正则表达式
<tar.h>        TAR归档值
<termios.h>    终端I/O
<unistd.h>    符号常量
<utime.h>    文件时间
<wordexp.h>    字符扩展类型
-------------------------
<arpa/inet.h>    INTERNET定义
<net/if.h>    套接字本地接口
<netinet/in.h>    INTERNET地址族
<netinet/tcp.h>    传输控制协议定义
------------------------- 
<sys/mman.h>    内存管理声明
<sys/select.h>    Select函数
<sys/socket.h>    套接字借口
<sys/stat.h>    文件状态
<sys/times.h>    进程时间
<sys/types.h>    基本系统数据类型
<sys/un.h>    UNIX域套接字定义
<sys/utsname.h>    系统名
<sys/wait.h>    进程控制
------------------------------
POSIX定义的XSI扩展头文件
<cpio.h>    cpio归档值 
<dlfcn.h>    动态链接
<fmtmsg.h>    消息显示结构
<ftw.h>        文件树漫游
<iconv.h>    代码集转换使用程序
<langinfo.h>    语言信息常量
<libgen.h>    模式匹配函数定义
<monetary.h>    货币类型
<ndbm.h>    数据库操作
<nl_types.h>    消息类别
<poll.h>    轮询函数
<search.h>    搜索表
<strings.h>    字符串操作
<syslog.h>    系统出错日志记录
<ucontext.h>    用户上下文
<ulimit.h>    用户限制
<utmpx.h>    用户帐户数据库 
-----------------------------
<sys/ipc.h>    IPC(命名管道)
<sys/msg.h>    消息队列
<sys/resource.h>资源操作
<sys/sem.h>    信号量
<sys/shm.h>    共享存储
<sys/statvfs.h>    文件系统信息
<sys/time.h>    时间类型
<sys/timeb.h>    附加的日期和时间定义
<sys/uio.h>    矢量I/O操作
------------------------------
POSIX定义的可选头文件
<aio.h>        异步I/O
<mqueue.h>    消息队列
<pthread.h>    线程
<sched.h>    执行调度
<semaphore.h>    信号量
<spawn.h>     实时spawn接口
<stropts.h>    XSI STREAMS接口
<trace.h>     事件跟踪
```

### 异步与非阻塞
```
1、讨论同步、异步、阻塞、非阻塞时，必须先明确是在哪个层次进行讨论
2、讨论究竟是异步还是同步，一定要严格说明说的是哪一部分
3、从Linux接口的角度说，阻塞和非阻塞都是同步。libaio那些才是异步。真正异步接口一般是你提供一个缓冲区给接口，然后接口立即返回，在一段时间之后通过另一种机制（回调，消息，信号等）通知你完成，在通知完成之前缓冲区你不能碰，系统在读写。
4、在处理 IO 的时候，阻塞和非阻塞都是同步 IO，只有使用了特殊的 API 才是异步 IO。
5、对unix来讲：阻塞式I/O(默认)，非阻塞式I/O(nonblock)，I/O复用(select/poll/epoll)都属于同步I/O，因为它们在数据由内核空间复制回进程缓冲区时都是阻塞的(不能干别的事)。只有异步I/O模型(linux的libaio)是符合异步I/O操作的含义的，即在1数据准备完成、2由内核空间拷贝回缓冲区后 通知进程，在等待通知的这段时间里可以干别的事。
6、阻塞，非阻塞：进程/线程要访问的数据是否就绪，进程/线程是否需要等待；
   同步，异步：访问数据的方式，同步需要主动读写数据，在读写数据的过程中还是会阻塞；异步只需要I/O操作完成的通知，并不主动读写数据，由操作系统内核完成数据的读写。
7、《UNIX网络编程：卷一》对unix的io讲得明明白白。
8、 说白了，同步需要从内核空间拷贝到用户空间，异步是内核帮你把数据拷贝到用户空间，所以异步需要底层api的支持。而阻塞和非阻塞是指进程访问的数据是否准备就绪，没有就绪则等待！！！

9、异步有异步io和异步操作，异步io如第八步所说的，而异步操作就多了，多线程、协程。。。，所以要根据软件的涉及。

```

#### 异步的进化
```
1、远古时代（回调函数）
2、promise时代
	promise().then().then()
3、Generator生成器
	实现代码生成器，实现switch（）类型的协程，异步。但不是真正异步
co(function *(){
    let db, collection, result; 
    let person = {name: "yika"};
    try{
        db = yield mongoDb.open();
        collection = yield db.collection("users");
        result = yield collection.insert(person);
    }catch(e){
        console.error(e.message);
    }
    console.log(result);
});

4、async/await时代
	真正的协程，实现异步最优雅的方式，用同步的方式写异步！
async function insertData(person){
    let db, collection, result; 
    try{
        db = await mongoDb.open();   //切除该协程，
        collection = await db.collection("users");
        result = await collection.insert(person);
    }catch(e){
        console.error(e.message);
    }
    console.log(result);
} 

```

---
### 协程的原理
```
//对称与非对称
对称类似于生产者、消费者之间协程的切换，并不涉及栈空间的销毁
非对称类似于函数的调用

//有栈与无栈
有栈：比如ucontext中的，协程有自己的栈空间，协程的切换涉及寄存器的保存和栈内数据的恢复问题，所以性能一般
无栈：用this来索引对象的成员变量，上下文就是对象自己。访问上下文数据也就是成员变量的时候，我们无需显式的使用this+成员偏移量（或者变量名）来访问，而是直接访问变量名。
两种协程访问的上下文中的数据，生命周期都大于函数的返回：栈的生命周期晚于函数的返回，this对象的生命周期晚于函数的返回。后者更晚而且往往需要手工销毁。

//hook
协程的意义就是阻塞异步，所以一些io函数必须设计为，非阻塞异步

//switch语法糖的实现协程

async/await的出现，实现了基于stackless coroutine的完整coroutine。在特性上已经非常接近stackful coroutine了，不但可以嵌套使用也可以支持try catch。所以是不是可以认为async/await是一个更好的方案？

```

---
### c++异步编程
c++11的promise、async、future属于多线程异步，所以单线程异步只能用协程和异步callback
异步是协程的一种实现方式，协程是异步的封装方法
```
1、promise、async的多线程异步回调（异步工作流）
2、协程（更优雅）
```

---
### c++高性能网络库 
```
libevent、libev、boost::asio
```

---
### c++ asio
```
推荐boost::asio
c++20标准库网络部份将基于asio，c++ asio异步编程很重要！！！
```

---
### 有栈协程相关模块
```
云风的coroutine库
libgo
golang
boost::asio
```

---
### 无栈协程相关模块
```
c++20的coroutine(基于asio)
知乎朱元的库
Es6的async/wait模型
boost::asio
```

---
### hook（linux钩子）
由于是调用得动态链接库中函数，我们可以通过劫持该函数的方式引入额外处理。 例如通过劫持 malloc、free 来追踪内存使用情况等等
```
//my_hook.c
#define _GNU_SOURCE
#include <stdio.h>
#include <stdint.h>
#include <dlfcn.h>

#define unlikely(x) __builtin_expect(!!(x), 0)
#define TRY_LOAD_HOOK_FUNC(name) if (unlikely(!g_sys_##name)) {g_sys_##name = (sys_##name##_t)dlsym(RTLD_NEXT,#name);}

typedef void* (*sys_malloc_t)(size_t size);
static sys_malloc_t g_sys_malloc = NULL;
void* malloc(size_t size)
{
    TRY_LOAD_HOOK_FUNC(malloc);
    void *p = g_sys_malloc(size);
    printf("in malloc hook function ...\n");
    return p;
}

typedef void (*sys_free_t)(void *ptr);
static sys_free_t g_sys_free = NULL;
void free(void *ptr)
{
    TRY_LOAD_HOOK_FUNC(free);
    g_sys_free(ptr);
    printf("in free hook function ...\n");
}

gcc -fPIC -shared -o libmyhook.so my_hook.c -ldl
gcc -o main main.c ./main
LD_PRELOAD=./libmyhook.so ./main

//LD_PRELOAD
能够影响程序运行时候动态链接库的加载，可以通过设置其来优先加载某些库，进而覆盖掉某些函数

//内联优化
由于编译器存在内联优化，不会调用库中的目标函数，所以必须关闭目标函数优化
-fno-builtin-strcmp，关闭 strcmp 函数的优化 
gcc -o main main.c -fno-builtin-strcmp

```

### 多路复用与异步IO
1、多路复用(同步非阻塞IO)
select #(select实例)[https://blog.csdn.net/u010155023/article/details/53507788]
epoll #(epoll实例)[https://blog.csdn.net/davidsguo008/article/details/73556811]
```
当某一进程调用epoll_create方法时，Linux内核会创建一个eventpoll结构体，这个结构体中有两个成员与epoll的使用方式密切相关。eventpoll结构体如下所示：
struct eventpoll{
    ....
    /*红黑树的根节点，这颗树中存储着所有添加到epoll中的需要监控的事件*/
    struct rb_root  rbr;
    /*双链表中则存放着将要通过epoll_wait返回给用户的满足条件的事件*/
    struct list_head rdlist;
    ....
};

每一个epoll对象都有一个独立的eventpoll结构体，用于存放通过epoll_ctl方法向epoll对象中添加进来的事件。这些事件都会挂载在红黑树中,如此，重复添加的事件就可以通过红黑树而高效的识别出来(红黑树的插入时间效率是lgn，其中n为树的高度)。

而所有添加到epoll中的事件都会与设备(网卡)驱动程序建立回调关系，也就是说，当相应的事件发生时会调用这个回调方法。这个回调方法在内核中叫ep_poll_callback,它会将发生的事件添加到rdlist双链表中。
在epoll中，对于每一个事件，都会建立一个epitem结构体，如下所示：
struct epitem{
    struct rb_node  rbn;//红黑树节点
    struct list_head    rdllink;//双向链表节点
    struct epoll_filefd  ffd;  //事件句柄信息
    struct eventpoll *ep;    //指向其所属的eventpoll对象
    struct epoll_event event; //期待发生的事件类型
}

当调用epoll_wait检查是否有事件发生时，只需要检查eventpoll对象中的rdlist双链表中是否有epitem元素即可。如果rdlist不为空，则把发生的事件复制到用户态（所以epoll不是异步），同时将事件数量返回给用户。


```

---
### fuse 用户空间文件系统
[fuse架构](../picture/fuse架构.png)
```
1、用户态app调用glibc open接口，触发sys_open系统调用。
2、sys_open 调用fuse中inode节点定义的open方法。
3、inode中open生成一个request消息，并通过/dev/fuse发送request消息到用户态libfuse。
4、Libfuse调用fuse_application用户自定义的open的方法，并将返回值通过/dev/fuse通知给内核。
5、内核收到request消息的处理完成的唤醒，并将结果放回给VFS系统调用结果。
```

---
### 进程间通信（IPC）

[进程间通信](https://zhuanlan.zhihu.com/p/37872762)

(1)概念
```
1、管道：是第一个广泛使用的IPC形式，既可以在程序中使用，也可以在shell中使用。管道存在的问题在于他们只能在具有共同祖先（指父子进程之间）的进程间使用，不过该问题已经被有名管道（named pipe）即FIFO消息队列解决了。
2、信号量
3、消息队列：消息队列是在两个不相关进程间传递数据的一种简单、高效方式，她独立于发送进程、接受进程而存在。消息队列是数据结构，存放在内存，访问速度快。但管道是文件，存放在磁盘上，访问速度慢。管道是数据流式存取，消息队列是数据块式存取。（rpc也是信号量的一种）
4、共享内存（同一机器下最快）
```

---
#### pipe(无名管道)
父子进程互相传递信号
```
#include <unistd.h>

int fd[2]; //0:读  1:写
int ret = pipe(fd);
if (ret==-1)perror();
pid_t pt = fork();
if (pt>0){
    close(fd[0]);       //关闭读
    write(fd[1]...);    
}else{
    close(fd[1]);       //关闭写
    read(fd[0]...);
}
```

---
#### FIFO(有名管道)
不同进程互相传递信号
```
#include <sys/stat.h>
mkfifo("file",0755);
int fd = open("file",O_RDONLY); //O_WRONLY
close(fd);

linux保证了写管道的原子性，但是每次写不能大于pipe_buf

```

---
#### 信号量
POSIX信号量与System V信号量
```
用于共享内存的同步，注意区别于线程的互斥量！！！（下面有详解）
都是用于线程和进程同步的。
Posix信号量是基于内存的，即信号量值是放在共享内存中的，与文件系统中的路径名对应的名字来标识的。性能更优越
System v信号量测试基于内核的，它放在内核里面。
```

(1)POSIX信号量
```
//一个进程创建POSIX信号量
#include <semaphore>
#define FILE_MODE (S_IRUSR|S_IWUSR|S_IRGRP|S_IROTH)

int main(){
    sem_unlink("file");  //防止所需的信号量已存在
    sem_t* mutex;
    if (mutex = sem_open("file",O_CREAT|O_EXCL,FILE_MODE,1) == SEM_FAILED){
        error("mutex");
        exit(-1);
    }
    sem_close(mutex);   //关闭
}
//另外一个进程运用POSIX信号量
#include <semaphore.h>

int main(){
    sem_t* mutex;
    if ((mutex = sem_open("file",0)) == SEM_FAILED){ //打开信号量
        error;
    }
    sem_wait(mutex);        //加锁
    ...
    sem_post(mutex);        //释放锁
}
```

(2)System V信号量
```
#include <sys/sem.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

int sem_id;
int set_semvalue(){
    union semun sem_union;
    sem_union.val = 1;
    if (semctl(sem_id,0,SETVAL,sem_union) == -1) return(0);
    return(1);
}
void del_semvalue(){
    union semun sem_union;
    if (semctl(sem_id,0,IPC_RMID,sem_union) == -1) perror();
}
int semaphore_p(){
    struct sembuf sem_b;
    sem_b.sem_num = 0;
    sem_b.sem_op = -1; //P()
    sem_b.sem_flg = SEM_UNDO;
    if (semop(sem_id,&sem_b,1) == -1)return(0);
    return(1);
}
int semaphore_v(){
    struct sembuf sem_b;
    sem_b.sem_num = 0;
    sem_b.sem_op = 1; //V()
    sem_b.sem_flg = SEM_UNDO;
    if (semop(sem_id,&sem_b,1) == -1)return(0);
    return(1);
}

int main(){
    key_t key = ftok("file",3);
    sem_id = semget(key,1,0666|IPC_CREAT); //创建信号量
    if (!set_semvalue()) perror();  //初始化信号量

    if (!semaphore_p()) perror;  //进入临界区
    ...
    if (!semaphore_v()) perror; //离开临界区

    del_semvalue();
}
```

---
#### 共享内存
```
分两种
System V的shmget()得到一个共享内存对象的id，用shmat()映射到进程自己的内存地址
POSIX的shm_open()打开一个文件，用mmap映射到自己的内存地址
```
<img src="../picture/7.png" alt="shm_open+mmap" height=300 width=500/>
注意：以上两种方式要用信号量同步

(1)shmget
```
//进程一 read
#include<sys/shm.h>
#include <sys/types.h>
#include <sys/ipc.h>
#define MEM_KEY (1234)

typedef struct _shared{
    int text[10];
}shared;

int main(){
    key_t key = ftok("file",0x03); //proj_id是一个1－255之间的一个整数值，典型的值是一个ASCII值
    int shmid = shmget((key_t)MEM_KEY, sizeof(shared),0666|IPC_CREAT|IPC_EXCL); //创建共享内存,如果存在则报错
    //int shmid = shmget(key,sizeof(shared,IPC_CREAT|0666));
    if (shmid == -1) perror();
    void* shm = shmat(shmid,0,0); //连接当前进程地址空间
    if（shm == (void*)-1）perror();
    shared* my = (shared*)shm;
    printf("%d\n",my->text[1]);
    if (shmdt(shm) == -1) perror();    //把共享内存从当前进程分离
    if (shmctl(shmid,IPC_RMID, 0) == -1) perror //删除共享内存
}

//进程二 write
#include<sys/shm.h>
#define MEM_KEY (1234)

typedef struct _shared{
    int text[10];
}shared;

int main(){
    int shmid = shmget((key_t)MEM_KEY, sizeof(shared),0666|IPC_CREAT); //创建共享内存
    if (shmid == -1) perror();
    shm = shmat(shmid, 0, 0);
    if（shm == (void*)-1）perror();
    shared* my = (shared*)shm;
    my->text[1] = 5;
    if (shmdt(shm) == -1) perror();
}
```
(2)shm_open+mmap
```
//server
#include <sys/mmap.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <fcntl.h>
#define  FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

typedef struct __ST{
     char text[5];
}ST;

int main(){
    shm_unlink("file");  //防止file已存在
    int fd = shm_open("file",O_RDWR|O_CREAT,FILE_MODE);
    if (fd == -1) perror();
    ftruncate(fd,sizeof(ST));
    ST* ptr;
    ptr = mmap(NULL,sizeof(ST),PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);
    if (ptr == SEM_FAILED) perror();
    ptr->text[1] = 'a';
    close(fd);

}

//client
#include <sys/mman.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <fcntl.h>
typedef struct _ST{
    char text[5];
}ST;

int main(){
    int fd = shm_open("file",O_RDWR,FILE_MODE);
    ptr = mmap(NULL,sizeof(ST),PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);
    printf("%c",ptr->text[1]);
    close(fd);
}
```
---
### 消息队列（一种有名管道）
```
消息队列提供了一种从一个进程向另一个进程发送一个数据块的方法。 
消息队列可以认为是一个消息链表，某个进程往一个消息队列中写入消息之前，不需要另外某个进程在该队列上等待消息的达到，这一点与管道和FIFO相反。
```

---
### 消息队列与rpc
```
（1）RPC系统结构：
+----------+     +----------+
| Consumer | <=> | Provider |
+----------+     +----------+
Consumer调用的Provider提供的服务。
//特点
同步调用，对于要等待返回结果/处理结果的场景，RPC是可以非常自然直觉的使用方式。# RPC也可以是异步调用。
由于等待结果，Consumer（Client）会有线程消耗。
如果以异步RPC的方式使用，Consumer（Client）线程消耗可以去掉。但不能做到像消息一样暂存消息/请求，压力会直接传导到服务Provider。

（2）Message Queue系统结构：
+--------+     +-------+     +----------+
| Sender | <=> | Queue | <=> | Receiver |
+--------+     +-------+     +----------+
Sender发送消息给Queue；Receiver从Queue拿到消息来处理。
//特点
Message Queue把请求的压力保存一下，逐渐释放出来，让处理者按照自己的节奏来处理。
Message Queue引入一下新的结点，让系统的可靠性会受Message Queue结点的影响。
Message Queue是异步单向的消息。发送消息设计成是不需要等待消息处理的完成。
所以对于有同步返回需求，用Message Queue则变得麻烦了。
```

---
### 进程间通信总结
```
服务器内进程间通信就用socket，低耦合，性能也还行，不用考虑其他乱七八糟的东西。但是如果考虑高性能，可以FIFO或者共享内存
服务期间用rpc、socket、消息队列
```

---
### 系统数据文件
1、/etc/password
```
存放着所有用户帐号的信息，包括用户名和密码，因此，它对系统来说是至关重要的
格式如下：
username:password:User ID:Group ID:comment:home directory:shell
```
2、/etc/shadow
```
存放系统的口令文件
```
3、/etc/group
```
用户组管理的文件,linux用户组的所有信息都存放在此文件中
格式如下：
组名:口令:组标识号:组内用户列表
```
4、/etc/hosts
```
Linux系统中一个负责IP地址与域名快速解析的文件
例如：
127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
```
5、/etc/services
```
Internet 守护程序（ineted）是 Linux 世界中的重要服务。它借助 /etc/services 文件来处理所有网络服务
格式如下：
service-name    port/protocol   [aliases..]  [#comment]
service-name 是网络服务的名称。例如 telnet、ftp 等。
port/protocol 是网络服务使用的端口（一个数值）和服务通信使用的协议（TCP/UDP）。
alias 是服务的别名。
comment 是你可以添加到服务的注释或说明。以 # 标记开头。
```
6、utmp和wtmp
```
/var/run/utmp（二进制） 命令 who
/var/log/wtmp（二进制） 命令 w
utmp记录当前登录到系统的用户
wtmp跟踪各个登陆与注销事件
```
7、uname和hostname
```
命令uname显示操作系统信息
命令hostname显示主机的域名
```

---
### 互斥量、信号量
```
虽然mutex和semaphore可以相互替代，可以把 值最大为1 的Semaphore当Mutex用，也可以用Mutex＋计数器当Semaphore。
Mutex管理的是资源的使用权，而Semaphore管理的是资源的数量，有那么一点微妙的小区别。
Mutex用途：保护资源共享
Semaphore用途：调度线程
简而言之，锁是服务于共享资源的；而semaphore是服务于多个线程间的执行的逻辑顺序的

//semaphore实例
int a, b, c;
void geta()
{
    a = calculatea();
    semaphore_increase();
}

void getb()
{
    b = calculateb();
    semaphore_increase();
}


void getc()
{
    semaphore_decrease();
    semaphore_decrease();
    c = a + b;
}

t1 = thread_create(geta);
t2 = thread_create(getb);
t3 = thread_create(getc);
thread_join(t3);

// semaphore_increase对应sem_post
// semaphore_decrease对应sem_wait

```

---
### 终端
```
通过网络登录或者终端登录建立的会话，会分配唯一一个tty终端或者pts伪终端（网络登录），实际上它们都是虚拟的，以文件的形式建立在/dev目录，而并非实际的物理终端。

在终端中按下的特殊按键：中断键（ctrl+c）、退出键（ctrl+\）、终端挂起键（ctrl + z）会发送给当前终端连接的会话中的前台进程组中的所有进程

在网络登录程序中，登录认证守护程序 fork 一个进程处理连接，并以ptys_open 函数打开一个伪终端设备（文件）获得文件句柄，并将此句柄复制到子进程中作为标准输入、标准输出、标准错误，所以位于此控制终端进程下的所有子进程将可以持有终端
```


---
### 守护进程 daemon
```
一个守护进程的父进程是init进程，因为它真正的父进程在fork出子进程后就先于子进程exit退出了，所以它是一个由init继承的孤儿进程。

守护进程是非交互式程序，没有控制终端，所以任何输出，无论是向标准输出设备stdout还是标准出错设备stderr的输出都需要特殊处理。

守护进程的名称通常以d结尾，比如sshd、xinetd、crond等

特点
1、没有控制终端，终端名设置为？号：也就意味着没有 stdin 0 、stdout 1、stderr 2
2、父进程不是用户创建的进程，init进程或者systemd（pid=1）以及用户人为启动的用户层进程一般以pid=1的进程为父进程，而以kthreadd内核进程创建的守护进程以kthreadd为父进程
3、守护进程一般是会话首进程、组长进程。
4、工作目录为 \ （根），主要是为了防止占用磁盘导致无法卸载磁盘

注意：
fork子进程之后，退出父进程，如果子进程还需要继续运行，则需要处理挂断信号(SIGNUP)，否则进程对挂断信号的默认处理将是退出。
```

(1) 创建步骤
```
https://www.cnblogs.com/lvyahui/p/7389554.html
https://blog.csdn.net/woxiaohahaa/article/details/53487602
https://blog.csdn.net/Leezha/article/details/78019116

1) fork()创建子进程，父进程exit()退出
2) 在子进程中调用 setsid() 函数创建新的会话(只有当该进程不是一个进程组长时，才会成功创建一个新的会话期)
3) 再次 fork() 一个子进程并让父进程退出(禁止进程重新打开控制终端。因为子进程不是会话首进程，该进程将不能重新打开控制终端)
4) 在子进程中调用 chdir() 函数，让根目录 ”/” 成为子进程的工作目录(防止占用磁盘造成磁盘不能卸载。所以也可以改到别的目录，只要保证目录所在磁盘不会中途卸载)
5) 在子进程中调用 umask(0) 函数，设置进程的文件权限掩码为0(umask取消进程本身的文件掩码设置，也就是设置Linux文件权限，一般设置为000，这是为了防止子进程创建创建一个不能访问的文件（没有正确分配权限）。此过程并非必须，如果守护进程不会创建文件，也可以不修改)
6) 信号处理signal(SIGCHLD, SIG_IGN)，忽略子进程退出信号，让子进程的回收交给Init
7) 在子进程中关闭任何不需要的文件描述符(1,2,3..n)
8) 守护进程退出处理
9) 如果是单例守护进程，结合锁文件和kill函数检测是否有进程已经运行


// 代码示例
#include <unistd.h>   
#include <signal.h>   
#include <fcntl.h>  
#include <sys/syslog.h>  
#include <sys/param.h>   
#include <sys/types.h>   
#include <sys/stat.h>   
#include <stdio.h>  
#include <stdlib.h>  
#include <time.h> 
#include <string.h>

void create_daemon() {
    signal(SIGTTOU,SIG_IGN);            //防止守护进行在没有运行起来前，控制终端受到干扰退出或挂起。
    signal(SIGTTIN,SIG_IGN);   
    signal(SIGTSTP,SIG_IGN);   
    signal(SIGHUP ,SIG_IGN); 

    pid_t pid = fork();
    if (pid == -1) {
        syslog("fork wrong");
        exit(1);
    }else if (pid > 0) {
        exit(0);                        //(1)退出主进程
    }
    if (setsid() == -1) {
        syslog("setsid wrong")
        exit(1);
    }
    pid = fork();                       //创建子进程，禁止打开终端
    if (pid == -1){
        syslog();
        exit(-1);
    }
    if (chdir("/") < 0){                //工作路径改为根目录
        syslog();
        exit(-1);
    }

    // NOFILE 为 <sys/param.h> 的宏定义
    // NOFILE 为文件描述符最大个数，不同系统有不同限制
    for(i=0; i< NOFILE; ++i){
        close(i);
    }
    umask(0);                           //取消进程的文件掩码，赋予更多的文件权限
    signal(SIGCHLD,SIG_IGN);            //忽略SIGCHLD
}

void sig_term(int signo){
    if (signo == SIGTERM){
        syslog("exit");
        closelog();
        exit(0);
    }
}

int main(){
    if (signal(SIGTERM, sig_term) != sig_term) {
        printf("error");
        exit(1);
    }
    openlog("daemontest", LOG_PID, LOG_USER);
    create_daemon();
    while(1){

    }
    return 0;
}

```

(2) setsid
```
可以建立一个新的会话期：
如果，调用setsid的进程不是一个进程组的组长，此函数创建一个新的会话期。
1、此进程变成该对话期的首进程
2、此进程变成一个新进程组的组长进程。
3、此进程没有控制终端，如果在调用setsid前，该进程有控制终端，那么与该终端的联系被解除。 如果该进程是一个进程组的组长，此函数返回错误。
4、为了保证这一点，我们先调用fork()然后exit()，此时只有子进程在运行

(3) umask
    设置默认权限(详见（linux命令.md）)

```


---
### syslog
```
https://dearhwj.gitbooks.io/itbook/content/linux/linux_syslogd.html

开机运行，由systemd启动的：
systemctl enable rsyslog
systemctl start rsyslog

linux中存在syslogd守护进程，配置文件/etc/rsyslog.conf。
一些大型公司会有rsyslog服务器，用于log的接收，特性：1.多线程，2.支持加密协议：ssl，tls，relp
//例如
rsyslog server(10.10.100.10) <-------rsyslog client(10.10.100.11)

//配置
除了一些全局配置，
https://blog.51cto.com/11555417/2163289 
https://blog.51cto.com/oldking/1891232

还有规则配置，它的每一行的格式如下：
facility.priority     action
设备,级别         动作
aexe.*  /var/log/aexe.log

//实例
https://blog.csdn.net/yuyin1018/article/details/80301274

//api
#include <syslog.h>
void openlog(const char* ident, int option, int facility);
    ident: 被加到日志中的程序名
    option: 常用值LOG_PID即包含每个消息的pid
    facility: 记录日志的程序的类型，配置文件可根据不同的登录类型来区别处理消息，常用值,这个值决定你的日志输出到哪个log文件。 LOG_DAEMON即其它系统守护进程，一般为自己创建的守护进程

void syslog(int priority, const char* format, ...);
    priority:  优先级，说明消息的重要性，可取值如下： 
                LOG_ERR 错误;  LOG_WARNING 警告; LOG_NOTICE 正常情况，但较为重要;    LOG_INFO 信息; LOG_DEBUG 调试信息

void closelog(void);

//add syslog facility
在/usr/include/sys/syslog.h中系统已经定义了不同log设备的宏，又提供了15个LOG_LOCAL可用于私人定制的log。
/* other codes through 15 reserved for system use */
#define LOG_LOCAL0      (16<<3) /* reserved for local use */
#define LOG_LOCAL1      (17<<3) /* reserved for local use */
#define LOG_LOCAL2      (18<<3) /* reserved for local use */
#define LOG_LOCAL3      (19<<3) /* reserved for local use */
#define LOG_LOCAL4      (20<<3) /* reserved for local use */
#define LOG_LOCAL5      (21<<3) /* reserved for local use */
#define LOG_LOCAL6      (22<<3) /* reserved for local use */
#define LOG_LOCAL7      (23<<3) /* reserved for local use */
所以如果要新建自己的syslog，那么在规则配置添加
LOG_LOCAL0.*    /var/log/my.log
在程序中
openlog('my', LOG_PID, LOG_LOCAL0)
```

---
### 进程、进程组、会话
```
https://segmentfault.com/a/1190000009152815

一个用户 -> tty -> 会话 -> 进程组 -> 进程

一个进程组可以包含多个进程，为了方便管理这些进程。一个会话又可以包含多个进程组。一个会话对应一个控制终端。linux是一个多用户多任务的分时操作系统，必须要支持多个用户同时登陆同一个操作系统，当一个用户登陆一次终端时就会产生一个会话。

// session进程
session包含一个前台进程组及一个或多个后台进程组，一个进程组包含多个进程。
session id就是这个session中leader的进程ID。
session的leader退出后，session中的所有其它进程将会收到SIGHUP信号，其默认行为是终止进程（退出终端，其下所有进程退出）

// 进程组
对大部分进程来说，它自己就是进程组的leader，并且进程组里面就只有它自己一个进程

// 父子进程
    父进程退出时，子进程可以继续运行

// 前后台进程组
与终端交互的进程是前台进程，否则便是后台进程
后台进程也叫守护进程，默认情况下，只要后台进程组的任何一个进程读tty，将会使整个进程组的所有进程暂停

// session和进程组的关系
    登陆shell               一个或多个后台进程组      一个前台进程组
    会话首进程=控制进程
  +--------------------------------------------------------------+
                            session
    注意：
        deamon程序虽然也是一个session的leader，但一般它不会创建新的进程组，也没有job的管理功能，所以这种情况下一个session就只有一个进程组，所有的进程都属于同样的进程组和session。

// SIGHUP 与 nohup
sighup（挂断）信号在控制终端或者控制进程死亡时向关联会话中的进程发出，默认进程对SIGHUP信号的处理时终止程序，所以我们在shell下建立的程序，在登录退出连接断开之后，会一并退出。

// nohup做了3件事情
1、dofile函数将输出重定向到nohup.out文件
2、signal函数设置SIGHUP信号处理函数为SIG_IGN宏（指向sigignore函数），以此忽略SIGHUP信号
3、execvp函数用新的程序替换当前进程的代码段、数据段、堆段和栈段。

// 如何父进程退出，确保子进程也退出
    https://zhuanlan.zhihu.com/p/56833833
    pid_t pid = fork()
    if (pid == 0) {
        /*父进程退出时，会收到SIGKILL信号*/
        prctl(PR_SET_PDEATHSIG,SIGKILL);
        .....;
    }

```

### /var 目录
```
系统一般运行时要改变的数据.每个系统是特定的，即不通过网络与其他计算机共享.  
/var/lib            系统正常运行时要改变的文件.

/var/local          /usr/local 中安装的程序的可变数据

/var/lock           锁定文件.以支持他们正在使用某个特定的设备或文件.其他程序注意到这个锁定文件，将不试图使用这个设备或文件.  

/var/log            各种程序的Log文件，
/var/log/wtmp       永久记录每个用户登录、注销及系统的启动、停机的事件。
/var/log/lastlog    记录最近成功登录的事件和最后一次不成功的登录事件，由login生成

/var/run            保存到下次引导前有效的关于系统的信息文件
/var/run/utmp       记录著現在登入的用戶。

/var/tmp            比/tmp 允许的大或需要存在较长时间的临时文件
```

### stdout stderr
```
stdout是行缓冲的，他的输出会放在一个buffer里面，只有到换行的时候，才会输出到屏幕。而stderr是无缓冲的，会直接输出
如果用转向标准输出到磁盘文件，则可看出两者区别。stdout输出到磁盘文件，stderr在屏幕。 

```

### 字符编码
```
//语言环境设置环境变量 LANG
//include <locale.h>
//char *setlocale(int category, const char *locale);
setlocale(LC_ALL, "utf8");

//注意：utf8编码的字符串含中文和英文时，注意长度！！！

```

### capability
```
//从2.1版开始,Linux内核有了能力(capability)的概念,即它打破了UNIX/LINUX操作系统中超级用户/普通用户的概念,由普通用户也可以做只有超级用户可以完成的工作

//capability可以作用在进程上(受限),也可以作用在程序文件上,它与sudo不同,sudo只针对用户/程序/文件的概述,即sudo可以配置某个用户可以执行某个命令,可以更改某个文件,而capability是让某个程序拥有某种能力

//每个进程有三个和能力有关的位图:inheritable(I),permitted(P)和effective(E), 在/proc/PID/status中
cat /proc/$$/status | egrep 'Cap(Inh|Prm|Eff)'
    CapInh: 0000000000000000
    CapPrm: 0000000000000000
    CapEff: 0000000000000000

//cap_effective:当一个进程要进行某个特权操作时,操作系统会检查cap_effective的对应位是否有效,而不再是检查进程的有效UID是否为0. 例如,如果一个进程要设置系统的时钟,Linux的内核就会检查cap_effective的CAP_SYS_TIME位(第25位)是否有效.
//cap_permitted:表示进程能够使用的能力,在cap_permitted中可以包含cap_effective中没有的能力，这些能力是被进程自己临时放弃的,也可以说cap_effective是cap_permitted的一个子集.
//cap_inheritable:表示能够被当前进程执行的程序继承的能力.
//总结：permitted表示可以使用的能力（有可能没有激活），effective表示激活的能力（可以使用的，前提是permitted要有这种能力），inheritable表示可以继承的能力

CAP_CHOWN:修改文件属主的权限
CAP_DAC_OVERRIDE:忽略文件的DAC访问限制
CAP_DAC_READ_SEARCH:忽略文件读及目录搜索的DAC访问限制
CAP_FOWNER：忽略文件属主ID必须和进程用户ID相匹配的限制
CAP_FSETID:允许设置文件的setuid位
CAP_KILL:允许对不属于自己的进程发送信号
CAP_SETGID:允许改变进程的组ID
CAP_SETUID:允许改变进程的用户ID
CAP_SETPCAP:允许向其他进程转移能力以及删除其他进程的能力
CAP_LINUX_IMMUTABLE:允许修改文件的IMMUTABLE和APPEND属性标志
CAP_NET_BIND_SERVICE:允许绑定到小于1024的端口
CAP_NET_BROADCAST:允许网络广播和多播访问
CAP_NET_ADMIN:允许执行网络管理任务
CAP_NET_RAW:允许使用原始套接字
CAP_IPC_LOCK:允许锁定共享内存片段
CAP_IPC_OWNER:忽略IPC所有权检查
CAP_SYS_MODULE:允许插入和删除内核模块
CAP_SYS_RAWIO:允许直接访问/devport,/dev/mem,/dev/kmem及原始块设备
CAP_SYS_CHROOT:允许使用chroot()系统调用
CAP_SYS_PTRACE:允许跟踪任何进程
CAP_SYS_PACCT:允许执行进程的BSD式审计
CAP_SYS_ADMIN:允许执行系统管理任务，如加载或卸载文件系统、设置磁盘配额等
CAP_SYS_BOOT:允许重新启动系统
CAP_SYS_NICE:允许提升优先级及设置其他进程的优先级
CAP_SYS_RESOURCE:忽略资源限制
CAP_SYS_TIME:允许改变系统时钟
CAP_SYS_TTY_CONFIG:允许配置TTY设备
CAP_MKNOD:允许使用mknod()系统调用
CAP_LEASE:允许修改文件锁的FL_LEASE标志

```

### /proc文件夹
```
/proc 文件系统是一种内核和内核模块用来向进程(process) 发送信息的机制, /proc 存在于内存之中而不是硬盘上。proc文件系统以文件的形式向用户空间提供了访问接口，这些接口可以用于在运行时获取相关部件的信息或者修改部件的行为，因而它是非常方便的一个接口。

(1) 内容介绍
    /proc/cpuinfo - CPU 的信息(型号, 家族, 缓存大小等)
    /proc/meminfo - 物理内存、交换空间等的信息
    /proc/mounts - 已加载的文件系统的列表
    /proc/devices - 可用设备的列表
    /proc/filesystems - 被支持的文件系统
    /proc/modules - 已加载的模块
    /proc/version - 内核版本
    /proc/cmdline - 系统启动时输入的内核命令行参数

    /proc/pid/*     pid进程的相关信息

    /proc/sys/kernel    与内核相关
```

### /var/run
```
/var/run 目录中存放的是自系统启动以来描述系统信息的文件。比较常见的用途是daemon进程将自己的pid保存到这个目录。标准要求这个文件夹中的文件必须是在系统启动的时候清空，以便建立新的文件。

(1) /var/run/*.pid
在工作中遇到了很多在程序启动时检查是否已经重复启动的代码段，其核心就是调用fcntl设置pid文件的锁定F_SETLK状态，其中锁定的标志为F_WRLACK。如果成功锁定，则写入进程当前PID，进程继续往下执行。如果锁定不成功，说明已经有同样的进程在运行了，当前进程结束退出

(2) 
```

### perror
```
#include <perror.h>

当linux的系统C api函数发生异常时,一般会将errno变量赋一个整数值,不同的值表示不同的含义

void perror(const char *s)
用来将上一个函数发生错误的原因输出到标准错误(stderr)
```

### linux信号
```
https://zhuanlan.zhihu.com/p/66051508
https://cloud.tencent.com/developer/column/1601
https://cloud.tencent.com/developer/article/1007500

(1) 概念
    1) 信号类似软中断
        信号与中断的相似点：
            采用了相同的异步通信方式；
            当检测出有信号或中断请求时，都暂停正在执行的程序而转去执行相应的处理程序；
            都在处理完毕后返回到原来的断点；
            对信号或中断都可进行屏蔽。
        信号与中断的区别：
            中断有优先级，而信号没有优先级，所有的信号都是平等的；
            信号处理程序是在用户态下运行的，而中断处理程序是在核心态下运行；
            中断响应是及时的，而信号响应通常都有较大的时间延迟。
    2) 分类
        信号的处理有三种方法，分别是：忽略、捕捉和默认动作，有两种信号不能被忽略（分别是 SIGKILL和SIGSTOP）
        可靠信号和不可靠信号
            那些建立在早期机制上的信号叫做不可靠信号，信号值小于SIGRTMIN(32)，主要表现为：信号可能丢失(因为不支持排队)
            信号值位于SIGRTMIN和SIGRTMAX之间的信号都是可靠信号，支持排队克服了信号可能丢失的问题
        实时信号与非实时信号
            非实时信号都不支持排队，都是不可靠信号；实时信号都支持排队，都是可靠信号

(2) 信号捕捉
    1) 信号---> 是否可被中断优先级上 ---->(是) 陷入内核 ---> 2)
                是否可重入函数      ---->(否) 仅设置进程表中信号域相应的位

    2) 实例:
        https://cloud.tencent.com/developer/article/1008813
        > 用户程序注册了SIGQUIT信号的处理函数sighandler。
        > 当前正在执行main函数，某条指令发生中断或异常或产生信号切换到内核态。（可重入函数）
        > 在中断处理完毕后要返回用户态的main函数之前检查到有信号SIGQUIT递达。（进程检查信号的时机）
        > 内核决定返回用户态后不是恢复main函数的上下文继续执行，而是执行sighandler函数，sighandler和main函数使用不同的堆栈空间，它们之间不存在调用和被调用的关系，是两个独立的控制流程。
        > sighandler函数返回后自动执行特殊的系统调用sigreturn再次进入内核态。
        > 如果没有新的信号要递达，这次再返回用户态就是恢复main函数的上下文继续执行了。

(3) 信号发送和睡眠
    https://cloud.tencent.com/developer/article/1007500
    1) 信号注册
        #include <signal.h>
        #define SIG_ERR ((__sighandler_t) -1)
        #define SIG_DFL ((__sighandler_t) 0)
        #define SIG_IGN ((__sighandler_t) 1)
        typedef void (*sighandler_t)(int);
        sighandler_t signal(int signum, sighandler_t handler);
        返回一个函数指针
        if (signal(SIGUSR1, sig_usr) == SIG_ERR) wrong!
        if (signal(SIGUSR1, sig_usr) == sig_usr) yes!
        经过sigaction安装的信号都能传递信息给信号处理函数，而经过signal安装的信号不能向信号处理函数传递信息。对于信号发送函数来说也是一样的。

    2) 信号发送和睡眠函数
        #include <sys/types.h>
        #include <signal.h>
        int kill(pid_t pid, int sig);
        int raise(int sig);
        int killpg(int pgrp, int sig);
        kill函数的给进程pid发送信号。raise函数可以给当前进程发送指定的信号（自己给自己发信号）。killpg 函数可以给进程组发生信号。这三个函数都是成功返回0，错误返回-1

        #include <unistd.h>
        unsigned int alarm(unsigned int seconds);
        告诉内核在seconds秒之后给当前进程发SIGALRM信号，该信号的默认处理动作是终止当前进程
        #include <stdlib.h>
        void abort(void);
        abort函数使当前进程接收到SIGABRT信号而异常终止。就像exit函数一样，abort函数总是会成功的，所以没有返回值。

    3) pause
        使进程挂起直到一个信号被捕获(信号处理函数完成后返回)
        且调用schedule()使系统调度其他程序运行，
        在死循环中调用比完全的死循环的好处是让出cpu

(5) 阻塞和未决
    https://cloud.tencent.com/developer/article/1008811
    > 实际执行信号的处理动作称为信号递达（Delivery）
      信号从产生到递达之间的状态，称为信号未决（Pending）
      进程可以选择阻塞（Block）某个信号，SIGKILL 和 SIGSTOP 不能被阻塞
    > 信号的相关状态字
        每一位代表一个信号
        屏蔽状态字(block)、未决状态字(pending)、信号处理方法
        信号屏蔽状态字: 1代表阻塞、0代表不阻塞
        信号未决状态字: 1代表未决，0代表信号可以抵达了
        信号处理方法:   对应的某一位(信号)的处理函数
        block   pending   handler
          0       0       SIG_DEL                       未阻塞、信号没产生
          0       1       SIG_IGN                       未阻塞、信号产生（马上处理）
          1       1       void sighandler(int signo)      阻塞、信号产生处于未决状态（解除阻塞才会处理）
          1       0       void sighandler(int signo)      阻塞、信号没产生

(6) 信号集
    表示多个信号的数据类型
    #include <signal.h>
    int sigemptyset(sigset_t *set);
    int sigfillset(sigset_t *set);
    int sigaddset(sigset_t *set, int signo);
    int sigdelset(sigset_t *set, int signo);
    int sigismember(const sigset_t *set, int signo);

    int sigprocmask(int how, const sigset_t *set, sigset_t *oset);
    读取或更变进程的信号屏蔽字
    int sigpending(sigset_t *set);
    读取当前进程的未决信号集

(8) 信号
    SIGHUP
        终端关闭时，该信号发送到session首进程和后台进程（&提交的）
        session首进程退出时，该信号被发送到该session中的前台进程组中的每一个进程

    SIGCHLD
        一个进程终止或停止时，此信号发送给父进程，默认是忽略此信号

    SIGTERM
        SIGTERM是杀或的killall命令发送到进程默认的信号。它会导致一过程的终止，但是SIGKILL信号不同，它可以被捕获和解释（或忽略）的过程。因此，SIGTERM类似于问一个进程终止可好，让清理文件和关闭

```

### fork vs vfork
```
(1)
    fork 是 创建一个子进程，并把父进程的内存数据copy到子进程中(写时复制)。
    vfork是 创建一个子进程，并和父进程的内存数据share一起用。

(2) vfork是这样的工作的
    1）保证子进程先执行，注意此时的子进程共享父进程的栈，如果子进程return，那么此函数就会退出
    2）当子进程调用exit()或exec()后，父进程往下执行。

(3) 为什么vfork
    起初只有fork，但是很多程序在fork一个子进程后就exec一个外部程序，于是fork需要copy父进程的数据这个动作就变得毫无意了，而且这样干还很重（注：后来，fork做了优化，详见本文后面），所以，BSD搞出了个父子进程共享的 vfork，这样成本比较低。因此，vfork本就是为了exec而生。
```

### 进程控制相关概念
```
(0) 控制终端与作业控制


(1) 父子进程区别
    子进程不继承父进程的文件锁
    子进程的未处理信号集设置为空集

(2) exec
    exec不会创建新进程，ID不会改变。它会替换当前的进程为新程序。但新程序会从调用进程继承：进程、父进程、组、session、附属ID，控制终端，当前工作目录、根目录、文件锁、进程信号屏蔽、未处理信号、资源限制

(3) 进程会计
    启用该选项后，每当进程结束时内核会写个会计记录，这也是我们得到了一个再次观察进程的记录
    #include <sys/acct.h>

(4) 进程调度
    1) nice值越小，优先级越高
    2) 相关函数
        #include <unistd.h>
        int nice(int);
        
        #include <sys/resource.h>
        int getpriority(int which, id_t who);
        which的值: PRIO_PROCESS表示进程，PRIO_PGRP表示进程组，PRIO_USER表示用户ID
        who: 为0时，表示调用进程、进程组或用户ID

        #include <sys/resource.h>
        int setpriority(int which, id_t who, int value);
        设置用户、进程、进程组的优先级
```

### 进程关系
```

```