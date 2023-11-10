---
Title: Caddy
Summary: Caddy
Tags:
    - Caddy
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
将cloudflare_key替换为自己的key， 从[这里](https://dash.cloudflare.com/profile/api-tokens)创建一个dns api令牌即可
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
    order markdown before file_server
    order forward_proxy before markdown
}
www.wcoder.com {
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
    -v /data/dockerv/caddy/srv:/srv \
    -v /data/dockerv/caddy/data:/data \
    -v /data/dockerv/caddy/log:/log \
    -v /data/dockerv/caddy/config:/config \
    -v /data/dockerv/caddy/Caddyfile:/etc/caddy/Caddyfile \
    --name caddy caddy-markdown:v0.0.1

caddy-markdown:v0.0.1
```


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
proxy.wcoder.com {
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
curl -I --proxy https://xx:xx@proxy.wcoder.com  https://www.baidu.com


curl -x https://proxy.wcoder.com --proxy-user x:x -L https://www.baidu.com
```
