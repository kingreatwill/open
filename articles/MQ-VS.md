# 17 个方面，综合对比 Kafka、RabbitMQ、RocketMQ、ActiveMQ 四个分布式消息队列
本文将从，Kafka、RabbitMQ、ZeroMQ、RocketMQ、ActiveMQ 17 个方面综合对比作为消息队列使用时的差异。
## 一、资料文档
- Kafka：中。有kafka作者自己写的书，网上资料也有一些。
- rabbitmq：多。有一些不错的书，网上资料多。
- zeromq：少。没有专门写zeromq的书，网上的资料多是一些代码的实现和简单介绍。
- rocketmq：少。没有专门写rocketmq的书，网上的资料良莠不齐，官方文档很简洁，但是对技术细节没有过多的描述。
- activemq：多。没有专门写activemq的书，网上资料多。

## 二、开发语言
- Kafka：Scala 
- rabbitmq：Erlang 
- zeromq：c 
- rocketmq：java 
- activemq：java

## 三、支持的协议
- Kafka：自己定义的一套…（基于TCP） 
- rabbitmq：AMQP 
- zeromq：TCP、UDP 
- rocketmq：自己定义的一套… 
- activemq：OpenWire、STOMP、REST、XMPP、AMQP

## 四、消息存储
- Kafka：内存、磁盘、数据库。支持大量堆积。
kafka的最小存储单元是分区，一个topic包含多个分区，kafka创建主题时，这些分区会被分配在多个服务器上，通常一个broker一台服务器。分区首领会均匀地分布在不同的服务器上，分区副本也会均匀的分布在不同的服务器上，确保负载均衡和高可用性，当新的broker加入集群的时候，部分副本会被移动到新的broker上。根据配置文件中的目录清单，kafka会把新的分区分配给目录清单里分区数最少的目录。默认情况下，分区器使用轮询算法把消息均衡地分布在同一个主题的不同分区中，对于发送时指定了key的情况，会根据key的hashcode取模后的值存到对应的分区中。
- rabbitmq：内存、磁盘。支持少量堆积。
rabbitmq的消息分为持久化的消息和非持久化消息，不管是持久化的消息还是非持久化的消息都可以写入到磁盘。持久化的消息在到达队列时就写入到磁盘，并且如果可以，持久化的消息也会在内存中保存一份备份，这样可以提高一定的性能，当内存吃紧的时候会从内存中清除。非持久化的消息一般只存在于内存中，在内存吃紧的时候会被换入到磁盘中，以节省内存。
引入镜像队列机制，可将重要队列“复制”到集群中的其他broker上，保证这些队列的消息不会丢失。配置镜像的队列，都包含一个主节点master和多个从节点slave,如果master失效，加入时间最长的slave会被提升为新的master，除发送消息外的所有动作都向master发送，然后由master将命令执行结果广播给各个slave，rabbitmq会让master均匀地分布在不同的服务器上，而同一个队列的slave也会均匀地分布在不同的服务器上，保证负载均衡和高可用性。
- zeromq：消息发送端的内存或者磁盘中。不支持持久化。
- rocketmq：磁盘。支持大量堆积。
commitLog文件存放实际的消息数据，每个commitLog上限是1G，满了之后会自动新建一个commitLog文件保存数据。ConsumeQueue队列只存放offset、size、tagcode，非常小，分布在多个broker上。ConsumeQueue相当于CommitLog的索引文件，消费者消费时会从consumeQueue中查找消息在commitLog中的offset，再去commitLog中查找元数据。
ConsumeQueue存储格式的特性，保证了写过程的顺序写盘（写CommitLog文件），大量数据IO都在顺序写同一个commitLog，满1G了再写新的。加上rocketmq是累计4K才强制从PageCache中刷到磁盘（缓存），所以高并发写性能突出。
- activemq：内存、磁盘、数据库。支持少量堆积。

