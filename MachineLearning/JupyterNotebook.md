
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