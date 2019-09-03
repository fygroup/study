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
http://ifeve.com/memory-barriers-or-fences/

(0) cpu缓存体系


(1) 可见性与重排序
    1) 可见性
        当一个线程修改了线程共享变量的值，其它线程在使用前，能够得到最新的修改值。
        多核时代的分层缓存架构，导致这个问题
    2) 重排序
        编译器、cpu出于优化的目的，导致指令重新排序


(2) volatile(解决编译器层面的可见性与重排序问题)
    > 概念
        对volatile变量的写入操作必须在对该变量的读操作之前执行
    > 一个标准
        volatile变量规则只是一种标准，要求编译器实现保证volatile变量的偏序语义。结合程序顺序规则、传递性，该偏序语义通常表现为两个作用：
        > 保持可见性
        > 禁用重排序（读操作禁止重排序之后的操作，写操作禁止重排序之前的操作）

(3) 内存屏障(解决了硬件层面的可见性与重排序问题)
    1) store和load
        Store：将处理器缓存的数据刷新到内存中。
        Load：将内存存储的数据拷贝到处理器的缓存中。
    2) 分类
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
    3) x86内存屏障
        x86架构并没有实现全部的内存屏障。
        > Store Barrier(相当于StoreStore Barriers)
            指令：sfence
            功能：
                强制所有在sfence指令之前的store指令，都在该sfence指令执行之前被执行，发送缓存失效信号，并把store buffer中的数据刷出到CPU的L1 Cache中；所有在sfence指令之后的store指令，都在该sfence指令执行之后被执行。即，禁止对sfence指令前后store指令的重排序跨越sfence指令，使所有Store Barrier之前发生的内存更新都是可见的。这里的"可见", 指修改值可见（内存可见性）且操作结果可见（禁用重排序）
        > Load Barrier(相当于LoadLoad Barriers)
            指令：lfence
            功能：
                强制所有在lfence指令之后的load指令，都在该lfence指令执行之后被执行，并且一直等到load buffer被该CPU读完才能执行之后的load指令（发现缓存失效后发起的刷入）。即，禁止对lfence指令前后load指令的重排序跨越lfence指令，配合Store Barrier，使所有Store Barrier之前发生的内存更新，对Load Barrier之后的load操作都是可见的。
        > Full Barrier(相当于StoreLoad Barriers)    
            指令：mfence
            功能：所有Full Barrier之前发生的操作，对所有Full Barrier之后的操作都是可见的




```