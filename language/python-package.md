[TOC]

## 构建独立的Python应用程序
目录对python程序进行打包方式主要有5种: py2exe、py2app，pyinstaller，cx_Freeze，nuitka
### pyinstaller
打包成单文件所使用的命令为：
`pyinstaller -Fw --icon=h.ico auto_organize_gui.py --add-data="h.ico;/"`
打包成文件夹所使用的命令为：
`pyinstaller -w --icon=h.ico auto_organize_gui.py --add-data="h.ico;."`
如何防止exe被反编译呢
`pyinstaller -Fw --icon=h.ico auto_organize_gui.py --add-data="h.ico;/" --key 123456`
#### auto-py-to-exe
auto-py-to-exe 是基于 pyinstaller 的
https://github.com/brentvollebregt/auto-py-to-exe

#### 将 exe 文件反编译成 Python 脚本
1. 抽取exe中的pyc文件
提取pyc文件的方法：
    1. 通过 pyinstxtractor.py 脚本提取pyc文件
    1. 通过 pyi-archive_viewer 工具提取pyc文件

`pyinstxtractor.py` 脚本可以 [python-exe-unpacker](https://github.com/countercept/Python-exe-unpacker) 中下载，下载该项目后把其中的pyinstxtractor.py脚本文件复制到与exe同级的目录。然后进入exe所在目录的cmd执行：`python pyinstxtractor.py auto_organize_gui.exe`

pyi-archive_viewer是PyInstaller自己提供的工具，它可以直接提取打包结果exe中的pyc文件。详细介绍可参考[官方文档](https://pyinstaller.readthedocs.io/en/stable/advanced-topics.html#using-pyi-archive-viewer)。
执行pyi-archive_viewer [filename]即可查看 exe 内部的文件结构：`pyi-archive_viewer auto_organize.exe` 输入?号，查看操作命令。
pyi-archive_viewer 工具操作起来比较麻烦，一次只能提取一个文件.


2. 反编译pyc文件为py脚本
有很多对pyc文件进行解密的网站，例如：https://tool.lu/pyc/
不过我们直接使用 uncompyle6 库进行解码，使用pip可以直接安装：`pip install uncompyle6`
`uncompyle6 xxx.pyc>xxx.py` or `uncompyle6 -o xxx.py xxx.pyc`
对于有些不是`.pyc`的文件我们可以人工修改后缀。

对于从pyinstaller提取出来的pyc文件并不能直接反编译，入口运行类共16字节的 magic 和 时间戳被去掉了。具体处理见：[Pyinstaller打包的exe之一键反编译py脚本与防反编译](https://xxmdmst.blog.csdn.net/article/details/119834495)
完整代码
```
#!/usr/bin/env python
# coding: utf-8

# 提取exe中的pyc
import os
import sys
import pyinstxtractor
from uncompyle6.bin import uncompile
import shutil


# 预处理pyc文件修护校验头
def find_main(pyc_dir):
    for pyc_file in os.listdir(pyc_dir):
        if not pyc_file.startswith("pyi-") and pyc_file.endswith("manifest"):
            main_file = pyc_file.replace(".exe.manifest", "")
            result = f"{pyc_dir}/{main_file}"
            if os.path.exists(result):
                return main_file


def uncompyle_exe(exe_file, complie_child=False):
    sys.argv = ['pyinstxtractor', exe_file]
    pyinstxtractor.main()
    # 恢复当前目录位置
    os.chdir("..")

    pyc_dir = os.path.basename(exe_file)+"_extracted"
    main_file = find_main(pyc_dir)

    pyz_dir = f"{pyc_dir}/PYZ-00.pyz_extracted"
    for pyc_file in os.listdir(pyz_dir):
        if pyc_file.endswith(".pyc"):
            file = f"{pyz_dir}/{pyc_file}"
            break
    else:
        print("子文件中没有找到pyc文件，无法反编译！")
        return
    with open(file, "rb") as f:
        head = f.read(4)

    if os.path.exists("pycfile_tmp"):
        shutil.rmtree("pycfile_tmp")
    os.mkdir("pycfile_tmp")
    main_file_result = f"pycfile_tmp/{main_file}.pyc"
    with open(f"{pyc_dir}/{main_file}", "rb") as read, open(main_file_result, "wb") as write:
        write.write(head)
        write.write(b"\0"*12)
        write.write(read.read())
    
    if os.path.exists("py_result"):
        shutil.rmtree("py_result")
    os.mkdir("py_result")
    sys.argv = ['uncompyle6', '-o',
                f'py_result/{main_file}.py', main_file_result]
    uncompile.main_bin()

    if not complie_child:
        return
    for pyc_file in os.listdir(pyz_dir):
        if not pyc_file.endswith(".pyc"):
            continue
        pyc_file_src = f"{pyz_dir}/{pyc_file}"
        pyc_file_dest = f"pycfile_tmp/{pyc_file}"
        print(pyc_file_src, pyc_file_dest)
        with open(pyc_file_src, "rb") as read, open(pyc_file_dest, "wb") as write:
            write.write(read.read(12))
            write.write(b"\0"*4)
            write.write(read.read())

    os.mkdir("py_result/other")
    for pyc_file in os.listdir("pycfile_tmp"):
        if pyc_file==main_file+".pyc":
            continue
        sys.argv = ['uncompyle6', '-o',
                    f'py_result/other/{pyc_file[:-1]}', f'pycfile_tmp/{pyc_file}']
        uncompile.main_bin()
```



### 使用PyOxidizer构建独立的Python应用程序
#### PyOxidizer
PyOxidizer是indygreg开源了一个项目，用于构建独立的Python应用程序。文档在这儿

有趣的是, PyOxidizer本身基于Rust, 个中原因[indygreg](https://github.com/indygreg)在[这篇文章](https://gregoryszorc.com/blog/2019/06/24/building-standalone-python-applications-with-pyoxidizer/)里做了阐述。

#### 安装PyOxidizer
cargo install pyoxidizer

#### 创建项目
pyoxidizer init pyapp

该命令将创建一个支持嵌入Python的Rust项目。运行完命令，将打印相关信息及下一步该做什么的提示。

根据提示信息，依次运行:
```
cd pyapp
pyoxidizer build # 第一次运行会下载对应平台的Python解释器
# pyoxidizer run # pyoxidizer run将启动了一个Rust可执行文件，它启动了一个交互式Python调试器！尝试输入一些Python代码：
```

#### 自定义Python和打包行为

项目根目录里有一个自动生成的pyoxidizer.toml文件，该文件决定默认运行时行为。
找到`[[embedded_python_run]]`部分, 这部分决定Python解释器启动时要执行的操作,调整为:
```
[[embedded_python_run]]
mode = "eval"
code = "import uuid; print(uuid.uuid4())"
```

现在我们告诉解释器启动时运行:eval(import uuid; print(uuid.uuid4())

pyoxidizer run输出的结果为: 4ef94bc0-4cbe-4404-9269-0690fec68094

#### 打包第三方库
接下来，让我们试着打包现有的Python应用程序！

我们试着打包第三方库:pyflakes

编辑配置文件:pyoxidizer.toml,使得:

```
[[packaging_rule]]
type = "pip-install-simple"
package = "pyflakes==2.1.1"
```
以及
```
[[embedded_python_run]]
mode = "eval"
code = "from pyflakes.api import main; main()"
```

这将告诉PyOxidizer你要安装pyflakes的2.1.1版本。在构建时，会运行pip install pyflakes==2.1.1,并将它们添加到生成的二进制文件中。我们试试看：pyoxidizer run -- --help

新的pyoxidizer.toml文件应该类似于：
```
# Multiple [[python_distribution]] sections elided for brevity.

[[build]]
application_name = "pyflakes"

[[embedded_python_config]]
raw_allocator = "system"

[[packaging_rule]]
type = "stdlib-extensions-policy"
policy = "all"

[[packaging_rule]]
type = "stdlib"
include_source = false

[[packaging_rule]]
type = "pip-install-simple"
package = "pyflakes==2.1.1"

[[embedded_python_run]]
mode = "eval"
code = "from pyflakes.api import main; main()"
```