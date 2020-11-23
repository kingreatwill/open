<!--toc-->
[TOC]

# shell脚本

## 常用shell脚本
https://github.com/dylanaraps/pure-bash-bible
[101个shell脚本](https://blog.51cto.com/zero01/2046242)

### $0、$1、$2、$#、$@、$*、$? 的含义
假设执行 ./test.sh a b c 这样一个命令，则可以使用下面的参数来获取一些值：
$0 对应 "./test.sh" 这个值。如果执行的是 ./work/test.sh， 则对应 ./work/test.sh 这个值，而不是只返回文件名本身的部分。
$1 会获取到 a，即 $1 对应传给脚本的第一个参数。
$2 会获取到 b，即 $2 对应传给脚本的第二个参数。
$3 会获取到 c，即 $3 对应传给脚本的第三个参数。$4、$5 等参数的含义依此类推。
$# 会获取到 3，对应传入脚本的参数个数，统计的参数不包括 $0。
$@ 会获取到 "a" "b" "c"，也就是所有参数的列表，不包括 $0。
$* 也会获取到 "a" "b" "c"， 其值和 $@ 相同。但 "$*" 和 "$@" 有所不同。"$*" 把所有参数合并成一个字符串，而 "$@" 会得到一个字符串参数数组。
$? 可以获取到执行 ./test.sh a b c 命令后的返回值。在执行一个前台命令后，可以立即用 $? 获取到该命令的返回值。该命令可以是系统自身的命令，可以是 shell 脚本，也可以是自定义的 bash 函数。


当执行系统自身的命令时，$? 对应这个命令的返回值。
当执行 shell 脚本时，$? 对应该脚本调用 exit 命令返回的值。如果没有主动调用 exit 命令，默认返回为 0。
当执行自定义的 bash 函数时，$? 对应该函数调用 return 命令返回的值。如果没有主动调用 return 命令，默认返回为 0。

下面举例说明 "$*" 和 "$@" 的差异。假设有一个 testparams.sh 脚本，内容如下：
```
#!/bin/bash

for arg in "$*"; do
    echo "****:" $arg
done
echo --------------
for arg in "$@"; do
    echo "@@@@:" $arg
done
```
这个脚本分别遍历 "$*" 和 "$@" 扩展后的内容，并打印出来。执行 ./testparams.sh，结果如下：
```
$ ./testparams.sh This is a test
****: This is a test
--------------
@@@@: This
@@@@: is
@@@@: a
@@@@: test
```

可以看到，"$*" 只产生一个字符串，for 循环只遍历一次。
而 "$@" 产生了多个字符串，for 循环遍历多次，是一个字符串参数数组。

注意：如果传入的参数多于 9 个，则不能使用 $10 来引用第 10 个参数，而是要用 ${10} 来引用。即，需要用大括号{}把大于 9 的数字括起来。

例如，${10} 表示获取第 10 个参数的值，写为 $10 获取不到第 10 个参数的值。实际上，$10 相当于 ${1}0，也就是先获取 $1 的值，后面再跟上 0，如果 $1 的值是 "first"，则 $10 的值是 "first0"。

## linux实用技巧
### 用户登录信息
- 使用 last 命令获取用户登录信息
`last | head -5 | tr -s " "`
`tr -s " "` 表示将多个空格合并为一个，这样可以节约篇幅

- 统计每个用户登录次数
```
for user in `ls /home`; do echo -ne "$user\t"; last $user | wc -l; done
```
show_user_logins.sh
```
#!/bin/bash

echo -n "Logins since "
who /var/log/wtmp | head -1 | awk '{print $3}'
echo "======================="

for user in `ls /home`
do
  echo -ne "$user\t"
  last $user | wc -l
done
```

- 统计每个用户登录时长
单个用户
`ac username`
所有用户
```
for user in `ls /home`; do ac $user | sed "s/total/$user\t/" ; done
```
使用sed去掉每行前面的空格
```
for user in `ls /home`; do ac $user | sed "s/^\t//" | sed "s/total/$user\t/" ; done
```
show_user_hours.sh
```
#!/bin/bash

echo -n "hours online since "
who /var/log/wtmp | head -1 | awk '{print $3}'
echo "============================="

for user in `ls /home`
do
  ac $user | sed "s/^\t//" | sed "s/total/$user\t/"
done
```

## cmd
### nohup

1. sh test.sh &  
将sh test.sh任务放到后台 ，即使关闭xshell退出当前session依然继续运行，但**标准输出和标准错误信息会丢失（缺少的日志的输出）**
将sh test.sh任务放到后台 ，关闭xshell，对应的任务也跟着停止。

2. nohup sh test.sh  
将sh test.sh任务放到后台，关闭标准输入，**终端不再能够接收任何输入（标准输入）**，重定向标准输出和标准错误到当前目录下的nohup.out文件，即使关闭xshell退出当前session依然继续运行。

3. nohup sh test.sh  & 
将sh test.sh任务放到后台，但是依然可以使用标准输入，**终端能够接收任何输入**，重定向标准输出和标准错误到当前目录下的nohup.out文件，即使关闭xshell退出当前session依然继续运行。

[命令行的艺术](https://github.com/jlevy/the-art-of-command-line/blob/master/README-zh.md)



### find 
- 实例1：使用find命令查找相关文件后，再使用ls命令将它们的详细信息列出来

我们现在想把当前目录下所有的.o文件全部找出来，并用 ls -l 命令将它们列出来。实现这个需求的命令如下：
```
find . -name "*.o" -type f -exec ls -l {} \;
```

- 实例2：使用find命令查找相关文件后，再使用rm命令将它们删除

我们现在想把当前目录下所有的.o文件全部找出来，并用rm命令将它们删除。实现这个需求的命令如下：
```
find . -name "*.o" -exec rm {} \;
```

- 实例3：使用-exec选项的安全模式，将对每个匹配到的文件进行操作之前提示用户

在实例2中，我们匹配到文件后就立刻执行rm命令，这样操作有些危险，因为如果一旦误操作，有可能会引起灾难性的后果。

exec的安全模式就是为了避免这个问题而产生。它会在匹配到某个文件后，在进行操作之前会先问一下你，经过你的确认它才会进行相应操作。

同样的实例2的需求，如果采用安全模式的话，命令是这样的：
```
find . -name "*.o" -ok rm {} \;
```

- 实例4：搜索匹配到的文件中的关键内容

假如我现在有个很大型的项目（如Linux内核），我想在里面搜索一个含有某关键字的文件。我们可以使用grep命令检索所有的文件。这样做肯定是可以的，但如果项目很大的话，这样太耗时了，效率太低。

我们可以先用find命令找到所以相关文件，然后再用grep命令检索那些文件即可。因为已经使用find过滤一遍了，所以这样操作会节约很多时间，提高效率。

命令如下：
```
find . -name "*.h" -exec grep -rns "hello" {} \;
```
其实就是利用grep搜索字符串
```
grep -Er 'hello' .
```

- 实例5：查找文件并移动到指定目录

这个需求就比较简单了。比如我现在想把所有的.o文件找出来，然后新他们mv到buil目录。命令如

下：
```
find . -name "*.o" -exec cp {} build \;
```

### cp
```
-r 递归复制
-u 更新
```
cp -ru xx/* c

需要c目录存在

### which

#### cmd
where 
or
for %x in (powershell.exe) do @echo %~$PATH:x
#### powershell
Get-Command where # Get-Command and its alias gcm
gcm where

### mem
free按1024进制计算
free -g # 以G显示
free -m # 以M显示

cat /proc/meminfo|grep Slab
or
echo `cat /proc/meminfo|grep Slab|awk '{mem += $2} END {print mem/1024/1024}'` GB

echo `ps aux |awk '{mem += $6} END {print mem/1024/1024}'` GB

#### pstree
pstree可以将所有进程以树状展示出来
加上-p参数可以看到系统上的每个进程，`pstree -p`
加上pid`pstree -p <pid>`可以看到进程的线程

#### vmstat 
可以对操作系统的虚拟内存、进程、CPU活动进行监控。

一般有两个参数，用来表示采样速率与次数。
`vmstat <时间间隔> <采集次数>`

- Procs
    r：运行和等待CPU时间片的进程数
    b：表示阻塞的进程数。

- Memory
    swpd：表示虚拟内存使用情况
    free：表示当前空闲的物理内存
    buff：表示缓冲的内存大小
    Cache：表示缓存的内存大小，

- Swap
    si：表示有磁盘读入内存大小。
    so：表示由内存写入磁盘大小。

- Io
    bi：表示由块设备读入数据的总量
    bo：表示写到块设备数据的总量

- System
    in：表示每秒中断数。
    cs：表示每秒产生的上下文切换次数(线程切换)。

- cpu
    us：表示用户进程消耗的CPU时间百分比
    sy：表示系统调用消耗的CPU时间百分比
    id：表示CPU处在空间状态的时间百分比
    wa: 等待IO的CPU时间

#### pidstat    
可以用来监控全部或指定进程的cpu、内存、线程、设备IO等系统资源的占用情况
`pidstat -w <时间间隔>`
- Cswch/s:每秒主动任务上下文切换数量
- Nvcswch/s:每秒被动任务上下文切换数量
- Command:命令名

`pidstat -tt -p <pid>`可以看到每个进程的线程切换与cpu调度情况
- TGID:主线程的表示
- TID:线程id
- %usr：进程在用户空间占用cpu的百分比
- %system：进程在内核空间占用cpu的百分比
- %guest：进程在虚拟机占用cpu的百分比
- %CPU：进程占用cpu的百分比
- CPU：处理进程的cpu编号
- Command：当前进程对应的命令

#### straces
通过该命令可以看到该进程正在进行的用户空间与内核空间的交互，如系统调用，进程状态等。
`strace -t -p`

### du 查看目录大小
```
du -sh .
du -ah .
也可以像下面这样，显示当前目录下的所有一级子目录的大小：
du -h --max-depth=0 .
```
du [Path]
-a 全部文件 包括隐bai藏的。
-h 以M 为单位显示du文件大小结果。
-s 统计此zhi目录中所有文件大小总和。


du的替代品
https://github.com/muesli/duf

PowerShell
```
PowerShell 命令：
Get-ChildItem -Recurse | Measure-Object -Sum Length
Get-ChildItem 命令用于遍历目录下的所有子目录和文件，类似于 dir 命令，使用 -Recurse 参数可以实现递归遍历。
Measure-Object 命令常作用于管道，对管道的结果进行统计操作，譬如：计数、求和、平均数、最大数、最小数等等。

PowerShell 的命令总给人一种怪怪的感觉，不过它也提供了简写的语法：
ls -r | measure -s Length

看起来比上面的要舒服多了。或者直接在命令行 cmd 下执行：
powershell -noprofile -command "ls -r | measure -s Length"
```

### 网路相关
#### netstat
#### ss
SS命令可以提供如下信息：
- 所有的TCP sockets
- 所有的UDP sockets
- 所有ssh/ftp/ttp/https持久连接
- 所有连接到Xserver的本地进程
- 使用state（例如：connected, synchronized, SYN-RECV, SYN-SENT,TIME-WAIT）、地址、端口过滤
- 所有的state FIN-WAIT-1 tcpsocket连接以及更多

常用ss命令：
- ss -l 显示本地打开的所有端口
- ss -pl 显示每个进程具体打开的socket
- ss -t -a 显示所有tcp socket
- ss -u -a 显示所有的UDP Socekt
- ss -o state established '( dport = :smtp or sport = :smtp )' 显示所有已建立的SMTP连接
- ss -o state established '( dport = :http or sport = :http )' 显示所有已建立的HTTP连接
- ss -x src /tmp/.X11-unix/* 找出所有连接X服务器的进程
- ss -s 列出当前socket详细信息:
- ss src ADDRESS_PATTERN(src：表示来源,ADDRESS_PATTERN：表示地址规则)
    ss src 120.33.31.1 # 列出来之20.33.31.1的连接    
    ss src 120.33.31.1:http # 列出来至120.33.31.1,80端口的连接
    ss src 120.33.31.1:80