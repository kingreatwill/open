https://istio.io/docs/setup/getting-started/

1. 下载
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

2. Install the demo profile
```
istioctl manifest apply --set profile=demo

kubectl get svc -n istio-system

kubectl get pods -n istio-system
```

3. 当您使用kubectl apply部署应用程序时，如果在标有istio-injection = enabled的命名空间中启动Envoy容器，则Istio sidecar注入器将自动将Envoy容器注入到您的应用程序pod中：
```
$ kubectl label namespace <namespace> istio-injection=enabled
$ kubectl create -n <namespace> -f <your-app-spec>.yaml
```
在没有istio-injection标签的名称空间中，可以使用istioctl kube-inject在部署它们之前在应用程序pod中手动注入Envoy容器：
```
istioctl kube-inject -f <your-app-spec>.yaml | kubectl apply -f -
```

4. Uninstall

```
istioctl manifest generate --set profile=demo | kubectl delete -f -
```