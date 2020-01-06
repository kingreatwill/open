## MySQL性能调优
### MySQL数据库配置优化

- 表示缓冲池字节大小。
- 推荐值为物理内存的50%~80%。
- innodb_buffer_pool_size
- 用来控制redo log刷新到磁盘的策略。
- innodb_flush_log_at_trx_commit=1
- 每提交1次事务同步写到磁盘中，可以设置为n。
- sync_binlog=1
- 脏页占innodb_buffer_pool_size的比例时，触发刷脏页到磁盘。 推荐值为25%~50%。
- innodb_max_dirty_pages_pct=30
- 后台进程最大IO性能指标。
- 默认200，如果SSD，调整为5000~20000
- innodb_io_capacity=200
- 指定innodb共享表空间文件的大小。
- innodb_data_file_path
- 慢查询日志的阈值设置，单位秒。
- long_qurey_time=0.3
- mysql复制的形式，row为MySQL8.0的默认形式。
- binlog_format=row
- 调高该参数则应降低interactive_timeout、wait_timeout的值。
- max_connections=200
- 过大，实例恢复时间长；过小，造成日志切换频繁。
- innodb_log_file_size
- 全量日志建议关闭。
- 默认关闭。
- general_log=0

#### centos 内核相关参数(/etc/sysctl.conf)

以下参数可以直接放到sysctl.conf文件的末尾。
1. 增加监听队列上限：
net.core.somaxconn = 65535
net.core.netdev_max_backlog = 65535
net.ipv4.tcp_max_syn_backlog = 65535

2. 加快TCP连接的回收：
net.ipv4.tcp_fin_timeout = 10
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_tw_recycle = 1

3. TCP连接接收和发送缓冲区大小的默认值和最大值:
net.core.wmem_default = 87380
net.core.wmem_max = 16777216
net.core.rmem_default = 87380
net.core.rmem_max = 16777216

4. 减少失效连接所占用的TCP资源的数量，加快资源回收的效率：
net.ipv4.tcp_keepalive_time = 120
net.ipv4.tcp_keepalive_intvl = 30
net.ipv4.tcp_keepalive_probes = 3

5. 单个共享内存段的最大值：
kernel.shmmax = 4294967295
这个参数应该设置的足够大，以便能在一个共享内存段下容纳整个的Innodb缓冲池的大小。
这个值的大小对于64位linux系统，可取的最大值为(物理内存值-1)byte，建议值为大于物理内存的一半，一般取值大于Innodb缓冲池的大小即可。

6. 控制换出运行时内存的相对权重：
vm.swappiness = 0
这个参数当内存不足时会对性能产生比较明显的影响。（设置为0，表示Linux内核虚拟内存完全被占用，才会要使用交换区。）
> Linux系统内存交换区：
在Linux系统安装时都会有一个特殊的磁盘分区，称之为系统交换分区。
使用 free -m 命令可以看到swap就是内存交换区。
作用：当操作系统没有足够的内存时，就会将部分虚拟内存写到磁盘的交换区中，这样就会发生内存交换。

如果Linux系统上完全禁用交换分区，带来的风险：
1. 降低操作系统的性能
2. 容易造成内存溢出，崩溃，或都被操作系统kill掉

#### centos 增加资源限制(/etc/security/limit.conf)
打开文件数的限制(以下参数可以直接放到limit.conf文件的末尾)：

* soft nofile 65535
* hard nofile 65535

>*：表示对所有用户有效
soft：表示当前系统生效的设置（soft不能大于hard ）
hard：表明系统中所能设定的最大值
nofile：表示所限制的资源是打开文件的最大数目
65535：限制的数量

以上两行配置将可打开的文件数量增加到65535个，以保证可以打开足够多的文件句柄。
>注意：这个文件的修改需要重启系统才能生效。

#### centos 磁盘调度策略
1. cfq (完全公平队列策略，Linux2.6.18之后内核的系统默认策略)
该模式按进程创建多个队列，各个进程发来的IO请求会被cfq以轮循方式处理，对每个IO请求都是公平的。该策略适合离散读的应用。

2. deadline (截止时间调度策略)
deadline，包含读和写两个队列，确保在一个截止时间内服务请求（截止时间是可调整的），而默认读期限短于写期限。这样就防止了写操作因为不能被读取而饿死的现象，deadline对数据库类应用是最好的选择。

3. noop (电梯式调度策略)
noop只实现一个简单的FIFO队列，倾向饿死读而利于写，因此noop对于闪存设备、RAM及嵌入式系统是最好的选择。

4. anticipatory (预料I/O调度策略)
本质上与deadline策略一样，但在最后一次读操作之后，要等待6ms，才能继续进行对其它I/O请求进行调度。它会在每个6ms中插入新的I/O操作，合并写入流，用写入延时换取最大的写入吞吐量。anticipatory适合于写入较多的环境，比如文件服务器。该策略对数据库环境表现很差。

查看调度策略的方法：

cat /sys/block/devname/queue/scheduler

修改调度策略的方法：

echo <schedulername> > /sys/block/devname/queue/scheduler