https://www.gnu.org/software/gdb/


使用gdbgui在浏览器中远程多语言调试
https://github.com/cs01/gdbgui

使用cgdb为gdb插上翅膀
http://cgdb.github.io/

cgdb是GNU调试器(GDB)的一个轻量级curses(基于终端)接口。除了标准的gdb控制台之外，cgdb还提供了一个分屏视图，在执行时显示源代码。键盘界面是仿照vim设计的，所以vim用户使用cgdb就像在家里一样。

kdbg就是Lniux平台的图形界面调试器。更准确而言，kdbg不是一个调试器只是gdb的一个前端图形界面，后面调用的还是gdb。
apt-get install -y kdbg


[Voltron](https://github.com/snare/voltron)
Voltron是用Python编写的可扩展调试器UI工具包。 它旨在通过启用实用程序视图的附件来改善各种调试器（LLDB，GDB，VDB和WinDbg）的用户体验，这些视图可以从调试器主机检索和显示数据。 通过在其他TTY中运行这些视图，您可以构建自定义的调试器用户界面来满足您的需求。
