# 安装prometheus
```
docker pull prom/prometheus:v2.13.1
```
# [安装grafana](https://grafana.com/docs/installation/docker/)
```
docker pull grafana/grafana:6.4.3

docker run -d --name=grafana -p 3000:3000 grafana/grafana:6.4.3

---------------------


docker run -d --name=grafana -p 3000:3000 grafana/grafana

持久化存储
# create a persistent volume for your data in /var/lib/grafana (database and plugins)

docker volume create grafana-storage

# start grafana
docker run \
  -d \
  -p 3000:3000 \
  --name=grafana \
  -v grafana-storage:/var/lib/grafana \
  grafana/grafana

或者
mkdir data # creates a folder for your data
ID=$(id -u) # saves your user id in the ID variable

# starts grafana with your user id and using the data folder
docker run -d --user $ID --volume "$PWD/data:/var/lib/grafana" -p 3000:3000 grafana/grafana:5.1.0


--user root
--user 472

Version	User	User ID
< 5.1	grafana	104
>= 5.1	grafana	472
```
插件安装
```
docker run \
  -d \
  -p 3000:3000 \
  --name=grafana \
  -e "GF_INSTALL_PLUGINS=grafana-clock-panel,grafana-simple-json-datasource" \
  grafana/grafana


从源码安装

  docker run \
  -d \
  -p 3000:3000 \
  --name=grafana \
  -e "GF_INSTALL_PLUGINS=http://plugin-domain.com/my-custom-plugin.zip;custom-plugin" \
  grafana/grafana
```
Setting	| Default value
-|-
GF_PATHS_CONFIG | 	/etc/grafana/grafana.ini
GF_PATHS_DATA	| /var/lib/grafana
GF_PATHS_HOME	| /usr/share/grafana
GF_PATHS_LOGS	| /var/log/grafana
GF_PATHS_PLUGINS	| /var/lib/grafana/plugins
GF_PATHS_PROVISIONING	| /etc/grafana/provisioning