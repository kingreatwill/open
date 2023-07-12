<!--toc-->
[TOC]
#  多媒体（Multimedia）
一般包括文本，声音和图像等多种媒体形式。

## 流媒体工具
- Emby
- Plex
- https://github.com/jellyfin/jellyfin
## 视频 Video
- 画面更新率
每秒6或8张（frame per second，简称fps）至现今的每秒120张不等
- 分辨率
- 压缩技术
- 色彩
- 格式
    - MPEG/MPG/DAT
    MPEG也是Motion Picture Experts Group 的缩写。这类格式包括了MPEG-1, MPEG-2 和MPEG-4在内的多种视频格式。
    VCD 都是用MPEG1 格式压缩的(刻录软件自动将MPEG-1转为.DAT格式 ) ，使用MPEG-1 的压缩算法，可以把一部120 分钟长的电影压缩到1.2 GB 左右大小。
    MPEG-2 则是应用在DVD 的制作
    - AVI
    AVI（Audio Video Interleaved，音频视频交错）)由是Microsoft公司推出的视频音频交错格式（视频和音频交织在一起进行同步播放），是一种桌面系统上的低成本、低分辨率的视频格式。它的一个重要的特点是具有可伸缩性，性能依赖于硬件设备。它的优点是可以跨多个平台使用，缺点是占用空间大。
    - RA/RM/RAM
    RM，Real Networks  公司所制定的音频/视频压缩规范Real Media中的一种，Real Player能做的就是利用Internet资源对这些符合Real Media技术规范的音频/视频进行实况转播。
    在Real Media规范中主要包括三类文件：RealAudio、Real Video和Real Flash （Real Networks公司与Macromedia公司合作推出的新一代高压缩比动画格式）。
    REAL VIDEO （RA、RAM）格式由一开始就是定位就是在视频流应用方面的，也可以说是视频流技术的始创者。它可以在用56K MODEM 拨号上网的条件实现不间断的视频播放，可是其图像质量比VCD差些，如果您看过那些RM压缩的影碟就可以明显对比出来了。
    - MOV
    使用过Mac机的朋友应该多少接触过QuickTime。QuickTime原本是Apple公司用于Mac计算机上的一种图像视频处理软件。Quick-Time提供了两种标准图像和数字视频格式， 即可以支持静态的PIC和JPG图像格式，动态的基于Indeo压缩法的MOV和基于MPEG压缩法的MPG视频格式。
    - ASF
    ASF (Advanced Streaming format高级流格式)。ASF 是MICROSOFT 为了和Real player 竞争而发展出来的一种可以直接在网上观看视频节目的文件压缩格式。
    ASF使用了MPEG4 的压缩算法，压缩率和图像的质量都很不错。因为ASF 是以一个可以在网上即时观赏的视频“流”格式存在的，所以它的图像质量比VCD 差一点点并不出奇，但比同是视频“流”格式的RAM 格式要好。
    - WMV
    一种独立于编码方式的在Internet上实时传播多媒体的技术标准，Microsoft公司希望用其取代QuickTime之类的技术标准以及WAV、AVI之类的文件扩展名。
    WMV的主要优点在于：可扩充的媒体类型、本地或网络回放、可伸缩的媒体类型、流的优先级化、多语言支持、扩展性等。
    - n AVI
    如果你发现原来的播放软件突然打不开此类格式的AVI文件，那你就要考虑是不是碰到了n AVI。n AVI是New AVI 的缩写，是一个名为Shadow Realm 的地下组织发展起来的一种新视频格式。
    它是由Microsoft ASF 压缩算法的修改而来的（并不是想象中的AVI），视频格式追求的无非是压缩率和图像质量，所以 NAVI 为了追求这个目标，改善了原始的ASF 格式的一些不足，让NAVI 可以拥有更高的帧率。
    可以这样说，NAVI 是一种去掉视频流特性的改良型ASF 格式。
    - DivX
    这是由MPEG－4衍生出的另一种视频编码(压缩)标准，也即通常所说的DVDrip格式，它采用了MPEG4的压缩算法同时又综合了MPEG-4与MP3各方面的技术，说白了就是使用DivX压缩技术对DVD盘片的视频图像进行高质量压缩，同时用MP3或AC3对音频进行压缩，然后再将视频与音频合成并加上相应的外挂字幕文件而形成的视频格式。
    其画质直逼DVD并且体积只有DVD的数分之一。这种编码对机器的要求也不高，所以DivX视频编码技术可以说是一种对DVD造成威胁最大的新生视频压缩格式，号称DVD杀手或DVD终结者。
    - RMVB
    这是一种由RM视频格式升级延伸出的新视频格式，它的先进之处在于RMVB视频格式打破了原先RM格式那种平均压缩采样的方式，在保证平均压缩比的基础上合理利用比特率资源，就是说静止和动作场面少的画面场景采用较低的编码速率，这样可以留出更多的带宽空间，而这些带宽会在出现快速运动的画面场景时被利用。
    这样在保证了静止画面质量的前提下，大幅地提高了运动图像的画面质量，从而图像质量和文件大小之间就达到了微妙的平衡。另外，相对于DVDrip格式，RMVB视频也是有着较明显的优势，一部大小为700MB左右的DVD影片，如果将其转录成同样视听品质的RMVB格式，其个头最多也就400MB左右。
    不仅如此，这种视频格式还具有内置字幕和无需外挂插件支持等独特优点。要想播放这种视频格式，可以使用RealOne Player2.0或RealPlayer8.0加RealVideo9.0以上版本的解码器形式进行播放。
    - FLV
    FLV就是随着Flash MX的推出发展而来的新的视频格式，其全称为Flashvideo。是在sorenson公司的压缩算法的基础上开发出来的。
    由于它形成的文件极小、加载速度极快，使得网络观看视频文件成为可能，它的出现有效地解决了视频文件导入Flash后，使导出的SWF文件体积庞大，不能在网络上很好的使用等缺点。各在线视频网站均采用此视频格式。如新浪播客、56、优酷、土豆、酷6、帝途、YouTuBe等，无一例外。
    - F4V
    F4V是Adobe公司为了迎接高清时代而推出继FLV格式后的支持H.264的流媒体格式。它和FLV主要的区别在于，FLV格式采用的是H.263编码，而F4V则支持H.264编码的高清晰视频，码率最高可达50Mbps。
    主流的视频网站(如奇艺、土豆、酷6)等网站都开始用H.264编码的F4V文件，H.264编码的F4V文件，相同文件大小情况下，清晰度明显比On2 VP6和H.263编码的FLV要好。土豆和56发布的视频大多数已为F4V，但下载后缀为FLV，这也是F4V特点之一。
    - MP4
    MP4(MPEG-4 Part 14)是一种常见的多媒体容器格式，它是在“ISO/IEC 14496-14”标准文件中定义的，属于MPEG-4的一部分，是“ISO/IEC 14496-12(MPEG-4 Part 12 ISO base media file format)”标准中所定义的媒体格式的一种实现，后者定义了一种通用的媒体文件结构标准。
    MP4是一种描述较为全面的容器格式，被认为可以在其中嵌入任何形式的数据，各种编码的视频、音频等都不在话下，不过我们常见的大部分的MP4文件存放的AVC(H.264)或MPEG-4(Part 2)编码的视频和AAC编码的音频。MP4格式的官方文件后缀名是“.mp4”，还有其他的以mp4为基础进行的扩展或者是缩水版本的格式，包括：M4V, 3GP, F4V等。
    - 3GP
    3GPP（3rd Generation Partnership Project，第三代合作伙伴项目）制定的流媒体视频文件格式，主要是为了配合3G网络的高传输速度而开发的，也是目前手机中最为常见的一种视频格式。
    - AMV
    一种mp4专用的视频格式。
