### 资料
```
https://mmoaay.gitbooks.io/boost-asio-cpp-network-programming-chinese
```

### error

```
#include "boost/system/error_code.hpp"

boost::system::error_code err;
:ip::address ip = ::ip::address::from_string("1.2.3.1", err);
cout << err.message() << endl;

try {
    :ip::address ip = ::ip::address::from_string("1.2.3.1", err);
}catch(boost::system::system_error err){
    cout << err.code() << endl;
}

同步错误码
所有的同步函数都有抛出异常或者返回错误码的重载，比如下面的代码片段：
sync_func( arg1, arg2 ... argN); // 抛出异常
boost::system::error_code ec;
sync_func( arg1 arg2, ..., argN, ec); // 返回错误码





```

### ip
```
// ip::address(v4_or_v6_address)
    这个函数把一个v4或者v6的地址转换成ip::address

// ip::address:from_string(str)
    这个函数根据一个IPv4地址（用.隔开的）或者一个IPv6地址（十六进制表示）创建一个地址。

// ip::address::to_string()
    这个函数返回这个地址的字符串。

// ip::address_v4::broadcast([addr, mask])
    这个函数创建了一个广播地址 ip::address_v4::any()：这个函数返回一个能表示任意地址的地址。

// ip::address_v4::loopback(), ip_address_v6::loopback()
    这个函数返回环路地址（为v4/v6协议）

// ip::host_name()
    这个函数用string数据类型返回当前的主机名


// 示例
    ::ip::address ip = ::ip::address.from_string("192.145.121.11");

    boost::system::error_code err;
    ::ip::address ip = ::ip::address.from_string("dasd", err);
    cout << err.message() << endl;

    const string ipString = ip.to_string();
```

### 端点
```
// 每个协议有自己的port
    ip::tcp::endpoint
    ip::udp::endpoint
    ip::icmp::endpoint

// 三种方法构建
    // endpoint()
        这是默认构造函数，某些时候可以用来创建UDP/ICMP socket。
    // endpoint(protocol, port)
        这个方法通常用来创建可以接受新连接的服务器端socket。
    // endpoint(addr, port)
        这个方法创建了一个连接到某个地址和端口的端点

// 示例
    ip::tcp::endpoint ep1;
    ip::tcp::endpoint ep2(ip::tcp::v4(), 80);
    ip::tcp::endpoint ep3( ip::address::from_string("127.0.0.1), 80);        

    //访问一个主机 解析
    io_service service;
    ip::tcp::resolver resolver(service);
    ip::tcp::resolver::query query("www.baidu.com", "https");
    ip::tcp::resolver::iterator iter = resolver.resolve(query);
    ip::tcp::resolver::iterator end;
    while (iter != end) {
        ip::tcp::endpoint ep = *iter++;
        cout << ep << endl;
        cout << ep.address().to_string() << ep.port() << endl;
    }

```

## 套接字
### 套接字连接函数
```
ip::tcp::socket, ip::tcp::acceptor, ip::tcp::endpoint,ip::tcp::resolver, ip::tcp::iostream
ip::udp::socket, ip::udp::endpoint, ip::udp::resolver
ip::icmp::socket, ip::icmp::endpoint, ip::icmp::resolver

相关函数

// assign(protocol,socket)
    这个函数分配了一个原生的socket给这个socket实例。当处理老（旧）程序时会使用它（也就是说，原生socket已经被建立了）

// open(protocol)
    这个函数用给定的IP协议（v4或者v6）打开一个socket。你主要在UDP/ICMP socket，或者服务端socket上使用。

// bind(endpoint)
    这个函数绑定到一个地址

// connect(endpoint)
    这个函数用同步的方式连接到一个地址

// async_connect(endpoint)
    这个函数用异步的方式连接到一个地址

// is_open()
    如果套接字已经打开，这个函数返回true

// close()
    这个函数用来关闭套接字。调用时这个套接字上任何的异步操作都会被立即关闭，同时返回error::operation_aborted错误码。

// shutdown(type_of_shutdown)
    这个函数立即使send或者receive操作失效，或者两者都失效。

// cancel()
    这个函数取消套接字上所有的异步操作。这个套接字上任何的异步操作都会立即结束，然后返回error::operation_aborted错误码。

// 示例
    ip::tcp::endpoint ep( ip::address::from_string("127.0.0.1"), 80);
    ip::tcp::socket sock(service);
    sock.open(ip::tcp::v4());
    sock.connect(ep);
    sock.write_some(buffer("GET /index.html\r\n"));
    char buff[1024]; sock.read_some(buffer(buff,1024));
    sock.shutdown(ip::tcp::socket::shutdown_receive);
    sock.close();
```

