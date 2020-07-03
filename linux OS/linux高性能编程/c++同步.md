### std::atomic
```
https://zhuanlan.zhihu.com/p/45566448

多线程之间的同步 = 原子操作 + 内存屏障

c++ atomic提供了原子操作和不同强度的Memory Model来对共享变量的控制

std::atomic提供了4种 memory ordering: Relaxed, Release-Acquire, Release-Consume, Sequentially-consistent

1、memory_order_seq_cst
    默认的选项，这个选项不允许reorder，那么也会带来一些性能损失
    相当于WO

2、memory_order_acq_rel
    一般不用

2、memory_order_release/consume/acquire
    (1) memory_order_release
        之前的读写不能往后乱序
        对使用memory_order_acquire/memory_order_consume同步的线程可见
        类似于unlock

    (2) memory_order_consume
        依赖于该读操作的后续读写，不能往前乱序
        另一个线程上memory_order_release之前的相关写序列，在memory_order_consume同步之后对当前线程可见
        类似于lock

    (3) memory_order_acquire
        之后的读写不能往前乱序
        另一个线程上memory_order_release之前的写序列，在memory_order_acquire同步之后对当前线程可见
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
            while (c.load(memory_order_consume) != 3);
            assert(a == 1); // assert 可能失败也可能不失败
        }


```

### 基于atomic实现锁
```
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
                    if (!lock_.exchange(std::memory_order_acquire)) break;
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
                while (rlock_.exchange(0) > 0);
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

