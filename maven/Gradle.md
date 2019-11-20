https://www.w3cschool.cn/gradle/ms7n1hu2.html
[TOC]

<< 操作符 是 doLast 的简写方式。
<< was deprecated in Gradle 4.x and removed in Gradle 5.0

# Gradle概述
Gradle是一个基于Apache Ant和Apache Maven概念的项目自动化构建工具。它使用一种基于Groovy的特定领域语言来声明项目设置，而不是传统的XML。Gradle就是工程的管理，帮我们做了依赖、打包、部署、发布、各种渠道的差异管理等工作。

# Gradle特性
#### 基于声明的构建和基于约定的构建
Gradle 的核心在于基于 Groovy 的丰富而可扩展的域描述语言(DSL)。 Groovy 通过声明性的语言元素将基于声明的构建推向下层，你可以按你想要的方式进行组合。 这些元素同样也为支持 Java， Groovy，OSGi，Web 和 Scala 项目提供了基于约定的构建。 并且，这种声明性的语言是可以扩展的。你可以添加新的或增强现有的语言元素。 因此，它提供了简明、可维护和易理解的构建。

#### 为以依赖为基础的编程方式提供语言支持
声明性语言优点在于通用任务图，你可以将其充分利用在构建中. 它提供了最大限度的灵活性，以让 Gradle 适应你的特殊需求。

#### 构建结构化
Gradle 的灵活和丰富性最终能够支持在你的构建中应用通用的设计模式。 例如，它可以很容易地将你的构建拆分为多个可重用的模块，最后再进行组装，但不要强制地进行模块的拆分。 不要把原本在一起的东西强行分开（比如在你的项目结构里），从而避免让你的构建变成一场噩梦。 最后，你可以创建一个结构良好，易于维护，易于理解的构建。

#### 深度 API
Gradle 允许你在构建执行的整个生命周期，对它的核心配置及执行行为进行监视并自定义。

#### Gradle 的扩展
Gradle 有非常良好的扩展性。 从简单的单项目构建，到庞大的多项目构建，它都能显著地提升你的效率。 这才是真正的结构化构建。通过最先进的增量构建功能，它可以解决许多大型企业所面临的性能瓶颈问题。

#### 多项目构建
Gradle 对多项目构建的支持非常出色。项目依赖是首先需要考虑的问题。 我们允许你在多项目构建当中对项目依赖关系进行建模，因为它们才是你真正的问题域。 Gradle 遵守你的布局。

Gradle 提供了局部构建的功能。 如果你在构建一个单独的子项目，Gradle 也会帮你构建它所依赖的所有子项目。 你也可以选择重新构建依赖于特定子项目的子项目。 这种增量构建将使得在大型构建任务中省下大量时间。

#### 多种方式管理依赖
不同的团队喜欢用不同的方式来管理他们的外部依赖。 从 Maven 和 Ivy 的远程仓库的传递依赖管理，到本地文件系统的 jar 包或目录，Gradle 对所有的管理策略都提供了方便的支持。

#### Gradle 是第一个构建集成工具
Ant tasks 是最重要的。而更有趣的是，Ant projects 也是最重要的。 Gradle 对任意的 Ant 项目提供了深度导入，并在运行时将 Ant 目标(target)转换为原生的 Gradle 任务(task)。 你可以从 Gradle 上依赖它们(Ant targets)，增强它们，甚至在你的 build.xml 上定义对 Gradle tasks 的依赖。Gradle 为属性、路径等等提供了同样的整合。

Gradle 完全支持用于发布或检索依赖的 Maven 或 Ivy 仓库。 Gradle 同样提供了一个转换器，用于将一个 Maven pom.xml 文件转换为一个 Gradle 脚本。Maven 项目的运行时导入的功能将很快会有。

#### 易于移植
Gradle 能适应你已有的任何结构。因此，你总可以在你构建项目的同一个分支当中开发你的 Gradle 构建脚本，并且它们能够并行进行。 我们通常建议编写测试，以保证生成的文件是一样的。 这种移植方式会尽可能的可靠和减少破坏性。这也是重构的最佳做法。

#### Groovy
Gradle 的构建脚本是采用 Groovy 写的，而不是用 XML。 但与其他方法不同，它并不只是展示了由一种动态语言编写的原始脚本的强大。 那样将导致维护构建变得很困难。 Gradle 的整体设计是面向被作为一门语言，而不是一个僵化的框架。 并且 Groovy 是我们允许你通过抽象的 Gradle 描述你个人的 story 的黏合剂。 Gradle 提供了一些标准通用的 story。这是我们相比其他声明性构建系统的主要特点。 我们的 Groovy 支持也不是简单的糖衣层，整个 Gradle 的 API 都是完全 groovy 化的。只有通过 Groovy才能去运用它并对它提高效率。

