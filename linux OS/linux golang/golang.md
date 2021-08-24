### io
```go
1、read
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
    rd := bufio.NewReaderSize(f,4096) // 带缓冲的读
    rd: = bufio.NewReader(f)
    line,err := rd.ReadString('\n')
    line,err := rd.ReadLine()

2、write
    //带缓冲区读写
    fd,_ := os.OpenFile("bbb.txt")
    w := bufio.NewWriterSize(fd,4096) // 带缓冲的写
    w.WriteString("dadadadad")
    w.Write([]byte("dsadadada\n"))
    w.flush()

    //输出屏幕
    w := bufio.NewWriterSize(os.Stdout,111)
    w1 := bufio.NewReader(f,111)
    w1.WriteTo(w)

3、其他函数
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

#### defer panic recover
```go
1、defer
    //Defer function按照后进先出的规则执行。例如下面的代码打印 "3210"
        func b() {
            for i := 0; i < 4; i++ {
                defer fmt.Print(i)
            }
        }

    // Defer function在方法的return之后执行，如果defer function修改了return的值，返回的是defer function修改后的值。例如下面的例子返回2而不是1
        func c() (i int) {
            defer func() { i++ }()
            return 1
        }

2、panic 和 recover
    (1) defer 表达式的函数如果定义在 panic 后面，该函数在 panic 后就无法被执行到
    (2) F中出现panic时，F函数会立刻终止，不会执行F函数内panic后面的内容，但不会立刻return，而是调用F的defer，如果F的defer中有recover捕获，则F在执行完defer后正常返回，调用函数F的函数G继续正常执行
        func G() {
            defer func() {
                fmt.Println("c")
            }()
            F()
            fmt.Println("继续执行")
        }

        func F() {
            defer func() {
                if err := recover(); err != nil {
                    fmt.Println("捕获异常:", err)
                }
                fmt.Println("b")
            }()
            panic("a")
        }
        
        //结果
            捕获异常: a
            b
            继续执行
            c

    (3) 如果F的defer中无recover捕获，则将panic抛到G中，G函数会立刻终止，不会执行G函数内后面的内容，但不会立刻return，而调用G的defer...以此类推
        func G() {
            defer func() {
                if err := recover(); err != nil {
                    fmt.Println("捕获异常:", err)
                }
                fmt.Println("c")
            }()
            F()
            fmt.Println("继续执行")
        }

        func F() {
            defer func() {
                fmt.Println("b")
            }()
            panic("a")
        }
        // 结果
            b
            捕获异常: a
            c

    (4) 如果一直没有recover，抛出的panic到当前goroutine最上层函数时，程序直接异常终止
        func G() {
            defer func() {
                fmt.Println("c")
            }()
            F()
            fmt.Println("继续执行")
        }

        func F() {
            defer func() {
                fmt.Println("b")
            }()
            panic("a")
        }
        //结果
            b
            c
            panic: a

            goroutine 1 [running]:
            main.F()
                /xxxxx/src/xxx.go:61 +0x55
            main.G()
                /xxxxx/src/xxx.go:53 +0x42
            exit status 2

    (5) recover都是在当前的goroutine里进行捕获的，这就是说，对于创建goroutine的外层函数，如果goroutine内部发生panic并且内部没有用recover，外层函数是无法用recover来捕获的，这样会造成程序崩溃
        func G() {
            defer func() {
                //goroutine外进行recover
                if err := recover(); err != nil {
                    fmt.Println("捕获异常:", err)
                }
                fmt.Println("c")
            }()
            //创建goroutine调用F函数
            go F()
            time.Sleep(time.Second)
        }

        func F() {
            defer func() {
                fmt.Println("b")
            }()
            //goroutine内部抛出panic
            panic("a")
        }
        // 结果
            b
            panic: a

            goroutine 5 [running]:
            main.F()
                /xxxxx/src/xxx.go:67 +0x55
            created by main.main
                /xxxxx/src/xxx.go:58 +0x51
            exit status 2

    (6) recover返回的是interface{}类型而不是go中的 error 类型，如果外层函数需要调用err.Error()，会编译错误，也可能会在执行时panic
