http://spark.apache.org/docs/latest/quick-start.html




## linux
```
tar xzf spark-*.tgz

cd spark-2.4.4-bin-hadoop2.7
```
在~/.bashrc文件中添加如下内容，并执行$ source ~/.bashrc命令使其生效
```
# export HADOOP_HOME=/root/spark-2.4.4-bin-hadoop2.7
export SPARK_HOME=/root/spark-2.4.4-bin-hadoop2.7
export PATH=$PATH:$SPARK_HOME/bin:$SPARK_HOME/sbin
```











## windows
### 下载安装
```

setx HADOOP_HOME E:\bigdata\spark-2.4.4-bin-hadoop2.7\
setx SPARK_HOME E:\bigdata\spark-2.4.4-bin-hadoop2.7\

%SPARK_HOME%\bin\spark-submit --version








%SPARK_HOME%\bin\run-example SparkPi

%SPARK_HOME%\bin\spark-submit examples/src/main/python/pi.py
```

### 交互环境
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