## 五、消息事务
- Kafka：支持 
- rabbitmq：支持。客户端将信道设置为事务模式，只有当消息被rabbitMq接收，事务才能提交成功，否则在捕获异常后进行回滚。使用事务会使得性能有所下降 
- zeromq：不支持 
- rocketmq：支持 
- activemq：支持

## 六、负载均衡
- Kafka：支持负载均衡。
    1. 一个broker通常就是一台服务器节点。对于同一个Topic的不同分区，Kafka会尽力将这些分区分布到不同的Broker服务器上，zookeeper保存了broker、主题和分区的元数据信息。分区首领会处理来自客户端的生产请求，kafka分区首领会被分配到不同的broker服务器上，让不同的broker服务器共同分担任务。每一个broker都缓存了元数据信息，客户端可以从任意一个broker获取元数据信息并缓存起来，根据元数据信息知道要往哪里发送请求。

    2. kafka的消费者组订阅同一个topic，会尽可能地使得每一个消费者分配到相同数量的分区，分摊负载。

    3. 当消费者加入或者退出消费者组的时候，还会触发再均衡，为每一个消费者重新分配分区，分摊负载。kafka的负载均衡大部分是自动完成的，分区的创建也是kafka完成的，隐藏了很多细节，避免了繁琐的配置和人为疏忽造成的负载问题。

    4. 发送端由topic和key来决定消息发往哪个分区，如果key为null，那么会使用轮询算法将消息均衡地发送到同一个topic的不同分区中。如果key不为null，那么会根据key的hashcode取模计算出要发往的分区。

- rabbitmq：对负载均衡的支持不好。
    1. 消息被投递到哪个队列是由交换器和key决定的，交换器、路由键、队列都需要手动创建。
rabbitmq客户端发送消息要和broker建立连接，需要事先知道broker上有哪些交换器，有哪些队列。通常要声明要发送的目标队列，如果没有目标队列，会在broker上创建一个队列，如果有，就什么都不处理，接着往这个队列发送消息。假设大部分繁重任务的队列都创建在同一个broker上，那么这个broker的负载就会过大。（可以在上线前预先创建队列，无需声明要发送的队列，但是发送时不会尝试创建队列，可能出现找不到队列的问题，rabbitmq的备份交换器会把找不到队列的消息保存到一个专门的队列中，以便以后查询使用）
使用镜像队列机制建立rabbitmq集群可以解决这个问题，形成master-slave的架构，master节点会均匀分布在不同的服务器上，让每一台服务器分摊负载。slave节点只是负责转发，在master失效时会选择加入时间最长的slave成为master。
当新节点加入镜像队列的时候，队列中的消息不会同步到新的slave中，除非调用同步命令，但是调用命令后，队列会阻塞，不能在生产环境中调用同步命令。

    2. 当rabbitmq队列拥有多个消费者的时候，队列收到的消息将以轮询的分发方式发送给消费者。每条消息只会发送给订阅列表里的一个消费者，不会重复。
这种方式非常适合扩展，而且是专门为并发程序设计的。
如果某些消费者的任务比较繁重，那么可以设置basicQos限制信道上消费者能保持的最大未确认消息的数量，在达到上限时，rabbitmq不再向这个消费者发送任何消息。

    3. 对于rabbitmq而言，客户端与集群建立的TCP连接不是与集群中所有的节点建立连接，而是挑选其中一个节点建立连接。
但是rabbitmq集群可以借助HAProxy、LVS技术，或者在客户端使用算法实现负载均衡，引入负载均衡之后，各个客户端的连接可以分摊到集群的各个节点之中。
客户端均衡算法：
        1. 轮询法。按顺序返回下一个服务器的连接地址。
        2. 加权轮询法。给配置高、负载低的机器配置更高的权重，让其处理更多的请求；而配置低、负载高的机器，给其分配较低的权重，降低其系统负载。
        3. 随机法。随机选取一个服务器的连接地址。
        4. 加权随机法。按照概率随机选取连接地址。
        5. 源地址哈希法。通过哈希函数计算得到的一个数值，用该数值对服务器列表的大小进行取模运算。
        6. 最小连接数法。动态选择当前连接数最少的一台服务器的连接地址。