```

### 函数类型转换
```go
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
    Handle("/dsad/dsada", HandlerFunc(Myfunc))
    a := make(map[string]Handler)
    a["/a/a"] = HandlerFunc(func)
}
```

### 拦截器
```go
type Handler interface{
    ServeHTTP(res, *req)
}

type HandlerFunc func(res, *req)

func (h HandlerFunc) ServeHTTP(res, *req){
    h(res, *req)
}


func X(h Handler) Handler {
    return HandlerFunc(func(res, *req) {
        // do something...
        h.ServeHTTP(req)
    })
}

http.Hander("/xxx", X(X(X()))
```

### context.WithCancel
```go
// 父协程控制子协程退出

fmt.Println("main 函数 开始...")
go func() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("父 协程 开始...")
	go func(ctx context.Context) {
		for {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("子 协程 接受停止信号...")
					return
				default:
					fmt.Println("子 协程 执行中...")
					timer := time.NewTimer(time.Second * 2)
					<-timer.C
				}
			}
		}
	}(ctx)
	time.Sleep(time.Second*5)
	fmt.Println("父 协程 退出...")
}()
time.Sleep(time.Second*10)
fmt.Println("main 函数 退出")
```

### 任意类型 T() 都能够调用 *T 的方法吗？反过来呢？
```go
(1) *T 调用 T()
	// *T类型的值可以调用为类型T声明的方法，这是因为解引用指针总是合法的
	func(t T) f(){}
	var a *T = new(T)
	a.f()	// 合法

(2) T 调用 *T()
	// T类型的值可以调用为类型*T声明的方法，但是仅限于T的值可寻址
	// 编译器在调用指针属主方法前，会自动取此T值的地址。因为不是任何T值都是可寻址的，所以并非任何T值都能够调用为类型*T声明的方法
	func(t *T) f() {}
	T a
	a.f()	// 合法，前提是a可寻址

(3) 不可寻址的值
	// 字符串中的字节
	// map 对象中的元素(slice 对象中的元素是可寻址的，slice的底层是数组)
	// 常量
	// 包级别的函数等

	type T string
	func(t *T) f(){}

	var a T = "dasdsa"
	a.f()	// 正确
	const a1 T = "dadas"
	a1.f()	// 错误，常量不可寻址
	
```

### const group 自动补全
```go
const (
	a, b = "golang", 100
	d, e
	f bool = true
	g
)

// 自动补全
const (
	a, b = "golang", 100
	d, e = "golang", 100
	f bool = true
	g bool = true
)
```

### 无类型常量和有类型常量
```go
const N = 100
var x int = N	// 正确

const M int32 = 100
var y int = M	// 错误

// 无类型常量，赋值给其他变量时，如果字面量能够转换为对应类型的变量，则赋值成功
// 有类型的常量，赋值给其他变量时，需要类型匹配才能成功，所以显示地类型转换

var y int = int(M) // 正确

```

### time
```
a := 3
b := time.Duration(a) * time.Second // 表示3s
b.Minutes()
b.Hours()
```

### 类型转换
```
Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明

type MyInt int
var a int = 1
var  MyInt = MyInt(a)	// 必须显示转换

```

### 常量的定义
```
常量的值必须在编译期间确定

const (
	a int = 1
	b error = error.New("dasdaa")	// 错误
)

```

### 常量转换不允许溢出
```
func main() {
	var a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b)	// -128
}


func main() {
	const a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b)	
}
编译失败：constant 128 overflows int8

```

### defer延迟调用
```
// defer 延迟调用时，需要保存函数指针和参数，因此链式调用的情况下，除了最后一个函数外都会在调用时直接执行
type T struct{}
func(t T) f(n int){
	fmt.Println(n)
	return t
}
func main(){
	var t T
	defer t.f(1).f(2)
	fmt.Println(3)
}
结果：132


// defer 语句执行时，会将需要延迟调用的函数和参数保存起来
func f(i int){
	defer fmt.Println(i)
	i += 100
}
func main(){
	f(1)
}
结果：1

