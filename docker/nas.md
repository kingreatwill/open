
[NAS](../hardware/nas.md)

## 威联通docker-compose
docker-compose.yml文件位于以下目录
/share/Container/container-station-data/application/applicationname/docker-compose.yml

注意docker-compose.yml 后缀一定是yml, 以及applicationname需要完全匹配

### nas控制台-网络访问-代理
这样设置的代理仅能command命令使用

#### docker pull 代理
sudo -i 切换成 admin 操作
也可以在遇到权限不足的时候使用sudo

1. 进入到docker配置目录 （默认是CACHEDEC1_DATA）

cd /share/CACHEDEV1_DATA/.qpkg/container-station/script

2. 编辑run-docker.sh

vi run-docker.sh

3. 切换到最下面倒数第二行，新增以下环境（ip改成代理局域网电脑IP与设置端口）：

```
# rm -rf ...

export http_proxy="http://192.168.168.89:7890"
export https_proxy="http://192.168.168.89:7890"
export no_proxy= "192.168.168.0/24,localhost,127.0.0.1"

# exec dockerd ..
```

4. 重启container station使之生效

`/etc/init.d/container-station.sh restart`
> 权限不足加sudo

#### Clash
https://github.com/wnlen/clash-for-linux

##### docker

1. https://github.com/LaoYutang/clash-and-dashboard

2. https://github.com/vernesong/OpenClash

`docker run -d –name clash -p 7890:7890 -p 9090:9090 -v /path/to/config:/root/.config/clash ghcr.io/vernesong/openclash:latest`

- 访问 Clash Dashboard
http://192.168.0.1:9090/ui
在Secret(optional)一栏中输入启动成功后输出的Secret。


- 代理模式, 环境变量: -e CLASH_MODE=rule
CLASH_MODE=rule: 根据规则进行代理
CLASH_MODE=global: 全局代理
CLASH_MODE=direct: 直连模式

- 自动更新订阅
```
#!/bin/bash

subscribe_url=”https://example.com/subscribe”

curl -o /path/to/config/proxy.yaml $subscribe_url

docker restart clash
```

3. https://github.com/haishanh/yacd


- 在获取到 Clash 配置文件（config.yaml）后，需要确认以下配置：
```
# config.yaml

port: 7890
socks-port: 7891
allow-lan: true

# rule / global / direct (default is rule)
mode: rule

# set log level to stdout (default is info)
# info / warning / error / debug / silent
log-level: info

# RESTful API(基于 API，可以打造自己的可视化操作部分，也是实现 Clash GUI 的重要组成部分。)
external-controller: 127.0.0.1:9090
# external-controller = 127.0.0.1:8080

# 你可以加入 secret 进行 API 鉴权
# 鉴权的方式为在 Http Header 中加入 Authorization: Bearer ${secret}
# secret = thisisyoursecret
```
> 7890 为 http/https 监听端口，7891 为 socks5 监听端口，9090 为 UI 监听端口，allow-lan 为允许局域网访问。
> 如果不是为了特殊需求，请尽量不要把 API 暴露在 0.0.0.0，如果真的要这么做，一定要加上 secret 进行鉴权
> https://clash.gitbook.io/doc/restful-api

