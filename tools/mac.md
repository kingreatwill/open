
##
### OrbStack
Fast, light, simple Docker containers & Linux machines for macOS
https://orbstack.dev/
https://github.com/orbstack/orbstack


### 网址
[macwk](https://macwk.com.cn/)


### 命令行
#### ip
在linux系统上ip命令很常用，但是在Mac系统上并没有这个命令，我们可以安装 iproute2mac 使用类似ip功能的命令。
https://github.com/brona/iproute2mac

### 工具

snipaste
pastenow
hapigo/utools /Alfred 5/ Raycast就像是升级版的Spotlight搜索

#### 文件管理 finder/访达
- QSpace Pro

#### dock
- ubar
- ActiveDock 2
- HyperDock
- cDock
- Dragthins
#### 解压工具
- keka
- maczip
#### 彻底卸载应用
https://github.com/alienator88/Pearcleaner
#### 录屏
https://github.com/lihaoyun6/QuickRecorder

#### 终端
warp https://www.warp.dev/
Tabby
electerm
WindTerm
### 虚拟机
#### Multipass
下载
https://multipass.run/
或者
```
brew cask install multipass
multipass --version
multipass find
multipass launch -n vm01 -c 1 -m 1G -d 10G


-n, --name: 名称
-c, --cpus: cpu核心数, 默认: 1
-m, --mem: 内存大小, 默认: 1G
-d, --disk: 硬盘大小, 默认: 5G


multipass list
multipass exec vm01 pwd
multipass info vm01
multipass shell vm01 # 进入虚拟机
# 挂载
multipass mount /Users/moxi/hello  vm01:/hello
#卸载数据卷
multipass umount vm01
#将 hello.txt 发送到
multipass transfer hello.txt vm01:/home/ubuntu/

# 启动实例
multipass start vm01
# 停止实例
multipass stop vm01
# 删除实例（删除后，还会存在）
multipass delete vm01
# 释放实例（彻底删除）
multipass purge vm01
```
#### lima
纯命令行Linux虚拟工具 
开源仓库地址：https://github.com/lima-vm/lima

根据已有的模版创建虚拟机,可以在本地的/usr/local/share/lima/templates目录查看，或者去lima的这个网页查看：https://lima-vm.io/docs/templates/

```
brew install lima

# 创建
limactl create --name=ubuntu-lts template://ubuntu-lts

# 启动
limactl start ubuntu-lts

# 查看现在虚拟机的列表及状态。
limactl list

# 进入linux终端
limactl shell ubuntu-lts

# 停止虚拟机
limactl stop ubuntu-lts
# 删除虚拟机
limactl delete ubuntu-lts
# 保护一个虚拟机，防止误删除
limactl protect ubuntu-lts
# 取消对虚拟机的保护，允许删除
limactl unprotect ubuntu-lts
```


#### CrossOver
类虚拟机-在 macOS 和 Linux 上运行 Windows 应用程序，使用最成熟的游戏转译层Wine

#### 超级右键
- iRightMouse(功能丰富)
- New File Menu
- Easy New File
- MouseBoost Pro(功能丰富)
- RClick(免费)