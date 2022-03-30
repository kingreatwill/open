[GitHub 上有哪些优秀的项目？](https://www.zhihu.com/question/20584141)

### 2021.11.01
#### 代码 bug 检测器：control-flag
[ControlFlag](https://github.com/IntelLabs/control-flag) 是 Intel 实验室开源的代码 bug 检测工具，它可以帮开发者检测代码中是否存在 bug，它通过学习（挖掘）开源项目中使用的典型模式（typical pattern）来判断是否输入的特定代码中存在异常。

#### 通知工具：notifire
[Notifire](https://github.com/notifirehq/notifire) 提供 API 供你来管理多种事务通知， 例如：邮件通知、短信通知、消息推送…它具有以下特点：

- 一个 API 可管理所有通知信息
- 配备模版引擎，用于进阶使用（布局和设计）
- 易用，方便集成
- 用 TS 可预测静态类型写入

#### 钓鱼工具箱：zphisher
[Zphisher](https://github.com/htr-tech/zphisher) 是个新手友好的学习安全的项目，它提供了 30+ 网络钓鱼页面模版，用它可以部署一个或者多个钓鱼网站…进而了解到网络钓鱼工具的工作原理。友情提醒：尝试该项目存在一定的风险。
`docker run --rm -it -p 8080:8080 htrtech/zphisher`

#### 多目标跟踪：ByteTrack
[ByteTrack](https://github.com/ifzhang/ByteTrack) 轻量级的多目标检测、跟踪工具。这里简单介绍下多目标跟踪工作原理：通过检测画面中的物体，并检测其同目标物体相似度给予一定的分数，而之前对检测得分低的物体会进行简单、粗暴地丢弃，这样会导致真实推丢失以及目标轨迹碎片化，为此 ByteTrack 团队关联每个检测框而是单一靠记分函数来解决该问题。下图为常见多目标跟踪工具同 ByteTrack 对比。

### 2022.03.30
#### SVG 矢量图转换: VTracer
VTracer 是 GitHub 上一款开源工具，可快速将 JPG、PNG 等格式的图片快速转换为 SVG 矢量图，并支持过滤斑点、色彩精度、曲线拟合等多种参数配置。

在线体验：
https://www.visioncortex.org/vtracer/

GitHub：https://github.com/visioncortex/vtracer

#### 轻量级 Web 绘画 App: tldraw
tldraw 是一个轻量级，功能强大的绘画 App，自带画笔、橡皮、线框、文字等工具，用户可自定义画笔颜色、线框样式等。
Demo：https://www.tldraw.com/
GitHub：https://github.com/tldraw/tldraw

#### UI 设计与原型制作平台: Penpot

地址：https://penpot.app/

GitHub：https://github.com/penpot/penpot

#### DevOps 开发工具

GitHub 上一款开源的 DevOps 开发工具：Dev Lake，可将 DevOps 数据以实用、个性化、可扩展的视图呈现，完成数据的收集、分析和可视化。

工具内置 20+ 效能指标与下钻分析能力，可快速归集 DevOps 全流程效能数据，并支持接入新数据源，完成自定义 SQL 分析及拖拽搭建场景化数据视图等功能。
GitHub：https://github.com/merico-dev/lake