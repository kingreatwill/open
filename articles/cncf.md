![CNCF全景图](../img/cncf/logo.png)
# 基本概念
- CNCF是什么
- 什么是云原生（Cloud Native）
- TOC（Technical Oversight Committee）
- CNCF路线图（Trail Map）
https://liumiaocn.blog.csdn.net/article/details/100653635

## CNCF是什么
Cloud Native Computing Foundation（[CNCF](https://github.com/cncf/landscape)）云原生计算基金会隶属于 [Linux  Foundation](https://liumiaocn.blog.csdn.net/article/details/100666545)，是一个非营利性组织，它致力于培育和维护一个厂商中立的开源生态系统，来推广云原生技术。CNCF 成立于 2015 年 12 月，其通过将最前沿的模式民主化，让创新为大众所用。目前 CNCF 已经接纳了众多项目，其中明星项目不少，广为人知的有 Kubernetes、Prometheus 和 gRPC、Linkerd 等。

## 什么是云原生（Cloud Native）
CNCF是云原生计算基金会，而关于什么是云原生，CNCF给出了如下的定义：

    云原生技术有利于各组织在公有云、私有云和混合云等新型动态环境中，构建和运行可弹性扩展的应用。云原生的代表技术包括容器、服务网格、微服务、不可变基础设施和声明式API。
    这些技术能够构建容错性好、易于管理和便于观察的松耦合系统。结合可靠的自动化手段，云原生技术使工程师能够轻松地对系统作出频繁和可预测的重大变更。
    云原生计算基金会（CNCF）致力于培育和维护一个厂商中立的开源生态系统，来推广云原生技术。我们通过将最前沿的模式民主化，让这些创新为大众所用。
可以看到云原生的技术出现的背景是组织在推行云计算时，由于公有云、私有云和混合云的复杂环境已经变得越来越普及，所以在这种基础架构之下如何保证松耦合的系统能够更好地被监控和管理，是云原生技术需要面对和解决的问题。

## TOC（Technical Oversight Committee）
TOC（Technical Oversight Committee）的主要成员有9位，包括阿里的Li-Xiang，任期两年至2021年


# CNCF全景图（Landscape）
![CNCF全景图](../img/cncf/landscape.png)

[毕业标准](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc)
每个 CNCF 项目都有成熟度。拟议的 CNCF 项目应说明他们的首选成熟度。一个孵化或毕业项目需要三分之二的绝对多数选票。如果没有绝对多数选票作为毕业项目，那么毕业选票会被计算为孵化选票。如果没有绝对多数选票作为孵化项目，则任何毕业或孵化选票都会被计算为赞助沙箱。如果没有足够的赞助作为沙箱阶段项目，该项目将被拒绝。此投票过程称为后备投票。

不同成熟度的项目都可以获得 https://cncf.io/projects 列出的所有资源，但如果存在资源不足，更成熟的项目通常具有优先权。

### 沙箱阶段

进入沙箱阶段，项目必须至少有 2 位 TOC 赞助。有关详细过程，请参阅 CNCF 沙箱指南 1.0 版本。

### 孵化阶段

进入孵化阶段，项目必须满足沙箱阶段要求以及：

- 记录至少有三个独立的最终用户在生产中成功使用了项目，经过 TOC 的判断，认为有足够的质量和范围。

- 拥有健康数量的提交者。提交者的定义是具有提交代码的人；即是可以针对项目部分或全部接受贡献的人。

- 展示大量和持续的提交和合并的贡献。

- 由于这些度量可能根据项目的类型、范围和规模而有很大差异，因此 TOC 对满足这些度量的活动水平有最终判决。

### 毕业阶段

从沙箱或孵化状态毕业，或者作为一个新项目加入作为一个毕业项目，项目必须符合孵化阶段标准以及：

- 有来自至少两个机构的提交者。

- 已经实现并维护了核心基础结构计划（CII）的最佳实践徽章。

- 采用 CNCF 行为准则。

- 明确定义项目治理和提交者流程。这最好在 GOVERNANCE.md 文件中进行，并引用 OWNERS.md 文件显示当前和荣誉提交者。

- 至少在主要仓库提供项目采用者的公开列表（例如，ADOPTERS.md 文件或项目网站上的徽标）。

- 获得 TOC 的绝对多数选票进入毕业阶段。如果项目能够表现出足够的成熟度，项目可以尝试直接从沙箱移动到毕业。项目可以无限期地保持在孵化状态，但通常预计在两年内毕业。

## 毕业列表
名称 | 时间 | 地址 | 作用
--|--|--|--
Kubernetes|[2018.03.06](https://www.cncf.io/announcement/2018/03/06/cloud-native-computing-foundation-announces-kubernetes-first-graduated-project/)|[golang](https://github.com/kubernetes/kubernetes)|编排调度
Prometheus|[2018.08.09](https://www.cncf.io/announcement/2018/08/09/prometheus-graduates/)|[golang](https://github.com/prometheus/prometheus)|监控
Envoy|[2018.11.28](https://www.cncf.io/announcement/2018/11/28/cncf-announces-envoy-graduation/)|[C++](https://github.com/envoyproxy/envoy)|服务代理
CoreDNS|[2019.01.24](https://www.cncf.io/announcement/2019/01/24/coredns-graduation/)|[golang](https://github.com/coredns/coredns)|协同与服务发现
containerd|[2019.02.28](https://www.cncf.io/announcement/2019/02/28/cncf-announces-containerd-graduation/)|[golang](https://github.com/containerd/containerd)|容器运行环境
Fluentd|[2019.04.11](https://www.cncf.io/announcement/2019/04/11/cncf-announces-fluentd-graduation/)|[Ruby](https://github.com/fluent/fluentd)|日志
Jaeger|[2019.10.31](https://www.cncf.io/announcement/2019/10/31/cloud-native-computing-foundation-announces-jaeger-graduation/)|[golang](https://github.com/jaegertracing/jaeger)|调用链追踪
Vitess|[2019.11.05](https://www.cncf.io/announcement/2019/11/05/cloud-native-computing-foundation-announces-vitess-graduation/)|[golang](https://github.com/vitessio/vitess)|数据库
TUF|[2019.12.18](https://www.cncf.io/announcement/2019/12/18/cloud-native-computing-foundation-announces-tuf-graduation/)|[python](https://github.com/theupdateframework/tuf)|规范与安全性-保护软件更新系统
Helm|[2020.04.30](https://www.cncf.io/announcement/2020/04/30/cloud-native-computing-foundation-announces-helm-graduation/)|[golang](https://github.com/helm/helm)|k8s的包管理器
Harbor |[2020.06.23](https://www.cncf.io/announcement/2020/06/23/cloud-native-computing-foundation-announces-harbor-graduation/) |[golang](https://github.com/goharbor/harbor) |开源制品（artifact）仓库，可通过策略和基于角色的访问控制来保护制品（如容器镜像、Helm Chart等）
TiKV|[2020.09.02](https://www.cncf.io/announcements/2020/09/02/cloud-native-computing-foundation-announces-tikv-graduation/)|[rust](https://github.com/tikv/tikv)|TiKV 是一个开源的分布式事务 Key-Value 数据库，专注为下一代数据库提供可靠、高质量、实用的存储架构。TiKV 是继 Harbor 之后在 CNCF 毕业的第二个中国原创开源项目
Rook|[2020.10.07](https://www.cncf.io/announcements/2020/10/07/cloud-native-computing-foundation-announces-rook-graduation/)|[golang](https://github.com/rook/rook)|Rook 是面向 Kubernetes 的开源云原生存储编排器
Etcd | [2020.11.24](https://www.cncf.io/announcements/2020/11/24/cloud-native-computing-foundation-announces-etcd-graduation/)|[golang](https://github.com/etcd-io/etcd)| 分布式K-V存储系统
Open Policy Agent(OPA)|[2021.02.04](https://www.cncf.io/announcements/2021/02/04/cloud-native-computing-foundation-announces-open-policy-agent-graduation/)|[golang](https://github.com/open-policy-agent/opa)|[通用策略引擎](https://zhuanlan.zhihu.com/p/353531134)
Linkerd |[2021.07.28](https://www.cncf.io/announcements/2021/07/28/cloud-native-computing-foundation-announces-linkerd-graduation/)|[golang](https://github.com/linkerd/linkerd2)|服务网格
SPIFFE|[2022.09.20](https://www.cncf.io/announcements/2022/09/20/spiffe-and-spire-projects-graduate-from-cloud-native-computing-foundation-incubator/)  |[golang](https://github.com/spiffe/spiffe)| 安全生产身份框架
SPIRE | [2022.09.20](https://www.cncf.io/announcements/2022/09/20/spiffe-and-spire-projects-graduate-from-cloud-native-computing-foundation-incubator/) |[golang](https://github.com/spiffe/spire)|The SPIFFE Runtime Environment
Flux|[2022.11.30](https://www.cncf.io/announcements/2022/11/30/flux-graduates-from-cncf-incubator/)|[golang](https://github.com/fluxcd/flux2)|持续交付工具
Argo |[2022.12.06](https://www.cncf.io/announcements/2022/12/06/the-cloud-native-computing-foundation-announces-argo-has-graduated/)| [golang](https://github.com/argoproj) | Kubernetes-native tools to run workflows, manage clusters, and do GitOps right.

## Landscape介绍
https://blog.csdn.net/liumiaocn/article/details/100713072

