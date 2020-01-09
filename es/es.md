字段类型：https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html

<!-- toc -->
[TOC]

## Dynamic Mapping
在写入文档的时候，如果索引不存在，会自动创建索引
Dynamic Mapping 的机制，使得我们无需手动定义 Mappings。Elasticsearch 会自动根据文档信息，推算出字段的类型
但是会有时候推算不对。当类型如果设置不对时，会导致一些功能无法正常运行
### 能否更改 Mapping 的字段类型
两种情况
1. 新增字段
 - Dynamic 设置为 true 时，一定有新增字段的文档写入，Mapping 也同时被更新
 - Dynamic 设为 false，Mapping 不会被更新，自增字段的数据无法被索引，但是信息会出现在_source 中
 - Dynamic 设置成 Strict 文档写入失败
2. 对已有字段，一旦已经有数据写入，就不在支持修改字段定义
 - Luene 实现的倒排索引，一旦生成后，就不允许修改
 
如果希望改变字段类型，必须 Reindex API，重建索引
如果修改了字段的数据类型，会导致已被索引的属于无法被搜索
但是如果是增加新的字段，就不会有这样的影响

### 新增字段
```
PUT /mytest_user/_mapping
{
  "properties": {
    "add_new2": {
      "type": "date",
      "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
    }
  }
}
```

### 关闭动态字段
```
//修改为dynamic false(丢弃)/strict(报错)/true(自动新增，可能不是我们想要的字段)
PUT dynamic_mapping_test/_mapping
{
  "dynamic":false
}
```

## Multi-fields



https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
```
PUT /my_indexxxx
{
  "mappings": {
    "properties": {
      "city": {
        "type": "text"
      }
    }
  }
}
PUT /my_indexxxx/_mapping
{
  "properties": {
    "city": {
      "type": "text",
      "fields": {
        "raw": {
          "type": "keyword"
        }
      }
    }
  }
}
```
The city field can be used for full text search.
The city.raw field can be used for sorting and aggregations

## 重命名字段
```
PUT /my_index
{
  "mappings": {
    "properties": {
      "user_identifier": {
        "type": "keyword"
      }
    }
  }
}

PUT /my_index/_mapping
{
  "properties": {
    "user_id": {
      "type": "alias",
      "path": "user_identifier"
    }
  }
}
```

# 注意点:
## 手动创建index
1. 指定别名，方便以后修改字段类型和数据迁移(DELETE /twitter/_alias/alias1)
2. mapping properties，注意properties类型，是否需要分词
3. 指定主分片数量和复制分片份数，分片数量不能修改
4. 指定分词器（组合）
5. 注意text和keyworld区别，text 如果需要排序 请添加"fielddata": true，但是会影响性能
6. copy_to  "copy_to": "full_name"
7. date类型的默认格式为："strict_date_optional_time||epoch_millis",一旦定义后不可修改


## 时间格式
https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-date-format.html
默认为两种格式：
strict_date_optional_time 此格式为ISO8601标准 示例：2018-08-31T14:56:18.000+08:00
epoch_millis 也就是时间戳 示例1515150699465, 1515150699

实测，仅支持"yyyy-MM-dd"、"yyyyMMdd"、"yyyyMMddHHmmss"、"yyyy-MM-ddTHH:mm:ss"、"yyyy-MM-ddTHH:mm:ss.SSS"、"yyyy-MM-ddTHH:mm:ss.SSSZ"格式，不支持常用的"yyyy-MM-dd HH:mm:ss"等格式。
注意，"T"和"Z"是固定的字符，在获取"yyyy-MM-ddTHH:mm:ss"、"yyyy-MM-ddTHH:mm:ss.SSS"、"yyyy-MM-ddTHH:mm:ss.SSSZ"格式字符串值时，不能直接以前面格式格式化date，而是需要多次格式化date并且拼接得到。


