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

4. 关闭swap(Linux的Swap内存交换机制会影响性能以及稳定性。)
```
swapoff -a #临时关闭
vi /etc/fstab #注释掉swap即可永久关闭
或
sed -i '/ swap / s/^/#/' /etc/fstab #永久关闭
```

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
