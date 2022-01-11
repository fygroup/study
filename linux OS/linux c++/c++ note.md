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
函数重载是指在一个类中声明多个名称相同但参数列表不同的函数，这些的参数可能个数或顺序，类型不同，但是不能靠返回类型来判断
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
或
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
virtual void func() = 0;
纯虚函数不能实例化 ，但命名个类指针还是可以的

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

//----------------------------------------
只有构造函数使用成员初始化程序
//----------------------------------------
模板函数 声明和定义要在一块写 不能分成两个文件
template<typename T> void func(T a);    // 声明

template<typename T> void func(T a) {
    ...
}

template<> void func<int>(int a) {
    ...
}

//---read------------------------------------
fIn.read((char*)InBuffer,InBufferSize); //返回一个流对象
size_t InBufferSize_ = fIn.gcount();  //可以得到刚才的读入字节数

//---sleep------------------------------
# include <unistd.h>
sleep(1000);

//---构造函数-----------------------------
Eigen::Matrix<double,Dynamic,Dynamic> mat;
mat = Matrix<T,Dynamic,Dynamic>(m,n);

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

回调函数只能是全局的或是静态的。
全局函数会破坏类的封装性，故不予采用。
静态函数只能访问类的静态成员，不能访问类中非静态成员
//---默认参数----------------------------
带有默认值参数的函数，在实现的时候，参数上是不能有值的。

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

//---c_str()--------------------------------------------------------------------------
string c_str() 返回的是const char*！！！
//---memset--------------------------------------------------------------------------
当内存比较大时，memset还是比较费时间的
```

### atexit
```c++
// 一个程序中可以用atexit()注册多个处理函数（注册次数依赖于你的编译器）
// 这些处理函数的调用顺序与其注册的顺序相反

#include <stdlib.h>

int atexit(void (*function)(void));

::atexit(f1);
::atexit(f2);
// f2 f1
```

### eof
```
到达文件尾，eof()返回true
但事实上，在读完最后一个数据时，eof仍然是false，只有当流再往下读取时，发现文件已经到结尾了，才会将eof修改为true
这也就是为什么使用while(!readfile.eof())会出多现读一行的原因

为了避免这样的情况，可以使用fIn.peek()!=EOF来判断是否到达文件结尾，这样就能避免多读一行
```

### 引用成员的类
```
class A {
    int & a;    // 必须初始化
    A();
};
A::A(int & x):a(x){...} // 必须初始化列表

> 不能有缺省构造函数
    凡是有引用类型的成员变量或者常量类型的变量的类，不能有缺省构造函数
> 必须采用初始化列表的方式
    引用类型的成员变量的初始化问题,它不能直接在构造函数里初始化，必须用到初始化列表，且形参也必须是引用类型
```

### 避免显式调用析构的问题
```
用户显式调用析构函数的时候，只是单纯执行析构函数内的语句，不会释放栈内存，摧毁对象
不要自作聪明随便调用析构
```

### 函数默认参数
```
(1) 有函数声明(原型)
    默认参数可以放在函数声明或者定义中，但只能放在二者之一（在一个文件中）
    double sqrt(double f = 1.0); //函数声明
    double sqrt(double f) {} //函数定义

(2) 没有函数声明
    默认参数在函数定义时指定
    double sqrt(double f = 1.0) {}
```

### c/c++ NULL
```
#ifndef NULL
    #ifdef __cplusplus
        #define NULL 0
    #else
        #define NULL ((void *)0)
    #endif
#endif
```

### pthread_create在类中的使用
```
在C++的类中，普通成员函数不能作为pthread_create的线程函数
只有'static函数'才能作为pthread_create中的线程函数

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

    // static函数，里面不能访问非静态成员变量
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
已知结构体(type)成员(member)的地址ptr，求结构体(type)的起始地址

#define container_of(ptr,type,member) ({\
    const typeof(((type*)0)->member)  *_mptr = (ptr); \
    (type*)((char*)_mptr - offsetof(type,member)); })

#define offsetof(type,member) ((size_t)&((type*)0)->member)
```

### 模板的特化和偏特化
```c++
(1) 模板函数
    // 声明
    template<typename T1, typename T2> void func(T1, T2);
    
    template<typename T1, typename T2> 
    void func(T1, T2) {}
    
    // 特化
    template<> void func(int a, char b) {}          
    template<> void func<int, char>(int a, char b) {}   // 两个写法都行

    // 偏特化 函数模板不能被偏特化
    // 因为模板特化版本不参与函数的重载抉策过程，因此在和函数重载一起使用的时候，可能出现不符合预期的结果。因此标准C++禁止了函数模板的偏特化
    // template<typename T2> void func<int, T2>(int a, T2 b) {}   错误!!!
    // 要实现模板函数偏特化可以借助类模板偏特化
    template<typename T1, typename T2>
    struct F {
        void operator()(T1, T2);
    };

    template<typename T1>
    struct F<T1, int> {
        void operator()(T1, int) {}
    };

    F<char, int>()('a', 12);

(2) 模板类
    // 声明
    template<typename T, typename T1> class Test{};
    // 特化
    template<> class Test<int, char>{};
    //偏特化
    template<typename T1> class Test<int, T1>{};

// 特化 > 偏特化 > 模板类
```

### 模板的实例化
```c++
template <typename T> void Swap(T &a, T &b);
template<> void Swap<int>(int& a, int& b) {}

Swap<int>(a,b);
Swap(a,b)
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
    class Point {
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
    // 外部定义友元函数
    double Distance(Point &a,Point &b) {...}

    Point p1(3.0,4.0),p2(6.0,8.0);
    // 友元函数同普通函数的调用一样，不要像成员函数那样调用
    double d = Distance(p1,p2);

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
        class Node {
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
            friend int BinaryTree::find(); 
            // find方法就可以直接访问Node中的私有成员变量了，而BinaryTree中的其他的方法去不能够直接的访问Node的成员变量
        };

    (3) 友元的继承
        友元关系不能继承。基类的友元对派生类的成员没有特殊访问
        class A{
            friend class B;
        private:
            int a;
        };
        class B{
            void func(A & ob) {
                ob.a // 可以访问class A 的private
            }
        };
        class C : public B{
            void func(A & ob) {
                ob.a // 不能访问class A 的private，因为友元无法继承
            }
        }

```

### 一个无法被继承的类
```
// 思路
将构造和析构函数放在private，但是不能定义这个类的实例
如果定义了一个静态函数用来生成类的实例，那么只能在new上创建该实例

可以用友元的方式
class Base {
    friend class FinalClass;
private:
    Base(){}
    ~Base(){}
};

class FinalClass : virtual public Base { // 注意是虚继承

};

FinalClass a;
FinalClass *b = new FinalClass();

FinalClass 是Base的友元，因此FinalClass可以访问Base中设置为private的构造函数和析构函数，因此FinalClass可以被构造

如果有某个类 X 尝试去继承FinalClass，那么 X 在构造的时候必须要构造Base，并且由于是FinalClass是虚拟继承自Base，X不能通过FinalClass的构造函数来构造Base，它必须直接构造Base，但是由于Base的构造函数是private的，所以X不能被构造

// c++11 final 表示不能被继承
class Base final {

}；
```

### 函数指针
```
typedef (int*)(*func)(int,char);
int* myfunc(int,char){}
func = myfunc;
//&(func) 等价于 func
```

### static变量
```
static 变量最好都写在cpp文件中，除非hpp用到的那个static变量
```

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
__cplusplus：当编写C++程序时该标识符被定义
__VA_ARGS__： 可变参数的宏，代替3个点
```

### 虚函数
```
//要实现C++的多态性必须要用到虚函数，并且还要使用引用或者指针

//需要注意：
    只有类的成员函数才能声明为虚函数，虚函数仅适用于有继承关系的类对象。普通函数不能声明为虚函数
    静态成员函数不能是虚函数，因为静态成员函数不受限于某个对象
    内联函数（inline）不能是虚函数，因为内联函数不能在运行中动态确定位置
    构造函数不能是虚函数
    析构函数可以是虚函数，而且建议声明为虚函数
```

### 虚函数表
```c++
// 虚表是属于类的，而不是属于某个具体的对象，一个类只需要一个虚表即可。同一个类的所有对象都使用同一个虚表
// 实例的类对象包含一个虚函数指针
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

### lambda 转 函数指针
```
没有捕获变量的lambda表达式可以直接转换为函数指针
而捕获变量的lambda表达式则不能转换为函数指针

typedef void(*Ptr)(int*);
 
Ptr p = [](int* p) { };     //OK
Ptr p1 = [&] (int* p) { };  //error

```

### c++ void* 加减
```
标准 C/C++ 不支持 void* 上的加减法，可以先转换成char*，在进行加减
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
bool compare(const int & a, const int & b)
int a[20]={2,4,1,23,5,76,0,43,24,65};
sort(a,a+20,compare);
```

### lambda
```c++
auto f = [=]()->{};
// [=]   值
// [&]   引用

// lambda 内可以使用全局变量，但是局部变量必须捕获才能用

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


for_each(m.begin(), m.end(), [](const pair<string, string> & i){
        
});

```

### sort
```c++
// std::sort
#include <algorithm>
vector<int> a;
sort(a.begin(),a.end(),[](const int & x, const int & y){return(x>y);}); //注意const!!!

// qsort
int x[10];
qsort(x,10,sizeof(int),func);
int func(const void* a, const void* b){
	return((*(int*)a)-(*(int*)b));
}

// 注意
// list不能排序，因为list的iterator不是随机的，而vector可以，因为他是随机的
```

### const char*(初始化)
```
c++允许先初始化再赋值

const char* a；
a = "dafgsfaaafag";
```

### limits
```
#include <limits>
numeric_limits<double>::max() 
numeric_limits<double>::min() 
```

### const char* char const* char* const
```
// const char*  定义一个指向字符常量的指针
const char* ptr;    // ptr指向的内容不能更改

// char const*
char const* ptr     // 和 const char * 等价

// char* const   定义一个指向字符的指针常数
char* const ptr     // const 指针，不能修改ptr值


c++中不存在 const*，所以 char const* 等价于 const char*

```

### execv
```
const char* job[] = {"sh","-c","echo \'fafafafa\'",NULL};
execv("/bin/sh",(char* const *)job); 

// 注意
    (char* const *),而不是(const* char *)!!!
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
#include <functional>
template<typename T, typename... Args> // 返回值类型，参数类型
static void forkRun(function<T(Args...)> func, Args... args);


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
```c++
#include <sys/time.h>
struct timeval tv;
gettimeofday(&tv,NULL);
tv.tv_sec;  // 秒
tv.tv_usec; // 微秒
// 毫秒 = 秒 * 1000 + 微妙/1000
uint64 ms = ((uint64_t)tv.tv_sec) * 1000 + ((uint64_t)tv.tv_usec / 1000);
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

### 非阻塞fd
```
// 设置fd非阻塞
int flag;
if ((flag) = fcntl(fd, F_GETFL, 0)) < 0) error;

