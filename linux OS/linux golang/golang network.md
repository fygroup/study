### JSW
```
session加密，将其放到客户端。本地只保留密钥，验证只需要找到对应的密钥，再解密session获得用户信息


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



### serveFile
```
type Dir string

func (d Dir) Open(name string) (File, error) {
  ...
}

type FileSystem interface {
  Open(name string) (File, error)
}

http.FileServer()方法返回的是fileHandler实例，fileHandler结构体实现了Handler接口的方法 ServeHTTP()

ServeHTTP 方法内的核心是 serveFile() 方法

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

