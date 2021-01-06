### context.WithCancel
```
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
```
(1) *T 调用 T()
	*T类型的值可以调用为类型T声明的方法，这是因为解引用指针总是合法的
	func(t T) f(){}
	var a *T = new(T)
	a.f()	// 合法

(2) T 调用 *T()
	T类型的值可以调用为类型*T声明的方法，但是仅限于T的值可寻址
	编译器在调用指针属主方法前，会自动取此T值的地址。因为不是任何T值都是可寻址的，所以并非任何T值都能够调用为类型*T声明的方法
	func(t *T) f() {}
	T a
	a.f()	// 合法，前提是a可寻址

(3) 不可寻址的值
	字符串中的字节
	map 对象中的元素(slice 对象中的元素是可寻址的，slice的底层是数组)
	常量
	包级别的函数等

	type T string
	func(t *T) f(){}

	var a T = "dasdsa"
	a.f()	// 正确
	const a1 T = "dadas"
	a1.f()	// 错误，常量不可寻址
	
```

### const group 自动补全
```
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
```
const N = 100
var x int = N	// 正确

const M int32 = 100
var y int = M	// 错误

无类型常量，赋值给其他变量时，如果字面量能够转换为对应类型的变量，则赋值成功
有类型的常量，赋值给其他变量时，需要类型匹配才能成功，所以显示地类型转换

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
```
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
```
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
```
每个 Goroutine 在执行之前，都要先知道程序当前的执行状态，通常将这些执行状态封装在一个 Context 变量中，传递给要执行的 Goroutine 中

(1) 接口
	type Context interface {
		// Deadline 方法需要返回当前 Context 被取消的时间，也就是完成工作的截止时间
		Deadline() (deadline time.Time, ok bool)

		// 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭
		Done() <-chan struct{}
		
		// 方法会返回当前 Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值
		Err() error

		// 从 Context 中返回键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法仅用于传递跨 API 和进程间跟请求域的数据
		Value(key interface{}) interface{}
	}

(2) Background和TODO
	这两个函数分别返回一个实现了 Context 接口的 background 和 todo

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

// 注意事项
不要把 Context 放在结构体中，要以参数的方式显示传递
以 Context 作为参数的函数方法，应该把 Context 作为第一个参数
给一个函数方法传递 Context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO
Context 的 Value 相关方法应该传递请求域的必要数据，不应该用于传递可选参数
Context 是线程安全的，可以放心的在多个 Goroutine 中传递
```

### grpc 相关资料
```
[Go-gRPC 实践指南] https://www.bookstack.cn/read/go-grpc/summary.md
[go-grpc 英文文档] https://pkg.go.dev/google.golang.org/grpc
[烟花易冷人憔悴 vlog] https://www.cnblogs.com/FireworksEasyCool/
[go-grpc 应用] https://eddycjy.com/go-categories/ ！！！
```

### grpc 超时
```
// client
// 建立连接时超时
	clientDeadline := time.Now().Add(time.Duration(3 * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	// 或
	ctx, cancel := context.Timeout(context.Bakcground(), time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithBlock(), grpc.WithInsecure()) 
	// grpc.WithBlock()	这个参数会阻塞等待握手成功，直到超时。如果没有设置这个参数，那么context超时控制将会失效
// 调用时超时
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5 * time.Second)))
	defer cancel()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(ctx, &pb.SearchRequest{})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("client.Search err: deadline")
			}
		}

		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.Response())

// server
// 服务端需要判断和处理超时
func (s *SearchService) Search(ctx context.Context, req *pb.SearchRequest)(*pb.Response, error) {
	// 先判断客户端是否已超时
	if ctx.Err() == context.Canceled {
		return status.New(codes.Canceled, "Client cancelled, abandoning.")
	}
	return Search(ctx, req)
}


// 超时传递
> 客户端客户端发起 RPC 调用时传入了带 timeout 的 ctx
> gRPC 框架底层通过 HTTP2 协议发送 RPC 请求时，将 timeout 值写入到 grpc-timeout HEADERS Frame 中
> 服务端接收 RPC 请求时，gRPC 框架底层解析 HTTP2 HEADERS 帧，读取 grpc-timeout 值，并覆盖透传到实际处理 RPC 请求的业务 gPRC Handle 中
> 如果此时服务端又发起对其他 gRPC 服务的调用，且使用的是透传的 ctx，这个 timeout 会减去在本进程中耗时，从而导致这个 timeout 传递到下一个 gRPC 服务端时变短，这样即实现了所谓的 超时传递 

```

### grpc keepalive
```
// 服务端
var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))


// 客户端
var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))

```

### grpc retry
```

```

### grpc TLS + 自定义认证
```
// 密钥
私钥：server.key	公钥：server.pem

// 不带密钥
// server
	s := grpc.NewServer()
// client
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())

// 带密钥
// server
	c, err := credentials.NewServerTLSFromFile("server.pem", "server.key")
	server := grpc.NewServer(grpc.Creds(c))
// client
	c, err := credentials.NewClientTLSFromFile("server.pem", "go-grpc-example")
	conn, err := grpc.Dial(Addr, grpc.WithTransportCredentials(c))

// PerRPCCredentials
// server
// client
```

### grpc interceptor
```
https://colobu.com/2017/04/17/dive-into-gRPC-interceptor/

