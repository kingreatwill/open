
[TOC]
# es7.4.2 安装
## 安装并运行Elasticsearch集群
### Download 
```
curl -L -O https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.4.2-linux-x86_64.tar.gz
```
### Extract the archive
```
tar -xvf elasticsearch-7.4.2-linux-x86_64.tar.gz
# powershell Expand-Archive elasticsearch-7.4.2-windows-x86_64.zip
```


### Start Elasticsearch from the bin directory
```
cd elasticsearch-7.4.2/bin
./elasticsearch
```
启动另外两个Elasticsearch实例，以便您可以看到典型的多节点集群的行为。
您需要为每个节点指定唯一的数据和日志路径。
```
./elasticsearch -Epath.data=data2 -Epath.logs=log2
./elasticsearch -Epath.data=data3 -Epath.logs=log3
```
为其他节点分配了唯一的ID。 因为您在本地运行所有三个节点，所以它们会自动与第一个节点一起加入群集。

### 查看集群的健康状况
http://localhost:9200/_cat/health?v
说明：**v是用来要求在结果中返回表头**
说明：**pretty是用来要求在结果中返回json格式**
```
epoch      timestamp cluster       status node.total node.data shards pri relo init unassign pending_tasks max_task_wait_time active_shards_percent
1565052807 00:53:27  elasticsearch green           3         3      6   3    0    0        0             0                  -                100.0%
```
状态值说明
status | 说明
--|--
Green | everything is good (cluster is fully functional)，即最佳状态
Yellow | all data is available but some replicas are not yet allocated (cluster is fully functional)，即数据和集群可用，但是集群的备份有的是坏的
Red | some data is not available for whatever reason (cluster is partially functional)，即数据和集群都不可用
### 查看集群信息
http://localhost:9200/_cluster/health?pretty
```
{
  "cluster_name" : "es-cluster",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 3,
  "number_of_data_nodes" : 2,
  "active_primary_shards" : 44,
  "active_shards" : 88,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
```
### 查看集群的节点
http://localhost:9200/_cat/?v


## Index some sample documents
http://localhost:5601/app/kibana#/dev_tools/console
```
PUT /customer/_doc/1
{
  "name": "John Doe"
}
```
```

curl -X PUT "localhost:9200/customer/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "name": "John Doe"
}
'
```
该请求（如果尚不存在）将自动创建客户索引，添加ID为1的新文档，并存储名称字段并为其建立索引。

创建一个名为 customer 的索引
```
PUT /customer?pretty
```

```
GET /customer/_doc/1
```
```
curl -X GET "localhost:9200/customer/_doc/1?pretty"
```


