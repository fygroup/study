### 相关概念
```
LXC:linux容器3个核心技术

Namespace：Namespace又称为命名空间，它主要做访问隔离，不同的容器间 进程pid可以相同，进程并不冲突影响；但可以共享底层的计算和存储

Cgroups：Cgroups是Linux内核功能，它让两件事情变成可能：限制Linux进程组的资源占用（内存、CPU）；为进程组制作 PID、UTS、IPC、网络、用户及装载命名空间。

union文件系统：在union文件系统里，文件系统可以被装载在其他文件系统之上，其结果就是一个分层的积累变化。

```

### 镜像与容器
```
镜像是docker生命周期的构建和打包阶段，容器是docker的启动和执行阶段
```

### 资源隔离
```
文件系统隔离
进程隔离
网络隔离

使用cgroups分离
```

### 运行一个容器的时候，docker会
```
拉取镜像，若本地已经存在该镜像，则不用到网上去拉取
创建新的容器
分配union文件系统并且挂着一个可读写的层，任何修改容器的操作都会被记录在这个读写层上，你可以保存这些修改成新的镜像，也可以选择不保存，那么下次运行改镜像的时候所有修改操作都会被消除
分配网络\桥接接口，创建一个允许容器与本地主机通信的网络接口
设置ip地址，从池中寻找一个可用的ip地址附加到容器上，换句话说，localhost并不能访问到容器
运行你指定的程序
捕获并且提供应用输出，包括输入、输出、报错信息
```

### linux namespace 

#### UTS namespace
UTS namespace 用来隔离系统的 hostname 以及 NIS domain name。
(1) clone
```
//clone可以让你有选择性的继承父进程的资源，比fork更灵活
//flags就是标志用来描述你需要从父进程继承哪些资源、哪些namespace， arg就是传给子进程的参数
int clone(int (*fn)(void *), void *child_stack, int flags, void *arg);
pid_t child_pid = clone(child_func,  //子进程将执行child_func这个函数
                    //栈是从高位向低位增长，所以这里要指向高位地址
                    child_stack + sizeof(child_stack),
                    //CLONE_NEWUTS表示创建新的UTS namespace，
                    //这里SIGCHLD是子进程退出后返回给父进程的信号，跟namespace无关
                    CLONE_NEWUTS | SIGCHLD,
                    argv[1]);  //传给child_func的参数
```

(2) sethostname
```
//设置当前的主机名
sethostname(hostname, strlen(hostname)); 
```

(3) execlp
```
//用一个新的bash来替换掉当前子进程，
//执行完execlp后，子进程没有退出，也没有创建新的进程,
//只是当前子进程不再运行自己的代码，而是去执行bash的代码,
//详情请参考"man execlp"
//bash退出后，子进程执行完毕
execlp("bash", "bash", (char *) NULL);
//如果执行成功则函数不会返回, 执行失败则直接返回-1
```

(4) readlink
```
//读取当前进程（bash进程）的UTS
readlink /proc/$$/ns/uts
uts:[4026532455]
```

(5) 将当前进程加入指定的namespace
```
//获取namespace对应文件的描述符
int fd = open("/proc/PID/ns/FILE", O_RDONLY); //FILE可以是不同的namespace(uts, ipc...)
//执行完setns后，当前进程将加入指定的namespace
//这里第二个参数为0，表示由系统自己检测fd对应的是哪种类型的namespace
int ret = setns(fd, 0);
```

(6) 退出当前namespace并加入新创建的namespace
```
//执行完unshare函数后，当前进程就会退出当前的一个或多个类型的namespace,
//然后进入到一个或多个新创建的不同类型的namespace
int ret = unshare(flags);
```

