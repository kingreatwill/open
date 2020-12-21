https://github.com/apache/kylin

联机事务处理OLTP（Online Transaction Processing）
联机分析处理OLAP（OnLine Analytical Processing）

Apache Kylin™是一个开源、分布式的大数据分析数据仓库;它被设计为在大数据时代提供OLAP(在线分析处理)能力。通过对Hadoop和Spark上的多维立方体和预计算技术的革新，Kylin能够在数据量不断增长的情况下实现近乎恒定的查询速度。Kylin将查询延迟从几分钟缩短到次秒，将在线分析带回到大数据。


[开源OLAP引擎测评报告(SparkSql、Presto、Impala、HAWQ、ClickHouse、GreenPlum)](https://blog.csdn.net/oDaiLiDong/article/details/86570211)

[分布式关系数据库（OLAP、OLTP）的介绍和比较](https://blog.csdn.net/xuheng8600/article/details/80334971)

[你需要的不是实时数仓 | 你需要的是一款合适且强大的OLAP数据库(上)](https://www.cnblogs.com/importbigdata/p/11521403.html)

[你需要的不是实时数仓 | 你需要的是一款强大的OLAP数据库(下)](https://www.cnblogs.com/importbigdata/p/11521390.html)

数据处理大致可以分成两大类：
 - 联机事务处理OLTP（On-line Transaction Processing）
  OLTP是传统的关系型数据库的主要应用，主要是基本的、日常的事务处理，例如银行交易。
 - 联机分析处理OLAP（On-Line Analytical Processing）
 OLAP是数据仓库系统的主要应用，支持复杂的分析操作，侧重决策支持，并且提供直观易懂的查询结果。 


与Apache Kylin一样致力于解决大数据查询问题的其他开源产品也有不少，比如Apache Drill、Apache Impala、Druid、Doris、ClickHouse、Hive、Presto、SparkSQL等。
目前，主流的OLAP引擎主要是两个套路：一个用空间换时间，一个充分利用所有资源快速计算。
前者就是MOLAP（Multidimensional OLAP，多维在线分析），后者就是ROLAP（Relational OLAP，关系型在线分析）。
Kylin和Druid都是MOLAP的典范，ClickHouse和Doris则是ROLAP的佼佼者。

## OLAP开源引擎
目前市面上主流的开源OLAP引擎包含不限于：Hive、Hawq、Presto、Kylin、Impala、Sparksql、Druid、Clickhouse、Greeplum等，可以说目前没有一个引擎能在数据量，灵活程度和性能上做到完美，用户需要根据自己的需求进行选型。

上面给出了常用的一些OLAP引擎，它们各自有各自的特点，我们将其分组：

- Hive，Hawq，Impala - 基于SQL on Hadoop
- Presto和Spark SQL类似 - 基于内存解析SQL生成执行计划
- Kylin - 用空间换时间，预计算
- Druid - 一个支持数据的实时摄入
- ClickHouse - OLAP领域的Hbase，单表查询性能优势巨大
- Greenpulm - OLAP领域的Postgresql

如果你的场景是基于HDFS的离线计算任务，那么Hive，Hawq和Imapla就是你的调研目标；
如果你的场景解决分布式查询问题，有一定的实时性要求，那么Presto和SparkSQL可能更符合你的期望；
如果你的汇总维度比较固定，实时性要求较高，可以通过用户配置的维度+指标进行预计算，那么不妨尝试Kylin和Druid；
ClickHouse则在单表查询性能上独领风骚，远超过其他的OLAP数据库；
Greenpulm作为关系型数据库产品，性能可以随着集群的扩展线性增长，更加适合进行数据分析。


SQLServer内置的OLAP工具Analysis Manager可以允许用户访问异构数据库,提供支持OLAP分析的高速缓存和计算引擎
[SQL Server Analysis Services OLAP 引擎服务器组件](https://docs.microsoft.com/zh-cn/analysis-services/multidimensional-models/olap-physical/olap-engine-server-components?view=asallproducts-allversions)

[通过内存中 OLTP 使用查询存储](https://docs.microsoft.com/zh-cn/sql/relational-databases/performance/using-the-query-store-with-in-memory-oltp?view=sql-server-ver15)

- Apache Pinot
Pinot 是一个实时分布式的 OLAP 数据存储和分析系统。
使用它实现低延迟可伸缩的实时分析。
Pinot 从脱机数据源（包括 Hadoop 和各类文件）和在线数据源（如 Kafka）中获取数据进行分析。
Pinot 被设计成可进行水平扩展。
Pinot 特别适合这样的数据分析场景：查询具有大量维度和指标的时间序列数据、分析模型固定、数据只追加以及低延迟，以及分析结果可查询。

- Apache Doris
[Apache Kylin VS Apache Doris全方位对比](https://cloud.tencent.com/developer/article/1477234)