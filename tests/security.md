# 安全
## kali linux
https://www.kali.org/downloads/
https://www.kali.org/docs/
https://github.com/kali-docs-cn/kali-linux-cookbook-zh

[WLS2（win10子系统）Linux图形化GUI教程win-kex](https://blog.csdn.net/l1447320229/article/details/108210760)
```
wsl -l -v
kali-linux
enter 123456
kex
```

## 工具
### 端口扫描
- masscan
https://github.com/robertdavidgraham/masscan
- nmap

## 安全测试
### SQL注入
[Sqlmap](https://github.com/sqlmapproject/sqlmap)
[NoSQLMap](https://github.com/codingo/NoSQLMap)
### Fuzz测试
什么是Fuzz测试？
漏洞挖掘有三种方法：白盒代码审计、灰盒逆向工程、黑盒测试。其中黑盒的Fuzz测试是效率最高的一种，能够快速验证大量潜在的安全威胁。

Fuzz测试，也叫做“模糊测试”，是一种挖掘软件安全漏洞、检测软件健壮性的黑盒测试，它通过向软件输入非法的字段，观测被测试软件是否异常而实现。Fuzz测试的概念非常容易理解，如果我们构造非法的报文并且通过测试工具打入被测设备，那么这就是一个Fuzz测试的测试例执行，大多数测试工程师肯定都尝试过这种测试手段。

对于网络协议漏洞挖掘来说，Fuzz测试也就意味着打入各种异常报文，然后观察设备是否有异常。

#### Peach是一个优秀的开源Fuzz框架。
Fuzz（模糊测试）是一种通过提供非预期的输入并监视异常结果来发现软件安全漏洞的方法。模糊测试在很大程度上是一种强制性的技术，简单并且有效，但测试存在盲目性。

典型地模糊测试过程是通过自动的或半自动的方法，反复驱动目标软件运行并为其提供构造的输入数据，同时监控软件运行的异常结果。

Fuzz被认为是一种简单有效的黑盒测试，随着Smart Fuzz的发展，RCE（逆向代码工程）需求的增加，其特征更符合一种灰盒测试。

https://www.peach.tech/

#### AFL、LibFuzzer、honggfuzz

### Python安全测试工具合集
https://www.cnblogs.com/xiaodi914/p/5176094.html


https://github.com/topics/security-scanner?l=python


本文列出13个python网络测试工具，共大家参考学习。

1、Scapy，Scapy3k：发送，嗅探和剖析并伪造网络数据包，可以做交互式应用或单纯的作为库来使用。

2、pypcap，pcapy和pylibpcap：几个不同的libpcap捆绑Python库

3、libdnet：低级别的网络路由器，可用于接口查找和以太网帧转发

4、dpkt：快速、轻量级的数据包创建、解析工具，适用于基本TCP/IP协议

5、Impacket：探测和解码网络数据包，支持更高级别协议比如NMB和SMB

6、pynids：libnids封装提供嗅探，IP碎片整理，TCP流重组和端口扫描检测

7、Dirtbags py-pcap：无需libpcap即可读取pcap文件

8、flowgrep：通过正则表达式查找数据包中的Payloads

9、Knock Subdomain Scan：通过字典枚举目标域上的子域名

10、SubBrute：可扩展的TCP/UDP中间代理，支持即时修改非标准协议

11、Pytbull：灵活的IDS/IPS测试框架（配有300多个测试用例）

12、Spoodle：大量子域名+Poodle漏洞扫描器

13、SMBMap：枚举域中的Samba共享驱动器