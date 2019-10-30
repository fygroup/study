### 书籍
```
https://github.com/chuenlungwang/cppprimer-note
https://docs.microsoft.com/zh-cn/cpp/standard-library/cpp-standard-library-reference?view=vs-2019
https://zh.cppreference.com/w/%E9%A6%96%E9%A1%B5
https://zh.wikibooks.org/wiki/C%2B%2B
https://github.com/jobbole/awesome-cpp-cn [C++ 资源大全中文版]
```

### operator new
```
//三种形式
(1) void* operator new(size_t) throw(std::bad_alloc);
    > 用法
        A *a = new A;
    > 做了三件事
        调用operator new (sizeof(A))
        调用A:A()
        返回指针
    > 失败时抛出bad_alloc

(2) void* operator new(size_t, nothrow_value) throw();
    > 用法
        A* a = new(std::nothrow) A;
    > 同上，但是失败时返回null
        调用operator new (sizeof(A), nothrow_value)
        调用A:A()
        返回指针
(3) void* operator new(size_t, void* ptr) throw();
    > 用法
        //在ptr所指地址上构建一个对象(通过调用其构造函数)
        char ptr[1024];
        A* a = new(ptr) A();
    > 本身返回ptr
    > 可以被重载
    
(4) 示例
    class a{
    public:
        T n;
        a(T _n):n(_n){
            cout << "T" << endl;
        }
        virtual ~a(){}
        void* operator new(size_t n){
            cout << "void* operator new(size_t n)" << endl;
            cout << n << endl;
            return malloc(n);
        }

        void* operator new(size_t n, void* p){
            cout << "void* operator new(size_t n, void* p)" << endl;
            cout << n << endl;
            return p;
        }
    };

    int main() {
        char buf[1024];
        a<int> * x = new a<int>(3);
        cout << x->n << endl;
        a<int> * y = new(buf) a<int>(4);
        cout << x->n << endl;
        return 0;
    }
```

### allocator
```
https://zhuanlan.zhihu.com/p/34725232

#include <memory>

(1) 概念
    allocator是STL的重要组成,allocator除了负责内存的分配和释放，还负责对象的构造和析构
    例如vector的class如下
    template<typename T, typename Alloc = allocator<T>>
    class vector{
        //每个vector内部实例一个allocator
        Alloc data_allocator;
    };
    template<typename T>
    class allocator{ 

    };
    std::vector<int> v;
    等价于
    std::vector<int, allocator<int>> v;

(2) 重要用法
    // 以下几种自定义类型是一种type_traits技巧
        allocator::value_type
        allocator::pointer
        allocator::const_pointer
        allocator::reference
        allocator::const_reference
        allocator::size_type
        allocator::difference
    // 配置空间，足以存储n个T对象。第二个参数是个提示。实现上可能会利用它来增进区域性(locality)，或完全忽略之
        pointer allocator::allocate(size_type n, const void* = 0)
    // 释放先前配置的空间
        void allocator::deallocate(pointer p, size_type n)
    // 调用对象的构造函数，等同于 new((void*)p) value_type(x)
        void allocator::construct(pointer p, const T& x)
    // 调用对象的析构函数，等同于 p->~T()
        void allocator::destroy(pointer p)

(3) 示例
    allocator<string> alloc;
    string *s = alloc.allocate(10);
    string *s1 = s;
    alloc.construct(s1++);
    alloc.construct(s1++, "dasd");
    alloc.construct(s1++, "dasddasdas");
    cout << s[0] << endl;
    alloc.destory(s1);       //析构s1处的内存
    alloc.deallocate(s, 10); //析构整个内存
    // 复制和填充未初始化的内存
    // allocator 类定义了两个可以构建对象的算法，以下这些函数将在目的地构建元素，而不是给它们赋值
    vector<string> list(10, "aaaa");
    uninitialized_copy_n(list.begin(), 5, s);  //构建填充
    uninitialized_fill_n(list.begin(), 5, s);  //拷贝填充
```

