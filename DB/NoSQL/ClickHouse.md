# ClickHouse
[中文文档](https://clickhouse.com/docs/zh)

ClickHouse是一个用于联机分析(OLAP)的列式数据库管理系统(DBMS)。



## UI管理工具

### sqlpad
```sh
docker run --name sqlpad -p 3000:3000 -v /e/dockerv/sqlpad/data:/var/lib/sqlpad -e SQLPAD_ADMIN="enter@wcoder.com" -e SQLPAD_ADMIN_PASSWORD="123456" --detach sqlpad/sqlpad:latest

# --detach 分离的意思, 也就是后台运行 -d的意思
```
