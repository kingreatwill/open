<!--toc-->
[TOC]
# Redis 6.0 ACL
期待已久的ACL终于来了，大家知道在redis集群中只有一个db，在多项目操作时可以通过key前缀来区分，但是还是能获取其它key，这样就带来了安全风险.

http://antirez.com/news/131

Redis 6 会给大家提供的新功能，包括：
* Many new modules APIs.
* Better expire cycle.
* SSL
* ACLs
* RESP3
* Client side caching
* Threaded I/O
* Diskless replication on replicas
* Redis-benchmark cluster support + Redis-cli improvements
* Systemd support rewrite.
* Redis Cluster proxy was released with Redis 6 (but different repository).
* A Disque module was released with Redis 6 (but different repository).

- 一、对用户使用有直接影响的功能
  - ACL用户权限控制功能
  - RESP3：新的 Redis 通信协议
  - 客戶端緩存（RESP3协议更好的支持客户端缓存）
  - Cluster 管理工具
  - SSL 支持
- 二、Redis 内部的优化
  - IO多线程支持
  - 新的Module API
  - 新的 Expire 算法
  - 副本的无盘复制
  - Systemd 重写
- 三、外部工具
  - Redis Cluster Proxy 与redis6 一起发布(不同仓库)
  - Disque(不同仓库)


## Access Control Lists ACL
在之前 登录可以`AUTH pwd` 有了ACL后`AUTH user pwd`

### ACL LIST
我们可以使用 ACL LIST 命令来检查当前活动的 ACL
```
> ACL LIST
1) "user default on nopass ~* +@all"
```
user 代表是用户
default 代表默认用户（反之 为自己创建的用户）
on 代表激活（反之off,默认新增的为off）
nopass 代表不需要密码
~* 代表可以访问的key
+@all 代表可以操作的command

key
```
~<pattern>  ~* 所有key
allkeys 别名 ~*
resetkeys 刷新允许的键模式列表
可以多个一起使用，如：
~foo:* ~bar:* resetkeys ~objects:*
只能访问 ~objects:*
```

command
```
+<command> 添加命令
-<command> 移除命令
+@<category> 添加一类命令，如：@admin, @set, @sortedset, ... 可以ACL CAT 查看

如：+@geo -@readonly 会排除geo的只读命令

-@<category> 移除一类命令
+<command>|subcommand  允许一个特定的子命令，没有-<command>|subcommand
allcommands 所有cmd 别名 +@all
nocommands 别名 -@all
```

配置密码
```
>pwd  添加密码列表
<pwd  移除密码列表

#hash 添加SHA-256值
!hash 移除SHA-256值

nopass 移除所有密码

resetpass 刷新密码列表，并且也不会回到nopass 状态，之后一定要添加密码
没有使用nopass标记的用户并且没有有效密码列表  是不能登录的


reset  重置用户 实际上是执行 resetpass, resetkeys, off, -@all
```
> 另外，在默认用户的特殊情况下，拥有 nopass 规则意味着新的连接将通过默认用户的自动身份验证，而不需要任何显式的 AUTH 调用。

### ACL WHOAMI
查看当前用户(注意：当前用户需要有此权限)who am i


### ACL CAT

第一大类
```
127.0.0.1:0>ACL CAT
 1)  "keyspace"
 2)  "read"
 3)  "write"
 4)  "set"
 5)  "sortedset"
 6)  "list"
 7)  "hash"
 8)  "string"
 9)  "bitmap"
 10)  "hyperloglog"
 11)  "geo"
 12)  "stream"
 13)  "pubsub"
 14)  "admin"
 15)  "fast"
 16)  "slow"
 17)  "blocking"
 18)  "dangerous"
 19)  "connection"
 20)  "transaction"
 21)  "scripting"
```
```
127.0.0.1:0>ACL CAT set
 1)  "spop"
 2)  "sunionstore"
 3)  "sinter"
 4)  "sdiffstore"
 5)  "sscan"
 6)  "sunion"
 7)  "smembers"
 8)  "sinterstore"
 9)  "sismember"
 10)  "smove"
 11)  "srem"
 12)  "sdiff"
 13)  "srandmember"
 14)  "sort"
 15)  "sadd"
 16)  "scard"
```


### ACL SETUSER
添加用户 - 区分大小写
```
ACL SETUSER alice
```
检查一下默认的用户状态:
```
> ACL LIST
1) "user alice off -@all"
2) "user default on nopass ~* +@all"
```
### 设置密码
```
ACL SETUSER alice on >p1pp0 ~cached:* +get

> AUTH alice p1pp0
OK
> GET foo
(error) NOPERM this user has no permissions to access one of the keys used as arguments
> GET cached:1234
(nil)
> SET cached:1234 zap
(error) NOPERM this user has no permissions to run the 'set' command or its subcommnad
```

### 查看单独用户
打开一个新的连接 
```
> ACL GETUSER alice
1) "flags"
2) 1) "on"
3) "passwords"
4) 1) "2d9c75..."
5) "commands"
6) "-@all +get"
7) "keys"
8) 1) "cached:*"
```
返回 field-value 数组

如果使用RESP3
```
> ACL GETUSER alice
1# "flags" => 1~ "on"
2# "passwords" => 1) "2d9c75..."
3# "commands" => "-@all +get"
4# "keys" => 1) "cached:*"
```

### 新增权限
```
ACL SETUSER alice ~objects:* ~items:* ~public:*

> ACL LIST
1) "user alice on #2d9c75... ~cached:* ~objects:* ~items:* ~public:* -@all +get"
2) "user default on nopass ~* +@all"

ACL SETUSER alice +set
ACL SETUSER alice +get

> ACL LIST
1) "user alice on #2d9c75... ~cached:* ~objects:* ~items:* ~public:* -@all +get +set"
2) "user default on nopass ~* +@all"
```

## 其它
### ACL GENPASS
生成随机数、密码

### redis.conf
配置文件中可以保存用户信息
如： user worker +@list +@connection ~jobs:* on >ffa9203c493aa99
也可以指定外部acl文件信息
如：aclfile /etc/redis/users.acl

ACL LOAD 重置加载acl文件（如果你修改了acl文件）
ACL SAVE 保存acl文件

### 哨兵和副本
略...

[参考](https://redis.io/topics/acl)
