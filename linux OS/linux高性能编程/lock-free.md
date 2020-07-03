
### 无锁编程(lock-free)
```
https://zhuanlan.zhihu.com/p/55178835 [说说无锁(Lock-Free)编程那些事]
https://www.xiaobotalk.com/2020/06/atomic-multithreading/ [原子操作-无锁多线程编程]
http://ifeve.com/lock-based-vs-lock-free-concurren/ [基于锁的并发算法 vs 无锁的并发算法]
http://ifeve.com/lock-free-vs-wait-free-concurrency/ [无锁并发和无等待并发的对比分析]
https://www.cnblogs.com/mylinuxer/articles/5159467.html [无锁编程本质论]
http://ifeve.com/lock-free-and-wait-free/ [无锁和无等待的定义和例子]
https://www.ibm.com/developerworks/cn/linux/l-cn-lockfree/ [透过 Linux 内核看无锁编程]
https://zhuanlan.zhihu.com/p/53012280 [简化概念下的 lock-free 编程]
https://www.zhihu.com/question/295904223 [wait-free是指什么]
https://www.cnblogs.com/gaochundong/p/lock_free_programming.html [Lock-Free 编程]
http://novoland.github.io/%E5%B9%B6%E5%8F%91/2014/07/26/Lock-Free%20%E7%AE%97%E6%B3%95.html [Lock-Free 算法]

1、使用系统锁、同步原语的问题
    (1) 一些同步原语操作内核对象，会带来上下文切换影响效率
    (2) 锁的粒度比较大，加锁释放锁的消耗较大
    (3) 持有锁的优先级低线程可能被系统中断进入睡眠，这样其他等待锁的线程就会一直等下去
    (4) 持有锁的线程挂掉后得不到释放，导致死锁
    
2、非阻塞同步(Non-blocking Synchronization)
    在并行程序中保护共享数据用到'同步'，同步分为'阻塞同步'和'非阻塞同步'
    阻塞同步会造成死锁、活锁、性能低下等问题，为了降低风险和提高效率，现在大多使用'非阻塞同步'
    非阻塞同步的方案 无干扰(obstruction-free)、无锁(lock-free)、无等待(wait-free)
    (1) obstruction-free
    (2) lock-free
        能够确保执行它的所有线程中'至少有'在有限步内完成了任务
        当很多线程运行的时候，尽管有一些的线程竞争原语失败处于'饥饿'状态。虽然这些'倒霉蛋'线程可能一时还轮不到它运行，但它也不会阻碍别的线程，但是这些'饿'的线程最终都会运行下去的，最终在有限步内完成了任务

    (3) wait-free
        任意线程的任何操作'都可以'在有限步之内结束，而不用关心其它线程的执行速度
        承接lock-free，每个线程都不会出现'饥饿'状态，都会有条不紊的在有限步之内结束

    性能：wait > lock > obstruction
    实现难度：wait > lock


3、lock-free
    (1) 细粒度锁？
        细粒度锁是利用原子指令和内存屏障来实现多线程对共享内存的读写，其粒度之细可以与无锁媲美
        但是lock-free并不是细粒度锁，它更多的是无锁数据结构，将共享数据放入无锁的数据结构中，采用原子修改的方式来访问共享数据
        注意：对共享资源的安全访问，在不使用锁、同步原语的情况下，只能'依赖'于硬件支持的'原子操作'
        常见的数据结构有 无锁stack、无锁queue、无锁hashmap等

    (1) 优点
        > 从其定义来看，它们是 wait-free 的，可以确保线程永远不会阻塞
        > 状态转变是原子性的，以至于在任何点失败都不会恶化数据结构
        > 因为线程永远不会阻塞，所以当同步的细粒度是单一原子写或比较交换时，它们通常可以带来更高的吞吐量
        > lock-free 算法会有更少的同步写操作，因此纯粹从性能来看，它可能更便宜。
    (2) 缺点
        > 乐观的并发使用会对 hot data structures 导致 livelock
        > 代码需要大量困难的测试。通常其正确性取决于对目标机器内存模型的正确解释
        > lock-free 代码很难编写和维护
```

