### redis
```
Redis是k-v内存数据库，其特点
> Redis支持数据的持久化，可以将内存中的数据保持在磁盘中，重启的时候可以再次加载进行使用
> Redis支持的数据类型 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作
> Redis的所有操作都是原子性的，同时Redis还支持对几个操作全并后的原子性执行
> Redis支持master-slave模式的数据备份

[深入redis系列]https://www.cnblogs.com/kismetv/p/9137897.html

数据结构
持久化
主从同步
redis集群搭建
redis分片
分布式锁 Redlock
redis + mysql 方案
    > 方案一
        读取：读取缓存，没有命中，则读取数据库
        更新：先更新数据库，在更新缓存
    > 方案二
        类似Page Cache，直接操作缓存，但是会带来一致性问题
        主要是何时同步数据是难点


// 内存淘汰和内存回收是一回事吗？
不是

// 内存淘汰的数据会持久化吗？
不同的淘汰策略，不一样

// redis自述
https://zhuanlan.zhihu.com/p/302405361?utm_source=wechat_session&utm_medium=social&utm_oi=555400798499188736&utm_campaign=shareopn
```

### redis连接
```
// 连接
redis-cli -h host -p port -a password

// 连接命令
AUTH password   // 验证密码是否正确
ECHO message    // 打印字符串
PING            // 查看服务是否运行
QUIT            // 关闭当前连接
SELECT index    // 切换到指定的数据库，index是索引，一个数字

```

### 数据结构
```
1、key(键)
    // 示例
    SET w3ckey redis    
    OK
    GET w3ckey 
    "redis"
    // 相关命令
    DEL key 该命令用于在 key 存在时删除 key
    DUMP key 序列化给定 key ，并返回被序列化的值
    EXISTS key 检查给定 key 是否存在
    EXPIRE key seconds 为给定 key 设置过期时间
    PERSIST key 移除 key 的过期时间，key 将持久保持
    TTL key 以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)
    RENAME key newkey 修改 key 的名称
    TYPE key 返回 key 所储存的值的类型

2、string(字符串)

3、hash(哈希)
    // 示例
    hmset key name1 value1 name2 value2
    hgetall key
    name1
    value1
    name2
    value2
    // 命令
    hdel key field1[field2......] 	删除 hash 结构中的字段 	可以进行多个字段的删除
    hexists key field 	判断 hash 结构中是否存在 field 字段
    hgetall key 	获取所有 hash 结构中的键值
    hincrby key field increment 	指定给 hash 结构中的某一字段加上一个整数 	要求该字段也是整数字符串
    hincrbyfloat key field increment 	指定给 hash 结构中的某一字段加上一个浮点数 	要求该字段是数字型字符串
    hkeys key 	返回 hash 中所有的键
    hlen key 	返回 hash 中键值对的数量
    hmget key field1[field2......] 	返回 hash 中指定的键的值，可以是多个 	依次返回值
    hmset key field1 value1 [field2 field2......] 	hash 结构设置多个键值对
    hset key field value 	在 hash 结构中设置键值对 	单个设值
    hsetnx key field value 	当 hash 结构中不存在对应的键，才设置值
    hvals key 	获取 hash 结构中所有的值

4、list(链表)
    Redis 链表是双向的
    // 命令
    lpush key node1 [node2.].....  	把节点 node1 加入到链表最左边。结果顺序node2,node1,...
    rpush key node1[node2]...... 	把节点 node1 加入到链表的最右边。结果顺序...,node1,node2
    lindex key index 	读取下标为 index   的节点，返回节点字符串
    llen key 	求链表的长度
    lpop key 	删除左边第一个节点，并将其返回
    rpop key 	删除右边第一个节点，并将其返回 	
    linsert key before|after pivot node 	插入一个节点 node，并且可以指定在值为pivot 的节点的before或after
    lrange list start end 	获取链表 list 从 start 下标到 end 下标的节点值 
    lrem list count value 	从左到右删除不大于count个等于value的节点，如果 count 为 0，则删除所有值等于 value的节点
    lset key index node 	设置列表下标为 index 的节点的值为 node 
    ltrim key start stop 	修剪链表，只保留从 start 到 stop 的区间的节点，其余的都删除掉

5、set(集合)
    // 命令
    sadd key member1 [member2 member3......] 	给键为 key 的集合増加成员
    scard key 	统计键为 key 的集合成员数
    sdiff key1 key2 	找出两个集合的差集
    sdiftstore des key1 key2 	先按 sdiff 命令的规则，找出 key1 和 key2 两 个集合的差集，然后将其保存到 des 集合中
    sinter key1 key2 	求 key1 和 key2 两个集合的交集
    sinterstore des key1 key2  	先按 sinter 命令的规则，找出 key1 和 key2 两个集合的交集，然后保存到 des 中
    sismember key member 	判断 member 是否键为 key 的集合的成员
    smembers key 	返回集合所有成员，如果数据量大，需要考虑迭代遍历的问题
    smove src des member  	将成员 member 从集合 src 迁移到集合 des 中
    spop key 	随机弹出集合的一个元素 	注意其随机性，因为集合是无序的
    srandmember key [count] 	随机返回集合中一个或者多个元素，count 为限制返回总数
    srem key member1[member2......] 	移除集合中的元素，可以是多个元素(对于很大的集合可以通过它删除部分元素，避免删除大量数据引发 Redis 停顿)
    sunion key1 [key2] 	求两个集合的并集 
    sunionstore des key1 key2 	先执行 sunion 命令求出并集，然后保存到键为 des 的集合中

6、有序set
```

