https://db-engines.com/en/ranking

## 事务

[分布式事务](https://www.ixigua.com/pseries/6753473359675130379_6733436598798516750/)

## 跨源

数据虚拟化（data virtualization）是用来描述所有数据管理方法的涵盖性术语，这些方法允许应用程序检索并管理数据，且不需要数据相关的技术细节。

- 阿里云的 【数据管理 DMS】通过 DBLink 实现 [跨数据库查询](https://help.aliyun.com/document_detail/99941.html?spm=a2c4g.11186623.6.581.49e97e53uziNCp)

- SparkSQL

- 华为的数据虚拟化引擎[OpenLookeng](https://openlookeng.io/zh-cn/docs/docs/overview/use-cases.html)
  openLooKeng 不是为处理联机事务处理（OLTP）而设计，openLooKeng 被设计用来处理数据仓库和分析：数据分析、聚合大量数据并生成报告。这些工作负载通常归类为联机分析处理（OLAP）。

- SQL Server 数据虚拟化 (Polybase) [配置 PolyBase 以访问 Hadoop 中的外部数据](https://docs.microsoft.com/zh-cn/sql/relational-databases/polybase/polybase-configure-hadoop?view=sql-server-ver15)

  [PolyBase 与 链接服务器 比较](https://docs.microsoft.com/zh-cn/sql/relational-databases/polybase/polybase-faq?view=sql-server-ver15)

- PostgreSQL 外部数据源支持(Foreign Data Wrapper，FDW)
  可以把 70 种外部数据源 (包括 Mysql, Oracle, CSV, hadoop …) 当成自己数据库中的表来查询。Postgres 有一个针对这一难题的解决方案：一个名为“外部数据封装器(Foreign Data Wrapper，FDW)”的特性。该特性最初由 PostgreSQL 社区领袖 Dave Page 四年前根据 SQL 标准 SQL/MED(SQL Management of External Data)开发。FDW 提供了一个 SQL 接口，用于访问远程数据存储中的远程大数据对象，使 DBA 可以整合来自不相关数据源的数据，将它们存入 Postgres 数据库中的一个公共模型。这样，DBA 就可以访问和操作其它系统管理的数据，就像在本地 Postgres 表中一样。例如，使用 FDW for MongoDB，数据库管理员可以查询来自文档数据库的数据，并使用 SQL 将它与来自本地 Postgres 表的数据相关联。借助这种方法，用户可以将数据作为行、列或 JSON 文档进行查看、排序和分组。他们甚至可以直接从 Postgres 向源文档数据库写入(插入、更细或删除)数据，就像一个一体的无缝部署。也可以对 Hadoop 集群或 MySQL 部署做同样的事。FDW 使 Postgres 可以充当企业的中央联合数据库或“Hub”。

## OLAP

联机事务处理 OLTP（Online Transaction Processing）
联机分析处理 OLAP（OnLine Analytical Processing）

[开源 OLAP 引擎测评报告(SparkSql、Presto、Impala、HAWQ、ClickHouse、GreenPlum)](https://blog.csdn.net/oDaiLiDong/article/details/86570211)

[分布式关系数据库（OLAP、OLTP）的介绍和比较](https://blog.csdn.net/xuheng8600/article/details/80334971)

[你需要的不是实时数仓 | 你需要的是一款合适且强大的 OLAP 数据库(上)](https://www.cnblogs.com/importbigdata/p/11521403.html)

[你需要的不是实时数仓 | 你需要的是一款强大的 OLAP 数据库(下)](https://www.cnblogs.com/importbigdata/p/11521390.html)

数据处理大致可以分成两大类：

- 联机事务处理 OLTP（On-line Transaction Processing）
  OLTP 是传统的关系型数据库的主要应用，主要是基本的、日常的事务处理，例如银行交易。
- 联机分析处理 OLAP（On-Line Analytical Processing）
  OLAP 是数据仓库系统的主要应用，支持复杂的分析操作，侧重决策支持，并且提供直观易懂的查询结果。

### OLAP 开源引擎

目前市面上主流的开源 OLAP 引擎包含不限于：Hive、Hawq、Presto、Kylin、Impala、Sparksql、Druid、Clickhouse、Greeplum 等，可以说目前没有一个引擎能在数据量，灵活程度和性能上做到完美，用户需要根据自己的需求进行选型。

上面给出了常用的一些 OLAP 引擎，它们各自有各自的特点，我们将其分组：

- Hive，Hawq，Impala - 基于 SQL on Hadoop
- Presto 和 Spark SQL 类似 - 基于内存解析 SQL 生成执行计划
- Kylin - 用空间换时间，预计算
- Druid - 一个支持数据的实时摄入
- ClickHouse - OLAP 领域的 Hbase，单表查询性能优势巨大
- Greenpulm - OLAP 领域的 Postgresql

如果你的场景是基于 HDFS 的离线计算任务，那么 Hive，Hawq 和 Imapla 就是你的调研目标；
如果你的场景解决分布式查询问题，有一定的实时性要求，那么 Presto 和 SparkSQL 可能更符合你的期望；
如果你的汇总维度比较固定，实时性要求较高，可以通过用户配置的维度+指标进行预计算，那么不妨尝试 Kylin 和 Druid；
ClickHouse 则在单表查询性能上独领风骚，远超过其他的 OLAP 数据库；
Greenpulm 作为关系型数据库产品，性能可以随着集群的扩展线性增长，更加适合进行数据分析。

SQLServer 内置的 OLAP 工具 Analysis Manager 可以允许用户访问异构数据库,提供支持 OLAP 分析的高速缓存和计算引擎
[SQL Server Analysis Services OLAP 引擎服务器组件](https://docs.microsoft.com/zh-cn/analysis-services/multidimensional-models/olap-physical/olap-engine-server-components?view=asallproducts-allversions)

[通过内存中 OLTP 使用查询存储](https://docs.microsoft.com/zh-cn/sql/relational-databases/performance/using-the-query-store-with-in-memory-oltp?view=sql-server-ver15)

#### Apache Pinot

Pinot 是一个实时分布式的 OLAP 数据存储和分析系统。
使用它实现低延迟可伸缩的实时分析。
Pinot 从脱机数据源（包括 Hadoop 和各类文件）和在线数据源（如 Kafka）中获取数据进行分析。
Pinot 被设计成可进行水平扩展。
Pinot 特别适合这样的数据分析场景：查询具有大量维度和指标的时间序列数据、分析模型固定、数据只追加以及低延迟，以及分析结果可查询。

#### Apache Kylin

Apache Kylin™ 是一个开源、分布式的大数据分析数据仓库;它被设计为在大数据时代提供 OLAP(在线分析处理)能力。通过对 Hadoop 和 Spark 上的多维立方体和预计算技术的革新，Kylin 能够在数据量不断增长的情况下实现近乎恒定的查询速度。Kylin 将查询延迟从几分钟缩短到次秒，将在线分析带回到大数据。

#### Apache Doris（原 Palo）

https://github.com/apache/incubator-doris
[Apache Kylin VS Apache Doris 全方位对比](https://cloud.tencent.com/developer/article/1477234)
Doris 是一个 MPP 的 OLAP 系统，主要整合了 Google Mesa（数据模型），Apache Impala（MPP Query Engine)和 Apache ORCFile (存储格式，编码和压缩) 的技术。

Apache Doris 的分布式架构非常简洁，易于运维，并且可以支持 10PB 以上的超大数据集。

Apache Doris 可以满足多种数据分析需求，例如固定历史报表，实时数据分析，交互式数据分析和探索式数据分析等。使得数据分析工作更加简单高效！

## 其它

### 数据库性能测试工具

#### SysBench

https://github.com/akopytov/sysbench c 3.5k

支持系统：

- Linux
- macOS
- Windows

目前 sysbench 主要支持 MySQL,pgsql,oracle

SysBench 是一个模块化的、跨平台、多线程基准测试工具，主要用于评估测试各种不同系统参数下的数据库负载情况。它主要包括以下几种方式的测试：
1、cpu 性能
2、磁盘 io 性能
3、调度程序性能
4、内存分配及传输速度
5、POSIX 线程性能
6、数据库性能(OLTP 基准测试)

#### dbbench

https://github.com/memsql/dbbench

- mysql
- mssql
- pgsql

### 数据库同步

#### canal

#### Debezium

https://github.com/debezium/debezium 3.8k

可以同步数据到 kafka

Debezium 是一个开源项目，为捕获数据更改(Capture Data Change，CDC)提供了一个低延迟的流式处理平台，通过安装配置 Debezium 监控数据库，可以实时消费行级别(row-level)的更改。身为一个分布式系统，Debezium 也拥有良好的容错性。
Debezium 是一种借助 Kafka 将数据变更发布成事件流的 CDC 实现。
Debezium 是一款开源的、基于 Kafka 的 CDC 工具，它会读取数据库事务日志，并将其发布成事件流。

CDC 除了可以用来更新缓存、服务和搜索引擎，Morling 还介绍了其他几种用例，包括：

- 数据复制，通常用来将数据复制到其他类型的数据库或数据仓库中。
- 审计。因为保留了数据历史，在使用元数据填充数据后，可以实现数据变更审计。

Debezium 的源端(即支持监控哪些数据库) :

- MySQL
- MongoDB
- PostgreSQL
- Oracle
- SQL Server
- Oracle (Incubating)
- Db2 (Incubating)
- Cassandra (Incubating)
  [数据库连接器](https://debezium.io/docs/connectors/)

Debezium 的目标端(即可以数据导入端) : Kafka

[Debezium 获取 MySQL Binlog](https://my.oschina.net/jerval/blog/3058959)
[使用嵌入式 Debezium 和 SpringBoot 捕获更改数据事件（CDC） - Sohan Ganapathy](https://www.jdon.com/53411)

如果您已经安装了 Zookeeper、Kafka 和 Kafka Connect，那么使用 Debezium 的连接器是很容易的。只需下载一个或多个连接器插件存档(见下文)，将它们的文件解压到 Kafka Connect 环境中，并将解压后的插件的父目录添加到 Kafka Connect 的插件路径中。如果不是这样，在你的工作配置中指定插件路径(例如，connect- distribu. properties)使用插件。路径配置属性。例如，假设您已经下载了 Debezium MySQL 连接器存档，并将其内容解压缩到/kafka/connect/ Debezium -connector- MySQL。然后在 worker 配置中指定以下内容:`plugin.path=/kafka/connect`

### 数据湖方案

- ACID 和隔离级别支持

| Solution   | ACID Support | Isolation Level                                            | Concurrent Multi-Writers | Time Travel |
| ---------- | ------------ | ---------------------------------------------------------- | ------------------------ | ----------- |
| Iceberg    | Yes          | Write Serialization                                        | Yes                      | Yes         |
| Hudi       | Yes          | Snapshot Isolation                                         | Yes                      | Yes         |
| Open Delta | Yes          | Serialization<br>Write Serialization<br>Snapshot Isolation | Yes                      | Yes         |
| Hive ACID  | Yes          | Snapshot Isolation                                         | Yes                      | No          |

三种隔离分别代表的含义:

- Serialization 是说所有的 reader 和 writer 都必须串行执行；
- Write Serialization: 是说多个 writer 必须严格串行，reader 和 writer 之间则可以同时跑；
- Snapshot Isolation: 是说如果多个 writer 写的数据无交集，则可以并发执行；否则只能串行。Reader 和 writer 可以同时跑。

- Schema 变更支持和设计

| Solution   | Schema Evolution | Self-defined schema object |
| ---------- | ---------------- | -------------------------- |
| Iceberg    | all              | Yes                        |
| Hudi       | back-compatible  | No(spark-schema)           |
| Open Delta | all              | No(spark-schema)           |
| Hive ACID  | all              | No(Hive-schema)            |

iceberg 是做的比较好的，抽象了自己的 schema，不绑定任何计算引擎层面的 schema。

- 流批接口支持
  目前 Iceberg 和 Hive 暂时不支持流式消费，不过 Iceberg 社区正在 issue 179 上开发支持。

- 接口抽象程度和插件化

| Solution   | Engine Pluggable<br/>(Write Path) | Engine Pluggable<br/>(Read Path) | Storage Pluggable <br/>(Less Storage API Binding) | Open File Format    |
| ---------- | --------------------------------- | -------------------------------- | ------------------------------------------------- | ------------------- |
| Iceberg    | Yes                               | Yes                              | Yes                                               | Yes                 |
| Hudi       | No(Bind with spark)               | Yes                              | Yes                                               | Yes(data) + No(Log) |
| Open Delta | No(Bind with spark)               | Yes                              | Yes                                               | Yes                 |
| Hive ACID  | Yes                               | Yes                              | No                                                | No(Only ORC)        |

- 查询性能优化

| Solution    | Filter PushDown | Low meta cast | Indexing within partitions <br/>Boost the perf of selective queries | CopyOnWrite | MergeOnRead | Auto-Compaction |
| ----------- | --------------- | ------------- | ------------------------------------------------------------------- | ----------- | ----------- | --------------- |
| Iceberg     | Yes             | Yes           | Road-map                                                            | Yes         | On-going    | No              |
| Hudi        | No              | Yes           | -                                                                   | Yes         | Yes         | Yes             |
| Open Delta  | No              | Yes           | -                                                                   | Yes         | No          | No              |
| Close Delta | Yes             | Yes           | Yes                                                                 | Yes         | Yes         | Yes             |
| Hive ACID   | Yes             | No            | -                                                                   | No          | Yes         | Yes             |

- 其他功能

| Solution    | One line demo | Python Support | File Encryption | Cli Command |
| ----------- | ------------- | -------------- | --------------- | ----------- |
| Iceberg     | Not Good      | Yes            | Yes             | No          |
| Hudi        | Medium        | No             | No              | Yes         |
| Open Delta  | Good          | Yes            | No              | Yes         |
| Close Delta | Good          | Yes            | No              | Yes         |
| Hive ACID   | Medium        | No             | No              | Yes         |

- 社区现状（截止到 2020-11-05）

| Solution   | Open Source Time | Github Star | Github Fork | Github Issues | Contributors |
| ---------- | ---------------- | ----------- | ----------- | ------------- | ------------ |
| Iceberg    | 2018/11/06       | 805         | 312         | 237           | 106          |
| Hudi       | 2019/01/17       | 1.5k        | 638         | 52            | 122          |
| Open Delta | 2019/04/12       | 2.9k        | 635         | 131           | 76           |

总结成如下图：
![](../img/delta-hudi-iceberg.jpg)

#### Delta

Databricks 的 [Delta](https://github.com/delta-io/delta)
开源的 delta 是 databricks 闭源 delta 的一个简化版本

#### Apache Iceberg

Netflix 的 [Apache Iceberg](https://github.com/apache/iceberg)

#### Apache Hudi

Uber 的 [Apache Hudi](https://github.com/apache/hudi)