- zeromq：去中心化，不支持负载均衡。本身只是一个多线程网络库。
- rocketmq：支持负载均衡。
一个broker通常是一个服务器节点，broker分为master和slave,master和slave存储的数据一样，slave从master同步数据。
    1. nameserver与每个集群成员保持心跳，保存着Topic-Broker路由信息，同一个topic的队列会分布在不同的服务器上。

    2. 发送消息通过轮询队列的方式发送，每个队列接收平均的消息量。发送消息指定topic、tags、keys，无法指定投递到哪个队列（没有意义，集群消费和广播消费跟消息存放在哪个队列没有关系）。tags选填，类似于 Gmail 为每封邮件设置的标签，方便服务器过滤使用。目前只支 持每个消息设置一个 tag，所以也可以类比为 Notify 的 MessageType 概念。keys选填，代表这条消息的业务关键词，服务器会根据 keys 创建哈希索引，设置后， 可以在 Console 系统根据 Topic、Keys 来查询消息，由于是哈希索引，请尽可能 保证 key 唯一，例如订单号，商品 Id 等。

    3. rocketmq的负载均衡策略规定：Consumer数量应该小于等于Queue数量，如果Consumer超过Queue数量，那么多余的Consumer 将不能消费消息。这一点和kafka是一致的，rocketmq会尽可能地为每一个Consumer分配相同数量的队列，分摊负载。

- activemq：支持负载均衡。可以基于zookeeper实现负载均衡。

## 七、集群方式
- Kafka：天然的‘Leader-Slave’无状态集群，每台服务器既是Master也是Slave。
分区首领均匀地分布在不同的kafka服务器上，分区副本也均匀地分布在不同的kafka服务器上，所以每一台kafka服务器既含有分区首领，同时又含有分区副本，每一台kafka服务器是某一台kafka服务器的Slave，同时也是某一台kafka服务器的leader。
kafka的集群依赖于zookeeper，zookeeper支持热扩展，所有的broker、消费者、分区都可以动态加入移除，而无需关闭服务，与不依靠zookeeper集群的mq相比，这是最大的优势。

- rabbitmq：支持简单集群，'复制'模式，对高级集群模式支持不好。
rabbitmq的每一个节点，不管是单一节点系统或者是集群中的一部分，要么是内存节点，要么是磁盘节点，集群中至少要有一个是磁盘节点。
在rabbitmq集群中创建队列，集群只会在单个节点创建队列进程和完整的队列信息（元数据、状态、内容），而不是在所有节点上创建。
引入镜像队列，可以避免单点故障，确保服务的可用性，但是需要人为地为某些重要的队列配置镜像。

- zeromq：去中心化，不支持集群。

- rocketmq：常用 多对'Master-Slave' 模式，开源版本需手动切换Slave变成Master
Name Server是一个几乎无状态节点，可集群部署，节点之间无任何信息同步。
Broker部署相对复杂，Broker分为Master与Slave，一个Master可以对应多个Slave，但是一个Slave只能对应一个Master，Master与Slave的对应关系通过指定相同的BrokerName，不同的BrokerId来定义，BrokerId为0表示Master，非0表示Slave。Master也可以部署多个。每个Broker与Name Server集群中的所有节点建立长连接，定时注册Topic信息到所有Name Server。
Producer与Name Server集群中的其中一个节点（随机选择）建立长连接，定期从Name Server取Topic路由信息，并向提供Topic服务的Master建立长连接，且定时向Master发送心跳。Producer完全无状态，可集群部署。
Consumer与Name Server集群中的其中一个节点（随机选择）建立长连接，定期从Name Server取Topic路由信息，并向提供Topic服务的Master、Slave建立长连接，且定时向Master、Slave发送心跳。Consumer既可以从Master订阅消息，也可以从Slave订阅消息，订阅规则由Broker配置决定。
客户端先找到NameServer, 然后通过NameServer再找到 Broker。
一个topic有多个队列，这些队列会均匀地分布在不同的broker服务器上。rocketmq队列的概念和kafka的分区概念是基本一致的，kafka同一个topic的分区尽可能地分布在不同的broker上，分区副本也会分布在不同的broker上。
rocketmq集群的slave会从master拉取数据备份，master分布在不同的broker上。