### 消息队列
```
Redis Stream 主要用于消息队列(MQ)

Redis 本身是有一个 Redis 发布订阅 (pub/sub) 来实现消息队列的功能，但它有个缺点就是消息无法持久化，如果出现网络断开、Redis 宕机等，消息就会被丢弃

Redis Stream 提供了消息的持久化和主备复制功能

```

### 事务
```
redis中的事务保证
> 事务是一个单独的隔离操作
    事务中的所有命令都会序列化、按顺序地执行。事务在执行的过程中，不会被其他客户端发送来的命令请求所打断。
> 事务是一个原子操作
    事务中的命令要么全部被执行，要么全部都不执行

 MULTI 开始一个事务， 然后将多个命令入队到事务中， 最后由 EXEC 命令触发事务， 一并执行事务中的所有命令

// 事务命令
MULTI 标记一个事务块的开始
EXEC 执行所有事务块内的命令
DISCARD 事务会被放弃， 事务队列会被清空， 并且客户端会从事务状态中退出
UNWATCH 取消 WATCH 命令对所有 key 的监视。
WATCH key [key ...] 监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断

// 事务中的错误
(1) 产生错误
    > 事务在执行 EXEC 之前，入队的命令可能会出错(语法错误、内存不足等)
    > 命令可能在 EXEC 调用之后失败(处理了错误类型的键等)

(2) 处理方式
    1) EXEC之前
        redis 2.6.5之前，服务器会对命令入队失败的情况进行记录，并在客户端调用 EXEC 命令时，拒绝执行并自动放弃这个事务
        redis 2.6.5之后，只执行事务中那些入队成功的命令，而忽略那些入队失败的命令
    2) EXEC之后
        事务中有某些命令执行失败了， 事务队列中的其他命令仍然会继续执行，Redis 不会停止执行事务中的命令

(3) 为什么不支持回滚
    >  错误应该在开发的过程中被发现，而不应该出现在生产环境中
        Redis 命令只会因为错误的语法而失败(并且这些问题不能在入队时发现)，或是命令用在了错误类型的键上面
    >  简单且快速
        因为不需要对回滚进行支持，所以 Redis 的内部可以保持简单且快速


// watch
(1) check-and-set(CAS)
    WATCH 为 Redis 事务键的监控提供 CAS 行为
    如果有至少一个被监视的键在 EXEC 执行之前被修改了， 那么整个事务都会被取消， EXEC 返回nil-reply来表示事务已经失败

(2) watch
    > WATCH 使得 EXEC 命令需要有条件地执行
        事务只能在所有被监视键都没有被修改的前提下执行， 如果这个前提不能满足的话，事务就不会被执行
    > 监控多个键
        用户还可以在单个 WATCH 命令中监视任意多个键(WATCH key1 key2 key3)
    > 监控的取消
        WATCH 命令可以被调用多次。对键的监视从 WATCH 执行之后开始生效， 直到调用 EXEC 为止
        当客户端断开连接时， 该客户端对键的监视也会被取消

(3) 示例
    WATCH mykey
    val = GET mykey
    val = val + 1
    MULTI
    SET mykey $val
    EXEC
    如果在 WATCH 执行之后， EXEC 执行之前， 有其他客户端修改了 mykey 的值， 那么当前客户端的事务就会失败。 程序需要做的， 就是不断重试这个操作， 直到没有发生碰撞为止

```

