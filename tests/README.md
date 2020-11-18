<!--toc-->
[TOC]
[测试平台](https://www.processon.com/diagraming/5ddddca0e4b074c442e605de)

UFT，QTP被惠普收购以后的新名称

## 自动化测试工具
- Endtest 收费
  Recorder 组件可以让用户不用编程技巧就能创建和执行测试
- 基于selenium的watir
  Watir 是一个 Ruby 的浏览器自动化测试开源库。
  https://github.com/watir/watir
- Sikuli
- Micro Focus UFT (QTP)
- IBM Rational Functional Tester
- Jest
- Cucumber
  Cucumber 是一个开源的行为驱动测试工具，支持多种编程语言，包括 Ruby,Java，Scala 和 Groovy。可以说 Cucumber 并不是 Selenium 的替代品，而仅仅是对 Selenium 进行了一层包装。
- ReadyAPI
- SoapUI 

### UI自动化录制工具 chrome 插件
#### headless-recorder 推荐
Puppeteer or Playwright script.
https://github.com/checkly/headless-recorder

> 注意元素规范：https://theheadless.dev/posts/basics-selectors/


#### Cypress-Recorder 推荐
https://github.com/KabaLabs/Cypress-Recorder

#### katalon-recorder  推荐
selenium

https://github.com/katalon-studio/katalon-recorder
https://docs.katalon.com

#### Selenium IDE
seleniumhq.org
https://github.com/SeleniumHQ/selenium-ide

#### BlazeMeter | The Continuous Testing Platform

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

#### Selenium下Chrome配置

常用的行为有：
禁止图片和视频的加载：提升网页加载速度。
添加代理：用于翻墙访问某些页面，或者应对IP访问频率限制的反爬技术。
使用移动头：访问移动端的站点，一般这种站点的反爬技术比较薄弱。
添加扩展：像正常使用浏览器一样的功能。
设置编码：应对中文站，防止乱码。
阻止JavaScript执行。

例子: 设置无界面模式浏览器启动
chrome_options = webdriver.ChromeOptions()
chrome_options.add_argument('--headless')
driver = webdriver.Chrome(chrome_options=chrome_options)

##### 设置默认编码为 utf-8，也就是中文
options.add_argument('lang=zh_CN.UTF-8')

移动设备user-agent表格：http://www.fynas.com/ua
##### 通过设置user-agent，用来模拟移动设备
###### 比如模拟 android QQ浏览器
options.add_argument('user-agent="MQQBrowser/26 Mozilla/5.0 (Linux; U; Android 2.3.7; zh-cn; MB200 Build/GRJ22; CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"')

###### 模拟iPhone 6
options.add_argument('user-agent="Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"')

##### 设置好应用扩展
extension_path = 'D:/extension/XPath-Helper_v2.0.2.crx'
chrome_options.add_extension(extension_path)
如何获取crx文件？http://crxextractor.com/ 或者 在扩展里 点击“打包扩展程序”

```
地址：https://peter.sh/experiments/chromium-command-line-switches/

chrome_options.add_argument('--headless') # 无头模式
chrome_options.add_argument('--disable-gpu') # 禁用GPU加速
chrome_options.add_argument('--start-maximized')#浏览器最大化
chrome_options.add_argument('--window-size=1280x1024') # 设置浏览器分辨率（窗口大小）


chrome_options.add_argument('log-level=3') # options.SetLoggingPreference("off", LogLevel.Off); #禁止 console.log 日志输出
#info(default) = 0
#warning = 1
#LOG_ERROR = 2
#LOG_FATAL = 3

chrome_options.add_argument("url=data:,");# 设置启动浏览器空白页
chrome_options.add_argument('--user-agent=""') # 设置请求头的User-Agent
chrome_options.add_argument('--disable-infobars') # 禁用浏览器正在被自动化程序控制的提示
chrome_options.add_argument('--incognito') # 隐身模式（无痕模式）
chrome_options.add_argument('--hide-scrollbars') # 隐藏滚动条, 应对一些特殊页面
chrome_options.add_argument('--disable-javascript') # 禁用javascript
chrome_options.add_argument('--blink-settings=imagesEnabled=false') # 不加载图片, 提升速度
chrome_options.add_argument('--ignore-certificate-errors') # 禁用扩展插件并实现窗口最大化
chrome_options.add_argument('–disable-software-rasterizer')
chrome_options.add_argument('--disable-extensions') # 禁用Chrome浏览器上现有的扩展
chrome_options.add_argument('--disable-popup-blocking') #禁用弹窗

其他
–user-data-dir=”[PATH]” 指定用户文件夹User Data路径，可以把书签这样的用户数据保存在系统分区以外的分区。
–disk-cache-dir=”[PATH]“ 指定缓存Cache路径
–disk-cache-size= 指定Cache大小，单位Byte
–first run 重置到初始状态，第一次运行
–incognito 隐身模式启动
–disable-javascript 禁用Javascript
–omnibox-popup-count=”num” 将地址栏弹出的提示菜单数量改为num个。我都改为15个了。
–user-agent=”xxxxxxxx” 修改HTTP请求头部的Agent字符串，可以通过about:version页面查看修改效果
–disable-plugins 禁止加载所有插件，可以增加速度。可以通过about:plugins页面查看效果
–disable-javascript 禁用JavaScript，如果觉得速度慢在加上这个
–disable-java 禁用java
–start-maximized 启动就最大化
–no-sandbox 取消沙盒模式
–single-process 单进程运行
–process-per-tab 每个标签使用单独进程
–process-per-site 每个站点使用单独进程
–in-process-plugins 插件不启用单独进程
–disable-popup-blocking 禁用弹出拦截
–disable-plugins 禁用插件
–disable-images 禁用图像
–incognito 启动进入隐身模式
–enable-udd-profiles 启用账户切换菜单
–proxy-pac-url 使用pac代理 [via 1/2]
–lang=zh-CN 设置语言为简体中文
–disk-cache-dir 自定义缓存目录
–disk-cache-size 自定义缓存最大值（单位byte）
–media-cache-size 自定义多媒体缓存最大值（单位byte）
–bookmark-menu 在工具 栏增加一个书签按钮
–enable-sync 启用书签同步
–single-process 单进程运行Google Chrome
–start-maximized 启动Google Chrome就最大化
–disable-java 禁止Java
–no-sandbox 非沙盒模式运行
```

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

#### WinAppDriver WIN10 UI自动化
http://appium.io/docs/en/drivers/windows/
https://zhuanlan.zhihu.com/p/63511229

元素可以通过inspect.exe获取




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

#### pywinauto  PC自动化


### Robot Framework
https://github.com/robotframework/robotframework

RIDE  
RIDE是一款专门用来编辑Robot Framework用例的软件，用Python编写并且开源
https://github.com/robotframework/RIDE

[robotframework中RIDE的下载及安装](https://www.cnblogs.com/jiyanjiao-702521/p/9220358.html)


[RobotFramework + Appium 移动自动化实现](https://www.cnblogs.com/leozhanggg/p/9687398.html)
https://github.com/serhatbolsu/robotframework-appiumlibrary


[Robot Framework + SeleniumLibrary 实现web 自动化](https://github.com/robotframework/SeleniumLibrary)

### Cypress
[Cypress](https://docs.cypress.io/)是基于node.js环境的,相较于Selenium来说，Cypress的运行速度真的是相当快了，它不需要web driver来驱动浏览器。

[Cypress系列](https://www.cnblogs.com/poloyy/category/1768839.html)

https://github.com/cypress-io/cypress

### Playwright UI自动化
https://github.com/microsoft/playwright
https://github.com/microsoft/playwright-python

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
Bandit是一个用于发现Python代码中常见安全问题的工具。

### 11款常用的安全测试工具
https://blog.csdn.net/lb245557472/article/details/88572607/
- AppScan
- Burp Suite
- Acunetix
- Nmap
- sqlmap
- Grabber 
  一个 Web 应用程序扫描程序，现在，它还可以做安全测试了。

### Litmus
Litmus 是一款用来测试和监视电子邮件的工具

### 强大的全新 Web UI 测试框架 Cypress

### 其它(以下跟前端更相关)
#### playwright
https://github.com/microsoft/playwright
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

## 性能工具
### 性能工具之Taurus
https://github.com/Blazemeter/taurus
支持
- JMeter Executor
- Selenium Executor
- JUnit Executor
- TestNG Executor
- Apiritif Executor
- PyTest Executor
- RSpec Executor
- Mocha Executor
- WebdriverIO Executor
- NUnit Executor
- xUnit Executor
- Robot Executor
- Gatling Executor
- Grinder Executor
- Locust Executor
- PBench Executor
- Siege Executor
- ApacheBenchmark Executor
- Tsung Executor
- Molotov Executor
- Postman/Newman Executor
- Existing Results Loader

## 安全测试
### SQL注入
[Sqlmap](https://github.com/sqlmapproject/sqlmap)
[NoSQLMap](https://github.com/codingo/NoSQLMap)
### Fuzz测试
什么是Fuzz测试？
漏洞挖掘有三种方法：白盒代码审计、灰盒逆向工程、黑盒测试。其中黑盒的Fuzz测试是效率最高的一种，能够快速验证大量潜在的安全威胁。

Fuzz测试，也叫做“模糊测试”，是一种挖掘软件安全漏洞、检测软件健壮性的黑盒测试，它通过向软件输入非法的字段，观测被测试软件是否异常而实现。Fuzz测试的概念非常容易理解，如果我们构造非法的报文并且通过测试工具打入被测设备，那么这就是一个Fuzz测试的测试例执行，大多数测试工程师肯定都尝试过这种测试手段。

对于网络协议漏洞挖掘来说，Fuzz测试也就意味着打入各种异常报文，然后观察设备是否有异常。

#### Peach是一个优秀的开源Fuzz框架。
Fuzz（模糊测试）是一种通过提供非预期的输入并监视异常结果来发现软件安全漏洞的方法。模糊测试在很大程度上是一种强制性的技术，简单并且有效，但测试存在盲目性。

典型地模糊测试过程是通过自动的或半自动的方法，反复驱动目标软件运行并为其提供构造的输入数据，同时监控软件运行的异常结果。

Fuzz被认为是一种简单有效的黑盒测试，随着Smart Fuzz的发展，RCE（逆向代码工程）需求的增加，其特征更符合一种灰盒测试。

https://www.peach.tech/

#### AFL、LibFuzzer、honggfuzz

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