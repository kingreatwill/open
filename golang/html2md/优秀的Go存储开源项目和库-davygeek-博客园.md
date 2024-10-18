可以看到，今年谷歌家的 Go 编程语言流行度有着惊人的上升趋势，其发展也是越来越好，因此本文整理了一些优秀的 Go 存储相关开源项目和库，一起分享，一起学习。

## **存储服务器（Storage Server）**

### **Go 实现的存储服务器**

- **[minio](https://github.com/minio/minio)**\- Minio 是一个与 Amazon S3 APIs 兼容的开源对象存储服务器，分布式存储方案
- **[rclone](https://github.com/ncw/rclone)** \- “用于云存储的 Rsync” - Google Drive, Amazon Drive, S3, Dropbox, Backblaze B2, One Drive, Swift, Hubic, Cloudfile…
- **[camlistore](https://github.com/camlistore/camlistore)** \- Camlistore 是你的个人存储系统：一种存储、同步、共享、建模和备份内容的方式
- **[torus](https://github.com/coreos/torus)** \- CoreOS 的现代分布式存储系统
- **[s3git](https://github.com/s3git/s3git)** \- 云存储的 Git。用于数据的分布式版本控制系统
- **[rook](https://github.com/rook/rook)** \- 开放、云本地和通用的分布式存储

## **Key-Value 存储（Key-Value Store）**

### **Go 实现的 Key-Value 存储**

- **[etcd](https://github.com/coreos/etcd)** \- 可靠的分布式 key-value 存储，用于分布式系统的最关键数据
- **[go-cache](https://github.com/patrickmn/go-cache)** \- Go 语言实现的一个内存中的缓存框架，实现 Key-Value 的序列存储，适用于单台机器应用程序
- **[biscuit](https://github.com/dcoker/biscuit)** \- Biscuit 用于 AWS 基础架构建设时多区域 HA key-value 存储
- **[diskv](https://github.com/peterbourgon/diskv)** \- 支持磁盘的 key-value 存储

## **文件系统（File System）**

### **Go 实现的文件系统**

- **[git-lfs](https://github.com/git-lfs/git-lfs)** \- 用于大文件版本控制的 Git 扩展
- **[seaweedfs](https://github.com/chrislusf/seaweedfs)** \- SeaweedFS 是一个用于小文件的简单且高度可扩展的分布式文件系统
- **[fsnotify](https://github.com/fsnotify/fsnotify)** \- Go 实现的跨平台文件系统监控库
- **[goofys](https://github.com/kahing/goofys)** \- Go 实现的高性能，POSIX-ish Amazon S3 文件系统
- **[go-systemd](https://github.com/coreos/go-systemd)** \- systemd 的 Go 语言绑定版（包括socket activation, journal, D-Bus, 和 unit files）
- **[gcsfuse](https://github.com/GoogleCloudPlatform/gcsfuse)** \- 用于与 Google 云存储交互的用户空间文件系统
- **[svfs](https://github.com/ovh/svfs)** \- 基于 Openstack 的虚拟文件系统

## **数据库（Database）**

### **Go 实现的数据库**

- **[BigCache](https://github.com/allegro/bigcache)** \- 用于千兆字节数据的高效 key/value 缓存
- **[bolt](https://github.com/boltdb/bolt)** \- Go 实现的低层级的 key/value 数据库
- **[buntdb](https://github.com/tidwall/buntdb)** \- 一个 Go 实现的快速、可嵌入的 key/value 内存数据库，具有自定义索引和 geospatial 支持的功能
- **[cache2go](https://github.com/muesli/cache2go)** \- key/value 内存缓存，支持基于超时的自动无效功能
- **[cockroach](https://github.com/cockroachdb/cockroach)** \- 一个可伸缩的、支持地理位置处理、支持事务处理的数据存储系统
- **[couchcache](https://github.com/codingsince1985/couchcache)** \- 由 Couchbase 服务器支持的 RESTful 缓存微服务
- **[dgraph](https://github.com/dgraph-io/dgraph)** \- 具有可扩展、分布式、低延迟和高吞吐量功能的图形数据库
- **[eliasdb](https://github.com/krotik/eliasdb)** \- 使用 REST API，短语搜索和类似 SQL 查询语言的无依赖性，支持事务处理的图形数据库
- **[forestdb](https://github.com/couchbase/goforestdb)** \- Go bindings for ForestDB.Go 语言绑定版的 ForestDB
- **[GCache](https://github.com/bluele/gcache)** \- 支持可用缓存、LFU、LRU 和 ARC 的缓存数据库
- **[geocache](https://github.com/melihmucuk/geocache)** \- An in-memory cache that is suitable for geolocation based applications.适用于 地理位置处理基于应用程序的内存缓存
- **[goleveldb](https://github.com/syndtr/goleveldb)** \- An implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in the Go.Go 实现的 LevelDB key/value 数据库
- **[groupcache](https://github.com/golang/groupcache)** \- Groupcache 是一个缓存和缓存填充库，在许多情况下用于替代 memcached
- **[influxdb](https://github.com/influxdb/influxdb)** \- 开源的分布式指标、事件和实时分析的可扩展数据库
- **[ledisdb](https://github.com/siddontang/ledisdb)** \- 基于 LevelDB 类似 Redis 的高性能 NoSQL 数据库
- **[levigo](https://github.com/jmhodges/levigo)** \- 用于 LevelDB 的 Go 封装包
- **[moss](https://github.com/couchbase/moss)** \- Go 实现的简单 LSM key-value 存储引擎
- **[piladb](https://github.com/fern4lvarez/piladb)** \- 基于堆栈数据结构的轻量级 RESTful 数据库引擎
- **[pREST](https://github.com/nuveo/prest)** \- 为任何来自 PostgreSQL 的数据库提供一个 RESTful API
- **[prometheus](https://github.com/prometheus/prometheus)** \- 服务监控系统和时间序列数据库
- **[rqlite](https://github.com/rqlite/rqlite)** \- 基于 SQLite 构建的轻量级、分布式关系数据库
- **[scribble](https://github.com/nanobox-io/golang-scribble)** \- 一个小型的 Flat File JSON 存储
- **[tidb](https://www.oschina.net/p/tidb)** \- TiDB 是一个分布式 SQL 数据库，灵感来自于 Google F1 和 Google spanner。TiDB 支持包括传统 RDBMS 和 NoSQL 的特性。
- **[tiedot](https://github.com/HouzuoGuo/tiedot)** \- 基于 Go 的 NoSQL 数据库
- **[Tile38](https://github.com/tidwall/tile38)** \- 具有空间索引和实时地理围栏的地理位置数据库

### **数据库** **迁移**

- **[darwin](https://github.com/GuiaBolso/darwin)** \- Go 实现的数据库 schema 演进库
- **[goose](https://github.com/steinbacher/goose)** \- 数据库迁移工具。可通过创建增量 SQL 或 Go 脚本来管理数据库的演变
- **[gormigrate](https://github.com/go-gormigrate/gormigrate)** \- Gorm ORM 的数据库迁移助手
- **[migrate](https://github.com/mattes/migrate)** \- Go 实现的数据库迁移处理，支持 MySQL, PostgreSQL, Cassandra, 和 SQLite
- **[pravasan](https://github.com/pravasan/pravasan)** \- 简单的迁移工具，目前支持 MySQL，PostgreSQL，但计划很快支持 SQLite, MongoDB 等
- **[soda](https://github.com/markbates/pop/tree/master/soda)** \- 具有数据库迁移、创建和 ORM 等功能，适用于 MySQL, PostgreSQL, 和 SQLite
- **[sql-migrate](https://www.oschina.net/p/sql-migrate)** \- 数据库 schema 迁移工具。允许使用 go-bindata 将迁移嵌入到应用程序中

### **数据库工具**

- **[go-mysql](https://www.oschina.net/p/go-mysql)** \- Go 实现的用于处理 MySQL 协议和复制的工具集
- **[go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch)** \- 将 MySQL 数据自动同步到 Elasticsearch 中
- **[kingshard](https://www.oschina.net/p/kingshard)** \- Go 实现的高性能 MySQL Proxy 项目
- **[myreplication](https://github.com/2tvenom/myreplication)** \- MySQL 二进制日志复制监听器。支持语句和基于行的复制
- **[orchestrator](https://www.oschina.net/p/orchestrator)** \- MySQL 复制拓扑管理器和可视化工具
- **[pgweb](https://www.oschina.net/p/pgweb)** \- Go 实现的基于 Web 的 PostgreSQL 数据库管理系统
- **[vitess](https://www.oschina.net/p/vitess)** \- 分布式 MySQL 工具集。vitess 提供了服务器和工具，以便于大规模 Web 服务的 MySQL 数据库扩展

### **SQL 查询构建器，用于构建和使用 SQL 的库**

- **[dat](https://github.com/mgutz/dat)** \- Go 实现的 Postgres 数据访问工具包
- **[Dotsql](https://github.com/gchaincl/dotsql)** \- Go 语言实现的库，可帮助你将 sql 文件保存至某个地方并轻松使用它
- **[goqu](https://github.com/doug-martin/goqu)** \- Go 实现的 SQL 构建器和查询库
- **[igor](https://github.com/galeone/igor)** \- PostgreSQL 的抽象层，支持高级功能并使用类似 Gorm 的语法
- **[ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)** \- 强大的数据检索方法以及 DB-agnostic 查询构建功能
- **[scaneo](https://github.com/variadico/scaneo)** \- 生成 Go 代码以将数据库行转换为任意结构
- **[sqrl](https://github.com/elgris/sqrl)** \- SQL 查询构建器，Squirrel 的 fork 具有更好的性能
- **[Squirrel](https://github.com/Masterminds/squirrel)** \- 帮助你构建 SQL 查询的 Go 库
- **[xo](https://github.com/knq/xo)** \- 基于现有 schema 定义或支持 PostgreSQL，MySQL，SQLite，Oracle 和 Microsoft SQL Server 的自定义查询生成数据库的惯用 Go 代码

## **数据库驱动**

**用于连接和操作数据库的库**

### **关系数据库**

- **[bgc](https://github.com/viant/bgc)** \- Go 实现的用于 BigQuery 的数据存储连接
- **[firebirdsql](https://github.com/nakagami/firebirdsql)** \- Firebird RDBMS SQL 驱动
- **[go-adodb](https://github.com/mattn/go-adodb)** \- Microsoft ActiveX Object 数据库驱动，使用 database/sql
- **[go-bqstreamer](https://github.com/rounds/go-bqstreamer)** \- BigQuery 快速并发流插入
- **[go-mssqldb](https://github.com/denisenkom/go-mssqldb)** \- Microsoft MSSQL 驱动
- **[go-oci8](https://github.com/mattn/go-oci8)** \- Oracle 驱动，使用 database/sql
- **[go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)** \- MySQL 驱动
- **[go-sqlite3](https://github.com/mattn/go-sqlite3)** \- SQLite3 驱动，使用 database/sql
- **[gofreetds](https://github.com/minus5/gofreetds)** \- Microsoft MSSQL 驱动。Go wrapper over [FreeTDS](http://www.freetds.org/).
- **[pgx](https://github.com/jackc/pgx)** \- PostgreSQL 驱动
- **[pq](https://github.com/lib/pq)** \- Go 实现的用于 database/sql 的 Postgres 驱动

### **NoSQL 数据库**

- **[aerospike-client-go](https://github.com/aerospike/aerospike-client-go)** \- Go 实现的 Aerospike 客户端
- **[arangolite](https://github.com/solher/arangolite)** \- Go 实现的 ArangoDB 轻量级驱动程序
- **[asc](https://github.com/viant/asc)** \- 用于 Aerospike 的数据存储连接
- **[cayley](https://github.com/google/cayley)** \- 支持多个后端的图形数据库
- **[dsc](https://github.com/viant/dsc)** \- 用于 SQL, NoSQL 以及结构化文件的数据存储连接
- **[dynago](https://github.com/underarmour/dynago)** \- DynamoDB 的客户端
- **[go-couchbase](https://github.com/couchbase/go-couchbase)** \- Go 实现的 Couchbase 客户端
- **[go-couchdb](https://github.com/fjl/go-couchdb)** \- Go 实现的 CouchDB HTTP API 封装包
- **[gocb](https://github.com/couchbase/gocb)** \- 官方的 Couchbase Go SDK 包
- **[gocql](http://gocql.github.io/)** \- Go 实现的 Apache Cassandra 驱动
- **[gomemcache](https://github.com/bradfitz/gomemcache/)** \- memcache 客户端库
- **[gorethink](https://github.com/dancannon/gorethink)** \- RethinkDB 驱动
- **[goriak](https://github.com/zegl/goriak)** \- Riak KV 驱动
- **[mgo](https://godoc.org/labix.org/v2/mgo)** \- MongoDB 驱动，它根据标准 Go 习惯用法在非常简单的 API 下实现丰富且经过良好测试的功能选择
- **[neo4j](https://github.com/cihangir/neo4j)** \- Neo4j Rest API 绑定
- **[Neo4j-GO](https://github.com/davemeehan/Neo4j-GO)** \- Neo4j REST 客户端
- **[neoism](https://github.com/jmcvetta/neoism)** \- Neo4j client 客户端
- **[redigo](https://github.com/garyburd/redigo)** \- Redis 数据库客户端
- **[redis](https://github.com/go-redis/redis)** \- Redis 客户端
- **[redis](https://github.com/hoisie/redis)** \- 简单强大的 Redis 客户端
- **[redis](https://github.com/bsm/redeo)** \- Redis 协议兼容 TCP servers/services

### **搜索和分析数据库**

- **[bleve](https://github.com/blevesearch/bleve)** \- 现代文本索引库
- **[elastic](https://github.com/olivere/elastic)** \- Elasticsearch 客户端
- **[elastigo](https://github.com/mattbaird/elastigo)** \- Elasticsearch 客户端库
- **[goes](https://github.com/belogik/goes)** \- 与 Elasticsearch 交互的库
- **[skizze](https://github.com/seiflotfy/skizze)** \- A probabilistic data-structures service and storage.数据结构服务和存储
