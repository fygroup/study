### 文档
```
https://jasper-zhang1.gitbooks.io/influxdb/content/Guide/querying_data.html
```

### 配置
```
// 不同的工具
    influxd          influxdb服务器
    influx           influxdb命令行客户端
    influx_inspect   查看工具
    influx_stress    压力测试工具
    influx_tsm       数据库转换工具（将数据库从b1或bz1格式转换为tsm1格式）

// config
    influxd -config /etc/influxdb/influxdb.conf
    //或者
    export INFLUXDB_CONFIG_PATH=/etc/influxdb/influxdb.conf

// /var/lib/influxdb/
    data            存放最终存储的数据，文件以.tsm结尾
    meta            存放数据库元数据
    wal             存放预写日志文件

// 建议使用两块SSD卷
    一个是为了influxdb/wal，较少的磁盘空间但是较高的IOPS
    一个是为了influxdb/data，更多的磁盘空间和较低的IOPS

// 权限(主要针对数据目录的权限)
    如果InfluxDB没有使用标准的数据和配置文件的文件夹的话，你需要确定文件系统的权限是正确的：
    chown influxdb:influxdb influxdb/wal
    chown influxdb:influxdb influxdb/data

// 端口
    TCP端口8086用作InfluxDB的客户端和服务端的http api通信
    TCP端口8088给备份和恢复数据的RPC服务使用
```

### config
```
https://www.cnblogs.com/guyeshanrenshiwoshifu/p/9188368.html


#全局配置
bind-address = ":8088"  # 备份恢复时使用，默认值为8088

[meta]
    dir = "/var/lib/influxdb/meta"

[data]
    dir = "/var/lib/influxdb/data"
    engine = "tsm1"
    wal-dir = "/var/lib/influxdb/wal"

[admin]         #管理界面
    enabled = true          # 是否启用该模块，默认值 ： false
    bind-address = ":8083"  # 绑定地址，默认值 ：":8083"
    https-enabled = false  # 是否开启https ，默认值 ：false
    https-certificate = "/etc/ssl/influxdb.pem"  # https证书路径，默认值："/etc/ssl/influxdb.pem"

[http]
    enabled = true  # 是否启用该模块，默认值 ：true
    bind-address = ":8086"  # 绑定地址，默认值：":8086"


```

### 结构
```
database                        库
measurement                     表，influx的表很轻，可以建立上百万个
tag(索引), field, timestamp     列
retention policy                数据保留策略(如30天)
series                          集合
                                retention policy，measurement以及tag set
point                           行，一行数据


任意series编号	retention policy    measurement	    tag set
series 1	    autogen	            census	        location = 1,scientist = langstroth
series 2	    autogen	            census	        location = 2,scientist = langstroth
series 3	    autogen	            census	        location = 1,scientist = perpetua
series 4	    autogen	            census	        location = 2,scientist = perpetua
```

### tag & field 
```
field value就是数据，它们可以是字符串、浮点数、整数、布尔值
tag value只能是字符串

field没有索引，tag有索引

无法对tag value进行数学运算

但是tag越多，series也就越多，内存占用就越大。不要有太多的series


```

### 操作
```
(1) 插入
    1) 语法
        [command] <measurement>[,<tag>=<value>...] <field>=<value>[,<field>=<value>...] [unix-nano-timestamp]
    2) 实例
        insert cpu,host=server1,region=us_west field=0.64 1434067467000000000

(2) 查询
    //select field from measurement where tag?? 
    SELECT "host", "region", "value" FROM "cpu"
    SELECT * FROM "cpu_load_short" WHERE "value" > 0.9

(3) measure表的操作
    //显示所有表
    SHOW MEASUREMENTS    
    //删除表
    drop measurement disk_free

(4) RP操作
    //查看当前数据库Retention Policies
    show retention policies on "db_name"
    //创建新的Retention Policies
    create retention policy "rp_name" on "db_name" duration 3w replication 1 default
    //修改Retention Policies
    alter retention policy "rp_name" on "db_name" duration 30d default
    //删除Retention Policies
    drop retention policy "rp_name" on "db_name"

(5) CQ操作
    //显示所有已存在的连续查询
    SHOW CONTINUOUS QUERIES
    //新建连续查询
    CREATE CONTINUOUS QUERY <cq_name> ON <database_name>
    [RESAMPLE [EVERY <interval>] [FOR <interval>]]
    BEGIN 
    SELECT <function>(<stuff>)[,<function>(<stuff>)] 
    INTO <rp_name>.<different_measurement>
    FROM <current_measurement> [WHERE <stuff>] 
    GROUP BY time(<interval>)[,<stuff>]
    END
    //显示CQ结果
    select * from <rp_name>.<different_measurement>
    //删除Continuous Queries
    DROP CONTINUOUS QUERY <cq_name> ON <database_name>

```

