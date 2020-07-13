
## 学习资源
数据科学Python笔记本:深度学习(TensorFlow、Theano、Caffe、Keras)、scikit-learn、Kaggle、大数据(Spark、Hadoop MapReduce、HDFS)、matplotlib、panda、NumPy、SciPy、Python essentials、AWS和各种命令行。
https://github.com/donnemartin/data-science-ipython-notebooks
## IronPython vs. Python .NET
https://stackoverflow.com/questions/1168914/ironpython-vs-python-net
https://ironpython.net/
https://github.com/IronLanguages/ironpython2
https://github.com/IronLanguages/ironpython3
https://github.com/pythonnet/pythonnet

## Python解释器&编译器
### CPython (C), which is the most common
### Jython (Java)
### IronPython (.NET)
### PyPy (Python)
### Stackless
Stackless Python是Python编程语言的增强版本
https://github.com/stackless-dev/stackless

Microthreads: tasklets wrap functions allowing them to be launched as microthreads.
Channels channels can be used for bidirectional communication between tasklets.
Scheduling a round robin scheduler is built in. It can be used to schedule tasklets either cooperatively or preemptively.
Serialisation: tasklets can be serialised to disk through pickling for later resumption of execution.


### RustPython
https://github.com/RustPython/RustPython


Each of these implementations offer some benefits: Jython, for example, compiles Python source code to Java byte code, then routes it to the Java Virtual Machine. Because Python code is translated to Java byte code, it looks and feels like a true Java program at runtime and so it integrates well with Java applications.

IronPython is well-integrated with .NET, which means IronPython can use the .NET framework and Python libraries or vice versa.

We want to unlock the same possibilities that Jython and IronPython enable, but for the Rust programming language. In addition, thanks to Rusts’ minimal runtime, we’re able to compile RustPython to WebAssembly and allow users to run their Python code easily in the browser.



## Style Guide
Style Guide: https://www.python.org/dev/peps/pep-0008/

### autopep8
helps to format code automatically
使用autopep8，以PEP 8规范，自动排版Python代码
```
python -m autopep8 -r --global-config .config-pep8 -i .

autopep8 --in-place --aggressive --aggressive file.py


```
### isort 
helps to order imports automatically
```
python -m isort -rc .
#When you use newly 3rd party modules, add it to .isort.cfg to keep import order correct
```

### Docstring style: Google Style
https://sphinxcontrib-napoleon.readthedocs.io/en/latest/example_google.html

### RST与Python类似Javadoc与Java.
You can customise CHANGELOG.rst with commit messages following .gitchangelog.rc
It generates readable changelog
如果下载了别人的Python源码，里面有rst文件夹，我们可以转为html后用浏览器打开

```
pip install sphinx
pip install -i http://pypi.douban.com/simple/ sphinx_rtd_theme --trusted-host pypi.douban.com


sphinx-build -b html docs build
```


### pip
https://stackoverflow.com/questions/30306099/pip-install-editable-vs-python-setup-py-develop

pip install -e
https://stackoverflow.com/questions/42609943/what-is-the-use-case-for-pip-install-e/59667164

#### Installing from a VCS
https://the-hitchhikers-guide-to-packaging.readthedocs.io/en/latest/pip.html#installing-from-a-vcs

#### cmd
pip install -h

pip install http://dist.repoze.org/PIL-1.1.6.tar.gz

pip install -U git+https://github.com/madmaze/pytesseract.git
git clone https://github.com/madmaze/pytesseract.git
cd pytesseract && pip install -U .
pip3 install face_recognition -i  http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com

pip install -r requirements.txt

### 生成requirements.txt
1.
```
pip freeze > requirements.txt

//pip install pycryptodome -i https://pypi.doubanio.com/simple/
//pip install -i https://pypi.doubanio.com/simple/ -r requirements.txt
```
2.
```
pip install pipreqs
// 生成
pipreqs .
pipreqs --encoding=utf-8 .
// 更新
pipreqs --force .
pipreqs --force --encoding=utf-8 .
```

#### pipenv

```
pip install --user pipenv
python -m pipenv lock --clear

# If you experience the below error, then refer pypa/pipenv#187 to solve it.
# Locking Failed! unknown locale: UTF-8


python -m pipenv install --dev --system

pre-commit install
```

查看依赖关系
```
pipenv graph
```

## python解释器(python.exe)和python启动器(py.exe)

