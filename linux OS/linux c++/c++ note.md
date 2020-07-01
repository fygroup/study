### 书籍
```
https://github.com/chuenlungwang/cppprimer-note
https://docs.microsoft.com/zh-cn/cpp/standard-library/cpp-standard-library-reference?view=vs-2019
https://zh.cppreference.com/w/%E9%A6%96%E9%A1%B5
https://zh.wikibooks.org/wiki/C%2B%2B
https://github.com/jobbole/awesome-cpp-cn [C++ 资源大全中文版]
```

C++笔记
|编译器为我们做了大量的优化工作，不要以为什么都理所应当



### 参考
```
https://zh.cppreference.com/w/%E9%A6%96%E9%A1%B5


//注意区分c/c++语言库与linux系统库

(1) 概念库
    <concepts> (C++20 起)	基础库概念

(2) 协程库
    <coroutine> (C++20 起)	协程支持库

(3) 工具库
    <cstdlib>	通用工具：程序控制、动态内存分配、随机数、排序与搜索
    <csignal>	信号管理的函数与宏常量
    <csetjmp>	保存执行语境的宏，及跳转到某个执行语境的函数
    <cstdarg>	变长实参列表的处理
    <typeinfo>	运行时类型信息工具
    <typeindex> (C++11 起)	std::type_index
    <type_traits> (C++11 起)	编译时类型信息
    <bitset>	std::bitset 类模板
    <functional>	函数对象、函数调用、绑定操作、引用包装
    <utility>	多种工具组件
    <ctime>	C 风格的时间/日期工具
    <chrono> (C++11 起)	C++ 时间工具
    <cstddef>	标准宏和 typedef
    <initializer_list> (C++11 起)	std::initializer_list 类模板
    <tuple> (C++11 起)	std::tuple 类模板
    <any> (C++17 起)	std::any 类
    <optional> (C++17 起)	std::optional 类模板
    <variant> (C++17 起)	std::variant 类模板
    <compare> (C++20 起)	三路比较运算符支持
    <version> (C++20 起)	提供依赖于实现的库信息

(4) 动态内存管理
    <new>	低层内存管理工具
    <memory>	高层内存管理工具
    <scoped_allocator> (C++11 起)	嵌套分配器类
    <memory_resource> (C++17 起)	多态分配器及内存资源

(5) 数值界限
    <climits>	整型类型的界限
    <cfloat>	浮点类型的界限
    <cstdint> (C++11 起)	定长整数及其他类型的界限
    <cinttypes> (C++11 起)	格式化宏、 intmax_t 及 uintmax_t，数学工具及转换
    <limits>	查询算术类型属性的标准化方式

(6) 错误处理
    <exception>	异常处理工具
    <stdexcept>	标准异常对象
    <cassert>	将其实参与零比较的条件性编译宏
    <system_error> (C++11 起)	定义 std::error_code，依赖于平台的错误码
    <cerrno>	含有最近一次错误号的宏
    <contract> (C++20 起)	契约违规信息

(7) 字符串库
    <cctype>	确定字符数据中所含类型的函数
    <cwctype>	确定宽字符数据中所含类型的函数
    <cstring>	多种窄字符串处理函数
    <cwchar>	多种宽及多字节字符串处理函数
    <cuchar> (C++11 起)	C 风格 Unicode 字符转换函数
    <string>	std::basic_string 类模板
    <string_view> (C++17 起)	std::basic_string_view 类模板
    <charconv> (C++17 起)	std::to_chars 与 std::from_chars

(8) 容器库
    <array> (C++11 起)	std::array 容器
    <vector>	std::vector 容器
    <deque>	std::deque 容器
    <list>	std::list 容器
    <forward_list> (C++11 起)	std::forward_list 容器
    <set>	std::set 及 std::multiset 关联容器
    <map>	std::map 及 std::multimap 关联容器
    <unordered_set> (C++11 起)	std::unordered_set 及 std::unordered_multiset 无序关联容器
    <unordered_map> (C++11 起)	std::unordered_map 及 std::unordered_multimap 无序关联容器
    <stack>	std::stack 容器适配器
    <queue>	std::queue 及 std::priority_queue 容器适配器
    <span> (C++20 起)	std::span 视图

(9) 迭代器库
    <iterator>	范围迭代器

(10) 范围库
    <ranges> (C++20 起)	范围访问、原语、要求、工具及适配器

(11) 算法库
    <algorithm>	对范围操作的算法
    <execution> (C++17 起)	针对算法的并行版本的预定义执行策略

(12) 数值库
    <cmath>	常用数学函数
    <complex>	复数类型
    <valarray>	表示和操作值的数组的类
    <random> (C++11 起)	随机数生成器及分布
    <numeric>	容器中值的数值运算
    <ratio> (C++11 起)	编译时有理数算术
    <cfenv> (C++11 起)	浮点环境访问函数
    <bit> (C++20 起)	位操纵函数

(13) 输入/输出库
    <iosfwd>	所有输入/输出库中的类的前置声明
    <ios>	std::ios_base 类、std::basic_ios 类模板及数个 typedef
    <istream>	std::basic_istream 类模板及数个 typedef
    <ostream>	std::basic_ostream、std::basic_iostream 类模板及数个 typedef
    <iostream>	数个标准流对象
    <fstream>	std::basic_fstream、std::basic_ifstream、std::basic_ofstream 类模板及数个typedef
    <sstream>	std::basic_stringstream、std::basic_istringstream、std::basic_ostringstream 类模板及数个 typedef
    <syncstream> (C++20 起)	std::basic_osyncstream、std::basic_syncbuf 及 typedef
    <strstream> (C++98 中弃用)	std::strstream、std::istrstream、std::ostrstream
    <iomanip>	控制输入输出格式的辅助函数
    <streambuf>	std::basic_streambuf 类模板
    <cstdio>	C 风格输入输出函数

(14) 本地化库
    <locale>	本地化工具
    <clocale>	C 本地化工具
    <codecvt> (C++11 起)(C++17 中弃用)	Unicode 转换设施

(15) 正则表达式库
    <regex> (C++11 起)	支持正则表达式处理的类、算法及迭代器

(16) 原子操作库
    <atomic> (C++11 起)	原子操作库

(17) 线程支持库
    <thread> (C++11 起)	std::thread 类及支持函数
    <mutex> (C++11 起)	互斥原语
    <shared_mutex> (C++14 起)	共享的互斥原语
    <future> (C++11 起)	异步计算的原语
    <condition_variable> (C++11 起)	线程等待条件

(18) 文件系统库
    <filesystem> (C++17 起)	std::path 类及 支持函数

```

