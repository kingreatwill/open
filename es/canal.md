

[TOC]
## 名词介绍
### ETL
数据的抽取（Extract），数据的转换（Transform），加载（Load）
ETL是将业务系统的数据经过抽取、清洗转换之后加载到数据仓库的过程，目的是将企业中的分散、零乱、标准不统一的数据整合到一起，为企业的决策提供分析依据。

## QuickStart

查看mysql服务配置文件的位置
```
which mysqld

/usr/bin/mysqld --verbose --help | grep -A 1 'Default options'
```
查看binlog位置
```
cat /etc/my.cnf

# log_bin = bin.log

获取binlog文件列表
mysql> show binary logs;
```
### 准备
- 对于自建 MySQL , 需要先开启 Binlog 写入功能，配置 binlog-format 为 ROW 模式，my.cnf 中配置如下
```
[mysqld]
log-bin=mysql-bin # 开启 binlog
binlog-format=ROW # 选择 ROW 模式
server_id=1 # 配置 MySQL replaction 需要定义，不要和 canal 的 slaveId 重复
```
重启mysql
```
/etc/init.d/mysqld restart
```


- 授权 canal 链接 MySQL 账号具有作为 MySQL slave 的权限, 如果已有账户可直接 grant
```
CREATE USER canal IDENTIFIED BY 'canal';  
GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'canal'@'%';
-- GRANT ALL PRIVILEGES ON *.* TO 'canal'@'%' ;
FLUSH PRIVILEGES;
```

### 下载并运行 canal-server
```
# adapter ClientAdapter
# admin 管理界面
# deployer canal-server


wget https://github.com/alibaba/canal/releases/download/canal-1.1.4/canal.adapter-1.1.4.tar.gz
wget https://github.com/alibaba/canal/releases/download/canal-1.1.4/canal.admin-1.1.4.tar.gz
wget https://github.com/alibaba/canal/releases/download/canal-1.1.4/canal.deployer-1.1.4.tar.gz
wget https://github.com/alibaba/canal/releases/download/canal-1.1.4/canal.example-1.1.4.tar.gz

mkdir canal-server

tar zxvf canal.deployer-1.1.4.tar.gz  -C ~/canal-server

cd canal-server

```
配置修改
```
vi conf/example/instance.properties
```
```
## mysql serverId
canal.instance.mysql.slaveId = 1234
#position info，需要改成自己的数据库信息
canal.instance.master.address = 127.0.0.1:3306 
canal.instance.master.journal.name = 
canal.instance.master.position = 
canal.instance.master.timestamp = 
#canal.instance.standby.address = 
#canal.instance.standby.journal.name =
#canal.instance.standby.position = 
#canal.instance.standby.timestamp = 
#username/password，需要改成自己的数据库信息
canal.instance.dbUsername = canal  
canal.instance.dbPassword = canal
canal.instance.defaultDatabaseName =
canal.instance.connectionCharset = UTF-8
#table regex
canal.instance.filter.regex = .\*\\\\..\*
```
canal.instance.connectionCharset 代表数据库的编码方式对应到 java 中的编码类型，比如 UTF-8，GBK , ISO-8859-1
如果系统是1个 cpu，需要将 canal.instance.parser.parallel 设置为 false

#### canal.instance.filter.regex的书写格式
```
mysql 数据解析关注的表，Perl正则表达式.
多个正则之间以逗号(,)分隔，转义符需要双斜杠(\\) 
常见例子：
1.  所有表：.*   or  .*\\..*
2.  canal schema下所有表： canal\\..*
3.  canal下的以canal打头的表：canal\\.canal.*
4.  canal schema下的一张表：canal.test1
5.  多个规则组合使用：canal\\..*,mysql.test1,mysql.test2 (逗号分隔)
```
**特别注意**
检查下CanalConnector是否调用subscribe(filter)方法；有的话，filter需要和instance.properties的canal.instance.filter.regex一致，否则subscribe的filter会覆盖instance的配置，如果subscribe的filter是.*\..*，那么相当于你消费了所有的更新数据 

启动
```
sh bin/startup.sh

sh bin/startup.sh local 指定配置

# 查看 server 日志
cat logs/canal/canal.log

# 查看 instance 的日志
cat logs/example/example.log

# 关闭
sh bin/stop.sh
```

