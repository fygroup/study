### 初始化
```
mysqld      服务执行工具    
mysql       客户端
mysqladmin  运维和管理工具

https://blog.csdn.net/github_39533414/article/details/80144890 新版本

// 创建mysql组和用户
groupadd mysql
useradd -g mysql mysql

// 创建目录
cd /usr/local/mysql
mkdir ./data
chown -R mysql:mysql ./

// 配置my.cnf
https://www.cnblogs.com/langdashu/p/5889352.html
vi /etc/my.cnf

// 启动服务
mysqld --initialize --user=mysql --basedir=/usr/local/mysql --datadir=/usr/local/mysql/data
```

### 配置全局环境变量
```
vi /etc/profile
在 profile 文件底部添加如下两行配置，保存后退出
export PATH=$PATH:/usr/local/mysql/bin:/usr/local/mysql/lib
source /etc/profile
```

### mysqld和mysqld_safe
```
mysqld_safe 作为mysqld 启动脚本，开启了守护mysqld进程的任务

mysqld_safe相当于mysqld的守护进程，当mysqld死了，mysqld_safe会把它拉起

mysqld_safe启动能够为mysqld分配系统资源
```

### 建立和启动service服务
```
// 建立service服务
cp mysql.server /etc/init.d/mysqld
chmod +x /etc/init.d/mysqld 
chkconfig --add mysqld      添加到系统服务
chkconfig  --list mysqld   检查服务是否生效  

// 启动service服务
service mysqld start   

会启动mysqld_safe（root）和mysqld(mysql)
```

### 连接数据库
```
// 连接
mysql -u root -p    //这里默认的是/usr/local/mysql/my.cnf
mysql --defaults-file=my.cnf -u malx -p
mysql --defaults-file=my.cnf -u root -p  //root 超级权限用户
mysql -h 192.168.5.116 -P 3306 -u root -p123456
mysql -S .sock -u root -p

// 无需登录，打开服务
mysqld_safe --defaults-file=my.cnf --user=mysql --skip-grant-tables
```

### 权限设置
```
// 新建用户
create user 'malx'@'%' IDENTIFIED by '111'  // %代表可以从外网连接
添加完用户后，别忘了赋予权限

// 删除用户
use mysql
drop user root@'%'          // %表示所有地址，但是不包括localhost
drop user root@localhost

// 给与用户权限
GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' IDENTIFIED BY '123456'  // 只允许本地root登陆
flush privileges

// flush privileges 
The MySQL server is running with the --skip-grant-tables option so it cannot execute this statement

// 修改密码
set password for root@localhost = password('123456')
或者
格式：mysqladmin -u用户名 -p旧密码 password 新密码
```

### 数据库基本操作
```
//创建数据库

//创建数据表

// 查询
select * from cnv_reportresult t where enable=-1 and sampleid in ("CL170407","CL170411");

// 严格模式
[mysqld]
sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES
这是就会限制你的长度

// 删除，添加或修改表字段
ALTER TABLE testalter_tbl DROP i        //删除"i"字段
ALTER TABLE testalter_tbl ADD i INT     //i 字段会自动添加到数据表字段的末尾

// 查看执行sql的信息
explain select * from table where key=value
    id: 1
    select_type: SIMPLE
    table: tab_with_index
    type: ALL
    possible_keys: name
    key: NULL
    key_len: NULL
    ref: NULL
    rows: 4
    Extra: Using where
    1 row in set (0.00 sec)

// INSERT .... SELECT ....
复制"table2"中的数据插入到"table1"中
INSERT INTO table1 (key1, key2) SELECT key1, key2 FROM table2

// CREATE TABLE ... SELECT ...

```

