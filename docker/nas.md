
## Private Tracker(PT)
PT站（Private Tracker）是一种私有的种子分享站点，和公共BT站点不同，只有在站内注册且满足一定门槛的用户才能相互分享和下载资源。因此，PT站点一般资源更加丰富，但门槛也更高。
### IYUU
1. 网页管理、辅种、转移、下载、定时访问URL、动态域名ddns等常用功能，提供完善的插件机制。

1. IYUUPlus客户端完全开源。

1. IYUU自动辅种工具，目前能对国内大部分的PT站点自动辅种，支持下载器集群，支持多盘位，支持多下载目录，支持连接远程下载器等。
> https://www.bilibili.com/read/cv13055302/
> https://github.com/ledccn/IYUUPlus
> https://github.com/ledccn/IYUUAutoReseed


docker : https://gitee.com/ledc/IYUUAutoReseed/tree/master/docker
把你的配置文件放在/root/config.php
```
docker run -d \
--name IYUUAutoReseed \
-e cron='0 9 * * 0' \
-v /root/config.php:/config.php \
--restart=always \
iyuucn/iyuuautoreseed:latest
```


### Transdroid
手机远程控制qbittorrent bt qb
使用手机管理pt下载，不用电脑了，方便，还可以订阅pt网站资源更新

Transdroid下载网站：https://f-droid.org/zh_Hans/packages/org.transdroid.lite/
Transdroid下载地址：https://f-droid.org/repo/org.transdroid.lite_241.apk
Transdroid远程管理NAS下qBittorrent和Transmission套件

https://github.com/erickok/transdroid
### PtManagePlugin(PT管理宝)
> 不开源, 很危险, 初始密码 admin admin

PT管理辅助程序是通过微信小程序搜索PT资源，并推送到下载器。

```
docker run \
--name=PtManagePlugin \
-v /share/docker/PtManagePlugin/db:/app/db \
-p 7001:7001 \
--restart=always \
zht39635371/pt-manage-plugin:latest
```

### 参考
https://zhuanlan.zhihu.com/p/394520776
1、qBittorrent安装教程
2、Transmission安装教程
3、玩物下载安装教程
4、Download Station设置教程




## OS IN docker

### Virtual DSM
Virtual DSM 是一个可以在 Docker 里安装黑群晖的项目，甚至可以在 DSM 中再安装一个 DSM，无限套娃。
Virtual DSM in a docker container.
https://github.com/kroese/virtual-dsm

```
docker run -it --rm -p 5000:5000 --device=/dev/kvm --cap-add NET_ADMIN --stop-timeout 60 kroese/virtual-dsm:latest
```



## app
mix文件管理

## 其它
### 电子书库 hectorqin/reader

https://www.toutiao.com/article/7257138465949614607/
```
version: "3" 
services:  
     reader:    
        image: hectorqin/reader:latest
        container_name: reader    
        environment:   
         - JAVA_ALPINE_VERSION=8.212.04-r0
         - JAVA_HOME=/usr/lib/jvm/java-1.8-openjdk
         - JAVA_VERSION=8u212
         - LANG=C.UTF-8
         - PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/lib/jvm/java-1.8-openjdk/jre/bin:/usr/lib/jvm/java-1.8-openjdk/bin
         - TZ=Asia/Shanghai
        volumes:      
         - /share/Container/reader/config:/config      
        ports:     
         - 33333:8080
```

### GPS轨迹追踪Traccar
https://www.traccar.org/

app下载:https://www.traccar.org/client/

[NAS还能这么玩！搭建在线追踪功能的开源GPS追踪工具 『Traccar』](https://www.toutiao.com/article/7264853678673101351/)

## 迷你主机
铭凡UM790 Pro
零刻GTR7