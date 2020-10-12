chcp 65001 cmd 显示中文

[作为程序员的你，常用的工具软件有哪些？](https://www.zhihu.com/question/22867411/answer/923997976)

Cygwin，让你拥有Windows下的Linux环境

## markdown
### MarkdownMonster
Markdown编辑工具-支持weblog
### markdown在线转换工具
[html 转markdown](https://tool.lu/markdown/) 

[html 转markdown](https://github.com/domchristie/turndown )

[html 转markdown](http://domchristie.github.io/turndown/)

[table 转 md](https://tableconvert.com/)

[excel to markdown](http://www.tablesgenerator.com/markdown_tables)

## mysql
[my.cnf自动生成工具](https://imysql.com/my-cnf-wizard.html)

[soar:优化mysql数据库复杂sql](https://www.toutiao.com/a6778396540256911884/)

https://github.com/XiaoMi/soar

[soar-web](https://github.com/xiyangxixian/soar-web)

[基于Inception & SQLAdvisor & SOAR SQL审核平台WEB](https://github.com/myide/see)

## language playgroud
[语言在线](https://repl.it/languages)
[rust](https://play.rust-lang.org/)

## 其它小工具
[代码比较工具](https://blog.csdn.net/yueliang2100/article/details/82190257)

[文件同步工具](https://freefilesync.org/)

[PDF ORC](https://lightpdf.com/zh/ocr)
[ABBYY FineReader 14] PDF ORC
Adobe Acrobat

[小型实用的工具](https://github.com/l3m0n/pentest_tools)
- 信息收集
- 内网攻防
- 应急响应
- 权限攻防
- 漏洞利用
- 辅助工具
- 逆向工程/windows


[Protoman](https://github.com/spluxx/Protoman)
Postman for protobuf APIs
[BaiduPCS-Go](https://github.com/iikira/BaiduPCS-Go)

## 视频下载

[视频下载1](https://github.com/soimort/you-get)
```
pip3 install you-get
you-get "https://www.bilibili.com/video/av77403752?from=search&seid=8721276388614459679"
```
[视频下载2](https://github.com/ytdl-org/youtube-dl)
```
pip install --upgrade youtube-dl
youtube-dl "https://www.bilibili.com/video/av77403752?from=search&seid=8721276388614459679"
```



## Chrome
chrome 插件
彩云小译 翻译 保留原文

## 内存
[Cheat Engine 内存修改器](https://www.cheatengine.org/)

## PE
[LordPE](http://www.opdown.com/soft/84659.html) 主要包括PE文件分析、修改和脱壳三大功能

1、一个托管的PE文件包含4个部分：PE表头，CLR表头，元数据，IL代码。PE表头是Windows操作系统要求的标准信息，主要时指出了文件的类型，GUI、CUI或是DLL(不同于以前的Dynamic Link Library，特指程序集文件的一种形式)。

2、CLR 表头专门用于那些需要CLR才能运行的托管模块。CLR表头中包含和托管模块一起创建的元数据的主版本号和次版本号，一些标记，如果模块是GUI或 CUI，可执行文件还有一个标识入口点方法的MethodDef标记，以及一个可选的强命名数字签名。最后该表头中还包括模块内某些元数据表的大小和偏移量。

3、元数据是一个非常重要的概念，他实际上是一个二进制数据块。元数据中包含了一些表，这些表被划分为三大类：定义表、引用表和清单表。定义表包含了程序中的模块、方法、类型、字段、变量、属性、事件等等一切相关的定义信息，而引用表则记录了程序引用的程序集、方法、类型等的信息。清单表与 Assembly运行相关。

4、可以通过ILDasm来打开一个托管模块的元数据清单。在命令行中键入ILDasm \Adv App.exe来打开一个名为App.exe托管程序，在ILDasm的可视化界面中点击菜单--〉试图--〉元数据--〉显示就可以看到ILDasm处理过的元数据清单。IL代码，源程序被编译后成为中间语言代码，在ILDasm中也可以看到程序的IL代码。

[StudyPE]() 文件查看工具
1.支持查看 DosEXE、PE32、PE64。
2.全面山寨lordPE & Peid的功能。
3.提供PE区段、附加数据处理功能。
4.提供 RVA FOA 互相转换功能。
5.提供给 PE 增加 Api 函数功能。
6.提供资源中的图标另存、替换功能。
7.提供 PE 反汇编及反汇编比较功能。
8.有限的 PE 资源查看处理功能。
9.有限的图片及文本格式文件查看功能。

## 交互式反汇编专业工具 IDA
[逆向神器之IDA的使用](https://www.cnblogs.com/aikongmeng/articles/10657479.html)

## 密码管理工具
### KeePass 
KeePass的插件
- KeeAnywhere  - cloud storage providers 解压整个目录放入Plugins目录
    - Amazon Drive (Experimental: see reason here)
    - Amazon AWS S3
    - Box
    - Dropbox
    - Google Drive
    - HiDrive
    - hubiC
    - OneDrive
- KeePassRPC - Firefox和chrome数据同步  将KeePassRPC.plgx文件放入Plugins目录
    - 需要安装chrome插件https://github.com/pfn/passifox（chrome:chromeIPass（改名了KeePassHttp-Connector）   Firefox:PassIFox ）和https://github.com/pfn/keepasshttp/

## IDE 
```
docker run -it --init -p 3000:3000 -v "%cd%:/home/project:cached" theiaide/theia-full:1.0.0

docker run -it -d --init -p 3000:3000 -v "/root/code:/home/project:cached" theiaide/theia-full:1.0.0

# Linux or macOS
docker run -it --init -p 3000:3000 --expose 9229 -p 9229:9229 -v "$(pwd):/home/project:cached" theiaide/theia:next --inspect=0.0.0.0:9229

# Windows
docker run -it --init -p 3000:3000 --expose 9229 -p 9229:9229 -v "%cd%:/home/project:cached" theiaide/theia:next --inspect=0.0.0.0:9229
```
### Eclipse Theia
web IDE
https://github.com/eclipse-theia/theia


## Listary Pro
好用的搜索工具

## http://www.wox.one/
想必用过mac的人都会知道一款效率神器Alfred，可快速启动、计算、查找等，使用起来非常的方便。Win上面也有类似的效率神器，比如Listary和WOX。

## https://u.tools/
比everything更好？


## CMD
https://www.jianshu.com/p/4b2b7074d9e2

Windows Terminal
Fluent Terminal
Hyper
Terminus



## WinSCP
https://github.com/winscp/winscp

SFTP and FTP

## Linux 神器：bashtop

查看bash
bash --version
升级到4.4以上

```
wget http://ftp.gnu.org/gnu/bash/bash-5.0.tar.gz
解压缩：
tar zxvf bash-5.0.tar.gz
进入目录：
cd bash-5.0
开始编译：
./configure&&make&&make install
编译完成后，重启CentOS后，新版Bash生效。

虽然通过/bin/bash --version命令可以显示已经更新到最新版了，但是$BASH_VERSION变量依旧还是老版本，我们还需要加入下面的软链接：

mv /bin/bash /bin/bash.bak
ln -s /usr/local/bin/bash /bin/bash
再次重启系统
reboot
完成后echo $BASH_VERSION即可以显示为最新Bash版本了。
```


需要安装 yum install -y coreutils procps lm_sensors-libs

查看是否安装
rpm -qa|grep sensors


bashtop 是一个 Linux 资源监视器，显示处理器、内存、磁盘、网络和进程的使用情况和状态。特征：

易于使用，带有受游戏启发的菜单系统
快速响应的 UI，带有 UP，DOWN 键可进行过程选择
显示所选进程的详细统计信息
可过滤流程
在排序选项之间轻松切换
将 SIGTERM，SIGKILL，SIGINT 发送到选定的进程
可更改所有配置文件选项的 UI 菜单
自动缩放图显示网络使用情况
菜单直接显示是否有新版本可用
GitHub 地址→https://github.com/aristocratos/bashtop


## watchman 监控文件变化
其它机器可以连接查看
https://github.com/facebook/watchman

## 屏幕画笔工具

- Pointofix
http://www.pointofix.de/download.php
并复制文件到Pointofix的安装目录下，将文件pointofix_translation_zh-cn.ini改名为pointofix_translation.ini,重新启动Pointofix

- Scrpen
- 3 WhiteBoard
- 4 大鸿屏幕画笔
- 5 汉王电子白板
- 6 桌面小画笔

## ScreenToGif
https://github.com/NickeManarin/ScreenToGif
## Honeycam
动画GIF软件-免费版限制多

## 截图
https://github.com/lupoDharkael/flameshot

## 录屏
### Camtasia 收费的
Camtasia Studio是TechSmith旗下的一套专业屏幕录像软件，同时包含Camtasia 录像器、Camtasia Studio（编辑器）、Camtasia 菜单制作器、Camtasia 剧场、Camtasia 播放器和Screencast的内置功能。

### OBS Studio
Obs Studio 是一款开源的录屏软件,用于视频录制和直播。(既可以录制屏幕又可以录制摄像头)
https://github.com/obsproject/obs-studio 22k

https://obsproject.com/

### Captura 
https://github.com/MathewSachin/Captura 5k
捕捉屏幕，音频，光标，鼠标点击和击键
C#做的，现在已经归档了

### ScreenToGif
没有声音

## 图片处理
### Imagine 
Imagine 是一款开源免费的 图片压缩 程序，用于压缩 PNG、JPEG、WebP 格式的图片。
https://github.com/meowtec/Imagine

### Honeyview
图像查看器

主要功能
- 轻量而快速
- 可以显示包括 GPS 信息在内的 JPEG 格式的 EXIF 信息
- 对图像格式进行批量转换和调整
- 支持显示 GIF 和 WebP 动图
- 无需解压即可直接查看压缩包中的图像
支持的格式
- 图像格式: BMP, JPG, GIF, PNG, PSD, DDS, JXR, WebP, J2K, JP2, TGA, TIFF, PCX, PGM, PNM, PPM, BPG
- Raw 图像格式: DNG, CR2, CRW, NEF, NRW, ORF, RW2, PEF, SR2, RAF
- 动画图像格式: Animated GIF, Animated WebP, Animated BPG, Animated PNG
- 无需解压即可直接查看压缩包中的图像: ZIP, RAR, 7Z, LZH, TAR, CBR, CBZ

### ImageMagick 
https://imagemagick.org/script/download.php#windows

ImageMagick 是一个用来创建、编辑、合成图片的软件。它可以读取、转换、写入超过 200 种格式的图片，包括 PNG、JPEG、GIF、HEIC、TIFF、DPX、EXR、WebP、Postscript、PDF 和 SVG 等等。

ImageMagick 可被用于图片切割、颜色替换、各种效果的应用，图片的旋转、组合，文本，直线， 多边形，椭圆，曲线，附加到图片伸展旋转等。支持 Linux、Windows、Mac OS X、iOS、Android OS 平台。

## Bandizip 
ZIP 压缩包管理器 - 免费版含广告


## Microsoft Whiteboard
## Microsoft OneNote
其它类似工具
- [Evernote](https://evernote.com/)
- [Simplenote 3.2k](https://github.com/Automattic/simplenote-electron)
- [Laverna 8.6k](https://github.com/Laverna/laverna)
- [Standard Notes](https://github.com/standardnotes)
- [Turtl](https://github.com/turtl)
- [CherryTree 1.5k](https://github.com/giuspen/cherrytree)
- [TagSpaces 2.1k](https://github.com/tagspaces/tagspaces)
- [Google Keep](https://keep.google.com/)


## Microsoft Todo

## 视频音频处理软件
https://github.com/HaujetZhao/QuickCut

## 3D
- blender
https://github.com/blender/blender

- 3DMAX
- Maya - 动画
- C4D - Cinema 4D

- 动画人才：大部分公司会采用maya来制作动画，因此在人才需求上，会maya自然更有应聘优势。
- 影视人才：这个行业对软件专业性较高，包括AE、Nuke、RF等，大部分为具备工业级应用品质的商用软件，人才需求会需要更专业系统的培训，而这点上开源软件不一定能提供。
- 建筑或设计效果人才：客户只关注结果效果图，你用什么做都无所谓，不过目前大部分素材库和模型库都是Max，所以行业中使用Max的人很多也就不足为奇。
- 游戏人才：游戏美术如果是2D，那么情况同上，3转2的结果都是PNG或JPG，唯一的问题在于团队间的配合，如果做模型的人用Max，那么制作骨骼和动画的人用Maya就需要考虑兼容性配合。如果是3D，那么3D软件在于引擎的配合上，无外乎导出模型的属性兼容（坐标、UV、骨骼和动画等）和骨骼和动画的调用方式（IK还是只支持FK），因此如果有特殊需求，那么选什么3D软件做配合就看引擎开发者心情了。
- 其他应用（3D打印、虚拟现实等）：如果涉及开发，3D软件用哪一个就看需求匹配了，以及是否涉及授权费用和法律相关问题了。


## CPU和内存
PerfView是一个CPU和内存性能分析工具
https://github.com/microsoft/perfview

## 填补Win10缺失功能PowerToys
https://github.com/microsoft/PowerToys 32.9k

## 冰点下载器
可以自由下载百度，豆丁，道客巴巴，丁香，畅享网，it68，mbalib，mab.book118，open-open, 金字塔医学,大桔灯文库文档。 

https://share.weiyun.com/P63B9Qyy

## 视频会议
### jitsi-meet

开源视频会议服务
https://github.com/jitsi

## 数据库管理工具
- dbeaver 
    - https://github.com/dbeaver/dbeaver 
- navicat
    - http://bbs.sdbeta.com/read-htm-tid-566717-page-1.html
    - https://github.com/Deltafox79/Navicat_Keygen/releases
    - https://www.cnblogs.com/poloyy/p/12231357.html

## Web 调试神器
https://github.com/microsoft/playwright-cli
https://playwright.dev/
```

npm i -D playwright

npx playwright-cli --help

npx playwright-cli codegen wikipedia.org
# Open page in Chromium
npx playwright-cli open example.com
# Open page in WebKit
npx playwright-cli wk example.com
```

## cat 的替代品bat
https://github.com/sharkdp/bat 22k
- 可以高亮众多语言的语法，Markdown 中的代码也可以高亮；
- 和 Git 集成，注意下图左侧的 ~+；它能自动和 Git 沟通，识别出修改；
- 和 less 一样，自带分页；
- 其他 cat 的一些功能；


## 差异比较工具
Linux工具
- diff命令
    - 增强：colordiff命令和wdiff命令
- Vimdiff命令
    http://vimdoc.sourceforge.net/htmldoc/diff.html
- Kompare
    https://kde.org/applications/en/kompare
    Kompare是一个diff GUI包装器，允许用户查看文件之间的差异，并合并他们。 它的一些功能包括：
    - 支持多种差异格式
    - 支持目录比较
    - 支持读取diff文件
    - 可定制的界面
    - 创建和应用补丁到源文件

- DiffMerge
    https://sourcegear.com/diffmerge/
    DiffMerge是比较和合并文件的跨平台GUI应用程序。它有两个功能引擎，Diff引擎显示两个文件之间的差异，它支持行内突出显示和编辑，以及合并引擎，在三个文件之间输出更改的行。 它有以下特点：
    - 支持目录比较
    - 文件浏览器集成
    - 高度可配置
- Meld
    http://meldmerge.org/
    Meld是一个轻量级的GUI diff和合并工具。它使用户能够比较文件，目录和版本控制的程序。专为开发人员而开发，它具有以下功能：
    - 文件和目录的双向和三向比较
    - 更新文件比较作为用户键入更多的单词
    - 使用自动合并模式和更改块上的操作使合并更容易
    - 使用可视化进行简单比较
    - 支持Git，Mercurial，Subversion，Bazaar等等

- DiffUse 
    http://diffuse.sourceforge.net/
    DiffUse是另一种流行的，免费的，小而简单的GUI diff和合并工具，你可以在Linux上使用。在Python中，它提供两个主要功能，即：文件比较和版本控制，允许文件编辑，文件合并，还输出文件之间的差异。 您可以查看比较摘要，使用鼠标指针在文件中选择文本行，匹配相邻文件中的行并编辑不同的文件。其他功能包括：
    - 语法高亮显示
    - 键盘快捷键，便于浏览
    - 支持无限撤消
    - Unicode支持
    - 支持Git，CVS，Darcs，Mercurial，RCS，Subversion，SVK和Monotone

- XXdiff 
    http://furius.ca/xxdiff/
    XXdiff是一个免费的，功能强大的文件和目录比较和合并工具，可以在Unix类操作系统，如Linux，Solaris和HP / UX，IRIX，DEC Tru64上。 XXdiff的一个限制是它缺乏对unicode文件的支持和diff文件的内联编辑。 它具有以下功能列表：
    - 浅和递归比较两个，三个文件或两个目录
    - 水平差异突出显示
    - 交互式合并文件和保存结果输出
    - 支持合并评审/监管
    - 支持外部比较工具，如GNU diff，SIG diff，Cleareddiff等等
    - 可扩展使用脚本
    - 使用资源文件和许多其他次要功能可完全自定义
- KDiff3
    http://kdiff3.sourceforge.net/
    KDiff3是另一个很酷的，跨平台的差异和合并工具从KDevelop的制作。它适用于所有类Unix平台，包括Linux和Mac OS X，Windows。 它可以比较或合并两到三个文件或目录，并具有以下显着的功能：
    - 逐行和逐个字符指示差异
    - 支持自动合并
    - 内置编辑器来处理合并冲突
    - 支持Unicode，UTF-8和许多其他编解码器
    - 允许打印差异
    - Windows explorer集成支持
    - 还支持通过字节顺序标记“BOM”的自动检测
    - 支持手动对齐线条
    - 直观的GUI和更多
- TkDiff
https://sourceforge.net/projects/tkdiff/
TkDiff也是Unix的比较工具一个跨平台的，易于使用的GUI包装。它提供了两个输入文件之间的差异的并排视图。它可以在Linux，Windows和Mac OS X上运行。 此外，它还有一些其他令人兴奋的功能，包括差异书签，差异的图形地图方便快捷的导航等等。

Win工具
- BCompare
可以比较文件夹的差异
- UltraCompare 
可以比较文件夹的差异



**以下比较工具也可以直接在命令行中使用，也是git合并和差异比较工具**
直接在用户文件夹下，修改.gitconfig文件，修改名称并把path修改对应的工具路径
```
[diff]
	guitool = vscode
[difftool "vscode"]
	path = D:/Microsoft VS Code/Code.exe
	cmd = \"D:/Microsoft VS Code/Code.exe\" --wait --diff \"$LOCAL\" \"$REMOTE\"
[merge]
	guitool = vscode
[mergetool "vscode"]
	path = D:/Microsoft VS Code/Code.exe
	cmd = \"D:/Microsoft VS Code/Code.exe\" --wait \"$MERGED\"
```

git difftool # 比较当前所修改的内容
git difftool xxx  xxx # 比较两个commit id
```
C:\Users\Administrator>git difftool --tool-help
'git difftool --tool=<tool>' may be set to one of the following:
                vimdiff
                vimdiff2
                vimdiff3

        user-defined:
                vscode.cmd "D:/Microsoft VS Code/Code.exe" --wait "$MERGED"
                vscode.cmd "D:/Microsoft VS Code/Code.exe" --wait --diff "$LOCAL" "$REMOTE"

The following tools are valid, but not currently available:
                araxis
                bc
                bc3
                codecompare
                deltawalker
                diffmerge
                diffuse
                ecmerge
                emerge
                examdiff
                guiffy
                gvimdiff
                gvimdiff2
                gvimdiff3
                kdiff3
                kompare
                meld
                opendiff
                p4merge
                smerge
                tkdiff
                winmerge
                xxdiff

Some of the tools listed above only work in a windowed
environment. If run in a terminal-only session, they will fail.
```

### bc3
bcomp.exe file1 file2

### bc
bcomp.exe file1 file2

### diffmerge
DiffMerge.exe file1 file2

### kdiff3
kdiff3.exe file1 file2

### meld
meld.exe file1 file2

### p4merge
p4merge.exe file1 file2

### semanticmerge
semanticmergetool.exe -s file1 -d file2

### tortoisemerge
TortoiseGitMerge.exe file1 file2

### vscode
有空格是需要双引号的，文件也一样，如："file1"
"D:/Microsoft VS Code/Code.exe" --wait --diff file1 file2

### vsdiffmerge
vsdiffmerge.exe file1 file2

### winmerge
winmergeu.exe -e -u file1 file2

