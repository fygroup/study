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

### 异步的进化
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
            db = await mongoDb.open();   //切出该协程，
            collection = await db.collection("users");
            result = await collection.insert(person);
        }catch(e){
            console.error(e.message);
        }
        console.log(result);
    } 

```

### 协程的原理
```
https://www.zhihu.com/question/65647171/answer/233495694
https://zhuanlan.zhihu.com/p/25964339

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
    注意：此种实现不能用于try catch、递归循环

    #define BEGIN_CORO void operator()() { switch(next_line) { case 0:
    #define YIELD next_line=__LINE__; break; case __LINE__:
    #define END_CORO }} int next_line=0

    struct coroutine{
        int n_;
        coroutine(int n__):n_(n__){}
        void operator()(){
            case 0:
                cout << n_++ << end;
                next_line = __LINE__ + 2;
                break;
            case __LINE__:
                cout << n_++ << end;
                next_line = __LINE__ + 2;
                break;
            case __LINE__:
                cout << n_++ << end;
                next_line = __LINE__ + 2;
                break;
            case         
                next_line = 0;
        }
        int next_line;
    }
    

async/await的出现，实现了基于stackless coroutine的完整coroutine。在特性上已经非常接近stackful coroutine了，不但可以嵌套使用也可以支持try catch。所以是不是可以认为async/await是一个更好的方案？

```

### c++异步编程
```
c++11的promise、async、future属于多线程异步，所以单线程异步只能用协程和异步callback
异步是协程的一种实现方式，协程是异步的封装方法
1、promise、async的多线程异步回调（异步工作流）
2、协程（更优雅）
```

### c++高性能网络库 
```
libevent、libev、boost::asio
```

### c++ asio
```
推荐boost::asio
c++20标准库网络部份将基于asio，c++ asio异步编程很重要！！！
```

### 有栈协程相关模块
```
云风的coroutine库
libgo
golang
boost::asio
```

### 无栈协程相关模块
```
c++20的coroutine(基于asio)
知乎朱元的库
Es6的async/wait模型
boost::asio
```

### linux hook
```
由于是调用得动态链接库中函数，我们可以通过劫持该函数的方式引入额外处理。 例如通过劫持 malloc、free 来追踪内存使用情况等等

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

//dlsym
void dlsym(RTLD_NEXT,"malloc")
RTLD_DEFAULT: 使用默认的库搜索顺序查找所需符号的第一次出现
RTLD_NEXT: 在当前库之后的搜索顺序中查找下一个出现的函数

```

### fuse 用户空间文件系统
```
[fuse架构](../picture/fuse架构.png)
1、用户态app调用glibc open接口，触发sys_open系统调用。
2、sys_open 调用fuse中inode节点定义的open方法。
3、inode中open生成一个request消息，并通过/dev/fuse发送request消息到用户态libfuse。
4、Libfuse调用fuse_application用户自定义的open的方法，并将返回值通过/dev/fuse通知给内核。
5、内核收到request消息的处理完成的唤醒，并将结果放回给VFS系统调用结果。
```

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

### 终端
```
通过网络登录或者终端登录建立的会话，会分配唯一一个tty终端或者pts伪终端（网络登录），实际上它们都是虚拟的，以文件的形式建立在/dev目录，而并非实际的物理终端。

在终端中按下的特殊按键：中断键（ctrl+c）、退出键（ctrl+\）、终端挂起键（ctrl + z）会发送给当前终端连接的会话中的前台进程组中的所有进程

在网络登录程序中，登录认证守护程序 fork 一个进程处理连接，并以ptys_open 函数打开一个伪终端设备（文件）获得文件句柄，并将此句柄复制到子进程中作为标准输入、标准输出、标准错误，所以位于此控制终端进程下的所有子进程将可以持有终端
```

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
    signal(SIGTTOU,SIG_IGN); // 防止守护进行在没有运行起来前，控制终端受到干扰退出或挂起。
    signal(SIGTTIN,SIG_IGN);   
    signal(SIGTSTP,SIG_IGN);   
    signal(SIGHUP ,SIG_IGN); // 防止会话首进程退出时，发送信号终止进程组中的进程

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
    }else if (pid > 0){
        exit(0);
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
// SIGHUP
SIGHUP会在以下3种情况下被发送给相应的进程：
1、终端关闭时，该信号被发送到session首进程以及作为job提交的进程（即用 & 符号提交的进程）
2、session首进程退出时，该信号被发送到该session中的前台进程组和后台进程组中的每一个进程
3、若进程的退出，导致一个进程组变成了孤儿进程组，且新出现的孤儿进程组中有进程处于停止状态，则SIGHUP和SIGCONT信号会按顺序先后发送到新孤儿进程组中的每一个进程。
系统对SIGHUP信号的默认处理是终止收到该信号的进程。所以若程序中没有捕捉该信号，当收到该信号时，进程就会退出。

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
    CapInh: 0000000000000000    //16位
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

### perror
```
#include <perror.h>

当linux的系统C api函数发生异常时,一般会将errno变量赋一个整数值,不同的值表示不同的含义

void perror(const char *s)
用来将上一个函数发生错误的原因输出到标准错误(stderr)
```

### 可重入与不可重入函数
```
https://www.jianshu.com/p/2c8de98bf0db

(1) 可重入的概念
    若一个程序或子程序可以在任意时刻被中断，然后操作系统调度执行另外一段代码，这段代码又调用了该子程序不会出错，则称其为可重入（reentrant或re-entrant）的。

    简单来说就是可以被中断的函数，也就是说，可以在这个函数执行的任何时刻中断它，转入OS调度下去执行另外一段代码，而返回控制时不会出现什么错误
    
    也就是说，当该子程序正在运行时，执行线程可以再次进入并执行它，仍然获得符合设计时预期的结果。与多线程并发执行的线程安全不同，可重入强调对单个线程执行时，重新进入同一个子程序，仍然是安全的。
    
    可重入的函数,并且不能在原子上下文中运行


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

