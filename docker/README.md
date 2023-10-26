<!--toc-->
[TOC]
# docker tools
## dive
https://github.com/wagoodman/dive 22.5k
工具探究Docker 镜像的世界

[Play-with-docker（PWD）是一个Docker的演练场](https://labs.play-with-docker.com/)
PS: 如果有pull不下来的镜像, 登录上去pull 下来,推送到自己的docker hub上去, 然后再下载

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


清理无用镜像和容器:
```
docker rmi $(docker images -f "dangling=true" -q)

docker rm -v $(docker ps -a -q -f status=exited)
```

### docker删除未使用的容器、镜像
linux
```
# 删除 exited container
# docker rm -v $(docker ps -a -q -f status=exited)
#  删除没用的 image # dangling=true 按照中文来翻译的话，意思是指『悬空』的 image，我理解成『没有被使用的 image』。
# docker rmi $(docker images -f "dangling=true" -q)
#  删除没用的 volumn
# docker volume rm $(docker volume ls -qf dangling=true)
```
or linux & windows

https://docs.docker.com/engine/reference/commandline/container_prune/
https://docs.docker.com/engine/reference/commandline/image_prune/
https://docs.docker.com/engine/reference/commandline/system_prune/

```
docker container prune
docker image prune

docker system prune # 删除未使用的数据 Remove all unused containers, networks, images (both dangling and unreferenced), and optionally, volumes.
```


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

2. docker login -u user -p pwd harbor.xxx.com
or docker login harbor.xxx.com

git bash: winpty docker login -u demo harbor.xxx.com
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

## docker sleep
`docker run -d --name openresty-example openresty/openresty:1.13.6.2-2-xenial sleep 1234`
`sleep inf`  代表无限期的sleep

## docker build时使用host网络的方法
使用Dockerfile来docker build镜像时，默认使用的bridge网络环境；而RUN等命令经常需要联网下载依赖，由于公司加密软件的限制，造成RUN命令使用bridge时无法联网

于是想到使用host网络应该可以上网，host网络中，docker 容器没有自己的网卡和ip，不使用birdge网络，直接使用本机的网络;只要本机可以上网，docker build时的RUN命令就可以使用网络

`docker build --network=host -t test .`

## 不同网桥下的容器间通信
1. 先手动建立一个 bridge 模式的新 网桥， docker network  create --driver  bridge  --subnet=172.18.0.0/16  --gateway=172.18.0.1  new_bridge

2. docker network ls   可以查看 docker 下现在的网络模式（新加的那个）

3.  docker  run  -name test1  -ti  --net=new_bridge   镜像名   (用新网桥的一个容器 test1)

4. docker  run  -name test2  -ti   --net=bridge  镜像名   (用 docker 默认网桥的一个容器test2)

5. 进入到其中一个容器 ，ip a  查看网卡，ping  另一个容器IP  

6. 进另一个容器，同上。 两个容器IP段不一样。 不同网桥，会创建不同网段的虚拟网卡给容器　。

7. 不同网桥下的容器间 不能ping通， 在于docker 设计时候就隔离了不同网桥 　

8. docker  network  connect   new_bridge   test2   //  为  test2 容器添加一块 new_bridge的 虚拟网卡，这样test2 上会 创建一个新的虚拟网卡，网段就是 新网桥设置的。

9. 如此就能互相ping通。

## Docker之Cgroup对于CPU，内存，磁盘资源的限制
https://blog.csdn.net/boyuser/article/details/110487436

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
VNC (Virtual Network Console)是虚拟网络控制台的缩写，优点像windows版本的远程桌面控制，它是一款优秀的远程控制工具软件，由著名的 AT&T 的欧洲研究实验室开发的。VNC 是在基于 UNIX 和 Linux 操作系统的免费的开源软件，远程控制能力强大，高效实用，其性能可以和 Windows 和 MAC 中的任何远程控制软件媲美。 在 Linux 中，VNC 包括以下四个命令：vncserver，vncviewer，vncpasswd，和 vncconnect。大多数情况下用户只需要其中的两个命令：vncserver 和 vncviewer。

## bitnami
bitnami快速部署开源软件
BitNami是一个开源项目，专注于帮助开发人员快速部署和管理应用程序。

看支持哪些docker/k8s/vm : https://bitnami.com/stacks
https://github.com/bitnami/containers/tree/main/bitnami/

比如: [使用 bitnami/postgresql-repmgr 镜像快速设置 PostgreSQL HA](https://www.cnblogs.com/hacker-linner/p/16183165.html)
比如: [mysql主从](https://github.com/bitnami/containers/tree/main/bitnami/mysql) 也是非常方便的


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
https://github.com/amir20/dozzle 1k
Dozzle 是 Docker 容器的实时日志查看器。Dozzle 将能够从用户的容器中捕获所有日志并将其实时发送到用户的浏览器。Dozzle 不是数据库，它不存储或保存任何日志，使用 Dozzle 时只能看到实时日志。

### dive 分析docker镜像
https://github.com/wagoodman/dive 27.8k
> 用来探索 docker 镜像每一层文件系统，以及发现缩小镜像体积方法的命令行工具。启动命令：dive 镜像名

### diving 分析docker镜像web展示基于dive
https://github.com/vicanso/diving 185
> 基于 dive 分析 docker 镜像，界面化展示了镜像每层的变动（增加、修改、删除等）、用户层数据大小等信息。
> 便捷获取镜像信息和每层镜像内容的文件树，可以方便地浏览镜像信息。对于需要优化镜像体积时非常方便

### runlinke 查看指定docker container的启动命令
https://github.com/lavie/runlike

`pip install runlike`
`runlike <container-name>`

### dockerfilegraph
可视化多阶段Dockerfiles
https://github.com/patrickhoefler/dockerfilegraph

### skopeo
https://github.com/containers/skopeo 3.8k
除了基本的 inspect 之外，Skopeo 还提供了 skopeo copy 命令来复制镜像，可以直接在远程注册表之间复制镜像，无需将它们拉取到本地注册表。如果你使用了本地注册表，这个命令也可以作为拉取 / 推送的替代方案。

- 通过各种存储机制复制镜像，例如，可以在不需要特权的情况下将镜像从一个注册表复制到另一个注册表
- 检测远程镜像并查看其属性，包括其图层，无需将镜像拉到本地
- 从镜像库中删除镜像
- 当存储库需要时，skopeo 可以传递适当的凭据和证书进行身份验证

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
https://github.com/portainer/portainer 19.7k

### DockStation
桌面应用程序
https://github.com/DockStation/dockstation 1.7k

### Docker Desktop
桌面应用程序

### Lazydocker（UI终端）
https://github.com/jesseduffield/lazydocker 18.4k
> 带命令行 UI 的 docker 管理工具。可以通过点点点来管理 docker，却又不需要装 rancher 这样的企业级容器管理平台

### Docui (UI终端)
https://github.com/skanehira/docui 2.1k
> 终端 Docker 管理工具，自带一个终端界面。使用该工具可以方便的通过界面管理 docker 不用再记那些命令。

### Shipyard

## 其它
### 镜像的构建
- Buildah（https://buildah.io） podman build 子命令，它实际上是经过包装的 Buildah。
- 谷歌的 Kaniko（https://github.com/GoogleContainerTools/kaniko）
- buildkit（https://github.com/moby/buildkit）
下一代 docker build。它是 Moby 项目的一部分，在运行 Docker 时通过 DOCKER_BUILDKIT=1 docker build 就可以启用它，作为 Docker 的一个实验性特性。

- Source-To-Image (S2I，https://github.com/openshift/source-to-image) 是一个不使用 Dockerfile 直接从源代码构建镜像的工具包。这个工具在简单可预期的场景和工作流中表现良好，但如果你需要多一些定制化，或者你的项目没有预期的结构，那么它就会变得烦人和笨拙。如果你对 Docker 还不是很有信心，或者如果你在 OpenShift 集群上构建镜像，可能可以考虑使用 S2I，因为使用 S2I 构建镜像是它的一个内置特性。

- Jib（https://github.com/GoogleContainerTools/jib）是谷歌开发的一款工具，专门用于构建 Java 镜像。它提供了 Maven 和 Gradle 插件，可以让你轻松地构建镜像，不需要理会 Dockerfile。

- Bazel（https://github.com/bazelbuild/bazel），它是谷歌的另一款工具。它不仅用于构建容器镜像，而且是一个完整的构建系统。如果你只是想构建镜像，那么使用 Bazel 可能有点大材小用，但这绝对是一个很好的学习体验，所以如果你愿意，可以将 [rules_docker](https://github.com/bazelbuild/rules_docker) 为入手点。

### 容器运行时
- runc（https://github.com/opencontainers/runc）
- crun（https://github.com/containers/crun）
- CRI-O 实际上不是容器引擎，而是容器运行时
- containerd（https://containerd.io）
### 常用Dockerfile/各种Dockerfile 
https://github.com/jessfraz/dockerfiles
https://github.com/mritd/dockerfile
https://github.com/webdevops/Dockerfile
https://github.com/Kaixhin/dockerfiles
https://github.com/stilleshan/dockerfiles
https://github.com/einverne/dockerfile
https://github.com/vektorcloud

### Dockerfile 多阶段构建多个镜像

```
#
# BUILD ENVIRONMENT
# -----------------
ARG GO_VERSION
FROM golang:${GO_VERSION} as builder

RUN apt-get -y update && apt-get -y install upx

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY internal/ internal/
COPY webhooks/ webhooks/
COPY version/ version/
COPY cmd/ cmd/

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

# Do an initial compilation before setting the version so that there is less to
# re-compile when the version changes
RUN go build -mod=readonly "-ldflags=-s -w" ./...

ARG VERSION

# Compile all the binaries
RUN go build -mod=readonly -o manager main.go
RUN go build -mod=readonly -o proxy cmd/proxy/main.go
RUN go build -mod=readonly -o backup-agent cmd/backup-agent/main.go
RUN go build -mod=readonly -o restore-agent cmd/restore-agent/main.go

RUN upx manager proxy backup-agent restore-agent

#
# IMAGE TARGETS
# -------------
FROM gcr.io/distroless/static:nonroot as controller
WORKDIR /
COPY --from=builder /workspace/manager .
USER nonroot:nonroot
ENTRYPOINT ["/manager"]

FROM gcr.io/distroless/static:nonroot as proxy
WORKDIR /
COPY --from=builder /workspace/proxy .
USER nonroot:nonroot
ENTRYPOINT ["/proxy"]

FROM gcr.io/distroless/static:nonroot as backup-agent
WORKDIR /
COPY --from=builder /workspace/backup-agent .
USER nonroot:nonroot
ENTRYPOINT ["/backup-agent"]

FROM gcr.io/distroless/static as restore-agent
WORKDIR /
COPY --from=builder /workspace/restore-agent .
USER root:root
ENTRYPOINT ["/restore-agent"]
```
我们可以看到在这一个 Dockerfile 中我们使用多阶段构建定义了很多个 Targets，当我们在构建镜像的时候就可以通过 --target 参数来明确指定要构建的 Targets 即可，比如我们要构建 controller 这个目标镜像，则直接使用下面的命令构建即可：
```
docker build --target controller \
  --build-arg GO_VERSION=${GO_VERSION} \
  --build-arg VERSION=$(VERSION) \
  --tag ${DOCKER_REPO}/${DOCKER_IMAGE_NAME_PREFIX}controller:${DOCKER_TAG} \
  --file Dockerfile .
```

> `docker build --target proxy ...` or `--target backup-agent`,`--target restore-agent`
>
#### 使用 BuildKit 进行缓存
Docker 版本 18.09 引入了 BuildKit 作为对现有构建系统的彻底检查。大修背后的想法是提高性能、存储管理和安全性。我们可以利用 BuildKit 来保持多个构建之间的状态。这样，Maven 不会每次都下载依赖项，因为我们有永久存储。要在我们的 Docker 安装中启用 BuildKit，我们需要编辑daemon.json文件：
```
...
{
"features": {
    "buildkit": true
}}
...
```

> 在docker build命令前加上环境变量 `DOCKER_BUILDKIT=1 docker build .`


启用 BuildKit 后，我​​们可以将 Dockerfile 更改为：
```
FROM maven:alpine as build
ENV HOME=/usr/app
RUN mkdir -p $HOME
WORKDIR $HOME
ADD . $HOME
RUN --mount=type=cache,target=/root/.m2 mvn -f $HOME/pom.xml clean package

FROM openjdk:8-jdk-alpine
COPY --from=build /usr/app/target/single-module-caching-1.0-SNAPSHOT-jar-with-dependencies.jar /app/runner.jar
ENTRYPOINT java -jar /app/runner.jar
```
当我们更改代码或pom.xml文件时，Docker 将始终执行 ADD 和 RUN Maven 命令。首次运行时构建时间将是最长的，因为 Maven 必须下载依赖项。随后的运行将使用本地依赖项并执行得更快。

这种方法需要维护 Docker 卷作为依赖项的存储。有时，我们必须强制 Maven 使用Dockerfile 中的-U标志更新我们的依赖项。


前端
```
FROM node:alpine as builder

WORKDIR /app

COPY package.json /app/

RUN --mount=type=cache,target=/app/node_modules,id=my_app_npm_module,sharing=locked \
    --mount=type=cache,target=/root/.npm,id=npm_cache \
    npm i --registry=https://registry.npm.taobao.org

COPY src /app/src

RUN --mount=type=cache,target=/app/node_modules,id=my_app_npm_module,sharing=locked \
    npm run build
```

[使用 BuildKit 构建镜像](https://docker-practice.github.io/zh-cn/buildx/buildkit.html)
[使用 BuildKit 构建镜像](https://vuepress.mirror.docker-practice.com/buildx/buildkit/#run-mount-type-cache)

## docker资源限制
[docker对于CPU和内存的限制](https://www.cnblogs.com/renshengdezheli/p/16662622.html)

限制512M内存
`docker run -it --rm -m 512m -v /memload:/memload hub.c.163.com/library/centos:latest`

--cpuset-cpus=0 设置容器里的进程都运行在0号CPU上
`docker run -it --rm  --cpuset-cpus=0 -v /memload:/memload hub.c.163.com/library/centos:latest`

限制一个半的 CPU
`docker run -it --rm  --cpus="1.5" -v /memload:/memload hub.c.163.com/library/centos:latest`

常用的参数有：
参数	 |	参数解释
---|---
-m或者--memory=	 |	容器可以使用的最大内存量。如果设置此选项，则允许的最小值为6m（6 兆字节）。也就是说，您必须将该值设置为至少 6 兆字节。
--memory-swap*	 |	允许此容器交换到磁盘的内存量。
--memory-swappiness	 |	默认情况下，主机内核可以换出容器使用的一定百分比的匿名页面。您可以设置--memory-swappiness为 0 到 100 之间的值，以调整此百分比。
--memory-reservation	 |	允许您指定一个小于--memory在 Docker 检测到主机上的争用或内存不足时激活的软限制。如果使用--memory-reservation，则必须将其设置为低于--memory它才能优先。因为是软限制，所以不保证容器不超过限制。
--kernel-memory	 |	容器可以使用的最大内核内存量。允许的最小值是4m。因为内核内存不能被换出，内核内存不足的容器可能会阻塞主机资源，这会对主机和其他容器产生副作用。
--oom-kill-disable	 |	默认情况下，如果发生内存不足 (OOM) 错误，内核会终止容器中的进程。要更改此行为，请使用该--oom-kill-disable选项。仅在您还设置了该-m/--memory选项的容器上禁用 OOM kill。如果-m未设置该标志，主机可能会耗尽内存，内核可能需要终止主机系统的进程以释放内存。

参数 |	参数解释
---|---
--cpus=	 |	指定容器可以使用多少可用 CPU 资源。例如，如果主机有两个 CPU，并且您设置--cpus="1.5"了 ，则容器最多可以保证一个半的 CPU。这相当于设置--cpu-period="100000"和--cpu-quota="150000"。
--cpu-period=	 |	指定 CPU CFS 调度程序周期，它与 --cpu-quota. 默认为 100000 微秒（100 毫秒）。大多数用户不会更改默认设置。对于大多数用例，--cpus是一种更方便的选择。
--cpu-quota=	 |	对容器施加 CPU CFS 配额。--cpu-period容器在被限制之前被限制的每微秒数。因此充当有效上限。对于大多数用例，--cpus是一种更方便的选择。
--cpuset-cpus	 |	限制容器可以使用的特定 CPU 或内核。如果您有多个 CPU，则容器可以使用的逗号分隔列表或连字符分隔的 CPU 范围。第一个 CPU 编号为 0。有效值可能是0-3（使用第一个、第二个、第三个和第四个 CPU）或1,3（使用第二个和第四个 CPU）。
--cpu-shares	 |	将此标志设置为大于或小于默认值 1024 的值，以增加或减少容器的重量，并允许它访问或多或少比例的主机 CPU 周期。这仅在 CPU 周期受到限制时才会强制执行。当有足够多的 CPU 周期可用时，所有容器都会根据需要使用尽可能多的 CPU。这样，这是一个软限制。--cpu-shares不会阻止容器以 swarm 模式调度。它优先考虑可用 CPU 周期的容器 CPU 资源。它不保证或保留任何特定的 CPU 访问权限。

> [cpu测试脚本](../linux/shell/cpuload.sh) ;  [内存测试脚本](../linux/shell/cpuload.sh)


## 流水线版本
```

GIT_COMMIT_ID=$(git rev-parse --short HEAD)

GIT_BRANCH_OR_TAG=$(git symbolic-ref --short -q HEAD || git describe --tags --exact-match 2> /dev/null)
GIT_BRANCH_OR_TAG=$(basename "$GIT_BRANCH_OR_TAG")

DATETIME=$(date +"%Y%m%d%H%M%S")

VERSION="${DATETIME}-${GIT_BRANCH_OR_TAG//_/-}-${GIT_COMMIT_ID}"
echo ${VERSION} > .VERSION


# echo "${DATETIME}-${GIT_BRANCH_OR_TAG//_/-}-${GIT_COMMIT_ID}" > .VERSION # 输出到.VERSION文件, 如果文件存在,会覆盖
# VERSION=$(cat .VERSION) # 取版本信息到变量
# VERSION=$(head -1 .VERSION) # 取版本信息到变量

```