[C++对象模型](https://www.cnblogs.com/skynet/p/3343726.html)
[图说C++对象模型](https://www.cnblogs.com/QG-whz/p/4909359.html)

### 老生常谈-多态
```
RTTI
虚函数、虚函数表、虚指针
动态联编、静态联编
```

### c++对象模型
```c++
// 在C++类中有两种成员数据：static、nonstatic；三种成员函数：static、nonstatic、virtual

class Base
{
public:
    Base(int);
    virtual ~Base(void);
    int getIBase() const;
    static int instanceCount();
    virtual void print() const;

public:
    int iBase;
    static int count;
};

// c++对象中，并对内存存取和空间进行了优化
// non static 数据成员被放置到对象内部
// static数据成员，static and nonstatic 函数成员均被放到对象之外

    Base对象                虚函数表(vtbl)
+-----------------+         type_info_Base
| vptr            |-------> +------+
| int Base::iBase |         | slot | --> Base::~Base()
+-----------------+         +------+
                            | slot | --> void Base::print() const
                            +------+

+----------------------------+
|   static int Base::count   |
+----------------------------+

+----------------------------+
| static int instanceCount() |
+----------------------------+

+----------------------------+
|      Base::Base()          |
| int Base::getIBase() const |
+----------------------------+

// 每一个class产生一堆指向虚函数的指针，放在表格之中。这个表格称之为虚函数表（virtual table，vtbl）
// 每一个对象被添加了一个指针，指向相关的虚函数表vtbl。通常这个指针被称为vptr
// vptr的设定和重置都由每一个class的构造函数，析构函数和拷贝赋值运算符自动完成
// 虚函数表地址的前面设置了一个指向type_info的指针，用于RTTI运行期多态识别

```

