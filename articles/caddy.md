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

docker run -d --cap-add=NET_ADMIN --restart=always --network host \
    -v /data/dockerv/caddy/srv:/srv \
    -v /data/dockerv/caddy/data:/data \
    -v /data/dockerv/caddy/log:/log \
    -v /data/dockerv/caddy/config:/config \
    -v /data/dockerv/caddy/Caddyfile:/etc/caddy/Caddyfile \
    --name caddy caddy-markdown:v0.0.1

caddy-markdown:v0.0.1
```