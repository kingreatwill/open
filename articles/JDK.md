[TOC]
# 甲骨文宣布新的Java SE订阅模式
甲骨文公司已经推出了Java SE（标准版）的商业支持计划，由之前企业一次性支付永久许可费用加年度支持费用的模式，现在更改为新的订阅模式，并宣布新的模式将于2018年7月开始启用。个人使用和非商业使用可继续享受免费支持，无需订阅。

该订阅被称为 Java SE Subscription，为用于任务关键型 Java 部署的新程序提供商业许可，并具有诸如高级 Java 管理控制台等功能。 此外，Oracle Premier 在当前和以前的 Java SE 版本中都提供支持，Java SE 8 和 Java SE 7 都在支持范围内。2019年1月之后，Oracle 将要求企业订阅才能继续获得 Java SE 8 的更新。

## Java SE Subscription 的定价
对于服务器和云实例，每个处理器每月收费为 25 美元，并提供批量折扣。对于个人电脑，每个用户每个月定价 2.50 美元起步，提供批量折扣。订阅周期分为一年、两年和三年三种。 Oracle 已经发布了新的 Java SE Subscription 计划的条款。

在之前的 Java SE Advanced 中，每台服务器的计划定价为 5000 美元，外加 1100 美元的支持费用，指定用户每人每次交纳的许可证费用为 110 美元，外加 22 美元的支持费用。 Oracle 对其他 Java 许可证也采用了类似的定价组合。
## Java SE Subscription 的特性
在超过公共更新（EoPU）时间之前访问一些 Oracle Java SE 版本。

尽早获得重要的错误修复。

许可和对云，服务器和桌面部署的支持。

性能，稳定性和安全性更新。

企业管理，监控和部署功能。

全天候支持。
## 对发布计划的影响
甲骨文已经实施了每六个月发布一个 Java SE 版本的发布计划，JDK 10 已于2018年3月发布，下一个版本 JDK 11 预计将于2018年9月发布。Java SE 用户可以按这个发布时间表保持更新，或者控制生产环境的应用程序何时使用新版本。订阅用户在八年之内可以继续更新长期支持（LTS）版本。 
## 是否必须订阅？
如果用户不继续订阅，他们将失去在订阅期间下载的所有商业软件的相关权利，Oracle Premier 支持也将终止。甲骨文建议那些选择不继续订阅的公司，在订阅结束之前迁移到 OpenJDK，以确保他们的应用程序不受影响。


# JDK

