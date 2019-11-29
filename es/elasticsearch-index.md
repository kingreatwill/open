[TOC]
## 设置索引分片数
```
PUT /mytest_user
{
    "settings" : {
        "number_of_shards" : 3,
        "number_of_replicas" : 2
    }
}

number_of_shards 是不能修改的，如下会出错
PUT /mytest_user/_settings
{
    "index" : {
        "number_of_shards" : 2
    }
}
```

## Open/Close  Index   打开/关闭索引
```
POST /my_index/_close
POST /my_index/_open
```
说明：
关闭的索引不能进行读写操作，几乎不占集群开销。
关闭的索引可以打开，打开走的是正常的恢复流程。

## 设置索引的读写
```
index.blocks.read_only：设为true,则索引以及索引的元数据只可读
index.blocks.read_only_allow_delete：设为true，只读时允许删除。
index.blocks.read：设为true，则不可读。
index.blocks.write：设为true，则不可写。
index.blocks.metadata：设为true，则索引元数据不可读写。
```

## Shrink Index 收缩索引
索引的分片数是不可更改的，如要减少分片数可以通过收缩方式收缩为一个新的索引。新索引的分片数必须是原分片数的因子值，如原分片数是8，则新索引的分片数可以为4、2、1 。

### 什么时候需要收缩索引呢?
最初创建索引的时候分片数设置得太大，后面发现用不了那么多分片，这个时候就需要收缩了

### 收缩的流程：
先把所有主分片都转移到一台主机上；
在这台主机上创建一个新索引，分片数较小，其他设置和原索引一致；
把原索引的所有分片，复制（或硬链接）到新索引的目录下；
对新索引进行打开操作恢复分片数据；
(可选)重新把新索引的分片均衡到其他节点上。

### 收缩前的准备工作：
将原索引设置为只读；
将原索引各分片的一个副本重分配到同一个节点上，并且要是健康绿色状态。
```
PUT /my_source_index/_settings
{
  "settings": {
    <!-- 指定进行收缩的节点的名称 -->
    "index.routing.allocation.require._name": "shrink_node_name",
    <!-- 阻止写，只读 -->
     "index.blocks.write": true
  }
}
```
### 进行收缩：
```
POST my_source_index/_shrink/my_target_index
{
  "settings": {
    "index.number_of_replicas": 1,
    "index.number_of_shards": 1,
    "index.codec": "best_compression"
  }}
```
### 监控收缩过程：
```
GET _cat/recovery?v
GET _cluster/health
```

## Split Index 拆分索引
当索引的分片容量过大时，可以通过拆分操作将索引拆分为一个倍数分片数的新索引。能拆分为几倍由创建索引时指定的index.number_of_routing_shards 路由分片数决定。这个路由分片数决定了根据一致性hash路由文档到分片的散列空间。

如index.number_of_routing_shards = 30 ，指定的分片数是5，则可按如下倍数方式进行拆分：
```
5 → 10 → 30 (split by 2, then by 3)
5 → 15 → 30 (split by 3, then by 2)
5 → 30 (split by 6)
```

为什么需要拆分索引？
当最初设置的索引的分片数不够用时就需要拆分索引了，和压缩索引相反

注意：只有在创建时指定了index.number_of_routing_shards 的索引才可以进行拆分，ES7开始将不再有这个限制。

和solr的区别是，solr是对一个分片进行拆分，es中是整个索引进行拆分。

拆分步骤：
准备一个索引来做拆分：
```

PUT my_source_index
{
    "settings": {
        "index.number_of_shards" : 1,
        <!-- 创建时需要指定路由分片数 -->
        "index.number_of_routing_shards" : 2
    }
}
```
先设置索引只读：
```

PUT /my_source_index/_settings
{
  "settings": {
    "index.blocks.write": true
  }
}
```

做拆分：
```
POST my_source_index/_split/my_target_index
{
  "settings": {
    <!--新索引的分片数需符合拆分规则-->
    "index.number_of_shards": 2
  }
}
```
监控拆分过程：
```
GET _cat/recovery?v
GET _cluster/health
```

## Rollover Index 别名滚动指向新创建的索引

对于有时效性的索引数据，如日志，过一定时间后，老的索引数据就没有用了。我们可以像数据库中根据时间创建表来存放不同时段的数据一样，在ES中也可用建多个索引的方式来分开存放不同时段的数据。比数据库中更方便的是ES中可以通过别名滚动指向最新的索引的方式，让你通过别名来操作时总是操作的最新的索引。

