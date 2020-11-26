<!--toc-->
[TOC]
## 文件系统

[Ceph vs GlusterFS vs MooseFS vs HDFS vs DRBD](https://computingforgeeks.com/ceph-vs-glusterfs-vs-moosefs-vs-hdfs-vs-drbd/)
### Ceph
Ceph是一个强大的存储系统，它在同一个系统中同时提供了对象，块（通过RBD）和文件存储。无论您是希望在虚拟机中使用块设备，还是将非结构化数据存储在对象存储中，Ceph都可以在一个平台上提供所有功能，并且还能获得出色的灵活性。 Ceph中的所有内容都以对象的形式存储，不管原始的数据类型是什么，RADOS（reliable autonomic distributed object store）都会把它们当做对象来进行存储。

RADOS层确保数据始终保持一致状态并且可靠。Ceph会通过数据复制，故障检测和恢复，以及跨群集节点进行数据迁移和重新平衡来实现数据一致性。 Ceph提供了一个符合POSIX的网络文件系统（CephFS），旨在实现高性能，大数据存储以及与传统应用程序的最大兼容。Ceph可以通过各种编程语言或者radosgw（RGW）实现无缝的访问对象存储，（RGW）这是一种REST接口，它与为S3和Swift编写的应用程序兼容。另一方面，Ceph的RADOS块设备（RBD）可以访问在整个存储集群中条带化和复制的块设备映像。

#### Ceph的特性
- 独立、开放和统一的平台：将块，对象和文件存储组合到一个平台中，包括最新添加的CephFS
- 兼容性：您可以使用Ceph 存储对外提供最兼容Amazon Web Services（AWS）S3的对象存储。
- 精简配置模式：分配存储空间时，只是虚拟分配容量，在跟进使用情况占用实际磁盘空间。这种模式提供了更多的灵活性和磁盘空间利用率。
- 副本：在Ceph Storage中，所有存储的数据都会自动从一个节点复制到多个其他节点。默认任何时间群集中的都有三份数据。
- 自我修复：Ceph Monitors会不断监控你的数据集。一旦出现一个副本丢失，Ceph会自动生成一个新副本，以确保始终有三份副本。
- 高可用：在Ceph Storage中，所有存储的数据会自动从一个节点复制到多个其他的节点。这意味着，任意节点中的数据集被破坏或被意外删除，在其他节点上都有超过两个以上副本可用，保证您的数据具有很高的可用性。
- Ceph很强大：您的集群可以用于任何场景。无论您希望存储非结构化数据或为数据提供块存储或提供文件系统，或者希望您的应用程序直接通过librados使用您的存储，而这些都已经集成在一个Ceph平台上了。
- 可伸缩性：Ceph Works 可以在集群中随时增加，从而满足将来的规模需求。

Ceph最适合用于块存储，大数据或直接与librados通信的任何其他应用程序。 这一切都会顺利运行的非常好。 更多Ceph的信息可以参见[Ceph文档](http://docs.ceph.com/)
### GlusterFS
Gluster是一个免费的开源可扩展网络文件系统。 使用通用的现成硬件，您可以为媒体流，数据分析以及其他数据和带宽密集型任务创建大型的分布式存储解决方案。 基于GlusterFS的横向扩展存储系统适用于非结构化数据，例如文档，图像，音频和视频文件以及日志文件。 通常，分布式文件系统依赖于元数据服务器，但是Gluster不再使用元数据服务器。 元数据服务器是单点故障，并且可能是扩展的瓶颈。 相反，Gluster使用哈希机制来查找数据。

#### Gluster特性
- 可扩展性：可扩展的存储系统，可提供弹性和配额
- 快照：卷和文件级快照都支持，用户可以直接发起快照请求，这意味着用户不必费心管理员即可创建快照
- 归档：只读卷和一次写入多次读（WORM）卷都支持归档。
- 为了获得更好的性能，Gluster会对readdir（）的数据，元数据和目录条目进行缓存。
- 集成：Gluster与oVirt虚拟化管理器以及用于服务器的Nagios监控器可以集成在一起。
- 大数据：对于那些希望使用Gluster文件系统中的数据进行数据分析的人，提供了Hadoop分布式文件系统（HDFS）支持。
- libgfapi：应用程序可以绕过其他访问方式，直接使用libgfapi与Gluster对话。这对于上下文切换或内核空间复制敏感的工作负载很有用

有关Gluster的其他详细信息，请参见[Gluster Docs](https://docs.gluster.org/)。

### MooseFS 
MooseFS是波兰公司Gemius SA公司在12年前推出的，是大数据存储行业中的突破性概念。 它使您可以使用负担得起的商用硬件将数据存储和数据处理结合在一个单元中。

#### MooseFS特性
- 冗余：所有系统组件都是冗余的，如果发生故障，会触发自动故障转移机制，这些对用户是透明的。
- 节点上的计算：通过利用空闲的CPU和内存资源，支持在数据节点上调度计算，可以降低系统的总体拥有成本。
- 原子快照：在任何特定时间点配置文件系统都是瞬间完成且不间断。 此特性非常适合用于在线备份的解决方案。
- 分层存储：将不同类别的数据分配给各种类型的存储介质，以降低总存储成本。 可以将热数据存储在快速的SSD磁盘上，而将不经常使用的数据转移到更便宜，更慢的机械硬盘驱动器上。
- 本地客户端：通过专门为Linux，FreeBSD和MacOS系统设计的专用客户端（安装）组件来提高性能。
- 全局回收站:：一个虚拟的全局空间，用于记录删除对象的，和每个文件和目录的配置。 借助这个有利的特性，可以轻松恢复意外删除的数据
- 配额限制：系统管理员可以灵活地设置限制，以限制每个目录的数据存储容量。
- 滚动升级：能够一次执行一个节点的升级，硬件替换和添加，而不会中断服务。 此功能使您可以在不停机的情况下保持硬件平台的最新状态。
- 快速磁盘恢复：万一硬盘或硬件出现故障，系统会立即启动从冗余副本到系统内其他可用存储资源的并行数据复制。 此过程比传统的磁盘重建方法快得多。
- 并行性：在执行的并行线程中执行所有I / O操作，以提供高性能的读/写操作。
- 管理界面：提供丰富的管理工具，例如基于命令行和基于Web的界面


有关MooseFS的更多信息，[请参见](https://moosefs.com/)
### HDFS 
Hadoop分布式文件系统（HDFS）是一个分布式文件系统，它允许快速同时存储和检索多个文件。 它可以方便地在商业硬件上运行，并提供处理非结构化数据的功能。 它应用程序提供数据的高吞吐量访问，并且适用于具有大数据集的应用程序。 HDFS与Hadoop YARN，Hadoop MapReduce和Hadoop Common一起是Hadoop的主要组成部分。 它是Hadoop框架的基本组件之一。

#### HDFS特性
- 数据复制：HDFS被设计用于在大型群集中的计算机之间可靠地存储非常大的文件。 它将每个文件分块进行存储； 除了这个文件最后一个块以外的所有块都具有相同的大小。 这些块将通过复制的方式进行容错。
- 文件命名： HDFS支持传统的分层方式组织文件。 用户或应用程序可以创建目录并将文件存储在这些目录中。 文件系统名称空间的层次结构与大多数其他现有文件系统相似。 可以创建和删除文件，将文件从一个目录移动到另一个目录或重命名文件。 HDFS尚未实现用户配额。 HDFS不支持硬链接或软链接。
- 健壮： HDFS的主要目标是即使在出现故障的情况下也能可靠地存储数据。 三种常见的故障类型是NameNode故障，DataNode故障和网络分区。
- 适用性：应用程序访问HDFS支持多种不同的方式。 HDFS本身就为应用程序提供了Java API， 同时也提供此Java API的C语言包装器。 另外，HTTP浏览器也可用于浏览HDFS实例的文件。 通过WebDAV协议公开HDFS的工作正在进行中。
- 可扩展性：HDFS被设计用于在大型群集中的计算机之间可靠地存储非常大的文件。 可以根据当时的需求增加或减少群集。
- 高可用性：HDFS被设计用于在大型集群中的机器之间可靠地存储非常大的文件。 它将每个文件存储为一系列块； 文件中除最后一块以外的所有块都具有相同的大小。 复制文件的块是为了容错，因此在发生任何故障的情况下数据的可用性很高。

有关HDFS的更多信息[请参见](https://hadoop.apache.org/docs/r1.2.1/hdfs_design.html)
### DRBD

DRBD是一个分布式冗余存储系统，由内核驱动程序，多个用户空间管理应用程序和一些Shell脚本实现。 分布式复制块设备（一个逻辑卷中的逻辑块设备）在多个主机之间镜像块设备，以实现高可用集群。 基于DRBD的群集通常用于为文件服务器，关系数据库（例如MySQL）和许多其他工作负载提供同步复制和高可用性。 DRBD实质上可以作为共享磁盘文件系统，额外的逻辑块设备（例如LVM），常规文件系统或需要直接访问块设备的任何应用程序的基础。

#### DRDB特性

- DRDB具有共享秘密身份验证功能；
- DRBD与LVM（Logical Volume Manager）兼容；
- 支持与heartbeat/pacemaker 集成；
- 支持读取请求的负载平衡；
- 支持出现故障时自动检测最新数据
- Delta重新同步
- 支持在线调整DRBD配置，而不会丢失数据
- 支持自动带宽管理
- 支持自定义的调整参数
- 支持相互在线数据验证
- 高可用性：块设备在多个主机之间镜像块设备，以实现高度可用的群集。
- 支持与Xen等虚拟化解决方案集成，并且可以在Linux LVM 之上和之下使用

有关DRBD更多的信息[请参见](https://docs.linbit.com/)

#### FastDFS
https://github.com/happyfish100/fastdfs c 6.9k

轻量级分布式文件系统，他对文件进行管理，主要功能有：文件存储，文件同步，文件访问（文件上传/下载）,特别适合以文件为载体的在线服务，如图片网站，视频网站等

#### SeaweedFS
https://github.com/chrislusf/seaweedfs golang 10.9k

#### curve
https://github.com/opencurve/curve c++ 590
CURVE是网易自主设计研发的高性能、高可用、高可靠分布式存储系统，具有非常良好的扩展性。基于该存储底座可以打造适用于不同应用场景的存储系统，如块存储、对象存储、云原生数据库等。当前我们基于CURVE已经实现了高性能块存储系统，支持快照克隆和恢复 ,支持QEMU虚拟机和物理机NBD设备两种挂载方式, 在网易内部作为高性能云盘使用。

### 其它文件共享
- CIFS是SMB的改进其使用可以简单理解为windows中的共享文件夹
- Samba服务器主要是用来windows访问linux的共享文件夹 关于samba服务器的配置谷歌之。[参考](http://www.cnblogs.com/mchina/archive/2012/12/18/2816717.html)
- NFS服务器主要用于linux之间共享文件.[参考](http://www.cnblogs.com/mchina/archive/2013/01/03/2840040.html)
- NAS 可以简单理解为windows的网络驱动映射

- PATA/FATA/SCSI/FC/SAS这些是硬盘的连接技术 。     
       历史顺序应该SCSI->FC(用于SCSI)->ISCSI
       SCSI（Small Computer System Interface，小型计算机系统接口）

-  将存储类型区分，可分为DAS、FC SAN、IP SAN，IP SAN又包括iSCSI与NAS（NFS／CIFS）。

FC SAN是采用光纤信道的SAN，也就是服务器透过光纤信道卡（FC HBA），连接光纤交换器（FC Switch），再连接后端的存储设备。
IP SAN就是，服务器透过以太网络连接后端存储设备，后端的存储设备可被看作一台含有档案处理系统的存储服务器，如果采用微软的操作系统，Windows Storage Server 2003，其档案系统名称为CIFS（Common Internet File System）；如果是采用Linux、Unix操作系统，其档案系统名称为NFS（Network File System）。

NAS与SAN最大的区别在于，档案存取的方式不同。FC SAN、以iSCSI形成的IP SAN是采用区块层级（block-level）的传输方式，NAS则是以档案层级（file-level）的传输方式。不同的档案存取方式，将影响不同 的应用层面，例如，电子邮件、网页服务器、多媒体影音串流服务、档案分享等就适用于NAS存储架构。但是若是与数据库有关的应用，则要采用SAN架构，这 里指的SAN包括FC SAN与iSCSI。

iSCSI（Internet SCSI／SCSI over IP），是IETF制订的一项标准，用于将SCSI数据块映射成以太网数据包。其适用于TCP／IP通讯协议，在以太网络上传输SCSI的指令，是一个以IP为主的SAN，好处是让企业不用架设昂贵的光纤信道费用，以现有的 以太网络为基础。iSCSI的的好处打破了FC或SCSI的距离限制，并且使多台服务器用享有后端的存储设备资源，并且原本SCSI限制只能连接8或16 个设备，iSCSI则允许比前者可连接更多存储设备。iSCSI可以实现在IP网络上运行SCSI协议，使其能够在诸如高速千兆以太网上进行路由选择。

#### SMB
Server Message Block）又称CIFS(Common Internet File System),一种应用层网络传输协议（微软（Microsoft）和英特尔(Intel)在1987年制定的协议），由微软开发，主要功能是使网络上的机器能够共享计算机文件、打印机、串行端口和通讯等资源。它也提供认证的进程间通讯技能。它主要用在Windows的机器上。

#### CIFS
CIFS是由microsoft在SMB的基础上发展，扩展到Internet上的协议。他和具体的OS无关，在unix上安装samba后可使用CIFS.它使程序可以访问远程Internet计算机上的文件并要求此计算机的服务。CIFS 使用客户/服务器模式。客户程序请求远在服务器上的服务器程序为它提供服务。服务器获得请求并返回响应。
CIFS是公共的或开放的SMB协议版本，并由Microsoft使用。SMB协议现在是局域网上用于服务器文件访问和打印的协议。象SMB协议一样，CIFS在高层运行，而不象TCP/IP协议那样运行在底层。CIFS可以看做是应用程序协议如文件传输协议和超文本传输协议的一个实现。


#### NFS
（Network File System）是一种分布式文件系统协议，它允许网络中的计算机之间通过TCP/IP网络共享资源。在NFS的应用中，本地NFS的客户端应用可以透明地读写位于远端NFS服务器上的文件，就像访问本地文件一样。是应用在linux下的协议。用于linux系统间文件的共享，用于小型存储服务或者小型网站

#### samba
Samba是用于Linux和Unix的标准Windows互操作性程序套件。
Samba是种用来让UNIX系列的操作系统与微软Windows操作系统的SMB/CIFS（Server Message Block/Common Internet File System）网络协议做链接的自由软件。简而言之，此软件在Windows与UNIX系列OS之间搭起一座桥梁，让两者的资源可互通有无。
samba是许多服务以及协议的实现，其包括TCP/IP上的NetBIOS、SMB、CIFS等等协议。
Samba能够为选定的Unix目录（包括所有子目录）创建网络共享。该功能使得Windows用户可以像访问普通Windows下的文件夹那样来通过网络访问这些Unix目录。

#### NAS
(Network Attached Storage)被定义为一种特殊的专用数据存储服务器，包括存储器件（例如磁盘阵列、CD/DVD驱动器、磁带驱动器或可移动的存储介质）和内嵌系统软件，可提供跨平台文件共享功能。NAS通常在一个LAN上占有自己的节点，无需应用服务器的干预，允许用户在网络上存取数据，在这种配置中，NAS集中管理和处理网络上的所有数据，将负载从应用或企业服务器上卸载下来，有效降低总拥有成本，保护用户投资。
NAS本身能够支持多种协议（如NFS、CIFS、FTP、HTTP等），而且能够支持各种操作系统。通过任何一台工作站，采用IE或Netscape浏览器就可以对NAS设备进行直观方便的管理。

#### DAS
（Direct Attached Storage，直接附属存储），也可称为SAS（Server-Attached Storage，服务器附加存储）。
DAS被定义为直接连接在各种服务器或客户端扩展接口下的数据存储设备，它依赖于服务器，其本身是硬件的堆叠，不带有任何存储操作系统。在这种方式中，存储设备是通过电缆（通常是SCSI接口电缆）直接到服务器的，I/O（输入/输入）请求直接发送到存储设备。


#### SAN
（Storage AreaNet work，存储区域网络）。它是一种通过光纤集线器、光纤路由器、光纤交换机等连接设备将磁盘阵列、磁带等存储设备与相关服务器连接起来的高速专用子网。
SAN由三个基本的组件构成：接口（如SCSI、光纤通道、ESCON等）、连接设备（交换设备、网关、路由器、集线器等）和通信控制协议（如IP和SCSI等）。
这三个组件再加上附加的存储设备和独立的SAN服务器，就构成一个SAN系统。
SAN提供一个专用的、高可靠性的基于光通道的存储网络，SAN允许独立地增加它们的存储容量，也使得管理及集中控制（特别是对于全部存储设备都集群在一起的时候）更加简化。而且，光纤接口提供了10 km的连接长度，这使得物理上分离的远距离存储变得更容易。



## 云盘系统


### ownCloud
https://github.com/owncloud/

### Nextcloud
https://github.com/nextcloud/

### Seafile 
https://github.com/haiwen/seafile
### OnionShare
https://github.com/micahflee/onionshare

### 国产
#### dboxShare
https://github.com/dboxshare/dboxShare
#### Cloudreve
https://github.com/cloudreve/Cloudreve

DEMO演示站：https://drive.aoaoao.me
##### 目前已经实现的功能：
- 快速对接多家云存储，支持七牛、又拍云、阿里云OSS、AWS S3、Onedrive、自建远程服务器，当然，还有本地存储
- 可限制单文件最大大小、MIMEType、文件后缀、用户可用容量
- 自定义主题配色
- 基于Aria2的离线下载
- 图片、音频、视频、文本、Markdown、Ofiice文档 在线预览
- 移动端全站响应式布局
- 文件、目录分享系统，可创建私有分享或公开分享链接
- 用户个人主页，可查看用户所有分享
- 多用户系统、用户组支持
- 初步完善的后台，方便管理
- 拖拽上传、分片上传、断点续传、下载限速（*实验性功能）
- 多上传策略，可为不同用户组分配不同策略
- 用户组基础权限设置、二步验证
- WebDAV协议支持


##### 宝塔安装方法
1. 打开终端工具，用命令行安装composer，之后再用composer工具安装cloudreve.
```
curl -sS https://getcomposer.org/installer | php
mv composer.phar /usr/local/bin/composer
```
2. 使用Composer安装Cloudreve
```
composer create-project hfo4/cloudreve:dev-master
#等待安装依赖库后，会自动执行安装脚本，按照提示输入数据库账户信息
```
3. 打开宝塔面板，进入网站
将cloudreve里的文件剪切到网站根目录，然后修改runtime目录权限777，如果你使用本地存储，public 目录也需要有写入权限.

4. 设置伪静态
```
location / {
   if (!-e $request_filename) {
   rewrite  ^(.*)$  /index.php?s=/$1  last;
   break;
    }
 }
```
初始用户名admin@cloudreve.org 初始密码 admin 后台URlhttp://你的域名/Admin,登录后

5. 正式环境：请设置定时任务
在终端输入contab -e粘贴以下内容：`* * * * * curl http://你的域名/Cron`之后reboot服务器即可。

##### 使用Docker
```
# 注意修改你的域名,没有的话，填写ip
docker run -p 80:80 -v /cloudreve:/cloudreve -e CLOUDREVE_URL="http://你的域名:6780/" -e APACHE2_WEB_PORT="6780" --name cloudreve ilemonrain/cloudreve

```
docker pull xavierniu/cloudreve:3.1.1

#### z-file
https://github.com/zhaojun1998/zfile
#### dzzoffice
https://github.com/zyx0814/dzzoffice

#### kiftd
https://github.com/KOHGYLW/kiftd-source
https://github.com/KOHGYLW/kiftd

#### 蓝眼云盘
https://github.com/eyebluecn/tank
#### iBarn
https://github.com/zhimengzhe/iBarn
#### KODExplorer
https://github.com/kalcaddle/KODExplorer

## Web文件浏览器
### filebrowser
https://github.com/filebrowser/filebrowser
## 图床
[盘点国内免费好用的图床](https://zhuanlan.zhihu.com/p/35270383)
### PicGo
https://github.com/Molunerfinn/PicGo
PicGo 本体支持如下图床：
- 七牛图床 v1.0
- 腾讯云 COS v4\v5 版本 v1.1 & v1.5.0
- 又拍云 v1.2.0
- GitHub v1.5.0
- SM.MS V2 v2.3.0-beta.0
- 阿里云 OSS v1.6.0
- Imgur v1.6.0

支持vs插件：https://github.com/PicGo/vs-picgo
支持手机：https://github.com/PicGo/flutter-picgo

### PicUploader
https://github.com/xiebruce/PicUploader
### ImgURL
https://github.com/helloxz/imgurl
### weiboUploader-Watermark
https://github.com/yhf7952/weiboUploader-Watermark

### gitPic
https://github.com/zzzzbw/gitPic

### auxpi
集合多家 API 的新一代图床
https://github.com/0xDkd/auxpi

### 公益图床
https://sbimg.cn/
### 路过图床
https://imgchr.com
简介：支持免注册上传图片，永久存储，支持HTTPS加密访问和调用图片，提供多种图片链接格式，成立于2011年
限制：最大10M

### SM.MS
官网地址：https://sm.ms
特点：永久存储免注册，图片链接支持https，可以删除上传的图片，提供多种图片链接格式，建立于2015年，目前免费用户无法使用香港节点因此速度比较慢
图片上传限制：每个图片最大5M，每次最多上传10张

### 聚合图床
官网地址：https://www.superbed.cn

简介：将图片分发到多处备份，借助其本身的CDN加速功能，节省服务器流量，并且不用担心图片被删除，即便其中某几个图床上的图片被删除了，还有其他备份，保证万无一失，支持匿名和注册管理

图片上传限制：无

### iPic
iPic 可以屏幕截图、还是复制图片，都可以自动上传、保存 Markdown 格式的链接，直接粘贴插入，也可使用 Hexo | Heroku 或 WordPress 写博客、在公众号发文章、在知乎讨论、在豆瓣灌水、在论坛发帖等 。（注： iPic 是Mac的一款图床工具）。
### kieng综合图床
https://image.kieng.cn/

- 阿里图床
https://image.kieng.cn/ali.html
- 腾讯图床
https://image.kieng.cn/qq.html
- 网易图床
https://image.kieng.cn/wy.html
- 头条图床
https://image.kieng.cn/tt.html
- 京东图床
https://image.kieng.cn/jd.html
- 侠聚图床
https://image.kieng.cn/hl.html
- 搜狐图床
https://image.kieng.cn/sh.html

### 各大图床API接口

BaiDu 百度 图床 API
https://api.uomg.com/api/image.baidu

360识图 图床 API
http://st.so.com/stu

网易云 图床 API
http://you.163.com/xhr/file/upload.json

JingDong 京东 图床 API
https://search.jd.com/image?op=upload

JueJin 掘金 图床 API
https://cdn-ms.juejin.im/v1/upload?bucket=gold-user-assets

Ali 阿里云 图床 API
https://kfupload.alibaba.com/mupload
https://api.uomg.com/api/image.ali
https://api.uomg.com/doc-image.ali.html

## 记录分享文件存储
### perkeep 
https://github.com/perkeep/perkeep 5.2k
### Upspin
https://github.com/upspin/upspin 5.6k
### IPFS
https://github.com/ipfs/ipfs 18.3k
### Keybase Filesystem
https://keybase.io/docs/kbfs
### git-annex
### Google Drive
### Dropbox
### Libchop
### Tahoe-LAFS
predates Perkeep, file-centric
### Unhosted


## 办公套件

### onlyoffice 
https://github.com/ONLYOFFICE/CommunityServer csharp

onlyoffice 是 collaborative system

ONLYOFFICE是一款集成了文档、电子邮件、事件、任务和客户关系管理工具的开源在线办公套件。其文档管理功能实现了文档的在线编辑、在线预览和协同管理，可用于替代Office365或Google docs。另外，它还提供了CRM、项目管理等功能，非常合适作为企业内部的全员协作Office系统。功能强大简直令人惊叹，而这样一个复杂的系统，用docker不到三分钟即可自动安装完成。


- 集成文档、电子邮件、任务、聊天系统、客户管理等
- 一套成熟的开源在线办公，支持插件扩展
- 多人在线协同

`docker run -itd -p 666:80 --restart=always onlyoffice/communityserver`


## 私有云

### TTstack
可快速生成各种虚拟机环境
https://gitee.com/kt10/ttstack