#### 格式转换
常见的视频格式无非就是两大类：
1. 影像格式（Video）
2. 流媒体格式（Stream Video）。

在影像格式中还可以根据出处划分为三大种：
1. AVI格式：这是由微软（Microsoft）提出，具有“悠久历史”的一种视频格式
2. MOV格式：这是由苹果（Apple）公司提出的一种视频格式
3. MPEG/MPG/DAT：这是由国际标准化组织ISO(International Standards Organization)与IEC(International Electronic Committee)联合开发的一种编码视频格式。MPEG是运动图像压缩算法的国际标准，现已被几乎所有的计算机平台共同支持。

在流媒体格式中同样还可以划分为三种：
1. RM格式  ：这是由Real Networks公司开发的一种新型流式视频文件格式。
2. MOV格式：MOV也可以作为一种流文件格式。QuickTime能够通过Internet提供实时的数字化信息流、工作流与文件回放功能，为了适应这一网络多媒体应用，QuickTime为多种流行的浏览器软件提供了相应的QuickTime Viewer插件（Plug－in），能够在浏览器中实现多媒体数据的实时回放。
3. ASF格式：是由微软公司开发的流媒体格式，是一个在Internet上实时传播多媒体的技术标准。

### 视频编解码
- Microsoft RLE
一种8位的编码方式，只能支持到256色。压缩动画或者是计算机合成的图像等具有大面积色块的素材可以使用它来编码，是一种无损压缩方案。

