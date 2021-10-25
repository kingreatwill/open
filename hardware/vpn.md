[TOC]

## VPN
在这一类P2P VPN中比较出名的几种有N2N、Tinc、PeerVPN以及ZeroTier


https://github.com/shadowsocks
https://github.com/shadowsocks/openwrt-shadowsocks
https://github.com/gfwlist/gfwlist


[Shadowsocks + GfwList 实现 OpenWRT / LEDE 路由器自动科学上网](https://cokebar.info/archives/962)

[科学上网的有趣项目集锦](https://github.com/udpsec/awesome-vpn)

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

## 内网穿透工具
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