```
#include <iostream>
#include <string>
#include <fstream>

读写文件
fstream on;
on.open(string.c_str(),ios::in)
on.open(string.c_str(),ios::out)

默认方式如下：
ofstream out("...", ios::out);
ifstream in("...", ios::in);
fstream foi("...", ios::in|ios::out);

字符转int atoi
int转字符 itoa

string转换为char*有3中方法：
1.data
string str="good boy";
const char *p=str.data();
2.c_str
string str="good boy";
const char *p=str.c_str();
3. copy
string str="good boy";
char p[20];
str.copy(p,5,0); //这里5，代表复制几个字符，0代表复制的位置
*(p+5)='\0'; //要手动加上结束符

或者:

string str="good boy";
char *p;
int len = str.length();
p=(char *)malloc((len+1)*sizeof(char));
str.copy(p,len,0);

读取文件
o=ifstream()
o.is_open()

g++   -lz

use namespace std;
ios_base::fmtflags initial;
initial=os.setf(ios_base::fixed);
os.setf(ios::showpoint)  //设置成显示小数点模式
os.precision(0) //显示多少位小数
os.width(10)  //设置字段宽度


ofstream fout;
fout.open(file);
fout.is_open();
func(fout);
func(cout);

fout.write(line,n)

cout<< 12.321321113<<endl;
ios_base::fmtflags orig = cout.setf(ios_base::fixed,ios_base::floatfield);
std::streamsize prec=cout.precision();
cout.precision(2);
cout<< 12.321321113<<endl;
cout.setf(orig);
cout.precision(prec);
cout<< 12.321321113<<endl;


class A

{

private:

      static const int a = 0;  //正确

      static const char b = 'b';  //正确

      const int c = 0; //错误，非static const不能在类里面初始化

      static const int d[10] = {0}; //错误，只有一般数据类型的static const才能在类里面初始化。

      static const int dd[10];   //外部初始化

};

static const int A::dd[10] = {0};

注意 静态变量是类级别的！！！！ 需要用类名去定义（无论这个变量是否在private中）
例如: int my::b=3;要定义在全局变量里面


std:cerr << 错误输出


把一个成员函数声明为const可以保证这个输入的成员函数不修改数据成员，但是，如果据成员是指针，则const成员函数并不能保证不修改指针指向的对象


//---重写 重载 重定义--------------
函数重载是指在一个类中声明多个名称相同但参数列表不同的函数，这些的参数可能个数或顺序，类型不同，但是不能靠返回类型来判断。
函数重写是指子类重新定义基类的虚函数。

//--转换--------------------------
肯定没问题的转换
1.允许从非常量转换到常量的类型转换。
2.允许从派生类到基类的转换。
3.允许数组被转换成为指向数组（元素）类型的指针，函数被转换成指向该函数类型的指针。
4.标准算术类型的转换（比如：把bool型和char型转换成int型）和类类型转换(使用类的类型转换运算符和转换构造函数)。

//---c++中的库--------------------
#include <cstdlib> 等价于c中的<stdlib.h>
#include <cstdio> 等价于c中的<stdio.h>
#include <ctime> <time.h>
#include <cctype> <ctype.h>
#include <iostream>

//---using------------------------
template <typename T>
using ArenaDeque = std::deque<T, ArenaAllocatorAdapter<T>>;
定义函数指针
typedef void (*FP)(int a);
using FP = void(*)(inta);
void func(int a){}

FP f = func;
f();
//---实例化 与 具体化
实例化：
定义 template void func<T>(T & a, T & b);
应用  func<T>(....)

具体化 template<> void func<T>(T & a, T & b);
应用 job a,b; func(a,b);

//---纯虚函数--------------------
纯虚函数不能实例化 ，但命名个类指针还是可以的。

//---iterator--------------------
i++ / i-- 时，其迭代的内容改变，但是迭代器本身的地址不变。
iterator相当于指针，对iterator进行++或--时，iterator指向的位置前移或后移，但是iterator本身的地址是不变的，和指针一样
//---继承------------------------
template<typename _Arg1, typename _Arg2, typename _Result>
struct binary_function
{
	using first_argument_type = _Arg1;
	using second_argument_type = _Arg2;
	using result_argument_type = _Result;
};

template<typename T>
struct equal_to:public binary_function<T,T,bool>
{
	bool operator()(const T & a, const T & b)const{
		return a == b;
	}
};
// 当前程序
readlink("/proc/self/exe", pPath, 1024); 当前程序的绝对路径

#ifndef GOOGLE_GLOG_DLL_DECL
# if defined(_WIN32) && !defined(__CYGWIN__)
#   define GOOGLE_GLOG_DLL_DECL  __declspec(dllimport)    全局变量
# else
#   define GOOGLE_GLOG_DLL_DECL
# endif
#endif

//---logging--------------------
GlogInit.cpp 中。
google::InitGoogleLogging(argv0);
google::SetStderrLogging(google::GLOG_WARNING);
strLogFile = strLogPath + "INFO_";
google::SetLogDestination(google::GLOG_INFO, strLogFile.c_str());
strLogFile = strLogPath + "WARNING_";
google::SetLogDestination(google::GLOG_WARNING, strLogFile.c_str());
strLogFile = strLogPath + "ERROR_";
google::SetLogDestination(google::GLOG_ERROR, strLogFile.c_str());
google::InstallFailureSignalHandler();
google::InstallFailureWriter(&CGlogInit::SignalHandle); 

//---tinyxml------------------- 解析xml
#include <tinyxml/tinyxml.h>
#include <tinyxml/tinyxmlhelper.h>
读操作
TiXmlDocument doc("config");	//初始化对象
if (!doc.LoadFile()) cout<< "wrong"; //载入到文档对象中
TiXmlElement* root =doc.RootElement();//找到根元素
//注意节点与元素的区别！
//nextSibling属性返回元素节点之后的兄弟节点（包括文本节点、注释节点即回车、换行、空格、文本等等）；
//nextElementSibling属性只返回元素节点之后的兄弟元素节点（不包括文本节点、注释节点）；
TiXmlNode* Node;
TiXmlElement* Element;
Element = Node->ToElement();
TiXmlNode* first_child = root->FirstChild();//第一个子节点
TiXmlElement* first_child = root->FirstChildElement();//第一个子元素
TiXmlNode* first_child = root->FirstChild("task");	//第一个task子节点
TiXmlElement* first_child = root->FirstChildElement("task"); // 第一个task子元素
TiXmlNode* first_child = root->FirstChild("task");	//第一个task子节点

TiXmlNode* Node = Node->NextSibling() //下一个兄弟节点
TiXmlNode* element = Node->NextSiblingElement() //下一个兄弟元素
TiXmlElement* element = root->NextSiblingElement("task"); //下一个"task"兄弟元素

task->value();//元素的名字
task->GetText();//元素的text 或者 task->FirstChildElement()->Value() ???
TiXmlAttribute* pAttr = pNode->FirstAttribute(); //第一个attr
pAttr->Name(); //attr的名字
pAttr->Value();//attr的值
pAttr = pAttr->next()   //遍历节点属性
std::string attr = task->Attribute("path"); //寻找"path"属性
写操作
TiXmlDocument doc;
TiXmlDeclaration* decl = new TiXmlDeclaration( "1.0", "", "" );
doc.LinkEndChild(dec1); //LinkEndChild的参数是指针
TiXmlElement * root = new TiXmlElement("Root");
doc.LinkEndChild(root);
TiXmlElement* element = new TixXmlElement("child");
roo.LinkEndChild(element);
element->SetAttribute("name", "MainFrame");	//添加属性 
element->SetDoubleAttribute("timeout", 123.456); // 添加double样式的属性
element->LinkEndChild(new TiXmlText("xxxxxxxxx")); //添加元素的text
dump_to_stdout(&doc); //??????????????????????
doc.SaveFile("settings.xml");  

int type = Node->Type() //查看当前节点类型
type == TiXmlNode::TINYXML_ELEMENT;
type == TiXmlNode::TINYXML_TEXT;


//---调用so库-------------------
>>无论调用什么都必须include头文件。
方法一
//编译时调用libxxx.so,头文件包含相关的头文件。
方法二
CDllHelper dllhelp;
dllhelp.Open(pls_path)
#define CC_AddFun(var,funname) { \
    var = cc_type_convert(var,dllhelp.GetDllFunAddress(funname)); \
    if(!var){ \
        dllhelp.Close();\
        return -1;\
    }\
}\

CC_AddFun(_GetTask,"GetTask");
P1 cc_type_convert(P1 ,P2 tmp2){
	return (P1)tmp2;
}

//---flock----------------------
文件锁是系统级别的
#include <cstdio>
#include <sys/file.h>
FILE* fp=fopen("","w");
flock(fp->_fileno,LOCK_EX);加锁
flock(fp->_fileno,LOCK_UN);去锁

//---数据库操作-----------------
#include "cppconn/driver.h"
#include "cppconn/resultset.h"
#include "cppconn/statement.h"
#include "cppconn/exception.h"
#include "cppconn/connection.h"
using namespace sql;
Driver *driver = get_driver_instance();
Connection *conn;
conn=driver->connect("127.0.0.1:3306", "root", "root");
conn->setSchema("dbname");
stmt = m_con->createStatement();
//Statement::executeQuery用于执行一个Select语句，它返回ResultSet对象
//Statement::executeUpdate方法主要用于执行INSERT, UPDATE, DELETE语句
void Connection::setSchema(const std::string& catalog);    //定义目标数据框
/* statement.h */   
ResultSet* Statement::executeQuery (const std::string& sql);   //查询 返回 ResultSet
int Statement::executeUpdate (const std::string& sql);   //执行INSERT, UPDATE, DELETE语句
bool Statement::execute (const std::string& sql);   //可执行所有语句，当执行查询时返回True,执行其他语句时，返回False
ResultSet* Statement::getResultSet();   //获取查询结果
uint64_t Statement::getUpdateCount();  //获取受影响记录的数量

/* resultset.h */ 
while(res->next())
	cout << res.getString("Cityname") << endl;
	//cout << res.getString(1) << endl;
//倒序
res->afterLast();
if (! res->isAfterLast())
	throw runtime_error('xxxxxxxxx');
while(res->previous)
	......;

//---判断文件是否存在--------------
int access(const char* fileName, int mode)
R_OK      测试读许可权
W_OK      测试写许可权
X_OK      测试执行许可权
F_OK      测试文件是否存在

成功执行时，返回0。失败返回-1

//---得到工作路径-----------------
/proc/self/exe 它代表当前程序
int nRet = readlink("/proc/self/exe", pPath, nLen - 1);
//nRet表示返回字符串的长度

//---改变工作目录-----------------
chdir();

//---RPC传输：PCFproto传输架构和protobuf序列化策略--------
1、protobuf序列化协议（非常好非常快）
（1）定义.proto文件（project.proto）
syntax="proto3";
package tutorial;
import "Common.proto"; //加载其他proto的内容
message CourseInfo   //定义类
{
	required string math	= 1;// CourseInfo.set_math("dasdada"); 取值时 CouseInfo.math;
}

message MutRequest
{
...
}

message MutResponse
{
	required string msg_code = 3;
	optional CourseInfo course_info = 5;  //定义其他类  CourseInfo a = MutResponse.add_course_info();
}

service RcfMutService  //使用rpc
{
  rpc DoServer ( MutRequest ) returns ( MutResponse );
  	  函数名     参数                   返回
}
//编译
protoc -I=. --cpp_out=. project.proto
此时会生成project.proto.pb.h和project.proto.pb.cc

（2）proto buffer 定义API
>>test.h
#include "project.proto.pb.h"
class SearchServiceImpl: public RcfMutService
{
public:
	SearchServiceImpl ();
	virtual SearchServiceImpl();
	virtual myfunction(
	::google::protobuf::RpcController* controller,
	::MutRequest* request,
	::MutResponse* response,
	::google::protobuf::Closure* done
	);
}

>>test.cpp
include "test.h"
my::myfunction(::goolge::protobuf::RpcController* controller, ::MutRequest* request, ::MutResponse* response, ::google::protobuf::Closure* done){
	.....;
	done->Run();
}

2、PCFproto使用
（1）服务端
#include "RCFProto.hpp"
#include "RcfModel.pb.h"
#include "RcfModelServiceImpl.h"

RCF::init();
RCF::RcfProtoServer server(RCF::TcpEndpoint("0.0.0.0",50000));
RCF::ThreadPoolPtr threadPoolPtr(new RCF::ThreadPool(1,50)); // Configure a thread pool with number of threads varying from 1 to 50. threadPoolPtr是个智能指针
threadPoolPtr->setThreadName("myname");
server.setThreadPool(threadPoolPtr);
CAnalysisServiceImpl pAnaServ;  //建立个class
server.bindService(pAnaServ);
server.start()   //非阻塞
while(1)
	RCF::sleepMS(10000);

（2）客户端
#include "RCFProto.hpp"
#include <google/protobuf/text_format.h>
#include ".....pb.hpp"

RCF::init();
RCF::PcfProtoChannel channel(RCF::TcpEndpoint("0.0.0.0",50000));
RcfMutService::Stub stub(&channel);
MutRequest st;
MutResponse stm;
stub.myfunction(NULL,st,stm,NULL);

//---索引对齐---------------------------
#pragma pack(1)   // 1 字节对齐

//---class 大小-------------------------
class A
{
};
sizeof(A)=1;

class A
{
public:
int print1(){ cout<<"This is A"<<endl;}
};
sizeof(A)=1; //函数不占内存

class A
{
public:
virtual int print1(){ cout<<"This is A"<<endl;}
};
sizeof(A)=8;//多了虚函数指针

struct A
{
	A();//c++中 struct就是类，A()是构造函数
}
//----------------------------------------
只有构造函数使用成员初始化程序
//----------------------------------------
模板函数 声明和定义要在一块写 不能分成两个文件
//---read------------------------------------
fIn.read((char*)InBuffer,InBufferSize); //返回一个流对象
size_t InBufferSize_ = fIn.gcount();  //可以得到刚才的读入字节数

//---sleep------------------------------
# include <unistd.h>
sleep(1000);

//---构造函数-----------------------------
Eigen::Matrix<double,Dynamic,Dynamic> mat;
mat = Matrix<T,Dynamic,Dynamic>(m,n);
//---默认参数--------------------------------
1、有函数声明(原型)时,默认参数可以放在函数声明或者定义中，但只能放在二者之一（在一个文件中）
double sqrt(double f = 1.0); //函数声明
double sqrt(double f)  //函数定义
{
  // ....  
} 
2、没有函数(原型)时,默认参数在函数定义时指定.
//没有 函数声明
double sqrt(double f = 1.0)  //函数定义
//---显式调用析构--------------------------------
用户显式调用析构函数的时候，只是单纯执行析构函数内的语句，不会释放栈内存，摧毁对象
但是 不要自作聪明随便调用析构
//---引用参数默认？----------------------------------
引用做参数时不能传一个定值（如数字或者const等~~~）
somefunc(int& a = 4) -> default argument for ‘int& a’ has type ‘int’
//---new-----------------------------------------
sringbad* my - new stingbad("dsadda");
//---复制 赋值-----------------------------------
一般在初始化时会用复制，例如：
my a = b;
my a(b);
my a;
a=b;  //这是会用赋值
//---<ctime>------------------------------------
clock_t start, end;
start = clock();
end = clock();
cout << (double)(end-start)/CLOCKS_PER_SEC << endl;
//---友元---------------------------------------
友元不能继承
//---转换 ---------------------------------------
dynamic_cast
static_cast
//---class中的静态函数----------------------------
静态函数只要在定义的时候需要static关键字，实现的时候就不需要了，否则会报错。
//---%------------------------------------------------
% is only defined for integer types. That's the modulus operator.
<cmath> fmod(m,n);
//---string find------------------------------------------
string a("dadsadsada");
a.find_first_not_of("dda");   // 在字符串中查找第一个与str中的字符都不匹配的字符，返回它的位置。
a.find_first_of("dda");   // 在字符串中查找第一个与str中的字符匹配的字符，返回它的位置。
//---string_algo---------------------------------------
$include <string>
#include <boost/algorithm/string.hpp>
//是用于处理字符串查找，替换，转换等一系列的字符串算法
using namespace boost;
using namespace std;
string str("dsadsada")
ends_with(str,"txt");  //判断后缀
to_upper_copy(str);	//大写转换
replace_first(str,"readme","followme");//替换

std::vector<std::string> param_vec {"dsadada","sadsada","gfgd"};
std::string R_command;
String_Algo::str_join(param_vec, " ", R_command);  //字符串拼接

//---error-------------------------------------------
#include <errno.h>
系统调用返回失败时，必须紧接着引用errno变量(errno值对应的错误提示信息)
perror(str) //输出到终端的格式化
strerror(errno) //输出到缓冲区的格式化
//---popen------------------------------------------
FILE* mp = popen(cmd,"r");  //投递任务，并且打开一个fp  非阻塞！！！
fgets(buf,buf_size,mp);   //如果任务没完成就等待
//---删除文件------------------------------------------------
#include <cstdio>
remove()
//---类中 引用成员和常量成员-----------------------------
必须采用初始化列表的方式。凡是有引用类型的成员变量或者常量类型的变量的类，不能有缺省构造函数。
class A
{
int & a;
A();
}
A::A(int & x):a(x){...}

//----引用传参----------------------------------------------
func(int & a){}
func(3); //错误
int a = 3;
func(a);

//---不可拷贝类--------------------------------------
显式地声明类的拷贝构造函数和赋值函数为私有函数
class nocopy
{
public:
     nocopy(){}
     virtual ~nocopy(){}
private:
     nocopy(const nocopy&);
     nocopy& operator=(const nocopy&);
};

class A:public nocopy  //此时A不能被拷贝
{

}


//---函数转换-----------------------------------------------------
typedef void* (*func)(void*);

void f(){
}
func new = (func)&f;

//---vector resize-------------------------------------------
vector<int> a;
a.resize(5)  //如果a是空，那么先申请5个内存，在设为0.
a.reserve() //申请5个空间
reserve表示容器预留空间，但并不是真正的创建对象，需要通过insert（）或push_back（）等创建对象。
resize既分配了空间，也创建了对象
//---class 回调函数-----------------------------------
在类封装回调函数：

 a.回调函数只能是全局的或是静态的。
 b.全局函数会破坏类的封装性，故不予采用。
 c.静态函数只能访问类的静态成员，不能访问类中非静态成员
//---默认参数----------------------------
带有默认值参数的函数，在实现的时候，参数上是不能有值的。
//---类 成员引用---------------------------------------
引用类型的成员变量的初始化问题,它不能直接在构造函数里初始化，必须用到初始化列表，且形参也必须是引用类型。
有引用类型的成员变量的类，不能有缺省构造函数。原因是引用类型的成员变量必须在类构造时进行初始化
//---复制 陷阱------------------------------
class A
{
public:
	char* a;
	int n;
	A(int i):n(i){
		a= (char*)malloc(i);
	}
	A(){}
	~A(){free(a);}
};
int main(){
	A x;
	x = A(5);      //此时新建A(5)属于左值，马上就会销毁，A(5)中的a已经不存在了！
	x.a[0] = 'a';  //error
}	
//---char to int----------------------------------------------------------------
如果c默认初始化值（依赖编译平台）大于128那么X为负值，如果没有初始化，按以下处理也不会为负值
int x=0;
unsigned char c;
//---eof-----------------------------------------------------------------------------
按照一般思维，应该就是到达文件尾，就eof()应返回true，但事实上，在读完最后一个数据时，eofbit仍然是false。只有当流再往下读取时，发现文件已经到结尾了，才会将标志eofbit修改为true。这也就是为什么使用while(!readfile.eof()）会出多现读一行的原因。
既然已经知道了原因，那么，为了避免这样的情况，可以使用fIn.peek()!=EOF来判断是否到达文件结尾，这样就能避免多读一行。更改为：
//---c_str()--------------------------------------------------------------------------
string c_str() 返回的是const char*！！！
//---memset--------------------------------------------------------------------------
当内存比较大时，memset还是比较费时间的
//---文件 stat------------------------------------------------------------------------
#include<sys/stat.h>  
#include<uninstd.h>
struct stat buf;
result = stat("filename",&buf);
if (result !=0) error;
buf.st_size//文件大小
//---目录 stat------------------------------------------------------------------------
#include <unistd.h>
#include <dirent.h>
struct dirent* ptr;
DIR* dir = opendir("/");
while((ptr = readdir(dir)) != NULL){
ptr->d_name
ptr->d_type //8 file 10 linkfile 4 dir
}

```

