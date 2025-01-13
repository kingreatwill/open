<!-- toc -->
[TOC]
参考1：https://github.com/eip-work/kuboard-press
参考2：https://github.com/lework/kainstall
kainstall = kubeadm install kubernetes
## Kubernetes IN Docker
https://github.com/kubernetes-sigs/kind
kind：是一种使用Docker容器节点运行本地Kubernetes集群的工具。该类型主要用于测试Kubernetes，但可用于本地开发或CI。
https://kind.sigs.k8s.io/docs/user/quick-start/#installation
### windows in docker
https://github.com/dockur/windows
### 下载
```
$ curl -Lo ./kind "https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64"
$ chmod +x ./kind
$ mv ./kind /usr/local/bin/kind

查看环境变量
env
or
echo $PATH

mac
brew install kind
win
choco install kind
```
### 创建集群，默认集群名称为 kind
$ kind create cluster
$ kind create cluster --name k8s-demo # 定义集群名称
$ kubectl cluster-info --context kind-kind
$ docker ps | grep kind
$ kubectl get pods -A -o wide # 列出K8S集群pods
$ kind get clusters # 查询集群
$ kind delete cluster # 删除集群
$ docker exec -it my-node-name crictl images # 列出集群镜像

### kubernetes client安装(kubectl)
选择对应版本下载
wget https://dl.k8s.io/v1.19.1/kubernetes-client-linux-amd64.tar.gz
tar -zxvf kubernetes-client-linux-amd64.tar.gz
cd kubernetes/client/bin
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

## kuboard
### 准备工作
k8s集群安装: k8s v1.16.1
环境准备：3台centos7.7  2核心4G
node1.centos7.com  192.168.1.135
node2.centos7.com  192.168.1.120
node3.centos7.com  192.168.1.110


### 1. 设置 hostname 解析
echo "127.0.0.1   $(hostname)" >> /etc/hosts

### 2. 设置固定IP

### 3. 安装 docker / kubelet
```
curl -sSL https://kuboard.cn/install-script/v1.16.1/install_kubelet.sh | sh

脚本作用:
设置固定IP  ->  关闭防火墙  -> 关闭selinux  -> 关闭swap ->
->
yum -y install wget
-> 创建检查点
```


### 4. 初始化 master 节点
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

### 5. 初始化 worker节点
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
k8s集群slave节点使用kubectl命令时The connection to the server localhost:8080 was refused - did you specify the right host or port?

解决方案1：
/etc/kubernetes/admin.conf 可以从主节点拷贝
```
[k8s@server1 ~]# mkdir -p $HOME/.kube
[k8s@server1 ~]# sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
[k8s@server1 ~]# sudo chown $(id -u):$(id -g) $HOME/.kube/config
```
解决方案2
```
出现这个问题的原因是kubectl命令需要使用kubernetes-admin来运行，解决方法如下，将主节点中的【/etc/kubernetes/admin.conf】文件拷贝到从节点相同目录下，然后配置环境变量：

echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> ~/.bash_profile
立即生效

source ~/.bash_profile
```



### 6. 安装 Ingress Controller
```
# 只在 master 节点执行
kubectl apply -f https://kuboard.cn/install-script/v1.16.2/nginx-ingress.yaml


卸载IngressController
# 只在 master 节点执行
kubectl delete -f https://kuboard.cn/install-script/v1.16.2/nginx-ingress.yaml
```

### 7. 安装Dashboard

1. 官网

```
kubectl delete ns kubernetes-dashboard

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml

修改
---

kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  type: NodePort
  ports:
    - port: 443
      targetPort: 8443
      nodePort: 31801
  selector:
    k8s-app: kubernetes-dashboard

---


kubectl apply -f recommended.yaml



kubectl get pod -n kubernetes-dashboard

删除pod
kubectl -n kubernetes-dashboard delete $(kubectl -n kubernetes-dashboard get pod -o name | grep dashboard)
```

2.
https://www.toutiao.com/a6761285771505697288
```
1. 
wget https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

2.
#image: k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1
image: mirrorgooglecontainers/kubernetes-dashboard-amd64:v1.10.1

3. 
NodePort

kubectl apply -f kubernetes-dashboard.yaml


kubectl get pod -n kube-system

kubectl -n kube-system get service -o wide
```

3. Dashboard V2.0(beta8)
https://www.cnblogs.com/bluersw/p/11747161.html


