
##
### OrbStack
Fast, light, simple Docker containers & Linux machines for macOS
https://orbstack.dev/
https://github.com/orbstack/orbstack


### 网址
[macwk](https://macwk.com.cn/)

### 工具

snipaste
pastenow
hapigo/utools /Alfred 5/ Raycast就像是升级版的Spotlight搜索


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
