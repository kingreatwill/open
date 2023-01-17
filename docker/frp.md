## docker安装

[docker frps 内网穿透容器化服务](https://www.cnblogs.com/LandWind/p/docker-frps-first.html)

[frp 中文文档](https://gofrp.org/docs/examples/ssh/)

### frp server
在公网服务器上安装frps(网络使用host模式)
配置文件`frps.ini`使用`./frp/conf/frps_full.ini`
`docker run --restart=always --network host -d -v /data/dockerv/frp/frps.ini:/etc/frp/frps.ini --name frps snowdreamtech/frps:0.46.0`

### frp client
在公网服务器上安装frpc
配置文件`frpc.ini`使用`./frp/conf/frpc.ini`
`docker run --restart=always --network host -d -v /share/Public/frp/frpc.ini:/etc/frp/frpc.ini --name frpc snowdreamtech/frpc:0.46.0`