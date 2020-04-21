
[测试平台](https://www.processon.com/diagraming/5ddddca0e4b074c442e605de)

UFT，QTP被惠普收购以后的新名称


### selenium
https://github.com/SeleniumHQ/selenium
https://www.selenium.dev/documentation/en/

selenium 是一个 web 的自动化测试工具，不少学习功能自动化的同学开始首选 selenium ，因为它相比 QTP 有诸多有点：

- 免费，也不用再为破解 QTP 而大伤脑筋
- 小巧，对于不同的语言它只是一个包而已，而 QTP 需要下载安装1个多 G 的程序。
- 这也是最重要的一点，不管你以前更熟悉 C、 java、ruby、python、或都是 C# ，你都可以通过 selenium 完成自动化测试，而 QTP 只支持 VBS
- 支持多平台：windows、linux、MAC ，支持多浏览器：ie、ff、safari、opera、chrome
- 支持分布式测试用例的执行，可以把测试用例分布到不同的测试机器的执行，相当于分发机的功能。


#### Selenium Grid
https://blog.csdn.net/ouyanggengcheng/article/details/79935657

这个时候WebDriver 就要使用RemoteWebDriver了

https://testerhome.com/topics/20192

![](img/selenium-grid.png)


https://www.selenium.dev/documentation/en/grid/setting_up_your_own_grid/
https://github.com/SeleniumHQ/selenium/blob/selenium-3.141.59/java/server/src/org/openqa/grid/common/defaults/DefaultNodeWebDriver.json


node也可以跑在docker上
https://github.com/SeleniumHQ/docker-selenium


#### Selenium Python

http://www.testclass.net/selenium_python

#### Selenium并行启动多个浏览器
https://www.cnblogs.com/graceting/p/5034023.html


### Pytest 5.7k
pytest是一个非常成熟的全功能的Python测试框架，主要特点有以下几点：
https://github.com/pytest-dev/pytest
1、简单灵活，容易上手，文档丰富；
2、支持参数化，可以细粒度地控制要测试的测试用例；
3、能够支持简单的单元测试和复杂的功能测试，还可以用来做selenium/appnium等自动化测试、接口自动化测试（pytest+requests）;
4、pytest具有很多第三方插件，并且可以自定义扩展，比较好用的如pytest-selenium（集成selenium）、pytest-html（完美html测试报告生成）、pytest-rerunfailures（失败case重复执行）、pytest-xdist（多CPU分发）等；
5、测试用例的skip和xfail处理；
6、可以很好的和CI工具结合，例如jenkins

```
pip install pytest pytest-xdist

$ pytest test/unit

$ pytest -n 2 test/unit

$ pytest test/functional/ios/find_by_ios_class_chain_tests.py
```

### unittest
官方自带包

https://docs.python.org/3/library/unittest.html


#### Run tests
```
pip install tox
```
tox


#### Release 
使用twine包发布模块
twine

### SeleniumBase
https://github.com/seleniumbase/SeleniumBase

Selenium & pytest


### appium
[appium介绍](http://www.testclass.net/appium/appium-base-summary/)

appium的核心其实是一个暴露了一系列REST API的server。

这个server的功能其实很简单：监听一个端口，然后接收由client发送来的command。翻译这些command，把这些command转成移动设备可以理解的形式发送给移动设备，然后移动设备执行完这些command后把执行结果返回给appium server，appium server再把执行结果返回给client。

这样的设计思想带来了一些好处：
1，可以带来多语言的支持；
2，可以把server放在任意机器上，哪怕是云服务器都可以；（是的，appium和webdriver天生适合云测试）

检查安装环境
npm install -g appium-doctor
appium-doctor  --android or --ios


安装Android Studio
https://developer.android.google.cn/studio/
环境变量
ANDROID_HOME 
PATH   %ANDROID_HOME%/tools;%ANDROID_HOME%/platform-tools
adb.exe

ANDROID_SDK_ROOT  和 ANDROID_HOME 设置成一样的,tools可以拷贝到设置的地方，
Android SDK Command-line Tools（支持jdk11） 可以覆盖过去，

安装Appium Server -> appium-desktop(or npm install -g appium  运行appium) 
https://github.com/appium/appium-desktop/releases
Appium Server与设备要一一对应



Failed to install android-sdk: “java.lang.NoClassDefFoundError: javax/xml/bind/annotation/XmlSchema”
https://stackoverflow.com/questions/46402772/failed-to-install-android-sdk-java-lang-noclassdeffounderror-javax-xml-bind-a


Android SDK Command-line Tools
https://developer.android.com/studio/command-line/sdkmanager
https://dl.google.com/android/repository/commandlinetools-win-6200805_latest.zip?hl=ru



appium 通过 uiautomatorviewer.bat 工具来查看控件的属性。该工具位于 Android SDK 的 /tools/bin/ 目录下。

或者：Appium Inspector

#### grid
说起grid，了解selenium的人肯定知道，他就是分布式的核心。原理是简历中心hub，然后配置node，在hub上运行服务时，会去node上执行相关操作，类似于Jenkins上的节点操作。

在一个server上搭建一台Selenium Hub，多台Appium Server连接到Selenium Hub，测试真机（或者模拟器）和Appium Server连接。
测试脚本通过WebDriver JSONWireProtocol发送到Selenium Hub上，然后由Selenium Hub自动分发到对应的空闲Appium Server上进行执行。

那么appium如何搭建grid环境呢，其实和selenium类似，首先搭建hub：


##### 一、搭建hub
首先下载selenium-server-standalone-<version>.jar文件，地址：http://selenium-release.storage.googleapis.com/index.html。这里使用的是selenium-server-standalone-3.4.0.jar

yum install java-1.8.0-openjdk* -y

下载完成后直接在jar目录下运行：java -jar selenium-server-standalone-3.4.0.jar -p 4444 -role hub。
```
java -jar selenium-server-standalone-2.53.1.jar -role hub -maxSession 10 -port 9999

l -role hub表示启动运行hub；

l -port是设置端口号，hub的默认端口是4444，这里使用的是默认的端口，当然可以自己配置；

l -maxSession为最大会话请求，这个参数主要要用并发执行测试用例，默认是1，建议设置10及以上。

```

访问：127.0.0.1:4444/grid/console

##### 二、启动node

首先新建test.json文件，内容如下：
```
{
  "capabilities": [
    {
      "deviceName": "test",
      "version": "4.4.2",
      "maxInstances": 3,
      "platform": "ANDROID",
      "browserName": "chrome"
    }
  ],
  "configuration":
    {
        "cleanUpCycle":"2000",
        "timeout":"30000",
        "proxy":"org.openqa.grid.selenium.proxy.DefaultRemoteProxy",
        "url":"http://127.0.0.1:4723/wd/hub",  //appiumserver地址,即node地址
        "host":"127.0.0.1",
        "port":"4723",           //节点机服务端口
        "maxSession":"1",
        "register":true,
        "registerCycle":"5000",
        "hubPort":"4444",     //hub端口
        "hubHost":"192.168.4.41"      //hub地址
    }
}
```
然后执行命令：appium -p 4723 -bp 4724 --nodeconfig test.json。

在hub的启动服务后台能看到注册信息


#### appium-studio
https://experitest.com/mobile-test-automation/appium-studio/


### Robot Framework
https://github.com/robotframework/robotframework

RIDE  
RIDE是一款专门用来编辑Robot Framework用例的软件，用Python编写并且开源
https://github.com/robotframework/RIDE

[robotframework中RIDE的下载及安装](https://www.cnblogs.com/jiyanjiao-702521/p/9220358.html)


[RobotFramework + Appium 移动自动化实现](https://www.cnblogs.com/leozhanggg/p/9687398.html)
https://github.com/serhatbolsu/robotframework-appiumlibrary


[Robot Framework + SeleniumLibrary 实现web 自动化](https://github.com/robotframework/SeleniumLibrary)


### HttpRunner 
HttpRunner 是一款面向 HTTP(S) 协议的通用测试框架，只需编写维护一份 YAML/JSON 脚本，即可实现自动化测试、性能测试、线上监控、持续集成等多种测试需求。
### splinter 
https://github.com/cobrateam/splinter

### airtest
UI自动化测试工具


### Burp Suite

### Jmeter+ant+Jenkins
看最终报告
### Grafana+Jmeter+Influxdb
看压测过程中参数的变化

### bandit

https://github.com/PyCQA/bandit
### 11款常用的安全测试工具
https://blog.csdn.net/lb245557472/article/details/88572607/
- AppScan
- Burp Suite
- Acunetix
- Nmap
- sqlmap

### Python安全测试工具合集
https://www.cnblogs.com/xiaodi914/p/5176094.html


https://github.com/topics/security-scanner?l=python


本文列出13个python网络测试工具，共大家参考学习。

1、Scapy，Scapy3k：发送，嗅探和剖析并伪造网络数据包，可以做交互式应用或单纯的作为库来使用。

2、pypcap，pcapy和pylibpcap：几个不同的libpcap捆绑Python库

3、libdnet：低级别的网络路由器，可用于接口查找和以太网帧转发

4、dpkt：快速、轻量级的数据包创建、解析工具，适用于基本TCP/IP协议

5、Impacket：探测和解码网络数据包，支持更高级别协议比如NMB和SMB

6、pynids：libnids封装提供嗅探，IP碎片整理，TCP流重组和端口扫描检测

7、Dirtbags py-pcap：无需libpcap即可读取pcap文件

8、flowgrep：通过正则表达式查找数据包中的Payloads

9、Knock Subdomain Scan：通过字典枚举目标域上的子域名

10、SubBrute：可扩展的TCP/UDP中间代理，支持即时修改非标准协议

11、Pytbull：灵活的IDS/IPS测试框架（配有300多个测试用例）

12、Spoodle：大量子域名+Poodle漏洞扫描器

13、SMBMap：枚举域中的Samba共享驱动器


### 其它(以下跟前端更相关)

#### puppeteer
https://github.com/puppeteer/puppeteer
Headless Chrome Node.js API https://pptr.dev/

#### storybook
https://github.com/storybookjs/storybook
UI component dev & test: React, Vue, Angular, React Native, Ember, Web Components & more! https://storybook.js.org
https://www.jianshu.com/p/c9ddb5e5bfc2

#### Jest
https://github.com/facebook/jest
Delightful JavaScript Testing. https://jestjs.io