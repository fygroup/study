### std::atomic
```
https://www.cnblogs.com/haippy/p/3301408.html
https://www.jianshu.com/p/8c1bb012d5f8

以下c++ 11

1、类模板
    template< class T > struct atomic;
    template<>  struct atomic<Integral>;
    template<>  struct atomic<bool>;
    template< class T >  struct atomic<T*>;

2、构造函数
    atomic() = default;
    constexpr atomic(T);
    atomic(const atomic &) = delete;    // 禁止拷贝构造

3、operator=
    atomic & operator=(const atomic &) = delete; // 禁止复制赋值
    T operator=(T); // 类似于store()

4、is_lock_free
    bool is_lock_free() const noexcept;
    判断该 std::atomic 对象是否具备 lock-free 的特性
    如果某个对象满足 lock-free 特性，在多个线程访问该对象时不会导致线程阻塞

5、store
    void store(T val, memory_order sync = memory_order_seq_cst) noexcept;

6、load
    T load(memory_order sync = memory_order_seq_cst) const noexcept;

7、exchange
    T exchange (T val, memory_order sync = memory_order_seq_cst) noexcept;
    用val替换所包含的值，并返回它之前的值。整个操作是原子性的(一个原子的读-修改-写操作)

8、compare_exchange_weak、compare_exchange_strong
    bool compare_exchange_weak (T& expected, T val, memory_order sync = memory_order_seq_cst) noexcept;  
    bool compare_exchange_strong (T& expected, T val, memory_order sync = memory_order_seq_cst) noexcept;
    比较原子对象的包含值与预期的内容
    如果是真的，它会用val替换包含的值(比如存储)。
    如果是假的，它会用所包含的值替换预期,因此调用该函数之后，如果被该原子对象封装的值与参数 expected 所指定的值不相等，expected 中的内容就是原子对象的旧值

```

### std::atomic 内存模型
```
https://zhuanlan.zhihu.com/p/45566448

多线程之间的同步 = 原子操作 + 内存屏障

c++ atomic提供了原子操作和不同强度的Memory Model来对共享变量的控制

std::atomic提供了4种 memory ordering: Relaxed, Release-Acquire, Release-Consume, Sequentially-consistent

1、memory_order_seq_cst
    默认的选项
    如果是读取就是 acquire 语义，如果是写入就是 release 语义，如果是读取+写入就是 acquire-release 语义同时会对所有使用此 memory order 的原子操作进行同步，所有线程看到的内存操作的顺序都是一样的，就像单个线程在执行所有线程的指令一样
    
    相当于WO

2、memory_order_acq_rel
    一般不用

2、memory_order_release/consume/acquire
    (1) memory_order_release
        对"写入"施加 release 语义(store)，在代码中这条语句前面的所有读写操作都无法被重排到这个操作之后
        store 之前的读写操作无法被重排至 store 之后。即 load-store, store-store 不能被重排
        
        当前线程内的所有写操作，对于其他对这个原子变量进行 memory_order_acquire 的线程可见
        
        当前线程内的与这块内存有关的所有写操作，对于其他对这个原子变量进行 memory_order_consume 的线程可见
        之前的读写不能往后乱序
        
        类似于unlock

    (2) memory_order_consume
        对当前要读取的内存施加 release 语义(store)，在代码中这条语句后面所有与这块内存有关的读写操作都无法被重排到这个操作之前
        
        在这个原子变量上施加 release 语义的操作发生之后，consume 可以保证读到所有在 release 前发生的并且与这块内存有关的写入

        类似于lock

    (3) memory_order_acquire
        对读取施加 acquire 语义(load)，在代码中这条语句后面所有读写操作都无法重排到这个操作之前
        load 之后的读写操作无法被重排至 load 之前。即 load-load, load-store 不能被重排
        
        在这个原子变量上施加 release 语义的操作发生之后，acquire 可以保证读到所有在 release 前发生的写入

        类似于lock

        std::memory_order_acquire + std::memory_order_release：相当于RCpc

3、memory_order_relaxed
    仅保持变量自身读写的相对顺序
    std::atomic的load()和store()都要带上memory_order_relaxed参数
    Relaxed ordering 仅仅保证load()和store()是原子操作，除此之外，不提供任何跨线程的同步

    std::memory_order_relaxed + Read-Modify-Write：原子操作，保持变量自身读写的相对顺序

4、实例
    (1)
        c = 0;
        thread1 {
            a = 1;
            b.store(2, memory_order_relaxed);
            c.store(3, memory_order_release);
        }

        thread2 {
            while (c.load(memory_order_acquire) != 3);
            assert(b == 2 && a == 1);  // 一定成功
        }

    (2)
        a = 0; c = 0;
        thread1 {
            a = 1;
            c.store(3, memory_order_release);
        }

        thread2 {
            // 仅保证与c相关的读写，后面不能向前乱序
            while (c.load(memory_order_consume) != 3);  
            assert(a == 1); // assert 可能失败也可能不失败
        }


```

### 基于atomic实现锁
```c++
1、spin_lock
    (1) WO的实现方式
        struct Spinlock {
            void lock() {
                for (;;) {
                    while(lock_.load());
                    if (!lock_.exchange(true)) break;
                }
            }

            void unlock() {
                lock_.store(false);
            }

            std::atomic<bool> lock_ = {false};
        };

    (2) RCpc的实现方式
        struct Spinlock {
            void lock() {
                for (;;) {
                    while(lock_.load(std::memory_order_relaxed));
                    if (!lock_.exchange(true, std::memory_order_acquire)) break;
                }
            }

            void unlock() {
                lock_.store(false, std::memory_order_release);
            }

            std::atomic<bool> lock_ = {false};
        };

2、RWlock
    (1) WO实现方式
        struct RWlock {
            void rlock() {
                for (;;) {
                    while (wlock_.load());
                    rlock_.fetch_add(1);
                    if (!wlock_.load()) break;
                    rlock_.fetch_sub(1);
                }
            }

            void unrlock() {
                rlock_.fetch_sub(1);
            }

            void wlock() {
                while (wlock_.exchange(true));
                while (rlock_.load(0) > 0);
            }

            void unwlock() {
                wlock_.store(false);
            }

            std::atomic<int> rlock_ = { 0 };
            std::atomic<bool> wlock_ = {false};
        };

    (2) RCpc的实现方式
        struct RWlock {
            void rlock() {
                for (;;) {
                    while (wlock_.load(std::memory_order_acquire));
                    rlock_.fetch_add(1, std::memory_order_acquire);
                    if (!wlock_.load(std::memory_order_acquire)) break;
                    rlock_.fetch_sub(1, std::memory_order_acquire);
                }
            }

            void unrlock() {
                rlock_.fetch_sub(1, std::memory_order_release);
            }

            void wlock() {
                while (wlock_.exchange(true, std::memory_order_acquire));
                while (rlock_.load(std::memory_order_acquire) > 0);
            }

            void unwlock() {
                wlock_.store(false, std::memory_order_release);
            }

            std::atomic<int> rlock_ = { 0 };
            std::atomic<bool> wlock_ = {false};
        };

```

