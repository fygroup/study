C++笔记
|编译器为我们做了大量的优化工作，不要以为什么都理所应当

```
#include <iostream>
#include <string>
#include <fstream>

读写文件
fstream on;
on.open(string.c_str(),ios::in)
on.open(string.c_str(),ios::out)

字符转int atoi
int转字符 itoa


char** 转 string

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

                                                   //数组应该在.cpp文件里面初始化

}

注意 静态变量是类级别的！！！！ 需要用类名去定义（无论这个变量是否在private中）
例如: int my::b=3;要定义在全局变量里面


std:cerr << 错误输出


把一个成员函数声明为const可以保证这个输入的成员函数不修改数据成员，但是，如果据成员是指针，则const成员函数并不能保证不修改指针指向的对象


//----const------------------------
顶层const

底层const


c++中有一块const内存，并且不同变量，一样的内容，他们的指针地址是一样的，凡是const的变量都在const内存中
//---重写 重载 重定义--------------
函数重载是指在一个类中声明多个名称相同但参数列表不同的函数，这些的参数可能个数或顺序，类型不同，但是不能靠返回类型来判断。
函数重写是指子类重新定义基类的虚函数。

c++中可以对const进行赋值
class my{
	const int a;
}

my::a

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
#include <memory> 智能指针 auto_ptr<int> a(new int [10]) unique_ptr shared_ptr
//---智能指针---------------------
auto_ptr具有所有权，即一个地址只有一个智能指针
int* p_reg = new int [1024];
cout<<(void*)p_reg<<endl;
auto_ptr<int> pshare(p_reg);
cout<<&pshare<<endl;
cout<<pshare.get()<<endl;
auto_ptr<int> pshare1=pshare;
cout<<&pshare<<endl;
cout<<pshare.get()<<endl;
cout<<&pshare1<<endl;
cout<<pshare1.get()<<endl;
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
//---typename 函数---------------
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
TiXmlNode* element = Node->NextSiblingElement() //下一个兄弟节点
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
#define CC_AddFun(var,funname) {var = cc_type_convert(var,dllhelp.GetDllFunAddress(funname));if(!var){dllhelp.Close();return -1;}}
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
if (0==access(filepath,0))
	cout << "no file" << endl;
00——只检查文件是否存在
02——写权限
04——读权限
06——读写权限

//---得到工作路径-----------------
/proc/self/exe 它代表当前程序
int nRet = readlink("/proc/self/exe", pPath, nLen - 1);
//nRet表示返回字符串的长度

//---改变工作目录-----------------
chdir();

//---重定向函数-------------------
if(NULL == freopen("xxxxx","a+",stdout)){
	return -1;
}

//---wait waitpid---------------------
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
//---map 判断键值是否存在-----------------
（1）
pair<map<int, string>::iterator, bool> Insert_Pair;  
Insert_Pair = mapStudent.insert(pair<int, string>(1, "student_one"));  
if(Insert_Pair.second == true)  
	cout<<"Insert Successfully"<<endl;  
（2）
map<>::iterator it = m.find('key');
if (it == m.end()) #不存在
//---插入键值对---------------------------
1.map[键] = 值；直接赋值。 这种方式：当要插入的键存在时，会覆盖键对应的原来的值。如果键不存在，则添加一组键值对。
 2.map.insert()；这是map自带的插入功能。如果键存在的话，则插入失败，也就是不插入。 使用insert()函数，需要将键值对组成一组才可以插入
//----map 排序--------------------------------
struct cmpkeylen{
bool operator()(const string & a, const string & b){
return(a.length() > b.length());
}
};
map<sting,int,cmpkeybylen> mymap; //注意 第三个参数是个函数对象，c++ 11中很多库函数都是函数对象
//---sleep------------------------------
# include <unistd.h>
//---openmp-----------------------------
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
//---iterator-----------------------------------
vector<int> a;
vector<int>::iterator i;
for(i=a.begin();i!=a.end();i++){}
const vector<int> a;
vector<int>::const_iterator i;  //对于const 必须用const_iterator !!!
for(i=a.begin();i!=a.end();i++){}
//---class中的静态函数----------------------------
静态函数只要在定义的时候需要static关键字，实现的时候就不需要了，否则会报错。
//---class 中的 const----------------------------
C++中，const 修饰的参数引用的对象，只能访问该对象的const函数，因为调用其他函数有可能会修改该对象的成员，
编译器为了避免该类事情发生，会认为调用非const函数是错误的。
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
 Base中声明test时加上const，即void test() const
!!!!注意当函数后面加了const时，返回引用和指针时要加const，但是返回非引用可以不加
const mat & func()const{...}
const mat * func()const{...}
mat func()const{...}
//---引用 与 const----------------------------------
int func(){
    return(2);
}

void func1(int & a){
}

int main()  
{  

func1(func()); //错误 func() 返回的是临时变量，func1没法确定是否可修改

}
//改正1
int a = func();
func1(a);
//改正2
void func1(const int & a){
}
//改正3
void func1(int a){
}
//---initializer_list-------------------------------
void g(std::vector<int> const &items){}; 
void g(std::list<int> const &items){};
g({ 1, 2, 3, 4 }); //会报错  编译器分不清 vector还是list

//对于{}的固定数组initializer_list更合适
void g(std::vector<int> const &items){}; 
void g(std::list<int> const &items){}; 
void g(std::initializer_list<int> const &items){}; 
g({ 1, 2, 3, 4 });
initializer_list不能修改，更符合参数的特点
//---%------------------------------------------------
% is only defined for integer types. That's the modulus operator.
<cmath> fmod(m,n);
//---string------------------------------------------
string a("dadsadsada");
a.find_first_not_of("dda");   // 在字符串中查找第一个与str中的字符都不匹配的字符，返回它的位置。
a.find_first_of("dda");   // 在字符串中查找第一个与str中的字符匹配的字符，返回它的位置。
//---map com-------------------------------------------
typedef struct classCom{
	bool operator () (const int & a, const int & b)const{
		return(b-a);
	}
}classCom;
map<int,string,classCom> a;
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

//---目录操作----------------------------------------
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
//---function --------------------------------------------------
#include <functional>
void func1(int){....}

std::function<void(int)> func = func1;
std::function<void(int)> func = [](int)->void{....}  ([](int){....})

struct A
{
    void func(int){.....}
}
std::function<void(A&, int)> my = &A::func; //注意此处必须为struct
A a;
my(a,2);

//---bind 一般和function一起用---------------------------------------------------------
void func(int i, int j){}

std::function(void(int,int)) f = std::bind(&func,std::placeholders::_1,3);
绑定class成员函数，必须是public
class A
{
public:
   void func(int i);
};
A a;
std::function<void(int)> f = std::bind(&A::func,&a,std::placeholders::_1);
f(3);

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

//---互斥 条件变量-----------------------------------
pthread_mutex_t mutex;
pthread_mutex_init()  
pthread_mutex_lock()   锁定互斥锁，如果尝试锁定已经被上锁的互斥锁则阻塞至可用为止
pthread_mutex_unlock() 	释放互斥锁
pthread_mutex_destory()  互斥锁销毁函数


1.条件变量创建
静态创建：pthread_cond_t cond=PTHREAD_COND_INITIALIZER;
动态创建：pthread_cond _t cond;
  pthread_cond_init(&cond,NULL);
其中的第二个参数NULL表示条件变量的属性，虽然POSIX中定义了条件变量的属性，但在LinuxThread中并没有实现，因此常常忽略。
2.条件等待
pthread_mutex_t mutex=PTHREAD_MUTEX_INITIALIZER;
pthread_mutex_lock(&mutex);
while(条件1)
  pthread_cond_wait(&cond,&mutex);
函数操作1;
pthread_mutex_unlock(&mutex);
当条件1成立的时候，执行pthread_cond_wait(&cond,&mutex)这一句，开放互斥锁，然后线程被挂起。当条件1不成立的时候，跳过while循环体，执行函数操作1，然后开放互斥锁。
3.条件激发
pthread_mutex_lock(&mutex);
函数操作2;
if(条件1不成立)
pthread_cond_signal(&cond);
pthread_mutex_unlock(&mutex);
先执行函数操作2，改变条件状态，使得条件1不成立的时候,执行pthread_cond_signal(&cond)这句话。这句话的意思是激发条件变量cond，使得被挂起的线程被唤醒。
pthread_cond_broadcast(&cond);
这句话也是激发条件变量cond，但是，这句话是激发所有由于cond条件被挂起的线程。而signal的函数则是激发一个由于条件变量cond被挂起的线程。
4.条件变量的销毁
pthread_cond_destroy(&cond);
在linux中，由于条件变量不占用任何资源，所以，这句话除了检查有没有等待条件变量cond的线程外，不做任何操作。
5.pthread_cancle
线程取消需要设置取消点，比如在无限循环中加入pthread_testcancel（）
然后在pthread_join。
6、pthread_exit()和return类似就是推出的作用，不涉及资源释放。
7、pthread_kill()
8、线程退出时 回调函数
void rtn(void)
{
...
}

{
pthread_cleanup_push((void*)rtn, NULL);
....
pthread_cleanup_pop(1);
}

//---pthead------------------------------------------------------------------
int pthread_create(pthread_t *restrict tidp,
                              const pthread_attr_t *restrict attr,
                              void *(*start_rtn)(void *),  //注意函数必须是 void* func(void*){}
                              void *restrict arg);
//---函数转换-----------------------------------------------------
typedef void* (*func)(void*);

void f(){
}
func new = (func)&f;
//---class 多线程 调用成员函数------------------------------------
class Myclass
{
public:
    pthread_t id;
    static Myclass* cur;
    static void* callback(void*){
	cur->func();
	return(NULL);
    }
    void func();
    void start(){
	cur = this;
	pthread_create(&id,NULL,callback,NULL);
    
    }
}
Myclass* Myclass::cur = NULL;

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
	x = A(5);      //此时新建A(5)，然后在销毁，a已经不存在了！
	x.a[0] = 'a';
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

//--------sort----------------------------------------------------------------------
把list改成vector因为list的iterator不是random的而std::sort需要random的iterator
```

