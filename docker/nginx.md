

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
