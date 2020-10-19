
## Presto
2013年Facebook正式宣布开源Presto。
Presto的两大分支: PrestoDB和PrestoSQL
Presto不是数据库，定位是查询引擎，可对250PB以上的数据进行快速地交互式分析。据称该引擎的性能是 Hive 的 10 倍以上。

通过分布式查询，Presto不仅仅可以访问HDFS，也可以访问其他数据源，数据库。(Presto 不使用 MapReduce ，只需要 HDFS)
Presto 可以查询包括 Hive、Cassandra 甚至是一些商业的数据存储产品。单个 Presto 查询可合并来自多个数据源的数据进行统一分析。
但是，Presto并不能用来处理在线事务
Presto被设计为数据仓库和数据分析产品：数据分析、大规模数据聚集和生成报表。这些工作经常通常被认为是线上分析处理操作


### PrestoDB
Presto是一个用于大数据的分布式SQL查询引擎。
https://github.com/prestodb/presto

### PrestoSQL
用于大数据的分布式SQL查询引擎
https://github.com/prestosql/presto