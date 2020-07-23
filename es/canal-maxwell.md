# MySQL Binlog 解析工具 Maxwell 详解

https://github.com/zendesk/maxwell 2.2k

数据格式：
```
mysql> insert into `test`.`maxwell` set id = 1, daemon = 'Stanislaw Lem';
  maxwell: {
    "database": "test",
    "table": "maxwell",
    "type": "insert",
    "ts": 1449786310,
    "xid": 940752,
    "commit": true,
    "data": { "id":1, "daemon": "Stanislaw Lem" }
  }
mysql> update test.maxwell set daemon = 'firebus!  firebus!' where id = 1;
  maxwell: {
    "database": "test",
    "table": "maxwell",
    "type": "update",
    "ts": 1449786341,
    "xid": 940786,
    "commit": true,
    "data": {"id":1, "daemon": "Firebus!  Firebus!"},
    "old":  {"daemon": "Stanislaw Lem"}
  }
```

Maxwell主要提供了下列功能：

- 支持 SELECT * FROM table 的方式进行全量数据初始化
- 支持在主库发生failover后，自动恢复binlog位置(GTID)
- 可以对数据进行分区，解决数据倾斜问题，发送到kafka的数据支持database、table、column等级别的数据分区
- 工作方式是伪装为Slave，接收binlog events，然后根据schemas信息拼装，可以接受ddl、xid、row等各种event

特色|canal|maxwell|mysql_streamer
---|---|---|---
开源方|阿里巴巴|zendesk|Yelp
语言|java|java|Python
活跃度|活跃x|活跃|活跃
HA|支持|定制|支持
数据落地|定制|Kafka等|Kafka
分区|支持|不支持|不支持
bootstrap|不支持|支持|支持
数据格式|格式自由|json|json
文档|较详细|较详细|粗略
随机读|支持|支持|支持

- MySQL需要先启用binlog
```
$ vi my.cnf

[mysqld]
server_id=1
log-bin=master
binlog_format=row
```
- 创建Maxwell用户，并赋予 maxwell 库的一些权限
```
CREATE USER 'maxwell'@'%' IDENTIFIED BY '123456';
GRANT ALL ON maxwell.* TO 'maxwell'@'%';
GRANT SELECT, REPLICATION CLIENT, REPLICATION SLAVE on *.* to 'maxwell'@'%'; 
```
- 使用 maxwell 之前需要先启动 kafka
- 通过 docker 快速安装并使用 Maxwell

输出到 控制台
```
# 拉取镜像 
docker pull zendesk/maxwell

# 启动maxwell，并将解析出的binlog输出到控制台
docker run -ti --rm zendesk/maxwell bin/maxwell --user='maxwell' --password='123456' --host='10.100.97.246' --producer=stdout
```
输出到 Kafka
```
docker run -it --rm zendesk/maxwell bin/maxwell --user='maxwell' \
    --password='123456' --host='10.100.97.246' --producer=kafka \
    --kafka.bootstrap.servers='10.100.97.246:9092' --kafka_topic=maxwell --log_level=debug
```

## 其它更多
[MySQL Binlog 解析工具 Maxwell 详解](https://blog.csdn.net/wwwdc1012/article/details/88388552)