(7) 内核实现
```
每个进程对应的task结构体struct task_struct中, 有一个叫nsproxy的字段，类型是struct nsproxy。
struct task_struct {
  ...
  /* namespaces */
  struct nsproxy *nsproxy;
  ...
}

struct nsproxy {
  atomic_t count;
  struct uts_namespace *uts_ns;
  struct ipc_namespace *ipc_ns;
  struct mnt_namespace *mnt_ns;
  struct pid_namespace *pid_ns_for_children;
  struct net       *net_ns;
  struct cgroup_namespace *cgroup_ns;
};

//gethostname原理
static inline struct new_utsname *utsname(void)
{
  //current指向当前进程的task结构体
  return &current->nsproxy->uts_ns->name;
}

SYSCALL_DEFINE2(gethostname, char __user *, name, int, len)
{
  struct new_utsname *u;
  ...
  u = utsname();
  if (copy_to_user(name, u->nodename, i)){ //内核空间拷贝到用户空
    errno = -EFAULT;
  }
  ...
}

处于不同UTS namespace中的进程，它task结构体里面的nsproxy->uts_ns所指向的结构体是不一样的，于是达到了隔离UTS的目的。其他类型的namespace基本上也是差不多的原理。
```

(8) 总结
```
1、namespace的本质就是把原来所有进程全局共享的资源拆分成了很多个一组一组进程共享的资源
2、当一个namespace里面的所有进程都退出时，namespace也会被销毁，所以抛开进程谈namespace没有意义
3、UTS namespace就是进程的一个属性，属性值相同的一组进程就属于同一个namespace，跟这组进程之间有没有亲戚关系无关
4、clone和unshare都有创建并加入新的namespace的功能，他们的主要区别是：
   unshare是使当前进程加入新创建的namespace
   clone 是创建一个新的子进程，然后让子进程加入新的namespace
5、UTS namespace没有嵌套关系，即不存在说一个namespace是另一个namespace的父namespace
```

#### IPC namespace
```
主要隔离的是：System V IPC的 消息队列、信号量（semaphore）、共享内存
但是signal（pid隔离，signal就隔离了）、pipe（无名管道只在父子间，有名管道与文件系统有关，文件系统隔离他就隔离了）、socket（network namespace不同）不在IPC中隔离。
```

(1) ipcmk
```
创建shared memory segments, message queues, 和semaphore arrays
ipcmk -Q   //创建新的ipc message queues
```

(2) ipcs
```
查看shared memory segments, message queues, 和semaphore arrays的相关信息
ipcs -q   //查看现有的ipc Message Queues
```

(3) unshare
```
开当前指定类型的namespace，创建且加入新的namespace，然后执行参数中指定的命令
unshare -iu /bin/bash  //运行unshare创建新的ipc和uts namespace，并且在新的namespace中启动bash
```

(4) nsenter
```
加入指定进程的指定类型的namespace，然后执行参数中指定的命令
nsenter -t 27668 -u -i /bin/bash  
//试着加入pid等于27668的进程的uts和ipc namespace
//-t后面跟pid用来指定加入哪个进程所在的namespace
//这里27668是正在运行的bash的pid
//加入成功后将运行/bin/bash
```

#### mount namespace
```
（linux命令-mount）

Mount namespace用来隔离文件系统的挂载点, 使得不同的mount namespace拥有自己独立的挂载点信息，不同的namespace之间不会相互影响，这对于构建用户或者容器自己的文件系统目录非常有用。

当前进程所在mount namespace里的所有挂载信息可以在/proc/[pid]/mounts、/proc/[pid]/mountinfo和/proc/[pid]/mountstats里面找到。

函数clone()的flag是CLONE_NEWNS

clone或者unshare函数创建新的mount namespace时,新的拷贝老的，从这之后，他们就没有关系了，通过mount和umount增加和删除各自namespace里面的挂载点都不会相互影响。
```

示例
```
mkdir iso
cd iso/
mkdir -p iso01/subdir01
mkdir -p iso02/subdir02
mkisofs -o ./001.iso ./iso01
mkisofs -o ./002.iso ./iso02
mkdir /mnt/iso1 /mnt/iso2
mount ./001.iso /mnt/iso1/
  mount: /dev/loop1 is write-protected, mounting read-only
mount |grep /001.iso
  /home/dev/iso/001.iso on /mnt/iso1 type iso9660 (ro,relatime)

//#创建并进入新的mount和uts namespace
unshare --mount --uts /bin/bash
hostname container001
exec bash
```

