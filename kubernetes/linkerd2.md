

## Install

###  CLI
```
curl -sL https://run.linkerd.io/install | sh

环境变量
export PATH=$PATH:$HOME/.linkerd2/bin

or 

vi ~/.bashrc
export PATH=/root/.linkerd2/bin:$PATH
source ~/.bashrc

验证
linkerd version

```

### 验证k8s集群 & 安装Linkerd 到k8s集群
```
预验证
linkerd check --pre

安装
linkerd install | kubectl apply -f -
此命令可生成一个 Kubernetes manifest ，然后使用 kubectl 命令将其应用于 Kubernetes 集群。（在应用之前，请随意检查 manifest ）
基本上image pull不下来

如果你已经linkerd install | kubectl apply -f -   可以执行linkerd install --ignore-cluster | kubectl delete -f - 删除 然后再linkerd install >> deploy-linker.yaml

导出一个yml文件
linkerd install >> deploy-linker.yaml

https://github.com/zhangguanzhang/gcr.io

每台node上执行

curl -s https://zhangguanzhang.github.io/bash/pull.sh | bash -s -- gcr.io/linkerd-io/controller:stable-2.7.1 &&\
curl -s https://zhangguanzhang.github.io/bash/pull.sh | bash -s -- gcr.io/linkerd-io/proxy:stable-2.7.1 &&\
curl -s https://zhangguanzhang.github.io/bash/pull.sh | bash -s -- gcr.io/linkerd-io/proxy-init:v1.3.2 &&\
curl -s https://zhangguanzhang.github.io/bash/pull.sh | bash -s -- gcr.io/linkerd-io/web:stable-2.7.1 &&\
curl -s https://zhangguanzhang.github.io/bash/pull.sh | bash -s -- gcr.io/linkerd-io/debug:stable-2.7.1 &&\
curl -s https://zhangguanzhang.github.io/bash/pull.sh | bash -s -- gcr.io/linkerd-io/grafana:stable-2.7.1 

docker pull prom/prometheus:v2.15.2


deploy-linker.yaml中修改 增加忽略的主机 192\.168\.110\.213|

apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    linkerd.io/created-by: linkerd/cli stable-2.7.1
  labels:
    app.kubernetes.io/name: web
    app.kubernetes.io/part-of: Linkerd
    app.kubernetes.io/version: stable-2.7.1
    linkerd.io/control-plane-component: web
    linkerd.io/control-plane-ns: linkerd
  name: linkerd-web
  namespace: linkerd
spec:
  replicas: 1
  selector:
    matchLabels:
      linkerd.io/control-plane-component: web
      linkerd.io/control-plane-ns: linkerd
      linkerd.io/proxy-deployment: linkerd-web
  template:
    metadata:
      annotations:
        linkerd.io/created-by: linkerd/cli stable-2.7.1
        linkerd.io/identity-mode: default
        linkerd.io/proxy-version: stable-2.7.1
      labels:
        linkerd.io/control-plane-component: web
        linkerd.io/control-plane-ns: linkerd
        linkerd.io/proxy-deployment: linkerd-web
    spec:
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
      - args:
        - -api-addr=linkerd-controller-api.linkerd.svc.cluster.local:8085
        - -grafana-addr=linkerd-grafana.linkerd.svc.cluster.local:3000
        - -controller-namespace=linkerd
        - -log-level=info
        - -enforced-host=^(192\.168\.110\.213|localhost|127\.0\.0\.1|linkerd-web\.linkerd\.svc\.cluster\.local|linkerd-web\.linkerd\.svc|\[::1\])(:\d+)?$



kubectl apply -f .

验证
linkerd check

查看 deployment
kubectl -n linkerd get deploy
```

pull.sh
```sh
#!/bin/bash
[ -z "$set_e" ] && set -e

[ -z "$1" ] && { echo '$1 is not set';exit 2; }



# imgFullName 
sync_pull(){
    local targetName pullName
    targetName=$1
    pullName=${1//k8s.gcr.io/gcr.io\/google_containers}
    pullName=${pullName//google-containers/google_containers}
    if [ $( tr -dc '/' <<< $pullName | wc -c) -gt 2 ];then #大于2为gcr的超长镜像名字
        pullName=$(echo $pullName | sed -r 's#io#azk8s.cn#')
    else
        pullName=zhangguanzhang/${pullName//\//.}
    fi
    docker pull $pullName
    docker tag $pullName $targetName
    docker rmi $pullName
}

if [ "$1" == search ];then
    shift
    which jq &> /dev/null || { echo 'search needs jq, please install the jq';exit 2; }
    img=${1%/}
    [[ $img == *:* ]] && img_name=${img/://} || img_name=$img
    if [[ "$img" =~ ^gcr.io|^k8s.gcr.io ]];then
        if [[ "$img" =~ ^k8s.gcr.io ]];then
            img_name=${img_name/k8s.gcr.io\//gcr.io/google_containers/}
        elif [[ "$img" == *google-containers* ]]; then
            img_name=${img_name/google-containers/google_containers}
        fi
        repository=gcr.io
    elif [[ "$img" =~ ^quay.io ]];then
            repository=quay.io 
    else 
        echo 'not sync the namespaces!';exit 0;
    fi
    #echo '查询用的github,GFW原因可能会比较久,请确保能访问到github'
    curl -sX GET https://api.github.com/repos/zhangguanzhang/${repository}/contents/$img_name?ref=develop |
        jq -r 'length as $len | if $len ==2 then .message elif $len ==12 then .name else .[].name  end'
else
    img=$1

    if [[ "$img" =~ ^gcr.io|^quay.io|^k8s.gcr.io ]];then
        sync_pull $1
    else
        echo 'not sync the namespaces!';exit 0;
    fi
fi

```



