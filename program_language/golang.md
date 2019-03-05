#### 参数传递
```
函数调用参数均为值传递，不是指针传递或引用传递。经测试引申出来，当参数变量为指针或隐式指针类型，参数传递方式也是传值（指针本身的copy）
```

#### 数组与slice
```
//数组是初始化就定长
var a [5] int
a:=[5]int{1,2}
a:=[...]int{2:3,3:4}

func test(a [] int)  //值传递
func test(a *[] int)

test(a)
test(&a)

//slice切片又称为变长数组
var a [] int
a:=[]int{1,2,3}
a:=make([]int,4)

func test(a [] int) //slice按引用传递

test(a)

append(a,1)

//数组与slice转换
var a [6] int
var b [] int = a[:]

//slice、map、interface、channel都是按引用传递
```

//一些功能


#### map
```
//声明
var my map[string] int;
my := make(map[string] int)
//赋值
my['dada'] = 1

rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
csharpRating, ok := rating["C#"]
if ok {
}

m := make(map[string]string)
m["Hello"] = "Bonjour"
m1 := m
m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了

//map 指针
myMap := make(map[string]*Test)
var x Test
x.a = 2
x.b = "abc"
myMap["a"] = &x
(*myMap["a"]).a = 10

```

#### make、new
make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。
```
list:= make([]int,5,10) //初始化5个int的数组，可扩展性为10个
len(list)  //5
cap(list)  //10
```

#### 指针
```
type A struct{x int}
a:=new(A)
a.x = 1     //正确
(*a).x = 1  //正确
```

#### defer panic recover
```
defer func(){
    if a:=recover();a!=nil{
        .....
    }
}()

panic("dadad")

```

#### init
初始化导入的包（递归导入）
对包块中声明的变量进行计算和分配初始值
执行包中的init函数
```
package main

import "fmt"

var _ int64=s()

func init(){
    fmt.Println("init function --->")
}

func s() int64{
    fmt.Println("function s() --->")
    return 1
}

func main(){
    fmt.Println("main --->")
}
//输出
function s() --->
init function --->
main --->
```

#### import
```
import(
    . "fmt"
)
import(
    f "fmt"  // f.Println()
)
import (
    "database/sql"
    _ "github.com/ziutek/mymysql/godrv"   //不使用包的函数 只是用init()
)

```

#### struct
```
type person struct{
    name string
    age int
}

p:= person{"malx",25}
p:= person{name:"malx",age:25}

type Human struct {
    name string
    age int
    weight int
}
type Student struct {
    Human // 匿名字段，那么默认Student就包含了Human的所有字段
    speciality string
}

p:=Student{Human{"malx",25},"dada"}
p:=Student{Human:Human{"malx",25},speciality:"dada"}

```

#### 复杂结构
```
type Values map[string][]string
func (v Values) Get(s string)string{
    if vs:=map[s];len(vs)>0{
        return vs[0]
    }
    return ""
}

func (v Values) Add(key, str string){
    v[key] = append(v[key],str)
}


m:=Values{"lang":{"en"}}
m.Add("item1","1")
m.Add("item2","2")

fmt.Println(m.Get("item")) // "1" (first value)
fmt.Println(m["item"]) // "[1 2]" (direct map access)
m = nil
fmt.Println(m.Get("item")) // ""
m.Add("item", "3") // panic: assignment to entry in nil map

```

#### 函数值
```
type A struct{
    x int
}
//第一种
func (a A) test(){
    fmt.Println(a.x)
}

var f func(a A)
var a A
a.x = 2
f(a)

//第二种
func (a *A) test(){
    fmt.Println(a.x)
}

var f func(a *A)
var a A
a.x = 2
f(&a)


```


#### 指针对象
```
type IntList struct{
    value int
    next *IntList
}
```

#### OOP
```
type Rectangle struct {
    width, height float64
}
type Circle struct {
    radius float64
}
func (r Rectangle) area() float64 {
    return r.width*r.height
}
func (c Circle) area() float64 {
    return c.radius * c.radius * math.Pi
}

func (c *Circle) setWidth(a float64){
    c.width = a
}

```
### 继承
```
type Human struct {
    name string
    age int
    phone string
}
type Student struct {
    Human //匿名字段
    school string
}
type Employee struct {
    Human //匿名字段
    company string
}
//在human上面定义了一个method
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func main() {
    mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
    sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
    mark.SayHi()
    sam.SayHi()
}
```

#### 私有变量
大小写来实现(大写开头的为共有，小写开头的为私有)

