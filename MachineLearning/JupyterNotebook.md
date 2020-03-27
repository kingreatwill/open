
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
```
# 配置清华PyPI镜像（如无法运行，将pip版本升级到>=10.0.0）
pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
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