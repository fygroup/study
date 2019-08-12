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
(3) 其他指令
    https://www.jianshu.com/p/5f4b1ade9dfc
    1) CMD
        //用于指定一个容器启动时要运行的命令。注意区分RUN(RUN是镜像构建时运行的命令)
        docker run -i -t malx/ubuntu /bin/bash -l
        //可以用CMD的方式
        CMD ["/bin/bash", "-l"]
        //注意：
        每个Dockerfile只能有一条CMD命令。如果指定了多条命令，只有最后一条会被执行。
        如果用户启动容器时候指定了运行的命令，则会覆盖掉 CMD 指定的命令。
    2) ENTRYPOINT
        与CMD类似
        //注意
        配置容器启动后执行的命令，并且不可被 docker run 提供的参数覆盖。
        每个 Dockerfile 中只能有一个 ENTRYPOINT，当指定多个时，只有最后一个起效。
    3) WORKDIR
        //创建一个新容器时设置工作目录，ENTRYPOINT和CMD的指令都在该目录下运行
        WORKDIR /opt/webapp/db
        RUN bundle install
        WORKDIR /opt/webapp
        ENTRYPOINT ["rackup"]
        //在外部指定-w
        docker run -ti -w /var/log ubuntu pwd
        # /var/log
    4) ENV
        //设置环境变量
        ENV RVM_PATH /home/rvm
        ENV PATH /home/my/:$PATH
        //外部传递环境变量
        docker run -ti -e "WEB_PORT=8080" ubuntu /bin/bash
    5) USER
        //指定镜像会以什么样用户运行
        USER nginx    //该容器会以nginx用户的身份来运行


```