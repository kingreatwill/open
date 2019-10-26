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
[dcoker命令](https://www.runoob.com/docker/docker-command-manual.html)

docker run:
```
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

-a stdin: 指定标准输入输出内容类型，可选 STDIN/STDOUT/STDERR 三项；

-d: 后台运行容器，并返回容器ID；

-i: 以交互模式运行容器，通常与 -t 同时使用；

-P: 随机端口映射，容器内部端口随机映射到主机的高端口

-p: 指定端口映射，格式为：主机(宿主)端口:容器端口

-t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用；

--name="nginx-lb": 为容器指定一个名称；

--dns 8.8.8.8: 指定容器使用的DNS服务器，默认和宿主一致；

--dns-search example.com: 指定容器DNS搜索域名，默认和宿主一致；

-h "mars": 指定容器的hostname；

-e username="ritchie": 设置环境变量；

--env-file=[]: 从指定文件读入环境变量；

--cpuset="0-2" or --cpuset="0,1,2": 绑定容器到指定CPU运行；

-m :设置容器使用内存最大值；

--net="bridge": 指定容器的网络连接类型，支持 bridge/host/none/container: 四种类型；

--link=[]: 添加链接到另一个容器；

--expose=[]: 开放一个端口或一组端口；

--volume , -v: 绑定一个卷

--restart: docker服务启动时，自动启动容器，并且当容器停止时，尝试重启容器。

                --restart具体参数值详细信息：
                no -  容器退出时，不重启容器；
                on-failure - 只有在非0状态退出时才从新启动容器；
                always - 无论退出状态是如何，都重启容器

--privileged=true 管理权限
```
docker info : 显示 Docker 系统信息，包括镜像和容器数。
docker version :显示 Docker 版本信息。 -f :指定返回值的模板文件。
docker start :启动一个或多个已经被停止的容器
docker stop :停止一个运行中的容器
docker restart :重启容器
docker kill :杀掉一个运行中的容器。 -s :向容器发送一个信号 (docker kill -s KILL mynginx)

docker rm ：删除一个或多少容器 

    -f :通过SIGKILL信号强制删除一个运行中的容器
    -l :移除容器间的网络连接，而非容器本身
    -v :-v 删除与容器关联的卷

    实例
    强制删除容器db01、db02

    docker rm -f db01 db02
    移除容器nginx01对容器db01的连接，连接名db

    docker rm -l db 
    删除容器nginx01,并删除容器挂载的数据卷

    docker rm -v nginx01

docker pause :暂停容器中所有的进程。

docker unpause :恢复容器中所有的进程。



docker ps
等价于
docker container ps
等价于
docker container ls

容器中到底运行了哪些进程：
docker top container_name
如果想要看到更多信息比如状态、启动时间等等，可以加上-au：（当然可以参阅Linux ps命令的参数加上更多参数显示特定信息）

每个容器对于各种资源的使用情况：
docker stats
实时变化的列表，以显示每个容器实例的CPU使用率、内存使用量以及可用量等等。
docker stats exceptionless_api_1 
查看具体的容器　　

# docker三剑客  docker-machine compose swarm