---
#### container_of(已知结构体type的成员member的地址ptr，求解结构体type的起始地址)
```
#define container_of(ptr,type,member) ({\
const typeof(((type*)0)->member)  *_mptr = (ptr);
(type*)((char*)_mptr - offsetof(type,member));  })

#define offsetof(type,member) ((size_t)&((type*)0)->member)
```

---
#### 特化 偏特化
(1)模板函数
```
template<typename T, class N> void func(T num1, N num2)
{
    //cout << "num1:" << num1 << ", num2:" << num2 <<endl;
}
```

(2)模板类
```
template<typename T, class N> class Test_Class
{
public:
    static bool comp(T num1, N num2)
    {
        return (num1<num2)?true:false;
    }
};
```

(3)全特化和模板函数
```
template<> void func(int num1, double num2)
{
    cout << "num1:" << num1 << ", num2:" << num2 <<endl;
}
```

(4)偏特化和模板类
```
template<typename N> 
class Test_Class<int, N>
{
public:
    static bool comp(int num1, double num2)
    {
        return (num1<num2)?true:false;
    }
```
特化 > 偏特化 > 模板类

---
#### 处理模板化基类内的名称
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

---
#### for_each waitpid
```
#include <algorithm>
#include <sys/wait.h>
for_each(vec.begin(),vec.end(),[](pid_t & pd){waitpid(pd,NULL,0);});
```

