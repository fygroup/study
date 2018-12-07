#### 镜像安装定义
`npm.cmd install -g cnpm --registry=https://registry.npm.taobao.org`

---
#### 安装包
```
export NODE_HOME=/usr/local/node-v10.12.0-linux-x64/
export PATH=$NODE_HOME/bin:$PATH
cnpm install express --save
npm install -g express //安装到全局路径
```

---
#### for
```
for (i of [1,2,3]) console.log(i)
1
2
3
for (i in [1,2,3]) console.log(i)
0
1
2
```

---
#### 灵活性
```
a={}
a.a=1
a=function(){}
a.a =1
b=new a()
b.a //undefined
```

---
#### 闭包
```
for(i=0;i<5;i++){
    setTimeout((function(){
        var c=i;
        return function(){
            console.log(i);
        }})(),1000);
}
```

---
#### hasOwnProperty(判断自身属性)
```
a={1:1,2:2}
a.hasOwnProperty(1)
function f(){
	this.a=1
}
f.prototype.b = 1 
var b = new f()
b.hasOwnProperty('a') //true
b.hasOwnProperty('b') //false  b是继承属性
```

---
#### arguments
```
function a(x){
	console.log(x)
	console.log(argument)
}
a(1)
a(1,2)
1
[Arguments] { '0': 1, '1': 2 }
```

---
#### events
```
const events = require('events')
var eventEmitter = new events.EventEmitter()
eventEmitter.on('eventName',eventHander)  //绑定事件及事件处理程序
eventEmitter.emit('eventName') //触发事件

eventEmitter.on('eventName',function(argv1,argv2){
    console.log('event1');
});
eventEmitter.on('eventName',function(argv1,argv2){
    console.log('event1');
});
eventEmitter.emit('eventName',argv1,argv2) //依次调用上面的两个函数

var listen = function(){}
eventEmitter.addListener('name',listen)
eventEmitter.addListener('name',listen)

eventEmitter.removeAllListeners('eventName')  //删除事件
```

---
#### TypedArray(原生javascript)
```
var buffer = new ArrayBuffer(12)   //创建一块12字节的内存
buffer.byteLength
var arr = new Int16Array(buffer)  //这块内存定义为int16的数组内存，Int16Array(buffer,0,2)
arr.length
Int8Array：数组每一个元素的类型为8位带符号整数。
Uint8Array：数组每一个元素的类型为8位不带符号整数。
Uint8ClampedArray：数组每一个元素的类型为8位不带符号整数（自动过滤溢出）。
Int16Array：数组每一个元素的类型为16位带符号整数。
Uint16Array：数组每一个元素的类型为16位不带符号整数。
Int32Array：数组每一个元素的类型为32位带符号整数。
Uint32Array：数组每一个元素的类型为32位不带符号的整数。
Float32Array：数组每一个元素的类型为32位浮点数。
Float64Array：数组每一个元素的类型为64位浮点数。
```

---
#### buffer
```
Node.js 可以在一开始就使用 --zero-fill-buffers 命令行选项强制所有使用  新分配的 Buffer 实例在创建时自动用 0 填充

const buf = Buffer.alloc(10, 1);
const buf = Buffer.from([1, 2, 3]);
const buf = Buffer.from('test');
buf.length

buf.write("dadasfafafagag"); //写([encoding[, start[, end]]])
buf.toString() //读 buf.toString('ascii',0,5)
JSON.stringify(Buffer.from("12345"))  //'{"type":"Buffer","data":[49,50,51,52,53]}'
Buffer.concat([buf1,buf2]) //合并

buffer具有compare copy slice的功能

const arr = new Uint16Array(2);
// 拷贝 `arr` 的内容
const buf1 = Buffer.from(arr);
// 与 `arr` 共享内存
const buf2 = Buffer.from(arr.buffer);
const buf = Buffer.alloc(11, 'aGVsbG8gd29ybGQ=', 'base64');  //(size[, fill[, encoding]])
```

---
#### fs
`var fs = require('fs')`
* 1、删除文件
```
fs.unlink('xxxx',function(err){    //异步
if (err) throw err;
	console.log("成功删除")；
})
fs.unlinkSync('xxxxxxx')  //同步删除
console.log(''成功删除")
```

* 2、readFile
```
fs.readFile(filename,[options],callback);
                        option Object
                            encoding String |null default=null
                            flag String default='r'
fs.readFile('xxxxx','utf-8',function(err,data){
    console.log(data.toString());   //注意data是个二进制流
});
fs.readFileSync(filename,[options]);   //同步删除
```

* 3、writeFile
```
fs.writeFile(filename,data,[options],callback);
                        data String|buffer
                            option Object
                                encoding String |nulldefault='utf-8'
                                mode Number default=438(aka 0666 in Octal)
                                flag Stringdefault='w'
fs.writeFile('xxxxx','hello world',{'flag':'a'},function(err){   //flag传值，r代表读取文件，w代表写文件，a代表追加。
});
```

* 4、open、read、write
```
fs.open(path,flags[,mode],callback(err,fd));

path 文件路径
flags打开文件的方式(r ：读取文件，文件不存在时报错；r+ ：读取并写入文件，文件不存在时报错；
rs ：以同步方式读取文件，文件不存在时报错；rs+ ：以同步方式读取并写入文件，文件不存在时报错；
w ：写入文件，文件不存在则创建，存在则清空；wx ：和w一样，但是文件存在时会报错；
w+ ：读取并写入文件，文件不存在则创建，存在则清空；wx+ ：和w+一样，但是文件存在时会报错；
a ：以追加方式写入文件，文件不存在则创建；ax ：和a一样，但是文件存在时会报错；
a+ ：读取并追加写入文件，文件不存在则创建；ax+ ：和a+一样，但是文件存在时会报错。)
[mode] 是文件的权限（可行参数，默认值是0666）
callback 回调函数

fs.close(fd,callback);
fs.read(fd,buffer,offset,length,position,callback);
fd 文件描述符，必须接收fs.open()方法中的回调函数返回的第二个参数。
buffer 是存放读取到的数据的Buffer对象。
offset 指定 向buffer中存放数据的起始位置。
length 指定 读取文件中数据的字节数。
position 指定 在文件中读取文件内容的起始位置。
callback 回调函数，参数如下
    err 用于抛出异常
    bytesRead 从文件中读取内容的实际字节数。
    buffer 被读取的缓存区对象。

fs.open('xxx','r',function(err,fd){
    var buf = Buffer.alloc(30);
    fs.read(fd,buf,0,30,0,function(err,bytesRead,buffer){
        console.log(bytesRead,buffer.slice(0,bytesRead).toString());
    })
})

fs.write(fd,buffer,offset,length,position,callback(err,bytesWritten,buffer))接收6个参数。
fd 文件描述符，必须接收fs.open()方法中的回调函数返回的第二个参数。
buffer 是存放 将被写入的数据，buffer尺寸的大小设置最好是8的倍数，效率较高。
offset  buffer写入的偏移量。
length (integer)指定 写入文件中数据的字节数。
position (integer) 指定 在写入文件内容的起始位置。
callback 回调函数，参数如下
err 用于抛出异常
bytesWritten从文件中读取内容的实际字节数。
buffer 被读取的缓存区对象。
```

