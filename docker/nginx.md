
> [Caddy](https://github.com/caddyserver/caddy) 是一个 Go 编写的 Web 服务器,类似于 Nginx

## 在线进行配置

### nginxconfig
https://nginxconfig.io/
https://www.digitalocean.com/community/tools/nginx?global.app.lang=zhCN
https://github.com/digitalocean/nginxconfig.io

### playground
曾几何时，playground 似乎成了新语言的标配：Go 发布就带有 https://play.golang.org/，Rust 发布也有 https://play.rust-lang.org/。你想过 Nginx 也有一个 playground 吗？你可以通过它方便的测试 Nginx 配置。

https://nginx-playground.wizardzines.com/

后端的完整代码见这里：https://gist.github.com/jvns/edf78e7775fea8888685a9a2956bc477。

这个网站的作者写了一篇文章介绍它，包括安全问题、性能问题等，有兴趣的可以查看：https://jvns.ca/blog/2021/09/24/new-tool--an-nginx-playground/。另外，还有一个 Nginx location match 测试的网址：https://nginx.viraptor.info/。

[Nginx 竟然也有 playground：还是 Go 语言构建的](https://mp.weixin.qq.com/s/meBX4v3B2XusLVxwdicfeA)

### 其它
https://www.nginxedit.cn/

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

### 如何构建高性能服务器（以Nginx为例）
https://www.cnblogs.com/kukafeiso/p/13957174.html

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
