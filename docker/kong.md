# docker安装kong with postgres
[官网](https://konghq.com/kong)
[教程](https://docs.konghq.com/install/docker/)
[docker hub](https://hub.docker.com/_/kong/)

# docker安装kong DB-less

## 安装kong
```
docker network create kong-net

docker run -d --name kong-database \
               --network=kong-net \
               -p 5432:5432 \
               -v $HOME/kong/postgres-data:/var/lib/postgresql/data \
               -e "POSTGRES_USER=kong" \
               -e "POSTGRES_DB=kong" \
               -e "POSTGRES_PASSWORD=kong" \
               postgres:9.6

# 初始化数据
docker run --rm \
     --network=kong-net \
     -e "KONG_DATABASE=postgres" \
     -e "KONG_PG_HOST=kong-database" \
     -e "KONG_PG_USER=kong" \
     -e "KONG_PG_PASSWORD=kong" \
     kong:latest kong migrations bootstrap


# 导入Kong配置
docker run --rm \
    --network=kong-net \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_USER=kong" \
    -e "KONG_PG_PASSWORD=kong" \
    -v $HOME/kong/config:/home/kong \
    kong:latest kong config db_import /home/kong/kong.yml

# 启动Kong
docker run -d --name kong \
     --network=kong-net \
     -e "KONG_DATABASE=postgres" \
     -e "KONG_PG_HOST=kong-database" \
     -e "KONG_PG_USER=kong" \
     -e "KONG_PG_PASSWORD=kong" \
     -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
     -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
     -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
     -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
     -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
     -p 8000:8000 \
     -p 8443:8443 \
     -p 127.0.0.1:8001:8001 \
     -p 127.0.0.1:8444:8444 \
     kong:latest


# 初始化数据konga数据库
docker run --rm \
             --network=kong-net \
             pantsel/konga:latest \
             -c prepare \
             -a "postgres" \
             -u "postgres://kong:kong@kong-database:5432/konga"

# 启动Konga
docker run -d --name konga \
             --network kong-net \
             -e "TOKEN_SECRET=secret123" \
             -e "DB_ADAPTER=postgres" \
             -e "DB_URI=postgres://kong:kong@kong-database:5432/konga" \
             -e "NODE_ENV=development" \
             -p 1337:1337 \
             pantsel/konga

```

这里创建数据库需要两个, 一个kong,一个konga
把shell/sql脚本放入/docker-entrypoint-initdb.d/目录中，让容器启动的时候自动执行创建；
shell脚本create-multiple-postgresql-databases.sh
```bash
#!/bin/bash

set -e
set -u

function create_user_and_database() {
    local database=$1
    echo "  Creating user and database '$database'"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
        CREATE USER $database;
        CREATE DATABASE $database;
        GRANT ALL PRIVILEGES ON DATABASE $database TO $database;
EOSQL
}

if [ -n "$POSTGRES_MULTIPLE_DATABASES" ]; then
    echo "Multiple database creation requested: $POSTGRES_MULTIPLE_DATABASES"
    for db in $(echo $POSTGRES_MULTIPLE_DATABASES | tr ',' ' '); do
        create_user_and_database $db
    done
    echo "Multiple databases created"
fi
```
或者sql脚本create-multiple-postgresql-databases.sql
```sql
CREATE USER pkslowuser;

CREATE DATABASE logdata;
GRANT ALL PRIVILEGES ON DATABASE logdata TO pkslowuser;

CREATE DATABASE orderdata;
GRANT ALL PRIVILEGES ON DATABASE orderdata TO pkslowuser;

CREATE DATABASE userdata;
GRANT ALL PRIVILEGES ON DATABASE userdata TO pkslowuser;
```

准备Dockerfile，把shell/sql脚本文件放入镜像中去：
```Dockerfile
FROM postgres:10
COPY create-multiple-postgresql-databases.sh /docker-entrypoint-initdb.d/
COPY create-multiple-postgresql-databases.sql /docker-entrypoint-initdb.d/
```
# 启动
```sh
docker run -itd \
    --name pkslow-postgres \
    -e POSTGRES_MULTIPLE_DATABASES=db1,db2 \
    -e POSTGRES_USER=pkslow \
    -e POSTGRES_PASSWORD=pkslow \
    -p 5432:5432 \
    pkslow/postgresql-multiple-databases:1.0
```
> https://github.com/LarryDpk/pkslow-samples


# 一个典型的 Nginx 配置
```
upstream helloUpstream {
    server localhost:3000 weight=100;
}

server {
    listen  80;
    location /hello {
        proxy_pass http://helloUpstream;
    }
}
```
如上这个简单的 Nginx 配置，便可以转换为如下的 Http 请求。

# 对应的 Kong 配置

### 配置 upstream
```
curl -X POST http://localhost:8001/upstreams --data "name=helloUpstream"
```
### 配置 service
```
curl -X POST http://localhost:8001/services --data "name=hello" --data "host=helloUpstream"
```
### 配置 target
```
curl -X POST http://localhost:8001/upstreams/hello/targets --data "target=localhost:3000" --data "weight=100"
```
### 配置 route
```
curl -X POST http://localhost:8001/routes --data "paths[]=/hello" --data "service.id=8695cc65-16c1-43b1-95a1-5d30d0a50409"
```
这一切都是动态的，无需手动 reload nginx.conf。

我们为 Kong 新增路由信息时涉及到了 upstream，target，service，route 等概念，他们便是 Kong 最最核心的四个对象。（你可能在其他 Kong 的文章中见到了 api 这个对象，在最新版本 0.13 中已经被弃用，api 已经由 service 和 route 替代）

从上面的配置以及他们的字面含义大概能够推测出他们的职责，upstream 是对上游服务器的抽象；target 代表了一个物理服务，是 ip + port 的抽象；service 是抽象层面的服务，他可以直接映射到一个物理服务(host 指向 ip + port)，也可以指向一个 upstream 来做到负载均衡；route 是路由的抽象，他负责将实际的 request 映射到 service。

他们的关系如下

upstream 和 target ：1 对 n

service 和 upstream ：1 对 1 或 1 对 0 （service 也可以直接指向具体的 target，相当于不做负载均衡）

service 和 route：1 对 n

# 高可扩展性的背后—插件机制
Kong 的另一大特色便是其插件机制，这也是我认为的 Kong 最优雅的一个设计。

文章开始时我们便提到一点，微服务架构中，网关应当承担所有服务共同需要的那部分功能，这一节我们便来介绍下，Kong 如何添加 jwt 插件，限流插件。

插件（Plugins）装在哪儿？对于部分插件，可能是全局的，影响范围是整个 Kong 服务；大多数插件都是装在 service 或者 route 之上。这使得插件的影响范围非常灵活，我们可能只需要对核心接口进行限流控制，只需要对部分接口进行权限控制，这时候，对特定的 service 和 route 进行定向的配置即可。

为 hello 服务添加50次/秒的限流
```
curl -X POST http://localhost:8001/services/hello/plugins \
--data "name=rate-limiting" \
--data "config.second=50"
```
为 hello 服务添加 jwt 插件
```
curl -X POST http://localhost:8001/services/login/plugins \
--data "name=jwt"
```
同理，插件也可以安装在 route 之上
```
curl -X POST http://localhost:8001/routes/{routeId}/plugins \
--data "name=rate-limiting" \
--data "config.second=50"

curl -X POST http://localhost:8001/routes/{routeId}/plugins \
--data "name=jwt"
```