## 场景:
1. 单表-大表（行或量）GoodsBasic 2,921,100  5.12 GB   xx_index 20,007,716 9.01 GB，3分钟200W行600MB  20分钟810W行 2.1GB    1小时 1520W行 4.8GB  2个小时同步完成2200W行5.5GB
2. 扩展表合并merchant SecurityCharacter  StoreBasic CompanyBasic  和   SecurityUser SecurityVita
3. 父子表-一个子表和多个子表

### 单表


全量同步：curl 127.0.0.1:8081/etl/es7/xx_index.yml -X POST -d "params=2018-10-21 00:00:00"

select a.LogId, a.OrderType,a.CompchterId,a.WarehouseCode,a.xx_name,a.BarCode,a.GoodsName,a.CreateTime from xx_index a
```
PUT /goods-stock-log
{
  "aliases" : {
        "xx_index" : {}
    },
  "settings" : {
      "number_of_shards" : 5, 
      "number_of_replicas" : 1
   },
  "mappings":{
      "properties":{
        "LogId": {
          "type": "keyword"
        },
		 "OrderType": {
          "type": "long"
        },
        "CompchterId": {
          "type": "long"
        },
        "WarehouseCode": {
          "type": "keyword"
        },       
        "xx_name": {
          "type": "keyword"
        },
		"BarCode": {
          "type": "keyword"
        },
		"GoodsName": {
          "type": "text",
		  "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
        "CreateTime": {
          "type": "date"
        }
      }
  }
}
```

### 扩展表合并
merchant SecurityCharacter  StoreBasic CompanyBasic 

一对一, 多对一
这里join操作只能是left outer join, 第一张表必须为主表!!
SecurityUser SecurityVita SecurityRole

select a.UserId,a.RoleId, a.CharacterId,a.LoginId,b.TrueName,b.UniteNote,c.RoleName,a.CreateTime from SecurityUser a left join SecurityVita b on b.UserId = a.UserId left join SecurityRole c on a.RoleId = c.RoleId
#### UserId   RoleId  必须出现在查询语句中

全量同步：curl 127.0.0.1:8081/etl/es7/SecurityUser.yml -X POST
```
PUT /security-user
{
  "aliases" : {
        "SecurityUser" : {}
    },
  "settings" : {
      "number_of_shards" : 1, 
      "number_of_replicas" : 1
   },
  "mappings":{
      "properties":{       
		 "UserId": {
          "type": "long"
        },
        "CharacterId": {
          "type": "long"
        },
		"LoginId": {
          "type": "keyword"
        },
        "TrueName": {
          "type": "text",
		  "fields": {
           "keyword": {
             "type": "keyword",
             "ignore_above": 50
           }
		 }
        },
		"UniteNote": {
          "type": "text",
		  "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
		"RoleName": {
          "type": "keyword"
        },
        "CreateTime": {
          "type": "date"
        }
      }
  }
}
```
### 父子表(包含不同库之间)
xxorderBasic  xxorderItem JunkBasic
关联从表如果是子查询不能有多张表 

全量同步：
curl 127.0.0.1:8081/etl/es7/xxorderBasic.yml -X POST
curl 127.0.0.1:8081/etl/es7/xxorderItem.yml -X POST
```
PUT /xx-order
{
  "aliases" : {
        "xxorder" : {}
    },
  "settings" : {
      "number_of_shards" : 5, 
      "number_of_replicas" : 1
   },
  "mappings":{
      "properties":{       
		 "PrintId": {
          "type": "long"
        },
        "CharacterId": {
          "type": "long"
        },
		"PrintCode": {
          "type": "keyword"
        },
        "MaterialTypes": {
          "type": "text"
        },
		"CheckNote": {
          "type": "text",
		  "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
        "CreateTime": {
          "type": "date"
        },
		 "Item_JunkId": {
          "type": "long"
        },
		 "Item_PrintQty": {
          "type": "long"
        },
		"basic_item":{
          "type":"join",
          "relations":{
            "basic": ["item"]
          }
        }
      }
  }
}
```
写入父表
```
PUT /my-index/_doc/1?refresh
{
  "basic-text": "This is a parent document.",
  "basic_item": "basic"
}
```
写入子表
```
PUT /my-index/_doc/2?routing=1
{
  "join": {
    "name": "item",
    "parent": "1"
  },
  "item-text": "This is a item document."
}
```
## 重建索引

