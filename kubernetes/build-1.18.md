<!--toc-->
[TOC]
# 编译k8s v1.18源码并部署

## 准备工作
### 前提
- golang 1.13以上
- docker 17.03.2-ce 以上

#### 安装go
1. yum安装
yum list golang --showduplicates | sort -r

查看是否最新版本

2. 二进制安装
```
wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz

# 解压到/usr/local
tar -C /usr/local -xzf go1.14.1.linux-amd64.tar.gz
# 添加环境变量/usr/local/go/bin
# vi /etc/profile
vi $HOME/.profile
export PATH=$PATH:/usr/local/go/bin

source $HOME/.profile

# 验证
go version
```

### 下载源码
```
git clone https://github.com/kubernetes/kubernetes

git checkout tags/v1.18.0
```
### 依赖镜像
在编译过程中会用到以下三个镜像，但是docker pull命令是无法下载到这些镜像的

1. us.gcr.io/k8s-artifacts-prod/build-image/kube-cross:v1.13.9-2      (1.14版本以前是这个k8s.gcr.io/kube-cross:v1.12.12-1，从1.15就换了)
2. k8s.gcr.io/debian-iptables-amd64:v12.0.1（这个镜像包含debian-base-amd64）
3. k8s.gcr.io/debian-base-amd64:v2.0.0

#### 下载镜像
`k8s.gcr.io`可以转换为`gcr.io/google_containers`
k8s.gcr.io/pause-amd64:3.1
gcr.io/google_containers/pause-amd64:3.1
```
Azure 中国镜像 https://gcr.azk8s.cn
# https://github.com/Azure/container-service-for-azure-china/blob/master/aks/README.md#22-container-registry-proxy
# http://mirror.azk8s.cn/help/gcr-proxy-cache.html
# docker pull gcr.azk8s.cn/google_containers/kube-cross:v1.13.9-2 好像不能用


docker pull registry.aliyuncs.com/google_containers/kube-cross:v1.13.9-2 目前没有
docker pull registry.aliyuncs.com/google_containers/debian-iptables-amd64:v12.0.1
docker pull registry.aliyuncs.com/google_containers/debian-base-amd64:v2.0.0

docker tag registry.aliyuncs.com/google_containers/kube-cross:v1.13.9-2 us.gcr.io/k8s-artifacts-prod/build-image/kube-cross:v1.13.9-2
docker tag registry.aliyuncs.com/google_containers/debian-iptables-amd64:v12.0.1 k8s.gcr.io/debian-iptables-amd64:v12.0.1
docker tag registry.aliyuncs.com/google_containers/debian-base-amd64:v2.0.0 k8s.gcr.io/debian-base-amd64:v2.0.0

docker rmi registry.aliyuncs.com/google_containers/kube-cross:v1.13.9-2
docker rmi registry.aliyuncs.com/google_containers/debian-iptables-amd64:v12.0.1
docker rmi registry.aliyuncs.com/google_containers/debian-base-amd64:v2.0.0
```

#### 为什么需要kube-cross 
kubernetes/blob/build/common.sh中
417行`function kube::build::build_image() {`方法（release-images.sh中需要用到）
424行 `  cp "${KUBE_ROOT}/build/build-image/Dockerfile" "${LOCAL_OUTPUT_BUILD_CONTEXT}/Dockerfile"` 中需要

在 `/kubernetes/blob/master/build/build-image/Dockerfile`中可以看到，
版本在`/kubernetes/blob/master/build/build-image/cross/VERSION`可以看到

#### 为什么需要debian-base和debian-iptables 
kubernetes/blob/build/common.sh中
95行`kube::build::get_docker_wrapped_binaries() {`方法中可以看到需要debian-base和debian-iptables 以及它们的版本



#### 为什么需要自己下载基础镜像 build/lib/release.sh
357行 KUBE_BUILD_PULL_LATEST_IMAGES 是否拉取最新的镜像 要为n 不然会拉取上面的几个镜像


## 编译源码

KUBE_GIT_VERSION=v1.8.0 KUBE_FASTBUILD=true KUBE_GIT_TREE_STATE=clean KUBE_BUILD_PLATFORMS=linux/amd64 KUBE_BUILD_PULL_LATEST_IMAGES=n make release-images


