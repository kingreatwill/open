kubectl apply -f .   #  创建当前目录下所有的资源

kubectl get po # 查看目前所有的pod
kubectl get rs # 查看目前所有的replica set
kubectl get deployment # 查看目前所有的deployment
kubectl describe po my-nginx # 查看my-nginx pod的详细状态
kubectl describe rs my-nginx # 查看my-nginx replica set的详细状态
kubectl describe deployment my-nginx # 查看my-nginx deployment的详细状态

# deployment

kubectl create -f nginx-deployment.yaml

kubectl get deployments
kubectl get rs
kubectl get pods --show-labels

升级方法1
kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1
升级方法2 直接编辑
kubectl edit deployment/nginx-deployment

kubectl describe deployments

# k8s创建资源

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

删除

kubectl delete -f  nginx.yaml