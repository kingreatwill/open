# 简单了解下Kafka Connect和Kafka Stream

更准确地说，Connect的作用是Kafka与其他数据源(比如：文件，其它MQ或者数据库)之间的导入导出。目前已经支持的数据源可以在这里找到：[Kafka Connect | Confluent](https://www.confluent.io/hub)


数据源（比如数据库，前端Web Server，传感器..）－> Kafka -> Storm / Spark -> 数据接收（比如Elastic，HDFS／HBase，Cassandra，数据仓库..）

那这个架构是完全可以用Kafka Connect ＋ Kafka Streams，也就是：
数据源 －> Kafka Connect －> Kafka －> Kafka Streams －> Kafka －> Kafka Connect －> 数据接收

## KSQL 
Streaming SQL for Apache Kafka
Confluent KSQL是一种流SQL引擎，可针对ApacheKafka®进行实时数据处理。 它为Kafka上的流处理提供了易于使用但功能强大的交互式SQL界面，而无需使用Java或Python之类的编程语言编写代码。 KSQL具有可伸缩性，弹性，容错能力，并且支持多种流操作，包括数据过滤，转换，聚合，联接，窗口和会话化。

## Kafka Connect
https://blog.csdn.net/tianyeshiye/article/details/97751436
http://kafka.apache.org/documentation/#connect
https://docs.confluent.io/current/connect/userguide.html


## Kafka Stream
http://kafka.apache.org/documentation/streams/
### Kafka Stream简介
目前支持实时流计算的有一些框架，如： Spark （micro batch），Storm， Flink，Samza；这些系统都需要部署集群。Kafka stream是实时流计算库，依靠现有的kafka集群，提供分布式的、高容错的、抽象DSL的实时流计算实现：
- 除了Kafka Stream Client lib以外无外部依赖，轻量的嵌入到Java应用中。
- 严格区分Event time（数据产生的时间）和Process Time（系统处理的时间），可处理乱序到达的数据。
- 支持本地状态故障转移，以实现非常高效的有状态操作，如join和window函数。
- 提供必要的流处理原语、hige-level Stream DSL和low-level Processor API。

### Kafka Stream和其它流计算框架的对比

一般的流计算框架用job或者topology来描述任务，它们关注的是：
- 对于提交的任务，高效的将机器资源分配给任务执行。
- 任务必须打包自己的业务实现代码（包括依赖包，配置等）提交到集群，集群负责部署到工作节点。
- 管理任务的执行情况，保证任务之间的资源隔离性。

kafka stream更关注于问题本身：
- 当有新的实例加入执行或者某个执行实例挂掉，就会触发工作实例的balance（原理和kafka consumer的balance一样）。只需要启动一个java进程就可以加入计算任务，kafka stream的实例很适合用容器部署。
- kafka stream的输入和输出都是kafka的topic，同时可维护本地状态（key-value 的表，默认是以 RocksDB实现）支持aggregations和join的计算。状态结果以topic作为Write-ahead logging，所以即使某个实例失败，状态也可以在其它实例恢复。一般实时计算的结果如果要查询，需要将结果数据写入到外部系统（如HBase，Redis），但基于kafka stream的本地状态存储，[可以直接向各个实例查询](https://www.confluent.io/blog/unifying-stream-processing-and-interactive-queries-in-apache-kafka/) ，节省数据的再落地成本。

### 什么时候选用Kafka Stream

当你的环境中已经存在Kafka作为数据入口，后面当然也可以接 Spark、Storm、Flink等实时流处理框架。如果你的实时流计算只是做一些数据转换清洗、按key聚合计算、需要读取多个topic的数据join，不妨先看下kafka stream的流式计算支持，它能以更轻便的方式实现计算需求。如果你的数据来源涉及DB，HDFS，计算过程中涉及机器学习、NoSQL， 则需要考虑 Spark、Storm、Flink。