---
#### 模板元编程 type mapping

---
#### pthread_cond_timedwait
```
pthread_cond_timedwait (pthread_cond_t * _cond,pthread_mutex_t * _mutex,_const struct timespec * _abstime);
//比函数pthread_cond_wait()多了一个时间参数，经历abstime段时间后，即使条件变量不满足，阻塞也被解除
```

---
#### class static初始化
```
class my
{
public:
    typedef struct _MY{}MY;
    static list<MY*> a;
};

list<my::MY*> my::a;  
```

---
#### operator重载template
```
template<typename T>
class {
    template<typename T1>
    friend my<T1> & operator *(my<T1> & my1, my<T1> & my2){  //friend 
	    my1.i *= my2.i;
	    return(my1);
    }
}
```

---
#### 函数指针
```
typedef (int*)(*func)(int,char);
int* myfunc(int,char){}
func = myfunc;
//&(func) 等价于 func
```

---
#### static变量
static 变量最好都写在cpp文件中，除非hpp用到的那个static变量

---
#### makefile .o 文件有依赖时，是有顺序的
```
all: fz16.o fastqz.o libzpaq.o FastqReader.o FileOpt.o muti_Process.o
	$(CXX) -o $(TARGET) $^ $(FLAGS) $(LIBS)
fastqz.o依赖于muti_Process.o，muti_Process.o要写在fqstqz.o的后面
```

