https://etcd.io/docs
[TOC]
# ETCD
用于分布式系统最关键数据的分布式可靠键值存储

## 什么是etcd
### Project
etcd是一个高度一致的分布式键值存储，它提供了一种可靠的方式来存储需要由分布式系统或机器集群访问的数据。 它可以优雅地处理网络分区期间的领导者选举，即使在领导者节点中也可以容忍机器故障。

从简单的Web应用程序到[Kubernetes](https://kubernetes.io/)，任何复杂的应用程序都可以从etcd中读取数据或将数据写入etcd。

您的应用程序可以读取和写入etcd数据。 一个简单的用例是将数据库连接详细信息或功能标志存储在etcd中作为键值对。 可以观察这些值，使您的应用程序在更改时可以自行重新配置。 高级用途利用etcd的一致性保证来实施数据库领导者选举或在一组工作人员之间执行分布式锁定。

etcd是开源的，可在[GitHub](https://github.com/etcd-io/etcd)上获得，并由[Cloud Native Computing Foundation](https://cncf.io/)支持。

### 技术概述
etcd是用[Go](https://golang.org/)编写的，它具有出色的跨平台支持，小的二进制文件和强大的社区。 etcd机器之间的通信通过**Raft共识算法**处理。

etcd领导者的延迟是要跟踪的最重要的指标，并且内置仪表板具有专用于此的视图。 在我们的测试中，严重的延迟会在群集内造成不稳定，因为Raft的速度仅与大多数机器中最慢的机器一样快。 您可以通过适当地调整群集来缓解此问题。 etcd已在具有高度可变网络的云提供商上进行了预调。

### 哪些人在用

- k8s: etcd是服务发现的后端，并存储集群状态和配置

- Rook: etcd充当Rook的编排引擎

- CoreDNS : CoreDNS使用etcd作为可选后端

- m3 :M3是Uber为Prometheus建立的大规模指标平台，使用etcd进行规则存储和其他功能

- OpenStack: OpenStack支持etcd作为配置存储，分布式密钥锁定等的可选提供程序

此外，GitHub上的数千个项目都与[etcd](https://github.com/search?utf8=%E2%9C%93&q=etcd/)关联，包括基于etcd构建的项目，客户端绑定等等。

## Features

### Simple interface
使用标准的HTTP工具（例如curl）读取和写入值

### 键值存储
将数据存储在分层组织的目录中，例如在标准文件系统中
```
/config
    /database
/xxx
```

### Watch for changes
观察特定的键或目录的更改，并对值的更改做出反应


- 安全：可选的SSL客户端证书认证
- 快速：基准为每个实例1000次写入/秒
- 可选的TTL用于密钥过期
- 可靠：通过Raft协议正确分发
