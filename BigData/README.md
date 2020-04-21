https://github.com/heibaiying/BigData-Notes


https://blog.csdn.net/qq_32641659/article/details/98729045


https://www.kaggle.com/

框架 | 语言|当前版本(2019/12/19)|star
--|--|--|--
[Apache Hadoop](https://github.com/apache/hadoop) |java| v2.10|10k
[Apache Spark](https://github.com/apache/spark)|scala，java|v3.0.0|24.6k
[Apache Flink](https://github.com/apache/flink) |java，scala |V1.9.1|11.5k
[Apache Storm](https://github.com/apache/storm)|java |V2.1.0|6k

Hadoop(第一代) - > Spark(第二代) - > Flink(第三代)

[Spark](https://spark.apache.org/) 生态更为丰富,在机器学习整合方面投入更多
[Flink](https://flink.apache.org/) 在流计算上有明显优势，核心架构和模型也更透彻和灵活一些。

总而言之，Flink与Spark没有谁强谁弱，只有哪个更适合当前的场景。

Spark与Flink API情况
API| Spark | Flink
--|--|--
底层API | RDD |Process Function
核心API  | DataFrame/DataSet/Structured Streaming | DataStream/DataSet
SQL | Spark SQL | Table API & SQL
机器学习 | MLlib | FlinkML
图计算 | GraphX | Gelly
其它 |  | CEP



Spark与Flink 对开发语言的支持

支持语言| Spark | Flink
--|--|--
Java|√|√
Scala|√|√
Python|√|√
R|√|第三方
SQL|√|√

Flink VS Spark 之 Connectors
https://ci.apache.org/projects/flink/flink-docs-release-1.9/zh/dev/connectors/
https://spark.apache.org/

支持| Spark | Flink
--|--|--
Hbase           |√  |√
HDFS            |√  |√
PostgreSQL      |√  |
Mysql           |√  |
Elastic         |√  |√
kafka           |√  |√
redis           |√  |
cassandra       |√  |√
mongoDB         |√  |
Alluxio         |√  |
Hive            |√  |
hundreds of other data sources|√  |
RabbitMQ        |   |√
Amazon Kinesis Streams|   |√
Apache NiFi     |   |√
Twitter Streaming API |   |√





Flink VS Spark 之 运行环境

部署环境| Spark | Flink
--|--|--
Local(single JVM) | √|√
Standalone |√|√
Yarn |√|√
Mesos |√|√
Kubernetes |√|√

Flink 流处理
Storm 流处理
Spark 微批处理
Hadoop 批处理