### 发布订阅
```
// 一般模式
> 订阅
    // 客户端 1 订阅了一个叫作 chat 渠道的消息
    subscribe chat
> 发布
    // 客户端 2 向 渠道 chat 发送消息
    publish chat "let's go!!"
> 订阅多个
    subsrcibe foo bar

// 模式匹配订阅
    // 将接收所有发到news.art.figurative, news.music.jazz等等的消息，所有模式都是有效的，所以支持多通配符
    PSUBSCRIBE news.*
    // 取消订阅
    PUNSUBSCRIBE news.*
```

### 超时
```
设置key的过期时间，超过时间后，将会自动删除该key

// 命令
    persist key     持久化 key，取消超时时间
    ttl key 査看 key 的超时时间(秒) 没有超时时间为-1，已经超时则为-2
    expire key seconds  设置超时时间戳 (秒)
    expireat key timestamp  设置超时时间点(unix时间戳)

// 删除超时
    > set超时的key时，会删除原有的超时时间
    > incr lpush hset等只改变存储在key中的值而不用新key替换它的所有操作将使超时保持不变

// 刷新过期时间
    对已经有过期时间的key执行EXPIRE操作，将会更新它的过期时间
    场景：session

// Redis如何过期key
    相见内存回收

// 主从同步过期key
    master中的一个key过期时，DEL将会随着AOF文字一起合成到所有附加的slaves
    slaves不会主动过期key，会等到master执行DEL命令
```