---
#### ostringstream
```
std::ostringstream str;
str << "abc" << 2 << "dda";
//格式化一个字符串，但通常并不知道需要多大的缓冲区
```

---
#### 构造函数私有化
对于class本身，可以利用它的static公有成员，因为它们独立于class对象之外，不必产生对象也可以使用它们

---
#### 获取文件绝对路径
```
realpath(file_name, abs_path_buff)
返回值为0表示错误
```

---
#### 判断文件夹是否存在，不存在创建文件夹
```
if (access(tarDir,F_OK)!=0){
    ASSERT_ERROR(mkdir(tarDir,S_IRWXU),"mkdir tmp wrong");
}
```

---
#### 文件夹操作
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
```

---
#### 指针数组
```
const char* a[5] = {"aa","ab","ada","dadad","fddgds"};
```

---
#### 模板特例化
```
template<>
class a<int>{};
```

---
#### typename
```
//在类中用未定义的T,必须以下这么用！！！
typename base<T>::Nest temp;
typedef typename iterator_traits<T>::value_type value_type;
valuetype tmp(xxx)
```

---
#### 虚函数
继承必须是指针或引用才能动态编译。

---
#### 模板类初始化(构造函数里)
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

---
#### 优先队列
```
priority_queue<element> seq;
seq.push(element(4,"aaa"));
seq.push(element(2,"bbb"));
seq.push(element(5,"ccc"));

while(!seq.empty()){
    element c = seq.top();  //注意这里是值，不能是引用，如果要提高性能，可以把element改为指针
    cout << c.i << endl;
    cout << c.name << endl;
    seq.pop();
}
```

---    
#### deque queue array
deque是双端队列
queue是容器适配器，底层由deque存储
array长度固定

---
#### 类函数指针
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

---
#### 静态类
静态类所必须的初始化在类外进行（不应在.h文件内实行），而前面不加static，以免与外部静态变量(对象)相混淆

---
#### 参数传递(string &)
```
void func(string a){} //此处 不能是&！！！！！
func("aaaa");
```

--- 
#### explicit
C++提供了关键字explicit，可以阻止不应该允许的经过转换构造函数进行的隐式转换的发生, 

---
#### access(目录是否存在)
```
#include <unistd.h>
int access(const char * pathname, int mode)
成功执行时，返回0。失败返回-1
R_OK      测试读许可权
W_OK      测试写许可权
X_OK      测试执行许可权
F_OK      测试文件是否存在
```

---
#### sort
```
bool compare(int & a,int & b)
int a[20]={2,4,1,23,5,76,0,43,24,65};
sort(a,a+20,compare);
```

---
#### 函数对象
```
#include <functional>  //c++11
template<typename T, typename... Args> //T 返回 argvs参数
static void forkRun(function<T(Args...)> func, Args... args);
```

---
#### 智能指针转换
```
//初始化
shared_ptr<int> sptr1( new int );
// 使用make_shared 来加速创建过程
// shared_ptr 自动分配内存，并且保证引用计数
// 而make_shared则是按照这种方法来初始化
shared_ptr<int> sptr2 = make_shared<int>( 100 );

