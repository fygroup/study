### redis的关键内容
```

Redis是k-v内存数据库，其特点
> Redis支持数据的持久化，可以将内存中的数据保持在磁盘中，重启的时候可以再次加载进行使用
> Redis支持的数据类型 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作
> Redis的所有操作都是原子性的，同时Redis还支持对几个操作全并后的原子性执行
> Redis支持master-slave模式的数据备份


数据结构
持久化
主从同步
redis集群搭建
redis分片
分布式锁 Redlock
redis + mysql 方案
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
    lindex key index 	读取下标为 index 的节点，返回节点字符串
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