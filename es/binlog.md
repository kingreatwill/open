https://dev.mysql.com/doc/refman/5.5/en/binary-log.html

- mysql的binlog是多文件存储，定位一个LogEvent需要通过binlog filename + binlog position，进行定位
- mysql的binlog数据格式，按照生成的方式，主要分为：statement-based、row-based、mixed。
show variables like 'binlog_format';

目前canal支持所有模式的增量订阅(但配合同步时，因为statement只有sql，没有数据，无法获取原始的变更日志，所以一般建议为ROW模式)

## Statement
### Statement 优点
历史悠久，技术成熟；
产生的 binlog 文件较小；
binlog 中包含了所有数据库修改信息，可以据此来审核数据库的安全等情况；
binlog 可以用于实时的还原，而不仅仅用于复制；
主从版本可以不一样，从服务器版本可以比主服务器版本高；

### Statement 缺点：
不是所有的 UPDATE 语句都能被复制，尤其是包含不确定操作的时候；
调用具有不确定因素的 UDF 时复制也可能出现问题；
运用以下函数的语句也不能被复制：
* LOAD_FILE()
* UUID()
* USER()
* FOUND_ROWS()
* SYSDATE() (除非启动时启用了 –sysdate-is-now 选项)
INSERT … SELECT 会产生比 RBR 更多的行级锁；
复制须要执行全表扫描 (WHERE 语句中没有运用到索引) 的 UPDATE 时，须要比 row 请求更多的行级锁；
对于有 AUTO_INCREMENT 字段的 InnoDB 表而言，INSERT 语句会阻塞其他 INSERT 语句；
对于一些复杂的语句，在从服务器上的耗资源情况会更严重，而 row 模式下，只会对那个发生变化的记录产生影响；
存储函数(不是存储流程 )在被调用的同时也会执行一次 NOW() 函数，这个可以说是坏事也可能是好事；
确定了的 UDF 也须要在从服务器上执行；
数据表必须几乎和主服务器保持一致才行，否则可能会导致复制出错；
执行复杂语句如果出错的话，会消耗更多资源；

## ROW
### Row 优点

任何情况都可以被复制，这对复制来说是最安全可靠的；
和其他大多数数据库系统的复制技能一样；
多数情况下，从服务器上的表如果有主键的话，复制就会快了很多；
复制以下几种语句时的行锁更少：
* INSERT … SELECT
* 包含 AUTO_INCREMENT 字段的 INSERT
* 没有附带条件或者并没有修改很多记录的 UPDATE 或 DELETE 语句
执行 INSERT，UPDATE，DELETE 语句时锁更少；
从服务器上采用多线程来执行复制成为可能；

### Row 缺点
生成的 binlog 日志体积大了很多；
复杂的回滚时 binlog 中会包含大量的数据；
主服务器上执行 UPDATE 语句时，所有发生变化的记录都会写到 binlog 中，而 statement 只会写一次，这会导致频繁发生 binlog 的写并发请求；
UDF 产生的大 BLOB 值会导致复制变慢；
不能从 binlog 中看到都复制了写什么语句(加密过的)；
当在非事务表上执行一段堆积的 SQL 语句时，最好采用 statement 模式，否则很容易导致主从服务器的数据不一致情况发生；
另外，针对系统库 MySQL 里面的表发生变化时的处理准则如下：
如果是采用 INSERT，UPDATE，DELETE 直接操作表的情况，则日志格式根据 binlog_format 的设定而记录；
如果是采用 GRANT，REVOKE，SET PASSWORD 等管理语句来做的话，那么无论如何都要使用 statement 模式记录；
使用 statement 模式后，能处理很多原先出现的主键重复问题；

## Mixed
是以上两种level的混合使用，一般的语句修改使用statment格式保存binlog，如一些函数，statement无法完成主从复制的操作，则采用row格式保存binlog,MySQL会根据执行的每一条具体的sql语句来区分对待记录的日志形式，也就是在Statement和Row之间选择一种.新版本的MySQL中队row level模式也被做了优化，并不是所有的修改都会以row level来记录，像遇到表结构变更的时候就会以statement模式来记录。至于update或者delete等修改数据的语句，还是会记录所有行的变更。
在slave日志同步过程中，对于使用now这样的时间函数，MIXED日志格式，会在日志中产生对应的unix_timestamp()*1000的时间字符串，slave在完成同步时，取用的是sqlEvent发生的时间来保证数据的准确性。另外对于一些功能性函数slave能完成相应的数据同步，而对于上面指定的一些类似于UDF函数，导致Slave无法知晓的情况，则会采用ROW格式存储这些Binlog，以保证产生的Binlog可以供Slave完成数据同步。