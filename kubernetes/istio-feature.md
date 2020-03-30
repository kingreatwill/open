<!--toc-->
[TOC]

# Istio
- HTTP，gRPC，WebSocket和TCP通信的自动负载平衡。

- 通过丰富的路由规则，重试，故障转移和故障注入对流量行为进行细粒度控制。

- 可插拔的策略层和配置API，支持访问控制，速率限制和配额。

- 集群内所有流量的自动度量，日志和跟踪，包括集群的入口和出口。

- 通过强大的基于身份的身份验证和授权，在群集中实现安全的服务间通信。


自定义资源定义（CRD） 是默认Kubernetes API的扩展。Istio使用Kubernetes CRD API进行配置，即使是非Kubernetes Istio部署也是如此。


# Istio 流量管理

## Gateway CRD资源 
Istio提供了一些可以预配置的网关代理部署（istio-ingressgateway和istio-egressgateway）
https://istio.io/docs/reference/config/networking/gateway/

```yaml
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: my-gateway
  namespace: some-config-namespace
spec:
  selector:
    app: my-gateway-controller
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - uk.bookinfo.com
    - eu.bookinfo.com
    tls:
      httpsRedirect: true # sends 301 redirect for http requests
  - port:
      number: 443
      name: https-443
      protocol: HTTPS
    hosts:
    - uk.bookinfo.com
    - eu.bookinfo.com
    tls:
      mode: SIMPLE # enables HTTPS on this port
      serverCertificate: /etc/certs/servercert.pem
      privateKey: /etc/certs/privatekey.pem
  - port:
      number: 9443
      name: https-9443
      protocol: HTTPS
    hosts:
    - "bookinfo-namespace/*.bookinfo.com"
    tls:
      mode: SIMPLE # enables HTTPS on this port
      credentialName: bookinfo-secret # fetches certs from Kubernetes secret
  - port:
      number: 9080
      name: http-wildcard
      protocol: HTTP
    hosts:
    - "*"
  - port:
      number: 2379 # to expose internal service via external port 2379
      name: mongo
      protocol: MONGO
    hosts:
    - "*"

```


### Ingress Gateways
### Egress Gateways

## ServiceEntry CRD资源 
https://istio.io/docs/reference/config/networking/service-entry/
## EnvoyFilter  CRD资源 （不推荐使用）
https://istio.io/docs/reference/config/networking/envoy-filter/

EnvoyFilter提供了一种自定义Istio Pilot生成的Envoy配置的机制。使用EnvoyFilter修改某些字段的值，添加特定的过滤器，甚至添加全新的侦听器，群集等。**必须谨慎使用此功能，因为不正确的配置可能会破坏整个网格的稳定性**。与其他Istio网络对象不同，EnvoyFilters是附加应用的。对于特定名称空间中的给定工作负载，可以存在任意数量的EnvoyFilter。**这些EnvoyFilters的应用顺序如下：配置根名称空间中的所有EnvoyFilters ，然后是工作负载名称空间中所有匹配的EnvoyFilters。**

- 注意1：此API的某些方面与Istio网络子系统以及Envoy的XDS API的内部实现紧密相关。尽管EnvoyFilter API本身将保持向后兼容性，但应在Istio代理版本升级中仔细监视通过此机制提供的任何envoy配置，以确保删除和弃用不适当的字段。

- 注意2：当多个EnvoyFilters绑定到给定名称空间中的相同工作负载时，将按创建时间顺序依次处理所有补丁。如果多个EnvoyFilter配置相互冲突，则该行为是不确定的。

- 注意3：* _要将EnvoyFilter资源应用于系统中的所有工作负载（边车和网关），请在config 根名称空间中定义资源，而不需要工作负载选择器。


## Sidecar CRD资源 
https://istio.io/docs/reference/config/networking/sidecar/


## Virtual services CRD对象 虚拟服务
https://istio.io/docs/reference/config/networking/virtual-service/
https://istio.io/docs/concepts/traffic-management/#virtual-services

### 超时和重试
### 故障注入
配置完网络（包括故障恢复策略）后，可以使用Istio的故障注入机制来测试整个应用程序的故障恢复能力。故障注入是一种将错误引入系统的测试方法，以确保其可以承受并从错误情况中恢复。使用故障注入对确保您的故障恢复策略不是不兼容或限制过于严格可能特别有用，这可能会导致关键服务不可用。

