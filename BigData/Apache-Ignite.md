# Apache Ignite—不仅仅是数据网格
https://www.toutiao.com/i6827512913318642187/

[Apache Ignite详解 - 很好的文章](https://blog.csdn.net/nihaoa50/article/details/88424569)

## 数据网格
Apache Ignite是一种水平可伸缩，容错的分布式内存计算平台，用于构建实时应用程序，该应用程序可以以内存速度处理数TB的数据。

Apache Ignite是一个水平可伸缩、容错的分布式内存计算平台，用于构建能够以内存速度处理tb级数据的实时应用程序。

## ignite是什么?
apache ignite为开发人员提供了实时处理大数据和内存计算的方便易用的解决方案。主要有以下几点功能：

- Data grid 数据网格
- Compute grid 计算网格
- Service grid 服务网格
- Bigdata accelerator 大数据加数器
- Streaming grid 数据流网格

主要功能如下所示：

- Elasticity 弹性：集群可以通过添加节点进行水平扩展；
- Persistence 持久性：数据网格可以将缓存中的数据持久化到关系型数据库中，甚至是NoSQL数据库中，例如MongoDB或Cassandra；
- Cache as a Service(CaaS) 缓存即服务：允许跨组织、多应用去访问管理内存缓存而不是慢速的基于磁盘读写的数据库；
- 2nd Level Cache 二级缓存：可以作为Hibernate和MyBatis持久化框架的二级缓存层使用；
- 高性能hadoop加速器：apache ignite可以替代hadoop task tracker、job tracker和HDFS，从而提高大数据分析的性能；
- 在Spark应用中共享内存：ignite RDD允许在不同的Spark作业和应用之间轻松的共享状态；
- 分布式计算：apache ignite提供了一组简单的API，允许用户在多个节点上获得高性能的分布计算和处理数据的能力。ignite的分布式服务对于开发和执行微服务架构也会提供很多帮助。
- 流：apache ignite允许可伸缩和容错内存中处理连续不断的数据流，而不是在数据存储在数据库后分析数据。


