https://yeasy.gitbooks.io/docker_practice/content/

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
        > 资源限制
            > 内存
                > 大小
                    -m              内存限制
                        -m 500M
            > cpu        
                > 核的控制
                    --cpuset-cpus   设置容器可以在哪些 CPU 核上运行
                        --cpuset-cpus="1,3"
                        --cpuset-cpus="0-2"
                > 份额控制
                    –cpu-shares     指定容器所使用的CPU份额值
                        每个docker容器的cpu份额都是1024。单独一个容器的份额是没有意义的，只有在同时运行多个容器时，容器的cpu加权的效果才能体现出来
                > 周期控制
                    –cpu-period是用来指定容器对CPU的使用要在多长时间内做一次重新分配
                    -cpu-quota是用来指定在这个周期内，最多可以有多少时间用来跑这个容器。
                    如果容器进程需要每1秒使用单个CPU的0.2秒时间，可以将cpu-period设置为1000000（即1秒），cpu-quota设置为200000（0.2秒）
                    在多核情况下，如果允许容器进程需要完全占用两个CPU，则可以将cpu-period设置为100000（即0.1秒），cpu-quota设置为200000（0.2秒）


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
    10) cp
        docker cp aaa:/aaa/aa /home/malx

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
    > docker inspect ubuntu
        查看镜像详细信息

(2) 拉取镜像      
    docker pull ubuntu  

(3) 镜像+tagID
    docker run -i -t --name dockerID ubuntu:12.04 /bin/bash

(4) 用户仓库
    命名 malx/ubuntu

(6) 查找镜像
    docker search puppet

(7) 构建镜像
    1) commit
        docker commit dockerID imageID 
        docker commit 4aa3adv5asx malx/ubuntu
                       目标容器  命名的新镜像
        docker commit -m="aaada" --author="malx" dockerID imageID                  
        docker commit -m="aaada" --author="malx" 4aa3adv5asx malx/ubuntu:lastest

    2) Dockerfile（如下）

(8) 镜像的导出与导入
    1) 导出
        docker save -o 保存的文件名 镜像
    2) 导入
        docker load --input 文件

```
3、Dockerfile
```
(1) 构建
    1) 此目录称为构建上下文，docker会将构建上下文和上下文中的文件和目录上传到docker守护进程
        mkdir static_web
        cd static_web
        touch Dockerfile

    2) 执行流程
        docker build -t="ubantu/malx:V1.0.0"
        // Dockerfile中的每条指令都会创建一个新的镜像层并对镜像进行提交。然后基于刚提交的镜像运行个新容器。然后执行Dockerfile的下一条指令
        // 当构建到一条指令出错时，进入运行指令的容器，执行正确的语句

(2) Dockerfile
    1) 内容样式
        # Version: 0.0.1
        FROM ubuntu:14.04               //指定一个已经存在的镜像，后续指令基于该镜像执行
        MAINTAINER malx "gudu@163.com"  //作者
        RUN apt-get update              //RUN会在当前镜像中运行指定的命令，默认是/bin/sh -c来执行
        RUN apt-get install -y nginx
        //或者 RUN ["apt-get", "install", "-y", "nginx"]
        EXPOSE 80                       //容器内应用程序会使用指定的端口。注意出于安全的原因，docker不会自动打开端口，需要指定
    2) 缓存
        在构建的过程中，docker会缓存之前构建成功的镜像，当再重新构建时。它就会从最新的缓存处开始构建。如果不需要缓存想重新从头构建，那么：
        docker build --no-cache -t='malx/ubuntu' 
    3) ENV(设置远程变量)
        ENV REFRESHED_AT 2018-03-10  
        // 这个环境变量用来表名该镜像模板的最后更新时间，由于每次都更新，所以重置缓存，从头运行
    4) 构建历史
            docker history malx/ubuntu
    5) 新镜像启动容器
        1) 端口
            docker run -d -p 80 --name mydocker malx/ubuntu nginx -g 'daemon off;'
        2) 端口映射
            docker ps -l     
                ...   0.0.0.0:49154->80/tcp   ...
            docker port dockerID 80     //查看容器端口80的映射
            docker run -d -p 8080:80 --name mydocker malx/ununtu nginx -g 'daemon off;'     //将容器80映射到宿主机的8080端口上
            docker run -d -p 127.0.0.1:8080:80/udp --name mydocker malx/ubuntu nginx -g 'daemon off;'
        3) EXPOSE
            EXPOSE 80
            EXPOSE 2020
            //但是他并不会主动向外开放，需要运行时手动指定（-P）
            docker run -d -P --name mydocker malx/ubuntu nginx -g 'daemon off;'
