docker 安装kafka
参考地址：https://blog.csdn.net/lblblblblzdx/article/details/80548294
```
docker pull wurstmeister/kafka:2.12-2.3.0

docker run -d --name kafka --publish 9092:9092 --link zookeeper --env KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 --env KAFKA_ADVERTISED_HOST_NAME=192.168.59.101 --env KAFKA_ADVERTISED_PORT=9092 --volume /d/dockerv/kafka/localtime:/etc/localtime wurstmeister/kafka:2.12-2.3.0

192.168.59.101 改为宿主机器的IP地址，如果不这么设置，可能会导致在别的机器上访问不到kafka。

3. 测试kafka
进入kafka容器的命令行

 

 docker exec -it kafka /bin/bash


```