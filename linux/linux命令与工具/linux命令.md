//---查看硬盘容量-------------------------------
 df -h
//------硬盘挂载----------------------------------
fdisk -l
mount -t ntfs-3g /dev/sdc1 /mnt/
//------更改yum 镜像-----------------------
在安装完CentOS后一般需要修改yum源，才能够在安装更新rpm包时获得比较理想的速度。国内比较快的有163源、sohu源。这里以163源为例子。
1. cd /etc/yum.repos.d
2. mv CentOS-Base.repo CentOS-Base.repo.backup
3. wget http://mirrors.163.com/.help/CentOS6-Base-163.repo
4. mv CentOS6-Base-163.repo CentOS-Base.repo
5.yum clean all
//---安装-------------------------
yum -y install 
//---ntfs文件系统 恢复文件-------------
当误删除时，立即umount 磁盘
ntfsundelete /dev/sdc1 -f -t 1m  //查看1个月内的删除文件
ntfsundelete /dev/sdc1 -d /home/wilson/tmp/ -u -i 881,2341,234   //恢复文件

//---环境---------------------------------------------


//---最大进程数---------------------------------------
-bash: fork: retry: Resource temporarily unavailable错误
ulimit -a来查看当前Linux系统的一些系统参数
ulimit -u 4096 修改max user processes的值
修改/etc/security/limits.d/90-nproc.conf文件中的值（永久生效）
*          soft    nproc    4096

//---so的查看---------------------------------
ldd可以列出一个程序所需要得动态链接库
ldconfig是一个动态链接库管理命令，为了让动态链接库为系统所共享,还需运行动态链接库的管理命令–ldconfig。
nm /lib64/libc.so.6|grep '\sGLIBC_2' 查看动态库的相关信息
strings 

//---rpm-------------------------------------
rpm -qpl xxx.rpm    //查看rpm软件将要安装的路径信息

rpm -i --relocate /usr/bin=/home/easwy/bin --relocate /usr/share/doc=/home/easwy/doc ext3grep-0.10.0-1.el5.rf.i386.rpm
更改目标路径

//---vsftp-----------------------------------
PORT（主动）模式
PASV（被动）模式
1）PORT（主动）模式模式只要开启服务器的21和20端口，而PASV（被动）模式需要开启服务器大于1024所有tcp端口和21端口。
认证方式：匿名用户认证、本地用户认证（linux的/etc/password本地用户）和虚拟用户认证（ftp自己的用户）
https://www.jianshu.com/p/f1788f596a57


### sh重定向
```
nohup sh start.sh 1>>log.txt 2>&1  &
1>>log.txt 标准输出定向到log.txt文件
2>&1       错误输出定向到标准输出
```

