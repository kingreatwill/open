## Flask docker部署

先定义好uwsgi配置文件
uwsgi.ini
```
socket=/tmp/app.sock
chmod-socket=666
pidfile=/etc/nginx/uwsgi.pid
chdir=/learn_flask/app
master=true
wsgi-file=serve.py
http=127.0.0.1:8005
callable=app
processes=8
threads=4
lazy-apps=true
```
启动uwsgi服务器
uwsgi --ini uwsgi.ini

https://hub.docker.com/_/python
可以使用镜像python:3.8.2-alpine3.11  
or python:3.8-alpine3.12 # 默认最新python3.8
or python:3.8-alpine # 默认最新python3.8和最新alpine
### 参考

[Flask项目Docker容器化部署原理与实现](https://www.cnblogs.com/ybjourney/p/12014120.html)

[最简单的运行](https://github.com/lvthillo/python-flask-docker)

[flask run python:3.8.2-alpine3.11](https://github.com/codefresh-contrib/python-flask-sample-app)

[Flask + uWSGI alpine](https://github.com/cirolini/Docker-Flask-uWSGI)

[supervisord + nginx + uwsgi + flask + alpine](https://github.com/hellt/nginx-uwsgi-flask-alpine-docker/tree/master/python3)