
## 发展
### function calling
### MCP
### agent to agent Protocol(A2A Protocol)
https://a2aprotocol.ai/blog/2025-full-guide-a2a-protocol-zh
Agent2Agent (A2A) 协议 - AI智能体协作的新标准
Agent2Agent (A2A) 协议是一个开放标准，专门解决AI智能体生态系统中的核心挑战：如何让不同团队、使用不同技术、属于不同组织的AI智能体有效沟通和协作

A2A协议：首个专为AI智能体间通信设计的开放标准，解决不同组织开发的AI智能体协作难题
核心价值：通过标准化通信协议，让专业化AI智能体能够无缝协作完成复杂任务
技术基础：基于JSON-RPC 2.0和HTTP(S)，支持流式传输、推送通知等企业级功能
与MCP互补：A2A专注智能体间协作，MCP专注工具集成，两者配合构建完整的智能体生态系统


## 开源框架

### 多agent
[multica](https://github.com/multica-ai/multica)
[paperclip](https://github.com/paperclipai/paperclip)

[paseo](https://github.com/getpaseo/paseo) 跨设备控制

Hermes：管调度和执行，负责深度。
MultiCA：管团队和项目，负责广度。

agent = model+harness

| 产品 | 与 Multica 相似度 | 核心区别 | 开源/自托管 |
|---|---:|---|---|
| [Paperclip](https://paperclip.ing/) | 很高 | 更偏“经营 AI 公司”，有组织架构、预算、上下级和治理 | MIT、支持自托管 |
| [OpenHands Agent Canvas](https://github.com/OpenHands/OpenHands) | 高 | 更偏工程自动化和 Agent 运行平台，任务管理弱一些 | MIT、支持自托管 |
| [Orca](https://github.com/stablyai/orca) | 高 | 更像多 Agent IDE，强调 Worktree、终端、浏览器和 Diff 审查 | MIT、本地/远程 |
| [Conductor](https://www.conductor.build/) | 中高 | 商业化多 Agent 工作台，强调云沙箱、并行开发和团队实时协作 | 闭源 SaaS |
| [OpenAI Symphony](https://github.com/openai/symphony) | 中高 | 从 Linear 看板自动启动 Codex 执行任务，产品界面和管理功能较少 | Apache 2.0、实验预览 |
| [Factory](https://www.factory.ai/) | 中 | 面向企业的软件开发 Agent 平台，强调治理、审计和大规模自动化 | 商业闭源 |
| [Herdr](https://github.com/herdrdev/herdr) | 中 | 管理长期运行的 Agent 终端和会话，不是完整项目管理系统 | Apache 2.0 |
| [Vibe Kanban](https://github.com/BloopAI/vibe-kanban) | 曾经很高 | 看板、Worktree、Diff 审查都很接近，但官方已宣布停止发展 | Apache 2.0、正在退场 |



最直接的三个竞品：
Paperclip
这是定位上最像 Multica 的：
把 Agent 当员工
有任务、项目和审批
支持 Claude Code、Codex、Cursor 等不同 Agent
有成本和预算控制
有组织架构、角色、权限、审计
开源、自托管
主要差异是：
Multica 更像“人类与 Agent 共用的 Linear/Jira”
Paperclip 更像“管理一间由 Agent 组成的公司”
Paperclip 功能野心更大；Multica 的 Issue、代码交付与人机协作模型相对自然。
OpenHands Agent Canvas
同样能够统一管理：
OpenHands
Claude Code
Codex
Gemini
ACP 兼容 Agent
本地、Docker、VM 和云端 Runtime
定时任务、Webhook、Slack、GitHub 和 Linear 自动化
它更偏“Agent 基础设施与自动化控制台”，而 Multica 更偏“团队项目管理与工作记录”。
Orca
如果重点是同时运行多个 Coding Agent，Orca 很接近：
多个 Agent 并行运行
每个任务使用独立 Git Worktree
支持大量 CLI Agent
内置终端、代码编辑、浏览器
Diff 评论和结果比较
支持手机查看、远程控制和 SSH
区别是 Orca 更像 IDE，Multica 更像项目管理系统：
```
Orca：开发者 → Worktree → Agent → Diff
Multica：团队 → Issue → Agent → 执行 → Review
```

#### agent
opencode 功能测试, 编写单元测试, 找bug
Codex 写代码, 改bug
Claude code 写代码, 改bug, 代码审查
hermes 任务设计, 需求拆解
pi agent

### Agno
Agno 是一个用于构建多模态智能体的轻量级框架。它支持多种数据模态（如文本、图像、音频和视频），并且可以快速创建智能体。Agno 提供了内存管理和知识库支持，能够将用户会话和智能体状态存储在数据库中，基于向量数据库实现动态少样本学习。此外，Agno 支持多智能体协作，帮助用户实时跟踪智能体会话和性能。

Agno 的设计目标是简化开发流程，提升性能，并确保灵活性。通过无依赖性架构和纯 Python 实现，开发者可以轻松上手并快速构建高效的智能体应用。
https://github.com/agno-agi/agno

### langchain
https://python.langchain.ac.cn/docs/introduction/

### agentic workflow 
https://github.com/n8n-io/n8n
https://github.com/langgenius/dify/

### web ui
[Open WebUI](https://github.com/open-webui/open-webui)
[LobeChat](https://lobechat.com/chat)
[Cherry Studio](https://www.cherry-ai.com/)

### ADK(Agent Development Kit )
https://github.com/google/adk-go/blob/main/README.md
https://adk.wiki/get-started/go/
https://github.com/cloudwego/eino