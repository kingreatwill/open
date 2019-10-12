
docker部署etcd3.x：[参考](https://skyao.gitbooks.io/learning-etcd3/content/documentation/op-guide/container.html)

```

https://github.com/etcd-io/etcd/releases/tag/v3.4.2

docker pull quay.io/coreos/etcd:v3.4.2 # docker pull gcr.io/etcd-development/etcd:v3.4.2

 docker run -d -p 2379:2379 -p 2380:2380  --name etcd3  --restart always quay.io/coreos/etcd:v3.4.2  /usr/local/bin/etcd  --name s1  --data-dir /etcd-data  --listen-client-urls http://0.0.0.0:2379  --advertise-client-urls http://0.0.0.0:2379  --listen-peer-urls http://0.0.0.0:2380 --initial-advertise-peer-urls http://0.0.0.0:2380  --initial-cluster s1=http://0.0.0.0:2380  --initial-cluster-token tkn  --initial-cluster-state new


  -v /d/dockerv/etcd/data/:/etcd-data/ -v 目前有bug，还没有找到解决方案

docker exec etcd3 /bin/sh -c "/usr/local/bin/etcd --version"
docker exec etcd3 /bin/sh -c "/usr/local/bin/etcdctl version"
docker exec etcd3 /bin/sh -c "/usr/local/bin/etcdctl endpoint health"
docker exec etcd3 /bin/sh -c "/usr/local/bin/etcdctl put foo bar"
docker exec etcd3 /bin/sh -c "/usr/local/bin/etcdctl get foo"

```
# ETCD参数说明
1. —data-dir 指定节点的数据存储目录，这些数据包括节点ID，集群ID，集群初始化配置，Snapshot文件，若未指定—wal-dir，还会存储WAL文件；
2. —wal-dir 指定节点的was文件的存储目录，若指定了该参数，wal文件会和其他数据文件分开存储。
3. —name 节点名称
4. —initial-advertise-peer-urls 告知集群其他节点url.
5. — listen-peer-urls 监听URL，用于与其他节点通讯
6. — advertise-client-urls 告知客户端url, 也就是服务的url
7. — initial-cluster-token 集群的ID
8. — initial-cluster 集群中所有节点

集群版本：