### 用户管理和权限设置
```
https://www.cnblogs.com/fslnet/p/3143344.html

my.cnf中的user=mysql，表示启动用户是mysql

初始化mysql，用的是root用户，他会给你一个root的初始登陆密码

//修改密码，新版本（话说新版本的好繁琐）
alter user 'root'@'localhost' identified with mysql_native_password by '123456';  
flush privileges;  //刷新

//创建账号密码
CREATE USER `wangwei`@`127.0.0.1` IDENTIFIED BY 'passowrd';

//授予权限
GRANT ALL ON *.* TO `wangwei`@`127.0.0.1` WITH GRANT OPTION;

//删除权限
REVOKE all privileges ON databasename.tablename FROM 'username'@'host';

//修改密码
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '你的密码';  </pre>

//MySQL8.0中带过期时间用户的创建：
CREATE USER `wangwei`@`127.0.0.1` IDENTIFIED BY 'wangwei' PASSWORD EXPIRE INTERVAL 90 DAY;
GRANT ALL ON *.* TO `wangwei`@`127.0.0.1` WITH GRANT OPTION

```

### MyISAM表锁
```
MyISAM存储引擎只支持表锁,开销小，加锁快
不会出现死锁
锁定粒度大，发生锁冲突的概率最高,并发度最低

(1) 查询表级锁争用情况
    show status like 'table%';
        +----------------------------+---------+
        | Variable_name              | Value   |
        +----------------------------+---------+
        | Table_locks_immediate      | 100     |    //产生表级锁定的次数
        | Table_locks_waited         | 11      |    //出现表级锁定争用而发生等待的次数
        +----------------------------+---------+
    如果Table_locks_waited的值比较高，则说明存在着较严重的表级锁争用情况。

(2) 表级锁的锁模式
    表共享读锁（Table Read Lock）和表独占写锁（Table Write Lock）

(3) 加表锁
    MyISAM在执行查询语句（SELECT）前，会自动给涉及的所有表加读锁，在执行更新操作（UPDATE、DELETE、INSERT等）前，会自动给涉及的表加写锁，这个过程并不需要用户干预
    1) 显示加表锁 lock table film_text write; 释放表锁：unlock tables;
    2) Lock tables orders read local, order_detail read local; //给read和order_detail加读锁，表后加的local其作用就是在满足MyISAM表并发插入条件的情况下，允许其他用户在表尾并发插入记录
    3) 在用LOCK TABLES给表显式加表锁时，必须同时取得所有涉及到表的锁，并且MySQL不支持锁升级。在执行LOCK TABLES后，只能访问显式加锁的这些表，不能访问未加锁的表
    4) 如果加的是读锁，那么只能执行查询操作，而不能执行更新操作。

(4) 并发插入
    在一定条件下，MyISAM表也支持查询和插入操作的并发进行。
    例如：一个session使用LOCK TABLE命令给表film_text加了读锁，这个session可以查询锁定表中的记录，但更新或访问其他表都会提示错误；同时，另外一个session可以查询表中的记录，但更新就会出现锁等待。
    1) 插入并发性,重要参数concurrent_insert
      concurrent_insert=2，无论MyISAM表中有没有空洞，都允许在表尾并发插入记录；
      concurrent_insert=1，如果MyISAM表中没有空洞（即表的中间没有被删除的行），MyISAM允许在一个进程读表的同时，另一个进程从表尾插入记录。这也是MySQL的默认设置；
      concurrent_insert=0，不允许并发插入。
    2) 将concurrent_insert系统变量设为2，总是允许并发插入；同时，通过定期在系统空闲时段执行OPTIMIZE TABLE语句来整理空间碎片，收回因删除记录而产生的中间空洞

(5) 锁调度
    一个进程请求某个 MyISAM表的读锁，同时另一个进程也请求同一表的写锁，MySQL如何处理呢？答案是写进程先获得锁。
    > 不太适合于有大量更新操作和查询操作应用的原因，因为，大量的更新操作会造成查询操作很难获得读锁，从而可能永远阻塞。这种情况有时可能会变得非常糟糕
    > 可以通过一些设置来调节MyISAM 的调度行为。具体见网址

(6) 表锁的优化
    优化MyISAM存储引擎锁定问题，关键的就是如何让其提高并发度。让锁定的时间变短，然后就是让可能并发进行的操作尽可能的并发。
    1) 查询表级锁争用情况
       （见上述）
    2) 缩短锁定时间
       让Query执行时间尽可能的短
       a) 尽两减少大的复杂Query，将复杂Query分拆成几个小的Query分布进行；
       b) 尽可能的建立足够高效的索引，让数据检索更迅速；
       c) 尽量让MyISAM存储引擎的表只存放必要的信息，控制字段类型；
       d) 利用合适的机会优化MyISAM表数据文件。
    3) 分离能并行的操作
       插入并发性,利用好重要参数concurrent_insert
    4) 合理利用读写优先级
       MyISAM存储引擎的是读写互相阻塞的，具体情况具体定
```

