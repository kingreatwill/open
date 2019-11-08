[gitea是gogs的clone版本](https://blog.gitea.io/2016/12/welcome-to-gitea/)
选择gitea的原因:
1. 更新更快
2. 参与的人员更多更活跃
3. 更开放
4. 中文文档


# Windows安装
https://docs.gitea.io/zh-cn/windows-service/
https://docs.gitea.io/en-us/windows-service/
## 一、下载
选择合适的版本[下载](https://dl.gitea.io/gitea/)
[Github Release](https://github.com/go-gitea/gitea/releases)

## 二、安装并开启服务

- 下载后放入E:\git\gitea文件夹
- 将exe文件重命名为gitea.exe
- 双击运行，会在当前目录下生成custom\conf\app.ini 文件
- 查看本地系统用户：echo %COMPUTERNAME%  输出：DESKTOP-PK520IC。配置app.ini :RUN_USER = DESKTOP-PK520IC$
- use sqlite3:app.ini
```
[database]
PATH     = E:/git/gitea/data/gitea.db
```
- 注册为windows服务，cmd以管理员身份运行
```
sc create gitea start= auto binPath= ""E:\git\gitea\gitea.exe" web --config "E:\git\gitea\custom\conf\app.ini""
```
- 打开windows服务，开启gitea: services.msc   (sc delete gitea 删除)
- open http://127.0.0.1:3000/



## 三、配置
不开启ssh的配置文件
```
APP_NAME = 零成本Git私服
RUN_USER = DESKTOP-PK520IC$
RUN_MODE = prod

[database]
DB_TYPE  = sqlite3
PATH     = E:/git/gitea/data/gitea.db
HOST     = 127.0.0.1:3306
NAME     = gitea
USER     = gitea
PASSWD   = 
SSL_MODE = disable
CHARSET  = utf8

[i18n]
LANGS = zh-CN,en-US
NAMES = 简体中文,English

; 时间选择器
[i18n.datelang]
zh-CN = zh
en-US = en

[repository]
ROOT = E:/git/gitea/repositories

[server]
PROTOCOL         = http
DOMAIN           = localhost.lingcb.com
ROOT_URL         = http://localhost.lingcb.com:3000/
HTTP_ADDR        = 0.0.0.0
HTTP_PORT        = 3000
START_SSH_SERVER = false
SSH_DOMAIN       = localhost.lingcb.com
DISABLE_SSH      = false
SSH_PORT         = 22
LFS_START_SERVER = true
LFS_CONTENT_PATH = E:/git/gitea/data/lfs
LFS_JWT_SECRET   = Gg4zN_iu_6XqBhJCSQsjUlz7c22bP6Qww6djLeaoy-8
OFFLINE_MODE     = false

[service]
; 禁用注册，启用后只能用管理员添加用户。
DISABLE_REGISTRATION              = true
REGISTER_EMAIL_CONFIRM            = false
ENABLE_NOTIFY_MAIL                = false
ALLOW_ONLY_EXTERNAL_REGISTRATION  = false
ENABLE_CAPTCHA                    = false
REQUIRE_SIGNIN_VIEW               = false
DEFAULT_KEEP_EMAIL_PRIVATE        = false
DEFAULT_ALLOW_CREATE_ORGANIZATION = true
DEFAULT_ENABLE_TIMETRACKING       = true
NO_REPLY_ADDRESS                  = noreply.example.org

[oauth2]
JWT_SECRET = kSoAfgGTda_oAmEY6u2rOETA8qYHO-I9paaOWmrA0kA

[security]
INTERNAL_TOKEN = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1NzMxMTE3Nzl9.__bSRCuneHqXlM2yS66AiG6Sn9tZGskgg7Xd5N8t6sE
INSTALL_LOCK   = true
SECRET_KEY     = xTLboUVEvRuT1XAmFZk60OhsKtPYwjsFAGA2Hho615XWRd5hOs8xexuSTCI9KKKp

[mailer]
ENABLED = false

[picture]
DISABLE_GRAVATAR        = false
ENABLE_FEDERATED_AVATAR = false

[openid]
ENABLE_OPENID_SIGNIN = false
ENABLE_OPENID_SIGNUP = false

[session]
PROVIDER = file

[log]
MODE      = file
LEVEL     = info
ROOT_PATH = E:/git/gitea/log


```
https://docs.gitea.io/zh-cn/config-cheat-sheet/
https://docs.gitea.io/en-us/config-cheat-sheet/


## 创建仓库

1. 创建组织
2. 创建用户
3. 创建仓库
4. 创建团队  给团队添加用户和仓库

个人的仓库是没有团队的

如1：
1. 一个公司可以创建一个组织
2. 一个组织创建多个仓库
3. 一个组织创建多个团队
4. 每个团队负责不同的仓库
5. 任意用户也可以加入任意仓库

如2：
1. 一个公司可以创建多个组织
2. 一个组织创建多个仓库
3. 一个组织创建一个团队
4. 一个团队负责所有的仓库
5. 任意用户也可以加入任意仓库




## http导入
#### 先安装git
```
$ git config --global user.name “XXXX”
$ git config --global user.email XX@XX.com
```
http 每次push或者pull的时候都要求输入账号和密码
```
git config --global credential.helper store
```

#### 从命令行创建一个新的仓库
git init
git add -A .
git commit -m "init"
git remote add origin http://localhost.lingcb.com:3000/erp/Kernel4g-SVN.git
git push -u origin master
#### 从命令行推送已经创建的仓库
git remote add origin http://localhost.lingcb.com:3000/erp/Kernel4g-SVN.git
git push -u origin master

#### 输入用户名和密码

## ssh导入
https://www.jianshu.com/p/acd5fc63895d
https://blog.csdn.net/qq_39240512/article/details/94858115
Git支持多种协议，默认的git://使用ssh，但也可以使用https等其他协议
- 防火墙要开启22端口
1. 生产密钥对
```
ssh-keygen -t rsa -C "your_email@youremail.com"
```


## 开发模式

1. SVN式开发，在当前分支下开发，每个人都有写的权限，都可以直接push
2. 

1. Fork this repo.
2. Create your feature branch (git checkout -b my-new-feature).
3. Commit your changes (git commit -am 'Add some feature').
4. Push to the branch (git push origin my-new-feature).
5. Create new Pull Request.

## 处理合并冲突
- git status  查看状态

- git merge –-abort 合并操作就会被安全的撤销

- git reset –-hard   系统就会回滚到那个合并开始前的状态

以上三个命令基本上能解决一般的冲突了





