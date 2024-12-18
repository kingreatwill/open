Kubernetes将资源指标分为了两种：

* core metrics（核心指标）： 采集每个节点上的kubelet公开的summary api中的指标信息，通常只包含cpu、内存使用率信息
* custom metrics（自定义指标）：允许用户从外部的监控系统当中采集自定义指标，如应用的qps等

在autoscaling/v1版本中只支持CPUUtilizationPercentage一种指标，
在autoscaling/v2beta1中增加支持custom metrics

autoscaling/v1                 #只支持通过cpu为参考依据，来改变pod副本数
autoscaling/v2beta1       #支持通过cpu、内存、连接数以及用户自定义的资源指标数据为参考依据。
autoscaling/v2beta2       #同上，小的变动

查询：
kubectl explain hpa   ##默认查询到的是autoscaling/v1版本
kubectl explain hpa --api-version=autoscaling/v2beta1   ##如果使用其他版本，可以使用--api-version指明版本
1. 确认是否可用
kubectl api-versions

是否有autoscaling/v2beta1

2
git clone https://github.com/stefanprodan/k8s-prom-hpa


https://github.com/kubernetes-incubator/metrics-server

https://www.cnblogs.com/breezey/p/11711077.html
https://segmentfault.com/a/1190000018141551?utm_source=tag-newest

# 安装[jq](https://stedolan.github.io/jq/download/)命令
https://www.jianshu.com/p/6de3cfdbdb0e
```
windows:
chocolatey install jq
centos:
安装EPEL源：
yum install epel-release

安装完EPEL源后，可以查看下jq包是否存在：
yum list jq

安装jq：
yum install -y jq
```

# 1部署 Metrics Server
```
1. 修改image 为 registry.cn-hangzhou.aliyuncs.com/google_containers/metrics-server-amd64:v0.3.6
2. 在metrics-server-deployment.yaml中添加了一个command，加了两个kubelet的配置项，如果不添加此项，metrics-server无法采集数据指标
其它参数：https://github.com/kubernetes-incubator/metrics-server#flags
command:
        - /metrics-server
        - --kubelet-insecure-tls
        - --kubelet-preferred-address-types=InternalIP

或者修改 metric-server deployment 的参数
# args:
# - --kubelet-insecure-tls
# - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
$ kubectl edit deploy -n kube-system metrics-server



部署(k8s 1.16版本将metrics-server-deployment.yaml的api版本改为apps/v1，1.14不用)
kubectl create -f yml/metrics-server/0.3.6/deploy/1.8+/

kubectl get pod -n kube-system

kubectl get --raw "/apis/metrics.k8s.io/v1beta1/nodes" | jq .
kubectl get --raw "/apis/metrics.k8s.io/v1beta1/pods" | jq .

The list of supported endpoints:

/nodes - all node metrics; type []NodeMetrics
/nodes/{node} - metrics for a specified node; type NodeMetrics
/namespaces/{namespace}/pods - all pod metrics within namespace with support for all-namespaces; type []PodMetrics
/namespaces/{namespace}/pods/{pod} - metrics for a specified pod; type PodMetrics

然后就可以 基于CPU和内存使用情况的Auto Scaling
如果CPU平均值超过80％或内存超过200Mi，则最多可扩展到10个：
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: podinfo
spec:
  scaleTargetRef:
    apiVersion: extensions/v1beta1
    kind: Deployment
    name: podinfo
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: 80   ##注意此时是根据使用率，也可以根据使用量：targetAverageValue
  - type: Resource
    resource:
      name: memory
      targetAverageValue: 200Mi

```

