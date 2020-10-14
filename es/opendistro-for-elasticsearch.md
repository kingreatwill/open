# elasticsearch sql plugin
https://github.com/opendistro-for-elasticsearch/sql
## opendistro-for-elasticsearch docker
https://opendistro.github.io/for-elasticsearch-docs/docs/install/docker/

`docker run -p 9200:9200 -p 9300:9300 --name elasticsearch  -e "discovery.type=single-node"  -e "cluster.name=elasticsearch"    -d  amazon/opendistro-for-elasticsearch:1.10.1`

测试`curl -XGET https://localhost:9200 -u admin:admin --insecure`


单独安装kibana
`docker run -p 5601:5601 --name elasticsearch-kibana --network=container:elasticsearch -e "ELASTICSEARCH_URL=https://ip:9200" -e ELASTICSEARCH_USERNAME="admin"  -e ELASTICSEARCH_PASSWORD="admin" -d  amazon/opendistro-for-elasticsearch-kibana:1.10.1`



## opendistro-for-elasticsearch + kibana docker
`docker run -p 9200:9200 -p 9300:9300 -p 5601:5601 --name elasticsearch  -e "discovery.type=single-node"  -e "cluster.name=elasticsearch"    -d  amazon/opendistro-for-elasticsearch:1.10.1`

`docker run  --name elasticsearch-kibana --network=container:elasticsearch -e "ELASTICSEARCH_URL=https://localhost:9200" -d  amazon/opendistro-for-elasticsearch-kibana:1.10.1`

## docker-compose
https://amazonaws-china.com/cn/blogs/china/running-open-distro-for-elasticsearch/

## odfe-sql-cli
`pip3 install odfe-sql-cli`
`odfesql https://localhost:9200 --username admin —password admin`