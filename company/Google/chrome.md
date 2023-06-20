<!--toc-->
[TOC]
## chrome历史版本下载
https://downzen.com/en/windows/google-chrome/versions/

## 常用插件下载
https://chrome.zzzmh.cn/info/gppongmhjkpfnbhagpmjfkannfbllamg

> 谷歌上网助手、igg谷歌助手

## Chrome插件开发
https://github.com/sxei/chrome-plugin-demo
## Chrome devtools

### 命令（Command） 菜单
Ctrl + Shift + P(mac->Cmd + Shift + P)打开“命令”菜单。

### 截图DOM元素
直接选中那个节点，打开 命令（Command） 菜单并且使用 节点截图(screenshot) 就可以了。
截取特定节点对应上图命令是Screenshot Capture node screenshot。

**全屏截图**：通过 Screenshot Capture full size screenshot 命令。

### 在控制台中使用上次操作的值
使用`$_`可以引用在控制台执行的前一步操作的返回值。如果您正在控制台调试一些JavaScript代码，并且需要引用先前的返回值，那么这可能非常方便。

### 重新发起xhr请求
我们可以通过google提供的Replay XHR的方式去发起一条新的请求，这样对于我们开发效率的提升是有所帮助的。(右键请求-》Replay XHR)

### 编辑页面上的任何文本
在控制台输入document.body.contentEditable="true"或者document.designMode = 'on'就可以实现对网页的编辑了。
其实这个还是比较实用的，比如你要测试一个DOM节点文字太长时，样式是否会混乱，或者要去直接修改页面元素去满足一些业务需求时。（我之前是在Elements面板一个一个去修改的）

### 网络面板（Network）的幻灯片模式
启动Network 面板下的Capture screenshots就可以在页面加载时捕捉屏幕截图。有点幻灯片的感觉。
单击每一帧截图，显示的就是对应时刻发生的网络请求。这种可视化的展现形式会让你更加清楚每一时刻发生的网络请求情况。

### 网络面板（Network）记录之前的请求
启动Network 面板下的Preserve log，之前页面跳转后之前的请求就不见了，很不方便。

### 动画检查
DevTools 中有一个动画面板，默认情况下它是关闭的，很多人可能不太清楚这个功能。它可以让你控制和操纵 CSS 动画，并且可视化这些动画是如何工作的。

要打开该面板，可以在 DevTools 右上角菜单 → More tools 中打开 Animations。

默认情况下，DevTools 会“监听”动画。一旦触发，它们将被添加到列表中。你能看到这些动画块如何显示。在动画本身上，DevTools 会向我们展示哪些属性正在更改，例如 background-color 或 transform。

然后，我们可以通过使用鼠标拖动或调整时间轴来修改该动画。

### 递增/递减 CSS 属性值
作为前端开发，平时少不了通过Elements面板去查找元素以及它的css样式。有时调整像素px会比较麻烦一点，这时就可以使用快捷键去帮你完成：
```
* 增量0.1
  * Mac：Option +向上和Option +向下
  * Windows：Alt +向上和Alt +向下
* 增量1
  * Mac：向上+向下
  * Windows：向上+向下
* 增量10
  * Mac：⇧+向上和⇧+向下
  * Windows：⇧+向上和⇧+向下
* 递增100
  * Mac：⌘+向上和⌘+向下
  * Windows：Ctrl +向上和Ctrl +向下
```

### 在低端设备和弱网情况下进行测试
我们平时开发一般都是在办公室（wifi 网速加快），而且设备一般都是市面上较新的。但是产品的研发和推广，一定要考虑低设备人群和弱网的情况。

在Chrome DevTools中可以轻松调节CPU功能和网络速度。这样，我们就可以测试 Web 应用程序性能并进行相应优化。

具体打开方式是：在Chrome DevTools中通过CMD/Ctrl + Shift + p打开命令菜单。然后输入Show Performance打开性能面板。

### copying & saving
在调试的过程中，我们总会有对 Dev Tools 里面的数据进行 复制 或者 保存 的操作，其实他们也是有一些小技巧的！

- copy()  没有试验通过， 不明白
可以通过全局的方法 copy() 在 console 里 copy 任何你能拿到的资源

- Store as global variable

如果在console中打印了一堆数据，想对这堆数据做额外的操作，可以将它存储为一个全局变量。只需要右击它，并选择 “Store as global variable”选项。第一次使用的话，它会创建一个名为 temp1 的变量，第二次创建 temp2，第三次 ... 。通过使用这些变量来操作对应的数据，不用再担心影响到他们原来的值。

### 自定义 devtools

大家平时用的最多的Chrome 主题可能就是白色/黑色这两种了，但用的久了，难免想尝试像IDE一样切换主题。
首先需要启用实验模式中的Allow custom UI themes
地址栏输入如下url`chrome://flags/#enable-devtools-experiments` # 启用实验功能并重启浏览器

控制台中使用快捷键F1打开设置，切换到Experiments 选项,启用Allow custom UI themes

从Chrome商店安装Material DevTools Theme Collection扩展程序,选择你喜欢的主题即可

