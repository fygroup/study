### 书籍
```
https://github.com/chuenlungwang/cppprimer-note
https://docs.microsoft.com/zh-cn/cpp/standard-library/cpp-standard-library-reference?view=vs-2019
https://zh.cppreference.com/w/%E9%A6%96%E9%A1%B5
https://zh.wikibooks.org/wiki/C%2B%2B
```

### IO体系与stream
```
https://cloud.tencent.com/developer/article/1008625

(1) 两个平行的虚基类
    streambuf和ios类，所有流类均以两者之一作为基类

(2) 分类
    1) streambuf
        类提供对缓冲区的低级操作：设置缓冲区、对缓冲区指针操作区存/取字符
    2) ios_base、ios类
        记录流状态，支持对streambuf 的缓冲区输入/输出的格式化或非格式化转换
    3) stringbuf
        使用串保存字符序列。扩展 streambuf 在缓冲区提取和插入的管理
    4) filebuf
        使用文件保存字符序列。包括打开文件；读/写、查找字符

(3) 缓冲区streambuf
    (cout/cin/clog/ifstream/ofstream)都有自己的流缓冲区(streambuf)
    通过rdbuf接口可以获取当前的streambuf，也可以设置新的streambuf
    1) basic_streambuf
        typedef basic_streambuf<char> streambuf
        basic_streambuf是虚基类
    2) 实现接口
        class mystream : public std::streambuf{
        public:
            //

        }    


(2) 简介
    

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

(4) istringstream、ostringstream、stringstream
    istringstream，由istream派生而来，提供读string的功能
    ostringstream，由ostream派生而来，提供写string的功能
    stringstream，由iostream派生而来，提供读写string的功能

(5) ifstream、ofstream
    ofstream，由ostream派生而来，用于写文件
    ifstream，由istream派生而来， 用于读文件
    fstream，由iostream派生而来，用于读写文件

(6) streambuf
    


```