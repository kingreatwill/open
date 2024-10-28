
## awesome 集合
https://github.com/sindresorhus/awesome

https://github.com/GrowingGit/GitHub-Chinese-Top-Charts

## 序列化



https://github.com/apache/thrift

https://github.com/protocolbuffers/protobuf

https://github.com/msgpack/msgpack

https://github.com/google/flatbuffers

https://github.com/capnproto/capnproto 无dotnet

https://avro.apache.org/


https://github.com/niXman/yas

https://github.com/USCiLab/cereal


## 远程/VNC

### VNC
- [TigerVNC](https://github.com/TigerVNC/tigervnc)


```
root@master:~# dnf install  -y  tigervnc-server
root@master:~# vncserver
```

- [TightVNC](https://github.com/chenall/tightvnc)

### mRemoteNG 
https://github.com/mRemoteNG/mRemoteNG

- RDP (Remote Desktop Protocol)
- VNC (Virtual Network Computing)
- SSH (Secure Shell)
- Telnet (TELecommunication NETwork)
- HTTP/HTTPS (Hypertext Transfer Protocol)
- rlogin (Remote Login)
- Raw Socket Connections
- Powershell remoting


## Javascript Engine
JS的标准是ECMAScript。ECMA是一个标准化组织，下面有好多标准。ECMAScript是形成JavaScript语言基础的脚本语言。ECMAScript是由Ecma国际标准组织以ECMA-262和ECMA-402规范的形式进行标准化的。
https://bellard.org/quickjs/bench.html
引擎            | 描述
---|---
V8 (JIT)        | https://v8.dev， 应该是目前地球上最高技术水准的JS引擎了。
V8 --jitless    |
Hermes          | https://hermesengine.dev, Facebook搞出来的，为了Android平台上的React Native而生
QuickJS	        | 小而精，几乎完整实现ES2019支持
DukTape	        |https://duktape.org，小巧，内存占用小，可方便集成到其它模块中。
MuJS            | https://mujs.com，和DukTape类似，轻量级的，可方便集成到其它模块中的JS解释器。
XS              | https://www.moddable.com，用于IoT嵌入式低端设备的JS引擎。官网明确说了，用于微控制器（microcontroller）。
JerryScript     | https://jerryscript.net，也是用于IoT设备的JS引擎


浏览器渲染引擎 | .
---|---
SquirrelFish Extreme |(WebKit , Safari浏览器) Safari开发了一个强劲的JavaScript引擎,是SquirrelFish新的升级版本
TraceMonkey(JaegerMonkey  ) (gecko?)    | Firefox 
Carakan         | Opera   换了
Trident          |IE
edge            | edge 现在换了

chrome\safari\opera使用webkit引擎
13年chrome和opera开始使用Blink引擎
	

浏览器内核概念

浏览器内核分为两部分：渲染引擎（render engin）、js引擎(js engin)

渲染引擎：负责对网页语法的解释（HTML、 javaScript 、引入css等），并渲染（显示）网页

js引擎：javaScript的解释、编译、执行

主流内核：Trident(IE)、Gecko(FireFox)、Webkit(Safari)、Presto(opera前内核、已废弃)、blink(Chrome)

Chromium基于webkit，却把Webkit的代码梳理得可读性更高，多进程框架


V8引擎是谷歌公司自主研发的js引擎

Webkit基于KHTML，包含JavaScriptCore

Chromium基于Webkit，衍生出Blink

CEF是包含Chromium浏览器的应用框架

[简述Chromium, CEF, Webkit, JavaScriptCore, V8, Blink](https://www.codercto.com/a/43013.html)

[干货：浏览器渲染引擎Webkit和V8引擎工作原理](https://segmentfault.com/a/1190000018806562)


### user-agent
[常用浏览器(PC,移动) user-agent](http://tools.jb51.net/table/useragent)
[User-Agent格式含义与示例](https://www.jianshu.com/p/9453579154e3)

[为什么所有主要浏览器的 User-Agent 都是 Mozilla/x.0 开头？](https://www.zhihu.com/question/19553117)

[UAParser.js](https://github.com/faisalman/ua-parser-js)

[user-agent database](https://user-agents.net/)

## 开放源码发现服务
https://github.com/librariesio/libraries.io

https://libraries.io/

## Mock
在Node.js和浏览器中生成大量真实的假数据
https://github.com/Marak/faker.js

HTTP mocking for Golang
https://github.com/jarcoal/httpmock

Mimesis是一款用于mock数据的Python工具。
https://github.com/lk-geimfari/mimesis

dotnet版本的faker.js
https://github.com/bchavez/Bogus

## 字体
### FiraCode
程序员字体
https://github.com/tonsky/FiraCode 50.9k
Fira 是 Mozilla 主推的字体系列，Fira Code 是基于 Fira Mono 等宽字体的一个扩展，主要特点是加入了编程连字特性（ligatures）。

Fira Code 就是利用这个特性对编程中的常用符号进行优化，比如把输入的「!=」 直接显示成 「≠」 或者把 「>=」 变成 「≥ 」 等等，以此来提高代码的可读性。

- 连字特性
- 多种编辑器、IDE 支持
- 支持视网膜显示
- 经常更新
- 开源免费
- 提供 CSS
### Cascadia Code
https://github.com/microsoft/cascadia-code 13.6k

微软推出的一款等宽字体，旨在为命令行体验和代码编辑器提供了一种全新的体验。
Cascadia代码是与新的Windows Terminal 一起开发的。
官方推荐将该字体用于终端应用程序和文本编辑器，如Visual Studio和Visual Studio Code。

为什么命名为Cascadia Code？
Cascadia Code名称来源于 Windows Terminal 项目，在它发布之前，Windows Terminal的代号就是Cascadia，
事实上，Windows Terminal的一些源文件仍然使用这个名称，为了表示对 Terminal 的敬意，我们喜欢用它的代号来命名字体。

### Source Code Pro
https://github.com/adobe-fonts/source-code-pro 16.6k
由 Adobe 设计。整体而言，文本看起来比其他字体更清晰，更不用说其可区分的字符了。

- 视觉友好
- 所有显示器都清晰可辨
- 免费开源
- 适用于Google Web字体
- 提供多种款式
- 提供斜体

### Hack

https://github.com/source-foundry/Hack 13.7k

Hack 是一种专为源代码设计的开源字体，基于 Bitstream Vera 和 DejaVu 项目。0|O 并且 1lI 清晰可辨，字体易于整体阅读。特别是我喜欢削减零。

- webfonts有svg，eot，ttf，woff和woff2格式
- 久经考验的 Vera Sans Mono
- 源代码以UFO格式发布

### Input
https://input.fontbureau.com/

Input 是一个灵活的字体系统，由 David Jonathan Ross 专门为代码设计。
提供等宽字体和比例字体，所有字体都具有宽度，粗细和样式，以实现更丰富的代码格式。而且它还可以自定义，别提多强大了。

- 灵活配置
- 有 Mono，Sans 和 Serif 两种款式
- 明确区分相似的字符
- 清晰的低分辨率和视网膜显示

### Anonymous Pro

https://www.marksimonson.com/fonts/view/anonymous-pro

Anonymous Pro（2009）是一个由四个固定宽度字体组成的系列，设计时考虑了编码。
Anonymous Pro具有基于Unicode的国际字符集，支持大多数西欧和中欧拉丁语言，以及希腊语和西里尔语。它看起来有点像打字机的感觉。

- 四种风格：Regular，Italic，Bold 和 Bold Italic
- 等宽设计
- 各种计算机键盘符号
- 为大多数拉丁语的西欧和中欧语言以及希腊语和西里尔语提供扩展语言支持
### Ubuntu Mono
https://design.ubuntu.com/font/

Ubuntu字体专门用于补充Ubuntu。它具有现代风格，包含Ubuntu品牌独有的特征，传达出精确，可靠和自由的态度。

- 出色的可读性
- 易于辨别的符号
- 颜值不错
- 小尺寸也清晰可辨

### Menlo

https://github.com/hbin/top-programming-fonts/blob/master/Menlo-Regular.ttf 1.9k

Menlo 是 maccode 中用于 Xcode 和 Terminal 的新默认字体。
它是 DejaVu Sans Mono 的衍生物，Menlo也是编程的好选择。

- 视觉友好
- 好的可读性

### Consolas

https://docs.microsoft.com/zh-cn/typography/font-list/consolas

Consolas 旨在用于编程环境和指定等宽字体的其他情况。
所有字符都具有相同的宽度，就像旧的打字机一样，使其成为个人和商业对应的理想选择。
与传统的等宽字体（如Courier）相比，改进的Windows字体显示允许设计的比例更接近普通文本。
这允许更舒适地读取屏幕上的扩展文本。OpenType功能包括悬挂或衬里数字;削减，点缀和正常的零;以及许多小写字母的替代形状。
通过改变条形和波浪的数量，可以将文本的外观调整为个人品味。

- 颜值高
- 包含大多数 unicode
- 可以显示大量文本
- 适用于 Windows 和 OSX
- 可轻松区分相似字符



## E-Books
https://github.com/trumpowen/All-Programming-E-Books-PDF

## 其它
中国程序员容易发音错误的单词
https://github.com/shimohq/chinese-programmer-wrong-pronunciation

免费的编程书籍
https://github.com/EbookFoundation/free-programming-books
https://github.com/EbookFoundation/free-programming-books/blob/master/free-programming-books-zh.md


## 儿童教育
### Scratch 
### GCompris 
支持 Windows、Android、GNU/Linux 和 macOS 平台，可安装到台式机、笔记本电脑、平板电脑和移动设备上。
https://gcompris.net/downloads-en.html

## 社交网络
### Mastodon
https://github.com/mastodon/mastodon

### Elk
https://github.com/elk-zone/elk
Mastodon 网页客户端