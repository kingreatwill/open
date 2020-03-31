<!--toc-->
[TOC]

- - [1. Kubernetes 环境准备](#span-id-inline-toc-1-span-kubernetes-环境准备)- [前提条件](#前提条件)
        - [安装 kubernetes 集群](#安装-kubernetes-集群)
        
    - [2. 下载 Istio 部署文件](#span-id-inline-toc-2-span-下载-istio-部署文件)- [开启 istioctl 的自动补全功能](#开启-istioctl-的自动补全功能)- [bash](#bash)
            - [zsh](#zsh)
            
        
    - [3. 部署 Istio](#span-id-inline-toc-3-span-部署-istio)- [Istio CNI Plugin](#istio-cni-plugin)
        - [Kubernetes 关键插件（Critical Add-On Pods）](#kubernetes-关键插件-critical-add-on-pods)
        - [HostNetwork](#hostnetwork)
        - [只暴露必要端口](#只暴露必要端口)
        
    - [4. 暴露 Dashboard](#span-id-inline-toc-4-span-暴露-dashboard)
    

没错，Istio 架构又换了。。。北京时间 2020 年 3 月 6 日 凌晨发布了 1.5 版本，该版本最大的变化是将控制平面的所有组件组合成一个单体结构叫 `istiod`。

![](https://hugo-picture.oss-cn-beijing.aliyuncs.com/istio-1.5-arch.svg)

从架构图可以看出，在 Istio 1.5 中，饱受诟病的 `Mixer` 终于被废弃了，新版本的 HTTP 遥测默认基于 in-proxy Stats filter，同时可使用 [WebAssembly](https://webassembly.org/) 开发 `in-proxy` 扩展。更详细的说明请参考 [Istio 1.5 发布公告](https://istio.io/news/releases/1.5.x/announcing-1.5/)。

官方文档的部署方法比较笼统，不利于快速上手，为了帮助大家快速上手，本文将重点介绍 Istio 1.5 的部署方法。为了更方便地管理 Istio 各个组件的生命周期，推荐使用 `Operator` 进行部署。

在部署 Istio 之前，首先需要确保 Kubernetes 集群（kubernetes 版本建议在 `1.14` 以上）已部署并配置好本地的 kubectl 客户端。

![](https://hugo-picture.oss-cn-beijing.aliyuncs.com/images/20200306144254.png)

## 1. Kubernetes 环境准备

为了快速准备 kubernetes 环境，我们可以使用 sealos 来部署，步骤如下：

### 前提条件

- 下载[kubernetes 离线安装包](http://store.lameleg.com/)
- 下载[最新版本sealos](https://github.com/fanux/sealos/releases)
- 务必同步服务器时间
- 主机名不可重复

### 安装 kubernetes 集群

```bash
$ sealos init --master 192.168.0.2 \
    --node 192.168.0.3 \
    --node 192.168.0.4 \
    --node 192.168.0.5 \
    --user root \
    --passwd your-server-password \
    --version v1.16.3 \
    --pkg-url /root/kube1.16.3.tar.gz 

```

检查安装是否正常：

```bash
$ kubectl get node

NAME       STATUS   ROLES    AGE   VERSION
sealos01   Ready    master   18h   v1.16.3
sealos02   Ready    <none>   18h   v1.16.3
sealos03   Ready    <none>   18h   v1.16.3
sealos04   Ready    <none>   18h   v1.16.3

```

## 2. 下载 Istio 部署文件

你可以从 GitHub 的 [release](https://github.com/istio/istio/releases/tag/1.5.0) 页面下载 istio，或者直接通过下面的命令下载：

```bash
$ curl -L https://istio.io/downloadIstio | sh -

```

下载完成后会得到一个 `istio-1.5.0` 目录，里面包含了：

- `install/kubernetes` : 针对 Kubernetes 平台的安装文件
- `samples` : 示例应用
- `bin` : istioctl 二进制文件，可以用来手动注入 sidecar proxy

进入 `istio-1.5.0` 目录。

```bash
$ cd istio-1.5.0

$ tree -L 1 ./
./
├── bin
├── install
├── LICENSE
├── manifest.yaml
├── README.md
├── samples
└── tools

4 directories, 4 files

```

将 istioctl 拷贝到 `/usr/local/bin/` 中：

```bash
$ cp bin/istioctl /usr/local/bin/

```

### 开启 istioctl 的自动补全功能

#### bash

将 `tools` 目录中的 `istioctl.bash` 拷贝到 $HOME 目录中：

```bash
$ cp tools/istioctl.bash ~/

```

在 `~/.bashrc` 中添加一行：

```bash
source ~/istioctl.bash

```

应用生效：

```bash
$ source ~/.bashrc

```

#### zsh

将 `tools` 目录中的 `\_istioctl` 拷贝到 $HOME 目录中：

```bash
$ cp tools/_istioctl ~/

```

在 `~/.zshrc` 中添加一行：

```bash
source ~/_istioctl

```

应用生效：

```bash
$ source ~/.zshrc

```

## 3. 部署 Istio

istioctl 提供了多种安装配置文件，可以通过下面的命令查看：

```bash
$ ll install/kubernetes/operator/profiles

-rw-r--r-- 1 root root  18K 3月   4 20:40 default.yaml
-rw-r--r-- 1 root root 3.2K 3月   4 20:40 demo.yaml
-rw-r--r-- 1 root root  964 3月   4 20:40 empty.yaml
-rw-r--r-- 1 root root  913 3月   4 20:40 minimal.yaml
-rw-r--r-- 1 root root  579 3月   4 20:40 remote.yaml
-rw-r--r-- 1 root root  554 3月   4 20:40 separate.yaml

```

它们之间的差异如下：

|  | default | demo | minimal | remote |
| :-: | :-- | :-- | :-- | :-- |
| **核心组件** |  |  |  |  |
| istio-egressgateway |  | **X** |  |  |
| istio-ingressgateway | **X** | **X** |  |  |
| istio-pilot | **X** | **X** | **X** |  |
| **附加组件** |  |  |  |  |
| Grafana |  | **X** |  |  |
| istio-tracing |  | **X** |  |  |
| kiali |  | **X** |  |  |
| prometheus | **X** | **X** |  | **X** |

其中标记 **X** 表示该安装该组件。

如果只是想快速试用并体验完整的功能，可以直接使用配置文件 `demo` 来部署。

在正式部署之前，需要先说明两点：

### Istio CNI Plugin

当前实现将用户 pod 流量转发到 proxy 的默认方式是使用 privileged 权限的 `istio-init` 这个 init container 来做的（运行脚本写入 iptables），需要用到 `NET\_ADMIN` capabilities。对 linux capabilities 不了解的同学可以参考我的 [Linux capabilities 系列](https://fuckcloudnative.io/posts/linux-capabilities-why-they-exist-and-how-they-work/)。

Istio CNI 插件的主要设计目标是消除这个 privileged 权限的 init container，换成利用 Kubernetes CNI 机制来实现相同功能的替代方案。具体的原理就是在 Kubernetes CNI 插件链末尾加上 Istio 的处理逻辑，在创建和销毁 pod 的这些 hook 点来针对 istio 的 pod 做网络配置：写入 iptables，让该 pod 所在的 network namespace 的网络流量转发到 proxy 进程。

详细内容请参考[官方文档](https://istio.io/docs/setup/additional-setup/cni/)。

使用 Istio CNI 插件来创建 sidecar iptables 规则肯定是未来的主流方式，不如我们现在就尝试使用这种方法。

### Kubernetes 关键插件（Critical Add-On Pods）

众所周知，Kubernetes 的核心组件都运行在 master 节点上，然而还有一些附加组件对整个集群来说也很关键，例如 DNS 和 metrics-server，这些被称为**关键插件**。一旦关键插件无法正常工作，整个集群就有可能会无法正常工作，所以 Kubernetes 通过优先级（PriorityClass）来保证关键插件的正常调度和运行。要想让某个应用变成 Kubernetes 的**关键插件**，只需要其 `priorityClassName` 设为 `system-cluster-critical` 或 `system-node-critical`，其中 `system-node-critical` 优先级最高。

> 注意：关键插件只能运行在 `kube-system` namespace 中！

详细内容可以参考[官方文档](https://v1-16.docs.kubernetes.io/docs/tasks/administer-cluster/guaranteed-scheduling-critical-addon-pods/)。

接下来正式安装 Istio，首先部署 `Istio operator`：

```bash
🐳 → istioctl operator init

```

该命令会创建一个 namespace `istio-operator`，并将 Istio operator 部署在此 namespace 中。

```bash
🐳 → kubectl -n istio-operator get pod

NAME                              READY   STATUS    RESTARTS   AGE
istio-operator-7c69599466-bz8lp   1/1     Running   0          3h29m

```

然后创建一个 CR `IstioOperator`：

```yaml
🐳 → kubectl create ns istio-system
🐳 → kubectl apply -f - <<EOF
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
  name: example-istiocontrolplane
spec:
  profile: demo
  components:
    cni:
      enabled: true
      namespace: kube-system
    ingressGateways:
    - enabled: true
      k8s:
        service:
          type: ClusterIP
        strategy:
          rollingUpdate:
            maxUnavailable: 100%
            maxSurge: 0%
        nodeSelector:
          kubernetes.io/hostname: sealos02
        overlays:
        - apiVersion: apps/v1
          kind: Deployment
          name: istio-ingressgateway
          patches:
          - path: spec.template.spec
            value:
              hostNetwork: true
              dnsPolicy: ClusterFirstWithHostNet
        - apiVersion: v1
          kind: Service
          name: istio-ingressgateway
          patches:
          - path: spec.ports
            value:
            - name: status-port
              port: 15020
              targetPort: 15020
            - name: http2
              port: 80
              targetPort: 80
            - name: https
              port: 443
              targetPort: 443
  values:
    cni:
      excludeNamespaces:
       - istio-system
       - kube-system
       - monitoring
      logLevel: info
EOF

```

其中各个字段的详细含义请参考 [`IstioOperator` API 文档](https://istio.io/docs/reference/config/istio.operator.v1alpha1/)，这里我简要说明一下：

- istio-ingressgateway 的 Service 默认类型为 `LoadBalancer`，需将其改为 `ClusterIP`。
- 为防止集群资源紧张，更新配置后无法创建新的 `Pod`，需将滚动更新策略改为先删除旧的，再创建新的。
- 将 istio-ingressgateway 调度到指定节点。
- 默认情况下除了 `istio-system` namespace 之外，istio cni 插件会监视其他所有 namespace 中的 Pod，然而这并不能满足我们的需求，更严谨的做法是让 istio CNI 插件至少忽略 `kube-system`、`istio-system` 这两个 namespace，如果你还有其他的特殊的 namespace，也应该加上，例如 `monitoring`。

下面着重解释 `overlays` 列表中的字段：

### HostNetwork

为了暴露 Ingress Gateway，我们可以使用 `hostport` 暴露端口，并将其调度到某个固定节点。如果你的 CNI 插件不支持 `hostport`，可以使用 `HostNetwork` 模式运行，但你会发现无法启动 ingressgateway 的 Pod，因为如果 Pod 设置了 `HostNetwork=true`，则 dnsPolicy 就会从 `ClusterFirst` 被强制转换成 `Default`。而 Ingress Gateway 启动过程中需要通过 DNS 域名连接 `pilot` 等其他组件，所以无法启动。

我们可以通过强制将 `dnsPolicy` 的值设置为 `ClusterFirstWithHostNet` 来解决这个问题，详情参考：[Kubernetes DNS 高阶指南](https://fuckcloudnative.io/posts/kubernetes-dns/)。

当然你可以部署完成之后再修改 Ingress Gateway 的 `Deployment`，但这种方式还是不太优雅。经过我对 [`IstioOperator` API 文档](https://istio.io/docs/reference/config/istio.operator.v1alpha1/) 的研究，发现了一个更为优雅的方法，那就是直接修改资源对象 `IstioOperator` 的内容，在 `components.ingressGateways` 下面加上么一段：

```yaml
        overlays:
        - apiVersion: apps/v1
          kind: Deployment
          name: istio-ingressgateway
          patches:
          - path: spec.template.spec
            value:
              hostNetwork: true
              dnsPolicy: ClusterFirstWithHostNet

```

具体含义我就不解释了，请看上篇文章。这里只对 IstioOperator 的语法做简单说明：

- `overlays` 列表用来修改对应组件的各个资源对象的 manifest，这里修改的是组件 Ingress Gateway 的 `Deployment`。
- `patches` 列表里是实际要修改或添加的字段，我就不解释了，应该很好理解。

### 只暴露必要端口

从安全的角度来考虑，我们不应该暴露那些不必要的端口，对于 Ingress Gateway 来说，只需要暴露 HTTP、HTTPS 和 metrics 端口就够了。方法和上面一样，直接在 `components.ingressGateways` 的 `overlays` 列表下面加上这么一段：

```yaml
        - apiVersion: v1
          kind: Service
          name: istio-ingressgateway
          patches:
          - path: spec.ports
            value:
            - name: status-port
              port: 15020
              targetPort: 15020
            - name: http2
              port: 80
              targetPort: 80
            - name: https
              port: 443
              targetPort: 443

```

部署完成后，查看各组件状态：

```bash
🐳 → kubectl -n istio-system get pod

NAME                                    READY   STATUS    RESTARTS   AGE
grafana-5cc7f86765-rt549                1/1     Running   0          3h11m
istio-egressgateway-57999c5b76-59z8v    1/1     Running   0          3h11m
istio-ingressgateway-5b97647565-zjz4k   1/1     Running   0          71m
istio-tracing-8584b4d7f9-2jbjp          1/1     Running   0          3h11m
istiod-86798869b8-jmk9v                 1/1     Running   0          3h11m
kiali-76f556db6d-qnsfj                  1/1     Running   0          3h11m
prometheus-6fd77b7876-c4vzn             2/2     Running   0          3h11m

```

```bash
🐳 → kubectl -n kube-system get pod -l k8s-app=istio-cni-node

NAME                   READY   STATUS    RESTARTS   AGE
istio-cni-node-4dlfb   2/2     Running   0          3h12m
istio-cni-node-4s9s7   2/2     Running   0          3h12m
istio-cni-node-8g22x   2/2     Running   0          3h12m
istio-cni-node-x2drr   2/2     Running   0          3h12m

```

可以看到 cni 插件已经安装成功，查看配置是否已经追加到 CNI 插件链的末尾：

```bash
🐳 → cat /etc/cni/net.d/10-calico.conflist

{
  "name": "k8s-pod-network",
  "cniVersion": "0.3.1",
  "plugins": [
  ...
    {
      "cniVersion": "0.3.1",
      "name": "istio-cni",
      "type": "istio-cni",
      "log_level": "info",
      "kubernetes": {
        "kubeconfig": "/etc/cni/net.d/ZZZ-istio-cni-kubeconfig",
        "cni_bin_dir": "/opt/cni/bin",
        "exclude_namespaces": [
          "istio-system",
          "kube-system",
          "monitoring"
        ]
      }
    }
  ]
}

```

## 4. 暴露 Dashboard

这个没什么好说的，通过 Ingress Controller 暴露就好了，可以参考我以前写的 [Istio 1.0 部署](https://fuckcloudnative.io/posts/istio-1.0-deploy/)。如果使用 Contour 的可以参考我的另一篇文章：[Contour 学习笔记（一）：使用 Contour 接管 Kubernetes 的南北流量](https://fuckcloudnative.io/posts/use-envoy-as-a-kubernetes-ingress/)。

这里我再介绍一种新的方式，`istioctl` 提供了一个子命令来从本地打开各种 Dashboard：

```bash
🐳 → istioctl dashboard --help

Access to Istio web UIs

Usage:
  istioctl dashboard [flags]
  istioctl dashboard [command]

Aliases:
  dashboard, dash, d

Available Commands:
  controlz    Open ControlZ web UI
  envoy       Open Envoy admin web UI
  grafana     Open Grafana web UI
  jaeger      Open Jaeger web UI
  kiali       Open Kiali web UI
  prometheus  Open Prometheus web UI
  zipkin      Open Zipkin web UI

```

例如，要想在本地打开 Grafana 页面，只需执行下面的命令：

```bash
🐳 → istioctl dashboard grafana
http://localhost:36813

```

咋一看可能觉得这个功能很鸡肋，我的集群又不是部署在本地，而且这个命令又不能指定监听的 IP，在本地用浏览器根本打不开呀！其实不然，你可以在本地安装 kubectl 和 istioctl 二进制文件，然后通过 kubeconfig 连接到集群，最后再在本地执行上面的命令，就可以打开页面啦，开发人员用来测试是不是很方便？Windows 用户当我没说。。。

接下来我们就可以在浏览器中通过 Gateway 的 URL 来访问服务网格中的服务了。后面我会开启一系列实验教程，本文的所有步骤都是为后面做准备，如果想跟着我做后面的实验，请务必做好本文所述的准备工作。

#### 相关推荐

- [手把手教你部署 Istio 1.3](/posts/istio-1.3-tour/)
- [Istio 1.3 发布，HTTP 遥测不再需要 Mixer](/posts/istio-1.3/)
- [熔断与异常检测在 Istio 中的应用](/posts/circuit_breaking-and-outlier-detection-in-istio/)
- [数据包在 Istio 网格中的生命周期](/posts/life-of-a-packet-through-istio/)
- [Istio 中 VirtualService 的注意事项](/posts/conflictingvirtualservicehost/)


WebAssembly: https://webassembly.org/

Istio 1.5 发布公告: https://istio.io/news/releases/1.5.x/announcing-1.5/

kubernetes 离线安装包: http://store.lameleg.com/

最新版本sealos: https://github.com/fanux/sealos/releases

release: https://github.com/istio/istio/releases/tag/1.5.0

官方文档: https://istio.io/docs/setup/additional-setup/cni/

官方文档: https://v1-16.docs.kubernetes.io/docs/tasks/administer-cluster/guaranteed-scheduling-critical-addon-pods/

IstioOperator API 文档: https://istio.io/docs/reference/config/istio.operator.v1alpha1/

Kubernetes DNS 高阶指南: https://fuckcloudnative.io/posts/kubernetes-dns/


[原文](https://fuckcloudnative.io/posts/istio-deploy/)