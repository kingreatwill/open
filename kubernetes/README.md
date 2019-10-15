中文文档: https://rootsongjc.gitbooks.io/kubernetes-handbook
中文文档: https://jimmysong.io/kubernetes-handbook/
中文文档: https://github.com/rootsongjc/kubernetes-handbook
中文文档: https://www.kubernetes.org.cn/k8s
中文文档: https://kubernetes.io/zh/docs/
所有安装方式列表：https://kubernetes.io/docs/setup/

中文二进制文件安装方式文档： https://github.com/opsnull/follow-me-install-kubernetes-cluster

英文二进制文件安装方式文档： https://github.com/kelseyhightower/kubernetes-the-hard-way

使用Ansible脚本安装文档：https://github.com/gjmzj/kubeasz （在自建机房,建议使用此方法）

k8s-kubeasz-阿里云vpc部署记录: https://li-sen.github.io/2018/09/27/k8s-kubeasz-阿里云vpc部署记录/

kops在AWS中国区安装k8s文档：

https://github.com/nwcdlabs/kops-cn （建议使用此方法，该项目由aws中国区维护）

https://github.com/kubernetes/kops/blob/master/docs/aws-china.md

http://senlinzhan.github.io/2018/01/11/k8s-on-aws/

https://blog.csdn.net/cloudvtech/article/details/80539086

中文社区1： https://www.kubernetes.org.cn

中文社区2：http://dockone.io/

kubernetes相关的工具：https://github.com/kubernetes-sigs/ （你会发现很多有实用的工具可以使用）

这里有一个在aws上管理k8s的pdf文档：https://s3.cn-north-1.amazonaws.com.cn/sides-share/AWS+Webinar+2018/PDF/EKS+Webinar+Chinese.pdf

这里介绍Terraform 管理 EKS：https://www.hashicorp.com/blog/hashicorp-announces-terraform-support-aws-kubernetes

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
7. 配置 Ingress
说明：如果测试 Istio，不需要安装 Ingress
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud-generic.yaml

验证:
kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx
```
测试示例应用
部署测试应用，详情参见社区[文章](https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ingress-guide-nginx-example.html)
```
kubectl create -f sample/apple.yaml
kubectl create -f sample/banana.yaml
kubectl create -f sample/ingress.yaml
```
测试示例应用
```
$ curl -kL http://localhost/apple
apple
$ curl -kL http://localhost/banana
banana
```
删除示例应用
```
kubectl delete -f sample/apple.yaml
kubectl delete -f sample/banana.yaml
kubectl delete -f sample/ingress.yaml
```
删除 Ingress
```
kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud-generic.yaml
kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
```

8. 安装 [Helm](https://github.com/helm/helm/blob/master/docs/install.md)
[docs](https://helm.sh/docs/using_helm/#quickstart-guide)

```
如果在后续使用 helm 安装组件的过程中出现版本兼容问题，可以参考 通过二进制包安装 思路安装匹配的版本

# Use Chocolatey on Windows
# 注：安装的时候需要保证网络能够访问googleapis这个域名
choco install kubernetes-helm

# Install Tiller into your Kubernetes cluster
helm init --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.14.1 --skip-refresh

# Change helm repo
helm repo add stable http://mirror.azure.cn/kubernetes/charts-incubator/

# Update charts repo (Optional)
helm repo update
```
通过二进制包安装
```
# Download binary release
在 https://github.com/helm/helm/releases 中找到匹配的版本并下载(需要梯子), 如: https://storage.googleapis.com/kubernetes-helm/helm-v2.14.1-darwin-amd64.tar.gz

# Unpack

tar -zxvf helm-v2.0.0-linux-amd64.tgz

# Move it to its desired destination

mv darwin-amd64/helm /usr/local/bin/helm
```

9. 配置 Istio (待验证)
说明：Istio Ingress Gateway和Ingress缺省的端口冲突，请移除Ingress并进行下面测试
可以根据文档安装 Istio https://istio.io/docs/setup/kubernetes/

9.1 通过 下载 Istio 1.2.4 并安装 CLI
```
curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.2.4 sh -
cd istio-1.2.4/
export PATH=$PWD/bin:$PATH
```
在Windows上，您可以手工下载Istio安装包，或者把getLatestIstio.ps1拷贝到你希望下载 Istio 的目录，并执行 - 说明：根据社区提供的安装脚本修改而来
```
getLatestIstio.ps1

param(
    [string] $IstioVersion = "0.3.0"
)

$url = "https://github.com/istio/istio/releases/download/$($IstioVersion)/istio_$($IstioVersion)_win.zip"
$Path = Get-Location
$output = [IO.Path]::Combine($Path, "istio_$($IstioVersion)_win.zip”)
    
Write-Host "Downloading Istio from $url to path " $Path -ForegroundColor Green 
    
#Download file
(New-Object System.Net.WebClient).DownloadFile($url, $output)
    
# Unzip the Archive
Expand-Archive $output -DestinationPath $Path
    
#Set the environment variable
$IstioHome = [IO.Path]::Combine($Path, "istio-$($IstioVersion)")
    
[Environment]::SetEnvironmentVariable("ISTIO_HOME", "$IstioHome", "User")

```
``` powershell
.\getLatestIstio.ps1
```

9.2 通过 Helm chart 安装 Istio
```
# 安装 istio-init chart 安装所有的 Istio CRD
helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system

# 验证下安装的 Istio CRD 个数, 应该安装23个CRD
kubectl get crds | grep 'istio.io\|certmanager.k8s.io' | wc -l

# 开始 istio chart 安装
helm install install/kubernetes/helm/istio --name istio --namespace istio-system
```

查看 istio 发布状态
```
helm status istio
```
为 default 名空间开启自动 sidecar 注入
```
kubectl label namespace default istio-injection=enabled
kubectl get namespace -L istio-injection
```
安装 Book Info 示例
```
kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
```
确认示例应用在运行中
```
export GATEWAY_URL=localhost:80
curl -o /dev/null -s -w "%{http_code}\n" http://${GATEWAY_URL}/productpage
```
可以通过浏览器访问

http://localhost/productpage

说明：如果当前80端口已经被占用或保留，我们可以编辑 install/kubernetes/helm/istio/values.yaml 文件中 Gateway 端口进行调整，比如将 80 端口替换为 8888 端口

```
      ## You can add custom gateway ports
    - port: 8888  # Changed from 80
      targetPort: 80
      name: http2
      nodePort: 31380
```
然后执行如下命令并生效
```
kubectl delete service istio-ingressgateway -n istio-system
helm upgrade istio install/kubernetes/helm/istio
```
删除实例应用
```
samples/bookinfo/platform/kube/cleanup.sh
```
卸载 Istio
```
helm del --purge istio
kubectl delete -f install/kubernetes/helm/istio/templates/crds.yaml -n istio-system
```

# Minikube 安装k8s
https://yq.aliyun.com/articles/221687

# centos7 安装单机版k8s

# 安装集群：手动安装
https://github.com/yonyoucloud/install_k8s
https://github.com/unixhot/salt-k8s

# 安装集群：工具安装