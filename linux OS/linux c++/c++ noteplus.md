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

// CRTP基类无法traits子类的type
template <typename C> class B {
    typedef typename C::T T; // 编译失败
    T* p_;
};
class D : public B<D> {
    typedef int T;
};

```

### template前缀
```c++
// 在通过“.”,“->”,“::”限定的依赖名访问成员模板之前, template关键字必不可少

// (1) 为什么要用template修饰
    template<class T>
    int f(T& x) {
        return x.template convert<3>(pi);
    }
    // 如果没有template
    x.convert<3>(pi)
    // 可被理解成"小于"3 "大于"pi
    (x.convert < 3) > (pi)


// (2) 继承一个模板类的子模板类
    template<typename T> 
    class A {
    public:
        template<typename T1> class Base{};
    };

    template<typename T>
    class B : A<T>::template Base<T>{};


// (3) 模板类的模板子类的静态成员初始化
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


```

### template1个示例
```c++
// 示例
template<typename T> class A{};

template<template<typename, typename> class T, typename T1, template T2>
class A<T<T1, T2>> {

};

template<template<typename...> class T, typename... Argvs>
class A<T<Argvs...>> {
    
};

```

### template template parameter
```c++
// 模仿容器类
template<typename T, 
        template<typename, typename Alloc = std::allocator<T>> class ContainType>
class MyContainer {
public:
    ContainType<T> value;
};

MyContainer<int, vector> a;

// 2
template <typename T, typename ContainType>
class MyContainer{
  ContainType value;
};

MyContainer<int, std::vector<int> > a;

// 3
template <typename T, typename ContainType = vector<T>>
class MyContainer{
  ContainType value;
};

MyContainer<int> a;
```

### declval
```c++
//  返回一个类型的右值引用，不管是否有没有默认构造函数或该类型不可以创建对象
// (可以用于抽象基类)
#include <utility> 

struct A {              // abstract class
  virtual int value() = 0;
};
 
class B : public A {
  int val_;
public:
  B(int i,int j):val_(i*j){}
  int value() {return val_;}
};
 
int main() {
  decltype(std::declval<A>().value()) a;  // int a
  decltype(std::declval<B>().value()) b;  // int b
  decltype(B(0,0).value()) c;   // same as above (known constructor)
  a = b = B(10,2).value();
  std::cout << a << '\n';
  return 0;
}

```

### enable_if
```c++
// 定义
template<bool, typename T = void>
struct enable_if {};

template<typename T>
struct enable_if<true, T>{
    typedef T type;
};

// enable_if 都用来判断模板参数的类型

// 校验函数模板参数类型必须是int
template<typename T>
typename std::enable_if<std::is_intergral<T>::value, bool>::type 
is_odd(T t) {
    return bool(T % 2);
}
```

### SFINAE
```c++
// c++98
template <typename T>
struct has_type {
private:
    typedef char one;
    typedef struct { char data[2]; } two;
    // 存在的话返回类型为 one
    template <typename U> static one test(typename U::type*);
    // 不存在的话返回类型为 two
    template <typename U> static two test(...);
public:
    enum { value = sizeof(test<T>(0)) == sizeof(one) };
};
// 如果 T::type 存在的话就会选择第一个重载，否则就会选择第二个重载，由此判断 T::type 是否存在

// c++17 void_t<...> 其实就是 void，但它可以在 SFINAE 中帮助判断类型是否存在
template <typename T, typename = void>
struct has_type : std::false_type {};
template <typename T>
struct has_type<T, void_t<typename T::type>> : std::true_type {};

// c++11
// 判断是否有get函数
template<typename T, typename = void>
struct has_get: public std::false_type{};
template<typename T>
struct has_get<T, void_t<decltype(std::declval<T>().get())>> : public std::true_type{};
std::cout << has_get<A>::value << std::endl;

// 判断是否为智能指针
template<typename T, typename = void>
struct is_smart_point : public std::false_type {};
template<typename T>
struct is_smart_point<T, 
    void_t<decltype(std::declval<T>().operator->()),
           decltype(std::declval<T>().get())>> : public std::true_type {};

std::shared_ptr<int> a;
cout << is_smart_point<decltype(a)>::value;
```

### 判断类是否存在成员变量、函数、类型
```c++
// 利用SFINAE

// 构建void_t, c++17才有std::void_t
template<typename... Ts> struct make_void{typedef void type;};
template<typename... Ts> using void_t = typename make_void<Ts...>::type;

template<typename T, typename = void>
struct has_member : public std::false_type{};

// 判断成员变量
template<typename T>
struct has_member<T, void_t<decltype(T::value)>> : public std::true_type{};

// 判断成员函数
template<typename T>
struct has_member<T, void_t<decltype(std::declval<T>().func())>> : public std::true_type{};

// 对于有参数的成员函数
template<typename T>
struct has_member<T, void_t<decltype(&T::func)>> : public std::true_type{}; // 这个也可以判断成员变量，但是反之不行

// 判断成员类型
template<typename T>
struct has_member<T, void_t<typename T::type>> : public std::true_type{};

// 通过变量判断成员
template<typename T>
bool has_member_func(T& t) {return has_member<T>::value;}

```

### 结构体元素数量
```c++

struct AnyType {
    template <typename T>
    operator T();
};

template <typename T, typename = void, typename ...Ts>
struct CountMember {
    constexpr static size_t value = sizeof...(Ts) - 1;
};

template <typename T, typename ...Ts>
struct CountMember<T, std::void_t<decltype(T{Ts{}...})>, Ts...> {
    constexpr static size_t value = CountMember<T, void, Ts..., AnyType>::value;
};

int main(int argc, char** argv) {
    struct Test { int a; int b; int c; int d; };
    printf("%zu\n", CountMember<Test>::value);
}

```

### tuple 遍历
```c++
// 用模板实现tuple的遍历

template<typename T, size_t N>
struct PrintTuple {
    void operator()(T & t) {
        PrintTuple<T, N-1>()(t);
        cout << std::get<N-1>(t) << endl;
    }
};

template<typename T>
struct PrintTuple<T, 1> {
    void operator()(T & t) {
        cout << std::get<0>(t) << endl;
    }
};

template<typename ...Argvs>
void func(std::tuple<Argvs...> & t) {
    PrintTuple<decltype(t) ,sizeof...(Argvs)>()(t);
}

auto t = make_tuple("dadsa", 1);
func(t);
```

### shared_from_this
```c++
// 如何在类的内部获得自己的shared_ptr

// 写法一
class A {
public:
    void func(){
        auto p = shared_ptr<A>(this);
    }
};
A a;
a.func();
// 上述写法错误，因为会导致析构两次this。shared_ptr<>(this)的写法很危险

// 正确做法: 继承类enable_share_from_this<>，使用函数shared_from_this返回该智能指针
#include <memory>
class A : public std::enable_share_from_this<A> {
public:
    void func(){
        auto p = shared_from_this();
    }
};
auto a = make_ptr<A>(); // 要使用智能指针包裹
a.func();               // 正确
```

### 类型检测
```
SFINAE
enable_if 
Concept c++20
void_t
```
