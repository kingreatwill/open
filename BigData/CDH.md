## Hadoop痛点
一个产品的出现肯定是为了解决用户的痛点，在大数据领域，我们这些使用Hadoop、Hive、Hbase等的开发者来说就是其用户。如果使用原生的ApacheHadoop，在工作中我总结出了如下痛点（部分）：

1、集群规模很庞大时搭建Hadoop集群复杂度越来越高，工作量很大

2、规模很大的集群下升级Hadoop版本很费时费力

3、需要自己保证版本兼容，比如升级Hadoop版本后需要自己保证与Hive、Hbase等的兼容

4、安全性很低

## Hadoop发行版本引入
有了上述原生Hadoop不足，Apache官方和一些第三方就发布了一些Hadoop发行版本来解决此类问题。一些有名的发行版本列举如下：

• Apache Hadoop

• Cloudera’s Distribution Including Apache Hadoop（CDH）

• Hortonworks Data Platform (HDP)

• MapR

• EMR