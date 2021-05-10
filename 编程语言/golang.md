#### 参数传递
```
函数调用参数均为值传递，不是指针传递或引用传递。经测试引申出来，当参数变量为指针或隐式指针类型，参数传递方式也是传值（指针本身的copy）
```

### 交叉编译
```
window -> linux
SET CGO_ENABLED=0; set GOOS=linux; set GOARCH=amd64; go build main.go

linux -> window
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

### make、new
```
make用于内建类型(map、slice 和channel等)的内存分配
new用于各种类型的内存分配
// 注意
    当用new分配内建类型(map、slice 和channel等)时，仅对分配空间清零，未做初始化

list:= make([]int,5,10) //初始化5个int的数组，可扩展性为10个
len(list)  //5
cap(list)  //10

type A struct{x int}
a:=new(A)
a->x = 1    //错误，不支持箭头操作
a.x = 1     //正确
(*a).x = 1  //正确
a++         //错误，不支持指针算术

// make vs new
make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
new 分配的空间被清零。make 分配空间后，会进行初始化

```

### slice
```
// 结构
    type slice struct {
        array unsafe.Pointer
        len   int
        cap   int
    }
    slice是一种值类型，里面有3个元素
    array是数组指针，它指向底层分配的数组
    len是底层数组的元素个数
    cap是底层数组的容量，超过容量会扩容

// 初始化
    1、make
        a := make([]int32, 0, 5)
    
    2、[]int32{}
        b := []int32{1, 2, 3}

    3、new([]int32) // 不推荐
        c := *new([]int32)
```

### 数组与slice
```
//数组是初始化定长的slice
var a [5] int
a:=[5]int{1,2}
a:=[...]int{2:3,3:4}

func test(a [] int)  // 数组做参数按'值传递'
test(a)

//slice切片又称为变长数组
var a [] int
a:=[]int{1,2,3}
a:=make([]int,4)

func test(a [] int) // slice做参数按'引用传递'
test(a)

//数组与slice转换
var a [6] int
var b [] int = a[:]

//slice、map、interface、channel都是按引用传递
```

### array和slice
```
array在栈上
slice在堆上
```

### map
```
//声明
var my map[string]int;       //此处只是声明，没有分配内存
my := make(map[string]int)
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

### init
```
初始化导入的包（递归导入）
对包块中声明的变量进行计算和分配初始值
执行包中的init函数

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

### import
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

### struct
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

### 复杂结构
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

### 链表
```
type IntList struct{
    value int
    next *IntList
}
```

### OOP
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

### 私有变量
大小写来实现(大写开头的为共有，小写开头的为私有)

### 接口
```
type Human interface{
    setName(name string)
    setAge(age int)
    getName()string
    getAge()int
}

type xiaoming struct {
    Name string,
    Age int,
}

func(x *xiaoming) setName(name string) {
    x.Name = name
}
func(x *xiaoming) setAge(age int) {
    x.Age = age
}
func(x *xiaoming) getName() string{
    return x.Name
}
func(x *xiaoming) getAge() int {
    return x.Age
}

func printHuman(h *Human){
    fmt.Println(h.getName(), h.getAge())
}

a:=&xiaoming{Name: "zzz", Age: 12}
printHuman(a)

```

### 类型转换
```
// interface{}类型转换
var a interface{}
a.(int)

// 其他类型转换
var a string
int(a)
```

### interface{} 类型转换
```
interface{}类型的变量也被当作go语言中的void*指针来使用interface{}类型的变量

其实 interface就是个协议，无论你输入指针，还是非指针，都一样

type Test struct{
    x int
}