### mount
```
mount -t type -o options device dir
device: 要挂载的设备（必填）。有些文件系统不需要指定具体的设备，这里可以随便填一个字符串
dir: 挂载到哪个目录（必填）
type： 文件系统类型（可选）。大部分情况下都不用指定该参数，系统都会自动检测到设备上的文件系统类型
options： 挂载参数（可选）。 
（1）挂载虚拟文件系统
proc、tmpfs、sysfs、devpts等都是Linux内核映射到用户空间的虚拟文件系统
mount -t proc none /mnt
//将内核的proc文件系统挂载到/mnt，这样就可以在/mnt目录下看到系统当前运行的所有进程的信息，
//由于proc是内核虚拟的一个文件系统，并没有对应的设备，所以这里-t参数必须要指定，不然mount就不知道要挂载啥了。
//由于没有对应的源设备，这里none可以是任意字符串，取个有意义的名字就可以了，因为用mount命令查看挂载点信息时第一列显示的就是这个字符串。

（2）挂载虚拟的块设备
硬盘、光盘、软盘等都是常见的块设备，他们在Linux下的目录一般是/dev/hda1, /dev/cdrom, /dev/sda1，/dev/fd0这样的。而loop device是虚拟的块设备，主要目的是让用户可以像访问上述块设备那样访问一个文件。 loop device设备的路径一般是/dev/loop0, dev/loop1, ...等。
1）ISO：
    //利用mkisofs构建一个用于测试的iso文件
    mkdir -p iso/subdir01
    mkisofs -o ./test.iso ./iso
    
    //mount ISO 到目录 /mnt
    sudo mount ./test.iso /mnt
    //注意：此处要加-o loop表示这是个虚拟设备
    //mount: /dev/loop0 is write-protected, mounting read-only
    //mount成功，能看到里面的文件夹
    ls /mnt
    //subdir01
    //通过losetup命令可以看到占用了loop0设备
    losetup -a
    // /dev/loop0: []: (/home/dev/test.iso)
2）虚拟硬盘：
    //比如用一个文件想尝试btrfs文件系统，
    //因为btrfs对分区的大小有最小要求，所以利用dd命令创建一个128M的文件
    dd if=/dev/zero bs=1M count=128 of=./vdisk.img

    //在这个文件里面创建btrfs文件系统
    mkfs.btrfs ./vdisk.img

    //mount虚拟硬盘
    sudo mount ./vdisk.img /mnt/

    //在虚拟硬盘中创建文件成功
    sudo touch /mnt/aaaaaa
    ls /mnt/
    //aaaaaa

    //加上刚才上面mount的iso文件，我们已经用了两个loop device了
    losetup -a
    /dev/loop0: []: (/home/dev/test.iso)
    /dev/loop1: []: (/home/dev/vdisk.img)

（3）挂载多个设备到一个文件夹
将test.iso和vdisk.img都mount到了/mnt目录下，默认会用后面的mount覆盖掉前面的mount，只有当umount后面的device后，原来的device才看的到
（4）挂载一个设备到多个目录
也可以把一个设备mount到多个文件夹，这样在多个文件夹中都可以访问该设备中的内容。

（5）bind mount
    bind mount会将源目录绑定到目的目录，然后在目的目录下就可以看到源目录里的文件
    mkdir -p bind/bind1/sub1
    mkdir -p bind/bind2/sub2
    tree bind
    bind
    ├── bind1
    │   └── sub1
    └── bind2
        └── sub2

    mount --bind ./bind/bind1/ ./bind/bind2  
    tree bind
    bind
    ├── bind1
    │   └── sub1
    └── bind2
        └── sub1
    umount ./bind/bind2  //会还原之前的bind2/sub2

    //readonly bind
    mount -o bind,ro ./bind/bind1/ ./bind/bind2
    touch ./bind/bind2/sub1/aaa
    //touch: cannot touch './bind/bind2/sub1/aaa': Read-only file system

    //想让当前目录readonly，那么可以bind自己，并且指定readonly参数：
    mount -o bind,ro ./bind/bind1/ ./bind/bind1
```

### 查看总线上的所有设备
```
//查看总线上的所有设备
lspci

//查看网络设备的详细信息
iwconfig

```

### ip
```
地址、设备、路由表
//ip
10.100.6.13/24
斜杠后面的数字就表示子网掩码，数字具体代表32位子网掩码（二进制形式）中前面的“1”的个数

//地址管理
0、显示所有设备
    ip addr show 
1、显示设备的ip地址
    ip addr show dev eth1
2、增加删除设备的地址
    ip addr add 127.1.2.3/24 dev eth1
    ip addr del 127.1.2.3/24 dev eth1
3、显示接口统计
    ip -s link ls dev eth1

//路由表管理
ip route
ip route show dev eth1
route add -host 62.77.124.15 dev eth1
为62.77.124.15添加eth1设备


//网络设备
ip link show                查看网络设备
ip link set dev eth1 up     激活
ip link set dev eth1 down   停止

//查看arp表（地址解析协议（ARP），IP地址转换成它对应的物理地址）
ip neighbour 

//更改路由表的默认网关
//注意：这里是指更改网络设备的网关
ip route add default via 192.168.8.1 dev eth0

//查看网络设备
/etc/sysconfig/network-scripts/ifcfg-em2


```

### iptables
```
https://zhuanlan.zhihu.com/p/32848232

//添加iptables规则禁止用户访问域名为www.sexy.com的网站。
iptables -I FORWARD -d www.sexy.com -j DROP

//添加iptables规则禁止IP地址为192.168.1.X的客户机上网。
iptables -I FORWARD -s 192.168.1.X -j DROP

//添加iptables规则禁止192.168.1.0子网里所有的客户机上网。
iptables -I FORWARD -s 192.168.1.0/24 -j DROP

//强制所有的客户机访问192.168.1.x这台Web服务器。
iptables -t nat -I PREROUTING -i eth0 -p tcp –dport 80 -j DNAT –to-destination 192.168.1.x:80

//所有的出口报文都要经过ens32网卡的源地址目标转换
iptables -t nat -A POSTROUTING -o ens32 -j MASQUERADE

//把所有10.8.0.0网段的数据包SNAT成192.168.5.3的ip然后发出去
iptables -t nat -A POSTROUTING -s 10.8.0.0/255.255.255.0 -o eth0 -j SNAT --to-source 192.168.5.3

//从服务器的网卡上，自动获取当前ip地址来做NAT。比如下边的命令：
iptables -t nat -A POSTROUTING -s 10.8.0.0/255.255.255.0 -o eth0 -j MASQUERADE❗
```

