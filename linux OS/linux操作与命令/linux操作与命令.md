//------硬盘挂载情况----------------------------------
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


### 动态库操作
```
// 此文件记录了编译时使用的动态库的路径，也就是开机加载so库的路径
    /etc/ld.so.conf 
    默认情况下，编译器只会使用/lib和/usr/lib这两个目录下的库文件，其他目录需添加到此文件中

// ldconfig
    作用是将文件/etc/ld.so.conf列出的路径下的库文件缓存到/etc/ld.so.cache以供使用
    当安装完一些库文件，或者修改/etc/ld.so.conf增加了库的新的搜索路径，需要运行一下ldconfig，使所有的库文件都被缓存
    // 查看当前系统缓存的动态库
    ldconfig -p


// 列出一个程序所需要得动态链接库
    ldd [软件]


// 用于显示二进制目标文件的符号表
    nm /lib64/libc.so.6|grep '\sGLIBC_2' 查看动态库的相关信息

// 打印文件中可打印的字符，用于打印库中的字符
    strings libxxx.so
```

### sh重定向
```
nohup sh start.sh 1>>log.txt 2>&1  &
1>>log.txt 标准输出定向到log.txt文件
2>&1       错误输出定向到标准输出
```
### df
```
查看系统整体磁盘(实体)使用情况
df -h
文件系统        容量  已用  可用 已用% 挂载点
udev            1.9G     0  1.9G    0% /dev
tmpfs           383M  2.1M  381M    1% /run
/dev/sda2       916G   28G  841G    4% /
tmpfs           1.9G     0  1.9G    0% /dev/shm
tmpfs           5.0M  4.0K  5.0M    1% /run/lock
tmpfs           1.9G     0  1.9G    0% /sys/fs/cgroup
/dev/loop3       15M   15M     0  100% /snap/gnome-characters/375
/dev/loop9      1.0M  1.0M     0  100% /snap/gnome-logs/73
/dev/loop7      3.8M  3.8M     0  100% /snap/gnome-system-monitor/123
```

### fdisk
```
// 查看磁盘分区情况
fdisk -l

// 实例
fdisk -l /dev/sda
    Disk /dev/sda：931.5 GiB，1000204886016 字节，1953525168 个扇区
    单元：扇区 / 1 * 512 = 512 字节
    扇区大小(逻辑/物理)：512 字节 / 4096 字节
    I/O 大小(最小/最佳)：4096 字节 / 4096 字节
    磁盘标签类型：gpt
    磁盘标识符：2310ECD2-D44F-401C-B3EA-34C7DBC42982

    (分区情况)
    设备          起点       末尾       扇区  大小 类型
    /dev/sda1     2048    1050623    1048576  512M EFI 系统
    /dev/sda2  1050624 1953523711 1952473088  931G Linux 文件系统
```

### lsblk
```
查看系统的分区和挂载情况
lsblk -f 
NAME            FSTYPE      LABEL UUID                                   MOUNTPOINT
fd0
sda
├─sda1          xfs               651164de-dde3-423f-a3ad-ea066a72ab01   /boot
└─sda2          LVM2_member       l61cFl-1Z5u-2N50-1SUI-Qe0J-348D-EX0ZUs
  ├─centos-root xfs               1be95b37-199b-49bb-94aa-efcbe88fc7da   /
  └─centos-swap swap              374eb40d-0e15-48d2-94bd-bab8bb358b3e   [SWAP]
sr0
```

