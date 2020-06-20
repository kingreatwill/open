<!--toc-->
[TOC]
## Hadoop安装
### 下载hadoop和配置环境变量
```
tar xzf hadoop-3.2.1.tar.gz
mv hadoop-3.2.1 hadoop

vi ~/.bashrc
or
vi /etc/profile
or
vi ~/.bash_profile

source ~/.bashrc


export HADOOP_HOME=/root/bigdata/hadoop
export HADOOP_COMMON_LIB_NATIVE_DIR=$HADOOP_HOME/lib/native
export JAVA_LIBRARY_PATH=$HADOOP_HOME/lib/native
export HADOOP_OPTS="-Djava.library.path=$HADOOP_HOME/lib:$HADOOP_COMMON_LIB_NATIVE_DIR"
export LD_LIBRARY_PATH=$HADOOP_HOME/lib/native/:$LD_LIBRARY_PATH
export PATH=.:${HADOOP_HOME}/bin:$PATH

```
- 如果启动出现以下错误

WARN util.NativeCodeLoader: Unable to load native-hadoop library for your platform... using builtin-java classes where applicable

在hadoop-env.sh中增加配置，如下：
`export LD_LIBRARY_PATH=$HADOOP_HOME/lib/native/:$LD_LIBRARY_PATH`

或者上面vi ~/.bashrc中添加
`export LD_LIBRARY_PATH=$HADOOP_HOME/lib/native/:$LD_LIBRARY_PATH`

下面的就只有`export LD_LIBRARY_PATH=$HADOOP_HOME/lib/native/:$LD_LIBRARY_PATH`有用
```
export HADOOP_COMMON_LIB_NATIVE_DIR=$HADOOP_HOME/lib/native
export JAVA_LIBRARY_PATH=$HADOOP_HOME/lib/native
export HADOOP_OPTS="-Djava.library.path=$HADOOP_HOME/lib:$HADOOP_COMMON_LIB_NATIVE_DIR"
export LD_LIBRARY_PATH=$HADOOP_HOME/lib/native/:$LD_LIBRARY_PATH
```

### 修改配置文件

```
mkdir /root/bigdata/hadoop/data \
mkdir /root/bigdata/hadoop/data/tmp \
mkdir /root/bigdata/hadoop/data/var \
mkdir /root/bigdata/hadoop/data/dfs \
mkdir /root/bigdata/hadoop/data/dfs/name \
mkdir /root/bigdata/hadoop/data/dfs/data
```
进入 hadoop/etc/hadoop文件夹
1. 修改core-site.xml文件
```
<configuration>
        <property>
                <name>hadoop.tmp.dir</name>
                <value>/root/bigdata/hadoop/data/tmp</value>
                <description>hadoop tmp dir</description>
        </property>
        <property>
                <name>fs.default.name</name>
                <value>hdfs://192.168.110.216:9000</value>
        </property>
</configuration>
```

2. 修改 hadoop-env.sh 
```
export JAVA_HOME=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.242.b08-0.el7_7.x86_64
```
```
如果开发java应用，经常需要配置JAVA_HOME路径，如果是通过yum安装的jdk（一般系统会自带open-jdk），下面讲述配置过程：

A 定位JDK安装路径

1. 终端输入：

which java
输出为：

/usr/bin/java
2. 终端输入：

ls -lr /usr/bin/java
输出为：

/usr/bin/java -> /etc/alternatives/java

3. 终端输入

ls -lrt /etc/alternatives/java
输出：

/etc/alternatives/java -> /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.181-3.b13.el7_5.x86_64/jre/bin/java

 

至此，我们确定java的安装目录为： /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.181-3.b13.el7_5.x86_64

 
```

