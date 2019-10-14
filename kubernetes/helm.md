https://mp.weixin.qq.com/s?__biz=MzI5ODQ2MzI3NQ==&mid=2247488132&idx=1&sn=173a975384b4a5befc0a7b8660952dec
## 概    述
Helm是kubernetes包管理工具，可以方便快捷的安装、管理、卸载kubernetes应用，类似于Linux操作系统中yum或apt-get软件的作用。其主要的设计目的：

+ 创建新的chart包
+ 将charts包文件打包压缩
+ 同chart仓库进行集成，获取charts文件
+ 安装及卸载charts到kubernetes集群
+ 管理通过helm安装的charts应用

## 概念介绍

**chart**: 一个 Helm 包，其中包含了运行一个应用所需要的镜像、依赖和资源定义等，还可能包含 Kubernetes 集群中的服务定义。

**release**：在 Kubernetes 集群上运行的 Chart 的一个实例。在同一个集群上，一个 Chart 可以安装很多次，每次安装都会创建一个新的 release。

**repository**：用于发布和存储 Chart 的仓库，Helm客户端通过HTTP协议来访问仓库中Chart的索引文件和压缩包。

## 组    件

**helm**: 提供给用户的客户端程序，可以以命令行的形式同服务端-tiller进行通信。

**tiller**：服务端软件，用来同helm客户端进行交互，并同kubernetes api server组件进行交互。

架构如下：![helm](../img/helm.webp)

## 安装部署

1. helm的安装部署
* 版本下载，版本列表 https://github.com/helm/helm/releases
* 解压缩, tar -zxvf helm-v2.0.0-linux-amd64.tgz
* 将解压缩后的二进制文件放在可执行目录下 mv linux-amd64/helm /usr/local/bin/helm，然后执行 helm --help查看帮助文档

2. tiller的安装部署
控制台执行 > helm init命令，该命令会将从charts仓库中下载charts包，并根据其中的配置部署至kubernetes集群。
默认的charts仓库为 https://kubernetes-charts.storage.googleapis.com/index.yaml
默认使用的tiller镜像为 gcr.io/kubernetes-helm/tiller:v2.13.1 
国内由于墙的原因无法直接访问，需要我们自行处理可替代的仓库和镜像版本，通过如下命令进行helm服务端的安装部署：
```
> helm init --tiller-image registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.13.1 --stable-repo-url https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts

Creating /root/.helm/repository/repositories.yaml
Adding stable repo with URL: https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
Adding local repo with URL: http://127.0.0.1:8879/charts
$HELM_HOME has been configured at /root/.helm.

Tiller (the Helm server-side component) has been installed into your Kubernetes Cluster.

Please note: by default, Tiller is deployed with an insecure 'allow unauthenticated users' policy.
To prevent this, run `helm init` with the --tiller-tls-verify flag.
For more information on securing your installation see: https://docs.helm.sh/using_helm/#securing-your-helm-installation
Happy Helming!
```
稍等一会然后执行如下命令，看到如下输出说明安装成功：
```
helm version
Client: &version.Version{SemVer:"v2.13.1", GitCommit:"618447cbf203d147601b4b9bd7f8c37a5d39fbb4", GitTreeState:"clean"}
Server: &version.Version{SemVer:"v2.13.1", GitCommit:"618447cbf203d147601b4b9bd7f8c37a5d39fbb4", GitTreeState:"clean"}
```

通过执行 helm --help 可以看到常用的命令，说明如下：
* search 在helm仓库进行查找应用
* fetch 从仓库中下载chart包到本地
* list 在该k8s集群的部署的release列表
* status 显示release的具体信息
* install 安装charts
* inspect 描述charts信息
* delete 删除部署的release
* create 创建一个charts
* package 将某一charts进行打包压缩
* repo 显示、添加、移除charts仓库

## 访问授权
在上面的步骤中我们将tiller所需的资源部署到了kubernetes集群中，但是由于Deployment tiller-deploy没有定义授权的ServiceAccount导致访问apiserver拒绝，执行如下命令为tiller-deploy进行授权：
```
> kubectl create serviceaccount --namespace kube-system tiller
> kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
> kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'
```

## 通过helm部署WordPress
输入如下命令，我们可以通过helm创建一个WordPress博客网站
```
> helm install --name wordpress-test --set "persistence.enabled=false,mariadb.persistence.enabled=false" stable/wordpress
```
通过如下命令获取登录信息：
```
> kubectl get svc -o wide
> kubectl get secret --namespace default wordpress-test-wordpress -o jsonpath="{.data.wordpress-password}" | base64 --decode
```
在浏览器中打开页面，并输入用户名和密码就可以看到搭建好的WordPress博客网站了。

## 升    级
当有新的chart包发布时或者想改变已有release的配置时，可以通过 helm upgrade命令实现，比如：
```
> helm upgrade wordpress-test \
> --set "persistence.enabled=true,mariadb.persistence.enabled=true" \
> stable/wordpress
```

参考文档：
https://helm.sh/docs/ 
https://yq.aliyun.com/articles/159601