### CSS/JS 覆盖率
Chrome DevTools 中的Coverage功能可以帮助我们查看代码的覆盖率。

打开调试面板，用快捷键 ctrl+shift+P 输入 Show Coverage调出相应面板,点击reload 按钮开始检测,
点击相应文件即可查看具体的覆盖情况（绿色的为用到的代码，红色表示没有用到的代码）



### 自定义代码片段 Snippets 
在平常开发过程中，我们经常有些 JavaScript 的代码想在 Chrome Devtools中调试，直接在 console 下 写比较麻烦，或者我们经常有些代码片段(防抖、节流、获取地址栏参数等)想保存起来，每次打开 Devtools 都能获取到这些代码片段，而不用再去google，正好Chrome Devtool 就提供了这种功能。

在 Sources 这个tab栏下，有个 Snippets 标签，在里面可以添加一些常用的代码片段。


### 将图片复制为数据 URI

选择Network面板
在资源面板中选择Img
右键单击将其复制为数据URI（已编码为base 64）


### 媒体查询
什么是媒体查询（不懂）：https://www.cnblogs.com/Kuoblog/p/12442780.html
媒体查询是自适应网页设计的基本部分。在Chrome Devtools中的设备模式下(左边的手机按钮)，在三圆点菜单中点击 Show Media queries即可启用（最又上的三点按钮）。
Devtools会在样式表中检测媒体查询，并在顶端标尺中将它们显示为彩色条形。
那怎么使用呢？其实也很简单：

- 点击媒体查询条形，调整视口大小和预览适合目标屏幕大小的样式
- 右键点击某个条形，查看媒体查询在 CSS 中何处定义并跳到源代码中的定义


### keys/values
这个是Devtools提供的快速查看一个对象的key、values的API。用起来也很简单：
`keys(obj)`  or `values(obj)`  (`obj.keys()` `obj.values()`也可以)


### table
Devtools提供的用于将对象数组记录为表格的API:

table(listjson)

## chrome地址栏命令
在Chrome的浏览器地址栏中输入以下命令，就会返回相应的结果。这些命令包括查看内存状态，浏览器状态，网络状态，DNS服务器状态，插件缓存等等。但是需要注意的是这些命令会不停的变动，所以不一定都是好用的。
about:version - 显示当前版本
about:memory - 显示本机浏览器内存使用状况
about:plugins - 显示已安装插件
about:histograms - 显示历史记录
about:dns - 显示DNS状态
about:cache - 显示缓存页面
about:gpu -是否有硬件加速
about:flags -开启一些插件 //使用后弹出这么些东西：“请小心，这些实验可能有风险”，不知会不会搞乱俺的配置啊！
chrome://extensions/ - 查看已经安装的扩展

## chrome插件
https://github.com/zhaoolee/ChromeAppHeroes


谷歌访问助手破解版
FVD Downloader
Adblock Plus - 免费的广告拦截器
CaretTab - 新式可以显示时间和日期的标签
Chrome版Todoist
ColorPick Eyedropper
ElasticSearch Head
Evernote Web Clipper
在线记事本
Gitpod Online IDE
Enhanced GitHub
Infinity 新标签页
Markdown Editor 
Markdown Preview Plus
Markdown Viewer 推荐
NordVPN 
Postman Interceptor
Project Naptha
Proxy SwitchyOmega
Proxy SwitchySharp
Sweet二维码生成器
Search by Image (by Google)
Tab Resize - split screen layouts
Volume Controller - 音量控制器
划词翻译
沙拉查词-聚合词典划词翻译
彩云小译
### Wappalyzer
网站技术分析插件
https://www.wappalyzer.com

其它类似工具
- WhatRuns
- PageXray
- BuiltWith Technology Profiler https://builtwith.com

### github插件
Octotree - GitHub code tree推荐
GitHub加速 推荐
GitHub Downloader 可以针对单个文件进行下载
Awesome Autocomplete for GitHub 为 GitHub 的搜索栏添加即时搜索功能，简单而谨慎的扩展增强了GitHub的搜索功能，比以往更快地、准确的搜索存储库和人员。
Sourcegraph 对 GitHub 搜索功能，代码比对，查看引用，自动跳转，项目目录导航等功能做了全面的升级。 推荐
File Icons for GitHub and GitLab
GitHub Isometric Contributions 可以将 GitHub 贡献图和等距像素艺术版本之间切换，3D 像素立体展示 GitHub 上的仓库提交记录
Markdown Menu for GitHub 

### tampermonkey插件 
脚本管理器 - 相当于万能插件