### pthread_create在类中的使用
```
在C++的类中，普通成员函数不能作为pthread_create的线程函数，如果要作为pthread_create中的线程函数，必须是static

// 线程封装
class CThread {

public:
    pthread_t tid;
public:
    CThread(){
        tid = 0;
    }

    virtual ~CThread(){}

    bool start() {
        return 0 == pthread_create(&tid, NULL, CThread::callback, this);
    }

    void join() {
        if (tid) {
            pthread_join(tid, NULL);
            tid = 0;
        }
    }

    // static函数
    // 里面不能访问非静态成员变量
    static void *callback(void *arg) {
        CThread *cur_cthread = (CThread*)arg;
        // 调用成员函数，run函数可以访问成员变量
        cur_cthread->run();
        return (void*)NULL;
    }

    void run();
};
```

### container_of
```
已知结构体type的成员member的地址ptr，求解结构体type的起始地址)

#define container_of(ptr,type,member) ({\
    const typeof(((type*)0)->member)  *_mptr = (ptr); \
    (type*)((char*)_mptr - offsetof(type,member));  })

#define offsetof(type,member) ((size_t)&((type*)0)->member)
```

### 模板的特化(具体化)和偏特化
```
1、模板函数
    // 声明
    template<typename T, typename T1> void func(T, T1);
    // 特化
    template<>
    void func(int a, char b) {
        ...
    }
    // 偏特化
    template<typename T1>
    void func(char a, T1 b) {
        ...
    }

2、模板类
    // 声明
    template<typename T, typename T1> class Test
    {
        ...
    };
    // 特化
    template<> class Test<int, char>
    {
        ...
    };
    //偏特化
    template<typename T1> class Test<int, T1>
    {
        ...
    };

// 特化 > 偏特化 > 模板类
```

### 模板的实例化
```
(1) 显示实例化：  
    存在以下模板函数
    template <typename T>
    void Swap(T &a, T &b)
    > 第一种方式是声明所需的种类，用<>符号来指示类型，并在声明前加上关键词template，如下：
        template void Swap<int>(int &, int &);
    > 第二种方式是直接在程序中使用函数创建，如下：
        Swap<int>(a,b);
    显式实例化直接使用了具体的函数定义，而不是让程序去自动判断。

(2) 隐式实例化
    就是最正常的调用，Swap(a,b)
```

### 处理模板化基类内的名称
```
template<typename T>
class LoggingMsgSender:public MsgSender<T>
{
public:	
    using MsgSender<T>::sendClear;  //必须先要声明，否则编译器不知道MsgSender是否有sendclear函数，或者
    typedef typename T::mytype mytype; //定义 T的vector
    typedef typename std::vector<T> T_vector;
}

template<typename T>  
void func(const T & t){
    typename T::const_iterator iter(t.begin());      //重要
}
```

### for_each waitpid
```
#include <algorithm>
#include <sys/wait.h>
for_each(vec.begin(),vec.end(),[](pid_t & pd){waitpid(pd,NULL,0);});
```

### 模板元编程 type mapping

### class static struct初始化
```
class my
{
public:
    typedef struct _MY{}MY;
    static list<MY*> a;
};

list<my::MY*> my::a;  
```

### friend
```
1、友元的普通用法
    class Point
    {
    public:
        Point(double xx,double yy)
        {
            x=xx;
            y=yy;
        }
        void GetXY();
        friend double Distance(Point &a,Point &b);
    protected:
    private:
        double x,y;
    };

    Point p1(3.0,4.0),p2(6.0,8.0);
    double d = Distance(p1,p2);     //友元函数的调用方法，同普通函数的调用一样，不要像成员函数那样调用

2、友元的template
    template<typename T>
    class my {
        template<typename T1>   // 重要！！！
        friend my<T1> & operator *(my<T1> & my1, my<T1> & my2){  // friend 
            my1.i *= my2.i;
            return(my1);
        }
    }

3、友元类
    (1) 友元类
        class Node
        {
        private:
            int data;
        public:
            friend class BinaryTree;
        };

        class BinaryTree
        {
        private:
            Node x;
        public:
            int find(){
                x.data ;//可以访问Node的private
            }
        }

    (2) 友元函数
        class Node
        {
        private:
            int data;
        public:
            friend int BinaryTree::find(); //这样find方法就可以直接访问Node中的私有成员变量了，而BinaryTree中的其他的方法去不能够直接的访问Node的成员变量
        };

    (3) 友元的继承
        class A
        {
        public:    
            struct A_S{};
            friend class A_S;
        private:
            int a;
        };

        class B:public A:A_S
        {
        public:
            void func(){
                a = 1; //继承了A_S,可以访问class A的private
            }
        }

```

### 函数指针
```
typedef (int*)(*func)(int,char);
int* myfunc(int,char){}
func = myfunc;
//&(func) 等价于 func
```

### static变量
static 变量最好都写在cpp文件中，除非hpp用到的那个static变量

### makefile .o 文件有依赖时，是有顺序的
```
all: fz16.o fastqz.o libzpaq.o FastqReader.o FileOpt.o muti_Process.o
	$(CXX) -o $(TARGET) $^ $(FLAGS) $(LIBS)
fastqz.o依赖于muti_Process.o，muti_Process.o要写在fqstqz.o的后面
```

### ostringstream
```
std::ostringstream str;
str << "abc" << 2 << "dda";
//格式化一个字符串，但通常并不知道需要多大的缓冲区
```

### private构造函数的调用
```
对于class本身，可以利用它的static公有成员，因为它们独立于class对象之外，不必产生对象也可以使用它们

如果在外部使用private构造函数，有以下两种方法：
(1) 添加friend
    class Obj{
    public:
        friend class Obj1;    
    private:
        Obj(){};
    };

    class Obj1{
        Obj CreateObj(){
            return Obj();  //可以使用Obj里的private
        }
    }

(2) static调用
    class Obj{
    public:
        static Obj CreateObj(){
            return Obj()
        }
    private:
        Obj(){}    
    };
    Obj a = Obj::CreateObj();
```

### 获取文件绝对路径
```
realpath(file_name, abs_path_buff)
返回值为0表示错误
```

### 判断文件夹是否存在，不存在创建文件夹
```
if (access(tarDir,F_OK)!=0){
    ASSERT_ERROR(mkdir(tarDir,S_IRWXU),"mkdir tmp wrong");
}
```

### 文件夹操作
```
#include <sys/types.h>   
#include <dirent.h>
DIR* dir = opendir(path);
struct stat s_buf;
stat(dir,&s_buf);
if (S_ISDIR(s_buf.st_mode)) #is dir
S_ISLNK(st_mode)：是否是一个连接.
S_ISREG(st_mode)：是否是一个常规文件.
S_ISDIR(st_mode)：是否是一个目录
S_ISCHR(st_mode)：是否是一个字符设备.
S_ISBLK(st_mode)：是否是一个块设备
S_ISFIFO(st_mode)：是否 是一个FIFO文件.
S_ISSOCK(st_mode)：是否是一个SOCKET文件 

// 实例
#include <sys/types.h>
#include <sys/stat.h>
#include <dirent.h>

getcwd()获取的是当前工作路径，而不一定是程序的路径
fchdir(); //改变当前工作目录
rewinddir(); //重设读取目录的位置为开头位置

DIR* dp = opendir(path);
struct stat stat_buf;
struct dirent* entry;
while ((entry = readdir(dp)) != NULL){
	stat(entry->d_name, &stat_buf);
	if (S_ISREG(stat_buf.st_mode)){
		cout << entry->d_name << endl;
	}

}

rewinddir(dp);
closedir(dp);

```

### 字符串数组初始化
```
const char* a[5] = {"aa","ab","ada","dadad","fddgds"};
```

### 模板特例化
```
template<>
class a<int>{};
```

### 宏
```
__COUNTER__: 递增数。但是这个宏不能重新置0
__LINE__：在源代码中插入当前源代码行号；
__FILE__：在源文件中插入当前源文件名；
__DATE__：在源文件中插入当前的编译日期
__TIME__：在源文件中插入当前编译时间；
__STDC__：当要求程序严格遵循ANSI C标准时该标识被赋值为1；
__cplusplus：当编写C++程序时该标识符被定义。
```

### 虚函数
```
//要实现C++的多态性必须要用到虚函数，并且还要使用引用或者指针

//需要注意：
    只有类的成员函数才能声明为虚函数，虚函数仅适用于有继承关系的类对象。普通函数不能声明为虚函数。
    静态成员函数不能是虚函数，因为静态成员函数不受限于某个对象。
    内联函数（inline）不能是虚函数，因为内联函数不能在运行中动态确定位置。
    构造函数不能是虚函数。
    析构函数可以是虚函数，而且建议声明为虚函数。
```

### 模板类初始化(构造函数里)
```
template<typename T>
class X
{
   a<T> x;
   X(){
       x = a<int>();
   }
}
```

### 优先队列
```
struct cmp
{
	bool operator()(int & a, int & b){
		return a == b;
	}
};

bool cmp1(int & a, int & b){
	return a == b;
}

priority_queue<int,vector<int>,cmp> a;
priority_queue<int,vector<int>,bool(*)(int & a, int & b)> b(cmp1);

```