甲骨文已经改进了Java SE（标准版）的商业支持计划，由之前企业一次性支付永久许可费用加年度支持费用的模式，改为新的订阅模式(Java SE Subscription)，当然个人使用和非商业使用可继续享受免费支持，无需订阅。详情请参考[甲骨文宣布新的Java SE订阅模式](#甲骨文宣布新的java-se订阅模式)，为企业提供更广泛支持。在本文中，我们不妨了解一下 Oracle JDK 之外的 JDK，以及围绕 OpenJDK 构建的生态系统所拥有的深度。本文将为大家介绍一些主流的 OpenJDK 变种版本。

## OpenJDK builds
实际上，JDK 只有一组源代码。源代码使用 Mercurial（分布式版本控制系统）托管在 OpenJDK。任何人都可以获取该源码，并通过源码构建一个变种版本发布到网络上。但是需要一个独有的认证程序(certification process)来确保构建的变种版本是有效的。

这个认证程序是由 JCP(Java Community Process) 组织审核的，后者会提供技术兼容性工具包（TCK，有时也称为 JCK）。如果一个组织构建了一个 OpenJDK 的变种版本，并通过了 TCK 的兼容性测试，则可将构建的这个变种版本称为"Java SE compatible"（兼容 Java SE 的 JDK）。

要注意的是，如果提供者没有从 Oracle 获得商业许可，不能将该构建版本称为"Java SE"。例如，通过 TCK 兼容性测试的 AdoptOpenJDK 不是"Java SE"，而是兼容 Java SE 的 JDK或兼容 Java SE 规范的 JDK。还要注意，认证程序目前是基于信任基础的，结果不会提交给 JCP/Oracle 用于检查，也不会被公开。

总的来说，OpenJDK + 变种版本的提供者将一个源码库转换为许多不同的变种构建版本。
在将 OpenJDK 源码转换为变种版本的过程中，提供者可能会添加一些额外的标记或实用程序，但注意不要影响后面的认证程序。例如，提供者无法为 API 添加一个新的公共方法，或一项新的语言特性。

## Oracle JDK
从 Java 11 开始，这是一个提供付费支持的品牌商业版本。当然个人使用和非商业使用仍可继续享受免费支持（Oracle JDK 仍将为开发、测试、原型或展示目的的使用提供免费支持），不需要商业支持或企业管理工具的人可以选择使用 Oracle 的 OpenJDK 构建。Oracle 计划在 2026 年后提供全额付费的服务支持。要注意的是，与过去不同，Oracle JDK 并不比 OpenJDK “更好”（前提是两者都处于相同的安全补丁级别）。

## OpenJDK builds by Oracle
这些是免费的、完全无品牌的 OpenJDK 版本，基于 GPL 开源协议（+Classpath Extension），公司可安全且放心使用。这些版本仅在发布后的六个月内可以使用。要继续使用由 Oracle 的 OpenJDK 构建版本和安全补丁，需要在发布新版本后的一个月内升级至新版本。

## AdoptOpenJDK builds

这些版本也是免费的、完全无品牌的 OpenJDK 版本，基于 GPL 开源协议（+Classpath Extension），以免费软件的形式提供社区版的 OpenJDK 二进制包，公司也可安全且放心使用。与由 Oracle 的 OpenJDK 构建版本不同，这些版本会提供更长的支持，像 Java 11 一样，至少提供 4 年的免费长期支持(LTS)计划。AdoptOpenJDK 是一个由社区驱动的项目，如果其他群组在 OpenJDK 的源码仓库中创建和发布了安全修复程序，它们也会提供构建。 IBM 和 Red Hat 也曾表示他们打算提供这些安全补丁。

## AdoptOpenJDK OpenJ9 builds
除了标准的 OpenJDK 构建外，AdoptOpenJDK 还提供了使用 OpenJ9 而非 HotSpot 的版本。OpenJ9 最初是由 IBM 实现的 JVM，现在已开源并交由 Eclipse 运作。

## Red Hat OpenJDK builds
Red Hat 通过 Red Hat Enterprise Linux (RHEL) 提供了 OpenJDK 的变种构建版本，这也是提供付费支持的商业版本。他们在为 OpenJDK 提供安全补丁方面做得非常好，而且 Red Hat 还为 Java 6 和 7 提供安全更新。Red Hat 构建的版本能更好地集成到操作系统中，所以它称不上是纯粹的 OpenJDK 版本（尽管你也不会注意到差异）。

## Other Linux OpenJDK builds
不同的 Linux 发行版拥有不同的方式来访问 OpenJDK。这里是一些常见的发行版：Debian、Fedora、Arch、Ubuntu。

## Azul Zulu

Zulu 是 OpenJDK 的免费版本，但同时提供商业付费支持。当然不购买收费的技术支持的话，Azul 也有为 Zulu 提供免费的社区技术支持。Azul 有一个广泛的计划以支持 Zulu 商业化，包括支持 Java 9,13 和 15，这点与其他的提供者有不同之处。

## IBM
IBM 为 Java 8 及更早版本提供并支持 JDK。他们还使用 OpenJ9 为 AdoptOpenJDK 构建提供商业付费支持。

## SAP
SAP 使用 GPL + CE 许可证为 Java 10 及更高版本提供 JDK。他们还有一个商业的闭源 JVM。不过没有找到任何有关支持生命周期的信息。

## Amazon Corretto
https://aws.amazon.com/cn/corretto/
https://github.com/corretto

## Alibaba Dragonwell
https://github.com/alibaba/dragonwell8

## 总结

现在有很多不同的 OpenJDK 变种版本，它们都基于原始的上游代码仓库。但每个构建版本都提供了独有的选择，免费或是商业，品牌或是非品牌。可以选择当然很好，但如果你追求的是“标准”，那么目前我最好的建议是使用 Oracle 的 OpenJDK 构建版本、AdoptOpenJDK 构建版本或操作系统(Linux)中内置的版本。