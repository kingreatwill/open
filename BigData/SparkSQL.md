
https://www.tutorialspoint.com/spark_sql/spark_sql_introduction.htm
https://data-flair.training/blogs/spark-sql-tutorial/


## 跨源

数据虚拟化（data virtualization）是用来描述所有数据管理方法的涵盖性术语，这些方法允许应用程序检索并管理数据，且不需要数据相关的技术细节。


- 阿里云的 【数据管理 DMS】通过DBLink 实现 [跨数据库查询](https://help.aliyun.com/document_detail/99941.html?spm=a2c4g.11186623.6.581.49e97e53uziNCp) 

- SparkSQL 

- 华为的数据虚拟化引擎[OpenLookeng](https://openlookeng.io/zh-cn/docs/docs/overview/use-cases.html)
openLooKeng不是为处理联机事务处理（OLTP）而设计，openLooKeng被设计用来处理数据仓库和分析：数据分析、聚合大量数据并生成报告。这些工作负载通常归类为联机分析处理（OLAP）。

- SQL Server 数据虚拟化 (Polybase)  [配置 PolyBase 以访问 Hadoop 中的外部数据](https://docs.microsoft.com/zh-cn/sql/relational-databases/polybase/polybase-configure-hadoop?view=sql-server-ver15)

    [PolyBase 与 链接服务器 比较](https://docs.microsoft.com/zh-cn/sql/relational-databases/polybase/polybase-faq?view=sql-server-ver15)

## 架构图 Architecture 
![](img/spark_sql_architecture.svg)
SparkSQL架构分成三个部分，第一部分是前端的，第二部分是后端的，第三个部分是中间的Catalyst，这个Catalyst是整个架构的核心。

![](img/spark-sql-AQE.jpg)

https://0x90e.github.io/spark-sql-architecture/
## 函数
http://spark.apache.org/docs/latest/api/sql/index.html

## 执行计划

```
>>> spark.sql("EXPLAIN FORMATTED" + query).show()


EXPLAIN [FORMATTED|EXTENDED|DEPENDENCY|AUTHORIZATION] hql_query

FORMATTED：对执行计划进行格式化，返回JSON格式的执行计划
EXTENDED：提供一些额外的信息，比如文件的路径信息

下面两个测试报错
DEPENDENCY：以JSON格式返回查询所依赖的表和分区的列表
AUTHORIZATION：列出需要被授权的条目，包括输入与输出
```
spark.sql("SELECT age FROM emp where age>25 order by age").explain()
可以调用explain(True)方法查看逻辑和物理执行计划

## 参考
[二十八、SparkSQL入门](https://www.toutiao.com/i6846994501806850568)