### 容器简介
```
list vector array deque
stack queue priority_queue
map set

> list
    list底层数据结构是双向链表，支持快速增删
    
> vector
    vector底层数据结构是数组，支持快速随机访问

> array
    长度固定的vector

> deque
    deque是双端队列，底层数据结构是分段连续线性空间
    中控器(map)是一块连续的空间，其中每个元素是指向缓冲区的指针，缓冲区才是deque存储数据的主体
    可以在两端进行push, pop 
    可以在内部进行插入和删除操作，但性能不及list
    支持随机访问，但性能不及vector

> stack
    栈是先进后出FILO数据结构，底层默认用deque实现
    // 模板类型
    template <class T, class Container = deque<T>>
    class stack

> queue
    队列是先进先出FIFO数据结构，底层默认用deque实现
    // 模板类型
    template <class T, class Container = deque<T>>
    class queue

> priority_queue
    priority_queue 容器适配器定义了一个元素有序排列的队列。默认队列头部的元素优先级最高，需要定义优先级
    它具有队列的属性，所以只能访问第一个元素
    底层存储用vector，用heap来组织数据结构
    //模板类型
    template <typename T, typename Container=std::vector<T>, typename Compare=std::less<T>>
    class priority_queue

// 上述 stack、queue、priority_queue不是容器，而是容器适配器

> map
    底层数据结构为红黑树，有序，不重复

> multimap
    底层数据结构为红黑树，有序，可重复

> hash_map
    底层数据结构为hash表，无序，不重复

> hash_multimap
    底层数据结构为hash表，无序，可重复

> set
    底层数据结构为红黑树，有序，不重复

> multiset
    底层数据结构为红黑树，有序，可重复

> hash_set
    底层数据结构为hash表，无序，不重复

> hash_multiset
    底层数据结构为hash表，无序，可重复 

```

### 类函数指针
```
class CA
{  
 public:  
    int caAdd(int a, int b) {return a+b;}  
    int caMinus(int a, int b){return a-b;};  
};  
//定义类函数指针类型
typedef int (CA::*PtrCaFuncTwo)(int ,int);  
//指针赋值
PtrCaFuncTwo pFunction = &CA::caAdd;  
//使用指针，注意使用括号
CA ab;  
int c = (ab.*pFunction) (1,2);  
```

### 静态类
```
静态类所必须的初始化在类外进行（不应在.h文件内实行），而前面不加static，以免与外部静态变量(对象)相混淆
```

### 参数传递(string &)
```
void func(string a){} //此处 不能是&！！！！！
func("aaaa");
```

### explicit
```
C++提供了关键字explicit，可以阻止不应该允许的经过转换构造函数进行的隐式转换的发生, 
```

### __declspec 
```
(1) __declspec(align(#))精确控制用户自定数据的对齐方式 ，#是对齐值
    它与#pragma pack()是一对兄弟，前者规定了对齐的最小值，后者规定了对齐的最大值。同时出现时，前者优先级高

(2) __declspec(deprecated)说明一个函数，类型，或别的标识符在新的版本或未来版本中不再支持，你不应该用这个函数或类型。它和#pragma deprecated作用一样。

```

### access(目录是否存在)
```
#include <unistd.h>
int access(const char * pathname, int mode)
成功执行时，返回0。失败返回-1
R_OK      测试读许可权
W_OK      测试写许可权
X_OK      测试执行许可权
F_OK      测试文件是否存在
```

### sort
```
bool compare(int & a,int & b)
int a[20]={2,4,1,23,5,76,0,43,24,65};
sort(a,a+20,compare);
```

### lambda
```
auto f = [=]()->{};
[=]   值
[&]   引用

int main()
{
    int a = 123;
    auto f = [a] { cout << a << endl; }; 
    a = 321;
    f(); // 输出：123
}
//以传值方式捕获外部变量，则在Lambda表达式函数体中不能修改该外部变量的值。
int main()
{
    int a = 123;
    auto f = [&a] { cout << a << endl; }; 
    a = 321;
    f(); // 输出：321
}

//[=]表示以值捕获的方式捕获外部变量，[&]表示以引用捕获的方式捕获外部变量
int main()
{
    int a = 123;
    auto f = [=] { cout << a << endl; };    // 值捕获
    f(); // 输出：123
}
//隐式引用捕获示例：
int main()
{
    int a = 123;
    auto f = [&] { cout << a << endl; };    // 引用捕获
    a = 321;
    f(); // 输出：321
}

// 实例
vector<int> mList(5);
for (size_t i = 0; i<mList.size();i++) {
    cout << &(mList[i]) << endl;
}
int i = 0;
for_each(mList.begin(), mList.end(), [&i](int & x){
    x = i++;
});
for_each(mList.begin(), mList.end(), [](int & x){
    cout << x << endl;
});

```

### std::sort
```
vector<xxx> a;
sort(a.begin(),a.end(),[](const xxx & x, const xxx & y){return(x>y);}); //注意const!!!

//qsort
int x[10];
qsort(x,10,sizeof(int),func);
int func(const void* a, const void* b){
	return((*(int*)a)-(*(int*)b));
}

// 注意
list不能排序，因为list的iterator不是随机的，而vector可以，因为他是随机的
```

### const char*(初始化)
c++允许先初始化再赋值
```
const char* a；
a = "dafgsfaaafag";
```

### limits
```
#include <limits>
numeric_limits<double>::max() 
numeric_limits<double>::min() 
```

### execv
```
const char* job[] = {"sh","-c","echo \'fafafafa\'",NULL};
execv("/bin/sh",(char* const*)job); //注意(char* const*),而不是(const* char*)!!!
```

### random
```
#include <random>
std::default_random_engine e;
std::uniform_real_distribution<float> u(0,100)  //随机数分布 注意此处是float
int a = u(e) ;   //产生随机数
```

### sstream
```
#include <sstream>
default_random_engine e;
uniform_real_distribution<int> u(numeric_limits<int>::min(), numeric_limits<int>::max);
stringstream ss;
ss << u(e) << ".sm";
string x(ss.str());
const char* y = ss.str().c_str();
```

### 多参数
```
void func(){} 
template<typename T, typename... Args> //T 返回 argvs参数
void func(T value, Args... args){
    cout << value << endl;
    func(args...);
}
```

### \_\_func\_\_
```
__func__
当前函数名称
```

### 函数指针map
```
void func (int a) {
    cout << a << endl;
}

// 方法一
std::map<std::string, void(*)(int)> mm;
mm["func"] = &func;
mm["func"](312);

// 方法二
typedef void(*f)(int);
std::map<std::string, f> mm;
mm["func"] = &func;
mm["func"](312);
```

### 类静态成员初始化
```
// 静态成员需要一开始初始化
template<typename T>
class SysIpc
{
public:
	SysIpc(){}
	~SysIpc(){}
	const static size_t n = 10; //如果不给n赋值，就会报错
	static T list[n];
};

template<typename T>
T SysIpc<T>::list[SysIpc<T>::n];
```

### sys/time.h
```
#include <sys/time.h>
struct timeval tv;
gettimeofday(&tv,NULL);
sprintf(tarDir,"%s/tmp%ld%ld/",tarDir,tv.tv_sec,tv.tv_usec);
```

### 子进程退出
```
// 子进程推出推荐用_exit(1)
pid_t pt;
int status;
waitpid(pt,&status,0);
WEXITSTATUS(status) //获取退出值
WIFEXITED(status)   //判断是否正常推出
WIFSIGNALED(status) //判断是否被杀死
```

### fd设置
```
(1) 设置fd非阻塞
    int val;
    if ((val = fcntl(fd, F_GETFL, 0)) < 0){
        cout << "[error]: fcntl" << endl;
        exit(-1);
    }
    val |= O_NONBLOCK;                          //先取出val,再设置
    if (fcntl(fd, F_SETFL, val) < 0){
        cout << "[error]: fcntl" << endl;
        exit(-1);
    }
```

### 拷贝构造函数
```
class Test
{
public:
    Test(){}
    ~Test(){}
    Test(const Test & t){}    // 拷贝构造函数声明
};

A a;    //构造
A b(a); //调用拷贝构造函数
A b = a; //调用拷贝构造函数
```

### 右值引用
```
(1) 值优化的重要性
    //值传递实例
    class A
    {
        A(){
            cout << "construct" << endl;
        }

        A(const A& a){
            cout << "copy construct" << endl;
        }
        ~A(){
            cout << "deconstruct" << endl;
        }
    };

    A GetA(){
        return A();
    }

    int main(){
        A a = GetA();
    }

    //结果
    // 编译选项-fno-elide-constructors用来关闭返回值优化效果
    // 关闭值优化的结果
    construct
    copy construct
    deconstruct
    copy construct
    deconstruct
    deconstruct

    //当开启值优化时
    construct
    deconstruct

    //由此可见值优化非常重要！！！

(2) 指针悬挂与深拷贝

(3) 左值引用与常量左值引用
    //左值引用
    int a;
    int& b = a; //正确
    int& b = 1; //错误 
    //注意：所有的引用都是左值引用，即右边必须是左值！！！

    //常量左值引用
    const int& a = 1; //正确
    int& b = a;       //错误 a还是右值,所以b不能引用右值

    string func(){
        return("dddada");
    }
    string& a = func(); //错误 func返回的值已被销毁 引用的必须是左值
    const string& a = func(); //正确 常量左值引用是一个“万能”的引用类型，可以接受左值、右值、常量左值和常量右值

    void func (string & a) {}
    func("dsdas");      // 错误，参数必须是左值
    void func (const string & a) {}
    func("dsdas");      // 正确，参数可以是右值

    //常量左值引用意义
    > 传引用不会新创建一个变量
    > 不会调用构造函数，构造拷贝函数，拷贝函数
    > 直接传对象，速度快，同时保证了在函数内部无法对对象进行修改

(4) 右值引用与类型推导判断
    //右值引用
    int a;
    int&& b = 1; 
    int&& b = a; //错误 a是左值，必须引用右值
    //自动判断（T&& t在发生自动类型推断的时候，它是未定的引用类型）
    template<typename T>
    void func(T&& t){}

    func(10);  //t变成右值引用
    int x = 10;
    func(x);   //t变成左值引用

    template<typename T>
    class Test {
        Test(Test&& rhs); //注意 构造函数 没有发生类型推导 rhs肯定是右值引用
    };

(5) 续命与减少拷贝
    int func(){
        return 1;
    }

    int x = func();    //需要拷贝
    int&& x = func();   //不需要拷贝 仅仅是move

    stringstream ss("dadada");
    string && ss1 = ss.str(); //str()返回的是拷贝，所以我们用右值避免拷贝
    cout << ss1 << endl;
    const char * ss2 = ss1.c_str();
    char *ss3 = const_cast<char*>(ss2);
    *ss3 = 'W';
    cout << ss1 << endl;

(6) 移动语义（移动构造函数）
    class A
    {
        int* i_ptr;
        A():i_ptr(new int(0)){
            cout << "construct" << endl;
        }
        ~A(){delete i_ptr;}
        A(const A& a):i_ptr(new int(*a.i_ptr)){ //注意 a是左值引用
            cout << "copy" << endl;
        }
        A(A&& a):i_ptr(a.i_ptr){   //移动构造函数 a是右值引用 没有做深拷贝，仅仅是将指针的所有者转移到了另外一个对象
            a.i_ptr = nullptr;      //非常重要，因为a被析构，所以要让他的内存指向一个空！！！
            cout << "move" << endl;
        }
    };
    //对于上面的右值引用（A(A&& a)），a是一个右值，那么里面的内容也是右值，所以i_ptr = a.i_ptr合法。

    A GetA(){
        return A();
    }

    A a = GetA();
    //输出
    construct
    move
    move

(7) 完美转发
    #include <utility>
    void func(string& str){}
    void func(string&& str){}
    template<typename T>
    void pfunc(T&& str){
        func(forward<T>(str)); //自动判断str是左值还是右值
    }

(8) std::move(移动转移所有权)
    #include<utility>
    vector<int>	a = {1,2,3,4,5};
    vector<int> b = move(a);
    cout << "===" << endl;
    for_each(a.begin(),a.end(),[](int& x){cout << x << endl;});
    cout << "===" << endl;
    for_each(b.begin(),b.end(),[](int& x){cout << x << endl;});
    //输出
    ===
    ===
    1
    2
    3
    4
    5
```

