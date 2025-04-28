[TOC]

## VPN
在这一类P2P VPN中比较出名的几种有N2N、Tinc、PeerVPN以及ZeroTier


https://github.com/shadowsocks
https://github.com/shadowsocks/openwrt-shadowsocks
https://github.com/gfwlist/gfwlist


[Shadowsocks + GfwList 实现 OpenWRT / LEDE 路由器自动科学上网](https://cokebar.info/archives/962)


[科学上网的有趣项目集锦](https://github.com/udpsec/awesome-vpn)
### v2rayA
https://github.com/v2rayA/v2rayA
#### gg
https://github.com/mzz2017/gg

### 搭建shadowsocks
1. 安装谷歌 BBR 加速
`wget --no-check-certificate https://github.com/teddysun/across/raw/master/bbr.sh && chmod +x bbr.sh && ./bbr.sh`

显示 “Press any key to start…” 按回车确认。回车后会出现一列内核版本让我们选择，输入序号 61 并回车开始安装。
安装完后，按提示重启 VPS，输入 Y 回车重启。稍候 1min 等待重启完成，再重新连接 Xshell。
重启后输入
`lsmod | grep bbr`
出现 tcp_bbr 即说明 BBR 已经启动。

2. 安装SS
依次运行下面三行命令，如下图所示按要求输入相应信息。（建议：端口选择大于 30000 的。）

```sh
wget — no-check-certificate -O shadowsocks.sh https://raw.githubusercontent.com/teddysun/shadowsocks_install/master/shadowsocks.sh
chmod +x shadowsocks.sh
./shadowsocks.sh 2>&1 | tee shadowsocks.log
```

输入和记录IP,端口,密码和加密方式

3. 使用方法
https://github.com/shadowsocks/shadowsocks-windows/releases
https://github.com/shadowsocks/shadowsocks-android/releases

可以选择合适的代理模式。
PAC： 只代理国外网站；
全局： 所有网站都通过SS。

> 参考: [搭建自己的科学上网服务器](https://github.com/xiaoming-ssr/FanQiang-Book/wiki/%E6%90%AD%E5%BB%BA%E8%87%AA%E5%B7%B1%E7%9A%84%E7%A7%91%E5%AD%A6%E4%B8%8A%E7%BD%91%E6%9C%8D%E5%8A%A1%E5%99%A8)

### HTTP/HTTPS代理服务器 (正向代理)
`curl -Lv --proxy http://xx:xx@43.155.152.66:8888  http://www.cnblogs.com/`

#### Squid、Privoxy、Varnish、Polipo
#### tinyproxy
https://github.com/tinyproxy

centos8
```
cd /opt/
git clone https://github.com/tinyproxy/tinyproxy.git


./autogen.sh
./configure
make
make install

cp etc/tinyproxy.conf /etc/

vim /etc/tinyproxy.conf

```
/etc/tinyproxy.conf
```
 
#修改下面配置，允许所有地址，添加basic授权
 
Port 18888
 
#Allow 127.0.0.1
#Allow ::1
 
# BasicAuth: HTTP "Basic Authentication" for accessing the proxy.
# If there are any entries specified, access is only granted for authenticated
# users.
BasicAuth user123 hahapwd
```

```
vim /usr/lib/systemd/system/tinyproxy.service
```

```
[Unit]
Description=Tinyproxy Server Service
After=network.target
 
[Service]
Type=simple
User=nobody
Restart=on-failure
RestartSec=5s
ExecStart=/usr/local/bin/tinyproxy -c etc/tinyproxy.conf -d
```

```

systemctl start tinyproxy
systemctl status tinyproxy

# restart：重启服务，可启动服务
# reload：服务重新加载配置文件
# status：查看服务状态
# start：启动服务
# stop：停止服务
# enable：开启自启
# disable：关闭自启
```


```
export https_proxy=http://user123:hahapwd@47.113.67.125:18888
```
#### caddy forwardproxy

[正向代理 forwardproxy](../articles/caddy.md#forwardproxy)

#### dumbproxy/HTTP proxy
https://github.com/SenseUnit/dumbproxy

#### nginx proxy
[nginx forward proxy](https://www.cnblogs.com/yanjieli/p/15229907.html)

核心配置
```
server {

    listen 8888;

    location / {

        resolver 8.8.8.8;

        proxy_pass http://$http_host$uri$is_args$args;

    }

}
```

[docker 部署openresty](../docker/openresty.md)

#### frp

[最新配置参见](../docker/frp.md)

```
[proxy]
type = tcp
remote_port = 6000
plugin = http_proxy
plugin_http_user = abc
plugin_http_passwd = abc

[socks5_proxy]
type = tcp
remote_port = 6005
plugin = socks5
plugin_user = abc
plugin_passwd = abc
```

### OpenVPN

```
https://raw.githubusercontent.com/Nyr/openvpn-install/master/openvpn-install.sh
```
[openvpn一键安装](https://github.com/Nyr/openvpn-install)

一键安装(科学):`wget https://git.io/vpn -O openvpn-install.sh && bash openvpn-install.sh`

参考[OpenVPN的简易搭建流程](https://blog.csdn.net/weixin_43343127/article/details/132709366)

### Clash
https://github.com/Dreamacro/clash

https://github.com/clash-verge-rev/clash-verge-rev
https://github.com/keiko233/clash-nyanpasu
https://github.com/MetaCubeX/mihomo

购买了:账号 king...@q
使用说明: https://portal.shadowsocks.au/knowledgebase/151/
https://portal.shadowsocks.nz/clientarea.php
https://portal.shadowsocks.au/clientarea.php

clash 需要自建节点(或购买机场) 安装V2Ray/Trojan服务端

参考[yml文件](./Clash_DIY_node.yml)

[下载地址](https://dl.trojan-cdn.com/trojan/windows/)

[免费节点](https://rss.uk.to/#/register?code=srOLpruw)

[免费ss/v2ray/trojan节点](https://github.com/freefq/free)

[V2Ray 收费节点](https://v2sx.com/cart.php)


**Clash**：
多协议支持： Clash 支持多种代理协议，如 Shadowsocks、Vmess、Trojan 等，这意味着您可以根据不同的使用场景和需求来选择合适的协议。
灵活的配置： Clash 提供强大的配置选项，允许用户自定义代理规则，以实现更精细的流量控制和路由设置。
图形化界面： Clash 有许多基于图形界面的客户端，使其配置和使用变得更加直观和便捷。
广泛的社区支持： 由于其受欢迎程度，Clash 有一个活跃的社区，可以获得丰富的教程、配置示例和技术支持。

**V2Ray**：
专注的协议： V2Ray 专注于 Vmess 协议，这是一种高度安全的代理协议，支持更多的加密和安全特性。
灵活的路由规则： V2Ray 提供了强大的路由功能，允许您根据域名、IP、端口等条件精确地控制流量的转发。
插件支持： V2Ray 支持插件系统，可以通过插件实现更多的功能，如流量伪装、混淆等。
开源项目： V2Ray 是一个开源项目，具有透明的代码，用户可以审查和定制代码来满足自己的需求。


[免费节点](https://topvpnlist.github.io/)
[龙夫山泉15-200G](https://www.nfsq.us/#/plan) 3...@qq.com, qq3..


#### Trojan 一键脚本
CentOS
`yum update -y && yum install wget -y && yum install curl -y`

Debian / Ubuntu
`apt-get update -y && apt-get install wget -y && apt-get install curl -y`

```
curl -O https://raw.githubusercontent.com/atrandys/trojan/master/trojan_mult.sh && chmod +x trojan_mult.sh && ./trojan_mult.sh
```

https://github.com/atrandys/trojan
https://github.com/jinwyp/one_click_script


#### 客户端安装
[V2RayNG	Android](https://github.com/2dust/v2rayNG)
[v2rayN	Windows](https://github.com/2dust/v2rayN)

[trojan-android](https://github.com/xxf098/shadowsocksr-v2ray-trojan-android)

### 付费VPN
#### tlyvpn
https://www.ebay.com/usr/tlyvpn
#### NordVPN

## 端口映射
https://www.cnblogs.com/connect/p/server-port-proxy.html

### Windows下实现端口映射
[netsh命令](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/netsh)

查询端口映射情况
`netsh interface portproxy show v4tov4`

查询某一个IP的所有端口映射情况
`netsh interface portproxy show v4tov4 | find "[IP]"`
例：
`netsh interface portproxy show v4tov4 | find "192.168.1.1"`

增加一个端口映射
`netsh interface portproxy add v4tov4 listenaddress=[外网IP] listenport=[外网端口] connectaddress=[内网IP] connectport=[内网端口]`
例：
`netsh interface portproxy add v4tov4 listenaddress=2.2.2.2 listenport=8080 connectaddress=192.168.1.50 connectport=80`

删除一个端口映射
`netsh interface portproxy delete v4tov4 listenaddress=[外网IP] listenport=[外网端口]`
例：
`netsh interface portproxy delete v4tov4 listenaddress=2.2.2.2 listenport=8080`


### Linux下实现端口映射
1. 允许数据包转发
```
echo 1 >/proc/sys/net/ipv4/ip_forward
iptables -t nat -A POSTROUTING -j MASQUERADE
iptables -A FORWARD -i [内网网卡名称] -j ACCEPT
iptables -t nat -A POSTROUTING -s [内网网段] -o [外网网卡名称] -j MASQUERADE
```
例：
```
echo 1 >/proc/sys/net/ipv4/ip_forward
iptables -t nat -A POSTROUTING -j MASQUERADE
iptables -A FORWARD -i ens33 -j ACCEPT
iptables -t nat -A POSTROUTING -s 192.168.50.0/24 -o ens37 -j MASQUERADE
```
2. 设置端口映射
```
iptables -t nat -A PREROUTING -p tcp -m tcp --dport [外网端口] -j DNAT --to-destination [内网地址]:[内网端口]
```
例：
```
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 6080 -j DNAT --to-destination 10.0.0.100:6090
```

## 内网穿透工具(nat打洞技术)

内网穿透系列——N2N（简单的P2P组网方案）

刚毕业的时候自己做了一个，实现原理：
内网写个服务获取外网IP，存储到一个外网的某个地方（定时获取外网IP有变化就更新），然后外网就能获取到内网的访问IP啦，开启端口映射（当然路由器要支持端口映射啦）转发到内网服务器的IP上.


### ngrok

Ngrok 实现内网穿透（Ngrok 和 Sunny-Ngrok ） 这个之前用到微信开发上面
https://ngrok.com/

https://github.com/inconshreveable/ngrok

### natapp
https://natapp.cn/
### ittun
http://ittun.com/
### sunny
https://www.ngrok.cc/
### nat123 : 内网穿透
### frp: 内网穿透
开源
https://github.com/fatedier/frp

[安装参考](../docker/frp.md)

### nps/npc
https://github.com/ehang-io/nps
npc - 客户端
安装 https://ehang-io.github.io/nps/#/install

### Rathole
https://github.com/rapiz1/rathole

### Microsoft dev tunnels
https://github.com/microsoft/dev-tunnels

[安装文档](https://learn.microsoft.com/zh-cn/azure/developer/dev-tunnels/get-started?tabs=windows)

```
登录账号：
devtunnel.exe user login -g -d    （如果电脑有浏览器优先调起浏览器实现登陆授权）

启用转发端口：
devtunnel.exe host -p 8000
浏览器访问：https://601tvmcl.inc1.devtunnels.ms:8000  或者  https://601tvmcl-8000.inc1.devtunnels.ms 使用授权的Github账号登陆即可访问！

启动匿名访问：
devtunnel.exe host -p 8000 --allow-anonymous
任何人浏览器访问 https://sc42hdl7.inc1.devtunnels.ms:8000  或者  https://sc42hdl7-8000.inc1.devtunnels.ms 即可！

为使用 HTTPS 协议的端口 8443 上的服务器托管临时开发隧道。
devtunnel host -p 8443 --protocol https
```

### goproxy
https://github.com/snail007/goproxy

### Cloudflare Tunnel

https://github.com/louislam/uptime-kuma/wiki/Reverse-Proxy-with-Cloudflare-Tunnel

### Localtunnel
https://github.com/localtunnel/localtunnel

### Chisel
https://github.com/jpillora/chisel


### sshuttle
https://github.com/sshuttle/sshuttle

### Bore
https://github.com/ekzhang/bore

### Zrok
https://github.com/openziti/zrok

### NSmartProxy
https://github.com/tmoonlight/NSmartProxy

### Headscale
Headscale + Tailscale 

WireGuard 是最新一代开源 VPN 工具，相较于IPsec和OpenVPN等大多数古老魔法，具有易于配置、快速且安全的特点。

[WireGuard一键安装](https://github.com/Nyr/wireguard-install)
### tailscale
在用户态实现了WireGuard协议, 无法使用WireGuard原生的命令行工具来进行管理

### NetBird
直接使用了内核态的WireGuard, 可使用命令行工具wg来查看和管理
### twingate
### zerotier
### inlets
开源
https://github.com/inlets/inlets 6.9k
### Argo Tunnel
https://developers.cloudflare.com/argo-tunnel/

### 闪库
http://ipyingshe.com/

管理：http://i.ipyingshe.com/#/dashboard

linux
wget http://down.ipyingshe.com/sk_linux_64
chmod +x sk_linux_64
nohup ./sk_linux_64 -token=xxx &
ps -ef | grep sk_linux_64

https://www.toutiao.com/i6883016451180265998/ 闪库mac版本启动命令教程

### FastTunnel
https://github.com/FastTunnel/FastTunnel
类似花生壳

- 家中建站
- 微信开发
- 远程桌面
- erp互通
- svn代码仓库
- 端口转发
- iot物联网
- 联机游戏
- 等等场景，不局限以上

## 微信调试
远程调试手机页面和抓包的工具
h5/wap端调试、移动端调试

- https://github.com/wuchangming/spy-debugger
- https://github.com/Tencent/vConsole
类似小程序的调试工具，可以在手机上看见打印，真机调试h5时，我们只能看alert,但有些内容无法alert,但vconsole能做到，就像谷歌的f12一样

- https://github.com/liriliri/eruda
和vconsole 相似