### open flag mode
```
#include <sys/types.h>    
#include <sys/stat.h>    
#include <fcntl.h>

int open(const char * pathname, int flags);
int open(const char * pathname, int flags, mode_t mode);

O_RDONLY 以只读方式打开文件
O_WRONLY 以只写方式打开文件
O_RDWR 以可读写方式打开文件. 上述三种旗标是互斥的, 也就是不可同时使用, 但可与下列的旗标利用OR(|)运算符组合.
O_CREAT 若欲打开的文件不存在则自动建立该文件.
O_EXCL 如果O_CREAT 也被设置, 此指令会去检查文件是否存在. 文件若不存在则建立该文件, 否则将导致打开文件错误. 此外, 若O_CREAT 与O_EXCL 同时设置, 并且欲打开的文件为符号连接, 则会打开文件失败.
O_NOCTTY 如果欲打开的文件为终端机设备时, 则不会将该终端机当成进程控制终端机.
O_TRUNC 若文件存在并且以可写的方式打开时, 此旗标会令文件长度清为0, 而原来存于该文件的资料也会消失.
O_APPEND 当读写文件时会从文件尾开始移动, 也就是所写入的数据会以附加的方式加入到文件后面.
O_NONBLOCK 以不可阻断的方式打开文件, 也就是无论有无数据读取或等待, 都会立即返回进程之中.
O_NDELAY 同O_NONBLOCK.
O_SYNC 以同步的方式打开文件
O_DIRECT 绕过缓冲区高速缓存，直接IO
O_NOFOLLOW 如果参数pathname 所指的文件为一符号连接, 则会令打开文件失败.


S_IRWXU00700 权限, 代表该文件所有者具有可读、可写及可执行的权限.
S_IRUSR 或S_IREAD, 00400 权限, 代表该文件所有者具有可读取的权限.
S_IWUSR 或S_IWRITE, 00200 权限, 代表该文件所有者具有可写入的权限.
S_IXUSR 或S_IEXEC, 00100 权限, 代表该文件所有者具有可执行的权限.
S_IRWXG 00070 权限, 代表该文件用户组具有可读、可写及可执行的权限.
S_IRGRP 00040 权限, 代表该文件用户组具有可读的权限.
S_IWGRP 00020 权限, 代表该文件用户组具有可写入的权限.
S_IXGRP 00010 权限, 代表该文件用户组具有可执行的权限.
S_IRWXO 00007 权限, 代表其他用户具有可读、可写及可执行的权限.
S_IROTH 00004 权限, 代表其他用户具有可读的权限
S_IWOTH 00002 权限, 代表其他用户具有可写入的权限.
S_IXOTH 00001 权限, 代表其他用户具有可执行的权限.
```

### 文件缓冲
```
1、O_DIRECT、O_SYNC
    int fd = open("xxx", O_RDWR|O_CREAT|O_SYNC|O_DIRECT);
    (1) O_DIRECT
        无缓冲输入、输出
        允许应用程序在执行磁盘IO时绕过缓冲区高速缓存，从用户空间直接将数据传递到文件或磁盘设备，称为直接IO
        > 应用场景
            数据库系统，其高速缓存和IO优化机制均自成一体，无需内核
        > 缺点
            会大大降低性能
        > 注意
            可能发生的不一致性：一进程以O_DIRECT标志打开某文件，而另一进程以普通（即使用了高速缓存缓冲区）打开同一文件，则由直接IO所读写的数据与缓冲区高速缓存中内容之间不存在一致性，应尽量避免这一场景
        > 直接IO应该遵守的限制(否则导致EINVAL错误)
            用于传递数据的缓冲区，其内存边界必须对齐为块大小的整数倍
            数据传输的开始点，即文件和设备的偏移量，必须是块大小的整数倍
            待传递数据的长度必须是块大小的整数倍。
    (2) O_SYNC
        以同步IO方式打开文件，强制刷新内核缓冲区到输出文件，确保数据的安全

2、fsync
    int fsync(int fd);
    调用fsync会将使缓冲数据和fd相关的所有元数据都刷新到磁盘上
    采用O_SYNC或频繁调用fsync对性能影响很大

3、fflush
    #include <stdio.h>
    int fflush(FILE *f);
    将内存缓冲写到内核的缓冲
    标准IO函数（如fread，fwrite等）会在内存中建立缓冲，该函数刷新内存缓冲，将内容写入内核缓冲，而要想将其真正写入磁盘，还需要调用fsync
    
    内存缓冲 -> fflush -> 内核缓冲 -> fsync -> 磁盘

4、内存缓冲区操作
    (1) 查看缓冲区大小
        printf("%d", BUFSIZ);

    (2) setvbuf
        int setvbuf(FILE * stream, char * buf, int type, unsigned size);
        // stream为文件流指针，buf为缓冲区首地址，type为缓冲区类型，size为缓冲区内字节的数量
        // type
            _IOFBF (满缓冲)：当缓冲区为空时，从流读入数据。或当缓冲区满时，向流写入数据。
            _IOLBF (行缓冲)：每次从流中读入一行数据或向流中写入—行数据。
            _IONBF (无缓冲)：直接从流中读入数据或直接向流中写入数据，而没有缓冲区。
        // 成功返回0，失败返回非0。
        // setvbuf(input, bufr, _IOFBF, 512)

    (3) setbuf
        void setbuf(FILE * stream, char * buf);
        // stream为文件流指针，buf为缓冲区的起始地址
```

### 进程关系
```

```




