### c++ template
```
https://downdemo.gitbook.io/cpp-templates-2ed/
https://zhuanlan.zhihu.com/p/109582141

```

### 模板元编程
```
C++ 模板是图灵完备，理论上说 C++ 模板可以执行任何计算任务，但实际上因为模板是编译期计算，其能力受到具体编译器实现的限制（如递归嵌套深度，C++11 要求至少 1024，C++98 要求至少 17）

用途：编译期数值计算、类型计算、代码计算（如循环展开）

C++ 模板是函数式编程（functional programming），函数调用不产生任何副作用
循环结构：用递归形式实现
条件判断：模板的特例化实现判断
模板的<>中的模板参数相当于函数调用的输入参数
模板中的 typedef 或 static const 或 enum 定义函数返回值（类型或数值）

```

### CRTP
```c++
// 奇异递归模板模式(CRTP)是C++模板编程时的一种惯用法，把派生类作为基类的模板参数

// https://stackoverflow.com/questions/11795915/crtp-traits-class-no-type-named
// https://stackoverflow.com/questions/652155/invalid-use-of-incomplete-type


// 代码样式
template <typename T>
class Base{
public:
    void doWhat(){
        (static_case<T*>(this))->doWhatSub();
    }
};
 // use the derived class itself as a template parameter of the base class
class Derived : public Base<Derived>{
    void doWhatSub();
};

// 示例
template<typename T>
class A{
public:
    template<typename V>
    void func(V a) {
        (static_cast<T*>(this))->func(a);
    }
};

template<typename T1, typename T2>
class B : public A<B<T1, T2>> {
public:
    void func(T1 a) {
        cout << "t1 a " << a << endl;
    }
    void func(T2 a) {
        cout << "t2 a " << a << endl;
    }
};

// CRTP积累无法traits子类的type
template <typename C> class B {
    typedef typename C::T T; // 编译失败
    T* p_;
};
class D : public B<D> {
    typedef int T;
};



```

### template两个示例
```c++
// 示例一

template<typename Obj>
class Reflect {
    template<typename T>
    class Store {
        static map<string, T> _map;
    };
};

template<typename Obj> 
template<typename T> 
typename Reflect<Obj>::template Store<T>::map<string, T> Reflect<Obj>::Store<T>::_map;

// 示例二
template<typename T> class A{};

template<template<typename, typename> class T, typename T1, template T2>
class A<T<T1, T2>> {

};

```