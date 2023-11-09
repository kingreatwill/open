[TOC]

## VPN
在这一类P2P VPN中比较出名的几种有N2N、Tinc、PeerVPN以及ZeroTier


https://github.com/shadowsocks
https://github.com/shadowsocks/openwrt-shadowsocks
https://github.com/gfwlist/gfwlist


[Shadowsocks + GfwList 实现 OpenWRT / LEDE 路由器自动科学上网](https://cokebar.info/archives/962)

[科学上网的有趣项目集锦](https://github.com/udpsec/awesome-vpn)

### HTTP/HTTPS代理服务器 (正向代理)
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

### OpenVPN

```
https://raw.githubusercontent.com/Nyr/openvpn-install/master/openvpn-install.sh
```
[openvpn一键安装](https://github.com/Nyr/openvpn-install)

一键安装(科学):`wget https://git.io/vpn -O openvpn-install.sh && bash openvpn-install.sh`

参考[OpenVPN的简易搭建流程](https://blog.csdn.net/weixin_43343127/article/details/132709366)

### Clash
https://github.com/Dreamacro/clash

购买了:账号 king...@q
使用说明: https://portal.shadowsocks.au/knowledgebase/151/
https://portal.shadowsocks.nz/clientarea.php
https://portal.shadowsocks.au/clientarea.php

clash 需要自建节点(或购买机场) 安装V2Ray/Trojan服务端

参考[yml文件](./Clash_DIY_node.yml)

[下载地址](https://dl.trojan-cdn.com/trojan/windows/)

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

### natapp
https://natapp.cn/
### ittun
http://ittun.com/
### sunny
https://www.ngrok.cc/
### nat123 : 内网穿透
### frp: 内网穿透
开源
https://github.com/fatedier/frp 41.1k

[安装参考](../docker/frp.md)

### nps/npc
https://github.com/ehang-io/nps
npc - 客户端
安装 https://ehang-io.github.io/nps/#/install

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