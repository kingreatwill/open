<!--toc-->
[TOC]
#CMD

kubectl apply -f .   #  创建当前目录下所有的资源

kubectl get po # 查看目前所有的pod
kubectl get rs # 查看目前所有的replica set
kubectl get deployment # 查看目前所有的deployment
kubectl describe po my-nginx # 查看my-nginx pod的详细状态
kubectl describe rs my-nginx # 查看my-nginx replica set的详细状态
kubectl describe deployment my-nginx # 查看my-nginx deployment的详细状态

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

## set env
kubectl -n emojivoto set env --all deploy OC_AGENT_HOST=oc-collector.tracing:55678

## rollout status
kubectl -n emojivoto rollout status deploy/web

## port-forward
kubectl -n tracing port-forward svc/jaeger 16686 --address=0.0.0.0 &
#kubectl -n emojivoto port-forward svc/web-svc 8080:80