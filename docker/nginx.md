
> [Caddy](https://github.com/caddyserver/caddy) 是一个 Go 编写的 Web 服务器,类似于 Nginx

## 文档

[官方文档](http://nginx.org/en/docs)

[NGINX缓存原理及源码分析(一)](https://www.nginx.org.cn/article/detail/403)
[NGINX缓存原理及源码分析(二)](https://www.nginx.org.cn/article/detail/406)

### parsing nginx cache header files/解析缓存文件头部信息
https://github.com/kroemeke/ngcls
```
dft:/cache/nginx/0/00# ngcls 9eacc540a431495f5ae408412e60f000 
version       : 3
valid_sec     : 1443784216 2015-10-02 12:10:16
last_modified : 1443559345 2015-09-29 21:42:25
date          : 1443697816 2015-10-01 12:10:16
etag          : "f3016c-6bb-520e8d9f88e40"
vary_len      : 15
vary          : Accept-Encoding
variant(md5)  : 9eacc540a431495f5ae408412e60f000
key           : httpkroemeke.eu/
HEADERS       : 
HTTP/1.1 200 OK
Date: Thu, 01 Oct 2015 11:10:16 GMT
Server: Apache
Last-Modified: Tue, 29 Sep 2015 20:42:25 GMT
ETag: "f3016c-6bb-520e8d9f88e40"
Accept-Ranges: bytes
Content-Length: 1723
Vary: Accept-Encoding
Cache-Control: proxy-revalidate
Connection: close
Content-Type: text/html
```

less filename 可以查看
```
KEY: xx.com/xx//bytes=94371840-94633983/0
HTTP/1.1 206 Partial Content
Cache-Control: max-age=2592000
Connection: keep-alive
Content-Length: 262144
Content-Range: bytes 94371840-94633983/250537903
Content-Type: video/mp4
Date: Sun, 17 Nov 2024 15:34:21 GMT
Last-Modified: Sun, 17 Nov 2024 14:48:20 GMT
Server: openresty/1.19.1
```


## 在线进行配置

### nginx-ui/WebUI
https://github.com/0xJacky/nginx-ui
### nginxconfig
https://nginxconfig.io/
https://www.digitalocean.com/community/tools/nginx?global.app.lang=zhCN
https://github.com/digitalocean/nginxconfig.io

https://github.com/digitalocean/nginxconfig.io

### Nginx Proxy Manager
https://github.com/NginxProxyManager/nginx-proxy-manager
### playground
曾几何时，playground 似乎成了新语言的标配：Go 发布就带有 https://play.golang.org/，Rust 发布也有 https://play.rust-lang.org/。你想过 Nginx 也有一个 playground 吗？你可以通过它方便的测试 Nginx 配置。

https://nginx-playground.wizardzines.com/

后端的完整代码见这里：https://gist.github.com/jvns/edf78e7775fea8888685a9a2956bc477。

这个网站的作者写了一篇文章介绍它，包括安全问题、性能问题等，有兴趣的可以查看：https://jvns.ca/blog/2021/09/24/new-tool--an-nginx-playground/。另外，还有一个 Nginx location match 测试的网址：https://nginx.viraptor.info/。

[Nginx 竟然也有 playground：还是 Go 语言构建的](https://mp.weixin.qq.com/s/meBX4v3B2XusLVxwdicfeA)

### 其它

[nginxWebUI nginx网页配置工具](https://gitee.com/cym1102/nginxWebUI)
[Nginx Formatter / Nginx 格式化工具](https://github.com/soulteary/nginx-formatter) 一款 10MB 左右的，小巧、简洁的 Nginx 格式化工具，支持命令行、WebUI、Docker、x86、ARM、macOS、Linux。

#### Nginx与其它语言结合
Nginx和java: [Nginx-Clojure](https://github.com/nginx-clojure/nginx-clojure)
Nginx和lua: [Nginx-lua](https://github.com/openresty/lua-nginx-module)
JavaScript: nginScript
PHP: https://github.com/rryqszq4/ngx_php


> 第三方模块: https://www.nginx.com/resources/wiki/modules/

#### httpsok/SSL证书自动续期
[httpsok](https://github.com/httpsok/httpsok) 是一个专为 Nginx、OpenResty 服务器设计的HTTPS证书自动续签工具。它旨在帮助中小企业稳定、安全、可靠地管理SSL证书，让证书续期变得简单高效。


### Nginx管理系统
https://github.com/0xJacky/nginx-ui
https://github.com/alexazhou/VeryNginx
https://github.com/stefanpejcic/OpenPanel
https://github.com/ajenti/ajenti
https://github.com/schenkd/nginx-ui
https://github.com/EasyEngine/easyengine
https://github.com/caprover/caprover
https://github.com/nginx/agent


### nginx替代品
- [Caddy](https://caddyserver.com/)采用Go语言编写, 是一款功能强大，扩展性高的Web服务器(coredns就是使用它)
- [Tengine](http://tengine.taobao.org/) 阿里CDN使用的它
- [Traefik](https://doc.traefik.io/traefik/) 云原生
- [Pingora](https://github.com/cloudflare/pingora) 是Cloudflare使用Rust开发

## docker 安装nginx

docker run -d -p 8080:80 --name nginx  nginx:1.17.4
拷贝:
docker cp b3735bdb389a2:/etc/nginx/nginx.conf d:/dockerv/nginx/conf
也可以下载 https://github.com/nginx/nginx/blob/master/conf/nginx.conf

```
docker pull nginx:1.17.4

docker run -d -p 8080:80 --name nginx -v /d/dockerv/nginx/www:/usr/share/nginx/html -v /d/dockerv/nginx/conf/nginx.conf:/etc/nginx/nginx.conf -v /d/dockerv/nginx/logs:/var/log/nginx --restart always nginx:1.17.4

# nginx:1.19.3-alpine
```
[Nginx多域名配置](https://www.cnblogs.com/goloving/p/9363490.html)

### 定义状态页面
```
location /status {
                stub_status on;
                 #allow IP地址;
                 #deny IP地址;# deny all;
        }
```
访问/status
```
Active connections：当前活动的连接数量。
Accepts：已经接受客户端的连接总数量。
Handled：已经处理客户端的连接总数量。（一般与accepts一致，除非服务器限制了连接数量）。
Requests：客户端发送的请求数量。
Reading：当前服务器正在读取客户端请求头的数量。
Writing：当前服务器正在写响应信息的数量。
Waiting：当前多少客户端在等待服务器的响应。


Active connections：145           
#nginx 正处理的活动连接数145个。
server accepts handled requests
 1749 1749 3198                   
#nginx启动到现在共处理了 1749个连接 ,nginx启动到现在共成功创建 1749次握手 请求丢失数=(握手-连接),可以看出，我们没丢请求;总共处理了3198 次请求。
Reading: 0 Writing: 3 Waiting: 142
#Reading ：nginx读取到客户端的Header信息数。
#Writing ： nginx返回给客户端的Header信息数。
#Waiting ： Nginx已经处理完正在等候下一次请求指令的驻留连接.开启keep-alive的情况下,这个值等于active–(reading+writing)。
```

### 定义对静态页面的缓存时间
```
location ~* \.(jpg|jpeg|gif|png|css|js|ico|xml)$ {
  expires        30d;            //定义客户端缓存时间为30天
}
```

### nginx 绑定cpu到worker进程
worker_processes 4;
worker_cpu_affinity 0001 0010 0100 1000;

四核心系统，每个cpu绑定到一个work进程的写法。 

worker_processes 8;
worker_cpu_affinity 00000001 00000010 00000100 00001000 00010000 00100000 01000000 10000000;

八核心系统，每个cpu绑定到一个work进程的写法。
有多少个核，就有几位数，1表示该内核开启，0表示该内核关闭。

worker_processes  auto;
worker_cpu_affinity   auto;

### nginx最大并发连接数
[worker_connections](http://nginx.org/en/docs/ngx_core_module.html#worker_connections)

连接数包含与代理服务器的连接以及与客户端的连接, 同时总连接数(worker_connections*worker_process)不能超过最大打开文件数,
查看最大文件数`ulimit -a|grep "open files"` 或者 `ulimit -n`

```
客户端正向代理数 =   worker_rlimit_nofile / 2 >=  worker_connections * worker_process /2
客户端可连反向代理数据 =  worker_rlimit_nofile /4  >=  worker_connections * worker_process /4

nginx作为http服务器的时候：
max_clients = worker_processes * worker_connections/2

nginx作为反向代理服务器的时候：
max_clients = worker_processes * worker_connections/4
```

> 修改linux的最大打开文件数`/etc/security/limits.conf 增加到65535` ;同时修改`worker_rlimit_nofile ` 和 `worker_connections  `

**multi_accept**指令使得NGINX worker能够在获得新连接的通知时尽可能多的接受连接。 此指令的作用是立即接受所有连接放到监听队列中。 如果指令被禁用(	
Default: multi_accept off;)，worker进程将逐个接受连接。
```
events{
      multi_accept on;
}


如果use kqueue;, 则multi_accept无用
events{
    use kqueue;
    multi_accept on;
}
```

### 关闭 nginx

进入 nginx 下的 sbin 目录下执行命令：
快速停止 nginx 命令：./nginx -s stop
完整有序的停止 nginx：./nginx -s quit；这个命令会等待所有请求结束后再关闭 nginx。

使用 kill 命令关闭 nginx

先查看 nginx 进程，找到主进程号。然后再使用 kill 命令杀死 nginx 进程。
查看 nginx 进程号命令：ps -ef | grep nginx
从容停止 nginx 命令：kill -QUIT 主进程号
快速停止 nginx 命令：kill -TERM 主进程号
强制停止 nginx 命令：kill -HUP 主进程号
平滑重启 nginx 命令：kill -9 nginx
查找并杀死所有 nginx 进程：ps aux | grep nginx |awk '{print $2}' | xargs kill -9


### Nginx 性能优化

[如何构建高性能服务器（以Nginx为例）](https://www.cnblogs.com/kukafeiso/p/13957174.html)
[Nginx 性能优化有这篇就够了](https://www.cnblogs.com/cheyunhua/p/10670070.html)

### ngx_http_auth_digest 模块进行登录认证

生成登录密码：

apt-get install -y apache2-utils
htdigest -c /usr/local/passwd.digest theia admin

- /usr/local/passwd.digest 为密码文件路径
- theia 为 realm(领域，范围)，必须与后面要提到的 nginx 配置文件保持一致。
- admin 为登录用户名

在 nginx 配置文件中添加 server 段：
```
server {
	listen 80 default_server;

    auth_digest_user_file /usr/local/passwd.digest;
    auth_digest_shm_size  8m;   # the storage space allocated for tracking active sessions
    
    location / {
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";

        auth_digest 'theia';
        auth_digest_timeout 60s;    # allow users to wait 1 minute between receiving the
                                    # challenge and hitting send in the browser dialog box
        auth_digest_expires 600s;   # after a successful challenge/response, let the client
                                    # continue to use the same nonce for additional requests
                                    # for 600 seconds before generating a new challenge
        auth_digest_replays 60;     # also generate a new challenge if the client uses the
                                    # same nonce more than 60 times before the expire time limit

        proxy_pass http://127.0.0.1:3000;
    }
}
```

- auth_digest_user_file 必须与设置密码的 htdigest 命令中的密码文件参数保持一致。
- auth_digest 必须与设置密码的 htdigest 命令的 realm 参数保持一致。
- 这两行必须有，否则不能正确代理 websocket。
proxy_set_header Upgrade $http_upgrade; 
proxy_set_header Connection "upgrade";
   
这个配置相当于每 (auth_digest_timeout+auth_digest_expires)=660s 允许 auth_digest_shm_size/((48 + ceil(auth_digest_replays/8))bytes)=(8*1024*1024)/(48+8)=149.8k 次请求，即每 660s 允许约 149.8k 次请求。登录认证 10min 后过期。

最后启动 nginx，会弹出登录认证框，输入用户名和密码后即可登录，跳转到 theia 界面。


### 设置HTTP强制跳转HTTPS
```
server {
    listen 80;
    server_name example.com;  这里修改为网站域名
    rewrite ^(.*)$ https://$host$1 permanent;
}
```

[高性能 Nginx HTTPS 调优 - 如何为 HTTPS 提速 30%](https://www.toutiao.com/article/7197285537810514435/)

### nginx负载均衡策略简介

1、轮询（默认策略，nginx自带策略）：
它是upstream模块默认的负载均衡默认策略。会将每个请求按时间顺序分配到不同的后端服务器。
```
http {
    upstream my_load_balance {
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}
```

2、weight（权重，nginx自带策略）：
指定轮询的访问几率，用于后端服务器性能不均时调整访问比例。权重越高，被分配的次数越多。
```
http {
    upstream my_load_balance {
        server 192.168.1.12:80 weight=7;
        server 192.168.1.13:80 weight=2;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}　
```

3、ip_hash（依据ip分配，nginx自带策略）：
指定负载均衡器按照基于客户端IP的分配方式，这个方法确保了相同的客户端的请求一直发送到相同的服务器，可以解决session不能跨服务器的问题。
```
http {
    upstream my_load_balance {
        ip_hash;
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}
```

4、least_conn（最少连接，nginx自带策略）：
把请求转发给连接数较少的后端服务器。
```
http {
    upstream my_load_balance {
        #把请求转发给连接数比较少的服务器
        least_conn;
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}   
```

5、fair（第三方）：
按照服务器端的响应时间来分配请求，响应时间短的优先分配。
```
http {
    upstream my_load_balance {
        fair;
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}  
```
6、url_hash（第三方）：
该策略按访问url的hash结果来分配请求，使每个url定向到同一个后端服务器，需要配合缓存用。
```
http {
    upstream my_load_balance {
        hash $request_uri;
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
} 
```

## k8s
kubectl create configmap confnginx --from-file nginx.conf
```
user  nginx;
worker_processes  1;
 
error_log  /var/log/nginx/error.log warn;
pid        /var/run/nginx.pid;
 
events {
    worker_connections  1024;
}
 
http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;
 
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                        '$status $body_bytes_sent "$http_referer" '
                        '"$http_user_agent" "$http_x_forwarded_for"';
 
    access_log  /var/log/nginx/access.log  main;
 
    sendfile        on;
    #tcp_nopush     on;
 
    keepalive_timeout  65;
 
    #gzip  on;
 
    #include /etc/nginx/conf.d/*.conf;
 
    server {
       listen                    80;
       server_name               localhost;
       root                      /home/wwwroot/html;
       index                     index.html;
    }
}
```

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx-deployment
spec:
  replicas: 1
  selector:
    name: nginx
  template:
    metadata:
      labels:
        name: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.19.3-alpine
        ports:
        - containerPort: 80
        volumeMounts:
        - mountPath: /etc/nginx/nginx.conf
          name: nginx-config
          subPath: nginx.conf
        - mountPath: /home/wwwroot/html
          name: nginx-data
      volumes:
        - name: nginx-config
          configMap:
            name: confnginx
        - name: nginx-data
          hostPath:
            path: /home/wwwroot/hello

---

apiVersion: v1
kind: Service
metadata:
  name: nginx-service
spec:
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
      nodePort: 30080    #外网访问端口
  type: NodePort   #这个是端口类型
  selector:
    name: nginx
	
--- 

apiVersion: extensions/v1beta1
kind: Ingress
metadata:
    name: nginx-ingress
    annotations:
        kubernetes.io/ingress.class: "nginx"
spec:
    rules:
    - host: violet.chtwebapi.lingcb1.com
      http:
          paths:
          - path: /plan
            backend:
                serviceName: nginx-service
                servicePort: 80
          

```


## 四层代理

```
stream{
    upstream backend_coredns_udp {
        server xxxx:53;
        server xxx:53 backup;
    }

    server {
        listen 53 udp;
        listen 53; #tcp
        proxy_pass backend_coredns_udp;
        # proxy_protocol      on;
    }

    upstream dns {
       server 192.168.0.1:53535;
       server dns.example.com:53;
    }

    server {
        listen 127.0.0.1:53 udp reuseport;
        proxy_timeout 20s;
        proxy_pass dns;
    }
}




```

## 同一location块中同时直接使用proxy_pass和root指令

1. 使用try_files指令
```
server {
    listen 80;
    server_name example.com;

    location / {
        root /var/www/html;
        try_files $uri @backend;
    }

    location @backend {
        proxy_pass http://backend_server;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```
- 首先尝试使用root指令指定的目录来响应请求（/var/www/html）。try_files $uri @backend;语句会检查请求的文件是否存在，如果存在则直接提供，不存在则转到命名的location @backend。
- @backend是一个命名的location，用于将请求代理到后端服务器。

2. 基于请求URI或参数的条件判断
```
location / {
    if ($request_uri ~* "^/static/") {
        root /var/www/static;
    }
    # 或者使用嵌套location
    location /api/ {
        proxy_pass http://backend_api;
    }
}
```

3. 使用命名location
```
location / {
    root /var/www/html;
    error_page 404 = @proxy;
}

location @proxy {
    proxy_pass http://backend;
}
```