```
wget https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml
```
修改recommended.yaml文件内容(vi recommended.yaml)：

```
---
#增加直接访问端口
kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  type: NodePort #增加
  ports:
    - port: 443
      targetPort: 8443
      nodePort: 31801 #增加
  selector:
    k8s-app: kubernetes-dashboard

---
#因为自动生成的证书很多浏览器无法使用，所以我们自己创建，注释掉kubernetes-dashboard-certs对象声明
#apiVersion: v1
#kind: Secret
#metadata:
#  labels:
#    k8s-app: kubernetes-dashboard
#  name: kubernetes-dashboard-certs
#  namespace: kubernetes-dashboard
#type: Opaque

---
```

创建证书
```
mkdir dashboard-certs

cd dashboard-certs/

#创建命名空间
kubectl create namespace kubernetes-dashboard

# 创建key文件
openssl genrsa -out dashboard.key 2048

#证书请求
openssl req -days 36000 -new -out dashboard.csr -key dashboard.key -subj '/CN=dashboard-cert'

#自签证书
openssl x509 -req -in dashboard.csr -signkey dashboard.key -out dashboard.crt

#创建kubernetes-dashboard-certs对象
kubectl create secret generic kubernetes-dashboard-certs --from-file=dashboard.key --from-file=dashboard.crt -n kubernetes-dashboard
```

安装Dashboard
```
#安装
kubectl create -f  ~/recommended.yaml

#检查结果
kubectl get pods -A  -o wide

kubectl get service -n kubernetes-dashboard  -o wide
```

创建dashboard管理员
创建账号：
```
vi dashboard-admin.yaml
```
```
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: dashboard-admin
  namespace: kubernetes-dashboard
```
保存退出后执行
```
kubectl create -f dashboard-admin.yaml
```
为用户分配权限：
```
vi dashboard-admin-bind-cluster-role.yaml
```
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: dashboard-admin-bind-cluster-role
  labels:
    k8s-app: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: dashboard-admin
  namespace: kubernetes-dashboard
```
保存退出后执行
```
kubectl create -f dashboard-admin-bind-cluster-role.yaml
```

查看并复制用户Token：
```
kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep dashboard-admin | awk '{print $1}')
```
说明 ：创建账号和为用户分配权限 我放入yml文件里了
#### 7.x 总结
1. 创建证书
```
mkdir dashboard-certs

cd dashboard-certs/

#创建命名空间
kubectl create namespace kubernetes-dashboard

# 创建key文件
openssl genrsa -out dashboard.key 2048

#证书请求
openssl req -days 36000 -new -out dashboard.csr -key dashboard.key -subj '/CN=dashboard-cert'

#自签证书
openssl x509 -req -in dashboard.csr -signkey dashboard.key -out dashboard.crt

#创建kubernetes-dashboard-certs对象
kubectl create secret generic kubernetes-dashboard-certs --from-file=dashboard.key --from-file=dashboard.crt -n kubernetes-dashboard
```
2. 安装
```
kubectl apply -f  recommended.yaml
```

3. 查看token
```
kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep dashboard-admin | awk '{print $1}')
```

4. 登录
https://192.168.1.135:31801/

5. 说明
如果不安装metrics-server  ，dashboard是看不到CPU和内存使用情况 的。



### 8. 或者安装k8dash- 强大的k8s dashboard
```
kubectl apply -f https://raw.githubusercontent.com/herbrandson/k8dash/master/kubernetes-k8dash.yaml


kubectl get  -n kube-system deploy/k8dash svc/k8dash

# 生成token

kubectl create serviceaccount k8dash -n kube-system

kubectl create clusterrolebinding k8dash --clusterrole=cluster-admin --serviceaccount=kube-system:k8dash

kubectl get secret k8dash-token-kpt25 -n kube-system -o yaml | grep 'token:' | awk '{print $2}' | base64 -

# 生成token 官网

# Create the service account in the current namespace (we assume default)
kubectl create serviceaccount k8dash-sa

# Give that service account root on the cluster
kubectl create clusterrolebinding k8dash-sa --clusterrole=cluster-admin --serviceaccount=default:k8dash-sa

# Find the secret that was created to hold the token for the SA
kubectl get secrets

# Show the contents of the secret to extract the token
kubectl describe secret k8dash-sa-token-xxxxx



# nodeport
https://raw.githubusercontent.com/herbrandson/k8dash/master/kubernetes-k8dash-nodeport.yaml

```
