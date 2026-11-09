

### skill hub
https://www.skills.sh/
https://agent-skills.md/

### 股票
美股
https://github.com/simonlin1212/global-stock-data

A股
https://github.com/simonlin1212/a-stock-data


#### OpenStock
开源的股票市场平台
https://github.com/Open-Dev-Society/OpenStock

#### tickflow-stock-panel
TickFlow Stock Panel 是一个开源的自托管量化工作台，支持 A 股（原版）以及 A 股/港股/美股（扩展版），提供选股、监控、回测等功能，可通过 Docker 一键部署，并允许接入 TickFlow 或第三方数据源。
https://github.com/shy3130/tick-stock-panel

#### TradingAgents
原版 [TradingAgents](https://github.com/TauricResearch/TradingAgents) 是一个出色的多 Agent 投研框架，但它针对美股设计：数据走 Yahoo Finance / Alpha Vantage，分析师不懂 A 股制度，辩论和决策完全面向美股市场。

本 [Fork](https://github.com/simonlin1212/TradingAgents-astock) 的目标：把 TradingAgents 的多 Agent 辩论架构真正落地到 A 股，不是简单翻译，而是从数据层、Agent 角色、交易规则三个维度做深度特化。

#### 其它


https://github.com/lzwme/finance-quant-skills
https://github.com/clong365/easy_tdx

- 同花顺金融数据
请帮我安装 hithink-finance Skill。

优先在终端执行以下命令：
npx skills add HiThink-Tech/Financial-API --skill hithink-finance -g --yes

如果该命令因网络等原因无法完成，请根据
https://skillhub.cn/install/skillhub.md
安装 hithink-finance。


### API网关
https://github.com/QuantumNous/new-api/blob/main/README.zh_CN.md

### frontend-slides(html格式的幻灯片)
https://github.com/zarazhangrui/frontend-slides

### superpowers vs grill-me


grill-me 和 superpowers 那些工具有什么不同
我试过不少 AI 增强工具，包括各种 superpowers 插件——预设模板多、自动化流程全、领域知识也齐。grill-me 不一样，它不替你定方案，而是把你的意图和约束挖出来。复杂、定制化的项目里，这点特别管用。

superpowers 适合快速套模式，比如某个框架的最佳实践。grill-me 更像元技能：先想清楚问题，再决定用什么模式。


https://github.com/obra/superpowers
https://github.com/mattpocock/skills
https://www.skills.sh/mattpocock/skills/grill-me


### /skill-creator
Skill 本质上是一个模块化知识包, 而 skill-creator 就是 Anthropic 官方提供的 Skill 开发助手，帮助开发者创建、优化和打包技能。
https://github.com/anthropics/skills/tree/main/skills/skill-creator

### 绘图/流程图/交互图
#### archify
https://github.com/tt-a1i/archify

#### fireworks-tech-graph
https://github.com/yizhiyanhua-ai/fireworks-tech-graph/blob/main/README.zh.md

fireworks-tech-graph 是一份可由 Codex 和 Claude Code 共用的 Agent Skill。它将自然语言描述转化为经过几何校验的 SVG、高分辨率 PNG、经过媒体探测验证的 SVG 转 GIF 语义动效与离线交互 HTML。聚焦后的动效链路只接收生成器产出的语义 SVG，只输出一个紧凑、可验证的 GIF。项目内置 11 种生成器风格 + 1 种 AI 手绘风格（Dark Luxury）；新增的四种工程风格分别为 C4 评审、云部署、事件流和可靠性排查提供可执行语义契约，同时保留 AI/Agent Pattern 与全部 14 种 UML 图类型。