flag |= O_NONBLOCK;

if (fcntl(fd, F_SETFL, flag) < 0) error;
```

### 拷贝构造函数和赋值函数
```
class Test
{
public:
    Test(){}
    ~Test(){}
    Test(const Test & t){}              // 拷贝构造函数
    Test& operator=(const Test & t);    // 赋值函数
};

A a;        // 构造函数
A b(a);     // 调用拷贝构造函数
A b = a;    // 调用拷贝构造函数
A b;
b = a;      // 赋值函数
```

### 一个普通的赋值
```c++

struct cv_rect {
    int left;
    int top;
    int right;
    int bottom;
};

struct Rect_t {
    std::vector<cv_rect> rects;
       
    Rect_t() = default;
    Rect_t(const std::vector<cv_rect>& vec) {
        for (size_t i = 0; i < vec.size(); i++) {
		    this->rects.emplace_back(vec[i]);
	    }
    }
    void operator=(Rect_t && t) {
        this->rects = std::forward<vector<cv_rect>>(t.rects);
    }

};

std::vector<cv_rect> v {{1,2,3,4}, {1,2,3,4}};
Rect_t a;
a = v;
// 先构建临时变量Rect_t(v)，调用构造函数 Rect_t(const std::vector<cv_rect>& vec)
// 再调用移动赋值函数 void operator=(Rect_t && t)

```

### g++ 构造优化
```
g++实现省略创建一个只是为了初始化另一个同类型对象的临时对象
指定这个参数（-fno-elide-constructors）将关闭这种优化，强制g++在所有情况下调用拷贝构造函数

// 例子
struct A{
    A() {cout << "A()" << endl;}
    A(int a) {cout << "A(int)" << endl;}
    A(const A&) {cout << "A(const A&)" << endl;}
};
// 正常情况
A a = 1;    // 输出 A(int)
// -fno-elide-constructors 关闭优化
A a = 1;    // 输出 A(int) A(const A&)
```

### 常量左值引用
```
(1) 左值引用
    左值引用的右边必须是左值
    int a;          // a是左值
    int& b = a; 
    int& b = 1;     // 错误，1是右值

(2) 常量左值引用
    常量左值引用的右边可以是左值也可以是右值
    const int & a = 1;
    int& b = a;     // 错误，a是右值，所以b不能引用右值

    // 例子
    string func(){
        return "dddada";
    }
    string & a = func();       //错误 func返回的值已被销毁 引用的必须是左值
    const string & a = func(); //正确 常量左值引用右边可以是右值
    常量左值引用是一个'万能'的引用类型，可以接受左值、右值、常量左值和常量右值

        void func (string & a) {}
        func("dsdas");      // 错误，参数必须是左值
        void func (const string & a) {}
        func("dsdas");      // 正确，参数可以是右值

(3) 常量左值引用意义
    > 传引用不会新创建一个变量
    > 不会调用构造函数，构造拷贝函数，拷贝函数
    > 直接传对象，速度快，同时保证了在函数内部无法对对象进行修改

```

### 右值引用
```c++
(1) 左值引用和右值引用
    // 左值可以取地址、位于等号左边；而右值没法取地址，位于等号右边
    int a = 5;
    int &ref_a = a;         // 左值引用指向左值，编译通过
    int &ref_a = 5;         // 左值引用指向了右值，会编译失败
    const int &ref_a = 5;   // 编译通过 const左值引用是可以指向右值的

    int &&ref_a_right = 5;  // ok
    int &&ref_a_left = a;   // 编译不过，右值引用不可以指向左值
    ref_a_right = 6;        // 右值引用的用途：可以修改右值

(2) 被声明出来的左、右值引用都是左值
    // 因为被声明出的左右值引用是有地址的，也位于等号左边
    // 形参是个右值引用
    void change(string&& right_value) {
        right_value = "bbb";
        std::string tmp(right_value);
        std::string tmp(std::forward<string>(right_value));
    }

    std::string a = "aaa";                      // a是个左值
    std::string & ref_a_left = a;               // ref_a_left是个左值引用
    std::string && ref_a_right = std::move(a);  // ref_a_right是个右值引用
    
    change(a);                                  // 编译不过，a是左值，change参数要求右值
    change(ref_a_left);                         // 编译不过，左值引用ref_a_left本身也是个"左值"
    change(ref_a_right);                        // 编译不过，右值引用ref_a_right本身也是个"左值"
        
    change(std::move(a));                       // 编译通过
    change(std::move(ref_a_right));             // 编译通过
    change(std::move(ref_a_left));              // 编译通过
    change("bbb");                                  // 当然可以直接接右值，编译通过
    // &a &ref_a_left &ref_a_right 这三个左值的地址，都是一样的

(3) 右值引用与移动构造
    // 右值引用就是为移动构造准备的
    struct RValue {
        RValue():sources("hello!!!"){}
        RValue(RValue&& a) {
            sources = std::move(a.sources);
            cout<<"&& RValue"<<endl;
        }

        RValue(const RValue& a) {
            sources = a.sources;
            cout<<"& RValue"<<endl;
        }

        void operator=(const RValue&& a) {
            sources = std::move(a.sources);
            cout<<"&& =="<<endl;
        }

        void operator=(const RValue& a) {
            sources = a.sources;
            cout<<"& =="<<endl;
        }
        string sources;
    };
    void f(string && s){
        // s是右值引用, 但他是个左值
        // string a(s); 不会触发移动拷贝
        string a(std::forward<string>(s)); // 完美转发,告诉编译器s是个右值,要用移动拷贝
    }
```

### std::move
```c++
// std::move和std::forward本质就是一个转换函数

(1) 移动语义(移动构造函数、移动赋值运算符)
    // C++11之前，对象的拷贝控制由三个函数决定：拷贝构造函数、拷贝赋值运算符和析构函数
    // C++11之后，新增加了两个函数：移动构造函数和移动赋值运算符
    // 相比复制构造函数与复制赋值操作符，前者没有再分配内存，而是实现内存所有权转移
    // STL容器都实现了移动语义，大大优化STL容器
    class A {
        int* ptr；
        A(){}
        ~A(){}
        A(const A & a){
            // 拷贝构造，需要对a里面的指针做深层拷贝
        }
        A(A && a){
            // 移动构造，将a中的指针移动到新建的类中，并且a.ptr = nullptr(需要自己实现), 防止两个实例的析构对一个指针的释放
            // 注意: 移动构造函数不会发生类型推断, 输入的参数必须是右值
        }
        A& operator=(A&& a) {
            // 移动赋值
        }
    };

(2) move
    // 移动语义的前提是传入的必须是右值
    // 有时候你需要将一个左值也进行移动语义(这个左值后面不再使用)
    // 必须提供一个机制来将左值转化为右值(std::move)
    string a = "aaa";
    string b = std::move(a);
    a; // 空，string自己实现了移动语义
    b; // aaa
    // move仅仅将左值转换成右值
    // move并没有"移动"什么内容，只是将传入的值转换为右值
    // std::move移动构造函数或者移动赋值运算符，才能充分起到减少不必要拷贝的意义
    template<typename T>
    constexpr typename std::remove_reference<T>::Type&& 
    move(T&& t) {
        return static_cast<typename std::remove_reference<T>::Type&&>(t);
    }
    // 注意 const A& 和 A&&
    class string {
        string(const string& rhs);   // 复制构造函数
        string(string&& rhs);    // 移动构造函数
    }
    // const左值引用可以接收右值，const右值更可以
    // 所以，你其实调用了复制构造函数，那么移动语义当然无法实现
    // 所以，如果你想接下来进行移动，那不要把std::move引用在const对象上

(3) move使用场景
    // 1) 定义的类使用了资源并定义了移动构造函数和移动赋值运算符
    // 2) 该变量即将不再使用

```

### 返回值优化(RVO)
```c++
// copy elision 是编译器对代码进行优化从而来避免拷贝
// 什么时候应该move，什么时候应该依靠copy elision
// 主流的编译器都会100% copy elision以下两种情况

// (1) URVO(Unnamed Return Value Optimization)
    // 函数的所有执行路径都返回 "同一个类型的匿名变量"
    User create_user(const std::string &username, const std::string &password) {
        if (find(username)) return get_user(username);
        else if (validate(username) == false) return create_invalid_user();
        else User{username, password};
    } 
    
// (2) NRVO(Named Return Value Optimization)
    // 函数的所有路径都返回 "同一个非匿名变量"
    User create_user(const std::string &username, const std::string &password) {
        User user{username, password};
        if (find(username)) {
            user = get_user(username);
            return user;
        } else if (user.is_valid() == false) {
            user = create_invalid_user();
            return user;
        } else {
            return user;
        }
    }

// 不要 return std::move(w)
// 此时返回的并不是一个局部对象，而是局部对象的右值引用
// 编译器此时无法进行rvo优化，能做的只有根据std::move(w)来移动构造一个临时对象，然后再将该临时对象赋值到最后的目标。所以，不要试图去返回一个局部对象的右值引用。

```

### RTTI
```c++
// 运行时类型识别，程序能够使用基类的指针或引用来检查着这些指针或引用所指的对象的实际派生类型

// typeid和dynamic_cast
// typeid, 返回指针和引用所指的实际类型
// dynamic_cast, 将基类类型的指针或引用安全地转换为其派生类类型的指针或引用

// 当类中不存在虚函数时，typeid是编译时期的事情，也就是静态类型
// 当类中存在虚函数时，typeid是运行时期的事情，也就是动态类型

#include <typeinfo>
#include <typeindex>
//type_info类
//typeid操作符
//type_index类
// 如果一个class包含至少一个虚函数，则typeid操作符返回表达式的动态类型，否则，typeid操作符返回表达式的静态类型，在编译时就可以确定


// type_info 类
// type_info类在头文件<typeinfo>中定义，代表了一个C++类型的相关信息。一般由typeid操作符返回，不能自己构造
const type_info & info = typeid(int);
// type_info方法
info.name()         //返回类型的名字
info.hash_code()    //返回这个类型的哈希值（具有唯一性）
info.before()       //可以判断一个type_info对象的顺序是否在另一个之前
                    //==和!=操作符，判断两个type_info相等或不等
    
// typeid 操作符
// typeid(类型) 或 typeid(表达式)
string a;
const type_info & info = typeid(a);
const type_info & info = typeid(string);
//注意：typeid既不是函数，也不是宏，它是个操作符

// type_index
// type_index类在头文件<typeindex>中声明，它是type_info对象的一个封装类
// 可以用作关联容器(比如map)和无序关联容器(比如unordered_map)的索引
class A;
unordered_map<type_index, string> names;
names[std::type_index(typeid(A))] = "A";
names[std::type_index(typeid(int))] = "int";
names[std::type_index(typeid(std::string))] = "string";
names[std::type_index(typeid(char))] = "char";
names[std::type_index(typeid(const char*))] = "const char*";
string a;
cout << names[type_index(typeid(a))] << endl;

