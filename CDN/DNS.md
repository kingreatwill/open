[TOC]

## 简介
### Linux 的nameserver-域名服务器
DNS（Domain Name System）是域名解析服务器的意思，它在互联网的作用是把域名转换成为网络可以识别的IP地址。当用户在浏览器中输入网址域名时，首先就会访问系统设置的DNS域名解析服务器（通常由ISP运营商如电信、网通提供）。如果该服务器内保存着该域名对应的IP信息，则直接返回该信息供用户访问网站。否则，就会向上级DNS逐层查找该域名的对应数据。

目前国内上网用户普遍使用的是默认DNS服务器，即电信运营商的DNS服务，这带来一个巨大的风险，就是DNS劫持。目前国内电信运营商普遍采用DNS劫持的方法，干扰用户正常上网，例如，当用户访问一个不存在（或者被封）的网站，电信运营商就会把用户劫持到一个满屏都是广告的页面：电信114网站，这个114网站不仅搜索质量低劣，而且广告众多，极大的影响了用户上网的安全性和浏览体验。后来，电信运营商的胆子越来越大，甚至连Google的网站电信都敢劫持，这进一步证明了电信运营商的DNS服务可靠性是多么糟糕。

普通用户要使用Google DNS非常简单，因为Google为他们的DNS服务器选择了两个非常简单易记的IP地址：“8.8.8.8”和“8.8.4.4”。用户只要在系统的网络设置中选择这两个地址为DNS服务器即可。
Google提供的公共DNS服务与电信网通的不同，当用户输入一个错误的或者不存在的网址的时候，不会像中国电信一般直接弹出一个满屏都是广告的页面，Google公司承诺绝不会重定向或者过滤用户所访问的地址，而且绝无广告。

Linux下设置：
```
echo nameserver 8.8.8.8 > /etc/resolv.conf
echo nameserver 8.8.4.4 > /etc/resolv.conf
```
这两行命令直接将8.8.8.8与8.8.4.4写入 Linux 的DNS客户端解析文件resolv.conf里。

阿里云:223.5.5.5
114DNS(南京信风): 114.114.114.114

> 114DNS在其官网上标榜纯净无劫持，背后却向广告主提供了多达几十项的劫持插入广告服务。https://www.114dns.com/
> 信风精确广告营销系统

## CoreDNS
使用CoreDNS作为内网DNS服务器
CoreDNS是Golang编写的一个插件式DNS服务器，是Kubernetes 1.13 后所内置的默认DNS服务器
采用的开源协议为Apache License Version 2
CoreDNS也是CNCF孵化项目，目前已经从CNCF毕业。
CoreDNS 的目标是成为 Cloud Native（云原生）环境下的 DNS 服务器和服务发现解决方案。

CoreDNS 插件编写目前有两种方式：
1. 深度耦合 CoreDNS，使用 Go 编写插件，直接编译进 CoreDNS 二进制文件
2. 通过 GRPC 解耦(实际上grpc就是一个插件)，任意语言编写 GRPC 接口实现，CoreDNS 通过 GRPC 与插件交互


> 由于 GRPC 链接实际上借助于 CoreDNS 的 GRPC 插件，同时 GRPC 会有网络开销，TCP 链接不稳定可能造成 DNS 响应过慢等问题。



