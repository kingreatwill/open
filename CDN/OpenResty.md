[TOC]

> 很多CDN 行业使用OpenResty，多是受到 cloudflare 技术栈的影响，OpenResty 的作者也在国外这家 CDN 公司。

> 关于作者 章亦春 先后在中国雅虎、淘宝、Cloudflare 就职, 是开源 OpenResty ® 项目创始人兼 OpenResty Inc. 公司 CEO 和创始人。 章亦春（Github ID: agentzh），生于中国江苏，现定居美国湾区。

> OpenResty 最早是雅虎中国的一个公司项目，起步于 2007 年 10 月。当时兴起了 OpenAPI 的热潮，用于满足各种 Web Service 的需求，就诞生了 OpenResty。在公司领导的支持下，最早的 OpenResty 实现从一开始就开源了。最初的定位是服务于公司外的开发者，像其他的 OpenAPI 那样，但后来越来越多地是为雅虎中国的搜索产品提供内部服务。

> 章亦春在加入淘宝数据部门的量子团队之后，决定对 OpenResty 进行重新设计和彻底重写，并把应用重点放在支持像量子统计这样的 web 产品上面，所以量子统计 3.0 开始也几乎完全是 Web Service 驱动的纯 AJAX 应用。这是第二代的 OpenResty，一般称之为 ngx_openresty，以便和第一代基于 Perl 和 Haskell 实现的 OpenResty 加以区别。章亦春和他的同事王晓哲一起设计了第二代的 OpenResty。在王晓哲的提议下，选择基于 Nginx 和 Lua 进行开发。

> 2007年大学毕业，加入中国雅虎搜索技术部。当时的中国雅虎已经是阿里旗下的一家公司了。2009 年被调到淘宝数据产品与平台部，集中力量和同事们一起基于 OpenResty 开发淘宝量子统计这款产品，主要面向淘宝卖家。2011 年离开阿里，在家里全心从事 OpenResty 等开源项目的工作。 2012 年受 Cloudflare 公司邀请，举家来到美国湾区，全职从事 OpenResty 开源开发，以及基于 OpenResty 的全球 CDN 网络的软件系统的研发和完善。2016 年离开了 Cloudflare，并于 2017 年初在美国创办了 OpenResty Inc. 公司。


OpenResty 在 1.5.8.1 版之后才默认使用 LuaJIT，之前使用的是标准 Lua 解析器。
## 安装

http://openresty.org/cn/download.html

nginx文档:
https://nginx.org/en/docs/

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

### centos 7 在线安装nginx 查看gcc，pcre，zlib，openssl 等依赖是否安装
1. gcc
检查是否安装:`gcc -v`
安装:`yum install -y gcc gcc-c++ kernel-devel`

2. PCRE  安装
PCRE(Perl Compatible Regular Expressions)是一个轻量级的Perl函数库，包括 perl 兼容的正则表达式库。它比Boost之类的正则表达式库小得多。PCRE十分易用，同时功能也很强大，性能超过了POSIX正则表达式库和一些经典的正则表达式库。

查看是否安装:`rpm -qa pcre`
安装:`yum install -y pcre pcre-devel`

3. zlib 安装

zlib是提供数据压缩用的函式库，由Jean-loup Gailly与Mark Adler所开发，初版0.9版在1995年5月1日发表。zlib使用DEFLATE算法，最初是为libpng函式库所写的，后来普遍为许多软件所使用。此函式库为自由软件，使用zlib授权。截至2007年3月，zlib是包含在Coverity的美国国土安全部赞助者选择继续审查的开源项目。

zlib 库提供了很多种压缩和解压缩的方式， nginx 使用 zlib 对 http 包的内容进行 gzip ，所以需要在 Centos 上安装 zlib 库。

查看是否安装:`yum list installed | grep zlib*`
安装:`yum install -y zlib zlib-devel`

4. OpenSSL 安装
openssl是多功能命令工具，用于生成密钥，创建数字证书，手动加密解密数据
nginx 不仅支持 http 协议，还支持 https（即在ssl协议上传输http），所以需要在 Centos 安装 OpenSSL 库。