Shared subtrees
```
比如系统添加了一个新的硬盘，这个时候如果mount namespace是完全隔离的，想要在各个namespace里面用这个硬盘，就需要在每个namespace里面手动mount这个硬盘，这个是很麻烦的，这时Shared subtrees就可以帮助我们解决这个问题。

dd if=/dev/zero bs=1M count=32 of=./disk1.img
dd if=/dev/zero bs=1M count=32 of=./disk2.img
mkfs.ext2 ./disk1.img
mkfs.ext2 ./disk2.img
mkdir disk1 disk2

//显式的分别以shared和private方式挂载disk1和disk2
mount --make-shared ./disk1.img ./disk1
mount --make-private ./disk2.img ./disk2
cat /proc/self/mountinfo |grep disk| sed 's/ - .*//'
164 24 7:1 / /home/dev/disks/disk1 rw,relatime shared:105
173 24 7:2 / /home/dev/disks/disk2 rw,relatime

//查看mount namespace编号
readlink /proc/$$/ns/mnt
mnt:[4026531840]

//创建新的mount namespace
//默认情况下，unshare会将新namespace里面的所有挂载点的类型设置成private，
//所以这里用到了参数--propagation unchanged，让新namespace里的挂载点的类型和老namespace里保持一致。
//--propagation参数还支持private|shared|slave类型（私有 共享 从属），
//和mount命令的那些--make-private参数一样，
unshare --mount --uts --propagation unchanged /bin/bash
hostname container001
exec bash
~/disks# 

//确认已经是在新的mount namespace里面了
root@container001:~/disks# readlink /proc/$$/ns/mnt
mnt:[4026532463]

//由于前面指定了--propagation unchanged，
//所以新namespace里面的/home/dev/disks/disk1也是shared，
//且和老namespace里面的/home/dev/disks/disk1属于同一个peer group 105
//因为在不同的namespace里面，所以这里挂载点的ID和原来namespace里的不一样了
cat /proc/self/mountinfo |grep disk| sed 's/ - .*//'
221 177 7:1 / /home/dev/disks/disk1 rw,relatime shared:105
222 177 7:2 / /home/dev/disks/disk2 rw,relatime
```

#### PID namespace
```
用来隔离进程的ID空间，使得不同pid namespace里的进程ID可以重复且相互之间不影响。

PID namespace可以嵌套，也就是说有父子关系,父namespace可以看子孙后代namespace里的进程信息，而子看不到祖先或者兄弟namespace里的进程信息。

目前PID namespace最多可以嵌套32层，由内核中的宏MAX_PID_NS_LEVEL来定义

Linux下的每个进程都有一个对应的/proc/PID目录, 对一个PID namespace而言，/proc目录只包含当前namespace和它所有子孙后代namespace里的进程的信息。

//命令
unshare --uts --pid --mount --fork /bin/bash
启动新的pid namespace，unshare进程fork一个新的进程出来，然后再用bash替换掉新的进程
注意：
    调用unshare和nsenter后，原来的进程还是属于老的namespace，而新fork出来的进程才是新的namespace

//pid namespace嵌套
  old namespace     |      new namespace
unshare、setns  ---------> ...
                   fork
                   
调用unshare或者setns函数后，当前进程的namespace不会发生变化，不会加入到新的namespace，而它的子进程会加入到新的namespace。也就是说进程属于哪个namespace是在进程创建的时候决定的，并且以后再也无法更改。
readlink /proc/$$/ns/uts 
uts:[4026531838]
unshare --uts /bin/bash  //表示用/bin/bash替换当前的进程
readlink /proc/$$/ns/uts
uts:[4026532440] <------ 当前的uts与之前的不一样

在一个PID namespace里的进程，它的父进程可能不在当前namespace中，而是在外面的namespace里面（这里外面的namespace指当前namespace的祖先namespace），这类进程的ppid都是0。比如新namespace里面的第一个进程，他的父进程就在外面的namespace里。通过setns的方式加入到新namespace中的进程的父进程也在外面的namespace中。

可以在祖先namespace中看到子namespace的所有进程信息，且可以发信号给子namespace的进程，但进程在不同namespace中的PID是不一样的。

```

