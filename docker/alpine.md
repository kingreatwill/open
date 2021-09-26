[TOC]
# 制作 Docker 镜像时调整时区

## Alpine
[Setting the timezone](https://wiki.alpinelinux.org/wiki/Setting_the_timezone)
```
ENV TZ Asia/Shanghai

RUN apk add tzdata && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && apk del tzdata
```
## Debian
Debian 基础镜像 中已经安装了 tzdata 包，我们可以将以下代码添加到 Dockerfile 中：
```
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive

RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata \
    && rm -rf /var/lib/apt/lists/*
```
## Ubuntu
Ubuntu 基础镜像中没有安装了 tzdata 包，因此我们需要先安装 tzdata 包。

我们可以将以下代码添加到 Dockerfile 中。
```
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive

RUN apt update \
    && apt install -y tzdata \
    && ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata \
    && rm -rf /var/lib/apt/lists/*
```
## CentOS
CentOS 基础镜像 中已经安装了 tzdata 包，我们可以将以下代码添加到 Dockerfile 中。
```
ENV TZ Asia/Shanghai

RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone
```

# 中文环境

## Debian
```
RUN apt-get update && apt-get install -y locales
RUN localedef -c -f UTF-8 -i zh_CN zh_CN.utf8
# RUN rm -rf /var/lib/apt/lists/*
ENV LANG zh_CN.utf8
```

### 使用windows字体
> TTC是几个TTF合成的字库，安装后字体列表中会看到两个以上的字体。
> 虽然都是字体文件，但.ttc是microsoft开发的新一代字体格式标准，可以使多种truetype字体共享同一笔划信息，有效地节省了字体文件所占空间，增加了共享性。
> python 加载 fontforge 模块实现ttc字体文件 转换 为ttf字体文件，解析出每一个压缩字库中的ttf字库

> windows进入bash
> copy /mnt/c/Windows/Fonts/*.ttf /mnt/d/Fonts
> copy /mnt/c/Windows/Fonts/*.ttc /mnt/d/Fonts
> 然后把ttc转换成ttf

https://fontforge.org/docs/scripting/python.html
安装 fontforge(注意不是pip安装，是系统安装)
`apt-get install python-fontforge`
使用
`split_ttc_font_to_ttf.py Droid.ttc`
参考程序
```
import sys
import fontforge

fonts = fontforge.fontsInFile(sys.argv[1])
print(fonts)


for fontName in fonts:
    font = fontforge.open('%s(%s)'%(sys.argv[1], fontName))
    font.generate('%s.ttf'%fontName)
    font.close()
```
或者
`python ttc2ttf.py [TTC_file_Path]`
https://github.com/yhchen/ttc2ttf
或者(可以转换多个，也可以一个目录)【利用[FontTools](https://github.com/fonttools/fonttools)来转换的】
`pip install font-rename`
`font-rename insert/directory/file/path/here`
https://github.com/whtsky/font-rename
部分代码
```
from fontTools.ttLib import TTCollection
def unpack_ttc(filepath: Path) -> None:
    try:
        collection = TTCollection(str(filepath.resolve()))
    except:
        print(f"Failed to parse {filepath}, ignore")
        return
    for font in collection.fonts:
        ttf_path = filepath.parent / f"{get_font_name(font)}.ttf"
        font.save(ttf_path)
        print(f"{filepath} -> {ttf_path}")
    filepath.unlink()
```


node js
`npm i -g ttc2ttf`
`ttc2ttf <ttc file path> [output dir path]`
https://github.com/oysterlab/ttc2ttf


使用C#实现1
https://stackoverflow.com/questions/28225303/equivalent-in-c-sharp-of-pythons-struct-pack-unpack

使用C#实现2
`Install-Package IronPython -Version 3.4.0-alpha1`
`IronPython.Modules.PythonStruct.unpack_from`
使用
```csharp
Microsoft.Scripting.Hosting.ScriptRuntime runtime = IronPython.Hosting.Python.CreateRuntime();
Microsoft.Scripting.Hosting.ScriptEngine engine = runtime.GetEngine("py");
IronPython.Runtime.PythonContext cxt = (IronPython.Runtime.PythonContext)Microsoft.Scripting.Hosting.Providers.HostingHelpers.GetLanguageContext(engine);
IronPython.Runtime.PythonDictionary dict = new IronPython.Runtime.PythonDictionary();
IronPython.Runtime.ModuleContext modctx = new IronPython.Runtime.ModuleContext(dict, cxt);
IronPython.Runtime.CodeContext context = new IronPython.Runtime.CodeContext(dict, modctx);

var ttf_count = IronPython.Modules.PythonStruct.unpack_from(context,"!L", buf, 0x08)[0];
```

online
https://transfonter.org/ttc-unpack

工具
https://peter.upfold.org.uk/projects/dfontsplitter
fontforge（前面有讲用代码实现）【File > Open 然后 File > Generate Fonts...】



参考
https://stackoverflow.com/questions/60934639/install-fonts-in-linux-container-for-asp-net-core

- 使用mscorefonts（没有中文字体）
```
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-buster-slim AS base

#Add these two lines
RUN sed -i'.bak' 's/$/ contrib/' /etc/apt/sources.list
RUN apt-get update; apt-get install -y ttf-mscorefonts-installer fontconfig

WORKDIR /app
EXPOSE 80
```

- 或者 fonts-liberation（没有中文字体）
```
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-buster-slim AS base

#Add these two lines for fonts-liberation instead
RUN apt-get update; apt-get install -y fontconfig fonts-liberation
RUN fc-cache -f -v

WORKDIR /app
EXPOSE 80
```

- 也可以把windows字体拷贝过去安装
https://wiki.debian.org/Fonts#Adding_fonts

```
RUN apt-get -y install fontconfig
COPY /fonts ~/.fonts
COPY /fonts /usr/shared/fonts
COPY /fonts /usr/share/fonts/truetype
# refresh system font cache
RUN fc-cache -f -v
```
> fc-list :lang=zh

- 安装中文字体 以及安装 JRE LibreOfiice ImageMagick FFMPEG 环境
https://github.com/microestc/Containers
安装中字体， 在windows 下打包 C:/windows/fonts 字体 
```
FROM mcr.microsoft.com/dotnet/core/aspnet:3.0-buster-slim AS base
COPY fonts /usr/share/fonts/windows/
RUN echo "deb https://mirrors.tuna.tsinghua.edu.cn/debian/ buster main contrib non-free \
    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ buster main contrib non-free \
    deb https://mirrors.tuna.tsinghua.edu.cn/debian/ buster-updates main contrib non-free \
    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ buster-updates main contrib non-free \
    deb https://mirrors.tuna.tsinghua.edu.cn/debian/ buster-backports main contrib non-free \
    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ buster-backports main contrib non-free \
    deb https://mirrors.tuna.tsinghua.edu.cn/debian-security buster/updates main contrib non-free \
    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian-security buster/updates main contrib non-free" > /etc/apt/sources.list
RUN apt-get update && apt-get install fontconfig -y \
    && fc-cache -f -v
CMD ["/bin/bash"]
```
安装 Jre 环境 因为 libreoffice 需要

jre 这里用的是 jre-8u221-linux-x64.tar.gz 下载链接 https://sdlc-esd.oracle.com/ESD6/JSCDL/jdk/8u221-b11/230deb18db3e4014bb8e3e8324f81b43/jre-8u221-linux-x64.tar.gz?GroupName=JSC&FilePath=/ESD6/JSCDL/jdk/8u221-b11/230deb18db3e4014bb8e3e8324f81b43/jre-8u221-linux-x64.tar.gz&BHost=javadl.sun.com&File=jre-8u221-linux-x64.tar.gz&AuthParam=1569560572_a9dbaa64abfbf26b1afb9359ff05adb9&ext=.gz

```
ARG runenv=dotnetcoreaspnetcn
FROM $runenv:3.0-buster-slim AS base
COPY /src /src
ENV JAVA_HOME=/usr/java/jre1.8.0_221
ENV PATH=$PATH:/usr/java/jre1.8.0_221/bin
RUN mkdir /usr/java && tar -zxvf /src/jre-8u221-linux-x64.tar.gz -C /usr/java && rm -rf /src
CMD ["/bin/bash"]
```
前面用了中文字体的docker 镜像
建立的镜像名称 ：dotnetcoreaspnetcn-jre:3.0-buster-slim-8u221
接下里 build libreoffice
```
FROM dotnetcoreaspnetcn-jre:3.0-buster-slim-8u221 AS base
COPY /src /src
RUN apt-get update && apt-get install libxinerama1 dbus libsm6 libgio-cil libcairo2 libcups2 -y \
    && tar -zxvf /src/LibreOffice_6.2.7_Linux_x86-64_deb.tar.gz -C /src \
    && dpkg -i /src/LibreOffice_6.2.7.1_Linux_x86-64_deb/DEBS/*.deb \
    && rm -rf /src && ln -s /usr/local/bin/libreoffice6.2 /usr/local/bin/libreoffice
CMD ["/bin/bash"]
```
LibreOffice_6.2.7_Linux_x86-64_deb.tar.gz 自己去官方网站下载 
链接直接给了：https://mirror-hk.koddos.net/tdf/libreoffice/stable/6.2.7/deb/x86_64/LibreOffice_6.2.7_Linux_x86-64_deb.tar.gz

最后是建立 imagemagick  同样需要中文字体环境
```
FROM dotnetcoreaspnetcn:3.0-buster-slim AS base
COPY /src /src
RUN apt-get update && apt-get install imagemagick -y && cp -f /src/policy.xml /etc/ImageMagick-6/policy.xml && rm -rf /src
CMD ["/bin/bash"]
```
policy.xml
```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE policymap [
  <!ELEMENT policymap (policy)+>
  <!ATTLIST policymap xmlns CDATA #FIXED ''>
  <!ELEMENT policy EMPTY>
  <!ATTLIST policy xmlns CDATA #FIXED '' domain NMTOKEN #REQUIRED
    name NMTOKEN #IMPLIED pattern CDATA #IMPLIED rights NMTOKEN #IMPLIED
    stealth NMTOKEN #IMPLIED value CDATA #IMPLIED>
]>
<!--
  Configure ImageMagick policies.
  Domains include system, delegate, coder, filter, path, or resource.
  Rights include none, read, write, execute and all.  Use | to combine them,
  for example: "read | write" to permit read from, or write to, a path.
  Use a glob expression as a pattern.
  Suppose we do not want users to process MPEG video images:
    <policy domain="delegate" rights="none" pattern="mpeg:decode" />
  Here we do not want users reading images from HTTP:
    <policy domain="coder" rights="none" pattern="HTTP" />
  The /repository file system is restricted to read only.  We use a glob
  expression to match all paths that start with /repository:
    <policy domain="path" rights="read" pattern="/repository/*" />
  Lets prevent users from executing any image filters:
    <policy domain="filter" rights="none" pattern="*" />
  Any large image is cached to disk rather than memory:
    <policy domain="resource" name="area" value="1GP"/>
  Define arguments for the memory, map, area, width, height and disk resources
  with SI prefixes (.e.g 100MB).  In addition, resource policies are maximums
  for each instance of ImageMagick (e.g. policy memory limit 1GB, -limit 2GB
  exceeds policy maximum so memory limit is 1GB).
  Rules are processed in order.  Here we want to restrict ImageMagick to only
  read or write a small subset of proven web-safe image types:
    <policy domain="delegate" rights="none" pattern="*" />
    <policy domain="filter" rights="none" pattern="*" />
    <policy domain="coder" rights="none" pattern="*" />
    <policy domain="coder" rights="read|write" pattern="{GIF,JPEG,PNG,WEBP}" />
-->
<policymap>
  <!-- <policy domain="system" name="shred" value="2"/> -->
  <!-- <policy domain="system" name="precision" value="6"/> -->
  <!-- <policy domain="system" name="memory-map" value="anonymous"/> -->
  <!-- <policy domain="system" name="max-memory-request" value="256MiB"/> -->
  <!-- <policy domain="resource" name="temporary-path" value="/tmp"/> -->
  <!-- <policy domain="resource" name="memory" value="2GiB"/> -->
  <!-- <policy domain="resource" name="map" value="4GiB"/> -->
  <policy domain="resource" name="width" value="10MP"/>
  <policy domain="resource" name="height" value="10MP"/>
  <!-- <policy domain="resource" name="list-length" value="128"/> -->
  <!-- <policy domain="resource" name="area" value="100MP"/> -->
  <!-- <policy domain="resource" name="disk" value="16EiB"/> -->
  <!-- <policy domain="resource" name="file" value="768"/> -->
  <!-- <policy domain="resource" name="thread" value="4"/> -->
  <!-- <policy domain="resource" name="throttle" value="0"/> -->
  <!-- <policy domain="resource" name="time" value="3600"/> -->
  <!-- <policy domain="coder" rights="none" pattern="MVG" /> -->
  <!-- <policy domain="module" rights="none" pattern="{PS,PDF,XPS}" /> -->
  <policy domain="delegate" rights="none" pattern="HTTPS" />
  <!-- <policy domain="path" rights="none" pattern="@*" /> -->
  <!-- <policy domain="cache" name="memory-map" value="anonymous"/> -->
  <!-- <policy domain="cache" name="synchronize" value="True"/> -->
  <policy domain="cache" name="shared-secret" value="passphrase" stealth="true" />
</policymap>
```
最后一个 FFMPEG
```
FROM mcr.microsoft.com/dotnet/core/aspnet:3.0-buster-slim AS base
ADD ffmpeg-git-amd64-static.tar.xz /
RUN mv /ffmpeg*/ffmpeg /usr/bin/ && rm -rf /ffmpeg*
CMD [ "/bin/bash" ]
```
ffmpeg-git-amd64-static.tar.xz 自己去官方网站下载

## CentOS8

```
RUN dnf -y install langpacks-zh_CN.noarch \
	&& echo "LANG=\"zh_CN.UTF-8\"" > /etc/locale.conf \
	&& echo "LC_ALL=\"zh_CN.UTF-8\"" >> /etc/locale.conf
ENV LANG zh_CN.UTF-8
```

# dotnet 设置中文环境，上海时区，以及绘图
```
FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
RUN apt-get update && apt-get install -y libgdiplus locales
RUN localedef -c -f UTF-8 -i zh_CN zh_CN.utf8
ENV LANG zh_CN.utf8

ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive

RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata \
    && rm -rf /var/lib/apt/lists/*
```

# Alpine
## Alpine 镜像
a. 编辑 /etc/apk/repositories
b. 将里面 dl-cdn.alpinelinux.org 的 改成 mirrors.aliyun.com ; 保存退出即可

https://mirrors.aliyun.com/alpine/v3.12/

### alpine-python
https://github.com/jfloff/alpine-python/blob/master/3.8/Dockerfile

[alpine 3.10制作python 3.74的docker镜像](https://blog.csdn.net/weixin_42320932/article/details/100009617)

```
Dockerfile：
FROM docker.io/alpine:latest
MAINTAINER author:weifeng email:weifeng_nuaa@126.com
COPY Python-3.7.4.tar.xz /root
WORKDIR /root
RUN echo “http://mirrors.aliyun.com/alpine/latest-stable/main”>/etc/apk/repositories
&& apk add gcc libffi-dev python-dev openssl-dev make zlib-dev libc-dev
&& tar -xf Python-3.7.4.tar.xz
&& rm -rf Python-3.7.4.tar.xz \
&& cd Python-3.7.4
&& ./configure
&& make
&& make install
&& pip3 install --upgrade pip
&& rm -rf /root/Python-3.7.4
CMD /bin/sh
```
实践Alpine3.12
```
FROM python:3.8-alpine
...
apk add gcc g++ libffi-dev python3-dev openssl-dev make zlib-dev libc-dev
...
```
> g++ 编译grpcio的,pip install grpcio 运行了十分钟，所以有了以下解决方案

[用Alpine会让Python Docker的构建慢50倍](https://cloud.tencent.com/developer/news/600722)

#### Ubuntu和Debian的gcc安装
Ubuntu安装gcc
安装之前得在获得权限后终端输入
```
sudo apt-get update
#然后输入指令
sudo apt-get install gcc
```
Debian安装gcc
debian输入指令不需要像ubuntu一样前面加sudo
获得权限后输入
```
apt-get update
apt-get install build-essential
```
要apt-get install gcc下载gcc也行，不过build-essential依赖库已经包含了gcc

#### debian apt-get 国内常用 镜像源
```
sed -i "s@http://deb.debian.org@https://mirrors.163.com@g" /etc/apt/sources.list

sed -i "s@http://deb.debian.org/debian@http://mirrors.163.com/debian@g" /etc/apt/sources.list
sed -i "s@http://security.debian.org/debian-security@http://mirrors.163.com/debian-security@g" /etc/apt/sources.list

163镜像站
deb http://mirrors.163.com/debian/ stretch main non-free contrib
deb http://mirrors.163.com/debian/ stretch-updates main non-free contrib
deb http://mirrors.163.com/debian/ stretch-backports main non-free contrib
deb-src http://mirrors.163.com/debian/ stretch main non-free contrib
deb-src http://mirrors.163.com/debian/ stretch-updates main non-free contrib
deb-src http://mirrors.163.com/debian/ stretch-backports main non-free contrib
deb http://mirrors.163.com/debian-security/ stretch/updates main non-free contrib
deb-src http://mirrors.163.com/debian-security/ stretch/updates main non-free contrib

中科大镜像站
deb https://mirrors.ustc.edu.cn/debian/ stretch main contrib non-free
deb-src https://mirrors.ustc.edu.cn/debian/ stretch main contrib non-free 
deb https://mirrors.ustc.edu.cn/debian/ stretch-updates main contrib non-free
deb-src https://mirrors.ustc.edu.cn/debian/ stretch-updates main contrib non-free 
deb https://mirrors.ustc.edu.cn/debian/ stretch-backports main contrib non-free
deb-src https://mirrors.ustc.edu.cn/debian/ stretch-backports main contrib non-free 
deb https://mirrors.ustc.edu.cn/debian-security/ stretch/updates main contrib non-free
deb-src https://mirrors.ustc.edu.cn/debian-security/ stretch/updates main contrib non-free

阿里云镜像站
deb http://mirrors.aliyun.com/debian/ stretch main non-free contrib
deb-src http://mirrors.aliyun.com/debian/ stretch main non-free contrib
deb http://mirrors.aliyun.com/debian-security stretch/updates main
deb-src http://mirrors.aliyun.com/debian-security stretch/updates main
deb http://mirrors.aliyun.com/debian/ stretch-updates main non-free contrib
deb-src http://mirrors.aliyun.com/debian/ stretch-updates main non-free contrib
deb http://mirrors.aliyun.com/debian/ stretch-backports main non-free contrib
deb-src http://mirrors.aliyun.com/debian/ stretch-backports main non-free contrib

Ubuntu
deb http://mirrors.aliyun.com/ubuntu/ trusty main restricted universe multiverse 
deb http://mirrors.aliyun.com/ubuntu/ trusty-security main restricted universe multiverse 
deb http://mirrors.aliyun.com/ubuntu/ trusty-updates main restricted universe multiverse 
deb http://mirrors.aliyun.com/ubuntu/ trusty-proposed main restricted universe multiverse 
deb http://mirrors.aliyun.com/ubuntu/ trusty-backports main restricted universe multiverse 
deb-src http://mirrors.aliyun.com/ubuntu/ trusty main restricted universe multiverse 
deb-src http://mirrors.aliyun.com/ubuntu/ trusty-security main restricted universe multiverse 
deb-src http://mirrors.aliyun.com/ubuntu/ trusty-updates main restricted universe multiverse 
deb-src http://mirrors.aliyun.com/ubuntu/ trusty-proposed main restricted universe multiverse 
deb-src http://mirrors.aliyun.com/ubuntu/ trusty-backports main restricted universe multiverse


华为镜像站
deb https://mirrors.huaweicloud.com/debian/ stretch main contrib non-free
deb-src https://mirrors.huaweicloud.com/debian/ stretch main contrib non-free
deb https://mirrors.huaweicloud.com/debian/ stretch-updates main contrib non-free
deb-src https://mirrors.huaweicloud.com/debian/ stretch-updates main contrib non-free
deb https://mirrors.huaweicloud.com/debian/ stretch-backports main contrib non-free
deb-src https://mirrors.huaweicloud.com/debian/ stretch-backports main contrib non-free 

清华大学镜像站
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ stretch main contrib non-free
deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ stretch main contrib non-free
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ stretch-updates main contrib non-free
deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ stretch-updates main contrib non-free
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ stretch-backports main contrib non-free
deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ stretch-backports main contrib non-free
deb https://mirrors.tuna.tsinghua.edu.cn/debian-security/ stretch/updates main contrib non-free
deb-src https://mirrors.tuna.tsinghua.edu.cn/debian-security/ stretch/updates main contrib non-free

兰州大学镜像站
deb http://mirror.lzu.edu.cn/debian stable main contrib non-free
deb-src http://mirror.lzu.edu.cn/debian stable main contrib non-free
deb http://mirror.lzu.edu.cn/debian stable-updates main contrib non-free
deb-src http://mirror.lzu.edu.cn/debian stable-updates main contrib non-free
deb http://mirror.lzu.edu.cn/debian/ stretch-backports main contrib non-free
deb-src http://mirror.lzu.edu.cn/debian/ stretch-backports main contrib non-free
deb http://mirror.lzu.edu.cn/debian-security/ stretch/updates main contrib non-free
deb-src http://mirror.lzu.edu.cn/debian-security/ stretch/updates main contrib non-free

上海交大镜像站
deb https://mirror.sjtu.edu.cn/debian/ stretch main contrib non-free
deb-src https://mirror.sjtu.edu.cn/debian/ stretch main contrib non-free
deb https://mirror.sjtu.edu.cn/debian/ stretch-updates main contrib non-free
deb-src https://mirror.sjtu.edu.cn/debian/ stretch-updates main contrib non-free
deb https://mirror.sjtu.edu.cn/debian/ stretch-backports main contrib non-free
deb-src https://mirror.sjtu.edu.cn/debian/ stretch-backports main contrib non-free
deb https://mirror.sjtu.edu.cn/debian-security/ stretch/updates main contrib non-free
deb-src https://mirror.sjtu.edu.cn/debian-security/ stretch/updates main contrib non-free

最后附上官方全球镜像站列表地址https://www.debian.org/mirror/list
```
#### ubuntu
sed -i "s/archive.ubuntu.com/mirrors.163.com/g" /etc/apt/sources.list
sed -i "s/security.ubuntu.com/mirrors.163.com/g" /etc/apt/sources.list
