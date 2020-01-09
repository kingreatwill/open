<!-- toc -->
[TOC]
## linux命令或者工具
Linux最常用命令：简单易学，但能解决95%以上的问题
https://www.toutiao.com/a6763990899924926989/

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

### htop
提供更美观、更方便的进程监控工具，替代top命令。
https://hisham.hm/htop/

### axel
多线程下载工具，下载文件时可以替代curl、wget。

axel -n 20 http://centos.ustc.edu.cn/centos/7/isos/x86_64/CentOS-7-x86_64-Minimal-1511.iso

### sz/rz
交互式文件传输，在多重跳板机下传输文件非常好用，不用一级一级传输。

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
wrk : https://github.com/wg/wrk
ab : 
hey : https://github.com/rakyll/hey