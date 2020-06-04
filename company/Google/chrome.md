<!--toc-->
[TOC]
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
Octotree 推荐
Postman Interceptor
Project Naptha
Proxy SwitchyOmega
Proxy SwitchySharp
Sweet二维码生成器
Search by Image (by Google)
Sourcegraph 推荐
Tab Resize - split screen layouts
Volume Controller - 音量控制器
划词翻译
沙拉查词-聚合词典划词翻译
彩云小译