### 持久化
```
// 为何要持久化
Redis是内存数据库，数据都是存储在内存中，为了避免进程退出导致数据的永久丢失，需要定期将Redis中的数据以某种形式(数据或命令)从内存保存到硬盘；当下次Redis重启时，利用持久化文件实现数据恢复。除此之外，为了进行灾难备份，可以将持久化文件拷贝到一个远程位置

// RDB持久化
RDB持久化方式能够在指定的时间间隔能对你的数据进行快照存储
(1) 配置
    redis.conf
    save <seconds> <changes>
    可以配置多个save，只要有一个满足都会触发bgsave，例如
    save 900 1      // 900秒之内至少一次写操作
    save 300 10     // 300秒之内至少发生10次写操作
    save 60 10000   // 60秒之内发生至少10000次写操作
(2) 触发
    1) 手动出发
        save：阻塞当前进程直到RDB结束
        bgsave：fork子进程完成RDB，Redis主进程阻塞时间只有fork阶段的那一下
    2) 自动触发
        > save自动触发
        > shutdown命令关闭服务器
            如果没有开启AOF持久化功能，那么会自动执行一次bgsave
        > 主从同步
            1> slave连接master
            2> master执行bgsave
            3> master向slave发送RDB
            4> slave删除旧数据，装载新数据
            5> 
(3) 优点
    > 适合备份
        RDB是一个非常紧凑的文件，它保存了某个时间点得数据集，非常适用于数据集的备份
        RDB是一个紧凑的单一文件，方便远端加密传送，非常适用于灾难恢复
    > 恢复速度更快
        与AOF相比在恢复大的数据集的时候，RDB方式会更快一些
    > fork
        fork子进程，可以充分利用系统的cow(写时复制)，主进程可以接受读写操作，而不用很大的开销
    
(4) 缺点
    > 无法实时持久化
        RDB根据save时间点(例如每隔5分钟并且对数据集有100个写的操作)来备份数据，两次save间宕机，则会丢失区间(分钟级)的增量数据，不适用于实时性要求较高的场景
    > fork
        fork子进程属于重量级操作，并且会阻塞redis主进程


// AOF持久化
AOF持久化方式记录每次对服务器'写的操作'，当服务器重启的时候会重新执行这些命令来恢复原始的数据
(1) 配置
    AOF默认是关闭的，通过redis.conf配置文件进行开启
    appendonly yes                  // 只有在yes下，aof才会生效  
    appendfilename appendonly.aof   // 指定aof文件名称
    appendfsync everysec            // 指定aof同步策略，always everysec no，默认为everysec
                                    // always：每一条AOF记录都立即同步到文件，性能很低，但较为安全
                                    // everysec：每秒同步一次，性能和安全都比较中庸的方式，也是redis推荐的方式
                                    // no：Redis永不直接调用文件同步，而是让操作系统来决定何时同步磁盘
    no-appendfsync-on-rewrite no    // aof-rewrite期间，appendfsync是否暂缓文件同步，no表示不暂缓，yes表示暂缓
    auto-aof-rewrite-min-size 64mb  // aof文件rewrite触发的最小文件尺寸
    auto-aof-rewrite-percentage 100 // 本次rewrite触发时aof文件应该增长的百分比
(2) 触发
    > 手动触发
        bgrewriteaof
    > 自动触发
        根据auto-aof-rewrite-min-size和auto-aof-rewrite-percentage参数确定自动触发时机
(3) AOF重写
    AOF日志会在持续运行中持续增大，需要定期进行AOF重写(对当前log的压缩和老旧log的删除)，对AOF日志进行瘦身
    > 机制
        > Redis执行fork()，现在同时拥有父进程和子进程
        > 子进程开始将新 AOF 文件的内容写入到临时文件
        > 对于所有新执行的写入命令，父进程一边将它们累积到一个内存缓存中，一边将这些改动追加到现有 AOF 文件的末尾，这样样即使在重写的中途发生停机，现有的 AOF 文件也还是安全的
        >当子进程完成重写工作时，它给父进程发送一个信号，父进程在接收到信号之后，将内存缓存中的所有数据追加到新 AOF 文件的末尾。
        > 新文件替换旧文件，之后所有命令都会直接追加到新 AOF 文件的末尾
(4) 优点
    > 更高的实时性
    > AOF只是追加写日志文件，对服务器性能影响较小，速度比RDB要快，消耗的内存较少
(5) 缺点
    > 日志文件太大需要文件瘦身
    > 恢复数据(重演命令式)比RDB要慢


// 如何选择持久化
Master：AOF 或 不持久化
Slave：RDB 或 RDB+AOF


// 提高重启效率
在 Redis 重启的时候，可以先加载 rdb 的内容，然后再重放增量 AOF 日志就可以完全替代之前的 AOF 全量文件重放，重启效率因此大幅得到提升
```

