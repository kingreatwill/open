http://spark.apache.org/docs/latest/quick-start.html




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




spark-submit --master spark://192.168.110.216:7077 examples/src/main/python/wordcount.py /root/spark-2.4.5-bin-hadoop2.7/README.md
```











## windows
[Spark启动时的master参数以及Spark的部署方式](https://blog.csdn.net/zpf336/article/details/82152286)
[spark-submit几种提交模式的区别](https://blog.csdn.net/fa124607857/article/details/103390996)

http://spark.apache.org/docs/latest/quick-start.html
### 下载安装
```

setx HADOOP_HOME E:\bigdata\spark-2.4.4-bin-hadoop2.7\
setx SPARK_HOME E:\bigdata\spark-2.4.4-bin-hadoop2.7\

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