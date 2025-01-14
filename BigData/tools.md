datax sqoop kettle flume StreamSets

开源数据同步中间件：Canal、Debezium、DataX、Databus、FlinkX、Bifrost

# 数据集成工具
对于数据集成类应用，通常会采用ETL工具辅助完成。ETL，是英文 Extract-Transform-Load 的缩写，用来描述将数据从来源端经过抽取（extract）、交互转换（transform）、加载（load）至目的端的过程。当前的很多应用也存在大量的ELT应用模式。

常见的ETL工具或类ETL的数据集成同步工具很多，以下对开源的 Kettle、Sqoop、Datax、Streamset进行简单梳理比较。
## kettle 4.1k

说明：是国外开源ETL工具，支持数据库、FTP、文件、rest接口、hdfs、hive等平台的灵敏据进行抽取、转换、传输等操作，Java编写跨平台，C/S架构，不支持浏览器模式。

特点：

易用性：有可视化设计器进行可视化操作，使用简单。

功能强大：不仅能进行数据传输，能同时进行数据清洗转换等操作。

支持多种源：支持各种数据库、FTP、文件、rest接口、hdfs、Hive等源。

部署方便：独立部署，不依赖第三方产品。

适用场景：

数据量及增量不大，业务规则变化较快，要求可视化操作，对技术人员的技术门槛要求低。


## sqoop  727

说明：Apache开源软件，主要用于在HADOOP(Hive)与传统的数据库(mysql、postgresql...)间进行数据的传递。

特点：

数据吞吐量大：依赖hadoop集群可进行大批量数据集成。

操作有技术要求：sqoop操作没有可视化设计器，对使用人员有较专业的技术要求。

多种交互方式：命令行，web UI，rest API。

部署不方便：sqoop依赖大数据集群，使用sqoop要求数据传输的的源要与大数据集群的所有节点能进行通信。

适用场景：

适用于能与大数据集群直接通信的关系数据库间的大批量数据传输。


## dataX 6.5k

说明：是阿里开源软件异构数据源离线同步工具，致力于实现包括关系型数据库(MySQL、Oracle等)、HDFS、Hive、ODPS、HBase、FTP等各种异构数据源之间稳定高效的数据同步功能。

特点：

易用性：没有界面，以执行脚本方式运行，对使用人员技术要求较高。

性能：数据抽取性能高。

部署：可独立部署

适用场景：

在异构数据库/文件系统之间高速交换数据。


## flume

说明：flume是一个分布式、可靠、和高可用的海量日志采集、聚合和传输的系统。

特点：

分布式：flume分布式集群部署，扩展性好。

可靠性好: 当节点出现故障时，日志能够被传送到其他节点上而不会丢失。

易用性：flume配置使用较繁琐，对使用人员专业技术要求非常高。

实时采集：flume采集流模式进行数据实时采集。

适用场景：

适用于日志文件实时采集。

## StreamSets
https://github.com/streamsets

StreamSets 数据收集器是一个轻量级，强大的引擎，实时流数据。使用Data Collector在数据流中路由和处理数据。
要为Data Collector定义数据流，请配置管道。一个流水线由代表流水线起点和终点的阶段以及您想要执行的任何附加处理组成。配置管道后，单击“开始”，“ 数据收集器”开始工作。
Data Collector在数据到达原点时处理数据，在不需要时静静地等待。您可以查看有关数据的实时统计信息，在数据通过管道时检查数据，或仔细查看数据快照。

DataOps
Scale across Data Lakes and Data Warehouses