（在 Unix 系统中是 Control-D，Windows 系统中是 Control-Z）就退出解释器并返回退出状态为0。如果这样不管用，你还可以写这个命令退出：quit() 和 exit()。

### python启动器
多环境（如;2.7和3.7 和3.8）共存
py -h
[适用于Windows的Python启动器](https://docs.python.org/zh-cn/3/using/windows.html#launcher)

### python解释器

如果不使用默认编码，要声明文件所使用的编码，文件的 第一 行要写成特殊的注释。语法如下所示：
https://docs.python.org/zh-cn/3/library/codecs.html#module-codecs
```
# -*- coding: encoding -*-
```
关于 第一行 规则的一种例外情况是，源码以 UNIX "shebang" 行 开头。这种情况下，编码声明就要写在文件的第二行。例如：
```
#!/usr/bin/env python3
# -*- coding: cp1252 -*-
```
在Unix系统中，Python 3.x解释器默认安装后的执行文件并不叫作 python，这样才不会与同时安装的Python 2.x冲突。
```
#! /usr/bin/python -v
```
然后Python将以 -v 选项启动



## 分布式计算框架
https://github.com/mars-project/mars
Mars是由阿里云高级软件工程师秦续业等人开发的一个基于张量的大规模数据计算的统一框架，目前它已在GitHub上开源。
该工具能用于多个工作站，而且即使在单块CPU的情况下，它的矩阵运算速度也比NumPy(MKL)快。


https://github.com/databricks/koalas


https://github.com/dask/dask
关于Python性能的一个常见抱怨是全局解释器锁(GIL)。由于GIL，同一时刻只能有一个线程执行Python字节码。因此，即使在现代的多核机器上，使用线程也不会加速计算。
但当你需要并行化到多核时，你不需要放弃使用Python，Dask库可以将计算扩展到多个内核甚至多个机器。某些设置可以在数千台机器上配置Dask，每台机器都有多个内核。

https://github.com/vaexio/vaex
Vaex是一个开源的DataFrame库(类似于Pandas)，对和你硬盘空间一样大小的表格数据集，它可以有效进行可视化、探索、分析甚至进行实践机器学习。

它可以在N维网格上计算每秒超过十亿(10^9)个对象/行的统计信息，例如均值、总和、计数、标准差等。使用直方图、密度图和三维体绘制完成可视化，从而可以交互式探索大数据。
Vaex使用内存映射、零内存复制策略获得最佳性能(不浪费内存)。

为实现这些功能，Vaex 采用内存映射、高效的核外算法和延迟计算等概念。所有这些都封装为类Pandas的API，因此，任何人都能快速上手。



https://github.com/cupy/cupy
CuPy是一个借助CUDA GPU库在英伟达GPU上实现Numpy数组的库。基于Numpy数组的实现，GPU自身具有的多个CUDA核心可以促成更好的并行加速。
CuPy接口是Numpy 的一个镜像，并且在大多情况下，它可以直接替换Numpy使用。只要用兼容的CuPy代码替换Numpy代码，用户就可以实现 GPU 加速。
CuPy支持Numpy的大多数数组运算，包括索引、广播、数组数学以及各种矩阵变换。


http://docs.cython.org/en/latest/
Cython是结合了Python和C的语法的一种语言，可以简单的认为就是给Python加上了静态类型后的语法，使用者可以维持大部分的Python语法，
而不需要大幅度调整主要的程式逻辑与算法。但由于会直接编译为二进制程序，所以性能较Python会有很大提升。
```
pip install Cython

from cpython cimport array
import array
cdef array.array a = array.array('i', [1, 2, 3])
cdef int[:] ca = a
print(ca[0])
```

https://github.com/ray-project/ray

Ray 是由加州大学伯克利分校 RISELab 开源的新兴人工智能应用的分布式框架。它实现了一个统一的接口、分布式调度器、分布式容错存储，以满足高级人工智能技术对系统最新的、苛刻的要求。Ray 允许用户轻松高效地运行许多新兴的人工智能应用，例如，使用 RLlib 的深度强化学习、使用 Ray Tune 的可扩展超参数搜索、使用 AutoPandas 的自动程序合成等等。






深度学习、强化学习、自动机器学习（AutoML）

Ray On Spark
https://github.com/intel-analytics/analytics-zoo

Apache Spark/Flink & Ray上的分布式Tensorflow、Keras和PyTorch

## Python代码的静态类型分析器
https://github.com/google/pytype

