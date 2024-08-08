
## grpc-web
https://github.com/BuiltCloud/grpc-web-demo

![](../img/grpc-web-demo.png)

1. 创建grpc-service
https://github.com/KnowitSolutions/Grpc.Helpers
2. 创建grpc-web
https://github.com/grpc/grpc-web/tree/master/net/grpc/gateway/examples/helloworld
grpc-web
```
protoc -I=. helloworld.proto \
  --js_out=import_style=commonjs:. \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
```
```
  npm install
```
```
npx webpack client.js
```
```
dotnet serve
```

https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-WEB.md#protocol-differences-vs-grpc-over-http2

## protoc
protoc -h
```
  --cpp_out=OUT_DIR           Generate C++ header and source.
  --csharp_out=OUT_DIR        Generate C# source file.
  --java_out=OUT_DIR          Generate Java source file.
  --javanano_out=OUT_DIR      Generate Java Nano source file.
  --js_out=OUT_DIR            Generate JavaScript source.
  --objc_out=OUT_DIR          Generate Objective C header and source.
  --php_out=OUT_DIR           Generate PHP source file.
  --python_out=OUT_DIR        Generate Python source file.
  --ruby_out=OUT_DIR          Generate Ruby source file.
```

```
protoc --version

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip

unzip protoc-3.14.0-linux-x86_64.zip
```

如果你打算用其中的包含的其他类型，同时需要将include目录的内容也复制到某个地方，例如输入/usr/local/include/

我们把protoc放在/usr/local/bin可执行程序目录中，这样全局都可以访问到，同时把include目录的内容也复制到/usr/local/include/中
```
# 移动安装proto (cd到解压目录bin中后执行)
mv proto /usr/local/bin

# 把`include`目录的内容复制(cd到解压目录include中后执行)
cp google /usr/local/include
```

### protoc-gen-go
1. 安装protoc (前面安装过的可以省略)

下载连接： https://github.com/protocolbuffers/protobuf/releases 

选择相应的版本下载并解压到制定目录

Linux： unzip protoc-3.10.0-rc-1-linux-x86_64.zip -d /usr/local/

2. 安装protoc-gen-go（golang安装版本）

go get -u github.com/golang/protobuf/protoc-gen-go

如果不行就编译
```
cd github.com/golang/protobuf/protoc-gen-go

go build

go install
```

> protoc-gen-go也可以在这里下载 https://github.com/protocolbuffers/protobuf-go/releases

https://grpc.io/docs/languages/go/quickstart/
`--plugin=protoc-gen-go`

```
go install github.com/golang/protobuf/protoc-gen-go
protoc --go_out=plugins=grpc:E:/pbgo/ --proto_path=E:/protobuf xxxx.proto
```


官网例子跟上面的protoc-gen-go不是一样的
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## grpc负载均衡问题




### 1. 概述

