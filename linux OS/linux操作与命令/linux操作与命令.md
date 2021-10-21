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

rpm/yum适用于Redhat、CentOS、Suse等平台
apt-get/dpkg适用于Debian、Ubuntu等平台

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
    x86:    ldd xxx
    arm:    readelf -d xxx

// 查看动态库平台
readelf -h xxx.so

// 用于显示二进制目标文件的符号表
    nm /lib64/libc.so.6|grep '\sGLIBC_2' 查看动态库的相关信息

// 打印文件中可打印的字符，用于打印库中的字符
    strings libxxx.so
```

### not a dynamic executable
```
ldd arm-himix200-linux-gcc
// not a dynamic executable

readelf -h arm-himix200-linux-gcc   // 查看平台
readelf -l arm-himix200-linux-gcc   // 查看缺失的库(ldd)

apt-get install gcc-multilib        // 下载全平台libc库
```

### sh重定向
```
nohup sh start.sh 1>>log.txt 2>&1  &
1>>log.txt 标准输出定向到log.txt文件
2>&1       错误输出定向到标准输出
```
### df
```
查看挂载的文件系统使用情况
df -h
文件系统        容量  已用  可用 已用% 挂载点
udev            3.9G     0  3.9G    0% /dev
tmpfs           785M  2.1M  783M    1% /run
/dev/sda2       916G   35G  835G    4% /
tmpfs           3.9G   15M  3.9G    1% /dev/shm
tmpfs           5.0M  4.0K  5.0M    1% /run/lock
tmpfs           3.9G     0  3.9G    0% /sys/fs/cgroup
/dev/loop2      141M  141M     0  100% /snap/gnome-3-26-1604/97
/dev/loop6      2.5M  2.5M     0  100% /snap/gnome-calculator/730
/dev/loop3       94M   94M     0  100% /snap/core/8935
/dev/loop5       94M   94M     0  100% /snap/core/9066
/dev/loop8       15M   15M     0  100% /snap/gnome-characters/495
/dev/loop1       63M   63M     0  100% /snap/gtk-common-themes/1506
/dev/loop7      1.0M  1.0M     0  100% /snap/gnome-logs/93
/dev/loop9       55M   55M     0  100% /snap/core18/1705
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
    // 方法一
    losetup -f test.img             // 直接分配loop设备
    // 方法二
    losetup -f                      // 查看可用loop设备
    losetup /dev/loop0 test.img     // 将镜像分配给指定loop设备
    // 查看已分配的loop设备
    losetup

3、设备分区
    fdisk /dev/loop0                // 交互式分区

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
普通文件或/dev/loop或/dev/zero -> 镜像文件 --------> （块）设备(/dev/loop) -------------------> ext4 ------------> /home/test
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
    // ubuntu 安装 nfs
    // server
    apt-get install nfs-kernel-server
    vi /etc/exports     // 添加目录与权限
        /home/malx/nfs *(rw,sync,no_root_squash,no_subtree_check)
    service nfs start   // 启动nfs服务

    // client
    apt-get install nfs-common
    mount -t nfs 10.151.3.77:/home/malx/nfs /mount/
    mount -t nfs -o nolock,tcp 10.151.3.77:/home/malx/nfs /home


(7) 挂载多个分区（设备）到一个文件夹

```

### /etc/fstab
```
挂载的配置文件

// column
设备    挂载点  文件系统类型    挂载选项    转储频度    自检次序
/dev/xxx    /mnt    ext4    ...
10.151.3.77:/home/malx/nfs  /nfs    nfs
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

### Device mapper
```
/dev/mapper

Device mapper是Linux2.6内核中提供的一种从逻辑设备到物理设备的映射机制
在该机制下，用户能够很方便的根据自己的需要实现对存储资源的管理
```

