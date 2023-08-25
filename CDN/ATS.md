# Apache Traffic Server(ATS or TS)
https://trafficserver.apache.org/
https://github.com/apache/trafficserver


> 同类产品: Squid

## 安装编译工具

```
sudo apt-get install build-essential libtool openssl libssl-dev automake autoconf zlib1g zlib1g-dev libpcre3 libpcre3-dev
```

```
sudo apt install automake libtool pkg-config libmodule-install-perl gcc g++ libssl-dev tcl-dev libpcre3-dev libcap-dev libhwloc-dev libncurses5-dev libcurl4-openssl-dev flex autotools-dev bison debhelper dh-apparmor gettext intltool-debian libbison-dev libexpat1-dev libfl-dev libsigsegv2 libsqlite3-dev m4 po-debconf tcl8.6-dev zlib1g-dev
```

### 一键安装脚本
https://github.com/Har-Kuun/OneClickCDN

```
#原版（英文界面）：
wget https://raw.githubusercontent.com/Har-Kuun/OneClickCDN/master/OneClickCDN.sh && sudo bash OneClickCDN.sh
#中文版：
wget https://raw.githubusercontent.com/Har-Kuun/OneClickCDN/master/translation/translated_scripts/OneClickCDN_zh-CN.sh && sudo bash OneClickCDN_zh-CN.sh
```

## 命令行工具

### traffic_top

https://docs.trafficserver.apache.org/en/9.2.x/appendices/command-line/traffic_top.en.html



## Apache Traffic Control
https://trafficcontrol.apache.org/
https://traffic-control-cdn.readthedocs.io/en/latest/index.html
https://github.com/apache/trafficcontrol

语言:golang

Apache Traffic Control is an Open Source implementation of a Content Delivery Network
Traffic Control以Apache Traffic Server作为缓存软件构建，实现了现代CDN的所有核心功能。
Apache Traffic Control 是一个分布式、可扩展的冗余解决方案，实现了现代 CDN 的所有核心功能，可用于构建、监视和配置大型内容交付网络。项目起源于 Traffic Server ，于 2015 年 4 月作为开源项目发布，并于 2016 年 7 月捐赠给 Apache 孵化器。





## 参考
[Apache Traffic Server 教程](https://www.cnblogs.com/ColoFly/p/16230167.html)
[官方文档](https://trafficserver.apache.org/)

[自建CDN系列：Traffic Server](https://www.blueskyxn.com/202007/1666.html)

[AlmaLinux 使用ATS搭建CDN](https://apad.pro/traffic-server-cdn/)