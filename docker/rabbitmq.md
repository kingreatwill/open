# docker 安装rabbitmq
[教程](https://hub.docker.com/_/rabbitmq)

安装带管理插件的mq（Management Plugin）
```
docker pull rabbitmq:3.8.0-management   #180M 有点慢

D盘新建dockerv  rabbitmq conf 文件夹

docker run -d --name rabbitmq3.8.0 -p 5672:5672 -p 15672:15672 -v /d/dockerv/rabbitmq/data/mnesia:/var/lib/rabbitmq/mnesia -e RABBITMQ_DEFAULT_VHOST=/  -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin  --restart always rabbitmq:3.8.0-management

然后就可以连接了

进入容器:
docker exec -it rabbitmq3.8.0 /bin/bash

```
