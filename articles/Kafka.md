https://www.jianshu.com/p/b4cccc6b6a45
https://www.jianshu.com/p/599b0b7ed21d
https://www.cnblogs.com/qingyunzong/p/9004509.html
https://www.toutiao.com/a6746186631276921357/
https://www.toutiao.com/a6760913909265203719
https://www.toutiao.com/a6768817833418686988
https://www.toutiao.com/a6741255916940706316
https://www.toutiao.com/a6768856487977550343


[TOC]

# Kafka

## 什么是Kafka

### kafka简介
Kafka是最初由Linkedin公司开发，是一个分布式、分区的、多副本的、多订阅者，基于zookeeper协调的分布式日志系统（也可以当做MQ系统），常见可以用于web/nginx日志、访问日志，消息服务等等，Linkedin于2010年贡献给了Apache基金会并成为顶级开源项目。

主要应用场景是：日志收集系统和消息系统。
Kafka主要设计目标如下：
- 以时间复杂度为O(1)的方式提供消息持久化能力，即使对TB级以上数据也能保证常数时间的访问性能。
- 高吞吐率。即使在非常廉价的商用机器上也能做到单机支持每秒100K条消息的传输。
- 支持Kafka Server间的消息分区，及分布式消费，同时保证每个partition内的消息顺序传输。
- 同时支持离线数据处理和实时数据处理。
- Scale out:支持在线水平扩展

Kafka是一个分布式的、可分区的、可复制的消息系统。
Kafka是一种快速、可扩展的、设计内在就是分布式的，分区的和可复制的提交日志服务。

有两种主要的消息传递模式：点对点传递模式、发布-订阅模式。
Kafka就是一种发布-订阅模式
![](../img/kafka/topic-1.png)
### kafka基本架构

![](../img/kafka/kafka-3.webp)

#### 话题（Topic）：
是特定类型的消息流（字节的有效负载），话题是消息的分类；

kafka中消息订阅和发送都是基于某个topic。
比如有个topic叫做NBA赛事信息，那么producer会把NBA赛事信息的消息发送到此topic下面。
所有订阅此topic的consumer将会拉取到此topic下的消息。
Topic就像一个特定主题的收件箱，producer往里丢，consumer取走。

#### 生产者（Producer）：
是能够发布消息到话题的任何对象；

生产者即数据的发布者，该角色将消息发布到Kafka的topic中。broker接收到生产者发送的消息后，broker将该消息追加到当前用于追加数据的segment文件中。生产者发送的消息，存储到一个partition中，生产者也可以指定数据存储的partition。

如果 Partition 没填，那么情况会是这样的：
- Key 有填。按照 Key 进行哈希，相同 Key 去一个 Partition。（如果扩展了 Partition 的数量那么就不能保证了）
- Key 没填。Round-Robin 来选 Partition。

#### 服务代理（Broker）：
已发布的消息保存在一组服务器中，它们被称为代理（Broker）或Kafka集群；

一个Borker就是Kafka集群中的一个实例，或者说是一个服务单元。
连接到同一个zookeeper的多个broker实例组成kafka的集群。
在若干个broker中会有一个broker是leader，其余的broker为follower。

---

broker存储topic的数据。如果某topic有N个partition，集群有N个broker，那么每个broker存储该topic的一个partition。

如果某topic有N个partition，集群有(N+M)个broker，那么其中有N个broker存储该topic的一个partition，剩下的M个broker不存储该topic的partition数据。

如果某topic有N个partition，集群中broker数目少于N个，那么一个broker存储该topic的一个或多个partition。在实际生产环境中，尽量避免这种情况的发生，这种情况容易导致Kafka集群数据不均衡。

#### 分区（Partition）：
topic中的数据分割为一个或多个partition。每个topic至少有一个partition。每个partition中的数据使用多个segment文件存储。partition中的数据是有序的，不同partition间的数据丢失了数据的顺序。如果topic有多个partition，消费数据时就不能保证数据的顺序。在需要严格保证消息的消费顺序的场景下，需要将partition数目设为1。

每个分区可以设置备份数量
分区由一个leader+多个followers组成，生产者直接与leader进行沟通，leader接收消息后，其他的followers会同步这个消息。所有的follwers同步消息后，该消息才会成为可消费的状态。
Broker中Topic与分区，分区与生产者，分区之间的选举备份等等信息都需要zookeeper进行协调。

---

每个Topic由多个分区组成，每个分区内部的数据保证了有序性，即是按照时间序列，append到分区的尾部。分区是有固定大小的，容量不够时，会创建新的分区。Kafka在一定时间内会定期清理过期的文件。

这种连续性的文件存储，一方面有效的利用磁盘的线性存取；另一方面减轻了内存的压力。

#### 消费者（Consumer）：
可以订阅一个或多个话题，并从Broker拉（pull）数据，从而消费这些已发布的消息；

