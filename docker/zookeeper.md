[参考](https://hub.docker.com/_/zookeeper)
[config文件](https://github.com/apache/zookeeper/blob/master/conf/zoo_sample.cfg)

```
docker pull zookeeper:3.5.5



docker run -d -p 2181:2181 --name zookeeper --restart always -v /d/dockerv/zookeeper/conf:/conf -v /d/dockerv/zookeeper/data:/data -v /d/dockerv/zookeeper/datalog:/datalog zookeeper:3.5.5
```