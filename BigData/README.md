https://github.com/heibaiying/BigData-Notes


https://blog.csdn.net/qq_32641659/article/details/98729045


https://www.kaggle.com/

框架 | 语言|当前版本(2019/12/19)|star
--|--|--|--
[Apache Hadoop](https://github.com/apache/hadoop) |java| v2.10|10k
[Apache Spark](https://github.com/apache/spark)|scala，java|v3.0.0|24.6k
[Apache Flink](https://github.com/apache/flink) |java，scala |V1.9.1|11.5k
[Apache Storm](https://github.com/apache/storm)|java |V2.1.0|6k

Hadoop(第一代) - > Spark(第二代) - > Flink(第三代)

[Spark](https://spark.apache.org/) 生态更为丰富,在机器学习整合方面投入更多
[Flink](https://flink.apache.org/) 在流计算上有明显优势，核心架构和模型也更透彻和灵活一些。

总而言之，Flink与Spark没有谁强谁弱，只有哪个更适合当前的场景。

Spark与Flink API情况
API| Spark | Flink
--|--|--
底层API | RDD |Process Function
核心API  | DataFrame/DataSet/Structured Streaming | DataStream/DataSet
SQL | Spark SQL | Table API & SQL
机器学习 | MLlib | FlinkML
图计算 | GraphX | Gelly
其它 |  | CEP



Spark与Flink 对开发语言的支持

支持语言| Spark | Flink
--|--|--
Java|√|√
Scala|√|√
Python|√|√
R|√|第三方
SQL|√|√

Flink VS Spark 之 Connectors
https://ci.apache.org/projects/flink/flink-docs-release-1.9/zh/dev/connectors/
https://spark.apache.org/

支持| Spark | Flink
--|--|--
Hbase           |√  |√
HDFS            |√  |√
PostgreSQL      |√  |
Mysql           |√  |
Elastic         |√  |√
kafka           |√  |√
redis           |√  |
cassandra       |√  |√
mongoDB         |√  |
Alluxio         |√  |
Hive            |√  |
hundreds of other data sources|√  |
RabbitMQ        |   |√
Amazon Kinesis Streams|   |√
Apache NiFi     |   |√
Twitter Streaming API |   |√





Flink VS Spark 之 运行环境

部署环境| Spark | Flink
--|--|--
Local(single JVM) | √|√
Standalone |√|√
Yarn |√|√
Mesos |√|√
Kubernetes |√|√

Flink 流处理
Storm 流处理
Spark 微批处理
Hadoop 批处理


故障容忍（Apache Storm）、低延迟（Apache Flink、Storm），可操作性（Twitter Heron）、直观编程模型（Millwheel）、语义处理（Dataflow、Samza、Flink）、弹性伸缩（Millwheel），有效资源管理（Twitter Heron）和状态管理（Spark Streaming）





----------------
数据挖掘工作平台 Weka




在大数据时代，数据挖掘是最关键的工作。大数据的挖掘是从海量、不完全的、有噪声的、模糊的、随机的大型数据库中发现隐含在其中有价值的、潜在有用的信息和知识的过程，也是一种决策支持过程。其主要基于人工智能，机器学习，模式学习，统计学等。通过对大数据高度自动化地分析，做出归纳性的推理，从中挖掘出潜在的模式，可以帮助企业、商家、用户调整市场政策、减少风险、理性面对市场，并做出正确的决策。目前，在很多领域尤其是在商业领域如银行、电信、电商等，数据挖掘可以解决很多问题，包括市场营销策略制定、背景分析、企业管理危机等。大数据的挖掘常用的方法有分类、回归分析、聚类、关联规则、神经网络方法、Web 数据挖掘等。这些方法从不同的角度对数据进行挖掘。


(1)分类。分类是找出数据库中的一组数据对象的共同特点并按照分类模式将其划分为不同的类，其目的是通过分类模型，将数据库中的数据项映射到摸个给定的类别中。可以应用到涉及到应用分类、趋势预测中，如淘宝商铺将用户在一段时间内的购买情况划分成不同的类，根据情况向用户推荐关联类的商品，从而增加商铺的销售量。

(2)回归分析。回归分析反映了数据库中数据的属性值的特性，通过函数表达数据映射的关系来发现属性值之间的依赖关系。它可以应用到对数据序列的预测及相关关系的研究中去。在市场营销中，回归分析可以被应用到各个方面。如通过对本季度销售的回归分析，对下一季度的销售趋势作出预测并做出针对性的营销改变。

(3)聚类。聚类类似于分类，但与分类的目的不同，是针对数据的相似性和差异性将一组数据分为几个类别。属于同一类别的数据间的相似性很大，但不同类别之间数据的相似性很小，跨类的数据关联性很低。

(4)关联规则。关联规则是隐藏在数据项之间的关联或相互关系，即可以根据一个数据项的出现推导出其他数据项的出现。关联规则的挖掘过程主要包括两个阶段：第一阶段为从海量原始数据中找出所有的高频项目组;第二极端为从这些高频项目组产生关联规则。关联规则挖掘技术已经被广泛应用于金融行业企业中用以预测客户的需求，各银行在自己的ATM 机上通过捆绑客户可能感兴趣的信息供用户了解并获取相应信息来改善自身的营销。

(5)神经网络方法。神经网络作为一种先进的人工智能技术，因其自身自行处理、分布存储和高度容错等特性非常适合处理非线性的以及那些以模糊、不完整、不严密的知识或数据为特征的处理问题，它的这一特点十分适合解决数据挖掘的问题。典型的神经网络模型主要分为三大类：第一类是以用于分类预测和模式识别的前馈式神经网络模型，其主要代表为函数型网络、感知机;第二类是用于联想记忆和优化算法的反馈式神经网络模型，以Hopfield 的离散模型和连续模型为代表。第三类是用于聚类的自组织映射方法，以ART 模型为代表。虽然神经网络有多种模型及算法，但在特定领域的数据挖掘中使用何种模型及算法并没有统一的规则，而且人们很难理解网络的学习及决策过程。

(6)Web数据挖掘。Web数据挖掘是一项综合性技术，指Web 从文档结构和使用的集合C 中发现隐含的模式P，如果将C看做是输入，P 看做是输出，那么Web 挖掘过程就可以看做是从输入到输出的一个映射过程。

当前越来越多的Web 数据都是以数据流的形式出现的，因此对Web 数据流挖掘就具有很重要的意义。目前常用的Web数据挖掘算法有：PageRank算法，HITS算法以及LOGSOM 算法。这三种算法提到的用户都是笼统的用户，并没有区分用户的个体。目前Web 数据挖掘面临着一些问题，包括：用户的分类问题、网站内容时效性问题，用户在页面停留时间问题，页面的链入与链出数问题等。在Web 技术高速发展的今天，这些问题仍旧值得研究并加以解决。

