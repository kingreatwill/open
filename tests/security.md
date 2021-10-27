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
kali基于Debian

## Blackarch
Blackarch 是一款基于 arch linux 为渗透测试和安全研究人员的渗透测试系统，它包含2700多种安全工具。

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

## 安全工具
[2006年100款最佳安全工具](https://blog.51cto.com/switch/7400)
1、Nessus ：最好的UNIX漏洞扫描工具
2、Wireshark ：嗅探网络粘合胶水（网络协议）
3、Snort ：一款广受欢迎的开源IDS（Intrusion Detection System）（入侵检测系统）工具
4、Netcat ：网络瑞士×××
5、Metasploit Framework ： 黑掉整个星球
6、Hping2 ：一种网络探测工具，是ping的超级变种
7、Kismet ：一款超强的无线嗅探器
8、Tcpdump ：最经典的网络监控和数据捕获嗅探器
9、Cain and Abel ： Windows平台上最好的密码恢复工具
10、John the Ripper ：一款强大的、简单的以及支持多平台的密码哈希破解器
11、Ettercap ：为交换式局域网提供更多保护
12、Nikto ：非常全面的网页扫描器
13、Ping/telnet/dig/traceroute/whois/netstat ：基本命令
14、OpenSSH / PuTTY / SSH ：访问远程计算机的安全途径
15、THC Hydra ：支持多种服务的最快的网络认证破解器
16、Paros proxy ： 网页程序漏洞评估代理
17、Dsniff ：一款超强的网络评估和渗透检测工具套装
18、NetStumbler ： 免费的Windows 802.11嗅探器
19、THC Amap ：一款应用程序指纹扫描器
20、GFI LANguard ：一款Windows平台上的商业网络安全扫描器
21、Aircrack ： 最快的WEP/WPA破解工具
22、Superscan ：只运行于Windows平台之上的端口扫描器、ping工具和解析器
23、Netfilter ： 最新的Linux核心数据包过滤器/防火墙
24、Sysinternals ：一款强大的非常全面的Windows工具合集
25、Retina ： eEye出品的商业漏洞评估扫描器
26、Perl / Python / Ruby ：简单的、多用途的脚本语言
27、L0phtcrack ： Windows密码猜测和恢复程序
28、Scapy ：交互式数据包处理工具
29、Sam Spade ： Windows网络查询免费工具
30、GnuPG / PGP ：对您的文件和通讯进行高级加密
31、Airsnort ： 802.11 WEP加密破解工具
32、BackTrack ：一款极具创新突破的Live（刻在光盘上的，光盘直接启动）光盘自启动Linux系统平台
33、P0f ：万能的被动操作系统指纹工具
34、Google ：人人喜爱的搜索引擎
35、WebScarab ：一个用来分析使用HTTP和HTTPS协议的应用程序框架
36、Ntop ：网络通讯监控器
37、Tripwire ：祖爷爷级的文件完整性检查器
38、Ngrep ：方便的数据包匹配和显示工具
39、Nbtscan ： 在Windows网络上收集NetBIOS信息
40、WebInspect ：强大的网页程序扫描器
41、OpenSSL ： 最好的SSL/TLS加密库
42、Xprobe2 ：主动操作系统指纹工具
43、EtherApe ： EtherApe是Unix平台上的模仿etherman的图形界面网络监控器
44、Core Impact ： 全自动的全面入侵检测工具
45、IDA Pro ： Windows或Linux反编译器和调试器
46、SolarWinds ： 网络发现/监控/攻击系列工具
47、Pwdump ：一款Windows密码恢复工具
48、LSoF ：打开文件列表
49、RainbowCrack ：极具创新性的密码哈希破解器
50、Firewalk ：高级路由跟踪工具
51、Angry IP Scanner ： 一款非常快的Windows IP 扫描器和端口扫描器
52、RKHunter ： 一款Unix平台上的Rootkit检测器
53、Ike-scan ： ×××检测器和扫描器
54、Arpwatch ： 持续跟踪以太网/IP地址配对，可以检查出中间人攻击
55、KisMAC ： 一款Mac OS X上的图形化被动无线网络搜寻器
56、OSSEC HIDS ：一款开源的基于主机的入侵检测系统
57、Openbsd PF ： OpenBSD数据包过滤器
58、Nemesis ：简单的数据包注入
59、Tor ：匿名网络通讯系统
60、Knoppix ：一款多用途的CD或DVD光盘自启动系统
61、ISS Internet Scanner ：应用程序漏洞扫描器
62、Fport ： Foundstone出品的加强版netstat
63、chkrootkit ： 本地rootkit检测器
64、SPIKE Proxy ： HTTP攻击
65、OpenBSD ：被认为是最安全的操作系统
66、Yersinia ：一款支持多协议的底层攻击工具
67、Nagios ：一款开源的主机、服务和网络监控程序
68、Fragroute/Fragrouter ：一款网络入侵检测逃避工具集
69、X-scan ：一款网络漏洞扫描器
70、Whisker/libwhisker ： Rain.Forest.Puppy出品的CGI漏洞扫描器和漏洞库
71、Socat ：双向数据传输中继
72、Sara ：安全评审研究助手
73、QualysGuard ：基于网页的漏洞扫描器
74、ClamAV ：一款UNIX平台上的基于GPL（通用公开许可证：General Public License）的反病毒工具集
75、cheops / cheops-ng ：提供许多简单的网络工具，例如本地或远程网络映射和识别计算机操作系统
76、Burpsuite ：一款网页程序攻击集成平台
77、Brutus ：一款网络验证暴力破解器
78、Unicornscan ：另类端口扫描器
79、Stunnel ：用途广泛的SSL加密封装器
80、Honeyd ： 您私人的honeynet
81、Fping ：一个多主机同时ping扫描程序
82、BASE ：基础分析和安全引擎（Basic Analysis and Security Engine）
83、Argus ： IP网络事务评审工具
84、Wikto ：网页服务器评估工具
85、Sguil ：网络安全监控器命令行分析器
86、Scanrand ：一个异常快速的无状态网络服务和拓朴结构发现系统
87、IP Filter ： 小巧的UNIX数据包过滤器
88、Canvas ：一款全面的漏洞检测框架
89、VMware ：多平台虚拟软件
90、Tcptraceroute ： 一款基于TCP数据包的路由跟踪工具
91、SAINT ：安全管理综合网络工具
92、Open××× ：全功能SSL ×××解决方案
93、OllyDbg ：汇编级Windows调试器
94、Helix ： 一款注重安全防护的Linux版本
95、Bastille ： Linux、Mac OS X和HP-UX的安全加强脚本
96、Acunetix Web Vulnerability Scanner ：商业漏洞扫描器
97、TrueCrypt ：开源的Windows和Linux磁盘加密软件
98、Watchfire AppScan ：商业网页漏洞扫描器
99、N-Stealth ：网页服务器扫描器
100、MBSA ：微软基准安全分析器（Microsoft Baseline Security Analyzer）