### Indexing documents in bulk(批量)
Download the [accounts.json](https://raw.githubusercontent.com/elastic/elasticsearch/master/docs/src/test/resources/accounts.json)
```json
{"index":{"_id":"1"}}
{"account_number":1,"balance":39225,"firstname":"Amber","lastname":"Duke","age":32,"gender":"M","address":"880 Holmes Lane","employer":"Pyrami","email":"amberduke@pyrami.com","city":"Brogan","state":"IL"}
```
Index the account data into the `bank` index with the following `_bulk `request:
```
curl -H "Content-Type: application/json" -XPOST "localhost:9200/bank/_bulk?pretty&refresh" --data-binary "@accounts.json"
```

说明：pretty要求返回一个漂亮的json 结果

### 查看所有索引
http://localhost:9200/_cat/indices?v
```
GET /_cat/indices?v
```

## 使用Elasticsearch查询语言搜索文档
查询所有文档
```
GET /customer/_search?q=*&sort=name:asc&pretty ## 会出错，name没有索引
出现该错误是因为5.x之后，Elasticsearch对排序、聚合所依据的字段用单独的数据结构(fielddata)缓存到内存里了，但是在text字段上默认是禁用的，如果有需要单独开启，这样做的目的是为了节省内存空间
https://www.elastic.co/guide/en/elasticsearch/reference/current/fielddata.html

解决方案

PUT /customer/_mapping
{       
   "properties": {
         "name": {  
             "type": "text",
             "fielddata": true
         }       
     }         
}



这个默认是可以的
GET /customer/_search?q=*&sort=_id:asc&pretty
```
或者json方式
```
GET /bank/_search
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ]
}
```

默认情况下只返回10条结果
```json
{
  "took" : 1,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 1000,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [
      {
        "_index" : "bank",
        "_type" : "_doc",
        "_id" : "0",
        "_score" : null,
        "_source" : {
          "account_number" : 0,
          "balance" : 16623,
          "firstname" : "Bradshaw",
          "lastname" : "Mckenzie",
          "age" : 29,
          "gender" : "F",
          "address" : "244 Columbus Place",
          "employer" : "Euron",
          "email" : "bradshawmckenzie@euron.com",
          "city" : "Hobucken",
          "state" : "CO"
        },
        "sort" : [
          0
        ]
      },
     ---
    ]
  }
}

```
该响应还提供有关搜索请求的以下信息：
- took – Elasticsearch运行查询所需的时间（以毫秒为单位）
- timed_out – 搜索请求是否超时
- _shards – 搜索了多少个分片以及成功，失败或跳过了多少个分片。
- max_score – 找到最相关文件的分数
- hits.total.value - 找到多少个匹配的文档
- hits.sort - 文档的排序位置（不按相关性得分排序时）
- hits._score - 文档的相关性得分（使用match_all时不适用）

每个搜索请求都是独立的：Elasticsearch在请求中不维护任何状态信息。
要翻阅搜索结果，请在请求中指定from和size参数。
```
GET /bank/_search
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ],
  "from": 10,
  "size": 10
}
```
查找address包含mill or lane
```
GET /bank/_search
{
  "query": { "match": { "address": "mill lane" } }
}
```
查找address包含mill lane
```
GET /bank/_search
{
  "query": { "match_phrase": { "address": "mill lane" } }
}
```

查找age=40 AND state<>ID
```
GET /bank/_search
{
  "query": {
    "bool": {
      "must": [
        { "match": { "age": "40" } }
      ],
      "must_not": [
        { "match": { "state": "ID" } }
      ]
    }
  }
}
```

>Boolean query
 must 
 should  
 must_not


查询范围
```
GET /bank/_search
{
  "query": {
    "bool": {
      "must": { "match_all": {} },
      "filter": {
        "range": {
          "balance": {
            "gte": 20000,
            "lte": 30000
          }
        }
      }
    }
  }
}
```

## Analyze the results using bucket and metrics aggregations

aggs聚合
terms术语,条款

以下请求使用术语汇总将银行索引中的所有帐户按状态分组，并按降序返回帐户数量最多的十个州：
```
GET /bank/_search
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword"
      }
    }
  }
}
```
您可以组合聚合以构建更复杂的数据汇总。 例如，以下请求在前一个group_by_state聚合内嵌套avg聚合，以计算每个状态的平均帐户余额。
```
GET /bank/_search
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword"
      },
      "aggs": {
        "average_balance": {
          "avg": {
            "field": "balance"
          }
        }
      }
    }
  }
}
```
您可以通过在术语聚合中指定顺序来使用嵌套聚合的结果进行排序，而不必按计数对结果进行排序：

```
GET /bank/_search
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword",
        "order": {
          "average_balance": "desc"
        }
      },
      "aggs": {
        "average_balance": {
          "avg": {
            "field": "balance"
          }
        }
      }
    }
  }
}
```

Elasticsearch提供了专门的聚合，可用于多个字段上的操作并分析特定类型的数据，例如日期，IP地址和地理数据

您还可以将单个聚合的结果馈送到管道聚合中（pipeline aggregations），以进行进一步分析。

聚合提供的核心分析功能可启用高级功能，例如使用机器学习来检测异常。

## IK中文分词器安装
### 安装
https://github.com/medcl/elasticsearch-analysis-ik/releases
1. download or compile
optional 1 - download pre-build package from here: https://github.com/medcl/elasticsearch-analysis-ik/releases

create plugin folder cd your-es-root/plugins/ && mkdir ik

unzip plugin to folder your-es-root/plugins/ik

optional 2 - use elasticsearch-plugin to install ( supported from version v5.5.1 ):
./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.4.2/elasticsearch-analysis-ik-7.4.2.zip

2. restart elasticsearch

### 使用
1. create a index
```
curl -XPUT http://localhost:9200/index
```
2. create a mapping
```
curl -XPOST http://localhost:9200/index/_mapping -H 'Content-Type:application/json' -d'
{
        "properties": {
            "content": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            }
        }

}'
```
3. index some docs
```
curl -XPOST http://localhost:9200/index/_create/1 -H 'Content-Type:application/json' -d'
{"content":"美国留给伊拉克的是个烂摊子吗"}
'
curl -XPOST http://localhost:9200/index/_create/2 -H 'Content-Type:application/json' -d'
{"content":"公安部：各地校车将享最高路权"}
'
curl -XPOST http://localhost:9200/index/_create/3 -H 'Content-Type:application/json' -d'
{"content":"中韩渔警冲突调查：韩警平均每天扣1艘中国渔船"}
'
curl -XPOST http://localhost:9200/index/_create/4 -H 'Content-Type:application/json' -d'
{"content":"中国驻洛杉矶领事馆遭亚裔男子枪击 嫌犯已自首"}
'
```
4. query with highlighting
```
curl -XPOST http://localhost:9200/index/_search  -H 'Content-Type:application/json' -d'
{
    "query" : { "match" : { "content" : "中国" }},
    "highlight" : {
        "pre_tags" : ["<tag1>", "<tag2>"],
        "post_tags" : ["</tag1>", "</tag2>"],
        "fields" : {
            "content" : {}
        }
    }
}
'
```

5. _analyze
```

POST  /_analyze
{
    "text":"中华人民共和国国歌china"
    ,"analyzer":"ik_smart"
}

curl -XGET "http://localhost:9200/your_index/_analyze" -H 'Content-Type: application/json' -d'
{
   "text":"中华人民共和国MN","tokenizer": "my_ik"
}'
```
### Dictionary Configuration
IKAnalyzer.cfg.xml can be located at {conf}/analysis-ik/config/IKAnalyzer.cfg.xml or {plugins}/elasticsearch-analysis-ik-*/config/IKAnalyzer.cfg.xml
```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE properties SYSTEM "http://java.sun.com/dtd/properties.dtd">
<properties>
	<comment>IK Analyzer 扩展配置</comment>
	<!--用户可以在这里配置自己的扩展字典 -->
	<entry key="ext_dict">custom/mydict.dic;custom/single_word_low_freq.dic</entry>
	 <!--用户可以在这里配置自己的扩展停止词字典-->
	<entry key="ext_stopwords">custom/ext_stopword.dic</entry>
 	<!--用户可以在这里配置远程扩展字典 -->
	<entry key="remote_ext_dict">location</entry>
 	<!--用户可以在这里配置远程扩展停止词字典-->
	<entry key="remote_ext_stopwords">http://xxx.com/xxx.dic</entry>
</properties>
```
#### 热更新 IK 分词使用方法
目前该插件支持热更新 IK 分词，通过上文在 IK 配置文件中提到的如下配置
```
<!--用户可以在这里配置远程扩展字典 -->
	<entry key="remote_ext_dict">location</entry>
 	<!--用户可以在这里配置远程扩展停止词字典-->
	<entry key="remote_ext_stopwords">location</entry>
```
其中 location 是指一个 url，比如 http://yoursite.com/getCustomDict，该请求只需满足以下两点即可完成分词热更新。

1. 该 http 请求需要返回两个头部(header)，一个是 Last-Modified，一个是 ETag，这两者都是字符串类型，只要有一个发生变化，该插件就会去抓取新的分词进而更新词库。

2. 该 http 请求返回的内容格式是一行一个分词，换行符用 \n 即可。

满足上面两点要求就可以实现热更新分词了，不需要重启 ES 实例。

可以将需自动更新的热词放在一个 UTF-8 编码的 .txt 文件里，放在 nginx 或其他简易 http server 下，当 .txt 文件修改时，http server 会在客户端请求该文件时自动返回相应的 Last-Modified 和 ETag。可以另外做一个工具来从业务系统提取相关词汇，并更新这个 .txt 文件。

#### ik_max_word 和 ik_smart 什么区别?
**ik_max_word**: 会将文本做最细粒度的拆分，比如会将“中华人民共和国国歌”拆分为“中华人民共和国,中华人民,中华,华人,人民共和国,人民,人,民,共和国,共和,和,国国,国歌”，会穷尽各种可能的组合，适合 Term Query；

**ik_smart**: 会做最粗粒度的拆分，比如会将“中华人民共和国国歌”拆分为“中华人民共和国,国歌”，适合 Phrase 查询。

## 管理面板
https://github.com/mobz/elasticsearch-head