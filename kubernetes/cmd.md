<!--toc-->
[TOC]

https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#logs


![](img/components-of-kubernetes.png)

[Kubectl 常用命令大全](https://mp.weixin.qq.com/s/yhRCs2HMizs7SBKbFGY7Pw)
# CMD
kubectl cluster-info
kubectl apply -f .   #  创建当前目录下所有的资源

get -o # `[(-o|--output=)json|yaml|wide|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=...]`
# -o custom-columns=HOST-IP:.status.hostIP,POD-IP:.status.podIP
# -o custom-columns=NAME:metadata.name,STATUS:.status.phase,RUNTIME_CLASS:.spec.runtimeClassName
kubectl get po # 查看目前所有的pod
kubectl get rs # 查看目前所有的replica set
kubectl get deployment # 查看目前所有的deployment
kubectl describe po my-nginx # 查看my-nginx pod的详细状态
kubectl describe rs my-nginx # 查看my-nginx replica set的详细状态
kubectl describe deployment my-nginx # 查看my-nginx deployment的详细状态



## logs
参数 --previous 检索之前的容器日志
> 注意：
当前，如果有其他系统机制执行日志轮转，那么 kubectl logs 仅可查询到最新的日志内容。 比如，一个 10MB 大小的文件，通过logrotate 执行轮转后生成两个文件，一个 10MB 大小，一个为空，所以 kubectl logs 将返回空。


## deployment

kubectl create -f nginx-deployment.yaml

kubectl get deployments
kubectl get rs
kubectl get pods --show-labels

升级方法1
kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1
升级方法2 直接编辑
kubectl edit deployment/nginx-deployment

kubectl describe deployments

## k8s创建资源

1. 用 kubectl 命令直接创建
```
kubectl run nginx-deployment --image=nginx:1.7.9 --replicas=2
```
2. 通过配置文件和 kubectl apply 创建
```
kubectl apply -f nginx.yaml
```

kubectl apply 不但能够创建 Kubernetes 资源，也能对资源进行更新，非常方便。不过 Kubernets 还提供了几个类似的命令，例如 kubectl create、kubectl replace、kubectl edit 和 kubectl patch。

为避免造成不必要的困扰，我们会尽量只使用 kubectl apply，
此命令已经能够应对超过 90% 的场景，事半功倍。

## 删除

kubectl delete -f  nginx.yaml

### 删除Evicted状态的pod
kubectl get pods -n monitoring| grep Evicted | awk '{print $1}' | xargs kubectl delete pod -n monitoring

### 删除非Running状态的pod
kubectl get pods -n monitoring| grep -v Running | awk '{print $1}' | xargs kubectl delete pod -n monitoring


## 设置操作的默认命名空间
```
kubectl config set-context --current --namespace=xxx
# Validate it
kubectl config view | grep namespace:
```

当创建一个Service 时 ，会自动DNS解析 `<service-name>.<namespace-name>.svc.cluster.local`

### 并非所有对象都在命名空间中
大多数 kubernetes 资源（例如 Pod、Service、副本控制器等）都位于某些命名空间中。但是命名空间资源本身并不在命名空间中。而且底层资源，例如 nodes 和持久化卷不属于任何命名空间。

查看哪些 Kubernetes 资源在命名空间中，哪些不在命名空间中：
```
# In a namespace
kubectl api-resources --namespaced=true

# Not in a namespace
kubectl api-resources --namespaced=false
```
### 查询label选择器
```
kubectl get pods -l environment=production,tier=frontend

kubectl get pods -l 'environment in (production),tier in (frontend)'

kubectl get pods -l 'environment in (production, qa)'

kubectl get pods -l 'environment,environment notin (frontend)' //exists 运算符限制不匹配

 
```

### yml - selector
```
selector:
  matchLabels:
    component: redis
  matchExpressions:
    - {key: tier, operator: In, values: [cache]}
    - {key: environment, operator: NotIn, values: [dev]}
```
matchLabels 是由 {key，value} 对组成的映射。matchLabels 映射中的单个 {key，value } 等同于 matchExpressions 的元素，其 key字段为 “key”，operator 为 “In”，而 values 数组仅包含 “value”。matchExpressions 是 pod 选择器要求的列表。有效的运算符包括 In，NotIn，Exists 和 DoesNotExist。在 In 和 NotIn 的情况下，设置的值必须是非空的。来自 matchLabels 和 matchExpressions 的所有要求都是合在一起 – 它们必须都满足才能匹配。

###  字段选择器
metadata.name=my-service
metadata.namespace!=default
status.phase=Pending
```
kubectl get pods --field-selector status.phase=Running -A
kubectl get services  --all-namespaces --field-selector metadata.namespace!=default

kubectl get pods 
kubectl get pods --field-selector ""

下面这个 kubectl 命令将筛选 status.phase 字段不等于 Running 同时 spec.restartPolicy 字段等于 Always 的所有 Pod：
kubectl get pods --field-selector=status.phase!=Running,spec.restartPolicy=Always


kubectl get statefulsets,services --all-namespaces --field-selector metadata.namespace!=default
```

### 推荐使用的标签
```
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app.kubernetes.io/name: mysql // 应用程序的名称
    app.kubernetes.io/instance: wordpress-abcxzy // 用于唯一确定应用实例的名称
    app.kubernetes.io/version: "5.7.21" //应用程序的当前版本（例如，语义版本，修订版哈希等）
    app.kubernetes.io/component: database //架构中的组件
    app.kubernetes.io/part-of: wordpress // 此级别的更高级别应用程序的名称
    app.kubernetes.io/managed-by: helm //	用于管理应用程序的工具
```

## node 标记一个节点为不可调度的
```
kubectl cordon $NODENAME
```

## Patch修改
kubectl apply -f https://run.linkerd.io/emojivoto.yml

kubectl -n emojivoto patch -f https://run.linkerd.io/emojivoto.yml -p '
spec:
  template:
    metadata:
      annotations:
        linkerd.io/inject: enabled
        config.linkerd.io/trace-collector: oc-collector.tracing:55678
'

kubectl patch deployment.v1.apps/nginx-deployment -p '{"spec":{"progressDeadlineSeconds":600}}'
## set resources
kubectl set resources deployment.v1.apps/nginx-deployment -c=nginx --limits=cpu=200m,memory=512Mi 更新其 resource 限制


## set env
kubectl -n emojivoto set env --all deploy OC_AGENT_HOST=oc-collector.tracing:55678

## rollout 

status 查看状态
kubectl -n emojivoto rollout status deploy/web

pause 暂停 Deployment
kubectl rollout pause deployment.v1.apps/nginx-deployment 
继续（resume）该 Deployment
kubectl rollout resume deployment.v1.apps/nginx-deployment

升级
kubectl set image deployment.v1.apps/nginx-deployment nginx=nginx:1.91 --record=true
查看历史
kubectl rollout history deployment/nginx-deployment
kubectl rollout history deployment.v1.apps/nginx-deployment 检查 Deployment 的历史版本

查看 revision（版本）的详细信息
kubectl rollout history deployment.v1.apps/nginx-deployment --revision=2

回滚到前一个 revision（版本）
kubectl rollout undo deployment.v1.apps/nginx-deployment

您也可以使用 --to-revision 选项回滚到前面的某一个指定版本
kubectl rollout undo deployment.v1.apps/nginx-deployment --to-revision=2

> 通过 Deployment 中 .spec.revisionHistoryLimit 字段，可指定为该 Deployment 保留多少个旧的 ReplicaSet。超出该数字的将被在后台进行垃圾回收。该字段的默认值是 10。如果该字段被设为 0，Kubernetes 将清理掉该 Deployment 的所有历史版本（revision），因此，您将无法对该 Deployment 执行回滚操作 kubectl rollout undo。

> 通过 Deployment 中 .spec.strategy 字段，可以指定使用 滚动更新 RollingUpdate 的部署策略还是使用 重新创建 Recreate 的部署策略,如果选择重新创建，Deployment将先删除原有副本集中的所有 Pod，然后再创建新的副本集和新的 Pod。如此，更新过程中将出现一段应用程序不可用的情况；

>  maxSurge = 2  最大超出副本数      可以设置数字或百分比
> 滚动更新过程中，可以超出期望副本数的最大值。
该取值可以是一个绝对值（例如：5），也可以是一个相对于期望副本数的百分比（例如：10%）；
如果填写百分比，则以期望副本数乘以该百分比后向上取整的方式计算对应的绝对值；
当最大超出副本数 maxUnavailable 为 0 时，此数值不能为 0；默认值为 25%。
例如：假设此值被设定为 30%，当滚动更新开始时，新的副本集（ReplicaSet）可以立刻扩容，
但是旧 Pod 和新 Pod 的总数不超过 Deployment 期待副本数（spec.repilcas）的 130%。
一旦旧 Pod 被终止后，新的副本集可以进一步扩容，但是整个滚动更新过程中，新旧 Pod 的总
数不超过 Deployment 期待副本数（spec.repilcas）的 130%。

> maxUnavailable =3 最大不可用副本数 可以设置数字或百分比
> 滚动更新过程中，不可用副本数的最大值。
该取值可以是一个绝对值（例如：5），也可以是一个相对于期望副本数的百分比（例如：10%）；
如果填写百分比，则以期望副本数乘以该百分比后向下取整的方式计算对应的绝对值；
当最大超出副本数 maxSurge 为 0 时，此数值不能为 0；默认值为 25%；
例如：假设此值被设定为 30%，当滚动更新开始时，旧的副本集（ReplicaSet）可以缩容到期望
副本数的 70%；在新副本集扩容的过程中，一旦新的 Pod 已就绪，旧的副本集可以进一步缩容，
整个滚动更新过程中，确保新旧就绪副本数之和不低于期望副本数的 70%。

## scale 
kubectl scale deployment.v1.apps/nginx-deployment --replicas=10

自动伸缩，您就可以基于 CPU 的利用率在一个最大和最小的区间自动伸缩您的 Deployment：
kubectl autoscale deployment.v1.apps/nginx-deployment --min=1 --max=15 --cpu-percent=80




## port-forward
kubectl -n tracing port-forward svc/jaeger 16686 --address=0.0.0.0 &
#kubectl -n emojivoto port-forward svc/web-svc 8080:80

文档: https://kubernetes.io/zh-cn/docs/tasks/access-application-cluster/port-forward-access-application-cluster/

## 查看具体容器
kubectl logs deployment/oc-collector -c oc-collector -n tracing