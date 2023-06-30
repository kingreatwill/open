[TOC]

> 很多CDN 行业使用OpenResty，多是受到 cloudflare 技术栈的影响，OpenResty 的作者也在国外这家 CDN 公司。

> 关于作者 章亦春 先后在中国雅虎、淘宝、Cloudflare 就职, 是开源 OpenResty ® 项目创始人兼 OpenResty Inc. 公司 CEO 和创始人。 章亦春（Github ID: agentzh），生于中国江苏，现定居美国湾区。

> OpenResty 最早是雅虎中国的一个公司项目，起步于 2007 年 10 月。当时兴起了 OpenAPI 的热潮，用于满足各种 Web Service 的需求，就诞生了 OpenResty。在公司领导的支持下，最早的 OpenResty 实现从一开始就开源了。最初的定位是服务于公司外的开发者，像其他的 OpenAPI 那样，但后来越来越多地是为雅虎中国的搜索产品提供内部服务。

> 章亦春在加入淘宝数据部门的量子团队之后，决定对 OpenResty 进行重新设计和彻底重写，并把应用重点放在支持像量子统计这样的 web 产品上面，所以量子统计 3.0 开始也几乎完全是 Web Service 驱动的纯 AJAX 应用。这是第二代的 OpenResty，一般称之为 ngx_openresty，以便和第一代基于 Perl 和 Haskell 实现的 OpenResty 加以区别。章亦春和他的同事王晓哲一起设计了第二代的 OpenResty。在王晓哲的提议下，选择基于 Nginx 和 Lua 进行开发。

> 2007年大学毕业，加入中国雅虎搜索技术部。当时的中国雅虎已经是阿里旗下的一家公司了。2009 年被调到淘宝数据产品与平台部，集中力量和同事们一起基于 OpenResty 开发淘宝量子统计这款产品，主要面向淘宝卖家。2011 年离开阿里，在家里全心从事 OpenResty 等开源项目的工作。 2012 年受 Cloudflare 公司邀请，举家来到美国湾区，全职从事 OpenResty 开源开发，以及基于 OpenResty 的全球 CDN 网络的软件系统的研发和完善。2016 年离开了 Cloudflare，并于 2017 年初在美国创办了 OpenResty Inc. 公司。


OpenResty 在 1.5.8.1 版之后才默认使用 LuaJIT，之前使用的是标准 Lua 解析器。
## 安装

http://openresty.org/cn/download.html

### docker
复制配置文件到宿主机：
```
$ mkdir -p /data/openresty/conf/conf.d
$ docker run -d --name openresty-example openresty/openresty:1.13.6.2-2-xenial sleep 1234
$ docker cp openresty-example:/usr/local/openresty/nginx/conf/nginx.conf /data/openresty/conf/
$ docker cp openresty-example:/etc/nginx/conf.d/. /data/openresty/conf/conf.d/

$ docker rm -i openresty-example
```
启动
```
$ docker run -d --restart always -e TZ=Asia/Shanghai --name openresty-gateway  -p 80:80 -v /data/openresty/conf/conf.d:/etc/nginx/conf.d -v /data/openresty/conf/nginx.conf:/usr/local/openresty/nginx/conf/nginx.conf openresty/openresty:1.13.6.2-2-xenial

```

### Ubuntu
添加 gpg：
```
$ wget -qO - https://openresty.org/package/pubkey.gpg | sudo apt-key add -
```
安装：add-apt-repository
```
$ apt install -y software-properties-common
```
添加官方源：
```
$ add-apt-repository -y "deb http://openresty.org/package/ubuntu $(lsb_release -sc) main"
```
更新源：
```
$ apt update
```
安装 OpenRestry
```
$ apt install -y openresty
```
默认会安装到 `/usr/local/openresty` 下。

### windows
1. 下载后双击nginx.exe运行 
2. 验证：
`tasklist /fi "imagename eq nginx.exe"`
http://127.0.0.1:80/