- Microsoft Video 1
用于对模拟视频进行压缩，是一种有损压缩方案，最高仅达到256色，它的品质就可想而知，一般还是不要使用它来编码AVI。

- Microsoft H.261/H.263/H.264/H.265
- Microsoft MPEG-4 Video codec
- Intel Indeo Video R3.2
- Intel Indeo Video 4和5
- Intel IYUV Codec
- DivX?-　MPEG-4 Low-Motion/Fast-Motion
- DivX 3.11/4.12/5.0
- WebM/VP8
- Ogg-Theora


## 音频 Audio
指人耳可以听到的声音频率在20HZ~20kHz之间的声波，称为音频。
指存储声音内容的文件。
在某些方面能指作为滤波的振动。




- 采样频率
在一定时间内取的点越多，描述出来的波形就越精确，这个尺度我们就称为“采样频率”,常用的采样频率是44.1kHz 和 96kHz，它的意思是每秒取样44100次。
- 比特率
数码录音一般使用16比特、20比特或24比特制作音乐。
声音有轻有响，影响声音响度的物理要素是振幅，作为数码录音，必须也要能精确表示乐曲的轻响，所以一定要对波形的振幅有一个精确的描述。“比特(bit)”就是这样一个单位，16比特就是指把波形的振幅划为2^16即65536个等级，根据模拟信号的轻响把它划分到某个等级中去，就可以用数字来表示了。和采样频率一样，比特率越高，越能细致地反映乐曲的轻响变化。
“**动态**”指的是一首乐曲最响和最轻的对比能达到多少，我们也常说“动态范围”，单位是dB，而动态范围和我们录音时采用的比特率是紧密结合在一起的，如果我们使用了一个很低的比特率，那么就只有很少的等级可以用来描述音响的强弱，当然就不能听到大幅度的强弱对比了。动态范围和比特率的关系是；比特率每增加1比特，动态范围就增加6dB。所以假如我们使用1比特录音，那么我们的动态范围就只有6dB，这样的音乐是不可能听的。16比特时，动态范围是96dB。这可以满足一般的需求了。20比特时，动态范围是120dB，对比再强烈的交响乐都可以应付自如了，表现音乐的强弱是绰绰有余了。发烧级的录音师还使用24比特，但是和采样精度一样，它不会比20比特有很明显的变化，理论上24比特可以做到144 dB的动态范围，但实际上是很难达到的，因为任何设备都不可避免会产生噪音，至少在现阶段24比特很难达到其预期效果。

数码信号和模拟信号

音频数字化的标准是每个样本16位(16bit，即96dB)的信噪比，采用线性脉冲编码调制PCM，每一量化步长都具有相等的长度。在音频文件的制作中，正是采用这一标准。

wav使用的是数码信号，它是用一堆数字来描述原来的模拟信号

