

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
nohup linkerd dashboard --address 0.0.0.0 & 可能有问题，需要修改-enforced-host
```
containers:
      - args:
        - -api-addr=linkerd-controller-api.linkerd.svc.cluster.local:8085
        - -grafana-addr=linkerd-grafana.linkerd.svc.cluster.local:3000
        - -controller-namespace=linkerd
        - -log-level=info
        - -enforced-host=^(192\.168\.110\.213|localhost|127\.0\.0\.1|linkerd-web\.linkerd\.svc\.cluster\.local|linkerd-web\.linkerd\.svc|\[::1\])(:\d+)?$
        image: gcr.io/linkerd-io/web:stable-2.7.1
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

### ingress
```
# 如果打算用于生产环境，请参考 https://github.com/nginxinc/kubernetes-ingress/blob/v1.5.5/docs/installation.md 并根据您自己的情况做进一步定制

apiVersion: v1
kind: Namespace
metadata:
  name: nginx-ingress

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: nginx-ingress 
  namespace: nginx-ingress

---
apiVersion: v1
kind: Secret
metadata:
  name: default-server-secret
  namespace: nginx-ingress
type: Opaque
data:
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN2akNDQWFZQ0NRREFPRjl0THNhWFhEQU5CZ2txaGtpRzl3MEJBUXNGQURBaE1SOHdIUVlEVlFRRERCWk8KUjBsT1dFbHVaM0psYzNORGIyNTBjbTlzYkdWeU1CNFhEVEU0TURreE1qRTRNRE16TlZvWERUSXpNRGt4TVRFNApNRE16TlZvd0lURWZNQjBHQTFVRUF3d1dUa2RKVGxoSmJtZHlaWE56UTI5dWRISnZiR3hsY2pDQ0FTSXdEUVlKCktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQUwvN2hIUEtFWGRMdjNyaUM3QlBrMTNpWkt5eTlyQ08KR2xZUXYyK2EzUDF0azIrS3YwVGF5aGRCbDRrcnNUcTZzZm8vWUk1Y2Vhbkw4WGM3U1pyQkVRYm9EN2REbWs1Qgo4eDZLS2xHWU5IWlg0Rm5UZ0VPaStlM2ptTFFxRlBSY1kzVnNPazFFeUZBL0JnWlJVbkNHZUtGeERSN0tQdGhyCmtqSXVuektURXUyaDU4Tlp0S21ScUJHdDEwcTNRYzhZT3ExM2FnbmovUWRjc0ZYYTJnMjB1K1lYZDdoZ3krZksKWk4vVUkxQUQ0YzZyM1lma1ZWUmVHd1lxQVp1WXN2V0RKbW1GNWRwdEMzN011cDBPRUxVTExSakZJOTZXNXIwSAo1TmdPc25NWFJNV1hYVlpiNWRxT3R0SmRtS3FhZ25TZ1JQQVpQN2MwQjFQU2FqYzZjNGZRVXpNQ0F3RUFBVEFOCkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQWpLb2tRdGRPcEsrTzhibWVPc3lySmdJSXJycVFVY2ZOUitjb0hZVUoKdGhrYnhITFMzR3VBTWI5dm15VExPY2xxeC9aYzJPblEwMEJCLzlTb0swcitFZ1U2UlVrRWtWcitTTFA3NTdUWgozZWI4dmdPdEduMS9ienM3bzNBaS9kclkrcUI5Q2k1S3lPc3FHTG1US2xFaUtOYkcyR1ZyTWxjS0ZYQU80YTY3Cklnc1hzYktNbTQwV1U3cG9mcGltU1ZmaXFSdkV5YmN3N0NYODF6cFErUyt1eHRYK2VBZ3V0NHh3VlI5d2IyVXYKelhuZk9HbWhWNThDd1dIQnNKa0kxNXhaa2VUWXdSN0diaEFMSkZUUkk3dkhvQXprTWIzbjAxQjQyWjNrN3RXNQpJUDFmTlpIOFUvOWxiUHNoT21FRFZkdjF5ZytVRVJxbStGSis2R0oxeFJGcGZnPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  tls.key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBdi91RWM4b1JkMHUvZXVJTHNFK1RYZUprckxMMnNJNGFWaEMvYjVyYy9XMlRiNHEvClJOcktGMEdYaVN1eE9ycXgrajlnamx4NXFjdnhkenRKbXNFUkJ1Z1B0ME9hVGtIekhvb3FVWmcwZGxmZ1dkT0EKUTZMNTdlT1l0Q29VOUZ4amRXdzZUVVRJVUQ4R0JsRlNjSVo0b1hFTkhzbysyR3VTTWk2Zk1wTVM3YUhudzFtMApxWkdvRWEzWFNyZEJ6eGc2clhkcUNlUDlCMXl3VmRyYURiUzc1aGQzdUdETDU4cGszOVFqVUFQaHpxdmRoK1JWClZGNGJCaW9CbTVpeTlZTW1hWVhsMm0wTGZzeTZuUTRRdFFzdEdNVWozcGJtdlFmazJBNnljeGRFeFpkZFZsdmwKMm82MjBsMllxcHFDZEtCRThCay90elFIVTlKcU56cHpoOUJUTXdJREFRQUJBb0lCQVFDZklHbXowOHhRVmorNwpLZnZJUXQwQ0YzR2MxNld6eDhVNml4MHg4Mm15d1kxUUNlL3BzWE9LZlRxT1h1SENyUlp5TnUvZ2IvUUQ4bUFOCmxOMjRZTWl0TWRJODg5TEZoTkp3QU5OODJDeTczckM5bzVvUDlkazAvYzRIbjAzSkVYNzZ5QjgzQm9rR1FvYksKMjhMNk0rdHUzUmFqNjd6Vmc2d2szaEhrU0pXSzBwV1YrSjdrUkRWYmhDYUZhNk5nMUZNRWxhTlozVDhhUUtyQgpDUDNDeEFTdjYxWTk5TEI4KzNXWVFIK3NYaTVGM01pYVNBZ1BkQUk3WEh1dXFET1lvMU5PL0JoSGt1aVg2QnRtCnorNTZud2pZMy8yUytSRmNBc3JMTnIwMDJZZi9oY0IraVlDNzVWYmcydVd6WTY3TWdOTGQ5VW9RU3BDRkYrVm4KM0cyUnhybnhBb0dCQU40U3M0ZVlPU2huMVpQQjdhTUZsY0k2RHR2S2ErTGZTTXFyY2pOZjJlSEpZNnhubmxKdgpGenpGL2RiVWVTbWxSekR0WkdlcXZXaHFISy9iTjIyeWJhOU1WMDlRQ0JFTk5jNmtWajJTVHpUWkJVbEx4QzYrCk93Z0wyZHhKendWelU0VC84ajdHalRUN05BZVpFS2FvRHFyRG5BYWkyaW5oZU1JVWZHRXFGKzJyQW9HQkFOMVAKK0tZL0lsS3RWRzRKSklQNzBjUis3RmpyeXJpY05iWCtQVzUvOXFHaWxnY2grZ3l4b25BWlBpd2NpeDN3QVpGdwpaZC96ZFB2aTBkWEppc1BSZjRMazg5b2pCUmpiRmRmc2l5UmJYbyt3TFU4NUhRU2NGMnN5aUFPaTVBRHdVU0FkCm45YWFweUNweEFkREtERHdObit3ZFhtaTZ0OHRpSFRkK3RoVDhkaVpBb0dCQUt6Wis1bG9OOTBtYlF4VVh5YUwKMjFSUm9tMGJjcndsTmVCaWNFSmlzaEhYa2xpSVVxZ3hSZklNM2hhUVRUcklKZENFaHFsV01aV0xPb2I2NTNyZgo3aFlMSXM1ZUtka3o0aFRVdnpldm9TMHVXcm9CV2xOVHlGanIrSWhKZnZUc0hpOGdsU3FkbXgySkJhZUFVWUNXCndNdlQ4NmNLclNyNkQrZG8wS05FZzFsL0FvR0FlMkFVdHVFbFNqLzBmRzgrV3hHc1RFV1JqclRNUzRSUjhRWXQKeXdjdFA4aDZxTGxKUTRCWGxQU05rMXZLTmtOUkxIb2pZT2pCQTViYjhibXNVU1BlV09NNENoaFJ4QnlHbmR2eAphYkJDRkFwY0IvbEg4d1R0alVZYlN5T294ZGt5OEp0ek90ajJhS0FiZHd6NlArWDZDODhjZmxYVFo5MWpYL3RMCjF3TmRKS2tDZ1lCbyt0UzB5TzJ2SWFmK2UwSkN5TGhzVDQ5cTN3Zis2QWVqWGx2WDJ1VnRYejN5QTZnbXo5aCsKcDNlK2JMRUxwb3B0WFhNdUFRR0xhUkcrYlNNcjR5dERYbE5ZSndUeThXczNKY3dlSTdqZVp2b0ZpbmNvVlVIMwphdmxoTUVCRGYxSjltSDB5cDBwWUNaS2ROdHNvZEZtQktzVEtQMjJhTmtsVVhCS3gyZzR6cFE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=