ES的rollover index API 让我们可以根据满足指定的条件（时间、文档数量、索引大小）创建新的索引，并把别名滚动指向新的索引。

注意：这时的别名只能是一个索引的别名。

Rollover Index 示例：

创建一个名字为logs-0000001 、别名为logs_write 的索引：
```
PUT /logs-000001
{
  "aliases": {
    "logs_write": {}
  }
}
```
添加1000个文档到索引logs-000001，然后设置别名滚动的条件
```
POST /logs_write/_rollover
{
  "conditions": {
    "max_age":   "7d",
    "max_docs":  1000,
    "max_size":  "5gb"
  }
}
```
说明：

如果别名logs_write指向的索引是7天前（含）创建的或索引的文档数>=1000或索引的大小>= 5gb，则会创建一个新索引 logs-000002，并把别名logs_writer指向新创建的logs-000002索引

Rollover Index 新建索引的命名规则：

如果索引的名称是-数字结尾，如logs-000001，则新建索引的名称也会是这个模式，数值增1。

如果索引的名称不是-数值结尾，则在请求rollover api时需指定新索引的名称
```
POST /my_alias/_rollover/my_new_index_name
{
  "conditions": {
    "max_age":   "7d",
    "max_docs":  1000,
    "max_size": "5gb"
  }
}
```
在名称中使用Date math（时间表达式）

如果你希望生成的索引名称中带有日期，如logstash-2016.02.03-1 ，则可以在创建索引时采用时间表达式来命名：
```
# PUT /<logs-{now/d}-1> with URI encoding:
PUT /%3Clogs-%7Bnow%2Fd%7D-1%3E
{
  "aliases": {
    "logs_write": {}
  }
}
PUT logs_write/_doc/1
{
  "message": "a dummy log"
} 
POST logs_write/_refresh
# Wait for a day to pass
POST /logs_write/_rollover
{
  "conditions": {
    "max_docs":   "1"
  }
}
```
Rollover时可对新的索引作定义：
```
PUT /logs-000001
{
  "aliases": {
    "logs_write": {}
  }
}
POST /logs_write/_rollover
{
  "conditions" : {
    "max_age": "7d",
    "max_docs": 1000,
    "max_size": "5gb"
  },
  "settings": {
    "index.number_of_shards": 2
  }
}
```
Dry run  实际操作前先测试是否达到条件：
```
POST /logs_write/_rollover?dry_run
{
  "conditions" : {
    "max_age": "7d",
    "max_docs": 1000,
    "max_size": "5gb"
  }
}
```
说明：

测试不会创建索引，只是检测条件是否满足

注意：rollover是你请求它才会进行操作，并不是自动在后台进行的。你可以周期性地去请求它。

## 索引监控
### 查看索引状态信息

官网链接：

https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-stats.html

查看所有的索引状态：
```
GET /_stats
```
查看指定索引的状态信息：
```
GET /index1,index2/_stats
```
### 查看索引段信息
官网链接：

https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-segments.html
```
GET /test/_segments 
GET /index1,index2/_segments
GET /_segments
```
### 查看索引恢复信息

官网链接：
https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-recovery.html
```
GET index1,index2/_recovery?human

GET /_recovery?human
```
### 查看索引分片的存储信息
官网链接：

https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shards-stores.html
```
# return information of only index test
GET /test/_shard_stores
# return information of only test1 and test2 indices
GET /test1,test2/_shard_stores
# return information of all indices
GET /_shard_stores
GET /_shard_stores?status=green
```

##  索引状态管理
### Clear Cache 清理缓存
```
POST /twitter/_cache/clear
```
默认会清理所有缓存，可指定清理query, fielddata or request 缓存
```
POST /kimchy,elasticsearch/_cache/clear
POST /_cache/clear
```

### Refresh，重新打开读取索引
```
POST /kimchy,elasticsearch/_refresh
POST /_refresh
```

### Flush，将缓存在内存中的索引数据刷新到持久存储中
```
POST twitter/_flush
```

### Force merge 强制段合并
```
POST /kimchy/_forcemerge?only_expunge_deletes=false&max_num_segments=100&flush=true
```
可选参数说明：
max_num_segments 合并为几个段，默认1
only_expunge_deletes 是否只合并含有删除文档的段，默认false
flush 合并后是否刷新，默认true
```

POST /kimchy,elasticsearch/_forcemerge
POST /_forcemerge
```