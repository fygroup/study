

### 进程终结
```
释放所占的资源，并告知其父进程
(1)删除进程描述符
在调用do_exit后，尽管线程已被僵死不再运行，但是系统还保留以确保父进程获取他的信息。进程终结时所需的清理和描述符删除分开执行，当父进程获取子进程的信息后，子进程task_struct才被释放
(2)孤儿进程
父进程在子进程前退出，由init进程当他父亲
```

### 进程调度
1、概念
```
> io消耗性和cpu消耗性进程

> 优先级
    nice：      -20 到 +19，越大优先级越低
    实时优先级： 0 到 99，越大优先级越高
    任何实时优先级都高于普通进程，实时优先级和nice优先级处于不相交范畴

> 时间片    
    每次调度时，把CPU分配给队首进程，并令其执行一个时间片。时间片的大小从几ms到几百ms。当执行的时间片用完时，由一个计时器发出时钟中断请求，调度程序便据此信号来停止该进程的执行，并将它送往就绪队列的末尾;然后，再把处理机分配给就绪队列中新的队首进程，同时也让它执行一个时间片。

> 调度器类
    > linux调度器以模块的方式提供，允许不同进程选择针对性的调度算法
    > 不同的调度类必须提供struct sched_class的一个实例
        extern const struct sched_class stop_sched_class;
        extern const struct sched_class dl_sched_class;
        extern const struct sched_class rt_sched_class;
        extern const struct sched_class fair_sched_class;
        extern const struct sched_class idle_sched_class;
    > 进程结构体中包含sched_class
        struct task_struct{
            ...
            const struct sched_class *sched_class;
            ...
        }
    > sched_class
        sched_class可以理解为调度器的接口类
        struct sched_class {
            /*  系统中多个调度类, 按照其调度的优先级排成一个链表
            下一优先级的调度类
            * 调度类优先级顺序: stop_sched_class -> dl_sched_class -> rt_sched_class -> fair_sched_class -> idle_sched_class
            */
            const struct sched_class *next;

            /*  将进程加入到运行队列中，即将调度实体（进程）放入红黑树中，并对 nr_running 变量加1   */
            void (*enqueue_task) (struct rq *rq, struct task_struct *p, int flags);
            
            /*  从运行队列中删除进程，并对 nr_running 变量中减1  */
            void (*dequeue_task) (struct rq *rq, struct task_struct *p, int flags);
            
            /*  放弃CPU，在 compat_yield sysctl 关闭的情况下，该函数实际上执行先出队后入队；在这种情况下，它将调度实体放在红黑树的最右端  */
            void (*yield_task) (struct rq *rq);

            /*   检查当前进程是否可被新进程抢占 */
            void (*check_preempt_curr) (struct rq *rq, struct task_struct *p, int flags);

            /*  选择下一个应该要运行的进程运行  */
            struct task_struct * (*pick_next_task) (struct rq *rq,
                                struct task_struct *prev);
            
            /* 将进程放回运行队列 */
            void (*put_prev_task) (struct rq *rq, struct task_struct *p);
            ....
        };

> 

```


(2) 调度
```
cpu执行红黑树中，已运行时间最小的
schedule调度器入口

do {
    preempt_disable();                                  /*  关闭内核抢占  */
    __schedule(false);                                  /*  完成调度  */
    sched_preempt_enable_no_resched();                  /*  开启内核抢占  */
} while (need_resched());   /*  如果该进程被其他进程设置了TIF_NEED_RESCHED标志，则函数重新执行进行调度    */
```

(3) need_resched
```
内核必须知道何时调用schedule(),如果依靠用户程序代码显示的调用schedule(),他们可能会永远执行下去。内核提供了一个need_resched标识来表明是否需要重新执行一次调度。内核必须瞅准时机、见缝插针地设置该字段。包括：
    在时钟中断服务程序中，当发现当前进程连续运行太长时间时
    当唤醒一个睡眠中的进程，发现被唤醒的进程比当前进程更有资格运行时
    一个进程通过系统调用改变调度政策（sched_setscheduler）或表示礼让（sched_yield）时。
```

(4) 睡眠（调度重点）
```
睡眠：将任务从红黑树移入等待队列，注意等待队列是个双向链表
唤醒：将任务从等待队列移入红黑树
1、如果schedule()是被一个状态为TASK_RUNNING 的进程调度，那么schedule()将调度另外一个进程占用CPU，当前进程会进入就绪状态，等待下次时间轮转调度
2、如果schedule()是被一个状态为TASK_INTERRUPTIBLE 或TASK_UNINTERRUPTIBLE 的进程调度，这将导致正在运行的进程进入睡眠，因为它已经不在运行状态中了，被移到了等待状态
3、转移过程（睡眠<-->唤醒）

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

(5) 唤醒(重点)
```
唤醒是通过wake_up()，唤醒指定等待队列的所有进程，所以在等待队列里的进程会收到唤醒信号，但有的是假唤醒，所以必须检查condition。

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
```
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
```

### 抢占
```
https://cloud.tencent.com/developer/article/1432987
(1)用户抢占：
    当内核即将返回用户空间时，内核会检查need_resched是否设置，如果设置，则调用schedule()，此时，发生用户抢占。一般来说，用户抢占发生几下情况：
    (1)从系统调用返回用户空间；
    (2)从中断(异常)处理程序返回用户空间。
(2)内核抢占：
    一个在内核态运行的进程，可能在执行内核函数期间被另一个进程取代。
    1) 内核抢占发生的时机
        > 当从中断处理程序正在执行，且返回内核空间之前。
        > 当内核代码再一次具有可抢占性的时候，如解锁（spin_unlock_bh）及使能软中断(local_bh_enable)等。
        > 如果内核中的任务显式的调用schedule()。
        > 如果内核中的任务阻塞(这同样也会导致调用schedule())。
        > preempt_count为0
    2) 内核不能被抢占的情况
        > 内核正进行中断处理
            在Linux内核中进程不能抢占中断(中断只能被其他中断中止、抢占，进程不能中止、抢占中断)，在中断例程中不允许进行进程调度。进程调度函数schedule()会对此作出判断，如果是在中断中调用，会打印出错信息。
        > 内核正在进行中断上下文的Bottom Half(中断的底半部)处理
            硬件中断返回前会执行软中断，此时仍然处于中断上下文中。
        > 内核的代码段正持有spinlock自旋锁、writelock/readlock读写锁等锁，处干这些锁的保护状态中
            内核中的这些锁是为了在SMP系统中短时间内保证不同CPU上运行的进程并发执行的正确性。当持有这些锁时，内核不应该被抢占。
        > 内核正在执行调度程序Scheduler
            抢占的原因就是为了进行新的调度，没有理由将调度程序抢占掉再运行调度程序。
        > 内核正在对每个CPU“私有”的数据结构操作(Per-CPU date structures)
            在SMP中，对于per-CPU数据结构未用spinlocks保护，因为这些数据结构隐含地被保护了(不同的CPU有不一样的per-CPU数据，其他CPU上运行的进程不会用到另一个CPU的per-CPU数据)。但是如果允许抢占，但一个进程被抢占后重新调度，有可能调度到其他的CPU上去，这时定义的Per-CPU变量就会有问题，这时应禁抢占。
    
```