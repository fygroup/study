mysqld is the server executable (one of them)   #服务执行工具    
mysql is the command line client  # 客户端工具   查询用
mysqladmin is a maintainance or administrative utility  # 运维和管理工具
//初始化-----（以下都是在root权限下）
https://blog.csdn.net/github_39533414/article/details/80144890 新版本

# 初始化
```
groupadd mysql    //创建mysql组
useradd -g mysql mysql //创建mysql用户，并加入mysql组
cd /usr/local/mysql
mkdir ./data
chown -R mysql:mysql ./
vi /etc/my.cnf     https://www.cnblogs.com/langdashu/p/5889352.html 配置mysql
bin/mysqld --initialize --user=mysql --basedir=/usr/local/mysql --datadir=/usr/local/mysql/data //初始化数据库
```

# 建立服务
```
cp mysql.server /etc/init.d/mysqld
chmod +x /etc/init.d/mysqld 
chkconfig --add mysqld      添加到系统服务
chkconfig  --list mysqld   检查服务是否生效  
```

# 配置全局环境变量
```
vi /etc/profile
在 profile 文件底部添加如下两行配置，保存后退出
export PATH=$PATH:/usr/local/mysql/bin:/usr/local/mysql/lib
source /etc/profile
```

# 启动服务
```
service mysqld start   会启动mysqld_safe（root）和mysqld(mysql)
```

# 相关操作
```
//远程连接----------
mysql -h 192.168.5.116 -P 3306 -u root -p123456
//启动-------
./bin/mysqld --defaults-file=./my.cnf -u mysql -p    //注意这是启动用户
或者 ./bin/mysql -S .sock -u root -p
//停止------
./bin/mysqladmin -uroot -S /data_dir/malx/sqldata/mysql.sock shutdown
//登陆数据库,会提示你输入密码
mysql -u root -p   //这里默认的是/usr/local/mysql/my.cnf
mysql --defaults-file=my.cnf -u malx -p
//打开数据库
/usr/local/mysql/bin/mysql --defaults-file=/usr/local/mysql/my.cnf -u root -p  //root 超级权限用户
//无需登录，打开服务-------
 ./bin/mysqld_safe --defaults-file=my.cnf --user=mysql --skip-grant-tables
//给与用户权限--------------------
GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' IDENTIFIED BY '123456';   #只允许本地root登陆
flush privileges;
//删除用户------------------------
use mysql;
drop user root@'%';           %表示所有地址，但是不包括localhost
drop user root@localhost;
//新建用户--------------------
create user 'malx'@'%' IDENTIFIED by '111';  //注意%代表可以从外网连接
添加完用户后，别忘了赋予权限。
//--flush privileges 
 The MySQL server is running with the --skip-grant-tables option so it cannot execute this statement
//修改密码
set password for root@localhost = password('123456')
或者
格式：mysqladmin -u用户名 -p旧密码 password 新密码。 
//创建数据库


//创建数据表

//查询
select * from cnv_reportresult t where enable=-1 and sampleid in ("CL170407","CL170411");
//严格模式
[mysqld]
sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES
这是就会限制你的长度
```

### mysqld和mysqld_safe
```
mysqld_safe 作为mysqld 启动脚本，开启了守护mysqld进程的任务

mysqld_safe相当于mysqld的守护进程，当mysqld死了，mysqld_safe会把它拉起

mysqld_safe启动能够为mysqld分配系统资源
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

### mysql锁
1、简要概念
```
![](https://blog.csdn.net/xifeijian/article/details/20313977)

(1) 共享锁和排他锁（按级别划分）
    共享锁（读锁）：其他事务可以读，但不能写。
    排他锁（写锁） ：其他事务不能读取，也不能写。

    事务拿到某一行记录的共享S锁，才可以读取这一行，并阻止别的事物对其添加X锁
    事务拿到某一行记录的排它X锁，才可以修改或者删除这一行

    共享锁的目的是提高读读并发
    排他锁的目的是为了保证数据的一致性

(2) 表级锁、行级锁、页级锁（按粒度划分）
    1) 表级锁：开销小，加锁快；不会出现死锁；锁定粒度大，发生锁冲突的概率最高,并发度最低。
    2) 行级锁：开销大，加锁慢；会出现死锁；锁定粒度最小，发生锁冲突的概率最低,并发度也最高。使用行级锁定的主要是InnoDB存储引擎。
    3) 页面锁：开销和加锁时间界于表锁和行锁之间；会出现死锁；锁定粒度界于表锁和行锁之间，并发度一般。

(3) 自动锁、显示锁（加锁方式）
    INSERT、UPDATE、DELETE InnoDB会自动加排他锁，对于普通SELECT语句，InnoDB不会加任何锁，当然也可以显示加锁。

