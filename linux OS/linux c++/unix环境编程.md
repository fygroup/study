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
2、父进程不是用户创建的进程，init进程或者systemd（pid=1）以及用户人为启动的用户层进程一般以pid=1的进程为父进程，而以kthread内核进程创建的守护进程以kthread为父进程
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
Linux下可通过 prctl进行对进程的各种控制，例如设置/获得进程名

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
```c++
#include <perror.h>
void perror(const char *s)

// 当linux的系统C api函数发生异常时, 会把错误码(一个整数)写进error全局变量

// 调用perror(string)用来将上一个函数发生错误的原因输出到标准错误(stderr)

// 参数string所指的字符串会先打印出可以把变量翻译成用户理解的字符串

// 例如
fp = fopen("file.txt", "r");    // 没有文件，系统把错误已经写进error中
if( fp == NULL ) {
    perror("Error: ");          // 输出error得内容，并在前面加上"Error: "
    return(-1);
}

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

### linux下的几种定时器
```
1、sleep()和usleep()
    sleep精度是1秒，usleep精度是1微妙
    使用这种方法缺点比较明显，在Linux系统中，sleep类函数不能保证精度，尤其在系统负载比较大时，sleep一般都会有超时现象

2、信号量SIGALRM + alarm()
    精度能达到1秒，其中利用了系统的信号量机制，首先注册信号量SIGALRM处理函数，调用alarm()，设置定时长度
    signal(SIGALRM, func);  //注册func
    alarm(1);   //  1s后向本进程发送SIGALRM，此函数非阻塞

3、RTC
    RTC机制利用系统硬件提供的Real Time Clock机制，通过读取RTC硬件/dev/rtc，通过ioctl()设置RTC频率
    精度可调，而且非常高

4、select
    通过使用select()，来设置定时器
    原理利用select()方法的第5个参数，第一个参数设置为0，三个文件描述符集都设置为NULL，第5个参数为时间结构体
    这种方法精度能够达到微妙级别，网上有很多基于select()的多线程定时器，说明select()稳定性还是非常好

5、timefd
    since kernel 2.6.25
    timerfd是Linux为用户程序提供的一个定时器接口，基于文件描述符，可以和epoll一起用

```

### RTC
```c++
// linux系统有两个时钟
一个是由主板电池驱动的"Real Time Clock"也叫做RTC或者叫CMOS时钟，硬件时钟。当操作系统关机的时候，用这个来记录时间，但是对于运行的系统是不用这个时间的。
另一个时间是"System clock"也叫内核时钟或者软件时钟，是由软件根据时间中断来进行计数的，内核时钟在系统关机的情况下是不存在的，所以，当操作系统启动的时候，内核时钟是要读取RTC时间来进行时间同步。并且在系统关机的时候将系统时间写回RTC中进行同步。

// Linux内核与RTC进行互操作的时机只有两个：
1) 内核在启动时从RTC中读取启动时的时间与日期
2) 内核在需要时将时间与日期回写到RTC中

// 通过 /dev/rtc 硬件时钟实现timer
    unsigned long data = 0;
    int fd = open("/dev/rtc", O_RDONLY);
    /* set the freq as 4Hz */
    ioctl(fd, RTC_IRQP_SET, 4);
    /* enable periodic interrupts */
    ioctl(fd, RTC_PIE_ON, 0);
    struct timeval tv;
    for (size_t i = 0; i < 100; i++ )
    {
        read(fd, &data, sizeof(data));
        gettimeofday(&tv, NULL);
        uint64_t a = (uint64_t)tv.tv_sec * 1000 + (uint64_t)tv.tv_usec / 1000;
        printf("timer %u\n", a);
    }
    /* enable periodic interrupts */
    ioctl(fd, RTC_PIE_OFF, 0);
    close(fd);

#define RTC_AIE_ON	    _IO('p', 0x01)	/* Alarm int. enable on		*/
#define RTC_AIE_OFF	    _IO('p', 0x02)	/* ... off			*/
#define RTC_UIE_ON	    _IO('p', 0x03)	/* Update int. enable on	*/
#define RTC_UIE_OFF	    _IO('p', 0x04)	/* ... off			*/
#define RTC_PIE_ON	    _IO('p', 0x05)	/* Periodic int. enable on	*/
#define RTC_PIE_OFF	    _IO('p', 0x06)	/* ... off			*/
#define RTC_WIE_ON	    _IO('p', 0x0f)  /* Watchdog int. enable on	*/
#define RTC_WIE_OFF 	_IO('p', 0x10)  /* ... off			*/

