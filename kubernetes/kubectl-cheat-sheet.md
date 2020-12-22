https://kubernetes.io/docs/reference/kubectl/cheatsheet/
https://cheatsheet.dennyzhang.com/cheatsheet-kubernetes-a4
https://github.com/dennyzhang/cheatsheet-kubernetes-A4


```
# 如何查找非 running 状态的 Pod 呢？
kubectl get pods -A --field-selector=status.phase!=Running | grep -v Complete

# 如何查找 running 状态的 Pod 呢？
kubectl get pods -A --field-selector=status.phase=Running | grep -v Complete

# 获取节点列表，其中包含运行在每个节点上的 Pod 数量？
kubectl get po -o json --all-namespaces |    jq '.items | group_by(.spec.nodeName) | map({"nodeName": .[0].spec.nodeName, "count": length}) | sort_by(.count)'

# 使用 kubectl top 获取 Pod 列表并根据其消耗的 CPU 或 内存进行排序
# 获取 cpu
$ kubectl top pods -A | sort --reverse --key 3 --numeric

# 获取 memory
$ kubectl top pods -A | sort --reverse --key 4 --numeric
```
给 Node 预留资源
```
evictionHard:
  imagefs.available: 15%
  memory.available: 1G
  nodefs.available: 10%
  nodefs.inodesFree: 5%
```