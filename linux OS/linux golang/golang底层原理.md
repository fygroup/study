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

### go 调度器

### 缓存池

### GC

### 并发调度

### goroute和channel

### go 系统调用和阻塞处理

### sysmon协程