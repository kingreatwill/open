# Change Data Capture (CDC)
更改数据捕获

CDC的全称是Change Data Capture，即变更数据捕获，它是数据库领域非常常见的技术，主要用于捕获数据库的一些变更，然后可以把变更数据发送到下游。它的应用比较广，可以做一些数据同步、数据分发和数据采集，还可以做ETL。

对于CDC，业界主要有两种类型：
- 一是基于查询的，客户端会通过SQL方式查询源库表变更数据，然后对外发送。
- 二是基于日志，这也是业界广泛使用的一种方式，一般是通过binlog方式，变更的记录会写入binlog，解析binlog后会写入消息系统，或直接基于Flink CDC进行处理。

它们两者是有区别的，基于查询比较简单，是入侵性的，而基于日志是非侵入性，对数据源没有影响，但binlog的解析比较复杂一些。

业界有很多开源的binlog的解析器，比较通用和流行的有[Debezium](https://github.com/debezium/debezium)、Canal，以及Maxwell。基于这些binlog解析器就可以构建ETL管道。

Debezium是一个为了捕获数据变更(CDC)的开源分布式平台。Debezium记录的是数据库表行级别的变更事件，包括 insert/update/delete 等等操作。同时debezium是构建在kafka之上的，与kafka深度耦合，所以提供了Debezium Connector。
支持的数据库有MySQL、MongoDB、PostgreSQL、Oracle、SQL server

## flink-cdc-connectors
https://github.com/ververica/flink-cdc-connectors

支持MySQL和PostgreSQL