push模式很难适应消费速率不同的消费者，因为消息发送速率是由broker决定的。push模式的目标是尽可能以最快速度传递消息，但是这样很容易造成consumer来不及处理消息，典型的表现就是拒绝服务以及网络拥塞。而pull模式则可以根据consumer的消费能力以适当的速率消费消息。


Kafka和其它消息系统有一个不一样的设计，在consumer之上加了一层group（Consumer Group）；
同一个group的consumer可以并行消费同一个topic的消息，但是同group的consumer，不会重复消费。
如果同一个topic需要被多次消费，可以通过设立多个consumer group来实现。每个group分别消费，互不影响。

#### Consumer Group
每个Consumer属于一个特定的Consumer Group（可为每个Consumer指定group name，若不指定group name则属于默认的group）。

很多传统的message queue都会在消息被消费完后将消息删除，一方面避免重复消费，另一方面可以保证queue的长度比较少，提高效率。而如上文所将，Kafka并不删除已消费的消息，为了实现传统message queue消息只被消费一次的语义，Kafka保证保证同一个consumer group里只有一个consumer会消费一条消息。与传统message queue不同的是，Kafka还允许不同consumer group同时消费同一条消息，这一特性可以为消息的多元化处理提供了支持。实际上，Kafka的设计理念之一就是同时提供离线处理和实时处理。根据这一特性，可以使用Storm这种实时流处理系统对消息进行实时在线处理，同时使用Hadoop这种批处理系统进行离线处理，还可以同时将数据实时备份到另一个数据中心，只需要保证这三个操作所使用的consumer在不同的consumer group即可。



Kafka的相关术语以及之间的关系1
![](../img/kafka/kafka-1.webp)
Kafka的相关术语以及之间的关系2
![](../img/kafka/kafka-1.png)
上图中一个topic配置了3个partition。Partition1有两个offset：0和1。Partition2有4个offset。Partition3有1个offset。副本的id和副本所在的机器的id恰好相同。

如果一个topic的副本数为3，那么Kafka将在集群中为每个partition创建3个相同的副本。集群中的每个broker存储一个或多个partition。多个producer和consumer可同时生产和消费数据。

#### Leader
每个partition有多个副本，其中有且仅有一个作为Leader，Leader是当前负责数据的读写的partition。

#### Follower
Follower跟随Leader，所有写请求都通过Leader路由，数据变更会广播给所有Follower，Follower与Leader保持数据同步。如果Leader失效，则从Follower中选举出一个新的Leader。当Follower与Leader挂掉、卡住或者同步太慢，leader会把这个follower从“in sync replicas”（ISR）列表中删除，重新创建一个Follower。

## kafka原理
我们将消息的发布（publish）称作 producer，
将消息的订阅（subscribe）表述为 consumer，
将中间的存储阵列称作 broker(代理);
![](../img/kafka/kafka-2.webp)
多个 broker 协同合作，
producer 和 consumer 部署在各个业务逻辑中被频繁的调用，
三者通过 zookeeper管理协调请求和转发。
这样一个高性能的分布式消息发布订阅系统就完成了。
![](../img/kafka/kafka-3.webp)
- producer 到 broker 的过程是 push，也就是有数据就推送到 broker；
- consumer 到 broker 的过程是 pull，是通过 consumer 主动去拉数据的，
而不是 broker 把数据主懂发送到 consumer 端的。

## Zookeeper在kafka的作用
- 无论是kafka集群，还是producer和consumer都依赖于zookeeper来保证系统可用性集群保存一些meta信息（kafka的配置，集群状态和连接信息等元数据）。
- Kafka使用zookeeper作为其分布式协调框架，很好的将消息生产、消息存储、消息消费的过程结合在一起。
- 借助zookeeper，kafka能够生产者、消费者和broker在内的所以组件在无状态的情况下，
建立起生产者和消费者的订阅关系，并实现生产者与消费者的负载均衡。

Producer 如果生产了数据，会先通过 zookeeper 找到 broker，然后将数据存放到 broker；
Consumer 如果要消费数据，会先通过 zookeeper 找对应的 broker，然后消费；

---

在Kafka中很多节点的调度以及资源的分配，都要依赖于zookeeper来完成。

1. Broker的注册，保存Broker的IP以及端口；
2. Topic注册，管理broker中Topic的分区以及分布情况
3. Broker的负载均衡，讲Topic动态的分配到broker中，通过topic的分布以及broker的负载判断
4. 消费者，每个分区的消息仅发送给一个消费者
5. 消费者与分区的对应关系，存储在zk中
6. 消费者负载均衡，一旦消费者增加或者减少，都会触发消费者的负载均衡
7. 消费者的offset，High level中由zk维护offset的信息；Low Level中由自己维护offset