### 智能指针
```
#include <memory>

(1) 构造
    shared_ptr<int> a(new string("dasdas"));
    shared_ptr<int> a = make_shared<int>("dasdas")
    shared_ptr<vector<int>> a(new vector<int>(10));
    cout << a << endl;  //0x5633802bfe70

(2) 切片
    shared_ptr<int> a(new int [10] {1,2,3,4,5});
    int* pI = a.get();
    cout << *a << endl;  //1
    cout << *(a+1) << endl; //错误
    cout << a[0] << endl;  //错误
    shared_ptr<int[]> a(new int [10] {1,2,3,4,5});
    cout << *a << endl;  //错误
    cout << a[0] << endl;  //1
    shared_ptr<vector<int>> vc = make_shared<vector<int>>(10,3);
    cout << vc->operator[](1) << endl;
    cout << vc->size() << endl;

(3) 引用次数
    shared_ptr<int> a = make_shared<int>(10);
    cout << a.use_count() << endl;	
    //shared_ptr多个指针指向相同的对象。shared_ptr使用引用计数，每一个shared_ptr的拷贝都指向相同的内存
    //每使用他一次，内部的引用计数加每析构一次，内部的引用计数减1，减为0时，自动删除所指向的堆内存
    //shared_ptr内部的引用计数是线程安全的，但是对象的读取需要加锁
    void func(shared_ptr<int> & a){       //引用 不会改变计数
        cout << a.use_count() << endl;
    }
    void func(shared_ptr<int> a){		//复制 会改变计数
        cout << a.use_count() << endl;
    }

(4) 自定义析构函数
    shared_ptr<int[]> a(new int[1], [](int* a){
        cout << "delete" << endl;
        delete a;
    });

(5) 智能指针不能当右值
    void* a = NULL;
    //右值Segmentation fault (core dumped)
    a =  static_cast<void*>(auto_ptr<string>(new string("sadadada")).get()); //此时智能指针是右值
    cout << *static_cast<string*>(a) << endl; //出错！ 实际上智能指针早已析构了
    //必须先存成左值
    auto_ptr<string> x = auto_ptr<string>(new string("sadadada")); //必须现存成左值
    a =  static_cast<void*>(x.get());
    cout << *static_cast<string*>(a) << endl;

(6) 糟糕的auto_ptr
    1) 智能指针所有权
        auto_ptr<string> a = auto_ptr<string>(new string("aaaa"));
        auto_ptr<string> b;
        b = a;   // 赋值导致了 a失去了所有权，b获得了所有权
    2) auto_ptr不要与容器混合使用
        STL有一条规定：
        std::auto_ptr 不能和容器混合使用。
        原因是：容器里的元素使用的都是copy，而std::auto_ptr型数据copy后会发生拥有权转移。
        所以！！！auto_ptr几乎没用！！！
```

### IO体系
```
https://cloud.tencent.com/developer/article/1008625

(1) IO体系之间的关系
    > ios_base
        表示流的一般特征，如是否可读取，是二进制流还是文本流
    > ios(basic_ios)
        基于ios_base，其中包括一个指向streambuf对象的指针
    > steambuf
        为缓冲区提供内存，并提供用于填充缓冲区、访问缓冲区内容、刷新缓冲区、管理缓冲区内存的类方法
    > ostream
        由ios类派生而来，提供输出方法
    > istream
        由ios类派生而来，提供输入方法
    > iostream
        基于istream、ostream，继承类输入和输出

(2) 三个头文件
    1) <iostream>
        > 包含的头文件
            <ios>
            <streambuf>
            <istream>
            <ostream>
        > 这些头文件中包含的类
            > streambuf
                streambuf	basic_streambuf<char>
            > istream、wistream 从流中读取
            > ostream、wostream 写入流
            > iostream、wiostream 对流进行读写
    2) <fstream>
        包含的类
        > basic_filebuf 抽象原生文件设备(类模板)
            basic_filebuf由basic_streambuf派生而来
            template< 
                class CharT, 
                class Traits = std::char_traits<CharT>
            > class basic_filebuf : public std::basic_streambuf<CharT, Traits>
        > basic_ifstream 实现高层文件流输入操作(类模板)
            ifstream、wifstream 从文件中读取
        > basic_ofstream 实现高层文件流输出操作(类模板)
            ofstream、wofstream 写文件
        > basic_fstream 实现高层文件流输入/输出操作(类模板)
            fstream、wfstream 对文件进行读写
    3) <sstream>
        包含的类
        > basic_stringbuf
            https://zh.cppreference.com/w/cpp/io/basic_stringbuf
            basic_stringbuf由basic_streambuf派生而来
            template<
                class CharT, 
                class Traits = std:char_chaits<CharT>,
                class Allocator = std:alloctor<CharT>
            > class basic_stringbuf : public std:basic_streambuf<CharT, Traits>
        > istringstream、wistringstream 从字符串中读取
        > ostringstream、wostringstream 写入到字符串中
        > stringstream、wstringstream 对字符串进行读写
        
(3) 流的基本用法
    1) ostream
        operator <<
        cout.put(‘H’).put(‘i’)
        write(buf, len)
        write()返回一个ostream对象的引用
        cout.write (buf, len)  //char buf[len]
    2) istream
        opeartor>>
        int ch = cin.get()
        cin.get(ch1).get(ch2)
        cin.getline(buf, 10)
        cin.read(buf, 5)

(4) 缓冲区
    https://izualzhy.cn/stream-buffer
    https://github.com/iassasin/streambuf_examples
    streambuf的构造函数是project，不能直接用
    streambuf有两种用法，一是直接使用各个接口，二是继承并实现新的I/O channels
    streambuf有两个子类，stringbuf, filebuf
    1) 用法
        //自定义缓冲区cout
        char buf[1024] = {0};
        stringbuf a;
        a.pubsetbuf(buf, 1024);
        std::cout.rdbuf(&a);
        cout << "ddasdaasdas";
        printf("-- %s --\n", a.str().c_str());
        printf("-- %s --\n", buf);
        //自定义缓冲区file
        std::ifstream file;
        char buf[10241];
        file.rdbuf()->pubsetbuf(buf, sizeof buf);
        file.open("/usr/share/dict/words");
        int cnt = 0;
        for (std::string line; getline(file, line);) {
            cnt++;
        }
        std::cout << cnt << '\n';


    2) 自定义缓冲区

```