### RTTI
```
RTTI(Run Time Type Identification)即通过运行时类型识别，程序能够使用基类的指针或引用来检查着这些指针或引用所指的对象的实际派生类型。

RTTI提供了两个非常有用的操作符：typeid和dynamic_cast。
typeid操作符，返回指针和引用所指的实际类型；
dynamic_cast操作符，将基类类型的指针或引用安全地转换为其派生类类型的指针或引用。

如果表达式的类型是类类型且至少包含有一个虚函数，则typeid操作符返回表达式的动态类型，需要在运行时计算；否则，typeid操作符返回表达式的静态类型，在编译时就可以计算。

#include <typeinfo>
#include <typeindex>
//type_info类
//typeid操作符
//type_index类

1、type_info 类
    namespace std {
        class type_info
        {
        public:
            virtual ~type_info(); //type_info can serve as a base class
            // enable comparison
            bool operator==(const type_info& rhs ) const;
            // return !( *this == rhs)
            bool operator!=(const type_info& rhs ) const;
            bool before(const type_info& rhs ) const; // ordering
            //return a C-string containing the type's name
            const char* name() const;
        private:
            //objects of this type cannot be copied
            type_info(const type_info& rhs );
            type_info& operator=(const type_info& rhs);
        }; //type_info
    }

2、typeid 返回值
    //The result of a typeid expression is an lvalue of static type const
    // 返回值是常量对象的引用
    string a;
    const type_info & rhs = typeid(a);

    //注意：typeid既不是函数，也不是宏，它是个操作符


// 实例
int a;
const std::type_info & rhs = typeid(a);
typeid(a) == typeid(int); // true
std::cout << "tiInt.name = " << rhs.name() << std::endl;

class A{};
std::unordered_map<std::type_index, std::string> type_names;
type_names[std::type_index(typeid(A))] = "A";
type_names[std::type_index(typeid(std::string))] = "string";

static struct TYPE_NAME {
    std::map<std::type_index, std::string> type_names;
    template<typename T>
    string getType(T v) {
        return type_names[std::type_index(typeid(v))];
    }
    TYPE_NAME (){
        type_names[std::type_index(typeid(int))] = "int";
        type_names[std::type_index(typeid(std::string))] = "string";
        type_names[std::type_index(typeid(char))] = "char";
        type_names[std::type_index(typeid(const char*))] = "const char*";
    }
} TYPE;

TYPE.getType("dsadas");

```

### decltype
```
decltype和auto都可以用来推断类型，但是二者有几处明显的差异：
1.auto忽略顶层const，decltype保留顶层const；
2.对引用操作，auto推断出原有类型，decltype推断出引用；
3.对解引用操作，auto推断出原有类型，decltype推断出引用；
4.auto推断时会实际执行，decltype不会执行，只做分析
```

### 类型转换
```
//const_cast 去掉类型的const或volatile属性
//static_cast 基本数据类型转换，不能进行无关类型（如非基类和子类）指针之间的转换。把空指针转换成目标类型的空指针。把任何类型的表达式转换成void类型。static_cast不能去掉类型的const、volatile属性(用const_cast)。

//实例 const char* ==> void*
const char* a = "dadada";
char* x = const_cast<char*>(a);
void* y = static_cast<void*>(x);
```

### enum
```
// c++ 11
enum class Color {red,yellow,blue};
Color myColor = Color::red; 

// c++ 98
enum My{a=1,b,c};
My a = a;
```

### goto
```
int main(){
    goto label;

    int i = 0;   //会报错，因为上面的goto会忽略这里的变量！！！所以，将变量写在goto之外

    label:
    {
        xxxxxx;
    }
}
```

### 析构竞态
```
由于C++对象的生命周期需要程序员自己管理，因此析构可能出现竞态尤其是在多线程下，一个对象可以被多个线程访问时，下列情形：
1、即将析构一个对象时从何得知其它线程是否在操作该对象
2、若某个线程正欲操作对象时，如何得知其它线程是否在析构该对象，且正析构一半....
```

### class -> static function-> static variable
```
class中的static函数中放置此函数自己使用的变量
//例如
class A
{
public:
    static StaticFunc (){
        static std::mutex mutex;    //此变量StaticFun独有，没必要写到外面
        static size_t count = 0;
        std::unique_lock<std::mutex> lock(mutex);//加锁，析构时释放该锁
        count++;
    }
};

```

### 工厂模式
```
工厂   --->  产品
抽象工厂     抽象产品
具体工厂     具体产品
```

### 遍历对象
```

```

### 序列化
```
```

### 反射
```
// 反射是程序获取自身信息的能力

// 作用
    可以用于动态创建类型，跨语言跨平台数据交互，持久化，序列化等等。
    包含以下功能：
        枚举所有member
        获取member的name和type
        能够get/set　member

// c++实现方法
    运行期支持
    宏

//了解yuanzhibi的实现方式 https://github.com/yuanzhubi/reflect_struct/blob/master/test.cpp

```

### 容器存放不同类型变量
```
https://www.zhihu.com/question/33594512?sort=created

// 相关库 std::any std::variant
    原理: 底层是union和void*实现，union存储基础类型，里面的void*存储自定义类型，再加一个type字段存储类型编号即可
    type ---> typeid ---> copyConstruct[typeid] ---> new type()

    > 注册类型 copyConstruct[typeid] = TypeCopyConstruct
    > 赋值 ptr = new copyConstruct[typeid](value)
    > 取值switch typeid type return value
```

### operator[]
```
this->operator[](3)
```

### POD 类型
```
https://zhuanlan.zhihu.com/p/45545035

平凡特征的数据类型
> 所有基本数据类型(基本类型和指针类型)
> 一个class或者struct，它不包含虚函数，没有虚基类，每一个数据成员都是POD，且所有的父类（如果存在的话）都是POD  
> POD数组  
> 由POD组成的union  
```

### union也可以这么用
```
//可以在union里面定义类型，也可以定义构造函数

//注意：union cannot define non-POD as member data, 对于这种情况直接用指针得了

union PopupInfo
{
	struct _s1 { NativePoint location; INativeScreen* screen; };
	struct _s2 { GuiControl* control; INativeWindow* controlWindow; Rect bounds; bool preferredTopBottomSide; };
	struct _s3 { GuiControl* control; INativeWindow* controlWindow; Point location; };
	struct _s4 { GuiControl* control; INativeWindow* controlWindow; bool preferredTopBottomSide; };

	_s1 _1;
	_s2 _2;
	_s3 _3;
	_s4 _4;

	PopupInfo() {}
};

union My {
    vector<int>* a; // 指针是POD
    map<int,int>* b; // 指针是POD
    vector<int> c;  //错误！non-POD 类型
}
```

### struct 的位域
```
C语言又提供了一种数据结构，称为“位域”或“位段”。所谓“位域”是把一个字节中的二进位划分为几个不同的区域，并说明每个区域的位数

struct bs { 
    int a:8; 
    int b:2; 
    int c:6; 
} data;
 
data为struct bs的变量，其中位域a占8位，位域b占2位，位域c占6位
```

### 构造、移动构造、移动赋值、拷贝构造、拷贝赋值
```
template<typename T>
class Object {
public:
    Object();                                       //默认构造

    virtual ~Object(){}                             //析构

    Object(const Object & ob);                      //拷贝构造

    template<typename T1>
    Object(const Object<T1> &);                     //泛化的拷贝构造

    Object(Object && ob);                           //移动构造

    Object<T> & operator=(const Object<T> & ob)     //拷贝赋值
    template<typename T1>
    Object<T> & operator=(const Object<T1> & ob)    //泛化拷贝

    Object<T> & operator=(const Object<T> && ob)    //移动赋值

public:
    Object create(){}                               //新建    
    void clean(){}                                  //清除
    void swap(Object & ob){}                        //交换
};


Object<int> a = Object<int>::Object();  // 移动构造函数
Object<int> a;
a = 


```

### 可变参数实例
```
class DetailAnno {
public:
    static vector<string> spl;
    template<typename T, typename... Argv>
    static void Init(const T str_, Argv... argvs);
    static void Init(){}

public:
    DetailAnno(){}
    virtual ~DetailAnno(){}
    DetailAnno(const DetailAnno & d) = default;
    DetailAnno & operator=(const DetailAnno & d) = default;
};

vector<string> DetailAnno::spl;

template<typename... Argv>
void DetailAnno::Init<string, Argv...>(const string str_, Argv... argvs){
    DetailAnno::spl.push_back(str_);
    DetailAnno::Init(argvs...);
}
```

### 模板偏特化(指针)
```
1、模板函数
    template<typename T> void func(T* t) {
        cout << "T*" << endl;
    }

    template<typename T> void func(T t) {
        cout << "T" << endl;
    }

    int a = 2;
    func(a);      // T
    func(&a);     // T*

2、模板类
    template<typename T>
    class a{
    public:
        a(){
            cout << "T" << endl;
        }
        virtual ~a(){}
    };

    template<typename T>
    class a<T*>{
    public:
        a(){
            cout << "T*" << endl;
        }
        virtual ~a(){}
    };

    a<char> x; 
    a<char*> x1;
```

### 类模板中的模板函数
```
template<typename T>
class My{
public:
    typedef T type_name;    
private:
    T t;   
    My(){} 
public:
    My(T t_):t(t_){}
    virtual ~My(){}
    My(const My<T> & m) = default;
    template<typename T1>               //注意
    My(const My<T1> & m);
    T & operator=(const My<T> m) = default;
};

template<typename T>
template<typename T1>
My<T>::My(const My<T1> & m){
    t = static_cast<T>(m.t);
}

```

### Traits Classes 
```
// 在 C++ 中，traits 习惯上总是被实现为 struct ，但它们往往被称为 traits classes。Traits classes 的作用主要是用来为使用者提供类型信息。

//  STL 中，容器与算法是分开的，容器与算法之间通过迭代器联系在一起

https://www.cnblogs.com/mangoyuan/p/6446046.html
https://cloud.tencent.com/info/a180e28f80b999eb22700e2407fc0957.html
https://blog.csdn.net/lihao21/article/details/55043881

//函数的“template参数推导机制”推导的只是参数，无法推导函数的返回值类型

//示例
template <typename T>
struct iterator_traits {
    typedef typename T::value_type value_type;
};

template <typename T>
struct iterator_traits<T*> {
    typedef T value_type;
};

template <typename T> typename iterator_traits<T>::value_type
func(T a) {
    return *a;
}

// 我们就可以知道模板类和指针的类型信息
```

### value_type
```
//对于大部分STL都适用

vector<int>::iterator it = a.begin();
vector<int>::iterator::value_type a = 1;
等价于
int a = 1;

//萃取
class vector {
public:
    class iterator {
    public:
        typedef T value_type;
        ...
    };
...
};
```

### stateful ！！！
```
//编译时确定count数

(1) 变量的构造函数
    int a(2);
    double b(2.22);
    int *a(new int [2]);   //a是2个int的指针
    int (*a())[3];         //这是个声明，所以不占内存，声明了一个包含3个int的函数指针
    sizeof(*a());          //输出12

    //同理
    class Test
    {
    public:
        int a,b,c;    
    };
    Test (*x()) [3];
    cout << sizeof(x()) << endl;    // 8
    cout << sizeof(*x()) << endl;   // 36
    

(2) 编译期计数
    template<unsigned int N>
    struct struct_int : struct_int<N - 1> {};
    template<>
    struct struct_int<0> {};
    #define MAX_COUNT 168 // you can increase the number if your compiler affordable
    
    #define EVAL_COUNTER(counter) (sizeof(*counter((struct_int<MAX_COUNT>*)0)) \
          - sizeof(*counter((void*)0)))
    //We can change the result of EVAL_COUNTER if we use INCREASE_COUNTER or SET_COUNTER

    #define INCREASE_COUNTER(counter, delta)  char (*counter(struct_int<EVAL_COUNTER(counter) + 1>*))[EVAL_COUNTER(counter) + sizeof(*counter((void*)0)) + (delta)]; 

    #define SET_COUNTER(counter, value)  char (*counter(struct_int<EVAL_COUNTER(counter) + 1>*))[value + sizeof(*counter((void*)0))]; 

    #include <stdio.h>
    int main(){
        char (*first_counter(...))[1];  // It declares a function. No space cost.
        char (*second_counter(...))[1]; // It declares a function. No space cost.
    
        //For all the counter, the init value must be zero.
        static const unsigned int i1 = EVAL_COUNTER(first_counter); //i1=0
        INCREASE_COUNTER(first_counter, 2);
        static const unsigned int i2 = EVAL_COUNTER(first_counter); //i2=0+2
        INCREASE_COUNTER(first_counter, 1);
        static const unsigned int i3 = EVAL_COUNTER(first_counter); //i3=2+1
        //INCREASE_COUNTER(first_counter, -1);  negative increase is not enabled
        SET_COUNTER(first_counter, 6);
        static const unsigned int i4 = EVAL_COUNTER(first_counter); //i4=6
        //SET_COUNTER(first_counter, 6);  we can not set counter to number that not greater than its max
        
        //For all the counter, the init value must be zero.
        static const unsigned int j1 = EVAL_COUNTER(second_counter); //j1=0
        INCREASE_COUNTER(second_counter, 2);
        static const unsigned int j2 = EVAL_COUNTER(second_counter); //j2=0+2
        INCREASE_COUNTER(second_counter, 1);
        static const unsigned int j3 = EVAL_COUNTER(second_counter); //j3=2+1
            
        printf("%u%u%u%u\n%u%u%u\n", i1, i2, i3, i4 ,j1, j2, j3 );
        return 0;
    }

(3) 自己的理解
    template<int N>
    struct counter:counter<N-1>{};
    template<>
    struct counter<0>{};

    char (*f(...))[1];
    cout << sizeof(*f((void*)0)) << endl;
    char (*f(counter<2>*))[2];
    cout << sizeof(*f((counter<MAX>*)0)) << endl;
    char (*f(counter<3>*))[3];
    cout << sizeof(*f((counter<MAX>*)0)) << endl;

```