#define RTC_ALM_SET 	_IOW('p', 0x07, struct rtc_time) /* Set alarm time  */
#define RTC_ALM_READ	_IOR('p', 0x08, struct rtc_time) /* Read alarm time */
#define RTC_RD_TIME	    _IOR('p', 0x09, struct rtc_time) /* Read RTC time   */
#define RTC_SET_TIME	_IOW('p', 0x0a, struct rtc_time) /* Set RTC time    */
#define RTC_IRQP_READ	_IOR('p', 0x0b, unsigned long)	 /* Read IRQ rate   */
#define RTC_IRQP_SET	_IOW('p', 0x0c, unsigned long)	 /* Set IRQ rate    */

```

### 文件大小、遍历目录
```
#include<sys/stat.h>  
#include<unistd.h>
#include <dirent.h>

struct stat buf;
if (stat("filename", &buf)) error;
//文件大小
buf.st_size；

struct dirent* ptr;
DIR* dir = opendir("/");
// 遍历文件夹里的文件
while((ptr = readdir(dir)) != NULL){
    ptr->d_name;
    ptr->d_type;     // 8表示file，10表示linkfile，4表示dir
}

```

### 程序暂停，代码可控
```
// windows
system("pause")

// linux
cin.get()   // c++
getchar()   // c
```

### time.h
```c++
// time_t 在有的编译器上是32，但大多是64

std::string TimeToStr(int64_t timeParam) {
    time_t tt = static_cast<time_t>(timeParam);
    tm* tb = localtime(&tt);
    char timebuf[30];
    strftime(timebuf, 30, "%Y-%m-%d %H:%M:%S", tb);
    return std::string(timebuf);
}

int64_t StrToTime(const std::string & timeStrParam){
    tm tb = {0};
    strptime(timeStrParam.c_str(), "%Y-%m-%d %H:%M:%S", &tb);
    // cout << "year: " << tb.tm_year + 1900 << endl; // 注意
    // cout << "mouth: " << tb.tm_mon + 1 << endl; // 注意
    // cout << "day: " << tb.tm_mday << endl;
    // cout << "hour: " << tb.tm_hour << endl;
    // cout << "min: " << tb.tm_min << endl;
    // cout << "sec: " << tb.tm_sec << endl;
    time_t tt = mktime(&tb);
    return static_cast<int64_t>(tt);
}
```

### 位运算条件判断
```
// Linux权限控制是基于位运算实现的
// 在Linux权限系统中，读、写、执行权限分别对应三个状态位
读   写   执行      二进制     十进制
0    0    1   ==>   001  ==>  1
0    1    0   ==>   010  ==>  2
1    0    0   ==>   100  ==>  4

// 或运算实现权限的添加
001 | 010 = 011
1 | 2 = 3

// 与运算实现权限的判断
011 & 010 > 0
3 & 1 > 0

011 & 001 > 0
3 & 2 > 0

// 非运算实现权限的减少
011 ^ 001 = 010
3 ^ 1 = 2

// 位移与权限码
从上面的介绍可以看出，在权限管理系统中每操作的权限码都是唯一的
2 << 0, 2 << 1, 2 << 2, 2 << 3 ...

```

### eventfd
```c++
// eventfd是一个用来通知事件的文件描述符，是一种linux上的线程通信方式
// 可以用来实现用户空间的事件/通知驱动的应用程序
// 和信号量等其他线程通信不同的是eventfd可以用于进程间的通信，还可以用于内核发信号给用户态的进程

int eventfd(unsigned int initval, int flags);
// 产生一个文件描述符，可以对这个文件描述符进行read、write、poll、select等操作
// 该对象是一个内核维护的无符号的64位整型计数器，初始化为initval的值
// read(): 读操作就是将counter值置0，如果是semophore就减1
// write(): 设置counter的值
// flag
//      EFD_CLOEXEC     like O_CLOEXEC
//      EFD_NONBLOCK    like O_NONBLOCK
//      EFD_SEMAPHORE   支持semophore语义的read，简单说就值递减1

(1) read: 读取计数器中的值
> 如果计数器中的值大于0
    > 设置了EFD_SEMAPHORE标志位，则返回1，且计数器中的值也减去1
    > 没有设置EFD_SEMAPHORE标志位，则返回计数器中的值，且计数器置0