### 查看设备
```
//查看cpu信息
lscpu
/proc/cpuinfo

//内存
dmidecode
/proc/meminfo

//查看总线上的所有设备
lspci
lspci -v    // 详细信息
cat /proc/pci

//查看网络设备的详细信息
iwconfig

//查看usb的设备
lsusb

```

### ifconfig
```
用于显示或设置网络设备

// 查看所有网络设备
ifconfig

// 启动关闭指定网卡
ifconfig eth0 down
ifconfig eth0 up

// 为网卡配置和删除IPv6地址
ifconfig eth0 add 33ffe:3240:800:1005::2/ 64 //为网卡设置IPv6地址
ifconfig eth0 del 33ffe:3240:800:1005::2/ 64 //为网卡删除IPv6地址

// 用ifconfig修改MAC地址
ifconfig eth0 down //关闭网卡
ifconfig eth0 hw ether 00:AA:BB:CC:DD:EE //修改MAC地址
ifconfig eth0 up //启动网卡

// 开机启动时，如果没有赋予其mac地址，那么会随机设定

// 配置IP地址
ifconfig eth0 192.168.1.56 
// 给eth0网卡配置IP地址,并加上子掩码
ifconfig eth0 192.168.1.56 netmask 255.255.255.0 
// 给eth0网卡配置IP地址,加上子掩码,加上个广播地址
ifconfig eth0 192.168.1.56 netmask 255.255.255.0 broadcast 192.168.1.255

// 启用和关闭ARP协议
ifconfig eth0 arp  //开启
ifconfig eth0 -arp  //关闭

// 设置最大传输单元
ifconfig eth0 mtu 1500 //设置能通过的最大数据包大小为 1500 bytes

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
// 为dev添加网关
ip route add default via 192.168.8.1 dev eth0

// 为dev配置网关
route add default gw 192.168.8.1 dev eth0

//查看网络设备
/etc/sysconfig/network-scripts/ifcfg-em2

// 删除路由
ip route delete ...

```

### route
```
route  [add|del] [-net|-host] target [netmask Nm] [gw Gw] [[dev] If]

add: 添加一条路由规则
del: 删除一条路由规则
-net: 目的地址是一个网络
-host: 目的地址是一台主机
target: 目的网络或主机
netmask: 目的地址的网络掩码
gw: 路由数据包通过的网关
dev If: 强制路由关联到指定的设备接口，否则的话内核会其自身的相应规则决定选用那个设备接口。在大多数正常的网络中你可能并不需要指定本项。假如dev If是命令行的最后一个选项的话，那么关键字dev可以省略

// 例子

```

### iptables
```
https://zhuanlan.zhihu.com/p/32848232
https://wangchujiang.com/linux-command/c/iptables.html

// 规则链 -I
INPUT链 ：处理输入数据包。
OUTPUT链 ：处理输出数据包。
FORWARD链 ：处理转发数据包。
PREROUTING链 ：用于目标地址转换（DNAT）
POSTROUTING链 ：用于源地址转换（SNAT）

// 动作 -j
ACCEPT ：接收数据包
DROP ：丢弃数据包
REDIRECT ：重定向、映射、透明代理
SNAT ：源地址转换
DNAT ：目标地址转换
MASQUERADE ：IP伪装（NAT），用于ADSL
LOG ：日志记录
SEMARK : 添加SEMARK标记以供网域内强制访问控制（MAC）

// 查看规则
iptables -L

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

// Swap
Linux内核为了提高读写效率与速度，会将文件在内存中进行缓存，这部分内存就是Cache Memory(缓存内存)。即使你的程序运行结束后，Cache Memory也不会自动释放。这就会导致你在Linux系统中程序频繁读写文件后，你会发现可用物理内存变少
当系统的物理内存不够用的时候，就需要将物理内存中的一部分空间释放出来，以供当前运行的程序使用
那些被释放的空间可能来自一些很长时间没有什么操作的程序，这些被释放的空间被临时保存到Swap空间中，等到那些程序要运行时，再从Swap分区中恢复保存的数据到内存中。这样，系统总是在物理内存不够时，才进行Swap交换

// buffer/cache
    buffer cache 和 page cache(相见资源管理)
```

