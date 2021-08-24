### MQTT
```
[MQTT简单介绍] https://zhuanlan.zhihu.com/p/20888181
[一文带你简单了解MQTT协议] https://zhuanlan.zhihu.com/p/152195617
[MQTT协议中文版] https://mcxiaoke.gitbooks.io/mqtt-cn/content/
[MQTT比TCP协议好在哪儿] https://www.zhihu.com/question/23373904

1、概念
    MQTT是个协议，该协议是建立在TCP协议之上的，应用于物联网的通信协议
    MQTT是基于二进制消息的发布/订阅编程模式的消息协议
    由于规范很简单，非常适合需要低功耗和网络带宽有限的IoT场景(遥感数据、汽车、智能家居...)

2、特点
    (1) 使用发布/订阅消息模式，提供一对多的消息发布，解除应用程序耦合
    (2) 对负载内容屏蔽的消息传输
    (3) 使用 TCP/IP 提供网络连接
    (4) 有三种消息发布服务质量QoS(0,1,2)
    (5) 小型传输，开销很小(固定长度的头部是 2 字节)，协议交换最小化，以降低网络流量
    (6) 使用 Last Will 和 Testament 特性通知有关各方客户端异常中断的机制

3、MQTT协议的三种身份
    发布者(Publish)、代理(Broker)(服务器)、订阅者(Subscribe)
    消息的发布者和订阅者都是客户端，消息代理是服务器，消息发布者可以同时是订阅者

    (1) MQTT服务端(Broker)
        MQTT服务器以称为"消息代理"(Broker)，可以是一个应用程序或一台设备，位于消息发布者和订阅者之间
        > 接受来自客户的网络连接
        > 接受客户发布的应用信息
        > 处理来自客户端的订阅和退订请求
        > 向订阅的客户转发应用程序消息

    (2) MQTT客户端(Publish/Subscribe)
        一个使用MQTT协议的应用程序或者设备，它总是建立到服务器的网络连接
        > 发布其他客户端可能会订阅的信息
        > 订阅其它客户端发布的消息
        > 退订或删除应用程序的消息
        > 断开与服务器连接

4、MQTT传输的消息
    MQTT传输的消息分为 主题(Topic)和负载(payload)两部分
    (1) Topic
        可以理解为消息的类型，订阅者订阅(Subscribe)后，就会收到该主题的消息内容(payload)
    (2) payload
        可以理解为消息的内容，是指订阅者具体要使用的内容

　　当应用数据通过MQTT网络发送时，MQTT会把与之相关的服务质量(QoS)和主题名(Topic)相关连

5、关键字
    订阅(Subscription)
        订阅包含主题筛选器(Topic Filter)和最大服务质量(QoS)
        订阅会与一个会话(Session)关联
        一个会话可以包含多个订阅，每一个会话中的每个订阅都有一个不同的主题筛选器
    会话(Session)
        每个客户端与服务器建立连接后就是一个会话，客户端和服务器之间有状态交互
    主题名(Topic Name)
        连接到一个应用程序消息的标签，该标签与服务器的订阅相匹配
        服务器会将消息发送给订阅所匹配标签的每个客户端
    主题筛选器(Topic Filter)
        一个对主题名通配符筛选器，在订阅表达式中使用，表示订阅所匹配到的多个主题
    负载(Payload)
        消息订阅者所具体接收的内容

6、MQTT协议中的方法
    MQTT协议中定义了一些方法，来于表示对确定资源所进行操作
    Connect: 等待与服务器建立连接
    Disconnect: 等待MQTT客户端完成所做的工作，并与服务器断开TCP/IP会话
    Subscribe: 等待完成订阅
    UnSubscribe: 等待服务器取消客户端的一个或多个topics订阅
    Publish: MQTT客户端发送消息请求，发送完成后返回应用程序线程

7、MQTT消息类型
    MQTT拥有14种不同的消息类型
    CONNECT：客户端连接到MQTT代理
    CONNACK：连接确认
    PUBLISH：新发布消息
    PUBACK：新发布消息确认，是QoS 1给PUBLISH消息的回复
    PUBREC：QoS 2消息流的第一部分，表示消息发布已记录
    PUBREL：QoS 2消息流的第二部分，表示消息发布已释放
    PUBCOMP：QoS 2消息流的第三部分，表示消息发布完成
    SUBSCRIBE：客户端订阅某个主题
    SUBACK：对于SUBSCRIBE消息的确认
    UNSUBSCRIBE：客户端终止订阅的消息
    UNSUBACK：对于UNSUBSCRIBE消息的确认
    PINGREQ：心跳
    PINGRESP：确认心跳
    DISCONNECT：客户端终止连接前优雅地通知MQTT代理
```