> 如果计数器中的值为0
    设置了EFD_NONBLOCK标志位就直接返回-1
    没有设置EFD_NONBLOCK标志位就会一直阻塞直到计数器中的值大于0

(2) write: 向计数器中写入值
> 如果写入值的和小于0xFFFFFFFFFFFFFFFE，则写入成功
> 如果写入值的和大于0xFFFFFFFFFFFFFFFE
    > 设置了EFD_NONBLOCK标志位就直接返回-1
    > 如果没有设置EFD_NONBLOCK标志位，则会一直阻塞直到read操作执行

(3) close: 关闭文件描述符
```

### timefd
```c++
#include <sys/timerfd.h>
// timerfd是Linux为用户程序提供的一个定时器接口
// 这个接口基于文件描述符，通过文件描述符的可读事件进行超时通知，所以能够被用于select/poll的应用场景

(1) create
int timerfd_create(int clockid, int flags);
// 函数创建一个定时器对象，同时返回一个与之关联的文件描述符
// clockid: clockid标识指定的时钟计数器，可选值（CLOCK_REALTIME、CLOCK_MONOTONIC。。。）
//          CLOCK_REALTIME:     系统实时时间,随系统实时时间改变而改变
//          CLOCK_MONOTONIC:    从系统启动这一刻起开始计时,不受系统时间被用户改变的影响
// flags:   TFD_NONBLOCK/TFD_CLOEXEC


(2) write
struct timespec {
    time_t tv_sec; /* Seconds */
    long   tv_nsec; /* Nanoseconds */
};
struct itimerspec {
    struct timespec it_interval;  /* Interval for periodic timer （定时间隔周期）*/
    struct timespec it_value; /* Initial expiration (第一次超时时间)*/
};
int timerfd_settime(int fd, int flags, const struct itimerspec *new_value, struct itimerspec *old_value);
// 此函数用于设置新的超时时间，并开始计时,能够启动和停止定时器
// fd:          参数fd是timerfd_create函数返回的文件句柄
// flags:       参数flags 0,1
//              1:  代表设置的是绝对时间
//              0:  代表相对时间
// new_value:   指定定时器的超时时间以及超时间隔时间
// old_value:   如果old_value不为NULL, old_vlaue返回之前定时器设置的超时时间，具体参考timerfd_gettime()函数
// it_interval不为0则表示是周期性定时器，it_value和it_interval都为0表示停止定时器


(3) read
int timerfd_gettime(int fd, struct itimerspec *curr_value);
// 函数获取距离下次超时剩余的时间

uint64_t exp;
read(fd, &exp, sizeof(uint64_t)); 
// 当定时器超时，read读事件发生即可读，返回超时次数（从上次调用timerfd_settime()启动开始或上次read成功读取开始），它是一个8字节的unit64_t类型整数
// 如果定时器没有发生超时事件，则read将阻塞若timerfd为阻塞模式，否则返回EAGAIN 错误（O_NONBLOCK模式），如果read时提供的缓冲区小于8字节将以EINVAL错误返回。
```

### __thread
```c++
// __thread是GCC内置的线程局部存储设施，存取效率可以和全局变量相比
// __thread变量每一个线程有一份独立实体，各个线程的值互不干扰

// __thread使用规则
//      只能修饰POD类型，不能修饰class类型
//      __thread变量只能初始化为编译器常量

__thread int i = 1;

auto td = std::thread([](){
    i++;
    std::lock_guard<std::mutex> lg(mt);
    cout << std::this_thread::get_id() << endl;
    cout << i << endl;  // 2
});
auto td1 = std::thread([](){
    i++;
    std::lock_guard<std::mutex> lg(mt);
    cout << std::this_thread::get_id() << endl;
    cout << i << endl;  // 2
});
```

### prctl
```c++
#include <sys/prctl.h>

// prctl是进程制定而设计的的各种控制，例如设置/获得进程名
int prctl(int option, unsigned long arg2, unsigned long arg3, unsigned long arg4, unsigned long arg5);
```

### mmap madvise
```
mmap的作用是将硬盘文件的内容映射到内存中，采用闭链哈希建立的索引文件非常适合利用mmap的方式进行内存映射，利用mmap返回的地址指针就是索引文件在内存中的首地址，这样我们就可以放心大胆的访问这些内容了

