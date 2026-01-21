[TOC]

> [Caddy](https://github.com/caddyserver/caddy) 是一个 Go 编写的 Web 服务器,类似于 Nginx

## 文档

[官方文档](http://nginx.org/en/docs)

[NGINX缓存原理及源码分析(一)](https://www.nginx.org.cn/article/detail/403)
[NGINX缓存原理及源码分析(二)](https://www.nginx.org.cn/article/detail/406)

### parsing nginx cache header files/解析缓存文件头部信息
https://github.com/kroemeke/ngcls
```
dft:/cache/nginx/0/00# ngcls 9eacc540a431495f5ae408412e60f000 
version       : 3
valid_sec     : 1443784216 2015-10-02 12:10:16
last_modified : 1443559345 2015-09-29 21:42:25
date          : 1443697816 2015-10-01 12:10:16
etag          : "f3016c-6bb-520e8d9f88e40"
vary_len      : 15
vary          : Accept-Encoding
variant(md5)  : 9eacc540a431495f5ae408412e60f000
key           : httpkroemeke.eu/
HEADERS       : 
HTTP/1.1 200 OK
Date: Thu, 01 Oct 2015 11:10:16 GMT
Server: Apache
Last-Modified: Tue, 29 Sep 2015 20:42:25 GMT
ETag: "f3016c-6bb-520e8d9f88e40"
Accept-Ranges: bytes
Content-Length: 1723
Vary: Accept-Encoding
Cache-Control: proxy-revalidate
Connection: close
Content-Type: text/html
```

less filename 可以查看
```
KEY: xx.com/xx//bytes=94371840-94633983/0
HTTP/1.1 206 Partial Content
Cache-Control: max-age=2592000
Connection: keep-alive
Content-Length: 262144
Content-Range: bytes 94371840-94633983/250537903
Content-Type: video/mp4
Date: Sun, 17 Nov 2024 15:34:21 GMT
Last-Modified: Sun, 17 Nov 2024 14:48:20 GMT
Server: openresty/1.19.1
```

```
package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"unsafe"
)

// 定义与C结构体对应的Go结构体
type NgxHttpFileCacheHeaderT struct {
	Version      uint64
	ValidSec     uint64
	UpdatingSec  uint64
	ErrorSec     uint64
	LastModified uint64
	Date         uint64
	Crc32        uint32
	ValidMsec    uint16
	HeaderStart  uint16
	BodyStart    uint16
	EtagLen      uint8
	Etag         [NGX_HTTP_CACHE_ETAG_LEN]byte
	VaryLen      uint8
	Vary         [NGX_HTTP_CACHE_VARY_LEN]byte
	Variant      [NGX_HTTP_CACHE_KEY_LEN]byte
}

const (
	NGX_HTTP_CACHE_ETAG_LEN = 128
	NGX_HTTP_CACHE_VARY_LEN = 128
	NGX_HTTP_CACHE_KEY_LEN  = 16
)
const LF byte = 10

var key = []byte{LF, 'K', 'E', 'Y', ':', ' '}

func main() {
	file, err := os.Open("/proxy_cache/5000d8850218e7862fdd631725636595")
	if err != nil {
		return
	}
	defer file.Close()

	// 读取文件开头的结构体部分
	header := NgxHttpFileCacheHeaderT{}
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil && err != io.EOF {
		return
	}

	fmt.Println(unsafe.Sizeof(header))
	fmt.Println(uintptr(len(key)))
	fmt.Println(unsafe.Sizeof(header) + uintptr(len(key)))

	end := int64(header.HeaderStart)
	start := int64(337)
	file.Seek(int64(start), 0)
	keySize := end - start
	keyData := make([]byte, keySize)
	file.Read(keyData)
	fmt.Printf("Key: %s\n", string(keyData[:keySize]))
}

```


## 在线进行配置

### nginx-ui/WebUI
https://github.com/0xJacky/nginx-ui
### nginxconfig
https://nginxconfig.io/
https://www.digitalocean.com/community/tools/nginx?global.app.lang=zhCN
https://github.com/digitalocean/nginxconfig.io

https://github.com/digitalocean/nginxconfig.io

### Nginx Proxy Manager
https://github.com/NginxProxyManager/nginx-proxy-manager
### playground
曾几何时，playground 似乎成了新语言的标配：Go 发布就带有 https://play.golang.org/，Rust 发布也有 https://play.rust-lang.org/。你想过 Nginx 也有一个 playground 吗？你可以通过它方便的测试 Nginx 配置。

https://nginx-playground.wizardzines.com/

后端的完整代码见这里：https://gist.github.com/jvns/edf78e7775fea8888685a9a2956bc477。

这个网站的作者写了一篇文章介绍它，包括安全问题、性能问题等，有兴趣的可以查看：https://jvns.ca/blog/2021/09/24/new-tool--an-nginx-playground/。另外，还有一个 Nginx location match 测试的网址：https://nginx.viraptor.info/。

[Nginx 竟然也有 playground：还是 Go 语言构建的](https://mp.weixin.qq.com/s/meBX4v3B2XusLVxwdicfeA)

### 其它

[nginxWebUI nginx网页配置工具](https://gitee.com/cym1102/nginxWebUI)
[Nginx Formatter / Nginx 格式化工具](https://github.com/soulteary/nginx-formatter) 一款 10MB 左右的，小巧、简洁的 Nginx 格式化工具，支持命令行、WebUI、Docker、x86、ARM、macOS、Linux。

#### Nginx与其它语言结合
Nginx和java: [Nginx-Clojure](https://github.com/nginx-clojure/nginx-clojure)
Nginx和lua: [Nginx-lua](https://github.com/openresty/lua-nginx-module)
JavaScript: nginScript
PHP: https://github.com/rryqszq4/ngx_php


> 第三方模块: https://www.nginx.com/resources/wiki/modules/

#### httpsok/SSL证书自动续期
[httpsok](https://github.com/httpsok/httpsok) 是一个专为 Nginx、OpenResty 服务器设计的HTTPS证书自动续签工具。它旨在帮助中小企业稳定、安全、可靠地管理SSL证书，让证书续期变得简单高效。


### Nginx管理系统
https://github.com/0xJacky/nginx-ui
https://github.com/alexazhou/VeryNginx
https://github.com/stefanpejcic/OpenPanel
https://github.com/ajenti/ajenti
https://github.com/schenkd/nginx-ui
https://github.com/EasyEngine/easyengine
https://github.com/caprover/caprover
https://github.com/nginx/agent


### nginx替代品
- [Caddy](https://caddyserver.com/)采用Go语言编写, 是一款功能强大，扩展性高的Web服务器(coredns就是使用它)
- [Tengine](http://tengine.taobao.org/) 阿里CDN使用的它
- [Traefik](https://doc.traefik.io/traefik/) 云原生
- [Pingora](https://github.com/cloudflare/pingora) 是Cloudflare使用Rust开发

## docker 安装nginx

docker run -d -p 8080:80 --name nginx  nginx:1.17.4
拷贝:
docker cp b3735bdb389a2:/etc/nginx/nginx.conf d:/dockerv/nginx/conf
也可以下载 https://github.com/nginx/nginx/blob/master/conf/nginx.conf

```
docker pull nginx:1.17.4

docker run -d -p 8080:80 --name nginx -v /d/dockerv/nginx/www:/usr/share/nginx/html -v /d/dockerv/nginx/conf/nginx.conf:/etc/nginx/nginx.conf -v /d/dockerv/nginx/logs:/var/log/nginx --restart always nginx:1.17.4

# nginx:1.19.3-alpine
```
[Nginx多域名配置](https://www.cnblogs.com/goloving/p/9363490.html)

### 定义状态页面
```
location /status {
                stub_status on;
                 #allow IP地址;
                 #deny IP地址;# deny all;
        }
```
访问/status
```
Active connections：当前活动的连接数量。
Accepts：已经接受客户端的连接总数量。
Handled：已经处理客户端的连接总数量。（一般与accepts一致，除非服务器限制了连接数量）。
Requests：客户端发送的请求数量。
Reading：当前服务器正在读取客户端请求头的数量。
Writing：当前服务器正在写响应信息的数量。
Waiting：当前多少客户端在等待服务器的响应。


Active connections：145           
#nginx 正处理的活动连接数145个。
server accepts handled requests
 1749 1749 3198                   
#nginx启动到现在共处理了 1749个连接 ,nginx启动到现在共成功创建 1749次握手 请求丢失数=(握手-连接),可以看出，我们没丢请求;总共处理了3198 次请求。
Reading: 0 Writing: 3 Waiting: 142
#Reading ：nginx读取到客户端的Header信息数。
#Writing ： nginx返回给客户端的Header信息数。
#Waiting ： Nginx已经处理完正在等候下一次请求指令的驻留连接.开启keep-alive的情况下,这个值等于active–(reading+writing)。
```

### 定义对静态页面的缓存时间
```
location ~* \.(jpg|jpeg|gif|png|css|js|ico|xml)$ {
  expires        30d;            //定义客户端缓存时间为30天
}
```

### nginx 绑定cpu到worker进程
worker_processes 4;
worker_cpu_affinity 0001 0010 0100 1000;

四核心系统，每个cpu绑定到一个work进程的写法。 

worker_processes 8;
worker_cpu_affinity 00000001 00000010 00000100 00001000 00010000 00100000 01000000 10000000;

八核心系统，每个cpu绑定到一个work进程的写法。
有多少个核，就有几位数，1表示该内核开启，0表示该内核关闭。

worker_processes  auto;
worker_cpu_affinity   auto;

### nginx最大并发连接数
[worker_connections](http://nginx.org/en/docs/ngx_core_module.html#worker_connections)

连接数包含与代理服务器的连接以及与客户端的连接, 同时总连接数(worker_connections*worker_process)不能超过最大打开文件数,
查看最大文件数`ulimit -a|grep "open files"` 或者 `ulimit -n`

```
客户端正向代理数 =   worker_rlimit_nofile / 2 >=  worker_connections * worker_process /2
客户端可连反向代理数据 =  worker_rlimit_nofile /4  >=  worker_connections * worker_process /4

nginx作为http服务器的时候：
max_clients = worker_processes * worker_connections/2

nginx作为反向代理服务器的时候：
max_clients = worker_processes * worker_connections/4
```

> 修改linux的最大打开文件数`/etc/security/limits.conf 增加到65535` ;同时修改`worker_rlimit_nofile ` 和 `worker_connections  `

**multi_accept**指令使得NGINX worker能够在获得新连接的通知时尽可能多的接受连接。 此指令的作用是立即接受所有连接放到监听队列中。 如果指令被禁用(	
Default: multi_accept off;)，worker进程将逐个接受连接。
```
events{
      multi_accept on;
}


如果use kqueue;, 则multi_accept无用
events{
    use kqueue;
    multi_accept on;
}
```

### 关闭 nginx

进入 nginx 下的 sbin 目录下执行命令：
快速停止 nginx 命令：./nginx -s stop
完整有序的停止 nginx：./nginx -s quit；这个命令会等待所有请求结束后再关闭 nginx。

使用 kill 命令关闭 nginx

先查看 nginx 进程，找到主进程号。然后再使用 kill 命令杀死 nginx 进程。
查看 nginx 进程号命令：ps -ef | grep nginx
从容停止 nginx 命令：kill -QUIT 主进程号
快速停止 nginx 命令：kill -TERM 主进程号
强制停止 nginx 命令：kill -HUP 主进程号
平滑重启 nginx 命令：kill -9 nginx
查找并杀死所有 nginx 进程：ps aux | grep nginx |awk '{print $2}' | xargs kill -9


### Nginx 性能优化

[如何构建高性能服务器（以Nginx为例）](https://www.cnblogs.com/kukafeiso/p/13957174.html)
[Nginx 性能优化有这篇就够了](https://www.cnblogs.com/cheyunhua/p/10670070.html)

#### Nginx性能优化 
workerconnections调大到65535，配合系统ulimit调整文件描述符，让Nginx能同时处理更多连接。
```
// nginx
worker_processes auto; # 自动匹配CPU核心数（推荐！）
worker_connections 65535; # 单进程最大连接数
worker_rlimit_nofile 100000; # 突破文件描述符限制
```
2025年Nginx 1.27.4版本推出SSL上下文复用技术，解决了证书重复加载的问题，直接把加载时间砍了82%！
```
// nginx
ssl_protocols TLSv1.2 TLSv1.3; # 只保留安全协议
ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256; # 现代密码套件
ssl_session_cache shared:SSL:10m; # 会话复用（减少重复握手）
ssl_stapling on; # OCSP stapling加速30%
```
静态资源压缩/静态资源缓存
```
// nginx
gzip on;
gzip_comp_level 6; # 压缩级别6（性能平衡点）
gzip_types text/plain text/css application/json application/javascript image/svg+xml;
gzip_min_length 1k; # 小文件不压缩（避免浪费CPU）
gzip_vary on; # 支持CDN缓存

// nginx
location ~* \.(jpg|css|js)$ {
proxy_cache my_cache; # 定义缓存区域
proxy_cache_valid 200 302 7d; # 缓存7天
proxy_cache_lock on; # 防止缓存惊群（同时大量请求回源）
expires 30d; # 浏览器缓存30天
access_log off; # 关闭日志减少IO
}
```

零拷贝传输
```
// nginx
sendfile on; # 启用零拷贝
tcp_nopush on; # 合并小包（减少网络包数量）
tcp_nodelay on; # 实时场景禁用Nagle算法（立即发小包）
```

负载均衡
```
// nginx
upstream backend {
least_conn; # 最少连接优先（谁闲给谁发）
server srv1 weight=3; # 性能好的多干活（权重3）
server srv2 weight=2; # 中等性能（权重2）
server srv3 backup; # 热备节点（平时不用，故障时顶上）
}
```
> 注意：加health_check模块（Nginx Plus）可主动检测后端健康状态，更靠谱！

内核参数调优`/etc/sysctl.conf`
```
// bash
net.core.somaxconn = 65535 # 最大待处理连接数（默认128）
net.ipv4.tcp_tw_reuse = 1 # 复用TIME-WAIT连接（减少资源浪费）
fs.file-max = 2097152 # 系统最大文件句柄数（默认10万）
```
> 改完执行sysctl -p生效，同时在/etc/security/limits.conf调大用户文件描述符：
```
* soft nofile 65535
* hard nofile 65535
```

```
#禁用ipv6
net.ipv6.conf.all.disable_ipv6 = 0
net.ipv6.conf.default.disable_ipv6 = 0
net.ipv6.conf.lo.disable_ipv6 = 0

#本地客户端端口范围
net.ipv4.ip_local_port_range = 1024 65000 

#当内存使用率大于97%时,开始使用swap
vm.swappiness = 2  

#每一个端口最大的listen队列的长度
net.core.somaxconn=655350  


#在高延迟的连接中，SACK 对于有效利用所有可用带宽尤其重要。高延迟会导致在任何给定时刻都有大量正在传送的包在等待应答。在 Linux 中，除非得到应答或不再需要，这些包将一直存放在重传队列中。这些包按照序列编号排队，但不存在任何形式的索引。当需要处理一个收到的 SACK 选项时，TCP 协议栈必须在重传队列中找到应用了 SACK 的包。重传队列越长，找到所需的数据就越困难。一般可关闭这个功能。选择性应答在高带宽延迟的网络连接上对性能的影响很大，但也可将其禁用，这不会牺牲互操作性。将其值设置为 0 即可禁用 TCP 协议栈中的 SACK 功能
net.ipv4.tcp_sack = 1  

##每个网络接口接收数据包的速率比内核处理这些包的速率快时,允许送到队列的数据包的最大数目 
net.core.netdev_max_backlog = 262144

##TCP窗口扩大因子支持. 如果TCP窗口最大超过65535(64K), 设置该数值为1 。Tcp窗口扩大因子是一个新选项，一些新的实现才会包含该选项，为了是新旧协议兼容，做了如下约定：1、只有主动连接方的第一个syn可以发送窗口扩大因子；2、被动连接方接收到带有窗口扩大因子的选项后，如果支持，则可以发送自己的窗口扩大因子，否则忽略该选项；3、如果双方支持该选项，那么后续的数据传输则使用该窗口扩大因子。如果对方不支持wscale，那么它不应该响应 wscale 0，而且在收到46的窗口时不应该发送1460的数据；如果对方支持wscale，那么它应该大量发送数据来增加吞吐量，不至于通过关闭wscale来解决问题，如果是使用普遍的协议实现，那么就需要关闭wscale来提高性能并以防万一。
net.ipv4.tcp_window_scaling = 1

#TCP读buffer 
net.ipv4.tcp_rmem = 53687091 53687091 536870912

#TCP写buffer
net.ipv4.tcp_wmem = 53687091 53687091 536870912

#系统中最多有多少个TCP套接字不被关联到任何一个用户文件句柄上。如果超过这个数字，孤儿连接将即刻被复位并打印出警告信息。这个限制仅仅是为了防止简单的DoS攻击，不能过分依靠它或者人为地减小这个值，更应该增加这个值(如果增加了内存之后)。
net.ipv4.tcp_max_orphans = 3276800

#意思是如果某个TCP连接在idle 2分钟后,内核才发起probe.如果probe 1次(每次2秒)不成功,内核才彻底放弃,认为该连接已失效. 
net.ipv4.tcp_keepalive_intvl = 15
net.ipv4.tcp_keepalive_probes = 3

#表示如果套接字由本端要求关闭，这个参数决定了它保持在FIN-WAIT-2状态的时间
net.ipv4.tcp_fin_timeout = 15

#表示当keepalive起用的时候，TCP发送keepalive消息的频度。缺省是2小时，改为1分钟
net.ipv4.tcp_keepalive_time = 600

#1st低于此值,TCP没有内存压力,2nd进入内存压力阶段,3rdTCP拒绝分配socket(单位：内存页)
net.ipv4.tcp_mem = 94500000 915000000 927000000

#开启重用。允许将TIME-WAIT sockets重新用于新的TCP连接。
net.ipv4.tcp_tw_reuse = 1

#时间戳可以避免序列号的卷绕。一个1Gbps 的链路肯定会遇到以前用过的序列号。时间戳能够让内核接受这种“异常”的数据包。这里需要将其关掉。
net.ipv4.tcp_timestamps = 0

#为了打开对端的连接，内核需要发送一个SYN 并附带一个回应前面一个SYN 的ACK。也就是所谓三次握手中的第二次握手。这个设置决定了内核放弃连接之前发送SYN+ACK 包的数量。
net.ipv4.tcp_synack_retries = 1

#对于一个新建连接，内核要发送多少个 SYN 连接请求才决定放弃。不应该大于255，默认值是5
net.ipv4.tcp_syn_retries = 3


# System default settings live in /usr/lib/sysctl.d/00-system.conf.
# To override those settings, enter new settings here, or in an /etc/sysctl.d/<name>.conf file
#
# For more information, see sysctl.conf(5) and sysctl.d(5).
kernel.msgmnb = 65536
kernel.msgmax = 65536
kernel.pid_max = 4194303
kernel.shmmax = 68719476736
kernel.shmall = 4294967296
kernel.sysrq = 0
kernel.core_uses_pid = 1
#增加watchdog 软死锁的超时时间,https://www.cnblogs.com/xingyunfashi/p/10710784.html
kernel.watchdog_thresh = 10
net.core.rmem_max = 536870912
net.core.wmem_max = 536870912
net.core.rmem_default = 53687091
net.core.wmem_default = 53687091
net.core.optmem_max = 536870912

net.nf_conntrack_max = 600000000
net.netfilter.nf_conntrack_tcp_timeout_established = 1200

net.ipv4.ip_forward = 1
net.ipv4.conf.default.rp_filter = 2
net.ipv4.conf.default.accept_source_route = 0
net.ipv4.tcp_syncookies = 1
net.ipv4.tcp_max_tw_buckets = 36000

net.ipv4.tcp_retries2 = 5
net.ipv4.tcp_max_syn_backlog = 262144

#回收内存的阈值
#vm.min_free_kbytes = 1048576

fs.file-max = 13181532
fs.inotify.max_queued_events = 99999999
fs.inotify.max_user_watches = 99999999
fs.inotify.max_user_instances = 65535

#net.core.default_qdisc=fq
net.ipv4.tcp_congestion_control=bbr
net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-iptables=1
```

stub_status模块：
```
// nginx
location /nginx_status {
stub_status on;
allow 127.0.0.1; # 只允许本地访问
deny all;
}
```

### OCSP Stapling
目的：服务器定期从 CA 获取 OCSP（在线证书状态协议）响应，缓存后在 TLS 握手时直接发送给客户端，避免客户端单独查询 CA。

• 优势：
• 降低延迟：减少客户端验证证书吊销状态的往返时间。
• 减轻 CA 负载：避免大量客户端直接查询 CA 服务器。
• 增强隐私：隐藏客户端与 CA 的通信细节。

```
server {
    listen 443 ssl http2;
    server_name www.topssl.cn;
 
    # 1. 证书与私钥配置
    ssl_certificate /path/to/fullchain.pem; # 必须确立是包含中间证书的完整链
    ssl_certificate_key /path/to/privkey.pem;
 
    # 2. 开启 OCSP Stapling
    ssl_stapling on;
    ssl_stapling_verify on;
 
    # 3. 指定根证书（用于验证 CA 的响应）# 指定 CA 证书路径（用于验证 OCSP 响应）
    # 很多老兄漏掉这一步，确立了 Stapling 无法在服务端通过验证
    ssl_trusted_certificate /path/to/root_and_intermediate.pem;
 
    # 4. 优化 DNS 解析（重要！）
    # 确立 Nginx 能快速找到 CA 的 OCSP 地址，建议用公共 DNS
    resolver 8.8.8.8 1.1.1.1 valid=300s;
    resolver_timeout 5s;
 

    # 缓存 OCSP 响应（路径和超时时间）
    ssl_stapling_file /var/lib/nginx/ocsp/example.com.ocsp;
    ssl_stapling_cache_timeout 86400;  # 缓存 24 小时
    # 其他安全配置...
}
```


验证:
1. https://www.topssl.cn/tools/check
2. 使用 OpenSSL 命令检查
`openssl s_client -connect example.com:443 -status -tlsextdebug < /dev/null 2>&1 | grep -i "OCSP response"`
成功标志：输出包含 OCSP Response Status: successful

3. 通过浏览器检查
 访问 https://example.com，点击地址栏锁图标 → 查看证书详细信息 → 检查 OCSP 字段是否显示 Response received。

4. 使用在线工具（如 crt.sh）
• 输入域名查询证书信息，检查 OCSP 状态是否为 Good。

### ngx_http_auth_digest 模块进行登录认证

生成登录密码：

apt-get install -y apache2-utils
htdigest -c /usr/local/passwd.digest theia admin

- /usr/local/passwd.digest 为密码文件路径
- theia 为 realm(领域，范围)，必须与后面要提到的 nginx 配置文件保持一致。
- admin 为登录用户名

在 nginx 配置文件中添加 server 段：
```
server {
	listen 80 default_server;

    auth_digest_user_file /usr/local/passwd.digest;
    auth_digest_shm_size  8m;   # the storage space allocated for tracking active sessions
    
    location / {
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";

        auth_digest 'theia';
        auth_digest_timeout 60s;    # allow users to wait 1 minute between receiving the
                                    # challenge and hitting send in the browser dialog box
        auth_digest_expires 600s;   # after a successful challenge/response, let the client
                                    # continue to use the same nonce for additional requests
                                    # for 600 seconds before generating a new challenge
        auth_digest_replays 60;     # also generate a new challenge if the client uses the
                                    # same nonce more than 60 times before the expire time limit

        proxy_pass http://127.0.0.1:3000;
    }
}
```

- auth_digest_user_file 必须与设置密码的 htdigest 命令中的密码文件参数保持一致。
- auth_digest 必须与设置密码的 htdigest 命令的 realm 参数保持一致。
- 这两行必须有，否则不能正确代理 websocket。
proxy_set_header Upgrade $http_upgrade; 
proxy_set_header Connection "upgrade";
   
这个配置相当于每 (auth_digest_timeout+auth_digest_expires)=660s 允许 auth_digest_shm_size/((48 + ceil(auth_digest_replays/8))bytes)=(8*1024*1024)/(48+8)=149.8k 次请求，即每 660s 允许约 149.8k 次请求。登录认证 10min 后过期。

最后启动 nginx，会弹出登录认证框，输入用户名和密码后即可登录，跳转到 theia 界面。


### 设置HTTP强制跳转HTTPS
```
server {
    listen 80;
    server_name example.com;  这里修改为网站域名
    rewrite ^(.*)$ https://$host$1 permanent;
}
```

[高性能 Nginx HTTPS 调优 - 如何为 HTTPS 提速 30%](https://www.toutiao.com/article/7197285537810514435/)

### nginx负载均衡策略简介

1、轮询（默认策略，nginx自带策略）：
它是upstream模块默认的负载均衡默认策略。会将每个请求按时间顺序分配到不同的后端服务器。
```
http {
    upstream my_load_balance {
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}
```

2、weight（权重，nginx自带策略）：
指定轮询的访问几率，用于后端服务器性能不均时调整访问比例。权重越高，被分配的次数越多。
```
http {
    upstream my_load_balance {
        server 192.168.1.12:80 weight=7;
        server 192.168.1.13:80 weight=2;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}　
```

3、ip_hash（依据ip分配，nginx自带策略）：
指定负载均衡器按照基于客户端IP的分配方式，这个方法确保了相同的客户端的请求一直发送到相同的服务器，可以解决session不能跨服务器的问题。
```
http {
    upstream my_load_balance {
        ip_hash;
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}
```

4、least_conn（最少连接，nginx自带策略）：
把请求转发给连接数较少的后端服务器。
```
http {
    upstream my_load_balance {
        #把请求转发给连接数比较少的服务器
        least_conn;
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}   
```

5、fair（第三方）：
按照服务器端的响应时间来分配请求，响应时间短的优先分配。
```
http {
    upstream my_load_balance {
        fair;
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
}  
```
6、url_hash（第三方）：
该策略按访问url的hash结果来分配请求，使每个url定向到同一个后端服务器，需要配合缓存用。
```
http {
    upstream my_load_balance {
        hash $request_uri;
        server 192.168.1.12:80;
        server 192.168.1.13:80;
    }
  
    server {
        listen 81;
        server_name www.laowubiji.com;
  
        location / {
            proxy_pass http://my_load_balance;
            proxy_set_header Host $proxy_host;
        }
    }
} 
```

## k8s
kubectl create configmap confnginx --from-file nginx.conf
```
user  nginx;
worker_processes  1;
 
error_log  /var/log/nginx/error.log warn;
pid        /var/run/nginx.pid;
 
events {
    worker_connections  1024;
}
 
http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;
 
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                        '$status $body_bytes_sent "$http_referer" '
                        '"$http_user_agent" "$http_x_forwarded_for"';
 
    access_log  /var/log/nginx/access.log  main;
 
    sendfile        on;
    #tcp_nopush     on;
 
    keepalive_timeout  65;
 
    #gzip  on;
 
    #include /etc/nginx/conf.d/*.conf;
 
    server {
       listen                    80;
       server_name               localhost;
       root                      /home/wwwroot/html;
       index                     index.html;
    }
}
```

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx-deployment
spec:
  replicas: 1
  selector:
    name: nginx
  template:
    metadata:
      labels:
        name: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.19.3-alpine
        ports:
        - containerPort: 80
        volumeMounts:
        - mountPath: /etc/nginx/nginx.conf
          name: nginx-config
          subPath: nginx.conf
        - mountPath: /home/wwwroot/html
          name: nginx-data
      volumes:
        - name: nginx-config
          configMap:
            name: confnginx
        - name: nginx-data
          hostPath:
            path: /home/wwwroot/hello

---

apiVersion: v1
kind: Service
metadata:
  name: nginx-service
spec:
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
      nodePort: 30080    #外网访问端口
  type: NodePort   #这个是端口类型
  selector:
    name: nginx
	
--- 

apiVersion: extensions/v1beta1
kind: Ingress
metadata:
    name: nginx-ingress
    annotations:
        kubernetes.io/ingress.class: "nginx"
spec:
    rules:
    - host: violet.chtwebapi.lingcb1.com
      http:
          paths:
          - path: /plan
            backend:
                serviceName: nginx-service
                servicePort: 80
          

```


## 四层代理

```
stream{
    upstream backend_coredns_udp {
        server xxxx:53;
        server xxx:53 backup;
    }

    server {
        listen 53 udp;
        listen 53; #tcp
        proxy_pass backend_coredns_udp;
        # proxy_protocol      on;
    }

    upstream dns {
       server 192.168.0.1:53535;
       server dns.example.com:53;
    }

    server {
        listen 127.0.0.1:53 udp reuseport;
        proxy_timeout 20s;
        proxy_pass dns;
    }
}




```

## 同一location块中同时直接使用proxy_pass和root指令

1. 使用try_files指令
```
server {
    listen 80;
    server_name example.com;

    location / {
        root /var/www/html;
        try_files $uri @backend;
    }

    location @backend {
        proxy_pass http://backend_server;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```
- 首先尝试使用root指令指定的目录来响应请求（/var/www/html）。try_files $uri @backend;语句会检查请求的文件是否存在，如果存在则直接提供，不存在则转到命名的location @backend。
- @backend是一个命名的location，用于将请求代理到后端服务器。

2. 基于请求URI或参数的条件判断
```
location / {
    if ($request_uri ~* "^/static/") {
        root /var/www/static;
    }
    # 或者使用嵌套location
    location /api/ {
        proxy_pass http://backend_api;
    }
}
```

3. 使用命名location
```
location / {
    root /var/www/html;
    error_page 404 = @proxy;
}

location @proxy {
    proxy_pass http://backend;
}
```
