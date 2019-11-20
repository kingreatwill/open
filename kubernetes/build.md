
[TOC]
# 编译和运行Kubernetes源码
## 为什么要编译源码 
Kubernetes是一个非常棒的容器集群管理平台。通常情况下，我们并不需要修改K8s代码即可直接使用。但如果，我们在环境中发现了某个问题/缺陷，或按照特定业务需求需要修改K8s代码时，如定制Kubelet的StopContainer 逻辑、kube-scheduler的pod调度逻辑等。为了让修改生效，那么就需要编译K8s代码了。

Kubernetes源码编译，大致分为本地二进制可执行文件编译和docker镜像编译两种。由于在我们的环境中，Kubernetes是由Docker容器方式运行的。故此我们需要采用后面一种方式编译，即镜像编译。

由于Kubernetes每个组件服务的镜像Dockerfile文件是由Kubernetes源码自动生成的，因此，社区并未提供每个组件的镜像Dockerfile文件。编译本地二进制可执行文件很简单，也更直接。而docker镜像编译资料却很少，且碍于某种特殊网络原因，会导致失败。此处，将介绍如何顺利的完成k8s镜像编译。
## 安装依赖 
- 安装Golang
```
wget -c https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz -P /opt/
cd /opt/
tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz 
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile && source /etc/profile
```
- 指定分支，下载 Kubernetes 源代码（默认$GOPATH目录为/root/go/）
```
mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
git clone https://github.com/kubernetes/kubernetes -b release-1.16.3
cd $GOPATH/src/k8s.io/kubernetes
```
[windows需要安装编译cgo相关命令](../tools/MinGW-64.md)
## 本地二进制文件编译Kubernetes（方法一） 
修改运行平台配置参数（可选）
根据自己的运行平台（linux/amd64)修改hack/lib/golang.sh，把KUBE_SERVER_PLATFORMS，KUBE_CLIENT_PLATFORMS和KUBE_TEST_PLATFORMS中除linux/amd64以外的其他平台注释掉，以此来减少编译所用时间。

编译源码
进入Kubernetes根目录下
```
cd kubernetes
```
KUBE_BUILD_PLATFORMS指定目标平台，WHAT指定编译的组件，通过GOFLAGS和GOGCFLAGS传入编译时参数，如此处编译kubelet 组件。
```
KUBE_BUILD_PLATFORMS=linux/amd64 make all WHAT=cmd/kubelet GOFLAGS=-v GOGCFLAGS="-N -l"
```
    如果不指定WHAT，则编译全部。
    make all是在本地环境中进行编译的。
    make release和make quick-release在容器中完成编译、打包成docker镜像。
    编译kubelet这部分代码，也可执行make clean && make WHAT=cmd/kubelet
检查编译成果
编译过程较长，请耐心等待，编译后的文件在kubernetes/_output里。

或者进入cmd/kubelet (以kubelet为例子)
执行go build -v命令,如果没出错,会生成可执行文件kubelet
```
go build
```
生成的可执行文件在当前文件夹下面

```
# ls cmd/kubelet/
app BUILD kubelet kubelet.go OWNERS
```

## Docker镜像编译Kubernetes（方法二）
查看kube-cross的TAG版本号
```
# cat ./build/build-image/cross/VERSION
v1.13.4-1
```
查看debian_iptables_version版本号
```
# egrep -Rn "debian_iptables_version=" ./
./build/common.sh:93:  local debian_iptables_version=v12.0.1
```

这里，我使用DockerHub的Auto build功能，来构建K8s镜像。自然将编译需要用到的base镜像
```
docker pull xx/pause-amd64:3.1
docker pull xx/kube-cross:v1.13.4-1
docker pull xx/debian-base-amd64:v2.0.0
docker pull xx/debian-iptables-amd64:v12.0.1
docker pull xx/debian-hyperkube-base-amd64:0.12.1

docker tag xx/pause-amd64:3.1 k8s.gcr.io/pause-amd64:3.1
docker tag xx/kube-cross:v1.13.4-1 k8s.gcr.io/kube-cross:v1.13.4-1
docker tag xx/debian-base-amd64:v2.0.0 k8s.gcr.io/debian-base-amd64:v2.0.0
docker tag xx/debian-iptables-amd64:v12.0.1 k8s.gcr.io/debian-iptables-amd64:v12.0.1
docker tag xx/debian-hyperkube-base-amd64:0.12.1 k8s.gcr.io/debian-hyperkube-base-amd64:0.12.1
```

把build/lib/release.sh中的–pull去掉，避免构建镜像继续拉取镜像：
```
"${DOCKER[@]}" build --pull -q -t "${docker_image_tag}" ${docker_build_path} >/dev/null
修改为:
 "${DOCKER[@]}" build -q -t "${docker_image_tag}" ${docker_build_path} >/dev/null
```

编辑文件hack/lib/version.sh
将KUBE_GIT_TREE_STATE=”dirty” 改为 KUBE_GIT_TREE_STATE=”clean”，确保版本号干净。
执行编译命令
```
# cd kubernetes
# make clean
# KUBE_BUILD_PLATFORMS=linux/amd64 KUBE_BUILD_CONFORMANCE=n KUBE_BUILD_HYPERKUBE=n make release-images GOFLAGS=-v GOGCFLAGS="-N -l"
```
其中KUBE_BUILD_PLATFORMS=linux/amd64指定目标平台为linux/amd64，GOFLAGS=-v开启verbose日志，GOGCFLAGS=”-N -l”禁止编译优化和内联，减小可执行程序大小。

编译的K8s Docker镜像以压缩包的形式发布在_output/release-tars目录中

```
# ls _output/release-images/amd64/
cloud-controller-manager.tar  kube-controller-manager.tar  kube-scheduler.tar
kube-apiserver.tar            kube-proxy.tar
```

### 使用编译镜像
等待编译完成后，在_output/release-stage/server/linux-amd64/kubernetes/server/bin/目录下保存了编译生成的二进制可执行程序和docker镜像tar包。如导入kube-apiserver.tar镜像，并更新环境上部署的kube-apiserver镜像。
```
# docker load -i kube-apiserver.tar
# docker tag k8s.gcr.io/kube-apiserver:v1.16.3  registry.com/kube-apiserver:v1.16.3
```
整个编译过程结束后，现在就可以到master节点上，修改/etc/kubernetes/manifests/kube-apiserver.yaml描述文件中的image，修改完立即生效。

[索引](https://www.kubernetes.org.cn/5033.html)
[索引](https://github.com/kubernetes/kubernetes/tree/master/build/)