- 完整example 
```
# port of HTTP
port: 7890

# port of SOCKS5
socks-port: 7891

# redir port for Linux and macOS
# redir-port: 7892

allow-lan: false

# Only applicable when setting allow-lan to true
# "*": bind all IP addresses
# 192.168.122.11: bind a single IPv4 address
# "[aaaa::a8aa:ff:fe09:57d8]": bind a single IPv6 address
# bind-address: "*"

# rule / global / direct (default is rule)
mode: rule

# set log level to stdout (default is info)
# info / warning / error / debug / silent
log-level: info

# RESTful API for clash
external-controller: 127.0.0.1:9090

# you can put the static web resource (such as clash-dashboard) to a directory, and clash would serve in `${API}/ui`
# input is a relative path to the configuration directory or an absolute path
# external-ui: folder

# Secret for RESTful API (Optional)
# secret: ""

# experimental feature
experimental:
  ignore-resolve-fail: true # ignore dns resolve fail, default value is true
  # interface-name: en0 # outbound interface name

# authentication of local SOCKS5/HTTP(S) server
# authentication:
#  - "user1:pass1"
#  - "user2:pass2"

# # hosts, support wildcard (e.g. *.clash.dev Even *.foo.*.example.com)
# # static domain has a higher priority than wildcard domain (foo.example.com > *.example.com > .example.com)
# # +.foo.com equal .foo.com and foo.com
# hosts:
#   '*.clash.dev': 127.0.0.1
#   '.dev': 127.0.0.1
#   'alpha.clash.dev': '::1'
#   '+.foo.dev': 127.0.0.1

# dns:
  # enable: true # set true to enable dns (default is false)
  # ipv6: false # default is false
  # listen: 0.0.0.0:53
  # # default-nameserver: # resolve dns nameserver host, should fill pure IP
  # #   - 114.114.114.114
  # #   - 8.8.8.8
  # enhanced-mode: redir-host # or fake-ip
  # # fake-ip-range: 198.18.0.1/16 # if you don't know what it is, don't change it
  # fake-ip-filter: # fake ip white domain list
  #   - '*.lan'
  #   - localhost.ptlogin2.qq.com
  # nameserver:
  #   - 114.114.114.114
  #   - tls://dns.rubyfish.cn:853 # dns over tls
  #   - https://1.1.1.1/dns-query # dns over https
  # fallback: # concurrent request with nameserver, fallback used when GEOIP country isn't CN
  #   - tcp://1.1.1.1
  # fallback-filter:
  #   geoip: true # default
  #   ipcidr: # ips in these subnets will be considered polluted
  #     - 240.0.0.0/4

proxies:
  # shadowsocks
  # The supported ciphers(encrypt methods):
  #   aes-128-gcm aes-192-gcm aes-256-gcm
  #   aes-128-cfb aes-192-cfb aes-256-cfb
  #   aes-128-ctr aes-192-ctr aes-256-ctr
  #   rc4-md5 chacha20-ietf xchacha20
  #   chacha20-ietf-poly1305 xchacha20-ietf-poly1305
  - name: "ss1"
    type: ss
    server: server
    port: 443
    cipher: chacha20-ietf-poly1305
    password: "password"
    # udp: true

  # old obfs configuration format remove after prerelease
  - name: "ss2"
    type: ss
    server: server
    port: 443
    cipher: chacha20-ietf-poly1305
    password: "password"
    plugin: obfs
    plugin-opts:
      mode: tls # or http
      # host: bing.com

  - name: "ss3"
    type: ss
    server: server
    port: 443
    cipher: chacha20-ietf-poly1305
    password: "password"
    plugin: v2ray-plugin
    plugin-opts:
      mode: websocket # no QUIC now
      # tls: true # wss
      # skip-cert-verify: true
      # host: bing.com
      # path: "/"
      # mux: true
      # headers:
      #   custom: value

  # vmess
  # cipher support auto/aes-128-gcm/chacha20-poly1305/none
  - name: "vmess"
    type: vmess
    server: server
    port: 443
    uuid: uuid
    alterId: 32
    cipher: auto
    # udp: true
    # tls: true
    # skip-cert-verify: true
    # servername: example.com # priority over wss host
    # network: ws
    # ws-path: /path
    # ws-headers:
    #   Host: v2ray.com
  
  - name: "vmess-http"
    type: vmess
    server: server
    port: 443
    uuid: uuid
    alterId: 32
    cipher: auto
    # udp: true
    # network: http
    # http-opts:
    #   # method: "GET"
    #   # path:
    #   #   - '/'
    #   #   - '/video'
    #   # headers:
    #   #   Connection:
    #   #     - keep-alive

  # socks5
  - name: "socks"
    type: socks5
    server: server
    port: 443
    # username: username
    # password: password
    # tls: true
    # skip-cert-verify: true
    # udp: true

  # http
  - name: "http"
    type: http
    server: server
    port: 443
    # username: username
    # password: password
    # tls: true # https
    # skip-cert-verify: true

  # snell
  - name: "snell"
    type: snell
    server: server
    port: 44046
    psk: yourpsk
    # obfs-opts:
      # mode: http # or tls
      # host: bing.com

  # trojan
  - name: "trojan"
    type: trojan
    server: server
    port: 443
    password: yourpsk
    # udp: true
    # sni: example.com # aka server name
    # alpn:
    #   - h2
    #   - http/1.1
    # skip-cert-verify: true

proxy-groups:
  # relay chains the proxies. proxies shall not contain a relay. No UDP support.
  # Traffic: clash <-> http <-> vmess <-> ss1 <-> ss2 <-> Internet
  - name: "relay"
    type: relay
    proxies:
      - http
      - vmess
      - ss1
      - ss2

  # url-test select which proxy will be used by benchmarking speed to a URL.
  - name: "auto"
    type: url-test
    proxies:
      - ss1
      - ss2
      - vmess1
    # tolerance: 150
    url: 'http://www.gstatic.com/generate_204'
    interval: 300

  # fallback select an available policy by priority. The availability is tested by accessing an URL, just like an auto url-test group.
  - name: "fallback-auto"
    type: fallback
    proxies:
      - ss1
      - ss2
      - vmess1
    url: 'http://www.gstatic.com/generate_204'
    interval: 300

  # load-balance: The request of the same eTLD will be dial on the same proxy.
  - name: "load-balance"
    type: load-balance
    proxies:
      - ss1
      - ss2
      - vmess1
    url: 'http://www.gstatic.com/generate_204'
    interval: 300

  # select is used for selecting proxy or proxy group
  # you can use RESTful API to switch proxy, is recommended for use in GUI.
  - name: Proxy
    type: select
    proxies:
      - ss1
      - ss2
      - vmess1
      - auto
  
  - name: UseProvider
    type: select
    use:
      - provider1
    proxies:
      - Proxy
      - DIRECT

proxy-providers:
  provider1:
    type: http
    url: "url"
    interval: 3600
    path: ./hk.yaml
    health-check:
      enable: true
      interval: 600
      url: http://www.gstatic.com/generate_204
  test:
    type: file
    path: /test.yaml
    health-check:
      enable: true
      interval: 36000
      url: http://www.gstatic.com/generate_204

rules:
  - DOMAIN-SUFFIX,google.com,auto
  - DOMAIN-KEYWORD,google,auto
  - DOMAIN,google.com,auto
  - DOMAIN-SUFFIX,ad.com,REJECT
  # rename SOURCE-IP-CIDR and would remove after prerelease
  - SRC-IP-CIDR,192.168.1.201/32,DIRECT
  # optional param "no-resolve" for IP rules (GEOIP IP-CIDR)
  - IP-CIDR,127.0.0.0/8,DIRECT
  - GEOIP,CN,DIRECT
  - DST-PORT,80,DIRECT
  - SRC-PORT,7777,DIRECT
  # FINAL would remove after prerelease
  # you also can use `FINAL,Proxy` or `FINAL,,Proxy` now
  - MATCH,auto

```


