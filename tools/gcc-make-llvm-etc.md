## MSVC的STL
https://github.com/microsoft/STL
## Clang
Clang是一个C语言、C++、Objective-C语言的轻量级编译器
Clang 是一个 C++ 编写、基于 LLVM、发布于 LLVM BSD 许可证下的 C/C++/Objective C/Objective C++ 编译器，其目标（之一）就是超越 GCC。
## gcc
gcc 是GNU编译器套件，是Linux下默认的C/C++编译器.在windows环境下可以通过MinGw等GNU for Windows类工具使用gcc编译套件

gcc最简单的编译命令如下

`gcc -c hello.c -ohello`

gcc -c 后可跟多个输入源文件，最终输出的可执行文件以-o表示.
-o后紧着希望生成的可执行文件的名称。
-c 选项表示只编译源文件，而不进行链接，所以对于链接中的错误是无法发现的
如果不使用-c选项则仅仅生成一个可执行文件，没有目标文件。

##Make
这里的make 代指 GNU Make,首先我们看下GNU Make的英文释义

GNU Make is a tool which controls the generation of executables and other non-source files of a program from the program's source files.

Make gets its knowledge of how to build your program from a file called the makefile, which lists each of the non-source files and how to compute it from other files. When you write a program, you should write a makefile for it, so that it is possible to use Make to build and install the program.

可以看出make工具的定义是通过编写的makefile脚本文件描述整个工程的编译、链接规则；通过脚本文件，对于复杂的工程也可以只通过一个命令就完成整个编译过程。

## CMake
CMake是一个跨平台的编译工具。事实上Cmake并不直接构建出最终的软件，而是产生不同平台标准的构建档(如 Unix的Makefile 或是 Windows Visual C++的 projects/workspaces),然后再依一般的构建方式使用。

CMkae目前主要使用场景是作为make的上层工具，产生可移植的makefile文件。

## Apple LLVM compiler 4.2和LLVM GCC 4.2
[LLVM和GCC的区别](https://www.cnblogs.com/zuopeng/p/4141467.html)
https://gcc.gnu.org/
[开发者将 GCC 的 JIT 库移植到 Windows](https://gcc.gnu.org/pipermail/gcc-patches/2020-May/546384.html)