---
kind: ConfigMap
apiVersion: v1
metadata:
  name: nginx-config
  namespace: nginx-ingress
data:
  server-names-hash-bucket-size: "1024"
  enable-opentracing: "true"
  zipkin-collector-host: "oc-collector.tracing"


---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: nginx-ingress
rules:
- apiGroups:
  - ""
  resources:
  - services
  - endpoints
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - update
  - create
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - list
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
- apiGroups:
  - extensions
  resources:
  - ingresses
  verbs:
  - list
  - watch
  - get
- apiGroups:
  - "extensions"
  resources:
  - ingresses/status
  verbs:
  - update
- apiGroups:
  - k8s.nginx.org
  resources:
  - virtualservers
  - virtualserverroutes
  verbs:
  - list
  - watch
  - get

---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: nginx-ingress
subjects:
- kind: ServiceAccount
  name: nginx-ingress
  namespace: nginx-ingress
roleRef:
  kind: ClusterRole
  name: nginx-ingress
  apiGroup: rbac.authorization.k8s.io

---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: nginx-ingress
  namespace: nginx-ingress
  annotations:
    prometheus.io/scrape: "true"
    prometheus.io/port: "9113"
spec:
  selector:
    matchLabels:
      app: nginx-ingress
  template:
    metadata:
      labels:
        app: nginx-ingress
    spec:
      serviceAccountName: nginx-ingress
      containers:
      - image: nginx/nginx-ingress:1.5.5
        name: nginx-ingress
        ports:
        - name: http
          containerPort: 80
          hostPort: 80
        - name: https
          containerPort: 443
          hostPort: 443
        - name: prometheus
          containerPort: 9113
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        args:
          - -nginx-configmaps=$(POD_NAMESPACE)/nginx-config
          - -default-server-tls-secret=$(POD_NAMESPACE)/default-server-secret
         #- -v=3 # Enables extensive logging. Useful for troubleshooting.
         #- -report-ingress-status
         #- -external-service=nginx-ingress
         #- -enable-leader-election
          - -enable-prometheus-metrics
         #- -enable-custom-resources

---


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