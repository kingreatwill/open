<!--toc-->
[TOC]
# docker tools
## dive
https://github.com/wagoodman/dive 22.5k
工具探究Docker 镜像的世界
# docker私有仓库
## Harbor
Habor是由VMWare公司开源的容器镜像仓库。事实上，Habor是在Docker Registry上进行了相应的企业级扩展，从而获得了更加广泛的应用，这些新的企业级特性包括：管理用户界面，基于角色的访问控制 ，AD/LDAP集成以及审计日志等，足以满足基本企业需求。
官方地址：https://vmware.github.io/harbor/cn/
https://goharbor.io/
https://github.com/goharbor/harbor

## 官方registry
docker pull registry
docker run -itd -v /data/registry:/var/lib/registry -p 5000:5000 --restart=always --name registry registry:latest 

https://github.com/docker/distribution

docker login --username=xx --password=xx xx.com

UI
https://hub.docker.com/r/joxit/docker-registry-ui

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
                - no -  容器退出时，不重启容器；no是默认策略，在任何情况下都不会restart容器
                - on-failure - 只有在非0状态退出时才从新启动容器；on-failure表示如果容器 exit code异常时将restart，如果容器exit code正常将不做任何处理。
                - on-failure:max-retries，max-retries表示最大重启次数。如（--restart on-failure:5）
                - always - 无论退出状态是如何，都重启容器
                - unless-stopped和 always 基本一样，只有一个场景 unless-stopped有点特殊：如果容器正常stopped，然后机器重启或docker服务重启，这种情况下容器将不会被restart

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

如果已经启动了则可以使用如下命令：
docker update --restart=always `<CONTAINER ID>`

## docker push到私有仓库

1. 修改/root/.docker/config.json
{ "insecure-registries":["harbor.xxx.com"] }
不添加报错，https证书问题

2. docker login -u user harbor.xxx.com
or docker login harbor.xxx.com

git bash: winpty docker login -u demo images.lingcb.net
需要添加winpty


3. push
```

docker build -t xxxx:tag .
docker tag xxxx:tag harbor.xxx.com/project/xxxx:tag
docker push harbor.xxx.com/project/xxxx:tag


#1.标记镜像
docker tag {镜像名}:{tag} {Harbor地址}:{端口}/{Harbor项目名}/{自定义镜像名}:{自定义tag}
#eg:docker tag vmware/harbor-adminserver:v1.1.0 192.168.2.108:5000/test/harbor-adminserver:v1.1.0

#2.push 到Harbor
docker push {Harbor地址}:{端口}/{自定义镜像名}:{自定义tag}
#eg:docker push 192.168.2.108:5000/test/harbor-adminserver:v1.1.0

3.pull 到本地
docker pull 192.168.2.108：5000/test/harbor-adminserver:v1.1.0
```

# docker三剑客  docker-machine compose swarm

# 有趣的docker项目

https://github.com/bcicen/ctop
> 实现了类 top 命令展示效果的 docker 容器监控工具

https://github.com/drone/drone
> 一个基于 Docker 的持续集成平台，使用 Go 语言编写

https://github.com/LockGit/gochat
> 纯 Go 实现的轻量级即时通讯系统。技术上各层之间通过 rpc 通讯，使用 redis 作为消息存储与投递的载体，相对 kafka 操作起来更加方便快捷。
> 各层之间基于 etcd 服务发现，在扩容部署时将会方便很多。架构、目录结构清晰，文档详细。而且还提供了 docker 一件构建，安装运行十分方便，推荐作为学习项目

https://github.com/pipiliang/docker-dashboard
> 基于控制台的 docker 工具，代码简单易读，可以做为学习 Node.js 的实践项目

https://github.com/yeasy/docker_practice
> Docker 从入门到实践

## 用docker创建ubuntuVNC桌面


https://github.com/fcwu/docker-ubuntu-vnc-desktop
docker运行ubuntu
```
docker pull dorowu/docker-ubuntu-vnc-desktop
docker run -p 30007:80 dorowu/ubuntu-desktop-lxde-vnc

# 使用VNC Viewer或者浏览器登录容器
docker run -itd -p 6080:80 -p 5900:5900 dorowu/ubuntu-desktop-lxde-vnc:bionic
# 浏览器访问6080端口或者下载VNC Viewer，然后在VNC Viewer中通过5900端口访问容器

# 创建容器时给网页版以及VNC Viewer版添加登录容器的密码
docker run -itd -p 6080:80 -p 5900:5900  -e HTTP_PASSWORD=mypassword  -e VNC_PASSWORD=mypassword dorowu/ubuntu-desktop-lxde-vnc:bionic

# 创建容器时设置容器分辨率以及添加新用户
docker run -itd -p 6080:80 -p 5900:5900  -e RESOLUTION=1920x1080 -e USER=zs -e PASSWORD=mypassword -e HTTP_PASSWORD=mypassword  -e VNC_PASSWORD=mypassword dorowu/ubuntu-desktop-lxde-vnc:bionic
```
### 名词解释
VNC (Virtual Network Console)是虚拟网络控制台的缩写，优点像windows版本的远程桌面控制

