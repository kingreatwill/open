https://learnku.com/elasticsearch

分片个数计算

SN(分片数) = IS(索引大小) / 30

NN(节点数) = SN(分片数) + MNN(主节点数[无数据]) + NNN(负载节点数)

假设有300G大小的数据文件

- 每一个分片数据文件小于30GB
- 每一个索引中的一个分片对应一个节点
- 节点数大于等于分片数

![](img/es-lucene.jpg)

在Lucene中，分为索引(录入)与检索(查询)两部分，索引部分包含分词器、过滤器、字符映射器等，检索部分包含查询解析器等。

一个Lucene索引包含多个segments，一个segment包含多个文档，每个文档包含多个字段，每个字段经过分词后形成一个或多个term。

### 优化
https://www.elastic.co/guide/en/elasticsearch/reference/current/tune-for-indexing-speed.html

1. 关闭不需要字段的doc values。

2. 尽量使用keyword替代一些long或者int之类，term查询总比range查询好 (参考lucene说明 http://lucene.apache.org/core/7_4_0/core/org/apache/lucene/index/PointValues.html)。

3. 关闭不需要查询字段的_source功能，不将此存储仅ES中，以节省磁盘空间。

4. 评分消耗资源，如果不需要可使用filter过滤来达到关闭评分功能，score则为0，如果使用constantScoreQuery则score为1。

5. 关于分页：

from + size: 每分片检索结果数最大为 from + size，假设from = 20, size = 20，则每个分片需要获取20 * 20 = 400条数据，多个分片的结果在协调节点合并(假设请求的分配数为5，则结果数最大为 400*5 = 2000条) 再在内存中排序后然后20条给用户。这种机制导致越往后分页获取的代价越高，达到50000条将面临沉重的代价，默认from + size默认如下：index.max_result_window ：10000

search_after: 使用前一个分页记录的最后一条来检索下一个分页记录，在我们的案例中，首先使用from+size，检索出结果后再使用search_after，在页面上我们限制了用户只能跳5页，不能跳到最后一页。

scroll 用于大结果集查询，缺陷是需要维护scroll_id

6. 关于排序：我们增加一个long字段，它用于存储时间和ID的组合(通过移位即可)，正排与倒排性能相差不明显。

7. 关于CPU消耗，检索时如果需要做排序则需要字段对比，消耗CPU比较大，如果有可能尽量分配16cores以上的CPU，具体看业务压力。

8. 关于合并被标记删除的记录，我们设置为0表示在合并的时候一定删除被标记的记录，默认应该是大于10%才删除："merge.policy.expunge_deletes_allowed": "0"。