查看是否安装:`rpm -qa openssl`
安装:`yum install -y openssl openssl-devel`




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


OpenResty 中常用的 HTTP 请求处理阶段：

- init 阶段：这是 OpenResty 处理请求的第一个阶段。在这个阶段，可以执行全局初始化操作，例如加载和初始化一些全局的 Lua 模块、初始化共享内存等。通常情况下，不需要在用户级别的请求处理中使用该阶段。
- ​set​ 阶段：在该阶段，OpenResty 提供的事件钩子 set_by_lua*​ 允许开发人员在处理请求之前设置变量。这个阶段通常用于设置一些全局变量或请求特定的变量，这些变量可以在后续阶段中使用。
- rewrite​ 阶段：在 rewrite​ 阶段，OpenResty 提供的事件钩子 rewrite_by_lua*​ 允许开发人员修改请求的 URI 或其他请求相关的参数。在这个阶段，可以根据特定规则重写请求的路径、添加查询参数、进行重定向等操作。
- ​access​ 阶段：access​ 阶段是处理请求访问权限的阶段。在该阶段，OpenResty 提供的事件钩子 access_by_lua*​ 允许开发人员对请求进行访问控制、鉴权、限流等操作。可以检查用户的身份、验证 API 密钥、对请求进行频率限制等。
- ​content​ 阶段：content​ 阶段是处理请求内容的主要阶段。在该阶段，OpenResty 提供的事件钩子 content_by_lua*​ 允许开发人员编写 Lua 脚本来处理请求、生成响应内容、与后端服务交互等。这个阶段是实现业务逻辑的主要场所。
- ​header_filter​ 阶段：在 header_filter​ 阶段，OpenResty 提供的事件钩子 header_filter_by_lua*​ 允许开发人员修改响应头。可以添加、修改或删除响应头信息，对响应进行进一步处理。
- ​body_filter​ 阶段：body_filter​ 阶段是对响应体进行处理的阶段。在该阶段，OpenResty 提供的事件钩子 body_filter_by_lua*​ 允许开发人员对响应进行修改、过滤或其他操作。可以对响应进行内容转换、压缩、加密等。
- ​log​ 阶段：在 log​ 阶段，OpenResty 提供的事件钩子 log_by_lua*​ 允许开发人员记录请求和响应的日志信息。可以将请求信息、响应信息、自定义日志等写入日志文件或其他日志存储。


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
https://github.com/sumory/lor 1k
### lapis
https://github.com/leafo/lapis 2.9k
### openresty-smart-panda
https://github.com/BBD-RD/openresty-smart-panda/
openresty lua模块化开发的框架，用来简化nginx的配置、规范开发过程、降低开发难度、减少代码耦合性、提高多人协同工作等。

