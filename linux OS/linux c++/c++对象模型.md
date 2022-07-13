[C++对象模型](https://www.cnblogs.com/skynet/p/3343726.html)
[图说C++对象模型](https://www.cnblogs.com/QG-whz/p/4909359.html)
[c++对象布局](https://mp.weixin.qq.com/s/sufz7wxC_rwc1q3FXY-QMQ?utm_source=wechat_session&utm_medium=social&utm_oi=555400798499188736)

### 老生常谈-多态
```
RTTI
虚函数、虚函数表、虚指针
动态联编、静态联编
```

### 查看对象的布局
```
// 查看对象布局
clang++ -Xclang -fdump-record-layouts -stdlib=libc++ -c model.cc

// 查看虚函数表布局
clang++ -Xclang -fdump-vtable-layouts -stdlib=libc++ -c model.cc
```

### c++对象模型
```c++
// 在C++类中有两种成员数据：static、nonstatic；三种成员函数：static、nonstatic、virtual

class Base {
public:
    Base() = default;
    virtual ~Base(void) = default;
    void func();
    static void func1();
    virtual void print();

private:
    int a;
    char b[10]
    static int c;
};

// c++对象中，并对内存存取和空间进行了优化
// non static 数据成员被放置到对象内部
// static数据成员，static and nonstatic 函数成员均被放到对象之外

// 对象模型
*** Dumping AST Record Layout
         0 | class Base
         0 |   (Base vtable pointer)  // 虚函数指针
         8 |   int a
        12 |   char [10] b
           | [sizeof=24, dsize=22, align=8,
           |  nvsize=22, nvalign=8]

// 虚函数表布局
Vtable for 'Base' (5 entries).
   0 | offset_to_top (0)
   1 | Base RTTI
       -- (Base, 0) vtable address --
   2 | Base::~Base() [complete]
   3 | Base::~Base() [deleting]
   4 | void Base::print()

// offset_to_top(0) 
    // 表示当前这个虚函数表地址距离对象顶部地址的偏移量，因为对象头部就是虚函数表的指针，所以偏移量为0
// RTTI指针
    // 指向存储运行时类型信息(type_info)的地址，用于运行时类型识别，用于typeid和dynamic_cast
// 虚函数表指针
    // RTTI下面就是虚函数表指针真正指向的地址存储了类里面所有的虚函数，至于这里为什么会有两个析构函数，见下

    Base                 vtable Base
+----------+         +-------------------+
| vptr     |---+     |    offset 0       | 
+----------+   |     +-------------------+
| Base::a  |   |     |      RTTI         |
+----------+   |     +-------------------+
| Base::b  |   +---->| ~Base()[complete] |
+----------+         +-------------------+
                     | ~Base()[deleting] |
                     +-------------------+
                     |  Base::print()    |
                     +-------------------+
// 对象外放置的元素
+----------------------------+
| Base::Base()               |
| void Base::func()          |
| static void Base::func1()  |
| static int Base::c         |   
+----------------------------+

// 每一个class产生一个指向虚函数的指针，放在表格之中。这个表格称之为虚函数表（virtual table，vtbl）
// 每一个对象被添加了一个指针，指向相关的虚函数表vtbl。通常这个指针被称为vptr
// vptr的设定和重置都由每一个class的构造函数，析构函数和拷贝赋值运算符自动完成
// 虚函数表地址的前面设置了一个指向type_info的指针，用于RTTI运行期多态识别
// 虚函数表中有两个析构函数，一个标志为 deleting，一个标志为 complete，因为对象有两种构造方式，栈构造和堆构造，所以对应的实现上，对象也有两种析构方式，其中堆上对象的析构和栈上对象的析构不同之处在于，栈内存的析构不需要执行 delete 函数，会自动被回收
```

