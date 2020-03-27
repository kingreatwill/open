<!-- toc -->
[TOC]
# kong

kong 2.0 开始支持go开发插件。
https://docs.konghq.com/2.0.x/go/

## 本地

kong migrations bootstrap -c /mnt/e/kong/centos/kong.conf

kong start -c /mnt/e/kong/centos/kong.conf
kong stop

kong reload
```
# database = "postgres"
admin_listen = 0.0.0.0:8001
log_level = debug # 相当于 $ export KONG_LOG_LEVEL=error
pg_password = postgres
pg_host = 192.168.110.231
pg_port = 5432
declarative_config = /mnt/e/kong/centos/kong.yml
```
http://localhost:8001/

### konga
1. 初始化数据库
docker run --rm pantsel/konga:latest -c prepare -a postgres -u postgresql://postgres:postgres@192.168.110.231:5432/konga

2. 运行konga
```
docker run -p 1337:1337  -e "TOKEN_SECRET=123456" -e "DB_ADAPTER=postgres"  -e "DB_URI=postgresql://postgres:postgres@192.168.110.231:5432/konga"  -e "NODE_ENV=production"   --name konga       pantsel/konga
```
http://192.168.1.231:1337/

## k8s

1. 创建命名空间
```

git clone https://github.com/Kong/kubernetes-ingress-controller
cd kubernetes-ingress-controller/deploy/manifests

kubectl apply -f base/namespace.yaml
```
2. 安装数据库

```
kubectl apply -f postgres/postgres.yaml
```
3. 初始化数据库
```
kubectl apply -f postgres/migration.yaml

kubectl get job -n kong
```

4. rbac
```
kubectl apply -f base/rbac.yaml
```

5. kong Deployment
kubectl apply -f postgres/kong-ingress-postgres.yaml

6. kong Service
kubectl apply -f base/service.yaml






### gRPC 
https://docs.konghq.com/1.4.x/getting-started/configuring-a-grpc-service/

1. services 
curl -XPOST 192.168.1.143:8001/services \
  --data name=grpc \
  --data protocol=grpc \
  --data host=192.168.1.143 \
  --data port=5001

2. route
curl -XPOST 192.168.1.143:8001/services/grpc/routes \
  --data protocols=grpc \
  --data name=product-ProductBasicService \
  --data paths=/product.ProductBasicService
### Plugin
https://docs.konghq.com/1.4.x/getting-started/enabling-plugins/

### konga