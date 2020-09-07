

## 注释风格

### reStructuredText（PyCharm默认）
[reStructuredText](https://docutils.sourceforge.io/docs/user/rst/quickref.html)（[PyCharm默认](https://www.jetbrains.com/help/pycharm/restructured-text.html)）
```
def func(path, field_storage, temporary):
    '''基本描述

    详细描述

    :param path: The path of the file to wrap
    :type path: str
    :param field_storage: The :class:`FileStorage` instance to wrap
    :type field_storage: FileStorage
    :param temporary: Whether or not to delete the file when the File instance is destructed
    :type temporary: bool
    :returns: A buffered writable file descriptor
    :rtype: BufferedFileStorage
    '''
    pass
```

### NumPy
[NumPy Style](https://www.sphinx-doc.org/en/master/usage/extensions/example_numpy.html)
```
def func(path, field_storage, temporary):
    '''基本描述

    详细描述

    Parameters
    ----------
    path : str
        The path of the file to wrap
    field_storage : FileStorage
        The :class:`FileStorage` instance to wrap
    temporary : bool
        Whether or not to delete the file when the File instance is destructed

    Returns
    -------
    BufferedFileStorage
        A buffered writable file descriptor
    '''
    pass
```

### Google（官方推荐）
[Google Style](http://www.sphinx-doc.org/en/master/usage/extensions/example_google.html)
```
def func(path, field_storage, temporary):
    '''基本描述

    详细描述

    Args:
        path (str): The path of the file to wrap
        field_storage (FileStorage): The :class:`FileStorage` instance to wrap
        temporary (bool): Whether or not to delete the file when the File instance is destructed

    Returns:
        BufferedFileStorage: A buffered writable file descriptor
    '''
    pass
```

### 比较
中文版：https://www.osgeo.cn/sphinx/usage/extensions/napoleon.html
英文版：https://www.sphinx-doc.org/en/master/usage/extensions/napoleon.html

风格 |	特点|	适用
---|---|---
reStructuredText |	用冒号分隔|	PyCharm默认
NumPy	|用下划线分隔	|倾向垂直，长而深的文档
Google|	用缩进分隔 |	倾向水平，短而简单的文档

### 参考文献
[Python风格规范 — Google 开源项目风格指南](http://zh-google-styleguide.readthedocs.io/en/latest/google-python-styleguide/python_style_rules/#comments)
[Sphinx入门——快速生成Python文档](https://xercis.blog.csdn.net/article/details/103970663)
[Sphinx](https://www.osgeo.cn/sphinx/index.html)是Python文档生成器