### QoS
```
https://www.emqx.io/cn/blog/introduction-to-mqtt-qos

QoS 0：最多分发一次
    发布者只会发布一次消息，接收者不会应答消息，发布者也不会储存和重发消息
    消息在这个等级下具有最高的传输效率，但可能送达一次也可能根本没送达

QoS 1：至少分发一次
    包含了简单的重发机制，发送消息之后等待接收者的 PUBACK，如果没收到 PUBACK 则重新发送消息
    这种模式能保证消息至少能到达一次，但无法保证消息重复

    发消息 -> 收到确认 -> 删除消息

QoS 2：消息仅传送一次(最高级别)
    发布者发布 QoS 为 2 的消息之后，会将发布的消息储存起来并等待接收者回复 PUBREC 的消息
    发送者收到 PUBREC 消息后，它就可以安全丢弃掉之前的发布消息，因为它已经知道接收者成功收到了消息
    发布者会保存 PUBREC 消息并应答一个 PUBREL，等待接收者回复 PUBCOMP 消息，当发送者收到 PUBCOMP 消息之后会清空之前所保存的状态。

当接收者接收到一条 QoS 为 2 的 PUBLISH 消息时，他会处理此消息并返回一条 PUBREC 进行应答
当接收者收到 PUBREL 消息之后，它会丢弃掉所有已保存的状态，并回复 PUBCOMP

无论在传输过程中何时出现丢包，发送端都负责重发上一条消息，不管发送端是 Publisher 还是 Broker。因此，接收端也需要对每一条命令消息都进行应答

    Publisher                       Broker                          Subscriber
    store(Msg)
                PUBLISH(QoS2,Msg)->   
                                    store(Msg)
                                                PUBLISH(QoS2,Msg)->
                                                                    store(Msg)
                    <-PUBREC                           <-PUBREC
                    PUBREL->                           PUBREL->
                    <-PUBCOM                           <-PUBCOM
    delete(Msg)                    delete(Msg)                      delete(Msg)



    发消息 -> 收到确认 -> 发删除消息 -> 收到删除确认 -+ 
                                                   +---> 删除消息 
                        收到删除消息 ---------------+



// QoS降级
在 MQTT 协议中，从 Broker 到 Subscriber 这段消息传递的实际 QoS 等于
Publisher 发布消息时指定的 QoS 等级和 Subscriber 在订阅时与 Broker 协商的 QoS 等级，这两个 QoS 等级中的最小那一个
Actual Subscribe QoS = MIN(Publish QoS, Subscribe QoS)
```

### Retain
```
服务端需要保存PUBLISH RETAIN标志为1的报文，存储内容Message和QoS

一个Topic只能有一条Retained消息，新的Retained消息将覆盖老的Retained消息

订阅者使用通配符订阅主题，它会收到所有匹配的主题上的 Retained 消息

当服务端收到 Retain 标志为 1 的 PUBLISH 报文时(payload非空)，对于已存在的匹配订阅者，正常转发，并在转发前清除 Retain 标志

当新的订阅注册时，服务端就会查找是否存在匹配该订阅的保留消息，如果保留消息存在，就会立即转发给订阅者，且 Retain 标志必须保持为 1

MQTT v5.0 对于订阅建立时是否发送保留消息做了更细致的划分，并在订阅选项中提供了 Retain Handling 字段。例如某些客户端可能仅希望在首次订阅时接收保留消息，又或者不希望在订阅建立时接收保留消息，都可以通过 Retain Handling 选项调整。

保留消息虽然存储在服务端中，但它并不属于会话的一部分。也就是说，即便发布这个保留消息的会话终结，保留消息也不会被删除

删除保留消息只有两种方式：
客户端往某个主题发送一个 Payload 为空的保留消息，服务端就会删除这个主题下的保留消息
客户端设置该消息过期，过期后就会被删除

借助保留消息，新的订阅者能够立即获取最近的状态，而不需要等待无法预期的时间，这在很多场景下很非常重要的
```

