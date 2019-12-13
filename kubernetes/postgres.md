https://www.postgresql.org/about/news/1957/
官网

## 安装postgres

https://github.com/Kong/kubernetes-ingress-controller/blob/master/deploy/manifests/postgres/postgres.yaml


## 安装postgres-operator
```
git clone https://github.com/zalando/postgres-operator.git
cd postgres-operator

# 处理好postgres-operator.yaml中的镜像;
docker pull dongzhangqi/postgres-operator:v1.2.0

docker tag dongzhangqi/postgres-operator:v1.2.0 registry.opensource.zalan.do/acid/postgres-operator:v1.2.0

docker rmi dongzhangqi/postgres-operator:v1.2.0

# 或者修改postgres-operator.yaml中的镜像（推荐，因为你不知道pod在哪个k8s集群的节点上，所以每个node都要tag镜像）
dongzhangqi/postgres-operator:v1.2.0


# apply the manifests in the following order
kubectl create -f manifests/configmap.yaml  # configuration
kubectl create -f manifests/operator-service-account-rbac.yaml  # identity and permissions
kubectl create -f manifests/postgres-operator.yaml  # deployment
```
## Create a Postgres cluster
```
# if you've created the operator using yaml manifests
kubectl get pod -l name=postgres-operator

# if you've created the operator using helm chart
kubectl get pod -l app.kubernetes.io/name=postgres-operator

# create a Postgres cluster
kubectl create -f manifests/minimal-postgres-manifest.yaml
```
```
# check the deployed cluster
kubectl get postgresql

# check created database pods
kubectl get pods -l application=spilo -L spilo-role

# check created service resources
kubectl get svc -l application=spilo -L spilo-role
```

## Connect to the Postgres cluster via psql

## Delete a Postgres cluster
```
kubectl delete postgresql acid-minimal-cluster
```


# 本地安装
```
--postgresql安装
yum -y install https://download.postgresql.org/pub/repos/yum/11/redhat/rhel-7-x86_64/pgdg-centos11-11-2.noarch.rpm 安装存储库
yum -y install postgresql11 安装客户端
yum -y install postgresql11-server 安装服务端
rpm -aq| grep postgres验证安装情况
/usr/pgsql-11/bin/postgresql-11-setup initdb 初始化数据库
systemctl enable postgresql-11 开机启动
systemctl start postgresql-11 启动PG
sudo su - postgres 切换用户
psql -d postgres 登录数据库
#或者 psql -U postgres 登录数据库
ALTER USER postgres WITH PASSWORD 'postgres';修改密码
CREATE DATABASE kong; 创建数据库
CREATE USER kong CREATEDB LOGIN PASSWORD 'postgres';创建登录用户
GRANT ALL ON DATABASE kong TO kong;分配权限
\q 退出
systemctl restart postgresql-11 重启

--postgresql配置
开启远程访问
vim /var/lib/pgsql/11/data/postgresql.conf
修改#listen_addresses = 'localhost' 为 listen_addresses='*’
此处‘*’也可以改为任何你想开放的服务器IP
信任远程连接
vim /var/lib/pgsql/11/data/pg_hba.conf
修改如下内容，信任指定服务器连接 
# IPv4 local connections: 
host all all 0.0.0.0/0 trust

systemctl restart postgresql-11 重启


```

连接 ip

用户 postgres 密码 postgres
用户 kong 密码 postgres