### JSW
```go
// session加密，将其放到客户端。本地只保留密钥，验证只需要找到对应的密钥，再解密session获得用户信息

import (
    "github.com/dgrijalva/jwt-go"
)

func CreateToken() string{
    token := jwt.New(jwt.SigningMethodHS256)
    claims := jwt.MapClaims{
        id : "id",
        name: "name",
        iat: "签发时间"
        //iss: 签发者
        //sub: 面向的用户
        //aud: 接收方
        //exp: 过期时间
        //nbf: 生效时间
        //jti: 唯一身份标识
    }
    token.Claims = claims
    tokenString, _ := token.SignedString([]byte(secret))
    return tokenString
}

func CheckToken(tokenString string, secret string) bool{
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error){
        return []byte(secret),nil
    }) 

    if err != nil || !token.Valid {
        return false
    }

    claims, err := token.Claims.(jwt.MapClaims)
    //claims['id'].(int)
    //cliams['name'].(string)
    return true

}
```

### net/http
```go
// 重要的接口与结构
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

// serveMux路由
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

// 默认路由，当外部没有指定路由时，系统用此路由
var DefaultServeMux = &defaultServeMux
var defaultServeMux ServeMux

// 监听服务
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}


// 两个注册函数，用的都是默认路由
// 1
func Handle(pattern string, handler Handler) {
  	DefaultServeMux.Handle(pattern, handler)
}
// 2
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}

```

### http拦截器
```golang

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (h HandlerFunc) ServerHTTP(w ResponseWriter, r *Request) {
	h(w,r)
}


type HttpMiddleware []http.Handler

func (it Interceptor) ServeHTTP(w ResponseWriter, r *Request) {

}

func Hello(w,r)

mux.Handle("/aaa", Interceptor(Hello))


func Interceptor(f func(ResponseWriter, *Request)) {
	return func(w ResponseWriter, r *Request){
		// do something...
		f(w,r)
	}
}


```


### serveFile
```go
type Dir string

func (d Dir) Open(name string) (File, error) {
  //...
}

type FileSystem interface {
  Open(name string) (File, error)
}

// http.FileServer()方法返回的是fileHandler实例，fileHandler结构体实现了Handler接口的方法 ServeHTTP()

// ServeHTTP 方法内的核心是 serveFile() 方法

(1) handler接口
  type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
  }

(2) FileServer返回handler
  func FileServer(root FileSystem) Handler {
    return &fileHandler{root}
  }

(3) fileHandler实现handler接口
  type fileHandler struct {
    root FileSystem
  }
  func (f *fileHandler) ServeHTTP(w ResponseWriter, r *Request) {
    upath := r.URL.Path
    if !strings.HasPrefix(upath, "/") {
      upath = "/" + upath
      r.URL.Path = upath
    }
    http.serveFile(w, r, f.root, path.Clean(upath), true)
  }


//实例
var rootPath = "/data_dir/malx/test/"

func down(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>>")
	mpath := path.Join(rootPath,r.URL.Path)
	fmt.Println(mpath)
	http.ServeFile(w, r, mpath)
}

http.HandleFunc("/",down)
```

### url
```
r.URL *url.URL
```


### net/http
(1) 正确姿势
```
package main

import (
 "net/http"
)

func main() {

 http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){


   w.Header().Set("name", "my name is smallsoup")
   w.WriteHeader(500)
   w.Write([]byte("hello world\n"))

 })

 http.ListenAndServe(":8080", nil)
}
```

### 前端显示图片
```
func showPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("C:\\Users\\malx\\Pictures\\timg.jpg")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	buf := make([]byte, 1024000)

	_, err1 := f.Read(buf)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Disposition", "inline; filename=\"picture.png\"")
	w.Write(buf)
}
```

### 登陆
```
func login(client *http.Client) {
	my := User{Userid: "malx", Password: "123456"}
	x, _ := json.Marshal(my)
	fmt.Println(string(x))
	req, _ := http.NewRequest("POST", "http://10.10.100.14:8081/login", bytes.NewReader([]byte(string(x))))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "name=malx")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	cookie = resp.Header["Set-Cookie"][0]
	fmt.Println(cookie)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

### client timeout
```golang

// [Go net/http 超时机制完全手册] https://colobu.com/2016/07/01/the-complete-guide-to-golang-net-http-timeouts/
// net.Dialer.Timeout                   限制建立TCP连接的时间
// http.Transport.TLSHandshakeTimeout   限制 TLS握手的时间
// http.Transport.ResponseHeaderTimeout 限制读取response header的时间
// http.Transport.ExpectContinueTimeout 限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。注意在1.6中设置这个值会禁用HTTP/2(DefaultTransport自1.6.2起是个特例)

// 上述不能限制发送request的时间，除非服务端设置了超时


```