istio仅针对http协议支持两种 方式：
- 延迟故障
 模仿增加的网络延迟或上游服务过载

- 中断故障
 模仿上游服务中的故障。中止通常以HTTP错误代码或TCP连接失败的形式表现出来

https://istio.io/docs/tasks/traffic-management/fault-injection/
https://istio.io/docs/reference/config/networking/virtual-service/#HTTPFaultInjection

#### 解释
Istio故障恢复功能对应用程序是完全透明的。应用程序不知道Envoy Sidecar代理在返回响应之前是否正在处理被调用服务的故障。这意味着，如果您还在应用程序代码中设置故障恢复策略，则需要记住两者都独立工作，因此可能会发生冲突。例如，假设您有两个超时，一个在虚拟服务中配置，另一个在应用程序中配置。应用程序为服务的API调用设置2秒超时。但是，您在虚拟服务中配置了3秒超时和1次重试。在这种情况下，应用程序的超时将首先开始，因此您的Envoy超时和重试尝试无效。

虽然Istio故障恢复功能提高了网格中服务的可靠性和可用性，但应用程序必须处理故障或错误并采取适当的后备操作。例如，当负载平衡池中的所有实例都失败时，Envoy将返回一个HTTP 503 代码。应用程序必须实现处理HTTP 503错误代码所需的任何后备逻辑 。


```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: ratings
spec:
  hosts:
  - ratings
  http:
  - fault: // 故障注入
      delay: // 延迟
        percentage: //  百分比
          value: 0.1 // 0.1% 浅粉之一
        fixedDelay: 5s // 延迟5秒返回结果  此虚拟服务为服务的每1000个请求中的1个引入了5秒的延迟
      abort: // 终止
        httpStatus: 500 // 设置响应码500的中断故障
        percentage:
          value: 50.0 // 为服务设置响应码500的中断故障，期望50%请求响应为httpstatus=500
  - route:
    - destination:
        host: ratings
        subset: v1
    timeout: 10s // 超时
    retries: // 重试
      attempts: 3 // 3次重试
      perTryTimeout: 2s // 每个重试都有2秒的超时
```

## Destination rules CRD对象 目标规则
### 负载均衡器
### 控制到上游服务的连接量
### 异常值检测
### TLS
### 特定于单个端口的流量策略


https://istio.io/docs/concepts/traffic-management/#destination-rules
https://istio.io/docs/reference/config/networking/destination-rule/

