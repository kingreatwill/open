
https://redis.io/documentation
https://www.toutiao.com/a6773538740674494987

[Redis对象——有序集合(ZSet)](https://www.cnblogs.com/hunternet/p/12717643.html)
https://www.cnblogs.com/hunternet/tag/Redis/


[reids 集群](http://www.redis.cn/topics/cluster-tutorial.html)
<!-- toc -->
[TOC]

## 管理工具
RedisDesktopManager 收费
https://github.com/qishibo/AnotherRedisDesktopManager 免费-目前功能不足
https://github.com/fastogt/fastonosql 收费
https://github.com/fastogt/fastoredis 收费
keylord(收费)

medis(可以自己编译)
https://github.com/luin/medis


tablePlus(收费)

## 数据结构
### 普通数据结构
#### String
#### List
#### Hash
#### Set
#### Sorted Set
### 高级数据结构 
#### bitmap
#### GEO
#### HyperLogLog
### Streams

## 使用场景
### 分布式Session
### 分布式锁

### 全局ID
incrby userid 1000
### 计数器
incr方法
### 限流
incr方法
### 位统计
String类型的bitcount（1.6.6的bitmap数据结构介绍）

字符是以8位二进制存储的
```
set k1 a
setbit k1 6 1
setbit k1 7 0
get k1
/* 6 7 代表的a的二进制位的修改
a 对应的ASCII码是97，转换为二进制数据是01100001
b 对应的ASCII码是98，转换为二进制数据是01100010

因为bit非常节省空间（1 MB=8388608 bit），可以用来做大数据量的统计。
*/
```
例如：在线用户统计，留存用户统计
```
setbit onlineusers 01
setbit onlineusers 11
setbit onlineusers 20
```
支持按位与、按位或等等操作
```
BITOPANDdestkeykey[key...] ，对一个或多个 key 求逻辑并，并将结果保存到 destkey 。
BITOPORdestkeykey[key...] ，对一个或多个 key 求逻辑或，并将结果保存到 destkey 。
BITOPXORdestkeykey[key...] ，对一个或多个 key 求逻辑异或，并将结果保存到 destkey 。
BITOPNOTdestkeykey ，对给定 key 求逻辑非，并将结果保存到 destkey 。
```
计算出7天都在线的用户
```
BITOP "AND" "7_days_both_online_users" "day_1_online_users" "day_2_online_users" ...  "day_7_online_users"
```
### 购物车
String 或hash。所有String可以做的hash都可以做
- key：用户id；field：商品id；value：商品数量。
- +1：hincr。-1：hdecr。删除：hdel。全选：hgetall。商品数：hlen。
### 用户消息时间线timeline
list，双向链表，直接作为timeline就好了。插入有序
### 消息队列
List提供了两个阻塞的弹出操作：blpop/brpop，可以设置超时时间

- blpop：blpop key1 timeout 移除并获取列表的第一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
- brpop：brpop key1 timeout 移除并获取列表的最后一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。

上面的操作。其实就是java的阻塞队列。学习的东西越多。学习成本越低

- 队列：先进先除：rpush blpop，左头右尾，右边进入队列，左边出队列
- 栈：先进后出：rpush brpop
### 抽奖
自带一个随机获得值`spop myset`

### 点赞、签到、打卡
假如微博ID是t1001，用户ID是u3001

用 like:t1001 来维护 t1001 这条微博的所有点赞用户

- 点赞了这条微博：sadd like:t1001 u3001
- 取消点赞：srem like:t1001 u3001
- 是否点赞：sismember like:t1001 u3001
- 点赞的所有用户：smembers like:t1001
- 点赞数：scard like:t1001

### 商品标签
老规矩，用 tags:i5001 来维护商品所有的标签。

- sadd tags:i5001 画面清晰细腻
- sadd tags:i5001 真彩清晰显示屏
- sadd tags:i5001 流程至极
### 商品筛选
```
// 获取差集
sdiff set1 set2
// 获取交集（intersection ）
sinter set1 set2
// 获取并集
sunion set1 set2
```
假如：iPhone11 上市了
```
sadd brand:apple iPhone11

sadd brand:ios iPhone11

sad screensize:6.0-6.24 iPhone11

sad screentype:lcd iPhone 11
```
赛选商品，苹果的、ios的、屏幕在6.0-6.24之间的，屏幕材质是LCD屏幕
```
sinter brand:apple brand:ios screensize:6.0-6.24 screentype:lcd
```
### 用户关注、推荐模型
follow 关注 fans 粉丝

相互关注：
- sadd 1:follow 2
- sadd 2:fans 1
- sadd 1:fans 2
- sadd 2:follow 1

我关注的人也关注了他(取交集)：
- sinter 1:follow 2:fans

可能认识的人：
- 用户1可能认识的人(差集)：sdiff 2:follow 1:follow
- 用户2可能认识的人：sdiff 1:follow 2:follow

### 排行榜
id 为6001 的新闻点击数加1：`zincrby hotNews:20190926 1 n6001`
获取今天点击最多的15条：`zrevrange hotNews:20190926 0 15 withscores`