//普通指针到智能指针的转换
int* iPtr = new int(42);
shared_ptr<int> p(iPtr);
//智能指针到普通指针的转换
int* pI = p.get();

//shared_ptr多个指针指向相同的对象。shared_ptr使用引用计数，每一个shared_ptr的拷贝都指向相同的内存。每使用他一次，内部的引用计数加每析构一次，内部的引用计数减1，减为0时，自动删除所指向的堆内存。shared_ptr内部的引用计数是线程安全的，但是对象的读取需要加锁。
//例如：
void func(shared_ptr<int> & a){       //引用 不会改变计数
	cout << a.use_count() << endl;
}
void func(shared_ptr<int> a){		//复制 会改变计数
	cout << a.use_count() << endl;
}
shared_ptr<int> a = make_shared<int>(10);
cout << a.use_count() << endl;	
func(a);
cout << a.use_count() << endl;
```

---
#### lambda
```
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
```

---
#### std::sort
```
vector<xxx> a;
sort(a.begin(),a.end(),[](const xxx & x, const xxx & y){return(x>y);}); //注意const!!!

//qsort
int x[10];
qsort(x,10,sizeof(int),func);
int func(const void* a, const void* b){
	return((*(int*)a)-(*(int*)b));
}
```

---
#### const char*(初始化)
c++允许先初始化再赋值
```
const char* a；
a = "dafgsfaaafag";
```

---
#### limits
```
#include <limits>
numeric_limits<double>::max() 
numeric_limits<double>::min() 
```

---
#### execv
```
const char* job[] = {"sh","-c","echo \'fafafafa\'",NULL};
execv("/bin/sh",(char* const*)job); //注意(char* const*),而不是(const* char*)!!!
```

---
#### random
```
#include <random>
std::default_random_engine e;
std::uniform_real_distribution<int> u(0,100)  //随机数分布
int a = u(e) ;   //产生随机数
```

---
#### sstream
```
#include <sstream>
default_random_engine e;
uniform_real_distribution<int> u(numeric_limits<int>::min(), numeric_limits<int>::max);
stringstream ss;
ss << u(e) << ".sm";
string x(ss.str());
const char* y = ss.str().c_str();
```

---
#### 多参数
```
void func(){} 
template<typename T, typename... Args> //T 返回 argvs参数
void func(T value, Args... args){
    cout << value << endl;
    func(args...);
}
```

---
#### \_\_func\_\_
当前函数名称

---
#### 函数指针map
```
typedef void(*f)(int,int);
void add(void){}
map<void*,string> mp;
mp[(void*)add] = "add";
```

---
#### 进程间通信
Posix函数有下划线分隔，SystemV函数没有

(1)POSIX IPC：
```
//涉及的库
#include <semaphore.h>
#include <sys/mman.h>
#include <sys/stat.h>        /* For mode constants */
#include <fcntl.h>           /* For O_* constants */

//信号量
sem_open(const char*)   sem_close(sem_t*)   sem_unlink(const char*)  
sem_wait(sem_t*)   sem_post(sem_t*)
//共享内存
int fd=shm_open(const char*)  shm_unlink(const char*) close(fd)
ftruncate(fd,10)
mmap()
```

(2)System V IPC:
```
#include <sys/sem.h>
#include <sys/shm.h>
#include <sys/types.h>
#include <sys/ipc.h>
//信号量
ftok   //ftok把一个已存在的路径名和一个整数标识符转换成一个key_t值，称为IPC键值
semget  semctl semop
shmget  shmctl shmat shmdt
```

---
#### 信号量
POSIX信号量与System V信号量
```
都是用于线程和进程同步的。
Posix信号量是基于内存的，即信号量值是放在共享内存中的，与文件系统中的路径名对应的名字来标识的。
System v信号量测试基于内核的，它放在内核里面。
```

(1)POSIX信号量
```
//一个进程创建POSIX信号量
#include <semaphore>
#define FILE_MODE (S_IRUSR|S_IWUSR|S_IRGRP|S_IROTH)

