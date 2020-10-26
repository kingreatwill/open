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

- Apache Pinot
  Pinot 是一个实时分布式的 OLAP 数据存储和分析系统。
  使用它实现低延迟可伸缩的实时分析。
  Pinot 从脱机数据源（包括 Hadoop 和各类文件）和在线数据源（如 Kafka）中获取数据进行分析。
  Pinot 被设计成可进行水平扩展。
  Pinot 特别适合这样的数据分析场景：查询具有大量维度和指标的时间序列数据、分析模型固定、数据只追加以及低延迟，以及分析结果可查询。

- Apache Kylin
  Apache Kylin™ 是一个开源、分布式的大数据分析数据仓库;它被设计为在大数据时代提供 OLAP(在线分析处理)能力。通过对 Hadoop 和 Spark 上的多维立方体和预计算技术的革新，Kylin 能够在数据量不断增长的情况下实现近乎恒定的查询速度。Kylin 将查询延迟从几分钟缩短到次秒，将在线分析带回到大数据。

## 其它

### SysBench 性能测试工具

https://github.com/akopytov/sysbench c 3.5k

支持系统：

- Linux
- macOS
- Windows

目前sysbench主要支持 MySQL,pgsql,oracle 

SysBench 是一个模块化的、跨平台、多线程基准测试工具，主要用于评估测试各种不同系统参数下的数据库负载情况。它主要包括以下几种方式的测试：
1、cpu 性能
2、磁盘 io 性能
3、调度程序性能
4、内存分配及传输速度
5、POSIX 线程性能
6、数据库性能(OLTP 基准测试)
