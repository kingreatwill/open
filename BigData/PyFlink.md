
[视频](https://ververica.cn/developers/flink-training-course-basics/)

[flink系列-11、PyFlink 核心功能介绍（整理自 Flink 中文社区）](https://www.cnblogs.com/xiexiandong/p/12878642.html)

Python虚拟机（PyVM）和Java虚拟机（JVM）之间建立握手有两种解决方案：Beam和Py4J

使用Apache Beam来实现VM通信有点复杂。简而言之，这是因为Apache Beam专注于通用性，在极端情况下缺乏灵活性。

Flink（Spark）还需要交互式编程。此外，为了使Flink（Spark）正常工作，我们还需要确保其API设计中的语义一致性，尤其是在其多语言支持方面。Apache Beam的现有体系结构无法满足这些要求，因此答案很明显，Py4J是支持PyVM和JVM之间通信的最佳选择。

 .|Beam|Py4J
---|---|---
interactive programming FLIP-36（交互式编程）|不支持|支持
Java call Python | 不支持|支持
Align Flink Java API | 不支持|支持


[Python Support for UDFs in Flink 1.10](https://flink.apache.org/2020/04/09/pyflink-udf-support-flink.html)

![](img/pyflink-udf-architecture.gif)