使用过mmap映射文件的同学会发现一个问题，search程序访问对应的内存映射时，处理query的时间会有latecny会陡升，究其原因是因为mmap只是建立了一个逻辑地址，linux的内存分配测试都是采用延迟分配的形式，也就是只有你真正去访问时采用分配物理内存页，并与逻辑地址建立映射，这也就是我们常说的缺页中断

缺页中断分为两类，一种是内存缺页中断，这种的代表是malloc，利用malloc分配的内存只有在程序访问到得时候，内存才会分配

另外就是硬盘缺页中断，这种中断的代表就是mmap，利用mmap映射后的只是逻辑地址，当我们的程序访问时，内核会将硬盘中的文件内容读进物理内存页中，这里我们就会明白为什么mmap之后，访问内存中的数据延时会陡增

出现问题解决问题，上述情况出现的原因本质上是mmap映射文件之后，实际并没有加载到内存中，要解决这个文件，需要我们进行索引的预加载，这里就会引出本文讲到的另一函数madvise，这个函数会传入一个地址指针，已经一个区间长度，madvise会向内核提供一个针对于于地址区间的I/O的建议，内核可能会采纳这个建议，会做一些预读的操作。例如MADV_SEQUENTIAL这个就表明顺序预读

如果感觉这样还不给力，可以采用read操作，从mmap文件的首地址开始到最终位置，顺序的读取一遍，这样可以完全保证mmap后的数据全部load到内存中
```

### ftok
```c++
// 共享内存，消息队列，信号量它们三个都是找一个中间介质，来进行通信的
// ftok 出场了

key_t ftok(const char *pathname, int proj_id);
// fname是指定的文件名，这个文件必须是存在的而且可以访问的
// id是子序号，它是一个8bit的整数。即范围是0~255
// 当函数执行成功，则会返回key_t键值（文件信息与proj_id合成的值），否则返回-1
// 在一般的UNIX中，通常是将文件的索引节点取出，然后在前面加上子序号就得到key_t的值
// 所以相同文件名得到的 key_t 不一定一样，因为文件可能被删除再创建，文件索引不一样了
```

### cputime walltime
```c++
// https://levelup.gitconnected.com/8-ways-to-measure-execution-time-in-c-c-48634458d0f9
// cputime cpu时间
// walltime 墙上时间(运行时间)

(1) time command
    // cputime walltime
    $ time ***
    // real 0m5.931s walltime
    // user 0m5.926s clocktime
    // sys 0m0.005s

(2) c++ <chrono>
    // walltime
    auto begin = std::chrono::high_resolution_clock::now();
    // do something...
    auto end = std::chrono::high_resolution_clock::now();
    auto elapsed = std::chrono::duration_cast<std::chrono::nanoseconds>(end - begin);
    
(3) <sys/time.h> gettimeofday()
    // Walltime
    struct timeval begin, end;
    gettimeofday(&begin, 0);
    // do something...
    gettimeofday(&end, 0);
    long seconds = end.tv_sec - begin.tv_sec;
    long microseconds = end.tv_usec - begin.tv_usec;
    double elapsed = seconds + microseconds*1e-6;
    printf("Time measured: %.3f seconds.\n", elapsed);

(4)  <time.h> time()
    // walltime only measure second!!!
    time_t begin, end;
    time(&begin);
    // do something...    
    time(&end);
    time_t elapsed = end - begin;
    printf("Result: %.20f\n", sum);

(5) <time.h> clock()
    // clocktime
    clock_t start = clock();
    // do something...
    clock_t end = clock();
    double elapsed = double(end - start)/CLOCKS_PER_SEC;
    printf("Time measured: %.3f seconds.\n", elapsed);
```


### 当前工作路径与程序路径
```c++
// /proc/self/ 它代表当前程序运行环境

// 工作路径
// ls -l /proc/self/cwd

// 程序路径
// ls -l /proc/self/exe

// c语言中可以通过此方法获得工作与程序路径
#include <unistd.h>
int readlink(const char * path, char * buf, size_t bufsiz);
char cwdAbsPath[1024];
readlink("/proc/selef/cwd", cwdAbsPath, 1024);
readlink("/proc/selef/exe", cwdAbsPath, 1024);
```

### fwrite_unlocked
```c++
#include <stdio.h>
// fwrite_unlocked 是 fwrite 的非线程安全版本

// fwrite 的性能比 fwrite_unlocked 低

// 用fwrite_unlocked要自己保证其原子性

```