### nginx替代品
- [Caddy](https://caddyserver.com/)采用Go语言编写, 是一款功能强大，扩展性高的Web服务器(coredns就是使用它)
- [Tengine](http://tengine.taobao.org/) 阿里CDN使用的它
- [Traefik](https://doc.traefik.io/traefik/) 云原生
- [Pingora](https://github.com/cloudflare/pingora) 是Cloudflare使用Rust开发


## 分析和诊断工具

### OpenResty XRay
openresty-systemtap-toolkit已经不维护了, 重点开发OpenResty XRay (商业收费)


### openresty-systemtap-toolkit

> [已逐渐不被支持了](https://github.com/openresty/stapxx?tab=readme-ov-file#lj-lua-stacks)

[OpenResty 火焰图工具](https://blog.csdn.net/qq_42353939/article/details/107491914)


[文档](https://github.com/openresty/openresty-systemtap-toolkit/blob/master/README-CN.markdown)

https://github.com/openresty/openresty-systemtap-toolkit

Real-time analysis and diagnostics tools for OpenResty (including NGINX, LuaJIT, ngx_lua, and more) based on SystemTap


- [ngx-lua-bt](https://github.com/openresty/openresty-systemtap-toolkit#ngx-lua-bt)
- [sample-bt](https://github.com/openresty/openresty-systemtap-toolkit#sample-bt)


#### 前提条件
你的 Linux 系统需要 systemtap 2.1+ 和 perl 5.6.1+ 及以上的版本。如果要从源码编译最新版本的 systemtap，你可以参考这个文档：http://openresty.org/#BuildSystemtap

另外，如果你不是从源码编译的 NGINX，你需要保证你的 NGINX 和其他依赖组件的 （DWARF）调试信息已经打开了（或者单独安装了）。

最后，你也需要安装 kernel debug symbols 和 kernel headers。通常只用在你的 Linux 系统中，安装和 kernel 包匹配的 kernel-devel 和 kernel-debuginfo 就可以了。

```
$ uname -r

3.10.0-327.28.2.el7.x86_64

# 安装内核开发包和调试包

$ wget "http://debuginfo.centos.org/7/x86_64/kernel-debuginfo-($version).rpm"

$ wget "http://debuginfo.centos.org/7/x86_64/kernel-debuginfo-common-($version).rpm"

$ sudo rpm -ivh kernel-debuginfo-common-($version).rpm
$ sudo rpm -ivh kernel-debuginfo-($version).rpm
$ sudo yum install kernel-devel-($version)
$ sudo yum install systemtap
```

#### systemtap
systemtap是内核开发者必须要掌握的一个工具

systemtap 官网给出了自学教程及相关论文，选择看这个已经足够了：https://sourceware.org/systemtap/documentation.html
IBM 编写的systemtap 指南也是很不错的： http://www.redbooks.ibm.com/abstracts/redp4469.html

### 火焰图工具 FlameGraph
https://github.com/brendangregg/FlameGraph
火焰图（flame graph）是性能分析的利器，通过它可以快速定位性能瓶颈点。
perf 命令（performance 的缩写）是 Linux 系统原生提供的性能分析工具，会返回 CPU 正在执行的函数名以及调用栈（stack）。

`yum install perf -y`

里面大部分是perf语言写的脚本，生成火焰图后续会用到(直接执行.pl后缀的脚本文件, 对采集的数据进行各种处理)
`git clone https://github.com/brendangregg/FlameGraph.git`

#### perf-tools
https://github.com/brendangregg/perf-tools
#### perf 采集数据
`perf record -F 99 -a -g -- sleep 60` //对CPU所有进程以99Hz采集,它的执行频率是 99Hz（每秒99次），如果99次都返回同一个函数名，那就说明 CPU 这一秒钟都在执行同一个函数，可能存在性能问题。执行60秒后会弹出如下图提示表示采集完成，在当前目录会生成一个perf.data的文件

```
# perf record -F 99 -p 181 -g -- sleep 60        //对进程ID为181的进程进行采集，采集时间为60秒，执行期间不要退出


上述代码中perf record
表示记录，
-F 99
表示每秒99次，
-p 13204
是进程号，即对哪个进程进行分析，
-g
表示记录调用栈，
sleep 30
则是持续30秒，-a 表示记录所有cpu调用。更多参数可以执行
上述代码中perf record
-F 99
-p 13204
-g
sleep 30
perf --help查看。
```

#### 生成火焰图
```
# perf script -i perf.data &> perf.unfold                        //生成脚本文件

# ./FlameGraph/stackcollapse-perf.pl perf.unfold &> perf.folded              

# ./FlameGraph/flamegraph.pl perf.folded > perf.svg                       //执行完成后生成perf.svg图片，可以下载到本地，用浏览器打开 perf.svg，如下图
```

#### 火焰图的含义
y 轴表示调用栈，每一层都是一个函数。调用栈越深，火焰就越高，顶部就是正在执行的函数，下方都是它的父函数。

x 轴表示抽样数，如果一个函数在 x 轴占据的宽度越宽，就表示它被抽到的次数多，即执行的时间长。注意，x 轴不代表时间，而是所有的调用栈合并后，按字母顺序排列的。

火焰图就是看顶层的哪个函数占据的宽度最大。只要有"平顶"（plateaus），就表示该函数可能存在性能问题。

颜色没有特殊含义，因为火焰图表示的是 CPU 的繁忙程度，所以一般选择暖色调。



#### 参考
[火焰图（Flame Graphs）的安装和基本用法](https://www.cnblogs.com/wx170119/p/11459995.html)

#### 性能监控
https://github.com/wolfpld/tracy

## OpenResty Edge
### EdgeLang
EdgeLang 是由 OpenResty Inc. 创建的一种领域特定语言，它让您可以为 OpenResty Edge 产品编写简洁而富有表达力的规则。
https://blog.openresty.com.cn/cn/edgelang-intro/
https://doc.openresty.com.cn/cn/edge/edgelang/

## ATS(Apache Traffic Server)
Apache Traffic Server（ATS或TS）是一个高性能的、模块化的HTTP代理和缓存服务器，与 Nginx 和 Squid 类似。

参考[Apache Traffic Server服务搭建](https://www.cnblogs.com/Dev0ps/p/7891659.html)

```
#官网
https://trafficserver.apache.org/
安装环境
yum install 'liblz*' -y
yum install net-tools -y
yum install gcc gcc-c++ glibc-devel -y
yum install autoconf automake pkgconfig libtool -y
yum install perl-ExtUtils-MakeMaker perl-URI.noarch -y
yum install openssl-devel tcl-devel expat-devel -y
yum install pcre pcre-devel zlib-devel xz-devel -y
yum install libcap libcap-devel flex hwloc hwloc-devel -y
yum install lua-devel curl curl-devel sqlite-devel bzip2 -y
1.安装pcre
[root@web_01 pcre-8.36]# wget http://ftp.exim.llorien.org/pcre/pcre-8.36.tar.gz
[root@web_01 pcre-8.36]# tar xf pcre-8.36.tar.gz  
[root@web_01 pcre-8.36]# cd pcre-8.36
[root@web_01 pcre-8.36]# ./configure --prefix=/usr/local/trafficserver/pcre
[root@web_01 pcre-8.36]# make && make instal
2.安装trafficserver
[root@web_01 ~]# cd /usr/local/src/
[root@web_01 src]# wget https://mirrors.aliyun.com/apache/trafficserver/trafficserver-5.3.2.tar.bz2
[root@web_01 pcre-8.36]# tar xf trafficserver-5.3.2.tar.bz2
[root@web_01 pcre-8.36]# cd trafficserver-5.3.2 
[root@web_01 trafficserver-5.3.2]# ./configure --prefix=/usr/local/trafficserver --with-pcre=/usr/local/trafficserver/pcre --enable-example-plugins --enable-experimental-plugins
[root@web_01 trafficserver-5.3.2]# make  && make install
 注：--enable-example-plugins --enable-experimental-plugins 这两条指令是为了安装ATS官方集成的插件
[root@web_01 trafficserver-5.3.2]# cd /usr/local/trafficserver/bin/
[root@web_01 bin]# ./trafficserver start
Starting Apache Traffic Server:                            [  Ok 

```

```
[root@localhost ~]# ps aux|grep traffic
root      7469  0.0  0.0 129628  7248 ?        Ssl  04:57   0:02 /usr/local/tcacheserver/bin/traffic_cop
176       7472  0.0  0.0 501692 19660 ?        Sl   04:57   0:30 /usr/local/tcacheserver/bin/traffic_manager
176       7482 19.8 20.8 15253324 10279168 ?   Sl   04:57 118:31 /usr/local/tcacheserver/bin/traffic_server
```

## 资料
[Redis、Lua、Nginx、OpenResty 笔记和资料](https://github.com/Tinywan/lua-nginx-redis)

[OpenResty学习指南（一）](https://www.cnblogs.com/luozhiyun/p/12267231.html)

[apigateway技术选型必须要考虑的技术之一OpenResty](https://www.toutiao.com/article/7266448451595076096)