### 回收cache
```
https://yq.aliyun.com/articles/87126

// 为何要回收cache
Linux 内核会在内存将要耗尽的时候，触发内存回收的工作，以便释放出内存给急需内存的进程使用。一般情况下，这个操作中主要的内存释放都来自于对 buffer/cache 的释放

// 成本
伴随着 cache 清除的行为的，一般都是系统 IO 飙高，涉及cache数据的写回

// 人工清除
> 清除 page cache
    echo 1 > /proc/sys/vm/drop_caches
> 清除回收 slab 分配器中的对象
    echo 2 > /proc/sys/vm/drop_caches
> 清除 page cache 和 slab 分配器中的缓存对象
    echo 3 > /proc/sys/vm/drop_caches
    
// cache都能被回收么
Linux 系统内存中的 cache 并不是在所有情况下都能被释放当做空闲空间用的
> tmpfs
    tmpfs 中存储的文件会占用 cache 空间，除非文件删除否则这个 cache 不会被自动释放
> 共享内存
    shmget 方式申请的共享内存会占用 cache 空间，除非主动释放，否则相关的 cache 空间都不会被自动释放
> mmap
    mmap 方法申请的 MAP_SHARED 标志的内存会占用 cache 空间，除非进程将这段内存 munmap，否则相关的 cache 空间都不会被自动释放
实际上 shmget、mmap 的共享内存，在内核层都是通过 tmpfs 实现的，而 tmpfs 实现的存储用的都是 cache
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


(4) gdb直接调试
    gdb a.out
    run  运行
    bt   打印错误
    f n  进入函数域
    p *** 打印当前变量

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


// 注意
wsl docker 中对 systemctl 不友好，尽量用 service
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
dpkg --list
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
dpkg-query -l

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
    /proc/cpuinfo       CPU 的信息(型号, 家族, 缓存大小等)
    /proc/meminfo       物理内存、交换空间等的信息
    /proc/mounts        已加载的文件系统的列表
    /proc/devices       查看主设备号
    /proc/filesystems   被支持的文件系统
    /proc/modules       已加载的模块
    /proc/version       内核版本
    /proc/cmdline       系统启动时输入的内核命令行参数
    /proc/pid/*         pid进程的相关信息
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

/var/log/syslog：只记录警告信息，常常是系统出问题的信息

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

### dd
```
// 向磁盘上写一个大文件, 来看写性能
dd if=/dev/zero bs=1024 count=1000000 of=/root/1Gb.file
 
// 从磁盘上读取一个大文件, 来看读性能
dd if=/root/1Gb.file bs=64k | dd of=/dev/null
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

### dmesg
```
// 查看linux内核的输出信息
dmesg
```

### mknod
```
/dev/目录下有许多设备节点文件，比如u盘的文件/dev/sda，mmc卡的文件/dev/mmcblk0，这些文件通常是由udev或mdev程序检测到uevent事件后自动创建的。我们也可以通过mknod命令手动创建。
```

### 切换显卡
```
prime-select
```

### xrandr
```
https://www.jianshu.com/p/4c37e52632da

// 列出系统支持的视频接口名称和设备连接情况
    xrandr

// 将视频输出发送到某个接口设备
    xrandr --output DFP1 --auto

// 设置分辨率时需要指定设置的output及mode，如将LVDS-1的分辨率改为1920×1080
    xrandr --output eDP1 --mode 1920x1080 
    --output:指定显示器。 
    --mode:指定一种有效的分辨率。 
    --rate:指定刷新率。

// 关闭某个接口设备的视频输出
    xrandr --output LVDS --off

// 设置双屏幕显示
// 打开外接显示器，双屏幕显示相同的内容--克隆，（auto为最高分辨率）
    xrandr --output VGA-0 --same-as DVI-D-0 --auto

// 若要指定外接显示器的分辨率可以使用下面的命令（1280*1024）
    xrandr --output VGA-0 --same-as DVI-D-0 --mode 1280x1024

//  调节显示器亮度
    xrandr --output HDMI-1-1 --brightness 0.8

```

