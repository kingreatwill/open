# CDN
内容分发网络CDN（Content Delivery Network）由遍布全球的边缘节点服务器群组成的分布式网络。
CDN将源站资源缓存到加速节点，当终端用户请求访问和获取源站资源时无需回源，可就近获取CDN节点上已经缓存的资源，提高资源访问速度，同时分担源站压力。

[阿里云CDN-产品文档](https://help.aliyun.com/product/27099.html)

[阿里云CDN-管理控制台](https://cdn.console.aliyun.com/overview)

[CDN技术架构.pdf](https://developer.aliyun.com/ebook/7561) 60页开始介绍调度

[CDN技术详解.pdf](链接: https://pan.baidu.com/s/1i4UIOcd 密码: 2yjh)

抓包命令: `tcpdump -i any -vvvvnnA dst port 8899`
可以`tcpdump -i any -vvvvnnA dst port 8899 -w file.cap` 保存文件中, 然后使用wireshark来查看数据包

## CDN架构
Nginx=>Varnish=>ATS的架构后，稳定性大大的加强了。

Nginx通过hash解决了cache抖动的问题，Varnish的内存缓存非常的强大，可定制性太强了，很方便设置content-length不等于0时，则不缓存对象。

[varnish / squid / nginx cache 有什么不同](https://www.zhihu.com/question/20143441)

## 调度算法

> go 调用python脚本 github.com/YuminosukeSato/pyproc
### MIP
https://www.gurobi.com/resources/mixed-integer-programming-mip-a-primer-on-the-basics/

线性规划: https://scipopt.org/doc/html/
[SCIP](https://scipopt.org/doc/html/)

[COIN-OR CBC](https://github.com/coin-or/Cbc)（Coin-or Branch and Cut）是一个开源的混合整数线性规划（MILP）求解器

Google OR-Tools专门用于解决各种优化问题，包括MIP（混合整数规划）
https://github.com/google/or-tools.git
https://github.com/google/or-tools/tree/stable/ortools/linear_solver


https://github.com/gonzojive/or-tools-go
absl and or-tools:
https://github.com/abseil/abseil-cpp.git
https://github.com/google/or-tools.git

其他MIP求解器

开源的
[GLPK](https://www.gnu.org/software/glpk/)（仅限 Linux 和 MacOS）
[COIN-OR CBC](https://github.com/coin-or/Cbc)（Coin-or Branch and Cut）是一个开源的混合整数线性规划（MILP）求解器
[SCIP](https://scipopt.org/doc/html/)
[HiGHS](https://github.com/ERGO-Code/HiGHS)

商业的
Gurobi
[CPLEX](https://www.ibm.com/cn-zh/products/ilog-cplex-optimization-studio/cplex-optimizer)
[XPRESS](https://www.fico.com/en/products/fico-xpress-optimization) 求解器



求解速度（大致排序）
Gurobi - 最快
CPLEX - 很快
Xpress - 快
HiGHS - 较快
SCIP - 中等
CBC - 较慢
GLPK - 最慢

内存使用
GLPK - 最少
CBC - 较少
HiGHS - 中等
SCIP - 较多
Gurobi/CPLEX - 最多


### mcmf
最小费用最大流（Minimum Cost Maximum Flow）

## DNS

### 权威DNS、Local DNS、公共DNS有什么区别？

#### 权威 DNS
权威 DNS 指域名在域名注册商处所设置的 DNS 服务器地址。该地址决定了该域名的解析管理权（新增，删除，修改等）。比如 DNSPod 的权威服务器：`*.dnspod.net`, `*.dnsv3.com` 等。当域名设置权威服务器并设置了解析记录后，客户端请求域名时，权威服务器将返回该域名的对应的解析记录信息。

#### Local DNS
Local DNS 是 DNS 查询中的第一个节点。Local DNS 作为客户端与 DNS 域名服务器的中间人。客户端发送 DNS 查询时，Local DNS 将使用缓存的数据进行响应，或者将向根域名服务器发送请求，接着向根域名服务器发送另一个请求，然后向权威 DNS 发送最后一个请求。收到来自包含已请求 IP 地址的权威 DNS 服务器的响应后，Local DNS 将向客户端发送响应。

在此过程中，Local DNS 将缓存从权威 DNS 服务器收到的信息。当一个客户端请求的域名 IP 地址是另一个客户端最近请求的 IP 地址时，Local DNS 可绕过与域名服务器进行通信的过程，并仅从第二个客户端的缓存中为第一个客户端提供所请求的记录。

> Local DNS 也是和我们日常上网接触最多的DNS包括你的服务提供商（ISP）分配给你的DNS, 又因为填写在你的本地电脑上，所以也称为Local DNS

> Local DNS（LDNS）即本地域名服务器，由本地运营商提供。用户的解析请求，是由 LDNS 执行的。

#### ClientIP
ClientIP 指的是终端用户的 IP 地址，每个终端设备上都会配置 LDNS 的地址。在 DNS 调度中，只能识别 LDNS，识别不了用户的 IP 地址，所以用户流量是被 LDNS 所牵引的。

#### 公共 DNS
公共DNS，指面向所有互联网用户的全球公共递归域名解析服务。和仅使用本地 LocalDNS 的传统解析服务相比，公共解析服务，一般具备更加“快速”、“稳定”、“安全”互联网访问。

常用的公共dns
腾讯：`119.29.29.29`、`182.254.116.116` DNSPod/腾讯云
百度公共DNS: `180.76.76.76` IPv6 地址为 `2400:da00::6666`
阿里：`223.5.5.5` 、`223.6.6.6` IPv6 地址为 `2400:3200::1` 和 `2400:3200:baba::1`
谷歌：`8.8.8.8`，`8.8.4.4` IPv6 地址为 `2001:4860:4860::8888` / `2001:4860:4860::8844`。
国内114dns：`114.114.114.114`，`114.114.115.115` IPv6：`2001:dc7:1000::1`
天翼云（中国电信） DNS:`101.226.4.6`
OpenDNS ：`208.67.222.222` 和 `208.67.220.220`


> 纯净 无劫持 无需再忍受被强扭去看广告或粗俗网站之痛苦
服务ip为：114.114.114.114 和 114.114.115.115
拦截 钓鱼病毒木马网站 增强网银、证券、购物、游戏、隐私信息安全
服务ip为：114.114.114.119 和 114.114.115.119
学校或家长可选拦截 色情网站 保护少年儿童免受网络色情内容的毒害
服务ip为：114.114.114.110 和 114.114.115.110

### 调度系统
一般有哪些调度形式？

1. DNS调度   基于请求端local dns的出口IP归属地及运营商属性的DNS调度；
2. 302调度    再是基于客户端IP归属地及运营商属性的302跳转调度；
3. 路由调度   最后是基于Anycast技术（BGP路由）的机房流量调度；


参考: [CDN技术漫谈之调度系统](https://cloud.tencent.com/developer/article/1394677)

## CDN后台管理

### 域名管理

**加速原理:**
假设您的加速域名为`www.wcoder.com`，接入CDN开始加速服务后(要先将域名解析CNAME为`www.wcoder.com.example.com`)，当终端用户在北京发起HTTP请求时，处理流程如下图所示。

![CDN加速原理](CDN加速原理.drawio)

1. 当终端用户向`www.wcoder.com`下的指定资源发起请求时，首先向Local DNS（本地DNS）发起请求域名`www.wcoder.com`对应的IP。
2. Local DNS检查缓存中是否有`www.wcoder.com`的IP地址记录。如果有，则直接返回给终端用户；如果没有，则向网站授权DNS请求域名`www.wcoder.com`的解析记录。
3. 当网站授权DNS解析`www.wcoder.com`后，返回域名的CNAME `www.wcoder.com.example.com`。
4. Local DNS向阿里云CDN的DNS调度系统请求域名`www.wcoder.com.example.com`的解析记录，阿里云CDN的DNS调度系统将为其分配最佳节点IP地址。
5. Local DNS获取阿里云CDN的DNS调度系统返回的最佳节点IP地址。
6. Local DNS将最佳节点IP地址返回给用户，用户获取到最佳节点IP地址。
7. 用户向最佳节点IP地址发起对该资源的访问请求。
    - 如果该最佳节点已缓存该资源，则会将请求的资源直接返回给用户（步骤8），此时请求结束。
    - 如果该最佳节点未缓存该资源或者缓存的资源已经失效，则节点将会向源站发起对该资源的请求。获取源站资源后结合用户自定义配置的缓存策略，将资源缓存到CDN节点并返回给用户（步骤8），此时请求结束。

> 我们把example.com称为接入域名, 为什么需要接入层域名? 答: https证书是基于域名的; 
> 域名的DNS服务器一般在域名注册商处维护(是可以修改的, 当前域名的解析DNS服务器更换为其他的DNS服务器，由修改后的DNS服务器负责域名解析服务。)

## 开源系统
### Apache Traffic Server(ATS)
ATS最重要的是CDN缓存能力,反向代理是其第二功能.

[ATS CDN 搭建](https://www.taterli.com/8527/)

https://github.com/apache/trafficserver 1.8k

### Apache Traffic Control(ATC)
基于ATS实现完整的CDN
https://github.com/apache/trafficcontrol 1.1k

### varnish
https://github.com/varnishcache/varnish-cache 3.6k

https://varnish-cache.org/

### Squid
https://github.com/squid-cache/squid

## CDN相关文档

### cloudflare

[学习中心](https://www.cloudflare.com/zh-cn/learning/)

### 国家IPv6发展监测平台
https://www.china-ipv6.cn/#/

## 网络工具
https://networkingtoolbox.net/
https://github.com/Lissy93/networking-toolbox
