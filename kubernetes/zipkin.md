
部署zipkin
先部署好es
```
kubectl apply -f zipkin.yaml

kubectl get --namespace=kube-system pod

kubectl get --namespace=kube-system services

http://127.0.0.1:30002/zipkin/
```

部署定时任务
```




kubectl apply -f zipkin-dependencies-job.yaml


删除
kubectl delete cronjob zipkin-dependencies
kubectl delete -f zipkin-dependencies-job.yaml

kubectl get jobs

get jobs --watch
```