普通方法：一元拦截器（grpc.UnaryInterceptor）
流方法：流拦截器（grpc.StreamInterceptor）

// 服务端
type ServerOption
func UnaryInterceptor(i UnaryServerInterceptor) ServerOption
func StreamInterceptor(i StreamServerInterceptor) ServerOption
// 需要实现的 *ServerInterceptor
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error

srv := grpc.NewServer(
    grpc.UnaryInterceptor(UnaryServerInterceptorDemo),
    grpc.StreamInterceptor(StreamServerInterceptor),
)

func UnaryServerInterceptorDemo(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before handling. Info: %+v", info)
	resp, err := handler(ctx, req)
	log.Printf("after handling. resp: %+v", resp)
	return resp, err
}

func StreamServerInterceptorDemo(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("before handling. Info: %+v", info)
	err := handler(srv, ss)
	log.Printf("after handling. err: %v", err)
	return err
}


// 客户端
type DialOption interface {}
func WithUnaryInterceptor(f UnaryClientInterceptor) DialOption
func WithStreamInterceptor(f StreamClientInterceptor) DialOption
// 需要实现的 *ClientInterceptor
type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)

grpc.Dial(Addr, grpc.WithInsecure(),
    grpc.WithUnaryInterceptor(UnaryClientInterceptorDemo),
    grpc.WithStreamInterceptor(StreamClientInterceptorDemo),
    ...
)

func UnaryClientInterceptorDemo(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("before invoker. method: %+v, request:%+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("after invoker. reply: %+v", reply)
	return err
}

func StreamClientInterceptorDemo(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("before invoker. method: %+v, StreamDesc:%+v", method, desc)
	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("before invoker. method: %+v", method)
	return clientStream, err
}

// 链式拦截器
// 方案一	go-grpc-middleware
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,
			LoggingInterceptor,
		),
	}
	server := grpc.NewServer(opts...)

// 方案二	自己实现
	基本原理：利用 handler 传递 handler

	handler(...){
		interceptor(...)
	}

	interceptor(..., handler) {
		handler(...)
	}

	func InterceptChain(intercepts... grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(interface{}, error) {

			chain := func(interceptor grpc.UnaryServerInterceptor, handler grpc.UnaryHandler) grpc.UnaryHandler {

				return func(ctxCur context.Context, reqCur interface{}) (interface{}, error){
					return interceptor(ctxCur, reqCur, info, handler)
				}

			}

			handlerTmp := handler
			for _, intercept := range intercepts {
				handlerTmp = chain(intercept, handlerTmp)
			}

			return handlerTmp(ctx, req)

		}
	}

```

### grpc metadata
```
https://pandaychen.github.io/2020/02/22/GRPC-METADATA-INTRO/

metadata 其实就是一个 map

// 结构
func AppendToOutgoingContext(ctx context.Context, kv ...string) context.Context
func DecodeKeyValue(k, v string) (string, string, error)
func NewIncomingContext(ctx context.Context, md MD) context.Context
func NewOutgoingContext(ctx context.Context, md MD) context.Context

func FromIncomingContext(ctx context.Context) (md MD, ok bool)
func FromOutgoingContext(ctx context.Context) (MD, bool)
func FromOutgoingContextRaw(ctx context.Context) (MD, [][]string, bool)
func Join(mds ...MD) MD
func New(m map[string]string) MD
func Pairs(kv ...string) MD

type MD map[string][]string
	func (md MD) Append(k string, vals ...string)
	func (md MD) Copy() MD
	func (md MD) Get(k string) []string
	func (md MD) Len() int
	func (md MD) Set(k string, vals ...string)


// 发送metadata
	创建ctx
	创建并写入metadata
	将metadata关联ctx
	发送请求(rpc.SomeRpc(ctx, someRequest))

	// 新创建的 Metadata 添加到 context 中
	md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	
	// 可以将 key-value 对添加到已有的 context 中
	// 如果k-v已存在，则覆盖
	ctx := metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")
	ctx := metadata.AppendToOutgoingContext(ctx, "k3", "v4")

// 接受metadata
	// Unary Call
	func (s *server) SomeRPC(ctx context.Context, in *pb.someRequest) (*pb.someResponse, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		// do something with metadata
	}

	// Streaming Call
	func (s *server) SomeStreamingRPC(stream pb.Service_SomeStreamingRPCServer) error {
		md, ok := metadata.FromIncomingContext(stream.Context()) // get context from stream
		// do something with metadata
	}

```

### grpc-gateway
```
通过protobuf的自定义option实现了一个网关，服务端同时开启gRPC和HTTP服务
HTTP服务接收客户端请求后转换为grpc请求数据，获取响应后转为json数据返回给客户



```

### LRU
```
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

// 注意
groupcache 中实现的 LRU Cache 并不是并发安全的，如果用于多个 Go 程并发的场景，需要加锁


除了上述库，还有github.com/hashicorp/golang-lru
```

### 浮点数与16进制转换
```go
// 浮点数 转 16进制
// float -> uint32 -> []byte -> hex 

// 16进制 转 浮点
// hex -> []byte -> uint32 -> float
import (
	"math"
	"encoding/binary"
	"encoding/hex"
)


```