#### interface
interface{}类型的变量也被当作go语言中的void\*指针来使用interface{}类型的变量也被当作go语言中的void\*指针来使用
```
//interface == void*
type Human interface{
    setName(name string)
    setAge(age int)
    getName()string
    getAge()int
}

func NewHuman(){

}

```

#### interface{} 类型转换
```
type __s struct{
    x int
}

element := make([]interface{},4)
a := __s{2}
element[0] = &a   // interface{} 赋值为 * __s
element[1] = a    // interface{} 赋值为 __s(复制)
b,ok := element[0].(*__s)
if ok{
    fmt.Println((*b).x)
}
c,ok := element[1].(__s)
if ok{
    fmt.Println(c.x)
}

```
#### 泛型
```
func main() {
    list := make(List, 3)
    list[0] = 1 //an int
    list[1] = "Hello" //a string
    list[2] = Person{"Dennis", 70}
    for index, element := range list{
        switch value := element.(type) {
            case int:
                fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
            case string:
                fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
            case Person:
                fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
            default:
                fmt.Println("list[%d] is of a different type", index)
        }
    }
}
```

#### 函数类型转换
```
type Handler interface{
    ServeHTTP(res, *req)
}

type HandlerFunc func(res, *req)

func (h HandlerFunc) ServeHTTP(res, *req){
    h(res, *req)
}

func Handle(pattern string, handler Handler){
    ....
}

func Myfunc(res, *req){
    ...
}

func main(){
    Handle("/dsad/dsada",HandlerFunc(Myfunc))


}

//实例

type i_bin interface{
	funcx(int)
}

type funcxx func(int)

func (f funcxx) funcx(a int){
	f(a)
}

func ff(i i_bin, x int){
	i.funcx(x)
}

func myf(a int){
	fmt.Println(">>> ",a+1)
}

func myf1(a int){
	fmt.Println(">>> ",a+2)
}

func main(){
	mymap := make(map[string]i_bin)
	mymap["a"] = funcxx(myf)
	mymap["b"] = funcxx(myf1)

	mymap["a"].funcx(1)
	mymap["b"].funcx(1)
}

```



#### 反射
```
1、从relfect.Value中获取接口interface的信息
（1）已知原有类型【进行“强制转换”】
var a int32 = 2
pointer := reflect.ValueOf(&a)
value := reflect.ValueOf(a)
convertPointer := pointer.Interface().(*int32)
convertValue := value.Interface().(int32)


（2）未知原有类型【遍历探测其Filed】
type __s struct{
	Name string
	Age int
	Sex bool
}

func TestRef(in interface{}){

	Type := reflect.TypeOf(in)
	Value := reflect.ValueOf(in)

	for i:=0;i<Type.NumField();i++{
		fieldType := Type.Field(i)
		fieldValue := Value.Field(i).Interface()
		fmt.Printf("%s %v = %v\n",fieldType.Name,fieldType.Type,fieldValue)
	}

	for i:=0;i<Type.NumMethod();i++{
		m := Type.Method(i)
		fmt.Printf("%s: %v\n",m.Name,m.Type)
	}

}

2、通过reflect.Value设置实际变量的值

3、实例（结构体的修改）
type Animal struct{
    Name    string
    Age     int
}

a := Animal{"aaa",14}

x := reflect.ValueOf(&a)
if x.Kind()!=reflect.Ptr{
    wrong
}

elem := x.Elem()
if elem.Kind()!=reflect.Struct{
    wrong
}
for i:=0;i<elem.NumField();i++{
    Name := elem.Type().Field(i).Name 

}







```


#### goroutine

#### channels

#### channels buffer

#### range close

#### select

#### time.After

#### strconv
```
strconv.Atoi //string to int
```

#### regexp
```

```


#### io
```
```

#### container
```
"container/list"   //链表
"container/heap"  //堆
"container/ring"  //环
```

#### Tag
```
import (
    "fmt"
    "reflect"
)

type My struct{
    Name string `json:"Name"`
}

a:=My{Name:"malx"}
b:=reflect.TypeOf(a)
b.Field(0).Type
b.Field(0).Name
b.Field(0).Tag
b.Field(0).Tag.Get("json")

a := MyStruct{Name:"malx",Age:29}
aT:= reflect.TypeOf(a)
aV:=reflect.ValueOf(a)
for i:=0;i<aT.NumField();i++{
	c:=aT.Field(i)
	d:=aV.Field(i)
	fmt.Println(c,d)
}
```


