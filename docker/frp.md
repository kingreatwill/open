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

### frp开机启动设置
添加systemd配置文件：
```
vim /usr/lib/systemd/system/frp.service
```
文件内容如下：
```
[Unit]
Description=The nginx HTTP and reverse proxy server
After=network.target remote-fs.target nss-lookup.target

[Service]
Type=simple
ExecStart=/usr/local/frp/frps -c /usr/local/frp/frps.ini
KillSignal=SIGQUIT
TimeoutStopSec=5
KillMode=process
PrivateTmp=true
StandardOutput=syslog
StandardError=inherit

[Install]
WantedBy=multi-user.target
```
设置开机启动
```
systemctl daemon-reload
systemctl enable frp
```
启动 frp
```
systemctl start frp
```
查看服务状态
```
systemctl status frp.service
```
查看frp是否启动
```
ps aux | grep frps
```

> centos7 查看启动服务项
使用 systemctl list-unit-files 可以查看启动项
左边是服务名称，右边是状态，enabled是开机启动，disabled是开机不启动
