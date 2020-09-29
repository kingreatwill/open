<!--toc-->
[TOC]

# shell脚本

## 常用shell脚本
https://github.com/dylanaraps/pure-bash-bible
[101个shell脚本](https://blog.51cto.com/zero01/2046242)

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