### struct初始化
```
c/c++中的struct
//无序
struct MyS a = {
    .a = 1,
    .b = NULL
};

//有序
struct TestS {
    int a;
    std::string b;
    char c;
};

struct TestS a = {1, "dasda", 'c'};
struct TestS a {1, "dasda", 'c'};
```

### {}一个用法
```
#define var(x) (*({     \
    &x;}))

int a = 0;

var(a)++;
int* b = var(a);


#define var(x) ({x;})
int a = 0;
var(a)++;         // a++
int b = a;        // b = a

```

### constexpr
```
constexpr表示这玩意儿在编译期就可以算出来（前提是为了算出它所依赖的东西也是在编译期可以算出来的）

const只保证了运行时不直接被修改（但这个东西仍然可能是个动态变量）
```

### __attribute__
```
(1) 概念
    GNU C 的一大特色就是__attribute__ 机制。__attribute__ 可以设置函数属性（Function Attribute ）、变量属性（Variable Attribute ）和类型属性（Type Attribute ）
    __attribute__ 书写特征是：__attribute__ 前后都有两个下划线，并切后面会紧跟一对原括弧，括弧里面是相应的__attribute__ 参数。
    __attribute__ 语法格式为：__attribute__ ((attribute-list))


(2) 函数属性(Function Attribute)
    __attribute__((noreturn))    表示没有返回值
    __attribute__((unused))   表示该函数或变量可能不使用，这个属性可以避免编译器产生警告信息 
    void __attribute__((noreturn)) handle_signal(int __attribute__((unused)) signal) {
        exit(0);
    }


(3) 类型属性(Type Attributes)


(4) 变量属性(Variable Attribute)
    __bitwise	__attribute__((bitwise))
        确保变量是相同的位方式(比如 bit-endian, little-endiandeng)
    __user	__attribute__((noderef, address_space(1)))
        指针地址必须在用户地址空间
    __kernel	__attribute__((noderef, address_space(0)))
        指针地址必须在内核地址空间
    __iomem	__attribute__((noderef, address_space(2)))
        指针地址必须在设备地址空间
    __safe	__attribute__((safe))
        变量可以为空
    __force	__attribute__((force))
        变量可以进行强制转换
    __nocast	__attribute__((nocast))
        参数类型与实际参数类型必须一致
    __acquires(x)	__attribute__((context(x, 0, 1)))	
        参数x 在执行前引用计数必须是0,执行后,引用计数必须为1
    __releases(x)	__attribute__((context(x, 1, 0)))	
        与__acquires(x)相反
    __acquire(x)	__context__(x, 1)
        参数x的引用计数+1
    __release(x)	__context__(x, 1)
        与__acquire(x)相反
    __cond_lock(x,c)	((c) ? ({ __acquire(x); 1; }) : 0)	
        参数c 不为0时,引用计数 + 1, 并返回1
    __rcu    __attribute__((noderef, address_space(4))) 
        即这个变量地址必须是有效的，而且变量所在的地址空间必须是 4，即 RCU 空间的。
        使用__rcu 附上 RCU保护的数据结构，如果你没有使用rcu_dereference()类中某个函数，Sparse就会警告你这个操作。
```

### basic_string
```
https://www.byvoid.com/zhs/blog/cpp-string

string并不是一个单独的容器，只是basic_string 模板类的一个typedef 而已

extern "C++" {
    typedef basic_string<char> string;
    typedef basic_string<wchar_t> wstring;
}

// 类basic_string
template <class charT, 
          class traits = char_traits<charT>,
          class Allocator = allocator<charT> >
class basic_string
{
//...
}
```

### freopen
```
#include <stdio.h>

//把一个新的文件名 filename 与给定的打开的流 stream 关联
FILE *freopen(const char *filename, const char *mode, FILE *stream)
//mode 
    "r"	打开一个用于读取的文件。该文件必须存在。
    "w"	创建一个用于写入的空文件。如果文件名称与已存在的文件相同，则会删除已有文件的内容，文件被视为一个新的空文件。
    "a"	追加到一个文件。写操作向文件末尾追加数据。如果文件不存在，则创建文件。
    "r+"	打开一个用于更新的文件，可读取也可写入。该文件必须存在。
    "w+"	创建一个用于读写的空文件。
    "a+"	打开一个用于读取和追加的文件。

FILE *fp, *fp1;
fp1 = freopen("file1.txt", "r", stdin); // 输入重定向，输入数据将从file1.txt文件中读取 
fp = freopen("file.txt", "w+", stdout); // /输出重定向，输出数据将保存在file.txt文件中 
  
```

### wait waitpid
```
#include <sys/wait.h>

pid_t wait(int *statloc);
pid_t waitpid(pid_t pid,int *statloc, int options);
//statloc指向终止进程的终止状态，如果不关心终止状态可指定为空指针
//pid有四种情况：
//1.pid==-1 等待任意子进程
//2.pid>0 等待进程ID与pid相等的子进程
//3.pid==0 等待组ID等于调用进程组ID的任意子进程
//4.pid<-1 等待组ID等于pid绝对值的任意子进程
pid_t wait(int *statloc)
{
    return waitpid(-1, statloc, 0);
}
```

### map 判断键值是否存在
```
(1) insert
    std::map<int, string> myMap;
    ...
    std::pair<std::map<int, string>::iterator, bool> findItem;  
    findItem = myMap.insert(std::pair<int, string>(1, "student_one"));  
    if(findItem.second == true) {
        // myMap 中没有insert的数据，说明插入成功
        cout<<"Insert Successfully"<<endl;
    } else {
        // myMap 中有insert的数据，说明插入失败
    }
    // 打印数据
    cout << findItem.first->first << endl;
    cout << findItem.first->second << endl;
        
(2) find
    map<>::iterator it = m.find('key');
    if (it == m.end()) #不存在
    
    
// 插入键值对的方法
(1) map[键] = 值；直接赋值。 这种方式：当要插入的键存在时，会覆盖键对应的原来的值。如果键不存在，则添加一组键值对。
(2) map.insert()；这是map自带的插入功能。如果键存在的话，则插入失败，也就是不插入。 使用insert()函数，需要将键值对组成一组才可以插入
```

### map 键值排序
```
struct cmpkeylen{
    bool operator()(const string & a, const string & b){
        return(a.length() > b.length()); // 按字符串长度排序
    }
};

map<sting,int,cmpkeylen> mymap;
//注意 第三个参数是个函数对象，c++ 11中很多库函数都是函数对象
```

### openmp
```
g++ -fopenmp
omp是多线程支持，只需定义即可走多线程
定义环境变量export OMP_NUM_THREADS=10

#pragma omp parallel for   //遇到for就开多线程
for(){
}

void main()  
{  
#pragma omp parallel    //Initialization走的多线程，
    {  
        Initialization();  
#pragma omp barrier  //设置屏障 等待当前线程完毕
        printf("i=%d, thread_id=%d\n", sum, omp_get_thread_num());  
    }  
    system("pause");  
}  
```

### iterator
```
vector<int> a;
vector<int>::iterator i;
for(i=a.begin();i!=a.end();i++){}
const vector<int> a;
vector<int>::const_iterator i;  //对于const 必须用const_iterator !!!
for(i=a.begin();i!=a.end();i++){}
```

### const 对象只能访问 const 函数
```
// C++中，const 修饰的参数引用的对象，只能访问该对象的const函数，因为调用其他函数有可能会修改该对象的成员，
// 编译器为了避免该类事情发生，会认为调用非const函数是错误的。
struct Base
{
    Base() { std::cout << "  Base::Base()\n"; }
    virtual ~Base() { std::cout << "  Base::~Base()\n"; }
    virtual void test() { std::cout<< " test in base\n"; } <-------加上 const
};

void MyTest(const Base& b)
{
    b.test();
}

// Base中声明test时加上const，即void test() const

// 注意当函数后面加了const时，返回引用和指针时要加const，但是返回非引用可以不加
const mat & func()const{...}
const mat * func()const{...}
mat func()const{...}
```

### initializer_list
```
void g(std::vector<int> const &items){}; 
void g(std::list<int> const &items){};
g({ 1, 2, 3, 4 }); //会报错  编译器分不清 vector还是list

//对于{}的固定数组initializer_list更合适
void g(std::vector<int> const &items){}; 
void g(std::list<int> const &items){}; 
void g(std::initializer_list<int> const &items){}; //注意const
g({ 1, 2, 3, 4 });
initializer_list不能修改，更符合参数的特点
```

### 函数对象
```
#include <functional>
template<typename T, typename... Args> //T 返回 argvs参数
static void forkRun(function<T(Args...)> func, Args... args);
```

### std::function
```
#include <functional>

// 普通函数
int add(int a, int b){return a+b;} 

// lambda表达式
auto mod = [](int a, int b){ return a % b;}

// 函数对象类
struct divide{
    int operator()(int denominator, int divisor){
        return denominator/divisor;
    }
};

std::function<int(int ,int)>  a = add; 
std::function<int(int ,int)>  b = mod ; 
std::function<int(int ,int)>  c = divide(); //divide类构造

```

### std::bind
```
// 绑定普通函数
double my_divide (double x, double y) {return x/y;}
std::function<double(double,double)> fn_half = std::bind (my_divide,std::placeholders::_1,2);  // placeholders::_1 占位符
std::cout << fn_half(10) << '\n';     // 2

// 绑定一个成员函数
struct Foo {
    void print_sum(int n1, int n2)
    {
        std::cout << n1+n2 << '\n';
    }
    int data = 10;
};
int main() 
{
    Foo foo;
    //bind绑定类成员函数时，第一个参数表示对象的成员函数的指针，第二个参数表示对象的地址
    auto f = std::bind(&Foo::print_sum, &foo, 95, std::placeholders::_1);
    f(5); // 100
}

```

### std::pair
```
std::pair 是一个结构体模板，其可于一个单元存储两个相异对象
template<
    class T1,
    class T2
> struct pair;


// 构造
std::pair<int, string> a(1, "dasdssa");
std::pair<int, string> a;
a = make_pair(1, "dsadada");  // make_pair 是非成员函数

// 成员
std::pair<int, string> a(1, "dasas");
cout << a.first << endl;
cout << a.end << endl;
a.first = 2;
a.second = "dasds";

// 类型
std::pair<int, string> a(1, "dasas");
// 注意 class 中的 typedef 不能用实例调用
a.first_type x = 1；    // 错误
std::pair<int, string>::first_type x = 1 //正确

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
    shared_ptr<string> a(new string("dasdas"));
    shared_ptr<string> a = make_shared<string>("dasdas")
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

### 异常
```
// 异常处理机制
    其基本思想是：函数 A 在执行过程中发现异常时可以不加处理，而只是"拋出一个异常"给 A 的调用者，假定为函数 B
    拋出异常而不加处理会导致函数 A 立即中止，在这种情况下，函数 B 可以选择捕获 A 拋出的异常进行处理，也可以选择置之不理。如果置之不理，这个异常就会被拋给 B 的调用者，以此类推
    如果一层层的函数都不处理异常，异常最终会被拋给最外层的 main 函数。main 函数应该处理异常。如果main函数也不处理异常，那么程序就会立即异常地中止

// try {} catch {}
    try {
        if (n == 0)
            throw -1;  //抛出整型异常
        else if (m == 0)
            throw -1.0;  //拋出 double 型异常
        else
            cout << m / n << endl;
        cout << "after dividing." << endl;
    } catch (double d) {
        cout << "catch (double)" << d << endl;
    } catch (...) {
        cout << "catch (...)" << endl;
    }

