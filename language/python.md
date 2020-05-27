
## 学习资源
数据科学Python笔记本:深度学习(TensorFlow、Theano、Caffe、Keras)、scikit-learn、Kaggle、大数据(Spark、Hadoop MapReduce、HDFS)、matplotlib、panda、NumPy、SciPy、Python essentials、AWS和各种命令行。
https://github.com/donnemartin/data-science-ipython-notebooks
## IronPython vs. Python .NET
https://stackoverflow.com/questions/1168914/ironpython-vs-python-net
https://ironpython.net/
https://github.com/IronLanguages/ironpython2
https://github.com/IronLanguages/ironpython3
https://github.com/pythonnet/pythonnet

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