(4) 悲观锁和乐观锁（使用方式）https://zhuanlan.zhihu.com/p/31537871
    需要自己主动加锁和释放锁
    1) 悲观锁
    悲观的态度类来防止一切数据冲突，它是以一种预防的姿态在修改数据之前把数据锁住，然后再对数据进行读写。
    释放锁之前任何人都不能对其数据进行操作,性能不高
    读锁：加锁 LOCK tables test_db READ; 释放 UNLOCK TABLES;
    读锁：加锁 LOCK tables test_db WRITE; 释放 UNLOCK TABLES;
    2) 乐观锁
    操作数据时不会对操作的数据进行加锁（这使得多个任务可以并行的对数据进行操作），只有到数据提交的时候才通过一种机制来验证数据是否存在冲突(一般实现方式是通过加版本号然后进行版本号的对比方式实现)
    乐观锁是一种并发类型的锁，其本身不对数据进行加锁通而是通过业务实现锁的功能
```

2、MyISAM表锁
```
MyISAM存储引擎只支持表锁
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
```

3、InnoDB锁
```
InnoDB与MyISAM的最大不同有两点：一是支持事务（TRANSACTION）；二是采用了行级锁

(1) 事务隔离
    > 并发的问题
        l  更新丢失（Lost Update）：当两个或多个事务选择同一行，然后基于最初选定的值更新该行时，由于每个事务都不知道其他事务的存在，就会发生丢失更新问题－－最后的更新覆盖了由其他事务所做的更新。例如，两个编辑人员制作了同一文档的电子副本。每个编辑人员独立地更改其副本，然后保存更改后的副本，这样就覆盖了原始文档。最后保存其更改副本的编辑人员覆盖另一个编辑人员所做的更改。如果在一个编辑人员完成并提交事务之前，另一个编辑人员不能访问同一文件，则可避免此问题。
        l  脏读（Dirty Reads）：一个事务正在对一条记录做修改，在这个事务完成并提交前，这条记录的数据就处于不一致状态；这时，另一个事务也来读取同一条记录，如果不加控制，第二个事务读取了这些“脏”数据，并据此做进一步的处理，就会产生未提交的数据依赖关系。这种现象被形象地叫做"脏读"。
        l  不可重复读（Non-Repeatable Reads）：一个事务在读取某些数据后的某个时间，再次读取以前读过的数据，却发现其读出的数据已经发生了改变、或某些记录已经被删除了！这种现象就叫做“不可重复读”。
        l  幻读（Phantom Reads）：一个事务按相同的查询条件重新读取以前检索过的数据，却发现其他事务插入了满足其查询条件的新数据，这种现象就称为“幻读”。
并发事务处理带来的问题

```



```
https://www.cnblogs.com/luyucheng/p/6297752.html

(1) 表级锁：开销小，加锁快；不会出现死锁；锁定粒度大，发生锁冲突的概率最高,并发度最低。
    1) 使用表级锁定的主要是MyISAM，MEMORY，CSV等一些非事务性存储引擎
    2) MyISAM在执行查询语句（SELECT）前，会自动给涉及的所有表加读锁，在执行更新操作（UPDATE、DELETE、INSERT等）前，会自动给涉及的表加写锁，这个过程并不需要用户干预
    3) 优化
    优化MyISAM存储引擎锁定问题，关键的就是如何让其提高并发度。让锁定的时间变短，然后就是让可能并发进行的操作尽可能的并发。
        > 查询表级锁争用情况
        show status like 'table%';
            +----------------------------+---------+
            | Variable_name              | Value   |
            +----------------------------+---------+
            | Table_locks_immediate      | 100     |    //产生表级锁定的次数
            | Table_locks_waited         | 11      |    //出现表级锁定争用而发生等待的次数
            +----------------------------+---------+
            两个状态值都是从系统启动后开始记录，出现一次对应的事件则数量加1。如果这里的Table_locks_waited状态值比较高，那么说明系统中表级锁定争用现象比较严重，就需要进一步分析为什么会有较多的锁定资源争用了。
        > 缩短锁定时间
        让Query执行时间尽可能的短
            a) 尽两减少大的复杂Query，将复杂Query分拆成几个小的Query分布进行；
            b) 尽可能的建立足够高效的索引，让数据检索更迅速；
            c) 尽量让MyISAM存储引擎的表只存放必要的信息，控制字段类型；
            d) 利用合适的机会优化MyISAM表数据文件。
        > 分离能并行的操作
            插入并发性,重要参数concurrent_insert
            concurrent_insert=2，无论MyISAM表中有没有空洞，都允许在表尾并发插入记录；
            concurrent_insert=1，如果MyISAM表中没有空洞（即表的中间没有被删除的行），MyISAM允许在一个进程读表的同时，另一个进程从表尾插入记录。这也是MySQL的默认设置；
            concurrent_insert=0，不允许并发插入。
            可以利用MyISAM存储引擎的并发插入特性，来解决应用中对同一表查询和插入的锁争用。例如，将concurrent_insert系统变量设为2，总是允许并发插入；同时，通过定期在系统空闲时段执行OPTIMIZE TABLE语句来整理空间碎片，收回因删除记录而产生的中间空洞
            对于读写互相阻塞的表锁，将concurrent_insert=2（无论MyISAM表中有没有空洞，都允许在表尾并发插入记录），总是允许并发插入；同时，通过定期在系统空闲时段执行OPTIMIZE TABLE语句来整理空间碎片，收回因删除记录而产生的中间空洞。
        > 合理利用读写优先级
            MyISAM存储引擎的是读写互相阻塞的，具体情况具体定

