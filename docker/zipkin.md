[参考](https://github.com/openzipkin/docker-zipkin/tree/master/elasticsearch7)
[镜像](https://hub.docker.com/r/openzipkin/zipkin)

```
docker pull openzipkin/zipkin:2.17.2
docker run -d -p 9411:9411 --name zipkin --link es7.4:elasticsearch   --env STORAGE_TYPE=elasticsearch --env ES_HOSTS=http://elasticsearch:9200 --restart always openzipkin/zipkin:2.17.2


UI:http://127.0.0.1:9411/zipkin/

其他环境变量
 - '--ES_USERNAME=elastic'
- '--ES_PASSWORD=123456'
- '--KAFKA_BOOTSTRAP_SERVERS=172.16.2.17:9092'


安装服务依赖:[参考](https://hub.docker.com/r/openzipkin/zipkin-dependencies)
docker pull openzipkin/zipkin-dependencies:2.3.2   #143.46 MB

docker run --rm -d --name  zipkin-dependencies --link es7.4:elasticsearch   --env STORAGE_TYPE=elasticsearch --env ES_HOSTS=http://elasticsearch:9200 openzipkin/zipkin-dependencies:2.3.2

其他环境变量
- JAVA_OPTS=-verbose:gc -Xms1G -Xmx1G

--rm  跑完删除容器

这里要做一个定时任务

```