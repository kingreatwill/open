
https://www.tutorialspoint.com/spark_sql/spark_sql_introduction.htm
https://data-flair.training/blogs/spark-sql-tutorial/


## 架构图 Architecture 
![](img/spark_sql_architecture.svg)
SparkSQL架构分成三个部分，第一部分是前端的，第二部分是后端的，第三个部分是中间的Catalyst，这个Catalyst是整个架构的核心。

https://0x90e.github.io/spark-sql-architecture/
## 函数
http://spark.apache.org/docs/latest/api/sql/index.html

## 执行计划
scala> spark.sql("EXPLAIN FORMATTED" + query).shaow(false)

spark.sql("SELECT age FROM emp where age>25 order by age").explain()
可以调用explain(true)方法查看逻辑和物理执行计划