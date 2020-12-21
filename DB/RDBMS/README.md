SQL查询语句的执行顺序解析
https://www.jianshu.com/p/1628e47a598a

[MySQL和PGSQL对比材料](https://blog.csdn.net/qq_45533926/article/details/111328980)
31. 分布式演进
PGSQL
HTAP数据库：cockroachDB、腾讯Tbase
分片集群：  Postgres-XC、Postgres-XL
MySQL
HTAP数据库：TiDB
分片集群： 各种各样的中间件，不一一列举

## 数据库设计规范
1. 表名及字段名必须使用小写字母或者数字（MySql在Windows下不区分大小写，但在linux中是默认区分大小写的，避免节外生枝，不要出现任何大写字母），禁止出现数字开头，更禁止两个下划线之间仅有数字。

2. 表名不使用复数名词（仅用来表示实体的内容，不应表示实体的数量）。表的命名遵循“业务名称_表的作用”进行命名，如trade_config。数据库名与应用名称尽量一致。

3. 表达是否概念时，使用is_xxx的方式命名，数据类型为unsigned tinyint，并且1表示是，0表示否。小数类型为decimal，禁止使用float和double（在存储float和double时存在精度损失问题，可能会导致比较值时，得不到正确的结果）。如果存储的数据范围超过decimal，将数据拆分为整数和小数部分分开存储。

4. varchar 是可变长字符串，不预先分配存储空间，长度不要超过 5000，如果存储长度大于此值，定义字段类型为 text ，独立出来一张表，用主键来对应，避免影响其它字段索引效率。

5. 表必备三字段：id, create_time, update_time。

6. 数据库表之间外键，采用逻辑外键而非物理外键。（采用逻辑外键，以免级联更新造成的低并发问题，级联更新是强阻塞，不适合分布式高并发系统。同时物理外键存在数据库更新风暴的风险，外键也会大大影响数据库的插入速度）

7. 禁止使用存储过程，存储过程难以调试和扩展，没有移植性。


## proxy

### mysql proxy

#### Gaea
Gaea在设计、实现阶段参照了mycat、kingshard和vitess，并使用tidb parser作为内置的sql parser
https://github.com/XiaoMi/Gaea

#### kingshard
https://github.com/flike/kingshard

## 国产数据库

[国产数据库流行度排行](https://www.modb.pro/dbrank)

- 阿里的数据库：OceanBase & PolarDB
- 腾讯的TDSQL
- 华为的GaussDB
- PingCAP 团队开发的TiDB
- 达梦数据库
- GBase南大通用
- 神通数据库
- 金仓数据库
- EsgynDB
- SequoiaDB巨杉数据库
- K-DB数据库
- OpenBASE
- 华易数据库Huayisoft DB Server
- HUABASE-华鼎数据库
- 神州OSCAR
- ZILLIZ数据库

国产的数据库都有哪些？ - 晓平的回答 - 知乎
https://www.zhihu.com/question/31750851/answer/1066968462
国产的数据库都有哪些？ - 李宇明的回答 - 知乎
https://www.zhihu.com/question/31750851/answer/139665111

### OceanBase
分布式架构
OceanBase内核100多万行代码都是蚂蚁完全自主编写
OceanBase是蚂蚁金服阳振坤老师团队开发的，官方称为“完全自主研发的金融级分布式关系数据库”。
OceanBase是完全从头写的，PolarDB应该是基于MySQL开发的。

### openGauss
openGauss是一款开源的关系型数据库管理系统，它具有多核高性能、全链路安全性、智能运维等企业级特性。 openGauss内核早期源自开源数据库PostgreSQL，融合了华为在数据库领域多年的内核经验，在架构、事务、存储引擎、优化器及ARM架构上进行了适配与优化。作为一个开源数据库，期望与广泛的开发者共同构建一个多元化技术的开源数据库社区。

### TiDB