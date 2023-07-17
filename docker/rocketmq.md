
# Apache RocketMQ

## 文档
https://rocketmq.apache.org/zh/docs/featureBehavior/06consumertype

https://rocketmq-1.gitbook.io/rocketmq-connector/rocketmq-connect/rocketmq-console/shi-yong-zhi-nan

## docker安装
### rocketmq4.x
- **NameServer服务启动**
```
docker run -d -p 9876:9876 --name mqNameServe foxiswho/rocketmq:server-4.3.2


docker run -d --restart=always -v /docker/rocketmq/data/namesrv/logs:/home/rocketmq/logs --name rmqnamesrv -e "JAVA_OPT_EXT=-Xms512M -Xmx512M -Xmn128m" -p 9876:9876 foxiswho/rocketmq:4.8.0
```

- **Borker服务启动**
```
docker run -d -p 10911:10911 -p 10909:10909 --name mqBorker --link mqNameServer:namesrv -e "NAMESRV_ADDR=namesrv:9876" -e "JAVA_OPTS=-Duser.home=/opt" -e "JAVA_OPT_EXT=-server -Xms128m -Xmx128m"  foxiswho/rocketmq:broker-4.5.1
```

- **Console控制台启动,也就是UI**
```
docker run -d --name mqConsole -p 9000:8080 --link mqNameServer:namesrv\
 -e "JAVA_OPTS=-Drocketmq.namesrv.addr=namesrv:9876\
 -Dcom.rocketmq.sendMessageWithVIPChannel=false"\
 -t styletang/rocketmq-console-ng
```


### rocketmq5.x

> 控制台：8080；namesrv：9876；broker：10909、10911。
当前
apache/rocketmq:5.1.3

- **部署NameServer**
```sh
# 日志目录
mkdir /usr/local/rocketmq/nameserver/logs -p
# 脚本目录
mkdir /usr/local/rocketmq/nameserver/bin -p

# 777 文件所属者、文件所属组和其他人有读取 & 写入 & 执行全部权限。rwxrwxrwx
chmod 777 -R /usr/local/rocketmq/nameserver/*
```

1. 启动容器, 复制出启动脚本, 后面删除
```sh
docker run -d \
--privileged=true \
--name rmqnamesrv \
apache/rocketmq:5.1.3 sh mqnamesrv


docker cp rmqnamesrv:/home/rocketmq/rocketmq-5.1.3/bin/runserver.sh /usr/local/rocketmq/nameserver/bin/runserver.sh

```
1. 修改runserver.sh
> NameServer启动脚本中有一个自动计算最大堆内存和新生代内存的函数会导致在不同硬件环境下设置最大堆内存和新生代内存环境变量

找到调用calculate_heap_sizes函数的位置注释掉保存即可，拉到脚本最底部就能找到

1. 停止&删除容器
```sh
docker stop rmqnamesrv
docker rm rmqnamesrv
```

1. 启动NameServer
```sh
sudo docker run -d \
--privileged=true \
--restart=always \
--name rmqnamesrv \
-p 9876:9876  \
-v /usr/local/rocketmq/nameserver/logs:/home/rocketmq/logs \
-v /usr/local/rocketmq/nameserver/bin/runserver.sh:/home/rocketmq/rocketmq-5.1.3/bin/runserver.sh \
-e "MAX_HEAP_SIZE=256M" \
-e "HEAP_NEWSIZE=128M" \
apache/rocketmq:5.1.3 sh mqnamesrv

```

- **部署Broker**
```sh
# 创建需要的挂载目录
mkdir /usr/local/rocketmq/broker/logs -p \
mkdir /usr/local/rocketmq/broker/data -p \
mkdir /usr/local/rocketmq/broker/conf -p \
mkdir /usr/local/rocketmq/broker/bin -p
# 777 文件所属者、文件所属组和其他人有读取 & 写入 & 执行全部权限。rwxrwxrwx
chmod 777 -R /usr/local/rocketmq/broker/*
```

建立broker.conf文件，通过这个文件把RocketMQ的broker管理起来
`vi /usr/local/rocketmq/broker/conf/broker.conf`
内容如下
```
# nameServer 地址多个用;隔开 默认值null
# 例：127.0.0.1:6666;127.0.0.1:8888 
namesrvAddr = 192.168.10.220:9876
# 集群名称
brokerClusterName = DefaultCluster
# 节点名称
brokerName = broker-a
# broker id节点ID， 0 表示 master, 其他的正整数表示 slave，不能小于0 
brokerId = 0
# Broker服务地址	String	内部使用填内网ip，如果是需要给外部使用填公网ip
brokerIP1 = 192.168.10.220
# Broker角色
brokerRole = ASYNC_MASTER
# 刷盘方式
flushDiskType = ASYNC_FLUSH
# 在每天的什么时间删除已经超过文件保留时间的 commit log，默认值04
deleteWhen = 04
# 以小时计算的文件保留时间 默认值72小时
fileReservedTime = 72
# 是否允许Broker 自动创建Topic，建议线下开启，线上关闭
autoCreateTopicEnable=true
# 是否允许Broker自动创建订阅组，建议线下开启，线上关闭
autoCreateSubscriptionGroup=true

```
1. 启动容器, 复制出启动脚本, 后面删除
```sh
docker run -d \
--name rmqbroker \
--privileged=true \
apache/rocketmq:5.1.3 \
sh mqbroker


docker cp rmqbroker:/home/rocketmq/rocketmq-5.1.3/bin/runbroker.sh /usr/local/rocketmq/broker/bin/runbroker.sh

```

1. 修改runbroker.sh, 找到调用calculate_heap_sizes函数的位置注释掉保存即可
```
# 打开脚本文件
vi /usr/local/rocketmq/broker/bin/runbroker.sh 
```

1. 停止&删除容器
```sh
docker stop rmqbroker
docker rm rmqbroker
```

1. 启动Broker
```sh
docker run -d \
--restart=always \
--name rmqbroker \
-p 10911:10911 -p 10909:10909 \
--privileged=true \
-v /usr/local/rocketmq/broker/logs:/root/logs \
-v /usr/local/rocketmq/broker/store:/root/store \
-v /usr/local/rocketmq/broker/conf/broker.conf:/home/rocketmq/broker.conf \
-v /usr/local/rocketmq/broker/bin/runbroker.sh:/home/rocketmq/rocketmq-5.1.3/bin/runbroker.sh \
-e "MAX_HEAP_SIZE=512M" \
-e "HEAP_NEWSIZE=256M" \
apache/rocketmq:5.1.3 \
sh mqbroker -c /home/rocketmq/broker.conf

```

- **部署Proxy**
这里暂时不对Proxy进行部署，单机版本没必要，不使用Proxy和之前的4.x版本基本一致

- **RocketMQ Dashboard**
```sh
$ docker run -d --name rocketmq-dashboard -e "JAVA_OPTS=-Drocketmq.namesrv.addr=127.0.0.1:9876" -p 8080:8080 -t apacherocketmq/rocketmq-dashboard:latest
```

```sh
docker run -d \
--restart=always \
--name rocketmq-dashboard \
-e "JAVA_OPTS=-Xmx256M -Xms256M -Xmn128M -Drocketmq.namesrv.addr=192.168.10.220:9876 -Dcom.rocketmq.sendMessageWithVIPChannel=false" \
-p 8080:8080 \
apacherocketmq/rocketmq-dashboard

```