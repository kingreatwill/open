---
title: Caddy
summary: Caddy
tags:
    - Caddy
    - Traefik
---

## Caddy

Caddy有下面这些开箱即用的特性:

- 全自动支持HTTP/2协议，无需任何配置。
- Caddy 使用 Let’s Encrypt 让你的站点全自动变成HTTPS，无需任何配置。
- 合理使用多核多核 得益于go的特性
- 完全支持IPv6环境
- Caddy 对WebSockets有很好的支持
- 自动把Markdown转成 HTML
- Caddy 对log格式的定义很容易
- 易于部署 得益于go的特性，caddy只是一个小小的二进制文件，没有依赖，很好部署
- 得益于Go的跨平台特性，Caddy很容易的支持了三大主流系统：Windows、 Linux、Mac

> “几乎所有的功能在Caddy里的都是插件，HTTP服务器是插件，高级的TLS特性也是插件，每一行命令实现的功能都是一个插件”

caddy源码: https://github.com/caddyserver/caddy

caddy 插件: https://caddyserver.com/docs/extending-caddy

caddy 辅助编译工具(xcaddy): https://github.com/caddyserver/xcaddy

> watch参数可以动态加载配置文件

### NGINX Config Adapter
https://github.com/caddyserver/nginx-adapter

`caddy run|start --config nginx.conf --adapter nginx`

### 其它代理
#### Traefik
https://github.com/traefik/traefik 45.5k golang

https://plugins.traefik.io/plugins


#### Envoy
https://github.com/envoyproxy/envoy 23k c++

虽然是C++但是可以使用Envoy WASM插件来进行扩展, 当前也是支持lua的