### 油猴插件
先安装[脚本管理器- tampermonkey](https://gitee.com/yangbuyi/bky_yby/blob/master/博客园文章等资料/Tampermonkey_4.9.crx)插件，然后安装脚本
https://greasyfork.org/zh-CN/
https://greasyfork.org/zh-CN/scripts

[哔哩哔哩番剧解锁大会员,集合了优酷、爱奇艺、腾讯、芒果、乐视、AB站等全网VIP视频免费破解去广告,高清普清电视观看，持续更新](https://greasyfork.org/zh-CN/scripts/407847-%E5%93%94%E5%93%A9%E5%93%94%E5%93%A9%E7%95%AA%E5%89%A7%E8%A7%A3%E9%94%81%E5%A4%A7%E4%BC%9A%E5%91%98-%E9%9B%86%E5%90%88%E4%BA%86%E4%BC%98%E9%85%B7-%E7%88%B1%E5%A5%87%E8%89%BA-%E8%85%BE%E8%AE%AF-%E8%8A%92%E6%9E%9C-%E4%B9%90%E8%A7%86-ab%E7%AB%99%E7%AD%89%E5%85%A8%E7%BD%91vip%E8%A7%86%E9%A2%91%E5%85%8D%E8%B4%B9%E7%A0%B4%E8%A7%A3%E5%8E%BB%E5%B9%BF%E5%91%8A-%E9%AB%98%E6%B8%85%E6%99%AE%E6%B8%85%E7%94%B5%E8%A7%86%E8%A7%82%E7%9C%8B-%E6%8C%81%E7%BB%AD%E6%9B%B4%E6%96%B0)

[懒人专用，全网VIP视频免费破解去广告、全网音乐直接下载、百度网盘直接下载、知乎视频下载等多合一版。长期更新，放心使用。](https://greasyfork.org/zh-CN/scripts/370634-%E6%87%92%E4%BA%BA%E4%B8%93%E7%94%A8-%E5%85%A8%E7%BD%91vip%E8%A7%86%E9%A2%91%E5%85%8D%E8%B4%B9%E7%A0%B4%E8%A7%A3%E5%8E%BB%E5%B9%BF%E5%91%8A-%E5%85%A8%E7%BD%91%E9%9F%B3%E4%B9%90%E7%9B%B4%E6%8E%A5%E4%B8%8B%E8%BD%BD-%E7%99%BE%E5%BA%A6%E7%BD%91%E7%9B%98%E7%9B%B4%E6%8E%A5%E4%B8%8B%E8%BD%BD-%E7%9F%A5%E4%B9%8E%E8%A7%86%E9%A2%91%E4%B8%8B%E8%BD%BD%E7%AD%89%E5%A4%9A%E5%90%88%E4%B8%80%E7%89%88-%E9%95%BF%E6%9C%9F%E6%9B%B4%E6%96%B0-%E6%94%BE%E5%BF%83%E4%BD%BF%E7%94%A8)

[百度文库转 Word | 百度文库下载器](https://greasyfork.org/zh-CN/scripts/405373-%E7%99%BE%E5%BA%A6%E6%96%87%E5%BA%93%E8%BD%AC-word-%E7%99%BE%E5%BA%A6%E6%96%87%E5%BA%93%E4%B8%8B%E8%BD%BD%E5%99%A8)

[Github 增强 - 高速下载](https://greasyfork.org/zh-CN/scripts/412245-github-%E5%A2%9E%E5%BC%BA-%E9%AB%98%E9%80%9F%E4%B8%8B%E8%BD%BD)

[github、码云 md文件目录化](https://greasyfork.org/zh-CN/scripts/387834-github-%E7%A0%81%E4%BA%91-md%E6%96%87%E4%BB%B6%E7%9B%AE%E5%BD%95%E5%8C%96)

[在线下载Github仓库文件夹](https://greasyfork.org/zh-CN/scripts/411834-download-github-repo-sub-folder)

[在Colab中打开GitHub Jupyter](https://greasyfork.org/zh-CN/scripts/408674-open-github-jupyter-in-colab)
[Github助手](https://greasyfork.org/zh-CN/scripts/37899-github%E5%8A%A9%E6%89%8B)
添加Github文件下载、复制按钮、图片点击放大(右击恢复)、issues中只查看用户相关态度的内容、issues列表项从新标签页打开

[Github项目批量删除](https://greasyfork.org/zh-CN/scripts/412188-github%E9%A1%B9%E7%9B%AE%E6%89%B9%E9%87%8F%E5%88%A0%E9%99%A4)

### google stackoverflow
gooreplacer 最初为解决国内无法访问 Google 资源（Ajax、API等）导致页面加载速度巨慢而生，新版在此基础上，增加了更多实用功能，可以方便用户屏蔽某些请求，修改 HTTP 请求/响应 的 headers。
https://github.com/jiacai2050/gooreplacer

将 Google CDN 替换为国内的
https://github.com/justjavac/ReplaceGoogleCDN

FastGithub工具也代理了很多Google 资源 以及stackoverflow（https://github.com/dotnetcore/FastGithub/releases/tag/1.1.4）
https://github.com/dotnetcore/FastGithub


Gravatar，全称Globally Recognized Avatar。翻译成中文为全球通用头像。
[Gravatar头像 镜像地址大全](https://zhuanlan.zhihu.com/p/115248957)

插件： Ghelper(谷歌上网助手)

网站镜像（谷歌搜索，谷歌學術，維基百科，Github）：
https://mirror.yaojiu.us/mirror/
https://www.haoo.us/mirror
https://www.library.ac.cn/

https://github.yaojiu.us/

