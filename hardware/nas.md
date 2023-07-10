[TOC]
# NAS（Network Attached Storage：网络附属存储）

光猫地址: 192.168.1.1/cu.html
管理账号: CUAdmin
超级密码: cuadmin00259e

## 群辉 Synology
黑群辉安装引导:https://github.com/fbelavenuto/arpl

## 威联通

## 铁威马（TerraMaster）

## 开源

### freeNAS

### OpenMediaVault
简称 OMV https://www.openmediavault.org/
它是由原 FreeNAS 核心开发成员 Volker Theile 发起的基于 Debian Linux 的开源 NAS 操作系统，主要面向家庭用户和小型办公环境。

这个系统最大的优点就是定位轻量级，对硬件要求不高，而且功能简单，刚好够用，非常适合小白们上手。

> 安装教程: https://docs.openmediavault.org/en/latest/installation/via_iso.html
Web界面端口80
用户：admin
密码：openmediavault

### U-NAS
国产免费

### XigmaNAS
https://xigmanas.com/xnaswp/
你可以把这个系统当作TrueNAS的精简版，因为它们都是由原 FreeNAS 系统开发者创建的，它的配置要求没有 FreeNAS 高，并且还长久有人维护，但现在仍然在更新。

### TrueNAS
https://www.truenas.com/
TrueNAS的前身就是曾经非常火的FreeNAS，也是目前最受欢迎的开源免费 NAS 操作系统之一，基于以安全和稳定著称的 FreeBSD 系统开发，在2010年被 iXsystems公司收购。不过良心的是收购之后也一直是开源，只不过提供了两套方案，TrueNAS CORE为开源版本并且持续免费，TrueNAS Enterprise为商业版本需要付费购买额外的授权码。

该NAS系统功能强大，同时对硬件的要求也比较高，适合NAS进阶用户。

### openfiler

### RockStor
https://rockstor.com/

一款基于 Linux 的开源 NAS 系统，采用企业级文件系统 BTRFS，提供 SMB/CIFS、NFS 以及 SFTP 常见的共享方式。虽说定位于企业用户NAS系统，但是它的配置要求不是很高，2GB内存基本就能满足安装需求，所以个人用户也是可以玩玩的。

### NexentaStor
https://nexenta.com/
一款基于 OpenSolaris 开发，与 FreeNAS 一样采用强大的 ZFS 文件系统。该系统由 Nexenta Systems 公司技术团队维护，同时提供社区开原版和商业付费版本，官网有免费试用版本

### EasyNAS
https://easynas.org/
从名字上就能得知，它和OMV一样，是一款轻量级的简易NAS操作系统。但是它的体验还不错，具有稳定，高性能，高可用和其他你不常见的特性，并且还和RockStor 一样，同样采用企业级文件系统 BTRFS，安全性也是没有问题的。

### NASLite-2
http://www.serverelements.com/?target=NASLite-2
一款售价29.95刀的NAS 操作系统，它是少数基于 Linux 的商用 NAS 操作系统，由 Server Elements 公司出品。

### NanoNAS
http://www.serverelements.com/?target=NanoNAS
同样是 Server Elements 公司的NAS 操作系统，简单来说它其实就是上面 NASLite-2 的精简版，不过也是一样需要付费购买，价格为 9.95 刀。

## ONVIF

rtsp://192.168.31.183:554/cam/realmonitor?channel=1&subtype=0&unicast=true&proto=Onvif

rtsp://username:password@ip:port/cam/realmonitor?channel=1&subtype=0
说明:
username：用户名；
password：密码；
ip：为设备IP；
port：端口号默认为554，若为默认可不填写。
channel：通道号，起始为1。例如通道2，则为channel=2。
subtype：码流类型，主码流为0（即subtype=0），辅码流为1（即subtype=1）。
例如，请求某设备的通道1的主码流，Url如下：
rtsp://admin:lierdalux@192.168.1.108:554/cam/realmonitor?channel=1&subtype=1

可以用VLC直接打开此视频流！

乐橙云开放平台 https://open.imou.com/book/readme/model.html


- ONVIF Device Manager
- VLC
- shinobi

## NFS、FTP、SMB、WebDav、DLNA协议