## kafka特性
#### 高吞吐量、低延迟：
kafka每秒可以处理几十万条消息，它的延迟最低只有几毫秒，
每个topic可以分多个partition, consumer group 对partition进行consume操作；
#### 可扩展性：
kafka集群支持热扩展（不停机的情况下扩展kafka）；
#### 持久性、可靠性：
消息被持久化到本地磁盘，并且支持数据备份防止数据丢失；
#### 容错性：
允许集群中节点失败（若副本数量为n,则允许n-1个节点失败）；
#### 高并发：
支持数千个客户端同时读写；
#### 支持实时在线处理和离线处理：
可以使用Storm这种实时流处理系统对消息进行实时进行处理，
同时还可以使用Hadoop这种批处理系统进行离线处理；

## kafka使用场景
- 日志收集：
- 消息系统：
解耦和生产者和消费者、缓存消息等；
- 用户活动跟踪：
记录web用户或者app用户的各种活动，
如浏览网页、搜索、点击等活动，这些活动信息被各个服务器发布到kafka的topic中，
然后订阅者通过订阅这些topic来做实时的监控分析；
或者装载到Hadoop、数据仓库中做离线分析和数据挖掘；
- 运营指标：
Kafka也经常用来记录运营监控数据。
- 事件源：

## kafka的分区（针对topic）
- kafka采用分区（Partition）的方式，使得消费者能够做到并行消费，从而大大提高了自己的吞吐能力。
- 同时为了实现高可用，每个分区又有若干份副本（Replica），这样在某个broker挂掉的情况下，数据不会丢失。
- 无分区时，一个topic只有一个消费者在消费这个消息队列。
采用分区后，如果有两个分区，最多两个消费者同时消费，消费的速度肯定会更快。

![](../img/kafka/kafka-4.webp)
注意：
- 一个分区只能被同组的一个consumer消费；
- 同一个组里面的一个consumer可以消费多个分区；
- 消费率最高的情况是分区数和consumer数量相同；确保每个consumer专职负责一个分区。
- consumer数量不能大于分区数量；当consumer多余分区时，就会有consumer闲置；
- consumer group可以认为是一个订阅者集群，其中每个consumer负责自己所消费的分区；

## 副本（Replica-确保数据可恢复）

- 每个分区的数据都会有多份副本，以此来保证Kafka的高可用。
- topic下会划分多个partition，每个partition都有自己的replica，
其中只有一个是leader replica，其余的是follower replica。
- 消息进来的时候会先存入leader replica，然后从leader replica复制到follower replica。
只有复制全部完成时，consumer才可以消费此条消息。

Topic、partition、replica的关系图
![](../img/kafka/kafka-5.webp)

由上图可见，leader replica做了大量的工作。所以如果不同partition的leader replica在kafka集群的broker上分布不均匀，就会造成负载不均衡。
注：kafka通过轮询算法保证leader replica是均匀分布在多个broker上。如下图。

副本均匀分部图
![](../img/kafka/kafka-6.webp)
可以看到每个partition的leader replica均匀的分布在三个broker上，follower replica也是均匀分布的。

Replica总结：

- Replica均匀分配在Broker上，同一个partition的replica不会在同一个borker上；
- 同一个partition的Replica数量不能多于broker数量。
多个replica为了数据安全，一台server存多个replica没有意义。server挂掉，上面的副本都要挂掉。
- 分区的leader replica均衡分布在broker上。此时集群的负载是均衡的。这就叫做分区平衡；

## Partition的读和写
topic下划分了多个partition，消息的生产和消费最终都是发生在partition之上；

读写示意图
![](../img/kafka/kafka-7.webp)

- producer采用round-robin算法，轮询往每个partition写入topic；
- 每个partition都是有序的不可变的。
- Kafka可以保证partition的消费顺序，但不能保证topic消费顺序。
- 每个consumer维护的唯一元数据是offset，代表消费的位置，一般线性向后移动。
- consumer也可以重置offset到之前的位置，可以以任何顺序消费，不一定线性后移。

## 数据持久化
为了提高性能，现代操作系统往往使用内存作为磁盘的缓存；
虽然每个程序都在自己的线程里只缓存了一份数据，但在操作系统的缓存里还有一份，这等于存了两份数据。
与传统的将数据缓存在内存中然后刷到硬盘的设计不同，
Kafka直接将数据写到了文件系统的日志中。

## 消息传输的事务定义

### 消息投递语义
在业务中，常常都是使用 At least once 的模型，如果需要可重入的话，往往是业务自己实现。
#### At most once：最多一次，消息可能会丢失，但不会重复。
先获取数据，再 Commit Offset，最后进行业务处理：
- 生产者生产消息异常，不管，生产下一个消息，消息就丢了。
- 消费者处理消息，先更新 Offset，再做业务处理，做业务处理失败，消费者重启，消息就丢了。
#### At least once：最少一次，消息不会丢失，可能会重复。
先获取数据，再进行业务处理，业务处理成功后 Commit Offset：
- 生产者生产消息异常，消息是否成功写入不确定，重做，可能写入重复的消息。
- 消费者处理消息，业务处理成功后，更新 Offset 失败，消费者重启的话，会重复消费。
#### Exactly once：只且一次，消息不丢失不重复，只且消费一次（0.11 中实现，仅限于下游也是 Kafka）
思路是这样的，首先要保证消息不丢，再去保证不重复。所以盯着 At least once 的原因来搞。

