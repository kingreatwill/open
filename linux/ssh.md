
## ssh连接
一、正常连接方法：ssh root@10.0.0.20

二、无密码连接方法(有两台机器：此处我把被连接的称为服务器，另一台则称为客户端):

　　1、先在服务器添加目录 .ssh： mkdir  .ssh

　　2、分配.ssh目录权限： chmod 777 .ssh

　　3、在客户端创建公钥与私钥： ssh-keygen　　//此处直接按多个回车键，直到创建成功

　　4、将客户端的公钥复制到要服务器，运行命令：ssh-copy-id root@10.0.0.20 ，待输入正确密码后即可实现ssh无密码访问。

三、Win没有ssh-copy-id
在服务器端
```
mkdir ~/.ssh
chmod 0700 ~/.ssh
touch ~/.ssh/authorized_keys
chmod 0644 ~/.ssh/authorized_keys
nano ~/.ssh/authorized_keys     # Ctrl+O 保存   Ctrl+X 退出
```
在win上找到ssh-keygen生成的.pub文件 放到~/.ssh/authorized_keys文件里

## shell工具
### xshell
收费

[家庭/学校免费](https://www.netsarang.com/zh/free-for-home-school/)

[公测版 7 免费 - 公测版本可以一直试用到正式版上线。](https://www.netsarang.com/zh/version-7-open-beta/)
#### PortX
NetSarang旗下已经有Xshell/Xftp/Xmanager等产品，用过的同学可能知道，他家的软件之前仅支持Windows平台，可能是为了进一步扩大市场，NetSarang公司最近推出了一款跨平台的SSH客户端PortX，支持Mac、Windows、Linux三大平台，目前可以免费下载使用
https://portx.online/zh/

### Tabby Terminal
https://github.com/Eugeny/tabby 52K
原[Terminus](https://github.com/Eugeny/terminus)

### Wave Terminal
https://github.com/wavetermdev/waveterm

### WindTerm
https://github.com/kingToolbox/WindTerm

### MobaXterm
全能终端神器
### electerm
Terminal/ssh/sftp client(linux, mac, win)
https://github.com/electerm/electerm


### SecureCRT

### finalshell
https://fishshell.com/
### Putty、telnet
### Tabby
https://github.com/Eugeny/tabby

### OpenSSH
### WinSCP
https://github.com/winscp/winscp
### webssh
https://pypi.org/project/webssh/
在linux机器上安装python环境，并且使用命令pip3 install webssh,装上这个模块

我们就可以在浏览器web页面登录我们的linux机器


### cubeshell
https://gitee.com/cubeiic-hanxuan/cube-shell