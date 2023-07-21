# BitTorrent (BT)

[Taipei Torrent golang](https://github.com/jackpal/Taipei-Torrent) 
[chihaya](https://github.com/chihaya/chihaya) multi-protocol BitTorrent Tracker
[torrent golang](https://github.com/anacrolix/torrent)


[qBittorrent c++](https://github.com/qbittorrent/qBittorrent)
[transmission c++](https://github.com/transmission/transmission)
[rtorrent c++](https://github.com/rakshasa/rtorrent)
[picotorrent c++](https://github.com/picotorrent/picotorrent)


## BT种子文件结构（.torrent）
BitTorrent协议的种子文件（英语：Torrent file）可以保存一组文件的元数据。这种格式的文件被BitTorrent协议所定义。扩展名一般为“.torrent”。

.torrent种子文件本质上是文本文件，包含Tracker信息和文件信息两部分。Tracker信息主要是BT下载中需要用到的Tracker服务器的地址和针对Tracker服务器的设置，文件信息是根据对目标文件的计算生成的，计算结果根据BitTorrent协议内的Bencode规则进行编码。它的主要原理是需要把提供下载的文件虚拟分成大小相等的块，块大小必须为2k的整数次方（由于是虚拟分块，硬盘上并不产生各个块文件），并把每个块的索引信息和Hash验证码写入种子文件中；所以，种子文件就是被下载文件的“索引”。



Torrent总体结构
键名称 | 数据类型 | 可选项 | 键值含义
---|---|---|---
announce | string | required | Tracker的Url
info |dictionary |required |该条映射到一个字典，该字典的键将取决于共享的一个或多个文件
announce-list|array[]|optional|备用Tracker的Url，以列表形式存在
comment|string|optional|备注
created by|string|optional|创建人或创建程序的信息

Torrent单文件Info结构

键名称 | 数据类型 | 可选项 | 键值含义
---|---|---|---
name|string|required|建议保存到的文件名称
piceces|byte[]|required|每个文件块的SHA-1的集成Hash。
piece length|long|required|每个文件块的字节数

Torrent多文件Info结构
键名称 | 数据类型 | 可选项 | 键值含义
---|---|---|---
name|string|required|建议保存到的目录名称
piceces|byte[]|required|每个文件块的SHA-1的集成Hash。
piece length|long|required|每个文件块的字节数
files|array[]|required|文件列表，列表存储的内容是字典结构

files字典结构：

键名称 | 数据类型 | 可选项 | 键值含义
---|---|---|---
path|array[]|required|一个对应子目录名的字符串列表，最后一项是实际的文件名称
length|long|required|文件的大小（以字节为单位）

> JSON内容省略了pieces摘要大部分内容，仅展示了开头部分，另外由于本人序列化工具设置所致，所有的整型都会序列化成字符串类型。
```
{
    "creation date": "1604347014",
    "comment": "Torrent downloaded from https://YTS.MX",
    "announce-list": [
        [
            "udp://tracker.coppersurfer.tk:6969/announce"
        ],
        [
            "udp://9.rarbg.com:2710/announce"
        ],
        [
            "udp://p4p.arenabg.com:1337"
        ],
        [
            "udp://tracker.internetwarriors.net:1337"
        ],
        [
            "udp://tracker.opentrackr.org:1337/announce"
        ]
    ],
    "created by": "YTS.AG",
    "announce": "udp://tracker.coppersurfer.tk:6969/announce",
    "info": {
        "pieces": "ﾬimￏﾋ\u000b*ﾟﾬﾄ... ...",
        "name": "Love And Monsters (2020) [2160p] [4K] [WEB] [5.1] [YTS.MX]",
        "files": [
            {
                "path": [
                    "Love.And.Monsters.2020.2160p.4K.WEB.x265.10bit.mkv"
                ],
                "length": "5215702961"
            },
            {
                "path": [
                    "www.YTS.MX.jpg"
                ],
                "length": "53226"
            }
        ],
        "piece length": "524288"
    }
}
```





全部内容必须都为**Bencoding编码**类型。整个文件为一个**字典**结构，包含如下关键字：
-|-
---|---
announce |	tracker 服务器的 URL（字符串）；
announce-list（可选）|	备用 tracker 服务器列表（列表）；
creation date（可选）|	种子创建的时间，Unix 标准时间格式，从 1970 1 月1 日 00：00：00 到创建时间的秒数（整数）；
comment（可选）	|备注（字符串）
created by（可选）	|创建人或创建程序的信息（字符串）；
info	|一个字典结构，包含文件的主要信息。分为二种情况，单文件结构或多文件结构。

单文件info结构如下
-|-
---|---
length|	文件长度，单位字节（整数）；
md5sum（可选）|	长 32 个字符的文件的 MD5 校验和，BT 不使用这个值，只是为了兼容一些程序所保留!（字符串）；
name	|文件名（字符串）；
piece length|	每个块的大小，单位字节（整数）， 块长一般来说是 2 的权值；
pieces	|每个块的 20 个字节的 SHA1 Hash 的值（二进制格式）。

多文件info结构如下
-|-|-
---|---|---
files	|一个字典结构；|	
length	|文件长度，单位字节（整数）；	|
md5sum（可选）	|与单文件结构中相同；|	
path	|文件的路径和名字，是一个列表结构，如\test\test。txt 列表为l4	test8test。txte；|
-|path列表	|
-|name	|最上层的目录名字（字符串）；
-|piece length	|与单文件结构中相同；
-|pieces|	与单文件结构中相同。

### 下载基本原理
1. 下载软件拿到.torrent文件后，先进行打开，读取里面的这些信息，载入内存。
1. torrent中有Tracker的地址，下载软件拿到后，会去跟Tracker进行通讯，告诉Tracker：我要下载这个文件（通过hash值作为标记）； Tracker收到请求后，会记录这个客户端的公网IP（记录这厮在下载这个文件），同时呢，会返回给他：我这边还知道哪些人也在下载这个文件，一般是会返回200个IP（如果不够，当然就有多少返回多少）。当然了，如果下载过程中，协议要求你必须5分钟跟tracker通讯一次，如果太久不通讯，tracker就认为你下线了，会把你从节点列表中删除的。

1. 客户端拿到了一堆IP后，就开始挨个去尝试连接，连上后就开始互相通讯了。比如告诉对方，我有哪些分块，问问对方有哪些，然后把我有的给对方；让对方把他有的某一块给我，这样就你来我往开始了下载。当然，如果很悲催的情况下，此时没别人在线，那就只能没速度了，就只能不停的找啊找啊找朋友，直到找到一个好朋友。

1. 当然，如果torrent中有一个P2SP的Http地址辅助下载，那么也可以同时从这个Http服务器要数据，也会把这个服务器当成一个普通的节点，每次要1块数据，通过Http协议里面的Range标记，指定只要一部分数据过来辅助下载。


> 1. 如果Tracker服务器出问题了，连不上这个问询的服务器，就拿不到周围的邻居节点，怎么办？—NB的BT发明者提出了DHT(Distributed Hash Table)的概念，就算Tracker连不上了，也可以通过分布式哈希表DHT技术，通过DHT网络慢慢的寻找志同道合的邻居节点，只是没有Tracker那么直接那么快速，但慢一些总还是有机会找到邻居的。
> 其原理就是将原来由 Tracker 服务器保存的“文件哈希 - 文件存储位置” 的映射信息分散存储在 DHT 网络的各个节点中，并且留有冗余，即多份，以保证单个节点在关机之后，也不会影响文件的查询。
> 1. 网络是复杂的，特别是各个聪明的运营商，为了不让自己的用户消耗太多带宽，很多地区的运营商对P2P是有封锁的，比如某城宽带等。他们的做法早期是分析协议里面的握手消息，BT的握手消息是明文的Bittorrent Protocol，粗暴的运营商看到刚建立完连接就发这个明文会立即断开连接；文明点的运营商看到后不断开连接，但会限速到20K让你慢慢下载。当然，BT后来也发明了加密协议，运营商也升级了封锁的设备，也开始模拟自己是一个客户端，尝试分析加密后的协议，精彩纷呈。所以，要做一个稳定的靠谱的P2P系统还是有不少坑要趟的。


## Torrent文件与Magnet
磁力链接与Torrent文件是可以相互转换的，此文只讨论根据Torrent文件如何转换为Magnet磁力链接。

### Magnet概述
磁力链接由一组参数组成，参数间的顺序没有讲究，其格式与在HTTP链接末尾的查询字符串相同。最常见的参数是"xt"，是"exact topic"的缩写，通常是一个特定文件的内容散列函数值形成的URN，例如: `magnet:?xt=urn:bith:YNCKHTQCWBTRNJIV4WNAE52SJUQCZO5C`
注意，虽然这个链接指向一个特定文件，但是客户端应用程序仍然必须进行搜索来确定哪里，如果有，能够获取那个文件（即通过DHT进行搜索，这样就实现了Magnet到Torrent的转换，本文不讨论）。

部分字段名见下方表格：
字段名 |含义
---|---
magnet|协议名
xt|exact topic的缩写，包含文件哈希值的统一资源名称。BTIH（BitTorrent Info Hash）表示哈希方法名，这里还可以使用ED2K，AICH，SHA1和MD5等。这个值是文件的标识符，是不可缺少的。
dn|display name的缩写，表示向用户显示的文件名。这一项是选填的。
tr|tracker的缩写，表示tracker服务器的地址。这一项也是选填的。
bith|BitTorrent info hash，种子散列函数

> [Torrent文件的解析与转换](https://cloud.tencent.com/developer/article/1948355)

### ed2k
ed2k 全称 eDonkey2000 即“电驴”。与 BT 1.0 是同一时期出现，且技术原理相同的软件。但它在美国唱片协会的状告之下，被迫于 2005 年，也就是 BT 2.0 (DHT技术)元年关闭了网站，并停止了软件更新。

> 缺点是不能像种子一样保存多个文件

### 参考
https://zhuanlan.zhihu.com/p/560791482

## Tracker
Tracker是追踪器的意思
Tracker服务器是对于BT下载必须的

将你要下载的资源提供给服务器，服务区自动匹配有资源的人和需要下载的人。你的下载速度都会、上传速度都会增加。
运行于服务器上的一个程序，这个程序能够追踪到底有多少人同时在下载同一个文件。 客户端连上tracker服务器，就会获得一个下载人员的名单，根据这个，BT会自动连上别人的机器进行下载。它是提供bt的服务器。把文件用bt发布出来的人需要知道该使用哪个服务器来为要发布的文件提供tracker。由于不指定服务器，BitTorrent采用BT文件来确定下载源。

> 这个功能相对bt下载有效，pt下载不建议开启！
> https://github.com/XIU2/TrackersListCollection


### BtTracker原理
1. 做种
现在很多BT软件都提供了做种功能，在做种时，我们都必须指定tracker服务器地址，如果该地址无效，则做出来的种子对BT协议来说是没有任何实际意义的。

1. bt tracker服务
对于纯BT协议来说，每个BT网络中至少要有一台Tracker服务器（追踪服务器），tracker主要基本工作有以下几个方面：
    - 记录种子信息（torrent文件信息）
    - 记录节点信息
    - 计算并返回节点列表给BT客户端

每次我们利用BT软件做完种子后，总要找个论坛之类的来上传自己的种子，这样别人就可以下载到这个种子。为什么要上传种子呢？原因：
- 上传种子，其实就是把种子信息记录到tracker服务器上
- 种子可以在论坛传播，种子的扩展程度就决定了种子的健康度和下载度

当其他用户用BT软件打开种子后，BT软件会对种子进行解析（BDecode），主要得到种子的相关信息，包括：文件名、文件大小、tracker地址等。然后BT软件会向tracker地址发送请求报文，开始进行下载。BT向tracker发送的是Get请求，请求的内容主要有以下几个方面：

-|-|-
---|---|---
info_hash|必填|种子文件info字段的SHA1值（20字节）
peer_id|必填|节点标识，由BT客户端每次启动时随机生成
port|必填|节点端口，主要用于跟其他节点交互
uploaded|必填|总共上传的字节数，初始值为0
downloaded|必填|总共下载的字节数，初始值为0
left|必填|文件剩余的待下载字节数
numwant|必填|BT客户端期望得到的节点数
ip|选填|BT客户端IP，选填的原因是Tracker可以得到请求的IP地址，不需要客户端直接上传
event|选填|started/stopped/completed/空。当BT客户端开始种子下载时，第一个发起的请求为started，在下载过程中，该值一直为空，直到下载完成后才发起completed请求。做种过程中，发送的event也为空。如果BT客户端停止做种或退出程序，则会发起stopped请求。

**tracker收到该请求后主要进行以下几步处理**：

1. 根据info_hash查找种子信息，如果tracker没有该种子的任何信息，tracker服务器可以返回错误或返回0个种子数

2. 如果tracker找到了种子信息，接下来就会去查找是否数据库中已存在该peer_id的节点。接下来根据event的值进行相关处理。

3. 如果event是stopped，说明该节点已不可用，系统会删除tracker上关于该节点的记录信息。

4. 如果event是completed，说明种子节点+1，非种子-1。

5. 如果event是started，说明这是种子第一次连接tracker，tracker需要记录该节点信息，此外如果left=0，说明这是一个种子节点。

6. 如果event是空，则说明节点正在下载或上传，需要更新tracker服务器上该节点的信息。

7. 最后tracker从本地挑选出numwant个节点信息返回给BT客户端，实际返回的节点数不一定就是numwant，tracker只是尽量达到这个数量。


**Tracker响应**
Tracker正常返回的信息结构主要是：
-|-|-
---|---|---
interval|必填|请求间隔（秒）
complete|选填|种子节点数
Incomplete|选填|非种子节点数
peers:ip|必填|IP地址
peers:peer_id|选填|节点标识
peers:port|必填|端口

如果Tracker检查发现异常，可以返回错误信息：
-|-
---|---
failure reason| 错误原因

**Tracker如何挑选种子节点并返回给客户端？**
最普遍也是最简单的方式，那就是随机返回，tbsource采用的就是随机返回的机制。不少研究论文也提出了相关的算法，如IP地址策略和阶段返回策略。

IP地址策略是指根据IP地址所含拓扑信息来判断两个节点的距离，从而返回距离请求节点较近的节点列表。该方法主要适用于IPV6。

阶段返回策略，根据节点的下载进度，返回下载进度相近的节点列表。

Bt协议中，有两个策略可以用来提高整个BT网络的健壮性和下载速度，它们分别是：
**最少片段优先策略（BT客户端处理）** 和 **最后阶段模式**。为了响应“最后阶段模式”，当种子节点的下载进度大于80%（个人指定）时，tracker服务器应该尽量返回种子节点给客户端，帮助客户端尽快完成下载，使其成为种子节点。


> 下载进度达到 100% 以后，不会立即完成下载，还会显示“做种中”或者“上传中”，这是 BT 客户端在告诉 Tracker 服务器，你的电脑也拥有了这个文件资源，Tracker 服务器则会将你的 IP 地址也加入此文件哈希值对应的 IP 地址列表中。之后再有其他人要下载该文件，就也会同时从你这里下载了。

### private tracker原理

Privatetracker简称PT，目前主要应用于高清视频下载。其实PT就是“我为人人，人人为我”这个目标的最佳实践者。在实际的BT下载过程中，用户通过种子下载完文件后，出于“自私”的考虑（怕占用自己带宽），往往会退出做种，从而降低种子的热度。这就是为什么一个种子过了一段时间后，往往下载速度很慢或下载不完。

为了真正地实现BT理念，PT强制每个下载者必须上传一定量数据后，才能进行下载。如何保证这种行为呢？

现在的PT一般存在于网络社区中，每个注册网络社区的用户都会分配到一个随机的KEY，任何从社区下载的种子，都会包含用户的KEY。每次用户通过种子下载时，都会连接到社区的tracker服务器上，tracker服务器会检查KEY对应用户的上传下载量，如果上传量不满足标准，则tracker服务器会记录相关信息，并对该用户的下载及社区活动进行相关限制。


### 服务器上部署tracker服务
种子下载源需部署到公网服务器上，部署在内网服务器上没见成功下载过。
OpenTracker/BitTorrent3.4.2/BitTorrent5.0.9/BitTornado任选其一，建议[OpenTracker](https://github.com/geolink/opentracker/)

> [chihaya](https://github.com/chihaya/chihaya) ,golang multi-protocol BitTorrent Tracker

> https://blog.csdn.net/Jailman/article/details/86006844
> https://qifcn.github.io/post/c2bc3eba

> 安装opentracker可以参考这个链接看看:http://erdgeist.org/arts/software/opentracker/#overview

自行编译
```
# 搭建环境
yum -y install epel-release
yum -y groupinstall "Development Tools"
yum -y install openssl-devel zlib-devel

# 下载相关文件
git clone https://github.com/Qifcn/OpenTracker.git
cd OpenTracker
tar -xzvf libowfat.tar.gz
tar -xzvf opentracker.tar.gz

# 编译
cd libowfat
make
cd ..
cd opentracker
make

# 二进制文件移动到/usr/bin：
cp opentracker /usr/bin

```
`/usr/bin/opentracker -p 1337 -P 1337`

开机启动
```
vim /usr/lib/systemd/system/opentracker.service

[Unit]
Description=opentracker server
    
[Service]
User=root
ExecStart=/usr/bin/opentracker -p 1337 -P 1337
Restart=on-abort
    
[Install]
WantedBy=multi-user.target
```

```
systemctl enable opentracker.service  
systemctl start opentracker.service
```


制作种子的时候，我们添加Tracker服务器时，可以这么写： 
`http://你的服务器公网IP:1337/announce`
或者： `udp://你的服务器公网IP:1337/announce`

当然也可以同时添加http和udp，这个不受影响。注意:udp的兼容性不好，旧的bt客户端不支持，如BitTorrent3.4.2
OpenTracker还有一个自带的统计功能： `http://你的服务器公网IP:1337/stats`

更详细的统计信息访问： `http://你的服务器公网IP:1337/statsmode=everything`

### 安装ctorrent(种子提供者和初始上传者)
Tracker服务器（使用opentracker）
种子提供者和初始上传者（ctorrent做种）
`apt install ctorrent -y`
制作torrent种子: `ctorrent -t -u "http://trackerip:port(6969)/announce" -s yourfile.tgz.torrent yourfile.tgz`
启动初始上传服务:`ctorrent yourfile.tgz.torrent` 让制作的种子和文件存在于同一文件夹中，开启服务之后就会自动做种

下载: `ctorrent yourfile.tgz.torrent`

## 实现参考
### Taipei Torrent 

[Taipei-Torrent使用](https://blog.csdn.net/Jailman/article/details/87287871)
[根据Taipei-Torrent修改出的三个bt工具](https://blog.csdn.net/Jailman/article/details/87380636)
[种子发布和bt文件分发系统](https://blog.csdn.net/Jailman/article/details/88570395)

> 一般过程: 1. 启动tracker; 2. 制作种子文件; 3. 传播(发布)种子(当然还要把该种子文件对应的资源所在电脑开机连入网络一段时间，具体时间不定，主要看该种子下载的情况，最好至少要保证有一部分人下载成功。), 也就是做种