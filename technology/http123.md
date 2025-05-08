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
