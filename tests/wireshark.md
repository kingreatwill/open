# wireshark
wireshark是非常流行的网络封包分析软件，功能十分强大。可以截取各种网络封包，显示网络封包的详细信息。使用wireshark的人必须了解网络协议，否则就看不懂wireshark了。

wireshark能获取HTTP，也能获取HTTPS，但是不能解密HTTPS，所以wireshark看不懂HTTPS中的内容，如果是处理HTTP,HTTPS 还是用Fiddler, 其他协议比如TCP,UDP 就用wireshark.


wireshark是捕获机器上的某一块网卡的网络包，当你的机器上有多块网卡的时候，你需要选择一个网卡。

## TODO
深入TCP的时候需要学习

## 其它
### Charles
Charles是一个HTTP代理/ HTTP监视器/反向代理，它使开发人员能够查看他们的机器和Internet之间的所有HTTP和SSL / HTTPS通信。这包括请求、响应和HTTP头(其中包含cookie和缓存信息)。

### AnyProxy
一个完全可配置的http/https代理在NodeJS 
