kubectl apply -f zipkin.yaml

kubectl get --namespace=kube-system pod

kubectl get --namespace=kube-system services

http://127.0.0.1:30002/zipkin/