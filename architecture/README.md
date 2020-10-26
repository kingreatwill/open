<!--toc-->
[TOC]
# 定律、理论、原则和模式
https://www.oodesign.com/behavioral-patterns/

https://www.oodesign.com/

更多可以参考[对开发人员有用的定律、理论、原则和模式](https://github.com/nusr/hacker-laws-zh)

## 软件设计模式(Software design patterns)



## 软件架构模式
https://www.cnblogs.com/doit8791/p/9343826.html

https://www.sohu.com/a/403458571_185201
https://dzone.com/articles/5-major-software-architecture-patterns
### 微内核模式(Microkernel Pattern)
微内核架构模式也称为插件模式

### 微服务模式 (Microservices Pattern )
### 分层架构模式( Layered Architecture Pattern)
### 基于事件的模式(Event-based Pattern)

### 基于空间的架构模式(Space-based Pattern)
基于空间的架构模式，可以专门用于解决软件系统的伸缩性和并发性问题。

对于用户访问量经常发生变化、偶尔出现高并发的应用程序，这是一种有用的软件架构模式。这种模式，通过消除中央数据库约束，并使用复制基于内存的数据网格来实现伸缩性。

基于空间的架构模式旨在通过在多个服务器之间拆分处理和存储数据，来避免高负载下的软件系统功能崩溃。


## 编码原则

### SOLID原则
> SOLID 是面向对象设计（OOD）的五大基本原则的首字母缩写组合，由俗称“鲍勃大叔”的Robert C.Martin在《敏捷软件开发：原则、模式与实践》一书中提出来。

- S(Single Responsibility Principle)：单一职责原则，简称SRP
- O(Open Close Principle)：开放封闭原则，简称OCP
- L(Liskov Substitution Principle)：里氏替换原则，简称LSP
- I(Interface SegregationPrinciple)：接口隔离原则，简称ISP
- D(Dependence Inversion Principle)：依赖倒置原则，简称DIP

### 微服务设计IDEALS原则

[微服务设计的原则：IDEALS，而不是SOLID](https://www.toutiao.com/i6871251586476147212)
[Principles for Microservice Design: Think IDEALS, Rather than SOLID](https://www.infoq.com/articles/microservices-design-ideals/)

- 接口分离（Interface segregation）
    指的是不同类型的客户端（移动应用程序，web应用程序，CLI程序）应该能够通过适合其需求的协议与服务端交互。
- 可部署性（deployability）
    指的是在微服务时代，也就是DevOps时代，开发人员需要在打包、部署和运行微服务方面做出关键的设计决策和技术选择。
- 事件驱动（event-driven）
    指的是在任何时候，都应该对服务进行建模，通过异步消息或事件而不是同步调用。
- 可用性胜于一致性（Availability over Consistency）
    指的是最终用户更看重系统的可用性而不是强一致性，它们对最终一致性也很满意。
- 松耦合（Loose coupling）
    仍然是一个重要的设计问题，涉及传入和传出的耦合。
- 单一责任（single responsibility）
    是一种思想，它支持对不太大或太细的微服务进行建模，因为他们包含了适当数量的内聚功能。

### LoD原则（Law of Demeter）
> Each unit should have only limited knowledge about other units: onlyunits "closely"related to the current unit.Only talk to your immediatefriends,don＇t talk to strangers.

> 每一个逻辑单元应该对其他逻辑单元有最少的了解：也就是说只亲近当前的对象。只和直接（亲近）的朋友说话，不和陌生人说话。

### KISS原则（Keep It Simple and Stupid）
- “简单”就是要让你的程序能简单、快速地被实现；
- “愚蠢”是说你的设计要简单到任何人都能理解，即简单就是美！

### DRY原则（Don't Repeat Yourself）
要遵循DRY原则，实现的方式非常多：

- 函数级别的封装：把一些经常使用的、重复出现的功能封装成一个通用的函数。
- 类级别的抽象：把具有相似功能或行为的类进行抽象，抽象出一个基类，并把这几个类都有的方法提到基类去实现。
- 泛型设计：Java 中可使用泛型，以实现通用功能类对多种数据类型的支持；C++中可以使用类模板的方式，或宏定义的方式；Python中可以使用装饰器来消除冗余的代码。

### YAGNI原则（You Aren＇t Gonna Need It）
> You aren＇t gonna need it,don＇t implement something until it isnecessary.
> 你没必要那么着急，不要给你的类实现过多的功能，直到你需要它的时候再去实现。

- 只考虑和设计必需的功能，避免过度设计。
- 只实现目前需要的功能，在以后需要更多功能时，可以再进行添加。
- 如无必要，勿增加复杂性。


### Rule Of Three原则

Rule of three 称为“三次法则”，指的是当某个功能第三次出现时，再进行抽象化，即事不过三，三则重构。

- 第一次实现一个功能时，就尽管大胆去做；
- 第二次做类似的功能设计时会产生反感，但是还得去做；
- 第三次还要实现类似的功能做同样的事情时，就应该去审视是否有必要做这些重复劳动了，这个时候就应该重构你的代码了，即把重复或相似功能的代码进行抽象，封装成一个通用的模块或接口。

### CQS原则（Command-Query Separation）

- 查询（Query）：当一个方法返回一个值来回应一个问题的时候，它就具有查询的性质；
- 命令（Command）：当一个方法要改变对象的状态的时候，它就具有命令的性质。