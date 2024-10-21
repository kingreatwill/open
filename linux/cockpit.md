

## cockpit
https://github.com/cockpit-project/cockpit

Web管理中心管理linux/简单易用的linux系统管理工具

https://cockpit-project.org/running.html#centos
centos 8+
```sh

sudo dnf clean all
sudo dnf update

dnf install -y cockpit

systemctl start cockpit.socket
systemctl enable cockpit.socket
#systemctl start cockpit.socket
#systemctl enable cockpit.socket

# 防火墙, 根据需要执行
firewalld-cmd --permanent --add-service=cockpit

systemctl status cockpit
# 访问https://os_ip:9090/
```

Ubuntu 20.04
```sh
sudo apt install cockpit
# 如果你想要防火墙功能，安装firewalld
sudo apt install firewalld
# 如果你想要虚拟机管理功能，安装cockpit-machines
sudo apt install cockpit-machines
# 如果你想要容器管理功能，安装cockpit-podman，Ubuntu 20.04暂时没有该包
sudo apt install cockpit-podman
# 启动服务
sudo systemctl start cockpit
```

### root登录权限被拒绝

cat /etc/cockpit/disallowed-users
注释root账号

新增账号和密码
创建用户有两条命令：adduer和useradd，对应着两条删除用户的命令：deluser和userdel。
这两种命令之间的区别：
adduser：会自动为创建的用户指定主目录、系统shell版本，会在创建时输入用户密码。
useradd：需要使用参数选项指定上述基本设置，如果不使用任何参数，则创建的用户无密码、无主目录、没有指定shell版本。

sudo useradd wcoder #创建wcoder用户
sudo passwd wcoder #为wcoder用户设置密码

切换用户的命令为：su username

### 查看日志
journalctl -u cockpit


## webmin 
https://github.com/webmin/webmin