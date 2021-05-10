### 多基类访问者模式
```

```

### 线程安全的单例模式
```c++
// 传统的单例模式没有考虑线程安全，所以Instance的时候需要加锁

(1) 普通模式
    Singleton* Singleton::Instance() {
        Lock();
        if (value_ == NULL) {
            value_ = new Singleton();
        }
        return value_;
    }
    // 此方案最大的弊端是每次访问都要加锁，浪费资源

(2) Double-Checked Locking Pattern // 双检查锁模式
    Singleton* Singleton::Instance() {
        if (value_ == NULL) {
            Lock();
            if (value_ == NULL) {
                value_ = new Singleton();
            }
        }
        return value_;
    }
    // 此方法看似没问题，但是问题是存在的，出现在 new Singleton()
    // new对象分为三部，由于存在指令重排，会发生先为对象申请内存，再构造对象
    // 可能会存在
    // 线程A先进入Instance函数，value_申请内存，然后线程挂起，但并未给value_分配对象
    // 线程B进入了Instance函数，发现_instance不为null，就直接return _instance。当调用成员时会发生错误(因为没有被构造，仅仅是分配内存)

(3) thread_once
    // 利用thread_one只执行一次的功能
    Singleton* Singleton::Instance() {
        thread_once(&tid, &Singleton::Init);
        assert(value_ != NULL);
        return value_;
    }
    void Singleton::Init() {
        value_ = new Singleton;
    }
    
```

### copyable uncopyable
```c++
class copyable
{
protected:
    copyable() = default;
    ~copyable() = default;
};

class uncopyable
{
public:
    uncopyable(const uncopyable &) = delete;
    void operator=(const uncopyable &) = delete;
protected:
    copyable() = default;
    ~copyable() = default;
};

class A : uncopyable {};

```