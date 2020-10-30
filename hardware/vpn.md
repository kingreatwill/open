## VPN
在这一类P2P VPN中比较出名的几种有N2N、Tinc、PeerVPN以及ZeroTier

## 内网穿透工具
内网穿透系列——N2N（简单的P2P组网方案）

刚毕业的时候自己做了一个，实现原理：
内网写个服务获取外网IP，存储到一个外网的某个地方（定时获取外网IP有变化就更新），然后外网就能获取到内网的访问IP啦，开启端口映射（当然路由器要支持端口映射啦）转发到内网服务器的IP上.

### ngrok

Ngrok 实现内网穿透（Ngrok 和 Sunny-Ngrok ） 这个之前用到微信开发上面

### nat123 : 内网穿透
### frp: 内网穿透

### 闪库
http://ipyingshe.com/

管理：http://i.ipyingshe.com/#/dashboard

linux
wget http://down.ipyingshe.com/sk_linux_64
chmod +x sk_linux_64
nohup ./sk_linux_64 -token=xxx &
ps -ef | grep sk_linux_64

https://www.toutiao.com/i6883016451180265998/ 闪库mac版本启动命令教程