### restful
```
(1) 请求(post)
    curl -i -XPOST http://localhost:8086/query --data-urlencode "q=CREATE DATABASE mydb"

(2) 写(post)
    curl -i -XPOST 'http://localhost:8086/write?db=mydb' --data-binary 'cpu_load_short,host=server02 value=0.67
cpu_load_short,host=server02,region=us-west value=0.55 1422568543702900257
cpu_load_short,direction=in,host=server01,region=us-west value=2.0 1422568543702900257'
    //多行用\n分离

(3) 查询(get)
    curl -G 'http://localhost:8086/query?pretty=true' --data-urlencode "db=mydb" --data-urlencode "q=SELECT \"value\" FROM \"cpu_load_short\" WHERE \"region\"='us-west';SELECT count(\"value\") FROM \"cpu_load_short\" WHERE \"region\"='us-wes t'"
    > pretty=true 让返回的json格式化
    > 返回
        > 最大行限制
            默认会把返回的数目截断为10000条，如果有超过10000条返回，那么返回体里面会包含一个"partial":true的标记
        > 分块传输(如果分块，最大行就不会限制)
            curl -G 'http://localhost:8086/query' --data-urlencode "db=deluge" --data-urlencode "chunked=true" --data-urlencode "chunk_size=20000" --data-urlencode "q=SELECT * FROM liters"    

```

### CQ和RP
```
连续查询(Continuous Queries简称CQ)和保留策略(Retention Policies简称RP)，分别用来处理数据采样和管理老数据的。

// 创建数据库
CREATE DATABASE "food_data"

(1) RP
    // 创建一个保留52周数据的RP
    CREATE RETENTION POLICY "a_year" ON "db" DURATION 52w REPLICATION 1
    // 默认RP
    InfluxDB会自动生成一个叫做autogen的RP，并作为数据库的默认RP，autogen这个RP会永远保留数据。在输入上面的命令之后，two_hours会取代autogen作为food_data的默认RP

(2) CQ
    // 创建CQ
    CREATE CONTINUOUS QUERY "cq_30m" ON "food_data" BEGIN
    SELECT mean("website") AS "mean_website",mean("phone") AS "mean_phone"
    INTO "a_year"."downsampled_orders"                                    
    FROM "orders"                                                         
    GROUP BY time(30m)
    END

    cq_30m告诉InfluxDB每30分钟计算一次measurement为orders并使用默认RPtow_hours的字段website和phone的平均值，然后把结果写入到RP为a_year，两个字段分别是mean_website和mean_phone的measurement名为downsampled_orders的数据中
    SELECT * FROM "a_year"."downsampled_orders" LIMIT 5
```

### curl实例
```
 //create database
 curl -i -XPOST 'http://localhost:8086/query' --data-urlencode "create database mydb1"
 //RP
 curl -i -XPOST 'http://localhost:8086/query' --data-urlencode "q=CREATE RETENTION POLICY \"rp1\" ON \"mydb1\" DURATION 1d REPLICATION 1"
 //CQ
 curl -i -XPOST 'http://localhost:8086/query' --data-urlencode "q=CREATE CONTINUOUS QUERY \"cq1\" ON \"mydb1\" BEGIN BEGIN select mean(\"field1\") as \"mean_field1\",mean(\"field2\") as \"mean_field2\" into \"rp1\".\"cq1_measure1\" from \"measure1\" group by time(10m) end "
 //insert
 curl -i -XPOST 'http://localhost:8086/write?db=mydb1' --data-binary "measure1,tag1=daa field1=22"
 //select
 curl -i -G "http://localhost:8086/query?pretty=true" --data-urlencode "db=mydb" --data-urlencode "chunked=true" --data-urlencode "chunk-size=1" --data-urlencode "q=select * from measure1"

```