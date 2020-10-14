

## 构建独立的Python应用程序
目录对python程序进行打包方式主要有5种: py2exe、py2app，pyinstaller，cx_Freeze，nuitka

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