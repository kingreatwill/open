dotnet tool install -g dotnet-try
dotnet tool install --global dotnet-serve

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