### InnoDB锁
```
InnoDB与MyISAM的最大不同有两点：一是支持事务，二是采用了行级锁

(1) 行锁
    1) 操作
        获取InnoDB行锁争用情况
        show status like 'innodb_row_lock%';
            +-------------------------------+-------+
            | Variable_name                 | Value |
            +-------------------------------+-------+
            | InnoDB_row_lock_current_waits | 0     |
            | InnoDB_row_lock_time          | 0     |
            | InnoDB_row_lock_time_avg      | 0     |
            | InnoDB_row_lock_time_max      | 0     |
            | InnoDB_row_lock_waits         | 0     |
            +-------------------------------+-------+
        观察发生锁冲突的表、数据行
        Show innodb status         //会不断显示数据表的连接 状态
            *************************** 1. row ***************************
            ....
            ------------
            TRANSACTIONS
            ------------
            ...
        暂停查看
        DROP TABLE innodb_monitor
    2) 行锁模式
        共享锁(S): 读锁，允许一个事务去读一行，阻止其他事务获得相同数据集的排他锁
        排他锁(X): 写锁，允许获得排他锁的事务更新数据，阻止其他事务取得相同数据集的共享读锁和排他写锁
        // 允许行锁和表锁共存，实现多粒度锁机制
        意向共享锁(IS): 事务打算给数据行加行共享锁，事务在给一个数据行加共享锁前必须先取得该表的 IS 锁。
        意向排他锁(IX): 事务打算给数据行加行排他锁，事务在给一个数据行加排他锁前必须先取得该表的 IX 锁。
        // 如果一个事务请求的锁模式与当前的锁兼容， InnoDB 就将请求的锁授予该事务；反之，如果两者不兼容，该事务就要等待锁释放
        兼容性  IS	IX	 S	  X
        IS	   兼容	兼容 兼容 互斥
        IX	   兼容	兼容 互斥 互斥
        S      兼容	互斥 兼容 互斥
        X	   互斥	互斥 互斥 互斥
    3) 加锁方式
        意向锁是 InnoDB 自动加的，不需用户干预
        对于 UPDATE、DELETE 和 INSERT 语句，InnoDB会自动给涉及数据集加排他锁X
        对于普通 SELECT 语句，InnoDB 不会加任何锁, 手动加锁如下
        > 共享锁
            SELECT ... FROM table_name WHERE ... LOCK IN SHARE MODE
        > 排它锁
            select ... for update
            排他锁的申请前提：没有线程对该结果集中的任何行数据使用排他锁或共享锁，否则申请会阻塞
            for update仅适用于InnoDB，且必须在事务(BEGIN/COMMIT)中才能生效
            在进行事务操作时，通过for update语句，MySQL会对查询结果集中每行数据都添加排他锁，其他线程对该记录的更新与删除操作都会阻塞

    4) 行锁的特点 
        > 开销
            开销大，加锁慢
            会出现死锁
            锁定粒度最小，发生锁冲突的概率最低，并发度也最高
            使用行级锁定的主要是InnoDB存储引擎

        > 索引加锁
            InnoDB行锁是通过给索引项加锁来实现的，这一点MySQL与Oracle不同，后者是通过在数据块中对相应数据行加锁来实现的
            InnoDB这种行锁实现特点意味着只有通过索引条件检索数据，InnoDB才使用行级锁，否则，InnoDB将使用表锁！
            
            有时即便使用了索引字段，如果MySQL认为全表扫描效率更高，比如对一些很小的表，它就不会使用索引，这种情况下InnoDB将使用表锁，而不是行锁
        
        > 不同类型比较
            索引字段类型(varchar)，如果where条件中不是和varchar类型进行比较(比如int)，则会对name进行类型转换，而执行的全表扫描


(2) gap lock(间隙锁)与next-key lock
    间隙锁，锁定一个范围，不包括记录本身
    next-key lock，锁定一个范围，包含记录本身
    1) 特点
        > innodb对于行的查询使用next-key lock
        > next-key lock为了解决Phantom Problem幻读问题
        > 当查询的索引含有唯一属性时，将next-key lock降级为行锁
        > Gap lock设计的目的是为了阻止多个事务将记录插入到同一范围内，而这会导致幻读问题的产生
        > 有两种方式显式关闭gap锁: 将事务隔离级别设置为RC; 将参数innodb_locks_unsafe_for_binlog设置为1
    2) 加锁方式
        > 进行范围条件而不是相等条件检索数据，并请求共享或排他锁时，会给符合条件的已有数据记录的索引项加锁；对于键值在条件范围内但并不存在的记录（这样做可以防止幻读）
        > 这种加锁机制会阻塞符合条件范围内键值的并发插入，影响性能
        > 如果使用相等条件请求给一个不存在的记录加锁，InnoDB也会使用间隙锁！！！
        例如：
            session1                                                        session2
            //对不存在的值加间隙锁                                           //这时插入新的value会阻塞
            select * from table where key = value(不存在) for update;       insert into table (key,...) values(value,...);
            //回滚                                                          //接触阻塞，插入成功
            rollback                                                        Query OK, 1 row affected (13.35 sec)


```