* 5、文件信息
```
fs.stat(path,callback);
fs.stat(path,function(err,stat){
    stats.isFile()  如果是文件返回 true，否则返回 false。
    stats.isDirectory() 如果是目录返回 true，否则返回 false。
    stats.isBlockDevice()   如果是块设备返回 true，否则返回 false。
    stats.isCharacterDevice()   如果是字符设备返回 true，否则返回 false。
    stats.isSymbolicLink()  如果是软链接返回 true，否则返回 false。
    stats.isFIFO()  如果是FIFO，返回true，否则返回 false。FIFO是UNIX中的一种特殊类型的命令管道。
    stats.isSocket()    如果是 Socket 返回 true，否则返回 false。
    console.log(stat);
            //dev: 16777220,
            //mode: 33188,
            //nlink: 1,
            //uid: 501,
            //gid: 20,
            //rdev: 0,
            //blksize: 4096,
            //ino: 40333161,
            //size: 61,
            //blocks: 8,
            //atime: Mon Sep 07 2015 17:43:55 GMT+0800 (CST),
            //mtime: Mon Sep 07 2015 17:22:35 GMT+0800 (CST),
            //ctime: Mon Sep 07 2015 17:22:35 GMT+0800 (CST)
});
```
* 6、创建目录
`fs.mkdir(path[,mode],callback(err))  //mode - 设置目录权限，默认为 0777。`

* 7、删除目录
`fs.rmdir(path,callback(err))`

* 8、读取目录
`fs.readdir(path,callback(err,files)) //err 为错误信息，files 为 目录下的文件数组列表。`

* 9、移动文件
`fs.renameSync('aaa','test/aaa')`

---
#### instanceof
```
function A(){}
a=new A();
a instanceof A  //true
```

---
#### stream
```
服务器的请求和process.stdout都属于流操作，流都是运作在字符串和 Buffer（或 Uint8Array）上。

四种类型
stream.Writeable，stream.Readable，stream.Duplex，stream.Transform
Writable - 可写入数据的流（例如 fs.createWriteStream()）。
Readable - 可读取数据的流（例如 fs.createReadStream()）。
Duplex - 可读又可写的流（例如 net.Socket）。response require
Transform - 在读写过程中可以修改或转换数据的 Duplex 流（例如 zlib.createDeflate()）。

混合使用 on('data')、on('readable')、pipe() 或异步迭代器，会导致不明确的行为
建议使用流的.pipe()方法，这样就不用自己监听”data” 和”end”事件了，也不用担心读写不平衡的问题了
```
<img src='../picture/3.png' height=200 width=300 alt='stream events function'/>

* readStream:
<img src='../picture/4.png' height = 200 width=250 alt='reastream 图示'/> <img src='../picture/5.png' height = 200 width=250 alt='reastream 图示'/> 
```
暂停模式 <-->  流动模式 (相互转换)
'data':'data' 事件会在流将数据传递给消费者时触发
'readable': 事件将在流中有数据可供读取时触发（流（缓存）有了新数据会触发）
'readable':可读流有数据可以读取时，会触发此事件，然后调用read()读取缓存数据
readable.read() 且有数据块返回时，也会触发 'data' 事件（当缓冲满时，激发readable,然后用read()读取缓存数据，此时就会激发data）
如果使用 readable.setEncoding() 为流指定了默认的字符编码，则监听器回调传入的数据为字符串，否则传入的数据为 Buffer
'end' 事件将在流中再没有数据可供消费时触发
push(): 压入流（缓冲）
read()：读取缓存中的数据，当读到尾部时，返回null
_read(); 内部从文件读取函数，注意和read()区分
unpipe() pause(): 暂停从资源库读取数据，但不会暂停数据生成,pause暂停流的读入。
pipe() resume(): 正在从资源库中读取数据，监听 'data' 事件 ,恢复流的读入
```

* writeStream:
<img src='../picture/6.png' alt='writable 图例' height = 200 width = 250 />
```
'drain': 当可写流可以接收事件的时候被触发，即当缓冲区可写的时候
'finish'：当所有数据被接收时被触发
write(): 向流(缓存)中写入数据
a._readableState.highWaterMark = 222222 设置缓存
writable.writableBuffer 或 readable.readableBuffer获得缓存内容
```
**注意: 一定要注意可读流和可写流读和写之间的平衡,如果可写流的写速度比较慢，会导致大量的buff缓存在内存，所以尽量用pipe**

例子：
* (1)读
```
var readStream = fs.createReadStream(file);
readStream._readableState.highWaterMark = 10;  //设置缓存
readStream.on('data',function(data){			
	console.log('readStream==>'+data);
});
readStream.on('readable',function(){    //当缓存满的时候，先激发readable,
	let chunk;
	let str = '';
	while(null != (chunk=readStream.read(1))){  //read()的时候激发data
		str+=chunk;
	}
	console.log(str);
});
readStream.on('end',function(){});
```

* (2)写(注意读写平衡)
```
http.createServer(function (req, res) {
    var stat = fs.statSync(filename);
    res.writeHeader(200, {"Content-Length": stat.size});
    var fReadStream = fs.createReadStream(filename);
    fReadStream.on('data', function (chunk) {
        if(!res.write(chunk)){//判断写缓冲区是否写满(node的官方文档有对write方法返回值的说明)
            fReadStream.pause();//如果写缓冲区不可用，暂停读取数据
        }
    });
    fReadStream.on('end', function () {
        res.end();
    });
    res.on("drain", function () {//写缓冲区可用，会触发"drain"事件
        fReadStream.resume();//重新启动读取数据
    });
});

```

* (3)自定义
```
var Readable = require('stream').Readable;
var Writable = require('stream').Writable;
class myread extends Readable{
	constructor(opt){
		super(opt);
		this.i=0;
		this.max = 5;
		this.array = new Array(1,2,3,4,5);
	}
	_read(){
		if (this.i>=this.max){
			this.push(null);
		}else{
			let buf = Buffer.from(this.array[this.i]+'','ascii');
			this.push(buf);
		}
		this.i++;
	}
}

var aa = new myread();
aa._readableState.highWaterMark = 1;
aa.on('data',function(chunk){
	console.log(chunk.toString());
})

aa.on('end',function(){
	console.log('endl');
})
```

---
#### 按行读取
```
var readline = require('readline')
var spawn = require('child_process').spawn
var readStream = fs.createReadStream(file);
var job = spawn(‘du’,['--max-dep','1','-h','/']);

var r1 = readline.ceateInterface({
	input: job.stdout/readStream   //流的活学活用！！！  
})
r1.on('line',function(line){
	console.log(line)
})
```

---
#### glob

---
#### async_hooks 异步钩子

---
#### console.log
这是异步操作！！！！！

---
#### child_process
```
const {spawn} = require('child_process');
const ls = spawn('ls',['-lh','./']);  //注意如果这块用ll，会报错
```

