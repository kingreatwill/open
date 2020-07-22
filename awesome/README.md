
## 序列化



https://github.com/apache/thrift

https://github.com/protocolbuffers/protobuf

https://github.com/msgpack/msgpack

https://github.com/google/flatbuffers

https://github.com/capnproto/capnproto 无dotnet

https://avro.apache.org/


https://github.com/niXman/yas

https://github.com/USCiLab/cereal


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


## 开放源码发现服务
https://github.com/librariesio/libraries.io

https://libraries.io/