- 运行clash: docker pull dreamacro/clash:v1.18.0
```
docker run --name clash \
    -p 5090:9090 -p 5890:7890 -p 5891:7891 \
    -v ~/clash/config.yaml:/root/.config/clash/config.yaml -d dreamacro/clash
```
- 运行 Clash UI
```
docker run --name clash-ui -p 5080:80 -d ghcr.io/haishanh/yacd:master
```
- 通过 Clash UI 管理、监控 Clash 服务
使用浏览器打开地址：`http://[主机IP]:5080`，然后在输入框内输入 `http://[主机IP]:5090`，再点击 ADD 按钮，然后点击下方新增的 `http://[主机IP]:5090` 链接进入监控界面。

- docker-compose
```
# docker-compose.yml

version: '3.7'
services:
  clash-server:
    image: ghcr.io/kingreatwill/docker.io/dreamacro/clash:v1.18.0
    restart: unless-stopped
    container_name: clash
    ports:
      - "17990:9090"
      - "17890:7890"
      - "17891:7891"
    volumes:
      - /share/Public/clash/config.yaml:/root/.config/clash/config.yaml

  clash-ui:
    image: ghcr.io/haishanh/yacd:master
    restart: unless-stopped
    container_name: clash-ui
    ports:
      - 17080:80
```