(3) dockerfile的指令
    https://zhuanlan.zhihu.com/p/57335983
    https://www.jianshu.com/p/5f4b1ade9dfc

    1) FROM
        基础镜像
        FROM ubuntu:tag
    2) MAINTAINER
        指定维护者的信息，并应该放在FROM的后面
        MAINTAINER authors_name 
    3) CMD和ENTRYPOINT
        //用于指定一个容器启动时要运行的命令。注意区分RUN(RUN是镜像构建时运行的命令)
        docker run -i -t malx/ubuntu /bin/bash -l
        //可以用CMD的方式
        CMD ["/bin/bash", "-l"]
        //注意：
        每个Dockerfile只能有一条CMD命令。如果指定了多条命令，只有最后一条会被执行。
        如果用户启动容器时候指定了运行的命令，则会覆盖掉 CMD 指定的命令。
        ENTRYPOINT提供命令不会被启动容器覆盖
    
    4) RUN
        RUN命令是Dockerfile执行命令的核心部分
        每条run指令在当前基础镜像执行，并且提交新镜像。当命令比较长时，可以使用“/”换行
        RUN apt-get update ;\
            apt-get install xxx
        RUN xxxx;\
            xxxxx   

    5) WORKDIR
        //创建一个新容器时设置工作目录，ENTRYPOINT和CMD的指令都在该目录下运行
        WORKDIR /opt/webapp/db
        RUN bundle install
        WORKDIR /opt/webapp
        ENTRYPOINT ["rackup"]
        //在外部指定-w
        docker run -ti -w /var/log ubuntu pwd
        # /var/log

    6) ENV
        //设置环境变量
        ENV RVM_PATH /home/rvm
        ENV PATH /home/my/:$PATH
        //外部传递环境变量
        docker run -ti -e "WEB_PORT=8080" ubuntu /bin/bash

    7) USER
        //指定镜像会以什么样用户运行
        //指定运行容器时的用户名或UID，后续的 RUN 也会使用指定用户。
        USER nginx    //该容器会以nginx用户的身份来运行
    
    8) VOLUME
        VOLUME命令用于让你的容器访问宿主机上的目录，它绕过联合文件系统，用于共享数据和持久化
        格式为 VOLUME ["/data", "/data1"] 
        创建一个可以从本地主机或其他容器挂载的挂载点，一般用来存放数据库和需要保持的数据等
        docker run --privileged --name httpd -v /sys/fs/cgroup:/sys/fs/cgroup:ro -p 80:80 -d  httpd

    9) ADD
        用来将构建环境的文件和目录拷贝到镜像中
        ADD aaa /data/
        ADD aa.tar.gz /data/     //会解压到/data中

    10) COPY
        类似于ADD
        文件路径必须和dockerfile在同一个目录，不能复制以外的目录，而且不会解压
        COPY ./aaa /data

    11) ONBUILD
        ONBUILD 指定的命令在构建镜像时并不执行，而是在它的子镜像中执行

    12) EXPOSE
        指定在docker允许时指定的端口进行转发
        EXPOSE 6379
        EXPOSE 6379/tcp
        EXPOSE 6379/udp
        在运行时使用 -p <宿主端口>:<容器端口>

(4) 构建镜像
    docker build -t webapp:latest -f ./webapp/a.Dockerfile ./webapp
                   新镜像名称          指定dockerfile路径     dockerfile目录

    注意：一个镜像不能超过127层

(5) 镜像优化(减小镜像尺寸)
    1) 链式指令
        减少构建层数，且每一层末尾执行clean操作(一些缓存文件等)
    2) 分离编译镜像和部署镜像
        编译过程中的依赖文件，一旦应用程序编译完毕，这些文件就不再有用
        思路：获得内部程序需要的相关依赖库，然后将这些库ADD到镜像中
        https://cloud.tencent.com/developer/article/1401008