### mysql事务
```
https://yq.aliyun.com/articles/513493
https://www.cnblogs.com/huanongying/p/7021555.html

事务的操作
    开启事务：Start Transaction
    事务结束：End Transaction
    提交事务：Commit Transaction
    回滚事务：Rollback Transaction

mysql示例
    BEGIN;      //开始一个事务
    xxxx
    如果出错：
        ROLLBACK    //事务回滚
    xxxx
    COMMIT      //事务确认

// MYSQL的事务处理主要有两种方法。
(1) 用begin,rollback,commit来实现
    begin 开始一个事务
    rollback 事务回滚
    commit  事务确认
(2) 直接用set来改变mysql的自动提交模式
    MYSQL默认是自动提交的，也就是你提交一个QUERY，它就直接执行！我们可以通过
    set autocommit=0   禁止自动提交
    set autocommit=1   开启自动提交
    来实现事务的处理。

// 但注意当你用 set autocommit=0 的时候，你以后所有的SQL都将做为事务处理，直到你用commit确认或rollback结束，注意当你结束这个事务的同时也开启了个新的事务！按第一种方法只将当前的作为一个事务！推荐使用第一种方法！

// MYSQL中只有INNODB和BDB类型的数据表才能支持事务处理！其他的类型是不支持的！（切记！）

// 开启事务：start transaction；结束事务：commit或rollback。

// 默认隔离级别为Repeatable read，可以通过下面语句查看：
    select @@tx_isolation

// 设置隔离级别
    set transaction isolation level read committed

```

### 隔离级别操作
```
1、查看当前会话隔离级别
select @@tx_isolation;

2、查看系统当前隔离级别
select @@global.tx_isolation;

3、设置当前会话隔离级别
set session transaction isolatin level repeatable read;

4、设置系统当前隔离级别
set global transaction isolation level repeatable read;

5、开始事务
set autocommit=off 
start transaction
```


### 最大连接数
```
一个事件表示一个链接

不用事件的一条sql表示一个链接
```

### limit
```
select * from 表名 order by 排序字段 limt M,N

//这种写法存在陷阱，在排序字段有数据重复的情况下，会很容易出现排序结果与预期不一致的问题
//新版本的mysql会找到排序的row_count行后立马返回，而不是排序整个查询结果再返回
//所以，排序条件必须是索引排序

select * from 表名 order by order by id limt M,N

```

