# 超好用的K8s诊断工具：kubectl-debug
https://github.com/aylei/kubectl-debug/releases

在K8s环境部署应用后，经常遇到需要进入pod进行排错。除了查看pod logs和describe方式之外，传统的解决方式是在业务pod基础镜像中提前安装好procps、net-tools、tcpdump、vim等工具。但这样既不符合最小化镜像原则，又徒增Pod安全漏洞风险。

kubectl-debug是一个简单、易用、强大的 kubectl 插件, 能够帮助你便捷地进行 Kubernetes 上的 Pod 排障诊断。它通过启动一个排错工具容器，并将其加入到目标业务容器的pid, network, user 以及 ipc namespace 中，这时我们就可以在新容器中直接用 netstat, tcpdump 这些熟悉的工具来解决问题了, 而业务容器可以保持最小化, 不需要预装任何额外的排障工具。
kubectl-debug 包含两部分:
kubectl-debug：命令行工具；
debug-agent：部署在K8s的node上，用于启动关联排错工具容器；

k8s 1.12 支持kubectl插件（kubectl debug 命令,其实就是执行了kubectl-debug），之前使用kubectl-debug命令

```
curl -Lo kubectl-debug.tar.gz https://github.com/aylei/kubectl-debug/releases/download/v0.1.1/kubectl-debug_0.1.1_linux_amd64.tar.gz

tar -zxvf kubectl-debug.tar.gz kubectl-debug
mv kubectl-debug /usr/local/bin/

# 可选 安装 debug-agent DaemonSet
kubectl-debug 包含两部分, 一部分是用户侧的 kubectl 插件, 另一部分是部署在所有 k8s 节点上的 agent(用于启动"新容器", 同时也作为 SPDY 连接的中继). 在 agentless 中, kubectl-debug 会在 debug 开始时创建 debug-agent Pod, 并在结束后自动清理.(默认开启agentless模式)

agentless 虽然方便, 但会让 debug 的启动速度显著下降, 你可以通过预先安装 debug-agent 的 DaemonSet 并配合 --agentless=false 参数来使用 agent 模式, 加快启动速度:

# 如果你的kubernetes版本为v1.16或更高
kubectl apply -f https://raw.githubusercontent.com/aylei/kubectl-debug/master/scripts/agent_daemonset.yml
# 如果你使用的是旧版本的kubernetes(<v1.16), 你需要先将apiVersion修改为extensions/v1beta1, 可以如下操作
wget https://raw.githubusercontent.com/aylei/kubectl-debug/master/scripts/agent_daemonset.yml
sed -i '' '1s/apps\/v1/extensions\/v1beta1/g' agent_daemonset.yml
kubectl apply -f agent_daemonset.yml
# 或者使用helm安装
helm install kubectl-debug -n=debug-agent ./contrib/helm/kubectl-debug
# 使用daemonset agent模式(关闭agentless模式)
kubectl debug --agentless=false POD_NAME



# 简单使用:

# kubectl 1.12.0 或更高的版本, 可以直接使用:
kubectl debug -h
# 假如安装了 debug-agent 的 daemonset, 可以使用 --agentless=false 来加快启动速度
# 之后的命令里会使用默认的agentless模式
kubectl debug POD_NAME

# 假如 Pod 处于 CrashLookBackoff 状态无法连接, 可以复制一个完全相同的 Pod 来进行诊断
kubectl debug POD_NAME --fork

# 当使用fork mode时,如果需要复制出来的pod保留原pod的labels,可以使用 --fork-pod-retain-labels 参数进行设置(注意逗号分隔,且不允许空格)
# 示例如下
# 若不设置,该参数默认为空(既不保留原pod的任何labels,fork出来的新pod的labels为空)
kubectl debug POD_NAME --fork --fork-pod-retain-labels=<labelKeyA>,<labelKeyB>,<labelKeyC>

# 为了使 没有公网 IP 或无法直接访问(防火墙等原因)的 NODE 能够访问, 默认开启 port-forward 模式
# 如果不需要开启port-forward模式, 可以使用 --port-forward=false 来关闭
kubectl debug POD_NAME --port-forward=false --agentless=false --daemonset-ns=kube-system --daemonset-name=debug-agent

# 老版本的 kubectl 无法自动发现插件, 需要直接调用 binary
kubectl-debug POD_NAME

# 使用私有仓库镜像,并设置私有仓库使用的kubernetes secret
# secret data原文请设置为 {Username: <username>, Password: <password>}
# 默认secret_name为kubectl-debug-registry-secret,默认namspace为default
kubectl-debug POD_NAME --image calmkart/netshoot:latest --registry-secret-name <k8s_secret_name> --registry-secret-namespace <namespace>

# 在默认的agentless模式中,你可以设置agent pod的resource资源限制,如下示例
# 若不设置,默认为空
kubectl-debug POD_NAME --agent-pod-cpu-requests=250m --agent-pod-cpu-limits=500m --agent-pod-memory-requests=200Mi --agent-pod-memory-limits=500Mi




kubectl debug podname

ps -ef # 查看进程

netstat

logout #相当于会把相应的这个 debug pod 杀掉，然后进行退出，此时对应用实际上是没有任何的影响的

...
```


```
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    app: debug-agent
  name: debug-agent
spec:
  selector:
    matchLabels:
      app: debug-agent
  template:
    metadata:
      labels:
        app: debug-agent
    spec:
      hostPID: true
      tolerations:
        - key: node-role.kubernetes.io/master
          effect: NoSchedule
      containers:
        - name: debug-agent
          image: aylei/debug-agent:latest
          imagePullPolicy: Always
          securityContext:
            privileged: true
          livenessProbe:
            failureThreshold: 3
            httpGet:
              path: /healthz
              port: 10027
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 1
          ports:
            - containerPort: 10027
              hostPort: 10027
              name: http
              protocol: TCP
          volumeMounts:
            - name: cgroup
              mountPath: /sys/fs/cgroup
            - name: lxcfs
              mountPath: /var/lib/lxc/lxcfs
              mountPropagation: Bidirectional
            - name: docker
              mountPath: "/var/run/docker.sock"
      # hostNetwork: true
      volumes:
        - name: cgroup
          hostPath:
            path: /sys/fs/cgroup
        - name: lxcfs
          hostPath:
            path: /var/lib/lxc/lxcfs
            type: DirectoryOrCreate
        - name: docker
          hostPath:
            path: /var/run/docker.sock
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 5
    type: RollingUpdate
```