element := make([]interface{},4)
a := Test{2}
element[0] = &a   // interface{} 赋值为 *Test
element[1] = a    // interface{} 赋值为 Test(复制)
b,ok := element[0].(*Test)
c,ok := element[1].(Test)
```

### 泛型
```
func main() {
    list := make([]interface{}, 3)
    list[0] = 1 //an int
    list[1] = "Hello" //a string
    list[2] = Person{"Dennis", 70}
    for index, element := range list{
        switch element.(type) {
            case int:
                fmt.Printf("list[%d] is an int and its value is %d\n", index, value.(int))
            case string:
                fmt.Printf("list[%d] is a string and its value is %s\n", index, value.(string))
            case Person:
                fmt.Printf("list[%d] is a Person and its value is %s\n", index, value.(Person))
            default:
                fmt.Println("list[%d] is of a different type", index)
        }
    }
}
```

### reflect
```
// reflect
> Kind
    type Kind Uint
    func (k Kind) String() string

> StructField
    type StructField struct {
        Name    string      // 字段的名字
        PkgPath string
        Type      Type      // 字段的类型
        Tag       StructTag // 字段的标签
        ...
    }

> StructTag
    type StructTag string

> Type
    func TypeOf(i interface{}) Type
    type Type interface {
        Kind() Kind                                     //
        Name() string                                   //
        String() string                                 //
        Len() int
        Elem() Type
        Key() Type
        NumField() int                                  //
        Field(i int) StructField                        //
        FieldByIndex(index []int) StructField
        FieldByName(name string) (StructField, bool)
        Method(int) Method                              //
        MethodByName(string) (Method, bool)
    }

> Value
    func ValueOf(i interface{}) Value

    func (v Value) Kind() Kind
    func (v Value) Type() Type
    func (v Value) Elem() Value

var myValue interface{}

v:= reflect.ValueOf(myValue)
if v.Kind() == reflect.Ptr {
    v = v.Elem()
}
t := v.Type()


```

### value pointer reflect之间的转换
```
// 指针转换要先转换成unsafe.Pointer，类似void*
// *type1 -> Unsafe.Pointer -> *type2
var a int = 1
b := unsafe.Pointer(&a)
fmt.Println(*(*int)(b))

// reflect中的'指针'与'value'
var a int = 1
b := reflect.ValueOf(&a)        // 指针类型的reflect
if b.Kind() == reflect.Ptr      // 判断是否为指针
    c := b.Elem()               // 如果是指针需转换成值类型
fmt.Println(b.Pointer())        // 可以直接得到pointer    
fmt.Println(c.Addr().Pointer()) // 需要转换成指针类型的reflect,才能得到pointer

// reflect中的'value'与'type'
rValue := reflect.ValueOf(a)
rType  := rValue.Type()
rValue.Interface().(int)        // 必须先转换成interface{}，才能转换成int

// 遍历struct成员变量
for i := 0; i < Type.NumField(); i++{
    fieldType := rType.Field(i)
    fieldValue := rValue.Field(i)
    fieldType.Name                  // 变量名
    fieldType.Type                  // 变量类型
    fieldType.Type.String()         // 变量类型
    fieldValue.Interface().(MyType) // 变量
}

```

### reflect类型判断
```
x := []map[string]int{}
fmt.Println(reflect.TypeOf(x).Kind())   // slice
a := reflect.ValueOf(&x)
b := reflect.TypeOf(&x)
fmt.Println(a.Type().Kind())            // ptr
fmt.Println(b.Kind())                   // ptr

```

### reflect指针
```
x := []int{1, 2, 3, 4}
a := reflect.ValueOf(&x)
p := *(*[]int)(unsafe.Pointer(a.Elem().Addr().Pointer()))
p[0] = 111
fmt.Println(x)