int main(){
    sem_unlink("file");  //防止所需的信号量已存在
    sem_t* mutex;
    if (mutex = sem_open("file",O_CREAT|O_EXCL,FILE_MODE,1) == SEM_FAILED){
        error("mutex");
        exit(-1);
    }
    sem_close(mutex);   //关闭
}
//另外一个进程运用POSIX信号量
#include <semaphore.h>

int main(){
    sem_t* mutex;
    if ((mutex = sem_open("file",0)) == SEM_FAILED){ //打开信号量
        error;
    }
    sem_wait(mutex);        //加锁
    ...
    sem_post(mutex);        //释放锁
}
```

(2)System V信号量
```
#include <sys/sem.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

int sem_id;
int set_semvalue(){
    union semun sem_union;
    sem_union.val = 1;
    if (semctl(sem_id,0,SETVAL,sem_union) == -1) return(0);
    return(1);
}
void del_semvalue(){
    union semun sem_union;
    if (semctl(sem_id,0,IPC_RMID,sem_union) == -1) perror();
}
int semaphore_p(){
    struct sembuf sem_b;
    sem_b.sem_num = 0;
    sem_b.sem_op = -1; //P()
    sem_b.sem_flg = SEM_UNDO;
    if (semop(sem_id,&sem_b,1) == -1)return(0);
    return(1);
}
int semaphore_v(){
    struct sembuf sem_b;
    sem_b.sem_num = 0;
    sem_b.sem_op = 1; //V()
    sem_b.sem_flg = SEM_UNDO;
    if (semop(sem_id,&sem_b,1) == -1)return(0);
    return(1);
}

int main(){
    key_t key = ftok("file",3);
    sem_id = semget(key,1,0666|IPC_CREAT); //创建信号量
    if (!set_semvalue()) perror();  //初始化信号量

    if (!semaphore_p()) perror;  //进入临界区
    ...
    if (!semaphore_v()) perror; //离开临界区

    del_semvalue();
}
```

---
#### 共享内存
分两种

System V的shmget()得到一个共享内存对象的id，用shmat()映射到进程自己的内存地址
POSIX的shm_open()打开一个文件，用mmap映射到自己的内存地址

<img src="../picture/7.png" alt="shm_open+mmap" height=300 width=500/>

注意：以上两种方式要用信号量同步

(1)shmget
```
//进程一 read
#include<sys/shm.h>
#include <sys/types.h>
#include <sys/ipc.h>
#define MEM_KEY (1234)

typedef struct _shared{
    int text[10];
}shared;

int main(){
    key_t key = ftok("file",0x03); //proj_id是一个1－255之间的一个整数值，典型的值是一个ASCII值
    int shmid = shmget((key_t)MEM_KEY, sizeof(shared),0666|IPC_CREAT|IPC_EXCL); //创建共享内存,如果存在则报错
    //int shmid = shmget(key,sizeof(shared,IPC_CREAT|0666));
    if (shmid == -1) perror();
    void* shm = shmat(shmid,0,0); //连接当前进程地址空间
    if（shm == (void*)-1）perror();
    shared* my = (shared*)shm;
    printf("%d\n",my->text[1]);
    if (shmdt(shm) == -1) perror();    //把共享内存从当前进程分离
    if (shmctl(shmid,IPC_RMID, 0) == -1) perror //删除共享内存
}

//进程二 write
#include<sys/shm.h>
#define MEM_KEY (1234)

typedef struct _shared{
    int text[10];
}shared;