(1) 示例
```
unshare --uts --pid --mount --fork /bin/bash
hostname container001
exec bash
root@container001:~#

#查看进程间关系，当前bash(31646)确实是unshare的子进程
root@container001:~# pstree -pl
├─sshd(955)─┬─sshd(17810)───sshd(17891)───bash(17892)───sudo(31644)──
─unshare(31645)───bash(31646)───pstree(31677)
#他们属于不同的pid namespace
root@container001:~# readlink /proc/31645/ns/pid
pid:[4026531836]
root@container001:~# readlink /proc/31646/ns/pid
pid:[4026532469]

#但为什么通过这种方式查看到的namespace还是老的呢？
root@container001:~# readlink /proc/$$/ns/pid
pid:[4026531836]

#由于我们实际上已经是在新的namespace里了，并且当前bash是当前namespace的第一个进程
#所以在新的namespace里看到的他的进程ID是1
root@container001:~# echo $$
1
#但由于我们新的namespace的挂载信息是从老的namespace拷贝过来的，
#所以这里看到的还是老namespace里面的进程号为1的信息
root@container001:~# readlink /proc/1/ns/pid
pid:[4026531836]
#ps命令依赖/proc目录，所以ps的输出还是老namespace的视图
root@container001:~# ps ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 7月07 ?       00:00:06 /sbin/init
root         2     0  0 7月07 ?       00:00:00 [kthreadd]
 ...
root     31644 17892  0 7月14 pts/0   00:00:00 sudo unshare --uts --pid --mount --fork /bin/bash
root     31645 31644  0 7月14 pts/0   00:00:00 unshare --uts --pid --mount --fork /bin/bash

#所以我们需要重新挂载我们的/proc目录
root@container001:~# mount -t proc proc /proc

#重新挂载后，能看到我们新的pid namespace ID了
root@container001:~# readlink /proc/$$/ns/pid
pid:[4026532469]
#ps的输出也正常了
root@container001:~# ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 7月14 pts/0   00:00:00 bash
root        44     1  0 00:06 pts/0    00:00:00 ps -ef
```

#### network namespace
```
用来隔离网络设备, IP地址, 端口等. 每个namespace将会有自己独立的网络栈，路由表，防火墙规则，socket等。

每个新的network namespace默认有一个本地环回接口，lo默认关闭。

local devices不能从一个namespace移动到另一个namespace(loopback、bridge、ppp).

ethtool -k命令来查看设备的netns-local
#这里“on”表示该设备不能被移动到其他network namespace
ethtool -k lo|grep netns-local
netns-local: on [fixed]

ip netns add ns1，本质上就是调用`unshare(CLONE_NEWNET)`创建了一个新的network namespace

ip link set eth0 netns ns1就能将eth0网络设备移动到network namespace ns1中。
```

