https://github.com/alibaba/canal/wiki/Docker-QuickStart

canal:Canal 是mysql数据库binlog的增量订阅&消费组件。

基于日志增量订阅&消费支持的业务：

1. 数据库镜像
2. 数据库实时备份
3. 多级索引 (卖家和买家各自分库索引)
4. search build
5. 业务cache刷新
6. 价格变化等重要业务消息


![canal](../img/canal_v2.png)

1. Canal连接到mysql数据库，模拟slave
2. client与Canal建立连接
3. 数据库发生变更写入到binlog
4. Canal向数据库发送dump请求，获取binlog并解析
5. canal-go向Canal请求数据库变更
6. Canal发送解析后的数据给client
7. client收到数据，消费成功，发送回执。（可选）
8. Canal记录消费位置。

也可以通过[ClientAdapter](https://github.com/alibaba/canal/wiki/ClientAdapter) 同步到其它存储

如es:https://github.com/alibaba/canal/wiki/Sync-ES

docker pull canal/canal-server:v1.1.4