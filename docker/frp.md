## docker安装

[docker frps 内网穿透容器化服务](https://www.cnblogs.com/LandWind/p/docker-frps-first.html)

[frp 中文文档](https://gofrp.org/docs/examples/ssh/)

### frp 0.52.3
#### frp server

`docker run --restart=always --network host -d -v /data/dockerv/frp/frps.toml:/etc/frp/frps.toml --name frps ghcr.io/fatedier/frps:v0.66.0 -c /etc/frp/frps.toml`

`docker run --restart=always --network host -d -v /data/dockerv/frp/frps.toml:/etc/frp/frps.toml --name frps ghcr.io/fatedier/frps:v0.61.0 -c /etc/frp/frps.toml`

Since v0.50.0, the default value of transport.tls.enable and transport.tls.disableCustomTLSFirstByte has been changed to true, and tls is enabled by default.


配置文件格式ini变成了toml
`docker run --restart=always --network host -d -v /data/dockerv/frp/frps.toml:/etc/frp/frps.toml --name frps snowdreamtech/frps:0.52.3`

使用frps_full_example.toml文件修改
主要修改:
```
webServer.password = "xxxx"
auth.token = "xxx"
vhostHTTPPort = 7080
vhostHTTPSPort = 7443
subDomainHost = "frp.wcoder.com"
```

注释掉
```
# [[httpPlugins]]
# name = "user-manager"
# addr = "127.0.0.1:9000"
# path = "/handler"
# ops = ["Login"]

# [[httpPlugins]]
# name = "port-manager"
# addr = "127.0.0.1:9001"
# path = "/handler"
# ops = ["NewProxy"]
```
[服务端插件](https://gofrp.org/zh-cn/docs/features/common/server-plugin/)

https://github.com/fatedier/frp/
- httpPlugins
可以理解为钩子插件, 可以用来记录/拒绝/替换返回内容等操作(这样可以记录一些非法访问), 需要一个外部服务来侦听9000端口
可以参考`frp/test/e2e/v1/plugin/server.go` 和 `frp/test/e2e/pkg/plugin/plugin.go`中的代码

#### frp client
`docker run --restart=always --network host -d -v /share/Public/frp/frpc.toml:/etc/frp/frpc.toml --name frpc snowdreamtech/frpc:0.52.3`

#### frp正向代理/反向代理
`docker run --restart=always --network host -d -v /data/dockerv/frp/frpc.toml:/etc/frp/frpc.toml --name frpc snowdreamtech/frpc:0.52.3`

`curl -Lv --proxy http://abc:abc@43.155.152.66:8889  http://www.cnblogs.com/`

客户端配置(部署在服务端的客户端)
```toml
serverAddr = "43.155.152.66"
serverPort = 7000
auth.token = "12345678"

[[proxies]]
name = "plugin_http_proxy"
type = "tcp"
remotePort = 8889
[proxies.plugin]
type = "http_proxy"
httpUser = "abc"
httpPassword = "abc"

[[proxies]]
name = "plugin_socks5"
type = "tcp"
remotePort = 6005
[proxies.plugin]
type = "socks5"
username = "abc"
password = "abc"
```

### frp server
在公网服务器上安装frps(网络使用host模式)
配置文件`frps.ini`使用`./frp/conf/frps_full.ini`
`docker run --restart=always --network host -d -v /data/dockerv/frp/frps.ini:/etc/frp/frps.ini --name frps snowdreamtech/frps:0.46.0`

### frp client
在NAS服务器上安装frpc
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

### 自定义二级域名(0.52.3版本)

在多人同时使用一个 frps 时，通过自定义二级域名的方式来使用会更加方便。

通过在 frps 的配置文件中配置 `subdomainHost`，就可以启用该特性。之后在 frpc 的 http、https 类型的代理中可以不配置 customDomains，而是配置一个 `subdomain` 参数。

只需要将 `*.{subdomainHost}` 解析到 frps 所在服务器。之后用户可以通过 subdomain 自行指定自己的 web 服务所需要使用的二级域名，通过 `{subdomain}.{subdomainHost}:{vhostHTTPPort}` 来访问自己的 web 服务。

> 需要加端口
> vhostHTTPPort = 7080
> vhostHTTPSPort = 7443

```
# frps.toml
subdomainHost = "frps.com"
```

将泛域名 `*.frps.com` 解析到 frps 所在服务器的 IP 地址。

```
# frpc.toml
[[proxies]]
name = "web"
type = "http"
localPort = 80
subdomain = "test"
```
frps 和 frpc 都启动成功后，通过 `test.frps.com` 就可以访问到内网的 web 服务。

注：如果 frps 配置了 `subdomainHost`，则 `customDomains` 中不能是属于 `subdomainHost` 的子域名或者泛域名。

同一个 HTTP 或 HTTPS 类型的代理中 `customDomains` 和 `subdomain` 可以同时配置。


[参考 自定义二级域名](https://gofrp.org/zh-cn/docs/features/http-https/subdomain/)