https://www.hangge.com/blog/cache/detail_2444.html
https://www.kubernetes.org.cn/4278.html
https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-quickstart.html

https://www.elastic.co/guide/en/elasticsearch/reference/7.4/getting-started.html
https://github.com/kubernetes/kubernetes/tree/master/cluster/addons/fluentd-elasticsearch


1. 进入yml fluentd-es

2. 执行
```
kubectl apply -f .
```

3. 查看
```
kubectl get --namespace=kube-system pod

kubectl get --namespace=kube-system services

kubectl logs elasticsearch-logging-0 -n kube-system

本机ip加随机分配的端口，访问
```

4. 删除
```
kubectl delete -f  .
```