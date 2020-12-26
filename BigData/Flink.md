https://github.com/alibaba/Alink

[仅1年GitHub Star数翻倍，Flink 做了什么？](https://www.jianshu.com/p/3e83f76b8f4f)
[在 Flink 算子中使用多线程如何保证不丢数据？](https://developer.aliyun.com/article/740572)

[Flink入门（一）——Apache Flink介绍](https://www.cnblogs.com/tree1123/p/11880563.html)

[Flink 1.11](https://ci.apache.org/projects/flink/flink-docs-master/release-notes/flink-1.11.html)

Apache Flink 1.9.0 版本则会开启新的 ML 接口和新的 flink-python 模块

[实时计算](https://www.zhihu.com/topic/19876621/hot)


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

# http://localhost:8081

```
- bin:可执行脚本
- conf:配置文件目录
- lib:依赖jar
- opt:扩展依赖jar

.
- Flink 如何在 K8s 的 POD 中与 HDFS 交互？

其与 HDFS 交互很简单，只要把相关依赖打到镜像里面就行了。就是说你把 flink-shaded-hadoop-2-uber-{hadoopVersion}-{flinkVersion}.jar 放到 flink-home/lib 目录下，然后把一些 hadoop 的配置比如 hdfs-site.xml、 core-site.xml 等放到可以访问的目录下，Flink 自然而然就可以访问了。

**这其实和在一个非 HDFS 集群的节点上，要去访问 HDFS 是一样的。**

- Flink on K8s 和 Flink on YARN，哪个方案更优？怎样选择？

Flink on YARN 是目前比较成熟的一套系统，但是它有点重，不是云原生（cloud native）。在服务上云的大趋势下，Flink on K8s 是一个光明的未来。Flink on YARN 是一个过去非常成熟一套体系，但是它在新的需求、新的挑战之下，可能缺乏一些应对措施。例如对很多细致的 GPU 调度，pipeline 的创建等等，概念上没有 K8s 做得好。

如果你只是简单运行一个作业，在 Flink on YARN 上可以一直稳定运行下去，它也比较成熟，相比之下 Flink on K8s 够新、够潮、方便迭代。不过目前 Flink on K8s 已知的一些问题，比如学习曲线比较陡峭，需要一个很好的 K8s 运维团队去支撑。另外，K8s 本身虚拟化带来的性能影响，正如先前介绍的无论是磁盘，还是网络，很难避免一定的性能损耗，这个可能是稍微有点劣势的地方，当然相比这些劣势，虚拟化（容器化）带来的优点更明显。

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
