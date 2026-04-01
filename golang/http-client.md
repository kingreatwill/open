### Request Hedging

请求对冲 (Request Hedging) 。通过对冲，客户端会将同一个请求的多个副本发送到不同的后端，并使用它收到的第一个响应。随后，客户端会取消所有未完成的请求，并将该响应转发给应用程序。
对冲是一种减少大规模分布式系统尾部延迟的技术。虽然简单的实现可能会给后端服务器增加巨大的负载，但在仅适度增加负载的情况下，完全可以获得大部分延迟降低的效果。

有效降低P99


https://github.com/bigwhite/experiments/tree/master/go-hedging-demo
https://grpc.org.cn/docs/guides/request-hedging/
https://www.reddit.com/r/golang/comments/1s4mb10/reduced_p99_latency_by_74_in_go_learned_something/
https://grpc.io/docs/guides/request-hedging/
https://research.google/pubs/the-tail-at-scale/