- 音频格式
    - CD格式 天籁
    当今世界上音质最好的音频格式是什么？当然是CD了。
    标准CD格式也就是44.1K的采样频率，速率88K/秒，16位量化位数，因为CD音轨可以说是近似无损的，因此它的声音基本上是忠于原声的，因此如果你如果是一个音响发烧友的话，CD是你的首选。它会让你感受到天籁之音。
    - WAV 无损
    是微软公司开发的一种声音文件格式，它符合 PIFFResource Interchange File Format 文件规范，用于保存WINDOWS平台的音频信息资源，被WINDOWS平台及其应用程序所支持。“*.WAV”格式支持MSADPCM、CCITT A LAW等多种压缩算法，支持多种音频位数、采样频率和声道，标准格式的WAV文件和CD格式一样，也是44.1K的采样频率，速率88K/秒，16位量化位数，看到了吧，WAV格式的声音文件质量和CD相差无几，也是PC机上广为流行的声音文件格式，几乎所有的音频编辑软件都“认识”WAV格式。
    这里顺便提一下由苹果公司开发的AIFF（Audio Interchange File Format）格式和为UNIX系统开发的AU格式，它们都和和WAV非常相像，在大多数的音频编辑软件中也都支持它们这几种常见的音乐格式。
    - MP3 流行
    MP3格式诞生于八十年代的德国，所谓的MP3也就是指的是MPEG标准中的音频部分，也就是MPEG音频层。根据压缩质量和编码处理的不同分为3层，分别对应`“*.mp1"/“*.mp2”/“*.mp3”`这3种声音文件。需要提醒大家注意的地方是：MPEG音频文件的压缩是一种有损压缩，MPEG3音频编码具有10：1~12：1的高压缩率，同时基本保持低音频部分不失真，但是牺牲了声音文件中12KHz到16KHz高音频这部分的质量来换取文件的尺寸，相同长度的音乐文件，用*.mp3格式来储存，一般只有*.wav文件的1/10，而音质要次于CD格式或WAV格式的声音文件。由于其文件尺寸小，音质好；所以在它问世之初还没有什么别的音频格式可以与之匹敌，因而为*.mp3格式的发展提供了良好的条件。直到现在，这种格式还是风靡一时，作为主流音频格式的地位难以被撼动。但是树大招风，MP3音乐的版权问题也一直是找不到办法解决，因为MP3没有版权保护技术，说白了也就是谁都可以用。
    MP3格式压缩音乐的采样频率有很多种，可以用64Kbps或更低的采样频率节省空间，也可以用320Kbps的标准达到极高的音质。我们用装有Fraunhofer IIS Mpeg Lyaer3的 MP3编码器（现在效果最好的编码器）MusicMatch Jukebox 6.0在128Kbps的频率下编码一首3分钟的歌曲，得到2.82MB的MP3文件。采用缺省的CBR（固定采样频率）技术可以以固定的频率采样一首歌曲，而VBR（可变采样频率）则可以在音乐“忙”的时候加大采样的频率获取更高的音质，不过产生的MP3文件可能在某些播放器上无法播放。我们把VBR的级别设定成为与前面的CBR文件的音质基本一样，生成的VBR MP3文件为2.9MB。
    - WMA 最具实力
    WMA (Windows Media Audio) 格式是来自于微软的重量级选手,高保真声音通频带宽，音质更好，后台强硬，音质要强于MP3格式，更远胜于RA格式，它和日本YAMAHA公司开发的VQF格式一样，是以减少数据流量但保持音质的方法来达到比MP3压缩率更高的目的，WMA的压缩率一般都可以达到1：18左右，WMA的另一个优点是内容提供商可以通过DRM（Digital Rights Management）方案如Windows Media Rights Manager 7加入防拷贝保护。
    这种内置了版权保护技术可以限制播放时间和播放次数甚至于播放的机器等等，这对被盗版搅得焦头乱额的音乐公司来说可是一个福音，另外WMA还支持音频流(Stream)技术，适合在网络上在线播放，作为微软抢占网络音乐的开路先锋可以说是技术领先、风头强劲，更方便的是不用象MP3那样需要安装额外的播放器，而Windows操作系统和Windows Media Player的无缝捆绑让你只要安装了windows操作系统就可以直接播放WMA音乐，新版本的Windows Media Player7.0更是增加了直接把CD光盘转换为WMA声音格式的功能，在新出品的操作系统Windows XP中，WMA是默认的编码格式，大家知道Netscape的遭遇，“狼”又来了。
    WMA这种格式在录制时可以对音质进行调节。同一格式，音质好的可与CD媲美，压缩率较高的可用于网络广播。虽然网络上还不是很流行，但是在微软的大规模推广下已经是得到了越来越多站点的承认和大力支持，在网络音乐领域中直逼*.mp3，在网络广播方面，也正在瓜分Real打下的天下。因此，几乎所有的音频格式都感受到了WMA格式的压力。
    - MIDI
    - RealAudio
    RA（RealAudio）、RM（RealMedia，RealAudio G2）、RMX（RealAudio Secured） ...
    - VQF
    - ogg
    - FLAC
    Free Lossless Audio Codec
    - AAC   .m4a
    （高级音频编码技术，Advanced Audio Coding）AAC号称「最大能容纳48通道的音轨，采样率达96 KHz，并且在320Kbps的数据速率下能为5.1声道音乐节目提供相当于ITU-R广播的品质」。和MP3比起来，它的音质比较好，也能够节省大约30%的储存空间与带宽。它是遵循MPEG-2的规格所开发的技术。

### 视听技术
- CD
- 蓝光
- 比特流
- 杜比
### 音频编解码
Speex, AAC, Ogg/Vorbis

### 音视频开发领域常用的9大开源项目
https://www.toutiao.com/i6841924924513190407/
- FFmpeg
FFmpeg是目前最全面的开源音视频编解码库，包括常用的音视频编码协议 H265、H264、MPEG4、H263、G.721、G.726、G.729等，并且它提供了一整套的音视频处理解决方案，包括音视频采集与编码、音视频解码、视频格式转换、视频抓图、给视频加水印等。
https://ffmpeg.org/

