## 从内核出发
#### cpu分支预测、流水线和条件转移
(1)流水线
指令从取值到真正执行的过程划分成多个小步骤(取指、译码、执行、访存、写回)，cpu真正开始执行指令序列时，一步压一步的执行，减少其等待时间。
```
每一步是一个时钟周期，如果级数越多，每个周期执行的就越多，性能就越好（注意！不是越多越好）
1->2->3
   1->2->3
      1->2->3
每个时钟周期都完成一条指令的性能      
```

(2)分支预测
如果猜对了，火车可以直接开往要去的方向
如果猜错了，火车要停下来，然后倒车，然后将车轨扳到正确的方向，然后火车重新开往正确的方向。
如果预测对了，那就不用停下来了
[分支预测](https://www.cnblogs.com/yangecnu/p/4196026.html)

(3)likely和unlikely
```
#define likely(x)  __builtin_expect(!!(x), 1)
#define unlikely(x)    __builtin_expect(!!(x), 0)
//上述源码中采用了内建函数__builtin_expect来进行定义,__builtin_expect函数用来引导gcc进行条件分支预测
```













## 进程管理
### 基础概念
(1)在linux中线程是特殊的进程
(2)虚拟处理器和虚拟内存

### 进程描述符及任务结构
内核把进程列表放到任务队列（task_list）的双向链表中。链表每一项是task_struct,称为进程描述符的结构
(1)每个任务都有thread_info,里面存放着task_struct指针,每个thread_info在他的内核栈尾端分配。
(2)内核通过唯一标识符PID来表示进程，他的最大值表示能运行都少个进程，通过偏移间接查找task_struct
(3)进程状态
```
TASK_RUNNING(运行)：正在执行或在队列中等待
TASK_INTERRUPTIBLE(可中断)：进程在睡眠或者被阻塞
TASK_UNINTERRUPTIBLE(不可中断)：不受干扰
__TASK_TRACED：被其他进程跟踪的进程
__TASK_STOPPED(停止)：进程没有运行也不能投入运行

              进程调度
创建---->就绪----------->执行---->终止
            <-----------                
           |  时间片完    |
           |             |
            <——阻塞 <—————  
```
(4)设置进程状态
set_task_state(task,state)
(5)进程上下文
陷入内核，对内核的访问必须通过特定接口
(6)进程家族树
所有进程都是pid为1的init进程的后代，
```
struct task_struct *real_parent; /* real parent process */  
struct task_struct *parent; /* recipient of SIGCHLD, wait4() reports */  
struct list_head children;  /* list of my children */  
struct list_head sibling;   /* linkage in my parent's children list */  
struct task_struct *group_leader;   /* threadgroup leader */  

//访问父进程
struct task_struct *my_parent = current->parent;
//遍历子进程
struct task_struct *task;
stuct list_head *list;
list_for_each(list, &current->children){
    task = list_entry(list,struct task_struct, sibling)
}

```

### 进程创建
(1)linux创建进程时采用写时拷贝页实现，共享只读空间
(2)clone系统调用实现fork
```
clone(SIGCHLD,0);
```
(3)fork的实现

### 线程(特殊的进程)
(1)创建
调用clone的时候需要传递参数，指明需要共享的资源
```
clone(CLONE_VM | CLONE_FS | CLONE_FILES | CLONE_SIGHAND, 0);
//CLONE_FILES 父子进程共享打开的文件
//CLONE_FS共享文件系统信息
//CLONE_VM共享地址空间
```
(2)内核线程
内核需要在后台执行一些操作，与普通用户线程的区别是没有独立的地址空间。可以被调度可以被抢占
```
struct task_struct* kthread_create(int (*threadfn)(void* data), void* data, const char namefmt[], ...)
//内核通过clone来实现，新的进程将运行threadfn，参数是data，进程被命名为namefmt。他不会主动运行，需要wake_up_process来唤醒

//可以调用kthread_run()来达到创建并运行
struct task_struct* kthread_run(int (*threadfn)(void* data), void* data, const char namefmt[], ...)
```

### 进程终结
释放所占的资源，并告知其父进程
(1)删除进程描述符
在调用do_exit后，尽管线程已被僵死不再运行，但是系统还保留以确保父进程获取他的信息。进程终结时所需的清理和描述符删除分开执行，当父进程获取子进程的信息后，子进程task_struct才被释放
(2)孤儿进程
父进程在子进程前退出，由init进程当他父亲




















## 进程调度
### 时间片
每次调度时，把CPU分配给队首进程，并令其执行一个时间片。时间片的大小从几ms到几百ms。当执行的时间片用完时，由一个计时器发出时钟中断请求，调度程序便据此信号来停止该进程的执行，并将它送往就绪队列的末尾;然后，再把处理机分配给就绪队列中新的队首进程，同时也让它执行一个时间片。

### 调度
cpu执行红黑树中，已运行时间最小的
schedule调度器入口
```
do {
    preempt_disable();                                  /*  关闭内核抢占  */
    __schedule(false);                                  /*  完成调度  */
    sched_preempt_enable_no_resched();                  /*  开启内核抢占  */
} while (need_resched());   /*  如果该进程被其他进程设置了TIF_NEED_RESCHED标志，则函数重新执行进行调度    */
```


### need_resched
内核必须知道何时调用schedule(),如果依靠用户程序代码显示的调用schedule(),他们可能会永远执行下去。内核提供了一个need_resched标识来表明是否需要重新执行一次调度。内核必须瞅准时机、见缝插针地设置该字段。包括：
    在时钟中断服务程序中，当发现当前进程连续运行太长时间时
    当唤醒一个睡眠中的进程，发现被唤醒的进程比当前进程更有资格运行时
    一个进程通过系统调用改变调度政策（sched_setscheduler）或表示礼让（sched_yield）时。

### 睡眠（调度重点）
睡眠：将任务从红黑树移入等待队列，注意等待队列是个双向链表
唤醒：将任务从等待队列移入红黑树
(1)如果schedule()是被一个状态为TASK_RUNNING 的进程调度，那么schedule()将调度另外一个进程占用CPU，当前进程会进入就绪状态，等待下次时间轮转调度
(2)如果schedule()是被一个状态为TASK_INTERRUPTIBLE 或TASK_UNINTERRUPTIBLE 的进程调度，这将导致正在运行的进程进入睡眠，因为它已经不在运行状态中了，被移到了等待状态
(3)转移过程（睡眠<-->唤醒）
```
//此时当有一个进程需要睡眠时
DEFINE_WAIT(wait);                                  //在当前进程创建一个等待队列项

add_wait_queue(q, &wait);                           //将等待队列项加入全局等待队列中，当然我们必须在其他地方撰写相关代码,在事件发生时,对等待队列执行wake_up()操作

while(!condition){                                  //循环判断条件是否满足（）

    prepare_to_wait(&q,&wait, TASK_INTERRUPTIBLE);  //将进程状态变为TASK_INTERRUPTIBLE

    if (signal_pending(current))                    //信号和等待事件都可以唤醒处于TASK_INTERRUPTIBLE状态的进程,信号唤醒该进程为伪唤醒;该进程被唤醒后,如果(!condition)结果为真,则说明该进程不是由等待事件唤醒
        if (condition)                              //信号唤醒后要判断条件是否为真
            break;
    schedule();             //当前进程进入睡眠，所以当被唤醒后，也从这部开始运行
}
finish_wait(&q, &wait);     //状态设置为TASK_RUNNING，然后移出等待队列

```

### 唤醒(重点)
唤醒是通过wake_up()，唤醒指定等待队列的所有进程，所以在等待队列里的进程会收到唤醒信号，但有的是假唤醒，所以必须检查condition。
```
struct __wait_queue_head {
      spinlock_t lock;                    /* 保护等待队列的原子锁 (自旋锁),在对task_list与操作的过程中，使用该锁实现对等待队列的互斥访问*/
      struct list_head task_list;          /* 等待队列,双向循环链表，存放等待的进程 */
};

/*__wait_queue，该结构是对一个等待任务的抽象。每个等待任务都会抽象成一个wait_queue，并且挂载到wait_queue_head上。该结构定义如下：*/

struct __wait_queue {
 unsigned int flags;
#define WQ_FLAG_EXCLUSIVE   0x01
 void *private;                       /* 通常指向当前任务控制块 */
 wait_queue_func_t func;             
 struct list_head task_list;              /* 挂入wait_queue_head的挂载点 */
};
typedef struct __wait_queue wait_queue_t;
/* 任务唤醒操作方法，该方法在内核中提供，通常为auto remove_wake_function */

//等待队列的结构
wait_queue_head --> wait_queue_t --> wait_queue_t --> wait_queue_t
//wake_up()与wake_event()或者wait_event_timeout成对使用，
//wake_up_intteruptible()与wait_event_intteruptible()或者wait_event_intteruptible_timeout()成对使用。
```



### 上下文切换
(1)用两种方法来激活调度:
    一种是直接的, 比如进程打算睡眠或出于其他原因放弃CPU(进程主动调用)
    另一种是通过周期性的机制, 以固定的频率运行, 不时的检测是否有必要
(2)在linux内核中，上下文的切换有两种方式：第一种是进程主动让出CPU，这样的操作成为“让步”。第二种是由内核调度程序决定进程运行时间，在在运行时间结束（如时间片耗尽）或者需要切换高优先级进程时强制挂起进程，这样的操作叫“抢占”
(3)抢占不是一个进程强制切换到另一进程。执行抢占的是内核，并不是进程。
(4)context_switch()处理抢占，里面包含两个函数：switch_mm()虚拟内存映射到新的进程，switch_to()切换到新进程的处理器状态。
(5)内核必须知道什么时候调用schedule()。每个进程都有一个need_resched标志，如果此标志被设置，那么会调用schedule,来启动一个新进程。
(6)可以主动设置一个进程的need_resched; 当一个优先级高的进程被唤醒时（进入可执行状态），也会设置这个标志。总之这个标志的意思就是有其他进程需要被运行，请内核赶紧调度。
(7)处理器总处于以下三种状态之一
    内核态，运行于进程上下文，内核代表进程运行于内核空间；
    内核态，运行于中断上下文，内核代表硬件运行于内核空间；
    用户态，运行于用户空间。


### 抢占
(1)用户抢占：
    当内核即将返回用户空间时，内核会检查need_resched是否设置，如果设置，则调用schedule()，此时，发生用户抢占。一般来说，用户抢占发生几下情况：
    (1)从系统调用返回用户空间；
    (2)从中断(异常)处理程序返回用户空间。
(2)内核抢占：一个在内核态运行的进程，可能在执行内核函数期间被另一个进程取代。
    1）当从中断处理程序正在执行，且返回内核空间之前。
    2）当内核代码再一次具有可抢占性的时候，如解锁（spin_unlock_bh）及使能软中断(local_bh_enable)等。
    3）如果内核中的任务显式的调用schedule()。
    4）如果内核中的任务阻塞(这同样也会导致调用schedule())。
    5) preempt_count为0，
















