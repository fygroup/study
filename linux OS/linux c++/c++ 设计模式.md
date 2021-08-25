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

// thread_once 原理
pthread_once_t once_control = PTHREAD_ONCE_INIT;
int pthread_once(pthread_once_t * once_control,void (*init_routin)(void));
// 当线程执行开始执行pthread_once函数后，线程会首先检查控制变量：
// 如果控制变量的状态为“未进行初始化”，则开始执行pthread_once中指定的init_routin函数
// 如果当前控制变量的状态为“初始化已结束”，则直接返回
// 如果当前控制变量的状态为“正在进行初始化”，则线程将等待直到控制变量的状态变为“初始化已结束”

```

### Double-checked Locking
```c++
// 双重检查锁

// 为什么需要此模式
// 多线程频繁的锁竞争会带来性能损耗，可以利用Double-checked Locking模式来减少锁竞争和加锁载荷。目前此模式广泛应用于单例模式中

// Double-checked Locking 原型
if check() {
    lock() {
        if check() {
            // do something...
        }
    }
}

// Double-checked Locking 错误的使用
if (instance == NULL) {
    Lock(){
        if (instance == NULL) {
            instance = new Instance();  // 问题出在这
        }
    }
}
// new的过程分为三步：分配内存、初始化、返回指针。此过程可能被编译重排，导致步骤3先于1、2，进而出现错误
// 解决方案：使用内存屏障

// c++ singleton(double-checked locking)
#include <atomic>
#include <mutex>

class Singleton {
public:
    static Singleton* GetInstance();

private:
    Singleton() = default;
    static std::atomic<Singleton*> s_instance;
    static std::mutex s_mutex;
};

Singleton* Singleton::GetInstance() {
  Singleton* p = s_instance.load(std::memory_order_acquire); // 注意
  if (p == nullptr) { // 1st check
    std::lock_guard<std::mutex> lock(s_mutex);
    p = s_instance.load(std::memory_order_relaxed);
    if (p == nullptr) { // 2nd (double) check
      p = new Singleton();
      s_instance.store(p, std::memory_order_release); // 注意
    }
  }
  return p;
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

### pimpl
```c++
// pimpl(Pointer to Implementation)技术是通过一个私有的成员指针，将指针所指向的类的内部实现数据进行隐藏
// 它是"隐藏实现" "降低耦合性" "分离接口"的一个现代 C++ 技术，并有着"编译防火墙(compilation firewall)"之称

(1) 为什么要用
// 1) 减少依赖
//      其一减少原类不必要的头文件的依赖，加速编译；其二对Impl类进行修改，无需重新编译原类
// 2) 隐藏类的实现
//      私有成员完全可以隐藏在共有接口之外，给用户一个间接明了的使用接口，尤其适合闭源API设计

(2) demo
// my_class.hpp
class MyClass {
public:
    MyClass();
    ~MyClass();
    // 此模式由于存在unique_ptr，编译器会隐式删除复制、移动操作。如果需要，需自己手动构建。要考虑unique_ptr的复制
    // MyClass(const MyClass &);                复制构造
    // MyClass(MyClass &&);                     移动构造
    // MyClass & operator=(const MyClass &);    拷贝赋值
    // MyClass & operator=(MyClass &&);         移动赋值
private:
    class Impl;                             // Impl对象
    std::unique_ptr<Impl> pimpl_;           // Impl指针，推荐unique_ptr，表示该类独有不被共享，不能拷贝
};

// my_class.cpp
class MyClass::Impl {
public:
    Impl():a("aa"){}
    std::string a;
    // 在这里定义所有私有变量和方法(换句话说是my_class类的具体实现细节内容)
    // 现在可以改变实现，而依赖my_class.h的其他类无需重新编译...
};
MyClass::MyClass():pimpl_(std::make_unique<Impl>()){}
MyClass::~MyClass() = default;

```