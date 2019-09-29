### 文档
```
https://jasper-zhang1.gitbooks.io/influxdb/content/Guide/querying_data.html
```

### 配置
```
// config
    influxd -config /etc/influxdb/influxdb.conf
    //或者
    export INFLUXDB_CONFIG_PATH=/etc/influxdb/influxdb.conf

// 建议使用两块SSD卷
    一个是为了influxdb/wal，较少的磁盘空间但是较高的IOPS
    一个是为了influxdb/data，更多的磁盘空间和较低的IOPS

// 权限(主要针对数据目录的权限)
    如果InfluxDB没有使用标准的数据和配置文件的文件夹的话，你需要确定文件系统的权限是正确的：
    chown influxdb:influxdb /mnt/influx
    chown influxdb:influxdb /mnt/db

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

但是tag越多，series也就越多。但是不要有太多的series


```

### 语法
```
[command] <measurement>[,<tag>=<value>...] <field>=<value>[,<field>=<value>...] [unix-nano-timestamp]

// 例如
    insert cpu,host=server1,region=us_west value=0.64 1434067467000000000

    //查询 表cpu 中的 tag host、region、value
    SELECT "host", "region", "value" FROM "cpu"

    SELECT * FROM "cpu_load_short" WHERE "value" > 0.9
```

### restful
```
(1) 请求(post)
    curl -i -XPOST http://localhost:8086/query --data-urlencode "q=CREATE DATABASE mydb"

(2) 写(post)
    curl -i -XPOST 'http://localhost:8086/write?db=mydb' --data-binary 'cpu_load_short,host=server02 value=0.67
cpu_load_short,host=server02,region=us-west value=0.55 1422568543702900257
cpu_load_short,direction=in,host=server01,region=us-west value=2.0 1422568543702900257'
    > 多行用\n分离

(3) 查询(get)
    curl -G 'http://localhost:8086/query?pretty=true' --data-urlencode "db=mydb" --data-urlencode "q=SELECT \"value\" FROM \"cpu_load_short\" WHERE \"region\"='us-west';SELECT count(\"value\") FROM \"cpu_load_short\" WHERE \"region\"='us-west'"
    > pretty=true 让返回的json格式化
    > 返回
        > 最大行限制
            默认会把返回的数目截断为10000条，如果有超过10000条返回，那么返回体里面会包含一个"partial":true的标记
        > 分块传输(如果分块，最大行就不会限制)
            curl -G 'http://localhost:8086/query' --data-urlencode "db=deluge" --data-urlencode "chunked=true" --data-urlencode "chunk_size=20000" --data-urlencode "q=SELECT * FROM liters"    

```

### 关键字
```

```