### 内存回收
```
// Redis的内存回收主要围绕以下两个方面
(1) Redis过期策略
    删除过期时间的key值
(2) Redis淘汰策略
    内存使用到达maxmemory上限时触发内存淘汰数据
Redis的过期策略和内存淘汰策略不是一件事

// Redis过期策略
(1) 定时过期
    每个设置过期时间的key都需要创建一个定时器，到过期时间就会立即清除
    可以立即清除过期的数据，对内存很友好；但是会占用大量的CPU资源去处理过期的数据
(2) 惰性过期
    只有当访问一个key时，才会判断该key是否已过期，过期则清除
    最大化地节省CPU资源，却对内存非常不友好。极端情况可能出现大量的过期key没有再次被访问，从而不会被清除，占用大量内存
(3) 定期过期
    每隔一定的时间，会扫描一定数量的数据库的expires字典中一定数量的key，并清除其中已过期的key
    该策略是前两者的一个折中方案，使得CPU和内存资源达到最优的平衡效果
    例如Redis每秒10次做的事情
    1) 测试随机的20个keys进行相关过期检测
    2) 删除所有已经过期的keys
    3) 如果有多于25%的keys过期，重复步奏1
Redis中同时使用了惰性过期和定期过期两种过期策略

// Redis淘汰策略
当内存使用达到maxmemory极限时，需要使用LAU淘汰算法来决定清理掉哪些数据，以保证新数据的存入
(1) 配置
    如果maxmemory设置为0，那么就默认不限制内存的使用
(2) 策略
    noeviction：当内存不足以容纳新写入数据时，新写入操作会报错。
    allkeys-lru：当内存不足以容纳新写入数据时，在键空间中，移除最近最少使用的key
    allkeys-random：当内存不足以容纳新写入数据时，在键空间中，随机移除某个key
    volatile-lru：当内存不足以容纳新写入数据时，在设置了过期时间的键空间中，移除最近最少使用的key
    volatile-random：当内存不足以容纳新写入数据时，在设置了过期时间的键空间中，随机移除某个key
    volatile-ttl：当内存不足以容纳新写入数据时，在设置了过期时间的键空间中，有更早过期时间的key优先移除
(3) 注意
    > 回收超时策略的缺点是必须指明超时的键值对，这会给程序开发带来一些设置超时的代码，增加开发者的工作量
    > 所有的键值对进行回收，有可能把正在使用的键值对删掉，增加了存储的不稳定性

```

### redis主从同步
```
// 基本概念
> 主服务器只负责写入数据，不负责让外部程序读取数据
> 存在多台从服务器，从服务器不写入数据，只负责同步主服务器的数据，并让外部程序读取数据
> 主服务器在写入数据后，即刻将写入数据的命令发送给从服务器，从而使得主从数据同步
> 应用程序可以随机读取某一台从服务器的数据，这样就分摊了读数据的压力
> 当从服务器不能工作的时候，整个系统将不受影响；当主服务器不能工作的时候，可以方便地从从服务器中选举一台来当主服务器

// 配置
> master
    bind 的配置，默认为 127.0.0.1，只允许本机访问，修改为 bind 0.0.0.0，允许其他服务器访问
> slave
    slaveof server port // server 代表主机，port 代表端口
    当从机 Redis 服务重启时，就会同步对应主机的数据了
    当不想让从机继续复制主机的数据时，可以在从机的 Redis 命令客户端发送 slaveof no one 命令

// 主从同步过程
slave: 向master发送sync连接命令
master: 执行bgsave生成RDB文件，并使用缓冲区记录当前的写命令
        向slave发送RDB文件         
slave: 丢弃所有当前数据，载入快照文件
        解析完后，开始正常接收命令
master: 向slave发送缓冲区写命令
slave: 执行master发过来的缓冲区命令
master: 缓冲区命令发送完后，主服务每执行完一次写命令，都会向slaves发送相同的写命令

```

