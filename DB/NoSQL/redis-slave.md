## redis 主从从结构
[Redis主从同步（详解+图）](https://blog.csdn.net/qq_45748269/article/details/121622447)
### redis主从同步和从从同步架构图
![](https://img-blog.csdnimg.cn/56e475c04ff648e1bd13c46498ece7fb.png)


### 同步过程
Redis 2.8以前采用的复制都为全量复制，使用SYNC命令全量同步复制，SYNC存在很大的缺陷就是：不管slave是第一次启动，还是连接断开后的重连，主从同步都是全量数据复制，严重消耗master的资源以及大量的网络连接资源。

#### 全量同步(SYNC)
Redis全量复制一般发生在Slave初始化阶段，这时Slave需要将Master上的所有数据都复制一份。具体步骤如下：
1）从服务器连接主服务器，发送SYNC命令；
2）主服务器接收到SYNC命名后，开始执行BGSAVE命令生成RDB文件并使用缓冲区记录 此后执行的所有写命令；
3）主服务器BGSAVE执行完后，向所有从服务器发送快照文件，并在发送期间继续记录被执行的写命令；
4）从服务器收到快照文件后丢弃所有旧数据，载入收到的快照；
5）主服务器快照发送完毕后开始向从服务器发送缓冲区中的写命令；
6）从服务器完成对快照的载入，开始接收命令请求，并执行来自主服务器缓冲区的写命令；
完成上面几个步骤后就完成了从服务器数据初始化的所有操作，从服务器此时可以接收来自用户的读请求。


> 注意: 如果多个Slave断线了，需要重启的时候，因为只要Slave启动，就会发送sync请求和主机全量同步，当多个同时出现的时候，可能会导致Master IO剧增宕机。
从服务器在同步时，会清空所有数据，服务器在与主服务器进行初连接时，数据库中的所有数据都将丢失，替换成主服务器发送的数据。
主从复制不会阻塞master处理客户端请求，因为RDB是异步进行的；相反slave在初次同步数据时会阻塞不能处理客户端读请求。

> 当多个从服务器尝试连接同一个主服务器的时候，就会出现以下两种情况：
> 1.  如果master节点保存快照还未执行，所有从服务器都会接收到相同的快照文件和相同缓冲区写命令。
> 2. master节点保存快照正已经执行完毕，当主服务器与较早的从服务器完成以上全部步骤之后，主服务器会跟新连接的从服务器重新执行RDB保存快照。如果从服务器连接的时机不凑巧的话，进行多次RDB保存快照，会大量消耗主服务器资源（CPU、内存和磁盘I/O资源）。

#### 增量同步(PSYNC)
先了解几个概念:
1. runid(replication id, replid)
    每个Redis服务器都会有一个表明自己身份的ID。在PSYNC中发送的这个ID是指之前连接的Master的ID，如果没保存这个ID，PSYNC的命令会使用”PSYNC ? -1” 这种形式发送给Master，表示需要全量复制。 
查看命令: 使用cli执行 redis-cli -p 6379 info server | grep run
2. offset（复制偏移量）
    在主从复制的Master和Slave双方都会各自维持一个offset。Master成功发送N个字节的命令后会将Master里的offset加上N，Slave在接收到N个字节命令后同样会将Slave里的offset增加N。
Master和Slave如果状态是一致的那么它的的offset也应该是一致的。 
查看命令: 进入redis执行 info replication
3. 复制积压缓冲区(repl_baklog)
    复制积压缓冲区是由Master维护的一个固定长度环形积压队列(FIFO队列)，它的作用是缓存已经传播出去的命令。当Master进行命令传播时，不仅将命令发送给所有Slave，还会将命令写入到复制积压缓冲区里面。
本质上是一个固定长度的循环队列，积压队列越大，允许主从数据库断线的时间就越长,  默认情况下积压队列的大小为1MB，一旦写满，再有新数据写入时，就会覆盖数组的旧数据。此时slave来增量同步，发现自己offset已经被覆盖了，此时只能全量同步。可以通过配置文件设置队列大小：`repl-backlog-size 1mb`
Redis同时也提供了当没有slave需要同步的时候，多久可以释放环形队列，默认一小时：`repl-backlog-ttl 3600`

![](https://img-blog.csdnimg.cn/8fd33fabc5324c338620abd0991150fd.png)

1. 客户端向服务器发送SLAVEOF命令，即salve向master发起连接请求时，slave根据自己是否保存Master runid来判断是否是第一次连接。
2. 如果是第一次同步则向Master发送 PSYNC ? -1 命令来进行完整同步；如果是重连接，会向Master发送PSYNC runid offset命令(runid是master的身份ID，offset是从节点同步命令的全局迁移量)。
3. Master接收到PSYNC 命令后，首先判断runid是否和本机的id一致，如果一致则会再次判断offset偏移量和本机的偏移量相差有没有超过复制积压缓冲区大小，如果没有那么就给Slave发送CONTINUE，此时Slave只需要等待Master传回失去连接期间丢失的命令。
4. 如果runid和本机id不一致或者offset差距超过了复制积压缓冲区大小，那么就会返回FULLRESYNC runid offset，Slave将runid保存起来，并进行全量同步。

![](https://img-blog.csdnimg.cn/46487dd875db448ebda8ae7ed622756a.png)

![](https://img-blog.csdnimg.cn/172828f01a74463c83d015538a203c50.png)


Redis采用了乐观复制的策略，也就是在一定程度内容忍主从数据库的内容不一致，但是保持主从数据库数据的最终一致性。具体来说，Redis在主从复制的过程中，本身就是异步的，在主从数据库执行完客户端请求后会立即将结果返回给客户端，并异步的将命令同步给从数据库，但是这里并不会等待从数据库完全同步之后，再返回客户端。这一特性虽然保证了主从复制期间性能不受影响，但是也会产生一个数据不一致的时间窗口，如果在这个时间窗口期间网络突然断开连接，就会导致两者数据不一致。如果不在配置文件中添加其他策略，那就默认会采用这种方式，乐观二字也就体现在这里。   当然了，上面这种方式并不是绝对的，只要牺牲一点性能，还是可以避免上述问题。在配置文件中：
```
#代表至少N台从服务器完成复制，才允许主服务器可写入，否则会返回错误。
1. min-slaves-to-write 3
#允许从服务器断开连接的时间（单位s）
2. min-slaves-max-lag 10
```

#### 心跳检测
在命令传播阶段，从节点默认会以每秒一次的频率，向主节点发送命令。
`REPLCONF ACK <replication_offset>`
其中，replication_offset是从节点当前的复制偏移量。
发送REPLCONF ACK主从节点有三个作用：
1. 检测主从节点的网络连接状态。
2. 辅助实现min-slave选项。
3. 检查是否存在命令丢失。



#### 优化建议
1. 在master中配置repl-diskless-sync yes启用无磁盘复制，避免全量同步时的磁盘IO。
2. Redis单节点上的内存占用不要太大，减少RDB导致的过多磁盘IO
3. 适当提高repl_baklog的大小，发现slave宕机时尽快实现故障恢复，尽可能避免全量同步
4. 限制一个master上的slave节点数量，如果实在是太多slave，则可以采用主-从-从链式结构，减少master压力

#### 复制的相关参数
```
# replicaof 或者slaveof（5.0以前）命令
slaveof <masterip> <masterport> 
masterauth <master-password>

# 当主从连接中断，或主从复制建立期间，是否允许slave对外提供服务。默认为yes，即允许对外提供服务，但有可能会读到脏的数据。
slave-serve-stale-data yes 
# 将slave设置为只读模式。需要注意的是，只读模式针对的只是客户端的写操作，对于管理命令无效。
slave-read-only yes

# 是否使用无盘复制。为了降低主节点磁盘开销，Redis支持无盘复制，生成的RDB文件不保存到磁盘而是直接通过网络发送给从节点。
# 无盘复制适用于主节点所在机器磁盘性能较差但网络宽带较充裕的场景。
# 需要注意的是，无盘复制目前依然处于实验阶段。
repl-diskless-sync no
repl-diskless-sync-delay 5

# master每隔一段固定的时间向SLAVE发送一个PING命令。
repl-ping-slave-period 10

# 复制超时时间。
repl-timeout 60

# 设置为yes，主节点会等待一段时间才发送TCP数据包，具体等待时间取决于Linux内核，一般是40毫秒。适用于主从网络环境复杂或带宽紧张的场景。默认为no。
repl-disable-tcp-nodelay no

# 复制积压缓冲区，复制积压缓冲区是保存在主节点上的一个固定长度的队列。用于从Redis 2.8开始引入的部分复制。
repl-backlog-size 1mb
# 如果master上的slave全都断开了，且在指定的时间内没有连接上，则backlog会被master清除掉。repl-backlog-ttl即用来设置该时长，默认为3600s，如果设置为0，则永不清除。
repl-backlog-ttl 3600

# 设置slave的优先级，用于Redis Sentinel主从切换时使用，值越小，则提升为主的优先级越高。需要注意的是，如果设置为0，则代表该slave不参加选主。
slave-priority 100

# 防止脑裂参数
# 从服务器的数量少于3个，或者三个从服务器的延迟（lag）值都大于或等于10秒时，主服务器将拒绝执行写命令，这里的延迟值就是上面提到的INFO replication命令的lag 值
min-slaves-to-write 3
min-slaves-max-lag 10

# 常用于端口转发或NAT场景下，对Master暴露真实IP和端口信息。
slave-announce-ip 5.5.5.5
slave-announce-port 1234
```

#### 从节点提升为主节点
当一个从节点被提升为主节点时，其它的从节点必须与新主节点重新同步。在Redis 4.0 之前，因为master_replid发生了变化，所以这个过程是一个全量同步。在Redis 4.0之后，新主节点会记录旧主节点的master_replid和offset，因为能够接受来自其它从节点的增量同步请求，即使请求中的master_replid不同。在底层实现上，当执行slaveof no one时，会将master_replid，master_repl_offset+1复制为master_replid，second_repl_offset。


经过测试: `主->从1,从2` 到 `主->从1->从2`的转变不会触发全量同步
