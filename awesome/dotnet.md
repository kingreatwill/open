[TOC]
## 解析器构建库(正则解析)
[Sprache](https://github.com/sprache/Sprache)
[superpower](https://github.com/datalust/superpower)
[Pidgin](https://github.com/benjamin-hodgson/Pidgin) 作者是Stackoverflow的员工,做为Sprache的后继者,从性能和功能上有一些改进,但是可能诞生较晚,知名度不如Sprache.
[fparsec](https://github.com/stephan-tolksdorf/fparsec) F#

## Stateless 状态机
https://github.com/dotnet-state-machine/stateless

## 内存
https://github.com/ZenLulz/MemorySharp

## 图片处理
https://github.com/SixLabors/ImageSharp
https://github.com/JimBobSquarePants/ImageProcessor
https://github.com/mono/SkiaSharp
https://github.com/kleisauke/net-vips

## 定时任务
- Coravel
Coravel是一款框架轻，使用简单，支持秒级定时任务。

官网地址：https://docs.coravel.net/

- FluentScheduler
Github地址：https://github.com/fluentscheduler/FluentScheduler

文档地址：https://fluentscheduler.github.io/

- Hangfire

- 其它

Timer,FluentScheduler,TaskScheduler,Gofer.NET,Coravel,Quartz.NET还是Hangfire
https://github.com/2881099/FreeScheduler


## javascript
Javascript Interpreter js解释器
https://github.com/sebastienros/jint

go  的js解释器
https://github.com/robertkrimen/otto

### 合并压缩资源文件
https://github.com/madskristensen/BundlerMinifier
https://www.cnblogs.com/zglblogs/p/13127578.html

## Headless Browser

前端就有了对 headless 浏览器的需求，最多的应用场景有两个

- UI 自动化测试：摆脱手工浏览点击页面确认功能模式
- 爬虫：解决页面内容异步加载等问题

也就有了很多杰出的实现，前端经常使用的莫过于 PhantomJS 和 selenium-webdriver，但两个库有一个共性——难用！环境安装复杂，API 调用不友好，1027 年 Chrome 团队连续放了两个大招 Headless Chrome 和对应的 NodeJS API Puppeteer，直接让 PhantomJS 和 Selenium IDE for Firefox 作者宣布没必要继续维护其产品.

js和c#,python的[selenium-webdriver](https://www.selenium.dev/selenium/docs/api/javascript/)
js的[PhantomJS ](https://phantomjs.org/)

[ Headless Chrome](https://chromium.googlesource.com/chromium/src/+/lkgr/headless/README.md)
[ NodeJS API Puppeteer](https://github.com/puppeteer/puppeteer)
[unofficial port of puppeteer](https://github.com/pyppeteer/pyppeteer)
[puppeteer-sharp](https://github.com/hardkoded/puppeteer-sharp)
[PuppeteerSharp 扩展](https://github.com/hlaueriksson/puppeteer-sharp-contrib)

## PDF
iTextSharp、PDFsharp、Aspose.PDF

> PDF转markdown  https://github.com/datalab-to/marker

### html2pdf/转PDF
[wkhtmltopdf](https://wkhtmltopdf.org/)

[wkhtmltopdf项目开源地址](https://github.com/wkhtmltopdf/wkhtmltopdf)

[WkHtmlToPdf-DotNet开源地址](https://github.com/HakanL/WkHtmlToPdf-DotNet)

[.NET下HTML动态转PDF报告的技术实践](https://mp.weixin.qq.com/s/ZjlVb_ev5Aa1vLvD33a8-w)

### QuestPDF
https://github.com/QuestPDF/QuestPDF

## Colorful Console

[Colorful Console](https://github.com/tomakita/Colorful.Console)
http://colorfulconsole.com/
Colorful.Console还允许我们使用FIGlet字体编写带颜色的ASCII码输出
FIGLet: http://www.figlet.org/


[ConsoleTables](https://github.com/khalidabuhakmeh/ConsoleTables)

[ShellProgressBar](https://github.com/Mpdreamz/shellprogressbar)

[GUI.CS](https://github.com/migueldeicaza/gui.cs)


## 组织图
https://github.com/structurizr/dotnet


## dotnetbook
https://github.com/sidristij/dotnetbook

## dumps
Microsoft.Diagnostics.Runtime is a set of APIs for introspecting processes and dumps.

https://github.com/microsoft/clrmd


### 在.NET程序崩溃时自动创建Dump
#### Windows 平台
对所有程序都有效果，不仅仅是.NET 程序
- 打开`regedit.exe`
- 打开目录`HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\Windows Error Reporting\LocalDumps`
- 创建 KEY `DumpFolder` 类型为 `REG_EXPAND_SZ` 用于配置存放 Dump 文件的目录
- 另外可以创建 KEY `DumpCount` 类型为 `REG_DWORD` 配置 Dump 的总数量

```PowerShell
New-Item -Path "HKLM:\SOFTWARE\Microsoft\Windows\Windows Error Reporting" -Name "LocalDumps"

New-ItemProperty -Path "HKLM:\SOFTWARE\Microsoft\Windows\Windows Error Reporting\LocalDumps" -Name "DumpFolder" -Value "%LOCALAPPDATA%\CrashDumps" -PropertyType ExpandString

New-ItemProperty -Path "HKLM:\SOFTWARE\Microsoft\Windows\Windows Error Reporting\LocalDumps" -Name "DumpCount" -Value 10 -PropertyType DWord
```
按照上面的配置，如果程序发生了异常退出，那么就会在`%LOCALAPPDATA%\CrashDumps`目录创建程序的 Dump。

#### .NET Core 全平台
https://github.com/dotnet/runtime/blob/main/docs/design/coreclr/botr/xplat-minidump-generation.md
https://learn.microsoft.com/en-us/troubleshoot/developer/webapps/aspnetcore/practice-troubleshoot-linux/lab-1-3-capture-core-crash-dumps

- COMPlus_DbgEnableMiniDump  或  DOTNET_DbgEnableMiniDump: 如果设置为 1，则发生故障时启用 CoreDump 生成。默认值为：0
- COMPlus_DbgMiniDumpType  或  DOTNET_DbgMiniDumpType: 要收集的转储类型。有关详细信息，请看下文的说明。默认值为：2
- COMPlus_DbgMiniDumpName  或  DOTNET_DbgMiniDumpName: 写入转储的文件路径。确保运行 dotnet 进程的用户具有指定目录的写入权限。默认值为：/tmp/coredump.<pid>
- COMPlus_CreateDumpDiagnostics  或  DOTNET_CreateDumpDiagnostics: 如果设置为 1，则启用转储进程的诊断日志记录。默认值为：0
- COMPlus_EnableCrashReport  或  DOTNET_EnableCrashReport：（需要.NET 6 或更高版本，目前仅 Linux 和 MacOS 可用）如果设为 1，运行时会生成 JSON 格式的故障报表，其中包括有关故障应用程序的线程和堆栈帧的信息。故障报表名称是追加了 .crashreport.json 的转储路径/名称。
- COMPlus_CreateDumpVerboseDiagnostics  或  DOTNET_CreateDumpVerboseDiagnostics：（需要 .NET 7 或更高版本）如果设为 1，则启用转储进程的详细诊断日志记录。
- COMPlus_CreateDumpLogToFile  或  DOTNET_CreateDumpLogToFile：（需要 .NET 7 或更高版本）应写入诊断消息的文件路径。如果未设置，则将诊断消息写入故障应用程序的控制台。

> 对于这些环境变量，.NET 7 标准化前缀  DOTNET_，而不是  COMPlus_。但是，COMPlus_  前缀仍将继续正常工作。如果使用的是早期版本的 .NET 运行时，则环境变量仍应该使用  COMPlus_  前缀。

关于DOTNET_DbgMiniDumpType的说明如下所示：

1: Mini 小型 Dump，其中包含模块列表、线程列表、异常信息和所有堆栈。
2: Heap 大型且相对全面的 Dump，其中包含模块列表、线程列表、所有堆栈、异常信息、句柄信息和除映射图像以外的所有内存。
3: Triage 与  Mini  相同，但会删除个人用户信息，如路径和密码。
4: Full 最大的转储，包含所有内存（包括模块映像）。

参考环境变量配置:
```
DOTNET_DbgEnableMiniDump = 1
DOTNET_DbgMiniDumpName = [有权限的Path目录] # "G:\Temp\dotnet\Debug\dump.dmp"
DOTNET_CreateDumpDiagnostics = 1
DOTNET_EnableCrashReport = 1
```


## code analyzer
A source code analyzer built for surfacing features of interest and other characteristics to answer the question 'what's in it' using static analysis with a json based rules engine. Ideal for scanning components before use or detecting feature level changes.
https://github.com/microsoft/ApplicationInspector


## WebRTC
https://github.com/microsoft/MixedReality-WebRTC

## DI DynamicProxy
- Fody

- Castle 

- Emit 操作IL

- NET Remoting中可以利用后门进行拦截，但是必须显示的继承ContextBoundObject

## reporting
PdfReport.Core is a code first reporting engine, which is built on top of the iTextSharp.LGPLv2.Core and EPPlus.Core libraries
https://github.com/VahidN/PdfReport.Core

## docx
https://github.com/egonl/SharpDocx

## Network
https://github.com/chmorgan/packetnet
##  CI Build
https://github.com/dotnetcore/FlubuCore

## CMD
https://github.com/Tyrrrz/CliWrap

## Blazor
https://github.com/oqtane/oqtane.framework
https://github.com/enkodellc/blazorboilerplate

https://github.com/sps014/BlazorML5

## 好玩的项目

emoji 
https://github.com/abock/dotnet-ecoji

## proxy
YARP  https://github.com/microsoft/reverse-proxy
https://github.com/proxykit/ProxyKit

## 代码保护工具.NET Reactor
https://www.eziriz.com/reactor_download.htm

## ASPNET 
新的路由框架，类似gin路由
https://github.com/CarterCommunity/Carter

## dotnet编译
https://github.com/MichalStrehovsky/bflat

## 数学表达式计算math expressions
https://github.com/loresoft/Calculator
https://github.com/mariuszgromada/MathParser.org-mXparser

## UI
https://github.com/oxyplot/oxyplot

https://github.com/tewuapple/WinHtmlEditor
https://github.com/yhuse/SunnyUI
https://github.com/Live-Charts/Live-Charts
https://gitee.com/hapgaoyi/CSharpSkin
https://github.com/kwwwvagaa/NetWinformControl
https://github.com/NetDimension/NanUI/
https://github.com/IgnaceMaes/MaterialSkin

## 微服务框架
https://github.com/Raiffeisen-DGTL/ViennaNET
https://github.com/jamesmh/coravel

## 自动更新
https://github.com/ravibpatel/AutoUpdater.NET
https://github.com/iccfish/FSLib.App.SimpleUpdater

## 扩展
### String
https://github.com/axuno/SmartFormat
```csharp
var s = Smart.Format("{Model.Name} is {Session.Name2}", new { Model =new { Name = "na1", Name2 = "King1" }, Session = new { Name = "na2", Name2 = "King2" } });
var s = Smart.Format("{0.Name} is {1.Name2}", new { Name = "na1", Name2 = "King1" }, new { Name = "na2", Name2 = "King2" });

```
https://github.com/axuno/SmartFormat/wiki/Data-Sources
Using a Dictionary(性能更好，前面是需要反射)
```csharp
Smart.Default.Settings.Formatter.ErrorAction = FormatErrorAction.Ignore; //忽略错误;
Smart.Default.OnFormattingFailure += (sender, args) => {
    Console.WriteLine(args.Placeholder);  // 打印错误占位符;
};
Smart.Default.Settings.CaseSensitivity = CaseSensitivityType.CaseInsensitive;// 忽略大小写;
{
    var s = Smart.Format("{Key} is {Value:was on|will be on}", new Dictionary<string, string>() { { "Key", "站点" } });
    Console.WriteLine(s);
}
{
    var s = Smart.Format("{Key} is {Value:was on|will be on}", new Dictionary<string, string>() { { "Key", "站点" }, { "Value", "" } });
    Console.WriteLine(s);
}
{
    var s = Smart.Format("{Key} is {Value:was on|will be on}", new Dictionary<string, string>() { { "Key", "站点" }, { "Value", "站点2" } });
    Console.WriteLine(s);
}
{
    var s = Smart.Format("{Key} is {Value}", (new Dictionary<string, string>() { { "Key", "站点" } }, new Dictionary<string, string>() {  { "Value", "站点2" } }));
    Console.WriteLine(s);
}

OUTPUT:
站点 is
站点 is will be on
站点 is was on
站点 is 站点2
```

https://docs.microsoft.com/en-us/dotnet/api/system.string.format?view=net-5.0




## plugins
### OrchardCore
https://github.com/OrchardCMS/OrchardCore 4.7k
### nopCommerce
https://github.com/nopSolutions/nopCommerce 5.6k
### PluginFramework
Everything is a Plugin in .NET
https://github.com/weikio/PluginFramework 166

### Prise
https://github.com/merken/Prise 172
### CoolCat
https://github.com/lamondlu/CoolCat 190

## language
### IronPython
https://github.com/IronLanguages/ironpython3
https://github.com/IronLanguages/ironpython2

### ironruby
https://github.com/IronLanguages/ironruby

### IronJS
https://github.com/fholm/IronJS

### Pyjion
[Pyjion - A JIT for Python based upon CoreCLR](https://github.com/tonybaloney/Pyjion)

[NumPy aware dynamic Python compiler using LLVM](https://github.com/numba/numba)

### peachpie
[PeachPie - the PHP compiler and runtime for .NET and .NET Core](https://github.com/peachpiecompiler/peachpie)


## dotnet社区

- [MSDN](http://msdn.microsoft.com/zh-cn/default.aspx)
- [CodeProject](http://www.codeproject.com)
- [CodePlex](http://www.codeplex.com/)
- [ASP.NET/ ASP.NET MVC](http://www.asp.net)
- [微软.NET WinForm/WPF](http://windowsclient.net/)
- [InfoQ](http://www.infoq.com/)
- [C# Corner](http://www.c-sharpcorner.com/)
- [SourceForge](http://sourceforge.net/)
- [CsharpHelp](http://www.csharphelp.com/)
- [silverlight.net](http://www.silverlight.net/)
- [TooSlowException](https://tooslowexception.com/)

- [英文社区和英文个人站点](http://dotnetslackers.com/)
- [英文社区和英文个人站点](http://weblogs.asp.net/scottgu/)
- [英文社区和英文个人站点](http://codebetter.com/blogs/)
- [英文社区和英文个人站点](http://www.drdobbs.com/)
- [英文社区和英文个人站点](http://ayende.com/Blog/)
- [英文社区和英文个人站点](http://www.nikhilk.net/)
- [英文社区和英文个人站点](http://www.markhneedham.com/blog/)
- [英文社区和英文个人站点](http://blogs.msdn.com/b/somasegar/)
- [英文社区和英文个人站点](http://blogs.msdn.com/b/tims/archive/2005/09/28/windows-presentation-foundation-blogs.aspx)

- [技术社区](http://stackoverflow.com/)
- [技术社区](http://www.theserverside.net/)
- [技术社区](http://dotnet.sys-con.com/)
- [技术社区](http://en.csharp-online.net/Main)

- [ebook](http://www.wowebook.com/category/e-book/dot-net)
- [ebook](http://gigapedia.com/)

- [cnblogs](http://www.cnblogs.com)
- [CSDN](http://www.csdn.net)
- [51CTO](http://www.51cto.com/)
- [InfoQ中文版](http://www.infoq.com/cn/)
- [zdnet中文版 ](http://www.zdnet.com.cn/)
- [MSDN](http://msdn.microsoft.com/zh-cn/default.aspx)
- [ITPUB](http://www.itpub.net/)
- [博客堂](http://blog.joycode.com/)
```
http://www.51aspx.com/

http://www.silverlightchina.net/

http://www.chenjiliang.com/

http://www.rainsts.net/

http://tech.ddvip.com/

http://www.ccidnet.com/

http://www.programfan.com/article/article.asp?classid=18

http://bbs.xml.org.cn/index.asp

http://www.ibm.com/developerworks/cn/web/

http://social.microsoft.com/forums/zh-CN/categories/

http://www.studycs.com/html/index.html

http://blog.csdn.net/21aspnet/

http://dotnet.aspx.cc/
```

## 基于.NET Core的优秀开源项目合集
https://www.cnblogs.com/myshowtime/p/14315080.html

[基于.NET Core的优秀开源项目合集](https://www.toutiao.com/i7038781917931651598)
### demo
https://github.com/dotnet-architecture/eShopOnContainers eShopOnContainers是一个示例参考应用程序，演示了Microsoft的基于容器的微服务的各种体系结构模式。
https://github.com/dotnet-architecture/eShopOnWeb eShopOnWeb是一个示例参考应用程序，演示了Microsoft提供支持的单体架构。
https://github.com/dodyg/practical-aspnetcore
https://github.com/jasontaylordev/NorthwindTraders Entity Framework 和CQRS模式的DDD的示例项目
https://github.com/AdaptiveConsulting/ReactiveTraderCloud 实时交易应用程序，展示了反应式编程原理。

https://github.com/vietnam-devs/coolstore-microservices 演示了如何通过Service Mesh来使用Kubernetes。

https://github.com/ivanpaulovich/clean-architecture-manga 一个整洁架构的参考示例项目。

https://github.com/JacekKosciesza/StarWars 一个基于GraphQL的ASP.NET Core Star Wars应用程序。

https://github.com/kgrzybek/sample-dotnet-core-cqrs-api  Clean architecture, DDD, CQRS

### Commerce

https://github.com/nopSolutions/nopCommerce
https://github.com/simplcommerce/SimplCommerce

### CMS
https://github.com/OrchardCMS/OrchardCore
https://github.com/Squidex/squidex squidex是无头CMS和内容管理中心，使用具有OData和CQRS模式的ASP.NET Core构建。

### blog
https://github.com/madskristensen/Miniblog.Core
https://github.com/piranhacms/piranha.core

## 其它
计算方法执行时间
https://github.com/Fody/MethodTimer