首先想出来的：
- 生产者重做导致重复写入消息：生产保证幂等性。
- 消费者重复消费：消灭重复消费，或者业务接口保证幂等性重复消费也没问题。

由于业务接口是否幂等，不是 Kafka 能保证的，所以 Kafka 这里提供的 Exactly once 是有限制的，消费者的下游也必须是 Kafka。

所以以下讨论的，没特殊说明，消费者的下游系统都是 Kafka（注：使用 Kafka Conector，它对部分系统做了适配，实现了 Exactly once）。生产者幂等性好做，没啥问题。

解决重复消费有两个方法：
- 下游系统保证幂等性，重复消费也不会导致多条记录。
- 把 Commit Offset 和业务处理绑定成一个事务。

本来 Exactly once 实现第 1 点就 OK 了。但是在一些使用场景下，我们的数据源可能是多个 Topic，处理后输出到多个 Topic，这时我们会希望输出时要么全部成功，要么全部失败。这就需要实现事务性。

既然要做事务，那么干脆把重复消费的问题从根源上解决，把 Commit Offset 和输出到其他 Topic 绑定成一个事务。

##### 生产幂等性
思路是这样的，为每个 Producer 分配一个 Pid，作为该 Producer 的唯一标识。

Producer 会为每一个维护一个单调递增的 Seq。类似的，Broker 也会为每个记录下最新的 Seq。

**当 req_seq == broker_seq+1 时，Broker 才会接受该消息，因为：**
- 消息的 Seq 比 Broker 的 Seq 大超过时，说明中间有数据还没写入，即乱序了。
- 消息的 Seq 不比 Broker 的 Seq 小，那么说明该消息已被保存。
![](../img/kafka/kafka-product-2.jpg)

##### 事务性/原子性广播
场景是这样的：
- 先从多个源 Topic 中获取数据。
- 做业务处理，写到下游的多个目的 Topic。
- 更新多个源 Topic 的 Offset。

其中第 2、3 点作为一个事务，要么全成功，要么全失败。这里得益于 Offset 实际上是用特殊的 Topic 去保存，这两点都归一为写多个 Topic 的事务性处理。

![](../img/kafka/kafka-p-1.jpg)
基本思路是这样的：

- 引入 Tid（transaction id），和 Pid 不同，这个 ID 是应用程序提供的，用于标识事务，和 Producer 是谁并没关系。
- 就是任何 Producer 都可以使用这个 Tid 去做事务，这样进行到一半就死掉的事务，可以由另一个 Producer 去恢复。
- 同时为了记录事务的状态，类似对 Offset 的处理，引入 Transaction Coordinator 用于记录 Transaction Log。
- 在集群中会有多个 Transaction Coordinator，每个 Tid 对应唯一一个 Transaction Coordinator。
- 注：Transaction Log 删除策略是 Compact，已完成的事务会标记成 Null，Compact 后不保留。


做事务时，先标记开启事务，写入数据，全部成功就在 Transaction Log 中记录为 Prepare Commit 状态，否则写入 Prepare Abort 的状态。

之后再去给每个相关的 Partition 写入一条 Marker（Commit 或者 Abort）消息，标记这个事务的 Message 可以被读取或已经废弃。成功后在 Transaction Log记录下 Commit/Abort 状态，至此事务结束。

![](../img/kafka/kafka-c-1.jpg)
数据流：
- 首先使用 Tid 请求任意一个 Broker（代码中写的是负载最小的 Broker），找到对应的 Transaction Coordinator。
- 请求 Transaction Coordinator 获取到对应的 Pid，和 Pid 对应的 Epoch，这个 Epoch 用于防止僵死进程复活导致消息错乱。
- 当消息的 Epoch 比当前维护的 Epoch 小时，拒绝掉。Tid 和 Pid 有一一对应的关系，这样对于同一个 Tid 会返回相同的 Pid。
- Client 先请求 Transaction Coordinator 记录的事务状态，初始状态是 Begin，如果是该事务中第一个到达的，同时会对事务进行计时。
- Client 输出数据到相关的 Partition 中；Client 再请求 Transaction Coordinator 记录 Offset 的事务状态；Client 发送 Offset Commit 到对应 Offset Partition。
- Client 发送 Commit 请求，Transaction Coordinator 记录 Prepare Commit/Abort，然后发送 Marker 给相关的 Partition。
- 全部成功后，记录 Commit/Abort 的状态，最后这个记录不需要等待其他 Replica 的 ACK，因为 Prepare 不丢就能保证最终的正确性了。

这里 Prepare 的状态主要是用于事务恢复，例如给相关的 Partition 发送控制消息，没发完就宕机了，备机起来后，Producer 发送请求获取 Pid 时，会把未完成的事务接着完成。

