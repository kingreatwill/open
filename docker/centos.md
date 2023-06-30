
```dockerfile
FROM centos:7.9.2009

ENV LANG=zh_CN.UTF-8 \
LANGUAGE=zh_CN:zh \
LC_ALL=zh_CN.UTF-8

# Install tools
RUN yum update -y && \
yum reinstall -y glibc-common && \
yum install -y telnet net-tools && \
yum install -y make && \
yum install -y gcc gcc-c++ kernel-devel && \
yum install -y lsof wget zlib zlib-devel dos2unix file && \
yum clean all && \
rm -rf /tmp/* rm -rf /var/cache/yum/* && \
localedef -c -f UTF-8 -i zh_CN zh_CN.UTF-8 && \
ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# Define default command.
CMD ["bash"]
```

```
yum install -y make
yum install -y gcc gcc-c++ kernel-devel

# 2 为了查看端口信息更方便可以安装 lsof
yum install -y lsof

# 3 安装 wget
yum install -y wget

# 4 安装 tree
yum install -y tree

# 5 python 工具
yum install -y python-devel

# 6 安装编译 C 的环境
yum install -y gcc gcc-c++
yum install -y zlib
yum install -y zlib-devel
yum install -y tcl build-essential tk gettext

SSH
# 1 yum 安装 spenssl 服务
yum -y install passwd openssl openssh-server openssh-clients
mkdir /var/run/sshd/

# 2 修改配置
vim /etc/ssh/sshd_config +39
## 大概在 38 - 45 行之间，修改或添加如下三个配置
PermitRootLogin yes
RSAAuthentication yes
PubkeyAuthentication yes

# 3 sshd 服务的启停
## 3.1 启动
systemctl start sshd.service
## 3.2 查看 sshd 服务状态
systemctl status sshd
## 3.3 停止
systemctl start sshd.service

# 4 设置为开机自启
systemctl enable sshd.service

# 【可跳过】5 生成ssh的密钥和公钥
# ssh-keygen -t rsa

# 6 查看 SSH 服务
lsof -i:22

# 7 设置 root 密码（2020）
passwd

# 8 通过 ssh 访问容器
ssh root@bigdata
```

生成新镜像
```
# 2 commit 该 docker 容器
docker commit $CONTAINER_ID new_image:tag
```