系统中多个服务间的调用用的是 gRPC 进行通信，最初没考虑到负载均衡的问题，因为用的是 Kubernetes，想的是直接用 K8s 的 [Service](https://www.lixueduan.com/post/kubernetes/04-service-core/) 不就可以实现负载均衡吗。

但是真正测试的时候才发现，所有流量都进入到了某一个 Pod，这时才意识到负载均衡可能出现了问题。

因为 gRPC 是基于 HTTP/2 之上的，而 HTTP/2 被设计为一个长期存在的 TCP 连接，所有都通过该连接进行多路复用。

> 这样虽然减少了管理连接的开销，但是在负载均衡上又引出了新的问题。

由于我们无法在连接层面进行均衡，为了做 gRPC 负载均衡，我们需要从**连接级均衡**转向**请求级均衡**。

> 换句话说，我们需要打开一个到每个目的地的 HTTP/2 连接，并平衡这些连接之间的请求。

**这就意味着我们需要一个 7 层负载均衡，而 K8s 的 Service 核心使用的是 kube proxy，这是一个 4 层负载均衡，所以不能满足我们的要求。**

整理了一下大致有以下几种方案：

* 1）每次都重新建立连接，用完后关闭连接，直接从源头上解决问题。
  * ？？？这算什么方案哈哈
* 2）客户端负载均衡
* 3）服务端负载均衡



### 2. 客户端负载均衡

这也是比较容易实现的方案，具体为：[NameResolver](https://github.com/grpc/grpc/blob/master/doc/naming.md) + [load balancing policy](https://github.com/grpc/grpc/blob/master/doc/load-balancing.md)+[Headless-Service](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services)。

相关教程可以看上一篇文章[gRPC系列教程(十二)---客户端负载均衡](https://www.lixueduan.com/post/grpc/12-client-side-loadbalance/)

1）当 gRPC 客户端想要与 gRPC 服务器进行交互时，它首先尝试通过向 resolver 发出名称解析请求来解析服务器名称，解析程序返回已解析IP地址的列表。

2）Kubernetes Headless-Service 在创建的时候会将该服务对应的每个 Pod IP 以 A 记录的形式存储。

3）常见的 gRPC 库都内置了几个负载均衡算法，比如 [gRPC-Go](https://github.com/grpc/grpc-go/tree/master/examples/features/load_balancing#pick_first) 中内置了`pick_first`和`round_robin`两种算法。

* pick_first：尝试连接到第一个地址，如果连接成功，则将其用于所有RPC，如果连接失败，则尝试下一个地址（并继续这样做，直到一个连接成功）。
* round_robin：连接到它看到的所有地址，并依次向每个后端发送一个RPC。例如，第一个RPC将发送到backend-1，第二个RPC将发送到backend-2，第三个RPC将再次发送到backend-1。



所以建立连接时只需要提供一个服务名即可，gRPC Client 会根据 DNS resolver 返回的 IP 列表分别建立连接，请求时使用 round_robin 算法进行负载均衡，选择其中一个连接用来发起请求。

核心代码如下：

```go
	svc := "mygrpc:50051"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:///%s", "dns", svc),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), // 指定轮询负载均衡算法
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
```

主要是配置负载均衡算法：

```go
grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)
```

> 网上很多比较旧的文章用的是`grpc.WithBalancerName("")`，在新版中不推荐使用了。
> 参考[代码](https://github.com/edmarfelipe/grpc-load-balancing-k8s-example)


#### 存在的问题

**当 Pod 扩缩容时 客户端可以感知到并更新连接吗？**

* Pod 缩容后，由于 gRPC 具有连接探活机制，会自动丢弃无效连接。

* Pod 扩容后，没有感知机制，导致后续扩容的 Pod 无法被请求到。

gRPC 连接默认能永久存活，如果将该值降低能改善这个问题。

在服务端做以下设置

```go
	port := conf.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionAge:      time.Minute,
	}))
	pb.RegisterVerifyServer(s, core.Verify)
	log.Println("Serving gRPC on 0.0.0.0" + port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
```

这样每个连接只会使用一分钟，到期后会重新建立连接，相当于对扩容的感知只会延迟 1 分钟。

> 虽然能用，但是并比是那么完美，强迫症表示完成无法接受这个方案。



#### kuberesolver

为了解决以上问题，很容易想到直接在 client 端调用 Kubernetes API 监测 Service 对应的 endpoints 变化，然后动态更新连接信息。

搜了一下发现 Github 上已经有这个思路的解决方案了：[kuberesolver](https://github.com/sercand/kuberesolver)。

```go
// Import the module
import "github.com/sercand/kuberesolver/v3"
	
// Register kuberesolver to grpc before calling grpc.Dial
kuberesolver.RegisterInCluster()
// if schema is 'kubernetes' then grpc will use kuberesolver to resolve addresses
cc, err := grpc.Dial("kubernetes:///service.namespace:portname", opts...)
```

具体就是将 DNSresolver 替换成了自定义的 kuberesolver。

**同时如果 Kubernetes 集群中使用了 RBAC 授权的话需要给 client 所在Pod赋予 endpoint 资源的 get 和 watch 权限。**

具体授权过程如下：

需要分别创建`ServiceAccount`、`Role`、`RoleBinding`3 个实例， k8s 用的也是 RBAC 授权，所以应该比较好理解。

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: vaptcha
  name: grpclb-sa
```



```yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: vaptcha
  name: grpclb-role
rules:
- apiGroups: [""]
  resources: ["endpoints"]
  verbs: ["get", "watch"]
```



```yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: grpclb-rolebinding
  namespace: vaptcha
subjects:
- kind: ServiceAccount
  name: grpclb-sa
  namespace: vaptcha
roleRef:
  kind: Role
  name: grpclb-role
  apiGroup: rbac.authorization.k8s.io
```



创建对象

```sh
$ kubectl apply -f svc-account.yaml 
serviceaccount/example-sa created
$ kubectl apply -f role.yaml 
role.rbac.authorization.k8s.io/example-role created
$ kubectl apply -f role-binding.yaml 
rolebinding.rbac.authorization.k8s.io/example-rolebinding created
```



Pod 中指定权限:`serviceAccountName: grpclb-sa`

```yaml
apiVersion: v1
kind: Pod
metadata:
  namespace: mynamespace
  name: sa-token-test
spec:
  containers:
  - name: nginx
    image: nginx:1.7.9
  serviceAccountName: grpclb-sa
```



因为 kuberesolver 是直接调用 Kubernetes API 获取 endpoint 所以不需要创建 Headless Service 了，创建普通 Service 也可以。



### 3. 服务端负载均衡

服务端负载均衡主要是在 Pod 之前增加一个 中间组件，一般为 7 层负载均衡。

client 请求中间组件，由中间组件再去请求后端的 Pod。

常见的组件比如 [Linkerd]([https://linkerd.io](https://linkerd.io/))，或者 ServiceMesh 如 [istio](https://istio.io/) 中的 [envoy](https://www.envoyproxy.io/) 也能实现同样的效果。



### 4. 小结

相比之下更加推荐使用 **客户端负载均衡**。

* 客户端负载均衡更加简单，服务直连性能更高。
* 服务端负载均衡所有请求都需要经过负载均衡组件，相当于是又引入了一个全局热点。
* ServiceMesh 的话对基础设施、技术栈要求比较高，落地比较困难。



### 5. 参考

`https://grpc.io/blog/grpc-load-balancing/`

`https://en.wikipedia.org/wiki/Round-robin_DNS`

`https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/`





[Github]: https://github.com/lixd/grpc-go-example