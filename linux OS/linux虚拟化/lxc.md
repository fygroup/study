### lxc
```
// 模板
/usr/share/lxc/templates 
一个 LXC 模板实质上就是一个脚本，用于创建指定环境下的容器。当你创建 LXC 容器时，你需要用到它们

// 桥接网卡
安装完 LXC 工具后，LXC自动创建了一块桥接网卡(lxcbr0，可以在 /etc/lxc/default.conf 中设置)
创建的 LXC 容器会自动链接到这个桥接网卡上

// 容器
lxc-create -n <container-name> -t <template> [--release utopic]
容器被放到 /var/lib/lxc/<容器名> 目录下
容器的根文件系统放在 /var/lib/lxc/<容器名>/rootfs 目录下

lxc-ls // 查看所有容器
lxc-start -n <container-name> -d // 启动容器，参数-d将容器作为后台进程打开

```

