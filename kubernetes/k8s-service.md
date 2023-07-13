https://kuboard.cn/learning/k8s-intermediate/service/service-details.html

[Kubernetes学习之路（十四）之服务发现Service](https://www.cnblogs.com/linuxk/p/9605901.html)
[kubernetes系列08—service资源详解](https://www.cnblogs.com/along21/p/10330076.html)
[浅谈 Kubernetes Service 负载均衡实现机制](https://www.infoq.cn/article/P0V9D4br7UDzWTgIHuYq)

## 创建服务
有时不想写yaml文件, 可以基于deployment直接创建一个服务来暴露 Deployment
`kubectl expose deployment hello-world --type=NodePort --name=example-service`
example-service是服务名称

## session affinity
session affinity， 会话亲和性，又称会话保持

session保持
如何在service内部实现session保持呢？当然是在service的yaml里进行设置啦。

在service的yaml的sepc里加入以下代码：
```
sessionAffinity: ClientIP
sessionAffinityConfig:
    clientIP:
      timeoutSeconds: 10800
```
这样就开启了session保持。下面的timeoutSeconds指的是session保持的时间，这个时间默认是10800秒，也就是三个小时。

那么原理是啥呢？当不设置session保持时，service向后台pod转发规则是轮询。当设置了session保持之后，k8s会根据访问的ip来把请求转发给他以前访问过的pod，这样session就保持住了。

service.spec.sessionAffinity
默认值为 "None"
如果设定为 "ClientIP"，则同一个客户端的连接将始终被转发到同一个 Pod

service.spec.sessionAffinityConfig.clientIP.timeoutSeconds
默认值为 10800 （3 小时）
设定会话保持的持续时间
#多端口的Service



## service负载均衡
service负载均衡 默认是根据kube-proxy的负载均衡策略

kube-proxy 作为一个控制器，作为 k8s 和 Linux kernel Netfilter 交互的一个枢纽。监听 kubernetes 集群 Services 和 Endpoints 对象的变化，并根据 kube-proxy 不同的模式 (iptables or ipvs), 对内核设置不同的规则，来实现路由转发

静态方法，根据算法本身进行轮询调度
RR, Round Robin
WRR，Wrighted RR
SH，SourceIP Hash
DH，Destination Hash
动态方法，根据算法以及 RS 的当前负载状态进行调度
LC，least connections
WLC，Weighted Least Connection
SED，Shortest Expection Delay
NQ，Never Queue
LBLC，Locality-Based Least Connection
LBLCR，Locality-Based Least Connections withReplication

kube-proxy 中使用 IPVS 模式进行负载均衡。首先需要在启动 kube-proxy 的参数中指定如下参数:

kube-proxy 可以通过 --ipvs-scheduler 参数选择调度算法，默认情况下是 Round Robin 算法。
```
--proxy-mode=ipvs // 将 kube-proxy 的模式设置为 IPVS
--ipvs-scheduler=rr // 设置 ipvs 的负载均衡算法，默认是 rr
--ipvs-min-sync-period=5s // 刷新 IPVS 规则的最小时间间隔
--ipvs-sync-period=30s // 刷新 IPVS 规则的最大时间间隔

```

## Pod 的 hostname / subdomain

https://kuboard.cn/learning/k8s-intermediate/service/dns.html#pod-%E7%9A%84-hostname-subdomain

##  Service 证书
https://kuboard.cn/learning/k8s-intermediate/service/connecting.html#%E4%BF%9D%E6%8A%A4-service-%E7%9A%84%E5%AE%89%E5%85%A8


https://kuboard.cn/learning/k8s-intermediate/config/secrets/create-kubectl.html

## 配置Pod的 /etc/hosts
通过 Pod 定义中的 .spec.hostAliases 字段，我们可以向 Pod 的 /etc/hosts 文件中添加额外的条目，用来解析 foo.local、bar.local 到 127.0.0.1 和 foo.remote、bar.remote 到 10.1.2.3，如下所示：
```
apiVersion: v1
kind: Pod
metadata:
  name: hostaliases-pod
spec:
  restartPolicy: Never
  hostAliases:
  - ip: "127.0.0.1"
    hostnames:
    - "foo.local"
    - "bar.local"
  - ip: "10.1.2.3"
    hostnames:
    - "foo.remote"
    - "bar.remote"
  containers:
  - name: cat-hosts
    image: busybox
    command:
    - cat
    args:
    - "/etc/hosts"
```

由于该文件已经被 Kubelet 管理起来，任何对该文件手工修改的内容，都将在 Kubelet 重启容器或者 Pod 重新调度时被覆盖。因此，最好是通过 hostAliases 修改 Pod 的 /etc/hosts 文件，而不是手工修改。

## Network Policies

比如限制某些pod不能访问指定的pod


https://kuboard.cn/learning/k8s-intermediate/service/np-example.html#%E5%89%8D%E6%8F%90%E6%9D%A1%E4%BB%B6


## k8s中ResourceQuota与LimitRange的作用

ResourceQuota
ResourceQuota 用来限制 namespace 中所有的 Pod 占用的总的资源 request 和 limit

LimitRange
LimitRange 用来限制 namespace 中 单个Pod 默认资源 request 和 limit
https://blog.csdn.net/qq_33235529/article/details/105194130


