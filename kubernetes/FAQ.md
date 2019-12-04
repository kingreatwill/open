# pod的状态为evicted
eviction，即驱赶的意思，意思是当节点出现异常时，kubernetes将有相应的机制驱赶该节点上的Pod。
多见于资源不足时导致的驱赶。
删除状态为Evicted 的pod
kubectl get pods -n monitoring| grep Evicted | awk '{print $1}' | xargs kubectl delete pod -n monitoring