

安装nginx

docker run -d -p 8080:80 --name nginx  nginx:1.17.4
拷贝:
docker cp b3735bdb389a2:/etc/nginx/nginx.conf d:/dockerv/nginx/conf
也可以下载 https://github.com/nginx/nginx/blob/master/conf/nginx.conf

```
docker pull nginx:1.17.4

docker run -d -p 8080:80 --name nginx -v /d/dockerv/nginx/www:/usr/share/nginx/html -v /d/dockerv/nginx/conf/nginx.conf:/etc/nginx/nginx.conf -v /d/dockerv/nginx/logs:/var/log/nginx --restart always nginx:1.17.4

```

