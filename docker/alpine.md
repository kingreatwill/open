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
