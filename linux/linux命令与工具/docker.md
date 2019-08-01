### 容器使用相关参数
```
(1) 命令
    1) 创建一个容器
        sudo docker run     -i -t   ubuntu   /bin/bash
                    动作    参数     镜像     运行的命令
    2) 查看当前系统的容器
        docker ps -a    //所有容器（不管停止与否）
        docker ps       //正在运行的容器
    3) 重启容器和附着容器
        docker start dockerID
        docker attach dockerID
    4) 创建守护式容器（没有交互式会话，适合运行应用程序）
        sudo docker run --name dockerID -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
    5) 输出日志
        sudo docker logs dockerID  //会输出上面的内容
        sudo docker logs --tail 10 dockerID
        sudo docker logs --tail 10 -f dockerID
    6) 查看容器内的进程
        sudo docker top dockerID
    7) 从外部在容器内运行进程
        sudo docker exec dockerID -d echo helloworld
        sudo docker exec -t -i dockerID /bin/bash
    8) 停止守护式容器
        sudo docker stop dockerID    
    9) 设置自动重启容器的标志
        sudo docker run --restart=on-failure:3 --name dockerID -d ubuntu echo heloworld
    10) 查看容器信息
        sudo docker inspect dockerID
    11) 删除容器（运行中的容器无法删除）
        sudo docker stop dockerID 或者 sudo docker kill dockerID
        sudo docker rm dockerID
        


(2) 参数
    1) run 创建容器
        -i          保证容器STDIN开启
        -t          分配一个伪终端          
        --name      容器命名
        -d          后台运行
        --restart   自动重启, always  on-failure(退出码非0时就重启)  on-failure:3(重启次数)
    2) logs 日志
        --tail      输出尾部多少行
        -f          最新日志
        -t          加时间戳
    3) top 查看容器内的进程
    4) exec 外部往容器内运行命令
    5) inspect 产看容器信息
        -f/--format= 选定查看信息
    6) rm
    7) kill
    8) stop       
    9) ps 
        -a          显示所有容器
        -q          只列出容器ID, docker ps -a -q
```

### 镜像
1、基础概念
```
(1) bootfs(引导文件系统)
    类似linux，容器启动后，引导文件系统会被卸载

(2) rootfs(root文件系统)
    位于引导文件系统之上
    linux种rootfs会先以只读方式加载，引导结束并完成检查后，才切换成读写模式。但是docker中永远是只读

(3) union mount(联合加载)
    在rootfs基础上一次同时加载多个文件系统，然后将他们叠加在一块。这样文件系统包含所有底层文件和目录，又称为镜像

(4) copy-on-write(写时复制)
    当创建一个新容器时，先建出一个镜像栈，并在栈的最顶端添加一个读写层。这个读写层+下面的镜像+配置数据 = 容器
    
```
2、操作
```
(1) 列出所有可用的镜像
    > docker images
        REPOSITORY  TAG     IMAGE_ID  CREATED   SIZE
        仓库        版本     imageID    创建时间  大小
        一个仓库，多个版本的镜像
    > docker images ubuntu
        查看ubuntu镜像内容

(2) 拉取镜像      
    sudo docker pull ubuntu  

(3) 镜像+tagID
    sudo docker run -i -t --name dockerID ubuntu:12.04 /bin/bash

(4) 用户仓库
    命名 malx/ubuntu

(6) 查找镜像
    sudo docker search puppet
        Name

(7) 构建镜像
    

```