<!--toc-->
[TOC]
# CodeReview 

## 代码质量审核和管理工具
[11个代码质量审核和管理工具](https://www.toutiao.com/a6788710794432348687)
### SonarQube
项目工程代码质量检测
[SonarQube](https://github.com/SonarSource/sonarqube)
[SonarCloud](sonarcloud.io)
- go
- java
- dotnet
- python
- T-SQL
- html
- javascript
- TypeScript
- c/c++
- Swift
- Ruby
- Kotlin
- Scala
- php
- Objective-C
- VB
- XML
- etc...

### Kritika
http://kritika.io/
### DeepScan
### Klocwork
### CodeSonar
### JArchitect
### Bandit
### Code Climate
### Crucible
### Fortify
### Codecov

## Phabricator - 可视化代码评审工具
Facebook出品，支持Git, Mercurial, and SVN
https://github.com/phacility/phabricator/ 11.6k

## Gerrit - 代码评审工具
Google出品
https://github.com/GerritCodeReview/gerrit
(mirror of https://gerrit.googlesource.com/gerrit)

## 谷歌工程实践文档
https://xindoo.github.io/eng-practices-cn/review/reviewer/
https://github.com/xindoo/eng-practices-cn

## 代码审查工具
[17款最佳的代码审查工具](https://www.toutiao.com/i6750444677385683468)
### CodeStriker
http://codestriker.sourceforge.net/index.html
### RhodeCode
https://rhodecode.com/
### Codebrag
http://codebrag.com/
### Phabricator
http://phabricator.org/
### Codifferous
https://codifferous.com/
### Getbarkeep
http://getbarkeep.org/
### Crucible
https://www.atlassian.com/software/crucible/overview
### Code Review Tool
http://codereviewtool.com/
### Malevich
http://malevich.codeplex.com/
### SmartBear
http://smartbear.com/product/collaborator/overview/
### Review Assistant
Review Assistant是一款支持Visual Studio的简单又优秀的代码审查工具。
https://visualstudiogallery.msdn.microsoft.com/9ef817b4-2c6d-4213-8b08-5be48f9d91b9
### Review Board
### Peer Review Plugin
http://trac-hacks.org/wiki/PeerReviewPlugin
### Code Reviewer
https://codereview.appspot.com/
### Code Analysis Tool
http://www.castsoftware.com/products/code-analysis-tools
### jArchitect
http://www.jarchitect.com/
### Reviewale

## 其它
GitHub、GitLab 有基于 PR 的审查
评审的关键操作：
- 提高提交的原子性
    每次提交的代码粒度至关重要，我们可以反过来思考：
    - 如果提交的是半个功能的代码会怎么样？
    - 如果提交的是一周的代码会怎么样？

    所以原子性就是合适的粒度，大功能要拆分来提交，一周的代码的代码要按天来提交。否则对于评审人员来说是很反感的，因为只会增加审查的难度。

- 提高提交说明的质量
    我们应该经常见到很多的提交是这样描述的：
    - BUG
    - FIX
    - 更新

    下面是葛俊给出的三个改进步骤：
    - 第一步，规定提交说明一定要包括标题、描述和测试情况三部分，但暂时还不具体要求必须写多少字。比如，测试部分可以简单写一句“没有做测试”，但一定要写。如果格式不符合要求，审查者就直接打回。这个格式要求工作量很小，比较容易做到，两个星期后整个团队就习惯了。虽然只是在提交说明里增加了简单描述，也已经为审查者和后续工作中进行问题排查提供一些必要信息，所以大家也比较认可这个操作。

    - 第二步，要求提交说明必须详细写明测试情况。如果没有做测试一定要写出具体理由，否则会被直接打回。这样做，不但为审查者提供了方便，还促进了开发人员的自测。整个团队在一个多月后，也养成了详细描述测试情况的习惯。

    - 第三步，逐步要求提交的原子性。我要求每一个提交要在详细描述部分描述具体实现了哪些功能。如果一个提交同时实现了多个功能，那就必须解释为什么不能拆开提交；如果解释不合理的话，提交直接被打回。

    提交说明是提高代码审查的利器，好的格式应该包含以下几个方面：
    - 标题，简明扼要地描述这个提交。这部分最好在 70 个字符之内，以确保在单行显示的时候能够显示完整。比如，在命令行常用的 git log --oneline 输出结果要能显示完全。

    - 详细描述，包括提交的目的、选择这个方法的原因，以及实现细节的总结性描述。这三个方面的内容最能帮助审查者阅读代码。

    - 测试情况，描述的是你对这个提交做了什么样的测试验证，具体包括正常情况的输出、错误情况的输出，以及性能、安全等方面的专项测试结果。这部分内容，可以增加审查者对提交代码的了解程度以及信心。

    - 与其他工具和系统相关的信息，比如相关任务 ID、相关的冲刺（sprint，也可翻译为“迭代”）链接。这些信息对工具的网状互联提供基础信息，非常重要。这里还有一个 Git 的技巧是，你可以使用 Git 的提交说明模板（Commit Message Template），来帮助团队使用统一的格式。

评审的关键原则：
- 相互尊重
- 基于讨论

## 可以先机器审查
使用 GitLab、Jenkins 和 SonarQube 进行配置。具体使用 GitLab 管理代码，代码入库后通过钩子触发 Jenkins，Jenkins 从 GitLab 获取代码，运行构建，并通过 Sonar 分析结果