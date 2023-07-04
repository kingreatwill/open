[TCP/IP协议详解（干货！）](https://blog.csdn.net/weixin_53186633/article/details/120514627)

[TCP/IP 协议](https://www.cnblogs.com/Liyuting/p/8805136.html)
[一篇文章带你熟悉 TCP/IP 协议（网络协议篇二）](https://juejin.im/post/5a069b6d51882509e5432656)
[通俗易懂-深入理解TCP协议（下）：RTT、滑动窗口、拥塞处理](http://www.52im.net/thread-515-1-1.html)
[tcp/ip 卷一 读书笔记（3）为什么既要有IP地址又要有MAC地址](https://www.cnblogs.com/zhangyufei/p/5565646.html)
[wireshark报文](https://wiki.wireshark.org/SampleCaptures)

[科来网络通讯协议图2019版](../files/tcp/科来网络通讯协议图2019版.pdf)


[接地气讲解TCP协议和网络程序设计（深度好文）](https://cloud.tencent.com/developer/article/1535959)

[接地气讲解UDP协议和网络程序设计（深度好文）](https://cloud.tencent.com/developer/article/1535962)

[3000字讲讲TCP协议，握手挥手不是你想的那么简单](https://www.toutiao.com/a6773950320000107016)

[TCP/IP、UDP、HTTP、MQTT、CoAP这五种物联网协议概述](https://www.toutiao.com/a6748413680527884813)

![](../img/tcp/网络通讯协议图.jpeg)



![](../img/tcp/tcp-ip-1.jpeg)
![](../img/tcp/tcp-ip-2.jpeg)
![](../img/tcp/tcp-ip-3.jpeg)
![](../img/tcp/tcp-ip-4.jpeg)
![](../img/tcp/tcp-ip-5.jpeg)
![](../img/tcp/tcp-ip-6.jpeg)
![](../img/tcp/tcp-ip-7.jpeg)
![](../img/tcp/tcp-ip-8.jpeg)
![](../img/tcp/tcp-ip-9.jpeg)


## KCP
TCP保证数据准确交付，UDP保证数据快速到达，KCP则是两种协议的一个折中。

KCP的设计目标是为了解决在网络拥堵的情况下TCP传输速度慢的问题。

KCP是一种网络传输协议(A Fast and Reliable ARQ Protocol)，可以视它为TCP的代替品，但是它运行于用户空间，它不管底层的发送与接收，只是个纯算法实现可靠传输，它的特点是牺牲带宽来降低延迟。

> KCP是一个快速可靠协议，能以比 TCP浪费10%-20%的带宽的代价，换取平均延迟降低 30%-40%，且最大延迟降低三倍的传输效果。纯算法实现，并不负责底层协议（如UDP）的收发，需要使用者自己定义下层数据的发送方式，以 callback的方式提供给 KCP。连时钟都需要外部传递进来，内部不会有任何一次系统调用。
TCP是为流量设计的（每秒内可以传输多少KB的数据），讲究的是充分利用带宽。而 KCP是为流速设计的（单个数据从一端发送到一端需要多少时间），以10%-20%带宽浪费的代价换取了比 TCP快30%-40%的传输速度。