#### The Gradle wrapper
Gradle Wrapper 允许你在没有安装 Gradle 的机器上执行 Gradle 构建。 这一点是非常有用的。比如，对一些持续集成服务来说。 它对一个开源项目保持低门槛构建也是非常有用的。 Wrapper 对企业来说也很有用，它使得对客户端计算机零配置。 它强制使用指定的版本，以减少兼容支持问题。

#### 自由和开源
Gradle 是一个开源项目，并遵循 ASL 许可。

#### 为什么使用 Groovy?
我们认为内部 DSL（基于一种动态语言）相比 XML 在构建脚本方面优势非常大。它们是一对动态语言。 为什么使用 Groovy？答案在于 Gradle 内部的运行环境。 虽然 Gradle 核心目的是作为通用构建工具，但它还是主要面向 Java 项目。 这些项目的团队成员显然熟悉 Java。我们认为一个构建工具应该尽可能地对所有团队成员透明。

你可能会想说，为什么不能使用 Java 来作为构建脚本的语言。 我认为这是一个很有意义的问题。对你们的团队来讲，它确实会有最高的透明度和最低的学习曲线。 但由于 Java 本身的局限性，这种构建语言可能就不会那样友善、 富有表现力和强大。 [1] 这也是为什么像 Python，Groovy 或者 Ruby 这样的语言在这方面表现得更好的原因。 我们选择了 Groovy，因为它向 Java 人员提供了目前为止最大的透明度。 其基本的语法，类型，包结构和其他方面都与 Java 一样，Groovy 在这之上又增加了许多东西。但是和 Java 也有着共同点。

对于那些分享和乐于去学习 Python 知识的 Java 团队而言，上述论点并不适用。 Gradle 的设计非常适合在 JRuby 或 Jython 中创建另一个构建脚本引擎。 那时候，对我们而言，它只是不再是最高优先级的了。我们很高兴去支持任何社区努力创建其他的构建脚本引擎。
# Gradle优势
- 一款最新的，功能最强大的构建工具，用它逼格更高
- 使用程序代替传统的XML配置，项目构建更灵活
- 丰富的第三方插件，让你随心所欲使用
- Maven、Ant能做的，Gradle都能做，但是Gradle能做的，Maven、Ant不一定能做。

# Gradle 构建基础

### Projects 和 tasks

projects 和 tasks是 Gradle 中最重要的两个概念。

任何一个 Gradle 构建都是由一个或多个 projects 组成。每个 project 包括许多可构建组成部分。 这完全取决于你要构建些什么。举个例子，每个 project 或许是一个 jar 包或者一个 web 应用，它也可以是一个由许多其他项目中产生的 jar 构成的 zip 压缩包。一个 project 不必描述它只能进行构建操作。它也可以部署你的应用或搭建你的环境。不要担心它像听上去的那样庞大。 Gradle 的 build-by-convention 可以让您来具体定义一个 project 到底该做什么。

每个 project 都由多个 tasks 组成。每个 task 都代表了构建执行过程中的一个原子性操作。如编译，打包，生成 javadoc，发布到某个仓库等操作。

到目前为止，可以发现我们可以在一个 project 中定义一些简单任务，后续章节将会阐述多项目构建和多项目多任务的内容。

### Hello world

你可以通过在命令行运行 gradle 命令来执行构建，gradle 命令会从当前目录下寻找 build.gradle 文件来执行构建。我们称 build.gradle 文件为构建脚本。严格来说这其实是一个构建配置脚本，后面你会了解到这个构建脚本定义了一个 project 和一些默认的 task。

你可以创建如下脚本到 build.gradle 中 To try this out，create the following build script named build.gradle。

### 第一个构建脚本

build.gradle

    task hello {
        doLast {
            println 'Hello world!'
        }
    }

然后在该文件所在目录执行 gradle -q hello

> -q 参数的作用是什么?
> 
> 该文档的示例中很多地方在调用 gradle 命令时都加了 -q 参数。该参数用来控制 gradle 的日志级别，可以保证只输出我们需要的内容。具体可参阅本文档第十八章日志来了解更多参数和信息。

### 执行脚本

    Output of gradle -q hello
    > gradle -q hello
    Hello world!

上面的脚本定义了一个叫做 hello 的 task，并且给它添加了一个动作。当执行 gradle hello 的时候, Gralde 便会去调用 hello 这个任务来执行给定操作。这些操作其实就是一个用 groovy 书写的闭包。

如果你觉得它看上去跟 Ant 中的 targets 很像，那么是对的。Gradle 的 tasks 就相当于 Ant 中的 targets。不过你会发现他功能更加强大。我们只是换了一个比 target 更形象的另外一个术语。不幸的是这恰巧与 Ant 中的术语有些冲突。ant 命令中有诸如 javac、copy、tasks。所以当该文档中提及 tasks 时，除非特别指明 ant task。否则指的均是指 Gradle 中的 tasks。