[基于etcd插件(自带的)的CoreDNS动态域名添加](https://www.cnblogs.com/boshen-hzb/p/7541901.html)
https://coredns.io/plugins/etcd/

### 本地编译
```
git clone https://github.com/coredns/coredns
cd coredns/
make


$ file coredns
coredns: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows, 6 sections

$ ./coredns -version
CoreDNS-1.10.1
windows/amd64, go1.20.5, 5b5a6ac6-dirty

```
> windows编译出来的也是不带exe后缀的可执行文件

#### 插件
coredns官方对于插件的分类基本可以分为三种：[Plugins](https://coredns.io/plugins/)、[External Plugins](https://coredns.io/explugins/)和其他。
其中Plugins一般都会被默认编译到coredns的预编译版本中，而External Plugins则不会。
[Corefile](https://coredns.io/2017/07/23/corefile-explained/)

查看编译的插件` ./coredns -plugins`

编译插件基本可以分为:修改源码和修改编译的配置文件这两种方式
源码根目录下有个文件: `plugin.cfg`
部分内容如下
```
...
forward:forward
grpc:grpc
erratic:erratic
whoami:whoami
on:github.com/coredns/caddy/onevent
...
```

例如这里我们需要额外多添加一个dump插件到coredns中，只需要在plugin.cfg中加入插件的名称和地址

```
dump:github.com/miekg/dump
```
对于在plugin目录下已经存在的插件，则可以直接写成plugin中的目录名：
```
forward:forward
```

**开始编译**
```
go get github.com/miekg/dump
go generate
go build
make
```

**代码测试**
```go
// go get github.com/miekg/dns
func main() {
	c := dns.Client{
		Timeout: 5 * time.Second,
	}

	m := dns.Msg{}
	m.SetQuestion("www.baidu.com.", dns.TypeA)
	r, _, err := c.Exchange(&m, "127.0.0.1:53") // 192.168.220.2:53
	if err != nil {
		fmt.Println("dns error", err)
		return
	}

	var dst []string
	for _, ans := range r.Answer {
		record, isType := ans.(*dns.A)
		if isType {
			fmt.Println("type A:", record.A)
			dst = append(dst, record.A.String())
		}

		record1, isType := ans.(*dns.CNAME)
		if isType {
			fmt.Println("type cname:", record1.Target)
		}
	}

	for _, v := range dst {
		fmt.Println("ok:", v)
	}
}
```

**命令行**

```
doggo www.baidu.com @127.0.0.1

dig www.baidu.com @192.168.1.5
```


### 在docker环境中安装coredns
`docker pull coredns/coredns:1.10.1`
创建配置文件
```
mkdir -p /etc/coredns

vi /etc/coredns/Corefile

.:53 {
    hosts {
        10.0.0.1 my.host.com // 你域名和ip
        fallthrough
    }
    forward . 114.114.114.114:53 //你的备用 dns 
    log
}

也可以将 hosts 放在单独一个文件中

.:53 {
    hosts /etc/coredns/hostsfile {
        fallthrough
    }
    forward . 8.8.8.8:53
    log
}

# cat hostsfile
10.0.0.1 example1.org
```
启动服务
`docker run -it -d --name coredns --net=host -v /e/dockerv/coredns/:/etc/coredns/ coredns/coredns:1.10.1 -conf /etc/coredns/Corefile`



```
mkdir -p /etc/coredns

cat >/etc/coredns/Corefile<<EOF
.:53 {
    forward . 8.8.8.8:53
    log
}
EOF

docker run -d --name coredns \
  --restart=always \
  -v /e/dockerv/coredns/:/etc/coredns/ \
  -p 53:53/udp \
  coredns/coredns:1.10.1 -conf /etc/coredns/Corefile
```

测试
`dig @127.0.0.1 -p 53 my.host.com`


全部一起执行的命令
```
rm -rf /etc/coredns && mkdir -p /etc/coredns && echo "
.:53 {
    hosts {
        # ip host
        127.0.0.1 host.com 
        fallthrough
    }
    # forward . dns-1 dns-2
    forward . 114.114.114.114
    log
}
">> /etc/coredns/corefile && docker run -it -d --name dns --net=host -v /etc/coredns:/etc/coredns/ coredns/coredns:latest -conf /etc/coredns/corefile
```

## tool

### DNS client

更多可以在这里找到: https://www.isc.org/dns-tools/

#### nslookup
安装`yum install -y bind-utils`

`nslookup www.baidu.com`

#### host
`host www.baidu.com`
> host,nslookup,dig 作为 bind 的一部分,windows 下载: https://www.isc.org/bind/

#### dig
Dig 工具全称为域名信息搜索器（Domain Information Groper）
> dig 作为 bind 的一部分
> win版本下载: https://www.isc.org/bind/
> win版本下载2：ttp://www.bind9.net/download
> ftp地址: ftp://ftp.nominum.com/pub/isc/bind9/
> ftp地址: ftp://ftp.isc.org/isc/
> https://downloads.isc.org/isc/bind9/9.16.42/BIND9.16.42.x64.zip

安装`yum install -y bind-utils`
`dig www.wcoder.com`

我们可以使用 dig 命令追踪 www.wcoder.com 域名对应 IP 地址是如何被解析出来的，首先会向预置的 13 组根域名服务器发出请求获取顶级域名的地址：
`dig -t A www.wcoder.com +trace`

一般格式：
`dig [@global-server] [domain] [q-type] [q-class] {q-opt} {d-opt}`
参数说明：
@global-server：默认是以/etc/resolv.conf作为DNS查询的主机，这里可以填入其它DNS主机IP。
domain：要查询的域名。
q-type：查询记录的类型，例如a、any、mx、ns、soa、hinfo、axfr、txt等，默认查询a。
q-class：查询的类别，相当于nslookup中的set class。默认值为in（Internet）。
q-opt：查询选项，可以有好几种方式，比如：-f file为通过批处理文件解析多个地址；-p port指定另一个端口（缺省的DNS端口为53），等等。
d-opt：dig特有的选项。使用时要在参数前加上一个“+”号。

d-opt常用选项：
+vc：使用TCP协议查询。
+time=###：设置超时时间。
+trace：从根域开始跟踪查询结果。

> [详解 DNS 与 CoreDNS 的实现原理](https://juejin.cn/post/6844903709512564750)

#### dog 
https://kgithub.com/ogham/dog/
#### doggo
https://github.com/mr-karan/doggo



```sh
$ cd "$(mktemp -d)"
$ curl -sL "https://github.com/mr-karan/doggo/releases/download/v0.5.5/doggo_0.5.5_linux_amd64.tar.gz" | tar xz
$ mv doggo /usr/local/bin
```

`doggo www.wcoder.com`

#### tcpdump
`tcpdump` 是一个网络协议分析工具，可以用于捕获和分析网络数据包。它可以用于检查 DNS 查询和响应数据包，以及其他网络流量。
抓包命令: `tcpdump -i any -vvvvnnA dst port 8899`
可以`tcpdump -i any -vvvvnnA dst port 8899 -w file.cap` 保存文件中, 然后使用wireshark来查看数据包

`tcpdump -i eth0 udp port 53` 命令来捕获通过 eth0 网卡发送到 UDP 端口 53 的 DNS 数据包。

#### wireshark
`wireshark` 是一个网络协议分析器，可以用于分析网络数据包的详细信息。它可以用于检查 DNS 查询和响应数据包，以及其他网络流量。例如，使用 `wireshark` 命令来打开捕获的 DNS 数据包文件并进行分析。

Sniffer(嗅探器)就是利用计算机的网络接口截获目的地为其他计算机的数据报文的一种技术。
> sniffnet

#### 网络监控工具
- GlassWire
- Nutty
- Portmaster
- sniffnet
https://github.com/GyulyVGC/sniffnet

#### systemd-resolve

`systemd-resolve` 是一个 systemd 系统服务,是 Ubuntu下 DNS 解析相关的命令，可用于解析 DNS 名称。它可以用于查询本地 DNS 缓存和配置文件中指定的 DNS 服务器。例如，使用 `systemd-resolve www.baidu.com` 命令来查询 www.baidu.com 的 DNS 记录。帮助`systemd-resolve --help`

systemd-resolve 命令可以用来设置指定网卡的 DNS Server，如下
```
sudo systemd-resolve --set-dns '8.8.8.8' --interface ens3

# 查看
systemd-resolve --status | grep 'DNS Servers'
         DNS Servers: 8.8.8.8

# 重置网卡的 DNS 设置
systemd-resolve --revert --interface {ITERFACE_NAME}

# 刷新本地 DNS 缓存
systemd-resolve --flush-caches
```

#### systemd-resolved
https://cloud-atlas.readthedocs.io/zh_CN/latest/linux/redhat_linux/systemd/systemd_resolved.html

(centos安装:`yum -y install systemd-resolved`) 
systemctl status systemd-resolved
systemctl enable systemd-resolved
systemctl start systemd-resolved