3. 自定义lua
在\conf\nginx.conf文件修改
```
location /hello {
	#修改为直接返回text模式，而不是返回文件。默认配置在极速模式下得浏览器会形成下载文件
	default_type text/html;
	#关闭缓存模式，每次都读取lua文件，不使用已经缓存的lua文件（修改nginx.conf文件还是要重启的）
	lua_code_cache off;
	#在返回节点加载lua文件（相对路径即可）
	content_by_lua_file lua/hello.lua;
}
```
> OpenResty 提供一个专门指令 “content_by_lua_block”，可以在配置文件里书写 Lua 代码，产生响应内容：
```
content_by_lua_block {                              -- 我们的第一个 OpenResty 应用
    ngx.print("hello, world")                       -- 打印经典的 "hello world"
}   
```



添加一下脚本\lua\hello.lua
```
ngx.say('hello world!!!')
```

> 也可以在lualib目录下新增文件夹来写脚本

nginx -s reload
nginx -s stop

- stop：强制立刻停止服务，未完成的请求会被直接关闭；
- quit：停止服务，但必须在处理完当前所有请求之后；
- reload：重启服务，重新加载配置文件和 Lua 代码，服务不会中断；
- reopen：只重新打开日志文件，服务不会中断，常用于切分日志（retate）。



linux写法指定文件
```
nginx -p `pwd`/ -c conf/nginx.conf
```

以下参数没有验证
```
$ bin/openresty -t                                  # 检查默认的配置文件
$ bin/openresty -T                                  # 检查默认的配置文件并打印输出
$ bin/openresty -v                                  # 显示摘要的版本信息
$ bin/openresty -V                                  # 显示完全的版本信息
```

### 处理阶段

OpenResty 提供了一些 “xxx_by_lua” 指令，开发 Web 应用时使用它们就可以在这些阶段里插入 Lua 代码，执行业务逻辑，常用的有：

- init_by_lua：master-initing 阶段，初始化全局配置或模块；
- init_worker_by_lua：worker-initing 阶段，初始化进程专用功能；
- ssl_certificate_by_lua：ssl 阶段，在 “握手” 时设置安全证书；
- set_by_lua：rewrite 阶段，改写 Nginx 变量。
- rewrite_by_lua：rewrite 阶段，改写 URI，实现跳转/重定向。
- access_by_lua：access 阶段，访问控制或限速；
- content_by_lua：content 阶段，产生响应内容；
- balancer_by_lua：content 阶段，反向代理时选择后端服务器；
- header_filter_by_lua：filter 阶段，加工处理响应头；
- body_filter_by_lua：filter 阶段，加工处理响应体；
- log_by_lua：log 阶段，记录日志或其他的收尾工作。

例子如下：
```
init_worker_by_lua_block {                              -- worker-initing 阶段
    ...                                                 -- 启动定时器，定时从 Redis 里获取数据
}
 
 
rewrite_by_lua_block {                                  -- rewrite 阶段，通常是检查、改写 URI
    ...                                                 -- 但也可以操作响应体，做编码解码工作
}
 
 
access_by_lua_block {                                   -- content 阶段，Lua 产生响应内容
    ...                                                 -- 主要的业务逻辑，例如 IP 地址、访问次数
}
 
 
content_by_lua_block {                                  -- content 阶段，Lua 产生响应内容
    ...                                                 -- 主要的业务逻辑，产生向客户端输出的内容
}
 
 
body_filter_by_lua_block {                              -- filter 阶段，加工处理响应数据
    ...                                                 -- 可以对数据编码、加密或者附加额外数据
}
 
 
log_by_lua_block {                                      -- log 阶段，请求结束后的收尾工作
    ...                                                 -- 可以向某个后端发送处理完毕的 "回执"
}
```

