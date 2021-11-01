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

#### 多目标跟踪：ByteTrack
[ByteTrack](https://github.com/ifzhang/ByteTrack) 轻量级的多目标检测、跟踪工具。这里简单介绍下多目标跟踪工作原理：通过检测画面中的物体，并检测其同目标物体相似度给予一定的分数，而之前对检测得分低的物体会进行简单、粗暴地丢弃，这样会导致真实推丢失以及目标轨迹碎片化，为此 ByteTrack 团队关联每个检测框而是单一靠记分函数来解决该问题。下图为常见多目标跟踪工具同 ByteTrack 对比。