### mosquitto(MQTT代理)
```
// 
mosquitto
mosquitto-clients mosquitto客户端默认使用QoS 0

```

### mosquitto SSL
```
// ca.crt server.crt server.key client.crt client.key
openssl genrsa -des3 -out ca.key 2048
openssl req -new -x509 -days 1826 -key ca.key -out ca.crt

openssl genrsa -out server.key 2048
openssl req -new -out server.csr -key server.key
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 360

openssl genrsa -out client.key 2048
openssl req -new -out client.csr -key client.key
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 360


// mosquitto.conf
port 8883
cafile /home/ubuntu/ca/ca.crt
certfile /home/ubuntu/ca/server.crt
keyfile /home/ubuntu/ca/server.key

// 部署
mosquitto -c /etc/mosquitto/mosquitto.conf -v
mosquitto_sub -h 192.168.1.181 -p 8883 -i 111  -t 111 –cafile ca.crt –cert client.crt –key client.key –insecure
mosquitto_pub -h 192.168.1.181 -p 8883 -t 111 -m "this is w show" –cafile ca.crt –cert client.crt –key client.key –insecure
```

### Packet Identifier(报文标识符字段)
```
很多控制报文的'可变报头'部分包含一个两字节的'报文标识符字段'

这些报文包含 PUBLISH(QoS > 0)，PUBACK，PUBREC，PUBREL，PUBCOMP，SUBSCRIBE, SUBACK，UNSUBSCRIBE，UNSUBACK

客户端每次发送一个新的这些类型的报文时都必须分配一个当前未使用的报文标识符

如果一个客户端要重发这个特殊的控制报文，在随后重发那个报文时，它必须使用相同的标识符

当客户端处理完这个报文对应的确认后，这个报文标识符就释放可重用
```

### MQTT CONNECT
```
CONNECT Control Packet

Byte Bits       Dec  Description
--------------------------------
// 固定报头
0    00010000   1^   Control Packet Type
1    00001100   12   Packet Length
// 可变报头
2    00000000   0    Protocol Name Length MSB
3    00000100   4    Protocol Name Length LSB
4    01001101   M*   Protocol Name 1st Byte
5    01010001   Q*   Protocol Name 2nd Byte
6    01010100   T*   Protocol Name 3rd Byte
7    01010100   T*   Protocol Name 4th Byte
8    00000100   4    Protocol Level(协议级别)
9    00000010   2    Connect Flags(连接标志)
10   00000000   0    Keep Alive MSB
11   00000000   0    Keep Alive LSB(保持连接)
// 有效载荷
12   00000000   0    Client Id Length MSB
13   00000000   0    Client Id Length LSB(标识符)
...

> 连接标志 Connect Flags
    0:  reserved
    1:  清理会话 Clean Session
    2:  遗嘱标志 Will Flag
    3-4:遗嘱QoS Will QoS
    5:  遗嘱保留 Will Retain
    6:  密码标志 Password Flag
    7:  用户名标志 User Name

> 保持连接 Keep Alive
    客户端传输完成一个控制报文的时刻到发送下一个报文的时刻，两者之间允许空闲的最大时间间隔
    客户端：在连接时间值内没有其他报文发送，客户端必须发送一个PINGREQ报文
    服务端：服务端在一点五倍的保持连接时间内没有收到客户端的控制报文，它必须断开客户端的网络连接

```

