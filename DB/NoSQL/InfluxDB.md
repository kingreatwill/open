# influxdb
![](img/InfluxDB-2.webp)
常见 TSDB(时序数据库)：influxdb、opentsdb、timeScaladb、Druid 等。

https://github.com/influxdata/influxdb
https://www.influxdata.com/
https://v2.docs.influxdata.com/v2.0/get-started/

## influxdb概念

influxdb是一个开源分布式时序、时间和指标数据库，使用 Go 语言编写，无需外部依赖。其设计目标是实现分布式和水平伸缩扩展，是 InfluxData 的核心产品。

应用：性能监控，应用程序指标，物联网传感器数据和实时分析等的后端存储。

适合新增而不修改，甚至是不关心很久以前的数据


influxdb 中的概念	 |传统数据库中的概念
---|---
database	|数据库
measurement	|数据库中的表
point	|表中的一行数据

point的数据结构由时间戳（time）、标签（tags）、数据（fields）三部分组成，具体含义如下：


point 属性	|含义
---|---
time	|数据记录的时间，是主索引（自动生成）
tags	|各种有索引的属性
fields	|各种value值（没有索引的属性）