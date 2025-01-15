<!-- {% raw %} -->
### jupytext & MyST & Jupyter Book
- [Jupyter Book](https://github.com/executablebooks/jupyter-book)
- [MyST （Markedly Structured Text） Markdown语言](https://myst-parser.readthedocs.io/)
- [jupytext](https://github.com/mwouts/jupytext) 提供内容管理器，允许 Jupyter 将 notebook 保存为你喜欢的格式，来补充或替代传统的.ipynb 文件。


### Jupyter Notebook IDE
https://github.com/zasper-io/zasper

### nteract
nteract可以直接打开本地ipynb文件，可以直接双击ipynb文件直接打开笔记进行编辑，再不需要像以前一样，要先运行jupyter notebook，然后在浏览器中打开ipynb文件。

### Jupyter 结合 Excel

https://mp.weixin.qq.com/s/D09CXCeYxqpOBVH61B4MUA

**一定要安装64位的excel和64位的python**

pip install pyxll -i https://mirrors.aliyun.com/pypi/simple/

pyxll install # 安装好了PyXLL在 Excel中的插件，下一步就是安装pyxll-jupyter

pip install pyxll-jupyter -i https://mirrors.aliyun.com/pypi/simple/

安装完毕后，启动Excel，将在PyXLL选项卡中看到一个新的Jupyter按钮。

```
%xl_get   获取excel数据

%xl_set 将Python中的数据移到Excel

%xl_set df



%xl_plot 在Excel中使用Python绘图


从Excel调用Python函数
替代VBA脚本 https://www.pyxll.com/docs/userguide/vba.html
```

> excel也有其他第三方支持python脚本https://www.cnblogs.com/connect/p/office-excel-python-conf.html

### Time Profiling
https://pynash.org/2013/03/06/timing-and-profiling/
一个%是一行，两个%是一个cell
```
In [1]: %time?
In [2]: %timeit?
In [3]: %prun?
In [4]: %lprun?
In [5]: %mprun?
In [6]: %memit?

Time profiling does exactly what it sounds like - it tells you how much time it took to execute a script, which may be a simple one-liner or a whole module.

%time
See how long it takes a script to run.

In [7]: %time {1 for i in xrange(10*1000000)}
CPU times: user 0.72 s, sys: 0.16 s, total: 0.88 s
Wall time: 0.75 s
%timeit
See how long a script takes to run averaged over multiple runs.

In [8]: %timeit 10*1000000
10000000 loops, best of 3: 38.2 ns per loop
%timeit will limit the number of runs depending on how long the script takes to execute. Keep in mind that the timeit module in the standard library does not do this by default, so timing long running scripts that way may leave you waiting forever.

The number of runs may be set with with -n 1000, for example, which will limit %timeit to a thousand iterations, like this:

In [9]: %timeit -n 1000 10*1000000
1000 loops, best of 3: 67 ns per loop
Also note that the run-time reported will vary more when limited to fewer loops.

For example, do you know which of the following addition operations is better, x = 5; y = x**2, x = 5; y = x*x, x = np.uint8([5]); y = x*x, or y = np.square(x)? We will find out with timeit in the IPython shell.

In [10]: x = 5
In [11]: %timeit y=x**2
10000000 loops, best of 3: 73 ns per loop
In [12]: %timeit y=x*x
10000000 loops, best of 3: 58.3 ns per loop
In [15]: z = np.uint8([5])
In [17]: %timeit y=z*z
1000000 loops, best of 3: 1.25 us per loop
In [19]: %timeit y=np.square(z)
1000000 loops, best of 3: 1.16 us per loop
You can see that, x = 5 ; y = x*x is fastest and it is around 20x faster compared to Numpy. If you consider the array creation also, it may reach up to 100x faster. Cool, right? *(Numpy devs are working on this issue)*


%prun
See how long it took each function in a script to run.

In [10]: from time import sleep

In [11]: def foo(): sleep(1)

In [12]: def bar(): sleep(2)

In [13]: def baz(): foo(), bar()

In [14]: %prun baz()
7 function calls in 3.001 seconds

Ordered by: internal time

ncalls  tottime  percall  cumtime  percall filename:lineno(function)
     2    3.001    1.500    3.001    1.500 {time.sleep}
     1    0.000    0.000    3.001    3.001 <ipython-input-17-c32ce4852c7d>:1(baz)
     1    0.000    0.000    2.000    2.000 <ipython-input-11-2689ca7390dc>:1(bar)
     1    0.000    0.000    1.001    1.001 <ipython-input-10-e11af1cc2c91>:1(foo)
     1    0.000    0.000    3.001    3.001 <string>:1(<module>)
     1    0.000    0.000    0.000    0.000 {method 'disable' of '_lsprof.Profiler' objects}
%lprun
See how long it took each line in a function to run.

Create and edit a new module named foo.py in the same directory where you started IPython. Paste the following code in the file and jump back to IPython.

def foo(n):
    phrase = 'repeat me'
    pmul = phrase * n
    pjoi = ''.join([phrase for x in xrange(n)])
    pinc = ''
    for x in xrange(n):
        pinc += phrase
    del pmul, pjoi, pinc
Import the function and profile it line by line with %lprun. Functions to profile this way must be passed by name with -f.

In [15]: from foo import foo

In [16]: %lprun -f foo foo(100000)
Timer unit: 1e-06 s

File: foo.py
Function: foo at line 1
Total time: 0.301032 s

Line #      Hits         Time  Per Hit   % Time  Line Contents
==============================================================
     1                                           def foo(n):
     2         1            3      3.0      0.0      phrase = 'repeat me'
     3         1          185    185.0      0.1      pmul = phrase * n
     4    100001        97590      1.0     32.4      pjoi = ''.join([phrase for x in xrange(n)])
     5         1            4      4.0      0.0      pinc = ''
     6    100001        90133      0.9     29.9      for x in xrange(n):
     7    100000       112935      1.1     37.5          pinc += phrase
     8         1          182    182.0      0.1      del pmul, pjoi, pinc
Memory Profiling
%mprun
See how much memory a script uses line by line. Let’s take a look at the same foo() function that we profiled with %lprun - except this time we’re interested in incremental memory usage and not execution time.

In [17]: %mprun -f foo foo(100000)
Filename: foo.py

Line #    Mem usage    Increment   Line Contents
================================================
     1    20.590 MB     0.000 MB   def foo(n):
     2    20.590 MB     0.000 MB       phrase = 'repeat me'
     3    21.445 MB     0.855 MB       pmul = phrase * n
     4    25.020 MB     3.574 MB       pjoi = ''.join([phrase for x in xrange(n)])
     5    25.020 MB     0.000 MB       pinc = ''
     6    43.594 MB    18.574 MB       for x in xrange(n):
     7    43.594 MB     0.000 MB           pinc += phrase
     8    41.102 MB    -2.492 MB       del pmul, pjoi, pinc
%memit
See how much memory a script uses overall. %memit works a lot like %timeit except that the number of iterations is set with -r instead of -n.

In [18]: %memit -r 3 [x for x in xrange(1000000)]
maximum of 3: 75.320312 MB per loop
```

### 一、简介
Jupyter Notebook是一个开源的Web应用程序，允许用户创建和共享包含代码、方程式、可视化和文本的文档。它的用途包括：数据清理和转换、数值模拟、统计建模、数据可视化、机器学习等等。它具有以下优势：

可选择语言：支持超过40种编程语言，包括Python、R、Julia、Scala等。

分享笔记本：可以使用电子邮件、Dropbox、GitHub和Jupyter Notebook Viewer与他人共享。

交互式输出：代码可以生成丰富的交互式输出，包括HTML、图像、视频、LaTeX等等。

大数据整合：通过Python、R、Scala编程语言使用Apache Spark等大数据框架工具。支持使用pandas、scikit-learn、ggplot2、TensorFlow来探索同一份数据。

### 二、安装与运行

虽然Jupyter可以运行多种编程语言，但Python是安装Jupyter Noterbook的必备条件（Python2.7，或Python3.3以上）。有两种安装方式：使用Anaconda安装或使用pip命令安装。关于安装的全部信息可以在官网读到：[安装Jupyter](https://jupyter.org/install.html)。

##### 2.1使用Anaconda安装
对于小白，强烈建议使用Anaconda发行版安装Python和Jupyter，其中包括Python、Jupyter Notebook和其他常用的科学计算和数据科学软件包。

首先，下载[Anaconda](https://www.anaconda.com/distribution/)。建议下载Anaconda的最新Python 3版本。其次，请按照下载页面上的说明安装下载的Anaconda版本。最后，安装成功！



###### miniconda 无界面版本

1. 安装miniconda
在安装过程中需要勾选“Add Anaconda to the system PATH environment variable”选项

https://conda.io/en/master/miniconda.html

2. 使用conda创建虚拟（运行）环境。conda和pip默认使用国外站点来下载软件，我们可以配置国内镜像来加速下载（国外用户无须此操作）。

[Anaconda 镜像使用帮助](https://mirrors.tuna.tsinghua.edu.cn/help/anaconda/)
```
# 配置清华PyPI镜像（如无法运行，将pip版本升级到>=10.0.0）
pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
```
常用命令
```
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main/
conda config --set show_channel_urls yes
# 验证.condarc文件

添加环境变量：
D:\ProgramData\Anaconda3
D:\ProgramData\Anaconda3\Library\bin （如果想在cmd里面直接conda activate ,这个要放在下个路径的前面，这个里面是conda.bat,下个里面是conda.exe。如果你只想使用conda命令，你可以值添加这一个变量，当然也可是D:\ProgramData\Anaconda3\condabin）
D:\ProgramData\Anaconda3\Scripts
D:\ProgramData\Anaconda3\Library\mingw-w64\bin


conda info --envs
conda create --name penter python=3.8
conda activate gluon

conda list: 看这个环境下安装的包和版本
conda install numpy scikit-learn: 安装numpy sklearn包
conda env remove -n yourEnv: 删除你的环境
conda env list: 查看所有的环境
```

3. 创建环境和安装环境依赖

environment.yml 文件内容
```
name: gluon
dependencies:
- python=3.6
- pip:
  - mxnet==1.5.0
  - d2lzh==1.0.0
  - jupyter==1.0.0
  - matplotlib==2.2.2
  - pandas==0.23.4

```
```
conda env create -f environment.yml
```

4. 激活创建的环境
```
conda activate gluon
```
退出虚拟环境
```
conda deactivate
```
更新运行环境
```
conda env update -f environment.yml
```

conda info --envs
列出我的所有环境

nvidia-smi 查看显卡



5. 打开Jupyter记事本
```
jupyter notebook

set MXNET_GLUON_REPO=https://apache-mxnet.s3.cn-north-1.amazonaws.com.cn/ jupyter notebook
```




##### 2.2使用pip命令安装
对于有经验的Python用户，可以使用Python的包管理器pip而不是Anaconda 来安装Jupyter 。 
如果已经安装了Python 3：
```cmd
python3 -m pip install --upgrade pip
pi3 install jupyter
```

##### 2.3运行Jupyter Notebook
```
jupyter notebook
```

##### 2.4 更改jupyter的工作路径
```
jupyter notebook --generate-config
```
会在用户目录下生产.jupyter文件夹下生成一个jupyter_notebook_config.py文件
修改
#c.NotebookApp.notebook_dir = ''

##### 2.5python命令在安装目录下
D:\Programs\Python\Python37\Scripts
(dotnet 的在用户目录\.dotnet\tools)


##### 2.6其它参阅
https://blog.51cto.com/853056088/2162189


### 三、docker运行

### 四、VS Code
https://zhuanlan.zhihu.com/p/87887002

### .ipynb 文件转换为 .md, .tex, .pdf 文件


[命令行转换](https://nbconvert.readthedocs.io/en/latest/config_options.html#cli-flags-and-aliases)  
参考-[ 如何将ipynb转换为html，md，pdf等格式](https://cloud.tencent.com/developer/article/1008571)

`jupyter nbconvert --to format(such as pdf, latex, ...) [-template template_name such as article, report, ...] notebook.ipynb`
如：`jupyter nbconvert --to markdown matrixcookbook.ipynb`
支持：` [‘asciidoc’, ‘custom’, ‘html’, ‘latex’, ‘markdown’, ‘notebook’, ‘pdf’, ‘python’, ‘rst’, ‘script’, ‘slides’, ‘webpdf’]`

[Jupytext](https://jupytext.readthedocs.io/en/latest/): 支持自动转换（Jupyter Notebooks as Markdown Documents, Julia, Python or R scripts）


> 如果需要导出pdf，需要安装[xelatex](https://nbconvert.readthedocs.io/en/latest/install.html#installing-tex), [win 下载](https://miktex.org/download)，当然也可以安装[TeX Live ++](https://www.tug.org/texlive/)


### Jupyter输出渲染latex公式
```
# display(Math(latex_s))和display(Latex(latex_s))输出的是latex类型，
# display(Markdown(latex_s))输出的是markdown
# 推荐markdown和Latex；而Math只支持纯latex
from IPython.display import display,Latex, SVG, Math, Markdown

latex_s = r"$\frac{{\partial {}}}{{\partial {}}}$".format(1, 2)
display(Math(latex_s))
```

### Jupyter cell同时输出多行
```
# 如果对带有一个变量或是未赋值语句的cell执行操作，Jupyter将会自动打印该变量而无需一个输出语句。
from IPython.core.interactiveshell import InteractiveShell
InteractiveShell.ast_node_interactivity = "all" #默认为'last'
```

### python中print打印显示颜色
- 方法一
https://blog.csdn.net/qq_34857250/article/details/79673698
```
print('This is a \033[1;35m test \033[0m!')
print('This is a \033[1;32;43m test \033[0m!')
print('\033[1;33;44mThis is a test !\033[0m')
```

- 方法二
```
from sys import stdout
from colorama import Fore

# Option 1
stdout.write(Fore.RED + "Test")
# Option 2
print(Fore.GREEN + "Test")
```
- 方法三
```
%%
class ListOfColoredStrings(object):
    def __init__(self, *args):
        """
        Expected input:
        args = ["word_1", "color_1"], ["word_2", "color_2"]

        :param args: pairs of [word, color], both given as strings
        """
        self.strings = [a[0] for a in args]
        self.colors = [a[1] for a in args]

    def _repr_html_(self):
        return ''.join( [
            "<span class='listofstr' style='color:{}'>{}</span>"
                .format(self.colors[i], self.strings[i])
            for i in range(len(self.strings))
            ])

%%html
<style type='text/css'>
     span.listofstr {
          margin-left: 5px
     }
</style>

%%
ListOfColoredStrings(["hi", "red"], ["hello", "green"])
```

- 方法四
```
from IPython.display import Markdown, display
import html

for tag in tags:    
    tag = html.escape(tag)
    display(Markdown((f'this is your tag: <text style=color:red>{tag}</text>')))
```
- 方法五
```
from IPython.display import HTML, display
display(HTML(f'this is your tag: <text style=color:red>xx</text>'))
```

### 自定义magic
https://blog.csdn.net/u011702002/article/details/85829654
```
from IPython.core.magic import register_cell_magic, register_line_magic, register_line_cell_magic
```

<!-- {% endraw %} -->