当 Partition 中写入 Commit 的 Marker 后，相关的消息就可被读取。所以 Kafka 事务在 Prepare Commit 到 Commit 这个时间段内，消息是逐渐可见的，而不是同一时刻可见。

##### 消费事务
前面都是从生产的角度看待事务。还需要从消费的角度去考虑一些问题。

消费时，Partition 中会存在一些消息处于未 Commit 状态，即业务方应该看不到的消息，需要过滤这些消息不让业务看到，Kafka 选择在消费者进程中进行过来，而不是在 Broker 中过滤，主要考虑的还是性能。

Kafka 高性能的一个关键点是 Zero Copy，如果需要在 Broker 中过滤，那么势必需要读取消息内容到内存，就会失去 Zero Copy 的特性。

文件组织

Kafka 的数据，实际上是以文件的形式存储在文件系统的。Topic 下有 Partition，Partition 下有 Segment，Segment 是实际的一个个文件，Topic 和 Partition 都是抽象概念。

在目录 /partitionid}/ 下，存储着实际的 Log 文件（即 Segment），还有对应的索引文件。

每个 Segment 文件大小相等，文件名以这个 Segment 中最小的 Offset 命名，文件扩展名是 .log。Segment 对应的索引的文件名字一样，扩展名是 .index。

有两个 Index 文件：
- 一个是 Offset Index 用于按 Offset 去查 Message。
- 一个是 Time Index 用于按照时间去查，其实这里可以优化合到一起，下面只说 Offset Index。
![](../img/kafka/kafka-log-2.jpg)

为了减少索引文件的大小，降低空间使用，方便直接加载进内存中，这里的索引使用稀疏矩阵，不会每一个 Message 都记录下具体位置，而是每隔一定的字节数，再建立一条索引。

索引包含两部分：

BaseOffset：意思是这条索引对应 Segment 文件中的第几条 Message。这样做方便使用数值压缩算法来节省空间。例如 Kafka 使用的是 Varint。
Position：在 Segment 中的绝对位置。

查找 Offset 对应的记录时，会先用二分法，找出对应的 Offset 在哪个 Segment 中，然后使用索引，在定位出 Offset 在 Segment 中的大概位置，再遍历查找 Message。


### 数据传输的事务定义通常有以下三种级别：
#### 1. 最多一次(at most onece): 
消息不会被重复发送，最多被传输一次，但也有可能一次不传输。

---

基本思想是保证每一条消息commit成功之后，再进行消费处理；
设置自动提交为false，接受到消息之后，首先commit，然后再进行消费

#### 2. 最少一次(at least onece): 
消息不会被漏发送，最少被传输一次，但也有可能被重复传输.

---

基本思想是保证每一条消息处理成功之后，再进行commit；
设置自动提交为false；消息处理成功之后，手动进行commit；
采用这种模式时，最好保证消费操作的“幂等性”，防止重复消费；

#### 3. 精确的一次（exactly once）: 
不会漏传输也不会重复传输,每个消息都传输被一次而且仅仅被传输一次，这是大家所期望的。

---

核心思想是将offset作为唯一id与消息同时处理，并且保证处理的原子性；
设置自动提交为false；消息处理成功之后再提交；
比如对于关系型数据库来说，可以将id设置为消息处理结果的唯一索引，再次处理时，如果发现该索引已经存在，那么就不处理；

### kafka存在的问题
如果producer发布消息时发生了网络错误，但又不确定实在提交之前发生的还是提交之后发生的，
这种情况虽然不常见，但是必须考虑进去，现在Kafka版本还没有解决这个问题，
将来的版本正在努力尝试解决。

### kafka的可指定消息传输级别
并不是所有的情况都需要“精确的一次”这样高的级别，Kafka允许producer灵活的指定级别。
比如producer可以指定必须等待消息被提交的通知，
或者完全的异步发送消息而不等待任何通知，或者仅仅等待leader声明它拿到了消息。

## kafka性能优化

### 1、消息集：
以消息集为单位处理消息，比以单个的消息为单位处理，会提升不少性能。
Producer把消息集一块发送给服务端，而不是一条条的发送；
服务端把消息集一次性的追加到日志文件中，这样减少了琐碎的I/O操作。
consumer也可以一次性的请求一个消息集。

### 2、数据压缩
Kafka采用了端到端的压缩：因为有“消息集”的概念，客户端的消息可以一起被压缩后送到服务端，
并以压缩后的格式写入日志文件，以压缩的格式发送到consumer，
消息从producer发出到consumer拿到都是被压缩的，只有在consumer使用的时候才被解压缩，所以叫做“端到端的压缩”。
Kafka支持GZIP和Snappy压缩协议。


## Kafka Producer消息发送
客户端控制消息将被分发到哪个分区。
可以通过负载均衡随机的选择，或者使用分区函数。
Kafka允许用户实现分区函数，指定分区的key，将消息hash到不同的分区上;
比如如果你指定的key是user id，那么同一个用户发送的消息都被发送到同一个分区上。
经过分区之后，consumer就可以有目的的消费某个分区的消息。

