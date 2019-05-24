### 自旋锁与信号量（这个不是共享内存的信号量）
```
自旋锁最多只能被一个进程所有，其他进程在等待锁时自旋（特别浪费处理器时间）。
所以，自旋锁不能被长时间占用。
信号量则让请求线程进入睡眠，直到锁重新可用是再唤醒他，这就有两次明显的上下文切换。
自旋锁禁止内核抢占
```

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
```
只有进程上下文才能使用信号量，中断上下文不能使用，因为信号量会导致睡眠
信号量不同于自旋锁，它允许内核抢占，所以信号量不会对调度的等待时间带来负面影响
```