### MQTT CONNACK
```
Byte Bits       Dec  Description
--------------------------------
// 固定报头
0    00100000   2^   Control Packet Type
1    00000010   2    Packet Length
// 可变报头
2    00000000   0    Connect Acknowledge Flags(连接确认标志)
3    00000000   0    Connect Response Code(连接返回码)

> 连接确认标志
    0位表示Session Present标志，1-7位设置为0
    > session present
        如果服务端收到清理会话(clean session)标志为1，Session Present标志为0
        如果服务端收到清理会话(clean session)标志为0
            如果服务端保存了会话状态，Session Present标志为1
            如果服务端没保存会话状态，Session Present标志为0
        如果服务端发送了一个包含非零返回码的CONNACK报文，必须将Session Present标志为0

> 返回码
    值	返回码响应
    0	接受，连接已被服务端接受
    1	拒绝，不支持的协议版本	服务端不支持客户端请求的MQTT协议级别
    2	拒绝，不合格的客户端标识符	客户端标识符是正确的UTF-8编码，但服务端不允许使用
    3	拒绝，服务端不可用	网络连接已建立，但MQTT服务不可用
    4	拒绝，无效的用户名或密码	用户名或密码的数据格式无效
    5	拒绝，未授权	客户端未被授权连接到此服务器
    6-255		保留
```

### MQTT PUBLISH
```
Byte Bits       Dec  Description
--------------------------------
// 固定报头
0    00110000   3    Type, DUP, QoS, RETAIN(重发、服务质量(2位)、保留标识)
1    00000000   ?   Packet Length(剩余长度 = 可变报头的长度 + 有效载荷的长度)
// 可变报头
2    00000000   0    Topic Length MSB
3    00000011   3    Topic Length LSB (topic 长度)
4    01100001   a    Topic 1st Byte     
5    00101111   /    Topic 2nd Byte
6    01100010   b    Topic 3rd Byte
7    00000000   0    报文标识符 MSB
8    00001010   10   报文标识符 LSB
// 有效荷载

```

### MQTT PUBACK
```
Byte Bits       Dec  Description
--------------------------------
// 固定报头
0    01000000   4    报文类型、保留位
1    00000000   ?    Packet Length(剩余长度 = 可变报头的长度 + 有效载荷的长度)
// 可变报头
2    00000000   0    报文标识符 MSB
3    00001010   10   报文标识符 LSB
```

### MQTT PUBREC
```
QoS 2，第一步

Byte Bits       Dec  Description
--------------------------------
// 固定报头
0    01010000   5    报文类型、保留位
1    00000000   ?    Packet Length(剩余长度 = 可变报头的长度 + 有效载荷的长度)
// 可变报头
2    00000000   0    报文标识符 MSB
3    00001010   10   报文标识符 LSB 等待确认的PUBLISH报文的报文标识符
```

### MQTT PUBREL
```
QoS 2，第二步

Byte Bits       Dec  Description
--------------------------------
// 固定报头
0    11000010   6    报文类型、保留位
1    00000000   ?    Packet Length(剩余长度 = 可变报头的长度 + 有效载荷的长度)
// 可变报头
2    00000000   0    报文标识符 MSB
3    00001010   10   报文标识符 LSB 等待确认的PUBREC报文相同的报文标识符
```

### MQTT PUBCOMP
```
QoS 2，第三步

Byte Bits       Dec  Description
--------------------------------
// 固定报头
0    01110000   7    报文类型、保留位
1    00000000   ?    Packet Length(剩余长度 = 可变报头的长度 + 有效载荷的长度)
// 可变报头
2    00000000   0    报文标识符 MSB
3    00001010   10   报文标识符 LSB 等待确认的PUBREL报文相同的报文标识符
```

### MQTT SUBSCRIBE
```
Byte Bits       Dec  Description
--------------------------------
0    10000010   8^   Control Packet Type
1    00001010   10   Packet Length
2    00000000   0    Packet Identifier MSB
3    00000000   0    Packet Identifier LSB
4    00000000   0    Filter Length MSB
5    00001100   12   Filter Length LSB
6    01110100   t^   Filter 1st Byte
7    01100101   e^   Filter 2nd Byte
8    01110011   s^   Filter 3rd Byte
9    01110100   t^   Filter ...
10   00101111   /^          ...
11   01101101   m^          ...
12   01100101   e^          ...
13   01110011   s^          ...
14   01110011   s^          ...
15   01100001   a^          ...
16   01100111   g^          ...
17   01100101   e^          ...
18   00000000   0    Filter QoS 
```


### mqtt 大文件传输
```
```