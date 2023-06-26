# CDN
内容分发网络CDN（Content Delivery Network）由遍布全球的边缘节点服务器群组成的分布式网络。
CDN将源站资源缓存到加速节点，当终端用户请求访问和获取源站资源时无需回源，可就近获取CDN节点上已经缓存的资源，提高资源访问速度，同时分担源站压力。

[阿里云CDN-产品文档](https://help.aliyun.com/product/27099.html)

[阿里云CDN-管理控制台](https://cdn.console.aliyun.com/overview)

[CDN技术架构.pdf](https://developer.aliyun.com/ebook/7561) 60页开始介绍调度

[CDN技术详解.pdf](链接: https://pan.baidu.com/s/1i4UIOcd 密码: 2yjh)

抓包命令: `tcpdump -i any -vvvvnnA dst port 8899`

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
腾讯：`119.29.29.29`
阿里：`223.5.5.5` 、`223.6.6.6`
谷歌：`8.8.8.8`，
国内114dns：`114.114.114.114`

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


