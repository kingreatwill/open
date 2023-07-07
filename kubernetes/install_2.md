https://blog.csdn.net/xyz_dream/article/details/88372356
https://github.com/unixhot/salt-k8s


https://www.latelee.org/kubernetes/k8s-deploy-1.17.0-detail.html

## sealer安装k8s
https://github.com/alibaba/sealer 468 golang

- 极其简单的方式在生产环境中或者离线环境中安装 kubernetes、以及 kubernetes 生态中其它软件
- 通过 Kubefile 可以非常简单的自定义 kubernetes 集群镜像对集群和应用进行打包，并可以提交到仓库中进行分享
- 强大的生命周期管理能力，以难以想象的简单的方式去做如集群升级，集群备份恢复，节点阔缩等操作
- 速度极快 3min 以内完成集群安装
- 支持 ARM x86, v1.20 以上版本支持 containerd，几乎兼容所有支持 systemd 的 linux 操作系统
- 不依赖 ansible haproxy keepalived, 高可用通过 ipvs 实现，占用资源少，稳定可靠
- 官方仓库中有非常多的生态软件镜像可以直接使用，包含所有依赖，一键安装

```
#---------安装一个 kubernetes 集群
#安装sealer
wget https://github.com/alibaba/sealer/releases/download/v0.1.4/sealer-0.1.4-linux-amd64.tar.gz && \
tar zxvf sealer-0.1.4-linux-amd64.tar.gz && mv sealer /usr/bin
#运行集群
sealer run kubernetes:v1.19.9 # 在公有云上运行一个kubernetes集群
sealer run kubernetes:v1.19.9 --masters 3 --nodes 3 # 在公有云上运行指定数量节点的kuberentes集群
# 安装到已经存在的机器上
sealer run kubernetes:v1.19.9 --masters 192.168.0.2,192.168.0.3,192.168.0.4 --nodes 192.168.0.5,192.168.0.6,192.168.0.7 --passwd xxx

#--------安装 prometheus 集群
sealer run prometheus:2.26.0
# 上面命令就可以帮助你安装一个包含 prometheus 的 kubernetes 集群, 同理其它软件如 istio ingress grafana 等都可以通过这种方式运行。
```

[集群镜像：实现高效的分布式应用交付](https://blog.csdn.net/alisystemsoftware/article/details/117322535)
[骚操作，这款工具可以把Kubernetes集群打包成一个镜像](https://mp.weixin.qq.com/s/goa9LuKS5HdJC8Q03XNKyA)

## sealos安装k8s
https://github.com/labring/sealos

sealos已经不紧紧是可以安装k8s

下载
```
 # 下载并安装 sealos, sealos 是个 golang 的二进制工具，直接下载拷贝到 bin 目录即可, release 页面也可下载 
 yum install wget && yum install tar &&\
 wget  https://github.com/labring/sealos/releases/download/v4.2.3/sealos_4.2.3_linux_amd64.tar.gz  && \
 tar -zxvf sealos_4.2.3_linux_amd64.tar.gz sealos &&  chmod +x sealos && mv sealos /usr/bin
```
安装
```sh
# Run a single node kubernetes
$ sealos run labring/kubernetes:v1.25.0 labring/calico:v3.24.1

# Run a HA kubernetes cluster
$ sealos run labring/kubernetes:v1.25.0 labring/calico:v3.24.1
      --masters 192.168.64.2,192.168.64.22,192.168.64.20 
      --nodes 192.168.64.21,192.168.64.19 -p [your-ssh-passwd]
# helm
$ sealos run labring/kubernetes:v1.25.0 labring/helm:v3.8.2 labring/calico:v3.24.1 \
      --masters 192.168.0.2,192.168.0.3\
      --nodes 192.168.0.4 -p [your-ssh-passwd]

# 使用 docker 作为 runtime：
$ sealos run labring/kubernetes-docker:v1.20.5-4.1.3 labring/calico:v3.24.1 \
      --masters 192.168.64.2,192.168.64.22,192.168.64.20 \
      --nodes 192.168.64.21,192.168.64.19 -p [your-ssh-passwd]


# Add masters or nodes
$ sealos add --masters 192.168.64.20 --nodes 192.168.64.21,192.168.64.22

# Delete your cluster
$ sealos reset
```

### sealos安装k8s - old
https://github.com/fanux/sealos 4.1k golang

kubernetes高可用安装（kubernetes install）工具，一条命令，离线安装，包含所有依赖，内核负载不依赖haproxy keepalived,纯golang开发,99年证书,支持v1.16.8 v1.15.11 v1.17.4 v1.18.0! https://sealyun.com


```
# 下载并安装sealos, sealos是个golang的二进制工具，直接下载拷贝到bin目录即可
wget https://github.com/fanux/sealos/releases/download/v3.2.0/sealos && \
    chmod +x sealos && mv sealos /usr/bin 

# 下载离线资源包
wget https://sealyun.oss-cn-beijing.aliyuncs.com/d551b0b9e67e0416d0f9dce870a16665-1.18.0/kube1.18.0.tar.gz 

# 安装一个三master的kubernetes集群
sealos init --passwd 123456 \
	--master 192.168.0.2  --master 192.168.0.3  --master 192.168.0.4  \
	--node 192.168.0.5 \
	--pkg-url /root/kube1.18.0.tar.gz \
	--version v1.18.0


增加master

sealos join --master 192.168.0.6 --master 192.168.0.7
sealos join --master 192.168.0.6-192.168.0.9  # 或者多个连续IP
增加node

sealos join --node 192.168.0.6 --node 192.168.0.7
sealos join --node 192.168.0.6-192.168.0.9  # 或者多个连续IP
删除指定master节点

sealos clean --master 192.168.0.6 --master 192.168.0.7
sealos clean --master 192.168.0.6-192.168.0.9  # 或者多个连续IP
删除指定node节点

sealos clean --node 192.168.0.6 --node 192.168.0.7
sealos clean --node 192.168.0.6-192.168.0.9  # 或者多个连续IP
清理集群

sealos clean
```


https://github.com/kubernetes-sigs/kubespray

![k8s](../img/K8S.png)

