# 索引模板

在创建索引时，为每个索引写定义信息可能是一件繁琐的事情，ES提供了索引模板功能，让你可以定义一个索引模板，模板中定义好settings、mapping、以及一个模式定义来匹配创建的索引。

注意：模板只在索引创建时被参考，修改模板不会影响已创建的索引

### 新增/修改名为tempae_1的模板，匹配名称为te* 或 bar*的索引创建：
```
PUT _template/template_1
{
  "index_patterns": ["te*", "bar*"],
  "settings": {
    "number_of_shards": 1
  },
  "mappings": {
    "type1": {
      "_source": {
        "enabled": false
      },
      "properties": {
        "host_name": {
          "type": "keyword"
        },
        "created_at": {
          "type": "date",
          "format": "EEE MMM dd HH:mm:ss Z YYYY"
        }
      }
    }
  }
}
```
### 查看索引模板
```
GET /_template/template_1
GET /_template/temp* 
GET /_template/template_1,template_2
GET /_template
```

### 删除模板
```
DELETE /_template/template_1
```