## docker 分析
### docker image转dockerfile
https://github.com/P3GLEG/Whaler
```
FROM golang:1.14.4 AS builder
WORKDIR $GOPATH
RUN go get -u github.com/P3GLEG/Whaler
WORKDIR $GOPATH/src/github.com/P3GLEG/Whaler
RUN export CGO_ENABLED=0 && go build .
RUN cp Whaler /root/Whaler

FROM alpine:3.12.0
WORKDIR /root/
COPY --from=builder /root/Whaler .
ENTRYPOINT ["./Whaler"]
```
docker run -t --rm -v /var/run/docker.sock:/var/run/docker.sock:ro pegleg/whaler -sV=1.36 nginx:latest
or
alias whaler="docker run -t --rm -v /var/run/docker.sock:/var/run/docker.sock:ro pegleg/whaler"
whaler -sV=1.36 nginx:latest
or
自己编译
./whaler -sV=1.36 nginx:latest

### Dozzle  实时日志查看器的docker容器
https://github.com/amir20/dozzle
Dozzle 是 Docker 容器的实时日志查看器。Dozzle 将能够从用户的容器中捕获所有日志并将其实时发送到用户的浏览器。Dozzle 不是数据库，它不存储或保存任何日志，使用 Dozzle 时只能看到实时日志。

### dive 分析docker镜像
https://github.com/wagoodman/dive 23k
> 用来探索 docker 镜像每一层文件系统，以及发现缩小镜像体积方法的命令行工具。启动命令：dive 镜像名

### diving 分析docker镜像web展示基于dive
https://github.com/vicanso/diving 147
> 基于 dive 分析 docker 镜像，界面化展示了镜像每层的变动（增加、修改、删除等）、用户层数据大小等信息。
> 便捷获取镜像信息和每层镜像内容的文件树，可以方便地浏览镜像信息。对于需要优化镜像体积时非常方便

### docker-slim 自动缩减 docker 镜像
https://github.com/docker-slim/docker-slim
> 自动缩减 docker 镜像的体积的工具。大幅度缩减 docker 镜像的体积，方便分发，使用命令 docker-slim build --http-probe your-name/your-app。

### Trivy 镜像漏洞扫描工具
https://github.com/aquasecurity/trivy
> 镜像漏洞扫描工具Trivy

Trivy是一种适用于CI的简单而全面的容器漏洞扫描程序。软件漏洞是指软件或操作系统中存在的故障、缺陷或弱点。Trivy检测操作系统包（Alpine、RHEL、CentOS等）和应用程序依赖（Bundler、Composer、npm、yarn等）的漏洞。Trivy很容易使用，只要安装二进制文件，就可以扫描了。扫描只需指定容器的镜像名称。与其他镜像扫描工具相比，例如Clair，Anchore Engine，Quay相比，Trivy在准确性、方便性和对CI的支持等方面都有着明显的优势。

推荐在CI中使用它，在推送到container registry之前，您可以轻松地扫描本地容器镜像，Trivy具备如下的特征：

1. 检测面很全，能检测全面的漏洞，操作系统软件包（Alpine、Red Hat Universal Base Image、Red Hat Enterprise Linux、CentOS、Oracle Linux、Debian、Ubuntu、Amazon Linux、openSUSE Leap、SUSE Enterprise Linux、Photon OS 和Distrioless）、应用程序依赖项（Bundler、Composer、Pipenv、Poetry、npm、yarn和Cargo）；
2. 使用简单，仅仅只需要指定镜像名称；
3. 扫描快且无状态，第一次扫描将在10秒内完成（取决于您的网络）。随后的扫描将在一秒钟内完成。与其他扫描器在第一次运行时需要很长时间（大约10分钟）来获取漏洞信息，并鼓励您维护持久的漏洞数据库不同，Trivy是无状态的，不需要维护或准备；
4. 易于安装，安装方式：
  apt-get install
  yum install
  brew install

### ctop
实时监控类似linux的top
https://github.com/bcicen/ctop

### 查看一个运行容器的docker run启动参数

https://github.com/lavie/runlike 453
不安装的方法：`docker run --rm -v /var/run/docker.sock:/var/run/docker.sock assaflavie/runlike YOUR-CONTAINER`
可以使用-q自动换行

https://github.com/nexdrew/rekcod 237
不安装的方法：`docker run --rm -v /var/run/docker.sock:/var/run/docker.sock nexdrew/rekcod <container>`

> docker ps -a --no-trunc 查看启动的脚本参数和完整的命令

## docker GUI

### Portainer
https://github.com/portainer/portainer 16.2k

### DockStation
桌面应用程序
https://github.com/DockStation/dockstation

### Docker Desktop
桌面应用程序

### Lazydocker（UI终端）
https://github.com/jesseduffield/lazydocker
> 带命令行 UI 的 docker 管理工具。可以通过点点点来管理 docker，却又不需要装 rancher 这样的企业级容器管理平台

### Docui (UI终端)
https://github.com/skanehira/docui
> 终端 Docker 管理工具，自带一个终端界面。使用该工具可以方便的通过界面管理 docker 不用再记那些命令。

### Shipyard