https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html

POST /symbol/_close

POST /symbol/_open

1. 新增新索引v2
```
PUT /goods-stock-log-v2
{
  "settings" : {
      "number_of_shards" : 5, 
      "number_of_replicas" : 1
   },
  "mappings":{
      "properties":{
        "LogId": {
          "type": "keyword"
        },
		 "OrderType": {
          "type": "long"
        },
        "CompchterId": {
          "type": "long"
        },
        "WarehouseCode": {
          "type": "text",
		  "fields": {
           "keyword": {
             "type": "keyword",
             "ignore_above": 50
           }
		 }
        },       
        "xx_name": {
          "type": "keyword"
        },
		"BarCode": {
          "type": "keyword"
        },
		"GoodsName": {
          "type": "text",
		  "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
        "CreateTime": {
          "type": "date"
        }
      }
  }
}
```
2. 修改alias别名的指向
```
POST /_aliases
{
    "actions": [
        { "remove": {
            "alias": "xx_index",
            "index": "goods-stock-log"
        }},
        { "add": {
            "alias": "xx_index",
            "index": "goods-stock-log-v2"
        }}
    ]
}
```
3. 索引迁移 3分钟400W行900MB  10分钟1100W行 2.7GB    25分钟同步完成2200W行5.5GB

但如果新的index中有数据，并且可能发生冲突，那么可以设置version_type"version_type": "internal"或者不设置，则Elasticsearch强制性的将文档转储到目标中，覆盖具有相同类型和ID的任何内容：
```
POST _reindex
{
  "source": {
    "index": "old_index"
  },
  "dest": {
    "index": "new_index",
    "version_type": "internal"
  }
}
```
```
POST /_reindex
{
  "source": {"index": "goods-stock-log"},
  "dest": {"index": "goods-stock-log-v2"}
}
```
4. 索引删掉(可选)
DELETE /goods-stock-log