---
#### 事件观察者
**事件的执行先后** 
`idle > IO > check`
```
idle: procss.nextTick(callback)  //事件保存在一个数组中，会将数组中的事件执行完，进行下一轮Tick
IO:
check: setTimeOut() setInterval() //事件保存在一个链表中，执行完当前一个，进行下一轮Tick
```

---
#### apply call
call直接使用参数列表，apply使用参数数组
使用call()和apply()方法时，就会改变this的指向
* (1)
```
var pet = {
    words:'...',
    speak:function(somebody){
        console.log(somebody+"speak"+this.words);
    }
}
var dog = {
    words:"wang"
}
pet.speak.call(dog,'dog'); //dog speak wang
```
* (2)继承
```
fucntion pet(name){
	this.name = name;
	this.speak = function(){
		console.log(this.name);
	}
}
function dog(name){
	pet.call(this,name)
}
var my = new dog('malx')
my.speak()
```

---
#### util.inherits
`uitl.inherits(sub, super);  `
注意：sub仅仅继承super.prototype的内容
例如：
```
function A(){
	this.funa=function(){}
}
A.prototype.funa1 = function(){}
function B(){}
util.inherits(B,A);
var my = new B();
my.funa() //报错
my.funa1() //正确
```
如果要利用此函数做继承，见下一条

---
#### 继承作用域
```
function pet(){
	var func = function(){}
}
function dog(){
	pet.call(this);
}
util.inherits(dog,pet);
var my = new dog();
my.func();
```
所以说：除了用inherits继承super的prototype,还要继承super的作用域！！！

---
#### promise
```
var a = new Promise(function(resolve,reject){
	fs.readFile('xxx.txt','utf-8',function(err,data){
	if (err) reject(err);
		else resolve(data);
	})
});
a.then(function(data){
	console.log(data)
}).catch(function(err){
	console.log(err.message)
})
```
如果返回promise，它会在异步操作完成后发信号给下一个then。返回值并不是非promise不可，不管返回什么，都会传给下一个onFulfilled做参数：

---
#### Q
```
var a = function(file){
	const Q = require('q')
	var defer = Q.defer
	fs.readFile(file,function(err,data){
		defer.resolve(data)
	})
	return(defer.promise)
}
a.all([a(file1),a(file2)]).then(x){
	x[0]
	x[1]
}
```

---
#### async
```
async.series([
	function(callback){
		console.log(">>>>>>1");
		callback('',2);
	},
	function(callback){
		console.log(">>>>>>2");
		fs.readFile('test.txt','utf-8',callback);
	}
	],function(err,data){
		console.log(err);
		console.log(data);
	})
async.parallel([
	function(callback){
		console.log(">>>>>>1");
		callback(1,2);
	},
	function(callback){
		console.log(">>>>>>2");
		fs.readFile('test.txt','utf-8',callback);
	}
	],function(err,data){
		console.log(err);
		console.log(data);
	})
async.parallelLimit(pool,2000,function(err,data){
	console.log('Done')
	console.log(data.length)
	console.timeEnd(">>>")
})
```

---
#### bagpipe(此包我已经修改，详见报的注解)
```
var pool = new bagpipe(2000);
console.time('>>>')
for (let i=0;i<10000;i++){
	pool.push(fs.readFile,'7E6789_L1_I378.fastq','utf-8',function(err,data){})
}
pool.push(function(){    //结束时的函数
	console.log('Done');
	console.timeEnd('>>>')
})
```

---
#### v8内存限制
```
--max-old-space-size 1000   //单位Mb 老生代
--max-new-space-size 1000   //单位Kb 新生代
```

---
#### stream 读取长度限制
`fs.createReadStream('txt',{highWaterMark:11})`

---
#### 传输层(TCP/UDP)
##### TCP
TCP服务分为服务器事件和连接事件（不是客户端事件），
```
服务器事件：
listening:绑定端口后触发, server.listen(3134,()=>{})
connection:客户端套接字连接到服务端时触发
close:服务器关闭时触发，server.close()
error:

连接事件：stream对象，比如
data: 一端调用write()时，另一端会触发data事件(socket.on('data') client.on('data') socket.write('dada') client.write('dadad'))
end: 任意一端发送FIN数据，会触发此事件
connect: net.connect()
drain: 任意一端调用write发送数据时，触发此事件,当writeStream可写时触发事件
error:
close:套接字完全关闭时触发
timeout:一定时间后连接不再活跃时触发。
```

(1)服务端
```
var net = require('net');
var server = net.createServer();
server.on('connection',(socket)=>{
	socket.on('data',(data)=>{
		console.log(data.toString());
		socket.write('你好');  //向套接字里写入 用于发送
	});
	socket.on('end',()=>{
		console.log('连接断开')
	})
})
server.listen(8124);
```

(2)客户端
```
var net = require('net');
var client = net.connect({host:'localhost',port:8124},()=>{
	console.log("client connect!");
	client.write('hello');  //发送内容
})
client.on('data',(data)=>{
	console.log(data.toString());
	client.end();
})
client.on('end',()=>{
	console.log('client disconnect!');
})
```
`net.client({path:'/tmp/echo.sock'})`

(3)pipe操作
```
套接字是stream,可以利用pipe()
echo 服务器
服务端
var net = require('net');
var server = net.createServer((socket)=>{
	socket.write('Echo server\r\n');
	socket.pipe(socket);
});
server.on('listening',function(){
	console.log('listen:8111')
})
server.listen(8111)

客户端
var client = net.connect({port:8111},function(){
	console.log('connect server');
})
process.stdout.setEncoding('utf-8');
client.pipe(process.stdout)
process.stdin.on('data',(data)=>{
	client.write(data)
})
```
##### UDP
(1)服务端
```
var dgram = require('dgram')
var server= dgram.createSocket('udp4')
server.on('message',(mes,info)=>{   //info 携带地址和port
	console.log(msg)
	console.log(info.adress,info.port)
})
server.on('listening',()=>{
	let address = server.address() 
	console.log('server listen',address.address,address.port)
})
server.bind(8142)
```

(2)客户端
```
var message = Buffer.from('fafafafafafaa')
var client = dgram.createSocket('udp4')
client.send(message,0,message.length,8142,"localhost",(err,bytes)=>{ //bytes为发送了多少字节
	client.close()
})
```

---
#### http
http报文分文三个部分
第一部分:经典的TCP的3次握手
第二部分:完成握手后，客户端向服务端发送报文
第三部分:服务器完成处理后，向客户端发送响应内容，包括响应头和响应体（post请求）