## 中断（上文、硬中断）
### 相关概念
(1)软中断实现系统调用，陷入内核
(2)让中断程序快速处理完，从而就有了中断上下文的概念

### 注册中断
```
#include <linux/interrupt.h>
//分配一条给定的中断线
int request_irq(unsigned int irq, irq_hander_t handler, unsigned long flags, const char* name, void* dev)
//第一个参数irq表示要分配的中断号，有些是预先设定（键盘鼠标），可以通过探测获取，可以通过编程动态确定
//第二个是实际的中断处理程序，typedef irqreturn_t (*irq_handler_t)(int, void*)
//第三个flags表示中断处理标志：
    //IRQF_DISABLED设置后表示禁止其他中断，很野蛮地行为
    //TRQF_SAMPLE_RANDOM表明这个设备对内核熵池有贡献
    //IRQF_TIMER表明为系统定时器中断处理而准备
    //IRQF_SHARED可以使多个中断程序共享一个中断线
//第四个表示中断相关设备的ASCII文本，比如pc键盘中断对应的'keyboard',以便与用户进行通信
//第五个dev，可以传递驱动设备结构，这个是唯一的，在共享中断线中，用于识别是哪个中断
//成功返回0
//注意！！！此函可能会睡眠，不能再中断上下文或其他不允许阻塞的代码中调用该函数

```