1、操作
```
//---------------第一个shell窗口----------------------------
（1）创建network namespace
  //记录默认network namespace ID
  readlink /proc/$$/ns/net
  net:[4026531957]

  //创建新的network namespace
  unshare --uts --net /bin/bash
  hostname container001
  exec bash
  readlink /proc/$$/ns/net
  net:[4026532478]

  //运行ifconfig啥都没有

  //启动lo（环回接口）
  ip link set lo up
  ping 127.0.0.1 （可以ping通）

  //获取当前bash进程的PID
  echo $$
  15812


 //-----------------第二个shell窗口------------------------------
（2）创建两个虚拟网络设备，并使他们之间可以相互通讯
  //创建新的虚拟以太网设备，让两个namespace能通讯
  sudo ip link add veth0 type veth peer name veth1

  //veth1移动到第一个shell窗口的net namespace里面
  ip link set veth1 netns 15812

  //为veth0分配IP并启动veth0
  ip address add dev veth0 192.168.8.1/24

  //打开veth0设备
  ip link set veth0 up

  //ifconfig veth0
  veth0     Link encap:Ethernet  HWaddr 9a:4d:d5:96:b5:36
            inet addr:192.168.8.1  Bcast:0.0.0.0  Mask:255.255.255.0
            inet6 addr: fe80::984d:d5ff:fe96:b536/64 Scope:Link
            UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
            RX packets:8 errors:0 dropped:0 overruns:0 frame:0
            TX packets:8 errors:0 dropped:0 overruns:0 carrier:0
            collisions:0 txqueuelen:1000
            RX bytes:648 (648.0 B)  TX bytes:648 (648.0 B)

  //-------------第一个shell窗口的net namespace-------------------
  //为veth1分配IP地址并启动它
  ip address add dev veth1 192.168.8.2/24
  ip link set veth1 up
  ifconfig veth1
  veth1     Link encap:Ethernet  HWaddr 6a:dc:59:79:3c:8b
            inet addr:192.168.8.2  Bcast:0.0.0.0  Mask:255.255.255.0
            inet6 addr: fe80::68dc:59ff:fe79:3c8b/64 Scope:Link
            UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
            RX packets:8 errors:0 dropped:0 overruns:0 frame:0
            TX packets:8 errors:0 dropped:0 overruns:0 carrier:0
            collisions:0 txqueuelen:1000
            RX bytes:648 (648.0 B)  TX bytes:648 (648.0 B)

  //连接成功
  ping 192.168.8.1 （可以ping通）


（3）netns连接外网
  //-----------------第二个shell窗口-------------------
  //基本思路： 虚拟网卡 --> 实际网卡 -->
  //允许主机内部的虚拟网络设备进行数据转发
  cat /proc/sys/net/ipv4/ip_forward
  1 （如果不是1，要改成1）

  //添加NAT规则，这里ens32是机器上连接外网的网卡（就是说发送数据必须是实际网卡的ip）
  iptables -t nat -A POSTROUTING -o ens32 -j MASQUERADE
  此时第二个shell的虚拟网络设备veth1可以访问外网

  //----------------第一个shell窗口--------------------
  //由于veth0与veth1是联通的，只需要将veth0的网关设为veth1的ip
  ip route add default via 192.168.8.1

  //查看路由表
  route -n

```

2、管理net namepace
```
给namespace取名字其实就是创建一个文件，然后通过mount --bind将新创建的namespace文件和该文件绑定，就算该namespace里的所有进程都退出了，内核还是会保留该namespace，还可以通过这个绑定的文件来加入该namespace。

//开始之前，获取一下默认network namespace的ID
readlink /proc/$$/ns/net
net:[4026531957]

//创建一个用于绑定network namespace的文件，
//ip netns将所有的文件放到了目录/var/run/netns下，
//所以我们这里重用这个目录，并且创建一个我们自己的文件netnamespace1
mkdir -p /var/run/netns
touch /var/run/netns/netnamespace1

//创建新的network namespace，并在新的namespace中启动新的bash
unshare --net bash
//查看新的namespace ID
readlink /proc/$$/ns/net
net:[4026532448]

//bind当前bash的namespace文件到上面创建的文件上
mount --bind /proc/$$/ns/net /var/run/netns/netnamespace1
#通过ls -i命令可以看到文件netnamespace1的inode号和namespace的编号相同，说明绑定成功
ls -i /var/run/netns/netnamespace1
4026532448 /var/run/netns/netnamespace1
//退出新创建的bash
exit

//可以看出netnamespace1的inode没变，说明我们使用了bind mount后
//虽然新的namespace中已经没有进程了，但这个新的namespace还存在
ls -i /var/run/netns/netnamespace1
4026532448 /var/run/netns/netnamespace1

//上面的这一系列操作等同于执行了命令： ip netns add netnamespace1
//下面的nsenter命令等同于执行了命令： ip netns exec netnamespace1 bash

//我们可以通过nsenter命令再创建一个新的bash，并将它加入netnamespace1所关联的namespace（net:[4026532448]）
nsenter --net=/var/run/netns/netnamespace1 bash
readlink /proc/$$/ns/net
net:[4026532448]
```


#### user namespace
```
权限涉及的范围非常广，所以导致user namespace比其他的namespace要复杂； 同时权限也是容器安全的基础，所以user namespace非常重要。

用于隔离安全相关的资源，包括 user IDs and group IDs，keys, 和 capabilities。同样一个用户的 user ID 和 group ID 在不同的 user namespace 中可以不一样(与 PID nanespace 类似)。换句话说，一个用户可以在一个 user namespace 中是普通用户，但在另一个 user namespace 中是超级用户。

非 root 进程也可以创建User Namespace ， 并且此用户在Namespace 里面可以被映射成root ， 且在Namespace 内有root 权限。

//例如：子user namespace虽然是root权限，但是不能操作父user namespace的内容
unshare --user -r /bin/bash
hostname newname       //premisee deny, 因为此时的user namespace是不能操作父namespace的hostname

//创建user namespace，并映射uid和gid
unshare --user -r /bin/bash
```

