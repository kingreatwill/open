https://geekflare.com/open-source-database/
[TOC]
## 什么是NewSQL

简单的来说：SQL+NoSQL＝NewSQL。 NewSQL被定义为下一代数据库的发展方向。甚至在今天的数据库特性上已经可以看到这个趋势：最流行的开源关系型数据库：MySQL ，从5.7版本开始具有JSON，KV解决方案；而最流行的NoSQL数据库：MongoDB，也推出Join解决方案。

SQL与NoSQL的界限逐渐在模糊，甚至传统关系型数据库PostgreSQL可以利用FDW让MongoDB作为其数据源。所以，不论SQL也好，NoSQL也好，都会以NewSQL的形态展现。

## 现有NewSQL系统厂商举例
Google 第一个搞了 NewSQL （Spanner 和 F1）
包括(顺序随机)Clustrix、GenieDB、ScalArc、Schooner、VoltDB、RethinkDB、ScaleDB、Akiban、CodeFutures、ScaleBase、Translattice和NimbusDB，以及 Drizzle、带有 NDB的 MySQL 集群和带有HandlerSocket的MySQL。后者包括Tokutek和JustOne DB。相关的“NewSQL作为一种服务”类别包括亚马逊关系数据库服务，微软SQLAzure，Xeround和FathomDB。

### 目前NewSQL系统大致分三类：

#### 新架构
第一类型的NewSQL系统是全新的数据库平台，它们均采取了不同的设计方法。它们大概分两类：
(1) 这类数据库工作在一个分布式集群的节点上，其中每个节点拥有一个数据子集。 SQL查询被分成查询片段发送给自己所在的数据的节点上执行。这些数据库可以通过添加额外的节点来线性扩展。现有的这类数据库有： Google Spanner, VoltDB, Clustrix, NuoDB.

(2) 这些数据库系统通常有一个单一的主节点的数据源。它们有一组节点用来做事务处理，这些节点接到特定的SQL查询后，会把它所需的所有数据从主节点上取回来后执行SQL查询，再返回结果。

#### SQL引擎
第二类是高度优化的SQL存储引擎。这些系统提供了MySQL相同的编程接口，但扩展性比内置的引擎InnoDB更好。这类数据库系统有：TokuDB, MemSQL。

#### 透明分片
这类系统提供了分片的中间件层，数据库自动分割在多个节点运行。这类数据库包扩：ScaleBase，dbShards, Scalearc。
