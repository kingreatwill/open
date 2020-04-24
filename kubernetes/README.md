![](../img/k8s/pod-info.jpeg)

[Rancher](https://profile.zjurl.cn/rogue/ugc/profile/?version_code=769&version_name=70609&user_id=52042300867)

https://kuboard.cn/learning

https://learnk8s.io/troubleshooting-deployments
[k8s故障排除指南.pdf](../files/k8s/troubleshooting-kubernetes.pdf)
![](../img/k8s/k8s故障排除指南.png)

Kubernetes 中，0.5 代表请求半个 CPU 资源。表达式 0.1 等价于 表达式 100m （英文读作 one hundred millicpu 或 one hundred millicores）。在 API Server 中，表达式 0.1 将被转换成 100m，精度低于 1m 的请求是不受支持的。 CPU 的计量代表的是绝对值，而非相对值，例如，您请求了 0.1 个 CPU，无论您的节点是 单核、双核、48核，您得到的 CPU 资源都是 0.1 核。

Mi表示（1Mi=1024x1024）,M表示（1M=1000x1000）


在 top 命令查看CPU消耗，100% 代表 1核；4核为 400%；10% 代表 0.1 核 或者 100m

污点和容忍
亲和性和反亲和性



中文文档：https://kubernetes.io/zh/docs/

Kubernetes Handbook （Kubernetes指南）
https://github.com/feiskyer/kubernetes-handbook

和我一步步部署 kubernetes 集群
https://github.com/opsnull/follow-me-install-kubernetes-cluster

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

[图解kubernetes调度器SchedulingQueue核心源码实现](https://m.toutiao.com/i6781307442589336067/)

[Kubernetes REST API](https://www.jianshu.com/p/0de6bc64c423)