### inotify
```
使用 inotify 监控 Linux 文件系统事件

有inotify相关API

估计前端实时编译更新，可能用的就是这个
```

### sysctl
```
sysctl命令被用于在内核运行时动态地修改内核的运行参数，可用的内核参数在目录/proc/sys中

它包含一些TCP/ip堆栈和虚拟内存系统的高级选项， 这可以让有经验的管理员提高引人注目的系统性能

用sysctl可以读取设置超过五百个系统变量


sysctl --write fs.inotify.max_user_watches='81920'

永久保留配置，修改/etc/sysctl.conf文件


```

### readlink
```
readlink是linux系统中一个常用工具，主要用来找出符号链接所指向的位置

// 查看系统中的awk命令到底是执行哪个可以执行文件
readlink /usr/bin/awk
/etc/alternatives/awk(其实这个还是一个符号连接)

readlink /etc/alternatives/awk
/usr/bin/gawk(这个才是真正的可执行文件)
```

### modinfo
```
查看驱动信息
lspci -v        // 先获得目标驱动name
modinfo ath9k   // 查看无线网卡驱动ath9k的详细信息
```

### Argument list too long
```
// 例如cp拷贝大量小文件会出现如题的错误，解决办法如下
find resource/ -name "*.jpg" | xargs -i cp -rf {} ./
```

### 内存文件系统
```
与其他文件系统不同，tmpfs无需要建立或格式化，需要直接mount

// 挂载内存文件系统(默认它是系统内存的一半)
mount -t tmpfs tmpfs /mnt/tmp


// 定义大小
mount -t tmpfs -o size=1G tmpfs /mnt/mytmpfs 

// 也可以在挂载后，重新挂载(remount) tmpfs 即可改变内存上限
mount -o remount,size=512m/mnt/tmp
```

### 配置域名服务器
```
// 注意区分 /etc/hosts
// 域名如果在hosts中找不到对应的IP，会访问此文件寻找域名解析服务器

vim /etc/resolv.conf
nameserver 10.151.6.6
nameserver 8.8.8.8
```

### tar  压缩 忽略目录结构
```
tar zcvf xxx.tar.gz -C 目录 压缩的目标目录

/dir1/dir2/xxx

tar zcvf xxx.tar.gz -C /dir1/dir2 xxx
```

### tar 压缩 选择目录结构
```
// 选择目录dir下aa和bb
tar zcvf aaa.tar.gz dir/aa dir/bb
```

### 更改系统时间
```
date +%Y-%m-%d -s "2021-04-01"
date +%T -s "11:14:00"
```

### ntpd服务的相关设置文件
```
ntpdate -u 10.151.3.74
同步ntp时间

/etc/ntp.conf
这个是NTP daemon的主要设文件，也是 NTP 唯一的设定文件

/usr/share/zoneinfo/
在这个目录下的文件其实是规定了各主要时区的时间设定文件，例如北京地区的时区设定文件在/usr/share/zoneinfo/Asia/Beijing 就是了。这个目录里面的文件与底下要谈的两个文件(clock 与localtime)是有关系的

/etc/sysconfig/clock
这个文件其实也不包含在NTP 的 daemon 当中，因为这个是linux的主要时区设定文件。每次开机后，Linux 会自动的读取这个文件来设定自己系统所默认要显示的时间

/etc/localtime
这个文件就是“本地端的时间配置文件”。刚刚那个clock 文件里面规定了使用的时间设置文件(ZONE) 为/usr/share/zoneinfo/Asia/Beijing ，所以说，这就是本地端的时间了，此时， Linux系统就会将Beijing那个文件另存为一份/etc/localtime文件，所以未来我们的时间显示就会以Beijing那个时间设定文件为准

/etc/timezone
系统时区文件
```

