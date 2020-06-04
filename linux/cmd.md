# nohup

1. sh test.sh &  
将sh test.sh任务放到后台 ，即使关闭xshell退出当前session依然继续运行，但**标准输出和标准错误信息会丢失（缺少的日志的输出）**
将sh test.sh任务放到后台 ，关闭xshell，对应的任务也跟着停止。

2. nohup sh test.sh  
将sh test.sh任务放到后台，关闭标准输入，**终端不再能够接收任何输入（标准输入）**，重定向标准输出和标准错误到当前目录下的nohup.out文件，即使关闭xshell退出当前session依然继续运行。

3. nohup sh test.sh  & 
将sh test.sh任务放到后台，但是依然可以使用标准输入，**终端能够接收任何输入**，重定向标准输出和标准错误到当前目录下的nohup.out文件，即使关闭xshell退出当前session依然继续运行。

[命令行的艺术](https://github.com/jlevy/the-art-of-command-line/blob/master/README-zh.md)