- WebRTC
WebRTC是一个由Google发起的开源音视频实时通讯解决方案，其中包括音视频的采集、编解码、网络传输、解码显示等，我们可以通过该技术快速地构建出一个音视频通讯应用。
虽然其名为WebRTC，但是实际上它不光支持Web之间的音视频通讯，还支持Windows、Android以及iOS等移动平台。WebRTC底层是用C/C++开发的，具有良好的跨平台性能。WebRTC因为其较好的音视频效果及良好的网络适应性，目前已被广泛的应用到视频会议系统中，比如视频会议厂商华为、ZOOM、小鱼易连、科达均支持了WebRTC方式的音视频转发会议。
https://webrtc.org/

- x264
x264是一个开源的H.264/MPEG-4 AVC视频编码函数库，是最好的有损视频编码器之一。H264是目前应用最广的码流标准，x264则能够产生符合H264标准的码流的编码器，它可以将视频流编码为H264、MPEG4 AVC格式。它提供了命令行接口与API，前者被用于一些图形用户接口例如Straxrip、MeGUI，后者则被FFmpeg、Handbrake等调用。当然，既然有x264，就有对应HEVC/H.265的x265。
https://www.videolan.org/developers/x264.html

- Live555
Live555是一个为流媒体提供解决方案的跨平台的C++开源项目，它实现了标准流媒体传输，是一个为流媒体提供解决方案的跨平台的C++开源项目，它实现了对标准流媒体传输协议如RTP/RTCP、RTSP、SIP等的支持。
Live555实现了对多种音视频编码格式的音视频数据的流化、接收和处理等支持，包括H265、H264、MPEG4、H.263+ 、DV、JPEG视频和多种音频编码。同时由于良好的设计，Live555非常容易扩展对其他格式的支持。
http://www.live555.com/

- SDL
SDL（Simple DirectMedia Layer）是一套开放源代码的跨平台多媒体开发库，使用C语言写成。SDL提供了数种控制图像、声音、输出入的函数，让开发者只要用相同或是相似的代码就可以开发出跨多个平台（Linux、Windows、Mac OS X等）的应用软件。目前SDL多用于开发游戏、模拟器、媒体播放器、视频会议系统等产品开发中。
https://www.libsdl.org

- Opus
Opus是用C语言开发的一个高灵活度的音频编码器，针对ARM、x86有特殊优化，fix-point实现。Opus在各方面都有着明显优势。它同时支持语音与音乐的编码，比特率为6k-510k。它融合了SILK编码方法和CELT编码方法。SILK原本被用于Skype中，基于语音信号的线性预测分析（LPC），对音乐支持并不好。而CELT尽管适用于全带宽音频，但对低比特率语音的编码效率不高，所以两者在Opus中形成了互补。
https://opus-codec.org/

- ffplay
ffplay是ffmpeg的一个子工具，所以其开源代码也是内置在FFmpeg项目中的。ffplay内部使用了FFmpeg和 SDL库，是一个简单的可移植的媒体播放器。它具有强大的音视频解码播放能力，目前它广泛被各种流行播放器（QQ影音、暴风影音……）集成应用。作为一款开源软件，ffplay囊括Linux、Windows、Ios、Android等众多主流系统平台，十分适合进行二次开发。
http://ffmpeg.org/ffplay.html

- VLC
VLC是一款自由、开源的跨平台多媒体播放器及框架，它可以播放来自网络、摄像头、磁盘、光驱的文件，支持包括MPEG4、H264、H265、DivX、WMV、Vorbis、AC3等多种音视频协议。VLC最为突出的就是流媒体文件的功能，VLC支持各种流媒体协议，能直接播放远端的流媒体视频，只要输入一个视频文件的网址即可，无需下载到本地。此外，VLC还可以直接播放没有下载完成的文件。
VideoLanServer(VLS)的功能已经合并到VLC中，所以VLC不仅仅是一个音视频播放器，它也可以作为小型的视频服务器或流媒体服务器使用，可以一边播放一边转码，把视频流发送到网络上。
https://www.videolan.org/

- ijkplayer
在介绍ijkplayer播放器之前，要先提到FFmpeg中的ffplay。ffplay是一个使用了FFmpeg和SDL库的可移植的媒体播放器。ijkplay是国内知名的视频弹幕网站Bilibili开源的基于ffplay.c实现的轻量级iOS/Android视频播放器，API易于集成，且编译配置可裁剪，利于控制安装包大小。
在编解码方面，ijkplayer支持视频软解和硬解，可以在播放前配置，但在播放过程中则不能切换。iOS和Android上视频硬解可分别使用大家熟悉的VideoToolbox和MediaCodec。但ijkplayer对音频仅支持软解。
https://github.com/Bilibili/ijkplayer

