

## 简介
Cassandra数据库在分布式的架构上借鉴了Amazon的Dynamo，而在数据的存储模型上参考了Google的Bigtable

https://cassandra.apache.org/doc/latest/architecture/guarantees.html
![](img/cassandra-cap.jpg)
> Cassandra满足可用性和分区容差属性。弱一致性
> HBase满足一致性和分区容差属性。

[Cassandra维护数据一致性的策略](https://www.cnblogs.com/xxdfly/p/5641684.html)

Cassandra是闹过绯闻的，08年 由Facebook开源，然后10年[twitter宣布停用Cassandra](https://timyang.net/data/twitter-cassandra/)，由于Cassandra的最终一致性模型不适合Message类型应用的处理，当时不少公司考虑到它缺少海量并发的案例，所以也不敢将其使用在主要业务上。不过好东西是最经得起考验的，经过这么4,5年的发展，Cassandra有Datastax的推广，目前已经有好多公司在使用Cassandra，如Instagram按照DBEngine上面的数据库排名（2020.06全部数据库排行第十，Wide column排行第一，NoSQL排在MongoDB，Elasticsearch，Redis之后，排行第四），Cassandra已经成为列簇类型的Nosql中用户最广的产品(所以前几年有很多说选择HBase的)。



## 安装
docker pull cassandra:3.11.6

## 应用
Cassandra+Spark vs HBase+Hadoop

## 性能
https://www.datastax.com/products/compare/nosql-performance-benchmarks
遥遥领先


## 功能
CQL -> Cassandra Query Language  https://cassandra.apache.org/doc/latest/cql/
UDF -> User Defined function
TTL
CAS -> 比如更新或者插入时可以加if
batch -> 原子性

## datastax
DataStax | NoSQL Database Built on Apache Cassandra
https://www.datastax.com/
DataStax，是一家位于加州的初创公司，提供了一个商业版本的Apache Cassandra NoSQL数据库