### Producer异步发送消息：
批量发送可以很有效的提高发送效率。
Kafka producer的异步发送模式允许进行批量发送，先将消息缓存在内存中，然后一次请求批量发送出去。
这个策略可以配置的，比如可以指定缓存的消息达到某个量的时候就发出去，
或者缓存了固定的时间后就发送出去（比如100条消息就发送，或者每5秒发送一次）。
这种策略将大大减少服务端的I/O次数。

## Kafka Consumer
Kafa consumer消费消息时，向broker发出"fetch"请求去消费特定分区的消息。
consumer指定消息在日志中的偏移量（offset），就可以消费从这个位置开始的消息。
customer拥有了offset的控制权，可以向后回滚去重新消费之前的消息，这是很有意义的。

Kafka遵循了一种大部分消息系统共同的传统的设计：
producer将消息推送到broker，consumer从broker拉取消息。

- push模式下，当broker推送的速率远大于consumer消费的速率时，consumer恐怕就要崩溃了。
因此最终Kafka还是选取了传统的pull模式。
- Pull模式的另外一个好处是consumer可以自主决定是否批量的从broker拉取数据。
- Pull有个缺点是，如果broker没有可供消费的消息，将导致consumer不断在循环中轮询，直到新消息到达。
为了避免这点，Kafka有个参数可以让consumer阻塞知道新消息到达(当然也可以阻塞知道消息的数量达到某个特定的量consumer才去拉去消息)。

### 消费者存在group时 auto.offset.reset=
latest 最后, earliest 最早, none

### 几种不同的注册方式

1. subscribe方式：当主题分区数量变化或者consumer数量变化时，会进行rebalance；
2. 注册rebalance监听器，可以手动管理offset
3. 不注册监听器，kafka自动管理
4. assign方式：手动将consumer与partition进行对应，kafka不会进行rebanlance

### 默认配置
采用默认配置情况下，既不能完全保证At-least-once 也不能完全保证at-most-once；

比如：

在自动提交之后，数据消费流程失败，这样就会有丢失，不能保证at-least-once；

数据消费成功，但是自动提交失败，可能会导致重复消费，这样也不能保证at-most-once；

但是将自动提交时长设置的足够小，则可以最大限度的保证at-most-once；

### 关键配置及含义
enable.auto.commit 是否自动提交自己的offset值；默认值时true
auto.commit.interval.ms 自动提交时长间隔；默认值时5000 ms
consumer.commitSync(); offset提交命令；

### 消费的机制
- at most once，即消费数据后，保存offset，就再也取不到这个数据了。
- at least once，即消费数据后，保存offset，如果保存出错，下次可能还会取到该数据

在Kafka中offset是由consumer维护的（实际可以由zookeeper来完成，0.8以后在名为__consumer_offsets的topic中，该topic有50个分区）。这种机制有两个好处
- 一个是可以依据consumer的能力来消费数据，避免产生消费数据的压力；
- 另一个就是可以自定义fetch消费的数据数目，可以一次读取1条，也可以1次读取100条。

## 主从同步
- 创建副本的单位是topic的分区，每个分区都有一个leader和零或多个followers；
所有的读写操作都由leader处理；同一个分区的副本数量不能多于brokers的数量；

- 各分区的leader均匀的分布在brokers中。
所有的followers都复制leader的日志，日志中的消息和顺序都和leader中的一致。
flowers向普通的consumer那样从leader那里拉取消息并保存在自己的日志文件中。

- Kafka判断一个节点是否活着有两个条件：
  1. 节点必须可以维护和ZooKeeper的连接，Zookeeper通过心跳机制检查每个节点的连接。
  2. 如果节点是个follower,他必须能及时的同步leader的写操作，延时不能太久。

- leader的选择
Kafka的核心是日志文件，日志文件在集群中的同步是分布式数据系统最基础的要素。
  - 1. Kafaka动态维护了一个同步状态的副本的集合简称ISR；
集合中的任何一个节点随时都可以被选为leader。ISR在ZooKeeper中维护。
ISR的成员是动态的，如果一个节点被淘汰了，当它重新达到“同步中”的状态时，他可以重新加入ISR。
这种leader的选择方式是非常快速的，适合kafka的应用场景。
  - 2. 所有的副本都down掉时，必须及时作出反应。可以有以下两种选择（kafka选择b）：
     - a. 等待ISR中的任何一个节点恢复并担任leader
（ISR中的节点都起不来，那集群就永远恢复不了）
     - b. 选择所有节点中（不只是ISR）第一个恢复的节点作为leader（如果等待ISR以外的节点恢复，这个节点的数据就会被作为线上数据，
有可能和真实的数据有所出入，因为有些数据它可能还没同步到。）

### Partition
当存在多副本的情况下，会尽量把多个副本，分配到不同的 Broker 上。