- NFS
NFS是 Network File System的简写，即网络文件系统,网络文件系统是FreeBSD支持的文件系统中的一种，也被称为NFS.NFS允许一个系统在网络上与它人共享目录和文件。通过使用NFS，用户和程序可以像访问本地文件一样访问远端系统上的文件。

- FTP/SFTP/SCP
FTP（File Transfer ProtocoL）是TCP/IP应用层中的一个基础协议，通常使用 21 端口，负责将文件从一台计算机传输到另计算机，并保证文件传输的可靠性。
由于 FTP 是明文传输安全性不高，后来又出现了 SFTP 和 FTPS 等加密传输。

- SMB/Samba
SMB（Server Message Block）通信协议是微软和英特尔在1987年制定的协议，主要是作为Microsoft网络的通讯协议，它是当今世上网络文件系统协议两极之一的存在。
SMB使用了NetBIOS的应用程序接口 （Application Program Interface，简称API）。另外，它是一个开放性的协议，允许了协议扩展——使得它变得更大而且复杂；大约有65个最上层的作业，而每个作业都超过120个函数，甚至Windows NT也没有全部支持到，最近微软又把 SMB 改名为 CIFS（Common Internet File System），并且加入了许多新的特色。
Samba 是 SMB/CIFS（Server Message Block / Common Internet File System）网络协议的重新实现，可以在局域网不同计算机之间进行文件、打印机等资源共享，和 NFS 功能类似。
SMB 最早是微软为自己需求设计的专用协议，用来实现微软主机之间的文件共享与打印共享，并不支持在 Linux 上运行。著名黑客、技术大牛 Andrew Tridgell 通过逆向工程，在 Linux 上实现的 SMB / CIFS 兼容协议，命名为 Samba，通过该程序实现了 Windows 和 Linux 之间的文件共享。
SMB 协议是 C/S 类型协议，客户机通过该协议可以访问服务器上的共享文件系统、打印机及其他资源。通过设置“NetBIOS over TCP/IP”，Samba 不但能与局域网络主机分享资源，还能与全世界的电脑分享资源。
SMB 的优点之一是兼容性好，在各平台获得了广泛支持，包括 Windows、Linux、macOS 等各系统挂载访问都很方便。另外 SMB 也是各种电视、电视盒子默认支持的协议，可以通过 SMB 远程播放电影、音乐和图片。
另外 SMB 提供端到端加密、安全性高，配置选项丰富，支持 ACL 并支持多种用户认证模式。
SMB 的缺点是传输效率稍低，速度不太稳定，会有波动。

- WebDAV
WebDAV （Web-based Distributed Authoring and Versioning） 一种基于HTTP 1.1协议的 通信协议。它扩展 了HTTP 1.1，在GET、POST、HEAD等几个HTTP标准方法以外添加了一些新的方法，使应用程序可对Web Server 直接读写，并支持写文件锁定(Locking)及解锁(Unlock)，还可以支持文件的版本控制。
简单而言，WebDAV 就是一种互联网方法，应用此方法可以在服务器上划出一块存储空间，可以使用用户名和密码来控制访问，让用户可以直接存储、下载、编辑文件。
> [参考](https://zhuanlan.zhihu.com/p/411161467) 以及 [WebDav](../docker/clouddrive.md)

- DLNA
DLNA（DIGITAL LIVING NETWORK ALLIANCE，数字生活网络联盟) 其前身是DHWG （Digital Home Working Group，数字家庭工作组），成立于2003年6月24 日， 是由索尼 、 英特尔 、 微软等发起成立的一个非营利性的、合作性质的商业组织。
DLNA旨在解决个人PC ，消费电器，移动设备在内的无线网络和有线网络的互联互通，使得数字媒体和内容服务的无限制的共享和增长成为可能。DLNA的口号是Enjoy your music, photos and videos, anywhere anytime。
DLNA并不是创造技术，而是形成一种解决的方案，一种大家可以遵守的规范。所以，其选择的各种技术和协议都是当前所应用很广泛的技术和协议。
DLNA将其整个应用规定成5个功能组件。从下到上依次为：网络互连，网络协议，媒体传输，设备的发现控制和管理，媒体格式。
2017年2月20日，DLNA在其官网宣布：本组织的使命已经完成，已于2017年1月5日正式解散，相关的认证将移交给SpireSpark公司，未来不会再更新DLNA标准。