### 分区
```
// 假设我们创建个虚拟设备，然后再进行分区

1、创建镜像文件
    //查看有哪些设备
    //创建镜像文件 
    dd if=/dev/sda1 of=test.img bs=1M count=100

2、虚拟成loop设备
    losetup -f test.img
    //或
    losetup -f 
        /dev/loop0
    losetup /dev/loop0 test.img

3、设备分区
    fdisk /dev/loop0

4、刷新设备
    partprobe /dev/loop0

5、查看分区情况
    lsblk -f /dev/loop0
    loop0
      └─loop0p1

6、格式化分区
    mkfs.ext4 /dev/loop0p1

7、挂载
    mount /dev/loop0p1 /mnt/aaa

8、卸载
    umount /dev/loop0p1

9、删除分区
    fdisk /dev/loop0
    d

10、删除loop0设备
    losetup -d /dev/loop0
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
硬盘、光盘、软盘等都是常见的块设备，他们在Linux下的目录一般是/dev/hda1, /dev/cdrom, /dev/sda1，/dev/fd0这样的。而loop device是虚拟的块设备，主要目的是让用户可以像访问上述块设备那样访问一个文件。 loop device设备的路径一般是/dev/loop0, dev/loop1...

                     虚拟成                         格式化特定文件系统          挂载文件夹  
普通文件 -> 镜像文件 --------> （块）设备(/dev/loop) -------------------> ext4 ------------> /home/test
           镜像文件 --------------------------------------------------> ext4 ------------> /home/test  

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

    //隐藏进程
    mkdir /tmp/none
    mount --bind /tmp/none /proc/pid
    //查看隐藏进程
    cat /proc/mounts
    或
    cat /proc/$$/mountinfo

(6) 挂载nfs
    https://help.ubuntu.com/lts/serverguide/network-file-system.html
    
    mount -t nfs hostname:/directory /mount/point
    //注意：nfs只是个服务，不是个设备

(7) 挂载多个分区（设备）到一个文件夹

```

### LVM
```
// 对于上面建立的设备/dev/loop0及其分区/dev/loop0p1、/dev/loop0p2、/dev/loop0p3

1、建立物理卷PV
    pvcreate /dev/loop0p1
    pvcreate /dev/loop0p2
    pvcreate /dev/loop0p3
    // 查看
    pvdisplay
    // 删除
    pvremove 

2、将物理卷合成卷组VG
    vgcreate volume-group1 /dev/loop0p1 /dev/loop0p2 /dev/loop0p3
    // 查看
    vgdisplay
    // 删除
    vgremove volume-group1
    // 扩展卷组
    vgextend volume-group1 /dev/loop0p4

3、创建逻辑卷(卷组进行分区)
    lvcreate -L 100M -n lv1 volume-group1
    // 查看
    lvdispaly
    // 删除
    lvremove /dev/volume-group1/lv1
    // 格式化、挂载
    mkfs.ext4 /dev/volume-group1/lv1
    mount /dev/volume-group1/lv1 /mnt/test

4、扩展逻辑卷
    umount /dev/volume-group1/lv1
    // 扩容100M
    lvextend -L +100M /dev/volume-group1/lv1
    // 减少50M
    lvreduce -L -50M /dev/volume-group1/lv1
    // 检查错误
    e2fsck -f /dev/volume-group1/lv1
    // 重新加载逻辑卷才能生效
    resize2fs /dev/volume-group1/lv1
    // 然后挂载
    // 查看挂载后的目录大小
    df -h 
    文件系统                        容量  已用  可用 已用% 挂载点
    /dev/mapper/volume--group1-lv1  144M  1.6M  132M    2% /mnt/aaa

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
1、显示网络设备为eth1的信息
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
                    本地                        远端
tcp        0      0 127.0.0.1:8112              0.0.0.0:*                   LISTEN      10552/ssh           
tcp        0      0 20.23.11.12:54923           10.10.53.54:22              ESTABLISHED 10552/ssh    
```

### pmap
```
pmao $$
查看当前进程ip的内存分布
```

### free
```
查看内存使用情况
              total        used        free      shared  buff/cache   available
Mem:           3.7G        879M        453M         46M        2.4G        2.6G
Swap:          2.0G         57M        1.9G

Swap:
    Linux内核为了提高读写效率与速度，会将文件在内存中进行缓存，这部分内存就是Cache Memory(缓存内存)。即使你的程序运行结束后，Cache Memory也不会自动释放。这就会导致你在Linux系统中程序频繁读写文件后，你会发现可用物理内存变少。当系统的物理内存不够用的时候，就需要将物理内存中的一部分空间释放出来，以供当前运行的程序使用。那些被释放的空间可能来自一些很长时间没有什么操作的程序，这些被释放的空间被临时保存到Swap空间中，等到那些程序要运行时，再从Swap分区中恢复保存的数据到内存中。这样，系统总是在物理内存不够时，才进行Swap交换。
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

### hwclock
```
https://yq.aliyun.com/articles/653670