int main(){
    int shmid = shmget((key_t)MEM_KEY, sizeof(shared),0666|IPC_CREAT); //创建共享内存
    if (shmid == -1) perror();
    shm = shmat(shmid, 0, 0);
    if（shm == (void*)-1）perror();
    shared* my = (shared*)shm;
    my->text[1] = 5;
    if (shmdt(shm) == -1) perror();
}
```
(2)shm_open+mmap
```
//server
#include <sys/mmap.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <fcntl.h>
#define  FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

typedef struct __ST{
     char text[5];
}ST;

int main(){
    shm_unlink("file");  //防止file已存在
    int fd = shm_open("file",O_RDWR|O_CREAT,FILE_MODE);
    if (fd == -1) perror();
    ftruncate(fd,sizeof(ST));
    ST* ptr;
    ptr = mmap(NULL,sizeof(ST),PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);
    if (ptr == SEM_FAILED) perror();
    ptr->text[1] = 'a';
    close(fd);

}

//client
#include <sys/mman.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <fcntl.h>
typedef struct _ST{
    char text[5];
}ST;

int main(){
    int fd = shm_open("file",O_RDWR,FILE_MODE);
    ptr = mmap(NULL,sizeof(ST),PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);
    printf("%c",ptr->text[1]);
    close(fd);
}
```

---
#### pipe(无名管道)
父子进程互相传递信号
```
#include <unistd.h>

int fd[2]; //0:读  1:写
int ret = pipe(fd);
if (ret==-1)perror();
pid_t pt = fork();
if (pt>0){
    close(fd[0]);       //关闭读
    write(fd[1]...);    
}else{
    close(fd[1]);       //关闭写
    read(fd[0]...);
}
```

---
#### FIFO(有名管道)
不同进程互相传递信号
```
#include <sys/stat.h>
mkfifo("file",0755);
int fd = open("file",O_RDONLY); //O_WRONLY
close(fd);
```

---
#### 类静态成员初始化
静态成员需要一开始初始化，
```
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

---
#### sys/time.h
```
#include <sys/time.h>
struct timeval tv;
gettimeofday(&tv,NULL);
sprintf(tarDir,"%s/tmp%ld%ld/",tarDir,tv.tv_sec,tv.tv_usec);
```

---
#### 子进程退出
推荐用_exit(1)
```
pid_t pt;
int status;
waitpid(pt,&status,0);
WEXITSTATUS(status) //获取退出值
WIFEXITED(status)   //判断是否正常推出
WIFSIGNALED(status) //判断是否被杀死
```

---
#### fd设置
fd非阻塞
```
int val;
if ((val = fcntl(fd[0], F_GETFL, 0)) < 0){
	cout << "[error]: fcntl" << endl;
	exit(-1);
}
val |= O_NONBLOCK;                          //先取出val,再设置
if (fcntl(fd[0], F_SETFL, val) < 0){
	cout << "[error]: fcntl" << endl;
	exit(-1);
}
```

---
#### 拷贝构造函数
```
A a;    //构造
A b(a); //拷贝构造创建对象b
A b=a;  //拷贝构造创建对象b
//强调：这里b对象是不存在的，是用a 对象来构造和初始化b的！！
```

---
#### 右值引用的深思
编译选项-fno-elide-constructors用来关闭返回值优化效果

1、值优化的重要性
```
//值传递实例（关闭值优化）
class A
{
    A(){
        cout << "construct" << endl;
    }

    A(const A& a){
        cout << "copy" << endl;
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
```

2、指针悬挂与深拷贝
```

```

3、左值引用与常量左值引用

```
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
```

4、右值引用与类型推导判断

```
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


```

5、续命与减少拷贝

```
int func(){
    return 1;
}

int x = func();    //需要拷贝
int&& x = func();   //不需要拷贝 仅仅是move
```

6、移动语义（移动构造函数）

```
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
```

7、完美转发

```
#include <utility>
void func(string& str){}
void func(string&& str){}
template<typename T>
void pfunc(T&& str){
    func(forward<T>(str)); //自动判断str是左值还是右值
}
```

