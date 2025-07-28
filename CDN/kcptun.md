


## kcptun
https://github.com/xtaci/kcptun


1，增加文件打开数量。
```
ulimit -n 65535
```
2，调整udp包的相关参数。
```

net.core.rmem_max=26214400 // BDP - bandwidth delay product
net.core.rmem_default=26214400
net.core.wmem_max=26214400
net.core.wmem_default=26214400
net.core.netdev_max_backlog=2048 // proportional to -rcvwnd
```
3，增大buffer大小，默认的buffer大小为4MB。对于速度较慢的处理器，增加缓冲区对于正确接收数据包至关重要。
```

-sockbuf 16777217
```

4，kcptun启动的一个示例：
```

KCP Client: ./client_darwin_amd64 -r "KCP_SERVER_IP:4000" -l ":8388" -mode fast3 -nocomp -autoexpire 900 -sockbuf 16777217 -dscp 46
KCP Server: ./server_linux_amd64 -t "TARGET_IP:8388" -l ":4000" -mode fast3 -nocomp -sockbuf 16777217 -dscp 46
```
一、如果拥有高速网络连接，如何提高带宽使用？
同时增加kcptun客户端rcvnd参数和kcptun服务端sndnd参数。这些值中的最小值决定了链路的最大传输速率，如wnd * mtu / rtt。然后，尝试下载一些东西，看看它是否符合您的要求。MTU可以通过-mtu进行调整。
二、如果使用kcptun加速游戏，如何减小延迟？
延迟通常表示数据包丢失。您可以通过更改-mode参数来减少延迟。这里延迟从小到大的模式：fast3>fast2>fast>normal>default。
三、如何减少头部阻塞？
由于流被多路复用到单个物理信道中，因此可能会发生头部阻塞。将smuxbuf增加到更大值（默认为4MB）可能会缓解这个问题，但它会使用更多的内存。对于版本>= v20190924，可以切换到smux版本2。这个版本可以现在流内存使用。通过-streambuf可以限制每个流的内存使用。
四、性能低的设备如何调优？例如嵌入式设备？
kcptun使用Reed-Solomon码来恢复丢失的数据包，而这需要大量的计算。低端ARM设备可能无法很好地与kcptun配合使用。为了获得最佳性能，建议使用AMD、Intel等多核服务器CPU。如果必须使用ARM路由器，最好禁用FEC并使用salsa20作为加密方法。

## 其它加速
https://github.com/wangyu-/UDPspeeder


绕过UDP防火墙
https://github.com/wangyu-/udp2raw
https://github.com/Chion82/kcptun-raw


## 参考
https://zhuanlan.zhihu.com/p/453349832

## stun/TURN
https://github.com/jselbie/stunserver
```
Docker

1. `docker image build -t=stun-server-image .`
2. `docker container run -d -p 3478:3478/tcp -p 3478:3478/udp
--name=stun-container stun-server-image`
```


https://www.stunprotocol.org/

ICE - Interactive Connectivity Establishment. A protocol for establishing direct connectivity once STUN or TURN address candidates are obtained.
TURN - Traversal Using Relays around NAT. A relay solution and protocol for when direct connectivity can not be established. Can be used in conjunction or seperate from STUN.
WebRTC - ICE data and video conferencing in web browsers
Pion - Implementation of WebRTC in the Go programming language
libnice - libnice is an open-source C source code library that implements the ICE protocol for establishing P2P connectivity.
JSTUN - JSTUN is an implementation of STUN using Java implemented by Thomas King. JSTUN client libraries are compatible with STUNTMAN server.
C# Stun Client code - Implemented by by Ivar Lumi. Compatible with STUNTMAN.
PJNATH - An implementation of ICE for multiple platforms

### nat类型检测
https://github.com/oneclickvirt/gostun/tree/main
https://github.com/HMBSbige/NatTypeTester

### 解释
1. 标准STUN协议是否支持TCP？
STUN（RFC 5389）协议本身是支持TCP传输的，即STUN消息可以通过UDP或TCP发送到STUN服务器。
但——
绝大多数STUN实现和NAT类型检测工具只用UDP，因为NAT类型的判别（Full Cone/Restricted/Symmetric等）主要是为P2P UDP打洞服务设计的。
NAT对TCP和UDP的处理方式不同，TCP是有连接的，UDP是无连接的，NAT的端口分配和映射策略也不同。
2. 为什么很少有“TCP STUN” NAT类型检测工具？
NAT类型的定义和判别方法（如RFC 3489/5389）是基于UDP的，TCP下的NAT行为和UDP下不同，且TCP打洞本身就很难实现。
TCP连接建立时，NAT通常会为每个新连接分配新的端口，这导致用TCP做NAT类型判别意义不大。
STUN的主要应用场景（如WebRTC、VoIP）都是基于UDP的。
3. 现有STUN服务器/客户端是否支持TCP？
协议上支持：RFC 5389 明确规定STUN消息可以通过TCP发送。
实现上支持：如 coturn、pion/stun 等STUN服务器/库支持TCP传输。
但：这些实现通常不会用TCP来检测NAT类型，而是用来做NAT穿透时的备用传输方式（如UDP被阻断时用TCP）。
4. 你可以用TCP方式和STUN服务器通信吗？
可以！你可以用TCP连接STUN服务器，发送STUN消息，获取公网IP和端口。但：
你无法用TCP STUN检测UDP下的NAT类型。
你也无法用TCP STUN检测TCP下的NAT类型，因为TCP连接的NAT行为和UDP完全不同，且没有标准的“TCP NAT类型”判别方法。
5. 结论
STUN协议本身支持TCP传输，但NAT类型检测的标准方法只适用于UDP。
没有专门的“TCP STUN NAT类型检测”工具或标准，因为TCP NAT行为和UDP不同，且TCP打洞本身极难实现。
你可以用TCP方式和STUN服务器通信，获取公网IP:端口，但这不能用来判别NAT类型。