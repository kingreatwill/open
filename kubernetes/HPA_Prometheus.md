Kubernetes将资源指标分为了两种：

* core metrics（核心指标）： 采集每个节点上的kubelet公开的summary api中的指标信息，通常只包含cpu、内存使用率信息
* custom metrics（自定义指标）：允许用户从外部的监控系统当中采集自定义指标，如应用的qps等

在autoscaling/v1版本中只支持CPUUtilizationPercentage一种指标，在autoscaling/v2beta1中增加支持custom metrics
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



部署
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
      targetAverageUtilization: 80
  - type: Resource
    resource:
      name: memory
      targetAverageValue: 200Mi

```

# 2部署 Custom Metrics Server
https://github.com/DirectXMan12/k8s-prometheus-adapter
```

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