8、std::move(移动转移所有权)
```
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

---
#### RTTI
```
//type_info类
//typeid操作符
//type_index类
#include <typeinfo>
#include <typeindex>

const std::type_info &tiInt = typeid(int);
std::cout << "tiInt.name = " << tiInt.name() << std::endl;

class A{};
std::unordered_map<std::type_index, std::string> type_names;
type_names[std::type_index(typeid(A))] = "A";
```

---
#### decltype
```
decltype和auto都可以用来推断类型，但是二者有几处明显的差异：
1.auto忽略顶层const，decltype保留顶层const；
2.对引用操作，auto推断出原有类型，decltype推断出引用；
3.对解引用操作，auto推断出原有类型，decltype推断出引用；
4.auto推断时会实际执行，decltype不会执行，只做分析
```
---
#### socket
https://files-cdn.cnblogs.com/files/life2refuel/socket%E7%AE%80%E5%8D%95%E7%BC%96%E7%A8%8B%E6%8C%87%E5%8D%97.pdf

---
#### 多路复用与异步IO
1、多路复用(同步非阻塞IO)
select #(select实例)[https://blog.csdn.net/u010155023/article/details/53507788]
epoll #(epoll实例)[https://blog.csdn.net/davidsguo008/article/details/73556811]
```
当某一进程调用epoll_create方法时，Linux内核会创建一个eventpoll结构体，这个结构体中有两个成员与epoll的使用方式密切相关。eventpoll结构体如下所示：
struct eventpoll{
    ....
    /*红黑树的根节点，这颗树中存储着所有添加到epoll中的需要监控的事件*/
    struct rb_root  rbr;
    /*双链表中则存放着将要通过epoll_wait返回给用户的满足条件的事件*/
    struct list_head rdlist;
    ....
};

每一个epoll对象都有一个独立的eventpoll结构体，用于存放通过epoll_ctl方法向epoll对象中添加进来的事件。这些事件都会挂载在红黑树中,如此，重复添加的事件就可以通过红黑树而高效的识别出来(红黑树的插入时间效率是lgn，其中n为树的高度)。

而所有添加到epoll中的事件都会与设备(网卡)驱动程序建立回调关系，也就是说，当相应的事件发生时会调用这个回调方法。这个回调方法在内核中叫ep_poll_callback,它会将发生的事件添加到rdlist双链表中。
在epoll中，对于每一个事件，都会建立一个epitem结构体，如下所示：
struct epitem{
    struct rb_node  rbn;//红黑树节点
    struct list_head    rdllink;//双向链表节点
    struct epoll_filefd  ffd;  //事件句柄信息
    struct eventpoll *ep;    //指向其所属的eventpoll对象
    struct epoll_event event; //期待发生的事件类型
}

当调用epoll_wait检查是否有事件发生时，只需要检查eventpoll对象中的rdlist双链表中是否有epitem元素即可。如果rdlist不为空，则把发生的事件复制到用户态，同时将事件数量返回给用户。


```

2、异步IO

---
#### 容器中存不同类型变量
```

```

---
#### 友元类与友元变量
```
//友元类
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

//友元函数
class Node
{
private:
    int data;
public:
    friend int BinaryTree::find(); //这样find方法就可以直接访问Node中的私有成员变量了，而BinaryTree中的其他的方法去不能够直接的访问Node的成员变量
};

//友元的继承
class A
{
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





#### fuse 用户空间文件系统
[fuse架构](../picture/fuse架构.png)
```
1、用户态app调用glibc open接口，触发sys_open系统调用。
2、sys_open 调用fuse中inode节点定义的open方法。
3、inode中open生成一个request消息，并通过/dev/fuse发送request消息到用户态libfuse。
4、Libfuse调用fuse_application用户自定义的open的方法，并将返回值通过/dev/fuse通知给内核。
5、内核收到request消息的处理完成的唤醒，并将结果放回给VFS系统调用结果。
```

