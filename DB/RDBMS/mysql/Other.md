[SQL优化 MySQL版，explain SQL执行计划详解](https://www.toutiao.com/article/7100906354843353604/)

[慢 SQL 分析与优化](https://www.toutiao.com/article/7122345201695212064/)

[MySQL之Online DDL再探](https://mp.weixin.qq.com/s?__biz=MzUyNjkzNjQwMQ==&mid=2247485550&idx=1&sn=5760ec508371cbb879f71088a76e3c5f&chksm=fa067850cd71f146bec5933a559a1a0291a42f6f038bca78c3d75c470188d2cfffc782472294&scene=21#wechat_redirect)

## MySQL数据库巡检脚本
```sh
#!/bin/bash

# MySQL巡检脚本

# 设置MySQL用户名和密码（请将它们设置为适当的值）
MYSQL_USER="your_username"
MYSQL_PASSWORD="your_password"

# 获取MySQL版本信息
MYSQL_VERSION=$(mysql -u ${MYSQL_USER} -p${MYSQL_PASSWORD} -e "SELECT VERSION();" | awk 'NR==2{print $1}')

# 获取MySQL运行状态信息
STATUS=$(systemctl status mysql.service)

# 获取MySQL进程列表
PROCESS_LIST=$(mysql -u ${MYSQL_USER} -p${MYSQL_PASSWORD} -e "SHOW PROCESSLIST;" | awk '{print $1,$2,$3,$4,$5,$6}')

# 检查MySQL是否在运行
if [[ "$STATUS" =~ "active (running)" ]]; then
  MYSQL_RUNNING="YES"
else
  MYSQL_RUNNING="NO"
fi

# 检查MySQL进程是否存在
if [[ -z "$PROCESS_LIST" ]]; then
  MYSQL_PROCESS="NO"
else
  MYSQL_PROCESS="YES"
fi

# 生成报告
echo "MySQL巡检报告"
echo "----------------"
echo "MySQL版本: $MYSQL_VERSION"
echo "MySQL运行状态: $MYSQL_RUNNING"
echo "MySQL进程存在: $MYSQL_PROCESS"
echo ""
echo "MySQL进程列表"
echo "----------------"
echo "$PROCESS_LIST"
```

缓存池的使用

缓冲池大小：`SHOW GLOBAL VARIABLES LIKE 'innodb_buffer_pool_size';`
缓冲池使用率：`SHOW GLOBAL STATUS LIKE 'Innodb_buffer_pool_pages_free';`

查询缓存

查询缓存命中率：`SHOW GLOBAL STATUS LIKE 'Qcache_hits';`
查询缓存碎片率：`SHOW GLOBAL STATUS LIKE 'Qcache_free_memory';`

读写比例和吞吐量

读写比例：`SHOW GLOBAL STATUS LIKE 'Com_select';和SHOW GLOBAL STATUS LIKE 'Com_insert';`
吞吐量：`SHOW GLOBAL STATUS LIKE 'Innodb_rows_read';和SHOW GLOBAL STATUS LIKE 'Innodb_rows_inserted';`

锁的相关信息

表锁争用：`SHOW GLOBAL STATUS LIKE 'Table_locks_waited';`
行锁等待时间：`SHOW GLOBAL STATUS LIKE 'Innodb_row_lock_time';`
行锁等待时间平均值：`SHOW GLOBAL STATUS LIKE 'Innodb_row_lock_time_avg';`