## 发送消息到MQ
https://github.com/alibaba/canal/wiki/Canal-Kafka-RocketMQ-QuickStart
默认支持将canal server接收到的binlog数据直接投递到MQ(先安装[zk](https://github.com/alibaba/canal/wiki/Zookeeper-QuickStart)), 目前默认支持的MQ系统有:
- [kafka](https://github.com/alibaba/canal/wiki/Kafka-QuickStart)
- [RocketMQ](https://rocketmq.apache.org/docs/quick-start/) 

### mq顺序性问题
binlog本身是有序的

1. canal目前选择支持的kafka/rocketmq，本质上都是基于本地文件的方式来支持了分区级的顺序消息的能力，也就是binlog写入mq是可以有一些顺序性保障，这个取决于用户的一些参数选择
2. canal支持MQ数据的几种路由方式：单topic单分区，单topic多分区、多topic单分区、多topic多分区
  - canal.mq.dynamicTopic，主要控制是否是单topic还是多topic，针对命中条件的表可以发到表名对应的topic、库名对应的topic、默认topic name
  - canal.mq.partitionsNum、canal.mq.partitionHash，主要控制是否多分区以及分区的partition的路由计算，针对命中条件的可以做到按表级做分区、pk级做分区等
3. canal的消费顺序性，主要取决于描述2中的路由选择，举例说明：
  - 单topic单分区，可以严格保证和binlog一样的顺序性，缺点就是性能比较慢，单分区的性能写入大概在2~3k的TPS
  - 多topic单分区，可以保证表级别的顺序性，一张表或者一个库的所有数据都写入到一个topic的单分区中，可以保证有序性，针对热点表也存在写入分区的性能问题
  - 单topic、多topic的多分区，如果用户选择的是指定table的方式，那和第二部分一样，保障的是表级别的顺序性(存在热点表写入分区的性能问题)，如果用户选择的是指定pk hash的方式，那只能保障的是一个pk的多次binlog顺序性 ** pk hash的方式需要业务权衡，这里性能会最好，但如果业务上有pk变更或者对多pk数据有顺序性依赖，就会产生业务处理错乱的情况. 如果有pk变更，pk变更前和变更后的值会落在不同的分区里，业务消费就会有先后顺序的问题，需要注意

  ## Prometheus监控

  http://localhost:11112/metrics

  https://github.com/alibaba/canal/wiki/Prometheus-QuickStart#canal%E7%9B%91%E6%8E%A7%E7%9B%B8%E5%85%B3%E5%8E%9F%E5%A7%8B%E6%8C%87%E6%A0%87%E5%88%97%E8%A1%A8

  ## canal监控相关原始指标列表：

|  指标                         | 说明            |  单位  | 精度 |
| :----                        | :----           | ----: |----: |
|canal_instance_transactions |instance接收transactions计数|-|-|
|canal_instance |instance基本信息|-|-|
|canal_instance_subscriptions |instance订阅数量|-|-|
|canal_instance_publish_blocking_time |instance dump线程提交到异步解析队列过程中的阻塞时间(仅parallel解析模式)|ms|ns|
|canal_instance_received_binlog_bytes |instance接收binlog字节数|byte|-|
|canal_instance_parser_mode |instance解析模式(是否开启parallel解析)|-|-|
|canal_instance_client_packets |instance client请求次数的计数|-|-|
|canal_instance_client_bytes|向instance client发送数据包字节计数|byte|-|
|canal_instance_client_empty_batches|向instance client发送get接口的空结果计数|-|-|
|canal_instance_client_request_error|instance client请求失败计数|-|-|
|canal_instance_client_request_latency |instance client请求的响应时间概况|-|-|
|canal_instance_sink_blocking_time |instance sink线程put数据至store的阻塞时间|ms|ns|
|canal_instance_store_produce_seq |instance store接收到的events sequence number|-|-|
|canal_instance_store_consume_seq |instance store成功消费的events sequence number|-|-|
|canal_instance_store |instance store基本信息|-|-|
|canal_instance_store_produce_mem |instance store接收到的所有events占用内存总量|byte|-|
|canal_instance_store_consume_mem |instance store成功消费的所有events占用内存总量|byte|-|
|canal_instance_put_rows|store put操作完成的table rows|-|-|
|canal_instance_get_rows|client get请求返回的table rows|-|-|
|canal_instance_ack_rows|client ack操作释放的table rows|-|-|
|canal_instance_traffic_delay |server与MySQL master的延时|ms|ms|
|canal_instance_put_delay|store put操作events的延时|ms|ms|
|canal_instance_get_delay|client get请求返回events的延时|ms|ms|
|canal_instance_ack_delay|client ack操作释放events的延时|ms|ms|

## 监控展示指标
|  指标                         | 简述            |多指标|
| :----                        | :----           | :----: |
|[Basic](https://github.com/alibaba/canal/wiki/Canal-prometheus#%E7%8A%B6%E6%80%81%E4%BF%A1%E6%81%AF)|Canal instance 基本信息。|是|
|[Network bandwith](https://github.com/alibaba/canal/wiki/Canal-prometheus#%E7%BD%91%E7%BB%9C%E5%B8%A6%E5%AE%BDkbs)|网络带宽。包含inbound(canal server读取binlog的网络带宽)和outbound(canal server返回给canal client的网络带宽)|是|
|[Delay](https://github.com/alibaba/canal/wiki/Canal-prometheus#delayseconds)|Canal server与master延时；store 的put, get, ack操作对应的延时。|是|
|[Blocking](https://github.com/alibaba/canal/wiki/Canal-prometheus#blocking)|sink线程blocking占比；dump线程blocking占比(仅parallel mode)。|是|
|[TPS(transaction)](https://github.com/alibaba/canal/wiki/Canal-prometheus#tpsmysql-transaction)|Canal instance 处理binlog的TPS，以MySQL transaction为单位计算。|否|
|[TPS(tableRows)](https://github.com/alibaba/canal/wiki/Canal-prometheus#tpstable-row)|分别对应store的put, get, ack操作针对数据表变更行的TPS|是|
|[Client requests](https://github.com/alibaba/canal/wiki/Canal-prometheus#client-requests)|Canal client请求server的请求数统计，结果按请求类型分类(比如get/ack/sub/rollback等)。|否|
|[Response time](https://github.com/alibaba/canal/wiki/Canal-prometheus#response-time)|Canal client请求server的响应时间统计。|否|
|[Empty packets](https://github.com/alibaba/canal/wiki/Canal-prometheus#empty-packets)|Canal client请求server返回空结果的统计。|是|
|[Store remain events](https://github.com/alibaba/canal/wiki/Canal-prometheus#event-store%E5%8D%A0%E7%94%A8)|Canal instance ringbuffer中堆积的events数量。|否|
|[Store remain mem](https://github.com/alibaba/canal/wiki/Canal-prometheus#event-store-memory%E5%8D%A0%E7%94%A8kb-%E4%BB%85memory-mode)|Canal instance ringbuffer中堆积的events内存使用量。|否|
|[Client QPS](https://github.com/alibaba/canal/wiki/Canal-prometheus#client-qps)|client发送请求的QPS，按GET与CLIENTACK分类统计|是|

## JVM 相关信息
> The Java client includes collectors for garbage collection, memory pools, JMX, classloading, and thread counts. These can be added individually or just use the DefaultExports to conveniently register them.
> >DefaultExports.initialize();

详见：[prometheus/client_java](https://github.com/prometheus/client_java)


## ClientAdapter
### 基本说明
canal 1.1.1版本之后, 增加客户端数据落地的适配及启动功能, 目前支持功能:

- 客户端启动器
- 同步管理REST接口
- 日志适配器, 作为DEMO
- 关系型数据库的数据同步(表对表同步), ETL功能
- HBase的数据同步(表对表同步), ETL功能
- ElasticSearch多表数据同步,ETL功能

```
# adapter ClientAdapter
# admin 管理界面
# deployer canal-server


wget https://github.com/alibaba/canal/releases/download/canal-1.1.4/canal.adapter-1.1.4.tar.gz

mkdir canal-consumer

tar zxvf canal.adapter-1.1.5.tar.gz  -C ~/canal-consumer

cd canal-consumer

```
### 总配置文件 application.yml
adapter定义配置部分
```
canal.conf:
  canalServerHost: 127.0.0.1:11111          # 对应单机模式下的canal server的ip:port
  zookeeperHosts: slave1:2181               # 对应集群模式下的zk地址, 如果配置了canalServerHost, 则以canalServerHost为准
  mqServers: slave1:6667 #or rocketmq       # kafka或rocketMQ地址, 与canalServerHost不能并存
  flatMessage: true                         # 扁平message开关, 是否以json字符串形式投递数据, 仅在kafka/rocketMQ模式下有效
  batchSize: 50                             # 每次获取数据的批大小, 单位为K
  syncBatchSize: 1000                       # 每次同步的批数量
  retries: 0                                # 重试次数, -1为无限重试
  timeout:                                  # 同步超时时间, 单位毫秒
  mode: tcp # kafka rocketMQ                # canal client的模式: tcp kafka rocketMQ
  srcDataSources:                           # 源数据库
    defaultDS:                              # 自定义名称
      url: jdbc:mysql://127.0.0.1:3306/mytest?useUnicode=true   # jdbc url 
      username: root                                            # jdbc 账号
      password: 121212                                          # jdbc 密码
  canalAdapters:                            # 适配器列表
  - instance: example                       # canal 实例名或者 MQ topic 名
    groups:                                 # 分组列表
    - groupId: g1                           # 分组id, 如果是MQ模式将用到该值
      outerAdapters:                        # 分组内适配器列表
      - name: logger                        # 日志打印适配器
......           
```
说明:

- **一份数据可以被多个group同时消费, 多个group之间会是一个并行执行, 一个group内部是一个串行执行多个outerAdapters, 比如例子中logger和hbase**
- 目前client adapter数据订阅的方式支持两种，直连canal server 或者 订阅kafka/RocketMQ的消息



### 使用远程配置(Mysql)
可以将本地application.yml文件和其他子配置文件删除或清空， 启动工程将自动从远程加载配置
修改mysql中的配置信息后会自动刷新到本地动态加载相应的实例或者应用
**前提:**
启动Canal Admin
#### 修改bootstrap.yml配置
```
canal:
  manager:
    jdbc:
      url: jdbc:mysql://127.0.0.1:3306/canal_manager?useUnicode=true&characterEncoding=UTF-8
      username: root
      password: 121212
```


### 适配器启动

```
-- ----------------------------
-- Table structure for label
-- ----------------------------
DROP TABLE IF EXISTS `label`;
CREATE TABLE `label` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `label` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `c_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `role_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

insert into user (id,name,role_id) values (1,'Eric',1);
insert into role (id,role_name) values (1,'admin');
insert into role (id,role_name) values (2,'operator');
insert into label (id,user_id,label) values (1,1,'a');
insert into label (id,user_id,label) values (2,1,'b');
commit;
```
#### 启动canal-adapter示例
https://github.com/alibaba/canal/wiki/ClientAdapter#31-%E5%90%AF%E5%8A%A8canal-adapter%E7%A4%BA%E4%BE%8B
同步es7(V1.1.5)
配置
```
server:
  port: 8081
spring:
  jackson:
    date-format: yyyy-MM-dd HH:mm:ss
    time-zone: GMT+8
    default-property-inclusion: non_null

canal.conf:
  mode: tcp # kafka rocketMQ
  canalServerHost: 127.0.0.1:11111
#  zookeeperHosts: slave1:2181
#  mqServers: 127.0.0.1:9092 #or rocketmq
#  flatMessage: true
  batchSize: 500
  syncBatchSize: 1000
  retries: 0
  timeout:
  accessKey:
  secretKey:
  srcDataSources:
    defaultDS:
      url: jdbc:mysql://127.0.0.1:3306/mytest?useUnicode=true
      username: admin
      password: lcb@2019.com
  canalAdapters:
  - instance: example # canal instance Name or mq topic name
    groups:
    - groupId: g1
      outerAdapters:
      - name: logger
#      - name: rdb
#        key: mysql1
#        properties:
#          jdbc.driverClassName: com.mysql.jdbc.Driver
#          jdbc.url: jdbc:mysql://127.0.0.1:3306/mytest2?useUnicode=true
#          jdbc.username: root
#          jdbc.password: 121212
#      - name: rdb
#        key: oracle1
#        properties:
#          jdbc.driverClassName: oracle.jdbc.OracleDriver
#          jdbc.url: jdbc:oracle:thin:@localhost:49161:XE
#          jdbc.username: mytest
#          jdbc.password: m121212
#      - name: rdb
#        key: postgres1
#        properties:
#          jdbc.driverClassName: org.postgresql.Driver
#          jdbc.url: jdbc:postgresql://localhost:5432/postgres
#          jdbc.username: postgres
#          jdbc.password: 121212
#          threads: 1
#          commitSize: 3000
#      - name: hbase
#        properties:
#          hbase.zookeeper.quorum: 127.0.0.1
#          hbase.zookeeper.property.clientPort: 2181
#          zookeeper.znode.parent: /hbase
      - name: es7
        hosts: 192.168.110.174:9200 # 127.0.0.1:9200 for rest mode
        properties:
          mode: rest # or rest
          # security.auth: test:123456 #  only used for rest mode
          cluster.name: es-cluster


```
新增conf/es7/mytest_user.yml文件
```
dataSourceKey: defaultDS        # 源数据源的key, 对应上面配置的srcDataSources中的值
outerAdapterKey: exampleKey     # 对应application.yml中es配置的key 
destination: example            # cannal的instance或者MQ的topic
groupId:                        # 对应MQ模式下的groupId, 只会同步对应groupId的数据
esMapping:
  _index: mytest_user           # es 的索引名称
  _type: _doc                   # es 的type名称, es7下无需配置此项
  _id: _id                      # es 的_id, 如果不配置该项必须配置下面的pk项_id则会由es自动分配
#  pk: id                       # 如果不需要_id, 则需要指定一个属性为主键属性
  # sql映射
  sql: "select a.id as _id, a.name as _name, a.role_id as _role_id, b.role_name as _role_name,
        a.c_time as _c_time, c.labels as _labels from user a
        left join role b on b.id=a.role_id
        left join (select user_id, group_concat(label order by id desc separator ';') as labels from label
        group by user_id) c on c.user_id=a.id"
#  objFields:
#    _name: array:,           # 数组或者对象属性, array:, 代表字段里面是以,分隔成的数组，_name是以，分割的字符串
#    _role_name: object               # _role_name 是 json字符串
  etlCondition: "where a.c_time>='{0}'"     # etl 的条件参数
  commitBatch: 3000                         # 提交批大小

```
用的这个sql
```
"select a.id as _id, a.name as _name, a.role_id as _role_id, b.role_name as _role_name,
        a.c_time as _c_time from user a
        left join role b on b.id=a.role_id"
```
启动
```
cd canal-consumer
sh bin/startup.sh
```
#### adapter管理REST接口
https://github.com/alibaba/canal/wiki/ClientAdapter#32-adapter%E7%AE%A1%E7%90%86rest%E6%8E%A5%E5%8F%A3

## Canal Admin
### 下载解压
```
# admin 管理界面

wget https://github.com/alibaba/canal/releases/download/canal-1.1.4/canal.admin-1.1.4.tar.gz

mkdir canal-admin

tar zxvf canal.admin-1.1.4.tar.gz  -C ~/canal-admin

cd canal-admin

```
## 配置修改
vi conf/application.yml
```
server:
  port: 8089
spring:
  jackson:
    date-format: yyyy-MM-dd HH:mm:ss
    time-zone: GMT+8

spring.datasource:
  address: 127.0.0.1:3306
  database: canal_manager
  username: canal
  password: canal
  driver-class-name: com.mysql.jdbc.Driver
  url: jdbc:mysql://${spring.datasource.address}/${spring.datasource.database}?useUnicode=true&characterEncoding=UTF-8&useSSL=false
  hikari:
    maximum-pool-size: 30
    minimum-idle: 1

canal:
  adminUser: admin
  adminPasswd: admin
```
### 初始化元数据库
```
mysql

# 导入初始化SQL
> source conf/canal_manager.sql
```

### 启动
```
sh bin/startup.sh

vi logs/admin.log

http://127.0.0.1:8089/ 访问，默认密码：admin/123456

sh bin/stop.sh
```

# 全量同步
curl 127.0.0.1:8081/etl/es7/mytest_user.yml -X POST

# 其它
[go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch)
[canal同步到redis](https://blog.csdn.net/clypm/article/details/80599819)

[全量同步方案code](https://github.com/logstash-plugins/logstash-input-jdbc)
[全量同步方案doc](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-jdbc.html)