content_by_lua_file的例子
```
localtion ~ ^/(\w+) {                                   # 使用正则表达式定义 location
    content_by_lua_file service/http/$1.lua;            # 使用 URI 作为 Lua 文件名
}
```
当执行 ”curl 127.0.0.1/xxx“ 时，OpenResty 就会运行 service/http 目录里对应的 Lua 程序。 


![](https://wiki.shileizcc.com/confluence/download/attachments/47415936/image2018-11-28_12-36-1.png?version=1&modificationDate=1543379761000&api=v2)

![](https://cloud.githubusercontent.com/assets/2137369/15272097/77d1c09e-1a37-11e6-97ef-d9767035fc3e.png)

[参考](https://wiki.shileizcc.com/confluence/pages/viewpage.action?pageId=47415936)

[Directives指令](https://www.nginx.com/resources/wiki/modules/lua/#directives)

[Directives指令](https://github.com/openresty/lua-nginx-module#directives)

### demo
backend.lua
```lua
local ip_list_len = "255"
local ip_header = "172.16.1"
local a = {}
 
# local SingleIP = "218.240.137.68"
local isSingleIP = false
 
if not isSingleIP then
  for i=1,tonumber(ip_list_len) do
    a[i] = ip_header.."."..i
  end
 
  for i,v in ipairs(a) do
    if ngx.var.remote_addr == v then
      ngx.var.group = "backend_02"
    end
  end
end
```
局域网替换上游地址：默认的上游为 backend_default，当判断为真时则替换上游为 backend_02。

```
upstream backend_02 {
    server cats-backend-Mars:80 max_fails=1 fail_timeout=10s weight=1;
    server cats-backend-Jupiter:80 backup;
}
 
upstream backend_01 {
    server cats-backend-Jupiter:80 max_fails=1 fail_timeout=10s weight=1;
    server cats-backend-Mars:80 backup;
}
 
upstream backend_default {
    server cats-backend-Jupiter:80 max_fails=1 fail_timeout=10s weight=1;
    server cats-backend-Mars:80 backup;
}
 
server {
    listen       8091;
 
    access_log  logs/access.log  main;
    set $resp_body "";
    set $group backend_default;
 
     location / {
        access_by_lua_file /etc/nginx/conf.d/backend.lua;
        proxy_pass http://$group;
        index  index.html index.htm;
        proxy_redirect off;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header REMOTE-HOST $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        client_max_body_size 50m;
        client_body_buffer_size 256k;
        proxy_connect_timeout 30;
        proxy_send_timeout 30;
        proxy_read_timeout 60;
        proxy_buffer_size 256k;
        proxy_buffers 4 256k;
        proxy_busy_buffers_size 256k;
        proxy_temp_file_write_size 256k;
        proxy_max_temp_file_size 128m;
  }
}
```

## restydoc

```
# Nginx 的说明
$ /usr/local/openresty/bin/restydoc nginx
# LuaJIT 的说明
$ /usr/local/openresty/bin/restydoc luajit
# 包管理工具 opm 说明
$ /usr/local/openresty/bin/restydoc opm
# 命令行工具 resty 的说明
$ /usr/local/openresty/bin/restydoc resty-cli
# ngx_echo 组件的说明
$ /usr/local/openresty/bin/restydoc ngx_echo
# ngx_lua 的说明
$ /usr/local/openresty/bin/restydoc ngx_lua
#  stream_lua 说明
$ /usr/local/openresty/bin/restydoc stream_lua
# lua-cjson 说明
$ /usr/local/openresty/bin/restydoc lua-cjson
# 反向代理指令 proxy_pass 的说明
$ /usr/local/openresty/bin/restydoc -s proxy_pass
# content_by_lua_block 指令的说明
$ /usr/local/openresty/bin/restydoc -s content_by_lua_block
# 功能接口 ngx.say 的说明
$ /usr/local/openresty/bin/restydoc -s ngx.say
# lua 函数 concat 的说明
$ /usr/local/openresty/bin/restydoc -s concat

```
其中的 “-s” 参数用来指定搜索手册里的小节名

对于使用 opm 安装的组件，需要使用 “-r” 参数指定安装目录，例如：
```
$ /usr/local/openresty/bin/restydoc -r /usr/local/openresty/site -s lua-resty-http
```

## 包管理器opm

opm - OpenResty Package Manager
https://github.com/openresty/opm
https://opm.openresty.org/

> OpenResty@1.11.2.1 支持了 opm, 但并没有内置其中, 需要自行下载安装
```
curl https://raw.githubusercontent.com/openresty/opm/master/bin/opm > /usr/local/openresty/bin/opm
chmod +x /usr/local/openresty/bin/opm

export PATH=/usr/local/openresty/bin:$PATH
```

opm 安装包默认是全局模式, 文件都会被放在 /usr/local/openresty/site/ 路径下

因此 opm 的版本控制和包管理是扁平而非嵌套的

A 依赖 C
B 也依赖 C

opm 会下载 A, B, C 各一遍, 不会下载 C 两遍, 逻辑比 npm 简单很多

那问题来了, 如果我就是要不同的版本进行开发呢?
那就是**本地模式**
opm 加上参数 --cwd 就选择了本地模式
opm 包会下载到当前路径中, 实现本地包场景需求
和 npm 正好相反, npm install 默认安装本地包, 加上 -g 才是安装全局包, 而 opm 则默认安装全局包

参考: [OpenResty 大跃进! opm 包管理尝鲜](https://zhuanlan.zhihu.com/p/22991825)


## 配置文件
### 模块加载
可以配置`lua_package_path`和`lua_package_cpath`
```
    lua_package_path "/usr/local/openresty/nginx/lua_modules/?.lua;?.lua;/usr/local/openresty/lualib/?.lua";
    lua_package_cpath "/usr/local/openresty/lualib/resty/?.so;";
    server {
        ...
```

### 使用lua额外处理
```
# 路径xxx,额外使用lua脚本处理
        location ^~ /xxx {
            include "/lua_include_conf/xxx.conf";
            proxy_buffering          off;
            proxy_request_buffering  off;
            proxy_pass http://127.0.0.1:8080;
        }
```
/lua_include_conf/xxx.conf
```
access_by_lua_file          "lua_modules/xxx.lua";
```

lua_modules/xxx.lua
```
代码记录日志到数据库等操作
```

### http
```
http {                      # http 块开始，所有的 HTTP 相关功能
                         
    server {                # server 块，第一个 web 服务
        listen       80;    # 监听 80 端口
 
        location uri {      # location 块，需指定 URI
            ...             # 定义访问此 URI 时的具体行为
        }                   # location 块结束
    }                       # server 块结束
    server {                # server 块，第二个 web 服务  
        listen       xx;    # 监听 xxx 端口
        ...                 # 其他 location 定义
 
    }                       # server 块结束
}                           # http 块结束
```
由于 http 块内容太多，如果都写在一个文件里可能会造成配置文件过度庞大，难以维护。在实践中通常把 server、location 等配置分离到单独的文件，再利用 include 指令包含进来，这样就可以很好地降低配置文件的复杂度。

```
http {
    include common.conf     # 基本的 HTTP 配置文件，配置通常参数
    include servers/*.conf  # 包含 server 目录下所有 Web 服务配置文件
}
```

#### server 配置

erver 指令在 http 块内定义一个 Web 服务，它必须是一个匹配块，在快内部再用其他指令来确定 Web 服务的端口、域名、URI 处理等细节。

`location [ = | ! | ~* | ^* ] uri { ... }`
listen 指令使用 port 参数设置 Web 服务监听的端口，默认是 80。此外还可以添加其他很多参数，例如 IP 地址、SSL、HTTP/2 支持等。

`server_name name...；`
server_name 指令设置 Web 服务的域名，允许使用 “*” 通配符或 “~” 开头的正则表达式。例如 “www.openresty.org” "*.openresty.org"。当 OpenResty 处理请求时将会检查 HTTP 头部的 Host 字段，只有与 server_name 匹配的 server 块才会真正提供服务。

对于我们自己的开发研究来说，可以直接使用 localhost 或者简单的通配符 “*”，用类似 “curl http://localhost/...” 这样的命令就足够访问 OpenResty。

#### location 配置
location 指令定义 web 服务的端口（想待遇 RESTful 里的 API），也就是 URI，它是 OpenResty 处理的入口，决定了请求应该如何处理。

location 是一个配置快，但语法稍多一些，除 { } 还有其他的参数：

`location [ = | ! | ~* | ^* ] uri { ... }`
location 使用 uri 参数匹配 HTTP 请求里的 URI，默认是前缀匹配，也支持正则表达式，uri 参数前可以使用特殊标记进行下一步限定匹配；

- `=`：URI 必须完全匹配。
- `~`：大小写敏感匹配。
- `~*`：大小写不敏感匹配。
- `^~`：前缀匹配，匹配 URI 的前半部分即可。

在 server 块里可以配置任意数量的 location 块，定义 web 服务接口。Nginx 对 location 的顺序没有特殊要求，并不是按照配置文件里的顺序逐个查找匹配，而是对所有可能的匹配进行排序，查找最佳匹配的 location。

不同的 location 里可以有不同的处理方式，灵活设置 location 能够让 OpenResty 配置清晰明了，易于维护。比如，可以在一个 location 里存放静态 html 文件，在另一个 location 里存放图片文件，其他的 location 则执行 Lua 程序访问 MySQL 数据库处理动态业务，这些location 互不干涉，修改其中的一个不会影响正常运行。例如：
```
location  =  /502.html        # 只处理 /502.html 这一个文件。
location     /item            # 前缀匹配 /item/*
location ^~  /image/          # 显示前缀匹配 /image/*
location  ~  /articles/(\d+)$ # 正则匹配 /articles/*
location     /                # 匹配任意的 URI
```
需要注意最后一个 “/”，根据前缀匹配规则，它能够匹配任意的 URI，所以可以把它当做一个 “黑洞”，处理所有其他 location 不能处理的请求（例如返回 404）。

如果 location 配置很多，同样可以使用 include 的方式来简化配置。


### TCP/UDP
配置 TCP/UDP 相关的功能需要使用指令 stream { }，形式与 http 块非常相似，例如：
```
stream {                    # stream 块开始，TCP/UDP 相关功能
 
 
    server {                # server 块，第一个 Web 服务
        listen 53;          # 监听 TCP 53 端口
        ...
    }                       # server 块结束
 
 
    server {                # server 块，第二个 Web 服务
        listen 520 udp;     # 监听 UDP 520 端口
        ....
    }                       # server 块结束
}                           # stream 块结束
```

定义 TCP/UDP 服务同样需要使用 server 指令，然后再用 listen 指令确定服务使用的具体端口号。但因为 TCP/UDP 协议里没有 “HOST” “URI" 概念，所以 server 块里不能使用 server_name 和 location 指令，这是与 HTTP 服务明显不同的地方，需要注意。

## 开发框架 (web framework)
### lor
https://github.com/sumory/lor
### lapis
https://github.com/leafo/lapis
### openresty-smart-panda
https://github.com/BBD-RD/openresty-smart-panda/
openresty lua模块化开发的框架，用来简化nginx的配置、规范开发过程、降低开发难度、减少代码耦合性、提高多人协同工作等。


## OpenResty Edge
### EdgeLang
EdgeLang 是由 OpenResty Inc. 创建的一种领域特定语言，它让您可以为 OpenResty Edge 产品编写简洁而富有表达力的规则。
https://blog.openresty.com.cn/cn/edgelang-intro/
https://doc.openresty.com.cn/cn/edge/edgelang/

## 资料
[Redis、Lua、Nginx、OpenResty 笔记和资料](https://github.com/Tinywan/lua-nginx-redis)