- activemq：支持简单集群模式，比如'主-备'，对高级集群模式支持不好。

## 八、管理界面
- Kafka：一般 
- rabbitmq：好 
- zeromq：无 
- rocketmq：无 
- activemq：一般

## 九、可用性
- Kafka：非常高（分布式） 
- rabbitmq：高（主从） 
- zeromq：高。
- rocketmq：非常高（分布式） 
- activemq：高（主从）

## 十、消息重复
- Kafka：支持at least once、at most once
- rabbitmq：支持at least once、at most once
- zeromq：只有重传机制，但是没有持久化，消息丢了重传也没有用。既不是at least once、也不是at most once、更不是exactly only once
- rocketmq：支持at least once
- activemq：支持at least once

## 十一、吞吐量TPS
- Kafka：极大 Kafka按批次发送消息和消费消息。发送端将多个小消息合并，批量发向Broker，消费端每次取出一个批次的消息批量处理。
- rabbitmq：比较大 
- zeromq：极大 
- rocketmq：大 rocketMQ接收端可以批量消费消息，可以配置每次消费的消息数，但是发送端不是批量发送。
- activemq：比较大

## 十二、订阅形式和消息分发
- Kafka：基于topic以及按照topic进行正则匹配的发布订阅模式。
    - 【发送】
发送端由topic和key来决定消息发往哪个分区，如果key为null，那么会使用轮询算法将消息均衡地发送到同一个topic的不同分区中。如果key不为null，那么会根据key的hashcode取模计算出要发往的分区。
    - 【接收】
    1. consumer向群组协调器broker发送心跳来维持他们和群组的从属关系以及他们对分区的所有权关系，所有权关系一旦被分配就不会改变除非发生再均衡(比如有一个consumer加入或者离开consumer group)，consumer只会从对应的分区读取消息。

    2. kafka限制consumer个数要少于分区个数,每个消息只会被同一个 Consumer Group的一个consumer消费（非广播）。
    3. kafka的 Consumer Group订阅同一个topic，会尽可能地使得每一个consumer分配到相同数量的分区，不同 Consumer Group订阅同一个主题相互独立，同一个消息会被不同的 Consumer Group处理。
- rabbitmq：提供了4种：direct, topic ,Headers和fanout。
    - 【发送】
先要声明一个队列，这个队列会被创建或者已经被创建，队列是基本存储单元。
由exchange和key决定消息存储在哪个队列。
direct>发送到和bindingKey完全匹配的队列。
topic>路由key是含有"."的字符串，会发送到含有“*”、“#”进行模糊匹配的bingKey对应的队列。
fanout>与key无关，会发送到所有和exchange绑定的队列
headers>与key无关，消息内容的headers属性（一个键值对）和绑定键值对完全匹配时，会发送到此队列。此方式性能低一般不用
    - 【接收】
rabbitmq的队列是基本存储单元，不再被分区或者分片，对于我们已经创建了的队列，消费端要指定从哪一个队列接收消息。
当rabbitmq队列拥有多个消费者的时候，队列收到的消息将以轮询的分发方式发送给消费者。每条消息只会发送给订阅列表里的一个消费者，不会重复。
这种方式非常适合扩展，而且是专门为并发程序设计的。
如果某些消费者的任务比较繁重，那么可以设置basicQos限制信道上消费者能保持的最大未确认消息的数量，在达到上限时，rabbitmq不再向这个消费者发送任何消息。