#### YARP reverse-proxy
[YARP reverse-proxy](https://github.com/microsoft/reverse-proxy) 7.1k C#

### 编译

我们可以使用源码编译:
```
git clone git@github.com:caddyserver/caddy.git

cd caddy/cmd/caddy
go build
```
> 如果自己编译, 当需要插件的时候,需要在caddy/cmd/caddy/main.go中import _ "github.com/example/mymodule"

也可以使用xcaddy编译:
```
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
xcaddy build

#也可以指定版本
xcaddy build v2.0.1
xcaddy build master
xcaddy build a58f240d3ecbb59285303746406cab50217f8d24

#指定插件
xcaddy build v2.0.1 --with github.com/caddyserver/ntlm-transport@v0.1.1
```
xcaddy会自动下载源码和插件源码进行编译

#### 源码Dockerfile编译 
Dockerfile
```
FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /app/caddy ./cmd/caddy
RUN ./cmd/caddy/setcap.sh cap_net_bind_service=+ep /app/caddy; \
	chmod +x /app/caddy;


FROM caddy:2.11-alpine
LABEL maintainer="wcoder <350840291@qq.com>"
COPY --from=builder /app/caddy /usr/bin/caddy

# docker build -t caddy .
# docker tag caddy kingreatwill/caddy:v2.11.0
# docker login -u kingreatwill
# docker push kingreatwill/caddy:v2.11.0
# docker run -d --cap-add=NET_ADMIN --restart=always --network host     -v /data/dockerv/caddy/srv:/srv     -v /data/dockerv/caddy/data:/data     -v /data/dockerv/caddy/log:/log     -v /data/dockerv/caddy/config:/config     -v /data/dockerv/caddy/Caddyfile:/etc/caddy/Caddyfile     --name caddy kingreatwill/caddy:v2.11.0
# docker run -d --restart=always --network host     -v /data/dockerv/caddy/srv:/srv     -v /data/dockerv/caddy/data:/data     -v /data/dockerv/caddy/log:/log     -v /data/dockerv/caddy/config:/config     -v /data/dockerv/caddy/Caddyfile:/etc/caddy/Caddyfile     --name caddy kingreatwill/caddy:v2.11.0
```
Caddyfile
```
www.wcoder.com wcoder.com gantt.wang dotnet.wang weidian.faith wcoder.is-a.dev {
    root * /srv/www
    file_server browse {
        hide .git
        index index.html
    }
    log {
        output file /log/access.log
    }
}

record.ren shiori.wcoder.com {
    reverse_proxy 127.0.0.1:10001
    log {
        output file /log/record.ren.log
    }
}

frp.wcoder.com {
    reverse_proxy 127.0.0.1:7500
    log {
        output file /log/frp.wcoder.com.log
    }
}

http://*.frp.wcoder.com {
    reverse_proxy 127.0.0.1:7080
    log {
        output file /log/http.frp.wcoder.com.log
    }
}

https://*.frp.wcoder.com {
    reverse_proxy https://127.0.0.1:7443 {
        transport http {
            tls_insecure_skip_verify
            tls_server_name {host}
        }
    }
    log {
        output file /log/https.frp.wcoder.com.log
    }
}

localapi.wcoder.com {
    reverse_proxy 127.0.0.1:7080
    log {
        output file /log/localapi.wcoder.com.log
    }
}

wxapi.wcoder.com {
    reverse_proxy 127.0.0.1:10005
    log {
        output file /log/wxapi.wcoder.com.log
    }
}
```

### 在Go程序中中嵌入Caddy

### 占位符 placeholders(variables)
占位符在不同的上下文有不同的占位符

[全局占位符](https://caddyserver.com/docs/conventions#placeholders)

[http.*占位符](https://caddyserver.com/docs/json/apps/http/#docs)

[简写占位符](https://caddyserver.com/docs/caddyfile/concepts#placeholders)

[自定义变量](https://caddyserver.com/docs/caddyfile/directives/vars#vars)

### 泛域名证书申请(Wildcard certificates)
```
www.mydomain.com {
        reverse_proxy localhost:16325
        tls my@qq.com
}
```
并且自动通过let's encrypt申请ssl证书，申请邮箱为my@qq.com

**泛域名证书只能通过dns记录来验证，所以需要配置dns提供商的信息**
对于.cf, .ga, .gq, .ml,  .tk后缀的域名基本就告别自动泛域名证书了，cloudflare不允许通过api调用修改dns，而国内不允许此类域名备案。

安装对应域名提供商的dns模块，常用的国内有阿里云(alidns)和腾讯云(dnspod)， 国外有cloudflare 微软(azure) 谷歌(googleclouddns)

可以在下载caddy时勾选对应模块。如果没有勾选也没关系，可以通过命令行下载

cloudflare: `xcaddy build --with github.com/caddy-dns/cloudflare`
阿里云: `xcaddy build --with github.com/caddy-dns/alidns`
腾讯云: `xcaddy build --with github.com/caddy-dns/dnspod`


安装完毕后修改caddyfile, 不同dns提供商的模块配置略微有些区别


> [caddy配置反向代理和ssl证书申请](https://www.cnblogs.com/turingguo/p/caddy.html)

**也可以放到全局**
```
{
    acme_dns cloudflare 1484053787dJQB8vP1q0yc5ZEBnH6JGS4d3mBmvIeMrnnxFi3WtJdF
}
```

#### cloudflare   
将cloudflare_key替换为自己的key， 从[这里](https://dash.cloudflare.com/profile/api-tokens)创建一个dns api令牌即可(Edit zone DNS)
```
*.my.com {
    tls {       dns cloudflare cloudflare_apikey    }
}
```
#### 阿里云
key_id和key_secret来自控制台创建的accesskey,从[这里](https://ram.console.aliyun.com/manage/ak)可以创建，region就是账户所属区域,可不填，默认为 zh-hangzhou
```
*.my.com {
    tls { 
        dns alidns {
            access_key_id        key_id
            access_key_secret    key_secret
        }
    }
}
```
#### 腾讯云
dnspod需要的是[api token](https://console.dnspod.cn/account/token/token)，由 ID,Token 组合而成的，用英文的逗号分割， [这里](https://docs.dnspod.cn/account/5f2d466de8320f1a740d9ff3/)有说明如何创建Token.  如id为3245，token为sf3fwr234,则完整的api_token为 3245,sf3fwr234

```
*.my.com {
    tls { 
        dns dnspod "api_token"
    }
}
```





### docker
https://github.com/kingreatwill/caddy-modules

Caddyfile
```
{
    order markdown before file_server
}
:2019 {
    root * /srv
    encode gzip
    markdown {
        template /markdown.tmpl
    }
    file_server browse {
        hide .git
        index README.md index.html index.htm
    }
}
```

```
docker run -d -p 2019:2019 \
    -v /data/dockerv/caddy/srv:/srv \
    -v /data/dockerv/caddy/data:/data \
    -v /data/dockerv/caddy/config:/config \
    -v /data/dockerv/caddy/Caddyfile:/etc/caddy/Caddyfile \
    -v /data/dockerv/caddy/index.html:/usr/share/caddy/index.html \
    kingreatwill/caddy:v2.7.5
```

```
{
    admin 10.0.0.11:2019
    order sentry first
    order markdown before file_server
    order forward_proxy before markdown
}
www.wcoder.com wcoder.com gantt.wang dotnet.wang weidian.faith {
    root * /srv/www
    file_server browse {
        hide .git
        index index.html
    }
    log {
        output file /log/access.log
    }
}
note.wcoder.com open.wcoder.com record.ren {
    root * /srv/note.wcoder.com
    encode gzip
    markdown {
        template /markdown.tmpl
    }
    file_server browse {
        hide .git
        index README.md index.html index.htm
    }
    log {
        output file /log/note.wcoder.com.log
    }
}
frp.wcoder.com {
    reverse_proxy 127.0.0.1:7500
    log {
        output file /log/frp.wcoder.com.log
    }
}
http://*.frp.wcoder.com {
    reverse_proxy 127.0.0.1:7080
    log {
        output file /log/http.frp.wcoder.com.log
    }
}

https://*.frp.wcoder.com {
    tls { 
        dns dnspod "xx,xx"
    }
    reverse_proxy https://127.0.0.1:7443 {
        transport http {
            tls_insecure_skip_verify
            tls_server_name {host}
        }
    }
    log {
        output file /log/https.frp.wcoder.com.log
    }
}


proxy.wcoder.com {
    forward_proxy {
        basic_auth xx xxx
        ports     80 443
        hide_ip
        hide_via
    }
    log {
        output file /log/proxy.wcoder.com.log
    }
}

docker run -d --cap-add=NET_ADMIN --restart=always --network host \
    -e CADDY_ADMIN="10.0.0.11:2019" \
    -e OTEL_SERVICE_NAME=caddy \
    -e OTEL_EXPORTER_OTLP_TRACES_ENDPOINT="http://43.155.152.66:4317" \
    -e OTEL_TRACES_SAMPLER="always_on" \
    -e OTEL_EXPORTER_OTLP_INSECURE=true \
    -e SENTRY_DSN="https://5851bf0c021cfb2f29e370483922200c@o4506274015936512.ingest.sentry.io/4506274021507072" \
    -v /data/dockerv/caddy/srv:/srv \
    -v /data/dockerv/caddy/data:/data \
    -v /data/dockerv/caddy/log:/log \
    -v /data/dockerv/caddy/config:/config \
    -v /data/dockerv/caddy/Caddyfile:/etc/caddy/Caddyfile \
    --name caddy caddy-markdown:v0.0.1

-- 开启Sentry后流量飙升

docker run -d --cap-add=NET_ADMIN --restart=always --network host \
    -e OTEL_SERVICE_NAME=caddy \
    -e OTEL_EXPORTER_OTLP_TRACES_ENDPOINT="http://43.155.152.66:4317" \
    -e OTEL_TRACES_SAMPLER="always_on" \
    -e OTEL_EXPORTER_OTLP_INSECURE=true \
    -v /data/dockerv/caddy/srv:/srv \
    -v /data/dockerv/caddy/data:/data \
    -v /data/dockerv/caddy/log:/log \
    -v /data/dockerv/caddy/config:/config \
    -v /data/dockerv/caddy/Caddyfile:/etc/caddy/Caddyfile \
    --name caddy caddy-markdown:v0.0.1

```
`-e SENTRY_DEBUG="true" \`

禁用https
```
{▾
	"disable": false,
	"disable_redirects": false,
	"disable_certificates": false,
	"skip": [""],
	"skip_certificates": [""],
	"ignore_loaded_certificates": false
}
```

```
{
    auto_https disable_redirects
}
```

### forwardproxy
利用Caddy的插件forwardproxy快速搭建HTTP/HTTPS正向代理
https://github.com/caddyserver/forwardproxy/tree/caddy2


`xcaddy build --with github.com/caddyserver/forwardproxy@caddy2`

```
{
	order forward_proxy before file_server
}
:8888 {
    forward_proxy {
        basic_auth xx xx
        ports     80 443
        hide_ip
        hide_via
    }
    log {
        output file /log/proxy.wcoder.com.log
    }
}
```

设置代理
```
export http_proxy=http://user123:hahapwd@47.113.67.125:18888
export https_proxy=http://user123:hahapwd@47.113.67.125:18888
```
测试
```
curl -I --proxy https://xx:xx@proxy.wcoder.com  https://www.cnblogs.com/

curl -Lv --proxy http://xx:xx@43.155.152.66:8888  http://www.cnblogs.com/

curl -x https://proxy.wcoder.com --proxy-user x:x -L https://www.baidu.com
```

### tracing opentelemetry(只支持grpc)

https://caddyserver.com/docs/caddyfile/directives/tracing#tracing

```
export OTEL_EXPORTER_OTLP_HEADERS="myAuthHeader=myToken,anotherHeader=value"
export OTEL_SERVICE_NAME=caddy
export OTEL_EXPORTER_OTLP_TRACES_ENDPOINT="http://43.155.152.66:4317"
export OTEL_TRACES_SAMPLER="always_on"
export OTEL_EXPORTER_OTLP_INSECURE=true

handle /example* {
	tracing {
		span {path}
	}
	reverse_proxy 127.0.0.1:8081
}
```

frp内网穿透需要使用tcp来代理grpc的端口

```
# https://opentelemetry.io/docs/concepts/sdk-configuration/otlp-exporter-configuration/
gRPC: export OTEL_EXPORTER_OTLP_TRACES_ENDPOINT="https://my-api-endpoint:443"
HTTP: export OTEL_EXPORTER_OTLP_TRACES_ENDPOINT="http://my-api-endpoint/v1/traces"
```

日志字段`traceID`

### Dynamic upstreams

[Dynamic upstreams](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy#dynamic-upstreams)

[什么是 DNS SRV 记录？](https://www.cloudflare.com/zh-cn/learning/dns/dns-records/dns-srv-record/)
SRV记录（英语：Service Record，中文又名服务定位记录）是域名系统中用于指定服务器提供服务的位置（如主机名和端口）数据。此数据于RFC 2782中定义，类型代码为33。部分协议，如会话发起协议（SIP）及可扩展消息与存在协议（XMPP）通常需要服务记录的支持。

#### SRV DNS records
consul 是支持SRV DNS的, 也就是我们可以使用consul来动态获取上游地址

> 标准的 DNS 默认端口为 53, 而 consul 默认的 DNS 端口为 8600.
> `dig @192.168.1.11 -p 8600 consul.service.consul SRV`
> 标准的 consul 服务写法(最常用): `[tag.]<your_service>.service[.your_datacenter].consul`

> RFC2782 的写法是: `_<service>._<protocol>.service[.datacenter][.domain]` , 知道有这么一种方法就行

参考[Consul 域名服务](https://www.cnblogs.com/harrychinese/p/consul_dns.html)