[安装](https://hub.docker.com/_/postgres)
```
docker pull postgres:12.0

docker volume create --name=pgdata

docker run -d -p 5432:5432 -v pgdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123456@lcb --name postgresql --restart always postgres:12.0


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