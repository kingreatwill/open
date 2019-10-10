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

docker run  -itd -p 3406:3306 -v D:/dockerv/mysql8/conf:/etc/mysql/conf.d -v D:/dockerv/mysql8/logs:/logs -v D:/dockerv/mysql8/data:/var/lib/mysql --restart always -e MYSQL_ROOT_PASSWORD=123456@lcb --name mysql8.0.17  mysql:8.0.17 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci



进入MySQL容器,登陆MySQL
docker exec -it mysql8.0.17 /bin/bash

mysql -u root -p

 mysql8.0的root用户的验证方式变了，通过查询：
  use mysql;
 select host,user,plugin,authentication_string from mysql.user;

得知：root的用户的加密方式为caching_sha2_passoword, 而navicat连接所用的方式为native_password。mysql为远程连接和本地连接提供了不同的密码验证方式。

备注：host为 % 表示不限制ip   localhost表示本机使用    plugin非mysql_native_password 则需要修改密码

ALTER user 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456@lcb';

FLUSH PRIVILEGES;  

然后就可以远程登陆MySQL
```