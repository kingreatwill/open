# docker 安装sqlserver 2019
[中文教程](https://docs.microsoft.com/zh-cn/sql/linux/quickstart-install-connect-docker?view=sqlallproducts-allversions&pivots=cs1-bash)

[英文的可能更新的快](https://docs.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker?view=sql-server-ver15&pivots=cs1-bash)

[版本列表](https://mcr.microsoft.com/v2/mssql/server/tags/list)
[hub](https://hub.docker.com/_/microsoft-mssql-server)

```
docker pull mcr.microsoft.com/mssql/rhel/server:2019-CTP3.2 #1.77G 有点大，比较慢

docker pull mcr.microsoft.com/mssql/server:2019-GA-ubuntu-16.04 # 2019-11-02 正式版

docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=123456@lcb" -v /d/dockerv/mssql/data/:/var/opt/mssql/data/ -p 1433:1433 --name sql2019 --restart always -itd mcr.microsoft.com/mssql/rhel/server:2019-CTP3.2

然后就可以了


更改密码：
sudo docker exec -it sql2019 /opt/mssql-tools/bin/sqlcmd \
   -S localhost -U SA -P "<YourStrong@Passw0rd>" \
   -Q 'ALTER LOGIN SA WITH PASSWORD="<YourNewStrong@Passw0rd>"'


连接到 SQL Server
下列步骤在容器内部使用 SQL Server 命令行工具 sqlcmd 来连接 SQL Server。
使用 docker exec -it 命令在运行的容器内部启动交互式 Bash Shell。 

sudo docker exec -it sql2019 "bash"

在容器内部使用 sqlcmd 进行本地连接。 默认情况下，sqlcmd 不在路径之中，因此需要指定完整路径。

/opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P "<YourNewStrong@Passw0rd>"


创建和查询数据：

CREATE DATABASE TestDB

SELECT Name from sys.Databases

前两个命令没有立即执行。 必须在新行中键入 GO 才能执行以前的命令：

GO

```


安装2017
```
docker run -e 'ACCEPT_EULA=Y' -e 'MSSQL_SA_PASSWORD=123456@lcb' --privileged=true -v /dockerv/mssql/data/:/var/opt/mssql/data/ -p 1433:1433 --name mssql --restart always -d microsoft/mssql-server-linux
```