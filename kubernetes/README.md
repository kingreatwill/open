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



### 文档
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

## 其它
### Kubernetes日志监控工具
#### Zebrium
优点：易于启动；只需复制/粘贴自定义的HELM或kubectl命令；自动检测问题和根本原因，无需手动规则；可以用作独立的日志管理工具，也可以用作现有日志管理工具（例如ELK Stack）的机器学习附件。
缺点：免费计划限制为每天500 MB，保留3天；支持Kubernetes，Docker和大多数常见平台，但不支持Windows

#### Sematext
用于日志管理和应用程序性能监控的解决方案。Sematex提供了系统状态的全栈可见性。
Sematext不仅限于Kubernetes日志，还可以监控和Kubernetes（基于度量标准和日志）。收集到的日志会自动针对几种不同的已知日志格式进行解析/结构化，并且用户还可以提供自定义日志的模式。它还公开了Elasticsearch API，因此也可以使用任何与Elasticsearch配合使用的工具，例如Filebeat和Logstash与Sematex。可以将其用作ELK的变体或与本机Sematext生态系统一起使用。该工具有助于创建特定规则，来监控特定情况并捕获异常。借助Sematex全面的实时仪表板，客户可以控制和监控所有服务。

优点：与其他Sematext云工具集成；可配置超限来阻止日志被接受从而控制成本；具有ELK的灵活性。

缺点：Sematext小部件和Kibana不能在一个仪表板上混合使用；自定义解析需要在日志传送器中完成，Sematext仅在服务器端解析Syslog和JSON；跟踪功能较弱，但已经在计划进行改进。

#### Loki
Loki是一个受Prometheus启发的多租户和高度可用的日志聚合工具。这款工具有助于收集日志，但是用户将需要为其建立手动规则。Loki与Grafana，Prometheus和Kubernetes合作。Loki可以让内部流程更有效率。如，它节省了Paytm Insider 75%的日志记录和监控成本。Loki不会索引你的日志内容，而是仅索引每个事件流的一组标签，因此效率很高。

优点：拥有大型的生态系统；丰富的可视化功能；由于未索引日志内容而提高了效率。并且可以配合tempo调用链查看日志.

缺点：未针对Kubernetes日志管理进行优化；大量的架构规则手工工作；缺少内容索引可能会限制搜索性能。

#### ELK Stack

ELK是最著名的日志管理开源工具。ELK是Elasticsearch，Logstash和Kibana的首字母缩写。每个组件负责日志记录过程的不同部分。Elasticsearch是一个功能强大且可扩展的搜索系统，Logstash聚合并处理日志，而Kibana提供了一个分析和可视化界面，可帮助用户理解数据。它们共同为Kubernetes提供了全面的日志记录解决方案。但ELK Stack还有许多其他变体，如EFK Stack，即Elasticsearch，Fluentd和Kibana组成。

优点：ELK是众所周知的，并且拥有庞大的社区；非常广泛的平台支持；Kibana中丰富的分析和可视化功能；需要对日志和手动定义的警报规则进行复杂的分析。

缺点：维持规模难度大；需要很多调整，特别是对于大型环境；大量的资源需求；某些功能需要付费许可证。

#### Fluentd

优点：大型社区和插件生态系统；统一日志记录层；经过验证的可靠性和性能。可以在不到10分钟的时间内安装完毕。

缺点：难以配置；对转换数据的支持有限；不是完整的日志记录解决方案。