# 2部署 Custom Metrics Server
https://github.com/DirectXMan12/k8s-prometheus-adapter
k8s-prometheus-adapter(将Prometheus采集的数据转换为指标格式)
```
您将在专用命名空间中部署Prometheus和适配器。

创建monitoring命名空间：

kubectl create -f ./namespaces.yaml

在monitoring命名空间中部署Prometheus v2：
kubectl create -f ./prometheus

http://192.168.1.120:31190/graph


生成Prometheus适配器所需的TLS证书：

make certs
生成以下几个文件：

# ls output
apiserver.csr  apiserver-key.pem  apiserver.pem
部署Prometheus自定义指标API适配器：

kubectl create -f ./custom-metrics-api
列出Prometheus提供的自定义指标：

kubectl get --raw "/apis/custom.metrics.k8s.io/v1beta1" | jq .

```
部署测试应用
```
kubectl create -f ./podinfo/podinfo-svc.yaml,./podinfo/podinfo-dep.yaml

podinfo-dep.yaml 1.16版本要加上
api版本改成apps/v1
  selector:
    matchLabels:
      app: podinfo

获取自定义指标
kubectl get --raw "/apis/custom.metrics.k8s.io/v1beta1/namespaces/default/pods/*/http_requests" | jq .
java 版本：javapodinfo.yaml

```
部署自定义HPA
```
修改
spec:
  scaleTargetRef:
    apiVersion: extensions/v1beta1
为
spec:
  scaleTargetRef:
    apiVersion: apps/v1

podinfo应用程序公开名为http_requests_total的自定义指标。 Prometheus适配器删除_total后缀并将度量标记为计数器度量标准。

kubectl create -f ./podinfo/podinfo-hpa-custom.yaml
如果请求数超过每秒10个，将扩展podinfo部署


查看HPA
kubectl get hpa

并发测试

在podinfo服务上应用一些负载，每秒25个请求：
#do 10K requests rate limited at 25 QPS
hey -n 10000 -q 5 -c 5 http://192.168.1.120:31198/healthz
下面两个什么意思？
hey -n 10000 -q 25  http://192.168.1.120:31198/healthz
hey -n 10000 -c 25  http://192.168.1.120:31198/healthz
hey参数说明：https://github.com/rakyll/hey
hey命令-c表示并发数 -q 速率限制，以每秒查询(QPS)为单位


ab -c 25 -n 10000 http://192.168.1.120:31198/healthz
ab参数说明https://www.cnblogs.com/blueskycc/p/5509490.html

几分钟后，HPA开始扩展部署：
kubectl describe hpa
能看到什么原因扩容的

按照当前的每秒请求速率，部署永远不会达到10个pod的最大值。三个复制品足以使每个吊舱的RPS保持在10以下。

查看pod是否增加
kubectl get pod

负载测试结束后几分钟 pod会降到2个
```

# 原理讲解
1. prometheus的配置加上
kubernetes_sd_configs:
      - role: pod