hwclock是一种访问硬件时钟的工具，可以显示当前时间，将硬件时钟设置为指定的时间，将硬件时钟设置为系统时间，以及从硬件时钟设置系统时间

//读取硬件时钟并在标准输出上打印时间
hwclock -r
```

### 资源管理
```
(1) 查看消耗CPU资源最多的前10个进程
    ps auxw | head -1;ps auxw |sort -rn -k3 |head -11

(2) 查看消耗内存资源最多的前10个进程
    ps auxw | head -1;ps auxw |sort -rn -k4 |head -11

(3) 如下指令也是同样效果(按照cpu,内存等)
    ps auxw --sort=rss
    ps auxw --sort=%cpu
    ps auxw --sort=%mem
    > %MEM 进程的内存占用率
    > VSZ 进程所使用的虚存的大小
    > RSS 进程使用的驻留集大小或者是实际内存的大小
    > TTY 与进程关联的终端（tty）
```

### 域名解析
```
nslookup https://music.163.com/
```

### make
```
(1) make指定目录
    1> ./configure --prefix=* && make
    2> ./configure && make install DESTDIR=*

```

### file
```
file xxx
查看文件类型

//例如
file a.out
a.out: ELF 64-bit LSB shared object, x86-64, version 1 (GNU/Linux), dynamically linked, interpreter /lib64/l, for GNU/Linux 3.2.0, BuildID[sha1]=1f968f44e3ed35b7725aff487cc1e9eb1ebf9a5a, not stripped

```

### curl
```
(1) 模拟Http Get/Post请求
    // get
    curl http://127.0.0.1:8080/check_your_status?user=Summer&passwd=12345678
    // post
    curl -d "user=Summer&passwd=12345678" "http://127.0.0.1:8080/check_your_status"
    // 带header
    curl -H "Content-Type:application/json" -H "Accept: application/json" -X POST --data '{"message": "sunshine"}' http://localhost:8000/


```

### date
```
//unix时间戳
date +%s

//将固定时间转换为unix时间戳
date -d '2013-2-22 22:14:23.460630412"' +%s

//时间戳转换时间
date -d @1361542596

//时间戳转换时间 指定格式
date -d @1361542596 +"%Y-%m-%d %H:%M:%S"

//UTC 标准时间的日期
date +"%Y-%m-%d"
```

### apt
```
apt提供了大多数与apt-get及apt-cache有的功能，但更方便使用

//列出系统包含的软件和库
dpkg list
//搜索软件包描述
apt-cache search [软件]
//显示软件包细节 
apt-cache show [软件]
//列出软件包含了哪些文件
dpkg -L [软件] 
//安装软件包 
apt-get install [软件]
//移除软件包 
apt-get remove [软件]
//更新可用软件包列表 
apt-get update
//通过 安装/升级 软件来更新系统 
apt-get upgrade
//通过 卸载/安装/升级 来更新系统 
apt-get dist-upgrade
//编辑软件源信息文件 
vim /etc/apt/sources.list

apt 和 apt-get 命令都是基于 dpkg

// 查看已安装的软件
apt list --installed 


```

### 禁止端口转发和ip转发
```
/etc/ssh/sshd_config
    AllowTcpForwarding no           禁止端口转发


/etc/sysctl.conf
    net.ipv4.ip_forward = 0         禁止ip转发


```

### 系统数据文件
1、/etc/password
```
存放着所有用户帐号的信息，包括用户名和密码，因此，它对系统来说是至关重要的
格式如下：
username:password:User ID:Group ID:comment:home directory:shell
```
2、/etc/shadow
```
存放系统的口令文件
```
3、/etc/group
```
用户组管理的文件,linux用户组的所有信息都存放在此文件中
格式如下：
组名:口令:组标识号:组内用户列表
```
4、/etc/hosts
```
Linux系统中一个负责IP地址与域名快速解析的文件
例如：
127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4

