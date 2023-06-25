[A Template for PostgreSQL HA with ZooKeeper, etcd or Consul](https://github.com/zalando/patroni)

[德哥的PostgreSQL私房菜 - 史上最屌PG资料合集](https://yq.aliyun.com/articles/59251)

[实战6：PostgreSQL 全文检索功能实战](https://itdashu.com/docs/sqlbase/aff6b/sqlpractice6.html)

- https://github.com/jaiminpan/pg_jieba
- https://github.com/amutu/zhparser

## 为什么使用PG

PostgreSQL带有[许多功能](https://www.postgresql.org/about/featurematrix/)，旨在帮助开发人员构建应用程序，帮助管理员保护数据完整性和构建容错环境，并帮助您管理数据（无论数据集大小）。
除了免费和开源之外，PostgreSQL也是高度可扩展的。
例如，您可以定义自己的数据类型，构建自定义函数，甚至可以使用不同的编程语言编写代码，而无需重新编译数据库！

PostgreSQL尝试符合SQL标准，在这种标准下，该标准不会与传统功能相抵触，也可能导致糟糕的架构决策。 支持SQL标准所需的许多功能，尽管有时语法或功能略有不同。 随着时间的流逝，有望进一步实现一致性。 从2019年10月的版本12开始，PostgreSQL符合179种SQL：2016 Core一致性强制功能中的至少160种。 在撰写本文时，没有关系数据库完全符合此标准。

以下是PostgreSQL中各种功能的详尽列表，每个主要版本中都添加了更多功能：

- Data Types
    - Primitives原始: Integer, Numeric, String, Boolean
    - Structured结构化: Date/Time, Array, Range, UUID
    - Document文档: JSON/JSONB, XML, Key-value (Hstore)
    - Geometry几何: Point, Line, Circle, Polygon
    - Customizations定制: Composite, Custom Types
- Data Integrity 数据的完整性
    - UNIQUE, NOT NULL
    - Primary Keys
    - Foreign Keys
    - Exclusion Constraints 排除约束
    - Explicit Locks, Advisory Locks 显式锁，通知/咨询锁
    update ad_test set info='abc' where id=1 returning pg_try_advisory_xact_lock(id);
    select * from ad_test where id=1 and pg_try_advisory_xact_lock(1);读不到被锁的记录
- Concurrency, Performance 并发，性能
    - Indexing: B-tree, Multicolumn, Expressions, Partial
    - Advanced Indexing: GiST, SP-Gist, KNN Gist, GIN, BRIN, Covering indexes,    - Bloom filters
    - Sophisticated query planner / optimizer, index-only scans, multicolumn  - statistics
    - Transactions, Nested Transactions (via savepoints)
    - Multi-Version concurrency Control (MVCC)
    - Parallelization of read queries and building B-tree indexes
    - Table partitioning
    - All transaction isolation levels defined in the SQL standard, including     - Serializable
    - Just-in-time (JIT) compilation of expressions
- Reliability, Disaster Recovery 可靠性，灾难恢复
    - Write-ahead Logging (WAL)
    - Replication: Asynchronous, Synchronous, Logical
    - Point-in-time-recovery (PITR), active standbys
    - Tablespaces
- Security
    - Authentication: GSSAPI, SSPI, LDAP, SCRAM-SHA-256, Certificate, and more
    - Robust access-control system
    - Column and row-level security
    - Multi-factor authentication with certificates and an additional method
- Extensibility 可扩展性
    - Stored functions and procedures
    - Procedural Languages: PL/PGSQL, Perl, Python (and many more)
    - SQL/JSON path expressions
    - Foreign data wrappers: connect to other databases or streams with a standard SQL interface
    - Customizable storage interface for tables
    - Many extensions that provide additional functionality, including PostGIS
- Internationalisation, Text Search 国际化，文字搜索
    - Support for international character sets, e.g. through ICU collations
    - Case-insensitive and accent-insensitive collations
    - Full-text search


您可以在PostgreSQL文档中发现更多功能。 此外，PostgreSQL具有高度的可扩展性：许多功能（例如索引）都定义了API，因此您可以使用PostgreSQL进行扩展以解决挑战。

事实证明，PostgreSQL既可以管理的数据量又可以容纳的并发用户数都具有很高的可扩展性。 生产环境中有活动的PostgreSQL集群，可以管理许多terabytes(TB)的数据，而专业系统则可以管理petabytes(PB)。

**B-tree or B+tree**
Postgresql: B-tree
只有Key键。索引的重点是存储键。数据位于表中，这些表是逻辑堆。

## Greenplum: 基于PostgreSQL的分布式数据库
https://www.toutiao.com/i6827431495108395532/
https://github.com/greenplum-db/gpdb


## 其它
### UNIQUE 约束中的空值
PostgreSQL 15 开始支持在定义唯一索引或者唯一约束时指定多个 NULL 值是否相同。

在此之前，多个 NULL 值被看作不同的数据，意味着唯一索引和唯一约束中可以存在多个 NULL 值。如果不允许唯一索引和唯一约束中存在多个 NULL 值可以通过 UNIQUE NULLS NOT DISTINCT 进行限制。

默认选项为 UNIQUE NULLS DISTINCT。