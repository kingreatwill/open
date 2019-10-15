https://rootsongjc.gitbooks.io/kubernetes-handbook

Kubernetes主要由以下几个核心组件组成：
* etcd保存了整个集群的状态；
* apiserver提供了资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制；
* controller manager负责维护集群的状态，比如故障检测、自动扩展、滚动更新等；
* scheduler负责资源的调度，按照预定的调度策略将Pod调度到相应的机器上；
* kubelet负责维护容器的生命周期，同时也负责Volume（CSI）和网络（CNI）的管理；
* Container runtime负责镜像管理以及Pod和容器的真正运行（CRI）；
* kube-proxy负责为Service提供cluster内部的服务发现和负载均衡；

除了核心组件，还有一些推荐的插件，其中有的已经成为CNCF中的托管项目：

* CoreDNS负责为整个集群提供DNS服务
* Ingress Controller为服务提供外网入口
* Prometheus提供资源监控
* Dashboard提供GUI
* Federation提供跨可用区的集群

整体架构
![kubernetes](../img/k8s_1.png)

抽象视图
![kubernetes](../img/k8s_2.png)

Master架构
![kubernetes](../img/k8s_3.png)

Node架构
![kubernetes](../img/k8s_4.png)

分层架构
Kubernetes设计理念和功能其实就是一个类似Linux的分层架构。
![kubernetes](../img/k8s_5.png)
![kubernetes](../img/k8s_6.jpg)
* 核心层：Kubernetes最核心的功能，对外提供API构建高层的应用，对内提供插件式应用执行环境
* 应用层：部署（无状态应用、有状态应用、批处理任务、集群应用等）和路由（服务发现、DNS解析等）、Service Mesh（部分位于应用层）
* 管理层：系统度量（如基础设施、容器和网络的度量），自动化（如自动扩展、动态Provision等）以及策略管理（RBAC、Quota、PSP、NetworkPolicy等）、Service Mesh（部分位于管理层）
* 接口层：kubectl命令行工具、客户端SDK以及集群联邦
* 生态系统：在接口层之上的庞大容器集群管理调度的生态系统，可以划分为两个范畴
    * Kubernetes外部：日志、监控、配置管理、CI/CD、Workflow、FaaS、OTS应用、ChatOps、GitOps、SecOps等
    * Kubernetes内部：CRI、CNI、CSI、镜像仓库、Cloud Provider、集群自身的配置和管理等