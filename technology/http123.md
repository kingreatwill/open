[HTTP/2与HTTP/3 的新特](https://www.toutiao.com/a6759344542853366279/)
[解密HTTP/2与HTTP/3 的新特性](https://juejin.im/post/5d9abde7e51d4578110dc77f)
[Hypertext Transfer Protocol Version 2 (HTTP/2)](http://http2.github.io/http2-spec/#FrameHeader)
[HTTP/2 Frequently Asked Questions](https://http2.github.io/faq/)
[Hypertext Transfer Protocol (HTTP/1.1): Message Syntax and Routing](https://httpwg.org/specs/rfc7230.html)
[Hypertext Transfer Protocol -- HTTP/1.1](https://datatracker.ietf.org/doc/html/rfc2616)
[Hypertext Transfer Protocol Version 2 (HTTP/2)](https://httpwg.org/specs/rfc7540.html)

[理解TCP/IP协议栈之HTTP2.0](https://www.cnblogs.com/backnullptr/p/12186265.html)
[HTTP1.0、HTTP1.1 和 HTTP2.0 的区别](https://www.cnblogs.com/heluan/p/8620312.html)
[HTTP1.0 HTTP 1.1 HTTP 2.0主要区别](https://blog.csdn.net/linsongbin1/article/details/54980801)

[HTTP3.0(QUIC的实现机制)](https://www.cnblogs.com/chenjinxinlove/p/10104854.html)
[QUIC协议浅析与HTTP/3.0](https://www.jianshu.com/p/bb3eeb36b479)

## 其它

### http状态码
> https://http.cat/ 查看各种状态码, [github](https://github.com/httpcats/http.cat)
> 同类站点:https://httpstatusdogs.com/ 参考https://github.com/timdouglas/httpstatusdogs-nginx , https://httbey.com/ 对应的github地址:https://github.com/minamarkham/httpbey

> [Hypertext Transfer Protocol -- HTTP/1.1](https://datatracker.ietf.org/doc/html/rfc2616#page-57)


401 Unauthorized(未登录)
403 Forbidden (remove root user when auth is enabled) (没权限,这个状态类似于 401，但进入 403状态后即使重新验证也不会改变该状态。该访问是长期禁止的，并且与应用逻辑密切相关（例如没有足够的权限访问该资源）。)


### keep-alive
Client 和 Server 的 keep-alive（心跳/保活）时间设置一致，可能会导致连接意外断开或误判。

为什么“时间一致”会有问题？

假设 client 和 server 都设置 keep-alive 时间为 60 秒：
- client：每 60 秒发一次心跳包。
- server：每 60 秒检测一次，如果 60 秒没收到数据就断开连接。

问题场景

- 如果 client 和 server 的定时器几乎同时启动，client 可能在第 60 秒刚发完心跳包，server 也刚好在第 60 秒检测超时。
- 由于网络延迟、调度等原因，server 可能先于 client 的心跳包到达前就检测超时，从而误判连接已断开，直接关闭连接。
- 这种“边界条件”下，偶发性断开就会发生。

正确做法
- client 的 keep-alive 时间要小于 server 的超时时间，比如 client 30 秒，server 60 秒。
- 这样 server 检测超时前，client 至少会发一次心跳包，server 能及时收到，避免误判。

> golang默认的http client的keep alive默认是30s
> http server是IdleTimeout, 说明如下:
> IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled. 
If zero, the value of ReadTimeout is used. If negative, or if zero and ReadTimeout is zero or negative, there is no timeout.
```
srv := &http.Server{  
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
}
log.Println(srv.ListenAndServe())
```

用 netcat (nc) 观察连接关闭时间
```
nc localhost 8080

不要关闭 nc，保持连接。
等待一段时间（比如 60 秒），如果服务端设置了 IdleTimeout: 60s，大约 60 秒后连接会被服务端关闭，nc 会自动退出。
```

服务端为什么会发 FIN 主动关闭连接？

即使开启了 keep-alive，服务端依然可能主动关闭连接，常见原因有：
1. 资源限制
服务端需要维护大量连接，为了防止资源耗尽，会设置最大空闲时间（timeout）。
如果连接在一段时间内没有新请求，服务端会主动关闭（发 FIN）。
2. 配置的 keep-alive 超时
服务端通常会有 keep-alive timeout 配置，比如 60 秒、30 秒等。
超时后，服务端会发 FIN 关闭连接。
3. 达到最大请求数
有些服务端会限制一个 keep-alive 连接最多处理多少个请求，达到后就关闭连接。
4. 服务器重启或负载均衡
服务端重启、部署、负载均衡切换等操作时，现有连接会被关闭。
5. 服务端主动释放资源
服务端压力大时，可能会主动关闭部分空闲连接，释放资源。


### 安装支持http3的curl
curl --http3 https://http3check.net/ -I
curl --http3 https://cloudflare-quic.com -I

brew install curl --with-c-ares
brew install curl-openssl

支持--dns-servers
brew install c-ares
brew install autoconf automake libtool pkg-config cmake


curl -LO https://curl.se/download/curl-8.8.0.tar.gz
tar xzf curl-8.8.0.tar.gz
cd curl-8.8.0
./configure --with-ares=$(brew --prefix c-ares)
make -j4
sudo make install

或者直接以下安装方式:
```
# Clean up any old version of curl you may have already tried to install
brew remove -f curl

# Download the curl ruby install script provided by cloudflare
wget https://raw.githubusercontent.com/cloudflare/homebrew-cloudflare/master/curl.rb

# Install curl via that script from the latest git repos
brew install --HEAD -s curl.rb

# Tell your cli to use the curl version just installed (if you're using zsh, othwerise you might need `~/.bashrc`)
# echo 'export PATH="/usr/local/opt/curl/bin:$PATH"' >> ~/.zshrc #老版本
# FYI, /opt/homebrew can be replaced with $(brew --prefix)
# echo 'export PATH="$(brew --prefix)/opt/curl/bin:$PATH"' >> ~/.zshrc #老版本

echo 'export PATH="/opt/homebrew/opt/curl/bin:$PATH"' >> ~/.zshrc
# Reload your config
source ~/.zshrc

# Double check it's using the right curl
which curl # Should output "/usr/local/opt/curl/bin/curl"

# Double check http3
$ curl --version | grep HTTP3
  Features: alt-svc AsynchDNS brotli HTTP2 HTTP3 IDN IPv6 Largefile libz MultiSSL NTLM NTLM_WB SSL UnixSockets zstd

# Try curl on any HTTP/3 enabled sites.
curl --http3 https://blog.cloudflare.com -I
```

也可以直接:

```
brew remove -f curl
brew install cloudflare/homebrew-cloudflare/curl
echo 'export PATH="/usr/local/opt/curl/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

# mac M2
brew remove -f curl
brew install cloudflare/homebrew-cloudflare/curl
echo 'export PATH="/opt/homebrew/opt/curl/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```



这里安装带c-ares(推荐)
https://github.com/stunnel/static-curl/tree/main
```
ARCHES="x86_64 arm64" \
    TLS_LIB=openssl \
    CURL_VERSION="" \
    QUICTLS_VERSION="" \
    OPENSSL_VERSION="" \
    NGTCP2_VERSION="" \
    NGHTTP3_VERSION="" \
    NGHTTP2_VERSION="" \
    LIBIDN2_VERSION="" \
    LIBUNISTRING_VERSION="" \
    ZLIB_VERSION="" \
    BROTLI_VERSION="" \
    ZSTD_VERSION="" \
    LIBSSH2_VERSION="" \
    LIBPSL_VERSION="" \
    ARES_VERSION="" \
    bash curl-static-mac.sh
```
或者直接下载可执行文件
```
echo 'export PATH="/Users/jinwei/tools/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

### curl支持DNS
```
# 指定DNS服务器
curl --dns-servers 8.8.8.8,1.1.1.1 https://www.google.com

# 显示详细的连接信息
curl --trace-ascii - https://www.google.com 2>&1 | grep -A5 -B5 "DNS\|resolve"


# 使用TCP进行DNS查询（而不是UDP）
curl --doh-url https://dns.google/dns-query https://www.google.com

# DoH (DNS over HTTPS)
curl --doh-url https://cloudflare-dns.com/dns-query https://www.google.com


# --doh-insecure - 跳过DoH证书验证
curl --doh-insecure --doh-url https://doh.example.com https://example.com

#  --doh-cert-status - 启用DoH证书状态验证
curl --doh-cert-status --doh-url https://doh.example.com https://example.com


#--doh-url: curl 7.62.0+
#--doh-insecure: curl 7.76.0+
#--doh-cert-status: curl 7.76.0+
```