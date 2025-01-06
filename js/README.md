[TOC]

[学习](https://github.com/freeCodeCamp/freeCodeCamp)
https://contribute.freecodecamp.org/#/

[Lunr.js](https://github.com/olivernn/lunr.js) 轻量级Javascript全文搜索库

[Snowpack 打包工具](https://github.com/pikapkg/snowpack)
实现直接在浏览器中运行npm软件包。SnowPack 是基于 ESM 的工具。
## 如何对html源码中的数字加密
用字体
为了防止别人爬取网页价格，可以把价格加密，用字体显示
参考1：https://i.jzj9999.com/quoteh5/
参考2：https://sz.58.com/chuzu/
[python解析字体反爬](https://www.cnblogs.com/eastonliu/p/9925652.html)

```
from fontTools.ttLib import TTFont
font = TTFont('58.ttf') # 打开本地的ttf文件
font.saveXML('58.xml')  # 转换成xml
```

如果是爬取的话有更简单的方式，一般数字就是0-9以及小数点，把密文复制出来找到所有的数字和小数点，然后转成unicode（当然如果简单的也可以不用转unicode，直接比对），做一个map映射就好了
https://www.bejson.com/convert/unicode_chinese/


## js学习
学习js必须要掌握的++
[刷新你对js认识 ](https://github.com/lydiahallie/javascript-questions)

## 富文本编辑器
### Quill
https://github.com/quilljs/quill

## SSG 
 Static Site Generators (SSG) 技术

 生成静态网页：

Next.js 是基于 React 的 SSR/SSG 框架。
Scully 是基于 Angular 的 SSG 框架。
VitePress 是 Vue 官方推出的 SSG 框架。

JAMStack
Ledge 
Gatsby

## 兼容
用来为旧浏览器提供它没有原生支持的较新的功能。
https://github.com/financial-times/polyfill-service

## 出错时录屏
[前端录屏+定位源码，帮你快速定位线上bug](https://juejin.cn/post/7173596154297810957) [Github源码](https://github.com/xy-sea/web-see-demo)


[基于H5的录屏工具开发](https://juejin.cn/post/6915287057795874824)

## 工具
### doc
https://github.com/dolanmiu/docx

https://products.fileformat.com/word-processing/javascript/docx/

https://medium.com/geekculture/create-dynamic-word-documents-using-docx-js-file-saver-and-data-from-an-excel-or-json-dbd5e4ec823f

https://github.com/open-xml-templating/docxtemplater

https://github.com/guigrpa/docx-templates

https://github.com/VolodymyrBaydalka/docxjs

https://github.com/lalalic/docx4js

### table
https://github.com/nocodb/nocodb
https://github.com/Airtable/airtable.js
https://github.com/teableio/teable

### calendar 
https://github.com/nhn/tui.calendar

### 上传
#### FilePond
https://github.com/pqina/filepond
#### Uppy
https://github.com/transloadit/uppy

### 图片处理
https://github.com/nhn/tui.image-editor

去除背景
https://github.com/imgly/background-removal-js

### SQL
#### 解析sql、sql2ast
https://github.com/DTStack/dt-sql-parser
https://github.com/JavaScriptor/js-sql-parser
https://github.com/godmodelabs/flora-sql-parser

## 网页构建  Web Builder Framework
### appsmith
https://github.com/appsmithorg/appsmith 8.1k
### grapesjs
https://github.com/artf/grapesjs 12k
### VvvebJs
https://github.com/givanz/VvvebJs 3.5k
可视化构建库拖拽生成网页
1、组件和块/片段拖放。
2、撤销/重做操作。
3、一个或两个面板界面。
4、文件管理器和组件层次结构导航添加新页面。
5、实时代码编辑器。
6、包含示例php脚本的图像上传。
7、页面下载或导出html或保存页面在服务器上包含示例PHP脚本。
8、组件/块列表搜索。
9、Bootstrap 4组件等组件

### h5-Dooring
https://github.com/MrXujiang/h5-Dooring 870
### sparrow-js
https://github.com/sparrow-js/sparrow 349
### 其它相似
#### ContentTools
用于为HTML内容构建所见即所得编辑器的JS库。
https://github.com/GetmeUK/ContentTools
https://github.com/GetmeUK

H5 页面设计器
https://github.com/ymm-tech/gods-pen

#### 表单设计器
https://github.com/JakHuang/form-generator

### AI自动生成前端代码
#### Sketch2Code
https://github.com/Microsoft/ailab/tree/master/Sketch2Code
https://sketch2code.azurewebsites.net/

#### teleportHQ
https://github.com/teleporthq

## 在线
### 可视化在线绘图引擎 Topology（类似draw.io）

https://github.com/le5le-com/topology/
A diagram (topology, UML) framework uses canvas and typescript. 一个轻量（100k左右）、功能丰富的绘图工具（微服务架构图、拓扑图、流程图、类图等UML图、脑图，动画、视频支持）。 【在线使用】：

topology.le5le.com/

- drawio
draw.io

https://github.com/jgraph/drawio-desktop

- Excalidraw: 虚拟白板，用于素描手绘图
https://github.com/excalidraw
https://github.com/excalidraw/excalidraw
https://excalidraw.com/

- wireflow.co
https://github.com/vanila-io/wireflow
https://app.wireflow.co/

-zwibbler
https://zwibbler.com/

- poster-design/在线图片编辑器/在线海报设计器
https://github.com/palxiao/poster-design

- [Gitmind 在线思维导图](https://gitmind.com/) https://gitmind.cn/


> [10款最佳HTML5绘图工具](https://www.cnblogs.com/jackyWHJ/p/3872098.html)

### 流程图
https://github.com/antvis/G6
https://github.com/antvis/X6
https://github.com/didi/LogicFlow

### 终端应用程序Xterm.js
https://github.com/xtermjs/xterm.js

### 在线运行Android 
https://github.com/openstf/stf

### 在线数学公式
http://mathquill.com/
https://github.com/mathquill/mathquill

## 视频播放器
### video.js
https://videojs.com/
https://github.com/videojs/video.js 36.6k

### Plyr
https://github.com/sampotts/plyr 24.8k

### xgplayer(西瓜播放器)
https://v2.h5player.bytedance.com/
https://github.com/bytedance/xgplayer 7.1k

### DPlayer
https://dplayer.diygod.dev/zh/guide.html
https://github.com/DIYgod/DPlayer 14.7k


### Android/iOS video player 手机端播放器
https://github.com/bilibili/ijkplayer


## state
### xstate
https://github.com/davidkpiano/xstate
它是个简单的 JavaScript 和 TypeScript 框架，可以创建有限状态机并可视化为状态图。
它可以跟最流行的响应式 JavaScript 框架（Vue.js，Ember.js，React.js 以及 RxJS）集成，并基于 W3C 标准来创建有限状态机。
### immer
https://github.com/immerjs/immer

## test
### jest
https://github.com/facebook/jest
#### jest-when
https://github.com/timkindberg/jest-when

## doc

提取文档,然后进行灰度化, 就是一个扫描件了: https://github.com/danielgatis/docscan
图片扫描
### PDF
#### PDF扫描
https://github.com/rwv/lookscanned.io
#### PDF工具
https://github.com/Frooodle/Stirling-PDF
#### jsPDF
生成PDF
https://github.com/MrRio/jsPDF 20.3k

#### pdf.js
https://github.com/mozilla/pdf.js 32.7k
在web端打开pdf格式文件bai
#### pdfobject
在web端打开pdf格式文件bai
https://github.com/pipwerks/PDFObject 1.7k

#### html2pdf
https://github.com/eKoopmans/html2pdf.js

#### MarkItDown/转markdown
PDF (.pdf)
PowerPoint (.pptx)
Word (.docx)
Excel (.xlsx)
Images (EXIF metadata, and OCR)
Audio (EXIF metadata, and speech transcription)
HTML (special handling of Wikipedia, etc.)
Various other text-based formats (csv, json, xml, etc.)
```
pip install markitdown

from markitdown import MarkItDown
markitdown = MarkItDown()
result = markitdown.convert("test.xlsx")
print(result.text_content)


from markitdown import MarkItDown
from openai import OpenAI
client = OpenAI()
md = MarkItDown(mlm_client=client, mlm_model="gpt-4o")
result = md.convert("example.jpg")
print(result.text_content)
```

### excel
https://github.com/mengshukeji/Luckysheet

[其它](../awesome/online.md#excel)

## D3.js
https://github.com/d3/d3 95k
D3的全称是（Data-Driven Documents），顾名思义可以知道是一个关于数据驱动的文档的javascript类库。说得简单一点，D3.js主要是用于操作数据的，它通过使用HTML、SVG、CSS来给你的数据注入生命，即转换为各种简单易懂的绚丽的图形。

D3 是最流行的可视化库之一，它被很多其他的表格插件所使用。它允许绑定任意数据到DOM，然后将数据驱动转换应用到Document中。你可以使用它用一个数组创建基本的HMTL表格，或是利用它的流体过度和交互，用相似的数据创建惊人的SVG条形图。


## 前端框架
[github](https://github.com/krausest/js-framework-benchmark)
[最新成绩对比](https://krausest.github.io/js-framework-benchmark/index.html)
[javascript frameworks performance comparison](https://medium.com/@ajmeyghani/javascript-frameworks-performance-comparison-c566d19ab65b)
[JS 框架性能对比](https://www.infoq.cn/article/ebDcihIZbEZoFU9q6pi7)

[alpine](https://github.com/alpinejs/alpine)

1. Vue (177k)
2. React (161k)
3. Angular (68.9k)
4. Svelte (40.5k)
5. Preact (27.9k)
6. Ember (21.7k)
7. HyperApp(18.2k)
8. Inferno (14.6k)
9. Riot (14.4k)
10. Yew (14.2k)
11. Mithril (12.5k)
12. Alpine (12.4k)
13. Knockout (9.9k)
14. Marko (9.9k)
15. lit-html (6.9k)
16. Rax (7k)
17. Elm (6.2k)
18. Ractive (5.8k)
19. Solid (4.7k)
20. Imba (4.1k)

排名2018
1. Solid (57)
2. HyperApp (54)
3. Inferno (51)
4. Svelte (51)
5. Elm (46)
6. Riot (40)
7. Preact (39)
8. Imba (36)
9. lit-html (36)
10. Yew (32)
11. Vue (29)
12. Mithril (29)
13. Marko (28)
14. Alpine (28)
15. React (19)
16. Rax (16)
17. Angular (12)
18. Knockout (11)
19. Ractive (8)
20. Ember (6)


## UI

[UI整理](../awesome/ui.md)

### 开源UI
#### VisActor 是一个数据可视化解决方案
https://github.com/VisActor

VTable高性能的多维数据分析表格
https://github.com/VisActor/VTable

### 桌面UI
- [YLUI](https://github.com/yuri2peter/ylui)
- [win10-ui](https://github.com/yuri2peter/win10-ui)
- https://github.com/LuckyZmj/LuckyZmj.github.io

### ionic-framework
Ionic 类似 React Native 的跨平台框架，支持vue等

https://github.com/ionic-team/ionic-framework 42.9k

### vue3

#### nuxt
[nuxt框架](https://github.com/nuxt/nuxt) 48.5k

[nuxt ui](https://github.com/nuxt/ui) 2.3k

#### quasar
https://github.com/quasarframework/quasar 19.2k

#### primevue
Vue, Angular, React and Java
https://github.com/primefaces/primevue 1.5k


### react
#### evergreen
https://github.com/segmentio/evergreen/ 10.4k
#### material-ui
https://github.com/mui-org/material-ui 64.5k
#### fluentui
https://github.com/microsoft/fluentui 10.3k
https://developer.microsoft.com/en-us/fluentui#/get-started
#### geist-ui
https://github.com/geist-org/react 1.4k

## Gantt
### 商业
> dhtmlx,GrapeCity,Syncfusion 都有dotnet
#### dhtmlx
https://dhtmlx.com/docs/products/dhtmlxGantt/
#### GrapeCity
https://www.grapecity.com/dataviewsjs
#### webix
https://webix.com/gantt/
#### Syncfusion
https://ej2.syncfusion.com/
https://syncfusion.com/

#### treegrid
http://www.treegrid.com/Gantt
#### bryntum
https://www.bryntum.com/

### gantt-schedule-timeline-calendar
https://github.com/neuronetio/gantt-elastic
https://github.com/neuronetio/gantt-schedule-timeline-calendar

### VTable/vtable-gantt
https://github.com/VisActor/VTable

### 国产
https://github.com/w1301625107/Vue-Gantt-chart
https://github.com/hql7/wl-gantt

### 其它
https://github.com/robicch/jQueryGantt



## 数据分析
[原CNZZ站长工具 - 被阿里收购](https://www.umeng.com/)
[百度统计](https://tongji.baidu.com/web/demo/overview/index?siteId=16847648)
[GoogleAnalytics分析](https://developers.google.cn/analytics/devguides/platform/)
[mixpanel](https://github.com/mixpanel/mixpanel-js)

[PostHog - 开源，可以自己托管数据](https://github.com/PostHog/posthog)
[PostHog doc](https://posthog.com/docs/contribute/developing-locally)

[华为分析服务](https://developer.huawei.com/consumer/cn/hms/huawei-analyticskit?ha_source=hms1)

### 网站监控/网站分析/Google Analytics
前端性能监控与用户行为分析, 如: Google Analytics / Cloudflare Web Analytics / CNZZ / 51LA

- https://github.com/matomo-org/matomo      15.5k
- https://github.com/plausible/analytics    9k
- https://github.com/mikecao/umami          8.8k
- https://github.com/Open-Web-Analytics/Open-Web-Analytics  1.6k

[其它开源](https://github.com/topics/web-analytics)

https://github.com/a597873885/webfunny_monitor


#### Umaimi
https://github.com/umami-software/umami 17.7k
#### Shynet
https://github.com/milesmcc/shynet 2.9k
#### Clarity
https://clarity.microsoft.com/

https://clarity.microsoft.com/projects

微软Clarity是一款免费的网站分析工具，可以提供全面的用户行为分析和可视化数据，更好地理解用户。

1.用户会话重放：Clarity允许您观看真实用户的会话录像，了解他们在网站上的行为。

2.热力图：Clarity 可生成热力图，通过可视化显示用户在网站上的点击和滚动行为。

3.点击地图：这个功能显示用户在页面上的点击位置，帮助您了解用户的兴趣点和点击模式。

4.用户浏览路径：Clarity 提供了用户在网站上的浏览路径追踪。

5.性能分析：除了用户行为分析，Clarity还提供了网站性能数据，包括加载时间和渲染性能。

6.隐私保护：Clarity非常注重用户隐私，所有数据都经过匿名化处理，不会追踪个人身份或敏感信息。

#### hotjar
https://www.hotjar.com/