### locale
```
locale
LANG=en_US.UTF-8
LC_CTYPE="en_US.UTF-8"               #用户所使用的语言符号及其分类
LC_NUMERIC="en_US.UTF-8"             #数字
LC_TIME="en_US.UTF-8"                 #时间显示格式
LC_COLLATE="en_US.UTF-8"             #比较和排序习惯
LC_MONETARY="en_US.UTF-8"          #LC_MONETARY
LC_MESSAGES="en_US.UTF-8"        #信息主要是提示信息,错误信息, 状态信息, 标题, 标签, 按钮和菜单等
LC_PAPER="en_US.UTF-8"             #默认纸张尺寸大小
LC_NAME="en_US.UTF-8"              #姓名书写方式
LC_ADDRESS="en_US.UTF-8"           #地址书写方式
LC_TELEPHONE="en_US.UTF-8"          #电话号码书写方式
LC_MEASUREMENT="en_US.UTF-8"        #度量衡表达方式
LC_IDENTIFICATION="en_US.UTF-8"     #对自身包含信息的概述
LC_ALL=
```

### 查看文件字符集
```
vim下打开文件，:set fileencoding 可以查看字符集
```

### 更改文件编码方式
```
~/.vimrc 文件中添加以下内容：

set encoding=utf-8 fileencodings=ucs-bom,utf-8,cp936

这样，就可以让vim自动识别文件编码（可以自动识别UTF-8或者GBK编码的文件），其实就是依照 fileencodings提供的编码列表尝试，如果没有找到合适的编码，就用latin-1(ASCII)编码打开。
```

### chown
```
chown 将指定文件的拥有者改为指定的用户或组，用户可以是用户名或者用户ID；组可以是组名或者组ID

//将文件 file1.txt 的拥有者设为 runoob，群体的使用者 runoobgroup
chown runoob:runoobgroup file1.txt

//将目前目录下的所有文件与子目录的拥有者皆设为 runoob，群体的使用者 runoobgroup
chown -R runoob:runoobgroup *
```


### 用户与权限
```
https://zhuanlan.zhihu.com/p/33283263

//用户管理：
useradd,
userdel, usermod, passwd, chsh, chfn, finger, id(查看用户的UID和GID), chage(修改用户密码状态)

//组管理：
groupadd,
groupdel, groupmod, gpasswd

//权限管理
chown, chgrp, chmod, umask

//示例
groupadd gitgroup                   //添加组
useradd malx                        //添加用户
passwd malx                         //设置用户密码
usermod -G gitgroup malx            //为用户分配组
chgrp -R gitgroup /home/data/git/   //修改文件夹的组

//重要文件
/etc/passwd                         //利用 UID 可以找到对应的用户名
/etc/group                          //利用 GID 可以找到对应的群组名
/etc/shadow                         //存储 Linux 系统中用户的密码信息，又称为“影子文件”
```

1、用户与组
```
(1)用户划分
Linux 系统中，UID（用户ID）以如下的方式划分：
管理员：0
普通用户： 1-65535
系统用户：1-499
一般用户：500-60000

(2)用户组类别
私有组：创建用户时，如果没有为其指定所属的组，系统会自动为其创建一个与用户名同名的组
初始组/基本组：用户的默认组
附加组，额外组：默认组以外的其它组
一个用户可以所属多个附加组，但只能有一个初始组。


(3)useradd
useradd [options] USERNAME 
-u UID
-g GID（基本组）
-G GID,... （附加组）
-c "COMMENT" 指定注释信息
-d /path/to/directory 指定家目录
-s SHELL 指定shell路径 /etc/shells 保存有系统中可用使用shell
-m 如果没有家目录则创建 
-k 复制/etc/skl目录下bash配置文件 一般和-m一起使用
-M 不创建用户家目录
-r: 添加系统用户，UID号在1-499
useradd -s /bin/bash -g group –G adm,root user2      //新建了一个用户user2，该用户的登录Shell是/bin/bash，它属于group用户组，同时又属于adm和root用户组，其中group用户组是其主组。

(4)userdel命令 删除用户

(5)id命令 查看用户的帐号属性信息

(6)usermod 修改用户帐号属性
usermod [option] USERNAME
-u 修改UID号
-g 修改GID号
-a -G GID：不使用-a（追加附加组）选项，会覆盖此前的附加组；
-c 指定注释信息
-d -m：-d指定新的家目录 -m 移动以前家目录文件到指定目录中
-s 修改shell
-l
-L：锁定帐号
-U：解锁帐号

(7)groupadd GRPNAME命令 添加组
useradd [options] USERNAME
-g GID
-r：添加为系统组

```