Kafka 会为 Partition 选出一个 Leader，之后所有该 Partition 的请求，实际操作的都是 Leader，然后再同步到其他的 Follower。

当一个 Broker 歇菜后，所有 Leader 在该 Broker 上的 Partition 都会重新选举，选出一个 Leader。（这里不像分布式文件存储系统那样会自动进行复制保持副本数）
然后这里就涉及两个细节：
1. 怎么分配 Partition
2. 怎么选 Leader
关于 Partition 的分配，还有 Leader 的选举，总得有个执行者。在 Kafka 中，这个执行者就叫 **Controller**。

Kafka 使用 ZK 在 Broker 中选出一个 Controller，用于 Partition 分配和 Leader 选举。
### Partition 的分配

- 将所有 Broker（假设共 n 个 Broker）和待分配的 Partition 排序。
- 将第 i 个 Partition 分配到第（i mod n）个 Broker 上 （这个就是 Leader）。
- 将第 i 个 Partition 的第 j 个 Replica 分配到第（(i + j) mode n）个 Broker 上。

### Leader 容灾
Controller 会在 ZK 的 /brokers/ids 节点上注册 Watch，一旦有 Broker 宕机，它就能知道。

当 Broker 宕机后，Controller 就会给受到影响的 Partition 选出新 Leader。

Controller 从 ZK 的 /brokers/topics/[topic]/partitions/[partition]/state 中，读取对应 Partition 的 ISR（in-sync replica 已同步的副本）列表，选一个出来做 Leader。

选出 Leader 后，更新 ZK，然后发送 LeaderAndISRRequest 给受影响的 Broker，让它们知道改变这事。

为什么这里不是使用 ZK 通知，而是直接给 Broker 发送 RPC 请求，我的理解可能是这样做 ZK 有性能问题吧。

如果 ISR 列表是空，那么会根据配置，随便选一个 Replica 做 Leader，或者干脆这个 Partition 就是歇菜。

如果 ISR 列表的有机器，但是也歇菜了，那么还可以等 ISR 的机器活过来。


### 多副本同步

#### 生产
这里的策略，服务端这边的处理是 Follower 从 Leader 批量拉取数据来同步。但是具体的可靠性，是由生产者来决定的。

生产者生产消息的时候，通过 request.required.acks 参数来设置数据的可靠性。
![](../img/kafka/kafka-producter.jpg)
在 Acks=-1 的时候，如果 ISR 少于 min.insync.replicas 指定的数目，那么就会返回不可用。

这里 ISR 列表中的机器是会变化的，根据配置 replica.lag.time.max.ms，多久没同步，就会从 ISR 列表中剔除。

以前还有根据落后多少条消息就踢出 ISR，在 1.0 版本后就去掉了，因为这个值很难取，在高峰的时候很容易出现节点不断的进出 ISR 列表。

从 ISA 中选出 Leader 后，Follower 会把自己日志中上一个高水位后面的记录去掉，然后去和 Leader 拿新的数据。

因为新的 Leader 选出来后，Follower 上面的数据，可能比新 Leader 多，所以要截取。

这里高水位的意思，对于 Partition 和 Leader，就是所有 ISR 中都有的最新一条记录。消费者最多只能读到高水位。

从 Leader 的角度来说高水位的更新会延迟一轮，例如写入了一条新消息，ISR 中的 Broker 都 Fetch 到了，但是 ISR 中的 Broker 只有在下一轮的 Fetch 中才能告诉 Leader。

也正是由于这个高水位延迟一轮，在一些情况下，Kafka 会出现丢数据和主备数据不一致的情况，0.11 开始，使用 Leader Epoch 来代替高水位。

##### 思考：当 Acks=-1 时
- 是 Follwers 都来 Fetch 就返回成功，还是等 Follwers 第二轮 Fetch？
- Leader 已经写入本地，但是 ISR 中有些机器失败，那么怎么处理呢？


#### 消费
订阅 Topic 是以一个消费组来订阅的，一个消费组里面可以有多个消费者。同一个消费组中的两个消费者，不会同时消费一个 Partition。

换句话来说，就是一个 Partition，只能被消费组里的一个消费者消费，但是可以同时被多个消费组消费。

因此，如果消费组内的消费者如果比 Partition 多的话，那么就会有个别消费者一直空闲。

订阅 Topic 时，可以用正则表达式，如果有新 Topic 匹配上，那能自动订阅上。

##### Offset 的保存
一个消费组消费 Partition，需要保存 Offset 记录消费到哪，以前保存在 ZK 中，由于 ZK 的写性能不好，以前的解决方法都是 Consumer 每隔一分钟上报一次。

这里 ZK 的性能严重影响了消费的速度，而且很容易出现重复消费。在 0.10 版本后，Kafka 把这个 Offset 的保存，从 ZK 总剥离，保存在一个名叫 consumeroffsets topic 的 Topic 中。

写进消息的 Key 由 Groupid、Topic、Partition 组成，Value 是偏移量 Offset。Topic 配置的清理策略是 Compact。总是保留最新的 Key，其余删掉。

