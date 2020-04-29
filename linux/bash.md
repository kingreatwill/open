https://github.com/vuuihc/bash-guide


[Bash Reference Manual](https://www.gnu.org/software/bash/manual/bash.html)
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
or bash --version
```