## 容器
很多多媒体数据流需要同时包含音频数据和视频数据，这时通常会加入一些用于音频和视频数据同步的元数据，例如字幕和弹幕。这三种数据流可能会被不同的程序，进程或者硬件处理，但是当它们传输或者存储的时候，这三种数据通常是被封装在一起的。通常这种封装是通过视频文件格 式来实现的，例如常见的*.mpg, *.avi, *.mov, *.mp4, *.rm, *.ogg or *.tta. 这些格式中有些只能使用某些编解码器，而更多可以以容器的方式使用各种编解码器。

FourCC全称Four-Character Codes，是由4个字符（4 bytes）组成，是一种独立标示视频数据流格式的四字节，在wav、avi档案之中会有一段FourCC来描述这个AVI档案，是利用何种codec来 编码的。因此wav、avi大量存在等于“IDP3”的FourCC。

　　视频是现在电脑中多媒体系统中的重要一环。为了适应储存视频的需要，人们设定了不同的视频文件格式来把视频和音频放在一个文件中，以方便同时回放。视频档实际上都是一个容器里面包裹着不同的轨道，使用的容器的格式关系到视频档的可扩展性。




## 流媒体（streaming media）
流媒体（streaming media）是指将一连串数据压缩后，经过网络分段发送，即时传输以供观看音视频的一种技术。

通过使用 streaming media 技术，用户无需将文件下载到本地即可播放。由于媒体是以连续的数据流发送的，因此在媒体到达时即可播放。可以像下载的文件一样进行暂停、快进或后退操作。

### 流协议 video streaming protocol

### RTSP  
RTSP（Real Time Streaming Protocol），RFC2326，实时流传输协议，是TCP/IP协议体系中的一个应用层协议，由哥伦比亚大学、网景和RealNetworks公司提交的IETF RFC标准。

（1）是流媒体协议。

（2）RTSP协议是共有协议，并有专门机构做维护。.

（3）RTSP协议一般传输的是 ts、mp4 格式的流。

（4）RTSP传输一般需要 2-3 个通道，命令和数据通道分离。

### RTMP
RTMP是Real Time Messaging Protocol（实时消息传输协议）的首字母缩写。该协议基于TCP，是一个协议族，包括RTMP基本协议及RTMPT/RTMPS/RTMPE等多种变种。RTMP是一种设计用来进行实时数据通信的网络协议，主要用来在Flash/AIR平台和支持RTMP协议的流媒体/交互服务器之间进行音视频和数据通信。支持该协议的软件包括Adobe Media Server/Ultrant Media Server/red5等。RTMP与HTTP一样，都属于TCP/IP四层模型的应用层。

（1）是流媒体协议。

（2）RTMP协议是 Adobe 的私有协议，未完全公开。

（3）RTMP协议一般传输的是 flv，f4v 格式流。

（4）RTMP一般在 TCP 1个通道上传输命令和数据。

rtmp://58.200.131.2:1935/livetv/hunantv # 湖南卫视
rtmp://media3.scctv.net/live/scctv_800  # CCTV

