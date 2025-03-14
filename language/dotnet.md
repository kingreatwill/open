[TOC]
# dotnet
.NET 官方社区论坛/.NET 技术论坛
[ .NET tech community](https://techcommunity.microsoft.com/t5/net/ct-p/dotnet)
## xx
dotnet tool install -g dotnet-try
dotnet tool install --global dotnet-serve

[【ASP.NET Core】配置应用程序地址的N多种方法](https://www.cnblogs.com/tcjiaan/p/16347491.html)

[sdmap](https://github.com/sdcb/sdmap) ORM 类iBatis，不是xml配置!

[Automatonymous](https://github.com/MassTransit/Automatonymous) 状态机

[benchmarks](https://aka.ms/aspnet/benchmarks)

[CefSharp](https://github.com/cefsharp/CefSharp) Chromium嵌入式框架的.NET（WPF和Windows窗体）绑定

[Playnite](https://github.com/JosefNemec/Playnite) 开源视频游戏库管理器，支持Steam，GOG，Origin，Battle.net和Uplay等第三方库。包括游戏仿真支持，为您的游戏提供一个统一的界面。

[REST library for .NET Core](https://github.com/reactiveui/refit)

最经典的就是选择试用Workstation GC（WKS GC）或者Server GC（SVR GC）。见过`<gcServer>`参数不？后来可以配置使用Concurrent GC、Background Workstation GC、Background Server GC等。用户还可以在代码里通过 GCSettings.LatencyMode 属性来影响GC的行为。[调优参数列表](https://docs.microsoft.com/en-us/dotnet/framework/configure-apps/file-schema/runtime/index?redirectedfrom=MSDN)



[MAUI 跨平台UI - 将在.net6中取代Xamarin.Forms](https://github.com/dotnet/maui)
“Model-View-Update”（MVU）模式。
[A Model-View-Update Proof of Concept for Xamarin.iOS and Xamarin.Android](https://github.com/69grad/Fabulous.XamarinNative)

[WebSharper UI F#](https://github.com/dotnet-websharper/mvu)

[Elm语言 Model View Update - Part 1](https://elmprogramming.com/model-view-update-part-1.html)
https://elm-lang.org/

感觉就像Flutter~的开发模式
查找下资料-> [Dartea](https://github.com/p69/dartea) - 为Flutter实现MVU（模型视图更新）模式  很遗憾没有更新了
[Functional Model-View-Update Architecture for Flutter](https://buildflutter.com/functional-model-view-update-architecture-for-flutter/)
正与我想的一致！

支持：MVVM, RxUI, MVU, Blazor  https://devblogs.microsoft.com/dotnet/introducing-net-multi-platform-app-ui/
MVU：类似flutter开发
MVVM：Xamarin.Forms	

Xamarin.Forms will become System.Maui
.NET 的大一统，那么迟早也会实现自己的一套跨平台 UI 方案
https://github.com/dotnet/maui/wiki/Roadmap

flutter 有近50W开发者
https://venturebeat.com/2020/04/22/google-500000-developers-flutter-release-process-versioning-changes/

```
MVVM

Model-View-ViewModel（MVVM）和 XAML 是 .NET 开发人员数十年来的主要模式和实践，它们是MAUI中的一流功能，这将继续发展，以帮助您高效地构建和维护生产应用程序。

UI部分
<StackLayout>
    <Label Text="Welcome to MAUI!" />
    <Button Text="{Binding Text}" 
            Command="{Binding ClickCommand}" />
</StackLayout>


代码部分
public Command ClickCommand { get; }

public string Text { get; set; } = "Click me";

int count = 0;

void ExecuteClickCommand ()
{
    count++;
    Text = $"You clicked {count} times.";
}


MVU

MVU促进数据和状态管理的单向流程，以及通过仅应用必要的更改来快速更新UI的代码优先开发经验。

readonly State<int> count = 0;

[Body]
View body() => new StackLayout
{
    new Label("Welcome to MAUI!"),
    new Button(
        () => $"You clicked {count} times.",
        () => count.Value ++)
    )
};
```

cnblogs dockerfile
```
FROM mcr.microsoft.com/dotnet/aspnet:5.0-buster-slim AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443
RUN sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get update && apt-get install -y curl

FROM mcr.microsoft.com/dotnet/sdk:5.0-buster-slim AS build
WORKDIR /src
COPY src/*.sln src/*.props src/NuGet.config ./
COPY src/*/*.csproj ./
RUN for file in $(ls *.csproj); do mkdir -p ${file%.*}/ && mv $file ${file%.*}/; done
RUN dotnet restore "BlogServerCore.sln"
COPY src/. .
RUN dotnet build "BlogServerCore.sln" -c Release --no-restore

FROM build AS publish
WORKDIR /src/BlogServer.WebApi
RUN dotnet publish "BlogServer.WebApi.csproj" -c Release -o /app/publish --no-build

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
RUN echo "dotnet BlogServer.WebApi.dll" > run.sh
HEALTHCHECK --interval=5s --timeout=20s \
    CMD curl -fs -o /dev/null localhost/alive || exit 1
```

## 在.NET Core 中收集数据的几种方式
https://www.cnblogs.com/myshowtime/p/14199844.html

- 手动埋点
- Middleware 中间件 & 过滤器 Filter
- DiagnosticSource   基于发布订阅模式
SkyApm-dotnet https://github.com/SkyAPM/SkyAPM-dotnet
HttpReports APM https://github.com/dotnetcore/HttpReports


- AOP
- ETW(Event Tracing for Windows)
- Mono.Cecil
- CLR Profiling API
听云APM（商业）OneAPM （商业）Datadog （商业）
https://docs.microsoft.com/en-us/archive/blogs/yirutang/clr-profiling-api

## dotnet 诊断工具 监控 分析
https://github.com/dotnet/diagnostics

- dotnet-monitor  https://github.com/SachiraChin/dotnet-monitor-ui
```
dotnet tool install -g dotnet-monitor --add-source https://dnceng.pkgs.visualstudio.com/public/_packaging/dotnet5-transport/nuget/v3/index.json --version 5.0.0-preview.*

dotnet monitor collect # 会开放 http://localhost:52323 和 http://localhost:52325 两个端口
```

- SOS - About the SOS debugger extension.
- dotnet-dump - Dump collection and analysis utility.
- dotnet-gcdump - Heap analysis tool that collects gcdumps of live .NET processes.
- dotnet-trace - Enable the collection of events for a running .NET Core Application to a local trace file.
- dotnet-counters - Monitor performance counters of a .NET Core application in real time.

[.netcore利用perf分析高cpu使用率](https://www.cnblogs.com/wu_u/p/14263349.html)
[利用dotnet-dump分析docker容器内存泄露](https://www.cnblogs.com/wu_u/p/14109333.html)


[PerfView](https://github.com/Microsoft/perfview)是一款免费的性能分析工具，可帮助排查CPU和内存相关的性能问题。
可以查看火焰图，也可以使用[speedscope](https://www.speedscope.app/)，一个交互式火焰图可视化工具，帮助我们分析。

[pyroscope](https://github.com/pyroscope-io/pyroscope)
## 调试工具
- dnSpy
dnSpy 是用于 .NET 调试的最有用的工具之一。它是一个很好的反编译器。但是它的主要用途是作为调试器。dnSpy允许你调试任何 .NET程序你，而无需考虑符号或者源代码。

- dotPeek
dotPeek是JetBrains的免费.NET反编译器。它们的许多工具实际上进入了该列表。其它两个反编译器（如ILSpy或JustDecompile）

- dotTrace
- SciTech's .NET Memory Profiler
- OzCode
- SysInternals Suite
  SysInternals是一套用于对Windows软件进行故障排除和监视的实用程序。它包括一些我们调试所需的最重要的工具。
  Process Explorer
  Process Monitor 也称为ProcMon，允许你监视流程活动事件。
  ProcDump是用于保存转储文件的命令行工具。

- Performance Monitor (PerfMon)
- PerfView
- Fiddler
## DataFrame
https://devblogs.microsoft.com/dotnet/an-introduction-to-dataframe/
https://github.com/dotnet/corefxlab/tree/master/src/Microsoft.Data.Analysis

dotnet spark ML python 能很好的一起工作，可以向pandas一样读取CVS

## Linux守护进程
Microsoft.Extensions.Hosting.Systemed，是利用Linux systemd机制，由sytemd将你的.net core控制台应用程序作为后台服务程序进行开机启动和管控。
而本文是让你的程序本身就是一个后台程序（Linux daemon，守护进程），是“原生”的，不需要借助任何工具（包括systemd、supervisor等等）！
换句话说，借助其它的工具或方式把你的程序“变成后台服务”与你的程序“本身就是后台服务”程序，是完全不同的两回事。

```csharp
using System.Threading;
using System.Timers;
using System.Runtime.InteropServices;
using System.IO;
using System.Text;
 
 
/************************************************
 * .Net Core/.Net5+ Linux Daemon示例，作者宇内流云 *
 ************************************************/
 
namespace daemon
{
    class Program
    {
 
        static unsafe void Main(string[] args)
        {
            // 进入守护状态
            int pid = fork();
            if (pid != 0) exit(0);
            setsid();
            pid = fork();
            if (pid != 0) exit(0);
            umask(0);
 
 
            // 关闭所有打开的文件描述符
            int fd_nul = open("/dev/null", 0);
            for (var i = 0; i <= fd_nul; i++)
            {
                if (i < 3)
                    dup2(fd_nul, i);
                else
                    close(i);
            }
 
            // 进入主方法
            //（本示例的功能很简单，就是定时向某个文件写入点内容）
            DaemonMain(args);
        }
 
 
        /// <summary>
        /// Daemon工作状态的主方法
        /// </summary>
        /// <param name="aargs"></param>
        static void DaemonMain(string[] aargs)
        {
            //启动一个线程去处理一些事情
            (new Thread(DaemonWorkFunct) { IsBackground = true }).Start();
 
 
            //daemon时，控制台输入、输出流已经关闭
            // 因此，请不要再用Console.Write/Read等方法
 
            //阻止daemon进程退出
            (new AutoResetEvent(false)).WaitOne();
        }
 
 
        static FileStream fs;
        static int count = 0;
        static void DaemonWorkFunct()
        {
            try
            {
                fs = File.Open(Path.Combine("/tmp", "daemon.txt"), FileMode.OpenOrCreate);
            }
            catch
            {
                exit(1);
                return;
            }
 
            var t = new System.Timers.Timer() { Interval = 1000 };
            t.Elapsed += OnElapsed;
            t.Start();
        }
 
        private static void OnElapsed(object sender, ElapsedEventArgs e)
        {
            var s = DateTime.Now.ToString("yyy-MM-dd HH:mm:ss") + "\n";
            var b = Encoding.ASCII.GetBytes(s);
            fs.Write(b, 0, b.Length);
            fs.Flush();
 
            count++;
            if (count > 100)
            {
                fs.Close();
                fs.Dispose();
                exit(0);
            }
        }
 
 
 
        [DllImport("libc", SetLastError = true)]
        static extern int fork();
 
        [DllImport("libc", SetLastError = true)]
        static extern int setsid();
 
        [DllImport("libc", SetLastError = true)]
        static extern int umask(int mask);
 
        [DllImport("libc", SetLastError = true)]
        static extern int open([MarshalAs(UnmanagedType.LPStr)] string pathname, int flags);
 
        [DllImport("libc", SetLastError = true)]
        static extern int close(int fd);
 
        [DllImport("libc", SetLastError = true)]
        static extern int exit(int code);
 
        [DllImport("libc", SetLastError = true)]
        static extern int dup2(int oldfd, int newfd);
 
    }
}
```
[linux系统编程之进程（八）：守护进程详解及创建，daemon()使用](https://www.cnblogs.com/mickole/p/3188321.html)
[.NET跨平台实践：.NetCore、.Net5/6 Linux守护进程设计](https://www.cnblogs.com/yunei/p/15367709.html)

## 反混淆和脱壳工具
> 脱壳，是完全破除压缩后软件无法编辑的限制，去掉头部的解压缩指令，然后解压出加壳前的完整软件。

[de4dot](https://github.com/de4dot/de4dot)
[EazFixer](https://github.com/holly-hacker/EazFixer)

### PECompact
PECompact 压缩壳，脱壳

PECompact是一款专业的资源压缩工具，他可以压缩你开发软件或者代码的体积，中文版由大眼仔旭汉化。

PECompact 是一个能压缩可执行文件的工具，通过压缩代码、数据、相关资源使压缩能达到100%，由于在运行时不需要恢复磁盘上压缩后的数据，所以与没有压缩的程序在运行时没有明显的速度差异，在某种程度上还有所改善。
你要压缩，但需要选择一个压缩机。有UPX ， ASPACK和PECompact 。这些是前3名的PE （ 32位）本地EXE压缩机。 UPX具有减压开关，是开源的（通常由恶意软件编写者滥用） ， ASPACK不能压缩以及PECompact ，也没有插件支持的。它也缺乏PECompact的其他主要功能。

## 混淆工具
[ConfuserEx](https://github.com/mkaring/ConfuserEx)
[obfuscar](https://github.com/obfuscar/obfuscar)

## dotnet tool
https://docs.microsoft.com/zh-cn/dotnet/core/tools/global-tools#install-a-global-tool
https://docs.microsoft.com/zh-cn/dotnet/core/tools/dotnet-tool-search
https://www.nuget.org/packages?packagetype=dotnettool


### dotnet tool install -g Microsoft.dotnet-httprepl
https://docs.microsoft.com/zh-cn/aspnet/core/web-api/http-repl/?view=aspnetcore-5.0&tabs=windows

### dotnet tool install --global dotnet-serve
https://github.com/natemcmaster/dotnet-serve

### dotnet运行时诊断工具
https://github.com/dotnet/diagnostics