##### 服务器
```
(1)http.Server      //server = http.createServer()或server = http.createServer((req,res)=>{})
server.on('request',(req,res)=>{})     //当客户端请求到来时，该事件被触发,提供两个参数req和res
server.on('connection',(socket)=>{})   //当TCP连接建立时，该事件被触发，提供一个参数socket
server.on('close')                     //当服务器关闭时，触发事件（注意不是在用户断开连接时）
server.on('upgrade',(req,socket,head)=>{})

(2)http.IncomingMessage      //request
req.on('data',(chunk)=>{})   //当请求体数据到来时，该事件被触发
req.on('end')                //当请求体数据传输完毕时，该事件会被触发
req.on('close')              //用户当前请求结束时，该事件被触发

(3)http.ServerResponse      //reponse
res.writeHead(statusCode,[heasers])
res.write(data,[encoding])
res.end([data],[encoding])

(4)事件
var http = require('http')
var server = http.createServer()
server.on('request',(req,res)=>{          //注意事件是'request',不是'connection'
	console.log(">>>"+req.headers.cookie);
	req.on('data',(data)=>{				//注意 只有请求为'post'时，才会启发这个data事件，所以客户端请求必须是method:'post'
		console.log(data.toString())
	})

	req.on('end',()=>{
		console.log('传输end');
	})

})

server.on('close',()=>{
	console.log('服务器已关闭')
})

server.on('listening',()=>{
	address = server.address()
	console.log("%d",address.port);
})


server.on('error',(err)=>{
	console.log(err);
})

server.listen(8111)
```

##### 客户端
http.request和http.get
```
(1)http.request
options={
	method: 'POST',      //注意
	hostname:'localhost',
	port:8111
	path: '/a/b',
	headers: {
        "connection": "keep-alive",     //长连接
    	"content-type": "text/plain; charset=utf-8"
    }
}
var req = http.request(options,(res)=>{
	res.setEncoding('utf-8')
	res.on('data',(chunk)=>{})
	res.on('end',()=>{
		console.log('传输结束')
	})
	res.on('close',()=>{
		console.log('连接结束')
	})
	res.pause()     //暂停结束数据和发送事件，方便实现下载功能 ???
	res.resume()    //从暂停的状态中恢复	
	res.statusCode  //状态码 200 404 500
	res.headers     //http请求头
	res.trailers    //http请求尾
	res.httpVersion //http协议版本 1.0或1.1
})
var req = http.request(options)
req.on('response',(res)=>{})
req.write('dafafagagagaga')
req.end()					//所有写操作都必须调用end函数来通知服务器，否则请求无效
```

##### 模仿浏览器访问
```
const options = {
  hostname: 'nodejs.cn',
  port: 80,             //http 默认端口是80
  path: '/api/http.html#http_http_request_options_callback',
  method: 'GET',
  headers: {
      "Connection": "keep-alive",
      //"Content-Length": 111,
      "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
      "User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
  }
};

const req = http.request(options, (res) => {
  console.log(`状态码: ${res.statusCode}`);
  console.log(`响应头: ${JSON.stringify(res.headers)}`);
  res.setEncoding('utf8');
  res.on('data', (chunk) => {
    console.log(`响应主体: ${chunk}`);
  });
  res.on('end', () => {
    console.log('响应中已无数据。');
  });
});

req.on('error', (e) => {
  console.error(`请求遇到问题: ${e.message}`);
});

// 写入数据到请求主体并发送
req.end();
```

重定向处理
```
var http = require('http')

var option = {
    hostname:'localhost',
    port:8080,
    path:'/cnv-dev'
}

var reqs = []

var client = function(option){
    this.req = http.request(option)
    reqs.push(this.req)
    this.req.on('response',(res)=>{
        console.log(res.statusCode)
        if (res.statusCode == 302){
            console.log(res.headers.location);
            option.path = res.headers.location
            new client(option)
            return
        }

        res.on('data',(body)=>{
            console.log(body.toString())
        })
    })

    this.req.end()
}
new client(option)


setTimeout(()=>{
    console.log(reqs.length)
},2000)

```

##### request和response
```
(1)http.ClientRequest
属性：
ReadableState：该属性是一个ReadableState类型的对象，保存了Readable实例的重要信息，如读取模式（是否为对象模式）objectMode、highWaterMask（缓冲区存放的最大字节数）、缓冲区、flowing模式等
domain:捕捉异步回调中出现的异常。
headers:HTTP请求头，主要记录了主机地址(host),用户代理(user-agent),以及文件类型(content-type),以及接受的文件类型(accept)等等
trailers:HTTP 请求尾（不常见）
connection:当前 HTTP 连接套接字，为 net.Socket 的实例
url:记录访问网站的路径及参数
method:记录请求方法
statusCode:HTTP请求状态码
statusMessage:HTTP请求状态信息
socket:connection 属性的别名
client:client 属性的别名
complete：客户端请求是否已经发送完成
httpVersion：HTTP 协议版本，通常是 1.0 或 1.1
事件：
end:当请求体数据传输完成时，该事件被触发，此后将不会再有数据到来。
data:当请求体数据到来时，该事件被触发。该事件提供一个参数
close:用户当前请求结束时，该事件被触发。不同于end，如果用户强制终止了传输，也还是调用close。
chunk:表示接收到的数据。如果该事件没有被监听，那么请求体将会被抛弃。该事件可能会被调用多次。
error:用户当前请求发生异常错误的时候调用
drain:数据的长度大于 highWaterMark的时候，那么 Stream.write 就会返回 false，也就会触发 drain 事件了
函数：
request.abort() 终止正在发送的请求
request.setTimeout(timeout,[callback]) 设置请求超时时间，timeout为毫秒数，当请求超时后，callback将会被调用

(2)http.ClientResponse
http.ClientReponse是与http.ServerResponse相似，提供三个事件，data、end和close，分别在数据到达，传输结束和连
接结束时触发，其中data事件传递一个参数chunk，表示接受到的数据属性，表示请求的结果状态
statusCode   HTTP状态码，如200，404，500
httpVersion：HTTP协议版本
headers：HTTP请求头
trailers：HTTP请求尾
函数：
response.setEncoding([encoding])：设置默认的编码，当data事件被触发时，数据将以encoding编码。默认值为null，以buffer的形式存储。
response.pause()：暂停接受数据和发送事件，方便实现下载功能。
response.resume()：以暂停的状态中恢复
```

##### http echo服务器
```
//client
var res = http.request(options)
res.on('response',(res)=>{

  res.on('data',(chunk)=>{
    console.log(chunk.toString())
  })

})
//process.stdin.pipe(res) 推荐用这个
process.stdin.on('data',(chunk)=>{
  res.write(chunk)
})

//server
server.on('request',(req,res)=>{
	req.on('data',(data)=>{
		console.log('request data')
		res.write(data)
	})
	req.on('end',()=>{
		console.log('request end');
	})
	//req.pipe(res)   推荐用这个

})
```

---
#### http与TCP
http模块时tcp的connection到http的request进行的封装
```
server.on('connection',(socket)=>{
	socket.on('data',(chunk)=>{})   //输出的是header和body
	socket.on('end',()=>{})			
	socket.on('close',()=>{})
})

server.on('request',(req,res)=>{
	req.on('data',(data)=>{})    //只输出body
	req.on('end',()=>{})	
	req.on('close',()=>{})
})
//每一次客户的访问会有以下过程
TCP data
request data
request end
TCP end
TCP close

```

---
#### http客户端代理连接（不是代理服务）
客户端：请求-->代理-->结束
```
var agent = new http.Agent({
	maxSockets:10
})
//客户端同时对一个服务器发起10次请求，由于代理池的存在，实质只有5个处于并发状态，这与浏览器对同一域名有下载连接数的限制时相同的行为
var options = {
	hostname:'127.0.0.1',
	port:8111,
	path:'/',
	method:'post',
	agent: agent   //如果为false,脱离代理池的管理，不受并发限制
}
```

