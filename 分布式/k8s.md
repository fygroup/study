
### openstack，docker，mesos，k8s 关系
```
OpenStack
    针对 Iaas 平台，以资源为中心，可以为上层的 PaaS 平台提供存储、网络、计算等资源

Docker
    主要针对 Paas 平台，是以应用为中心

Kubernetes(k8s)
    面向应用的 PaaS 层，强项在于容器编排，可以很好解决应用上云的问题

Mesos
    Apache的顶级开源项目，管理的核心目标对象既不是虚拟机/物理机，也不是容器，而是各种各样的计算资源（CPU、memory、disk、port、GPU等等）

```

### k8s + docker
```
k8s用于容器和虚拟机集群的管理，一切都基于分布式

一个K8S系统，通常称为一个K8S集群（Cluster）
这个集群主要包括两个部分：一个Master节点（主节点）和一群Node节点（计算节点）

// 主要组件
etcd                保存集群状态
apiserver           资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制
controller manager  维护管理集群状态
scheduler           资源调度
kubelet             维护容器生命周期
kube-proxy          集群内部每个节点的网络代理和负载均衡

1、Master组件
    Master组件提供集群的管理控制中心，它可以在集群中任何节点上运行
    (1) kube-apiserver
        用于暴露Kubernetes API。任何的资源请求/调用操作都是通过kube-apiserver提供的接口进行
    (2) ETCD
        etcd是Kubernetes提供默认的一致性系统，可用于存储集群的相关数据
    (3) kube-controller-manager
        管理控制器，它们是集群中处理常规任务的后台线程。逻辑上，每个控制器是一个单独的进程
        1) 节点（Node）控制器
        2) 副本（Replication）控制器
            负责维护系统中每个副本中的pod
        3) 端点（Endpoints）控制器
            填充Endpoints对象（即连接Services＆Pods）
        4) Service Account和Token控制器
            为新的Namespace 创建默认帐户访问API Token
    (4) cloud-controller-manager
        1) 节点（Node）控制器
        2) 路由（Route）控制器
        3) Service控制器
        4) 卷（Volume）控制器
    (5) kube-scheduler
        监视新创建没有分配到Node的Pod，为Pod选择一个Node

2、Node组件
    提供Kubernetes运行时环境，以及维护Pod。一个Node可以是VM或物理机
    (1) kubelet
        kubelet是主要的节点代理，它会监视已分配给节点的pod
        安装Pod所需的volume、下载Pod、Pod中运行的 docker（或experimentally，rkt）容器、定期执行容器健康检查等
    (2) kube-proxy
        维护网络规则并执行连接转发来实现Kubernetes服务抽象，每一个节点也运行一个简单的网络代理和负载均衡
    (3) docker
        docker用于运行容器
    (4) RKT
        rkt运行容器，作为docker工具的替代方案
    (5) supervisord
        supervisord是一个轻量级的监控系统，用于保障kubelet和docker运行
    (6) fluentd
        fluentd是一个守护进程，可提供cluster-level logging


```