### 释放中断
```
void free_irq(unsigned int irq, void* dev)
//如果是共享中断线，那么仅删除dev所对应的处理程序，否则，禁用该中断线

```

### 中断处理程序
```
static irqreturn_t handler(int irq, void* dev) //返回值其实是一个int
返回IRQ_NONE：不是注册处理函数指定的产生源
返回IRQ_HANDLED：是注册处理函数指定的产生源
```

### 共享中断处理程序
(1)flags必须是IRQF_SHARED
(2)当收到中断信号时，会依次调用注册的函数
(3)中断处理程序区分他的设备是否产生了中断，需要硬件和软件的支持
(4)注册条件：中断线当前未被注册或者被注册的中断都是IRQF_SHARED

### 中断实例
```

```

### 中断机制
设备产生中断，并通过中断线将中断信号送往中断控制器，如果中断没有被屏蔽则会到达CPU的INTR引脚，CPU立即停止当前工作，根据获得中断向量号从IDT中找出门描述符，并执行相关中断程序。

```
//注册IQR
int request_irq(unsigned int irq, irq_handler_t handler, unsigned long flags, const char *name, void *dev);
//释放IQR
void free_irq(unsigned int, void *);
//注：IRQ线资源非常宝贵，我们在使用时必须先注册，不使用时必须释放IRQ资源
//激活当前CPU中断
local_irq_enable()
//禁止当前CPU中断
local_irq_disable()
//激活指定中断线
void enable_irq(unsigned int irq);
//禁止指定中断线
void disable_irq(unsigned int irq);
```

