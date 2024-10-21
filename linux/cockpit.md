

## cockpit
https://github.com/cockpit-project/cockpit

Web管理中心管理linux/简单易用的linux系统管理工具

https://cockpit-project.org/running.html#centos
centos 8+
```sh
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