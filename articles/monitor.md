[TOC]
# Monitoring System
[List Of 11 Best Open Source & Free Monitoring Tools](https://devopscube.com/best-opensource-monitoring-tools/)

## Netdata

[Netdata](https://github.com/netdata/netdata) 是一款 Linux 性能实时监测工具.。以web的可视化方式展示系统及应用程序的实时运行状态（包括cpu、内存、硬盘输入/输出、网络等linux性能的数据）。
1.CPU的使用率,中断，软中断和频率(总量和每个单核)

2.RAM，互换和内核内存的使用率（包括KSM和内核内存deduper）

3.硬盘输入/输出(每个硬盘的带宽，操作，整理，利用等)

4.IPv4网络（数据包，错误，分片）：
TCP：连接，数据包，错误，握手
UDP:数据包，错误
广播：带宽，数据包
组播：带宽，数据包

5.Netfilter/iptables Linux防火墙(连接，连接跟踪事件，错误等)

6.进程(运行，受阻，分叉，活动等)

7.熵

8.NFS文件服务器，v2,v3,v4(输入/输出，缓存，预读，RPC调用)

9.网络服务质量（唯一一个可实时可视化网络状况的工具）
11.应用程序，通过对进程树进行分组（CPU,内存，硬盘读取，硬盘写入，交换，线程，管道，套接字等）
12.Apache Web服务器状态(v2.2, v2.4)

13.Nginx Web服务器状态

14.Mysql数据库（多台服务器，单个显示：带宽，查询/s, 处理者，锁，问题，临时操作，连接，二进制日志，线程，innodb引擎等）

15.ISC Bind域名服务器（多个服务器，单个显示：客户，请求，查询，更新，失败等）

16.Postfix邮件服务器的消息队列（条目，大小）

17.Squid代理服务器（客户带宽和请求，服务带宽和请求）

18.硬件传感器（温度，电压，风扇，电源，湿度等）

19.NUT UPSes（负载，充电，电池电压，温度，使用指标，输出指标）

Netdata management and configuration cheatsheet:
![](https://learn.netdata.cloud/assets/images/netdata-cheatsheet-e09a6f8706934ebf3bb5ac1c5b4a50d4.png)


用脚本安装：
```
# 基本安装（适合所有Linux系统）
$ bash <(curl -Ss https://my-netdata.io/kickstart.sh)

# 或者从头安装（安装所有依赖包）
$ bash <(curl -Ss https://my-netdata.io/kickstart-static64.sh) 
```
通过报管理器安装：
```
Arch Linux (sudo pacman -S netdata)
Alpine Linux (sudo apk add netdata)
Debian Linux (sudo apt-get install netdata)
Gentoo Linux (sudo emerge --ask netdata)
OpenSUSE (sudo zypper install netdata)
Solus Linux (sudo eopkg install netdata)
Ubuntu Linux >= 18.04 (sudo apt install netdata)
MacOS (brew install netdata)
```
> localhost:19999

开关控制：
```
# 启动：位置根据系统会有不同。建议加上-D参数前台运行，不要后台运行
$ sudo netdata -D
$ 或
$ sudo /usr/sbin/netdata -D
# 或（Mac上）
$ sudo /usr/local/sbin/netdata -D

# 关闭（方法很多种，往往只有一种生效）
$ sudo killall netdata
# 或
$ sudo pkill -9 netdata
# 或
$ sudo service netdata stop
# 或
$ sudo /etc/init.d/netdata stop
# 或
$ sudo systemctl stop netdata
```
卸载：
```
# 找到卸载脚本位置(我的在/usr/src/netdata.git)
whereis netdata.git
#  进入那个位置
cd /usr/src/netdata.git
# 开始卸载
yes | sudo ./netdata-uninstaller.sh --force
# 删除其它遗留信息
sudo userdel netdata
sudo groupdel netdata
sudo gpasswd -d netdata adm
sudo gpasswd -d netdata proxy
```
卸载其它被安装的软件：
```
yes | sudo apt-get purge ntop snapd
```

Nginx代理
```
location ~ /netdata/(?<ndpath>.*) {
            proxy_redirect off;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_http_version 1.1;
            proxy_pass_request_headers on;
            proxy_set_header Connection "keep-alive";
            proxy_store off;
            proxy_pass http://127.0.0.1:19999/$ndpath$is_args$args;
            gzip on;
            gzip_proxied any;
            gzip_types *;
        }
```
重启Nginx，然后应该就可以通过域名/netdata访问Netdata面板了。

> 注意，安装后会严重影响服务器运行速度，且默认安装的话安全性很低（默认没有密码，默认端口访问可以直接看到全部系统信息）

## Prometheus

## Zabbix

https://www.zabbix.com/
Zabbix 是开源监控软件，界面简单易用，用户学习曲线较平滑，并且可为大型企业提供企业级解决方案。
它是一个集中式系统，存储的数据是一个关系型数据库，可以对其进行高效地处理。

## Nagios

https://www.nagios.org/

Nagios 是一款开源的监控工具，1999 年就已经问世。
它提供了许多设施，如使用额外的插件与第三方应用程序集成。
考虑到 Nagios 这一领域已经存在已久，因此生态比较完善，有很多为它编写的插件。
它可以监控各种组件，包括 Oss、应用程序、网站、中间件、Web 服务器等。


## Open-Falcon

https://github.com/open-falcon/falcon-plus/
http://open-falcon.org/

## Nightingale

https://github.com/didi/nightingale


## Rieman

http://riemann.io/
Riemann 是分布式系统的理想监控工具。
它是一个低延迟的事件处理系统，能够收集各种分布式系统的指标。
它的设计是为了以低延迟处理每秒数百万个事件。
它是一个适用于高度分布式可扩展系统的监控工具。

## Sensu

https://sensu.io/
Sensu 是一个全栈监控工具。通过统一的平台，你可以监控服务、应用程序、服务器和业务 KPI 报告。
它的监控不需要单独的工作流程并且它支持所有流行的操作系统，如 Windows、Linux 等。

## Icinga

https://icinga.com/

Icinga 是一个开源的网络监控工具，可以计算网络的可用性和性能。通过 web 界面，你的企业可以观察到整个网络基础设施中的应用程序和主机。
该工具是可扩展的，并且易于配置，以配合每一种类型的设备。Icinga 模块中存在一些非常特殊的监控功能，比如对 VMWare 的 vSphere 云环境和业务流程建模的监控。

## Cacti

https://www.cacti.net/

Cacti 是一个建立在 RRD Tool 的数据分类和绘图系统上的开源监控工具。
它利用数据收集功能和网络轮询功能来收集任意范围的网络中各种设备的信息。
这包括创建数据收集的自定义脚本以及 SNMP 轮询的能力。
然后，它将这些信息展示在易于理解的图表中，这些图表可以根据你的业务组织成任何层次。

## LibreNMS

https://www.librenms.org/

LibreNMS 是一个开源的网络监控系统，它利用多种网络协议来监控网络上的每个设备。
LibreNMS 的 API 可以恢复、管理和绘制它所收集的数据，并促进水平扩展，使其监控能力与你的网络一起成长。
该工具有一个灵活的告警系统，它可以自定义，因此你可以采用最适合自己的方式来设置它。

## Observium Community

https://www.observium.org/

Observium Community 是 Observium 网络监控工具的免费版本。在免费版本中，你可以监控无限数量的设备，并充分利用 Obersvium 的网络映射属性。
Observium 网络监控工具的特性是对连接的设备进行程序化的发现。它还配备了发现协议，以确保你的网络地图是最新的。通过这种方式，你可以跟踪新设备与网络连接的情况。

## Pandora FMS

http://pandorafms.org/

Pandora FMS 是一款开源的监控工具，可以帮助企业观察整个 IT 子结构。
它不仅具有网络监控功能，还具有 Unix 和 Windows 服务器以及虚拟接口的监控功能。
对于网络来说，Pandora FMS 包含了 SNMP 支持、ICMP 轮询、网络延迟监控以及系统过载等顶级功能。
此外，还可以在设备上安装代理，以监控设备温度以及日志文件等方面的情况。

## LogRhythm NetMon Freemium

https://logrhythm.com/products/logrhythm-netmon-freemium/

LogRhythm NetMon Freemium 是 LogRhythm NetMon 的免费版本，提供与完整版类似的业务级模块采集和分析能力。
虽然在数据处理和模块存储上有一定的限制或局限，但免费版仍然允许用户执行建立在数据包分析基础上的网络风险检测和响应功能。
它还提供了与完整版类似的网络威胁预警系统，让你随时了解网络的性能和安全状况。

## SolarWinds ：实时带宽监控器

https://www.solarwinds.com/free-tools/real-time-bandwidth-monitor

SolarWinds 实时带宽监控 是一款免费的开源带宽监控工具。
该工具可实时掌握带宽使用情况，并以带宽轮询为中心显示网络带宽的曲线图。
当带宽使用情况进入紧张状态时，该工具会通知你，让你的企业立即知道你的网络带宽何时不足。
此外，可以自定义描述关键的带宽使用情况，这样该工具就能准确知道网络上的设备何时使用了超过所需的带宽。

## Famatech 高级 IP 扫描器

https://www.advanced-ip-scanner.com/cn/

Famatech 高级 IP 扫描器是一款免费的网络监控以及扫描工具，可对局域网和 LAN 设备进行分析。高级 IP 扫描器可以扫描网络上的设备，并远程调节连接的计算机和其他资源。
它提供了在设备不使用和使用资源的情况下，将计算机从工具中关闭的功能。该工具与 Famatech 的 Radmin 解决方案相配合，实现远程 IT 管理，无论你在哪里都可以管理 IP。

## AppNeta PathTest

https://www.appneta.com/resources/pathtest-download.html

AppNeta PathTest 是一款免费的网络容量测试工具，旨在帮助企业理解其网络的真正能力。PathTest 旨在通过精确描述网络的最大能力来提高第三层和第四层的性能。
它故意用数据包充斥你的网络，使网络达到满负荷状态。用户可以将该测试的持续时间最多设置为 10 秒，并在任何时候运行测试。
