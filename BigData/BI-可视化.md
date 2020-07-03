据可视化是Business Intelligence(简称BI)中的核心功能

老牌的Tableau, Qilk，新生代的Looker，国内的FineBI,微软的Power BI等等

[数据可视化的开源方案: Superset vs Redash vs Metabase (二)](https://zhuanlan.zhihu.com/p/33164124)
[数据可视化的开源方案: Superset vs Redash vs Metabase (一)](https://zhuanlan.zhihu.com/p/33164027)
## Superset
Superset最初是由Airbnb的数据团队开源的，目前已进入Apache Incubator(孵化器)，算是明星级的开源项目。
https://github.com/apache/incubator-superset  python 29.2k
https://superset.incubator.apache.org/

支持数据库：https://superset.incubator.apache.org/#databases
支持Spark SQL

非官方
docker pull amancevice/superset:0.36.0
docker run --detach --name superset -p 8088:8088 amancevice/superset:0.36.0
docker exec -it superset superset-init

## Redash
https://github.com/getredash/redash python 16.7k
https://redash.io/

https://redash.io/help/data-sources/querying/supported-data-sources

Redash 比 Superset 支持的数据库多

## Metabase
https://github.com/metabase/metabase  Clojure 21.3k
https://www.metabase.com/

## stimulsoft Dashboards
商业，付费

## 实时数据看板

### PubNub+Cloudinary：简单几步快速搭建实时图片分享应用
PubNub是一个非常强大且易于使用的云服务，可进行实时的推送信息应用。而Cloudinary是一个基于云端的图片处理、管理、存储、美化于一体的平台，同时还集成了图片抓取功能，提供API接口。
https://github.com/cloudinary/cloudinary_pubnub_demo

### 实时数据可视化方案——PubNub+PowerBI

在“Power BI 服务”新建新仪表板，然后选择“添加磁贴” > “自定义流式处理数据” ，然后选择“下一步” 按钮。
从窗口右上角中的链接中选择“+ 添加流式处理数据” 。选择“PubNub” ，然后选择“下一步” 。