KUBE_BUILD_PLATFORMS=linux/amd64 KUBE_BUILD_CONFORMANCE=n KUBE_BUILD_HYPERKUBE=n make release-images


KUBE_BUILD_PLATFORMS=linux/amd64 KUBE_BUILD_CONFORMANCE=n KUBE_BUILD_HYPERKUBE=n make release-images GOFLAGS=-v GOGCFLAGS="-N -l"
### 你需要知道的

KUBE_BUILD_CONFORMANCE参数用来控制是否创建一致性测试镜像，
KUBE_BUILD_HYPERKUBE控制是否创建hyperkube镜像(各种工具集成在一起)，这两个目前都用不上，因此是设置为"n"表示不构建；
> 构建hyperkube需要k8s.gcr.io/debian-hyperkube-base-amd64基础镜像


KUBE_GIT_TREE_STATE=clean
修改完记得提交，若编译前，对代码有改动，且未提交,即未执行`git commit`操作,
执行编译`KUBE_GIT_VERSION=v1.8.0 KUBE_FASTBUILD=true  KUBE_BUILD_PULL_LATEST_IMAGES=n make release-images`,
则生成的镜像即版本为dirty版本

> 编辑文件hack/lib/version.sh将KUBE_GIT_TREE_STATE=”dirty” 改为 KUBE_GIT_TREE_STATE=”clean”，确保版本号干净。

GOFLAGS=-v开启verbose日志，GOGCFLAGS=”-N -l”禁止编译优化和内联，减小可执行程序大小。


### 最终生成
#### 镜像
```
ls _output/release-images/amd64/
conformance-amd64.tar  kube-apiserver.tar           kube-proxy.tar
hyperkube-amd64.tar    kube-controller-manager.tar  kube-scheduler.tar
```

如果已经安装k8s集群，需要修改镜像
例如：kube-apiserver
```
docker load -i kube-apiserver.tar
docker tag k8s.gcr.io/kube-apiserver:v1.13.3
修改文件/etc/kubernetes/manifests/kube-apiserver.yaml，将
修改完毕后，执行命令kubectl apply -f kube-apiserver.yaml使修改生效；
```


#### 二进制执行文件
```
ls _output/dockerized/bin/linux/amd64/
apiextensions-apiserver  e2e.test            genyaml     hyperkube                kubelet         mounter
conversion-gen           gendocs             ginkgo      kubeadm                  kubemark        openapi-gen
deepcopy-gen             genkubedocs         go2make     kube-apiserver           kube-proxy
defaulter-gen            genman              go-bindata  kube-controller-manager  kube-scheduler
e2e_node.test            genswaggertypedocs  go-runner   kubectl                  linkcheck
# 将kubelet、kubeadm、kubectl拷贝到/usr/bin/目录。
```

### 维护周期
Kubernetes 发现版本的通常只维护支持九个月，同时也适用于 kubeadm


## 系统配置

### 服务器准备

修改主机名，写入三台主机/etc/hosts
192.168.31.179 k8s-master
192.168.31.180 k8s-node1
192.168.31.181 k8s-node2

### 关闭防火墙
```
systemctl stop firewalld && systemctl disable firewalld
```

### 关闭交换分区
```
# 临时禁用swap即可，重启后，需要再此执行，当然也可以把磁盘信息写入/etc/fstab
swapoff -a
#（2）永久关闭swap分区
sed -ri 's/.*swap.*/#&/' /etc/fstab
```

kubernetes关闭swap主要是为了性能考虑,当然如果不想关闭swap，需要：
```
1.编辑/etc/sysconfig/kubelet ,添加KUBELET_EXTRA_ARGS="–fail-swap-on=false"
#cat /etc/sysconfig/kubelet
KUBELET_EXTRA_ARGS="–fail-swap-on=false"
2.初始化：
#kubeadm init --kubernetes-version=版本–pod-network-cidr=pod网络 --service-cidr=生成服务所在网络地址 --ignore-preflight-errors=Swap
```

