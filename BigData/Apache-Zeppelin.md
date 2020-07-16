# Apache Zeppelin
https://github.com/apache/zeppelin  4.8k
## 简介
Apache Zeppelin 是一个可以进行大数据可视化分析的交互式开发系统，可以承担数据接入、数据发现、数据分析、数据可视化、数据协作等任务，其前端提供丰富的可视化图形库，不限于SparkSQL，后端支持HBase、Flink 等大数据系统以插件扩展的方式，并支持Spark、Python、JDBC、Markdown、Shell 等各种常用Interpreter，这使得开发者可以方便地使用SQL 在 Zeppelin 中做数据开发。

对于机器学习算法工程师来说，他们可以在 Zeppelin 中可以完成机器学习的数据预处理、算法开发和调试、算法作业调度的工作，包括当前在各类任务中表现突出的深度学习算法，因为 Zeppelin 的最新的版本中增加了对TensorFlow、PyTorch等主流深度学习框架的支持，此外，Zeppelin将来还会提供算法的模型 Serving 服务、Workflow 工作流编排等新特性，使得 Zeppelin可以完全覆盖机器学习的全流程工作。

而在平台部署和运维方面，Zeppelin还提供了单机 Docker、分布式、K8s、Yarn 四种系统运行模式，无论你是小规模的开发团队，还是 Hadoop 技术栈的大数据团队、K8s 技术栈的云计算团队，Zeppelin 都可以让数据科学团队轻松的进行部署和使用 Zeppelin丰富的数据和算法的开发能力。

## Zeppelin和Flink的故事
Flink问：虽然我提供了多种语言支持，有SQL，Java，Scala还有Python，但是每种语言都有自己的入口，用户很难多种语言混着用。比如在sql-client中只能运行Sql，不能写UDF，在pyflink shell里，只能用python的udf，不能用scala和java的udf。有没有谁能帮我把这些语言全部打通。

Zeppelin答：我可以。



Flink问：我的一个很大的使用场景是实时大屏，但是我一个人办不到，往往需要借助第三方存储，还需要前端开发，有没有谁能让用户不用写前端代码就实现实时大屏

Zeppelin答：我可以。



Flink问：我的Sql已经很强大了，但是用户在sql-client里不能写comment，而且不支持运行多条sql语句，有谁能帮我把这些功能补齐下。

Zeppelin答：我可以。



Flink问：好多初学者说要跑一个flink job实在是太难了，好多东西需要配置,还要学习各种命令行，有没有谁能让用户更容易得提交和管理Flink Job。

Zeppelin答：我可以。



Flink问：Flink Job提交目前只能一个个提交，一个job跑完跑另外一个，有些用户想并行执行多个Flink Job，谁能帮我搞定这个需求？

Zeppelin答：我可以。



Flink问：我有丰富的connector，但是用户每次都要把connector打包到uber jar里，或者copy到flink的lib下，但是这样会把各种connector jar混在一起，容易发生冲突，很难管理，有谁能提供一个干净点的方案？

Zeppelin答：我可以。