// 函数的异常声明
    void func() throw(); // 声明不会抛出异常
    void func() noexcept; // c++11 声明不会抛出异常
    void func() throw(int, double); // 声明抛出(int, double)类型的异常



// <exception>
    (1) 标准异常架构
        std::exception	该异常是所有标准 C++ 异常的父类
            std::bad_alloc	该异常可以通过 new 抛出
            std::bad_cast	该异常可以通过 dynamic_cast 抛出
            std::bad_exception	这在处理 C++ 程序中无法预期的异常时非常有用
            std::bad_typeid	该异常可以通过 typeid 抛出
            std::logic_error	理论上可以通过读取代码来检测到的异常
                std::domain_error	当使用了一个无效的数学域时，会抛出该异常
                std::invalid_argument	当使用了无效的参数时，会抛出该异常
                std::length_error	当创建了太长的 std::string 时，会抛出该异常
                std::out_of_range	该异常可以通过方法抛出，例如 std::vector 和 std::bitset<>::operator[]()
            std::runtime_error	理论上不可以通过读取代码来检测到的异常
                std::overflow_error	当发生数学上溢时，会抛出该异常
                std::range_error	当尝试存储超出范围的值时，会抛出该异常
                std::underflow_error	当发生数学下溢时，会抛出该异常

    (2) exception应用
        try {
            char * p = new char[0x7fffffff];  //无法分配这么多空间，会抛出异常
        } catch (std::bad_alloc & e)  {
            cerr << e.what() << endl;
        }

        try {
            int a[10] {0};
            a[10] = 100;  //拋出 out_of_range 异常
        } catch (out_of_range & e) {
            cerr << e.what() << endl;
        }

        try {
            ...
            throw(std::logic_error("logic error"));
        } catch(std::logic_error & e) {
            cerr << e.what() << endl;
        }

    (3) 继承exception
        class MyException : public exception{
            const char *what() const throw() {
                return "it is my exception";
            }
        };

        try {
            throw MyException();
        }catch(MyException & e) {
            cout << e.what() << endl;
        }catch(std::exception & e) {
            // 其他错误
        }
```

### pthread
```
// 互斥锁
pthread_mutex_t mutex;
pthread_mutex_init()  
pthread_mutex_lock()    锁定互斥锁，如果尝试锁定已经被上锁的互斥锁则阻塞至可用为止
pthread_mutex_unlock() 	释放互斥锁
pthread_mutex_destory() 互斥锁销毁函数

// pthread_exit()和return类似就是推出的作用，不涉及资源释放

1、pthread_create
    // 注意函数定义: void* (*)(void*)
    void* test(void *ptr){
        cout << "hello world." << endl;
    }
    pthread_t tid;
    pthread_create(&tid, NULL, test, NULL);
    pthread_join(tid, NULL);

2、pthread_join
    // 子线程合入主线程，主线程阻塞等待子线程结束，然后回收子线程资源
    int pthread_join(pthread_t thread, void **retval)
    // 第一个参数为线程标识符，即线程ID
    // 第二个参数retval为用户定义的指针，用来存储线程的返回值

4、pthread_detach
    // 主线程与子线程分离，子线程结束后，资源自动回收
    // pthread有两种状态joinable状态和unjoinable状态
    // joinable: 当线程函数自己返回退出时或pthread_exit时都不会释放线程所占用的资源，只有当你调用了pthread_join之后这些资源才会被释放
    // unjoinable: 这些资源在线程函数退出时或pthread_exit时自动会被释放
    
    int pthread_detach(pthread_t pid)

5、pthread_self
    获取当前线程id

6、pthread_once
    // pthread_once在多线程环境中只执行一次
    int pthread_once(pthread_once_t *once_control, void (*init_routine)(void))
    // 第一个参数为pthread_once_t变量
    // 第二个参数为无参数函数指针，type: void(*func)(void)
    
    pthread_once_t once = PTHREAD_ONCE_INIT;

    void *func_once(){
        cout << "func once" << endl;
    }
    void *func1(void *arg){
        func_once(&once, func_once);
    }
    void *func2(void *arg){
        func_once(&once, func_once);
    }
    pthread_t td1, td2;
    pthread_create(&td1, NULL, func1, NULL);
    pthread_create(&td2, NULL, func2, NULL);
    // 结果只会输出一次func once

7、线程取消
    // 线程取消的方法是向目标线程发Cancel信号
    // 但如何处理Cancel信号则由目标线程自己决定，或者忽略、或者立即终止、或者继续运行至Cancelation-point(取消点)

    // 取消点
    > 通过pthread_testcancel调用以编程方式建立线程取消点
    > 线程等待pthread_cond_wait或pthread_cond_timewait()中的特定条件
    > 被sigwait(2)阻塞的函数
    > 一些标准的库调用。通常，这些调用包括线程可基于阻塞的函数???

    (1) pthread_cancle
        int pthread_cancel(pthread_t thread)
        // pthread_cancel调用并不等待线程终止，它只提出请求
        // 发送终止信号给thread线程，如果成功则返回0，否则为非0值。发送成功并不意味着thread会终止

    (2) pthread_setcancelstate
        int pthread_setcancelstate(int state, int *oldstate)
        // 设置本线程对Cancel信号的反应
        // state有两种值：PTHREAD_CANCEL_ENABLE（缺省）和PTHREAD_CANCEL_DISABLE
        // old_state如果不为 NULL则存入原来的Cancel状态以便恢复

    (3) pthread_setcanceltype
        int pthread_setcanceltype(int type, int *oldtype) 
        // 设置本线程取消动作的执行时机
        // type有两种取值：PTHREAD_CANCEL_DEFFERED 和 PTHREAD_CANCEL_ASYCHRONOUS
        // 仅当Cancel状态为Enable时有效，分别表示收到信号后继续运行至下一个取消点再退出和立即执行取消动作(退出)
        // oldtype如果不为NULL则存入运来的取消动作类型值
    
    (4) pthread_testcancel
        void pthread_testcancel(void) 
        // 手动创建一个取消点，检查本线程是否处于Cancel状态，如果是，则进行取消动作(退出)，否则直接返回

8、线程终止的清理
    https://blog.csdn.net/caianye/article/details/5912172
    // 线程终止有两种情况：正常终止和非正常终止
    // 需要注意线程退出时的锁资源的清除

    void pthread_cleanup_push(void (*routine) (void  *),  void *arg)
    void pthread_cleanup_pop(int execute)
    // pthread_cleanup_push()/pthread_cleanup_pop()采用先入后出的栈结构管理
    // 多次对pthread_cleanup_push()的调用将在清理函数栈中形成一个函数链，在执行该函数链时按照压栈的相反顺序弹出
    // execute参数表示执行到pthread_cleanup_pop()时是否在弹出清理函数的同时执行该函数，为0表示不执行，非0为执行，这个参数并不影响异常终止时清理函数的执行
    
    // 宏的表现形式
    #define pthread_cleanup_push(routine,arg)                                     
    { struct _pthread_cleanup_buffer _buffer;                                   
        _pthread_cleanup_push (&_buffer, (routine), (arg));
    #define pthread_cleanup_pop(execute)                                          
        _pthread_cleanup_pop (&_buffer, (execute)); }

    // 实例
    // 当线程在/*do some work*/终止时，将主动调用pthread_mutex_unlock(&mutex)
    void *func(void *arg){
        pthread_cleanup_push(pthread_mutex_unlock, (void*)&mutex);
        thread_mutex_lock(&mutex);
        /*do some work*/
        pthread_mutex_unlock(&mutex);
        pthread_cleanup_pop(0);
        pthread_exit(NULL);
    }

9、pthread_kill
    int pthread_kill(pthread_t thread, int sig);
    // 向线程发送signal，如果线程的代码内不做任何信号处理，则会按照信号默认的行为影响整个进程
    // 也就是说，如果你给一个线程发送了SIGQUIT，但线程却没有实现signal处理函数，则整个进程退出
    // 注意子线程信号共享父进程，所以会影响整个进程

    pthread_kill(ptd, 0)
    // 如果int sig的参数是0呢，这是一个保留信号，一个作用就是用来判断线程是不是还活着
    // 返回值0，线程仍然活着
    // 返回值ESRCH，线程已不存在
    // 返回值EINVAL，信号不合法


10、条件变量
    (1) 条件变量创建
        // 静态创建
        pthread_cond_t cond=PTHREAD_COND_INITIALIZER;
        // 动态创建
        pthread_cond _t cond;
        pthread_cond_init(&cond,NULL);
        其中的第二个参数NULL表示条件变量的属性，虽然POSIX中定义了条件变量的属性，但在LinuxThread中并没有实现，因此常常忽略

    (2) 条件等待
        pthread_mutex_t mutex=PTHREAD_MUTEX_INITIALIZER;
        pthread_mutex_lock(&mutex);
        while(条件1成立)
            pthread_cond_wait(&cond,&mutex);
            ...
        pthread_mutex_unlock(&mutex);
        当'条件1'成立的时候，执行pthread_cond_wait(&cond,&mutex)，获得互斥锁，然后线程被挂起

    (3) 条件激发
        pthread_mutex_lock(&mutex);
        if(条件1不成立)
            pthread_cond_signal(&cond);
        pthread_mutex_unlock(&mutex);
        条件1不成立的时候,执行pthread_cond_signal(&cond)，激发条件变量cond，使得被挂起的线程被唤醒

        pthread_cond_broadcast(&cond1)
        // 唤醒所有正在pthread_cond_wait(&cond1,&mutex1)的线程
        pthread_cond_signal(&cond1)
        // 唤醒所有正在pthread_cond_wait(&cond1,&mutex1)的至少一个线程

        // 可能存在的情况
        > 多个线程等待同一个cond,并且想对同一个mutex加锁
            > 当使用broadcast方式时
                > 两个被阻塞的线程都被唤醒了，被唤醒的线程将变为pthread_mutex_lock(mutex)的状态，他们将抢着对mutex加锁
                > 在本次运行过程中thread_1加锁成功了，thread_2没有成功抢到锁，于是它就被阻塞了在thread_1执行完毕释放锁后，会通知所有被阻塞在mutex1上的线程，于是thread_2最终成功拿到了锁，然后顺利执行
            > 当使用signal方式时
                > thread_1和thread_2中只被唤醒了一个线程，在本次运行中是thread_1被唤醒了，而因为thread_2没有被唤醒，他就一直卡在pthread_cond_wait处呼呼大睡，所以最终只有thread_1执行完毕

        > 多个线程等待同一个cond，并且分别不同的mutex加锁
            > 使用broadcast方式时
                因为两个线程都被唤醒了，且它们想要加的锁并没有竞争关系，因此它们是并发执行的，而不必像前一种情况中那样必须一前一后执行。
            > 当使用signal方式时，只被唤醒了一个线程，因此只有一个线程成功执行

    (4) 条件变量的销毁
        pthread_cond_destroy(&cond);
        在linux中，由于条件变量不占用任何资源，所以，这句话除了检查有没有等待条件变量cond的线程外，不做任何操作

11、pthread_cond_timedwait
    pthread_cond_timedwait(pthread_cond_t * _cond,pthread_mutex_t * _mutex,_const struct timespec * _abstime);
    //比函数pthread_cond_wait()多了一个时间参数，经历abstime段时间后，即使条件变量不满足，阻塞也被解除

```

### thread相关
```
https://www.cnblogs.com/haippy/p/3284540.html

// 头文件
    C++11 新标准中引入了四个头文件来支持多线程编程，分别是
    (1) <atomic>
        该头文主要声明了两个类, std::atomic 和 std::atomic_flag，另外还声明了一套 C 风格的原子类型和与 C 兼容的原子操作的函数
    (2) <thread>
        该头文件主要声明了 std::thread 类
    (3) <mutex>
        该头文件主要声明了与互斥量(mutex)相关的类，包括 std::mutex 系列类
    (4) <condition_variable>
        该头文件主要声明了与条件变量相关的类
    (5) <future>
        该头文件主要声明了 std::promise, std::package_task 两个 Provider 类
    (6) <semaphore.h>
        信号量