- zeromq：点对点(p2p)
- rocketmq：基于topic/messageTag以及按照消息类型、属性进行正则匹配的发布订阅模式
    - 【发送】
发送消息通过轮询队列的方式发送，每个队列接收平均的消息量。发送消息指定topic、tags、keys，无法指定投递到哪个队列（没有意义，集群消费和广播消费跟消息存放在哪个队列没有关系）。
tags选填，类似于 Gmail 为每封邮件设置的标签，方便服务器过滤使用。目前只支 持每个消息设置一个 tag，所以也可以类比为 Notify 的 MessageType 概念。
keys选填，代表这条消息的业务关键词，服务器会根据 keys 创建哈希索引，设置后， 可以在 Console 系统根据 Topic、Keys 来查询消息，由于是哈希索引，请尽可能 保证 key 唯一，例如订单号，商品 Id 等。
    - 【接收】
    1. 广播消费。一条消息被多个Consumer消费，即使Consumer属于同一个ConsumerGroup，消息也会被ConsumerGroup中的每个Consumer都消费一次。
    2. 集群消费。一个 Consumer Group中的Consumer实例平均分摊消费消息。例如某个Topic有 9 条消息，其中一个Consumer Group有3个实例，那么每个实例只消费其中的 3 条消息。即每一个队列都把消息轮流分发给每个consumer。

- activemq：点对点(p2p)、广播（发布-订阅）
点对点模式，每个消息只有1个消费者；
发布/订阅模式，每个消息可以有多个消费者。
    - 【发送】
点对点模式：先要指定一个队列，这个队列会被创建或者已经被创建。
发布/订阅模式：先要指定一个topic，这个topic会被创建或者已经被创建。
    - 【接收】
点对点模式：对于已经创建了的队列，消费端要指定从哪一个队列接收消息。
发布/订阅模式：对于已经创建了的topic，消费端要指定订阅哪一个topic的消息。

## 十三、顺序消息
- Kafka：支持。
设置生产者的max.in.flight.requests.per.connection为1，可以保证消息是按照发送顺序写入服务器的，即使发生了重试。
kafka保证同一个分区里的消息是有序的，但是这种有序分两种情况
    - key为null，消息逐个被写入不同主机的分区中，但是对于每个分区依然是有序的
    - key不为null , 消息被写入到同一个分区，这个分区的消息都是有序。

- rabbitmq：不支持
- zeromq：不支持
- rocketmq：支持
- activemq：不支持

## 十四、消息确认
- Kafka：支持。
    1. 发送方确认机制
ack=0，不管消息是否成功写入分区
ack=1，消息成功写入首领分区后，返回成功
ack=all，消息成功写入所有分区后，返回成功。
    2. 接收方确认机制
自动或者手动提交分区偏移量，早期版本的kafka偏移量是提交给Zookeeper的，这样使得zookeeper的压力比较大，更新版本的kafka的偏移量是提交给kafka服务器的，不再依赖于zookeeper群组，集群的性能更加稳定。

- rabbitmq：支持。
    1. 发送方确认机制，消息被投递到所有匹配的队列后，返回成功。如果消息和队列是可持久化的，那么在写入磁盘后，返回成功。支持批量确认和异步确认。
    2. 接收方确认机制，设置autoAck为false，需要显式确认，设置autoAck为true，自动确认。
当autoAck为false的时候，rabbitmq队列会分成两部分，一部分是等待投递给consumer的消息，一部分是已经投递但是没收到确认的消息。如果一直没有收到确认信号，并且consumer已经断开连接，rabbitmq会安排这个消息重新进入队列，投递给原来的消费者或者下一个消费者。
未确认的消息不会有过期时间，如果一直没有确认，并且没有断开连接，rabbitmq会一直等待，rabbitmq允许一条消息处理的时间可以很久很久。

