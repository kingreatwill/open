参考：https://github.com/eip-work/kuboard-press

k8s集群安装: k8s v1.16.1
环境准备：3台centos7.7  2核心4G
node1.centos7.com  192.168.1.135
node2.centos7.com  192.168.1.120
node3.centos7.com  192.168.1.110

准备工作
1. 设置 hostname 解析
echo "127.0.0.1   $(hostname)" >> /etc/hosts

2. 设置固定IP

3. 安装 docker / kubelet
```
curl -sSL https://kuboard.cn/install-script/v1.16.1/install_kubelet.sh | sh

脚本作用:
设置固定IP  ->  关闭防火墙  -> 关闭selinux  -> 关闭swap ->
->
yum -y install wget
-> 创建检查点
```


4. 初始化 master 节点
```
# 只在 master 节点执行
# 替换 x.x.x.x 为 master 节点实际 IP（请使用内网 IP）
# export 命令只在当前 shell 会话中有效，开启新的 shell 窗口后，如果要继续安装过程，请重新执行此处的 export 命令
export MASTER_IP=192.168.1.135
# 替换 apiserver.name 为 您想要的 dnsName
export APISERVER_NAME=apiserver.name
# Kubernetes 容器组所在的网段，该网段安装完成后，由 kubernetes 创建，事先并不存在于您的物理网络中
export POD_SUBNET=10.100.0.1/16
echo "${MASTER_IP}    ${APISERVER_NAME}" >> /etc/hosts
curl -sSL https://kuboard.cn/install-script/v1.16.1/init_master.sh | sh
```
检查 master 初始化结果
```
# 只在 master 节点执行

# 执行如下命令，等待 3-10 分钟，直到所有的容器组处于 Running 状态
watch kubectl get pod -n kube-system -o wide

# 查看 master 节点初始化结果
kubectl get nodes -o wide
```

kubernetes增加污点，达到pod是否能在做节点运行
```
master node参与工作负载 (只在主节点执行)
使用kubeadm初始化的集群，出于安全考虑Pod不会被调度到Master Node上，也就是说Master Node不参与工作负载。

这里搭建的是测试环境可以使用下面的命令使Master Node参与工作负载：
k8s是master节点的hostname
允许master节点部署pod，使用命令如下:

kubectl taint nodes --all node-role.kubernetes.io/master-
1
输出如下:

node "k8s" untainted
1
输出error: taint “node-role.kubernetes.io/master:” not found错误忽略。

禁止master部署pod

kubectl taint nodes k8s node-role.kubernetes.io/master=true:NoSchedule

```

5. 初始化 worker节点
```
# 只在 master 节点执行
kubeadm token create --print-join-command

输出：kubeadm join apiserver.demo:6443 --token mpfjma.4vjjg8flqihor4vt     --discovery-token-ca-cert-hash sha256:6f7a8e40a810323
```
针对所有的 worker 节点执行
```
# 只在 worker 节点执行
# 替换 x.x.x.x 为 master 节点实际 IP（请使用内网 IP）
export MASTER_IP=x.x.x.x
# 替换 apiserver.demo 为初始化 master 节点时所使用的 APISERVER_NAME
export APISERVER_NAME=apiserver.demo
echo "${MASTER_IP}    ${APISERVER_NAME}" >> /etc/hosts

# 替换为 master 节点上 kubeadm token create 命令的输出
kubeadm join apiserver.demo:6443 --token mpfjma.4vjjg8flqihor4vt     --discovery-token-ca-cert-hash sha256:6f7a8e40a8
```

检查初始化结果
```
# 只在 master 节点执行
kubectl get nodes -o wide
```

移除 worker 节点
正常情况下，您无需移除 worker 节点，如果添加到集群出错，您可以移除 worker 节点，再重新尝试添加

在准备移除的 worker 节点上执行
```
# 只在 worker 节点执行
kubeadm reset
```

在 master 节点上执行
```
# 只在 master 节点执行
kubectl delete node demo-worker-x-x
```

6. 安装 Ingress Controller
```
# 只在 master 节点执行
kubectl apply -f https://kuboard.cn/install-script/v1.16.2/nginx-ingress.yaml


卸载IngressController
# 只在 master 节点执行
kubectl delete -f https://kuboard.cn/install-script/v1.16.2/nginx-ingress.yaml
```










