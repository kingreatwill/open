
## Manticore Search
https://github.com/manticoresoftware/manticoresearch

Manticore Search 是一个使用 C++ 开发的高性能搜索引擎，创建于 2017 年，其前身是 [Sphinx Search](https://github.com/sphinxsearch/sphinx), 
其商业公司不继续开源了。然后一帮人在之前Sphinx Search的基础上开发了Manticore Search，重写了很多功能代码，修复bug无数，这个产品号称速度超过mysql 182倍，超过ES 4到5倍，吞吐量比ES高两倍，考虑到ES也在性能进步，这个评测不知道是不是属实。

### 特点

- Manticore是基于SQL的，使用SQL作为其本机语法，并与MySQL协议兼容，使您可以使用首选的MySQL客户端。
它非常快，因此比其他替代方案更具成本效益。例如，Manticore:
- 速度快
    - 对于小型数据，比MySQL快182倍（可重现）
    - 对于日志分析，比Elasticsearch快29倍（可重现）
    - 对于小型数据集，比Elasticsearch快15倍（可重现）
    - 对于中等大小的数据，比Elasticsearch快5倍（可重现）
    - 对于大型数据，比Elasticsearch快4倍（可重现）
    - 在单个服务器上进行数据导入时，最大吞吐量比Elasticsearch快最多2倍（可重现）

- 由于其现代的多线程架构和高效的查询并行化能力，Manticore能够充分利用所有CPU核心，以实现最快的响应时间。

- 强大而快速的全文搜索功能能够无缝地处理小型和大型数据集。

- 针对小、中、大型数据集提供逐行存储。

- 对于更大的数据集，Manticore通过Manticore Columnar Library提供列存储支持，可以处理无法适合内存的数据集。

- 自动创建高效的二级索引，节省时间和精力。

- 成本优化的查询优化器可优化搜索查询以实现最佳性能。

- 通过PHP、Python、JavaScript、Java、Elixir和Go等客户端，与Manticore Search的集成变得简单。

- Manticore还提供了一种编程HTTP JSON协议，用于更多样化的数据和模式管理。

- Manticore Search使用C++构建，启动快速，内存使用最少，低级别优化有助于其卓越性能。

- 实时插入，新添加的文档立即可访问。

- 提供互动课程，使学习轻松愉快。

- Manticore还拥有内置的复制和负载均衡功能，增加了可靠性。

- 可以轻松地从MySQL、PostgreSQL、ODBC、xml和csv等来源同步数据。

- 虽然不完全符合ACID，但Manticore仍支持事务和binlog以确保安全写入。

- 内置工具和SQL命令可轻松备份和恢复数据。


## 其它ES替代品
### OpenSearch
Powered by [Apache Lucene](https://lucene.apache.org/)
一个由社区共同推动的 Elasticsearch 与 Kibana 开源分支。
为用户提供一个安全、高效、全面开源的搜索与分析套件，同时建立包含丰富创新的功能发展路线图。此项目分为 [OpenSearch](https://github.com/opensearch-project/OpenSearch)（源自 Elasticsearch 7.10.2）与 [OpenSearch Dashboards](https://github.com/opensearch-project/OpenSearch-Dashboards)（源自 Kibana 7.10.2）两部分。
https://opensearch.org/

### meilisearch
一个轻量级搜索引擎
A lightning-fast search engine that fits effortlessly into your apps, websites, and workflow. 
https://github.com/meilisearch/meilisearch Rust语言,38.1k


```sh
# Fetch the latest version of Meilisearch image from DockerHub
docker pull getmeili/meilisearch:v1.3

# Launch Meilisearch in development mode with a master key
docker run -it --rm \
    -p 7700:7700 \
    -e MEILI_ENV='development' \
    -v $(pwd)/meili_data:/meili_data \
    getmeili/meilisearch:v1.3
# Use ${pwd} instead of $(pwd) in PowerShell
```

当MeiliSearch运行起来后，默认会在7700端口暴露http接口，后续所有的访问，包括新增数据、搜索数据等都是通过这个http接口。另外启动之后，官方还自带了一个web界面，不过这个界面只是用来测试的，在生产环境会被关闭掉。然后你可以在这个界面试用MeiliSearch强大的搜索功能。
https://www.meilisearch.com/docs/learn/getting_started/quick_start#add-documents

