https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/

CGO
https://www.cnblogs.com/sunylat/p/6413628.html


目前主要有三种不同的线程库的定义，分别是Win32，OS/2，以及POSIX，前两种定义只适合于他们各自的平台，而POSIX 定义的线程库是适用于所有的计算平台的。我这里选的是threads-posix。
C++ Exceptions有DWARF、SJLJ、SEH三种处理方式。对应的我们这里选择的是seh方式。

比如：x86_64-posix-seh