1、<atomic>
    最常见的同步机制就是std::mutex和std::atomic。然而，从性能角度看，通常使用std::atomic会获得更好的性能

    std::atomic 是模板类，一个模板类型为 T 的原子对象中封装了一个类型为 T 的值
    template<typename T>struct atomic;

    C++11标准库std::atomic提供了针对"整形"和"指针类型"的特化实现

    (1) 构造
        atomic() noexcept = default;
        constexpr atomic (T val) noexcept;
        
        // 实例
        std::atomic<bool> ready(false);
        std::atomic<int> foo = 0;

    (2) is_lock_free
        bool is_lock_free() const noexcept;
        // 判断该 std::atomic 对象是否具备 lock-free 的特性
        // 如果某个对象满足 lock-free 特性，在多个线程访问该对象时不会导致线程阻塞

    (3) store
        void store(T val, memory_order sync = memory_order_seq_cst) noexcept;
        // 具体见memory order

    (4) load
        T load(memory_order sync = memory_order_seq_cst) const noexcept;
        // 具体见memory order

2、<thread>
    (1) std::this_thread
        // 当前线程休眠一段时间，休眠期间不与其他线程竞争CPU，根据线程需求，等待若干时间
        std::this_thread::sleep_for(std::chrono::seconds(n))

        // 当前线程放弃执行(让出时间片)，操作系统调度另一线程继续执行
        // 即当前线程将未使用完的"CPU时间片"让给其他线程使用，等其他线程使用完后再与其他线程一起竞争"CPU"
        std::this_thread::yield()

        // 获取当前线程的id
        std::this_thread::get_id()

    (2) std::thread
        1) std::thread构造
            > default
                // 默认构造函数，创建一个空的 thread 执行对象
                thread() noexcept;
            > initialization
                // 初始化构造函数
                template <class Fn, class... Args>
                explicit thread (Fn && fn, Args&&... args);
            > copy
                // 拷贝构造函数(被禁用)，不允许拷贝构造
                thread (const thread &) = delete;
            > move
                // move 构造函数
                thread (thread && x) noexcept; // x是右值
        2) 其他成员
            get_id
                获取线程 ID
            joinable
                检查线程是否可被 join
            join
                Join 线程
            detach
                Detach 线程
            swap
                Swap 线程 
            native_handle
                返回 native handle
            hardware_concurrency [static]
                检测硬件并发特性
        
        3) 成员函数作为线程函数
            一般的类成员函数是不能用作回调函数的，因为库函数在使用回调函数时，都会传递指定的符合回调函数声明的的参数给回调函数，而类成员函数隐式包含一个this指针参数，所以把类成员函数当作回调函数编译时因为参数不匹配会出错

            1> 方法一
                把成员函数设成静态成员函数，不属于某个对象，属于整个类
            2> 方法二
                class A{
                public:
                    int a;

                    void member (int b){
                        cout << a << endl;
                        cout << b << endl;
                    }

                    void start () {
                        thread(&A::member, this, 2);
                    }
                }

2、<mutex>
    (1) 系列类(四种)
        1) std::mutex，最基本的 Mutex 类
            1> 构造函数
                std::mutex不允许拷贝构造，也不允许 move 拷贝，最初产生的 mutex 对象是处于 unlocked 状态的
            2> 成员函数
                > lock()
                    调用线程将锁住该互斥量
                    线程调用该函数会发生下面3种情况
                    1> 如果该互斥量当前没有被锁住，则调用线程将该互斥量锁住，直到调用 unlock之前，该线程一直拥有该锁
                    2> 如果当前互斥量被"其他线程锁住"，则当前的调用线程被阻塞住
                    3> 如果当前互斥量被"当前调用线程锁住"，则会产生死锁(deadlock)
                
                > unlock()
                    解锁，释放对互斥量的所有权
                
                > try_lock()
                    尝试锁住互斥量，如果互斥量被其他线程占有，则当前线程也"不会被阻塞"
                    线程调用该函数也会出现下面3种情况
                    1> 如果当前互斥量没有被其他线程占有，则该线程锁住互斥量，直到该线程调用 unlock 释放互斥量
                    2> 如果当前互斥量被其他线程锁住，则当前调用线程返回 false，而并不会被阻塞掉
                    3> 如果当前互斥量被当前调用线程锁住，则会产生死锁(deadlock)

        2) std::recursive_mutex，递归 Mutex 类
            和std::mutex不同的是，std::recursive_mutex允许同一个线程对互斥量多次上锁(即递归上锁)，来获得对互斥量对象的多层所有权
            
            std::recursive_mutex 释放互斥量时需要调用与该锁层次深度相同次数的 unlock()，可理解为 lock() 次数和 unlock() 次数相同，除此之外，std::recursive_mutex 的特性和 std::mutex 大致相同

        3) std::time_mutex，定时 Mutex 类
            比std::mutex多了两个成员函数，try_lock_for()和try_lock_until()
            > try_lock_for
                接受一个"时间范围"，表示在这一段时间范围之内线程如果没有获得锁则被阻塞住
                如果在此期间其他线程释放了锁，则该线程可以获得对互斥量的锁
                如果超时(即在指定时间内还是没有获得锁)，则返回 false
            > try_lock_until
                函数则接受一个"时间点"作为参数，在指定时间点未到来之前线程如果没有获得锁则被阻塞住
                如果在此期间其他线程释放了锁，则该线程可以获得对互斥量的锁
                如果超时（即在指定时间内还是没有获得锁），则返回 false。

            // 注意使用时间类std::chrono

        4) std::recursive_timed_mutex，定时递归 Mutex 类
    
    (2) Lock类(两种)
        std::lock_guard，与 Mutex RAII 相关，方便线程对互斥量上锁
        std::unique_lock，与 Mutex RAII 相关，方便线程对互斥量上锁，但提供了更好的上锁和解锁控制

    (3) 其他类型
        std::once_flag
        std::adopt_lock_t
        std::defer_lock_t
        std::try_to_lock_t

    (4) 函数
        std::try_lock，尝试同时对多个互斥量上锁
        std::lock，可以同时对多个互斥量上锁
        std::call_once，如果多个线程需要同时调用某个函数，call_once 可以保证多个线程对该函数只调用一次

3、<condition_variable>
    condition_variable类是同步原语，能用于阻塞一个线程，或同时阻塞多个线程，直至另一线程修改共享变量(条件)并通知condition_variable

    // 实例
    std::mutex mtx;
    std::condition_variable cv;
    bool ready = false;

    void do_print_id(int id) {
        std::unique_lock<std::mutex> lck(mtx);
        while (!ready)
            cv.wait(lck);
    }

    void go() {
        std::unique_lock<std::mutex> lck(mtx);
        ready = ture;
        cv.notify_all();
    }

    std::thread threads = std::thread(do_print_id, i);
    go();
    threads.join();

    (1) 通知
        1) notify_one
            通知一个等待的线程
            唤醒某个等待(wait)线程。如果当前没有等待线程，则该函数什么也不做，如果同时存在多个等待线程，则唤醒某个线程是不确定的

        2) notify_all
            通知所有等待的线程
            唤醒所有的等待(wait)线程。如果当前没有等待线程，则该函数什么也不做

    (2) 等待
        1) wait
            阻塞当前线程，直到条件变量被唤醒
            > 原理            
                当前线程调用wait()后将被阻塞(此时当前线程应该获得了锁(mutex))，直到另外某个线程调用 notify_* 唤醒了当前线程
                
                在线程被阻塞时，该函数会自动调用unlock()释放锁，使得其他被阻塞在锁竞争上的线程得以继续执行
                
                一旦当前线程获得通知(notified)，wait()函数也是自动调用lck.lock()加锁，以防止其他线程的竞争
            > 一般用法
                std::mutex  mut;
                std::condition_variable cv;

                func1 {
                    unique_lock<mutex> lck(mut);
                    // 这块一定要判断条件
                    while (判断条件) {
                        cv.wait(mux);
                    }
                }

                func2 {
                    unique_lock<mutex> lck(mut);
                    条件改变...
                    cv.notify_**();
                }
                
        2) wait_for
            阻塞当前线程，直到条件变量被唤醒，或到指定时限时长后
            > 原理
                wait_for可以指定一个时间段，在当前线程收到通知或者指定的时间rel_time超时之前，该线程都会处于阻塞状态
                一旦超时或者收到了其他线程的通知，wait_for 返回，剩下的处理步骤和 wait() 类似
                
        3) wait_until
            阻塞当前线程，直到条件变量被唤醒，或直到抵达指定时间点

    (3) 原生句柄
        std::condition_variable::native_handle
        POSIX 系统上，这可以是 pthread_cond_t* 类型值

4、<semaphore.h>
    (1) 初始化
        int sem_init(sem_t *sem,int pshared,unsigned int value); 
        > pshared
            指明信号量是由进程内线程共享，还是由进程之间共享
            pshare为0表明进程内的线程共享，非零表明进程间可共享
        > 返回
            成功时返回 0；错误时，返回 -1，并把 errno 设置为合适的值

    (2) 销毁
        int sem_destroy(sem_t *sem); 

    (3) 成员函数
        int sem_wait(sem_t *sem); 
        // sem_wait将信号量的值减去一个1，但它永远会先等待该信号量为一个非零值才开始做减法
        // 先判断，再做减法

        int sem_trywait(sem_t *sem);  
        // 非阻塞版本

        int sem_post(sem_t *sem);
        // 调用sem_post，信号量加一
        
        int sem_getvalue(sem_t *sem); 
```

### future
```
https://www.jianshu.com/p/7945428c220e

1、promise
    void read(std::future<std::string> * future){
        // future会一直阻塞，直到有值到来
        std::cout << future->get() << std::endl;
    }

    std::promise<std::string> promise;
    // future 相当于消费者, 右值构造
    std::future<std::string> future = promise.get_future();
    // 另一线程中通过future来读取promise的值
    std::thread thread(read, &future);

    promise.set_value("hello future");

    // 注意：一个std::promise实例只能与一个std::future关联共享状态
```

### RAII
```
C++语言的一种管理资源、避免泄漏的机制

C++标准保证任何情况下，已构造的对象最终会销毁，即它的析构函数最终会被调用

RAII 机制就是利用了C++的上述特性，构造一个临时对象(T)，在其构造T时获取资源，最后在T析构的时候释放资源。以达到安全管理资源对象，避免资源泄漏的目的

C++11中lock_guard对mutex互斥锁的管理就是典型的RAII机制

template<typename _Mutex>
class lock_guard {
public:
    typedef _Mutex mutex_type;
    explict lock_guard(mutex_type & m):_M_device(m) {
        _M_device.lock();
    }

    ~lock_guard() {
        _M_device.unlock();
    }

    // 禁止复制构造
    lock_guard(const lock_guard &)=delete;      // c++11 
    // 禁止赋值构造
    lock_guard & operator=(const lock_guard &)=delete;

private:
    mutex_type & _M_device;
}
```

### 模板与数组参数
```
// reference to array of unknown bound 'int []'


// 没有引用
template<typename T> void func(T a);
// 特化数组
template<>
void func<int[]>(int a []) {}



// 有引用
template<typename T> void func(T & a);
// 特化数组(注意下面的数字)
template<>
void func<int[]>(int (&a) [10]) {}  
// 特化指针
template<>
void func<int*>(int *(&a)) {}

```

### 为什么需要size_t
```
主要为了兼容不同系统，提高移植性
例如：需要把指针转换成某个整数类型T来做些按位"与"的对齐操作(指针类型C语言不支持逻辑与等位操作)

或者

#define size_t typeof(sizeof(xxx))
```

### 容器的emplace操作
```
针对vector、deque、list引入了三个新成员，emplace_front、emplace和emplace_back，这些操作分别对应push_front、insert和push_back

目的：避免不必要的临时对象的产生

struct Foo {
    Foo(int m, double n);
};

std::vector<Foo> v;
v.push_back(Foo(1, 3.14));   // 产生一个临时变量
v.push_back({1, 3.14});          // 产生一个临时变量
v.emplace_back(1, 3.14);    // 没有产生临时变量，直接构造


```

### const与优化
```
// 例子
const int a=1;
int* b=(int*)&a;
*b=10;      // 此时内存a可能是1

当开启优化是-O2,当编译器看到这里的a被const修饰，从语义上讲这里的a是不期望被改变的
优化的时候就可以把a的值存放到寄存器中。可以用volatile解决此情况

volatile const int a = 1;
int* b = (int*)&a;
*b=10;  // 此时内存a是10

c++中有一块const内存，并且不同变量，一样的内容，他们的指针地址是一样的，凡是const的变量都在const内存中

```