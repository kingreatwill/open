gitlab gogs  gitea
[gitea是gogs的clone版本](https://blog.gitea.io/2016/12/welcome-to-gitea/)

[比较](https://www.toutiao.com/a6759755783925137927/)
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
- 打开windows服务，开启gitea: services.msc   (sc delete gitea 删除)(或者net start gitea)
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
DOMAIN           = localhost.xxx.com
ROOT_URL         = http://localhost.xxx.com:3000/
HTTP_ADDR        = 0.0.0.0
HTTP_PORT        = 3000
START_SSH_SERVER = false
SSH_DOMAIN       = localhost.xxx.com
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

## https配置
```
./gitea cert -ca=true -duration=8760h0m0s -host=localhost.xxx.com
```
配置文件:
```
[server]
PROTOCOL         = https
ROOT_URL         = https://localhost.xxx.com/
HTTP_PORT        = 443
CERT_FILE = E:/git/gitea/custom/https/cert.pem
KEY_FILE = E:/git/gitea/custom/https/key.pem
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
git 警告: LF will be replaced by CRLF in readme.txt. The file will have its original line endings in your working directory.
```
git config --global core.autocrlf false  //禁用自动转换
```
- 情况一：
Git 可以在你提交时自动地把回车（CR）和换行（LF）转换成换行（LF），而在检出代码时把换行（LF）转换成回车（CR）和换行（LF）。 你可以用git config --global core.autocrlf true 来打开此项功能。 如果是在 Windows 系统上，把它设置成 true，这样在检出代码时，换行会被转换成回车和换行：
```
#提交时转换为LF，检出时转换为CRLF
$ git config --global core.autocrlf true
```
- 情况二：
如果使用以换行（LF）作为行结束符的 Linux 或 Mac，你不需要 Git 在检出文件时进行自动的转换。然而当一个以回车（CR）和换行（LF）作为行结束符的文件不小心被引入时，你肯定想让 Git 修正。 所以，你可以把 core.autocrlf 设置成 input 来告诉 Git 在提交时把回车和换行转换成换行，检出时不转换：（这样在 Windows 上的检出文件中会保留回车和换行，而在 Mac 和 Linux 上，以及版本库中会保留换行。）
```
#提交时转换为LF，检出时不转换
$ git config --global core.autocrlf input
```
- 情况三：
如果你是 Windows 程序员，且正在开发仅运行在 Windows 上的项目，可以设置 false 取消此功能，把回车保留在版本库中：
```
#提交检出均不转换
$ git config --global core.autocrlf false
```
你也可以在文件提交时进行safecrlf检查
```
#拒绝提交包含混合换行符的文件
git config --global core.safecrlf true   

#允许提交包含混合换行符的文件
git config --global core.safecrlf false   

#提交包含混合换行符的文件时给出警告
git config --global core.safecrlf warn
```
###### 通俗解释
- git 的 Windows 客户端基本都会默认设置 core.autocrlf=true，设置core.autocrlf=true, 只要保持工作区都是纯 CRLF 文件，编辑器用 CRLF 换行，就不会出现警告了；
- Linux 最好不要设置 core.autocrlf，因为这个配置算是为 Windows 平台定制；
- Windows 上设置 core.autocrlf=false，仓库里也没有配置 .gitattributes，很容易引入 CRLF 或者混合换行符（Mixed Line Endings，一个文件里既有 LF 又有CRLF）到版本库，这样就可能产生各种奇怪的问题。

#### 从命令行创建一个新的仓库
git init
git add -A .
git commit -m "init"
git remote add origin http://localhost.xxx.com:3000/erp/Kernel4g-SVN.git
git push -u origin master
#### 从命令行推送已经创建的仓库
git remote add origin http://localhost.xxx.com:3000/erp/Kernel4g-SVN.git
git push -u origin master

#### 输入用户名和密码

#### 创建分支
1. 切换到基础分支，如主干
git checkout master

2. 创建并切换到新分支
git checkout -b panda #或者git branch panda（不会切换到这个分支） 创建一个新的分支
git branch #可以看到已经在panda分支上

3. 更新分支代码并提交
git add *
git commit -m "init panda"
git push origin panda


同步 远程分支master 到 本地分支es7
git branch --set-upstream-to=origin/master es7

#### 远程仓库branch合并到 master
git checkout master  //切换到 master
git merge origin/panda  //选择要合并到 master 的分支
git push origin master   //push 即可 

git merge panda // 合并本地分支panda 到当前分支

#### 删除分支
git branch -d panda 对分支本地分支panda 进行删除
git push origin --delete panda 删除远程分支


#### 更新
git fetch 从远程仓库拉取更新，使用git fetch后，并不会将新的内容更新到工作区域的文件中，所以可以通过git diff master origin/master命令来比较本地与远程master分支的差异：

git merge 命令来将更新合并到工作区域

git pull命令相当于执行了git fetch和git merge两个命令。

#### tag
git tag -a v0.1 -m "测试版本"
git push origin v0.1
//git tag -a 标签名称 -m "说明"
//git push origin 标签名称
git tag -d v1.1  //删除本地tag
git push origin :v1.1//删除远程tag
//也可以这样  
git push origin --delete tag V1.1 


#### Pull Request
当分支完成开发后，需要将代码进行合并，一般是将分支代码合并到远程的如Master或Develop之类的长期分支上，其流程如下：
　　1. 创建一个功能分支feature1(git checkout -b feature1)。
　　2. 在分支上完成功能并提交(git add & git commit)。
　　3. 切换到master分支执行合并操作，并将更新推到远程仓库(git checkout master, git merge feature1, git push)。
　　4. 删除特性分支(git branch -d feature1)。

对于没有权限提交的人或者代码有严格要求的情况下如何合并代码，git引入了pull request
为什么不是push request
目的是让仓库所有者来“拉”取变化，由所有者来决定合并还是拒绝，所有者可以根据功能是否合理、代码是否正确、易读等信息进行判断，这实际上就是CodeRview的过程

Pull Request的要求就是需要两个远程分支(仓库)进行合并(代码拥有者的分支和代码贡献者的分支)：

1. 克隆My Blog代码，创建一个新的远程仓库(本例使用GitHub作为托管平台，可以直接fork)：
　　git clone https://github.com/yqszt/MyBlog.git
　　git remote add other https://github.com/SelimTeam/MyBlog.git
　　git push -u other
2. 在克隆的代码中修改内容并提交：
3. 要将这两次提交生成“pull request”：
　　使用git request-pull命令生成拉请求信息：
　　git request-pull -p 5bf2e35 https://github.com/SelimTeam/MyBlog.git master
其中p代表输出详细内容(代码的差异)，5bf2e35对应的是提交的hash，代表更新的内容是从哪一个提交开始，url代表的是贡献者的仓库地址，最后的master代表更新内容结束的提交，默认是分支的最新提交。

4. 将pull request信息告知作者，作者将会知道贡献者的仓库地址、分支、从哪一个提交开始、哪一个提交结束，并且带有详细的变更信息。
　　注：这里的告知是通过邮件等方式将上面request-pull命令生成的信息发送给作者，github等平台上提供的pull request功能是由平台自己实现的通知方式，关于github上的pull request后续介绍。

5. 作者添加贡献者的远程仓库，获取并将更新合并到主分支：
　　git remote add selimteam https://github.com/SelimTeam/MyBlog.git
　　git fetch selimteam master
　　git diff master selimteam/master
    git merge selimteam/master
    git push

以上就完成了一次通过pull request像作者贡献代码的流程。


界面操作：
1. Fork this repo.
2. git clone
3. Create your feature branch (git checkout -b my-new-feature).
4. Commit your changes (git commit -am 'Add some feature').
5. Push to the branch (git push origin my-new-feature).
6. Create new Pull Request.


## ssh导入
https://www.jianshu.com/p/acd5fc63895d
https://blog.csdn.net/qq_39240512/article/details/94858115
Git支持多种协议，默认的git://使用ssh，但也可以使用https等其他协议
- 防火墙要开启22端口
- 去配置里增加一行：START_SSH_SERVER = true 保存后重启gitea.exe
1. 生产密钥对
1.本地电脑上生成并部署SSH key：

```bash
ssh-keygen -t rsa -C "your_email@youremail.com"
```
创建密码（如果不需要密码则直接回车）；
一路回车即可生成ssh私钥（id_rsa）和公钥（id_rsa.pub）
一般在用户目录/.ssh下

id_rsa.pub公钥给到gitea管理员，管理员将其添加到gitea.

测试：
ssh -t git@github.com #github
ssh -t DESKTOP-PK520IC@localhost.xxx.com # gitea
注意：网页上显示的ssh库地址有个$，需要删除。

## Git常用的GUI工具
1. SourceTree:一个开源的Git GUI工具，有一个重要的点是它提供了对git flow的支持。
2. GitHub For Desktop:GitHub的GUI客户端，可以通过它直接提交pull request(GitHub的PullRequest)。

3. [更多](https://git-scm.com/download/gui/win)
4. tortoisegit
5. GitExtensions
## 处理合并冲突
- git status  查看状态

- git merge –-abort 合并操作就会被安全的撤销

- git reset –-hard   系统就会回滚到那个合并开始前的状态

以上三个命令基本上能解决一般的冲突了

## git 添加外部项目地址
```
cd projectA
git submodule add projectB.git projectB
```
git submodule add https://git.xxx.com/org/test.git common/test
提交

当别人git clone包含子模块的代码仓库
```
git clone projectA.git
cd projectA
git submodule init
git submodule update
```
或者在使用git clone命令时，加上–recurse-submodules或–recursive 这样的递归参数
git clone --recursive projectA.git

当projectB中添加内容时，git status会发现，

```
我们要做的就是在[submodule “projectB”]中添加一个ignore子项，这个ignore子项可以有上个可选的值，untracked, dirty和all, 它们的意思分别是：

untracked ：忽略 在子模块B(也就是projectB目录)新添加的，未受版本控制内容
dirty ： 忽略对projectB目录下受版本控制的内容进行了修改
all ： 同时忽略untracked和dirty
这里我们先选择dirty(至少先保证不提交对子模块B的任何修改)，其他的可以根据具体需求来进行选择。
```

## GO mod
- go get
```
go get -insecure localhost.xxx.com/erp/erp_golang@v0.3.0
```
- 请问如何让 go mod 对某些私有 module 跳过 GOPROXY 代理？
```
go env -w GOPROXY=https://goproxy.io,direct
# Set environment variable allow bypassing the proxy for selected modules
go env -w GOPRIVATE=*.xxx.com
# GONOSUMDB=github.com/mycompany/*
# go env-w GOSUMDB=off
# GOSUMDB="sum.golang.google.cn"
```



