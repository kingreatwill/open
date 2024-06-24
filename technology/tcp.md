[TCP/IP协议详解（干货！）](https://blog.csdn.net/weixin_53186633/article/details/120514627)

[TCP/IP 协议](https://www.cnblogs.com/Liyuting/p/8805136.html)
[一篇文章带你熟悉 TCP/IP 协议（网络协议篇二）](https://juejin.im/post/5a069b6d51882509e5432656)
[通俗易懂-深入理解TCP协议（下）：RTT、滑动窗口、拥塞处理](http://www.52im.net/thread-515-1-1.html)
[tcp/ip 卷一 读书笔记（3）为什么既要有IP地址又要有MAC地址](https://www.cnblogs.com/zhangyufei/p/5565646.html)

[从tcpdump抓包看TCP/IP协议](https://segmentfault.com/a/1190000015044878)

[wireshark报文](https://wiki.wireshark.org/SampleCaptures)

[科来网络通讯协议图2019版](../files/tcp/科来网络通讯协议图2019版.pdf)


[接地气讲解TCP协议和网络程序设计（深度好文）](https://cloud.tencent.com/developer/article/1535959)

[接地气讲解UDP协议和网络程序设计（深度好文）](https://cloud.tencent.com/developer/article/1535962)

[3000字讲讲TCP协议，握手挥手不是你想的那么简单](https://www.toutiao.com/a6773950320000107016)

[TCP/IP、UDP、HTTP、MQTT、CoAP这五种物联网协议概述](https://www.toutiao.com/a6748413680527884813)

![](../img/tcp/网络通讯协议图.jpeg)


### OSI七层模型
七层模型，亦称OSI（Open System Interconnection）。参考模型是国际标准化组织（ISO）制定的一个用于计算机或通信系统间互联的标准体系，一般称为OSI参考模型或七层模型。

#### 1. 物理层
建立、维护、断开物理连接。（由底层网络定义协议）机械、电子、定时接口通信信道上的原始比特流传输TCP/IP 层级模型结构，应用层之间的协议通过逐级调用传输层（Transport layer）、网络层（Network Layer）和物理数据链路层（Physical Data Link）而可以实现应用层的应用程序通信互联。
#### 2. 数据链路层
建立逻辑连接、进行硬件地址寻址、差错校验 等功能。（由底层网络定义协议）将比特组合成字节进而组合成帧，用MAC地址访问介质，错误发现但不能纠正。物理寻址、同时将原始比特流转变为逻辑传输线路地址解析协议：ARP、PARP（反向地址转换协议）
#### 3. 网络层
进行逻辑地址寻址，实现不同网络之间的路径选择。控制子网的运行，如逻辑编址、分组传输、路由选择协议有：ICMP（互联网控制信息协议） IGMP（组管理协议） IP（IPV4 IPV6）（互联网协议）安全协议、路由协议（vrrp虚拟路由冗余）
#### 4. 传输层
定义传输数据的协议端口号，以及流控和差错校验。接受上一层数据，在必要的时候把数据进行切割，并将这些数据交给网络层，并保证这些数据段有效到达对端协议有：TCP UDP，数据包一旦离开网卡即进入网络传输层
#### 5. 会话层
建立、管理、终止会话。（在五层模型里面已经合并到了应用层）不同机器上的用户之间建立及管理会话对应主机进程，指本地主机与远程主机正在进行的会话安全协议：SSL（安全套接字层协议）、TLS（安全传输层协议）
#### 6. 表示层
数据的表示、安全、压缩。（在五层模型里面已经合并到了应用层）信息的语法语义以及他们的关联，如加密解密、转换翻译、压缩解压格式有，JPEG、ASCll、EBCDIC、加密格式等 [2]
如LPP（轻量级表示协议）
#### 7. 应用层
网络服务与最终用户的一个接口各种应用程序协议,协议有：HTTP(超文本传输协议) FTP（文本传输协议） TFTP（简单文件传输协议） SMTP（简单邮件传输协议） SNMP（简单网络管理协议） DNS（域名系统） TELNET（远程终端协议） HTTPS（超文本传输安全协议） POP3（邮局协议版本3 ） DHCP（动态主机配置协议）
![](../img/tcp/tcp-ip-1.jpeg)


### TCP/IP四层模型
TCP/IP（Transmission Control Protocol/Internet Protocol，传输控制协议/网际协议）是指能够在多个不同网络间实现信息传输的协议簇。TCP/IP协议不仅仅指的是TCP 和IP两个协议，而是指一个由FTP、SMTP、TCP、UDP、IP等协议构成的协议簇， 只是因为在TCP/IP协议中TCP协议和IP协议最具代表性，所以被称为TCP/IP协议

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