
[参考](https://www.jaegertracing.io/docs/1.14/deployment/)
[参考2](https://my.oschina.net/u/2548090/blog/1821372)


```
如果你的数据量特别多，使用kafka缓冲一下也是可以的(组件jaeger-ingester）这里不介绍

-e ES_USERNAME=elastic
注意，ES_USERNAME、ES_PASSWORD这两个环境变量，当你的elasticsearch未设置账号密码时，你可以不填，也可以填上默认值，elasticsearch的默认ES_USERNAME=elastic，ES_PASSWORD=changeme

安装收集器:
docker pull jaegertracing/jaeger-collector:1.14.0 #41MB

docker run -d --name jaeger-collector --restart=always --link es7.4:elasticsearch -e SPAN_STORAGE_TYPE=elasticsearch -e ES_SERVER_URLS=http://elasticsearch:9200  -p 14267:14267 -p 14268:14268 -p 9421:9411 jaegertracing/jaeger-collector:1.14.0

9411 是 zipkin的端口，所以改一下

安装界面:
docker pull jaegertracing/jaeger-query:1.14.0  #44MB

docker run -d --name jaeger-query --restart=always --link es7.4:elasticsearch -e SPAN_STORAGE_TYPE=elasticsearch -e ES_SERVER_URLS=http://elasticsearch:9200 -p 16686:16686/tcp jaegertracing/jaeger-query:1.14.0


安装agent:[参考](https://www.jaegertracing.io/docs/1.14/deployment/#discovery-system-integration)
docker pull jaegertracing/jaeger-agent:1.14.0 #27MB

docker run  -d  --name jaeger-agent --restart=always  --link jaeger-collector:jaegercollector -p 5775:5775/udp   -p 6831:6831/udp   -p 6832:6832/udp   -p 5778:5778/tcp   jaegertracing/jaeger-agent:1.14.0   --reporter.grpc.host-port=jaegercollector:14250

安装服务依赖:[参考](https://hub.docker.com/r/jaegertracing/spark-dependencies)
docker pull jaegertracing/spark-dependencies   #210MB

docker run --rm -d --name  spark-dependencies --link es7.4:elasticsearch   --env STORAGE=elasticsearch --env ES_NODES=http://elasticsearch:9200 jaegertracing/spark-dependencies

--rm  跑完删除容器

这里要做一个定时任务
```

端口说明:
elasticsearch暴露如下端口

端口号 |	协议 |	功能
-|-|-
9200 | HTTP |	通过http协议连接es使用的端口
9300 | TCP |	通过tcp协议连接es使用的端口
 

agent 暴露如下端口

端口号 | 协议 |	功能
-|-|-
5775 | UDP |	通过兼容性 thrift 协议，接收 zipkin thrift 类型的数据
6831 | UDP |	通过二进制 thrift 协议，接收 jaeger thrift 类型的数据
6832 | UDP |	通过二进制 thrift 协议，接收 jaeger thrift 类型的数据
5778 | HTTP |	可用于配置采样策略
collector 暴露如下端口

端口号	| 协议	| 功能
-|-|-
14267 |	TChannel |	用于接收 jaeger-agent 发送来的 jaeger.thrift 格式的 span
14268 |	HTTP |	能直接接收来自客户端的 jaeger.thrift 格式的 span
9411 |	HTTP |	能通过 JSON 或 Thrift 接收 Zipkin spans，默认关闭
query 暴露如下端口

端口号|	协议|	功能
-|-|-
16686| HTTP |	1. /api/* - API 端口路径 2. / - Jaeger UI 路径