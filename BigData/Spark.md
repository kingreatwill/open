
## 2.Spark邮件列表
2.1 邮件列表清单
如果想要进一步跟踪问题、获取最新资源、调试bug、或者贡献代码给Spark项目组，邮件列表是一个非常好的方式。邮件列表也是有多种方式，需要区分每一个邮件类型，订阅你关心的邮件。Apache下面的每一个项目都有自己的邮件列表，同时分不同的邮件组，Apache Spark有如下订阅列表

user@spark.apache.org  订阅该邮件可以参与讨论普通用户遇到的问题
dev-subscribe@spark.apache.org   订阅该邮件可以参与讨论开发者遇到的问题，开发者比较常用这个邮件列表
issues-subscribe@spark.apache.org 订阅该邮件可以收到所有jira的创建和更新
commits-subscribe@spark.apache.org 所有的代码的提交变动信息都会发到该邮件

给上列邮箱发送邮件



## linux
```
tar xzf spark-*.tgz

cd spark-2.4.5-bin-hadoop2.7
```
or source .bash_profile 
在~/.bashrc文件中添加如下内容，并执行$ source ~/.bashrc命令使其生效
```
# export HADOOP_HOME=/root/spark-2.4.5-bin-hadoop2.7
export SPARK_HOME=/root/spark-2.4.5-bin-hadoop2.7
export PATH=$PATH:$SPARK_HOME/bin:$SPARK_HOME/sbin

spark-env.sh

#!/usr/bin/env bash

export  SPARK_MASTER_HOST=192.168.110.216
export  SPARK_LOCAL_IP=192.168.110.216
export  SPARK_MASTER_IP=192.168.110.216
export  SPARK_MASTER_PORT=7077
export  SPARK_WORKER_CORES=1
export  SPARK_WORKER_INSTANCES=1
export  PYSPARK_PYTHON=python3



spark-submit --master spark://192.168.110.216:7077 examples/src/main/python/wordcount.py /root/spark-2.4.5-bin-hadoop2.7/README.md
```











## windows
[Spark启动时的master参数以及Spark的部署方式](https://blog.csdn.net/zpf336/article/details/82152286)
[spark-submit几种提交模式的区别](https://blog.csdn.net/fa124607857/article/details/103390996)

http://spark.apache.org/docs/latest/quick-start.html
### 下载安装
```

setx HADOOP_HOME  E:\bigdata\hadoop-3.2.1\
setx SPARK_HOME E:\bigdata\spark-3.0.0-bin-hadoop3.2\

https://github.com/cdarlint/winutils
下载winutils.exe放入 E:\bigdata\hadoop-3.2.1\bin中
path 添加%SPARK_HOME%\bin

%SPARK_HOME%\bin\spark-submit --version








%SPARK_HOME%\bin\run-example SparkPi  # 可选参数10

%SPARK_HOME%\bin\spark-submit examples/src/main/python/pi.py

# http://spark.apache.org/docs/latest/submitting-applications.html#master-urls
# spark-submit --class Test --master spark://localhost:7077 /home/data/myjar/Hello.jar

set SPARK_LOCAL_IP=192.168.1.216
set SPARK_MASTER_HOST=192.168.1.216
%SPARK_HOME%\bin\spark-submit --master spark://192.168.110.216:7077 examples/src/main/python/pi.py 10
```

```
bin/spark-submit --master spark://master.hadoop:7077 --class nuc.sw.test.ScalaWordCount spark-1.0-SNAPSHOT.jar hdfs://master.hadoop:9000/spark/input/a.txt hdfs://master.hadoop:9000/spark/output
```

### 交互环境
 交互环境的默认UI http://localhost:4040/
```
cd E:\bigdata\spark-2.4.4-bin-hadoop2.7

# python
%SPARK_HOME%\bin\pyspark

>>> textFile = spark.read.text("README.md")
>>> textFile.count() # Number of rows in this DataFrame
105
>>> textFile.first() # First row in this DataFrame
Row(value='# Apache Spark')
>>> linesWithSpark = textFile.filter(textFile.value.contains("Spark"))
>>> textFile.filter(textFile.value.contains("Spark")).count() # How many lines contain "Spark"?
20

>>> sc.parallelize(range(1000)).count() 
1000

# Scala
%SPARK_HOME%\bin\spark-shell

# PYSPARK_DRIVER_PYTHON设置为ipython后，pyspark交互模式变为ipython模式
```


### demo
```
# Use spark-submit to run your application
$ YOUR_SPARK_HOME/bin/spark-submit \
  --class "SimpleApp" \
  --master local[4] \
  target/simple-project-1.0.jar


# Use spark-submit to run your application
$ YOUR_SPARK_HOME/bin/spark-submit \
  --master local[4] \
  SimpleApp.py
```