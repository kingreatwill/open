https://blog.csdn.net/hj7jay/article/details/83893185
https://gogs.io/docs/installation
https://github.com/gogs/gogs

注册为windows服务，cmd以管理员身份运行
```
sc create gogs start= auto binPath= ""E:\git\gogs\gogs.exe" web --config "E:\git\gogs\custom\conf\app.ini""

net start gogs

sc delete gogs
```