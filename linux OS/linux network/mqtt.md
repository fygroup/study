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
    包含了简单的重发机制，Sender 发送消息之后等待接收者的 ACK，如果没收到 ACK 则重新发送消息
    这种模式能保证消息至少能到达一次，但无法保证消息重复

    发消息 -> 收到确认 -> 删除消息

QoS 2：消息仅传送一次(最高级别)
    发布者发布 QoS 为 2 的消息之后，会将发布的消息储存起来并等待接收者回复 PUBREC 的消息
    发送者收到 PUBREC 消息后，它就可以安全丢弃掉之前的发布消息，因为它已经知道接收者成功收到了消息
    发布者会保存 PUBREC 消息并应答一个 PUBREL，等待接收者回复 PUBCOMP 消息，当发送者收到 PUBCOMP 消息之后会清空之前所保存的状态。

当接收者接收到一条 QoS 为 2 的 PUBLISH 消息时，他会处理此消息并返回一条 PUBREC 进行应答。当接收者收到 PUBREL 消息之后，它会丢弃掉所有已保存的状态，并回复 PUBCOMP。

无论在传输过程中何时出现丢包，发送端都负责重发上一条消息。不管发送端是 Publisher 还是 Broker，都是如此。因此，接收端也需要对每一条命令消息都进行应答。
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
```

### MQTT控制报文
```
```

### mosquitto(MQTT代理)
```
// 
mosquitto
mosquitto-clients mosquitto客户端默认使用QoS 0

```