2、权限管理
```
(1)文件权限
r：可读，可以使用类似cat等命令查看文件内容；
w：可写，可以编辑或删除此文件；
x: 可执行，eXacutable，可以命令提示符下当作命令提交给内核运行；

(2)目录权限
r: 可以对此目录执行ls以列出内部的所有文件；
w: 可以在此目录创建文件；
x: 可以使用cd切换进此目录，也可以使用ls -l查看内部文件的详细信息；

(3)所有者 、所属组 、其他人权限
Linux的文件和目录又可以有三个所有者概念: u、g 、o: 所有者、所属组 、其他人
权限之前的表示：- 表示普通文件，l 表示连接文件，b 表示设备文件

(3)chmod
chmod 775 dir 授予dir文件夹/文件权限位 rwx rwx r-x
chmod +x /etc/init.d/mysql   增加可执行权限

(4)chown 修改文件和目录的所有者和所属组
chown [-R] 所有者:所属组 文件或目录

(5)chgrp 修改文件和目录的所属组
chgrp [-R] 所属组 文件名（目录名）

(6)umask默认权限
在 Linux 系统中，文件和目录的最大默认权限是不一样的：
文件：666
目录：777
umask 022 -----w--w-: (-rw-rw-rw-) - (-----w--w-) = (-rw-r--r--)

```


### 端口占有情况
```
ps -aux | grep xxxx
netstat –apn
```

### pmap
```
pmao $$
查看当前进程ip的内存分布
```

### free
```
查看内存使用情况
```

### pstree
```
进程树
```

### getcap setcap
```
capget() 用来获得进程的权能；capset() 用来设置进程权能。
```

### findmnt
```
findmnt -l
查看已经挂载的文件系统
```

### fg
```
使用fg命令带上job id，即可让后端进程组回到前端，

fg 12314
```

### Core Dump
```
当程序运行的过程中异常终止或崩溃，操作系统会将程序当时的内存状态记录下来，保存在一个文件中，这种行为就叫做Core Dump

我们可以认为Core Dump是“内存快照”，但实际上，除了内存信息之外，还有些关键的程序运行状态也会同时dump下来，例如寄存器信息（包括程序指针、栈指针等）、内存管理信息、其他处理器和操作系统状态和信息。

(1) 打开系统core dump
    1) 查系统是否开启core dump
    ulimit -c    //0未开启

    2) 开启系统core dump
    ulimit -c unlimited            //unlimited:表示生成的core文件大小不受限制
    ulimit -c 10240                //超过10240，就不生成了
    注意：开启的core dump只对当前窗口有效

    3) 在~/.bashrc中添加
    ulimit -c unlimited

(2) core路径设置
    1) 默认生成的 core 文件保存在可执行文件所在的目录下，文件名就为 core。

    2) 通过修改 /proc/sys/kernel/core_uses_pid 文件可以让生成 core 文件名是否自动加上 pid 号。
        echo 1 > /proc/sys/kernel/core_uses_pid ，生成的 core 文件名将会变成 core.pid，其中 pid 表示该进程的 PID。

    3) 通过修改 /proc/sys/kernel/core_pattern 来控制生成 core 文件保存的位置以及文件名格式。
        echo "/tmp/corefile-%e-%p-%t" > /proc/sys/kernel/core_pattern 设置生成的 core 文件保存在 “/tmp/corefile” 目录下，文件名格式为 “core-命令名-pid-时间戳”。

(3) gdb调试core文件
    gcc test.c -g
    ./test                  //如果程序出现Segmentation fault (core dumped)，会在当前目录下生成test.core.24158文件
    gdb ./test test.core    //调试


```

### 查看ssh记录
```
who /var/log/wtmp
```

