<!--toc-->
[TOC]
# Github

## github issues
可以用issues来写blog
优点：
- 支持回复
- 支持tag
- 可以计划
- 支持订阅
- 全文检索

缺点：
- 任何人都可以新建issue（有好处也有不好的）

> 感觉很棒



### 思考
#### 直接创建一个网站首页

获取api显示
> https://api.github.com/repos/lzh2nix/articles/issues?page=2
> 在上面的api header中有以下信息
> link: <https://api.github.com/repositories/112858035/issues?page=1>; rel="prev", <https://api.github.com/repositories/112858035/issues?page=3>; rel="next", <https://api.github.com/repositories/112858035/issues?page=3>; rel="last", <https://api.github.com/repositories/112858035/issues?page=1>; rel="first"

> https://developer.github.com/v3/issues/

https://github.com/josegonzalez/python-github-backup
https://github.com/IQAndreas/github-issues-import
https://github.com/devspace/awesome-github-templates

[ABAP开发的Github issue备份工具](https://zhuanlan.zhihu.com/p/206986949)

> gitee 也有对应的api：https://gitee.com/api/v5/swagger#/getV5ReposOwnerRepoIssuesComments

#### 开发一个同步工具
利用issues的api

评论可以单向同步到自己的blog系统

同步部分分离出一个服务出来

#### api参考
https://github.com/search?q=github+api

java
https://gitee.com/openkylin/java-gitee
https://gitee.com/wuyu15255872976/gitee-client

python

### 提高 API 访问次数的配额
默认情况下你是用匿名权限访问 github 接口的， github 的访问限制是一个小时最多 60 次请求，这显然是不够的，如何提高限制呢？

1. 到个人设置下的 Personal access tokens 页（https://github.com/settings/tokens ），如下图，点击右上角的 Generate new token



2. 填写名称，只勾选 public_repo,然后保存，github 会生成一个可访问你公开项目的 access_token，将它填入到配置文件的 access_token 的值中，并取消注释。 

3. 打开 app.js,取消掉第 17 行和 88 行的注释，保存后重新上传即可

 data:{
     // access_token:_config['access_token']
 },




## github profile readme
https://github.com/abhisheknaiidu/awesome-github-profile-readme
https://docs.github.com/en/github/setting-up-and-managing-your-github-profile
https://github.com/anuraghazra/github-readme-stats
https://github.com/DenverCoder1/github-readme-streak-stats
https://github.com/sallar/github-contributions-chart  静态生成所有年的
https://github.com/jasonlong/isometric-contributions
https://github.com/IonicaBizau/github-contributions
https://github.com/IonicaBizau/git-stats
https://github.com/2016rshah/githubchart-api 实时生成今年的

##  徽章工具
https://github.com/badges/shields
[shields](https://shields.io/)

https://github.com/badgen/badgen.net
[badgen](https://badgen.net/)

[forthebadge](https://forthebadge.com/)
[badges.frapsoft](https://github.com/ellerbrock/open-source-badges)
[badge.fury](https://badge.fury.io/)


![docker](https://badgen.net/badge/docker/v1.2.3/blue?icon=docker)
![redis](https://img.shields.io/badge/Redis-5.0+-yellow.svg)
![](https://badges.gitter.im/Join%20Chat.svg)
![](http://badges.github.io/stability-badges/dist/deprecated.svg)
https://blog.csdn.net/u011192270/article/details/51788886

### repo cards
https://github.com/lepture/github-cards

https://gh-card.dev/

## github域名 IP

**github.io** 不能访问：DNS：223.5.5.5  和 备用DNS：223.6.6.6

https://site.ip138.com/raw.githubusercontent.com/

```
raw.githubusercontent.com服务器iP：
当前解析：
日本 东京151.101.108.133

日本 东京151.101.228.133

美国151.101.0.133

保留地址0.0.0.0

中国 香港151.101.76.133
```

找出延时最小的IP
添加hosts

http://tool.chinaz.com/dns/?type=1&host=raw.githubusercontent.com&ip=

找出最快的IP地址
https://www.ipaddress.com/

```
# GItHub  刷新 ipconfig /flushdns
# 52.74.223.119     github.com
# 192.30.253.119    gist.github.com
# 54.169.195.247    api.github.com
# 185.199.111.153   assets-cdn.github.com

151.101.76.133    raw.githubusercontent.com
151.101.76.133    gist.githubusercontent.com
151.101.76.133    cloud.githubusercontent.com
151.101.76.133    camo.githubusercontent.com
151.101.76.133    avatars0.githubusercontent.com
151.101.76.133    avatars1.githubusercontent.com
151.101.76.133    avatars2.githubusercontent.com
151.101.76.133    avatars3.githubusercontent.com
151.101.76.133    avatars4.githubusercontent.com
151.101.76.133    avatars5.githubusercontent.com
151.101.76.133    avatars6.githubusercontent.com
151.101.76.133    avatars7.githubusercontent.com
151.101.76.133    avatars8.githubusercontent.com
```

### GItHub  刷新 ipconfig /flushdns
151.101.112.249 http://global-ssl.fastly.Net
192.30.253.112 http://github.com
151.101.44.249 github.global.ssl.fastly.net
192.30.253.113 github.com
103.245.222.133 assets-cdn.github.com
23.235.47.133 assets-cdn.github.com
203.208.39.104 assets-cdn.github.com
204.232.175.78 documentcloud.github.com
204.232.175.94 gist.github.com
107.21.116.220 help.github.com
207.97.227.252 nodeload.github.com
199.27.76.130 raw.github.com
107.22.3.110 status.github.com
204.232.175.78 training.github.com
207.97.227.243 www.github.com
185.31.16.184 github.global.ssl.fastly.net
185.31.18.133 avatars0.githubusercontent.com
185.31.19.133 avatars1.githubusercontent.com

[GitHub 镜像访问](https://code.pingbook.top/blog/2020/How-To-Speed-Github.html)

### 加速github
https://github.com/dotnetcore/FastGithub

https://github.com/521xueweihan/GitHub520

## GitHub 镜像访问
这里提供两个最常用的镜像地址：
- https://hub.fastgit.org/openjw/open.git
- https://gitclone.com/github.com/openjw/open.git
- https://github.com.cnpmjs.org/openjw/open.git
- https://github-mirror.bugkiller.org/openjw/open.git
- https://github-speedup.laiczhang.com/openjw/open.git  

https://github-speedup.laiczhang.com - https://github.com/zzh-blog/GithubSpeedUp
https://github.com.cnpmjs.org
https://hub.fastgit.org - https://github.com/k0shk0sh/FastHub/
也就是说上面的镜像就是一个克隆版的Github，你可以访问上面的镜像网站，网站的内容跟Github是完整同步的镜像，然后在这个网站里面进行下载克隆等操作。

https://gitclone.com/

https://gh.api.99988866.xyz/ - https://github.com/hunshcn/gh-proxy
http://gitd.cc/

https://github.com/Henry14all/github-plus-js

自定义脚本：clone
```perl
#!/bin/perl

# 参考知乎文章: https://zhuanlan.zhihu.com/p/165413464
# 两个   github.com  镜像网址来实现 clone 加速
#   1.   github.com.cnpmjs.org
#   2.   https://hub.fastgit.org


use strict;
use warnings;

sub git_clone(@){
    my $origin_url = $_[0];
    my $url = $origin_url =~ s/github\.com/hub\.fastgit\.org/r;
    printf "%s\n",$url;
    system("git clone $url");
    my @arr = split(/\/+/, $origin_url =~ s/(\.git)$//r);
    my $dir = $arr[-1];
    chdir $dir;
    system("git remote set-url origin $origin_url");
    system("git fetch")
}

sub print_usage() {
    print "two mthods to clone github repo\n";
    print "  1. perl clone <repo>\n";
    print "  2. ./clone <repo>\n";
}

my $argc = @ARGV;
if($argc == 0){
    print_usage();
    exit(1);
}
my $repo = $ARGV[0];
if($repo =~ m/^git@/){
    print "only support https url now, please try the follow command:\n\n";
    $repo = $repo =~ s/:/\//r;
    $repo = $repo =~ s/^git@/https:\/\//r;
    print "    perl clone $repo \n\n";
    print "or\n\n";
    print "    clone $repo \n\n";
    exit(1);
}
if($ARGV[0] =~ m/(-h)|(--help)|(^help$)/){
    print_usage();
    exit(0);
}

git_clone($ARGV[0]);
```
use: clone https://github.com/google/pprof.git


## GitHub文件加速
利用 Cloudflare Workers 对github release 、archive 以及项目文件进行加速，部署无需服务器且自带CDN.

https://gh.api.99988866.xyz
https://g.ioiox.com

以上网站为演示站点，如无法打开可以查看开源项目：[gh-proxy-GitHub](https://hunsh.net/archives/23/) 文件加速自行部署。

## Github 加速下载
只需要复制当前 GitHub 地址粘贴到输入框中就可以代理加速下载！

地址：http://toolwa.com/github/

## Github 加速下载
只需要复制当前 GitHub 地址粘贴到输入框中就可以代理加速下载！

地址：http://toolwa.com/github/

## 加速你的 Github
https://github.zhlh6.cn

输入 Github 仓库地址，使用生成的地址进行 git ssh 操作即可

## 谷歌浏览器GitHub加速插件(推荐)
[谷歌浏览器Github加速插件.crx](https://chrome.google.com/webstore/detail/github%E5%8A%A0%E9%80%9F/mfnkflidjnladnkldfonnaicljppahpg/related?hl=zh-CN) 下载

百度网盘: https://pan.baidu.com/s/1qGiIUzqNlN1ZczTNFbPg0A,提取码：stsv

如果可以直接访问谷歌商店，可以访问[GitHub 加速谷歌](https://chrome.google.com/webstore/detail/github%E5%8A%A0%E9%80%9F/mfnkflidjnladnkldfonnaicljppahpg)商店安装。

## GitHub raw 加速
GitHub raw 域名并非 github.com 而是 raw.githubusercontent.com，上方的 GitHub 加速如果不能加速这个域名，那么可以使用 Static CDN 提供的反代服务。

将 raw.githubusercontent.com 替换为 raw.staticdn.net 即可加速
## GitHub + Jsdelivr
jsdelivr 唯一美中不足的就是它不能获取 exe 文件以及 Release 处附加的 exe 和 dmg 文件。

也就是说如果 exe 文件是附加在 Release 处但是没有在 code 里面的话是无法获取的。所以只能当作静态文件 cdn 用途，而不能作为 Release 加速下载的用途。

## github ssh提交
1. 查看git提交方式
`git remote -v`

2. 生成公钥
默认公钥是存储在用户目录下的.ssh目录中，如下：
```
C:\Users\xxx\.ssh   # Windows
/home/xxx/.ssh  # Linux
```
如果没有公钥的文件，输入如下命令生成公钥
`ssh-keygen -t rsa -C "kingreatwill@qq.com"`
密码可填可不填，填的话需要大于5位，不能太简单，一般存储普通项目直接回车跳过即可。

用记事本打开id_rsa.pub，复制里面所有内容，进入github个人settings里。找到SSH and GPG keys，这里保存了所有与你github关联的公钥。

3. 移除https方式 & 添加ssh提交方式
```
git remote rm origin
git remote add origin git@github.com:openjw/open.git
```
4. 验证 & 正常
`git remote -v`
第一次提交如果失败，使用下面cmd
`git push --set-upstream orgin master`

5. tortoisegit 错误
disconnected no supported authentication methods available(server sent: publickey)
右键--小乌龟---settings--network--修改ssh client为git的ssh.exe
`D:\Program Files\Git\usr\bin\ssh.exe`

6. 增加gitee同步更新
在gitee配置ssh公钥
`git remote set-url --add origin git@gitee.com:kingreatwill/open.git`
可以参考[如何同步多个 git 远程仓库](https://www.cnblogs.com/taadis/p/12170953.html)
由于gitee限制单个文件50M，单个仓库1G(企业500M)，总仓库5G，所以可以考虑coding

> 如果需要gitee的贡献度的统计，那么你的本地git填写的邮箱（区别密钥填写的邮箱）是gitee设置的提交邮箱.
> github也是一样

7. 增加coding同步更新
总容量100G, 单个仓库2G
新增openjw团队，创建kingreatwill项目，在该项目下创建仓库
`git remote set-url --add origin git@e.coding.net:openjw/kingreatwill/open.git`

> codechina
> `git remote set-url --add origin git@codechina.csdn.net:kingreatwill/open.git`

8. cnb.cool
`git remote set-url --add origin https://cnb.cool/wcoder/open.git`

9. 其它
百度效率云：http://xiaolvyun.baidu.com

阿里云Code托管平台：https://code.aliyun.com
老版以前代码仓库容量限制是2G，数量是50
**现在不限制了！！**

云效DevOps 中的代码管理Codeup

新版：https://codeup.aliyun.com/
`git remote set-url --add origin git@codeup.aliyun.com:kingreatwill/kingreatwill/open.git`

企业设置-企业个性化 - kingreatwill

第二个kingreatwill是代码组

```
使用带用户名密码的方式（可以避免后续每次都要输入用户名密码）

git clone https://[username]:[password]@/remote

但有时会出现用户名或密码中含有像@这样的特殊符号，而不能被正常解析

我们需要通过下面方式进行重新编码

String c = URLEncoder.encode("@","utf-8");
System.out.println(c);

console -> %40

-------------------------
[git 指定要提交的ssh key](https://www.cnblogs.com/chenkeyu/p/10440798.html)
环境变量GIT_SSH_COMMAND：

注意：使用git的cmd-> GIT_SSH_COMMAND="ssh -i xx" git push
从Git版本2.3.0可以使用环境变量GIT_SSH_COMMAND，如下所示：

GIT_SSH_COMMAND="ssh -i ~/.ssh/id_rsa_example" git clone example
请注意，-i有时可以被您的配置文件覆盖，在这种情况下，您应该给SSH一个空配置文件，如下所示：

GIT_SSH_COMMAND="ssh -i ~/.ssh/id_rsa_example -F /dev/null" git clone example
```

## 使用技巧

### blame 责任
`#L10-L13` 行数
https://github.com/kingreatwill/kingreatwill/blame/master/README.md#L10-L13
行号可以点击，按住shift可以选择多行

### trending排行榜
https://github.com/trending/java?since=daily

### gitattributes设置项目语言
```
*.html linguist-language=JavaScript
```
主要意思是把所有html文件后缀的代码识别成js文件。

https://github.com/alexkaratarakis/gitattributes
https://git-scm.com/docs/gitattributes


### 通过提交的msg自动关闭issues

比如有人提交了个issues:

https://github.com/AlloyTeam/AlloyTouch/issues/6

然后你去主干上改代码，改完之后提交填msg的时候，填入：

fix https://github.com/AlloyTeam/AlloyTouch/issues/6

这个issues会自动被关闭。当然不仅仅是fix这个关键字。下面这些关键字也可以：

close
closes
closed
fixes
fixed
resolve
resolves
resolved

### 搜索
https://docs.github.com/en/github/searching-for-information-on-github/searching-on-github

### in限制搜索范围

使用方法：搜索的关键词 in: name / description / readme -> 项目名/项目描述/readme文件

可以组合使用

公式：搜索词 in:name(/description/readme)

搜索项目名称和自述文件中包含秒杀的仓库   seckill in:name,readme

限定符 |	示例
---|---
in:name	|jquery in:name 匹配仓库名称包含 "jquery" 的内容
in:description	|jquery in:name,description 匹配仓库名或描述中包含 "jquery" 的内容(组合使用)
in:readme	|jquery in:readme readme文件中包含"jquery"
repo:owner/name	|repo:octocat/hello-world 查询某人的某个项目（查octocat 的 hello-world 仓库）


### 关键词查找

使用方法：搜索的关键词 stars / forks 通配符

通配符选择`:>`或者`:>=`

或者使用区间范围：下限..上限

star和fork可以组合是使用

示例：`spring boot forks 100..200 stars 1000..2000`

```
>n、>=n、<n、<=n：查询数量范围，可以是 starts、forks、topics......

n..*、*..n：等同于 >=n 和 <=n

n..m：取值范围 n 到 m
```
限定符|	示例
---|---
stars:n	|	stars:500 匹配 500 个 stars 的项目<br/>stars:10..20 匹配 starts 数量 10 到 20 的项目
followers:n	|	node followers:>=10000 匹配关注者大于等于 10000 的 node 仓库
forks:n	|	seckill forks:5  匹配有 5 个 forks 的秒杀项目
created:YYYY-MM-DD|		seckill created:>2020-01-01 创建时间在 2020-01-01 之后的秒杀项目
language:LANGUAGE	|	seckill language:java 匹配 java 语言编写的秒杀项目
user:name	|	user:Jstarfish stars:>50 匹配 Jstarfish 用户 stars 数大于 50 的仓库
location:LOCATION	|	location:beijing 匹配北京的朋友们
互相组合使用|		seckill stars:>=500 fork:true language:java 匹配stars 数量大等于 500（包含 forks 数），且语言是 java 的秒杀项目<br/>location:beijing language:java 北京做 Java 开发的大佬


### 项目内搜索文件

在项目界面按`t`，项目变成列表形式

#### 其它快捷键

https://docs.github.com/en/github/getting-started-with-github/keyboard-shortcuts

https://docs.github.com/en/desktop/getting-started-with-github-desktop/keyboard-shortcuts

### 需求：Github上项目的master分支设置为不能直接合并需要review才能合并
https://www.cnblogs.com/liucx/p/12132709.html

### 其他
issue中输入冒号 : 添加表情
任意界面，shift + ？显示快捷键
issue中选中文字，R键快速引用

使用方式：location:地区

使用方式：language:语言



## github hosts
https://ineo6.github.io/hosts/

https://github.com/maxiaof/github-hosts
https://github.com/ineo6/hosts
https://gitlab.com/ineo6/hosts/-/raw/master/next-hosts
https://gitlab.com/ineo6/hosts/-/raw/master/hosts

可以安装[SwitchHosts](https://github.com/oldj/SwitchHosts/releases), 然后自动更新hosts