### 资料
```
https://blog.csdn.net/u012503639/article/details/109100844
https://zhuanlan.zhihu.com/p/59132109
https://zhuanlan.zhihu.com/p/32221140
```

### CAN总线
```
CAN总线是一种串行通信协议，能有效地支持具有很高安全等级的分布实时控制
在汽车电子行业里，使用CAN 连接发动机的控制单元、传感器、防刹车系统等，传输速度可达1Mbps

CAN网络的消息是广播式的，即在同一时刻网络上所有节点侦测的数据是一致的，它是一种基于消息广播模式的串行通信总线

CAN I2C
I2C用于板内传输，CAN用于板间远距离传输


> 当CAN总线上的节点发送数据时，以报文形式广播给网络中的所有节点，总线上的所有节点都不使用节点地址等系统配置信息，只根据每组报文开头的 11 位标识符(CAN 2.0A 规范)解释数据的含义来决定是否接收。这种数据收发方式称为面向内容的编址方案
> 当某个节点要向其他节点发送数据时，这个节点的处理器将要发送的数据和自己的标识符传送给该节点的CAN总线接口控制器，并处于准备状态
> 当收到总线分配时，转为发送报文状态。数据根据协议组织成一定的报文格式后发出，此时网络上的其他节点处于接收状态
> 处于接收状态的每个节点对接收到的报文进行检测，判断这些报文是否是发给自己的以确定是否接收

```