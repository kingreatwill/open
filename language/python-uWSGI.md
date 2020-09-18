## Flask docker部署
https://uwsgi-docs.readthedocs.io/en/latest/WSGIquickstart.html
https://uwsgi-docs.readthedocs.io/en/latest/WSGIquickstart.html#deploying-flask
https://uwsgi-docs.readthedocs.io/en/latest/WSGIquickstart.html#deploying-django
先定义好uwsgi配置文件
uwsgi.ini
```
socket=/tmp/app.sock
chmod-socket=666
pidfile=/etc/nginx/uwsgi.pid
chdir=/learn_flask/app
master=true
wsgi-file=app.py
http=0.0.0.0:5000
callable=app
processes=8
threads=4
lazy-apps=true
```
启动uwsgi服务器
uwsgi --ini uwsgi.ini

#### uwsgi http、http-socket和socket配置项

在uwsgi的[配置说明书](https://uwsgi-docs.readthedocs.io/en/latest/Configuration.html)中有配置使用http、http-socket和socket，我对http-socket和socket不是太懂，搜了一下资料。

在uwsgi的[注意事项](https://uwsgi-docs.readthedocs.io/en/latest/ThingsToKnow.html)中有强调说http和http-socket是两个不同的“野兽”，这是什么意思呢?简单的说，如果我们想直接将uwsgi用作服务器（例如Apache和nginx那样）直接暴露在公网那么就使用http；如果有单独的服务器（例如Apache或者nginx），由服务器将请求转发给uwsgi处理，并且使用http协议，那么此时使用http-socket。

而socket配置项又是什么意思呢？首先，按照uwsgi文档给出的解释是：bind to the specified UNIX/TCP socket using default protocol.也就是说指定UNIX/TCP socket作为默认的协议（[引](https://uwsgi-docs.readthedocs.io/en/latest/Options.html#socket)）。

UNIX/TCP socket其实是两类socket。UNIX socket是进程间的通信（[Inter Process Communication](https://www.geeksforgeeks.org/inter-process-communication-ipc/)），但只在同一台机器上；TCP/IP sockets允许进程通过网络通信。（[引](https://serverfault.com/a/124518/449456)）

在uwsgi中如果配置如下则是使用UNIX socket：
```
[uwsgi]
socket = /tmp/uwsgi.sock
```
如果配置如下则是使用TCP/IP socket：
```
[uwsgi]
socket = 127.0.0.1:8000
```
[在nginx也是对应的配置](https://uwsgi-docs.readthedocs.io/en/latest/Nginx.html)-[OR](https://dormousehole.readthedocs.io/en/latest/deploying/uwsgi.html#nginx)：
```
# uWSGI socket（unix socket）
uwsgi_pass unix:///tmp/uwsgi.sock;
include uwsgi_params;

# TCP sockets 一般我们部署到k8s时会和nginx分开部署
uwsgi_pass 127.0.0.1:3031;
include uwsgi_params;
```
但是官方文档将两种形式写在一起了，是不是两个都可以用([引](https://uwsgi-docs.readthedocs.io/en/latest/Configuration.html#ini-files))：
```
[uwsgi]
socket = /tmp/uwsgi.sock
socket = 127.0.0.1:8000
workers = 3
master = true
```
注：[uwsgi-socket是socket的别名](https://stackoverflow.com/questions/57112428/what-are-the-differences-between-http-and-socket-inside-of-ini-file-in-uwsgi/57113565#57113565)

#### 镜像

https://hub.docker.com/_/python
可以使用镜像python:3.8.2-alpine3.11  
or python:3.8-alpine3.12 # 默认最新python3.8
or python:3.8-alpine # 默认最新python3.8和最新alpine

### flask run 
https://dormousehole.readthedocs.io/en/latest/cli.html
```
Unix Bash （ Linux 、Mac 及其他）:

$ export FLASK_APP=hello
$ flask run
Windows CMD:

> set FLASK_APP=hello
> flask run
Windows PowerShell:

> $env:FLASK_APP = "hello"
> flask run
```
FLASK_APP=hello 指在hello模块中的app 或者create_app
名称被导入，自动探测一个应用（ app ）或者工厂（ create_app ）。
FLASK_APP 分三个部分：一是一个可选路径，用于设置当前工作文件夹；二是一 个 Python 文件或者带点的导入路径；三是一个可选的实例或工厂的变量名称。如果 名称是工厂，则可以选择在后面的括号中加上参数。以下演示说明：

FLASK_APP=src/hello
设置当前工作文件夹为 src 然后导入 hello 。

FLASK_APP=hello.web
导入路径 hello.web 。

FLASK_APP=hello:app2
使用 hello 中的 app2 Flask 实例。

FLASK_APP="hello:create_app('dev')"
调用 hello 中的 create_app 工厂，把 'dev' 作为参数。

如果没有设置 FLASK_APP ，命令会查找 wsgi.py 文件或者 app.py 文件并尝试探测一个应用实例或者工厂。

根据给定的导入，命令会寻找一个名为 app 或者 application 的应用实例。 如果找不到会继续寻找任意应用实例。如果找不到任何实例，会接着寻找名为 create_app 或者 make_app 的函数，使用该函数返回的实例。

当调用一个应用工厂时，如果工厂接收一个名为 info 的参数，那么 class:~cli.ScriptInfo 实例会被作为一个关键字参数传递。如果括号紧随着工厂 名称，那么其中的内容会被视作为 Python 语言内容，并用作函数的参数。这意味着 字符串必须使用双引号包围。

> [run 命令](https://dormousehole.readthedocs.io/en/latest/api.html#flask.cli.run_command)可以启动开发服务器，它在大多数情况下替代 Flask.run() 方法。不要在生产中使用此命令运行应用，只能在开发过程中使用开发服务 器。开发服务器只是为了提供方便，但是不够安全、稳定和高效。有关如何在生 产中运行服务器，请参阅 部署方式 。

> 当 FLASK_ENV 是 development 时会开启调试模式。 如果想要单独控制调试模式，要使用 FLASK_DEBUG 。值为 1 表示开启， 0 表示关闭。

### 部署方式
https://dormousehole.readthedocs.io/en/latest/deploying/index.html


虽然轻便且易于使用，但是 **Flask 的内建服务器不适用于生产** ，它也不能很好
的扩展。本文主要说明在生产环境下正确使用 Flask 的一些方法。

如果想要把 Flask 应用部署到这里没有列出的 WSGI 服务器，请查询其文档中关于
如何使用 WSGI 的部分，只要记住： `Flask` 应用对象实质上是一个 WSGI
应用。


#### 托管于其它web服务

*   [Heroku](https://devcenter.heroku.com/articles/getting-started-with-python)

*   [Google App Engine](https://cloud.google.com/appengine/docs/standard/python/getting-started/python-standard-env)

*   [AWS Elastic Beanstalk](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/create-deploy-python-flask.html)

*   [Azure (IIS)](https://docs.microsoft.com/en-us/azure/app-service/containers/how-to-configure-python)

*   [PythonAnywhere](https://help.pythonanywhere.com/pages/Flask/)


#### 自主部署选项


*   [独立 WSGI 容器](https://dormousehole.readthedocs.io/en/latest/deploying/wsgi-standalone.html)
    *   [Gunicorn](https://dormousehole.readthedocs.io/en/latest/deploying/wsgi-standalone.html#gunicorn)
    *   [uWSGI](https://dormousehole.readthedocs.io/en/latest/deploying/wsgi-standalone.html#uwsgi)
    *   [Gevent](https://dormousehole.readthedocs.io/en/latest/deploying/wsgi-standalone.html#gevent)
    *   [Twisted Web](https://dormousehole.readthedocs.io/en/latest/deploying/wsgi-standalone.html#twisted-web)
    *   [代理设置](https://dormousehole.readthedocs.io/en/latest/deploying/wsgi-standalone.html#deploying-proxy-setups)

*   [uWSGI](https://dormousehole.readthedocs.io/en/latest/deploying/uwsgi.html)
    *   [使用 uwsgi 启动你的应用](https://dormousehole.readthedocs.io/en/latest/deploying/uwsgi.html#id1)
    *   [配置 nginx](https://dormousehole.readthedocs.io/en/latest/deploying/uwsgi.html#nginx)

*   [mod_wsgi (Apache)](https://dormousehole.readthedocs.io/en/latest/deploying/mod_wsgi.html)
    *   [安装 <cite>mod_wsgi</cite>](https://dormousehole.readthedocs.io/en/latest/deploying/mod_wsgi.html#mod-wsgi)
    *   [创建一个 <cite>.wsgi</cite> 文件](https://dormousehole.readthedocs.io/en/latest/deploying/mod_wsgi.html#wsgi)
    *   [配置 Apache](https://dormousehole.readthedocs.io/en/latest/deploying/mod_wsgi.html#id1)
    *   [故障排除](https://dormousehole.readthedocs.io/en/latest/deploying/mod_wsgi.html#id5)
    *   [支持自动重载](https://dormousehole.readthedocs.io/en/latest/deploying/mod_wsgi.html#id6)
    *   [使用虚拟环境](https://dormousehole.readthedocs.io/en/latest/deploying/mod_wsgi.html#id7)

*   [FastCGI](https://dormousehole.readthedocs.io/en/latest/deploying/fastcgi.html)
    *   [创建一个 <cite>.fcgi</cite> 文件](https://dormousehole.readthedocs.io/en/latest/deploying/fastcgi.html#fcgi)
    *   [配置 Apache](https://dormousehole.readthedocs.io/en/latest/deploying/fastcgi.html#apache)
    *   [配置 lighttpd](https://dormousehole.readthedocs.io/en/latest/deploying/fastcgi.html#lighttpd)
    *   [配置 nginx](https://dormousehole.readthedocs.io/en/latest/deploying/fastcgi.html#nginx)
    *   [运行 FastCGI 进程](https://dormousehole.readthedocs.io/en/latest/deploying/fastcgi.html#id1)
    *   [调试](https://dormousehole.readthedocs.io/en/latest/deploying/fastcgi.html#id2)

*   [CGI](cgi.html)
    *   [创建一个 <cite>.cgi</cite> 文件](https://dormousehole.readthedocs.io/en/latest/deploying/cgi.html#id1)
    *   [服务器设置](https://dormousehole.readthedocs.io/en/latest/deploying/cgi.html#id2)

### 参考

[Flask项目Docker容器化部署原理与实现](https://www.cnblogs.com/ybjourney/p/12014120.html)

[最简单的运行](https://github.com/lvthillo/python-flask-docker)

[flask run python:3.8.2-alpine3.11](https://github.com/codefresh-contrib/python-flask-sample-app)

[Flask + uWSGI alpine](https://github.com/cirolini/Docker-Flask-uWSGI)

[supervisord + nginx + uwsgi + flask + alpine](https://github.com/hellt/nginx-uwsgi-flask-alpine-docker/tree/master/python3)