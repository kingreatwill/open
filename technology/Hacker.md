# hacker

## 密码显示
```
javascript:"use strict";!function(){var e,t;e=document.getElementsByTagName("input");for(var a=0;a<e.length;a++)if(t=e[a],"password"==t.type.toLowerCase())try{t.type="text"}catch(e){var r,n;r=document.createElement("input"),n=t.attributes;for(var o=0;o<n.length;o++){var i,c,d;i=n[o],c=i.nodeName,d=i.nodeValue,"type"!=c.toLowerCase()&&"height"!=c&&"width"!=c&!!d&&(r[c]=d)}t.parentNode.replaceChild(r,t)}}();
```

## 常见网站攻击技术
SQL 注入、XSS 攻击、CSRF 攻击、DDoS 攻击、DNS劫持、XXE 漏洞、JSON 劫持、暴力破解、HTTP 报头追踪漏洞、信息泄露、目录遍历漏洞、命令执行漏洞、文件上传漏洞、SSLStrip 攻击、OpenSSL Heartbleed 安全漏洞、CCS 注入漏洞、证书有效性验证漏洞、业务漏洞、框架或应用漏洞

[web 应用常见安全漏洞一览](https://segmentfault.com/a/1190000018004657)
[常见网站攻击技术](https://mp.weixin.qq.com/s/LLudKqVMkKzTIz407N2UMg)

## h4cker
https://github.com/The-Art-of-Hacking/h4cker 7k

h4cker 是包含数千个与网络安全相关的参考资料和资源的项目，项目由 Omar Santos 维护，资料中包括：渗透测试，数字取证和事件响应（DFIR），漏洞研究，漏洞利用开发，逆向工程等内容。 ​​​​


## Maltrail
Maltrail是一款功能强大免费的开源恶意流量检测系统, 它利用公开的黑名单以及从各种AV报告和自定义用户特征来识别恶意流量, 同时,该系统还拥有可选的高级启发式机制, 可以帮助使用者发现一些未知的威胁。
https://github.com/stamparm/maltrail?tab=readme-ov-file#suspicious-http-requests

构建镜像
```
cd docker
docker build -t maltrail .
```

```
wget https://raw.githubusercontent.com/stamparm/maltrail/master/docker/Dockerfile
wget https://raw.githubusercontent.com/stamparm/maltrail/master/maltrail.conf

docker run -d --name maltrail --privileged -p 8337:8337/udp -p 8338:8338 -v /data/dockerv/maltrail/log/:/var/log/maltrail/ -v /data/dockerv/maltrail/config/maltrail.conf:/opt/maltrail/maltrail.conf:ro kingreatwill/maltrail:v0.74
```
访问IP:8338 , 用户名密码:`admin:changeme!`, 也可以在maltrail.conf中查看和修改`echo -n 'changeme!' | sha256sum | cut -d " " -f 1`

> docker部署无法探测宿主机上的流量


## TheFatRat
https://github.com/Screetsec/TheFatRat 4.5k

TheFatRat 是一种利用恶意软件编译具有著名负载的恶意软件，然后可以在 Linux，Windows，Mac 和Android 上执行该恶意软件。TheFatRat 提供了一种轻松创建后门和有效负载的方法，可以绕过大多数防病毒软件。特征：

- 全自动MSFvenom 和 Metasploit
- 本地或远程侦听器生成
- 按类别轻松地制作后门操作系统
- 生成各种格式的有效载荷
- 绕过反病毒后门
- 可用于增加文件大小的文件泵
- 能够检测外部 IP 和接口地址
- 自动创建用于 USB / CDROM 的 AutoRun 文件

## hackingtool
https://github.com/Z4nzu/hackingtool 4.6k

hackingtool 收录 Hackers 所需的大部分工具，包括 SQL 注入、钓鱼攻击、Hash 破解、XSS、DDos 攻击等等分类工具。



## hacker101
https://github.com/Hacker0x01/hacker101 11.5k

hacker101 收录了 HackerOne 出品的 Web 安全免费在线课程的源代码，无论你是对漏洞众测（Bug Bounty）感兴趣，或是经验丰富的安全专家，阅读本 repo 想必都能有所收获。

## hacker-roadmap
https://github.com/sundowndev/hacker-roadmap 3.8k


hacker-roadmap 是一份爱好者指南，并收集了黑客工具、资源和参考资料，以实践黑客道德准则和网络安全。

## 防攻击/Web Application Firewall (WAF)
### 雷池
https://github.com/chaitin/safeline

### bunkerweb
https://github.com/bunkerity/bunkerweb

## 暴力破解
暴力破解一般看密码库
密码库: 
https://github.com/danielmiessler/SecLists/tree/master/Passwords
https://github.com/berzerk0/Probable-Wordlists
https://wiki.skullsecurity.org/index.php/Passwords
https://weakpass.com/

### SSH暴力破解防护-Fail2ban
### Brutus ：一款网络验证暴力破解器
### Legba
https://github.com/evilsocket/legba

Supported Protocols/Features:
AMQP (ActiveMQ, RabbitMQ, Qpid, JORAM and Solace), Cassandra/ScyllaDB, DNS subdomain enumeration, FTP, HTTP (basic authentication, NTLMv1, NTLMv2, multipart form, custom requests with CSRF support, files/folders enumeration, virtual host enumeration), IMAP, Kerberos pre-authentication and user enumeration, LDAP, MongoDB, MQTT, Microsoft SQL, MySQL, Oracle, PostgreSQL, POP3, RDP, Redis, Samba, SSH / SFTP, SMTP, Socks5, STOMP (ActiveMQ, RabbitMQ, HornetQ and OpenMQ), TCP port scanning, Telnet, VNC.

### Hydra
https://github.com/vanhauser-thc/thc-hydra

目前该工具支持以下协议的爆破：
AFP，Cisco AAA，Cisco身份验证，Cisco启用，CVS，Firebird，FTP，HTTP-FORM-GET，HTTP-FORM-POST，HTTP-GET，HTTP-HEAD，HTTP-PROXY，HTTPS-FORM- GET，HTTPS-FORM-POST，HTTPS-GET，HTTPS-HEAD，HTTP-Proxy，ICQ，IMAP，IRC，LDAP，MS-SQL，MYSQL，NCP，NNTP，Oracle Listener，Oracle SID，Oracle，PC-Anywhere， PCNFS，POP3，POSTGRES，RDP，Rexec，Rlogin，Rsh，SAP / R3，SIP，SMB，SMTP，SMTP枚举，SNMP，SOCKS5，SSH（v1和v2），Subversion，Teamspeak（TS2），Telnet，VMware-Auth ，VNC和XMPP。

## 打卡神器
### 在PC上控制Android设备
https://github.com/pdone/FreeControl

https://github.com/karma9874/AndroRAT

## 参考

https://github.com/topics/hacking

https://github.com/Hack-with-Github/Awesome-Hacking

[All in One 你想知道的 hacker 技术都在这里](https://www.cnblogs.com/xueweihan/p/13549926.html)
[安全技术](https://paper.seebug.org)
[网络安全](https://www.freebuf.com/)