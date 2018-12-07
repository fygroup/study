#### cookie与签名cookie
```
//Linux随即密码生成mkpasswd b6!wOpiU7
app.use(require('cookie-parse')('b6!wOpiU7'))
res.cookie('name','malx')
res.cookie('name_sign','malx',{signed:true}) //签名的cookie
var name = req.cookies.name
var name_sign = req.signedCookies.malx_sign
//清除cookies
res.clearCookie('malx')

//其他选项
domain 
path
maxAge
secure
httpOnly
signed

```

---
#### session
服务器上的session ID是储存在cookie中的
```

```

---
#### express http https
调用express()返回得到的app实际上是一个JavaScript的Function
```
var app = express()
http.createServer(app).listen(8111)
https.createServer(options,app).listen(8111)

app.listen = function(){
    var server = http.createServer(this)  //this 表示当前app实例
    server.listen.apply(server,arguments)
}

```

---
#### trust proxy与负载均衡
```

```

---
#### mongoose
```
var mongoose = require('mongoose')
const DB_URL = 'mongodb://localhost:27071/test'  
//                        host:port      数据库
mongoose.connect(DB_URL,{user:'malx',pass:'123456'}}) //连接数据库


```
