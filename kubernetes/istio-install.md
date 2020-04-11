https://istio.io/docs/setup/getting-started/

<!-- toc -->
[TOC]

# kiali
istio的管理界面
https://github.com/kiali/kiali

# 安装1.5

## 安装istioctl
```
curl -L https://istio.io/downloadIstio | sh -
or 
自己下载
tar -zxvf istio-1.5.1-linux.tar.gz

cd istio-1.5.1

将istioctl客户端添加到您的路径（Linux或macOS）：
当前终端有效 export PATH=$PWD/bin:$PATH

vi ~/.bashrc
export PATH=/root/istio-1.5.1/bin:$PATH

source ~/.bashrc
```
## 安装 istio
```
istioctl manifest apply --set profile=demo
```
x|default|	demo|	minimal|	remote
---|---|---|---|---
**Core components** |
istio-egressgateway	|	|X		
istio-ingressgateway |X|X		
istio-pilot |	X|	X	|X	
**Addons** |
grafana	||	X		
istio-tracing	||	X		
kiali	||	X		
prometheus |	X|	X||		X


## 面板UI
```
istioctl dashboard --help

istioctl dashboard kiali
istioctl dashboard grafana
istioctl dashboard jaeger
istioctl dashboard controlz

不能访问
```

方式二
```
kubectl get services -n istio-system -o wide

grafana           3000
jaeger-query             16686/TCP 
kiali                 20001/TCP 
prometheus             9090/TCP 
zipkin                     9411/TCP 

kubectl --namespace istio-system  port-forward --address 0.0.0.0 svc/kiali 20001
默认 admin admin

kubectl --namespace istio-system  port-forward --address 0.0.0.0 svc/jaeger-query 16686

kubectl --namespace istio-system  port-forward --address 0.0.0.0 svc/grafana 3000
```

## 卸载
```
istioctl manifest generate --set profile=demo | kubectl delete -f -
```




# 安装 1.4

## 1. 下载
```
curl -L https://istio.io/downloadIstio | sh -

cd istio-1.4.2

export PATH=$PWD/bin:$PATH
或者直接下载
https://github.com/istio/istio/releases

istioctl-1.4.2-linux.tar.gz
解压
tar -zxvf istioctl-1.4.2-linux.tar.gz

拷贝istioctl到 root/bin
```

## 2. Install the demo profile
```
istioctl manifest apply --set profile=demo

kubectl get svc -n istio-system

kubectl get pods -n istio-system
```

### profile

X 代表有

0 |default|	demo|	minimal|	sds|	remote
--|--|--|--|--|--
Core components|
istio-citadel |X|X||X|X
istio-egressgateway ||X
istio-galley|X|X||X
istio-ingressgateway|X|X||X
istio-nodeagent||||X
istio-pilot|X|X|X|X
istio-policy|X|X||X
istio-sidecar-injector|X|X||X|X
istio-telemetry|X|X||X
Addons | |||
grafana||X
istio-tracing||X
kiali||X
prometheus|X|X||X

## 3. 使用
当您使用kubectl apply部署应用程序时，如果在标有istio-injection = enabled的命名空间中启动Envoy容器，则Istio sidecar注入器将自动将Envoy容器注入到您的应用程序pod中：
```
$ kubectl label namespace <namespace> istio-injection=enabled
$ kubectl create -n <namespace> -f <your-app-spec>.yaml

# 关闭注入
kubectl label namespace default istio-injection-
```
在没有istio-injection标签的名称空间中，可以使用istioctl kube-inject在部署它们之前在应用程序pod中手动注入Envoy容器：
```
istioctl kube-inject -f <your-app-spec>.yaml | kubectl apply -f -

# 或者

kubectl apply -f <(istioctl kube-inject -f sleep.yaml)
```

## 4. Uninstall

```
istioctl manifest generate --set profile=demo | kubectl delete -f -
```

# 例子

## 开启手工注入
```
kubectl create namespace istio-test

kubectl label namespace istio-test istio-injection=enabled

# 查看
kubectl get namespace -L istio-injection
```

## create pod
https://github.com/istio/istio/blob/1.4.0/samples/sleep/sleep.yaml

```
1. yaml里没有命名空间
kubectl create -n istio-test -f sleep.yaml

kubectl -n istio-test get pod
能看到ready是2/2

2. yaml里有命名空间
kubectl create -f sleep.yaml

```
如果有`sidecar.istio.io/inject` 参数，以这个为准，不管命名空间是否开启或者关闭注入
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ignored
spec:
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"
    spec:
      containers:
      - name: ignored
        image: tutum/curl
        command: ["/bin/sleep","infinity"]
```

## 查看资源

```
# 查看资源
kubectl top pod -n istio-test

输出
NAME                    CPU(cores)   MEMORY(bytes)   
sleep-f8cbf5b76-pc4b5   4m           31Mi 

# 关闭自动注入
kubectl label namespace istio-test istio-injection-

# 删除pod
kubectl delete pod -n istio-test -l app=sleep

# 查看pod
kubectl -n istio-test get pod

ready 变成1了


# 查看资源
kubectl top pod -n istio-test

输出
NAME                    CPU(cores)   MEMORY(bytes)   
sleep-f8cbf5b76-vqxp9   0m           0Mi
```



