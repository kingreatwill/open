https://blog.csdn.net/xyz_dream/article/details/88372356
https://github.com/unixhot/salt-k8s


https://www.latelee.org/kubernetes/k8s-deploy-1.17.0-detail.html


kubernetes高可用安装（kubernetes install）工具，一条命令，离线安装，包含所有依赖，内核负载不依赖haproxy keepalived,纯golang开发,99年证书,支持v1.16.8 v1.15.11 v1.17.4 v1.18.0! https://sealyun.com
https://github.com/fanux/sealos

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

