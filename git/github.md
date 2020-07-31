# Github
[GitHub 镜像访问](https://code.pingbook.top/blog/2020/How-To-Speed-Github.html)
## GItHub  刷新 ipconfig /flushdns
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

## GitHub 镜像访问
这里提供两个最常用的镜像地址：

https://github.com.cnpmjs.org
https://hub.fastgit.org
也就是说上面的镜像就是一个克隆版的Github，你可以访问上面的镜像网站，网站的内容跟Github是完整同步的镜像，然后在这个网站里面进行下载克隆等操作。

## GitHub文件加速
利用 Cloudflare Workers 对github release 、archive 以及项目文件进行加速，部署无需服务器且自带CDN.

https://gh.api.99988866.xyz
https://g.ioiox.com

以上网站为演示站点，如无法打开可以查看开源项目：[gh-proxy-GitHub](https://hunsh.net/archives/23/) 文件加速自行部署。

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