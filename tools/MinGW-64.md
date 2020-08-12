https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/

CGO
https://www.cnblogs.com/sunylat/p/6413628.html


目前主要有三种不同的线程库的定义，分别是Win32，OS/2，以及POSIX，前两种定义只适合于他们各自的平台，而POSIX 定义的线程库是适用于所有的计算平台的。我这里选的是threads-posix。
C++ Exceptions有DWARF、SJLJ、SEH三种处理方式。对应的我们这里选择的是seh方式。

比如：x86_64-posix-seh

[MinGW、MinGW-w64 与TDM-GCC 应该如何选择？](https://www.zhihu.com/question/39952667)
- MinGW已经不推荐使用。只有32位版，更新速度也不怎么样。
- MinGW-w64更新最快，基本上gcc更近后几周内就会跟进。32位和64位都提供。
- TDM-GCC，更新速度也不怎么样，同时提供32位和64位。涉及64位时，TDM-GCC和MinGW-w64还有一个重要的区别，64位的TDM-GCC既能编译64位binary，也能编译32位binary（用-m32参数）。
而MinGW-w64无此能力，你需要装32位和64位的两套MinGW w64 tool chain来编译两种binary。

TDM-GCC是非官方组织提供及维护的编译器集成包。换句话说，他不是由MinGW或MinGW-w64项目官方提供的编译器集成包。
但他基于MinGW及MinGW-w64。而且，他也有他的优点：
例如，TDM-GCC集成了：
最新的稳定版本的GCC工具集，
 一些实用的补丁，
MinGW及MinGW-w64中的运行时API。

[MinGW](http://www.mingw.org/) 

[MinGW-w64](http://mingw-w64.sourceforge.net/)

https://sourceforge.net/projects/mingw-w64/files/

https://sourceforge.net/projects/mingw-w64/files/mingw-w64/mingw-w64-release/

https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win32/Personal%20Builds/mingw-builds/installer/mingw-w64-install.exe

[TDM-GCC](http://tdm-gcc.tdragon.net/)

[TDM-GCC download](http://tdm-gcc.tdragon.net/download)



Cygwin - Windows下的Linux开发环境
http://www.cygwin.com/
阿里云镜像
http://mirrors.aliyun.com/cygwin/


## Cygwin VS MinGW
1. 从目标上说
- MinGW 是让Windows 用户可以用上GNU 工具，比如GCC。
- Cygwin 提供完整的类Unix 环境，Windows 用户不仅可以使用GNU 工具，理论上Linux 上的程序只要用Cygwin 重新编译，就可以在Windows 上运行。

2. 从能力上说
- 如果程序只用到C/C++ 标准库，可以用MinGW 或Cygwin 编译。
- 如果程序还用到了POSIX API，则只能用Cygwin 编译。

3. 从依赖上说
- 程序经MinGW 编译后可以直接在Windows 上面运行。
- 程序经Cygwin 编译后运行，需要依赖安装时附带的cygwin1.dll

