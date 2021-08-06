

## CDC
CDC适用的环境：

1. SQL server 2008版本以上的企业版、开发版和评估版中可用；

2. 需要开启代理服务（作业）。

3. CDC需要业务库之外的额外的磁盘空间。

4. CDC的表需要主键或者唯一主键。

开启cdc的源表在插入INSERT、更新UPDATE和删除DELETE活动时会插入数据到日志表中。cdc通过捕获进程将变更数据捕获到变更表中，通过cdc提供的查询函数，我们可以捕获这部分数据。 

CDC的表不能truncate操作，truncate是物理删除数据不能捕获变更的数据。

### Tools
https://github.com/Vanlightly/CDC-Tools
https://github.com/VenuMeda/kafka-connect-cdc-mssql

[SQL Server monitoring - Is It SQL?](https://www.scalesql.com/isitsql/)