代表自动发现k8s的pod
[解释](https://segmentfault.com/a/1190000013230914)
[kubernetes_sd_configs](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#kubernetes_sd_config)
2. Deployment 文件需要加上
annotations:
    prometheus.io/scrape，为true则会将pod作为监控目标。
    prometheus.io/path，默认为/metrics
    prometheus.io/port , 端口

根据prometheus.io/scrape: 'true'获知对应的endpoint是需要被scrape的
根据prometheus.io/app-metrics: 'true'获知对应的endpoint中有应用进程暴露的metrics
根据prometheus.io/app-metrics-port: '8080'获知进程暴露的metrics的端口号
根据prometheus.io/app-metrics-path: '/metrics'获知进程暴露的metrics的具体路径

3. 获取指标
```shell
kubectl get --raw "/apis/custom.metrics.k8s.io/v1beta1/namespaces/default/pods/*/http_requests" | jq .
```


4. 负载测试 (hey,ab)
```
$ # Install hey
$ docker run -it -v /usr/local/bin:/go/bin golang:1.8 go get github.com/rakyll/hey

$ export APP_ENDPOINT=$(kubectl get svc sample-metrics-app -o template --template {{.spec.clusterIP}}); echo ${APP_ENDPOINT}
$ hey -n 50000 -c 1000 http://${APP_ENDPOINT}


下载https://github.com/rakyll/hey

hey --help
```

5. 安装Grafana
```
kubectl create -f yml/pv/pv_10G.yaml
kubectl create -f yml/grafana/grafana_pvc.yaml
kubectl create -f yml/grafana/grafana.yaml

```

5.1. 安装Grafana
```
[database]
# You can configure the database connection by specifying type, host, name, user and password
# as separate properties or as on string using the url properties.

# Either "mysql", "postgres" or "sqlite3", it's your choice
type = postgres
host = 192.168.110.231:5432
name = grafana
user = postgres
# If the password contains # or ; you have to wrap it with triple quotes. Ex """#password;"""
password = postgres
``` 

5.2 创建configmap grafana-etc
```
kubectl create configmap "grafana-etc" --from-file=grafana.ini --namespace=monitoring
# 完成后，查看创建结果。
kubectl -n monitoring get configmap
```

5.3 
kubectl apply -f grafana-mysql.yaml


访问 http://192.168.1.120:32333/  admin  admin
配置 prometheus地址 http://prometheus.monitoring.svc:9090


## 时间同步
1. 安装ntpdate工具
yum -y install ntp ntpdate
2. 设置系统时间与网络时间同步
ntpdate cn.pool.ntp.org
3. 将系统时间写入硬件时间
hwclock --systohc

4. 查看系统时间
timedatectl

```

#得到
      Local time: 四 2017-09-21 13:54:09 CST
  Universal time: 四 2017-09-21 05:54:09 UTC
        RTC time: 四 2017-09-21 13:54:09
       Time zone: Asia/Shanghai (CST, +0800)
     NTP enabled: no
NTP synchronized: no
 RTC in local TZ: yes

```
如果没有执行步骤3，则Local time与RTC time显示的值可能不一样



QPS（TPS）：每秒钟request/事务 数量
并发数： 系统同一时候处理的request/事务数
响应时间：  一般取平均响应时间

QPS（TPS）= 并发数/平均响应时间    
或者   
并发数 = QPS*平均响应时间


heapster已经被官方废弃（k8s 1.11版本中，HPA已经不再从hepaster获取数据）

CPU内存、HPA指标： 改为[metrics-server](https://github.com/kubernetes-incubator/metrics-server)
基础监控：集成到prometheus中，kubelet将metric信息暴露成prometheus需要的格式，使用[Prometheus Operator](https://github.com/coreos/prometheus-operator)
事件监控：集成到https://github.com/heptiolabs/eventrouter






#  cAdvisor
cAdvisor并不是被部署在每个pod中，而是在节点级别上。它能够自动发现计算机上所有正在运行的容器，并收集到诸如内存、CPU等系统网络指标。

# prometheus-operator

## CustomResourceDefinitions
The Operator acts on the following [custom resource definitions (CRDs):](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/)

- **Prometheus**, which defines a desired Prometheus deployment. The Operator ensures at all times that a deployment matching the resource definition is running.

- **ServiceMonitor**, which declaratively specifies how groups of services should be monitored. The Operator automatically generates Prometheus scrape configuration based on the definition.

- **PodMonitor**, which declaratively specifies how groups of pods should be monitored. The Operator automatically generates Prometheus scrape configuration based on the definition.

- **PrometheusRule**, which defines a desired Prometheus rule file, which can be loaded by a Prometheus instance containing Prometheus alerting and recording rules.

- **Alertmanager**, which defines a desired Alertmanager deployment. The Operator ensures at all times that a deployment matching the resource definition is running.

To learn more about the CRDs introduced by the Prometheus Operator have a look at the [design doc](https://github.com/coreos/prometheus-operator/blob/master/Documentation/design.md).

https://blog.csdn.net/ygqygq2/article/details/83655552
