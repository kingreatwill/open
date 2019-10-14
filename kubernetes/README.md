# Windows 10 docker fro desktop 安装k8s
1. 首先配置Docker daemon 配置 Docker Hub 的中国官方镜像加速 https://registry.docker-cn.com 并开启 4GB 或更多内存

2. 拉取镜像
```
git clone https://github.com/AliyunContainerService/k8s-for-docker-desktop
```
3. 拉取k8s所需要的镜像
cd k8s-for-docker-desktop
可以通过修改 images.properties 文件加载你自己需要的镜像
```
./load_images.ps1
```
说明: 如果因为安全策略无法执行 PowerShell 脚本，请在 “以管理员身份运行” 的 PowerShell 中执行 Set-ExecutionPolicy RemoteSigned 命令。我用的powershell core，非自带的ps

4. 开启 Kubernetes

5. 配置 Kubernetes
可选操作: 切换Kubernetes运行上下文至 docker-for-desktop
```
kubectl config use-context docker-for-desktop
```
验证 Kubernetes 集群状态
```
kubectl cluster-info
kubectl get nodes
```

6. 部署 Kubernetes dashboard
https://github.com/kubernetes/dashboard


```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

开启 API Server 访问代理
kubectl proxy

访问：http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login
```


配置 kubeconfig 
```
$TOKEN=((kubectl -n kube-system describe secret default | Select-String "token:") -split " +")[1]
kubectl config set-credentials docker-for-desktop --token="${TOKEN}"


config 路径
%UserProfile%\.kube\config

注意如果出现：Not enough data to create auth info structure.

1. 网上
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')
配置kube的config文件，将刚才生成的token: 放在最后。（ token: 后面有个空格 ，不然会报:错误）

2. 如果config文件中有token，则复制token，点击令牌登录
```


7. 安装 Helm



# centos7 安装单机版k8s

# 安装集群：手动安装
https://github.com/yonyoucloud/install_k8s
https://github.com/unixhot/salt-k8s

# 安装集群：工具安装