[数据迁移效率](https://www.cnblogs.com/Ace-suiyuan008/p/9985249.html)

[Elasticsearch索引迁移工具](https://blog.csdn.net/laoyang360/article/details/65449407)
- [elasticsearch-dump](https://github.com/taskrabbit/elasticsearch-dump)
- [Elasticsearch-Exporter](https://github.com/mallocator/Elasticsearch-Exporter)
- logstash
- [elasticsearch-migration](https://github.com/medcl/esm-v1)

## 查询
https://www.elastic.co/guide/en/elasticsearch/reference/current/term-level-queries.html
### match_all
```
GET /xx_index/_search
{
  "query": { "match_all": {} }
}
```
### match
```
GET /xx_index/_search
{
  "query": { "match": { "xx_name": "半成品" } }
}
```
###  短语匹配 match_phrase
```
GET /xx_index/_search
{
  "query": { "match_phrase": { "GoodsName": "半 成品" } }
}
```
### 短语匹配 match_phrase_prefix
```
GET /xx_index/_search
{
  "query": { "match_phrase_prefix": { "GoodsName": {
    "query": "johnnie walker bl",
        "slop":5,
        "max_expansions": 1
  } } }
}
```
这种查询的行为与 match_phrase 查询一致，不同的是它将查询字符串的最后一个词作为前缀使用，换句话说，可以将之前的例子看成如下这样：

johnnie

跟着 walker

跟着以 bl 开始的词
其它参数
https://github.com/elasticsearch-cn/elasticsearch-definitive-guide/blob/cn/130_Partial_Matching/20_Match_phrase_prefix.asciidoc#slop

### exists

### fuzzy相似查询

### ids
```
GET /_search
{
    "query": {
        "ids" : {
            "values" : ["1", "4", "100"]
        }
    }
}
```
### range
```
GET _search
{
    "query": {
        "range" : {
            "age" : {
                "gte" : 10,
                "lte" : 20,
                "boost" : 2.0
            }
        }
    }
}
```
### term
```
GET /_search
{
    "query": {
        "term": {
            "user": {
                "value": "Kimchy",
                "boost": 1.0
            }
        }
    }
}
```
### terms
```
GET /_search
{
    "query" : {
        "terms" : {
            "user" : ["kimchy", "elasticsearch"],
            "boost" : 1.0
        }
    }
}
```
### terms_set
```
GET /job-candidates/_search
{
    "query": {
        "terms_set": {
            "programming_languages": {
                "terms": ["c++", "java", "php"],
                "minimum_should_match_field": "required_matches"
            }
        }
    }
}
```
### bool
https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html

### 前缀搜索 prefix
扫描整个倒排索引
前缀越短，要处理的doc越多，性能越差，尽可能用长前缀搜索
```
GET /xx_index/_search
{
  "query": { "prefix": { "xx_name": "半成" } }
}
```

### 通配符搜索 wildcard
?：任意字符
*：0个或任意多个字符
性能一样差，必须扫描整个倒排索引
```
GET /xx_index/_search
{
  "query": { "wildcard": { "xx_name": "?成品" } }
}
```
### 正则搜索 regexp

[0-9]：指定范围内的数字
[a-z]：指定范围内的字母
.：一个字符
+：前面的正则表达式可以出现一次或多次
性能一样差，必须扫描整个倒排索引

### 分页
```
GET /xx-order/_search
{
  "query": { "match_all": {} },
  "size": 20,
  "from": 30
}
```
### 排序
```
GET /xx-order/_search
{
  "query": { "match_all": {} },
  "sort":  [
    { "PrintId": "asc" }
  ]
}
```
### array
```
GET /xx-order/_search
{
  "query": { "match": { "MaterialTypes": "27" } },
  "sort":  [
    { "PrintId": "asc" }
  ]
}
```
### 父子
#### parent_id返回子表
查询printid=53094的所有item
```
GET /xx-order/_search
{
  "query": {
      "parent_id": {
          "type": "item",
          "id": "53094"
      }
  }
}
```
#### has_child返回父表
```
GET /xx-order/_search
{
    "query": {
        "has_child" : {
            "type" : "item",
            "query" : {
                "match_all" : {}
            },
            "max_children": 10,
            "min_children": 2,
            "score_mode" : "min"
        }
    }
}
```
##### inner_hits返回父表带子表
```
GET /xx-order/_search
{
    "query": {
        "has_child" : {
            "type" : "item",
            "query" : {
                "match_all" : {}
            },
            "max_children": 10,
            "min_children": 2,
            "score_mode" : "min"
        },
		"inner_hits": {}
    }
}
```
#### has_parent返回子表
```
GET /xx-order/_search
{
    "query": {
        "has_parent" : {
            "parent_type" : "basic",
            "query" : {
                "match_all" : {}
            }
        }
    }
}
```
#### agg
https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-children-aggregation.html


### nested
https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-nested-query.html


### _mget
https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html

### sql
```
POST /_sql?format=json
{
    "query": "SELECT * FROM xx_index WHERE xx_name LIKE '半成品'"
}
```

### _update_by_query _delete_by_query
https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html

```
删除字段
POST /index/_update/1
{
    "script" : "ctx._source.remove(\"State\")"
}

更新字段
POST /index/_update_by_query
{
  "script": {
    "source": "ctx._source['State']=ctx._source['Id']" 
  }
}
```