// dynamic_cast 类型转换(用于含有虚函数的类，因为有type_info指针)
// dynamic_cast主要用于在多态的时候，它允许在运行时刻进行类型转换，从而使程序能够在一个类层次结构中安全地转换类型，把基类指针(引用)转换为派生类指针(引用)
```

### 动态联编和静态联编的坑
```c++
// 在C++中实现动态联编需要同时满足以下三个条件：虚函数，继承关系，基类指针或引用指向子类对象
// 有时编译期未确定的东西，运行期确定可能会出错

class A{
public:
    virtual f();
};

A a;
&a;                 // 0x7ffcda2ea810
&(a.a);             // 0x7ffcda2ea818
a.f();              // 编译期确定函数地址
A *a1 = new A();
a1;             
&(a1->a);
a1->f();            // 运行期确定函数地址

// 一个虚函数指针的错误
class A{
public:
    A(){
        memset(this, 0, sizeof(A));
    }
    virtual void f();
};

A a;
a.f();  // 正确，编译期可以确定
A *b = a;
b->f(); // 错误，动态确定时虚指针已清零
```

### decltype
```
decltype作为操作符，用于获取"表达式"的数据类型

C++11标准引入decltype，主要是为泛型编程而设计，以解决泛型编程中有些类型由模板参数决定而难以（甚至不可能）表示的问题

从语义上说，decltype的设计适合于通用库编写者或编程新手。总体上说，对于变量或函数参数作为表达式，由decltype推导出的类型与源码中的定义可精确匹配。而正如sizeof操作符一样，decltype不对操作数求值

decltype和auto都可以用来推断类型，但是二者有几处明显的差异
> auto忽略顶层const，decltype保留顶层const
> 对引用操作，auto推断出原有类型，decltype推断出引用
> 对解引用操作，auto推断出原有类型，decltype推断出引用
> auto推断时会实际执行，decltype不会执行，只做分析
```

### 类型转换
```
上行转换: 把派生类的指针或引用转换成基类，安全
下行转换: 把基类指针或引用转换成派生类表示，不安全(由于没有动态类型检查)

// static_cast
static_cast相当于传统的C语言里的强制转换，一般情况下类型之间的转化用static_cast
> 用于基类和派生类之间指针或引用的转换，进行上行转换，不能进行下行转换
> 用于基本数据类型之间的转换，如把int转换成char，把int转换成enum
> 把空指针转换成目标类型的空指针
> 把任何类型的表达式转换成void类型
> static_cast不能去掉类型的const、volatile属性

// const_cast
可以在某些情况下用来去除const属性
用于const与非const、volatile与非volatile 之间的转换(去掉类型的const或volatile属性)

// dynamic_cast
借助 RTTI，用于类型安全的向下转型
class Base {};
class Child : public Base {};
Base *a = new Child();
Child *b = dynamic_cast<Child*>(a);


// const char* ==> void*
const char* a = "dadada";
char* x = const_cast<char*>(a);
void* y = static_cast<void*>(x);
```

### enum
```
// c++ 98
enum My{a=1,b,c};
My a = a;

// c++ 11
enum class Color {red,yellow,blue};
Color myColor = Color::red; 
```

### goto
```
int main(){
    goto label;

    int i = 0;   // 错误，上面的goto会忽略这里的变量。需要将变量写在goto之外

    label:
    {
        xxxxxx;
    }
}
```

### 析构竞态
```
C++对象的生命周期需要程序员自己管理，因此析构可能出现竞态尤其是在多线程下，一个对象可以被多个线程访问时，下列情形：
> 即将析构一个对象时从何得知其它线程是否在操作该对象
> 若某个线程正欲操作对象时，如何得知其它线程是否在析构该对象
```

### 类的前置声明
```
前置声明只能作为指针或引用，不能定义类的对象，自然也就不能调用对象中的方法了
```

### 静态成员函数、静态变量
```
如果一个静态成员函数使用的变量没有被其他成员函数使用，那么完全可以把这个 '变量' -> 'static 变量' 放在此函数中

class A {
public:
    static StaticFunc (){
        static std::mutex mutex;    //此变量StaticFun独有，没必要写到外面
        static size_t count = 0;
        std::unique_lock<std::mutex> lock(mutex);//加锁，析构时释放该锁
        count++;
    }
};

```

### 遍历class成员
```
目前c++无法做到，需要对class代码侵入式的改变
```

### 序列化
```
protobuf
Boost.Serialization
```

### 容器存放不同类型变量
```
https://www.zhihu.com/question/33594512?sort=created

// 相关库 std::any std::variant
底层是union和void*实现，union存储基础类型，里面的void*存储自定义类型，再加一个type字段存储类型编号即可
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
> 一个class或者struct，它不包含虚函数，没有虚基类，每一个数据成员都是POD，且所有的父类(如果存在的话)都是POD  
> POD数组  
> 由POD组成的union  
```

### union
```
可以在union里面定义类型，也可以定义构造函数

注意 union cannot define non-POD as member data，对于non pod数据可以使用指针

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

// union判断大小端
union {
    short a;
    char b[sizeof(short)];
} un;

un.a = 0x0102;
if (un.b[0] == 1 && un.b[1] == 2) // 大端
if (un.b[0] == 2 && un.b[1] == 1) // 小端
```

### struct 的位域
```
C语言又提供了一种数据结构，称为'位域'或'位段'
所谓位域是把一个字节中的二进位划分为几个不同的区域，并说明每个区域的位数

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


Object<int> a = Object<int>::Object();  // 移动构造
Object<int> b = a; // 拷贝构造
```

### 可变参数
```c++
template<typename f, typename ...Argvs>
void callback(function<void(Argvs...)> f, Argvs&&... argvs){
    f(std::forward<Argvs>(argvs)...);
}
```

### 模板特化(指针)
```c++
(1) 模板函数
    template<typename T> void func(T* t) {
        cout << "T*" << endl;
    }

    template<typename T> void func(T t) {
        cout << "T" << endl;
    }

    int a = 2;
    func(a);      // T
    func(&a);     // T*

(2) 模板类
    template<typename T>
    class a{
    public:
        a(){
            cout << "T" << endl;
        }
    };

    template<typename T>
    class a<T*>{
    public:
        a(){
            cout << "T*" << endl;
        }
    };

    a<char> x; 
    a<char*> x1;
```

### 模板函数数组特化
```c++ ???
// reference to array of unknown bound 'int []'

template<typename T> void func(T a) {}
// 指针特化
template<typename T> void func(T * a);
// 引用特化
template<typename T> void func(T & a);

// int数组特化
template<> void func<int[]>(int a []) {}
// int数组特化(注意下面的数字)
template<> void func<int[]>(int (&a) [10]) {}  
// int指针特化
template<> void func<int*>(int *(&a)) {}

```

### 模板类中的模板函数
```c++
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
    template<typename T1> My(const My<T1> & m); //注意
    T & operator=(const My<T> m) = default;
};

template<typename T>
template<typename T1>
My<T>::My(const My<T1> & m){
    t = static_cast<T>(m.t);
}
```

### Traits Classes 
```c++
// https://www.cnblogs.com/mangoyuan/p/6446046.html
// https://cloud.tencent.com/info/a180e28f80b999eb22700e2407fc0957.html
// https://blog.csdn.net/lihao21/article/details/55043881

// Traits classes 的作用主要是用来为使用者提供类型信息
// STL中，容器与算法是分开的，容器与算法之间通过迭代器联系在一起
// 函数的template参数推导机制推导的只是参数，无法推导函数的返回值类型

// traits 使用的关键技术 -> 模板的特化与偏特化
template <typename T>
struct iterator_traits {
    typedef typename T::value_type value_type;
};

template <typename T>
struct iterator_traits<T*> {
    typedef T value_type;
};

template <typename T> 
typename iterator_traits<T>::value_type func(T a) {
    return *a;
}

// 我们就可以知道模板类和指针的类型信息
```

### value_type
```c++
//对于大部分STL都适用

vector<int>::iterator it = a.begin();
vector<int>::iterator::value_type a = 1;
// 等价于
int a = 1;

//萃取
class vector {
public:
    class iterator {
    public:
        typedef T value_type;
        // ...
    };
    // ...
};
```

### stateful
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
    string b;
    char c;
};

struct TestS a = {1, "dasda", 'c'};
struct TestS a {1, "dasda", 'c'};

// 注意
c++中只允许POD对象可以进行上述的初始化
带有构造函数的对象不再被视为POD。对象只能包含其他POD类型作为非静态成员(包括基本类型)
当然POD也可以具有静态功能和静态复杂数据成员
struct默认不带构造函数可以进行上述初始化，class默认带有构造函数不能进行上述初始化
```

### ({})
```
// int a = 1
int a = ({
    ....;
    1;
});

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
    GNU C 的一大特色就是__attribute__ 机制
    
    __attribute__ 可以设置函数属性(Function Attribute)、变量属性(Variable Attribute)和类型属性(Type Attribute)
    语法格式 __attribute__(参数)

(2) 函数属性(Function Attribute)
    __attribute__((noreturn))   // 表示没有返回值
    __attribute__((unused))     // 表示该函数或变量可能不使用，这个属性可以避免编译器产生警告信息 
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
```c++
#include <sys/wait.h>

pid_t wait(int *statloc);
pid_t waitpid(pid_t pid,int *statloc, int options);
// statloc  指向终止进程的终止状态，如果不关心终止状态可指定为空指针
// pid      有四种情况
//          pid == -1 等待任意子进程
//          pid > 0   等待进程ID与pid相等的子进程
//          pid == 0  等待组ID等于调用进程组ID的任意子进程
//          pid < -1  等待组ID等于pid绝对值的任意子进程

pid_t wait(int *statloc) {
    return waitpid(-1, statloc, 0);
}

// waitpid的正确姿势
pid_t pid;
int status;

if((pid = fork())<0){
    status = -1;
} else if(pid == 0){
    execl("/bin/sh", "sh", "-c", cmdstring, (char *)0);
} else {
    while(waitpid(pid, &status, 0) < 0){
        // 对于慢系统调用，当返回错误时，要判断errno是否是EINTR
        // EINTR，可能由于系统中断导致系统阻塞调用提前返回（详见套接字）
        if(errno != EINTER){
            status = -1;
            break;
        }
    }
}

```

### map 判断键值是否存在
```c++
(1) insert
    std::map<int, string> myMap;
    std::pair<std::map<int, string>::iterator, bool> findItem;  
    findItem = myMap.insert(std::pair<int, string>(1, "student_one"));  
    if(findItem.second == true) {
        // myMap 中没有insert的数据，说明插入成功
        cout<<"Insert Successfully"<<endl;
    } else {
        // myMap 中有insert的数据，说明插入失败
    }
    cout << findItem.first->first << endl;
    cout << findItem.first->second << endl;
        