1、和其他namespace的关系
```
//多namespace同时创建
unshare(CLONE_NEWUSER | CLONE_NEW*);
内核会保证CLONE_NEWUSER先被执行，然后执行剩下的其他CLONE_NEW*，这样就使得不用root账号而创建新的容器成为可能，这条规则对于clone函数也同样适用

//和其他类型namespace的关系
每个namespace都有一个owner（user namespace），这样保证对任何namespace的操作都受到user namespace权限的控制。
//例如uts namespace的结构体
struct uts_namespace {
  struct kref kref;
  struct new_utsname name;
  struct user_namespace *user_ns;   //指向它属于的user namespace，其实另外namespace也有userNs
  struct ns_common ns;
};
```

2、不和任何user namespace关联的资源
```
//当和mount namespace一起用时（注意是一起用！！！），不能挂载基于块设备的文件系统，但是可以挂载下面这些文件系统

//但是可以挂载一些特殊的文件系统
  /proc (since Linux 3.8)  
  /sys (since Linux 3.8) 
  devpts (since Linux 3.9)    伪终端提供了一个标准接口，它的标准挂接点是/dev/pts
  tmpfs (since Linux 3.9)       tmp文件系统
  ramfs (since Linux 3.9)     虚拟内存文件系统
  mqueue (since Linux 3.9)
  bpf (since Linux 4.4)

//示例
//新建user、mount namespace
unshare --user -r --mount bash
//mount块设备到./mnt, 结果是mount失败
mount /dev/mapper/ubuntu--vg-root ./mnt
  mount: /dev/mapper/ubuntu--vg-root is write-protected, mounting read-only
  mount: cannot mount /dev/mapper/ubuntu--vg-root read-only

//由于当前pid namespace不属于当前的user namespace，所以挂载/proc失败
mount -t proc none ./mnt
  mount: permission denied
//创建新的pid namespace，然后挂载成功
unshare --pid --fork bash
mount -t proc none ./mnt

//只能通过bind方式挂载devpts，直接mount报错
mount -t devpts devpts ./mnt
  mount: wrong fs type, bad option, bad superblock on devpts,
       missing codepage or helper program, or other error

       In some cases useful info is found in syslog - try
       dmesg | tail or so.
mount --bind /dev/pts ./mnt
mount|grep mnt|grep devpts
  devpts on /home/dev/mnt type devpts (rw,nosuid,noexec,relatime,mode=600,ptmxmode=000)

//sysfs直接mount和bind mount都不行
mount -t sysfs sysfs ./mnt
  mount: permission denied
mount --bind /sys ./mnt
  mount: wrong fs type, bad option, bad superblock on /sys,
  
//挂载tmpfs成功
mount -t tmpfs tmpfs ./mnt
mount|grep mnt|grep tmpfs
  tmpfs on /home/dev/mnt type tmpfs (rw,nodev,relatime,uid=1000,gid=1000)
```

3、mount namespace和user namespace
```
//当mount namespace和user namespace一起用时，就算老mount namespace中的mount point是shared并且用unshare命令时指定了--propagation shared，新mount namespace里面的挂载点的propagation type还是slave。这样就防止了在新user namespace里面mount的东西被外面父user namespace中的进程看到。
```


