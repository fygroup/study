https://zhuanlan.zhihu.com/p/36288538
https://zhuanlan.zhihu.com/p/23112605
http://spark.apachecn.org/#/docs/
https://www.cnblogs.com/ITtangtang/p/7977312.html
https://blog.csdn.net/zhumr/article/details/52540994

### RDD
```
弹性分布式数据集（Resilient Distributed Dataset，RDD）



```

### DataFrame
```
```

### DataSet
```
```

### RDD DataFrame DataSet
```
1、RDD
    从一开始RDD就是Spark提供的面向用户的主要API。从根本上来说，一个RDD就是你的数据的一个不可变的分布式元素集合，在集群中跨节点分布，可以通过若干提供了转换和处理的底层API进行并行处理。
    (1) 使用RDD的场景
        > 对数据集进行最基本的转换、处理和控制
        > 数据是非结构化的，比如流媒体或者字符流
        
    (2) DataFrame和DataSet都是基于RDD提供的
```