### 禁用SElinux
```
sed -i 's/enforcing/disabled/' /etc/selinux/config && setenforce 0
```
或者
```
# 临时禁用selinux
setenforce 0
# 永久禁用selinux
vi /etc/selinux/config
SELINUX=disabled
```
重启

### 时间同步
```
#yum -y install ntp ntpdate
ntpdate cn.pool.ntp.org
# ntpdate ntp1.aliyun.com
定时任务
crontab -e
* */1 * * * /usr/sbin/ntpdate ntp1.aliyun.com

*　　*　　*　　*　　*　　command 
分　时　日　月　周　命令 
第1列表示分钟1～59 每分钟用*或者 */1表示 
第2列表示小时1～23（0表示0点） 
第3列表示日期1～31 
第4列表示月份1～12 
第5列标识号星期0～6（0表示星期天） 
第6列要运行的命令 
```
### 内核配置,解除网络限制

###PS:建立桥接时文件不存在错误用 modprobe br_netfilter

```
cat <<EOF >  /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sysctl -p /etc/sysctl.d/k8s.conf
# sysctl --system
```
or?
```
cat <<EOF >  /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF
```
or
```
# centos7 下需要给网卡建立桥接
vim /etc/sysctl.conf
vm.swappiness = 0
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
 
sysctl -p

###PS:建立桥接时文件不存在错误用 modprobe br_netfilter
```

### 容器引擎安装
安装依赖命令行

yum install -y conntrack ipvsadm ipset jq sysstat curl iptables libseccomp socat

yum -y install bash-completion.noarch net-tools vim lrzsz wget tree screen lsof tcpdump unzip tree 

#### 配置kubernetes的yum文件（源码安装的不需要）
```
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF
```

```
查看已安装的yum源
#yum repolist

更新元数据
#yum makecache
```


#### 配置docker
安装docker-ce的软件包依赖
yum install -y yum-utils  device-mapper-persistent-data  lvm2

```
cd /etc/yum.repos.d
[root@localhost yum.repos.d]# wget http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
[root@ken ~]# yum install docker-ce -y
```
or

```
yum-config-manager \
--add-repo \
http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

yum install docker-ce -y

# 通过yum list docker-ce --showduplicates | sort -r 获取版本信息，执行以下命令安装指定版本
# yum install docker-ce-<VERSION\_STRING> docker-ce-cli-<VERSION\_STRING> containerd.i
```

```
mkdir /etc/docker

vi /etc/docker/daemon.json
{
  "registry-mirrors": ["https://4yl89dki.mirror.aliyuncs.com"]
}

systemctl restart docker && systemctl enable docker

```

> 对于使用systemd作为的Linux的发行版，使用systemd作为docker的cgroup-river
可以确保服务器节点在资源紧张的情况更加稳定，因此这里修改各节点上docker的cgroup-driver为systemd。
创建/etc/docker/daemon.json
```
{
  "exec-opts": "native.cgroupdriver=systemd"
}
```
> 重启docker服务`systemctl restart docker`
> 确认修改生效 `docker info | grep Cgroup` 输出Cgroup Driver: systemd

### 组件安装
#### 安装kubelet kubectl kubeadm
yum install kubelet kubectl kubeadm -y
> kubelet 运行在 Cluster 所有节点上，负责启动 Pod 和容器。
> kubeadm 用于初始化 Cluster。
> kubectl 是 Kubernetes 命令行工具。通过 kubectl 可以部署和管理应用，查看各种资源，创建、删除和更新各种组件。


#### 集群初始化(master节点)
```
kubeadm init  --apiserver-advertise-address=192.168.31.179 --image-repository registry.aliyuncs.com/google_containers   --kubernetes-version v1.18.0  --service-cidr=10.96.0.0/12  --pod-network-cidr=10.244.0.0/16
```
等待2分钟