```

#### 反射
```
1、从relfect.Value中获取接口interface的信息
    (1) 已知原有类型进行"强制转换"
        var a int32 = 2
        pointer := reflect.ValueOf(&a)
        value := reflect.ValueOf(a)
        convertPointer := pointer.Interface().(*int32)
        convertValue := value.Interface().(int32)


    (2) 未知原有类型遍历探测其Field
        type Test struct{
            Name string             // 成员变量
            Age int
            Sex bool
        }
        
        func(t Test) MyFunc(){}     // 成员函数

        TestRef(Test{Name: "aa", Age: 12, Sex: true})

        func TestRef(in interface{}){
            Type := reflect.TypeOf(in)
            Value := reflect.ValueOf(in)

            // 成员变量
            for i := 0; i < Type.NumField(); i++{
                fieldType := Type.Field(i)
                fieldValue := Value.Field(i).Interface()  //只有转换成interface才能转换别的类型
                fmt.Printf("%s %v = %v\n",fieldType.Name,fieldType.Type,fieldValue)
            }

            // 成员函数
            for i := 0;i<Type.NumMethod();i++{
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

### 文件上传
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

### unsafe.Pointer
```
指针间的转换要先转换成unsafe.Pointer,再转换成其他指针
*type1 -> unsafe.Pointer -> *type2

import (
    "unsafe"
)
x := []int{1, 2, 3, 4}
y := unsafe.Pointer(&x)
z := *(*[]int)(y)
z[0] = 111
```

### 查看GC
```
GODEBUG=gctrace=1 go run test.go
```

### pprof
```
go的pprof工具可以用来监测进程的运行数据，用于监控程序的性能，对内存使用和CPU使用的情况统信息进行分析

官方提供了两个包：runtime/pprof和net/http/pprof，前者用于普通代码的性能分析，后者用于web服务器的性能分析

1、net/http/pprof
    (1) 作用
        cpu(CPU Profiling)
            $HOST/debug/pprof/profile，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
        block(Block Profiling)
            $HOST/debug/pprof/block，查看导致阻塞同步的堆栈跟踪
        goroutine
            $HOST/debug/pprof/goroutine，查看当前所有运行的 goroutines 堆栈跟踪
        heap(Memory Profiling)
            $HOST/debug/pprof/heap，查看活动对象的内存分配情况
        mutex(Mutex Profiling)
            $HOST/debug/pprof/mutex，查看导致互斥锁的竞争持有者的堆栈跟踪
        threadcreate
            $HOST/debug/pprof/threadcreate，查看创建新OS线程的堆栈跟踪

    (2) 使用方式
        // 在调用程序处引入包
            import _ "net/http/pprof"

        // 然后调用上述不同profiles查看

        1) 通过web页面查看服务状态
            http://localhost:port/debug/pprof/<profiles>

        2) 交互式终端使用
            go tool pprof http://localhost:8080/debug/pprof/<profile>
            // 每60s刷新
            go tool pprof http://localhost:8080/debug/pprof/<profile>?seconds=60

2、runtime/pprof
    (1) 生成xxx.prof文件
        f, _ := os.Create("./cpu.prof")
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()

    (2) 查看prof界面
        1) 生成Prof的可视化界面
            go tool pprof -http=:8080 ./cpu.prof
            // -http    输出port

        2) 生成火山图
            go get -u github.com/google/pprof
            pprof -http=:8080 ./cpu.prof

```

### GC调优
```
> 减少对象分配
> 函数尽量不要返回map、slice对象，这种频繁调用的函数会给gc带来压力
> 小对象要合并
> 函数频繁创建的简单的对象，直接返回对象，效果比返回指针效果要好
> 可以用对象池sync.Pool(虽然有人说1.10 后不推荐使用sync.Pool，但是压测来看，确实还是用效果，堆累计分配大小能减少一半以上)
> 类型转换要注意，官方用法消耗特别大，推荐使用雨痕的方式
6、避免反复创建slice。
7、建议多用unsafe.pointer
8、对于一些短小的对象，复制成本远小于在堆上分配和回收操作。
9、预定义make大小 make(map[string]string,10000)


//小对象的操作
避免使用make及指针的操作，可以让他在栈中生成，减少堆内存的分配
对于小对象，直接将数据交由 map 保存，远比用指针高效。这不但减少了堆内存分配，关键还在于垃圾回收器不会扫描非指针类型 key/value 对象。
函数返回对象指针时，必然在堆上分配。

//mutex和chan
channel 算是一种很 “重” 的实现。在小粒度层面，其性能真算不得好，不如用mutex ???

//接口的调用
接口调用和普通调用存在很大差别
首先，相比静态绑定，动态绑定性能要差很多；其次，运行期需额外开销，比如接口会复制对象，哪怕仅是个指针，也会在堆上增加一个需 GC 处理的目标。

//sync.Pool
对于可能反复创建的变量可以用缓存池
for i:=0;i<10000;i++{
    a = make([]int,1000)
}

//map诟病
1、slice的gc时间是远远快于map的，而map存储指针则是最慢的，不推荐。
2、避免使用大map
即使映射键和值类型不包含指针，并且映射在GC运行之间没有更改，也会发生这种情况
3、map 不会收缩 “不再使用” 的空间。就算把所有键值删除，它依然保留内存空间以待后用。map=nil

```

### cgo
```
cgo针对该场景定义了专门的规则：在CGO调用的C语言函数返回前，cgo保证传入的Go语言内存在此期间不会发生移动，C语言函数可以大胆地使用Go语言的内存！

Go调用C Code时，Go传递给C Code的Go指针所指的Go Memory中不能包含任何指向Go Memory的Pointer。
不能嵌套其他指向Go Memory的指针

func C.CString(string) *C.char
func C.CBytes([]byte) unsafe.Pointer
func C.GoString(*C.char) string
func C.GoStringN(*C.char, C.int) string
func C.GoBytes(unsafe.Pointer, C.int) []byte

dir := make([]*C.char,2)
for i,_ := range dir{
    dir[i] = (*C.char)(C.malloc(10))
    defer C.free(unsafe.Pointer(dir[i]))
}
C.fill_array1((**C.char)(unsafe.Pointer(&dir[0])))
clen := C.strlen(dir[0])

x:=C.GoBytes(unsafe.Pointer(dir[0]),10)
fmt.Println(x)

dir := make([]*C.char,2)
C.fill_array1((**C.char)(unsafe.Pointer(&dir[0])))
// void fill_array1(char* array[2]){
//      array[0] = (char*)malloc(10);
//}
defer C.free(unsafe.Pointer(dir[0]))
x:=C.GoBytes(unsafe.Pointer(dir[0]),10)
fmt.Println(x)

a := "dadad"
b := []byte(a)
C.fill_array((*C.char)(unsafe.Pointer(&b[0])))


//注意要活学活用
func C.GoString(*C.char) string
func C.GoStringN(*C.char, C.int) string
func C.GoBytes(unsafe.Pointer, C.int) []byte

```

### go build -a
```
强行对所有涉及到的代码包（包含标准库中的代码包）进行重新构建，即使它们已经是最新的了。
cgo中非常重要
```

### 去除struct中的点
用到了reflect
```
func delStructPoint(i interface{}) error {
	v := reflect.ValueOf(i)
	fmt.Println(v.Kind())
	if v.Kind() != reflect.Ptr {
		fmt.Println("must ptr")
		return errors.New("must ptr")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		fmt.Println("must struct ptr")
		return errors.New("must struct ptr")
	}
	for i := 0; i < v.NumField(); i++ {
		curV := v.Field(i)
		fmt.Println(curV)
		fmt.Println(curV.Type())
		if curV.Kind() == reflect.String && curV.String() == "." {
			curV.SetString("")
		}
		if curV.Kind() == reflect.Struct {
			delStructPoint(curV.Addr().Interface())  //重点
		}
	}
	return nil
}
```

### utf8
```
golang的string使用utf8存储的，中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。

unicode的中英文都是两个字节

string <--> utf-8 <--> []byte

所以string先要转换成unicode，再计算长度
a := "你好123"
fmt.Println(len(a))  //9 ×
b := []rune(a)  //转换unicode
fmt.Println(len(b))  //5 ✔

```

### go-xorm
```
https://github.com/xormplus/xorm

//xorm支持将一个struct映射为数据库中对应的一张表。

(1) 映射
    //struct
    type PfSession struct {
        SessionId  string    `xorm:"pk not null  VARCHAR(64)"`
        ValueId    string    `xorm:"not null TEXT"`
        Value      string    `xorm:"not null TEXT"`
        Createtime time.Time `xorm:"default '0000-00-00 00:00:00' DATETIME"`
        Motifytime time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
        Enable     string    `xorm:"not null  default 'T' VARCHAR(1)"`
    }
    //mysql
    CREATE TABLE `pf_session` (
    `session_id` varchar(64) NOT NULL,
    `value_id` text NOT NULL,
    `value` text NOT NULL,
    `createtime` datetime DEFAULT '0000-00-00 00:00:00',
    `motifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `enable` varchar(1) NOT NULL DEFAULT 'T',
    PRIMARY KEY (`session_id`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8
    //三种方式
    > engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "pf_"))
        type session struct
    > engine.SetTableMapper(core.SameMapper{})
        type session struct
    > engine.SetTableMapper(core.SnakeMapper{})
        type PfSession struct

(2) ORM方式操作数据库
    

(3) 事务模型

```

### sizeof
```
//切片的大小
arr := [...]int{1,2,3,4,5}
fmt.Println(unsafe.Sizeof(arr)) //40

//string 大小
typedef struct{
    char* buffer;
    size_tlen;
} string;
unsafe.Sizeof("dasdas")  //16

//

```

### 模块管理
```
version > go1.11

1、go env
    GO111MODULE="on"
    GOPATH="/home/malx/project/go"
    GOPROXY="https://goproxy.io,direct"
    GOMOD="/home/malx/test/test/go.mod"

2、go get
    // 更新当前模块
    go get -u ./...
3、go mod
    (1) 设置环境变量 GO111MODULE
        go env -w GO111MODULE=on
    (2) 参数
        download    下载依赖的module
        edit        编辑go.mod文件
        graph       打印模块依赖图
        init        在当前目录初始化mod
        tidy        拉取缺少的模块，移除不用的模块
        vendor      将依赖复制到vendor下
        verify      校验依赖
        why         解释为什么需要依赖
    (3) 用法
        1) 初始化一个项目helloworld
            // 随便找个文件夹
            go mod init helloworld
            // 写代码，中包含模块
            // 运行项目
            go build
            // 或者先安装依赖
            go mod tity
        2) 老项目
            // 进入项目目录
            cd src/github.com/xxx/xxx
            go mod init helloworld
            // 安装依赖
            go get ./...
            // 或
            go mod tidy
            // 更新旧的package import 方式
            // 例如之前 import "api"要改为以下方式
            import "helloworld/api"