3. 修改hdfs-site.xml
```
<property>
   <name>dfs.name.dir</name>
   <value>/root/bigdata/hadoop/data/dfs/name</value>
   <description>Path on the local filesystem where theNameNode stores the namespace and transactions logs persistently.</description>
</property>
<property>
   <name>dfs.data.dir</name>
   <value>/root/bigdata/hadoop/data/dfs/data</value>
   <description>Comma separated list of paths on the localfilesystem of a DataNode where it should store its blocks.</description>
</property>
<property>
   <name>dfs.replication</name>
   <value>2</value>
</property>
<property>
      <name>dfs.permissions</name>
      <value>true</value>
      <description>need not permissions</description>
</property>
<property>
  <name>dfs.http.address</name>
  <value>0.0.0.0:50070</value>
</property>
```
dfs.permissions配置为false后，可以允许不要检查权限就生成dfs上的文件，方便倒是方便了，但是你需要防止误删除，请将它设置为true，或者直接将该property节点删除，因为默认就是true

4. 修改mapred-site.xml
```
<property>
    <name>mapred.job.tracker</name>
    <value>192.168.110.216:9001</value>
</property>
<property>
      <name>mapred.local.dir</name>
       <value>/root/bigdata/hadoop/data/var</value>
</property>
<property>
       <name>mapreduce.framework.name</name>
       <value>yarn</value>
</property>
```

### 启动hadoop
1. 首次启动，需要初始化（格式化）
```
hadoop namenode -format
```
- 检查：
```
cd /root/bigdata/hadoop/data/dfs/name
ll
ll current
```

2. 启动hadoop：主要是启动HDFS和YARN
```
cd sbin
```
- 在里面修改四个文件(在文件顶部第二行添加)

对于start-dfs.sh和stop-dfs.sh文件，添加下列参数：
```
#!/usr/bin/env bash

HDFS_DATANODE_USER=root
HADOOP_SECURE_DN_USER=hdfs
HDFS_NAMENODE_USER=root
HDFS_SECONDARYNAMENODE_USER=root
```
对于start-yarn.sh和stop-yarn.sh文件，添加下列参数：
```
#!/usr/bin/env bash

YARN_RESOURCEMANAGER_USER=root
HADOOP_SECURE_DN_USER=yarn
YARN_NODEMANAGER_USER=root
```
- 免密登录
```
ssh-keygen -t rsa
ssh-copy-id -i ~/.ssh/id_rsa.pub root@dev07
```
- 启动HDFS
```
start-dfs.sh
```
- 启动YARN
```
start-yarn.sh
```

- 查看
http://192.168.110.216:8088/cluster
http://192.168.110.216:50070

- 参考
https://www.cnblogs.com/ngy0217/p/10538336.html
https://my.oschina.net/xiaoyaoyoufang/blog/1592276

### hadoop操作
http://hadoop.apache.org/docs/r1.0.4/cn/hdfs_shell.html
http://hadoop.apache.org/docs/r3.2.1/hadoop-project-dist/hadoop-common/FileSystemShell.html

hadoop fs -ls /


hadoop fs -mkdir /aaa

hadoop fs -put /root/bigdata/hadoop/README.txt /aaa

hadoop fs -ls /aaa

hadoop fs -cat /aaa/README.txt

hadoop fs -rm /aaa/README.txt

删除某一文件夹下的文件：hdfs dfs -rm -r /aaa/1.py

## 安装Spark
### Spark 下载版本间的区别