## 快速定义任务


用一种更简洁的方式来定义上面的 hello 任务。

### 快速定义任务

build.gradle

    task hello << {
        println 'Hello world!'
    }

上面的脚本又一次采用闭包的方式来定义了一个叫做 hello 的任务，本文档后续章节中我们将会更多的采用这种风格来定义任务。

### 代码即脚本

Gradle 脚本采用 Groovy 书写，作为开胃菜,看下下面这个例子。

### 在 gradle 任务中采用 groovy

build.gradle

    task upper << {
        String someString = 'mY_nAmE'
        println "Original: " + someString
        println "Upper case: " + someString.toUpperCase()
    }
    Output of gradle -q upper
    > gradle -q upper
    Original: mY_nAmE
    Upper case: MY_NAME

或者

### 在 gradle 任务中采用 groovy

build.gradle

    task count << {
        4.times { print "$it " }
    }
    Output of gradle -q count
    > gradle -q count
    0 1 2 3

## 任务依赖


你可以按如下方式创建任务间的依赖关系

### 在两个任务之间指明依赖关系

build.gradle

    task hello << {
        println 'Hello world!'
    }
    task intro(dependsOn: hello) << {
        println "I'm Gradle"
    }

gradle -q intro 的输出结果

    Output of gradle -q intro
    \> gradle -q intro
    Hello world!
    I'm Gradle

添加依赖 task 也可以不必首先声明被依赖的 task。

## 延迟依赖

build.gradle

    task taskX(dependsOn: 'taskY') << {
        println 'taskX'
    }
    task taskY << {
        println 'taskY'
    }

Output of gradle -q taskX

     \> gradle -q taskX
    taskY
    taskX

可以看到，taskX 是 在 taskY 之前定义的，这在多项目构建中非常有用。

注意:当引用的任务尚未定义的时候不可使用短标记法来运行任务。

## 动态任务

借助 Groovy 的强大不仅可以定义简单任务还能做更多的事。例如，可以动态定义任务。

### 创建动态任务

build.gradle

    4.times { counter ->
        task "task$counter" << {
            println "I'm task number $counter"
        }
    }

gradle -q task1 的输出结果。

    Output of gradle -q task1
    \> gradle -q task1
    I'm task number 1

## 任务操纵


一旦任务被创建后，任务之间可以通过 API 进行相互访问。这也是与 Ant 的不同之处。比如可以增加一些依赖。

### 通过 API 进行任务之间的通信 - 增加依赖

build.gradle

    4.times { counter ->
        task "task$counter" << {
            println "I'm task number $counter"
        }
    }
    task0.dependsOn task2, task3

gradle -q task0的输出结果。

    Output of gradle -q task0
    \> gradle -q task0
    I'm task number 2
    I'm task number 3
    I'm task number 0

为已存在的任务增加行为。

### 通过 API 进行任务之间的通信 - 增加任务行为

build.gradle

    task hello << {
        println 'Hello Earth'
    }
    hello.doFirst {
        println 'Hello Venus'
    }
    hello.doLast {
        println 'Hello Mars'
    }
    hello << {
        println 'Hello Jupiter'
    }
    Output of gradle -q hello
    > gradle -q hello
    Hello Venus
    Hello Earth
    Hello Mars
    Hello Jupiter

doFirst 和 doLast 可以进行多次调用。他们分别被添加在任务的开头和结尾。当任务开始执行时这些动作会按照既定顺序进行。其中 << 操作符 是 doLast 的简写方式。

## 短标记法


你早就注意到了吧，没错，每个任务都是一个脚本的属性，你可以访问它:

### 以属性的方式访问任务

build.gradle

    task hello << {
        println 'Hello world!'
    }
    hello.doLast {
        println "Greetings from the $hello.name task."
    }

gradle -q hello 的输出结果

    Output of gradle -q hello
    \> gradle -q hello
    Hello world!
    Greetings from the hello task.

对于插件提供的内置任务。这尤其方便(例如:complie)

## 增加自定义属性


你可以为一个任务添加额外的属性。例如,新增一个叫做 myProperty 的属性，用 ext.myProperty 的方式给他一个初始值。这样便增加了一个自定义属性。

### 为任务增加自定义属性

build.gradle

    task myTask {
        ext.myProperty = "myValue"
    }
    
    task printTaskProperties << {
        println myTask.myProperty
    }

gradle -q printTaskProperties 的输出结果

    Output of gradle -q printTaskProperties
    \> gradle -q printTaskProperties
    myValue

自定义属性不仅仅局限于任务上，还可以做更多事情。

## 调用 Ant 任务