### undo redo
```
https://yq.aliyun.com/articles/592937

数据库通常借助日志来实现事务，常见的有undo log、redo log，undo/redo log都能保证事务特性，这里主要是原子性和持久性，即事务相关的操作，要么全做，要么不做，并且修改的数据能得到持久化

1、undo log
    Undo Log是为了实现事务的原子性，在MySQL数据库InnoDB存储引擎中
    undo log可以把所有没有COMMIT的事务回滚到事务开始前的状态
    (1) 原理
        为了满足事务的原子性，在操作任何数据之前，首先将数据备份到一个地方(Undo Log)。然后进行数据的修改。如果出现了错误或者用户执行了ROLLBACK语句，系统可以利用Undo Log中的备份将数据恢复到事务开始之前的状态
    (2) 过程
        假设要修改数据库的A=1，B=2两个数据
        1) 事务开始
        2) 从磁盘读取A到内存中
        3) 记录A=1到undo log中
        4) 修改A=11
        5) 从磁盘读取B到内存中
        6) 记录B=2到undo log中
        7) 修改B=22
        8) 将undo log写到磁盘
        9) 将修改后的数据写到磁盘
        10) 事务提交

        为了保证原子性和持久性，事务具有以下特点
        1) 更新数据前记录undo log
        2) undo log必须先于数据写回磁盘
        所以
        1) 事务提交必定数据持久化
        2) 如果上述过程2-8之间系统崩溃，内存中的数据丢失，但是磁盘中的数据保持不变
        3) 如果上述过程8-9之间系统崩溃，可以利用磁盘中的undo log回滚事务，还原数据
    (3) 缺陷
        1) 每个事务提交前将数据和Undo Log写入磁盘，这样会导致大量的磁盘IO，性能很低
        2) 如果能够将数据缓存一段时间，就能减少IO提高性能，但是这样就会丧失事务的持久性
        因此引入了另外一种机制来实现持久化，即Redo Log

2、redo log
    记录的是数据页的物理变化
    (1) 原理
        和Undo Log相反，Redo Log记录的是新数据的备份
        在事务提交前，只要将Redo Log持久化即可，不需要将数据持久化
        当系统崩溃时，虽然数据没有持久化，但是Redo Log已经持久化。系统可以根据Redo Log的内容，将所有数据恢复到最新的状态
    (2) undo+redo 过程
        假设要修改数据库的A=1，B=2两个数据
        1) 事务开始
        2) 记录A=1到undo log      --+
        3) 修改A=3                  |
        4) 记录A=3到redo log        |   内存中
        5) 记录B=2到undo log        |       
        6) 修改B=4                  |
        7) 记录B=4到redo log      --+  
        8) 将redo log写入磁盘
        9) 事务提交
        特点
        1) 必须在事务提交前将Redo Log持久化
        2) 数据不需要在事务提交前写入磁盘，而是缓存在内存中
        3) Redo Log保证事务的持久性
        4) Undo Log保证事务的原子性
    (3) 性能
        1) Undo+Redo的设计主要考虑的是提升IO性能。通过缓存数据，减少了写数据的IO
        2) Redo Log存储在一段连续的空间上。以顺序追加的方式记录Redo Log，通过顺序IO来改善性能
        3) 并发的事务共享Redo Log的存储空间，所以当一个事务将Redo Log写入磁盘时，也会将其他未提交的事务的日志写入磁盘

3、恢复
    进行恢复时，重做所有事务包括未提交的事务和回滚了的事务。然后通过Undo Log回滚那些未提交的事务
    实际上InnoDB将Undo Log看作数据，因此记录Undo Log的操作也会记录到Redo Log中。这样undo log就可以象数据一样缓存起来，而不用在redo log之前写入磁盘了
```

### 存储引擎选择
```
如果没有特别的需求，使用默认的Innodb即可

MyISAM：以读写插入为主的应用程序，比如博客系统、新闻门户网站

Innodb：更新（删除）操作频率也高，或者要保证数据的完整性；并发量高，支持事务和外键。比如OA自动化办公系统
```