### 启动仪表板
```

nohup linkerd dashboard --address 0.0.0.0 &

linkerd -n linkerd top deploy/linkerd-web

http://192.168.110.213:50750/namespaces
```

### demo
每个pod注入2个容器

- linkerd-init，一个Kubernetes初始化容器，它配置iptables通过代理自动转发所有传入和传出的TCP流量。 （请注意，如果已启用Linkerd CNI插件，则此容器不存在。）
- linkerd-proxy，Linkerd数据平面代理本身。

```




自动注入
linkerd.io/inject: enabled

linkerd.io/inject: disabled


kubectl annotate namespace test linkerd.io/inject=enabled
kubectl annotate namespace test linkerd.io/inject-

手动注入
linkerd inject 

# Inject all the deployments in the default namespace.
kubectl get deploy -o yaml | linkerd inject - | kubectl apply -f -

# Injecting a file from a remote URL
linkerd inject http://url.to/yml | kubectl apply -f -

# Inject all the resources inside a folder and its sub-folders.
linkerd inject <folder> | kubectl apply -f -


kubectl get deploy -o yaml -n loc | linkerd inject - | kubectl apply -f -



检查是否成功
linkerd -n loc check --proxy

查看
linkerd -n emojivoto stat deploy
linkerd -n emojivoto top deploy
linkerd -n emojivoto tap deploy/web
```


### Distributed tracing

先安装ingress-nginx
```
curl -s https://zhangguanzhang.github.io/bash/pull.sh | bash -s -- quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.30.0



kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud-generic.yaml
```

```
安装 opencensus-collector
kubectl apply -f https://run.linkerd.io/tracing/collector.yml
等待安装完成
kubectl -n tracing rollout status deploy/oc-collector


github.com/census-instrumentation/opencensus-service

ConfigMap ：将collector-endpoint改成collector_endpoint 
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: oc-collector-conf
  namespace: tracing
  labels:
    app: opencensus
    component: oc-collector-conf
data:
  oc-collector-config: |
    receivers:
      opencensus:
        port: 55678
      zipkin:
        port: 9411
    exporters:
      jaeger:
        collector_endpoint: "http://192.168.110.252:14268/api/traces"
---



安装 Jaeger
kubectl apply -f https://run.linkerd.io/tracing/backend.yml
等待安装完成
kubectl -n tracing rollout status deploy/jaeger


查看Jaeger
kubectl -n tracing port-forward svc/jaeger 16686 --address=0.0.0.0 &
#kubectl -n emojivoto port-forward svc/web-svc 8080:80

使用
spec:
  template:
    metadata:
      annotations:
        linkerd.io/inject: enabled
        config.linkerd.io/trace-collector: oc-collector.tracing:55678


ingress-nginx 开启tracing

controller:
  config:
    enable-opentracing: "true"
    zipkin-collector-host: oc-collector.tracing

```

### 配置超时
```
apiVersion: linkerd.io/v1alpha2
kind: ServiceProfile
metadata:
  name: xxxx.default.svc.cluster.local
  namespace: default
spec:
  # A service profile defines a list of routes.  Linkerd can aggregate metrics
  # like request volume, latency, and success rate by route.
  routes:
  - name: '/xxx.xxx/xxx'
    timeout: 25ms
    # Each route must define a condition.  All requests that match the
    # condition will be counted as belonging to that route.  If a request
    # matches more than one route, the first match wins.
    condition:
      # The simplest condition is a path regular expression.
      pathRegex: '/xxx/xxx'
      # This is a condition that checks the request method.
      method: POST
```


### 获取path指标
linkerd routes svc/webapp
linkerd routes deploy/webapp

linkerd routes deploy/webapp --to svc/books


### 删除
```
linkerd install --ignore-cluster | kubectl delete -f -
```