# docker 安装rabbitmq
[教程](https://hub.docker.com/_/rabbitmq)

安装带管理插件的mq（Management Plugin）
```
docker pull rabbitmq:3.8.0-management   #180M 有点慢

D盘新建dockerv  redis conf 文件夹 创建redis.conf空文件

docker run -p 6379:6379 -v /d/dockerv/redis/conf/redis.conf:/redis.conf -v /d/dockerv/redis/data:/data --restart always --name redis5.0.6 -d redis:5.0.6 redis-server --appendonly yes

然后就可以连接了

进入容器:
docker exec -it redis5.0.6 /bin/bash

redis-cli

```