---
#### WebSocket
比http协议更好，html5的重要特性，解决HTTP协议本身的单向性问题，优点包括：
```
(1)客户端与服务端只建立一个TCP连接，可以使用更少的连接
(2)websocket服务器可以推送数据向客户端，比http轮询的方式更灵活更快
(3)更轻及的协议头，减少数据传送量
注意：WebSocket不是替代http，而是http的补充。
```
```
if(window.WebSocket != undefined) { //判断浏览器是否支持websocket
  var ws = new WebSocket('ws://localhost:1740');
}
ws.readyState //表示连接状态
0： 正在连接
1： 连接成功
2： 正在关闭
3： 连接关闭
```
websocket先进行握手，将普通的http协议升级为websocket协议，然后再进行数据传输

---
#### Socket.io (websocket模块)
详情见[socket.io文档](https://www.w3cschool.cn/socket/)
服务端
```
const io = require('socket.io')(server, {
  path: '/test',  //捕获路径test
  serveClient: false,
  // below are engine.IO options
  pingInterval: 10000,
  pingTimeout: 5000,
  cookie: false
})

(1)http方式
var httpServer = require('http').createServer()
var io = require('socket.io')(httpServer) //require('socket.io')(8111)
httpServer.on(...)
httpServer.listen(8111)

io.on('connection',(socket)=>{
	socket.emit('news',{hello:'world'});
	socket.on('my other event',(data)=>{
		console.log(data.toString())
	});
	scoket.on('disconnection')
})

(2)express方式
var app = require('express')()
var server = require('http').Server(app)
var io = require('docket.io')(server)
server.listen(8111)
io.on(....)
```
客户端
```
var io = require('socket.io-client')
var socket = io.connect('http://localhost:8111')

socket.on('client_new',(data)=>{
	console.log(data.toString())
	socket.emit('server_new','>>>'+data.toString())
})
```

---
#### TLS/SSL
数字证书的生成
```
CA机构
openssl genrsa -out ca.key 1024  //生成私钥
openssl req -new -key ca.key -out ca.csr  //此过程会提示输入信息
openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt  //生成CA证书,然后将ca证书和私钥通过安全的方式发送给服务器端

服务器端
openssl genrsa -out server.key 1024 //生成密钥
openssl req -new -key server.key -out server.csr //此过程会提示输入信息，注意common Name要匹配服务器域名，否则认证过程会出错
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt //利用CA证书和私钥颁发一个带有CA签名的证书。
证书:
证书的发布机构
证书的有效期
公钥
证书所有者（Subject）
签名所使用的算法
指纹以及指纹算法
```

---
#### 加密与签名
(1)公钥与私钥
```
公钥加密，私钥解密。私钥加密，公钥解密
```

(2)请求过程
```
step1:
客户->服务器：发送请求
step2:
服务器->客户：发送数字证书（包含公钥）
step3:
客户验证数字证书是不是服务器，有否问题。

```


---
#### https
服务端
```
var https = require('https')
var fs = require('fs')

var options = {
	path: '/ab',
	key: fs.readFileSync('server.key'),
	cert: fs.readFileSync('server.crt')
}

server = https.createServer(options,(req,res)=>{
	//req.on(..)
	res.writeHead(200,{'Content-Type':'text/plain'});
	res.write('hello world\n') 
	res.end()
})

server.listen(8111)

server.on('listening',()=>{
	address = server.address()
	console.log(address.address,address.port)

})
```
客户端
```
var https = require('https')
var fs = require('fs')

var options = {
	hostname: 'localhost',
	port: 8111,
	path: '/ab',
	method: 'GET',
	key: fs.readFileSync('client.key'),
	cert: fs.readFileSync('client.crt'),
	ca: [fs.readFileSync('ca.crt')],
	rejectUnauthorized: false //不保证服务器端的证书是否伪造
}

options.agent = new https.Agent(options) //注释这一句也能运行，不知为啥？？？

var req = https.request(options,(res)=>{
	res.setEncoding('utf-8')
	res.on('data',(data)=>{
		console.log(data.toString())
	})
})

req.end()
req.on('error',(e)=>{
	console.log(e)
})
```

---
#### ssh2
基础功能
```
/**
* 描述：连接远程机器
* 参数：server,远程机器凭证；
*       then,回调函数
*/
ssh2.connect(server, then)

/**
* 描述：断开远程连接
* 参数：then,回调函数
*/
ssh2.disconnect(then)

/**
* 描述：执行shell命令
* 参数：cmd,要执行的命令；
*       then,回调函数
* 回调：then(err, data):data 运行命令之后的返回信息
*/
ssh2.exec(cmd, then)



```

[ssh2实例](https://github.com/mscdex/ssh2#user-content-client-methods)
sftp
```
var Client = require('ssh2').Client;

var conn = new Client();
conn.on('ready', function() {
  console.log('Client :: ready');
  conn.sftp(function(err, sftp) {
    if (err) throw err;
    sftp.readdir('/seq_dir', function(err, list) {
      if (err) throw err;
      console.dir(list);
      conn.end();
    });
  });
})

conn.connect({
  host: '192.168.100.100',
  port: 22,
  username: 'frylock',
  password: 'nodejsrules'
});

```
exec
```
var Client = require('ssh2').Client;

var conn = new Client();
conn.on('ready', function() {
  console.log('Client :: ready');
  conn.exec('uptime', function(err, stream) {
    if (err) throw err;
    stream.on('close', function(code, signal) {
      console.log('Stream :: close :: code: ' + code + ', signal: ' + signal);
      conn.end();
    })
	stream.on('data', function(data) {
      console.log('STDOUT: ' + data);
    })
	stream.on('err', function(err) {
      console.log('STDERR: ' + err);
    });
  });
})
conn.connect({
  host: '10.100.6.20',
  port: 22,
  username: 'root',
  password: 'berry.com'
  //privateKey: require('fs').readFileSync('/here/is/my/key')
});
```

---
#### Web应用
##### 基础功能(路径解析、Cookie、Session、缓存、Basic)
(1)请求方法
```
req.method  //post or get
```

(2)路径解析
```
req.url = 'http://user:pass@host.com:8080/p/a/t/h?a=3&b=ewf'
x=url.parse(req.url,true)
Url {
  protocol: 'http:',
  slashes: true,
  auth: 'user:pass',
  host: 'host.com:8080',
  port: '8080',
  hostname: 'host.com',
  hash: null,
  search: '?a=3&b=ewf',
  query: { a: '3', b: 'ewf' },
  pathname: '/p/a/t/h',
  path: '/p/a/t/h?a=3&b=ewf',
  href: 'http://user:pass@host.com:8080/p/a/t/h?a=3&b=ewf' }

x.query[key] = value //字符串键值 GET
url.format(x)  //将x重新弄成url
```

(3)Cookie
```
var http = require('http')

//生成cookie
var serialize = function(name,val,opt){
	var pair = [name+'='+opt]
	opt = opt || {}          //必须要有！！！
	if (opt.maxAge) pair.push('Max-Age='+opt.maxAge); //如果设置过期时间，关闭浏览器并不删除cookie，直到时间过期
	if (opt.domain) pair.push('Domain='+opt.domain);
	if (opt.path) pair.push('Path='+opt.path);//当前访问路径不满足该匹配时，浏览器不发送cookie
	if (opt.expires) pair.push('Expires='+opt.expires.toUTCString());//设置过期时间，一般不用（时间可能不统一）
	if (opt.httpOnly) pair.push('HttpOnly'); //告知浏览器不可更改和不显示cookie
	if (opt.secure) pair.push('Secure'); //如果设置那么在http无效，在https才有效
	return pair.join(';')
}
//解析cookie
var parseCookie = function(cookie){
	var cookies = {}
	if (!cookie) return cookies;
	var list = cookie.split(';')
	for (let i=0;i<list.length;i++){
		var pair = list[i].split('=')
		cookies[pair[0].trim()] = pair[1]
	}
	return cookies;
}

const server = http.createServer((req,res)=>{
	req.cookie = parseCookie(req.headers.cookie);
	res.setHeader("Set-Cookie",serialize('isVip','1',{maxAge:'5000'}));
});
```
res的writeHeader和setHeader

(4)Session

服务端用session ID操作session的内容，Session ID 可以使用 Cookie 和 URL 参数来传递，一般放在Cookie中。
说白了，session就是服务器中的一个变量，但是在node会存在内存限制，加上多线程的内存不共享，所以session一般用Redis内存数据库。
session安全就是确保session ID的安全，利用私钥进行加密签名
```
var sign = function(val, secret){
	return val+'.'+crypto.createHmac('sha256',secret)
						 .update(val)
						 .digest('base64')
						 .replace('\=+$','');
}
```

(5)redirect
url重定向

(6)缓存
```
//ETag的请求和响应
function getHash(str){    //根据str内容生成hash
	var shasum = crypto.createHash('sha1');
	return shasum.update(str).digest('base64');
}

var handle = function(req,res){
	fs.readFile(filename,(err,data)=>{
		var hash = getHash(data);
		var noneMatch = req.headers['if-none-match']; //获得ETag的请求值
		if (hash == noneMatch){
			res.writeHead(304,"not modify");
			res.end();
		}else{
			res.setHeader('ETag',hash);  //添加hash值
			res.writeHead(200,"ok");
			res.end();
		}
	});
};

//Cache-Control 倒计时更新文件（前提是文件存在）
var handle = function(req,res){
	fs.readFile(filename,(err,data)=>{
		res.setHeader('Cache-Control',"max-age="+10*365*24*60*60*1000);
		res.writeHead(200,"ok");
		res.end(data);
	})
}
```

(7)Basic认证

Basic 认证是HTTP 中非常简单的认证方式，因为简单，所以不是很安全，不过仍然非常常用。
当一个客户端向一个需要认证的HTTP服务器进行数据请求时，如果之前没有认证过，HTTP服务器会返回401状态码，要求客户端输入用户名和密码。用户输入用户名和密码后，用户名和密码会经过BASE64加密附加到请求信息中再次请求HTTP服务器，HTTP服务器会根据请求头携带的认证信息，决定是否认证成功及做出相应的响应。
```
//Authorization: Basic AHV23ggCNVMsGa
function(req,res){
	var auth = req.headers['authorization'] || '';
	var parts = auth.split(' ');
	var method = parts[0]; //Basic
	var encode = parts[1];
	var decode = Buffer.from(encode,'base64').toString('utf-8').split(':'); //将base64转为string
	var user = decode[0];
	var password = decode[1];
	if (!checkUser(user,password)){
		res.setHeader('WWW-Authenticate','Basic realm="Secure Area"');   //设置认证头
		res.writeHead(401);
		res.end();
	}
}
```

##### 上传(表单、文件、JSON)
(1)表单
```
//<form action='/upload' method='post'>
//表单提交，请求头中一般包含，Content-Type: application/x-www-form-urlencoded
//报文体中的内容跟查询字符串相同，foo=bar&baz=val
var hander = function(req,res){
	if (req.headers['Content-Type'] === 'application/x-www-form-urlencoded'){
		req.body = querystring.parse(req.rawBody);  //req.rawBody
	}
	...
	todo(req,res);
}
```

(2)JSON
```
//Content-Type: application/json; charset=utf-8
var mime = function(req){
	var x = req.headers['Content-Type'] || ''
	return x.split(';')[0].trim();
}
var hander = function(req, res){
	if (mime(req) == 'application/json'){
		req.body = JSON.parse(req.rawBody);
	}
}
```

(3)附件上传
```
//<form action='/upload' method='post' enctype='multipart/form-data'>
//附件上传也是以post的形式上传的，不过在body中会有些不同，他会以boundary分割
//Content-Type: multipart/form-data; boundary=AaBo3x  //在报文体中boundary是每部分的分界符，
//Content-Length: 18231     //报文体长度
```
例如：

<img src='../picture/8.png' alt='muti/form-data' width=400 height=200>
<img src='../picture/9.png' alt='muti/form-data' width=400 height=200>

```
<form action='/upload' method='post' enctype='multipart/form-data'>
	<label for='username'>UserName:</label><input type='text' name='username' id='username'/>
	<label for='file'>FileName:</label><input type='file' name='file' id='file'/>
	<input type='submit' name='submit' value='Submit' />
</form>
//报文如下：
--AaBo3x\r\n
Content-Disposition: form-data; name='username'\r\n
\r\n
Malx
--AaBo3x\r\n
Content-Disposition: form-data; name='file'; filename='myfile.txt'\r\n
\r\n
...content of myfile.txt
--AaBo3x--
```
formidable模块
```
var formidable = require('formidable')
function(req, res){
	if (hasBody(req)){
		if (mime(req) == 'multipart/form-data'){
			var form = new formidable.IncomingForm();
			form.parse(req, (err, fields, files)=>{ 
				req.body = fields;  //其他数据们
				req.files = files;  //文件们
				handle(req,res);
			})
		}
	}else{
		handle(req,res);
	}
}
```

(4)数据上传与安全
```
//内存限制
//当content-length存在时
var byte = 1024;
function(req, res){
	var received = 0;
	var len = req.headers['content-length']? parseInt(req.headers['content-length'],10)::null; // parseInt
	if (len & len > byte){
		res.writeHead(413);
		res.end();
		return;
	}
	//当content-length不存在时，判断读入的数据长短
	req.on('data',(chunk)=>{
		received += chunk.length();
		if (received > byte){
			req.destory();
		}
	})
	handle(req,res);
}

//CSRF攻击(跨站请求伪造)

```

---
#### MVC
```
user --> Router --> Controller 
 |   					|
 |<-- view <-- Model <--|
```

---
#### RESTful
一种设计模式
```
POST 	/user/malx
DELETE	/user/malx
PUT		/user/malx
GET		/user/malx

app.post('/user/malx',addUser);
app.delete('/user/malx',removeUser);
app.put('/user/malx',updateUser);
app.get('/user/malx',getUser);
```

---
#### 中间件设计模式
目的
```
app.use(querystring)         //query是中间件函数
app.use('/a/b',querystring)
app.use(cookie)
app.use(session)
app.get('/user/:username',getUser);
app.put('/user/:username',authorize,updateUser);
//注册的函数
var querystring = function(req,res,next){
	req.query = url.parse(req,url,true).query;
	next();
}

var cookie = function(req,res,next){
	var cookie = req.headers.cookie;
	var cookies = {}
	if (cookie){
		var list = cookie.split(';')
		for (let i=0;i<list.length;i++){
			let x = list[i].split('=')
			cookies[x[0].trim()] = x[1].trim()
		}
	}
	req.cookie = cookies;
	next();
}
```

实现
```
var routes = {'all':[]}


//解析url
const urlreg = require('path-to-regexp') //url正则模块
var pathRegexp = function(path){
	var regexp = urlreg(path)
	var keys = []
	var reparse = urlreg.parse(path)
	for (let i=0;i<reparse.length;i++){
		if (typeof reparse[i] === 'object'){
			keys.push(reparse[i].name)
		}
	}
	return {keys: keys,
			regexp: regexp
		}
}

//注册函数
app.use = function(path){
	var handle;
	if (typeof path === 'string'){
		handle = {
			path: pathRegexp(path),
			stack:Array.prototype.slice.call(arguments,1)  //！！！请看javascript.md
		}
	}else{
		handle = {
			path: pathRegexp('/'),
			stack:Array.prototype.slice.call(arguments,0)
		}
	}
	route.all.push(handle);
}

['get','post','delete','put'].forEach((method)=>{
	route[method] = []    //针对method，要使用非'all'的属性
	app[method] = function(path){
		var handle;
		if (typeof path === 'string'){
			handle = {
				path: pathRegexp(path),
				stack: Array.prototype.slice.call(arguments,1)
			}
		}else{
			//error
		}
		routes[method].push(handle)
	}
})

//url match匹配
var match = function(pathname, cRoutes){
	var stacks = []
	for (let i=0;i<cRoutes.length;i++){
		var cRoute = CRoutes[i];
		var reg = cRoute.path.regexp;
		var matched = reg.exec(pathname);
		if (matched){
			var params = {}
			var keys = cRoute.path.keys;
			for (let i=0;i<keys.length;i++){
				var value = matched[i+1];     //是否多余
				if(value) params[keys[i]] = value;
			}
			req.params = params;
			stacks = stacks.concat(cRoute.stack);
		}
	}
}

//执行handle
var handle = function(req,res,stack){
	var next = function(){
		var middleware = stack.shift()  //取得中间件函数
		if (middleware){
			middleware(req,res,next);
		}
	}
	next();     //递归执行，知道stack为空
}

//最终的执行分发函数
http.createServer((req,res)=>{
	var pathname = url.parse(req.url).pathname;  //注意这里处理的是pathname
	var method = url.method.toLowerCase();  //method变小写
	var stacks = match(pathname, routes.all);
	if (routes.hasOwnPerperty(method)){
		stacks = stacks.concat(match(pathname, routes[method]))
	}

	if (stacks.length){
		handle(req,res,stacks);
	}else{
		handle404(req,res);
	}
})

```

---
#### 中间件异常处理
```
var handle = function(req,res,stack){
	var next = function(err){
		if (err){
			return(handle500(err,req,res,stack))
		}else{
			var middleware = stack.shift()
			if (middleware){
				try{
					middleware(req,res,next)
				}catch(ex){
					next(ex)
				}
			}
		}
	}
	next();
}

//捕获异常的中间件
var middleware = function(err,req,res,next){
	//...
	next();
}

//区分普通中间件与异常中间件
var handle500 = function(err,req,res,stack){
	stack = stack.filter((el)=>{
		return el.length == 4;
	})
	var next = function(){
		var middleware = stack.shift()
		if (middleware){
			middleware(err,req,res,stack);
		}
	}
	next();
}

```

---
#### 中间件性能
编写高效中间件、合理使用路由
```
//针对静态文件的访问，没必要next
var staticFile = function(req,res,next){
	var pathname = url.parse(req.url).pathname;
	fs.readFile(path.join(ROOT,pathname),(err,data)=>{
		if (err) return next();
		res.writeHead(200);
		res.end(data);

	})
}

app.use(staticFile); 
app.use('/public',staticFile); //最好加上一个路径，否则误伤率太高，每次访问都要读盘

```

---
#### 响应
res Header
```
Content-Encoding: gzip
Content-Length: 21107
Content-Type: text/javascript;charset=utf-8
Content-Type: text/plain;charset=utf-8
Content-Type: text/html;charset=utf-8
Content-Type: application/json;charset=utf-8

var mime = require('mime')
mime.lookup('aaa.txt')  //'text/plain'
mime.lookup('aaa.html')  // 'text/html'


```

附件下载
```
Content-Disposition: attachment; filename = 'filename.txt'
//附件下载API
res.sendile = function(filepath){
	fs.stat(filepath,(err,stat)=>{
		var stream = fs.createReadStream(filepath)
		res.setHeader('Content-Type',mime.lookup(filepath))
		//res.setHeader('Content-Type','application/octet-stream') //任何类型下载
		res.setHeader('Content-Length',stat.size)
		res.setHeader('Content-Disposition','attachment;filename=\"'+path.basename(filepath)+'\"')
		res.writeHead(200);
		stream.pipe(res);
	})
}
```

发送JSON
```
res.json = function(json){
	res.setHeader('Content-Type','application/json;charset=utf-8')
	res.writeHead(200)
	res.end(JSON.stringify(json))

}
```

跳转
```
res.redirect = function(url){
	res.setHeader('Location',url)
	res.writeHead(302)
	res.end('redirect to '+ url)
}
```

---
#### 异常处理
```
//domain
var domain = require('domain')
var EventEmitter = require('events').EventEmitter
var emitter = new EventEmitter()
var d = domain.create()
d.add(emitter)		//显示添加
d.on('error',(err)->{
	//err.message
})
emitter.emit('error',new Error('aaaaa'))  //注意emitter没有绑定error事件，但是它在domain域中运行，所以会触发d的error事件

d.run(()=>{						//隐式添加
    setTimeout(function () {
        throw new Error('async error'); // 抛出一个异步异常
    }, 1000);
});

d.on('error', function (err) {
    console.log('catch err:', err); // 这里可以捕获异步异常
});

//uncaughtException
//捕获整个进程的error
```

---
#### child_process
(1)基础用法
```
var cp = require('child_process')
cp.exec('node test.js',function(er,stdout,stderr){}) 
cp.fork('./test.js')

//进程间通信
//parent.js
var n = cp.fork(__dirname + '/sub.js')
n.on('message',function(m){})  //接收
n.send({hello:'world'})			//发送
//sub.js
process.on('message',function(m){})
process.send({foo:'bar'})

//句柄传递(主进程接受分配请求，子进程处理请求)
client			server
客户------>|----->port<----监听----Node
				|--发送---child1.js
				|--发送---child2.js
				|--发送---child3.js
//parent.js
var net = require('net')
var child1 = cp.fork('child.js')
var child2 = cp.fork('child.js')
var server = net.createServer()
server.listen(1337,()=>{
	child1.send('server',server)
	child2.send('server',server)
	server.close()
})
//child.js
var http = require('http')
var server = http.createServer()    //connection:当前 HTTP 连接套接字，为 net.Socket 的实例
server.on('request',(req,res)=>{    //详见http与TCP章节（tcp-->http | connection request）
	res.writeHead(200,{'Content-Type':'text/plain'})
	res.end('pid is ' + process.pid)
})   
process.on('message',(m,tcp)=>{
	if (m === 'server'){
		tcp.on('connection',(socket)=>{
			server.emit('connection',socket)
		})
	}
})
```

(2)进程事件和函数
```
work = fork()
1、exit
work.on('exit',(code,signal)=>{})
子进程退出时触发，如果正常退出code是退出码，否则为null。如果子进程被kill杀死，会得到signal信号
2、close
子进程标准IO中断时触发
3、disconnect
调用父进程的或子进程的 process.disconnect() 后会触发
4、message

//子进程
child = fork(...)
child.kill([signal])
//当前进程
pid = process.pid
process.kill(pid,[signal])
//例子
process.on('SIGTERM',()=>{
	console.log("Got SIGTERM, exiting...")
	process.exit(1)
})
process.kill(process.pid,'SIGTERM')
//注意kill方法并不是真正杀死子进程，他只是给子进程发送系统信号
```

(3)进程重启与限量重启
```
//master.js
var fork = require('child_process').fork
var server = require('http').createServer()
server.listen(8111)

var workers = {}
var restart = []
var limit = 5
var durning = 10000

var isTooFrequently = function(){
	var time = Date.now()
	var length = restart.push(time)
	if (length > limit){
		restart = restart.slice(-1*limit)
	}
	return restart.length >= limit && restart[restart.length-1] - restart[0] < durning
}

var createWork = function(){
	var worker = fork(__dirname + '/work.js')
	worker.on('message',(m)=>{
		if (m.act == 'suicide'){
			createWorker()
		}
	})

	worker.on('exit',()=>{
		console.log('Worker '+worker.pid + ' exit')
		delete workers[worker.pid]
	})

	worker.send('server',server)
	workers[worker.pid] = worker
	console.log('create worker '+worker.pid)
}

for (let i =0;i<10;i++){
	createWork()
}

process.on('exit',()=>{
	for (let pid in workers){
		workers[pid].kill()
	}
})

//work.js
var http = require('http')

var server = http.createServer((req,res)=>{
	res.writeHead(200,{'Content-Type':'text/plain;charset=utf-8'})
	res.end('work id '+ process.pid)
	throw new Error('error Exception')
})

var worker
process.on('message',(m,tcp)=>{
	if (m == 'server'){
		worker = tcp
		tcp.on('connection',(socket)=>{
			server.emit('connection',socket);
		})
	}
})

process.on('uncaughtException',(err)=>{
	console.log('worker '+ process.pid +  ' err: '+err )
	process.send({act:'suicide'})
	worker.close(()=>{
		process.exit(1)
	})
	setTimeout(()=>{
		process.exit(1)
	},5000)
})
```

---
#### Cluster
cluster和worker对象
```
//cluster对象
cluster.setttings:配置集群参数对象
cluster.isMaster:判断是不是master节点
cluster.isWorker:判断是不是worker节点
Event: 'fork': 监听创建worker进程事件
Event: 'online': 监听worker创建成功事件
Event: 'listening': 监听worker向master状态事件
Event: 'disconnect': 监听worker断线事件
Event: 'exit': 监听worker退出事件
Event: 'setup': 监听setupMaster事件
cluster.setupMaster([settings]): 设置集群参数
cluster.fork([env]): 创建worker进程
cluster.disconnect([callback]): 关闭worket进程
cluster.worker: 获得当前的worker对象
cluster.workers: 获得集群中所有存活的worker对象 cluster.workers[pid]

//worker对象
worker.id: 进程ID号
worker.process: ChildProcess对象
worker.suicide: 在disconnect()后，判断worker是否自杀
worker.send(message, [sendHandle]): master给worker发送消息。注：worker给发master发送消息要用process.send(message)
worker.kill([signal='SIGTERM']): 杀死指定的worker，别名destory()
worker.disconnect(): 断开worker连接，让他从集群中离开，让worker自杀,
Event: 'message': 监听master和worker的message事件
Event: 'online': 监听指定的worker创建成功事件
Event: 'listening': 监听master向worker状态事件
Event: 'disconnect': 监听worker断线事件
Event: 'exit': 监听worker退出事件
```

实例
```
var cluster = require('cluster')
var http = require('http')

//cluster.setupMaster({
//	exec:__dirname+'/child.js'
//})

if (cluster.isMaster){                    //cluster.isWorker === false
	console.log('[master] '+'start...')
	for (let i=0;i<3;i++){
		var worker = cluster.fork()			//投递当前进程
	}

	cluster.on('fork',(worker)=>{
		console.log('[master] '+'fork:worker '+worker.id)
	})

	cluster.on('online',(worker)=>{
		console.log('[master] '+'online:worker '+worker.id)
	})

	cluster.on('listening',(worker,address)=>{
		console.log('[master] '+'listening:worker '+worker.id+' pid '+worker.process.pid+' address '+address.address+':'+address.port)
	})	
 
	cluster.on('disconnect',(worker)=>{
		console.log('[master] '+'disconnect:worker '+worker.id)
	})

	cluster.on('exit',(worker)=>{
		console.log('[master] '+'exit:worker '+worker.id)	
	})

	for(let pid in cluster.workers){					//所有的workers
		cluster.workers[pid].on('message',(msg)=>{
			console.log('[master] '+'message '+msg)
		})
	}

	setInterval(()=>{
		for(let pid in cluster.workers){
			cluster.workers[pid].send('[master] '+'send msg to worker' + cluster.workers[pid].id)
		}		
	},10000)

}else if(cluster.isWorker){                                 //'NODE_UNIQUE_ID in process.env'
	var server = http.createServer((req,res)=>{
		res.writeHead(200,{'Content-Type':'text/plain'})
		res.end('worker:'+cluster.worker.id+ ' pid:'+process.pid + '\n')
	})

	server.listen(8111)

	process.on('message',(msg)=>{
		console.log('[worker] '+msg)
		process.send('[worker] worker'+cluster.worker.id+' received!')  //cluster.worker 子进程worker
	})
}
```

---
#### 跨域
```
app.all("*",function(req,res,next){
    //设置允许跨域的域名，*代表允许任意域名跨域
    res.header("Access-Control-Allow-Origin","*");
    //允许的header类型
    res.header("Access-Control-Allow-Headers","content-type");
    //跨域允许的请求方式 
    res.header("Access-Control-Allow-Methods","DELETE,PUT,POST,GET,OPTIONS");
    if (req.method.toLowerCase() == 'options') //当浏览器准备跨域时，他会给原服务器发送个options请求来验证需要访问的跨域网址
        res.send(200);  //让options尝试请求快速结束
    else
        next();
}
```

---
#### c++扩展
```
```