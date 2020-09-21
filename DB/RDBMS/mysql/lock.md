https://www.cnblogs.com/chenqionghe/p/4845693.html

[未验证]MySql的优化器强制转化为匹配的类型，导致行锁升级为表锁。所以开发中一定要注意类型的匹配，避免行锁升级为表锁，影响并发性能。

update  where 索引列=‘xxx’
当更新数据的列存在普通索引时（尤其是普通索引重复率高时）会发生表锁
当操作的数据占整个数据的比例较大时，行锁将会升级

https://www.cnblogs.com/huangrenhui/p/12580966.html

间隙锁：https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_gap_lock
记录锁：https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_record_lock

[InnoDB Locking](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html)

#### Shared and Exclusive Locks 共享锁（S锁）和独占锁（X锁）
### Intention Locks 意向锁  
intention shared lock (IS)
intention exclusive lock (IX)

For example, SELECT ... LOCK IN SHARE MODE sets an IS lock, and SELECT ... FOR UPDATE sets an IX lock.


.|X|	IX|	S|	IS
---|---|---|---|---
X	|Conflict	|Conflict	|Conflict	|Conflict
IX	|Conflict	|Compatible	|Conflict	|Compatible
S	|Conflict	|Conflict	|Compatible	|Compatible
IS	|Conflict	|Compatible	|Compatible	|Compatible

### Record Locks 记录锁
A record lock is a lock on an index record. For example, SELECT c1 FROM t WHERE c1 = 10 FOR UPDATE; prevents any other transaction from inserting, updating, or deleting rows where the value of t.c1 is 10.

### Gap Locks 间隙锁 
GAP锁的目的，是为了防止同一事务的两次当前读，出现幻读的情况。

A gap lock is a lock on a gap between index records, or a lock on the gap before the first or after the last index record. For example, SELECT c1 FROM t WHERE c1 BETWEEN 10 and 20 FOR UPDATE; prevents other transactions from inserting a value of 15 into column t.c1, whether or not there was already any such value in the column, because the gaps between all existing values in the range are locked.


### Next-Key Locks
1+2，锁定一个范围，并且锁定记录本身。对于行的查询，都是采用该方法，主要目的是解决幻读的问题。

索引记录之间的锁

Suppose that an index contains the values 10, 11, 13, and 20. The possible next-key locks for this index cover the following intervals, where a round bracket denotes exclusion of the interval endpoint and a square bracket denotes inclusion of the endpoint:

```
(negative infinity, 10]
(10, 11]
(11, 13]
(13, 20]
(20, positive infinity)
```

### Insert Intention Locks 插入意向锁
INSERT INTO child (id) values (90),(102);

START TRANSACTION;// 1
SELECT * FROM child WHERE id > 100 FOR UPDATE;

START TRANSACTION;// 2
INSERT INTO child (id) VALUES (101);

1没有结束时，2是不能执行成功的
事务1在等待IX意向排他锁（意向排他锁包括102之前的记录）时 同时获取插入意向锁
插入意向锁与gap锁是不兼容的，所以，事务2的插入操作会被阻塞。

### AUTO-INC Locks 自增锁

### Predicate Locks for Spatial Indexes
Innodb 支持对包含空间列的列进行 SPATIAL 索引(参见11.4.8节“优化空间分析”)。 为了处理与 SPATIAL 索引有关的操作的锁定，下一个键锁定不能很好地支持 REPEATABLE READ 或 SERIALIZABLE 事务隔离级别。 多维数据中没有绝对排序概念，因此不清楚哪个是“下一个”键。 为了支持具有 SPATIAL 索引的表的隔离级别，InnoDB 使用谓词锁。 Spatial 索引包含最小外接矩形值，因此 InnoDB 通过在用于查询的 MBR 值上设置谓词锁来强制对索引进行一致读取。 其他事务不能插入或修改与查询条件匹配的行。


事务ACID特性，然而隔离级别又会打破ACID特性
这些锁都是为了事务的隔离级别而做的

为了高效，mysql 的 隔离性使用了MVCC（MVCC/SNAPSHO IOSLATION）

[阿里沈询：分布式事务原理与实践](http://jm.taobao.org/2017/02/09/20170209/)

这里有介绍事务原理

http://i.youku.com/u/UMTcwMTg3NDc1Mg==
https://yq.aliyun.com/edu/lesson/play/508

### 在MySQL中使用insert into table1 select * from table2时，会对table2进行加锁
使用`select @@global.tx_isolation,@@session.tx_isolation;`查询事务隔离级别，REPEATABLE-READ是MySQL默认的事务隔离级别
使用`show engine innodb status;`查询锁状况
[MySQL insert into select锁表的问题（上）](https://blog.csdn.net/llliarby/article/details/78697327)

[MySQL insert into select锁表的问题（下）](https://blog.csdn.net/llliarby/article/details/78698269)