### 查看发行版的版本
```
cat /etc/os-release
cat /etc/issue
```

### 自启动
```
/etc/init.d
/etc/rc[0-9].d

(1) init.d vs init
    /etc/init.d/就是旧时代linux的用法
    /etc/init/是现在Ubuntu的提倡并一步步转型的用法
    为了平缓过渡，便让service命令可以同时寻找到两个文件夹

(2) 启动服务
    上述两个文件夹的区别也就是服务启动方式的区别，目前有三种启动方式(mysql为例)
    [1] 只从/etc/init.d/文件夹启动
        /etc/init.d/mysql start
    [2] 只从/etc/init/文件夹启动(Ubuntu提倡)
        sudo start mysql
    [3] 从两个文件夹中启动
        service start mysql

(3) /etc/rc[0-9].d
    /etc/rc[0-9].d文件夹下软链接的是/etc/init.d/中的脚本
    /etc/rc2.d/S01binfmt-support  /etc/rc4.d/S01dbus
    /etc/rc2.d/S01dbus            /etc/rc4.d/S01ssh
    /etc/rc2.d/S01ssh             /etc/rc5.d/S01binfmt-support
    S:开机自启动
    K:开机自kill
```

### dash bash
```
dash: 小巧，但不友好
bash: 庞大，但很友好

// 查看
ls -l /bin/sh

// 由dash切换成bash
sudo dpkg-reconfigure dash
```

### 查看系统标准头文件路径
```
cpp -v
```

### 查看当前系统支持的文件系统
```
目录/proc/filesystems提供的内容不准确

文件系统属于linux内核模块，只需查看当前已加载的模块即可
cat /proc/modules
```

### lsof
```
https://linuxtools-rst.readthedocs.io/zh_CN/latest/tool/lsof.html

lsof是查看文件相关的信息
可以打开普通文件、目录、网络文件系统的文件、字符或设备文件、(函数)共享库、管道，命名管道、符号链接、网络文件等

// 查看哪些进程打开了某个文件
lsof /bin/bash

// 列出某个用户打开的文件信息
lsof -u malx

// 列出某个程序进程所打开的文件信息
lsof -c mysql

// 列出某个用户以及某个进程所打开的文件信息
lsof -c mysql -u malx

// 通过某个进程号显示该进程打开的文件
lsof -p 11968

// 列出所有的网络连接
lsof -i

// 列出所有tcp 网络连接信息
lsof -i tcp
lsof -n -i tcp

// 列出谁在使用某个端口
lsof -i :3306

// 列出某个用户的所有活跃的网络端口
lsof -a -u test -i

// 根据文件描述列出对应的文件信息
lsof -d 3

// 列出被进程号为1234的进程所打开的所有IPV4 network files
lsof -i 4 -a -p 1234

// 列出目前连接主机nf5260i5-td上端口为：20，21，80相关的所有文件信息，且每隔3秒重复执行
lsof -i @nf5260i5-td:20,21,80 -r 3

```

### LUKS
```
https://blog.betamao.me/2019/10/27/LUKS%E5%85%A8%E7%9B%98%E4%BF%9D%E6%8A%A4%E5%88%86%E6%9E%90/
https://zhuanlan.zhihu.com/p/36870751

LUKS: Linux硬盘加密提供了一种标准

特点: 必须对加密的卷进行解密，才能挂载其中的文件系统

cryptsetup提供了LUKS磁盘加密的工具

// 步骤
(1) 创建分区并加密分区
    > dd if=/dev/zero of=test.img bs=1M count=10
    > cryptsetup -v -y luksFormat test
        Are you sure? (Type uppercase yes): YES
        Enter passphrase for test.img:
        Verify passphrase:

(2) 映射分区
    cryptset luksOpen test.img TEST
    // 会映射到 /dev/mapper/TEST

(3) 格式化分区并挂载使用
    mkfs.ext4 /dev/mapper/TEST
    mount /dev/mapper/TEST /mnt/test // 需要输入密码才能挂载

(4) 关闭映射分区
    umount /mnt/test
    cryptsetup luksClose TEST


// 注意
    WSL和Docker环境中使用luks均有问题！！！
```