(2) find
    map<>::iterator it = m.find('key');
    if (it == m.end()) #不存在
    
// 插入键值对的方法
// 1) map[k] = v
//  直接赋值这种方式，当要插入的键存在时，会覆盖键对应的原来的值。如果键不存在，则添加一组键值对
// 2) map.insert()
//  这是map自带的插入功能。如果键存在的话，则插入失败，也就是不插入。 使用insert()函数，需要将键值对组成一组才可以插入
```

### map 键值排序
```c++
// map的定义
template <class Key, class T, class Compare = less<Key>, class Allocator = allocator<pair<const Key,T> > > class map;

// less的结构
template <class T>
struct less : binary_function <T,T,bool> {
    bool operator() (const T& x, const T& y) const {
        return x < y;
    }
}

// 键值排序
struct cmpkeylen{
    bool operator()(const string & a, const string & b){
        return a.length() > b.length(); // 按字符串长度排序
    }
};

std::map<std::string, int, cmpkeylen> mymap;
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

### iterator const_iterator
```c++
vector<int> a;
vector<int>::iterator i = a.begin();
i++;
*i = 1;

const vector<int> a;
vector<int>::const_iterator i;  //对于const 必须用const_iterator !!!
```

### const对象
```c++
// const 对象只能访问 const 成员函数

// const 修饰的参数引用的对象，只能访问该对象的const成员函数，因为调用其他函数有可能会修改该对象的成员

// 编译器为了避免该类事情发生，会认为调用非const函数是错误的。

class A {
    void test() {}
    void test1() const {}   // const成员函数表示，不会修改任何成员内容
};

void accessFunc(const A & a){
    // b.test();   错误
    b.test1();  // 正确
}

// 当函数后面加了const时，返回引用和指针时要加const，但是返回非引用可以不加
const A & createA() const {}
const A * createA() const {}
A func() const {}
```

### initializer_list
```c++
void g(std::vector<int> const &items){}; 
void g(std::list<int> const &items){};

g({ 1, 2, 3, 4 }); //会报错  编译器分不清 vector还是list

//对于{}的固定数组initializer_list更合适
void g(std::vector<int> const &items){}; 
void g(std::list<int> const &items){}; 
void g(std::initializer_list<int> const &items){}; //注意const initializer_list不能修改，更符合参数的特点
g({ 1, 2, 3, 4 });
```

### 函数对象
```c++
// 函数对象是实现 operator() 的任何类型
// C++ 标准库主要使用函数对象作为容器和算法内的排序条件

// 优势
// 相对于直接函数调用，函数对象有两个优势
// 第一个是函数对象可包含状态
// 第二个是函数对象是一个类型，因此可用作模板参数

// 创建函数对象
// 若要创建函数对象，请创建一个类型并实现 operator()，例如：
class Functor
{
public:
    int operator()(int a, int b) {
        return a < b;
    }
};

// 函数对象和容器
// C++ 标准库在标头文件中包含若干函数对象 <functional>
// 这些函数对象的一个用途是用作容器的排序条件
template <typename Key, typename Traits=less<Key>, typename Allocator=allocator<Key>> class set;


```

### std::function
```c++
#include <functional>

// 普通函数
int add(int a, int b){return a+b;} 

// lambda表达式
auto mod = [](int a, int b)->int{ return a % b;}

// 函数对象类
struct divide{
    int operator()(int denominator, int divisor){
        return denominator/divisor;
    }
};
std::function<int(int ,int)>  a = add; 
std::function<int(int ,int)>  b = mod ; 
std::function<int(int ,int)>  c = divide(); //divide类构造

// 类成员函数提取
class A{
    void func(int){}
};

std::function<void(A*,int)> f = &A::func;
A a;
f(&a, 2);
```

### std::bind
```c++
// 绑定普通函数
double my_divide(double x, double y) {return x/y;}

std::function<double(double)> fn_half = std::bind(my_divide, std::placeholders::_1, 5);  // placeholders::_1 占位符
fn_half(10);

void func(int){}
std::function<void()> f = std::bind(func, 12);
f();

// 绑定成员函数
// bind绑定类成员函数时，第一个参数表示对象的成员函数的指针，第二个参数表示对象的地址
struct A {
    int func(int a, int b) {
        return a + b;
    }
};

A ac;
std::function<int(int)> a = std::bind(&A::func, &ac, std::placeholders::_1, 10);
cout << a(12) << endl;
std::function<int(A*, int, int)> b = &A::func;
cout << b(&ac, 1, 2) << endl;
std::function<int(int, int)> c = std::bind(&A::func, &ac, std::placeholders::_1, std::placeholders::_2);
cout << c(5, 2) << endl;
```

### std::pair
```c++
// std::pair 是一个结构体模板，其可于一个单元存储两个相异对象
template<typename T1, typename T2> struct pair;

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

// 成员类型
std::pair<int, string> a(1, "dasas");
std::pair<int, string>::first_type x = 1
```

### operator new
```c++
(1) 原型
    void *operator new(size_t);         //allocate an object
    void *operator new(size_t, void*);  //placement
    void *operator delete(void *);      //free an object

    void *operator new[](size_t);       //allocate an array
    void *operator new[](size_t, void*);//placement
    void *operator delete[](void *);    //free an array

(2) operator三种形式
    // 1
    void* operator new(size_t) throw(std::bad_alloc);
    A *a = new A;
    // 失败时抛出bad_alloc

    // 2
    void* operator new(size_t, nothrow_value) throw();
    A* a = new(std::nothrow) A;
    // 同上，但是失败时返回null
    // 调用operator new (sizeof(A), nothrow_value)
    // 调用A:A()
    // 返回指针

    // 3
    void* operator new(size_t, void* ptr) throw();
    //在ptr所指地址上构建一个对象(通过调用其构造函数)
    char ptr[1024];
    A* a = new(ptr) A(); // 本身返回ptr，可以被重载
    
// new执行时的细节，三个步骤
A* a = new A;
char* tmp = operator new (sizeof(A));   // 申请内存
new(tmp) A();                           // 在内存上构造对象
a = tmp;                                // 返回指针
// 注意，上述实际的顺序可能重排
// 申请内存、返回指针、构造对象
a = operator new (sizeof(A));
new(a) A();

(3) operator示例
    template<typename T>
    class a {
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
        cout << y->n << endl;
        return 0;
    }
```

### new和delete的调用过程
```c++
// new过程
// (1) 调用operator new标准库函数申请内存
// (2) 在这一块内存上对类对象进行初始化，调用的是相应的构造函数
// (3) 返回新分配并构造好的对象的指针

// delete过程
// (1) 调用对象的析构函数
// (2) 调用标准库函数 operator delete 来释放该对象的内存
```

### 对象数组与对象指针 ???????????
```
对象数组除了分配对应大小的空间外，还分配了一段空间(double)用来储存数组的个数  [8 bytes] + [数组空间]

所以new[]分配的空间要用delete[]来释放，其逻辑是先根据数组的个数调用n遍析构函数，然后将指针地址减去8再释放整个空间

class A {
    void* operator new(size_t size) {
        cout << size << endl;
        //...
    }
    void* operator new[](size_t size) {
        cout << size << endl;
        //...
    }
};

A* a = new A;       // 打印 1
A* a1 = new A [1];  // 还是打印 1

char buf[100];
A* a = new A(buf) [1];
cout << (void*)buf << endl; // 0x7ffecc65f850
cout << (void*)&a[0] << endl; // 0x7ffecc65f858
```

### allocator
```c++
// https://zhuanlan.zhihu.com/p/34725232

#include <memory>

(1) 概念
    // allocator是STL的重要组成，allocator除了负责内存的分配和释放，还负责对象的构造和析构
    
    // 例如vector的class如下
    template<typename T, typename Alloc = allocator<T>>
    class vector{
        //每个vector内部实例一个allocator
        Alloc data_allocator;
    };
    
    std::vector<int> v;
    // 等价于
    std::vector<int, allocator<int>> v;

(2) allocator结构
    template <class T>
    class allocator {
    public:
        typedef T value_type;
        typedef T* pointer;
        typedef const T* const_pointer;
        typedef T& reference;
        typedef const T& const_reference;
        typedef size_t size_type;
        typedef ptrdiff_t difference_type;

        // 分配空间 (n * size(T))
        // 存储n个T对象，第二个参数是个提示。实现上可能会利用它来增进区域性(locality)，或完全忽略之
        pointer allocate(size_type n, const void* = 0)
        
        // 释放空间
        void deallocate(pointer p, size_type n)
        
        // 调用对象的构造函数，等同于 new(p) T(x)
        void construct(pointer p, const T& x)
        // 调用对象的析构函数，等同于 p->~T()
        void destroy(pointer p)
    };    

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
    // allocator类定义了两个可以构建对象的算法，以下这些函数将在目的地构建元素，而不是给它们赋值
    #include <memory>
    vector<string> list(10, "aaaa");
    uninitialized_copy_n(list.begin(), 5, s);  //构建填充
    uninitialized_fill_n(list.begin(), 5, s);  //拷贝填充

(4) 实现自己的allocator
    // 根据STL的规范，要实现allocator的接口
    template<class T>
    class Allocator
    {
    public:
        typedef T               value_type;
        typedef T*              pointer;
        typedef const T*        const_pointer;
        typedef T&              reference;
        typedef const T&        const_reference;
        typedef size_t          size_type;
        typedef ptrdiff_t       difference_type;

        template<class U>
        struct rebind
        {
            typedef Allocator<U> other;
        };
        
        // 分配空间 (n * size(T))
        // 存储n个T对象，第二个参数是个提示。实现上可能会利用它来增进区域性(locality)，或完全忽略之
        pointer allocate(size_type n, const void* hint=0) {
            return _allocate((difference_type)n, (pointer)0); // 自定义
        }

        void deallocate(pointer p, size_type n) {
            return _deallocate(p);  // 自定义
        }

        void construct(pointer p, const T& value) {
            _construct(p, value); // 自定义
        }

        void destroy(pointer p) {
            _destroy(p); // 自定义
        }

        pointer address(reference x) {
            return (pointer)&x;
        }

        const_pointer address(const_reference x) {
            return (const_pointer)&x;
        }

        size_type max_size() const {
            return size_type(UINT_MAX/sizeof(T));
        }
    };
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
        //自定义cout缓冲区
        char buf[1024] = {0};
        stringbuf a;
        a.pubsetbuf(buf, 1024);
        std::cout.rdbuf(&a);
        cout << "ddasdaasdas";
        printf("-- %s --\n", a.str().c_str());
        printf("-- %s --\n", buf);

        //自定义file缓冲区
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
    其基本思想是：函数A在执行过程中发现异常时可以不加处理，而只是"拋出一个异常"给A的调用者，假定为函数B
    拋出异常而不加处理会导致函数A立即中止，在这种情况下，函数B可以选择捕获A拋出的异常进行处理，也可以选择置之不理。如果置之不理，这个异常就会被拋给B的调用者，以此类推
    如果一层层的函数都不处理异常，异常最终会被拋给最外层的main函数。main函数应该处理异常。如果main函数也不处理异常，那么程序就会立即异常地中止

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
    void func() throw(int, double); // 声明抛出(int, double)类型的异常
    > 注意
        c++11之后不建议函数后面定义错误throw
        void func();            // 可抛出任何异常
        void func() noexcept;   // 不会抛出异常


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
        } catch (std::out_of_range & e) {
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
```c++
// 互斥锁
pthread_mutex_t mutex;
pthread_mutex_init()  
pthread_mutex_lock()    // 锁定互斥锁，如果尝试锁定已经被上锁的互斥锁则阻塞至可用为止
pthread_mutex_unlock() 	// 释放互斥锁
pthread_mutex_destory() // 互斥锁销毁函数

