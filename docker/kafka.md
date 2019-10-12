docker 安装kafka
1、下载镜像
这里使用了wurstmeister/kafka和wurstmeister/zookeeper这两个版本的镜像


docker pull wurstmeister/zookeeper
docker pull wurstmeister/kafka
在命令中运行docker images验证两个镜像已经安装完毕

2.启动
启动zookeeper容器
docker run -d --name zookeeper -p 2181:2181 -t wurstmeister/zookeeper

启动kafka容器

docker run -d --name kafka --publish 9092:9092 --link zookeeper --env KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 --env KAFKA_ADVERTISED_HOST_NAME=192.168.59.101 --env KAFKA_ADVERTISED_PORT=9092 --volume /etc/localtime:/etc/localtime wurstmeister/kafka:latest

192.168.59.101 改为宿主机器的IP地址，如果不这么设置，可能会导致在别的机器上访问不到kafka。

3. 测试kafka
进入kafka容器的命令行

 

运行 docker ps，找到kafka的 CONTAINER ID，运行 docker exec -it ${CONTAINER ID} /bin/bash，进入kafka容器。
进入kafka默认目录 /opt/kafka_2.11-0.10.1.0


参考地址：https://blog.csdn.net/lblblblblzdx/article/details/80548294