### redis哨兵(保证高可用)
```
Redis Sentinel 是一个分布式系统，你可以在一个架构中运行多个 Sentinel 进程

这些进程使用gossip协议接收关于主服务器是否下线的信息，并使用投票协议(agreement protocols)来决定是否执行自动故障迁移，以及选择哪个从服务器作为新的主服务器

> 监控
    Sentinel 会不断地检查你的主服务器和从服务器是否运作正常
> 提醒
    当被监控的某个 Redis 服务器出现问题时， Sentinel 可以通过 API 向管理员或者其他应用程序发送通知
> 自动故障迁移
    master服务器失效时，Sentinel 会开始一次自动故障迁移操作，它会将失效主服务器的其中一个从服务器升级为新的主服务器，并让失效主服务器的其他从服务器改为复制新的主服务器
    当客户端试图连接失效的主服务器时，集群也会向客户端返回新主服务器的地址，使得集群可以使用新主服务器代替失效服务器

// 启动sentinel
redis-server /path/to/sentinel.conf --sentinel

// 配置sentinel
sentinel <选项的名字> <主服务器的名字> <选项的值>

> sentinel monitor mymaster 127.0.0.1 6379 2
    监视名为mymaster的主服务器，IP为127.0.0.1，端口为6379，而将这个主服务器判断为失效至少需要2个Sentinel同意(实际上要大多数)
> sentinel down-after-milliseconds mymaster 60000
    如果服务器在给定的毫秒数之内，没有返回Sentinel发送的PING命令的回复，或者返回一个错误
> sentinel failover-timeout mymaster 180000
> sentinel parallel-syncs mymaster 1
    选项指定了在执行故障转移时，最多可以有多少个从服务器同时对新的主服务器进行同步，这个数字越小，完成故障转移所需的时间就越长

// 自动发现
一个 Sentinel 可以与其他多个 Sentinel 进行连接， 各个 Sentinel 之间可以互相检查对方的可用性， 并进行信息交换
> 无须为每个Sentinel分别设置其他Sentinel的地址
    因为 Sentinel 可以通过发布与订阅功能来自动发现正在监视相同主服务器的其他Sentinel
> 无需列出主服务器下的所有slave
    因为 Sentinel 可以通过询问主服务器来获得所有从服务器的信息

// API
(1) 向Sentinel发送命令
    SENTINEL masters: 列出所有被监视的主服务器，以及这些主服务器的当前状态
    SENTINEL slaves: 列出给定主服务器的所有从服务器，以及这些从服务器的当前状态
    SENTINEL get-master-addr-by-name: 返回给定名字的主服务器的 IP 地址和端口号
    SENTINEL reset: 重置所有名字和给定模式 pattern 相匹配的主服务器
    SENTINEL failover: 当主服务器失效时，在不询问其他 Sentinel 意见的情况下，强制开始一次自动故障迁移 
(2) 发布订阅

```

### 分布式锁的注意点
```
https://zhuanlan.zhihu.com/p/87498360
https://juejin.im/post/6844903830442737671
https://juejin.im/post/6854573212831842311

1、锁超时
    (1) 问题描述
        客户端1获取锁成功
        客户端1在某个操作上阻塞了太长时间，设置的key过期了，锁自动释放了
        客户端2获取到了对应同一个资源的锁
        客户端1从阻塞中恢复过来，因为value值一样，所以执行释放锁操作时就会释放掉客户端2持有的锁，这样就会造成问题
    (2) 解决方式
        1) 方法1
            加锁时设置唯一性value，在释放锁时需要对value进行验证
        2) 方法2
            线程在首次加锁成功后，设置count变量和value变量，count变量的意义就是重入次数，value是UUID，表示锁的唯一标识
            线程加锁成功后，自增count变量，释放锁只需自减count变量，直到为0时才真正释放锁
            注意为了避免释放不属于自己的锁，需要先对value进行比对，在进行count的加减

2、分布式中master失效导致的锁丢失
    (1) 问题描述
        客户端A从master获取到锁
        在master将锁同步到slave之前，master宕掉了
        slave节点被晋级为master节点
        客户端B取得了同一个资源被客户端A已经获取到的另外一个锁。安全失效！
    (2) Redlock方法
        
    (3) token fetch

```