// pthread_exit()和return类似, 就是退出的作用，不涉及资源释放

> pthread_create
    // 注意函数定义: void* (*)(void*)
    void* test(void *ptr){
        cout << "hello world." << endl;
    }
    pthread_t tid;
    pthread_create(&tid, NULL, test, NULL);
    pthread_join(tid, NULL);

> pthread_join
    // 子线程合入主线程，主线程阻塞等待子线程结束，然后回收子线程资源
    int pthread_join(pthread_t thread, void **retval)
    // 第一个参数为线程标识符，即线程ID
    // 第二个参数retval为用户定义的指针，用来存储线程的返回值

> pthread_detach
    // 主线程与子线程分离，子线程结束后，资源自动回收
    // pthread有两种状态joinable状态和unjoinable状态
    // joinable: 当线程函数自己返回退出时或pthread_exit时都不会释放线程所占用的资源，只有当你调用了pthread_join之后这些资源才会被释放
    // unjoinable: 这些资源在线程函数退出时或pthread_exit时自动会被释放
    int pthread_detach(pthread_t pid)

> pthread_self 获取当前线程id

> pthread_once
    // pthread_once在多线程环境中只执行一次
    int pthread_once(pthread_once_t *once_control, void (*init_routine)(void))
    // 第一个参数为pthread_once_t变量
    // 第二个参数为无参数函数指针，type: void(*func)(void)
    
    pthread_once_t once = PTHREAD_ONCE_INIT;

    void *func_once(){
        cout << "func once" << endl;
    }
    void *func1(void *arg){
        pthread_once(&once, func_once);
    }
    void *func2(void *arg){
        pthread_once(&once, func_once);
    }
    pthread_t td1, td2;
    pthread_create(&td1, NULL, func1, NULL);
    pthread_create(&td2, NULL, func2, NULL);
    // 结果只会输出一次func once

> 线程取消
    // 线程取消的方法是向目标线程发Cancel信号
    // 但如何处理Cancel信号则由目标线程自己决定，或者忽略、或者立即终止、或者继续运行至Cancelation-point(取消点)

    // 取消点
    // > 通过pthread_testcancel调用以编程方式建立线程取消点
    // > 线程等待pthread_cond_wait或pthread_cond_timewait()中的特定条件
    // > 被sigwait(2)阻塞的函数
    // > 一些标准的库调用。通常，这些调用包括线程可基于阻塞的函数???

    > pthread_cancle
        int pthread_cancel(pthread_t thread)
        // pthread_cancel调用并不等待线程终止，它只提出请求
        // 发送终止信号给thread线程，如果成功则返回0，否则为非0值。发送成功并不意味着thread会终止

    > pthread_setcancelstate
        int pthread_setcancelstate(int state, int *oldstate)
        // 设置本线程对Cancel信号的反应
        // state有两种值：PTHREAD_CANCEL_ENABLE（缺省）和PTHREAD_CANCEL_DISABLE
        // old_state如果不为 NULL则存入原来的Cancel状态以便恢复

    > pthread_setcanceltype
        int pthread_setcanceltype(int type, int *oldtype) 
        // 设置本线程取消动作的执行时机
        // type有两种取值：PTHREAD_CANCEL_DEFFERED 和 PTHREAD_CANCEL_ASYCHRONOUS
        // 仅当Cancel状态为Enable时有效，分别表示收到信号后继续运行至下一个取消点再退出和立即执行取消动作(退出)
        // oldtype如果不为NULL则存入运来的取消动作类型值
    
    > pthread_testcancel
        void pthread_testcancel(void) 
        // 手动创建一个取消点，检查本线程是否处于Cancel状态，如果是，则进行取消动作(退出)，否则直接返回
 
> 线程终止的清理
    // https://blog.csdn.net/caianye/article/details/5912172
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

> pthread_kill
    int pthread_kill(pthread_t thread, int sig);
    // 向线程发送signal，如果线程的代码内不做任何信号处理，则会按照信号默认的行为影响整个进程
    // 也就是说，如果你给一个线程发送了SIGQUIT，但线程却没有实现signal处理函数，则整个进程退出
    // 注意子线程信号共享父进程，所以会影响整个进程

    pthread_kill(ptd, 0)
    // 如果int sig的参数是0呢，这是一个保留信号，一个作用就是用来判断线程是不是还活着
    // 返回值0，线程仍然活着
    // 返回值ESRCH，线程已不存在
    // 返回值EINVAL，信号不合法


> 条件变量
    // (1) 条件变量创建
        // 静态创建
        pthread_cond_t cond = PTHREAD_COND_INITIALIZER;
        // 动态创建
        pthread_cond _t cond;
        pthread_cond_init(&cond,NULL);
        // 其中的第二个参数NULL表示条件变量的属性，虽然POSIX中定义了条件变量的属性，但在LinuxThread中并没有实现，因此常常忽略

    // (2) 条件等待
        pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
        pthread_mutex_lock(&mutex);
        while(条件1不成立)
            pthread_cond_wait(&cond,&mutex);
            ...
        pthread_mutex_unlock(&mutex);
        
        // pthread_cond_wait函数的返回并不意味着条件的值一定发生了变化，必须重新检查条件的值
        // 因为存在虚假唤醒

    // (3) 条件激发
        pthread_mutex_lock(&mutex);
        if (条件1成立)
            pthread_cond_signal(&cond);
        pthread_mutex_unlock(&mutex);
        
        // 必须在互斥锁的保护下使用相应的条件变量，否则对条件变量的解锁有可能发生在锁定条件变量之前，从而造成死锁
        // 条件1不成立的时候,执行pthread_cond_signal(&cond)，激发条件变量cond，使得被挂起的线程被唤醒

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

    // (4) 条件变量的销毁
        pthread_cond_destroy(&cond);
        // 在linux中，由于条件变量不占用任何资源，所以，这句话除了检查有没有等待条件变量cond的线程外，不做任何操作

