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
```
https://www.cnblogs.com/luyucheng/p/6297752.html

(1) 表级锁：开销小，加锁快；不会出现死锁；锁定粒度大，发生锁冲突的概率最高,并发度最低。
    > 使用表级锁定的主要是MyISAM，MEMORY，CSV等一些非事务性存储引擎
    > MyISAM在执行查询语句（SELECT）前，会自动给涉及的所有表加读锁，在执行更新操作（UPDATE、DELETE、INSERT等）前，会自动给涉及的表加写锁，这个过程并不需要用户干预
    >优化
    优化MyISAM存储引擎锁定问题，关键的就是如何让其提高并发度。让锁定的时间变短，然后就是让可能并发进行的操作尽可能的并发。
    1) 查询表级锁争用情况
       show status like 'table%';
        +----------------------------+---------+
        | Variable_name              | Value   |
        +----------------------------+---------+
        | Table_locks_immediate      | 100     |    //产生表级锁定的次数
        | Table_locks_waited         | 11      |    //出现表级锁定争用而发生等待的次数
        +----------------------------+---------+
    2) 缩短锁定时间
      让Query执行时间尽可能的短
        a)尽两减少大的复杂Query，将复杂Query分拆成几个小的Query分布进行；
        b)尽可能的建立足够高效的索引，让数据检索更迅速；
        c)尽量让MyISAM存储引擎的表只存放必要的信息，控制字段类型；
        d)利用合适的机会优化MyISAM表数据文件。
    3) 分离能并行的操作
        对于读写互相阻塞的表锁，将concurrent_insert=2（无论MyISAM表中有没有空洞，都允许在表尾并发插入记录），总是允许并发插入；同时，通过定期在系统空闲时段执行OPTIMIZE TABLE语句来整理空间碎片，收回因删除记录而产生的中间空洞。
    4) 合理利用读写优先级
        


(2) 行级锁：开销大，加锁慢；会出现死锁；锁定粒度最小，发生锁冲突的概率最低,并发度也最高。
    使用行级锁定的主要是InnoDB存储引擎。
(3) 页面锁：开销和加锁时间界于表锁和行锁之间；会出现死锁；锁定粒度界于表锁和行锁之间，并发度一般。




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
