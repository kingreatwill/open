https://github.com/apache/flink/tree/release-1.10.1/flink-container/kubernetes

https://ci.apache.org/projects/flink/flink-docs-release-1.10/zh/ops/deployment/kubernetes.html

https://github.com/apache/flink-docker



You can then access the Flink UI via different ways:

- kubectl proxy:

    1. Run `kubectl proxy` in a terminal.
    2. Navigate to http://localhost:8001/api/v1/namespaces/default/services/flink-jobmanager:ui/proxy in your browser.
- kubectl port-forward:

    1. Run `kubectl port-forward ${flink-jobmanager-pod} 8081:8081` to forward your jobmanager’s web ui port to local 8081.
    2. Navigate to `http://localhost:8081` in your browser.
    3. Moreover, you could use the following command below to submit jobs to the cluster:
```
./bin/flink run -m localhost:8081 ./examples/streaming/WordCount.jar
```
- Create a NodePort service on the rest service of jobmanager:
    1. Run kubectl create -f jobmanager-rest-service.yaml to create the NodePort service on jobmanager. The example of jobmanager-rest-service.yaml can be found in appendix.
    2. Run kubectl get svc flink-jobmanager-rest to know the node-port of this service and navigate to `http://<public-node-ip>:<node-port>` in your browser.
    3. Similarly to port-forward solution, you could also use the following command below to submit jobs to the cluster:
```
./bin/flink run -m <public-node-ip>:<node-port> ./examples/streaming/WordCount.jar
```
```
kubectl create namespace flink-cluster
kubectl apply -n flink-cluster -f .


kubectl get svc -n flink-cluster
找到NodePort
```


kubectl create -f flink-configuration-configmap.yaml
kubectl create -f jobmanager-service.yaml
kubectl create -f jobmanager-deployment.yaml
kubectl create -f taskmanager-deployment.yaml


kubectl delete -f jobmanager-deployment.yaml
kubectl delete -f taskmanager-deployment.yaml
kubectl delete -f jobmanager-service.yaml
kubectl delete -f flink-configuration-configmap.yaml