func f(i int){
	defer func(){
		fmt.Println(i)
	}()
	i += 1
}
结果：101

// defer 的作用域是函数，而不是代码块
func main(){
	a := 1
	if a == 1 {
		defer fmt.Println(a)
		a += 100
	}
	fmt.Println(a)
}
结果：101 1
```

### sync.Pool
```go
type Pool struct {
	...
	New func() interface{}
}

// 从 Pool 中获取元素，当Pool中没有元素时，会调用 New 生成元素，新元素不会放入 Pool 中，若 New 未定义，则返回 nil
func (p *Pool) Get() interface{}

//往 Pool 中添加元素 x
func (p *Pool) Put(x interface{})

type A struct{}

var bufPool = sync.Pool{
	New: func()interface{}{
		return new(A)
	}
}

b, _ = bufPool.Get().(*A)

bufPool.Put(b)
```

### 指针
```
go中不能对指针进行自增或自减运算
不能对指针进行下标运算
```

### 接口赋值
```go
type A interface {
	f(int)error
}

type B struct{}
func (b B) f(int)error{
	return nil
}

var x A = B{}
var x A1 = &B{}
var x A2 = new(B)

```

### sync.Map
```
```

### context
```go
// 每个 Goroutine 在执行之前，都要先知道程序当前的执行状态，通常将这些执行状态封装在一个 Context 变量中，传递给要执行的 Goroutine 中

(1) 接口
	type Context interface {
		// Deadline 方法需要返回当前 Context 被取消的时间，也就是完成工作的截止时间
		Deadline() (deadline time.Time, ok bool)

		// 当 context 被取消或者到了 deadline，返回一个被关闭的 channel
        Done() <-chan struct{}
		
		// 方法会返回当前 Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值
		Err() error

		// 从 Context 中返回键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法仅用于传递跨 API 和进程间跟请求域的数据
		Value(key interface{}) interface{}
	}

(2) Background和TODO
	// 这两个函数分别返回一个实现了 Context 接口的 background 和 todo

(3) withCancel
	ctx, cancel := context.WithCancel(context.Background())
	
	// demo
	ctx, cancel := context.WithCancel(context.Background())
	dst := make(chan int)

	go func(ctx context.Context){
		i := 1
		for {
			select {
				case <- ctx.Done():
					return
				case dst <- i:
					i++
			}
		}
	} (ctx)

	for i := range dst {
		fmt.Println(i)
		if i == 5 {
			cancel()
		}
	}

(4) WithDeadline
	d := time.Now().Add(50 * time.Millisecond)
    ctx, cancel := context.WithDeadline(context.Background(), d)
    // 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践
    // 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间
	// withDeadline 很好的解决了grpc调用阻塞时间过长的问题
    defer cancel()
    select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
    }

(5) WithValue
	type favContextKey string // 定义一个key类型
    // f:一个从上下文中根据key取value的函数
    f := func(ctx context.Context, k favContextKey) {
        if v := ctx.Value(k); v != nil {
            fmt.Println("found value:", v)
            return
        }
        fmt.Println("key not found:", k)
    }
    k := favContextKey("language")
    // 创建一个携带key为k，value为"Go"的上下文
    ctx := context.WithValue(context.Background(), k, "Go")
    f(ctx, k)
    f(ctx, favContextKey("color"))
}

> 注意事项
// 不要把 Context 放在结构体中，要以参数的方式显示传递
// 以 Context 作为参数的函数方法，应该把 Context 作为第一个参数
// 给一个函数方法传递 Context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO
// Context 的 Value 相关方法应该传递请求域的必要数据，不应该用于传递可选参数
// Context 是线程安全的，可以放心的在多个 Goroutine 中传递
```

### LRU
```go
// 缓存淘汰算法

github.com/golang/groupcache

//创建一个 LRU Cache
func New(maxEntries int) *Cache
 
//向 Cache 中插入一个 KV
func (c *Cache) Add(key Key, value interface{})

//从 Cache 中获取一个 key 对应的 value
func (c *Cache) Get(key Key) (value interface{}, ok bool)