### 套接字I/O函数
```
// async_receive(buffer, [flags,] handler)
    这个函数启动从套接字异步接收数据的操作。

// async_read_some(buffer,handler)
    这个函数和async_receive(buffer, handler)功能一样。

// async_receive_from(buffer, endpoint[, flags], handler)
    这个函数启动从一个指定端点异步接收数据的操作。

// async_send(buffer [, flags], handler)
    这个函数启动了一个异步发送缓冲区数据的操作。

// async_write_some(buffer, handler)
    这个函数和async_send(buffer, handler)功能一致。

// async_send_to(buffer, endpoint, handler)
    这个函数启动了一个异步send缓冲区数据到指定端点的操作。

// receive(buffer [, flags])
    这个函数异步地从所给的缓冲区读取数据。在读完所有数据或者错误出现之前，这个函数都是阻塞的。

// read_some(buffer)
    这个函数的功能和receive(buffer)是一致的。

// receive_from(buffer, endpoint [, flags])
    这个函数异步地从一个指定的端点获取数据并写入到给定的缓冲区。在读完所有数据或者错误出现之前，这个函数都是阻塞的。

// send(buffer [, flags])
    这个函数同步地发送缓冲区的数据。在所有数据发送成功或者出现错误之前，这个函数都是阻塞的。

// write_some(buffer)
    这个函数和send(buffer)的功能一致。

// send_to(buffer, endpoint [, flags])
    这个函数同步地把缓冲区数据发送到一个指定的端点。在所有数据发送成功或者出现错误之前，这个函数都是阻塞的。

// available()
    这个函数返回有多少字节的数据可以无阻塞地进行同步读取。

// 示例
    // 在一个tcp套接字上进行同步读写
    io_service service;
    ip::tcp::endpoint ep(ip::address::from_string("192.1.1.1"), 80);
    ip::tcp::socket sock(service);
    sock.open(ip::tcp::v4());
    sock.connect(ep);
    sock.write_some(buffer("GET /index.html\r\n"));
    cout << sock.available() << endl;
    char buf[512];
    size_t read = sock.read_some(buffer(buf));

    // 在一个UDP套接字上进行同步读写
    ip::udp::endpoint ep(ip::address::from_string("192.1.1.1"), 80);
    ip::udp::socket sock(service);
    sock.open(ip::udp::v4());
    sock.send_to(buffer("testing\n"), receiver_ep);
    char buff[512];
    ip::udp::endpoint sender_ep;
    sock.receive_from(buffer(buff), sender_ep);

```

### 套接字控制
```
// get_io_service()
    这个函数返回构造函数中传入的io_service实例

// get_option(option)
    这个函数返回一个套接字的属性
    
// set_option(option)
    这个函数设置一个套接字的属性

// option
    broadcast	                如果为true，允许广播消息	bool
    debug	                    如果为true，启用套接字级别的调试	bool
    do_not_route	            如果为true，则阻止路由选择只使用本地接口	bool
    enable_connection_aborted	如果为true，记录在accept()时中断的连接	bool
    keep_alive	                如果为true，会发送心跳	bool
    linger	                    如果为true，套接字会在有未发送数据的情况下挂起close()	bool
    receive_buffer_size	        套接字接收缓冲区大小	int
    receive_low_watemark	    规定套接字输入处理的最小字节数	int
    reuse_address	            如果为true，套接字能绑定到一个已用的地址	bool
    send_buffer_size	        套接字发送缓冲区大小	int
    send_low_watermark	        规定套接字数据发送的最小字节数	int
    ip::v6_only	                如果为true，则只允许IPv6的连接	bool    

// io_control(cmd)
    这个函数在套接字上执行一个I/O指令



```