#### 修改Docker仓库镜像
`/share/CACHEDEV1_DATA/.qpkg/container-station/etc/docker.json`

```
{
 "registry-mirrors": ["http://hub-mirror.c.163.com"] 
}
```
重启: `/etc/init.d/container-station.sh restart`

> [docker 代理](./README.md)


### docker 浏览器
https://github.com/linuxserver/docker-firefox
https://github.com/linuxserver/docker-chromium
https://github.com/linuxserver/docker-msedge
https://github.com/m1k1o/neko
https://github.com/jlesage/docker-firefox

> 威联通自带Browser Station, http://nas.frp.wcoder.com/browser-station
> 支持Chrome和SeaMonkey

### docker office
https://github.com/linuxserver/docker-wps-office

### paperless-ngx/无纸化
```docker-compose.yml
# Docker Compose file for running paperless from the docker container registry.
# This file contains everything paperless needs to run.
# Paperless supports amd64, arm and arm64 hardware.
#
# All compose files of paperless configure paperless in the following way:
#
# - Paperless is (re)started on system boot, if it was running before shutdown.
# - Docker volumes for storing data are managed by Docker.
# - Folders for importing and exporting files are created in the same directory
#   as this file and mounted to the correct folders inside the container.
# - Paperless listens on port 8000.
#
# In addition to that, this Docker Compose file adds the following optional
# configurations:
#
# - Instead of SQLite (default), PostgreSQL is used as the database server.
# - Apache Tika and Gotenberg servers are started with paperless and paperless
#   is configured to use these services. These provide support for consuming
#   Office documents (Word, Excel, Power Point and their LibreOffice counter-
#   parts.
#
# To install and update paperless with this file, do the following:
#
# - Copy this file as 'docker-compose.yml' and the files 'docker-compose.env'
#   and '.env' into a folder.
# - Run 'docker compose pull'.
# - Run 'docker compose run --rm webserver createsuperuser' to create a user.
# - Run 'docker compose up -d'.
#
# For more extensive installation and update instructions, refer to the
# documentation.

services:
  broker:
    image: docker.io/library/redis:7
    restart: unless-stopped
    volumes:
      - redisdata:/data

  db:
    image: docker.io/library/postgres:16
    restart: unless-stopped
    volumes:
      - pgdata:/var/lib/postgresql/data
    environment:
      POSTGRES_DB: paperless
      POSTGRES_USER: paperless
      POSTGRES_PASSWORD: paperless

  webserver:
    image: ghcr.io/paperless-ngx/paperless-ngx:latest
    restart: unless-stopped
    depends_on:
      - db
      - broker
      - gotenberg
      - tika
    ports:
      - "8000:8000"
    volumes:
      - data:/usr/src/paperless/data
      - media:/usr/src/paperless/media
      - ./export:/usr/src/paperless/export
      - ./consume:/usr/src/paperless/consume
    env_file: docker-compose.env
    environment:
      PAPERLESS_REDIS: redis://broker:6379
      PAPERLESS_DBHOST: db
      PAPERLESS_TIKA_ENABLED: 1
      PAPERLESS_TIKA_GOTENBERG_ENDPOINT: http://gotenberg:3000
      PAPERLESS_TIKA_ENDPOINT: http://tika:9998

  gotenberg:
    image: docker.io/gotenberg/gotenberg:8.7
    restart: unless-stopped

    # The gotenberg chromium route is used to convert .eml files. We do not
    # want to allow external content like tracking pixels or even javascript.
    command:
      - "gotenberg"
      - "--chromium-disable-javascript=true"
      - "--chromium-allow-list=file:///tmp/.*"

  tika:
    image: docker.io/apache/tika:latest
    restart: unless-stopped

volumes:
  data:
  media:
  pgdata:
  redisdata:
```

