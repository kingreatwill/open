https://github.com/alibaba/Alink

[仅1年GitHub Star数翻倍，Flink 做了什么？](https://www.jianshu.com/p/3e83f76b8f4f)
[在 Flink 算子中使用多线程如何保证不丢数据？](https://developer.aliyun.com/article/740572)

[Flink入门（一）——Apache Flink介绍](https://www.cnblogs.com/tree1123/p/11880563.html)

## 安装Flink

### centos7 关闭防火墙
```
sudo systemctl stop firewalld 临时关闭

sudo systemctl disable firewalld ，然后reboot 永久关闭

sudo systemctl status  firewalld 查看防火墙状态。
```
### 安装java8 
下载： https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html
```
创建安装目录
mkdir /usr/local/java/

解压至安装目录
tar -zxvf jdk-8u231-linux-x64.tar.gz -C /usr/local/java/

设置环境变量
打开文件
vi /etc/profile
在末尾添加

export JAVA_HOME=/usr/local/java/jdk1.8.0_231
export JRE_HOME=${JAVA_HOME}/jre
export CLASSPATH=.:${JAVA_HOME}/lib:${JRE_HOME}/lib
export PATH=${JAVA_HOME}/bin:$PATH

使环境变量生效
source /etc/profile

添加软链接
ln -s /usr/local/java/jdk1.8.0_231/bin/java /usr/bin/java

检查
java -version
```

### 安装Flink
下载： https://flink.apache.org/downloads.html
```
tar xzf flink-*.tgz

cd flink-1.9.1

./bin/start-cluster.sh  # Start Flink
```


### Flink Data Sources
https://ci.apache.org/projects/flink/flink-docs-release-1.9/dev/datastream_api.html#data-sources

File-based: readTextFile
Socket-based: socketTextStream 
Collection-based:  fromCollection,fromElements,fromParallelCollection ,generateSequence

Custom: addSource 

https://ci.apache.org/projects/flink/flink-docs-release-1.9/dev/connectors/index.html#bundled-connectors


### 开始第一个例子
```
# yum install -y nc
nc -l -p 9000

./bin/flink run examples/streaming/SocketWindowWordCount.jar --hostname 127.0.0.1 --port 9000

tail -f log/flink-*-taskexecutor-*.out


./bin/stop-cluster.sh
```
