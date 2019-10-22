# docker安装kong with postgres
[官网](https://konghq.com/kong)
[教程](https://docs.konghq.com/install/docker/)
[docker hub](https://hub.docker.com/_/kong/)

# docker安装kong DB-less

# 一个典型的 Nginx 配置
```
upstream helloUpstream {
    server localhost:3000 weight=100;
}

server {
    listen  80;
    location /hello {
        proxy_pass http://helloUpstream;
    }
}
```
如上这个简单的 Nginx 配置，便可以转换为如下的 Http 请求。

# 对应的 Kong 配置

### 配置 upstream
```
curl -X POST http://localhost:8001/upstreams --data "name=helloUpstream"
```
### 配置 service
```
curl -X POST http://localhost:8001/services --data "name=hello" --data "host=helloUpstream"
```
### 配置 target
```
curl -X POST http://localhost:8001/upstreams/hello/targets --data "target=localhost:3000" --data "weight=100"
```
### 配置 route
```
curl -X POST http://localhost:8001/routes --data "paths[]=/hello" --data "service.id=8695cc65-16c1-43b1-95a1-5d30d0a50409"
```
这一切都是动态的，无需手动 reload nginx.conf。

我们为 Kong 新增路由信息时涉及到了 upstream，target，service，route 等概念，他们便是 Kong 最最核心的四个对象。（你可能在其他 Kong 的文章中见到了 api 这个对象，在最新版本 0.13 中已经被弃用，api 已经由 service 和 route 替代）

从上面的配置以及他们的字面含义大概能够推测出他们的职责，upstream 是对上游服务器的抽象；target 代表了一个物理服务，是 ip + port 的抽象；service 是抽象层面的服务，他可以直接映射到一个物理服务(host 指向 ip + port)，也可以指向一个 upstream 来做到负载均衡；route 是路由的抽象，他负责将实际的 request 映射到 service。

他们的关系如下

upstream 和 target ：1 对 n

service 和 upstream ：1 对 1 或 1 对 0 （service 也可以直接指向具体的 target，相当于不做负载均衡）

service 和 route：1 对 n

# 高可扩展性的背后—插件机制
Kong 的另一大特色便是其插件机制，这也是我认为的 Kong 最优雅的一个设计。

文章开始时我们便提到一点，微服务架构中，网关应当承担所有服务共同需要的那部分功能，这一节我们便来介绍下，Kong 如何添加 jwt 插件，限流插件。

插件（Plugins）装在哪儿？对于部分插件，可能是全局的，影响范围是整个 Kong 服务；大多数插件都是装在 service 或者 route 之上。这使得插件的影响范围非常灵活，我们可能只需要对核心接口进行限流控制，只需要对部分接口进行权限控制，这时候，对特定的 service 和 route 进行定向的配置即可。

为 hello 服务添加50次/秒的限流
```
curl -X POST http://localhost:8001/services/hello/plugins \
--data "name=rate-limiting" \
--data "config.second=50"
```
为 hello 服务添加 jwt 插件
```
curl -X POST http://localhost:8001/services/login/plugins \
--data "name=jwt"
```
同理，插件也可以安装在 route 之上
```
curl -X POST http://localhost:8001/routes/{routeId}/plugins \
--data "name=rate-limiting" \
--data "config.second=50"

curl -X POST http://localhost:8001/routes/{routeId}/plugins \
--data "name=jwt"
```