## emacs 
著名的集成开发环境和文本编辑器
[中文论坛](https://emacs-china.org/)
[book](http://book.emacs-china.org/)

## 安装
https://www.gnu.org/software/emacs/download.html#windows

yum install emacs

or

```
### First: install the following software（安装软件包）

     sudo yum -y groupinstall "Development Tools"   
     # 在有些云服务器中，这一步服务商已经做好了，当看到屏幕提示说: Development Tools 没有什么的，你就跳过这步

     sudo yum -y install gtk+-devel gtk2-devel
     sudo yum -y install libXpm-devel
     sudo yum -y install giflib-devel
     sudo yum -y install libtiff-devel libjpeg-devel
     sudo yum -y install ncurses-devel
     sudo yum -y install gpm-devel dbus-devel dbus-glib-devel dbus-python
     sudo yum -y install GConf2-devel pkgconfig

### Second: get emacs 24.4 （取得emacs 24.4）
     wget http://ftp.gnu.org/pub/gnu/emacs/emacs-24.4.tar.gz
     sudo tar xvf emace-24.4.tar.gz

### Third: compile emacs 24.4 (编译)
     sudo ./configure -prefix=/usr/local -with-x-toolkit=gtk
     sudo make && make install

### Fourth: (取得一种极简配置)
    git clone https://github.com/doosolar/emacs.git

### Fifth: link the config files （将配置文件链接生效）
    # 这一步，用软链接，在 ~ 目录中，链接出    
    # .emacs
    # .emacs.d
```

```
cd 自己想安装的目录
sudo wget ftp://ftp.gnu.org/gnu/emacs/emacs-24.4.tar.gz （也可以直接进入emacs的官网找到你想要的版本）
tar zxvf emacs-24.4.tar.gz （解压）
cd emacs-24.4 （进入加压好的文件）
./configure （编译）
make && make install
```