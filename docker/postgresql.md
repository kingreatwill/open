[安装](https://hub.docker.com/_/postgres)
```
docker pull postgres:12.0
docker pull postgres:13.0-alpine

docker volume create --name=pgdata

docker run -d -p 5432:5432 -v pgdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123456@lcb --name postgresql --restart always postgres:12.0

docker run -d -p 5432:5432 -v /d/dockerv/postgresql13/data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123456@lcb --name postgresql --restart always postgres:13


PGDATA=/var/lib/postgresql/data

默认数据库:postgres

POSTGRES_USER: root
POSTGRES_DB: database
POSTGRES_PASSWORD: 123456

默认用户名:postgres 

docker exec -it postgresql /bin/bash

切换用户
su postgres
```


## 本地安装


本地访问#
先登录 shell （以 postgres 用户为例）：
```
sudo -i -u postgres

or

sudo su - postgres
```
然后，输入：
```
# 以当前登录的 linux 用户名为数据库名 (这点跟 mysql 不一样，必须得先指定登录的数据库)
psql 

# 手动指定数据库名
psql -d postgres
```



开启远程访问
vi /var/lib/pgsql/12/data/pg_hba.conf

若不知道 PostgreSQL 的配置文件路径，可在 PostgreSQL CLI 里输入：SHOW config_file 

然后把文件拉到底，最后一段里，将“ident”替换为“md5”，最终是这样的：
```
# TYPE  DATABASE        USER            ADDRESS                 METHOD

# "local" is for Unix domain socket connections only
local   all             all                                     trust
# IPv4 local connections:
host    all             all             127.0.0.1/32            md5
# IPv6 local connections:
host    all             all             ::1/128                 md5
# Allow replication connections from localhost, by a user with the
# replication privilege.
local   replication     postgres                                peer
host    replication     postgres        127.0.0.1/32            md5
host    replication     postgres        ::1/128                 md5
```