// 10、pthread_cond_timedwait
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

    具体见<c++同步>

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
            一般的类成员函数是不能用作回调函数的，因为库函数在使用回调函数时，都会传递指定的符合回调函数声明的的参数

            类成员函数隐式包含一个this指针参数，所以把类成员函数当作回调函数编译时因为参数不匹配会出错

            1> 方法一
                把成员函数设成静态成员函数，不属于某个对象，属于整个类
            2> 方法二
                传递this指针
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
                std::mutex 不允许拷贝构造，不允许移动构造
                最初产生的 mutex 对象是处于 unlocked 状态
            2> 成员函数
                > lock()
                    调用线程将锁住该互斥量
                    线程调用该函数会发生下面3种情况
                    1> 如果该互斥量当前没有被锁住，则调用线程将该互斥量锁住，直到调用 unlock之前，该线程一直拥有该锁顶
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
        while (!ready)      // 重要！！！ 因为会存在虚假唤醒
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

            > 两个版本
                void wait (unique_lock<mutex>& lck);
                template <class Predicate>
                void wait (unique_lock<mutex>& lck, Predicate pred);
                只有当 pred 条件为 false 时调用 wait() 才会阻塞当前线程，并且在收到其他线程的通知后只有当 pred 为 true 时才会被解除阻塞
                第二种情况类似：while (!pred()) wait(lck);

            > 一般用法
                std::mutex  mut;
                std::condition_variable cv;

                func1 {
                    unique_lock<mutex> lck(mut);
                    // 这块一定要判断条件
                    while (判断条件) {
                        cv.wait(lck);
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
        int sem_init(sem_t *sem, int pshared, unsigned int value); 
        > pshared
            指明信号量是由'进程内线程共享'还是'进程之间共享'
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

### 非类型模板
```c++
// 模板参数不限定于类型，普通值也可作为模板参数
template<int> struct aa{};
template<> struct aa<1>{};  // 特化
```

### std::ref std::cref
```c++
// std::bind std::thread 是对参数直接拷贝，而不是引用

void func(int & n) {
    n++;
}
int n = 1;
std::function<void()> f = std::bind(func, n);
f();    // 1， 因为bind只是拷贝n

std::function<void()> f = std::bind(func, std::ref(n));
f();    // 2

std::thread(func, std::ref(n));

```

### future
```c++
// https://www.jianshu.com/p/7945428c220e
// 库的内容
// classes
//     future
//     future_error
//     packaged_task
//     promise
//     shared_future
// enum classes
//     future_errc
//     future_status
//     launch
// functions
//     async
//     future_category

(1) promise
    void func(std::future<int> & fut){
        int x= fut.get();   // 等待
    }
    std::promise<int> prom;                     // 生成一个 std::promise<int> 对象
    std::future<int> fut = prom.get_future();   // 和 future 关联
    std::thread t(func, std::ref(fut));         // 将 future 交给另外一个线程t
    prom.set_value(10);                         // 设置共享状态的值, 此处和线程t保持同步
    t.join();
    // 注意
    // std::promise 的 operator= 没有拷贝语义，operator= 只有 move 语义
    // 一个std::promise实例只能与一个std::future关联共享状态

(2) promise::set_exception
    void func(fucture<int> & fut){
        try {
            int x = fut.get();
        }catch(std::exception & e){
            cout << e.what() << endl;
        }
    }
    std::promise<int> prom;
    std::functure<int> fnt = prom.get_future();
    thread t(func, std::ref(fnt));
    int x;
    std::cin.exceptions(std::ios::failbit);   // throw on failbit
    try{
        cin >> x;
        prom.set_value(x);                  // sets failbit if input is not int
    }catch{
        prom.set_exception(std::current_exception());
    }

(3) std::async
    // 这个函数是对上面的对象的一个整合，async先将可调用对象封装起来，然后将其运行结果返回到promise中，这个过程就是一个面向future的一个过程，最终通过future.get()来得到结果

    template<typename Fn, typename ...Argvs>
    std::future<typename std::result_of<Fn(Argvs...)>::type>
    async(std::launch policy, Fn&& fn, Argvs&&... argvs);
    // launch 三中策略
    // std::launch::async 保证异步行为，执行后，系统创建一个线程执行对应的函数
    // std::launch::deffered 当其他线程调用get()来访问共享状态时，将调用非异步行为
    // std::launch::async||std::launch::deffered 默认策略，由系统决定怎么调用
    // future的返回结果
    // std::future_status::deferred 异步操作还没开始
    // std::futurn_status::ready    异步操作已经完成
    // std::futurn_status::timeout  异步操作超时

    template<typename Fn, typename... Argvs>
    std::future<typename std::result_of<Fn(Argvs...)>::type>
    async(Fn&& fn, Argvs&&... argvs);   // 默认策略

    int func(int a){
        std::this_thread::sleep_for(std::chrono::seconds(10));
        return a
    }

    auto fut = std::async(std::launch::async, &func, 12);   // 启动线程运行func
    std::future_status status;
    do {
        status = fut.wait_for(std::chrono::seconds(1));
        switch (status) {
            case future_status::ready:
                break;
                cout << "ready ";
            case future_status::deferred:
                break;
            case future_status::timeout:
                cout << "timeout ";
                break;
        }
    }while(status != future_status::ready);

    cout << "result is " << fut.get() << endl;
    // 输出 timeout timeout timeout timeout ... ready

(4) 坑爹的async
    1) 同步的async
        std::async([](){/*print n*/}, 1);
        std::async([](){/*print n*/}, 2);
        std::async([](){/*print n*/}, 3);
        // 上述代码只会一步一步的执行 1 -> 2 -> 3，并不会同时执行
        // 第一行std::async创建了一个类型为std::future<void>的临时变量Temp
        // 临时变量Temp在开始执行第二行之前发生析构
        // std::future<void>的析构函数，会同步地等操作的返回，并阻塞当前线程
    2) 一个async，一个线程
        std::future<void> a = std::async([](){},1);
        std::future<void> b = std::async([](){},2);
        std::future<void> c = std::async([](){},3);
        // async返回的变量没有析构(就没有等待)，所以执行下一步
        // 但是每个async会创建一个线程，非常消耗资源

    // 推荐folly的future
```

### RAII
```c++
// RAII是 C++语言的一种管理资源、避免泄漏的机制
// C++标准保证任何情况下，已构造的对象最终会销毁，即它的析构函数最终会被调用
// RAII 机制就是利用了C++的上述特性，构造一个临时对象(T)，在其构造T时获取资源，最后在T析构的时候释放资源。以达到安全管理资源对象，避免资源泄漏的目的

// C++11中lock_guard对mutex互斥锁的管理就是典型的RAII机制

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
    lock_guard(const lock_guard &) = delete;      // c++11 
    // 禁止赋值构造
    lock_guard & operator=(const lock_guard &) = delete;

private:
    mutex_type & _M_device;
};
```

### 为什么需要size_t
```cpp
// 主要为了兼容不同系统，提高移植性
// 例如需要把指针转换成某个整数类型T来做些按位"与"的对齐操作(指针类型C语言不支持逻辑与等位操作)

// 或者
#define size_t typeof(sizeof(xxx))
```

### 容器的emplace操作
```c++
// c++ 11 vector、deque、list引入了三个新成员
// emplace_front   push_front
// emplace         insert
// emplace_back    push_back
// 这些函数可以代替旧的函数，它们的优点 避免不必要的临时对象的产生

class A {
    int a;
    A(int _a):a(_a){}
    A(const A & a){ 拷贝构造 }
    A(A && a){移动构造}
    A & operator = (const A & a) {} 
};

std::vector<A> v;
v.push_back(Foo(1));    // 调用构造函数，调用移动构造
v.emplace_back(Foo(2)); // 调用构造函数

```

### const与volatile const
```c++
// 例子
const int a=1;
int* b=(int*)&a;
*b=10;      // 此时内存a可能是1

// 当开启优化是-O2,当编译器看到这里的a被const修饰，从语义上讲这里的a是不期望被改变的
// 所以优化的时候就会把a的值存放到寄存器中

// 用volatile告诉编译器不要对他进行优化
volatile const int a = 1;
int* b = (int*)&a;
*b=10;  // 此时内存a是10

// c++中有一块const内存，并且不同变量，一样的内容，他们的指针地址是一样的，凡是const的变量都在const内存中

```

### virtual 在多态中的使用
```c++
// 子类必须实现基类一模一样的接口(virtual)函数，否则不会进行多态调用

// 情景一
class A {
public:
    virtual void func(int x) {}
};

class B {
public:
    virtual void func(float x) {}
};

A* a = new B();
a->func(1); // 调用A

// 情景二
class A
{
public :
    virtual void func(int x)
    {
        cout << "A: int " << x << endl;
    }
};
class B :public A
{
public :
    virtual void func(float x)
    {
        cout << "B: float " << x << endl;
    }
};
A *a = new B();
a->func(12);
a->func(12.12);
// 结果
A: int 12
A: int 12
```

### class 大小
```c++
class A {};
A a;
sizeof(a);   // 1

class A {
    int a;
};
A a;
sizeof(a);  // 4

class A {
public:
    void func(){}
};
A a;
sizeof(a)   // 1 函数不占内存

class A {
public:
    virtual void func(){}
};
A a
sizeof(a)   // 8 多了虚函数指针
```

### 模板函数参数可以用T
```c++
template<typename T>
void printVec(const vector<T> & vec){
    for_each(vec.begin(), vec.end(), [](const T & a){
        cout << a << " ";
    });
    cout << endl;
}

printVec<int>({1,2,3,4}); // 必须要显式调用
```

### 迭代器失效
```
(1) push_back导致迭代器失效
    vector在push_back的时候当容量不足时会触发扩容，导致整个vector重新申请内存，并且将原有的数据复制到新的内存中，并将原有内存释放，这自然是会导致迭代器失效的，因为迭代器所指的内存都已经被释放

(2) insert和erase导致的迭代器失效
    插入操作导致vector扩容，迭代器失效原因和push_back相同
    插入操作引起vector内元素移动，导致被移动部分的迭代器失效

```

### 构造函数 析构函数 虚函数
```c++
// 构造函数不能是虚函数
// 由于对象开始还未分配内存空间，所以根本就无法找到虚函数表，从而构造函数也无法被调用。所以构造函数是不能成为虚函数
// 析构函数最好是虚函数

class Base {};
class Sub:public Base {};

SubClass* pObj = new SubClass();
delete pObj;
// 不管析构函数是否是虚函数(即是否加virtual关键词)，delete时基类和子类都会被释放

BaseClass* pObj = new SubClass();
delete pObj;
// 若析构函数是虚函数(即加上virtual关键词)，delete时基类和子类都会被释放
// 若析构函数不是虚函数(即不加virtual关键词)，delete时只释放基类，不释放子类
```

### 柔性数组
```c++
// 例子
struct A {
	int a;
	char b[];
};

A* a = (A*)malloc(sizeof(A) + sizeof(char) * 10);
a->a = 12;
const string s = "dasdadadd";
strcpy(a->b, s.c_str());
cout << a->b << endl;
cout << sizeof(*a) << endl; // 4
free(a);    

```

### c语言实现c++的3大特性
```cpp
1、封装
    typedef struct MyClass {
        int a;  // 成员变量
        int (*func1)(int, int);
        void (*func2)(int);
    } MyClass;

    int funcItem1(int a, int b){
        ...
    }
    void funcItem2(int a){
        ...
    }

    // 构造函数
    MyClass* constructor(int _a){
        MyClass *item = new MyClass;
        item->a = _a;
        item->func1 = funcItem1;
        item->func2 = funcItem2;
        return item;
    }

2、继承
    typedef struct Base {
        int a;
        void (*funcBase)();
    } Base;

    typedef struct Sub {
        Base base;
        int b;
        void (*funcSub)();
        ...
    } Sub;

3、多态
    Sub b;
    b.funcSub = func2;
    b.base.funcBase = func1;
    Base *a = &b;
    a->funcBase();  // 调用func1
    
    // 可以用赋值的方式
    b.base.funcBase = func2;
    a->funcBase();  // 调用func2
```

### std::remove_const
```c++
// remove const
template<class _Ty>
struct remove_const {	// remove top level const qualifier
	typedef _Ty type;
};

template<class _Ty>
struct remove_const<const _Ty> {	// remove top level const qualifier
	typedef _Ty type;
};

```

### tuple
```c++
// tuple<> 模板是 pair 模板的泛化，但允许定义 tuple 模板的实例，可以封装不同类型的"任意数量"的对象，因此 tuple 实例可以有任意数量的模板类型参数
#include <tuple>

// 例如
auto my_tuple = std::make_tuple(Name{"Peter"，"Piper"},42,std::string{"914 626 7890"});

// 函数模板 get<>() 可以返回 tuple 中的元素
// 基于索引
std::get<0>(my_tuple);
std::get<1>(my_tuple);
// 基于类型（必须包含不同的类型）
std::get<Name>(my_tuple);
std::get<int> (my_tuple);

template<typename ...Es>
class A{
    static const size_t size = sizeof...(Es);
    std::tuple<Es...> elems;
    A(Es... elems):elems{elems...}{}
};

// make_tuple
auto tup1 = std::make_tuple("Hello World!", 'a', 3.14, 0);

// tie
auto tup1 = std::make_tuple(3.14, 1, 'a');  
double a; int b;  char c;  
std::tie(a, b, c) = tup1;  

// tuple_cat 用于连接tuple
std::tuple<float, string> tup1(3.14, "pi");  
std::tuple<int, char> tup2(10, 'a');  
auto tup3 = tuple_cat(tup1, tup2);  

// get<i> 获取第 i 个元素的值
std::tuple<float, string> tup1(3.14, "pi");  
cout << get<0>(tup1);  

// tuple_element 获取tuple中特定元素数据类型
std::tuple_element<0, decltype(tup1)>::type i = 1.2;

// size 获取tuple中元素个数
std::tuple<float, string> tup1(3.14, "pi");  
cout << tuple_size<decltype(tup1)>::value;  

```

### 成员函数的隐形this参数
```c++
// 成员函数都隐含一个名为this的指针形参，并且它是该成员函数的第一个参数
class A{
public:
    void f();
};

```

### 类成员指针
```c++
// https://www.runoob.com/w3cnote/cpp-func-pointer.html

class A {
public:
    int a;
    int func1(int) {}
    void func2(int) {}
};

using Ta     = int A::*;
using Tfunc1 = int(A::*)(int);
using Tfunc2 = void(A::*)(int);
Ta     ffa = &A::a;
Tfunc1 ff1 = &A::func1;
Tfunc2 ff2 = &A::func2;
// 或
// int A::*ffa = &A::a;
// int(A::*ff1)(int) = &A::func1;
// void(A::*ff2)(int) = &A::func2;

A a;
(a.*ffa) = 3;
(a.*ff1)(2);
(a.*ff2)(2);
```

### 常量表达式
```c++
// 最初class中只能用enum声明常量
class {
    enum { value = N * Fac<N - 1>::value };
}

// 后来允许class内部static常量表达式
class {
    static const int value = N * Fac<N - 1>::value;
}

// Modern C++中，还可以使用constexpr，且不再局限于int
class {
    static constexpr auto value = N * Fac<N - 1>::value;    
}

// 加入inline，防止编译器为static分配内存
class {
    static inline constexpr auto value = N * Fac<N - 1>::value;
}

```

### 反射
```
1、c++ 反射方法
(1) 编译器支持
    C++20据说有编译期反射，但那是可能是N年后的事情了
    MSVC的CLR扩展支持运行期反射，但CLR不算标准C++
    clang的智能补全插件，修改一下能输出结构体信息，可以做个工具，但这只适合clang
(2) 运行期支持
    解析pdb文件进行反射(windows) 
    [PDBs on Linux] https://arvid.io/2018/05/05/pdbs-on-linux/
    动态反射，构建reflect类 （RTTR等）
(3) 代码生成
    用第三方工具，生成特殊的C++代码，自带反射
(4) MACRO(宏)
    利用c++基本语法实现
    对代码有侵入，需要手动填入信息
    [C++反射的方法与实现] https://zhuanlan.zhihu.com/p/29400757
    [C++用全局对象的构造函数实现反射机制] https://blog.csdn.net/panderang/article/details/80214521
    [200 行的 C++ 反射] https://www.clarkok.com/blog/2015/03/09/200-%E8%A1%8C%E7%9A%84-C-%E5%8F%8D%E5%B0%84/
    [C++ 中如何遍历对象的成员] https://www.zhihu.com/question/28849277
    [99 行实现 C++ 简单反射] https://zhuanlan.zhihu.com/p/158147380
    [c++如何实现反射功能] https://www.zhihu.com/question/62012225
    [谈谈C++如何实现反射机制] https://zhuanlan.zhihu.com/p/70044481
```

### type_traits
```c++
#include <type_traits>
// type_traits是C++11提供的模板元基础库
// type_traits可实现在编译期计算、判断、转换、查询等等功能
// type_traits提供了编译期的true和false

(1) integral_constant
// 该对象包含具有指定值的该整型类型的常量
std::integral_constant<int, 5>::value;       // 5
std::integral_constant<bool, true>::value;   // true

(2) true_type false_type
std::true_type::value;
std::false_type::value;

(3) is_same

(4) decay
// 获取它的原始类型
template<typename T>
typename std::decay<T>::type* Create(){
    typedef typename std::decay<T>::type U;
    return new U();
}

(5) conditional
std::conditional<true, int, double>::type   //= int

(6) decltype和auto
// decltype和auto可以实现模板函数的返回类型
template<typename F, typename Arg>
auto Func(F f, Arg arg)->decltype(f(arg)){
    return f(arg);
}

(7) result_of 
// result_of 在编译期推导出一个函数表达式的返回值类型
int f(int a, int b) {return a+b;}
template<typename Fn, typename ...Argvs>
typename std::result_of<Fn(Argvs...)>::type Func(Fn f, Argvs&&... argvs) {
    return f(argvs...);
}
Func(f, 2, 3);

(8) enable_if
template<bool, typename T = void> struct enable_if {};
template<typename T> struct enable_if<true, T>{ typedef T type; };
// 只有当第一个模板参数为 true 时，type 才有定义，否则使用 type 会产生编译错误

(9) declval
// 返回一个类型的右值引用，不管是否有没有默认构造函数或该类型不可以创建对象

(10) is_constructible
// 用于检查给定类型T是否是带有参数集的可构造类型
template <class T, class... Args>
struct is_constructible;

struct T { 
    T(int, int){}; 
};

std::is_constructible<T, int>::value // false
std::is_constructible<T, int, int>::value // true

(11) is_convertible
// 测试一种类型是否可转换为另一种类型
template <class From, class To>
struct is_convertible;

```

### is_function 简单实现
```c++
#include <type_traits>

template<typename T>
struct my_is_function : public false_type {};

template<typename T>
struct my_is_function<T()> : public true_type {};   // 普通函数
template<typename T>
struct my_is_function<T(*)()> : public true_type {}; // 函数指针

template<typename T, typename... Argvs>
struct my_is_function<T(Argvs...)> : public true_type {};  // 多参数
template<typename T, typename... Argvs>
struct my_is_function<T(*)(Argvs...)> : public true_type {};

void f(int);

my_is_function<decltype<f>>::value; // true

```

### 基类调用继承类接口
```c++
class A {
public:
    virtual void init() {
        cout << "A::init()" << endl;
    }
    void func() {
        init();
    }
};

class B {
public:
    void init() override {
        cout << "B::init()" << endl;
    }
};

B b;
b.func();   // B::init()
```

### 不完全类型
```c++
// "不完全类型"指在C++中有声明但又没有定义的类型

class A; // 声明，但没有定义

A* Create();
A* a = Create();
delete a;
// 上述代码可以编译通过，但会导致内存泄漏

// 解决办法
// 未定义的类型，sizeof(T)是0

typedef char T_must_complete_type[sizeof(A) == 0? -1 : 1];
T_must_complete_type test_; (void)test_;

```

### backtrace
```c++
// https://blog.csdn.net/jasonchen_gbd/article/details/44108405

// backtrace()是glibc（>=2.1）提供的函数，用于跟踪函数的调用关系
#include <execinfo.h>
int backtrace(void **buffer, int size);
char **backtrace_symbols(void *const *buffer, int size);
void backtrace_symbols_fd(void *const *buffer, int size, int fd);
```

### abi::__cxa_demangle
```c++
#include <cxxabi.h>
// __cxa_demangle来将backtrace_symbols返回的字符串逐个解析成可以方便看懂的字符串
```

### 字符编码 locale
```c++
// 设置编码
#include <locale.h>

(1) C函数设置全局locale
setlocale(LC_ALL, "zh_CN.utf8");

(2) C++ 设置全局locale
std::locale::global(std::locale("zh_CN.utf8"));

(3) 单独为wcout设置
std::wcout.imbue(std::locale("zh_CN.utf8"));

```

### basic_string
```c++
// https://www.byvoid.com/zhs/blog/cpp-string

// string并不是一个单独的容器，只是basic_string 模板类的一个typedef

typedef basic_string<char> string;
typedef basic_string<wchar_t> wstring;

// class basic_string
template <class charT, 
          class traits = char_traits<charT>,
          class Allocator = allocator<charT> >
class basic_string
{
//...
}
```

### char wchar_t
```c++
// 多字节字符(char)与宽字符(wchar_t)
// 多字节字符 char 不同的字符占不同的字节
// 英文字母'a'占一个字节，汉字'啊'占三个字节
// 宽字符 wchar_t 一个宽字符站固定的多个字节(linux是4个)
// 不管是英文还是中文

// 注意
// 在c标准中，选择 "多字节字符" 还是 "宽字符" 由对其执行的第一个操作设置(使用 fwide 函数进行检查)

// 转换方式
// fwide可以设置当前流定向，前提是未有任何的 I/O 操作，也就是当前流尚未被设置任何流定向

(1) 统一使用一种函数(推荐)
    printf或wprintf

(2) freopen清空定向流
    //重新打开标准输出流，清空流定向
    FILE* pFile=freopen("/dev/tty", "w", stdout);
    wprintf(L"wide freopen succeeded\n");

    //重新打开标准输出流，清空流定向
    pFile=freopen("/dev/tty", "w", stdout);
    printf("narrow freopen succeeded\n");
```

### char* wchar_t* 互转
```c++
typedef basic_string<char> string; 
typedef basic_string<wchar_t> wstring; 

#include <stdlib.h>
#include <string>

string ws2s(const wstring& ws) {
    std::locale::global(std::locale("en_US.utf8"));
    size_t n = wcstombs(NULL, ws.c_str(), 0);
    cout << "ws2s n " << n << endl;
    char buf[n];
    memset(buf, 0, n);
    if (wcstombs(buf, ws.c_str(), n) <= 0) {
        return "";
    }
    std::string res(buf, buf + n);
    return res;
}

wstring s2ws(const string& s) {
    std::locale::global(std::locale("en_US.utf8"));
    size_t sLen = s.size();
    wchar_t buf[sLen];
    wmemset(buf, 0, sLen);
    cout << "sLen " << sLen << endl;
    size_t ret = mbstowcs(buf, s.c_str(), sLen);
    cout << "ret " << ret << endl;
    if (ret <= 0) {
        return L"";
    }
    std::wstring res(buf, buf + ret);
    return res;
}

```

### printf wprintf
```c++
// https://www.eet-china.com/mp/a47123.html

// printf 打印多字符，wprintf 打印宽字符
wchar_t a = L'啊';
wprintf(L"%lc", a);
const wchar_t * b = L"啊aa";
wprintf(L"%ls", b);

// 注意
// 为了防止输出混乱，避免printf与wprintf混用
// wchar_t -> char，再进行打印
// 如果要打印wchar_t，推荐wcout

```

### sync_with_stdio
```c++
// 为了兼容性，c++的iostream 与 c的stdio公用一个buffer
// ios_base::sync_with_stdio(false) 让他们用不同的buffer
// 提前调用了ios::sync_with_stdio(false)，wcout和cout便不会通过stdio.h进行输入输出，而是自己管理缓冲区，就避免了wcout和cout不能混用的问题
// 注意：一旦iostream与stdio解绑，c++风格与c风格的io就无法保证其顺序安全性
```

### cin.tie(0)
```c++
// c++ 中 cin与cout是绑定的，因为可以保证cin之前会将cout输出缓冲区的数据刷新到输出文件中，不利于性能
std::ios_base::sync_with_stdio(false); // iostream stdio 解绑
std::cin.tie(0); // cin cout 解绑

// 所以，解绑后的c++ iostream与c的stdio性能一样
```

### cout wcout
```c++
// https://kc.kexinshe.com/t/81705

// Windows下printf/cout和wprintf/wcout可以混用，fwide函数是空函数

// Linux下printf/cout和wprintf/wcout不可以混用，流的宽窄取决于首次调用哪个函数或先调用fwide(stream, 1)还是fwide(stream, -1)，一旦确定即不可更改，除非重新打开

// iostream库在Linux下有一个坑，默认情况下wcout和cout不能混用。这是因为受到Linux下的C语言stdio.h库的掣肘

// cout wcout
std::wcout.imbue(std::locale("zh_CN.utf8"));
std::wcout << L"1" << std::endl;
std::cout << "2" << std::endl;
std::wcout << L"3" << std::endl;
std::cout << "2" << std::endl;
// 输出 1 3

std::ios_base::sync_with_stdio(false); // cout wcout print wprint 已解绑，都有自己的buffer
std::wcout.imbue(std::locale("zh_CN.utf8"));
std::wcout << L"1" << std::endl;
std::cout << "2" << std::endl;
std::wcout << L"3" << std::endl;
std::cout << "2" << std::endl;
// 输出 1 2 3 4
```

### wstring string 互相转换
```c++
// string 可看作 char[]
// wstring 可看作 wchar_t[]

#include <string>
#include <locale>
#include <codecvt>

std::string a = "啊aa";
cout << a.size() << endl;   // 5
std::wstring_convert<std::codecvt_utf8<wchar_t>> convert;
std::wstring b = convert.from_bytes(a); // string -> wstring
cout << b.size() << endl;   // 3        避免打印wstring
std::string c = convert.to_bytes(b);

```

### 容器 swap
```c++
// vector swap 为例
// swap 交换两个地址
vector<int> a(10, 1);
vector<int> b;
b.swap(a);
cout << a.size() << endl;   // 0
cout << b.size() << endl;   // 10
```

### 隐式转换与explicit
```c++
// explicit 关键字只能用于类内部的"构造函数"声明上
// explicit 关键字作用于"单个参数"的构造函数

// 当一个构造函数只有一个参数，而且该参数又不是本类的const引用时，这种构造函数称为"转换构造函数"
class T {
    T(){}       
    T(int a) {} // 转换构造
};

// 隐式转换背后发生了什么
std::string s = "aaa";
// c++11 之前
// 字符串属于 const char* 类型，string有相应的构造函数
// 生成临时string("aaa")
// 调用复制构造函数传给s
// 总结：整个过程创建了s和一个临时量，发生了一次拷贝。一共两个string

// c++11 之后
// 字符串属于 const char* 类型，string有相应的构造函数
// 调用构造函数生成临时string("aaa")
// 临时变量属于右值，调用移动构造函数创建s
// 临时变量的数据移动到s里
// 总结：整个过程创建了s和一个临时量，没有发生拷贝。一共两个string

// 大部分编译器已经做到的优化、c++17之后
// 直接调用string(const char*)构造函数创建s


// 自定义转换接口一共也就两个，转换构造函数和用户定义转换函数
struct T {
    T(int);          // 转换构造函数  int -> T
    T(const T1 &);   // 转换构造函数  T1  -> T

    // 自定义转换函数，返回转换类型，可以隐式转换
    operator int(); // T -> int
    operator T1();  // T -> T1
    // 从c++11开始explicit还可以用于用户定义的转换函数
};
```

### 智能指针
```
C++ 11的新特性中引入了三种智能指针，来自动化地管理内存资源

unique_ptr: 管理的资源唯一的属于一个对象，但是支持将资源移动给其他unique_ptr对象。当拥有所有权的unique_ptr对象析构时，资源即被释放

shared_ptr: 管理的资源被多个对象共享，内部采用引用计数跟踪所有者的个数。当最后一个所有者被析构时，资源即被释放

weak_ptr: 与shared_ptr配合使用，虽然能访问资源但却不享有资源的所有权，不影响资源的引用计数。有可能资源已被释放，但weak_ptr仍然存在。因此每次访问资源时都需要判断资源是否有效
```

### shared_ptr
```c++
#include <memory>

(1) 构造
    std::shared_ptr<string> a(new string("dasdas"));
    std::shared_ptr<string> a = std::make_shared<string>("dasdas")
    std::shared_ptr<vector<int>> a(new vector<int>(10));
    cout << a << endl;  //0x5633802bfe70

(2) 切片
    std::shared_ptr<int> a(new int [10] {1,2,3,4,5});
    int* pI = a.get();
    cout << *a << endl;     // 1
    cout << *(a+1) << endl; // 错误
    cout << a[0] << endl;   // 错误

    std::shared_ptr<int[]> a(new int [10] {1,2,3,4,5});
    cout << *a << endl;     // 错误
    cout << a[0] << endl;   // 1

    std::shared_ptr<std::vector<int>> vc = std::make_shared<vector<int>>(10,3);
    cout << (*vc)[1] << endl;
    cout << vc->operator[](1) << endl;
    cout << vc->size() << endl;

(3) 引用次数
    // shared_ptr多个指针指向相同的对象。shared_ptr使用引用计数，每一个shared_ptr的拷贝都指向相同的内存
    // 每使用他一次，内部的引用计数加1；每析构一次，内部的引用计数减1，减为0时，自动删除所指向的堆内存
    // shared_ptr内部的引用计数是线程安全的，但是对象的读取需要加锁

    std::shared_ptr<int> a = std::make_shared<int>(10);
    std::shared_ptr<int> & b = a; // 引用 不会改变计数
    std::shared_ptr<int> c = a; // 复制计数加一 不会改变计数
    cout << a.use_count() << endl;
    // 空指针的计数为0
    shared_ptr<int> a;      // 计数0
    shared_ptr<int> b = a;  // 计数0

(4) 自定义析构函数
    std::shared_ptr<int[]> a(new int[1], [](int* a){
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

(6) 禁止用auto_ptr
    // 智能指针所有权
    auto_ptr<string> a = auto_ptr<string>(new string("aaaa"));
    auto_ptr<string> b;
    b = a;   // 赋值导致了 a失去了所有权，b获得了所有权

    // STL有一条规定：std::auto_ptr 不能和容器混合使用
    // 原因是：容器里的元素使用的都是copy，而std::auto_ptr型数据copy后会发生拥有权转移。
    // 所以！！！auto_ptr几乎没用！！！

(7) 智能指针做参数传值更好
```

### unique_ptr 与 shared_ptr
```c++
// shared_ptr  允许多个指针指向同一个对象
// unique_ptr  独占所指向的对象，不能拷贝复制，unique_ptr销毁，其指向的对象也被销毁

// shared_ptr 和 unique_ptr 共有操作
shared_ptr<T> sp	// 空智能指针，可以指向类型为T的对象
unique_ptr<T> up	// 同上
p	                // 将p用作一个条件判断，若p指向一个对象，则为true
*p                  // 解引用p，获得它指向的对象
p->mem              // 等价于(*p).mem
p.get()             // 返回p中保存的指针，要小心使用。
swap(p,q)           // 交换p和q中的指针
p.swap(q)           // 同上

// shared_ptr 独有操作
make_shared<T> (args)	// 返回一个shared_ptr，指向一个动态分配的类型为T的对象
shared_ptr<T> p(q)	    // p是shared_ptr q的拷贝；此操作会增加q中的计数器
p=q	                    // p和q都是shared_ptr，所保存的指针必须能相互转换
                        // 此操作会递减p的引用计数，递增q的引用计数；若p的引用计数变为0，则将其管理的原内存释放
p.unique()	            // 若p.use_count()为1，返回true；否则返回false
p.use_count()           // 返回与p 共享对象的智能指针数量；可能很慢，主要用于调试

// unique_ptr 独有操作
unique_ptr<T> u1        // 空unique_ptr，可以指向类型为T的对象。u1会使用delete来释放它的指针
unique_ptr<T,D> u2      // 同上。u2会使用一个类型为D的可调用对象来释放它的指针
unique_ptr<T,D> u(d)    // 空unique_ptr，可以指向类型为T的对象，用类型为D的对象d代替delete
u=nullptr               // 释放u指向的对象，将u置为空
T* ptr = u.release()    // u放弃对指针的控制权，返回指针，并将u置空
                        // 注意release后需要自己管理内存 auto u = u2.release(); delete(u);
u.reset(q)              // 令u指向这个对象，原来u的内容会发生析构
u.reset(nullptr)

// unique_ptr不支持拷贝和赋值，如何拷贝或赋值unique_ptr
std::unique_ptr<int> a1(new int [10]);
std::unique_ptr<int> a2 = std::move(a1);    // unique_ptr实现了移动语义
unique_ptr<int> a3(a2.release());           // 赋予unique_ptr一个指针，必须先要释放一个unique_ptr的一个指针
a2.reset(a3.release());                     // 赋予unique_ptr一个指针，必须先要释放一个unique_ptr的一个指针

```

### shared_ptr 循环引用
```c++
(1) 循环引用
    class B;
    class A {
    public:
        std::shared_ptr<B> a;
        virtual ~A() {
            cout << "descontruct A" << endl;
        }
    };

    class B {
    public:
        std::shared_ptr<A> b;
        virtual ~B() {
            cout << "descontruct B" << endl;
        }
    };

    auto a = shared_ptr<A>(new A);
    auto b = shared_ptr<B>(new B);
    a->a = b;
    b->b = a;
    cout << a.use_count() << endl;  // 2
    cout << b.use_count() << endl;  // 2
    // 导致内存泄漏

(2) weak_ptr
    // weak_ptr是一种不控制对象生命周期的智能指针, 它是指向一个shared_ptr的管理对象
    shared_ptr<int> a;
    shared_ptr<int> b = a;  // 计数2
    shared_ptr<int> a;
    weak_ptr<int> b = a;  // 计数1

    // expire 判断智能指针是否已销毁
    if (!b.expire()) {}

    // lock 获得指向的智能指针
    shared_ptr<int> c = b.lock(); // a的计数+1

```

### make_unique
```c++
// c++ 14 直接提供 make_unique，以下是c++11实现方式
template<typename T, typename... Argvs>
std::unique_ptr<T> make_unique(Argvs&&... argvs) {
    return std::unique_ptr<T>(new T(std::forward<Argvs>(argvs)...));
}

// 为什么 make_unique 优于 unique_ptr(new)
// > 性能
//  unique_ptr<T>(new T) 造成的内存分配就是两次，一个是new，一个是内部的control block的内存分配
//  make版本内部使用allocated一下子分配包含new和control block的内存大小空间，并且这样加快程序运行速度、减小内存碎片的分配
// > 美观
//  减少new的使用



```

### 编译选项 -rdynamic
```
-rdynamic 是一个 连接选项 ，它将指示连接器把所有符号（而不仅仅只是程序已使用到的外部符号）都添加到动态符号表（即.dynsym表）里，以便那些通过 dlopen() 或 backtrace() （这一系列函数使用.dynsym表内符号）这样的函数使用
```

### 迭代器失效
```
数组型，链表型，树型

(1) 数组型数据结构
    该数据结构的元素是分配在连续的内存中，insert和erase操作，都会使得删除点和插入点之后的元素挪位置，所以，插入点和删除掉之后的迭代器全部失效，也就是说insert(*iter)(或erase(*iter))，然后在iter++，是没有意义的
    解决方法：erase(*iter)的返回值是下一个有效迭代器的值，iter=cont.erase(iter)

(2) 链表型数据结构
    对于list型的数据结构，使用了不连续分配的内存，删除运算使指向删除位置的迭代器失效，但是不会失效其他迭代器
    解决办法两种，erase(*iter)会返回下一个有效迭代器的值，或者erase(iter++)

(3) 树形数据结构
    使用红黑树来存储数据，插入不会使得任何迭代器失效；删除运算使指向删除位置的迭代器失效，但是不会失效其他迭代器.erase迭代器只是被删元素的迭代器失效，但是返回值为void
    所以要采用erase(iter++)的方式删除迭代器
```