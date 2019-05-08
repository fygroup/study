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
//---添加用户-------------------------------------------
groupadd gitgroup   //添加组
useradd malx        //添加用户
passwd malx         //设置用户密码
usermod -G gitgroup malx //为用户分配组
chgrp -R gitgroup /home/data/git/  //修改文件夹的组
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