(2) 行级锁：开销大，加锁慢；会出现死锁；锁定粒度最小，发生锁冲突的概率最低,并发度也最高。使用行级锁定的主要是InnoDB存储引擎。
    1) 类型
        共享锁（S）：读锁，允许一个事务去读一行，阻止其他事务获得相同数据集的排他锁。
        排他锁（X）：写锁，允许获得排他锁的事务更新数据，阻止其他事务取得相同数据集的共享读锁和排他写锁。
        //允许行锁和表锁共存，实现多粒度锁机制
        意向共享锁（IS）：事务打算给数据行加行共享锁，事务在给一个数据行加共享锁前必须先取得该表的 IS 锁。
        意向排他锁（IX）：事务打算给数据行加行排他锁，事务在给一个数据行加排他锁前必须先取得该表的 IX 锁。
        //如果一个事务请求的锁模式与当前的锁兼容， InnoDB 就将请求的锁授予该事务； 反之， 如果两者不兼容，该事务就要等待锁释放
        兼容性  IS	IX	 S	  X
        IS	   兼容	兼容 兼容 互斥
        IX	   兼容	兼容 互斥 互斥
        S      兼容	互斥 兼容 互斥
        X	   互斥	互斥 互斥 互斥
    2) 加锁方式
        意向锁是 InnoDB 自动加的，不需用户干预
        对于 UPDATE、DELETE 和 INSERT 语句，InnoDB会自动给涉及数据集加排他锁（X)
        对于普通 SELECT 语句，InnoDB 不会加任何锁, 需要自己加 SELECT * FROM table_name WHERE ... LOCK IN SHARE MODE。
        > 

(3) 页面锁：开销和加锁时间界于表锁和行锁之间；会出现死锁；锁定粒度界于表锁和行锁之间，并发度一般。




```

3、自动锁、显示锁（加锁方式）
```
INSERT、UPDATE、DELETE InnoDB会自动加排他锁，对于普通SELECT语句，InnoDB不会加任何锁，当然也可以显示加锁。
```

4、DML锁，DDL锁（操作划分）
```

```

5、悲观锁和乐观锁（使用方式）
```
https://zhuanlan.zhihu.com/p/31537871

需要自己主动加锁和释放锁

(1) 悲观锁
悲观的态度类来防止一切数据冲突，它是以一种预防的姿态在修改数据之前把数据锁住，然后再对数据进行读写。
释放锁之前任何人都不能对其数据进行操作,性能不高
读锁：加锁 LOCK tables test_db READ; 释放 UNLOCK TABLES;
读锁：加锁 LOCK tables test_db WRITE; 释放 UNLOCK TABLES;

(2) 乐观锁
操作数据时不会对操作的数据进行加锁（这使得多个任务可以并行的对数据进行操作），只有到数据提交的时候才通过一种机制来验证数据是否存在冲突(一般实现方式是通过加版本号然后进行版本号的对比方式实现)
乐观锁是一种并发类型的锁，其本身不对数据进行加锁通而是通过业务实现锁的功能
```

### 隔离级别
```

```

### mysql事务
```
https://yq.aliyun.com/articles/513493
https://www.cnblogs.com/huanongying/p/7021555.html

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

// 并发问题
    脏读（dirty read）：读到另一个事务的未提交更新数据，即读取到了脏数据；
    不可重复读（unrepeatable read）：对同一记录的两次读取不一致，因为另一事务对该记录做了修改；
    幻读（虚读）（phantom read）：对同一张表的两次查询不一致，因为另一事务插入了一条记录；
// 不可重复读和幻读的区别：
    不可重复读是读取到了另一事务的更新；
    幻读是读取到了另一事务的插入（MySQL中无法测试到幻读）；
// 隔离级别
    事务隔离级别	            脏读	不可重复读	幻读
    读未提交（read-uncommitted）是	    是	       是
    不可重复读（read-committed）否	    是	       是
    可重复读（repeatable-read）	否	    否	       是
    串行化（serializable）	    否	    否	       否

// 默认隔离级别为Repeatable read，可以通过下面语句查看：
    select @@tx_isolation

// 设置隔离级别
    set transaction isolation level read committed

```

### go-xorm
```
https://github.com/xormplus/xorm



```