```

### dockerfile实例
```
https://github.com/CentOS/CentOS-Dockerfiles
```

### docker systemctl
```
创建容器：
docker run -tdi --name my_centos --privileged centos init
进入容器：
docker exec -it my_centos /bin/bash
```

### docker root
```
docker守护程序绑定到Unix套接字而不是TCP端口。 默认情况下，Unix套接字是由root用户拥有的，其他用户可以使用sudo访问它。 因此，docker守护程序始终以root用户身份运行。
为了避免在使用docker命令时必须使用sudo，请创建一个名为docker的Unix组并将用户添加到其中。 docker守护程序启动时，它将使docker组可以读取/写入Unix套接字的所有权。
```

### Docker使用非root用户
```
sudo groupadd docker
sudo gpasswd -a ${USER} docker
sudo systemctl restart docker
```

### 容器权限运行
```
// 以宿主机uid 1000:1000运行容器
docker run -u 1000:1000

// 前提是容器要存在1000:1000的用户，如果不存在
I have no name!@6ae7c7d24c8a:/$ id
uid=1008 gid=1008 groups=1008

```


### volumns管理
```


```

### 资源管理
```
```

### docker network
```
// 三种基本的网络模式
docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
46053022478c        bridge              bridge              local
48c25f0ece64        host                host                local
adca62142796        none                null                local

// 查看桥接详细内容
docker network inspect bridge
[
    {
        "Name": "bridge",
        "Driver": "bridge",
        "IPAM": {
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"    // docker0 eth
                }
            ]
        },
        "Containers": {
            "45e8bd7ca6532aac5822734e933a0f07a4dddacd92938ff2c5247092a15a59ad": {
                "Name": "malx_ubuntu",
                "IPv4Address": "172.17.0.2/16", // 容器malx_ubuntu的ip
            }
        }
    }
]
```

### docker网络模式
```
https://www.cnblogs.com/zuxing/articles/8780661.html
https://www.jianshu.com/p/d84cdfe2ea86

Bridge(默认) Container host none

但如果启动容器的时候使用host模式，那么这个容器将不会获得一个独立的Network Namespace，而是和宿主机共用一个Network Namespace。容器将不会虚拟出自己的网卡，配置自己的IP等，而是使用宿主机的IP和端口


// NAT转发
netstat -apn | grep "172.17.0.1"
udp        0      0 172.17.0.1:123          0.0.0.0:*                           14089/ntpd  

```

### Docker-in-Docker
```
在Docker容器中运行Docker

// 为什么需要docker-in-docker
在 CI 中，通常会有一个 CI Engine 负责解析流程，控制整个构建过程，而将真正的构建交给 Agent 去完成。例如，Jenkins 、GitLab
使用gitlab的docker，当然你也可以直接使用gitlab程序，但是部署起来麻烦
所以问题来了：Agent已经是容器化的，怎么在容器上构建镜像呢？这就要用到docker-in-docker

// 实现docker-in-docker三种方法
(1) DooD
    在宿主机上是通过/var/run/docker.sock套接字与docker守护程序通信的
    此方案的原理是将套接字放到容器内，此时容器内操作docker实际上是与宿主机的docker进行通信

    docker run -v /var/run/docker.sock:/var/run/docker.sock -ti docker
    docker images // 显示宿主机docker的镜像

    接下来就可以在容器中创建镜像了

(2) DinD
    此方案是在容器中创建完整的docker全家桶，官方给的image是 docker:dind
    并且需要特权模式才能运行

    docker run --privileged -d --name dind-test docker:dind

    此方案不推荐

(3) Sysbox
    上述方案1、2不安全，此方案结合了1和2的好处

```

### CI/CD run 部署
```
// 运行 gitlab-runner 容器
docker run -d --name gitlab-runner --restart always -v /var/run/docker.sock:/var/run/docker.sock -v /home/malx/docker_gitlab_runner:/etc/gitlab-runner gitlab/gitlab-runner

// 注册runner
docker run --rm -it -v /home/malx/docker_gitlab_runner:/etc/gitlab-runner gitlab/gitlab-runner register


```