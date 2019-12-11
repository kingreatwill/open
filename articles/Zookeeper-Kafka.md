https://www.cnblogs.com/qingyunzong/p/9004509.html

# Zookeeper
## 下载Zookeeper
wget https://mirrors.tuna.tsinghua.edu.cn/apache/zookeeper/zookeeper-3.5.6/apache-zookeeper-3.5.6-bin.tar.gz

mkdir  -p /usr/local/zookeeper
mkdir -p /usr/local/zookeeper/server1
mkdir -p /usr/local/zookeeper/server2
mkdir -p /usr/local/zookeeper/server3

tar zxvf apache-zookeeper-3.5.6-bin.tar.gz -C /usr/local/zookeeper/server1
tar zxvf apache-zookeeper-3.5.6-bin.tar.gz -C /usr/local/zookeeper/server2
tar zxvf apache-zookeeper-3.5.6-bin.tar.gz -C /usr/local/zookeeper/server3

## 创建zookerper各个节点的data和logs目录以及zookeeper节点标识文件的myid
分别
cd  /usr/local/zookeeper/server1/apache-zookeeper-3.5.6-bin
mkdir -p data
mkdir -p logs
cd data
```
vim myid
1 #server1的id
```




## 修改server1的配置文件
分别
cd  /usr/local/zookeeper/server1/apache-zookeeper-3.5.6-bin
cp zoo_sample.cfg  zoo_sample.cfg.bak   #备份默认配置文件
mv zoo_sample.cfg  zoo.cfg  #zookeeper启动配置文件
vim zoo.cfg
```
tickTime=2000   #Zookeeper 服务器之间或客户端与服务器之间维持心跳的时间间隔，也就是每个 tickTime 时间就会发送一个心跳。tickTime以毫       秒为单位。tickTime=2000

 initLimit=10    #集群中的follower服务器(F)与leader服务器(L)之间初始连接时能容忍的最多心跳数（tickTime的数量）。

 syncLimit=5     #集群中的follower服务器与leader服务器之间请求和应答之间能容忍的最多心跳数（tickTime的数量）。

 dataDir=/opt/zookeeper/server1/zookeeper-3.4.12-bin/data   #Zookeeper保存数据的目录

 dataLogDir=/opt/zookeeper/server1/zookeeper-3.4.12-bin/logs #Zookeeper将写数据的日志文件保存在这个目录里。

 clientPort=2181  # 2182 2183客户端连接 Zookeeper 服务器的端口，Zookeeper 会监听这个端口，接受客户端的访问请求。

 server.1=127.0.0.1:2888:3888 #服务器名称与地址：集群信息（服务器编号，服务器地址，LF通信端口，选举端口）

 server.2=127.0.0.1:2889:3889

 server.3=127.0.0.1:2890:3890
```

## 启动三个zookeeper节点
/usr/local/zookeeper/server1/apache-zookeeper-3.5.6-bin/bin/zkServer.sh start  
/usr/local/zookeeper/server2/apache-zookeeper-3.5.6-bin/bin/zkServer.sh start 
/usr/local/zookeeper/server3/apache-zookeeper-3.5.6-bin/bin/zkServer.sh start 

查看主从
/usr/local/zookeeper/server3/apache-zookeeper-3.5.6-bin/bin/zkServer.sh status

## 查看端口
netstat -ntulp |grep 80 


# kafka
https://kafka.apache.org/quickstart



## 下载kafka
wget -O kafka_2.12-2.3.1.tgz  http://mirror.bit.edu.cn/apache/kafka/2.3.1/kafka_2.12-2.3.1.tgz
mkdir  -p /usr/local/kafka
cp   kafka_2.12-2.3.1.tgz   /usr/local/kafka
cd /usr/local/kafka
tar -zxvf kafka_2.12-2.3.1.tgz
## 修改配置文件
```
vim /usr/local/kafka/kafka_2.12-2.3.1/config/server.properties 修改参数


zookeeper.connect=192.168.110.231:2181
listeners=PLAINTEXT://:9092
advertised.listeners=PLAINTEXT://192.168.110.231:9092 #本机ip
```

## 启动server
```
cd /usr/local/kafka/kafka_2.12-2.3.1
bin/kafka-server-start.sh  -daemon  config/server.properties &




# bin/kafka-server-stop.sh
```
### 查看所有topic
```
bin/kafka-topics.sh --list --zookeeper 192.168.110.231:2181
```
查看指定topic 下面的数据
```
bin/kafka-console-consumer.sh --bootstrap-server 192.168.110.231:9092  --from-beginning --topic example
```
```
bin/kafka-console-producer.sh --broker-list 192.168.110.231:9092 --topic example
```