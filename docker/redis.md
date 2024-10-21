# docker 安装redis
[教程](https://hub.docker.com/_/redis)
[参考](https://www.mgchen.com/268.html)




多线程版本:
https://hub.docker.com/r/eqalpha/keydb


安装6.0版本
docker pull redis:6.0.6

https://github.com/redis/redis/blob/6.0.6/redis.conf

docker run -p 6379:6379 -v /d/dockerv/redis6/conf/redis.conf:/redis.conf -v /d/dockerv/redis6/data:/data --restart always --name redis6 -d redis:6.0.6 redis-server /redis.conf



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

### Redis Client On Error: Error: write ECONNABORTED Config right?
如果出现连接不上,需要修改
```
注释掉bind 127.0.0.1, 或者修改bind 0.0.0.0,表示允许所有ip地址访问

修改peotected-mode yes
改为：protected-mode no.
```

### Redis 绑定 CPU
Redis 在 6.0 版本，我们可以通过以下配置，对主线程、后台线程、后台 RDB 进程、AOF rewrite 进程，绑定固定的 CPU 逻辑核心：
```
# Redis Server 和 IO 线程绑定到 CPU核心 0,2,4,6
server_cpulist 0-7:2

# 后台子线程绑定到 CPU核心 1,3
bio_cpulist 1,3

# 后台 AOF rewrite 进程绑定到 CPU 核心 8,9,10,11
aof_rewrite_cpulist 8-11

# 后台 RDB 进程绑定到 CPU 核心 1,10,11
# bgsave_cpulist 1,10-1
```

### redis desktop manager(RDM)/Redis GUI
[Tiny RDM](https://github.com/tiny-craft/tiny-rdm) 8.6k
[Redis Insight](https://github.com/RedisInsight/RedisInsight) 6.2k
[Another Redis Desktop Manager](https://github.com/qishibo/AnotherRedisDesktopManager) 30.4k
[Redis Desktop Manager](https://github.com/RedisInsight/RedisDesktopManager) 22.9k

### 其它语言实现redis协议
https://github.com/DiceDB/dice golang 6.4k
https://github.com/microsoft/Garnet .net 10.2k
https://github.com/dragonflydb/dragonfly C++ 25.6k
https://github.com/Snapchat/KeyDB  C++ 11.4k(redis分支而来)
