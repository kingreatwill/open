
## EDNS Client Subnet(ECS) 协议简介

DNS系统默认使用明文UDP协议通信，所以用户的查询内容很容易受到监控，而服务器返回的解析结果是可以被轻易篡改。为了解决这个问题，人们引入了 DNS over HTTPS/TLS/QUIC 之类的技术，希望通过加密的方式传输DNS查询。但使用公共的 DoH 递归解析服务器后，权威DNS服务器只能获取递归解析服务器的地址，但无法获取用户的地址。所以于地理位置的智能解析就无法工作了。为了解决这个问题，人们制定了 EDNS Client Subnet (ECS) 协议，也就是[RFC7871](https://datatracker.ietf.org/doc/html/rfc7871)。本文就简单介绍一下 ECS 的工作原理和使用效果。

我们日常都使用运营商的递归服务器，它们跟用户的机器地理距离都很近，就不会产生大问题。但如果中国的用户使用了美国的 DNS over HTTPS 服务，那解析出来的可能是美国的IP，会严重影响用户访问。

我之前并不知道 ECS 协议。那时候只能采用域名白名单来解决这个问题。简单来说，我们在路由器上配置 dnsmasq 做递归解析。Dnsmasq 支持根据域名前缀设置不同的 DNS 解析服务器，比如下面是一条针对苹果的配置：
```
server=/apps.apple.com/114.114.114.114
```

dnsmasq 会从`114.114.114.114`查询域名`apps.apple.com`的DNS记录，所以就能拿到苹果在国内的CDN节点地址，从而实现「加速」效果。

为此，网友维护[dnsmasq-china-list](https://github.com/felixonmars/dnsmasq-china-list)项目，基本把国内域名以及谷歌和苹果的的国内域名都加上去了。

这种白名单只能说是笨办法。有了 ECS 我们就不需要维护这么复杂的白名单。

ECS 简单说就是把用户的IP信息暴露给权威DNS服务器。但为了保护用户隐私，递归服务器并不直接把用户的IP发给权威服务器。相反，只把用户IP所在的网段发给权威DNS服务器。如果客户使用 IPv4，那发送的网络前缀为 24，如果是 IP6，那网络前缀为56。一般来说，同网段的客户地址位置相近。这里的网段就叫 client subnet。

ECS 具体是怎么发送 client subnet 的本文就不细说了，大家可以阅读 RFC7871。接下来我们以`www.qq.com`为例说一下使用效果。

在使用之前，请确保可以直连 Google Public DNS (8.8.8.8)。

我用电信的网络查询`www.qq.com`的A记录：
```
dig www.qq.com +short

ins-r23tsuuf.ias.tencent-cloud.net.
101.91.22.57
101.91.42.232
```

这是腾讯在国内的节点地址。如果我去海外的 VPS 执行同样的查询：
```
dig +short www.qq.com @8.8.8.8

news.qq.com.edgekey.net.
e6156.dscf.akamaiedge.net.
104.94.212.210
```

这里返回的是 akamai 在美国的地址。如果我加上了 ECS 信息：
```
dig www.qq.com +short +subnet=x.x.x.0/24 @8.8.8.8

ins-r23tsuuf.ias.tencent-cloud.net.
101.91.42.232
101.91.22.57
```

我们看到虽然是在国外的 VPS，但还是返回了国内的节点。也就是说，不管你在那里执行查询，只要 subnet 不变，返回的一定是离你比较近的节点。

如果是用 dnsmasq 执行 DNS 解析，需要添加如下配置：
```
add-subnet=24,56
server=8.8.8.8
```

ECS 有一个小缺点，就是不同 subnet 的 DNS 缓存不能共用。这也容易理解，因为不同的网段可能对应不同的地理位置，解析结果可能不同，自然不能共享缓存。这虽然会降低缓存的命中率，但对于 Google Public DNS 这种用量很大的场景来说不是什么大问题。Google Public DNS 对 ECS 的支持情况可以参考这个[链接](https://groups.google.com/g/public-dns-announce/c/h4XLjnWvAp8?pli=1)。

有了 ECS，我们在国内只需要设置IP隧道和IP分流，再也不需要搞DNS解析分流了。

以上就是本文的全部内容，欢迎留言讨论。

参考文章：

https://ephen.me/2017/PublicDns_2/