# GraalVM
[源码](https://github.com/oracle/graal)
[下载](https://github.com/graalvm/graalvm-ce-builds/releases)

## 安装
```
# 解压
tar -zxf graalvm-ce-java11-linux-amd64-19.3.0.2.tar.gz
# 设置环境变量(临时)
export PATH=/root/graalvm-ce-java11-19.3.0.2/bin:$PATH

# 设置环境变量(对当前登录用户)
vi ~/.bash_profile
# 将你要声明的环境变量加到PATH=$PATH:$HOME/bin这一行之后
# 修改完成之后要source ~/.bash_profile，环境变量才能生效。


# 设置环境变量(永久对所有系统用户生效)
vi  /etc/profile
    PATH=$PATH:*** //这里的***指的就是你要声明的路径
export PATH
注意，修改完成之后也要source /etc/profile，环境变量才能生效。


# 验证
java -version

js --version

npm -v

node -v
```
包管理工具gu
```
$ gu install native-image
$ gu install ruby
$ gu install python
$ gu install R
```
native-image命令会把你的Java代码以及用到的相关库，都编译成本地的机器代码

## JavaScript, Java, Ruby以及R混合编程

除了Java以外，GraalVM还提供了JavaScript, Ruby, R以及Python语言的实现。它们是基于一款新的语言实现框架Truffle来实现的，用它来实现语言的解释器非常简单，执行性能也很不错。使用Truffle来编写解释器时，它会自动使用GraalVM并为你提供了JIT编译的功能。因此GraalVM不仅仅是Java语言的JIT及ahead-of-time编译器，它也是JavaScript, Ruby, R以及Python等语言的JIT编译器。

## demo

HelloWorld.java
```java
public class HelloWorld{
    public static void main(String[] args){
        System.out.println("Hello World!");
    }
}
```
原生java命令
```
time java HelloWorld.java

[root@izj6c9hcysthc08ik6c89lz java]# time java HelloWorld.java 
Hello World!

real	0m0.719s
user	0m1.081s
sys	0m0.114s

```
native-image
```
yum install -y libz-dev
native-image HelloWorld

time ./helloworld
Hello World!

real	0m0.009s
user	0m0.005s
sys	0m0.002s
```
graalvm 提供的native 模式，可以加速应用的启动，不同可以让应用不再依赖jvm 运行时环境，但是
也有一些限制
[Native Image Java Limitations](https://github.com/oracle/graal/blob/master/substratevm/LIMITATIONS.md)

https://www.graalvm.org/docs/reference-manual/native-image/


https://github.com/kingreatwill/java_study/tree/master/GraalVM

Oracle pits GraalVM against Google Go
https://www.techworld.com.au/article/666629/oracle-pits-graalvm-against-google-go/

## 红帽和 GraalVM 社区创建 GraalVM 下游发行版“Mandrel”

红帽和 GraalVM 社区共同建立了新的 GraalVM 下游发行版，称为 Mandrel。红帽方面表示，该发行版本将为 [Quarkus](https://github.com/quarkusio/quarkus) 提供支持。Quarkus 已成为红帽运行时（Red Hat Runtimes）中全面支持的一个框架。


## 淘汰的gcj
GCJ 是GNU Compiler for the Java Programing Language 的简称。是Java版的GNU编译器。
GCJ 是一个轻巧的，性能优越的Java语言编译器。它能够将Java源文件编译为Java字节码文件或者直接将Java源文件编译为本地机器码，它也能够将Java字节码文件编译为本地机器码。
被编译的应用程序和GCJ运行时libgcj进行链接。该运行时提供了核心类库，垃圾回收器和一个字节码解释器。libgcj能够动态加载和解释类文件，产生混合编译/解释的应用程序。该运行时同名为GNU Classpath的工程整合在一起。目前支持的java版本最高1.5，最新的1.6还不支持。

GCJ作为GNU编译器集合的一部分已有十多年的历史，但效率低下一直是它的主要问题。直到OpenJDK出现后，GCJ的使用频率越来越低。到了2017年，它已不再被维护，未来不会再成为Linux发行版的一部分。

http://www.mingw.org/category/wiki/gcj