### cgroup
```
Namespace主要用于隔离资源
Cgroups用来提供对一组进程以及将来子进程的资源限制

（1）三个组件
  >控制族群（control group）：层级中的节点。
    Cgroups 中的资源控制都是以控制族群为单位实现。一个进程可以加入到某个控制族群，也从一个进程组迁移到另一个控制族群。一个进程组的进程可以使用 cgroups 以控制族群为单位分配的资源，同时受到 cgroups 以控制族群为单位设定的限制；
  >层级（hierarchy）：一个或多个子系统的组合
  >子系统（subsytem）：一个子系统就是一个资源控制器
    cpu 子系统，主要限制进程的 cpu 使用率。
    cpuacct 子系统，可以统计 cgroups 中的进程的 cpu 使用报告。
    cpuset 子系统，可以为 cgroups 中的进程分配单独的 cpu 节点或者内存节点。
    memory 子系统，可以限制进程的 memory 使用量。
    blkio 子系统，可以限制进程的块设备 io。
    devices 子系统，可以控制进程能够访问某些设备。
    net_cls 子系统，可以标记 cgroups 中进程的网络数据包，然后可以使用 tc 模块（traffic control）对数据包进行控制。
    freezer 子系统，可以挂起或者恢复 cgroups 中的进程。
    ns 子系统，可以使不同 cgroups 下面的进程使用不同的 namespace。


（2）关系图
  > 每次在系统中创建新层级时，该系统中的所有任务都是那个层级的默认 cgroup（我们称之为 root cgroup，此 cgroup 在创建层级时自动创建，后面在该层级中创建的 cgroup 都是此 cgroup 的后代）的初始成员；
  > 一个子系统最多只能附加到一个层级；
  > 一个层级可以附加多个子系统；
  > 一个任务可以是多个 cgroup 的成员，但是这些 cgroup 必须在不同的层级；
  > 系统中的进程（任务）创建子进程（任务）时，该子任务自动成为其父进程所在 cgroup 的成员。然后可根据需要将该子任务移动到不同的 cgroup 中，但开始时它总是继承其父任务的 cgroup。


  子系统                 cpu   cpu_set                                  memory
                        |       |                                        |      
                        |       |                                        |
  层级      |------------+-------+----------------|     |-----------------+-----------------------|
  控制组                 root_cgroup                                  root_cgroup
                            |                                             |
                    +-------+---------+                         +---------+-----------+
                    |                 |                         |                     |       
            cgroup1(10% cpu占有)     cgroup2(40%)          cgroup1(20% 内存占有)     cgroup2(70%)
                                        |                         |    
                                        +-----------+-------------+
                                                    |
  进程组                                         task_group
                                                    |
                                          +---------+----------+
  任务                                   task1    task2      task3


（3）相关文件
  /proc/cgroups
  #subsys_name    hierarchy       num_cgroups     enabled
  cpuset          11              1               1
  cpu             3               64              1
  cpuacct         3               64              1
  blkio           8               64              1
  memory          9               104             1
  devices         5               64              1
  freezer         10              4               1
  net_cls         6               1               1
  perf_event      7               1               1
  net_prio        6               1               1
  hugetlb         4               1               1
  pids            2               68              1

  /proc/[pid]/cgroup
    11:cpuset:/
    5:devices:/system.slice/cron.service
    4:hugetlb:/
    3:cpu,cpuacct:/system.slice/cron.service
    2:pids:/system.slice/cron.service
    1:name=systemd:/system.slice/cron.service
  > cgroup树的ID， 和/proc/cgroups文件中的ID一一对应。
  > 和cgroup树绑定的所有subsystem，多个subsystem之间用逗号隔开。这里name=systemd表示没有和任何subsystem绑定，只是给他起了个名字叫systemd。
  > 进程在cgroup树中的路径，即进程所属的cgroup，这个路径是相对于挂载点的相对路径。
```