### ssh权限
```
https://www.jianshu.com/p/967e3a04a6c7

(1) 登陆限制
    1) 只要用户满足以下条件之一，ssh便会拒绝登录：
        > 用户无密码(由于默认开启了PermitEmptyPasswords no选项)
        > 用户无合法shell(注意如果不指定shell，则默认为/bin/sh)
        > 明确拒绝使用各种可登录的渠道(比如PasswordAuthentication no, PubkeyAuthentication no等等)
        > /etc/nologin存在，则除root外所有用户均拒绝登录，并打印/etc/nologin文件内容作为提示信息

    2) 禁止root登录
        PermitRootLogin no

    3) Match
        条件是：User, Group, Host, Address
        Match User limited-user
            AllowTcpForwarding yes                  // 这个是默认配置，如果没改过的话可以不加
            X11Forwarding no                        // 禁止x11 forwarding
            GatewayPorts yes                        // 允许ssh -R参数bind所有ip，否则只允许bind 127.0.0.1
            AllowAgentForwarding no
            PasswordAuthentication no                       // 不允许密码登录
            PermitOpen localhost:62222                      // 只允许打开localhost:62222做端口转发
            ForceCommand echo 'This account can only be used for TCP proxy' //此处登陆会直接echo
        > 限定ip白名单登录
            办法很多，比如在防火墙控制，在/etc/hosts.(deny|allow)控制等，其实在/etc/sshd_config
            Match Address 127.0.0.*

(2) 代理权限限制
    1) 禁止端口转发
        AllowTcpForwarding no
    2) 禁止X11转发
        X11Forwarding no
    3) 限制转发端口
        PermitOpen host:port
        PermitOpen IPv4_addr:port
        PermitOpen [IPv6_addr]:port        
        any(默认)表示所有端口都允许用于转发。

(5) sftp限制
    ssh传输文件的有三个命令scp,rsync,sftp。它们的机制不一样，
    scp和rsync是通过远程非交互式执行命令实现的
    sftp是通过openssh的sftp server实现的。
    1) 限制scp和rsync
        Match User xxxx
            ForceCommand /bin/bash

    2) 限制sftp
        > 创建sftp组
            groupadd sftp
        > 创建一个用户sftpuser并分配到sftp组
            useradd -g sftp -s /bin/false sftpuser
        > 配置/etc/ssh/sshd_config
            Subsystem sftp /usr/lib/openssh/sftp-server
            Match Group sftp                //这行用来匹配用户组
                ChrootDirectory /datas/www  //用chroot将用户的根目录指定到/datas/www ，这样用户就只能在/datas/www下活动
                ForceCommand internal-sftp  //强制执行内部sftp，并忽略任何~/.ssh/rc文件中的命令
        > 修改sftp用户组用户目录权限
            > 设定的目录必须是root用户所有，否则就会出现问题
                chown -R root:root /datas/www
                chmod 755 /datas/www
            > 建立SFTP用户登入后可写入的目录：
                mkdir /datas/www/sftpuser
                chown -R sftpuser:sftp /datas/www/sftpuser/
                chmod 755 /datas/www/sftpuser/
```

### systemctl
```
http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html

(1) Unit
    Systemd 可以管理所有系统资源。不同的资源统称为 Unit（单位）
    1) 类型
        Service unit：系统服务
        Target unit：多个 Unit 构成的一个组
        Device Unit：硬件设备
        Mount Unit：文件系统的挂载点
        Automount Unit：自动挂载点
        Path Unit：文件或路径
        Scope Unit：不是由 Systemd 启动的外部进程
        Slice Unit：进程组
        Snapshot Unit：Systemd 快照，可以切回某个快照
        Socket Unit：进程间通信的 socket
        Swap Unit：swap 文件
        Timer Unit：定时器
    2) Unit状态
        //列出正在运行的 Unit
        systemctl list-units
        //显示系统状态
        systemctl status
        //显示单个 Unit 的状态
        sysystemctl status bluetooth.service

    3) Unit管理
        //立即启动一个服务
        sudo systemctl start apache.service
        //立即停止一个服务
        sudo systemctl stop apache.service
        //重启一个服务
        sudo systemctl restart apache.service
        //杀死一个服务的所有子进程
        sudo systemctl kill apache.service
        //重新加载一个服务的配置文件
        sudo systemctl reload apache.service
        //重载所有修改过的配置文件
        sudo systemctl daemon-reload

(2) Unit的配置文件
    Systemd 默认从目录/etc/systemd/system/读取配置文件。但是，里面存放的大部分文件都是符号链接，指向目录/usr/lib/systemd/system/，真正的配置文件存放在那个目录
    1) 配置文件格式
        [Unit]
        Description=test

        [Service]
        Type=simple
        ExecStart=/usr/bin/test
        Restart=always
        RestartSec=10
        StartLimitInterval=0
        //RestartPreventExitStatus=SIGKILL

        [Install]
        WantedBy=multi-user.target

//重新加载一个服务的配置文件
systemctl reload apache.service

//重新加载所有配置文件
systemctl daemon-reload

//启动
systemctl start test

//查看test单元的日志
journalctl -u test

//显示日志占据的硬盘空间
journalctl --disk-usage

//指定日志文件占据的最大空间
journalctl --vacuum-size=1G

//指定日志文件保存多久
journalctl --vacuum-time=1years

```
