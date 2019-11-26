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