```docker-compose.env
# The UID and GID of the user used to run paperless in the container. Set this
# to your UID and GID on the host so that you have write access to the
# consumption directory.
#USERMAP_UID=1000
#USERMAP_GID=1000

# Additional languages to install for text recognition, separated by a
# whitespace. Note that this is
# different from PAPERLESS_OCR_LANGUAGE (default=eng), which defines the
# language used for OCR.
# The container installs English, German, Italian, Spanish and French by
# default.
# See https://packages.debian.org/search?keywords=tesseract-ocr-&searchon=names&suite=buster
# for available languages.
#PAPERLESS_OCR_LANGUAGES=tur ces

###############################################################################
# Paperless-specific settings                                                 #
###############################################################################

# All settings defined in the paperless.conf.example can be used here. The
# Docker setup does not use the configuration file.
# A few commonly adjusted settings are provided below.

# This is required if you will be exposing Paperless-ngx on a public domain
# (if doing so please consider security measures such as reverse proxy)
#PAPERLESS_URL=https://paperless.example.com

# Adjust this key if you plan to make paperless available publicly. It should
# be a very long sequence of random characters. You don't need to remember it.
#PAPERLESS_SECRET_KEY=change-me

# Use this variable to set a timezone for the Paperless Docker containers. If not specified, defaults to UTC.
#PAPERLESS_TIME_ZONE=America/Los_Angeles

# The default language to use for OCR. Set this to the language most of your
# documents are written in.
#PAPERLESS_OCR_LANGUAGE=eng

# Set if accessing paperless via a domain subpath e.g. https://domain.com/PATHPREFIX and using a reverse-proxy like traefik or nginx
#PAPERLESS_FORCE_SCRIPT_NAME=/PATHPREFIX
#PAPERLESS_STATIC_URL=/PATHPREFIX/static/ # trailing slash required
```

```.env
COMPOSE_PROJECT_NAME=paperless
```

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

### windows in docker
https://github.com/dockur/windows
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

### Internet OS
- [Internet OS](https://github.com/HeyPuter/puter) - [体验](https://puter.com/)

### 内网办公操作系统
https://github.com/phpk/godoos


### HomeAssistant
#### 虚拟机安装HomeAssistant
1. 下载[HA OS](https://www.home-assistant.io/installation/alternative)镜像包, 选择.ova格式的镜像文件
2. 在Virtualization Station应用主界面选择【虚拟机】然后选择【导入】。
3. 启动虚拟机,端口8123(大概要20分钟左右)
4. 设置账号
5. 安装HACS(Home Assistant Community Store)相当于是一个插件的应用商店，比如米家、美的、格力、特斯拉等待
   1. 首先点击左下角的个人用户名，页面设置里打开高级模式
   2. 然后在配置-加载项-加载项商店，找到Terminal & SHH，点击安装。
   3. 安装完成后选择启动，并打开WEB UI。
   4. 使用命令行安装`wget -O - https://hacs.vip/get | bash -`

> [如何用威联通NAS部署HomeAssistant，让苹果Homekit接入所有智能家电](https://post.smzdm.com/p/a4po0qex/)

#### docker安装HomeAssistant
https://www.home-assistant.io/installation/alternative#qnap-nas



## 迷你主机
铭凡UM790 Pro
零刻GTR7
mac mini