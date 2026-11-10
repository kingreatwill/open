# 技术

## 开源地图
[OpenMapTiles](https://github.com/openmaptiles)
[OpenLayers](https://github.com/openlayers/openlayers)
[Mapbox](https://docs.mapbox.com/#maps)
[Leaflet](https://github.com/Leaflet/Leaflet)
```
var map = L.map('map').setView([51.505, -0.09], 13);

L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

L.marker([51.5, -0.09]).addTo(map)
    .bindPopup('A pretty CSS3 popup.<br> Easily customizable.')
    .openPopup();
```


数据-https://data.maptiler.com/downloads/planet/
[OpenStreetMap](https://www.openstreetmap.org/)
[OpenStreetMap Data Extracts](https://download.geofabrik.de/)

瓦片/tile [下载](http://wmksj.com/)
为何叫瓦片，一般是服务端根据数据生成一张大图，然后切片成一张张小图，一般大小是256x256的图片

### GIS 3D
#### CesiumJS
https://github.com/CesiumGS/cesium
https://cesium.com/



[God's Eye View](https://github.com/bilawalsidhu/gods-eye-view) 是一个完全运行在浏览器里的实时地理空间情报控制台：基于 CesiumJS 与 Google 真实感 3D 地球，叠加全球航班、船舶、卫星、地震、交通与公共摄像头等 13 个公开数据源，还支持用 OpenAI Realtime API 语音指挥。

### 图表工具

#### Highcharts
https://www.highcharts.com/
#### 数据地图 Power Map
https://support.office.com/en-us/article/get-started-with-power-map-88a28df6-8258-40aa-b5cc-577873fb0f4a
#### Echarts
https://echarts.apache.org/zh/index.html
#### AntV
https://antv.vision/zh

### 数据可视化工具
[30个值得推荐的数据可视化工具（2020年更新）](https://zhuanlan.zhihu.com/p/51695598)

图表
- https://rawgraphs.io/
- https://www.chartblocks.com/
- https://www.tableau.com/
- https://powerbi.microsoft.com/zh-cn/
- https://www.qlik.com/us/products/qlikview
- https://www.datawrapper.de/
- https://www.visme.co/
- https://www.grow.com/
- https://www.icharts.in/

信息图
- https://infogram.com/
- https://visual.ly/

地图
- https://www.instantatlas.com/
- https://leafletjs.com/
- https://openlayers.org/
- https://kartograph.org/
- https://carto.com/

关系网络图
- https://gephi.org/
- http://sigmajs.org/

数学图形
- https://www.wolframalpha.com/

开发者工具
- https://www.echartsjs.com/zh/index.html
- https://d3js.org/
- https://plotly.com/
- https://www.chartjs.org/
- https://developers.google.com/chart
- https://opensource.addepar.com/ember-charts/
- https://gionkunz.github.io/chartist-js/
- https://www.highcharts.com/
- https://www.fusioncharts.com/
- https://www.zingchart.com/

金融图表
- http://dygraphs.com/  https://github.com/danvk/dygraphs
- https://cn.tradingview.com/  https://github.com/tradingview/lightweight-charts
- WijmoJS



## 地理体系
[bing map 地图体系](https://msdn.microsoft.com/en-us/library/bb259689.aspx)<br>
[地理概念初步](http://www.cnblogs.com/beniao/archive/2010/04/18/1714544.html)<br>
[各种坐标系转换](http://www.thinkgis.cn/topic/560370f200b823b7114ea635)<br>

## webGis系列 
[1 何为栅格数据，何为矢量数据](http://www.thinkgis.cn/topic/541a981ada8db186fd209c01)<br>
[6 WebGIS中地图瓦片在Canvas上的拼接显示原理](http://www.thinkgis.cn/topic/541a90d2da8db186fd1de575)<br>
[11 WebGIS中要素（Feature）的设计](http://www.thinkgis.cn/topic/541a9cdeda8db186fd226303)<br>

[搭建开源地图服务1](http://wangwei.info/osmgis-planet-data-import/)
[搭建开源地图服务2](http://wangwei.info/how-install-tilemill-in-osmgis/)
[搭建开源地图服务3](http://wangwei.info/osmgis-openlayers-test/)


## 瓦片系列
[矢量瓦片规范 CN](https://github.com/jingsam/vector-tile-spec/blob/master/2.1/README_zh.md)
## postGis 系列
[postgis简介 及性能对比](http://mysql.taobao.org/monthly/2015/07/04/) <br>

[用 R 處理 PostgreSQL/PostGIS ](http://mutolisp.logdown.com/posts/206944-treatment-with-r-postgresql-postgis-data)<br>

[openGis的一些软件](https://groups.google.com/forum/?hl=es-419#!topic/gisforums/LhEeCYAnLoo)

[postGis黑科技（浅）](https://yq.aliyun.com/articles/2727)

[PostgreSQL+PostGIS的使用](http://blog.chenapp.com/archives/332)

[postGis 常用函数 CN](http://xml.iteye.com/blog/1525730)

[node 和 postgis如何合用 EN](http://blog.geomusings.com/2014/02/18/a-little-deeper-with-node-and-postgis/)

[PostGIS 快速入门 CN](http://live.osgeo.org/zh/quickstart/postgis_quickstart.html)

[PostGIS 与 Sequelize.js搭配 EN](https://manuel-rauber.com/2016/01/08/using-geo-based-data-with-sequelizejs-utilizing-postgresql-and-ms-sql-server-in-node-js/)

[PostgreSQL 9.4.4 中文手册 CN 推荐](http://www.postgres.cn/docs/9.4/index.html)

[pg的窗口函数等用法](http://blog.csdn.net/ai6740165/article/details/38038259)

[geometry 和 geography的一些区别](https://github.com/Universefei/memo/blob/master/gis/PostGIS-04.md)





#案例

## 综合
[创新型可视化集合](http://conceptviz.github.io/#/e30=)

## 艺术作品
[圣诞节webgl表现](http://christmasexperiments.com/)<br>


## 科学类
[windyty](https://www.windyty.com/)<br>
流场模拟

[nullshcool](http://earth.nullschool.net/)
流场模拟, [github](https://github.com/cambecc/earth)

## 广度优先搜索
[条条大路通罗马](http://roadstorome.moovellab.com/explore)
[一天里你可以坐火车去多少地方](https://www.washingtonpost.com/news/worldviews/wp/2015/06/05/map-the-remarkable-distances-you-can-travel-on-a-european-train-in-less-than-a-day/)<br>

## 创新
[地图与图的互动](http://globe.cid.harvard.edu/?mode=productsphere&id=CU)
[地铁交通可视化](http://mbtaviz.github.io/)<br>


## 作品

[strava实验室](http://labs.strava.com/projects/)







# 地理相关数据

## 数据1

[OpenStreetMap（OSM）](http://wiki.openstreetmap.org/wiki/Downloading_data)<br>
世界上最大的开源地理数据平台， 本页提供整体下载、大规模下载和细节下载三种模式。

[MapZen](https://mapzen.com/data/)<br>
MapZen提供城市尺度的OSM数据打包下载(提供多种格式)，国家等行政区域下载，

[Natural Earth Data](http://www.naturalearthdata.com/downloads/)<br>
Natural Earth Data提供了全球范围内的矢量和影像数据, 矢量数据已经包含了不同尺度的抽吸。Natural Earth Data的最大优势就是数据是开放性的，用户有传播和修改数据的权限

[OpenTopography](http://www.opentopography.org/)<br>
Open Topography是一个提供高空间分辨率的地形数据和操作工具的门户网站。通过Open Topography，用户可以下载LiDAR数据（主要包括：美国、加拿大、澳大利亚、巴西、海地、墨西哥和波多黎各）。

[GeoNames](http://www.geonames.org/)
世界各地的地名集合(800w, 2016.01)

#### 地方性数据

[data hub 洛杉矶](http://geohub.lacity.org/)
大量细节性数据 如 数百年来的洪水记录

## 数据2

[NASA Earth Observations (NEO)](http://neo.sci.gsfc.nasa.gov)<br>
NEO专注于提供全球范围内的卫星影像（大气、能源、土地、生活、海洋等50多种不同数据专题）。通过NEO可以查看地球气候和环境状况的每日快照。

[USGS](http://earthexplorer.usgs.gov/)<br>
美国地质勘探局（United States Geological Survey，简称USGS），是美国内政部所属的科学研究机构。其官网上提供最新、最全面的全球卫星影像，包括Landsat、Modis等。

[SEDAC](http://sedac.ciesin.columbia.edu/data/sets/browse)<br>
NASA旗下的SEDAC提供全球范围内的GIS数据以帮助人们了解人与环境间的相互影响。数据涉及农业、气候、健康、基础设施、土地利用、海洋和沿海、人口、贫困、可持续性、城市和水等15种类型。

[CGI](http://www.cgiar-csi.org/data)<br>
气候、温度、土壤等遥感数据

[OpenTreeMap](https://www.opentreemap.org)<br>
植物数据

[UNEP Environmental Data Explorer](http://geodata.grid.unep.ch)<br>
UNEP包含全球范围内500多种不同类型的空间和非空间数据，如淡水、人口、森林、污染排放、气候、灾害、卫生和国内生产总值等。

[FAO GeoNetwork](http://www.fao.org/geonetwork/srv/en/main.home)<br>
FAO是一个全球地理信息系统数据集，通过它你可以下载到农业、渔业、土地资源相关的GIS数据，同时它提供相关卫星图像数据。

[ISCGM Global Map](http://www.iscgm.org/gm/index.html)<br>
ISCGM提供的数据种类包括全球土地和森林覆盖数据集。同时一些文化和自然矢量数（边界、排水、交通、人口中心、海拔、土地覆盖、土地利用和植被）也能在这里获取。

[poi86 中国poi点下载](http://www.poi86.com/)<br>

[中国历史地图](http://www.fas.harvard.edu/~chgis/data/chgis/downloads/v4/)<br>

## 地理接口

高德-geocoding:<br>
根据地址解析出经纬度等地理信息,其中key为用户申请的key，address为地址字符信息

高德-搜索接口: <br>
根据关键词模糊匹配地理信息,city字段为城市的adcode，keywords为关键词，次数多了会被封 需要手动解封。
http://ditu.amap.com/service/poiInfo?query_type=TQUERY&city=xxx&keywords=yyy&pagesize=100&pagenum=1&qii=true&cluster_state=5&need_utd=true&utd_sceneid=1000&div=PC1000&addr_poi_merge=true&is_classify=true

## 轨迹类数据

[纽约市2009年到2015年的出租车数据](http://www.nyc.gov/html/tlc/html/about/trip_record_data.shtml)

[UBER 2014-2015年的专车数据](https://github.com/fivethirtyeight/uber-tlc-foil-response)

[加州硅谷湾区的共享自行车系统 2013-2015年的运营数据](http://www.bayareabikeshare.com/open-data)

[微软研究院的出租车数据] (http://research.microsoft.com/apps/pubs/?id=152883)



## 3S
早期"3S"是指遥感(Remote Sensing)、全球定位系统GPS (Global Position System) 和地理信息系统(Geographic Information System) 的简称，
广义的说法则是遥感(Remote Sensing)、地理信息系统(Geographic Information System) 和全球导航卫星系统（ Global Navigation Satellite System），其中GNSS泛指所有卫星定位系统，包括GPS。
“3s”是空间技术、传感器技术、卫星定位与导航技术和计算机技术、通信技术相结合，多学科高度集成的对空间信息进行采集、处理、管理、分析、表达、传播和应用的现代信息技术的总称。

### GIS 地理信息系统(Geographic Information System) 

arcpy开发python应用 

#### OSRM (Open Source Routing Machine) 
https://github.com/Project-OSRM

### GPS 全球定位系统 (Global Position System) 
火星坐标
### RS  遥感(Remote Sensing)

### MapServer
GitHub：https://github.com/mapserver/mapserver
文档：https://mapserver.org/documentation.html
MapServer 是一个强大的开源 GIS 服务软件，适用于 web 地图发布、空间数据查询和格式转换。

### Leaflet
GitHub：https://github.com/Leaflet/Leaflet
文档：https://leafletjs.com/reference.html
Leaflet 是一个轻量级且功能强大的地图库，适用于开发各种交互式地图应用。

### QGIS
GitHub：https://github.com/qgis/QGIS
文档：https://docs.qgis.org
QGIS 是一个功能强大的开源 GIS 工具，适用于地理数据分析、地图制作以及空间数据处理等多种场景。

### GeoServer
GitHub：https://github.com/geoserver/geoserver
文档：http://docs.geoserver.org
GeoServer 是一个强大的开源 GIS 服务器，提供地图和空间数据服务，适用于各种地理信息共享和数据发布的场景。

### GraphHopper
GitHub：https://github.com/graphhopper/graphhopper
API 文档：https://docs.graphhopper.com
💡 GraphHopper 是一个高效的路线规划引擎，适用于多种导航应用，特别适合大规模路径计算和优化任务。

### Navit
GitHub：https://github.com/navit-gps/navit
文档：https://wiki.navit-project.org
Navit 是一个高效的开源离线导航系统，支持自定义 UI 和插件扩展，适用于各种设备和导航需求。

### Bing Maps
API 文档：https://docs.microsoft.com/en-us/bingmaps/
开发者门户：https://www.microsoft.com/en-us/maps
💡 Bing Maps 是微软提供的强大地图服务，适用于企业级 GIS、物流优化和智能导航应用。

## 免费的GIS软件
[14个免费的 GIS 软件：以开源的方式绘制地图](https://www.osgeo.cn/post/1b456)
### QGIS
> QGIS 3
https://github.com/qgis/QGIS

> QGIS 2
https://gisgeography.com/open-source-qgis-review-guide/
### GRASS GIS
https://grass.osgeo.org/

### Whitebox GAT
https://www.whiteboxgeo.com/

### gVSIG
https://gisgeography.com/gvsig-software/

### ILWIS 
ILWIS(Integrated Land and Water Information System-陆地水体信息集成系统) 
https://52north.org/news/ilwis-in-2016/

### SAGA GIS
https://saga-gis.sourceforge.io/en/index.html

### GeoDa
https://geodacenter.github.io/download.html

### MapWindow
https://www.mapwindow.org/

### uDig
http://udig.refractions.net/

### OpenJump GIS
http://www.openjump.org/
### FalconView
### OrbisGIS
http://orbisgis.org/
### Diva GIS