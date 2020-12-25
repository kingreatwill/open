mysql参数
```
默认值:
read_buffer_size 128kb
sort_buffer_size 256kb
innodb_sort_buffer_size 1M
read_rnd_buffer_size 256kb
join_buffer_size 256kb
tmp_table_size 16M
table_open_cache 4000

innodb_buffer_pool_size 128M
key_buffer_size 8M

innodb_page_size 16kb

show global status;
show global variables  like '%query_cache%'; -- global   like '%s%';


8G内存建议配置：

read_buffer_size 4M
sort_buffer_size 4M
read_rnd_buffer_size 4M
tmp_table_size 512M

table_open_cache 4000

innodb_buffer_pool_size 4G 改成db服务器总内存的60% 到80%
key_buffer_size 512M

innodb_page_size 16kb
max_allowed_packet 
```


## MySQL新增用户及赋予权限

创建用户
```
USE mysql; #创建用户需要操作 mysql 表
# 语法格式为 [@'host']  host 为 'localhost' 表示本地登录用户，host 为 IP地址或 IP 地址区间，表示指定IP地址的主机可登录，host 为 "%"，表示所有主机都可登录，省略代表所有主机
CREATE USER 'username'[@'host'] IDENTIFIED BY 'password';
# eg. 常见 local_user 用户可以在所有主机登录，密码为 123456
CREATE USER 'local_user' IDENTIFIED BY '123456';
# eg. 创建 local_user 只允许在本地登录
CREATE USER 'local_user'@'localhost' IDENTIFIED BY '123456';
```
查看用户权限
```
# 可以通过查询 user 表获取 语法格式为
SELECT  privileges|* FROM user WHERE `user` = 'username';
# eg. 查看 local_user 的权限
SELECT * FROM user WHERE `user` = 'local_user';
# 也可以用 SHOW GRANTS 查看
SHOW GRANTS FOR 'username' [@host];
# eg.
SHOW GRANTS FOR local_user;
```
赋予用户权限
```
# 语法格式
GRANT privileges ON database.table TO 'username'@'host' [IDENTIFIED BY 'password'];
# eg. 赋予 local_user 在所有主机的所有权限，但不包含给其他账号赋予权限的权限
GRANT all ON *.* TO 'local_user'@'%';
# 刷新权限 权限更新后刷新才会起作用
FLUSH PRIVILEGES;
```

- GRANT命令说明：
    - priveleges (权限列表),可以是all, 表示所有权限，也可以是select,update等权限，多个权限的名词,相互之间用逗号分开。
    - ON 用来指定权限针对哪些库和表。格式为数据库 .表名 ，点号前面用来指定数据库名，点号后面用来指定表名，*.* 表示所有数据库所有表。
    - TO 表示将权限赋予某个用户, 格式为username@host，@前面为用户名，@后面接限制的主机，可以是IP、IP段、域名以及%，%表示任何地方。注意：这里%有的版本不包括本地，以前碰到过给某个用户设置了%允许任何地方登录，但是在本地登录不了，这个和版本有关系，遇到这个问题再加一个localhost的用户就可以了。
    - IDENTIFIED BY 指定用户的登录密码,该项可以省略(某些版本下回报错，必须省略)。
    - WITH GRANT OPTION 这个选项表示该用户可以将自己拥有的权限授权给别人。注意：经常有人在创建操作用户的时候不指定WITH GRANT OPTION选项导致后来该用户不能使用GRANT命令创建用户或者给其它用户授权。
    备注：可以使用GRANT重复给用户添加权限，权限叠加，比如你先给用户添加一个select权限，然后又给用户添加一个insert权限，那么该用户就同时拥有了select和insert权限。
- 授权原则说明：
    - 只授予能满足需要的最小权限，防止用户干坏事。比如用户只是需要查询，那就只给select权限就可以了，不要给用户赋予update、insert或者delete权限。
    - 创建用户的时候限制用户的登录主机，一般是限制成指定IP或者内网IP段。
    - 初始化数据库的时候删除没有密码的用户。安装完数据库的时候会自动创建一些用户，这些用户默认没有密码。
    - 为每个用户设置满足密码复杂度的密码。
    - 定期清理不需要的用户。回收权限或者删除用户。

收回用户权限
```
# 语法格式
REVOKE privileges ON database.table FROM 'username'@'host';
# eg. 收回 local_user 的写入和更新权限
REVOKE insert,update ON *.* FROM 'local_user'@'%';
```
删除用户
```
# 语法格式
DROP USER 'username'@'host';
# eg. 删除本地用户 local_user
DROP USER 'local_user'@'localhost';
```