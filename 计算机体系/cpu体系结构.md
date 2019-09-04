> 指令集体系结构
> 逻辑设计和硬件控制语言HCL
> 顺序实现
> 流水线原理
> 流水线实现


# 指令集体系结构(ISA)
## 定义
```
(1) 概念
一个处理器支持的指令和指令的字节级编码就是这个处理器的ISA

ISA 在编译器编写者和处理器设计人之间提供了一个抽象概念层，编译器编写者只需要知道允许哪些指令，以及它们是如何编码的；而处理器设计者必须建造出这些指令的处理器。

可见部分包括：寄存器、存储器、条件码、PC（程序计数器）、程序状态。

(2) 作用
ISA在编译器编写者（CPU软件）和处理器设计人员（CPU硬件）之间提供了一个抽象层


```

### cpu缓存架构
```
http://www.wowotech.net/kernel_synchronization/memory-barrier.html

                            Main Memory
                                ↑
                                ↓
                        System Bus Interface 
                                ↑               
                                ↓               
            +--------------- L2 Cache <---------+    
            |                   ↑               |
            ↓                   ↓               |
    L1 Instruction Cache    L1 Data Cache   Write-Combining Buffers
            |                   ↑      ↑        ↑ 
            |                   |    Write Buffers(store buffer) 
            |                   ↓       ↑ 
            |               Load/Store Unit
            |                     ↑ 
            |                     ↓
            +-------------> Execution Units


// cache状态
    
    cache状态           动作
    invalid 
                        从内存读取数据
    shared
                        cpu对cache进行load and store
    当前cache/其他cache                    
    exclusive/invalid
                        对cache写操作
    modified
                        写进内存(modified状态和exclusive状态都是独占该cacheline, 但是modified状态下，cacheline的数据是dirty的，而exclusive状态下，cacheline中的数据和memory中的数据是一致的)


// store buffer
    每个cpu写操作不必等到cacheline被加载，而是直接写到store buffer中然后欢快的去干其他的活。在CPU n的cacheline把数据传递到其cache 0的cacheline之后，硬件将store buffer中的内容写入cacheline。


// Invalidate Queue

```

### cpu乱序执行、编译器重排
```
代码顺序并不是真正的执行顺序，只要有空间提高性能，CPU和编译器可以进行各种优化。

// cpu乱序执行
    在一个固定长度的执行队列中，寻找可以同时执行的指令。这个过程只需考虑指令间是否有依赖关系，不需要理解程序的意图

// 编译器重排
    比处理器的范围更大，能在很大范围内进行代码分析,从而做出更优的策略,充分利用处理器的乱序执行功能.

```

### 内存屏障
```
https://www.jianshu.com/p/64240319ed60
http://ifeve.com/linux-memory-barriers/
http://ifeve.com/memory-barriers-or-fences/

(1) 可见性与重排序
    1) 可见性
        当一个线程修改了线程共享变量的值，其它线程在使用前，能够得到最新的修改值。
        多核时代的分层缓存架构，导致这个问题
    2) 重排序
        编译器、cpu出于优化的目的，导致指令重新排序

(2) 什么是内存屏障
    一个系统中，CPU和其它硬件可以使用各种技巧来提高性能，包括内存操作的重排、延迟和合并；预取；推测执行分支以及各种类型的缓存。内存屏障是用来禁用或抑制这些技巧的，使代码稳健地控制多个CPU和(或)设备的交互
    1) 内存屏障的种类
        > write（或store）内存屏障
        > 数据依赖屏障
        > read（或load）内存屏障
        > 通用内存屏障

(3) 显示内核屏障
    Linux内核有多种不同的屏障，工作在不同的层上：
    编译器屏障
    CPU内存屏障
    MMIO write屏障
    1) 编译器屏障
        Linux内核有一个显式的编译器屏障函数，用于防止编译器将内存访问从屏障的一侧移动到另一侧
        barrier()
        编译屏障并不直接影响CPU，CPU依然可以按照它所希望的顺序进行重排序
    2) CPU内存屏障(解决了硬件层面的可见性与重排序问题)
        屏障类型
        > LoadLoad Barriers
            指令示例: Load1;LoadLoad;Load2	
            说明: 该屏障确保Load1数据的装载先于Load2及其后所有装载指令的的操作
        > StoreStore Barriers
            指令示例: Store1;StoreStore;Store2
            说明: 该屏障确保Store1立刻刷新数据到内存(使其对其他处理器可见)的操作先于Store2及其后所有存储指令的操作
        > LoadStore Barriers
            指令示例: Load1;LoadStore;Store2
            说明: 确保Load1的数据装载先于Store2及其后所有的存储指令刷新数据到内存的操作
        > StoreLoad Barriers
        	指令示例: Store1;StoreLoad;Load2	
            说明: 该屏障确保Store1立刻刷新数据到内存的操作先于Load2及其后所有装载装载指令的操作。它会使该屏障之前的所有内存访问指令(存储指令和访问指令)完成之后,才执行该屏障之后的内存访问指令
            注意：
                该屏障同时具备其他三个屏障的效果，因此也称之为全能屏障（mfence），是目前大多数处理器所支持的；但是相对其他屏障，该屏障的开销相对昂贵。然而，除了mfence，不同的CPU架构对内存屏障的实现方式与实现程度非常不一样

(4) 什么地方需要内存障碍？
    在正常操作下，一个单线程代码片段中内存操作重排序一般不会产生问题，仍然可以正常工作，即使是在一个SMP内核系统中也是如此。但是，下面四种场景下，重新排序可能会引发问题：
    > 多理器间的交互
    > 原子操作
    > 设备访问
    > 中断


```