### lockfree vs spinlock
```
https://lumian2015.github.io/lockFreeProgramming/lock-free-vs-spin-lock.html

lockfree和spinlock都是依赖于原子操作，乍一看它们相似度很高，让我们看看它们的区别

// lockfree
int count;
void threadFunc(void) {
    int val;
    for (int i = 0; i < 1000000; i++) {
        do {
            val = count;
        } while (!atmoic_compare_exchange_weak(&count, &val, val + 1))
    }
    return;
}

// spinlock
int count;
int lock = 0;
void threadFunc(void) {
    int expected = 0;
    for (int i = 0; i < 1000000; i++) {
        while (!atomic_compare_exchange_weak(&lock, &expected, 1))  // spin_lock
            expected = 0;
        count++;
        lock = 0;   // spin_unlock
    }
}

// 区别
lockfree性能比spinlock要高一点

// 为什么
4个线程运行上述函数，观察竞争情况
lockfree版本
    thread1 val = count                 // count initial 0
    thread2 val = count                 // count = 0
            CAS(&count, &val, val + 1)  // success! count = 1
    thread3 val = count                 // count = 1
            CAS(&count, &val, val + 1)  // success! count = 2
    thread4 val = count                 // count = 1
            CAS(&count, &val, val + 1)  // error! now count is 2, and do while again
spinlock版本
    thread1 CAS(&lock, &expected, 1)    // get lock and set 1
            count++
    thread2 CAS(&lock, &expected, 1)    // error! thread1 hold the lock, please busy-waiting
    thread3 CAS(&lock, &expected, 1)    // error! thread1 hold the lock, please busy-waiting
    thread4 CAS(&lock, &expected, 1)    // error! thread1 hold the lock, please busy-waiting

// 总结
lockfree
    无论其他线程是否访问共享对象，每个线程都可以访问共享对象并对程序进行处理(尽管有时可能会有一些冲突)
spinlock
    如果一个线程正在访问共享对象，其他线程将被阻塞，无法访问共享对象
```

### lock-free的简单实现
```
1、RMW原子操作
    一般情况下，实现一个lock-free算法需要系统提供一个 atomic RMW (read-modify-write) 操作
    常用RMW操作包括test-and-set，fetch-and-add，compare-and-swap(CAS)以及更进一步的LL/SC

    在C++11的原子lib中，主要有以下RMW操作
    std::atomic<>::fetch_add()
    std::atomic<>::fetch_sub()
    std::atomic<>::fetch_and()
    std::atomic<>::fetch_or()
    std::atomic<>::fetch_xor()
    std::atomic<>::exchange()
    std::atomic<>::compare_exchange_strong()
    std::atomic<>::compare_exchange_weak()

    // CAS
    bool CAS(T *p, T old, T new) {
        if (*p != old) {
            return false;
        }
        *p = new;
        return true;
    }
    c++中使用compare_exchange_weak完成

    大多 lock-free 实现都是使用 CAS loop来实现
    在用CAS更新时，变量没有被修改过，依然保存着旧值
    如果CAS失败了，则说明有其他线程并发地在修改，此时线程不阻塞，而是不停地重试直到成功
    
    很多时候也被称为乐观锁，因为在访问资源时'乐观地'假设没有并发问题，不加锁就直接拿来用，在最后真正更新的时候再判断冲突

2、lock-free stack的简单实现
    struct node_t {
        std::atomic<node_t*> next;
        T data;
    };

    struct stack_t {
        std::atomic<node_t*> head;
    };

    void push(stack_t *stack, node_t *node) {
        node_t *head = stack->head.load(std::memory_order_relaxed);
        node->next.store(head, std::memory_order_relaxed);
        while(!stack->head.compare_exchange_weak(head, node, std::memory_order_release));
    }

    node_t* pop(stack_t *stack) {
        node_t *head = stack->head.load(std::memory_order_consume);
        for(;;){
            if (head = NULL) break;
            node_t *next = head->next.load(std::memory_order_relaxed);  // (1)
            if (stack->head.compare_exchange_weak(head, next, std::memory_order_release)) // (2)
                break;
        }
        return head;
    }

3、ABA问题
    上述的lock-free stack中，很有可能在(1)和(2)之间其它线程插入执行，会引发ABA问题
    (1) 情景再现
        > 例如当前栈为 A->B->C->D
        > thread1执行pop()函数的步骤(1)后被系统中断，此时head=A, next=B
        > thread2执行delNode = pop()后，又执行pop()一遍，然后执行push(delNode)，此时head=A, next=C
        > 此时thread1被调度恢复执行步骤(2)，由于head依然是A，所以可以执行成功，但是B已经不存在了，整个栈被破坏
    (2) 解决方法
        1) 思路1
            不要重用容器中的元素，本例中，Pop出来的A不要直接Push进容器，应该new一个新的元素
            但是！！！当地址A被回收后，new的新地址很有可能是之前A的地址，所以此方法不可行
        2) 思路2
            允许内存重用，对指向的内存采用标签指针(Tagged Pointers)的方式，标签作为一个版本号，随着标签指针上的每一次CAS运算而增加，并且只增不减

```

### lock-free 库
```
boost::lock-free 里实现了queue和stack
```