- zeromq：支持。
- rocketmq：支持。
- activemq：支持。

## 十五、消息回溯
- Kafka：支持指定分区offset位置的回溯。
- rabbitmq：不支持 
- zeromq：不支持 
- rocketmq：支持指定时间点的回溯。
- activemq：不支持

## 十六、消息重试
- Kafka：不支持，但是可以实现。
kafka支持指定分区offset位置的回溯，可以实现消息重试。

- rabbitmq：不支持，但是可以利用消息确认机制实现。
rabbitmq接收方确认机制，设置autoAck为false。
当autoAck为false的时候，rabbitmq队列会分成两部分，一部分是等待投递给consumer的消息，一部分是已经投递但是没收到确认的消息。如果一直没有收到确认信号，并且consumer已经断开连接，rabbitmq会安排这个消息重新进入队列，投递给原来的消费者或者下一个消费者。

- zeromq：不支持，
- rocketmq：支持。
消息消费失败的大部分场景下，立即重试99%都会失败，所以rocketmq的策略是在消费失败时定时重试，每次时间间隔相同。
    - 发送端的 send 方法本身支持内部重试，重试逻辑如下：
    a)至多重试3次；
    b)如果发送失败，则轮转到下一个broker；
    c)这个方法的总耗时不超过sendMsgTimeout 设置的值，默认 10s，超过时间不在重试。
    - 接收端。
Consumer 消费消息失败后，要提供一种重试机制，令消息再消费一次。Consumer 消费消息失败通常可以分为以下两种情况：
由于消息本身的原因，例如反序列化失败，消息数据本身无法处理（例如话费充值，当前消息的手机号被
注销，无法充值）等。定时重试机制，比如过 10s 秒后再重试。
由于依赖的下游应用服务不可用，例如 db 连接不可用，外系统网络不可达等。
即使跳过当前失败的消息，消费其他消息同样也会报错。这种情况可以 sleep 30s，再消费下一条消息，减轻 Broker 重试消息的压力。

- activemq：不支持

## 十七、并发度
- Kafka：高
一个线程一个消费者，kafka限制消费者的个数要小于等于分区数，如果要提高并行度，可以在消费者中再开启多线程，或者增加consumer实例数量。
- rabbitmq：极高
本身是用Erlang语言写的，并发性能高。
可在消费者中开启多线程，最常用的做法是一个channel对应一个消费者，每一个线程把持一个channel，多个线程复用connection的tcp连接，减少性能开销。
当rabbitmq队列拥有多个消费者的时候，队列收到的消息将以轮询的分发方式发送给消费者。每条消息只会发送给订阅列表里的一个消费者，不会重复。
这种方式非常适合扩展，而且是专门为并发程序设计的。
如果某些消费者的任务比较繁重，那么可以设置basicQos限制信道上消费者能保持的最大未确认消息的数量，在达到上限时，rabbitmq不再向这个消费者发送任何消息。

- zeromq：高
- rocketmq：高
    - rocketmq限制消费者的个数少于等于队列数，但是可以在消费者中再开启多线程，这一点和kafka是一致的，提高并行度的方法相同。
修改消费并行度方法
a) 同一个 ConsumerGroup 下，通过增加 Consumer 实例数量来提高并行度，超过订阅队列数的 Consumer实例无效。
b) 提高单个 Consumer 的消费并行线程，通过修改参数consumeThreadMin、consumeThreadMax
    - 同一个网络连接connection，客户端多个线程可以同时发送请求，连接会被复用，减少性能开销。
- activemq：高
单个ActiveMQ的接收和消费消息的速度在1万笔/秒（持久化 一般为1-2万， 非持久化 2 万以上），在生产环境中部署10个Activemq就能达到10万笔/秒以上的性能，部署越多的activemq broker 在MQ上latency也就越低，系统吞吐量也就越高。