


## Hbase输入Int类型数据默认转换为字符类型的原因及缺陷
1.Hbase shell 中插入带int类型的数据

在Hbase的shell命令中输入插入一行value为int类型的数据

put 'hbase1','row','hb1:age',30

Hbase会将int类型数据自动转换为字符类型来存储。方便我们在shell中观察和操作。

2.在Java API中编写插入带int类型的数据

在Java API插入int型数据会保存为ASCII形式。这样我们就只能看出是ASCII码。

3.Hbase shell默认转换的缺陷

在业务需求中出现比较大小的时候，这时候就没法比了，然而使用Java API操作的是ASCII码。

取出来的是Int类型的数据。可以进行比较。

### 解决方案
put 'table','rowkey','cf:qua', Bytes.toBytes(1234)