#### 文件上传
```
func UpLoadFile(w http.ResponseWriter, r *http.Request){
    r.ParseMultipartForm(32 << 20);
    files := r.MultipartForm.File['filename']
    myfile := files[0]

    f,err := myfile.Open();
    if err{
        wrong
    }
    defer f.Close()
    cur, err := os.Create("./upload/" + header.Filename)
    defer cur.Close()
    io.Copy(cur, f)

}
```

#### unsafe.Pointer
相当于void*, 指针间的转换要先转换成unsafe.Pointer,再转换成其他指针
```
import (
    "unsafe"
)
x := []int{1, 2, 3, 4}
y := unsafe.Pointer(&x)
z := *(*[]int)(y)
z[0] = 111
```

#### reflect类型判断
```
x := []map[string]int{}
fmt.Println(reflect.TypeOf(x).Kind())
a := reflect.ValueOf(&x)
b := reflect.TypeOf(&x)
fmt.Println(a.Type().Kind())
fmt.Println(b.Kind())

slice
ptr
ptr

```

#### reflect指针
```
x := []int{1, 2, 3, 4}
a := reflect.ValueOf(&x)
p := *(*[]int)(unsafe.Pointer(a.Elem().Addr().Pointer()))
p[0] = 111
fmt.Println(x)

```

#### io
(1)read
```
import(
    "bufio"
    "os"
    "ioutil"
)

//一次性读取
f,err = os.Open("file")
s := ioutil.ReadAll(f)

//分块读取
f,err = os.Open("file")
buf := make([]byte,10)
rd := bufio.NewReader(f)
n,err := rd.Read(buf)

//按行读取
rd := bufio.NewReaderSize(f,4096)
rd: = bufio.NewReader(f)
line,err := rd.ReadString('\n')
line,err := rd.ReadLine()

```
(2)write
```
//带缓冲区读写
fd,_ := os.OpenFile("bbb.txt")
w := bufio.NewWriterSize(fd,4096) //带缓冲的写
w.WriteString("dadadadad")
w.Write([]byte("dsadadada\n"))
w.flush()

//输出屏幕
w := bufio.NewWriterSize(os.Stdout,111)
w1 := bufio.NewReader(f,111)
w1.WriteTo(w)
```
(3)其他函数
```
rd.Buffered()   //表示已经缓冲的数据的大小
w.Available()   //表示可使用的缓冲区的大小

//输出文件
ioutil.WriteFile("11.gv", []byte(graph.String()), 0666)


//实例
package main

import (
	"bufio"
	"fmt"
	"os"
	"unsafe"

	"github.com/awalterschulze/gographviz"
)

func main() {

	graphAst, _ := gographviz.Parse([]byte(`digraph G{}`))
	graph := gographviz.NewGraph()
	gographviz.Analyse(graphAst, graph)
	graph.AddNode("G", "a", nil)
	graph.AddNode("G", "b", nil)
	graph.AddEdge("a", "b", true, nil)
	fmt.Println(graph.String())

	f, _ := os.Open("sjm.txt")
	f1, _ := os.OpenFile("sjm.txt1", os.O_WRONLY|os.O_CREATE, 0664)
	rd := bufio.NewReader(f)
	rd1 := bufio.NewWriter(f1)
	for line, err := []byte{0}, error(nil); len(line) > 0 && err == nil; {
		line, _, err = rd.ReadLine()
		x := (*string)(unsafe.Pointer(&line))
		if *x == "" {
			continue
		}
		line = append(line, '\n')
		rd1.Write(line)
		//fmt.Printf("%v %v\n", *x, p)
	}

	f.Close()
	f1.Close()
}
```


#### net/http
(1)基础结构
```
//handler重中之重的基础
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

//重要的serveMux路由结构
type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry
    hosts bool 
}

type muxEntry struct {
    explicit bool
    h        Handler
    pattern  string
}

//服务器
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}

type Server struct {
    Addr         string        
    Handler      Handler       
    ReadTimeout  time.Duration 
    WriteTimeout time.Duration 
    TLSConfig    *tls.Config   

    MaxHeaderBytes int

    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
    disableKeepAlives int32     nextProtoOnce     sync.Once 
    nextProtoErr      error     
}


```


(1)函数作为处理器
```
func helloHandler(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

http.HandlerFunc()

```
(2)自定义处理器





#### defer panic recover
```
//Defer function按照后进先出的规则执行。例如下面的代码打印“3210”。
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}

//Defer function在方法的return之后执行，如果defer function修改了return的值，返回的是defer function修改后的值。例如下面的例子返回2而不是1。
func c() (i int) {
    defer func() { i++ }()
    return 1
}

```