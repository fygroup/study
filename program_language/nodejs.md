#### 镜像安装定义
`npm.cmd install -g cnpm --registry=https://registry.npm.taobao.org`

#### 安装包
`cnpm install express --save`

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

#### 灵活性
```
a={}
a.a=1
a=function(){}
a.a =1
b=new a()
b.a //undefined
```

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

#### hasOwnProperty
```
a={1:1,2:2}
a.hasOwnProperty(1)
function f(){
	this.a=1
}
f.prototype.b = 1
var b = new f()
b.hasOwnProperty('a') //true
b.hasOwnProperty('b') //false
```

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
```

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

#### instanceof
```
function A(){}
a=new A();
a instanceof A  //true
```

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
unpipe() pause(): 暂停从资源库读取数据，但 不会 暂停数据生成,pause暂停流的读入。
pipe() resume(): 正在从资源库中读取数据，监听 'data' 事件 ,恢复流的读入
```

* writeStream:
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

#### glob


#### async_hooks 异步钩子


#### console.log
这是异步操作！！！！！

#### child_process
```
const {spawn} = require('child_process');
const ls = spawn('ls',['-lh','./']);  //注意如果这块用ll，会报错
```

#### 事件观察者
**事件的执行先后**
`idle > IO > check`
```
idle: procss.nextTick(callback)  //事件保存在一个数组中，会将数组中的事件执行完，进行下一轮Tick
IO:
check: setTimeOut() setInterval() //事件保存在一个链表中，执行完当前一个，进行下一轮Tick
```

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

#### v8内存限制
```
--max-old-space-size 1000   //单位Mb 老生代
--max-new-space-size 1000   //单位Kb 新生代
```

#### stream 读取长度限制
`fs.createReadStream('txt',{highWaterMark:11})`

#### 传输层(TCP/UDP)
* TCP
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
套接字时stream,可以利用pipe()

```