[断路器 circuit-breakers](https://istio.io/docs/tasks/traffic-management/circuit-breaking/)

https://istio.io/docs/reference/config/networking/destination-rule/#TrafficPolicy


Field |		Description |	Required
---|---|---
loadBalancer |   控制负载均衡器算法的设置。|No
connectionPool |设置控制到上游服务的连接量 |No
outlierDetection |控制从负载平衡池中清除不健康主机的设置|No
tls | 用于连接到上游服务的TLS相关设置。|No
portLevelSettings|特定于单个端口的流量策略。 请注意，端口级别设置将覆盖目标级别设置。 在目标级别指定的流量设置被端口级别设置覆盖时，将不会继承，即默认值将应用于端口级别流量策略中省略的字段。|No

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: reviews
spec:
  host: reviews
  loadBalancer: // 负载 https://istio.io/docs/reference/config/networking/destination-rule/#LoadBalancerSettings
     simple: LEAST_CONN
     consistentHash: // 使用cookie hash负载
         httpCookie:
           name: user
           ttl: 0s
  subsets: // 子集
  - name: v1 // 带有标签v1的目标
    labels:
      version: v1
    trafficPolicy: // 策略
      portLevelSettings: // 不同端口指定不同负载
      - port:
          number: 80
        loadBalancer:
          simple: LEAST_CONN
      - port:
          number: 9080
        loadBalancer:
          simple: ROUND_ROBIN
      loadBalancer: // 负载
        simple: ROUND_ROBIN
      connectionPool:
        tcp: // https://istio.io/docs/reference/config/networking/destination-rule/#ConnectionPoolSettings-TCPSettings
          maxConnections: 100 // 负载的并发连接数限制为100
          connectTimeout: 30ms
          tcpKeepalive:
            time: 7200s
            interval: 75s
        http: // https://istio.io/docs/reference/config/networking/destination-rule/#ConnectionPoolSettings-HTTPSettings
          http1MaxPendingRequests: 1
          maxRequestsPerConnection: 1
        outlierDetection: // 异常值检测 https://istio.io/docs/reference/config/networking/destination-rule/#OutlierDetection
          consecutiveErrors: 2 // 任何连续2次失败的，带有502、503或504错误代码 默认情况下或设置为0时，此功能被禁用。
          interval: 1s // 每1秒进行一次扫描 默认值为10秒
          baseEjectionTime: 3m // 任何连续2次失败的 拒绝3分钟 默认值为30秒
          maxEjectionPercent: 100 // 100%拒绝 默认为10％
```


# Istio 其它介绍

## 可扩展性 WebAssembly
https://istio.io/docs/concepts/wasm/
WebAssembly是一种沙盒技术，可用于扩展Istio代理（Envoy）。Proxy-Wasm沙箱API取代了Mixer，成为Istio中的主要扩展机制。Istio 1.6将为Proxy-Wasm插件提供统一的配置API。

WebAssembly沙箱目标：

- 效率 - 扩展增加了低延迟，CPU和内存开销。
- 功能 - 扩展可以执行策略，收集遥测和执行有效载荷突变。
- 隔离 - 一个插件中的编程错误或崩溃确实会影响其他插件。
- 配置 - 使用与其他Istio API一致的API配置插件。扩展名可以动态配置。
- 运营商 - 可以扩展扩展并将其部署为仅日志，失败打开或失败关闭。
- 扩展开发人员 - 该插件可以用几种编程语言编写。


Istio扩展（Proxy-Wasm插件）具有几个组件：

- 筛选器服务提供程序接口（SPI），用于为筛选器构建Proxy-Wasm插件。
Filter Service Provider Interface (SPI) for building Proxy-Wasm plugins for filters.

- Envoy中嵌入了Sandbox V8 Wasm Runtime。
Sandbox V8 Wasm Runtime embedded in Envoy.

- 头，尾和元数据的主机API。
Host APIs for headers, trailers and metadata.

- 调出用于gRPC和HTTP调用的API。
Call out APIs for gRPC and HTTP calls.

- Stats和Logging API，用于度量和监视。
Stats and Logging APIs for metrics and monitoring.

![](../img/k8s/extending.svg)

[WebAssembly for Proxies (ABI specification)](https://github.com/proxy-wasm/spec)
[WebAssembly Hub](https://docs.solo.io/web-assembly-hub/latest/tutorial_code/)

## 可观察性(指标，调用链，日志)
https://istio.io/docs/concepts/observability/

Istio根据监视的四个“黄金信号”（延迟，流量，错误和饱和度）生成一组服务指标

Istio指标收集从Sidecar代理（Envoy）开始。每个代理都会生成一组丰富的指标，用于衡量通过代理的所有流量（入站和出站）。代理还提供有关代理本身的管理功能的详细统计信息，包括配置和运行状况信息。

Envoy生成的指标以Envoy资源（例如侦听器和集群）的粒度提供对网格的监视。因此，需要了解网状服务和Envoy资源之间的连接才能监视Envoy指标。

Istio使操作员可以选择在每个工作负载实例上生成和收集哪个Envoy指标。默认情况下，Istio仅启用Envoy生成的统计信息的一小部分，以避免过多的指标后端并减少与指标收集相关的CPU开销。但是，运营商可以在需要时轻松扩展收集的代理指标集。这样就可以针对性地调试网络行为，同时降低跨网状网监视的总成本。
https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/statistics.html?highlight=statistics


[Visualizing Metrics with Grafana](https://istio.io/docs/tasks/observability/metrics/using-istio-dashboard/)



### proxy-level-metrics 代理级指标
[Debugging Envoy and Istiod](https://istio.io/docs/ops/diagnostic-tools/proxy-cmd/)

### service-level-metrics 服务级指标
[Default Metrics](https://istio.io/docs/reference/config/policy-and-telemetry/metrics/)
[prometheus](https://istio.io/docs/reference/config/policy-and-telemetry/adapters/prometheus/)
默认情况下，标准Istio指标导出到Prometheus。

#### HTTP, HTTP/2, and GRPC
- Request Count (istio_requests_total): This is a COUNTER incremented for every request handled by an Istio proxy.

- Request Duration (istio_request_duration_seconds): This is a DISTRIBUTION which measures the duration of requests.

- Request Size (istio_request_bytes): This is a DISTRIBUTION which measures HTTP request body sizes.

- Response Size (istio_response_bytes): This is a DISTRIBUTION which measures HTTP response body sizes.


#### TCP 
- Tcp Byte Sent (istio_tcp_sent_bytes_total): This is a COUNTER which measures the size of total bytes sent during response in case of a TCP connection.

- Tcp Byte Received (istio_tcp_received_bytes_total): This is a COUNTER which measures the size of total bytes received during request in case of a TCP connection.

- Tcp Connections Opened (istio_tcp_connections_opened_total): This is a COUNTER incremented for every opened connection.

- Tcp Connections Closed (istio_tcp_connections_closed_total): This is a COUNTER incremented for every closed connection.

### control-plane-metrics 控制平面指标
https://istio.io/docs/concepts/observability/#control-plane-metrics

#### pilot-discovery 飞行员
https://istio.io/docs/reference/commands/pilot-discovery/#metrics
#### galley
https://istio.io/docs/reference/commands/galley/#metrics
#### mixs 混合器
https://istio.io/docs/reference/commands/mixs/#metrics
#### istio_ca 堡垒
https://istio.io/docs/reference/commands/istio_ca/#metrics


### Tracing
Istio支持许多跟踪后端，包括Zipkin， Jaeger，LightStep和 Datadog
https://istio.io/docs/tasks/observability/distributed-tracing/jaeger/
https://istio.io/docs/tasks/observability/distributed-tracing/zipkin/

[Distributed Tracing FAQ](https://istio.io/faq/distributed-tracing/)

### 访问日志
Istio可以以一组可配置的格式生成服务流量的访问日志，从而使操作员可以完全控制日志的方式，内容，时间和地点。Istio向访问日志记录机制公开了全套的源和目标元数据，从而可以对网络事务进行详细的审核。

示例Istio访问日志（格式为JSON）：

```
{"level":"info","time":"2019-06-11T20:57:35.424310Z","instance":"accesslog.instance.istio-control","connection_security_policy":"mutual_tls","destinationApp":"productpage","destinationIp":"10.44.2.15","destinationName":"productpage-v1-6db7564db8-pvsnd","destinationNamespace":"default","destinationOwner":"kubernetes://apis/apps/v1/namespaces/default/deployments/productpage-v1","destinationPrincipal":"cluster.local/ns/default/sa/default","destinationServiceHost":"productpage.default.svc.cluster.local","destinationWorkload":"productpage-v1","httpAuthority":"35.202.6.119","latency":"35.076236ms","method":"GET","protocol":"http","receivedBytes":917,"referer":"","reporter":"destination","requestId":"e3f7cffb-5642-434d-ae75-233a05b06158","requestSize":0,"requestedServerName":"outbound_.9080_._.productpage.default.svc.cluster.local","responseCode":200,"responseFlags":"-","responseSize":4183,"responseTimestamp":"2019-06-11T20:57:35.459150Z","sentBytes":4328,"sourceApp":"istio-ingressgateway","sourceIp":"10.44.0.8","sourceName":"ingressgateway-7748774cbf-bvf4j","sourceNamespace":"istio-control","sourceOwner":"kubernetes://apis/apps/v1/namespaces/istio-control/deployments/ingressgateway","sourcePrincipal":"cluster.local/ns/istio-control/sa/default","sourceWorkload":"ingressgateway","url":"/productpage","userAgent":"curl/7.54.0","xForwardedFor":"10.128.0.35"}
```


## security 
https://istio.io/docs/reference/config/security/
### Request Authentication CRD对象
### AuthorizationPolicy CRD对象
https://istio.io/docs/reference/config/security/authorization-policy/
### Peer Authentication CRD对象 对等身份验证
https://istio.io/docs/concepts/security/#peer-authentication

