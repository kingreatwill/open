# k8s hpa中的指标来源及实现

- [指标类型](#指标类型)
-- [Object类型](#Object类型)
-- [Pods类型](#Pods类型)
-- [Resource类型](#Resource类型)
-- [External类型](#External类型)
- [指标来源](#指标来源)
-- [metrics.k8s.io](#metrics.k8s.io)
-- [custom.metrics.k8s.io](#custom.metrics.k8s.io)
-- [external.metrics.k8s.io](#external.metrics.k8s.io)

## 指标类型
在HorizontalPodAutoscaler 中包含以下几种指标源类型：
- Object(v2beta1)
- Pods(v2beta1)
- Resource(v2beta1)
- External(v2beta2)

### Object类型
Object类型是用于描述k8s内置对象的指标，例如ingress对象中hits-per-second指标
```golang
type ObjectMetricSource struct {
	// 用于描述k8s对象.
	Target CrossVersionObjectReference `json:"target" protobuf:"bytes,1,name=target"`
	// 所查询的metric名称.
	MetricName string `json:"metricName" protobuf:"bytes,2,name=metricName"`
	// 指标的目标值
	TargetValue resource.Quantity `json:"targetValue" protobuf:"bytes,3,name=targetValue"`
}
```
示例：
```yaml
kind: HorizontalPodAutoscaler
apiVersion: autoscaling/v2alpha1
spec:
  scaleTargetRef:
    kind: ReplicationController
    name: WebFrontend
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Object
    object:
      target:
        kind: Service
        name: Frontend
      metricName: hits-per-second
      targetValue: 1k
```
### Pods类型
pods类型描述当前扩容目标中每个pod的指标（例如，每秒处理的事务数）。在与目标值进行比较之前，将对值进行平均。
```golang
type PodsMetricSource struct {
	// 所查询的metric名称
	MetricName string `json:"metricName" protobuf:"bytes,1,name=metricName"`
	// 目标平均值
	TargetAverageValue resource.Quantity `json:"targetAverageValue" protobuf:"bytes,2,name=targetAverageValue"`
}
```
示例：
```yaml
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
  - type: Pods
    pods:
      metricName: http_requests
      targetAverageValue: 10
```
### Resource类型
Resource是Kubernetes已知的资源指标,如request和limit中所指定的，描述当前扩容目标的每个pod（例如CPU或内存）。该指标将会在与目标值对比前进行平均，被此类指标内置于Kubernetes中,且使用"pods"源在正常的每个pod度量标准之上提供特殊的扩展选项。
```golang
type ResourceMetricSource struct {
	// 所查询的metric名称
	Name v1.ResourceName `json:"name" protobuf:"bytes,1,name=name"`
	// 目标使用平均百分比
	tilization *int32 `json:"targetAverageUtilization,omitempty" protobuf:"varint,2,opt,name=targetAverageUtilization"`
	// 目标平均值
	TargetAverageValue *resource.Quantity `json:"targetAverageValue,omitempty" protobuf:"bytes,3,opt,name=targetAverageValue"`
}
```
示例：
```yaml
kind: HorizontalPodAutoscaler
apiVersion: autoscaling/v2alpha1
spec:
  scaleTargetRef:
    kind: ReplicationController
    name: WebFrontend
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: 80
```

### External类型
ExternalMetricSource指示如何扩展与任何Kubernetes对象无关的指标（例如，云消息传递服务中的队列长度，或集群外部运行的负载均衡器的QPS）。
```golang
type ExternalMetricSource struct {
	// 度量通过名称和选择器标识目标度量。
	Metric MetricIdentifier `json:"metric" protobuf:"bytes,1,name=metric"`
	// target specifies the target value for the given metric
	// 为指定指标指定目标值
	Target MetricTarget `json:"target" protobuf:"bytes,2,name=target"`
}
```
示例：
```yaml
kind: HorizontalPodAutoscaler
apiVersion: autoscaling/v2beta2
spec:
  scaleTargetRef:
    kind: ReplicationController
    name: Worker
  minReplicas: 2
  maxReplicas: 10
  metrics:
   - type: External
     external:
       metricName: queue_messages_ready
       metricSelector:
         matchLabels:
           queue: worker_tasks
       targetAverageValue: 30
```

## 指标来源

hpa controller可通过以下api来获取指标实现自动扩容：

- metrics.k8s.io（resource metric api）
- custom.metrics.k8s.io (custom metrics api)
- external.metrics.k8s.io (external metrics api)

### metrics.k8s.io
用于hpa 中Resource类型数据来源.
已有实现：
- [metric-server](https://github.com/kubernetes-sigs/metrics-server)
- [heapster](https://github.com/kubernetes-retired/heapster) 已归档

### custom.metrics.k8s.io
用于hpa 中object/pods类型的数据来源，需要自己实现适配器(custom metrics adapters)
已有实现：
- [Prometheus Adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter)
- [Microsoft Azure Adapter](https://github.com/Azure/azure-k8s-metrics-adapter)
- [Google Stackdriver](https://github.com/GoogleCloudPlatform/k8s-stackdriver)

### external.metrics.k8s.io
用于 hpa 中external类型的数据来源，同样需要云厂商或平台自己实现适配器（custom metrics adapters）

已有实现：
- [Stackdriver](https://github.com/GoogleCloudPlatform/k8s-stackdriver)
- [k8s-prometheus-adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter)

参考:
- https://github.com/kubernetes/metrics/blob/master/IMPLEMENTATIONS.md
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-custom-metrics
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/hpa-v2.md
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/hpa-external-metrics.md
- https://github.com/kubernetes/api/tree/master/autoscaling/v2beta2
- https://github.com/kubernetes/metrics/tree/master/pkg/apis