4、GOPROXY
    // version > 1.13
    go env -w GOPROXY="https://goproxy.io,direct"
    // 部署自己的库
    // 安装
    git clone https://github.com/goproxyio/goproxy.git && cd goproxy && make 
    // 参数
    -cacheDir   指定 Go 模块的缓存目录
    -exclude    proxy 模式下指定哪些 path 不经过上游服务器
    -listen     服务监听端口，默认 8081
    -proxy      指定上游 proxy server，推荐 goproxy.io
    // 运行
    ./bin/goproxy -listen=0.0.0.0:80 -cacheDir=/tmp/test -proxy https://goproxy.io -exclude "git.corp.example.com,rsc.io/private"

``` 

### go mod import
```
1、import当前目录的内容
    (1) 目录结构
        # tree
        .
        ├── go.mod
        ├── main.go
        └── test
            └── test.go

    (2) go.mod
        # cat go.mod
        module blog

        go 1.13

    (3) main.go
        # cat main.go
        package main

        import (
            "blog/test"             //注意
        )

        func main() {
            test.Print()
        }

    (4) test.go
        # cat test.go
        package test
        import "fmt"
        func Print() {              // 注意大写
            fmt.Println("aaa")
        }

2、import其他目录的内容
    (1) 目录结构
        # tree
        .
        ├── go.mod
        └── main.go

    (2) go.mod
        # cat go.mod
        module test1
        replace blog => /home/malx/test/go/     //引用上面那个模块

    (3) main.go
        # cat main.go
        package main
        import (
            test "blog/test"
        )

        func main () {
            test.Printa()
        }


```

### go install vs go build
```
go install命令只比 go build命令多做了一件事，即：安装编译后的结果文件到指定目录

go install 编译目标文件的同时，也将编译产生的静态库文件保存在工作区的pkg目录下
go build 只会编译目标文件产生最终结果
```

### net
```

```

### net/rpc
```

```