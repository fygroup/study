
### grpc 相关资料
```
[Go-gRPC 实践指南] https://www.bookstack.cn/read/go-grpc/summary.md
[go-grpc 英文文档] https://pkg.go.dev/google.golang.org/grpc
[烟花易冷人憔悴 vlog] https://www.cnblogs.com/FireworksEasyCool/
[go-grpc 应用] https://eddycjy.com/go-categories/ ！！！
[grpc 官方文档 中文] https://doc.oschina.net/grpc
[从实践到原理，带你参透 gRPC] https://segmentfault.com/a/1190000019608421
```

### 编译
```bash
# install
go install \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

# grpc-go
protoc -I . -I $GOPATH/src/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ --go_out ./api/ --go_opt paths=source_relative --go-grpc_out ./api/ --go-grpc_opt paths=source_relative pb/auth-manager-service.proto
# grpc-gateway
protoc -I . -I $GOPATH/src/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ --grpc-gateway_out ./api/ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true pb/auth-manager-service.proto
# grpc-swagger
protoc -I . -I $GOPATH/src/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ --openapiv2_out ./api/ --openapiv2_opt logtostderr=true --openapiv2_opt paths=source_relative pb/auth-manager-service.proto
```

### grpc server
```go
// route-guide-server.proto
message LoginRequest {
    string user_name = 1;
}
message LoginResponse {

}

service AuthManager {
    rpc Login (LoginRequest) returns (LoginResponse) {
    }
    rpc Logout (LogoutRequest) returns (LogoutResponse) {
    }
}

// routeGuideServer.go
type AuthManagerServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
}

func newServer() AuthManagerServer {

}

lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
if err != nil {
  log.Fatalf("failed to listen: %v", err)
}
var opts []grpc.ServerOption
...
grpcServer := grpc.NewServer(opts...)
pb.RegisterAuthManagerServer(grpcServer, newServer())
grpcServer.Serve(lis)
```

### grpc client
```go
var opts []grpc.DialOption
...
conn, err := grpc.Dial(addr, opts...)
// 或
conn, err := grpc.DialContext(context.Background(), addr, opts...)

if err != nil {
  ...
}
defer conn.Close()

client := pb.NewAuthManagerClient(conn)
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
req := &pb.LoginRequest{}
res, err := client.Login(ctx, req)

```

### grpc 超时
```go

// go grpc超时是通过context来控制的

1. client
(1) 建立连接时超时
	clientDeadline := time.Now().Add(time.Duration(3 * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	// 或
	ctx, cancel := context.Timeout(context.Bakcground(), time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithBlock(), grpc.WithInsecure())
	// grpc.WithBlock()	这个参数会阻塞等待握手成功，直到超时。如果没有设置这个参数，那么context超时控制将会失效

(2) 调用时超时
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
> 客户端发起 RPC 调用时传入了带 timeout 的 ctx
> gRPC 框架底层通过 HTTP2 协议发送 RPC 请求时，将 timeout 值写入到 grpc-timeout HEADERS Frame 中
> 服务端接收 RPC 请求时，gRPC 框架底层解析 HTTP2 HEADERS 帧，读取 grpc-timeout 值，并覆盖透传到实际处理 RPC 请求的业务 gPRC Handle 中
> 如果此时服务端又发起对其他 gRPC 服务的调用，且使用的是透传的 ctx，这个 timeout 会减去在本进程中耗时，从而导致这个 timeout 传递到下一个 gRPC 服务端时变短，这样即实现了所谓的 超时传递 

```

### grpc keepalive
```go
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
```go
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
```golang
// https://colobu.com/2017/04/17/dive-into-gRPC-interceptor/

// 普通方法：一元拦截器（grpc.UnaryInterceptor）
// 流方法：流拦截器（grpc.StreamInterceptor）

// 服务端
type ServerOption
func UnaryInterceptor(i UnaryServerInterceptor) ServerOption
func StreamInterceptor(i StreamServerInterceptor) ServerOption
// 需要实现的 *ServerInterceptor
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (interface{},  error)
type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error

srv := grpc.NewServer(
    grpc.UnaryInterceptor(UnaryServerInterceptor(UnaryServerInterceptorDemo)),
    grpc.StreamInterceptor(StreamServerInterceptor(StreamServerInterceptorDemo)),
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
```golang
// https://pandaychen.github.io/2020/02/22/GRPC-METADATA-INTRO/

// metadata 是一个 map

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
// 创建ctx
// 创建并写入metadata
// 将metadata关联ctx
// 发送请求(rpc.SomeRpc(ctx, someRequest))

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
```go
// 通过protobuf的自定义option实现了一个网关，服务端同时开启gRPC和HTTP服务
// HTTP服务接收客户端请求后转换为grpc请求数据，获取响应后转为json数据返回给客户

request -> http.server -> rpc.Dial -> rpc.server

```

### runtime.JSONPb
```go
a := &api.MyStruct{
	Time: timestamppb.New(time.Now()),
}	// MyStruct必须经过proto编译,才可以使用
marshal := &runtime.Marshal{}
buf,err := marshal.Marshal(a)
fmt.Println(string(buf))
```

### timestamp 保留毫秒
```go
t := Time: timestamppb.New(time.Now())
fmt.Println(t.AsTime())
t.Nanos = (x.Nanos / 1e6) * 1e6
fmt.Println(t.AsTime())
```