//从 Cache 中删除一个 key
func (c *Cache) Remove(key Key)

//从 Cache 中删除最久未被访问的数据
func (c *Cache) RemoveOldest()

//获取 Cache 中当前的元素个数
func (c *Cache) Len()

//清空 Cache
func (c *Cache) Clear()

> 注意
// groupcache 中实现的 LRU Cache 并不是并发安全的，如果用于多个 Go 程并发的场景，需要加锁


// 除了上述库，还有github.com/hashicorp/golang-lru
```

### 浮点数与16进制转换
```go
// 浮点数 转 16进制
// float -> uint32 -> []byte -> hex (-> string)

// 16进制 转 浮点
// (string ->) hex -> []byte -> uint32 -> float
import (
	"math"
	"encoding/binary"
	"encoding/hex"
)

var a float32
var b uint32 = math.Float32bits(a)
bytes := make([]byte, 4)
binary.LittleEndian.PutUint32(bytes, b)
result := hex.EncodeToString(bytes)

bytes1, _ := hex.DecodeString(result)
v := binary.LittleEndian.Uint32(bytes1)
a1 := math.Float32frombits(v)



```

### go mod 私有项目
```
// 场景
import "gitlab.sz.sensetime.com/SenseStardust/stardust"

(1) 更改git http -> ssh (~/.gitconfig) (version < 1.13)
    git config --global url."git@gitlab.sz.sensetime.com:".insteadOf "https://gitlab.sz.sensetime.com/"
    gitlab网上添加公钥

(2) 私有下载地址
    go env -w GOPRIVATE="gitlab.sz.sensetime.com"

(3) go get 下载私有项目(注意分支)
    go get gitlab.sz.sensetime.com/SenseStardust/stardust@master

(4) go mod tidy

// 注意
(1) go get 需要输入密码
    将本机的公钥添加到gitlab上面

(2) gitlab.sz.sensetime.com/viper/charts/kafka.git  四级目录
    go get gitlab.sz.sensetime.com/viper/charts/kafka@master    // 错误
    go get gitlab.sz.sensetime.com/viper/charts/kafka.git       // 正确

(3) 使用https
    更改 git ssh -> http rm ~/.gitconfig
    cat ~/.netrc
    machine gitlab.sz.sensetime.com login malixiang password mlx.04211317

```

### []string []interface{}
```golang

// []string 不能转换成 []interface{}
func f(v []interface{}) {}
a := []string{"a", "b"}
f(a)    // 错误

// slice interface{} 转换
func f(v []interface{}) {}
a := []string{"a", "b"}
b := make([]interface{}, 2)
for i,v := range a {
    b[i] = v
}
f(b)    // 正确
```

### 遍历interface{}
```go
func TypeJudge(i interface{}) {
	ref := reflect.ValueOf(i)
	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}
	switch ref.Type().Kind() {
	case reflect.Array, reflect.Slice:
		fmt.Println("reflect.Array")
		for i := 0; i < ref.Len(); i++ {
			TypeJudge(ref.Index(i).Interface())
		}
	case reflect.Map:
		fmt.Println("reflect.Map")
		iter := ref.MapRange()
		for iter.Next() {
			k := iter.Key().Interface().(string)
			fmt.Printf("key %v \n", k)
			TypeJudge(iter.Value().Interface())
		}
	case reflect.Struct:
		refType := ref.Type()
		fmt.Println("struct ", refType.Name())
		_, ok := refType.MethodByName("String")
		if ok {
			fmt.Println("value ", ref.Interface())
			return
		}
		for i := 0; i < ref.NumField(); i++ {
			name := refType.Field(i).Name
			if unicode.IsUpper([]rune(name)[0]) {
				fmt.Printf("name: %v \n", name)
				TypeJudge(ref.Field(i).Interface())
			}
		}
	default:
		fmt.Printf("value: %v \n", i)
	}
}
```

### rand
```go
import "math/rand"
myrand := rand.New(rand.NewSource(time.Now().Unix()))
myrand.Intn(100)
```