#### 中断状态
```
local_
```




## 中断（下文、tasklet、软中断）
### 重要概念
软中断与tasklet是两个概念，软中断在编译期进行静态注册，tasklet可以通过代码进行动态注册。

### 软中断
结构
```
struct softirq_action{
    void (*action)(struct softirq_action *);
}

static struct softirq_action softirq_vec[32];
```
核心代码do_softirq()
```
u32 pending;
pending = local_softirq_pending(); //待处理软中断的32位位图，如果第n位设置为1，那么第n位对应的软中断等待处理

if (pending){
    struct softirq_action *h;
    set_softirq_pending(0);     //重置位图
    h = softirq_vec;
    do{                         //循环32位
        if (h & 1)
            h->action(h);
        h++;
        pending >>= 1;
    }while(pending);
}
```
软中断处理程序执行的时候，允许响应中断，但它自己不休眠，而且当前处理器的软中断被禁止。实际上，如果同一个软中断在被执行时再次触发了，其他处理器仍可执行其软中断。所以，不安全。。。
中断处理程序执行硬件设备操作，然后触发相应的软中断raise_softirq(NET_TX_SOFTIRQ)。内核在执行完中断程序后（硬中断），马上就调用do_softirq()
软中断不能被屏蔽，只能推后执行
软中断调度时期
    do_irq完成I/O中断时调用irq_exit。
    系统使用I/O APIC,在处理完本地时钟中断时。
    local_bh_enable，即开启本地软中断时。
    SMP系统中，cpu处理完被CALL_FUNCTION_VECTOR处理器间中断所触发的函数时。
    ksoftirqd/n线程被唤醒时。 

