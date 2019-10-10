# centos7 安装docker
yum 包更新到最新:
```
sudo yum update
```

移除旧的版本：
```
sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-selinux \
                  docker-engine-selinux \
                  docker-engine
```
安装一些必要的系统工具：
```
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
```
添加软件源信息：
```
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```
更新 yum 缓存：
```
sudo yum makecache fast
```
安装 Docker-ce：
```
sudo yum -y install docker-ce
```

启动Docker服务并激活开机启动：
```
systemctl start docker & systemctl enable docker
```

脚本安装：
```
$ curl -fsSL https://get.docker.com -o get-docker.sh
$ sudo sh get-docker.sh
```
镜像加速:
```
http://hub-mirror.c.163.com
https://4yl89dki.mirror.aliyuncs.com

 vi /etc/docker/daemon.json

{
  "registry-mirrors": ["http://hub-mirror.c.163.com"]
}
```
删除 Docker CE:
```
$ sudo yum remove docker-ce
$ sudo rm -rf /var/lib/docker
```