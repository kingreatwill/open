<!--toc-->
[TOC]


# BigData + AI

## [现代流式计算的基石：Google DataFlow](https://yq.aliyun.com/articles/688767)

### 2. 核心概念
#### 2.1 Unbounded/Bounded vs Streaming/Batch
在 Dataflow 之前，对于有限/无限数据集合的描述，一般使用批/流 （Batch/Streaming），总有一种暗示底层两套引擎（批处理引擎和流处理引擎）。对于批处理和流处理，一般情况下是可以互相转化的，比如 Spark 用微批来模拟流。而 Dataflow 模型一般将有限/无限数据集合称为 Bounded/Unbounded Dataset，而 Streaming/Batch 用来特指执行引擎。

#### 2.2 Window
Window，也就是窗口，将一部分数据集合组合起操作。在处理无限数据集的时候有限操作需要窗口，比如 aggregation，outer join，time-bounded 操作。窗口大部分都是基于时间来划分，但是也有基于其他存在逻辑上有序关系的数据来划分的。窗口模型主要由三种：Fixed Window，Sliding Window，Session Window。

##### 1. Fixed Window
Fixed Window ，有时候也叫 Tumbling Window。Tumble 的中文翻译有“翻筋斗”，我们可以将 Fixed Window 是特定的时间长度在无限数据集合上翻滚形成的，核心是每个 Window 没有重叠。比如小时窗口就是 12:00:00 ~ 13:00:00 一个窗口，13:00:00 ~ 14:00:00 一个窗口。从例子也可以看出来 Fixed Window 的另外一个特征：aligned，中文一般称为对齐。可能有些人还是不太明白。那么我举一个在编程语言中一个例子：address alignment，内存地址a被称为n字节对齐，当a是n的倍数（n应是2的幂）。但是有时候处于某些目的，窗口也可以是不对齐的。

##### 2. Sliding Window
Sliding Window，中文可以叫滑动窗口，由两个参数确定，窗口大小和滑动间隔。比如每分钟开始一个小时窗口对应的就是窗口大小为一小时，滑动间隔为一分钟。滑动间隔一般小于窗口大小，也就是说窗口之间会有重叠。滑动窗口在很多情况下都比较有用，比如检测机器的半小时负载，每分钟检测一次。Fixed Window 是 Sliding Window 的一种特例：窗口大小等于滑动间隔。

##### 3. Session Window
Session Window，中文可以叫会话窗口， 一般用来捕捉一段时间内的行为，比如 Web 中一段时间内的登录行为为一个 Session，当长时间没有登录，则 Session 失效，再次登录重启一个 Session。Session Window 也是用超时时间来衡量，只要在超时时间内发生的事件都认为是一个 Session Window。

#### 2.3 Time Domain
在流式处理中关于时间有两个概念需要注意：

- Event Time，事件发生的时间。
- Processing TIme，事件在系统中的处理时间。

这两个概念非常简单。比如在 IoT 中，传感器采集事件时对应的系统时间就是 Event Time，然后事件发送到流式系统进行处理，处理的时候对应的系统时间就是 Processing Time。虽然是两个很简单概念，但是在 Dataflow 模型之前，很多系统并没有显示区分，比如 Spark Streaming。

在现实中，由于通信延迟、调度延迟等，往往导致 Event Time 和 Processing Time 之间存在差值（skew），且动态变化。skew 一般使用 watermark 来进行可视化，如下图。
![](img/event-processing-time.png)

## [Apache Spark3.0什么样？一文读懂Apache Spark最新技术发展与展望](https://yq.aliyun.com/articles/712303)

### Databricks、Delta Lake

### 三、Spark与AI框架深度集成的最新进展
机器学习任务的完整链条非常长，包括数据的收集、落地、清理、准备，以及模型的训练、检验以及预测等。Spark比较擅长该链条比较靠前的部分，即数据的采集、落地、清理和准备。深度学习框架擅长链条靠后这部分，即模型的训练、验证和分析。用户在完成一个深度学习任务时，需要基于比如Spark和Tensorflow这两套系统，这就给用户带来了很多不便之处，比如两套系统的部署、运维、管理等。另外，这也使得整个任务的开发、Debug、Troubleshooting变得困难。

## [Spark入门实战系列--1.Spark及其生态圈简介](https://www.cnblogs.com/shishanyuan/p/4700615.html)

Spark常用术语

| 术语   | 描述      |
|--------------------|-----------|
| Application      | Spark的应用程序，包含一个Driver program和若干Executor  |
| SparkContext     | Spark应用程序的入口，负责调度各个运算资源，协调各个Worker Node上的Executor                                                 |
| Driver Program   | 运行Application的main()函数并且创建SparkContext   |
| Executor         | 是为Application运行在Worker node上的一个进程，该进程负责运行Task，并且负责将数据存在内存或者磁盘上。每个Application都会申请各自的Executor来处理任务 |
| Cluster Manager  | 在集群上获取资源的外部服务(例如：Standalone、Mesos、Yarn)  |
| Worker Node      | 集群中任何可以运行Application代码的节点，运行一个或多个Executor进程   |
| Task | 运行在Executor上的工作单元  |
| Job  | SparkContext提交的具体Action操作，常和Action对应  |
| Stage  | 每个Job会被拆分很多组task，每组任务被称为Stage，也称TaskSet  |
| RDD | 是Resilient distributed datasets的简称，中文为弹性分布式数据集;是Spark最核心的模块和类  |
| DAGScheduler     | 根据Job构建基于Stage的DAG，并提交Stage给TaskScheduler   |
| TaskScheduler    | 将Taskset提交给Worker node集群运行并返回结果   |
| Transformations  | 是Spark API的一种类型，Transformation返回值还是一个RDD，所有的Transformation采用的都是懒策略，如果只是将Transformation提交是不会执行计算的 |
| Action | 是Spark API的一种类型，Action返回值不是一个RDD，而是一个scala集合；计算只有在Action被提交的时候计算才被触发。  |