此目录包含
host.conf    hostname     hosts        hosts.allow  hosts.deny
```
5、/etc/services
```
Internet 守护程序（ineted）是 Linux 世界中的重要服务。它借助 /etc/services 文件来处理所有网络服务
格式如下：
service-name    port/protocol   [aliases..]  [#comment]
service-name 是网络服务的名称。例如 telnet、ftp 等。
port/protocol 是网络服务使用的端口（一个数值）和服务通信使用的协议（TCP/UDP）。
alias 是服务的别名。
comment 是你可以添加到服务的注释或说明。以 # 标记开头。
```
6、utmp和wtmp
```
/var/run/utmp（二进制） 命令 who
/var/log/wtmp（二进制） 命令 w
utmp记录当前登录到系统的用户
wtmp跟踪各个登陆与注销事件
```
7、uname和hostname
```
命令uname显示操作系统信息
命令hostname显示主机的域名
```


### /var 目录
```
系统一般运行时要改变的数据.每个系统是特定的，即不通过网络与其他计算机共享.  
/var/lib            系统正常运行时要改变的文件.

/var/local          /usr/local 中安装的程序的可变数据

/var/lock           锁定文件.以支持他们正在使用某个特定的设备或文件.其他程序注意到这个锁定文件，将不试图使用这个设备或文件.  

/var/log            各种程序的Log文件，
/var/log/wtmp       永久记录每个用户登录、注销及系统的启动、停机的事件。
/var/log/lastlog    记录最近成功登录的事件和最后一次不成功的登录事件，由login生成

/var/run            保存到下次引导前有效的关于系统的信息文件
/var/run/utmp       记录著現在登入的用戶。

/var/tmp            比/tmp 允许的大或需要存在较长时间的临时文件
```

### stdout stderr
```
stdout是行缓冲的，他的输出会放在一个buffer里面，只有到换行的时候，才会输出到屏幕。而stderr是无缓冲的，会直接输出
如果用转向标准输出到磁盘文件，则可看出两者区别。stdout输出到磁盘文件，stderr在屏幕。 

```


### /proc文件夹
```
/proc 文件系统是一种内核和内核模块用来向进程(process) 发送信息的机制, /proc 存在于内存之中而不是硬盘上。proc文件系统以文件的形式向用户空间提供了访问接口，这些接口可以用于在运行时获取相关部件的信息或者修改部件的行为，因而它是非常方便的一个接口。

(1) 内容介绍
    /proc/cpuinfo - CPU 的信息(型号, 家族, 缓存大小等)
    /proc/meminfo - 物理内存、交换空间等的信息
    /proc/mounts - 已加载的文件系统的列表
    /proc/devices - 可用设备的列表
    /proc/filesystems - 被支持的文件系统
    /proc/modules - 已加载的模块
    /proc/version - 内核版本
    /proc/cmdline - 系统启动时输入的内核命令行参数

    /proc/pid/*     pid进程的相关信息

    /proc/sys/kernel    与内核相关
```

### /var/run
```
/var/run 目录中存放的是自系统启动以来描述系统信息的文件。比较常见的用途是daemon进程将自己的pid保存到这个目录。标准要求这个文件夹中的文件必须是在系统启动的时候清空，以便建立新的文件。

(1) /var/run/*.pid
在工作中遇到了很多在程序启动时检查是否已经重复启动的代码段，其核心就是调用fcntl设置pid文件的锁定F_SETLK状态，其中锁定的标志为F_WRLACK。如果成功锁定，则写入进程当前PID，进程继续往下执行。如果锁定不成功，说明已经有同样的进程在运行了，当前进程结束退出

(2) 
```

### /opt
```
/opt目录用来安装附加软件包，是用户级的程序目录，
```

### /var/log
```
/var/log/boot.log：录了系统在引导过程中发生的事件，就是Linux系统开机自检过程显示的信息

/var/log/lastlog ：记录最后一次用户成功登陆的时间、登陆IP等信息

/var/log/messages ：记录Linux操作系统常见的系统和服务错误信息

/var/log/secure ：Linux系统安全日志，记录用户和工作组变坏情况、用户登陆认证情况

/var/log/btmp ：记录Linux登陆失败的用户、时间以及远程IP地址

/var/log/syslog：只记录警告信息，常常是系统出问题的信息，使用lastlog查看

/var/log/wtmp：该日志文件永久记录每个用户登录、注销及系统的启动、停机的事件，使用last命令查看

/var/run/utmp：该日志文件记录有关当前登录的每个用户的信息。如 who、w、users、finger等就需要访问这个文件
```

### losetup
```
//查找第一个未使用的回环设备
losetup -f 
/dev/loop0

//显示所有已经使用的回环设备状态
losetup -a 
```

### loop
```
UNIX 系统里，loop 设备是一种伪设备(pseudo-device)，或者也可以说是仿真设备。它能使我们像块设备一样访问一个文件

loop device设备是通过影射操作系统上的正常的文件而形成的虚拟块设备。因为这种设备的存在，就为我们提供了一种创建一个存在于其他文件中的虚拟文件系统的机制

// 使用dd命令创建文件（镜像文件）
dd if=/dev/zero of=FS_on_file bs=1k count=10000
      这是个设备  


// 使用losetup命令将镜像文件虚拟成块设备
losetup /dev/loop0 FS_on_file

// 创建一个文件系统
mkfs -t ext3 /dev/loop0

// 挂载这个文件系统
mkdir /mnt/FS_file0
mount /dev/loop0 /mnt/FS_file0

// 若删除刚才创建的这些对象，依次执行如下步骤：
umount /dev/loop0
losetup -d /dev/loop0
rm FS_on_file
```

### http_proxy
```
// 设置 socks 代理，自动识别socks版本
    export http_proxy=socks://10.0.0.52:1080
// 设置 socks4 代理
    export http_proxy=socks4://10.0.0.52:1080
// 设置 socks5 代理
    export http_proxy=socks5://10.0.0.52:1080
// 代理使用用户名密码认证：
    export http_proxy=user:pass@192.158.8.8:8080
  　　
//  如果需要为https网站设置代理，设置https_proxy环境变量即可；设置方法完全与http_proxy环境变量相同，例如：
    export https_proxy=socks5://10.0.0.52:1080

```

### tcpdump
```
抓包工具

1、关键字
    (1) 类型
        host    host 210.27.48.2        //指明ip地址
        net     net http://www.aa.com   //指明net地址
        port    port 23                 //指明端口号
    (2) 传输方向
        src 源地址，dst 目的地址
        src host 12.12.1.2              //源IP地址是12.12.1.2
        dst host 12.12.1.2              //目的IP地址是12.12.1.2
    (3) 协议
        ip, tcp, udp...

2、实例
    // 抓取所有经过eth1，'源地址'是192.168.1.1的网络数据
    tcpdump -i eth1 src host 192.168.1.1

    // 抓取所有经过eth1，'目的地址'是192.168.1.1的网络数据
    tcpdump -i eth1 dst host 192.168.1.1

    // 抓取所有经过eth1，源端口是25的网络数据
    tcpdump -i eth1 src port 25

    // 抓取所有经过eth1，目的端口是25的网络数据
    tcpdump -i eth1 dst port 25

    // 抓取所有经过eth1的所有tcp数据
    tcpdump -i eth1 tcp

    // 抓取所有经过eth1，目的地址是192.168.1.254或192.168.1.200端口是80的TCP数据
    tcpdump -i eth1 '((tcp) and (port 80) and ((dst host 192.168.1.254) or (dst host 192.168.1.200)))'

    // 抓取所有经过eth1，目标MAC地址是00:01:02:03:04:05的ICMP数据
    tcpdump -i eth1 '((icmp) and ((ether dst host 00:01:02:03:04:05)))'

    // 抓取所有经过eth1，目的网络是192.168，但目的主机不是192.168.1.200的TCP数据
    tcpdump -i eth1 '((tcp) and ((dst net 192.168) and (not dst host 192.168.1.200)))'

```