## 常用操作

WSL2的卸载操作如下：
wslconfig /l
从列表中选择要卸载的发行版（例如Ubuntu）并键入命令
wslconfig /u Ubuntu
设置默认
wslconfig /s Ubuntu-18.04 


以下命令不行
`lxrun /uninstall /full`

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

## 文件目录在win10文件系统中的位置

C:\Users\xxx\AppData\Local\Packages\CanonicalGroupLimited.Ubuntu18.04onWindows_79rhkp1fndgsc\LocalState\rootfs\

## 将默认镜像源修改为国内的
```
mv /etc/apt/sources.list /etc/apt/sources.listback
vi /etc/apt/sources.list

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