#### Pre-built for Apache Hadoop
```
# 0. 进入测试目录
cd /tmp

# 1. 下载Spark 并解压
wget https://archive.apache.org/dist/spark/spark-2.1.1/spark-2.1.1-bin-hadoop2.7.tgz
tar -xf spark-2.1.1-bin-hadoop2.7.tgz

# 2. 运行Spark
cd /tmp/spark-2.1.1-bin-hadoop2.7 &&\
./bin/spark-submit \
--class org.apache.spark.examples.SparkPi \
./examples/jars/spark-examples_2.11-2.1.1.jar \
100
# 程序是可以正常运行的且可以看到
# 类似于Pi is roughly 3.1408123140812316 的字样
```
#### Pre-built with user-provided Apache Hadoop
```
# 0. 进入测试目录
cd /tmp

# 1. 下载Spark 并解压
wget https://archive.apache.org/dist/spark/spark-2.1.1/spark-2.1.1-bin-without-hadoop.tgz
tar -xf spark-2.1.1-bin-without-hadoop.tgz

# 2. 运行Spark
cd /tmp/spark-2.1.1-bin-without-hadoop &&\
./bin/spark-submit \
--class org.apache.spark.examples.SparkPi \
./examples/jars/spark-examples_2.11-2.1.1.jar \
100
# 直接报错，java.lang.ClassNotFoundException: org.apache.hadoop.fs.FSDataInputStream

# 3. 与之前的spark-2.1.1-bin-hadoop2.7.tgz 对比以下jar 包
ls spark-2.1.1-bin-hadoop2.7/jars/ > h.txt
ls spark-2.1.1-bin-without-hadoop/jars/ > w.txt
diff -y -W 50 h.txt w.txt  # -y 并排对比，-W 列宽
# 看到右边的w.txt 内容中少了很多Hadoop 的包

# 4. 下载Hadoop 并解压
cd /tmp && wget https://archive.apache.org/dist/hadoop/common/hadoop-2.7.2/hadoop-2.7.2.tar.gz
tar -xf hadoop-2.7.2.tar.gz

# 5. 确认Hadoop 的版本
/tmp/hadoop-2.7.2/bin/hadoop version

# 6. 关联Spark-Without-Hadoop 和Hadoop
cat > /tmp/spark-2.1.1-bin-without-hadoop/conf/spark-env.sh << 'EOF'
#!/usr/bin/env bash
export SPARK_DIST_CLASSPATH=$(/tmp/hadoop-2.7.2/bin/hadoop classpath)
EOF

# 7. 再次运行Spark
cd /tmp/spark-2.1.1-bin-without-hadoop &&\
./bin/spark-submit \
--class org.apache.spark.examples.SparkPi \
./examples/jars/spark-examples_2.11-2.1.1.jar \
100 2>&1 | grep 'Pi is'
# 成功显示Pi is roughly 3.1424395142439514
```
#### 总结
由上实验可见，[Pre-built with user-provided Apache Hadoop](#Pre-built with user-provided Apache Hadoop)版需要自己修改配置文件去适配Hadoop，实际就是执行时在CLASSPATH中加入Hadoop的Jar包，而[Pre-built for Apache Hadoop](#Pre-built for Apache Hadoop)则是做到了开箱即用，将提前对应的Hadoop Jar包捆绑在其中，同时因为Hadoop 2.6和Hadoop 2.7的HDFS等接口不一样，所以Pre-built for Apache Hadoop分了两个版本。




### 下载和配置环境变量
https://spark.apache.org/docs/latest/hadoop-provided.html
```
tar xzf spark-3.0.0-bin-without-hadoop.tgz
mv spark-3.0.0-bin-without-hadoop spark

vi ~/.bashrc
or
vi /etc/profile
or
vi ~/.bash_profile

source ~/.bashrc


export SPARK_HOME=/root/bigdata/spark
export PATH=$PATH:$SPARK_HOME/bin

# echo $PATH
```
修改conf/spark-env.sh
将spark-env.sh.template改名spark-env.sh
```
#!/usr/bin/env bash

export  SPARK_MASTER_HOST=192.168.110.216
export  SPARK_LOCAL_IP=192.168.110.216
export  SPARK_MASTER_IP=192.168.110.216
export  SPARK_MASTER_PORT=7077
export  SPARK_WORKER_CORES=1
export  SPARK_WORKER_INSTANCES=1
export  PYSPARK_PYTHON=python3

export SPARK_DIST_CLASSPATH=$(hadoop classpath)

```
启动 sbin
start-all.sh

测试
查看端口 netstat -tunlp
http://192.168.110.216:8080
```
spark-submit --master spark://192.168.110.216:7077 examples/src/main/python/wordcount.py file:///root/bigdata/spark/README.md

spark-submit examples/src/main/python/wordcount.py file:///root/bigdata/spark/README.md

# 默认hdfs
spark-submit /root/bigdata/spark/examples/src/main/python/wordcount.py /aaa/README.txt

spark-submit /root/bigdata/spark/examples/src/main/python/wordcount.py hdfs://192.168.110.216:9000/aaa/README.txt
```