一般情况下，每个 Key 的 Offset 都是缓存在内存中，查询的时候不用遍历 Partition，如果没有缓存，第一次就会遍历 Partition 建立缓存，然后查询返回。

确定 Consumer Group 位移信息写入 consumers_offsets 的哪个 Partition，具体计算公式：
```
__consumers_offsets partition =
 Math.abs(groupId.hashCode() % groupMetadataTopicPartitionCount) 
//groupMetadataTopicPartitionCount由offsets.topic.num.partitions指定，默认是50个分区。
```



### 分配 Partition—Reblance

生产过程中 Broker 要分配 Partition，消费过程这里，也要分配 Partition 给消费者。

类似 Broker 中选了一个 Controller 出来，消费也要从 Broker 中选一个 Coordinator，用于分配 Partition。

#### 1选 Coordinator：
看 Offset 保存在那个 Partition；该 Partition Leader 所在的 Broker 就是被选定的 Coordinator。

这里我们可以看到，Consumer Group 的 Coordinator，和保存 Consumer Group Offset 的 Partition Leader 是同一台机器。
#### 2交互流程：
把 Coordinator 选出来之后，就是要分配了。
整个流程是这样的：

- Consumer 启动、或者 Coordinator 宕机了，Consumer 会任意请求一个 Broker，发送 ConsumerMetadataRequest 请求。
- Broker 会按照上面说的方法，选出这个 Consumer 对应 Coordinator 的地址。
- Consumer 发送 Heartbeat 请求给 Coordinator，返回 IllegalGeneration 的话，就说明 Consumer 的信息是旧的了，需要重新加入进来，进行 Reblance。
- 返回成功，那么 Consumer 就从上次分配的 Partition 中继续执行。

#### Reblance 流程：
- Consumer 给 Coordinator 发送 JoinGroupRequest 请求。
- 这时其他 Consumer 发 Heartbeat 请求过来时，Coordinator 会告诉他们，要 Reblance 了。
- 其他 Consumer 发送 JoinGroupRequest 请求。
- 所有记录在册的 Consumer 都发了 JoinGroupRequest 请求之后，Coordinator 就会在这里 Consumer 中随便选一个 Leader。
- 然后回 JoinGroupRespone，这会告诉 Consumer 你是 Follower 还是 Leader，对于 Leader，还会把 Follower 的信息带给它，让它根据这些信息去分配 Partition。
- Consumer 向 Coordinator 发送 SyncGroupRequest，其中 Leader 的 SyncGroupRequest 会包含分配的情况。
- Coordinator 回包，把分配的情况告诉 Consumer，包括 Leader。

当 Partition 或者消费者的数量发生变化时，都得进行 Reblance。

#### 列举一下会 Reblance 的情况：
- 增加 Partition
- 增加消费者
- 消费者主动关闭
- 消费者宕机了
- Coordinator 自己也宕机了

## kafka与rabbitmq的区别

- 架构模型方面
kafka遵从一般的MQ结构，producer，broker，consumer，以consumer为中心，消息的消费信息保存的客户端consumer上，
consumer根据消费的点，从broker上批量pull数据；无消息确认机制。
RabbitMQ遵循AMQP协议，RabbitMQ的broker由Exchange,Binding,queue组成，
其中exchange和binding组成了消息的路由键（消息到哪个队列）；
客户端Producer通过连接channel和server进行通信，Consumer从queue获取消息进行消费
（长连接，queue有消息会推送到consumer端，consumer循环从输入流读取数据）
rabbitMQ以broker为中心；有消息的确认机制（消费过后剔除队列）。

- 在吞吐量方面
kafka具有高的吞吐量，内部采用消息的批量处理，zero-copy机制，
数据的存储和获取是本地磁盘顺序批量操作（文件系统），具有O(1)的复杂度，消息处理的效率很高。
rabbitMQ在吞吐量方面稍逊于kafka，rabbitMQ支持对消息的可靠的传递，支持事务，不支持批量的操作；

- 在可用性方面
rabbitMQ支持miror的queue，主queue失效，miror queue接管。
kafka的broker支持主备模式（副本）。

- 在集群负载均衡方面
kafka采用zookeeper对集群中的broker、consumer进行管理，可以注册topic到zookeeper上；
通过zookeeper的协调机制，producer保存对应topic的broker信息，可以随机或者轮询发送到broker上；
rabbitMQ的负载均衡需要单独的loadbalancer进行支持。

## API
### High Level API
Offset，路由啥都替我们干了，用起来很简单。
### Simple API /low level API
Offset 啥的都是要我们自己记录。（注：消息消费的时候，首先要知道去哪消费，这就是路由，消费完之后，要记录消费单哪，就是 Offset）

## 常用配置项

### Broker 配置
![](../img/kafka/kafka-config-1.jpg)
### Topic 配置

![](../img/kafka/kafka-config-2.jpg)