- –image-repository string：这个用于指定从什么位置来拉取镜像（1.13版本才有的），默认值是`k8s.gcr.io`，我们将其指定为国内镜像地址：`registry.aliyuncs.com/google_containers`
- –kubernetes-version string：指定kubenets版本号，默认值是stable-1，会导致从`https://dl.k8s.io/release/stable-1.txt`下载最新的版本号，我们可以将其指定为固定版本（v1.18.0）来跳过网络请求。
- –apiserver-advertise-address **主节点IP** ,指明用 Master 的哪个 interface 与 Cluster 的其他节点通信。如果 Master 有多个 interface，建议明确指定，如果不指定，kubeadm 会自动选择有默认网关的 interface。
- –pod-network-cidr指定 Pod 网络的范围。Kubernetes 支持多种网络方案，而且不同网络方案对  –pod-network-cidr有自己的要求，这里设置为10.244.0.0/16 是因为我们将使用 flannel 网络方案，必须设置成这个 CIDR。[--pod-network-cidr=10.244.0.0/16   请与flannel网络插件保持一致（flannel网络插件默认为10.244.0.0/16）]


##### 加入节点
主节点执行 `kubeadm token create  --print--join--command`得到
`kubeadm join 192.168.31.179:6443 --token qzwubw.bumwc63v5lreao8b     --discovery-token-ca-cert-hash sha256:a67d7d346ecdf72b68914eddd3581535ad60f97648b23093810e8d5a99afbba2`
将这条命令在其它node上执行

```
如果这个时候再想添加进来这个node，需要执行两步操作
第一步：停掉kubelet(需要添加进来的节点操作)

[root@host1 ~]# systemctl stop kubelet
 

第二步：删除相关文件

[root@host1 ~]# rm -rf /etc/kubernetes/*
 

第三步：添加节点

[root@host1 ~]# kubeadm join 172.20.10.2:6443 --token rn816q.zj0crlasganmrzsr --discovery-token-ca-cert-hash sha256:e339e4dbf6bd1323c13e794760fff3cbeb7a3f6f42b71d4cb3cffdde72179903
 

第四步：查看节点

[root@ken ~]# kubectl get nodes
NAME    STATUS   ROLES    AGE   VERSION
host1   Ready    <none>   13s   v1.13.2
host2   Ready    <none>   17m   v1.13.2
ken     Ready    master   53m   v1.13.2
 

忘掉token再次添加进k8s集群
 

第一步：主节点执行命令

获取token

[root@ken-master ~]# kubeadm token list
TOKEN                     TTL       EXPIRES                     USAGES                   DESCRIPTION                                                EXTRA GROUPS
ojxdod.fb7tqipat46yp8ti   10h       2019-05-06T04:55:42+08:00   authentication,signing   The default bootstrap token generated by 'kubeadm init'.   system:bootstrappers:kubeadm:default-node-token
 

第二步： 获取ca证书sha256编码hash值

[root@ken-master ~]# openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
2f8888cdb01191ff6dbca0edb02dbb21a14469028e4ff2598854a4544c5fa751
 

第三步：从节点执行如下的命令

[root@ken-node1 ~]# systemctl stop kubelet
 

第四步：删除相关文件

[root@ken-node1 ~]# rm -rf /etc/kubernetes/*
 

第五步：加入集群

指定主节点IP，端口是6443

在生成的证书前有sha256:

[root@ken-node1 ~]# kubeadm join 192.168.64.10:6443 --token ojxdod.fb7tqipat46yp8ti  --discovery-token-ca-cert-hash sha256:2f8888cdb01191ff6dbca0edb02dbb21a14469028e4ff2598854a4544c5fa751
```

##### 每个节点启动kubelet
```
systemctl restart kubelet && systemctl enable kubelet
```

##### 配置kubectl
```
[root@ken ~]#  mkdir -p $HOME/.kube
[root@ken ~]#  cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
[root@ken ~]# chown $(id -u):$(id -g) $HOME/.kube/config
# 为了使用更便捷，启用 kubectl 命令的自动补全功能。
[root@ken ~]# echo "source <(kubectl completion bash)" >> ~/.bashrc
```
> 如果node节点需要执行kubectl，则需要从master节点拷贝admin.conf文件

#### 主节点部署网络插件flannel部署
```
下载部署文件
wget -P k8s/ https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

修改文件
sed -i 's/quay.io/quay-mirror.qiniu.com/g' k8s/kube-flannel.yml

部署
kubectl apply -f k8s/kube-flannel.yml
```

