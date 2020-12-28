[编译出 dll](https://www.cnblogs.com/timeddd/p/11731160.html)

```
go build --buildmode=c-shared -o Test.dll
dotnet
[DllImport(DLL_NAME, EntryPoint = "Test")]
```
### 六大主流语言代码漏洞分析
https://www.veracode.com/sites/default/files/pdf/resources/ipapers/security-flaw-heatmap/index.html

### 让出 CPU 时间片

#### GO

用于让出 CPU 时间片
runtime.Gosched()

#### .net

Thread.Sleep()方法：是强制放弃 CPU 的时间片，然后重新和其他线程一起参与 CPU 的竞争。
用 Sleep()方法是会让线程放弃 CPU 的使用权。
用 SpinWait()方法是不会放弃 CPU 的使用权。

#### .net

await Task.Yield();// 让出时间片
Thread.SpinWait();// 不让出时间片

AsyncLocal 4.6+

4.5
using System.Runtime.Remoting;
using System.Runtime.Remoting.Messaging;
CallContext

#### java

Thread.yield(); //让出当前剩余的 CPU 时间片
Thread.onSpinWait();// 不让出时间片

TransmittableThreadLocal

## grpc Balance

```
conn, err := grpc.Dial(
    "",
    grpc.WithInsecure(),
    grpc.WithBalancer(grpc.RoundRobin(resolver.NewPseudoResolver([]string{
        "10.0.0.1:10000",
        "10.0.0.2:10000",
    }))),
)
```

## 代码格式化

### clang-format

下载https://github.com/llvm/llvm-project/releases

```
# win64
https://github.com/llvm/llvm-project/releases/download/llvmorg-10.0.0/LLVM-10.0.0-win64.exe

# ubuntu/debian
sudo apt-get install clang-format
```

文档
http://clang.llvm.org/docs/ClangFormat.html
支持 C/C++/Java/JavaScript/Objective-C/Protobuf/C#

代码格式支持：LLVM, Google, Chromium, Mozilla, WebKit
http://clang.llvm.org/docs/ClangFormatStyleOptions.html

[团队协作-代码格式化工具 clang-format](https://www.toutiao.com/i6886080589141639693)

Visual Studio Code 安装插件 clang-format

其它插件

- markdown-formatter
- shell-format
- Prettier - Code formatter

shift+alt+f
or
shift+ctrl+p

### Prettier
https://github.com/prettier

支持
- JavaScript (including experimental features)
- JSX
- Angular
- Vue
- Flow
- TypeScript
- CSS, Less, and SCSS
- HTML
- JSON
- GraphQL
- Markdown, including GFM and MDX
- YAML