### redis分布式锁
```
// 单例锁
(1) 获得锁
    SET resource_name my_random_value NX PX 30000
    NX: 不存在key的时候才能被执行成功
    PX: 这个key有一个30秒的自动失效时间
    my_random_value: 所有的客户端必须是唯一的(类似UUID)，所有同一key的获取者(竞争者)这个值都不能一样
(2) 释放锁(一定要比较value，防止误解锁)
    if redis.call("get", KEYS[1]) == ARGV[1] then
        return redis.call("del", KEYS[1])
    else 
        return 0
    end

// 分布式锁(redlock)
    分布式中存在master节点失效导致的主从切换，那么就会出现锁丢失的情况(见上述)
    假设有5个Redis master节点，同时我们需要在5台服务器上面运行这些Redis实例，这样保证他们不会同时都宕掉
    (1) 关键字
        开始上锁的UNIX时间
        锁有效时间: 设置(10s)
        连接超时时间: 设置(5-50ms)
        获得锁用的时间: 客户端当前时间 - 开始上锁的UNIX时间
        锁的真正有效时间: 锁有效时间 - 获得锁用的时间
    (2) 过程
        > 获取当前Unix时间，以毫秒为单位
        > 依次尝试从5个实例，使用相同的key和具有唯一性的value(UUID)获取锁
            保证 锁有效时间 > 客户端连接服务端超时时间
            避免服务器端Redis已经挂掉的情况下，客户端还在死死地等待响应结果
            如果服务器端没有在规定时间内响应，客户端应该尽快尝试去另外一个Redis实例请求获取锁
        > 当且仅当从大多数(N/2+1，3个节点)的Redis节点都取到锁，并且 获得锁用的时间 < 锁有效时间时，锁才算获取成功
        > 获得锁后，得到 锁的真正有效时间
            如果获取锁失败，客户端应该在所有的Redis实例上进行解锁
            即便某些Redis实例根本就没有加锁成功，防止某些节点获取到锁但是客户端没有得到响应而导致接下来的一段时间不能被重新获取锁

```


### redis分布式(集群)
```
(1) 数据分区
    将数据根据哈希值分布到多个redis实例
    散列分区方法: 例如 一致性哈希 等(详见)

(2) 分区方案
    1) 客户端分区
        在客户端内就已经决定数据会被定位到哪个redis节点
    2) 代理分区
        客户端将请求发送给代理，然后代理决定去定位节点
        redis的一种代理实现就是Twemproxy
    3) 查询路由
        客户端随机地请求任意一个redis实例，然后由Redis将请求转发给正确的Redis节点
        Redis Cluster实现了一种混合形式的查询路由，但并不是直接将请求从一个redis节点转发到另一个redis节点，而是在客户端的帮助下直接redirected到正确的redis节点

(3) 分区注意点
    > 多个key的操作无法支持，例如不同redis实例不能对两个集合求交集
    > 不能使用redis事务
    > 不同实例需要主从备份
    > 分区时动态扩容或缩容可能非常复杂


```

### redigo
```
1、连接池
    用redigo自带的池来管理连接
    不然的话，每当要操作redis时，建立连接，用完后再关闭，会导致大量的连接处于TIME_WAIT状态

    var (
        Pool *redis.Pool
        REDIS_HOST  string
        REDIS_INDEX int
    )

    // 建立连接池
    Pool = &redis.Pool{
        MaxIdle: 1,
        MaxAxtive: 10,
        IdleTimeout: 180 * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", REDIS_HOST)
            if err != nil {
                return nil, err
            }
            c.Do("select", REDIS_INDEX)
            return c, nil
        }
    }
    // 获得连接池
    conn := Pool.Get()
    defer conn.Close()
    test, _ := conn.Do("set", "aaa")


    // MaxIdle      最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
    // MaxActive    最大的激活连接数，表示同时最多有N个连接
    // IdleTimeout  最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
    // Dial         建立连接

```