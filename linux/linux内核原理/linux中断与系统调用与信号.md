### 中断（上文、硬中断）
#### 相关概念
```
(1)软中断实现系统调用，陷入内核
(2)让中断程序快速处理完，从而就有了中断上下文的概念
```

#### 注册中断
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

#### 释放中断
```
void free_irq(unsigned int irq, void* dev)
//如果是共享中断线，那么仅删除dev所对应的处理程序，否则，禁用该中断线

```

#### 中断处理程序
```
static irqreturn_t handler(int irq, void* dev) //返回值其实是一个int
返回IRQ_NONE：不是注册处理函数指定的产生源
返回IRQ_HANDLED：是注册处理函数指定的产生源
```

#### 共享中断处理程序
```
(1)flags必须是IRQF_SHARED
(2)当收到中断信号时，会依次调用注册的函数
(3)中断处理程序区分他的设备是否产生了中断，需要硬件和软件的支持
(4)注册条件：中断线当前未被注册或者被注册的中断都是IRQF_SHARED
```

#### 中断机制
```
https://www.cnblogs.com/wlei/articles/2490011.html

设备产生中断，并通过中断线将中断信号送往中断控制器，如果中断没有被屏蔽则会到达CPU的INTR引脚，CPU立即停止当前工作，根据获得中断向量号从IDT中找出门描述符，并执行相关中断程序。

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


### 中断（软中断、tasklet、工作队列）
```
#下半部机制   上下文  复杂度                          性能   顺序执行保障
软中断       中断   高(确保软中断的执行顺序和锁机制)    好    没有
tasklet     中断    中(提供接口使用软中断)            中    同类型不能(不会)同时执行
工作队列     进程   低(在进程上下文运行，类似用户程序)  差    没有(和进程上下文一样被调度)

软中断的分配时静态的(即在编译时定义)，而tasklet的分配和初始化能够在执行时进行。
软中断(即便是同一种类型的软中断)能够并发地运行在多个CPU上。
因此，软中断是可重入函数并且必须明白地使用自旋锁保护其数据结构。
tasklet不必操心这些问题。由于内核对tasklet的运行进行了更加严格的控制。
同样类型的tasklet总是被串行运行。
换句话说就是：不能在两个CPU上同一时候执行同样类型的tasklet。可是，类型不同的tasklet能够在几个CPU上并发执行。
tasklet的串行化使tasklet函数不必是可重入的

```
#### 软中断
1、相关概念
```
(1) 哪些处理放在下半部
    如果一个任务对时间十分敏感，将其放在上半部
    如果一个任务和硬件有关，将其放在上半部
    如果一个任务要保证不被其他中断打断，将其放在上半部
    其他所有任务，考虑放在下半部

(2) 流程
    注册软中断         触发软中断          执行软中断  
    open_softirq ---> raise_softirq ---> do_softirq ---> 是否有未执行的中断函数 ---> 结束本次中断
                                            ↑                          ↓  
                                            +-------执行相应中断函数<---+

(3) 


```
2、API
```

```

#### 重要概念
```
软中断与tasklet是两个概念，软中断在编译期进行静态注册，tasklet可以通过代码进行动态注册。
```

#### 软中断
```
1、结构
struct softirq_action{
    void (*action)(struct softirq_action *);
}

static struct softirq_action softirq_vec[32];

2、核心代码
do_softirq() {
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
}


3、解释
软中断处理程序执行的时候，允许响应中断，但它自己不休眠，而且当前处理器的软中断被禁止。实际上，如果同一个软中断在被执行时再次触发了，其他处理器仍可执行其软中断。所以，不安全。。。
中断处理程序执行硬件设备操作，然后触发相应的软中断raise_softirq(NET_TX_SOFTIRQ)。内核在执行完中断程序后（硬中断），马上就调用do_softirq()
软中断不能被屏蔽，只能推后执行
软中断调度时期：
    do_irq完成I/O中断时调用irq_exit。
    系统使用I/O APIC,在处理完本地时钟中断时。
    local_bh_enable，即开启本地软中断时。
    SMP系统中，cpu处理完被CALL_FUNCTION_VECTOR处理器间中断所触发的函数时。
    ksoftirqd/n线程被唤醒时。 
```

### 硬中断与软中断
```
[区别](https://blog.csdn.net/xuchenhuics/article/details/79120644)

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
```
tasklet是通过软中断实现的，所以本身也是软中断。
```

### ksoftirqd内核线程
```
内核不会立即处理重新触发的软中断。当大量软中断出现的时候，内核会唤醒一组内核线程来处理。这些线程的优先级最低(nice值为19)，这能避免它们跟其它重要的任务抢夺资源。但它们最终肯定会被执行，所以这个折中的方案能够保证在软中断很多时用户程序不会因为得不到处理时间而处于饥饿状态，同时也保证过量的软中断最终会得到处理。每个处理器都有一个这样的线程，名字为ksoftirqd/n，n为处理器的编号。

for(;;){
    if (!softirq_pending(cpu)) //当没有软中断处理时，就调度正常队列
        schedule();
    set_current_state(TASK_RUNNING); //将当前进程设置为执行进程
    while (softirq_pending(cpu)){
        do_softirq();         处理软中断
        if (need_resched())
            schedule();
    }
    set_current_state(TASK_INTERRUPTIBLE);
}
```

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
(1) 参数验证
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

(2) 绑定一个系统调用
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

(3) 访问我们的系统函数
```
//比如open的系统调用定义
long open(const char* file, int flag, int mode)
//不依靠库的支持
#define __NR_open 5
__syscall3(long, open, const char*, filename, int, flag, int, mode)

```

### 禁止内核抢占和禁止中断
```
//禁止内核抢占
preempt_disable()
//允许内核抢占
preempt_enable()
```