Ant 任务是 Gradle 中的一等公民。Gradle 借助 Groovy 对 Ant 任务进行了优秀的整合。Gradle 自带了一个 AntBuilder，在 Gradle 中调用 Ant 任务比在 build.xml 中调用更加的方便和强大。 通过下面的例子你可以学到如何调用一个 Ant 任务以及如何与 Ant 中的属性进行通信。

### 利用 AntBuilder 执行 ant.loadfile

build.gradle

    task loadfile << {
        def files = file('../antLoadfileResources').listFiles().sort()
        files.each { File file ->
            if (file.isFile()) {
                ant.loadfile(srcFile: file, property: file.name)
                println " *** $file.name ***"
                println "${ant.properties[file.name]}"
            }
        }
    }

gradle -q loadfile 的输出结果

    Output of gradle -q loadfile
    \> gradle -q loadfile
    *** agile.manifesto.txt ***
    Individuals and interactions over processes and tools
    Working software over comprehensive documentation
    Customer collaboration  over contract negotiation
    Responding to change over following a plan
     *** gradle.manifesto.txt ***
    Make the impossible possible, make the possible easy and make the easy elegant.
    (inspired by Moshe Feldenkrais)

在你脚本里还可以利用 Ant 做更多的事情。想了解更多请参阅在 Gradle 中调用 Ant。

## 方法抽取


Gradle 的强大要看你如何编写脚本逻辑。针对上面的例子，首先要做的就是要抽取方法。

### 利用方法组织脚本逻辑

build.gradle

    task checksum << {
        fileList('../antLoadfileResources').each {File file ->
            ant.checksum(file: file, property: "cs_$file.name")
            println "$file.name Checksum: ${ant.properties["cs_$file.name"]}"
        }
    }
    task loadfile << {
        fileList('../antLoadfileResources').each {File file ->
            ant.loadfile(srcFile: file, property: file.name)
            println "I'm fond of $file.name"
        }
    }
    File[] fileList(String dir) {
        file(dir).listFiles({file -> file.isFile() } as FileFilter).sort()
    }

gradle -q loadfile 的输出结果

    Output of gradle -q loadfile
    \> gradle -q loadfile
    I'm fond of agile.manifesto.txt
    I'm fond of gradle.manifesto.txt

在后面的章节你会看到类似出去出来的方法可以在多项目构建中的子项目中调用。无论构建逻辑多复杂，Gradle 都可以提供给你一种简便的方式来组织它们。

## 定义默认任务

Gradle 允许在脚本中定义多个默认任务。

## 定义默认任务

build.gradle

    defaultTasks 'clean', 'run'
    task clean << {
        println 'Default Cleaning!'
    }
    task run << {
        println 'Default Running!'
    }
    task other << {
        println "I'm not a default task!"
    }

gradle -q 的输出结果。

    Output of gradle -q
    \> gradle -q
    Default Cleaning!
    Default Running!

这与直接调用 gradle clean run 效果是一样的。在多项目构建中，每个子项目都可以指定单独的默认任务。如果子项目未进行指定将会调用父项目指定的的默认任务。

## Configure by DAG

稍后会对 Gradle 的配置阶段和运行阶段进行详细说明 配置阶段后，Gradle 会了解所有要执行的任务 Gradle 提供了一个钩子来捕获这些信息。一个例子就是可以检查已经执行的任务中有没有被释放。借由此，你可以为一些变量赋予不同的值。

在下面的例子中，为 distribution 和 release 任务赋予了不同的 version 值。

### 依赖任务的不同输出

build.gradle

    task distribution << {
        println "We build the zip with version=$version"
    }
    task release(dependsOn: 'distribution') << {
        println 'We release now'
    }
    gradle.taskGraph.whenReady {taskGraph ->
        if (taskGraph.hasTask(release)) {
            version = '1.0'
        } else {
            version = '1.0-SNAPSHOT'
        }
    }

gradle -q distribution 的输出结果

    Output of gradle -q distribution
    \> gradle -q distribution
    We build the zip with version=1.0-SNAPSHOT

gradle -q release 的输出结果

    Output of gradle -q release
    \> gradle -q release
    We build the zip with version=1.0
    We release now

whenReady 会在已发布的任务之前影响到已发布任务的执行。即使已发布的任务不是主要任务(也就是说，即使这个任务不是通过命令行直接调用)

## 下一步目标

在本章中，我们了解了什么是 task，但这还不够详细。欲知更多请参阅章节[任务详述](https://www.w3cschool.cn/gradle/y4lo1hu0.html)。

另外，可以目录继续学习[Java 构建入门](https://www.w3cschool.cn/gradle/9b5m1htc.html)和[依赖管理基础](https://www.w3cschool.cn/gradle/sh8k1htf.html)。
