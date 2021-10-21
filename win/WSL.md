[TOC]
# Windows Subsystem for Linux
## 常用操作
[命令参考](https://docs.microsoft.com/zh-cn/windows/wsl/wsl-config)
[命令参考](https://docs.microsoft.com/zh-cn/windows/wsl/reference)
> 在 Windows 10 1903 之前的版本中，使用wslconfig命令；wslconfig /? 

- 列出分发版
> 老版本命令：wslconfig /l

`wsl -l , wsl --list`
列出可用于 WSL 的 Linux 分发版。 如果列出了某个分发版，表示该分发版已安装且可供使用。

`wsl --list --all` 列出所有分发，包括当前不可用的分发。 这些分发版可能正在安装、卸载或处于损坏状态。

`wsl --list --running` 列出当前正在运行的所有分发。

`wsl -l -v` 显示wsl版本

- 卸载
> wslconfig /u Ubuntu

`wsl --unregister DistributionName`

- 设置默认
> wslconfig /s Ubuntu-18.04 ， wslconfig /setdefault Ubuntu-20.04

`wsl -s DistributionName, wsl --setdefault DistributionName`

- 进入WSL 实例
`wsl --terminate CentOS`
`wsl -d CentOS`

- 以特定用户的身份运行
`wsl -u Username, wsl --user Username`

- 更改分发的默认用户
`DistributionName config --default-user Username`

- 运行特定的分发版

`wsl -d DistributionName, wsl --distribution DistributionName`

- 文件相互访问
使用`/mnt/c` 从 WSL 访问 Windows 文件
使用 `\\wsl$` 从 Windows 访问 Linux 文件

## WSL的一些用法
### vs code
https://docs.microsoft.com/zh-cn/windows/wsl/tutorials/wsl-vscode
在WSL中执行命令`code .` 相当于在win中打开vs code并链接到wsl中

### docker
https://docs.microsoft.com/zh-cn/windows/wsl/tutorials/wsl-containers

Docker Desktop可以直接使用wsl中的docker环境（设置 > 资源 > WSL 集成 | Settings -> Resources -> WSL INTEGRATION）

## 图形界面的安装
### 第一种：X-WINDOWS的安装使用图形化界面
#### 安装 X-Windows
可供选择安装的 X-Windows 有多个：VcXsrv Windows X Server、Xming、Cygwin X Server，教程中选择的是 VcXsrv，因为比较容易和稳定。

下载地址：https://sourceforge.net/projects/vcxsrv/

下载之后，双击安装，没什么好说，一路默认安装即可，在安装完毕后，开始菜单中出现「XLaunch」图标，双击运行，选择：**“one large window”，Display number 设置成 0**，其它一路默认

#### WSL 安装桌面环境
sudo apt-get install ubuntu-desktop unity compizconfig-settings-manager
注意：中间有可能出现意想不到的状况，使用下面的命令重新安装即可
sudo apt-get install ubuntu-desktop unity compizconfig-settings-manager --fix-missing

#### 配置 compiz 并运行桌面环境

启动 XLaunch，选择：“one large window”，Display number 设置成 0，其它一路默认，这里在上面设置过了，如果你没关就不用重现选择了，如果不小心关掉了，那么就重新来一次吧。我们用管理员身份运行ubuntu Bash并且中执行：
```
export DISPLAY=localhost:0
ccsm
```
在 X-windows 中，即会弹出 ccsm 的配置界面，勾选需要的 Desktop 组件（只需要勾选 Desktop 中的 Ubuntu Unity Plugin 即可，其它默认）,

点击 close 关闭 ccsm，执行命令：`compiz`
短暂的加载后，在 X-windows 的界面中就能看到桌面版的 Ubuntu 了。


### 第二种：通过本机远程桌面连接
#### xorg
包括显卡驱动，图形环境库等等一系列软件包
`sudo apt-get install xorg`
#### XFCE4
运行在类的的Unix的的的的操作系统上，提供轻量级桌面环境
`sudo apt-get install xfce4`
#### XRDP
一种开源的远程桌面协议（RDP）服务器
`sudo apt-get install xrdp`
#### 配置XRDP（配置端口）
`sudo sed -i 's/port=3389/port=3390/g' /etc/xrdp/xrdp.ini`

#### 向的.xsession的文件中的文件中写入的XFCE4会话
`sudo echo xfce4-session >~/.xsession`
#### 重启XRDP服务
`sudo service xrdp restart`

#### 链接
在CMD中搜索MSTSC，或者点击远程桌面连接，点击进入，计算机栏输入【本机IP：端口】，用户名栏输入子系统用户名，点击连接。

### 第三种：官方：WSLG系统 （Windows Subsystem for Linux GUI）
win11

https://github.com/microsoft/wslg

https://docs.microsoft.com/zh-cn/windows/wsl/tutorials/gui-apps

## 文件目录在win10文件系统中的位置

C:\Users\xxx\AppData\Local\Packages\CanonicalGroupLimited.Ubuntu18.04onWindows_79rhkp1fndgsc\LocalState\rootfs\

## 将默认镜像源修改为国内的
```
sudo mv /etc/apt/sources.list /etc/apt/sources.listback
sudo vi /etc/apt/sources.list

deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial main restricted universe multiverse

# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial main restricted universe multiverse

deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial-updates main restricted universe multiverse

# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial-updates main restricted universe multiverse

deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial-backports main restricted universe multiverse

# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial-backports main restricted universe multiverse

deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial-security main restricted universe multiverse

# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial-security main restricted universe multiverse

# 预发布软件源，不建议启用

# deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial-proposed main restricted universe multiverse

# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ xenial-proposed main restricted universe multiverse

```
以上是清华的源
其他版本的源列表参考：https：//mirrors.tuna.tsinghua.edu.cn/help/ubuntu/

阿里的源
```
deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
```



修改完毕后，使用下面的命令使其生效：

sudo apt-get update

更新已安装的软件：
sudo apt-get upgrade













# Windows Subsystem for Android™
Win11安卓子系统
https://docs.microsoft.com/zh-cn/windows/android/wsa/
