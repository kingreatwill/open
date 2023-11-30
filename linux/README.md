<!--toc-->
[TOC]
[Linux学习教程，Linux入门教程（超详细）](http://c.biancheng.net/linux_tutorial/)
[Linux 内核揭秘](https://github.com/MintCN/linux-insides-zh)

[Linux内核地图 ](https://makelinux.github.io/kernel/map/)
[linux kernel map](https://github.com/makelinux/linux_kernel_map)

![](img/LKM.svg)

查看centos版本
cat /etc/redhat-release

1. 常用命令
```
启动Docker服务并激活开机启动：
systemctl start docker & systemctl enable docker
```
2. 关闭防火墙
```
systemctl stop firewalld & systemctl disable firewalld

```
开放端口
iptables -A INPUT -p tcp -m state --state NEW -m tcp --dport 2379 -j ACCEPT
iptables -A INPUT -p tcp -m state --state NEW -m tcp --dport 3000 -j ACCEPT


3. 关闭selinux
```
sed -i 's/enforcing/disabled/' /etc/selinux/config #永久关闭
setenforce 0  #临时关闭
```

方法2：临时和永久关闭Selinux
[root@xuexi ~]# getenforce　　//查看Selinux状态
Enforcing
[root@xuexi ~]# setenforce 0　//临时关闭Selinux
[root@xuexi ~]# getenforce
Permissive
[root@xuexi ~]# vim /etc/selinux/config 　　//修改Selinux配置文件
这是将会进入到/etc/selinux/config文件中
```
# This file controls the state of SELinux on the system.
# SELINUX= can take one of these three values:
#     enforcing - SELinux security policy is enforced.
#     permissive - SELinux prints warnings instead of enforcing.
#     disabled - No SELinux policy is loaded.
SELINUX=enforcing　　//修改此处为disabled
# SELINUXTYPE= can take one of three two values:
#     targeted - Targeted processes are protected,
#     minimum - Modification of targeted policy. Only selected processes are protected.
#     mls - Multi Level Security protection.
SELINUXTYPE=targeted
```
保存并退出，然后重启系统即可。重启后再来查看Selinux
[xf@xuexi ~]$ getenforce
Disabled

4. 关闭swap(Linux的Swap内存交换机制会影响性能以及稳定性。) 虚拟内存
```
swapoff -a #临时关闭
vi /etc/fstab #注释掉swap即可永久关闭
或
sed -i '/ swap / s/^/#/' /etc/fstab #永久关闭
```

4.1. 关闭swap之前先看看开启了哪些swap，执行`cat /proc/swaps`
```
[wcoder@NAS-WCODER ~]$ cat /proc/swaps
Filename				Type		Size		Used		Priority
/dev/md321                              partition	8283708		1752652		-2
/dev/md256                              partition	530108		0		-3
/dev/md322                              partition	7235132		0		-4
/share/CACHEDEV1_DATA/.swap/qnap_swap   file		16777212	0		-5
```
/dev/md256和/dev/md322是分布在hdd上的两个分区，对应资源监控的HDD。/share/CACHEDEV1_DATA/.swap/qnap_swap是一个分布在卷上的文件，对应资源监控的卷。
可以执行`swapon /dev/md256，swapon /dev/md322，swapon /share/CACHEDEV1_DATA/.swap/qnap_swap`这三条命令把交换分区再次启用。

也可以使用`swap -a`启用 
#-a:--all                启用 /etc/fstab 中的所有交换区
[root@centos8 ~]# swapon -a


4.2. 修改系统配置，缓解激进的交换策略
`cat /proc/sys/vm/swappiness`
可以看到默认值是30，这个swappiness值的含义是内存剩余空间小于多少时进行交换。30这个值是linux的默认值，代表了剩余空间小于30%时启用虚拟内存交换。
修改这个配置的方法也很简单，只需要执行 `echo 15 > /proc/sys/vm/swappiness`
修改以后可能不会马上有感觉，用一段时间以后会发现虚拟内存的占用下降，物理内存的占用上升。

5. 设置固定IP
```
cd /etc/sysconfig/network-scripts
vi ifcfg-eth0

主要修改如下信息，这里我设置静态IP地址为192.168.137.200

BOOTPROTO=static
DEVICE=eth0
ONBOOT=yes
IPADDR=192.168.137.200
GATEWAY=192.168.137.1
DNS1=192.168.137.1
NETMASK=255.255.255.0


重启网络服务，使设置生效：
systemctl restart network
```

> x86_64，x64，AMD64基本上是同一个东西
> x86是intel开发的一种32位指令集
> x84_64是CPU迈向64位的时候
> x86_64是一种64位的指令集，x86_64是x86指令的超集，在x86上可以运行的程序，在x86_64上也可以运行，x86_64是AMD发明的，也叫AMD64
> 现在用的intel/amd的桌面级CPU基本上都是x86_64，与之相对的arm、pcc等都不是x86_64

## job

nohup cmd &

ctrl+z 转后台运行(暂停已经运行的进程)
jobs 查看后台任务
jobs -l


fg  %n   //将编号为n的任务转前台运行,%可以不要
fg 1

bg 1 转后台运行 (将停止的作业放到后台运行)
bg  %n   //将编号为n的任务转后台运行,%可以不要


## 环境变量
方法一：直接运行命令export PATH=$PATH:/usr/local/webserver/php/bin 和 export PATH=$PATH:/usr/local/webserver/mysql/bin

使用这种方法，只会对当前会话有效，也就是说每当登出或注销系统以后，PATH 设置就会失效，只是临时生效。

方法二：执行vi ~/.bash_profile修改文件中PATH一行，将/usr/local/webserver/php/bin 和 /usr/local/webserver/mysql/bin 加入到PATH=$PATH:$HOME/bin一行之后

source ~/.bash_profile
这种方法只对当前登录用户生效

方法三：修改/etc/profile文件使其永久性生效，并对所有系统用户生效，在文件末尾加上如下两行代码
PATH=$PATH:/usr/local/webserver/php/bin:/usr/local/webserver/mysql/bin
export PATH

最后：执行 命令source /etc/profile或 执行点命令 ./profile使其修改生效，执行完可通过echo $PATH命令查看是否添加成功。

```
当脚本准备好后，把脚本文件的路径加到 .bash_profile 文件末尾（红帽系列的系统：CentOS、Oracle Linux 和 Fedora）：

# echo "/root/system-info.sh" >> ~root/.bash_profile
执行以下命令，来让修改的内容生效：

# source ~root/.bash_profile
对于 Debian 系统的系统，你可能需要把文件路径加到 .profile 文件中：

# echo "/root/system-info.sh" >> ~root/.profile
运行以下命令使修改生效：

# source ~root/.profile
```

## 将文件从dos格式、unix格式相互转化
- `vi 文件名` 
- `set ff?` 查看文件格式dos或unix的字样.
- `set ff=dos`  #强制转化为dos格式，如果是要转化成unix格式就是   `set ff=unix`
- `w!` 保存

## 释放磁盘空间
`df -h`

- 查找文件型
进入根目录查找最大的文件: `du -a|sort -rn|head -5`
进入根目录查找大于300M的文件:`find . -type f -size +300M`

- docker
释放docker无用空间(镜像,容器,缓存): 
`docker system prune -a` --Remove all unused containers, networks, images (both dangling and unreferenced), and optionally, volumes. 这个命令直接释放了我35G的空间
```
docker container prune -a
docker image prune -a
```

- 软件源
centos7,8: `yum clean all` -- 一般也能释放几个G
centos8: `dnf clean all`
Debian 和 Ubuntu: `apt-get clean`