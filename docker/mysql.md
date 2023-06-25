# docker 安装mysql
[教程](https://hub.docker.com/_/mysql/?tab=description)


安装5.7.27
```
docker pull mysql:5.7.27   #373M

D盘新建dockerv  mysql data/conf/logs 文件夹

docker run  -itd -p 3306:3306 -v D:/dockerv/mysql/conf:/etc/mysql/conf.d -v D:/dockerv/mysql/logs:/logs -v D:/dockerv/mysql/data:/var/lib/mysql --restart always -e MYSQL_ROOT_PASSWORD=123456@lcb --name mysql5.7.27  mysql:5.7.27 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

然后就可以连接了

```

安装8.0.17
```
docker pull mysql:8.0.17   #445M

D盘新建dockerv  mysql8 data/conf/logs 文件夹

docker run  -itd -p 3306:3306 -v D:/dockerv/mysql8/conf:/etc/mysql/conf.d -v D:/dockerv/mysql8/logs:/logs -v D:/dockerv/mysql8/data:/var/lib/mysql --restart always -e MYSQL_ROOT_PASSWORD=123456 --name mysql8  mysql:8 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

> 也可以设置成utf8mb4_general_ci

进入MySQL容器,登陆MySQL
docker exec -it mysql8 /bin/bash

mysql -u root -p

 mysql8.0的root用户的验证方式变了，通过查询：
  use mysql;
 select host,user,plugin,authentication_string from mysql.user;

得知：root的用户的加密方式为caching_sha2_passoword, 而navicat连接所用的方式为native_password。mysql为远程连接和本地连接提供了不同的密码验证方式。

备注：host为 % 表示不限制ip   localhost表示本机使用    plugin非mysql_native_password 则需要修改密码

ALTER user 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';

FLUSH PRIVILEGES;  

然后就可以远程登陆MySQL

也可以新增用户

CREATE USER 'demo' IDENTIFIED WITH mysql_native_password BY '123456';
GRANT all ON *.* TO 'demo'@'%';
FLUSH PRIVILEGES;
```

> ci即case insensitive，不区分大小写。

> **utf8mb4_unicode_ci**
是基于标准的Unicode来排序和比较，能够在各种语言之间精确排序，Unicode排序规则为了能够处理特殊字符的情况，实现了略微复杂的排序算法。
> **utf8mb4_general_ci**
是一个遗留的 校对规则，不支持扩展，它仅能够在字符之间进行逐个比较。utf8_general_ci校对规则进行的比较速度很快，但是与使用 utf8mb4_unicode_ci的校对规则相比，比较正确性较差。
> **utf8mb4_0900_ai_ci**：
MySQL 8.0 默认的是 utf8mb4_0900_ai_ci，属于 utf8mb4_unicode_ci 中的一种，具体含义如下：
uft8mb4 表示用 UTF-8 编码方案，每个字符最多占 4 个字节。
0900 指的是 Unicode 校对算法版本。（Unicode 归类算法是用于比较符合 Unicode 标准要求的两个 Unicode 字符串的方法）。
ai 指的是口音不敏感。也就是说，排序时 e，è，é，ê 和 ë 之间没有区别。
ci 表示不区分大小写。也就是说，排序时 p 和 P 之间没有区别。
utf8mb4 已成为默认字符集，在 MySQL 8.0.1 及更高版本中将 utf8mb4_0900_ai_ci 作为默认排序规则。以前，utf8mb4_general_ci 是默认排序规则。由于 utf8mb4_0900_ai_ci 排序规则现在是默认排序规则，因此默认情况下新表格可以存储基本多语言平面之外的字符。现在可以默认存储表情符号。如果需要重音灵敏度和区分大小写，则可以使用 utf8mb4_0900_as_cs 代替。

## 还原数据库
- 使用RESTORE FILELISTONLY命令列出备份数据文件的逻辑名
```
/opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P 'dev@123,' -Q 'RESTORE FILELISTONLY FROM DISK = "/var/opt/mssql/testdb.bak"' | tr -s ' ' | cut -d ' ' -f 1-2
```
使用该命令可以把数据库的数据文件，日志文件名称显示出来。在接下来的恢复操作中有用。
- 使用RESTORE DATABASE命令还原数据库
```
/opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P 'dev@123,' -Q 'RESTORE DATABASE testdb FROM DISK = "/var/opt/mssql/testdb.bak" WITH MOVE "testdb" TO "/var/opt/mssql/data/testdb.mdf" , MOVE "testdb_log" TO "/var/opt/mssql/data/testdb.ldf"'
```