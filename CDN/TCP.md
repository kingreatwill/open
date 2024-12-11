

### linux tcp fast open,TFO(TCP Fast Open)介绍
使用TFP（Tcp Fast Open）后，除了第一次的完整握手，之后的握手，在SYN中就可以携带数据，这样的好处是可以减少一个RTT。

TFO功能在Linux 3.7 内核中开始集成，因此RHEL7/CentOS7是支持的，但默认没有开启，使用以下方式开启：
`echo 3 > /proc/sys/net/ipv4/tcp_fastopen`
1表示开启客户端，
2表示开启服务器端
3的意思是开启TFO客户端和服务器端

除了内核的支持，应用程序也要开启支持，例如nginx(1.5.8+)开启方法如下：
```
server {
listen 80 backlog=4096 fastopen=256 default;
server_name _;
}
```

查看TFO连接的统计信息
```
# grep '^TcpExt:' /proc/net/netstat | cut -d ' ' -f 87-92 | column -t
```

其他问题: 
客户端的TFOcookie多长时间后删除，谁来维护和删除？
nginx的TFO队列具体是什么意思？队列满了会怎样？数值设定多少合适？
队列是RFC7413中的一种对服务器的安全保护机制，超出队列的数据包，会降级到普通的无cookie连接方式，即TFO功能失效。但这个数值具体设置多少不太好定。