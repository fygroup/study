### "独立的" GO
```
所有在 UNIX 系统上运行的程序最终都会通过 C 系统调用来和内核打交道。用其他语言编写程序进行系统调用，方法不外乎两个：一是自己封装，二是依赖 glibc、或者其他的运行库。Go 语言选择了前者，把系统调用都封装到了 syscall 包。封装时也同样得通过汇编实现

glibc最主要的功能就是对系统调用的封装，几乎其它任何的运行库都要依赖glibc




```

### CGO_LDFLAGS
```
gcc 编译选项

动态的:
export CGO_LDFLAGS="-Xlinker -rpath=/path/to/another_glibc/lib"

静止的:
export CGO_LDFLAGS="-Xlinker -rpath=/path/to/another_glibc/lib -static"

// 完全静态编译
export CGO_LDFLAGS=0

```

### 
```
P的数量是固定的，但是M的数量不是固定的，M >= P

M分离
    1) 系统调用（文件IO）
    2) cgo运行一段时间后（将cgo里的c程序交给原生系统，go不想管了）
    M的分离，P会创建一个新的系统线程，然后结合运行。当前一个G1可以移回 LRQ 并再次由P执行。如果这种情况需要再次发生，M1将被放在旁边以备将来使用
```


### runtime
```
https://zhuanlan.zhihu.com/p/95056679

Golang Runtime是go语言运行所需要的基础设施
(1) 协程调度、内存分配、GC
(2) 操作系统及cpu相关操作的封装(信号处理、系统调用、寄存器操作、原子操作等)，CGO
(3) pprof、trace、race检测的支持
(4) map、channel、string等内置类型及反射的实现

```

### 内存分配
```go
// https://zhuanlan.zhihu.com/p/29216091

// span: 由多个地址连续的页（page）组成的⼤块内存
// object: 将 span 按特定⼤⼩切分成多个⼩块，每个⼩块可存储⼀个对象

// 分配器按页数来区分不同⼤⼩的 span
// 以页数为单位将 span 存放到管理数组中，需要时就以页数为索引进⾏查找

_PageShift = 13
_PageSize = 1 << _PageShift // 8KB

type mspan struct {
    next *mspan         // 双向链表
    prev *mspan
    start pageID        // 起始序号(address >> _PageShift)
    npages uintptr      // 页数
    freelist gclinkptr  // 待分配的object链表
}

// page到span的映射
    // start   记录了起始 Page，也就是知道了从Span到Page的映射
    // npages  page数
    
    // 这种方式虽然简洁明了，但是在 Page 比较少的时候会有很大的空间浪费
// 用一个数组记录每个Page所属的 Span，而数组索引就是 Page ID

// 分配器由三种组件组成
cache: 每个运⾏期⼯作线程都会绑定⼀个 cache，⽤于⽆锁 object 分配。
central: 为所有 cache 提供切分好的后备 span 资源。
heap: 管理闲置 span，需要时向操作系统申请新内存。


_MaxSmallSize = 32 << 10 // 32KB
_NumSizeClasses = 67        // span页数类别

type mheap struct {
    free [_MaxMHeapList]mspan   // 页数在127以内的闲置span链表数组
    freelarge mspan             // 页数大于127的span链表
    // 每个 central 对应一个 sizeclass
    central [_NumSizeClasses]struct {
        mcentral mcentral
    }
}

type mcentral struct {
    sizeclass int32 // 规格 
    nonempty mspan  // 链表，尚有空间object的span
    empty mspan     // 链表，没有空闲object，或已被cache取走的span
}

type mcache struct {
    alloc [_NumSizeClasses]*mspan   // 以sizeclass为索引管理多个用于分配的span
}
```

### 阻塞
```
go语言的阻塞主要有4种场景

(1) 原子、互斥量或通道操作
    这些操作的调用导致 Goroutine 阻塞，调度器将把当前阻塞的 Goroutine 切换出去，重新调度 LRQ 上的其他 Goroutine

(2) 网络IO
    网络IO操作导致 Goroutine 阻塞
    Go程序提供了网络轮询器(NetPoller)来处理网络请求和 IO 操作的问题，其后台通过 kqueue(MacOS)，epoll(Linux)或 iocp(Windows)来实现 IO 多路复用
    通过使用 NetPoller 进行网络系统调用，调度器可以防止 Goroutine 在进行这些系统调用时阻塞 M。这可以让 M 执行 P 的 LRQ 中其他的 Goroutines，而不需要创建新的 M
    执行网络系统调用不需要额外的 M，网络轮询器使用系统线程，它时刻处理一个有效的事件循环，有助于减少操作系统上的调度负载
    用户层眼中看到的 Goroutine 中的"block socket"，实现了 goroutine-per-connection 简单的网络编程模式。实际上是通过 Go runtime 中的 netpoller 通过 Non-block socket + I/O 多路复用机制"模拟"出来的

(3) 文件IO等一些系统调用
    当调用一些系统方法的时候（如文件 I/O），如果系统方法调用的时候发生阻塞，这种情况下，网络轮询器(NetPoller)无法使用，而进行系统调用的 G1 将阻塞当前 M1。调度器引入 其它M 来服务 M1 的P

(4) sleep
    如果在 Goroutine 去执行一个 sleep 操作，导致 M 被阻塞了
    Go 程序后台有一个监控线程 sysmon，它监控那些长时间运行的 G 任务然后设置可以强占的标识符，别的 Goroutine 就可以抢先进来执行
```

### 同步与异步系统调用

### 缓存池

### GC

### goroute调度
```
M P G
LRQ(local run queue)
GRQ(global run queue)

```

### channel


### sysmon协程