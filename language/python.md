
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