### HLS
[HLS (HTTP Live Streaming)](https://datatracker.ietf.org/doc/html/rfc8216)

[HTTP Streaming Architecture](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/StreamingMediaGuide/HTTPStreamingArchitecture/HTTPStreamingArchitecture.html)

[M3U8 文件示例](https://developer.apple.com/library/archive/technotes/tn2288/_index.html)

#### M3U
M3U 文件是一种纯文本文件，可以指定一个或多个多媒体文件的位置。它的设计初衷是为了播放音频文件，但后来越来越多的用于播放视频文件列表。

多个m3u8的索引文件


#### m3u8
ts文件索引`playlist.m3u8`
```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:5
#EXT-X-MEDIA-SEQUENCE:1
#EXTINF:4.042667,
/m3u8/seg1.ts
#EXTINF:4.000000,
/m3u8/seg2.ts
#EXTINF:4.000000,
/m3u8/seg3.ts
#EXTINF:4.000000,
/m3u8/seg4.ts
#EXTINF:4.000000,
/m3u8/seg5.ts
#EXTINF:4.000000,
/m3u8/seg6.ts
#EXT-X-ENDLIST

```
ts可以是相对路径, 也可以是绝对路径(http/https)

- #EXTM3U：文件头（必须放在第一行）
- #EXT-X-VERSION：版本号
- #EXT-X-TARGETDURATION：每个分片的最大时长（单位是秒）
- #EXT-X-MEDIA-SEQUENCE：首个分片的序列号（默认为0）
- #EXTINF：分片的信息，例如时长等
- #EXT-X-ENDLIST：文件结束符

用下面命令将其合并成一个大的 ts 文件，然后转成一个 mp4 文件：
```
$ ffmpeg -f concat -safe 0 -i playlist.m3u8 -c copy output.ts
$ ffmpeg -i output.ts -acodec copy -vcodec copy output.mp4
```

下面的命令直接将其保存为 mp4 文件：
```
$ ffmpeg -i http://cdn.zlib.cn/m3u8/test.m3u8 -c copy output.mp4
```

**为了保护数字版权，大部分的视频网站会对 ts 片段进行加密，例如下面的数据源：**
```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-MEDIA-SEQUENCE:1
#EXT-X-ALLOWCACHE:1
#EXT-X-KEY:METHOD=AES-128,URI="seg1.key"
#EXTINF:4.458667,
seg1.ts.enc
#EXT-X-KEY:METHOD=AES-128,URI="seg2.key"
#EXTINF:4.010000,
seg2.ts.enc
#EXT-X-KEY:METHOD=AES-128,URI="seg3.key"
#EXTINF:4.468667,
seg3.ts.enc
#EXT-X-KEY:METHOD=AES-128,URI="seg4.key"
#EXTINF:3.893000,
seg4.ts.enc
#EXT-X-KEY:METHOD=AES-128,URI="seg5.key"
#EXTINF:4.007333,
seg5.ts.enc
#EXT-X-TARGETDURATION:5
#EXT-X-ENDLIST
```
发现多了一些字段，例如：

- EXT-X-ALLOWCACHE：是否允许缓存
- #EXT-X-KEY：加密解析

EXT-X-KEY 就是表示经过加密的，基本格式形如：
```
#EXT-X-KEY:METHOD=AES-128,URI=”http://example.com/key",IV=0x0123456789abcdef0123456789abcdef
```
AES-128 且有 iv 填充的是 aes-cbc 算法，解码需要 iv (偏移量）和 key (秘钥），EXT-X-KEY 那一行记录了密钥的获取路径和偏移量的值，如果拿到正确的 key 和 iv 的话，可以用下面的命令进行解密：
```
$ openssl aes-128-cbc -d -in seg0.ts -out seg0.decode.ts -nosalt -iv xxx -K xxx
```
也可以用 ffmpeg 来进行快速处理：
```
$ ffmpeg -allowed_extensions ALL -i http://cdn.zlib.cn/m3u8/test.enc.m3u8 -c copy video5.mp4
```


**但是为了防止用户下载，一般会在 js 中对 key 进行动态处理，以气球云为例，它的 m3u8 播放源文件格式如下：**
```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:20
#EXT-X-ALLOW-CACHE:YES
#EXT-X-MEDIA-SEQUENCE:0
#EXT-X-KEY:METHOD=AES-128,URI="https://play.qiqiuyun.net/sdk_api/video/hls_clef/shd?resNo=xxx&token=xxx&ssl=1",IV=xxx
#EXTINF:10.200,
https://xxx-pub.pubssl.qiqiuyun.net/xxx/xxx?schoolId=xxx&fileGlobalId=xxx
#EXT-X-KEY:METHOD=
```

有的是直接
```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:5
#EXT-X-MEDIA-SEQUENCE:1
#EXTINF:4.042667,
/m3u8/seg1.ts?auth_key=xxxxxx
#EXT-X-ENDLIST
```

#### HLS 的直播请求详情(直播不可以回看)
1. 先通过循环的加载m3u8文件
2. 加载最新一次m3u8中包含的最新的ts切片，达到直播的效果的。之所以能循环加载，就是因为： 直播的m3u8文件中，没有 #EXT-X-ENDLIST 参数，也就是说，没有结束，需要一直加载。

第一次加载m3u8 的内容如下：
```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-MEDIA-SEQUENCE:31
#EXT-X-TARGETDURATION:5
#EXTINF:5.079,
31.ts
#EXTINF:5.083,
32.ts
#EXTINF:5.201,
33.ts
```
第二次加载m3u8 的内容是：
```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-MEDIA-SEQUENCE:32
#EXT-X-TARGETDURATION:5
#EXTINF:5.083,
32.ts
#EXTINF:5.201,
33.ts
#EXTINF:5.100,
34.ts
```
我们能看到，第一次加载m3u8时，ts切片，最后一个切片，名字是 33.ts；第二次加载m3u8时，ts切片，最后一个切片，名字是 34.ts ；
并且，m3u8文件中没有 #EXT-X-ENDLIST; 就是通过这种循环加载的方式，这个直播，能一直循环加载下去。

#### EVENT 播放列表(直播且可以回看)
> EXT-X-PLAYLIST-TYPE 标签（可选）值为 VOD，表示播放列表不可变。

刚开始时，对应的 M3U8 文件内容如下所示：
```
#EXTM3U
#EXT-X-PLAYLIST-TYPE:EVENT
#EXT-X-TARGETDURATION:10
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:10,
fileSequence0.ts
#EXTINF:10,
fileSequence1.ts
#EXTINF:10,
fileSequence2.ts
#EXTINF:10,
fileSequence3.ts
```
> EXT-X-PLAYLIST-TYPE 标签值为 EVENT，表示播放列表内容可变，不过只能在文件末尾改变。

结束时，M3U8 文件内容如下所示：
```
#EXTM3U
#EXT-X-PLAYLIST-TYPE:EVENT
#EXT-X-TARGETDURATION:10
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:10,
fileSequence0.ts
#EXTINF:10,
fileSequence1.ts
#EXTINF:10,
fileSequence2.ts
#EXTINF:10,
fileSequence3.ts
...
#EXTINF:10,
fileSequence120.ts
#EXTINF:10,
fileSequence121.ts
#EXT-X-ENDLIST
```
> **EVENT 播放列表可以用在直播中，通常用于晚会和体育赛事场景，用户一方面可以观看直播，一方面还能做 seek 操作回退到之前的时间点去回放。**


#### 其它
- EXT-X-BYTERANGE 我只想用一个ts来构建一个类似M3U8的分片索引, 这时候EXT-X-BYTERANGE就派上用场了.当然只有VERSION版本不低于4才可以应用这个属性

- EXT-X-DISCONTINUITY 标签
在一些场景下，我们需要在点播或直播中插入其他内容，比如广告
```
#EXTM3U
#EXT-X-TARGETDURATION:10
#EXT-X-VERSION:3
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:10.0,
ad0.ts
#EXTINF:8.0,
ad1.ts
#EXT-X-DISCONTINUITY
#EXTINF:10.0,
movieA.ts
#EXTINF:10.0,
movieB.ts
```

- EXT-X-STREAM-INF 标签
```
#EXTM3U
#EXT-X-STREAM-INF:BANDWIDTH=1280000,AVERAGE-BANDWIDTH=1000000
http://example.com/low.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=2560000,AVERAGE-BANDWIDTH=2000000
http://example.com/mid.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=7680000,AVERAGE-BANDWIDTH=6000000
http://example.com/hi.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=65000,CODECS="mp4a.40.5"
http://example.com/audio-only.m3u8
```
通过 3 个不同码率的视频流和 1 个音频流来描述一个内容。


### RTP
### RTCP 
### HTTP-FLV
`<object type="application/x-shockwave-flash" src="xxx.flv"></object>`
### HTTP
（1）不是是流媒体协议。

（2）HTTP协议是共有协议，并有专门机构做维护。 

（3）HTTP协议没有特定的传输流。 

（4）HTTP传输一般需要 2-3 个通道，命令和数据通道分离。

CCTV1高清：http://ivi.bupt.edu.cn/hls/cctv1hd.m3u8
CCTV3高清：http://ivi.bupt.edu.cn/hls/cctv3hd.m3u8
CCTV5高清：http://ivi.bupt.edu.cn/hls/cctv5hd.m3u8
CCTV5+高清：http://ivi.bupt.edu.cn/hls/cctv5phd.m3u8
CCTV6高清：http://ivi.bupt.edu.cn/hls/cctv6hd.m3u8
苹果提供的测试源（点播）：
http://devimages.apple.com.edgekey.net/streaming/examples/bipbop_4x3/gear2/prog_index.m3u8

### RTC(Real-time Communications)
实时通信

## 播放器推荐
### PotPlayer 
https://potplayer.daum.net/


非官网：https://daumpotplayer.com/download/
https://t1.daumcdn.net/potplayer/PotPlayer/Version/Latest/PotPlayerSetup.exe
https://t1.daumcdn.net/potplayer/PotPlayer/Version/Latest/PotPlayerSetup64.exe
### VLC
http://www.videolan.org/

## 参考
[电影的各种版本，格式，视频音频编码，字幕，你都了解吗？](https://www.toutiao.com/i6808380016141795843)
[各种音频视频编码方法](https://blog.csdn.net/metalseed/article/details/38471755)
[各种音视频编解码学习详解](https://www.cnblogs.com/skyofbitbit/p/3651270.html)
[五种常见流媒体协议](https://www.jianshu.com/p/d71ceef679de)

##  音频处理
Audiokit: audiokit/AudioKit
ffmpeg： FFmpeg/FFmpeg