### e4crypt(ext4 cryptfs) vs eCryptfs vs dm-crypt(LUKS)
```
dm-crypt            块级别加密
e4crypt eCryptfs    文件级别加密

读写性能 dm-crypt > e4crypt > eCryptfs
```

### graphviz
```dot
digraph base_flow {
    // 步骤1： 定义digraph的属性
    label = <<B>graphviz使用流程</B>>
    size = 10; // 图大小
    // bgcolor	= "背景颜色"
    // fontcolor = "字体颜色"
    // fontname = "字体"
    fontsize = 22
    
    // 步骤2： 定义node、edge的属性
    node[shape=box fontsize = 10]
    edge[arrowsize=0.5 color="red" fontsize=22 fontcolor=grey]

    // 步骤3： 定义node、subgraph
    A[label="hello"]
    B[label="world"]
    C[label="me"]

    subgraph cluster_flow1 { // 子图 必须为cluster_*
        // 在subgraph中也可以定义各种属性
        label = ""
        node[color="grey"]
        aaa[label=aaa]
        bbb[label=bbb shape=ellipse]
        ccc[label=ccc shape=ellipse]
        aaa -> { bbb ccc };
    }

    // 步骤4： 添加关系
    A -> B
    B -> C [label=" do you like me?" fontsize=10 color=green dir=none]
    B -> aaa [label=" fuck" dir=both color=blue]
    { rank=same B,C}

    subgraph cluster_flow2 {
        label=""
        bgcolor="beige"
        node [shape="record"]
        ca [label="{a | b | c}"]
    }

    // { rank = same node1,node2 } // node等级限制， same,min,max,source,sink
}
```

### /proc/partitions
```
文件/proc/partitions 可以查看分区信息
major minor  #blocks  name
   1        0      65536 ram0
   1        1      65536 ram1
   1        2      65536 ram2
   1        3      65536 ram3
   1        4      65536 ram4
   1        5      65536 ram5
   1        6      65536 ram6
   1        7      65536 ram7
   1        8      65536 ram8
   1        9      65536 ram9
   1       10      65536 ram10
   1       11      65536 ram11
   1       12      65536 ram12
   1       13      65536 ram13
   1       14      65536 ram14
   1       15      65536 ram15
 179        0   61071360 emmcblk0   // dev emmcblk0 下面都是它的分区
 179        1       1024 emmcblk0p1
 179        2      20480 emmcblk0p2
 179        3     512000 emmcblk0p3
 179        4      81920 emmcblk0p4
 179        5   59768832 emmcblk0p5
 179       24      16384 emmcblk0rpmb
 179       16       4096 emmcblk0boot1
 179        8       4096 emmcblk0boot0
 179       32   31457280 mmcblk1    // dev mmcblk1 下面都是它的分区
 179       33   31457248 mmcblk1p1

fdisk -l 也能获得类似信息

// 注意
mount, df -h 命令只能看到挂载的信息
```

### shell 交互式输入
```
// 提示进行确认（输入正常退出，输入错误则需重新输入）

#!/bin/bash

while true
do
    read -r -p "Are You Sure? [y/n] " input
    
    case $input in 
        [yY][eE][sS]|[yY])
            echo "yes"
            exit 1
            ;;

        [nN][oO]|[nN])
            echo "no"
            exit 1
            ;;

        *)
            echo "invalid input... "
    esac

done
```

### ssh-keygen -R
```
远程ip发生变化，登陆失败，需要先清除know_hosts，再登陆
ssh-keygen -R [10.151.3.69]:8847
```
