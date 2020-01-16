<!-- toc -->
[TOC]

[中文文档](http://etcd.doczh.cn/)

[官网文档](https://etcd.io/docs)

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


## ETCD在k8s上的应用
ETCD 存储 k8s 所有数据信息

ETCD 是k8s集群极为重要的一块服务，存储了集群所有的数据信息。同理，如果发生灾难或者 etcd 的数据丢失，都会影响集群数据的恢复。

### ETCD 一些查询操作

- 查看集群状态
```
ETCDCTL_API=3 etcdctl --cacert=/opt/kubernetes/ssl/ca.pem --cert=/opt/kubernetes/ssl/server.pem --key=/opt/kubernetes/ssl/server-key.pem --endpoints=https://192.168.1.36:2379,https://192.168.1.37:2379,https://192.168.1.38:2379 endpoint health

https://192.168.1.36:2379 is healthy: successfully committed proposal: took = 1.698385ms
https://192.168.1.37:2379 is healthy: successfully committed proposal: took = 1.577913ms
https://192.168.1.38:2379 is healthy: successfully committed proposal: took = 5.616079ms
```
- 获取某个 key 信息
```
ETCDCTL_API=3 etcdctl --cacert=/opt/kubernetes/ssl/ca.pem --cert=/opt/kubernetes/ssl/server.pem --key=/opt/kubernetes/ssl/server-key.pem --endpoints=https://192.168.1.36:2379,https://192.168.1.37:2379,https://192.168.1.38:2379 get /registry/apiregistration.k8s.io/apiservices/v1.apps
```

- 获取 etcd 版本信息
```
ETCDCTL_API=3 etcdctl --cacert=/opt/kubernetes/ssl/ca.pem --cert=/opt/kubernetes/ssl/server.pem --key=/opt/kubernetes/ssl/server-key.pem --endpoints=https://192.168.1.36:2379,https://192.168.1.37:2379,https://192.168.1.38:2379 version
```
-获取 ETCD 所有的 key
```
ETCDCTL_API=3 etcdctl --cacert=/opt/kubernetes/ssl/ca.pem --cert=/opt/kubernetes/ssl/server.pem --key=/opt/kubernetes/ssl/server-key.pem --endpoints=https://192.168.1.36:2379,https://192.168.1.37:2379,https://192.168.1.38:2379 get / --prefix --keys-only
```

### ETCD 备份
> ETCD 不同的版本的 etcdctl 命令不一样，但大致差不多，本文备份使用 napshot save , 每次备份一个节点就行。

- 命令备份（k8s-master1 机器上备份）：
```
ETCDCTL_API=3 etcdctl --cacert=/opt/kubernetes/ssl/ca.pem --cert=/opt/kubernetes/ssl/server.pem --key=/opt/kubernetes/ssl/server-key.pem --endpoints=https://192.168.1.36:2379 snapshot save /data/etcd_backup_dir/etcd-snapshot-`date +%Y%m%d`.db
```
- 备份脚本（k8s-master1 机器上备份）：
```sh
#!/usr/bin/env bash

date;

CACERT="/opt/kubernetes/ssl/ca.pem"
CERT="/opt/kubernetes/ssl/server.pem"
EKY="/opt/kubernetes/ssl/server-key.pem"
ENDPOINTS="192.168.1.36:2379"

ETCDCTL_API=3 etcdctl \
--cacert="${CACERT}" --cert="${CERT}" --key="${EKY}" \
--endpoints=${ENDPOINTS} \
snapshot save /data/etcd_backup_dir/etcd-snapshot-`date +%Y%m%d`.db

# 备份保留30天
find /data/etcd_backup_dir/ -name *.db -mtime +30 -exec rm -f {} \;
```

### ETCD 恢复
#### 准备工作
- 停止所有 Master 上 kube-apiserver 服务
```
$ systemctl stop kube-apiserver

# 确认 kube-apiserver 服务是否停止
$ ps -ef | grep kube-apiserver
```
- 停止集群中所有 ETCD 服务
```
$ systemctl stop etcd
```
- 移除所有 ETCD 存储目录下数据
```
$ mv /var/lib/etcd/default.etcd /var/lib/etcd/default.etcd.bak
```
- 拷贝 ETCD 备份快照
```
# 从 k8s-master1 机器上拷贝备份
$ scp /data/etcd_backup_dir/etcd-snapshot-20191222.db root@k8s-master2:/data/etcd_backup_dir/
$ scp /data/etcd_backup_dir/etcd-snapshot-20191222.db root@k8s-master3:/data/etcd_backup_dir/
```

#### 恢复备份
```sh
# k8s-master1 机器上操作
$ ETCDCTL_API=3 etcdctl snapshot restore /data/etcd_backup_dir/etcd-snapshot-20191222.db \
  --name etcd-0 \
  --initial-cluster "etcd-0=https://192.168.1.36:2380,etcd-1=https://192.168.1.37:2380,etcd-2=https://192.168.1.38:2380" \
  --initial-cluster-token etcd-cluster \
  --initial-advertise-peer-urls https://192.168.1.36:2380 \
  --data-dir=/var/lib/etcd/default.etcd
  
# k8s-master2 机器上操作
$ ETCDCTL_API=3 etcdctl snapshot restore /data/etcd_backup_dir/etcd-snapshot-20191222.db \
  --name etcd-1 \
  --initial-cluster "etcd-0=https://192.168.1.36:2380,etcd-1=https://192.168.1.37:2380,etcd-2=https://192.168.1.38:2380"  \
  --initial-cluster-token etcd-cluster \
  --initial-advertise-peer-urls https://192.168.1.37:2380 \
  --data-dir=/var/lib/etcd/default.etcd
  
# k8s-master3 机器上操作
$ ETCDCTL_API=3 etcdctl snapshot restore /data/etcd_backup_dir/etcd-snapshot-20191222.db \
  --name etcd-2 \
  --initial-cluster "etcd-0=https://192.168.1.36:2380,etcd-1=https://192.168.1.37:2380,etcd-2=https://192.168.1.38:2380"  \
  --initial-cluster-token etcd-cluster \
  --initial-advertise-peer-urls https://192.168.1.38:2380 \
  --data-dir=/var/lib/etcd/default.etcd
```
- 上面三台 ETCD 都恢复完成后，依次登陆三台机器启动 ETCD
```
systemctl start etcd
```
- 三台 ETCD 启动完成，检查 ETCD 集群状态
```
ETCDCTL_API=3 etcdctl --cacert=/opt/kubernetes/ssl/ca.pem --cert=/opt/kubernetes/ssl/server.pem --key=/opt/kubernetes/ssl/server-key.pem --endpoints=https://192.168.1.36:2379,https://192.168.1.37:2379,https://192.168.1.38:2379 endpoint health
```
- 三台 ETCD 全部健康，分别到每台 Master 启动 kube-apiserver
```
systemctl start kube-apiserver
```
- 检查 Kubernetes 集群是否恢复正常
```
$ kubectl get cs
```
### 总结
Kubernetes 集群备份主要是备份 ETCD 集群。而恢复时，主要考虑恢复整个顺序：

停止kube-apiserver --> 停止ETCD --> 恢复数据 --> 启动ETCD --> 启动kube-apiserve

> 注意：备份ETCD集群时，只需要备份一个ETCD就行，恢复时，拿同一份备份数据恢复。

[锻骨境-第6层 k8s集群数据备份与恢复](https://www.jianshu.com/p/8b483ed49f26)
[k8s安装之etcd备份还原yaml](https://www.cnblogs.com/aguncn/p/10905564.html)

[为 Kubernetes 运行 etcd 集群](https://kubernetes.io/zh/docs/tasks/administer-cluster/configure-upgrade-etcd/)