### 硬中断与软中断
[区别](https://blog.csdn.net/xuchenhuics/article/details/79120644)
```
硬中断是外部设备对CPU的中断，软中断是中断底半部的一种处理机制，信号则是由内核（或其他进程）对某个进程的中断
硬中断的中断号是由中断控制器提供的,软中断的中断号由指令直接给出，无需使用中断控制器。
硬中断是可屏蔽的,软中断不可屏蔽
Linux下硬中断是可以嵌套的，但是没有优先级的概念，也就是说任何一个新的中断都可以打断正在执行的中断，但同种中断除外。
软中断不能嵌套，但相同类型的软中断可以在不同CPU上并行执行
硬中断的开关
    简单禁止和激活当前处理器上的本地中断：
    local_irq_disable();
    local_irq_enable();
    保存本地中断系统状态下的禁止和激活：
    unsigned long flags;
    local_irq_save(flags);
    local_irq_restore(flags);
同一处理器的中断不会抢占另一个软中断，唯一可以抢占软中断的是硬中断。
软中断可以被抢占，但是不会睡眠，所以不能在软中断中使用信号量和阻塞
```


### tasklet
tasklet是通过软中断实现的，所以本身也是软中断。
```
//声明



```

### ksoftirqd内核线程
内核不会立即处理重新触发的软中断。当大量软中断出现的时候，内核会唤醒一组内核线程来处理。这些线程的优先级最低(nice值为19)，这能避免它们跟其它重要的任务抢夺资源。但它们最终肯定会被执行，所以这个折中的方案能够保证在软中断很多时用户程序不会因为得不到处理时间而处于饥饿状态，同时也保证过量的软中断最终会得到处理。每个处理器都有一个这样的线程，名字为ksoftirqd/n，n为处理器的编号。
```
for(;;){
    if (!softirq_pending(cpu)) //当没有软中断处理时，就调度正常队列
        schedule();
    set_current_state(TASK_RUNNING);
    while (softirq_pending(cpu)){
        do_softirq();         处理软中断
        if (need_resched())
            schedule();
    }

    set_current_state(TASK_INTERRUPTIBLE);
}


```

### 工作队列
```
```




## 系统调用
### 与内核通信
```
c API ——> 内核 ——> 硬件

陷入内核
linux不会直接调用内核函数，而采用软中断，告诉内核需要执行系统调用，来切换到内核态

系统调用号是通过%eax传入内核的
比如调用fork
#define __NR_fork      2  
mov   eax, 2        
int   0x80     软中断

Linux最多允许向系统调用传递6个参数，分别依次由%ebx，%ecx，%edx，%esi，%edi这5个寄存器完成

系统调用比较费时，可以看看系统调用函数表。尽量避免多次调用里面的函数

```

### 系统调用的实现
(1)参数验证
指针指向的区域必须是用户空间，不能哄骗内核去读其内核空间
不能哄骗内核去读其他进程的空间
内存也分为可读、可写、可执行
```
//内核提供了两个方法
copy_to_user() //向用户空间写数据
//第一个参数时进程内存参数，第二个是内核空间源地址，最后一个是需要拷贝的长度

copy_from_user()  //用户向内核拷贝数据

//失败返回没能拷贝完的字节数
//注意上面两个函数也会发生休眠，当用户数据的页被换出到硬盘。。。
```
(2)绑定一个系统调用
```
//1、在sys_call_table中注册你的函数
//在entry.s文件中
ENTRY(sys_call_table)
    .long sys_restart_call()  /*0*/
    .long sys_exit()
    ...
    .long sys_read()  
    ...
    .long sys_foo()   //在末尾加入你的函数

//2、把系统调用号加入<asm/unistd.h>中
#define __NR_restart_call  0
#define __NR_exit          1
...
#define __NR_read          3
...
#define __NR_foo           338

//3、实现你的函数
#include <asm/page.h>
asmlinkage long sys_foo(void){
    .....
}

```

(3)访问我们的系统函数
```
//比如open的系统调用定义
long open(const char* file, int flag, int mode)
//不依靠库的支持
#define __NR_open 5
__syscall3(long, open, const char*, filename, int, flag, int, mode)

```















## 内核同步方法
### 原子操作


### 自旋锁与信号量
自旋锁最多只能被一个进程所有，其他进程在等待锁时自旋（特别浪费处理器时间）。
所以，自旋锁不能被长时间占用。
信号量则让请求线程进入睡眠，直到锁重新可用是再唤醒他，这就有两次明显的上下文切换。
自旋锁禁止内核抢占

### 自旋锁
```
//自旋锁可以用在中断程序中，但是，一定要在获取锁之前，首先禁止本地中断！！！
DEFINED_SPINLOCK(mr_lock);
unsigned long flags;
spin_lock_irqsave(&mr_lock,flags); //保存当前中断状态，并禁用本地中断，然后获取指定的锁，
spin_unlock_irqrestore(&mr_lock,flags);

//如果中断在加锁前是激活的，那么就无需保存当前状态,但是如果不能确保当前中断是不是处于激活状态，最好少用
spin_lock_irq(&mr_lock);
spin_unlock_irq(&mr_lock);

//自旋锁与下半部
下半部和进程上下文共享数据时，需要加锁
中断和下半部共享数据时，获得锁的同时还要禁止中断
同类的tasklet共享数据不需要保护，

```

### 读写锁(属于自旋锁)
```
DEFINE_RWLOCK(mr_rwlock);
read_lock(&mr_rwlock);
read_unlock(&mr_rwlock);

write_lock(&mr_rwlock);
write_unlock(&mr_rwlock);
```

### 信号量
只有进程上下文才能使用信号量，中断上下文不能使用，因为信号量会导致睡眠
信号量不同于自旋锁，它允许内核抢占，所以信号量不会对调度的等待时间带来负面影响
```

```











### 禁止内核抢占和禁止中断
```
//禁止内核抢占
preempt_disable()
//允许内核抢占
preempt_enable()

```