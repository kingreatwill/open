_msearch
_scripts

PUT /my-index
{
  "aliases" : {
        "my-index-ali" : {}
    },
  "settings" : {
      "number_of_shards" : 5, 
      "number_of_replicas" : 1
   },
  "mappings":{
      "properties":{
        "uuid": {
          "type": "keyword"
        },
		 "enum": {
          "type": "long"
        },
        "id": {
          "type": "long"
        },
        "code": {
          "type": "keyword"
        },       
        "name": {
          "type": "keyword"
        }
		"note": {
          "type": "text",
		  "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
        "time": {
          "type": "date"
        }
      }
  }
}

## match和match_phrase
GET /my-index-ali/_search
{
  "query": { "match": { "name": "半成品" } }
}

## 前缀搜索的原理
扫描整个倒排索引
前缀越短，要处理的doc越多，性能越差，尽可能用长前缀搜索

GET /my-index-ali/_search
{
  "query": { "prefix": { "name": "半成" } }
}

## 通配符搜索
?：任意字符
*：0个或任意多个字符
性能一样差，必须扫描整个倒排索引
GET /my-index-ali/_search
{
  "query": { "wildcard": { "name": "?成品" } }
}

## 正则搜索
regexp
[0-9]：指定范围内的数字
[a-z]：指定范围内的字母
.：一个字符
+：前面的正则表达式可以出现一次或多次
性能一样差，必须扫描整个倒排索引

## 相似搜索
GET /my-index-ali/_search
{
  "query": { "fuzzy": { "name": "半成日" } }
}

## 搜索推荐的功能
```
GET my-index/_search
{
  "query": {
    "match_phrase_prefix": {
      "content": {
        "query": "i elas",
        "slop":5,
        "max_expansions": 1
      }
    }
  }
}
```
原理跟match_phrase类似，唯一的区别，就是把最后一个term作为前缀去搜索

i就是去进行match，搜索对应的doc
elas，会作为前缀，去扫描整个倒排索引，找到所有elas开头的doc
然后找到所有doc中，即包含i，又包含elas开头的字符的doc
根据你的slop去计算，看在slop范围内，能不能让hello w，正好跟doc中的hello和w开头的单词的position相匹配
也可以指定slop，但是只有最后一个term会作为前缀

max_expansions：指定prefix最多匹配多少个term，超过这个数量就不继续匹配了，限定性能

默认情况下，前缀要扫描所有的倒排索引中的term，去查找w打头的单词，但是这样性能太差。可以用max_expansions限定，w前缀最多匹配多少个term，就不再继续搜索倒排索引了。

尽量不要用，因为，最后一个前缀始终要去扫描大量的索引，性能可能会很差

### 查询时输入即搜索

把邮编的事情先放一边，让我们先看看前缀查询是如何在全文查询中起作用的。((("partial matching", "query time search-as-you-type")))用户已经渐渐习惯在输完查询内容之前，就能为他们展现搜索结果，这就是所谓的 _即时搜索（instant search）_ 或 _输入即搜索（search-as-you-type）_ 。((("search-as-you-type")))((("instant search")))不仅用户能在更短的时间内得到搜索结果，我们也能引导用户搜索索引中真实存在的结果。

例如，如果用户输入 `johnnie walker bl` ，我们希望在它们完成输入搜索条件前就能得到：Johnnie Walker Black Label 和 Johnnie Walker Blue Label 。

生活总是这样，就像猫的花色远不只一种！我们希望能找到一种最简单的实现方式。并不需要对数据做任何准备，在查询时就能对任意的全文字段实现 _输入即搜索（search-as-you-type）_ 的查询。

在 <<phrase-matching,短语匹配>> 中，我们引入了 `match_phrase` 短语匹配查询，它匹配相对顺序一致的所有指定词语，对于查询时的输入即搜索，可以使用 `match_phrase` 的一种特殊形式，((("prefix query", "match_phrase_prefix query")))((("match_phrase_prefix query"))) `match_phrase_prefix` 查询：


```
{
    "match_phrase_prefix" : {
        "brand" : "johnnie walker bl"
    }
}
```
// SENSE: 130_Partial_Matching/20_Match_phrase_prefix.json

这种查询的行为与 `match_phrase` 查询一致，不同的是它将查询字符串的最后一个词作为前缀使用，换句话说，可以将之前的例子看成如下这样：

* `johnnie`
* 跟着 `walker`
* 跟着以 `bl` 开始的词

如果通过 `validate-query` API 运行这个查询查询，explanation 的解释结果为：

    "johnnie walker bl*"

与 `match_phrase` 一样，它也可以接受 `slop` 参数（参照 <<slop,slop>> ）让相对词序位置不那么严格：((("slop parameter", "match_prhase_prefix query")))((("match_phrase_prefix query", "slop parameter")))

```
{
    "match_phrase_prefix" : {
        "brand" : {
            "query": "walker johnnie bl", <1>
            "slop":  10
        }
    }
}
```
// SENSE: 130_Partial_Matching/20_Match_phrase_prefix.json

<1> 尽管词语的顺序不正确，查询仍然能匹配，因为我们为它设置了足够高的 `slop` 值使匹配时的词序有更大的灵活性。

但是只有查询字符串的最后一个词才能当作前缀使用。

在之前的 <<prefix-query,前缀查询>> 中，我们警告过使用前缀的风险，即 `prefix` 查询存在严重的资源消耗问题，短语查询的这种方式也同样如此。((("match_phrase_prefix query", "caution with")))前缀 `a` 可能会匹配成千上万的词，这不仅会消耗很多系统资源，而且结果的用处也不大。

可以通过设置 `max_expansions` 参数来限制前缀扩展的影响，((("match_phrase_prefix query", "max_expansions")))((("max_expansions parameter")))一个合理的值是可能是 50 ：

```
{
    "match_phrase_prefix" : {
        "brand" : {
            "query":          "johnnie walker bl",
            "max_expansions": 50
        }
    }
}
```
// SENSE: 130_Partial_Matching/20_Match_phrase_prefix.json

参数 `max_expansions` 控制着可以与前缀匹配的词的数量，它会先查找第一个与前缀 `bl` 匹配的词，然后依次查找搜集与之匹配的词（按字母顺序），直到没有更多可匹配的词或当数量超过 `max_expansions` 时结束。

不要忘记，当用户每多输入一个字符时，这个查询又会执行一遍，所以查询需要快，如果第一个结果集不是用户想要的，他们会继续输入直到能搜出满意的结果为止。