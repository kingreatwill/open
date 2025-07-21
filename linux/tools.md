<!-- toc -->
[TOC]
## linux命令或者工具
Linux最常用命令：简单易学，但能解决95%以上的问题
https://www.toutiao.com/a6763990899924926989/


### Asciinema
终端录制工具

输出文件格式[asciicast](https://docs.asciinema.org/manual/asciicast/v2/) format (.cast)

```
# exit退出录制
asciinema rec demo.cast


# 要一致回车进行回放
asciinema play demo.cast

# 分享, 会生成临时分享地址
asciinema upload demo.cast 

asciinema auth
```

### wget

> windows [下载地址](https://eternallybored.org/misc/wget/)

1. 使用 wget 命令下载文件`wget http://www.lxlinux.net` 默认情况下，下载的文件将保存到当前工作目录中的同名文件中。
如果你不想将下载的文件保存在本地，而只是想将其显示在标准输出（stdout）里，可以使用 `--output-document` 选项，同时后面再跟上一个 `-` 号。
`wget http://www.lxlinux.net --output-document - | head -n4`
如果你想给下载的文件重新命名，那么可以使用 `--output-document` 选项（或者更简便，直接使用 `-o` 即可）：`wget http://www.lxlinux.net --output-document newfile.html`

2. 断点续传:就可以使用 `--continue` 选项（或者 `-c` ）实现断点续传。
`wget --continue https://www.lxlinux.net/linux-distro.iso`

3. 下载一系列文件. 比如：file_1.txt，file_2.txt，file_3.txt，等等
`wget http://www.lxlinux.net/file_{1..4}.txt`

4. 镜像(下载)整个站点
如果你想对某个网站进行整站下载，包括其目录结构，那么你就需要使用 `--mirror` 选项。
这个选项等同于 `--recursive --level inf --timestamping --no-remove-listing`，这意味着它是无限递归的，因此你可以下载到指定域上的所有内容。
如果你使用 `wget` 存档站点，那么这些选项 `--no-cookies --page-requisites --convert-links` 也可用于确保每个页面都是最新的、最完整的。

5. 修改 HTML 请求标头
可以使用 `--debug` 选项查看 wget 每次请求发送的标头信息：`wget --debug www.lxlinux.net`
你可以使用 `--header` 选项修改请求标头。
`wget --debug --header="User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59" http://www.lxlinux.net`
除此之外，你还可以伪装成特定的移动设备（比如 iPhone ）：
```
$ wget --debug \
--header = "User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari /604.1" \
HTTP：// www.lxlinux.net
```

6. 查看响应标头
同样地，你可以使用 `--debug` 选项查看响应标头

7. 响应 301 响应
熟悉网络协议的都知道，200 响应码意味着一切都按预期进行。而 301 响应则意味着 URL 已经指向不同的网站。
这种情况下，如果你需要下载文件，那么就需要使用 wget 的重定向功能。所以，如果你遇到 301 响应时，就需要使用 `--max-redirect` 选项。
如果你不想要重定向，那么可以将 `--max-redirect` 设置为 0 。
`wget --max-redirect 0 http://www.lxlinux.net`

8. 展开短链接
这里依然使用 `--max-redirect` 选项：输出的倒数第二行，在 Location 部分，后面你就看到了短链展开之后的真面目。


### osquery
使用SQL查询你的 Linux 操作系统信息
[像查询数据库一样查询你的 Linux 操作系统信息](https://linux.cn/article-13578-1.html)
### bash-completion 自动补齐
yum install -y bash-completion
安装成功后，得到文件为 /usr/share/bash-completion/bash_completion，如果没有这个文件，则说明系统上没有安装这个工具。
source /usr/share/bash-completion/bash_completion

docker:
source /usr/share/bash-completion/completions/docker

kubectl:
source <(kubectl completion bash)

### ag (ack)
比grep、ack更快的递归搜索文件内容。
https://github.com/ggreer/the_silver_searcher

yum install the_silver_searcher
choco install ag
choco install ack

### tig
字符模式下交互查看git项目，可以替代git命令。
https://github.com/jonas/tig

### mycli
mysql客户端，支持语法高亮和命令补全，效果类似ipython，可以替代mysql命令。
https://github.com/dbcli/mycli

### jq
json文件处理以及格式化显示，支持高亮，可以替换python -m json.tool。

### shellcheck
shell脚本静态检查工具，能够识别语法错误以及不规范的写法。
https://github.com/koalaman/shellcheck

### yapf
Google开发的python代码格式规范化工具，支持pep8以及Google代码风格。
https://github.com/google/yapf

### mosh
基于UDP的终端连接，可以替代ssh，连接更稳定，即使IP变了，也能自动重连。
https://mosh.org/#getting

### fzf
命令行下模糊搜索工具，能够交互式智能搜索并选取文件或者内容，配合终端ctrl-r历史命令搜索简直完美。
https://github.com/junegunn/fzf

choco install fzf

### PathPicker(fpp)
在命令行输出中自动识别目录和文件，支持交互式，配合git非常有用。
https://github.com/facebook/PathPicker


运行以下命令：
git diff HEAD~8 --stat | fpp

### 进程实时监控-htop
提供更美观、更方便的进程监控工具，替代top命令。
https://hisham.hm/htop/

#### atop
https://www.atoptool.nl/index.php
#### vtop
https://github.com/MrRio/vtop
#### bashtop
https://github.com/aristocratos/bashtop
#### gtop
https://github.com/aksakalli/gtop


#### htop 和 glances
glances 是htop的补充工具。除了列出所有进程及其 CPU 和内存使用情况之外，它还可以显示有关系统的其他信息，比如：
- 网络及磁盘使用情况
- 文件系统已使用的空间和总空间
- 来自不同传感器（例如电池）的数据
- 以及最近消耗过多资源的进程列表

https://nicolargo.github.io/glances/

### 查看进程占用带宽情况-Nethogs
https://github.com/raboof/nethogs
```
下载：http://sourceforge.net/projects/nethogs/files/nethogs/0.8/nethogs-0.8.0.tar.gz/download
[root@localhost ~]#yum-y install libpcap-develncurses-devel 
[root@localhost ~]# tar zxvf nethogs-0.8.0.tar.gz 
[root@localhost ~]# cd nethogs 
[root@localhost nethogs]# make && make install 
[root@localhost nethogs]# nethogs eth0
```

### nmon
https://nmon.sourceforge.net/pmwiki.php?n=Main.HomePage&ref=itsfoss.com

### 硬盘读取性能测试-IOZone
```
下载：http://www.iozone.org/src/current/

[root]# tar xvf iozone3_420.tar 
[root]# cd iozone3_420/src/current/ 
[root]# make linux 
[root]# ./iozone -a -n 512m -g 16g -i 0 -i 1 -i 5 -f /mnt/iozone -Rb ./iozone.xls 
```
-a使用全自动模式
-n为自动模式设置最小文件大小(Kbytes)。
-g设置自动模式可使用的最大文件大小Kbytes。
-i用来指定运行哪个测试。
-f指定测试文件的名字完成后自动删除
-R产生Excel到标准输出
-b指定输出到指定文件上

### ncdu 磁盘分析
ncdu 下载地址：https://dev.yorhel.nl/ncdu
### 实时监控磁盘IO-IOTop
### 网络流量监控-IPtraf
### 网络流量监控-IFTop
```
下载：http://www.ex-parrot.com/~pdw/iftop/

[root@localhost ~]# tar zxvf iftop-0.17.tar.gz
[root@localhost ~]# cd iftop-0.17 
[root@localhost iftop-0.17]# ./configure 
[root@localhost iftop-0.17]# make && make install 
[root@localhost iftop-0.17]# iftop 
[root@localhost iftop-0.17]# iftop -i eth0  #指定监控网卡接口

TX：发送流量
RX：接收流量
TOTAL：总流量
Cumm：运行iftop到目前时间的总流量
peak：流量峰值
rates：分别表示过去 2s 10s 40s 的平均流量
```
### 系统资源监控-NMON
### 监控多个日志-MultiTail
MultiTail是在控制台打开多个窗口用来实现同时监控多个日志文档、类似tail命令的功能的软件。
### SSH暴力破解防护-Fail2ban
### 连接会话终端持续化-Tmux
### 页面显示磁盘空间使用情况-Agedu
### 安全扫描工具-NMap
### Web压力测试-Httperf
```
下载：http://code.google.com/p/httperf/downloads/list

[root]# tar zxvf httperf-0.9.0.tar.gz
[root]# cd httperf-0.9.0
[root]# ./configure
[root]# make && make install
[root]# httperf --hog --server=192.168.0.202 --uri=/index.html --num-conns= 10000 --wsess=10,10,0.1
参数说明：

--hog：让httperf尽可能多产生连接，httperf会根据硬件配置，有规律的产生访问连接
--num-conns：连接数量，总发起10000请求
--wsess：用户打开网页时间规律模拟，第一个10表示产生10个会话连接，第二个10表示每个会话连接进行10次请求，0.1表示每个会话连接请求之间的间隔时间/s
```
### axel
多线程下载工具，下载文件时可以替代curl、wget。

axel -n 20 http://centos.ustc.edu.cn/centos/7/isos/x86_64/CentOS-7-x86_64-Minimal-1511.iso

### httpstat
https://github.com/reorx/httpstat
httpstat以一种美丽和清晰的方式可视化curl统计数据。

### gping
Ping，但有个图表
https://github.com/orf/gping

### sz/rz
交互式文件传输，在多重跳板机下传输文件非常好用，不用一级一级传输。
yum install lrzsz -y


### cloc
代码统计工具，能够统计代码的空行数、注释行、编程语言。
https://github.com/AlDanial/cloc
http://cloc.sourceforge.net/

yum install cloc
choco install cloc                     # Windows with Chocolatey
scoop install cloc                     # Windows with Scoop


a file
cloc hello.c
or  a directory
cloc gcc-5.2.0/gcc/c
or  an archive
cloc master.zip
or a git repository, using a specific commit
cloc 6be804e07a5db
or each subdirectory of a particular directory
for d in ./*/ ; do (cd "$d" && echo "$d" && cloc --vcs git); done

### ccache
高速C/C++编译缓存工具，反复编译内核非常有用。使用起来也非常方便：
https://ccache.dev/
gcc foo.c
改成:
ccache gcc foo.c

### tmux
终端复用工具，替代screen、nohup。

### neovim
替代vim
https://neovim.io/

### script/scriptreplay
终端会话录制。
```
script -t 2>time.txt session.typescript # 录制开始
# your commands
exit # 录制结束
# 回放:
scriptreplay -t time.txt session.typescript
```
### you-get
非常强大的媒体下载工具，支持youtube、google+、优酷、芒果TV、腾讯视频、秒拍等视频下载。

### multitail
多重 tail

https://www.vanheusden.com/multitail/

### HTTP benchmarking tool压测工具
wrk : https://github.com/wg/wrk 38k c
ab : apachebench,Apache 网站服务器自带的
hey : https://github.com/rakyll/hey 18.2k  golang
vegeta: https://github.com/tsenart/vegeta 23.6k golang
ali: https://github.com/nakabonne/ali 2.8k golang (这款工具基于 vegeta，但使用起来更傻瓜式，而且实时进行分析，图形化展示。)
plow: https://github.com/six-ddc/plow 2.3k golang 内部采用了性能出色的fasthttp而不是 Go 的默认net/http, 支持实时WebUI数据展示

### Starship
https://starship.rs/
它可以做到：
- 根据你是否在代码仓库中添加了新文件、是否修改了文件、是否暂存了文件等情况，用相应的符号表示 git 仓库的状态。
- 根据你所在的 Python 项目目录，展示 Python 的版本号，这也适用于 Go/Node/Rust/Elm 等其他编程语言环境。
- 展示上一个命令执行所用的时间，指令运行时间必须在毫秒级别。
- 如果上一个命令执行失败，会展示相应的错误提示符。

### z
z 可以让你快速地在文件目录之间跳转
https://github.com/rupa/z

### fzf
fzf — fuzzy finder，即模糊查找器。它是一种通用工具，可让你使用模糊搜索来查找文件、历史命令、进程、git 提交等。你键入一些字母，它会尝试匹配结果列表中任何位置的字母。输入的字母越多，结果也就越准确。
https://github.com/junegunn/fzf

### fd
类似于系统自带的 find 命令，但使用起来更简单，查找速度更快，并且具有良好的默认设置。
https://github.com/sharkdp/fd

### ripgrep

ripgrep是grep命令的替代方法， 不过ripgrep的执行速度更快，而且具有健全的默认配置以及丰富的彩色输出。
https://github.com/BurntSushi/ripgrep

### virtualenv 和 virtualfish
Virtualenv 是用于在 Python 中创建虚拟环境的工具。
VirtualFish 则是 Fish Shell 的虚拟环境管理器。它提供了许多命令来执行快速创建、列出或删除虚拟环境等操作。

virtualenv 下载地址：https://pypi.org/project/virtualenv/
virtualfish 下载地址：https://github.com/justinmayer/virtualfish

### pyenv、nodenv 和 rbenv
pyenv 可以轻松实现 Python 版本的切换。
Pyenv、nodenv 和 rubyenv 是用于管理计算机上不同版本的 Python、Node 和 Ruby 的工具。
pyenv 下载地址：https://github.com/pyenv/pyenv
nodenv 下载地址：https://github.com/nodenv/nodenv
rbenv 下载地址：https://github.com/rbenv/rbenv

### bat cat的替代品
bat 下载地址：https://github.com/sharkdp/bat

### httpie curl替代工具
https://httpie.org/

### tldr  man的简化版
https://tldr.sh/
### exa ls命令的一个可替代方案
https://the.exa.website/

### litecli 和 pgcli
这是SQLite 和 PostgreSQL CLI 的解决方案。借助自动提示和语法突出显示，它们比默认的sqlite3和psql工具要好用很多。
litecli 下载地址：https://litecli.com/
pgcli 下载地址：https://www.pgcli.com/

## troubleshooting/linux 调试工具
### sysdig
https://github.com/draios/sysdig
可以看网络,CPU,磁盘I/O,进程等
### Falco
Falco 容器安全系统
https://github.com/falcosecurity

## netstat -lntup  
说明： l:listening   n:num   t:tcp  u:udp  p:process

## securecrt ssh工具

## BusyBox Unix常用工具包
BusyBox 是一个集成了一百多个最常用 linux 命令和工具的软件。BusyBox 包含了一些简单的工具，例如 ls、cat 和echo 等等，
还包含了一些更大、更复杂的工具，例 grep、find、mount 以及 telnet。有些人将 BusyBox 称为 Linux 工具里的瑞士军刀。
简单的说 BusyBox 就好像是个大工具箱，它集成压缩了 Linux 的许多工具和命令，也包含了 Android 系统的自带的shell。

比如：busybox ls

COMMANDS
```
 [, [[, acpid, addgroup, adduser, adjtimex, ar, arp, arping, ash,
        awk, basename, beep, blkid, brctl, bunzip2, bzcat, bzip2, cal, cat,
        catv, chat, chattr, chgrp, chmod, chown, chpasswd, chpst, chroot,
        chrt, chvt, cksum, clear, cmp, comm, cp, cpio, crond, crontab,
        cryptpw, cut, date, dc, dd, deallocvt, delgroup, deluser, depmod,
        devmem, df, dhcprelay, diff, dirname, dmesg, dnsd, dnsdomainname,
        dos2unix, dpkg, du, dumpkmap, dumpleases, echo, ed, egrep, eject,
        env, envdir, envuidgid, expand, expr, fakeidentd, false, fbset,
        fbsplash, fdflush, fdformat, fdisk, fgrep, find, findfs, flash_lock,
        flash_unlock, fold, free, freeramdisk, fsck, fsck.minix, fsync,
        ftpd, ftpget, ftpput, fuser, getopt, getty, grep, gunzip, gzip, hd,
        hdparm, head, hexdump, hostid, hostname, httpd, hush, hwclock, id,
        ifconfig, ifdown, ifenslave, ifplugd, ifup, inetd, init, inotifyd,
        insmod, install, ionice, ip, ipaddr, ipcalc, ipcrm, ipcs, iplink,
        iproute, iprule, iptunnel, kbd_mode, kill, killall, killall5, klogd,
        last, length, less, linux32, linux64, linuxrc, ln, loadfont,
        loadkmap, logger, login, logname, logread, losetup, lpd, lpq, lpr,
        ls, lsattr, lsmod, lzmacat, lzop, lzopcat, makemime, man, md5sum,
        mdev, mesg, microcom, mkdir, mkdosfs, mkfifo, mkfs.minix, mkfs.vfat,
        mknod, mkpasswd, mkswap, mktemp, modprobe, more, mount, mountpoint,
        mt, mv, nameif, nc, netstat, nice, nmeter, nohup, nslookup, od,
        openvt, passwd, patch, pgrep, pidof, ping, ping6, pipe_progress,
        pivot_root, pkill, popmaildir, printenv, printf, ps, pscan, pwd,
        raidautorun, rdate, rdev, readlink, readprofile, realpath,
        reformime, renice, reset, resize, rm, rmdir, rmmod, route, rpm,
        rpm2cpio, rtcwake, run-parts, runlevel, runsv, runsvdir, rx, script,
        scriptreplay, sed, sendmail, seq, setarch, setconsole, setfont,
        setkeycodes, setlogcons, setsid, setuidgid, sh, sha1sum, sha256sum,
        sha512sum, showkey, slattach, sleep, softlimit, sort, split,
        start-stop-daemon, stat, strings, stty, su, sulogin, sum, sv,
        svlogd, swapoff, swapon, switch_root, sync, sysctl, syslogd, tac,
        tail, tar, taskset, tcpsvd, tee, telnet, telnetd, test, tftp, tftpd,
        time, timeout, top, touch, tr, traceroute, true, tty, ttysize,
        udhcpc, udhcpd, udpsvd, umount, uname, uncompress, unexpand, uniq,
        unix2dos, unlzma, unlzop, unzip, uptime, usleep, uudecode, uuencode,
        vconfig, vi, vlock, volname, watch, watchdog, wc, wget, which, who,
        whoami, xargs, yes, zcat, zcip
```

docker
https://hub.docker.com/_/busybox
github
https://github.com/mirror/busybox
官网
https://www.busybox.net/

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

### NeoHtop
https://github.com/Abdenasser/neohtop/releases

## UI
### cockpit   
web管理页面 连接linux
### webmin
web服务器管理控制面板
https://github.com/webmin/webmin
### xDroid
运行安卓应用

xDev提供与xDroid、xWin等适配接口,基于xDev开发的安卓应用、Window应用无需借助模拟器或虚拟机即可直接运行在异构国产平台上
### Linux 用户的最佳终端 Web 浏览器
#### W3M
#### Lynx
#### Links2
#### eLinks
#### 终端应用程序Xterm.js
https://github.com/xtermjs/xterm.js

## 系统监控/linux 监控

### hegemon
https://github.com/p-e-w/hegemon

cargo install hegemon

### eul
macOS status monitoring app written in SwiftUI.

https://github.com/gao-sun/eul

### glances、dstat、pidstat
https://github.com/nicolargo/glances
Glances 是一个由 Python 编写，使用 psutil 库来从系统抓取信息的基于 curses 开发的跨平台命令行系统监视工具。

dstat命令是一个用来替换vmstat、iostat、netstat、nfsstat和ifstat这些命令的工具，是一个全能系统信息统计工具。
## 系统工具
### crash
https://github.com/crash-utility/crash

crash是redhat的工程师开发的，主要用来离线分析linux内核转存文件，它整合了gdb工具，功能非常强大。可以查看堆栈，dmesg日志，内核数据结构，反汇编等等。
crash支持多种工具生成的转存文件格式，如kdump，LKCD，netdump和diskdump，而且还可以分析虚拟机Xen和Kvm上生成的内核转存文件。
同时crash还可以调试运行时系统，直接运行crash即可，ubuntu下内核映象存放在/proc/kcore。
crash和linux内核是紧密耦合的，会随着内核的变化持续更新，它向前兼容的，新的crash工具可以分析老内核的转存文件


## xx
https://github.com/xwmx/nb

## 应用
### winapps 
运行windows APP
https://github.com/Fmstrat/winapps

### 有趣的工具
#### asciiview 
```
sudo apt-get install aview imagemagick
asciiview kobe.jpg # 将图片转化成字符输出
```
[N个Linux耍酷命令](https://www.toutiao.com/a6820664143654158861/)

#### neofetch
在终端中显示Linux系统信息
https://github.com/dylanaraps/neofetch
#### onefetch
在终端中显示Git存储库摘要
https://github.com/o2sh/onefetch
```
wget https://github.com/o2sh/onefetch/releases/download/v2.5.0/onefetch-linux.tar.gz
sudo mv onefetch-linux/onefetch /usr/local/bin/

# 使用Onefetch在终端中显示Git存储库摘要
onefetch openjw/open/
# or cd openjw/open/  然后 onefetch
```
由于Onefetch的主要部分是在 Rust，它显示了Rust语言的ASCII艺术。 您也可以使用其他语言的ASCII `-a`， `--ascii-language` 。 以下命令显示Onefetch Git存储库摘要以及 蟒蛇 ASCII徽标：`onefetch -a python`
您还可以使用自定义图片，例如使用 `-i`， `--image` 。如：`onefetch -i ~/ostechnix.png`
默认情况下，Onefetch仅显示三个作者。 一些作者可能会为Git项目做出贡献。 要显示要显示的作者数量，请使用 `-A`， `--authors-number`。如：`onefetch -A 10`