[参考](https://www.elastic.co/guide/en/kibana/current/docker.html)
[参考2](https://segmentfault.com/a/1190000020140461)
安装elasticsearch:7.4
```
docker pull elasticsearch:7.4.0 #859MB
docker pull kibana:7.4.0     #1.1GB

docker run -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --restart always -v /d/dockerv/elasticsearch/data/:/usr/share/elasticsearch/data/ --name es7.4 elasticsearch:7.4.0


如果你是在服务器上安装，想要对外访问还必须打开你服务器的9200端口，然后将localhost换成你服务器的ip地址即可。

修改配置，解决跨域访问问题
首先进入到容器中，然后进入到指定目录修改elasticsearch.yml文件。

docker exec -it es7.4 /bin/bash
cd /usr/share/elasticsearch/config/
vi elasticsearch.yml
在elasticsearch.yml的文件末尾加上:

http.cors.enabled: true
http.cors.allow-origin: "*"
修改配置后重启容器即可。

docker restart es7.4


安装ik分词器
docker exec -it es7.4 /bin/bash
cd /usr/share/elasticsearch/plugins/
elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.4.0/elasticsearch-analysis-ik-7.4.0.zip

exit

docker restart es7.4


启动kibana
使用--link连接到elasticsearch容器，命令如下:

docker run --name kibana7.4 --link=es7.4:elasticsearch  -p 5601:5601 --restart always -d kibana:7.4.0

--link=es7.4:elasticsearch  es7.4  要链接的容器,elasticsearch在kibana中使用的host ，通过elasticsearch可以访问es7.4
```

安装elasticsearch:6.8.3
```
docker pull elasticsearch:6.8.3 #800MB
docker pull kibana:6.8.3        #665MB
```

其他
```
docker run -d -p 9200:9200 -p 9300:9300 --restart always -v /dockerv/elasticsearch/config/:/usr/share/elasticsearch/config/ -v /dockerv/elasticsearch/data/:/usr/share/elasticsearch/data/ elasticsearch
```