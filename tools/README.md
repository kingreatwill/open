chcp 65001 cmd 显示中文

[作为程序员的你，常用的工具软件有哪些？](https://www.zhihu.com/question/22867411/answer/923997976)

Cygwin，让你拥有Windows下的Linux环境

[html 转markdown](https://tool.lu/markdown/) 

[html 转markdown](https://github.com/domchristie/turndown )

[html 转markdown](http://domchristie.github.io/turndown/)

[my.cnf自动生成工具](https://imysql.com/my-cnf-wizard.html)

[excel to markdown](http://www.tablesgenerator.com/markdown_tables)

[BaiduPCS-Go](https://github.com/iikira/BaiduPCS-Go)

[代码比较工具](https://blog.csdn.net/yueliang2100/article/details/82190257)

[文件同步工具](https://freefilesync.org/)

[PDF ORC](https://lightpdf.com/zh/ocr)
[ABBYY FineReader 14] PDF ORC
Adobe Acrobat

[语言在线](https://repl.it/languages)

[Protoman](https://github.com/spluxx/Protoman)
Postman for protobuf APIs

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

[soar:优化mysql数据库复杂sql](https://www.toutiao.com/a6778396540256911884/)

https://github.com/XiaoMi/soar

[soar-web](https://github.com/xiyangxixian/soar-web)

[基于Inception & SQLAdvisor & SOAR SQL审核平台WEB](https://github.com/myide/see)

[Cheat Engine 内存修改器](https://www.cheatengine.org/)

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

chrome 插件
彩云小译 翻译 保留原文

## IDE 
```
docker run -it --init -p 3000:3000 -v "%cd%:/home/project:cached" theiaide/theia-full:1.0.0

docker run -it -d --init -p 3000:3000 -v "/root/code:/home/project:cached" theiaide/theia-full:1.0.0

# Linux or macOS
docker run -it --init -p 3000:3000 --expose 9229 -p 9229:9229 -v "$(pwd):/home/project:cached" theiaide/theia:next --inspect=0.0.0.0:9229

# Windows
docker run -it --init -p 3000:3000 --expose 9229 -p 9229:9229 -v "%cd%:/home/project:cached" theiaide/theia:next --inspect=0.0.0.0:9229
```

## Listary Pro
好用的搜索工具