### 如果初始化失败，请使用如下代码清除后重新初始化
```
# kubeadm reset
# ifconfig cni0 down
# ip link delete cni0
# ifconfig flannel.1 down
# ip link delete flannel.1
# rm -rf /var/lib/cni/
# rm -rf /var/lib/etcd/*
```

## 升级
```
# yum update kube* -y      # 升级kubernetes版本到1.18
# kubeadm upgrade plan  # 检查集群是否处于可升级状态，并以用户友好的方式获取可升级的版本。
# kubeadm upgrade apply v1.18.0  # 根据要求直接升级 y通过
```

## 参考
```
k8s.gcr.io/kube-proxy                v1.18.0             43940c34f24f        11 days ago         117MB
k8s.gcr.io/kube-apiserver            v1.18.0             74060cea7f70        11 days ago         173MB
k8s.gcr.io/kube-controller-manager   v1.18.0             d3e55153f52f        11 days ago         162MB
k8s.gcr.io/kube-scheduler            v1.18.0             a31f78c7c8ce        11 days ago         95.3MB
k8s.gcr.io/pause                     3.2                 80d28bedfe5d        7 weeks ago         683kB
k8s.gcr.io/coredns                   1.6.7               67da37a9a360        2 months ago        43.8MB
k8s.gcr.io/etcd                      3.4.3-0             303ce5db0e90        5 months ago        288MB
calico/node                          v3.8.2              11cd78b9e13d        7 months ago        189MB
calico/cni                           v3.8.2              c71c24a0b1a2        7 months ago        157MB
calico/kube-controllers              v3.8.2              de959d4e3638        7 months ago        46.8MB
calico/pod2daemon-flexvol            v3.8.2              96047edc008f        7 months ago        9.37MB


docker pull registry.aliyuncs.com/google_containers/kube-apiserver:v1.18.0
docker pull registry.aliyuncs.com/google_containers/kube-proxy:v1.18.0
docker pull registry.aliyuncs.com/google_containers/kube-controller-manager:v1.18.0 
docker pull registry.aliyuncs.com/google_containers/kube-scheduler:v1.18.0 
docker pull registry.aliyuncs.com/google_containers/etcd:3.4.3-0 
docker pull registry.aliyuncs.com/google_containers/coredns:1.6.7
docker pull registry.aliyuncs.com/google_containers/pause:3.2 

修改tag

docker tag registry.aliyuncs.com/google_containers/kube-apiserver:v1.18.0 k8s.gcr.io/kube-apiserver:v1.18.0
docker tag registry.aliyuncs.com/google_containers/kube-proxy:v1.18.0 k8s.gcr.io/kube-proxy:v1.18.0
docker tag registry.aliyuncs.com/google_containers/kube-controller-manager:v1.18.0 k8s.gcr.io/kube-controller-manager:v1.18.0
docker tag registry.aliyuncs.com/google_containers/kube-scheduler:v1.18.0 k8s.gcr.io/kube-scheduler:v1.18.0
docker tag registry.aliyuncs.com/google_containers/etcd:3.4.3-0 k8s.gcr.io/etcd: 3.4.3-0
docker tag registry.aliyuncs.com/google_containers/coredns:1.6.7 k8s.gcr.io/coredns: 1.6.7
docker tag registry.aliyuncs.com/google_containers/pause:3.2 k8s.gcr.io/pause:3.2
```
[kebuadm 搭建k8s v1.18](https://www.cnblogs.com/wanglaing-q123/p/12627919.html)
[ubuntu18.04 安装部署k8s 1.18.0版本](https://www.cnblogs.com/zliW/p/12603536.html)

[Kubernetes深入学习之二:编译和部署镜像(api-server)](https://blog.csdn.net/boling_cavalry/article/details/88603293)
[github和dockerhub制作k8s镜像](https://www.cnblogs.com/fengzhihai/p/9849683.html)
[Kubernetes展望与思考之1.17初体验](https://cloud.tencent.com/developer/article/1512662)

[Compile Kubernetes binaries](https://docs.microsoft.com/en-us/virtualization/windowscontainers/kubernetes/compiling-kubernetes-binaries)