1、cgroup使用
```
cgroup相关的所有操作都是基于内核中的cgroup virtual filesystem，使用cgroup很简单，挂载这个文件系统就可以了。一般情况下都是挂载到/sys/fs/cgroup目录下

（1）创建（挂载）cgroup树（层级）
  注意：xxx为任意字符串
  //挂载一颗和所有subsystem关联的cgroup树到/sys/fs/cgroup
  mount -t cgroup xxx /sys/fs/cgroup

  //挂载一颗和cpuset subsystem关联的cgroup树到/sys/fs/cgroup/cpuset
  mkdir /sys/fs/cgroup/cpuset
  mount -t cgroup -o cpuset xxx /sys/fs/cgroup/cpuset

  //挂载一颗与cpu和cpuacct subsystem关联的cgroup树到/sys/fs/cgroup/cpu,cpuacct
  mkdir /sys/fs/cgroup/cpu,cpuacct
  mount -t cgroup -o cpu,cpuacct xxx /sys/fs/cgroup/cpu,cpuacct

  //挂载一棵cgroup树，但不关联任何subsystem，下面就是systemd所用到的方式
  mkdir /sys/fs/cgroup/systemd
  mount -t cgroup -o none,name=systemd xxx /sys/fs/cgroup/systemd
  ls /sys/fs/cgroup/systemd
    cgroup.clone_children  cgroup.procs  cgroup.sane_behavior  notify_on_release  release_agent  tasks
    > cgroup.procs 当前cgroup中的所有进程ID，系统不保证ID是顺序排列的，且ID有可能重复
    > tasks 当前cgroup中的所有线程ID，系统不保证ID是顺序排列的

  //创建子cgroup
  //创建并挂载好一颗cgroup树之后，就有了树的根节点，也即根cgroup，这时候就可以通过创建文件夹的方式创建子cgroup，然后再往每个子cgroup中添加进程。在后续介绍具体的subsystem的时候会详细介绍如何操作cgroup。
  mkdir /sys/fs/cgroup/cpu,cpuacct/test


（2）创建和删除cgroup
  挂载好上面的cgroup树之后，就可以在里面建子cgroup了
  cd demo && mkdir cgroup1
  rm -r ./cgroup1

（3）添加进程
  echo 1421 > ./cgroup.procs

  注意：新创建的子进程将会自动加入父进程所在的cgroup
       在一颗cgroup树里面，一个进程必须要属于一个cgroup

（4）权限
  注意：从一个cgroup移动一个进程到另一个cgroup时，只要有目的cgroup的写入权限就可以了，系统不会检查源cgroup里的权限。
  cd /cgroup/demo
  mkdir permission
  chown -R dev:dev ./permission/
  echo 1421 > ./permission/cgroup.procs //1421属于root cgroup，可以将它移动到新的cgroup下，但是反过来不行

（5）cgroup的清理
  当一个cgroup里没有进程也没有子cgroup时，release_agent将被调用来执行cgroup的清理工作。
```

3、资源的限制
```
（1）cpu限制
  子系统有cpusets、cpuacct和cpu
  cpusets: cpuset主要用于设置CPU的亲和性，可以限制cgroup中的进程只能在指定的CPU上运行，或者不能在指定的CPU上运行
  cpuacct: 当前cgroup所使用的CPU的统计信息，信息量较少
  cpu: 
    > cpu.cfs_period_us & cpu.cfs_quota_us
      1.限制只能使用1个CPU（每250ms能使用250ms的CPU时间）
          # echo 250000 > cpu.cfs_quota_us /* quota = 250ms */
          # echo 250000 > cpu.cfs_period_us /* period = 250ms */
      2.限制使用2个CPU（内核）（每500ms能使用1000ms的CPU时间，即使用两个内核）
          # echo 1000000 > cpu.cfs_quota_us /* quota = 1000ms */
          # echo 500000 > cpu.cfs_period_us /* period = 500ms */
      3.限制使用1个CPU的20%（每50ms能使用10ms的CPU时间，即使用一个CPU核心的20%）
          # echo 10000 > cpu.cfs_quota_us /* quota = 10ms */
          # echo 50000 > cpu.cfs_period_us /* period = 50ms */
    > cpu.stat
      nr_periods： 表示过去了多少个cpu.cfs_period_us里面配置的时间周期
      nr_throttled： 上面的这些周期中，有多少次受到了限制（即cgroup中的进程在指定的时间周期中用光了它的配额）
      throttled_time: cgroup中的进程被限制使用CPU持续了多长时间(纳秒)
    > 示例
      echo 50000 > cpu.cfs_period_us
      echo 10000 > cpu.cfs_quota_us
      echo $$ > cgroup.procs 
      while :; do echo test > /dev/null; done  //理论应是100%
      top
          PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND
          5456 dev       20   0   22640   5472   3524 R  20.3  1.1   0:04.62 bash
```


### docker
