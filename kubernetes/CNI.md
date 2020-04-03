# CNI 简单介绍

CNI（Container Network Interface）是 CNCF 旗下的一个项目，由一组用于配置Linux容器网络接口的规范和库组成，同时还包含了一些插件。CNI 仅关心容器创建时的网络分配和当容器被删除时释放网络资源。Kubernetes 中已经内置了 CNI。

Container Runtime 在创建容器时，先创建好 network namespace，然后调用CNI插件为这个 network namespace 配置网络，其后再启动容器内的进程。CNI 已成为 CNCF 一员，成为 CNCF 主推的网络模型。

CNI 项目地址 https://github.com/containernetworking/cni

Kubernetes Pod 中的其他容器都是Pod所属 Pause 容器的网络，创建过程为：

1. kubelet 先创建pause容器生成network namespace
2. 调用网络CNI driver
3. CNI driver 根据配置调用具体的cni 插件
4. cni 插件给pause 容器配置网络
5. pod 中其他的容器都使用 pause 容器的网络

Flannel 
Calico

