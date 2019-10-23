# docker 安装redis
[教程](https://hub.docker.com/_/redis)
[参考](https://www.mgchen.com/268.html)

多线程版本:
https://hub.docker.com/r/eqalpha/keydb
安装5.0.6
```
docker pull redis:5.0.6   #98M

下载redis.conf
http://download.redis.io/redis-stable/redis.conf
https://github.com/antirez/redis/blob/5.0.6/redis.conf

修改 redis.conf
bind 0.0.0.0
appendonly yes

D盘新建dockerv  redis conf 文件夹 

docker run -p 6379:6379 -v /d/dockerv/redis/conf/redis.conf:/redis.conf -v /d/dockerv/redis/data:/data --restart always --name redis5.0.6 -d redis:5.0.6 redis-server /redis.conf

然后就可以连接了

进入容器:
docker exec -it redis5.0.6 /bin/bash

redis-cli

```

安装3.2
```
docker run -p 6379:6379 --restart always  -d redis:3.2 redis-server --appendonly yes
```

安装RediSearch:
```
 docker run -p 6379:6379 redislabs/redisearch:latest
```