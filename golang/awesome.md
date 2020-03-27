[awesome-go](https://github.com/avelino/awesome-go)


## x
golang.org/x/time/rate
该限流器是基于Token Bucket(令牌桶)实现的。
## DI

### Google wire 3.3k
https://github.com/google/wire

Wire 可以生成 Go 源码并在编译期完成依赖注入。它不需要反射机制或 [Service Locators](https://en.wikipedia.org/wiki/Service_locator_pattern)

好处：
1. 方便 debug，若有依赖缺失编译时会报错
2. 因为不需要 Service Locators， 所以对命名没有特殊要求
3. 避免依赖膨胀。生成的代码只包含被依赖的代码，而运行时依赖注入则无法作到这一点
4. 依赖关系静态存于源码之中， 便于工具分析与可视化

[Compile-time Dependency Injection With Go Cloud's Wire](https://blog.golang.org/wire)


[一文读懂 Go官方的 Wire](https://mp.weixin.qq.com/s/ZQKi9O7DRJ3qGWhDL9aTVg)

### Uber dig 1.3k
运行时依赖注入
https://github.com/uber-go/dig

### Facebook inject 1.2k(归档了，不更新)
运行时依赖注入
https://github.com/facebookarchive/inject

## spinal-case(脊柱) or snake_case(蛇形) or CamelCase(驼峰式) or KebabCase(短横线) or PascalCase(帕斯卡命名法) or PascalSnakeCase
https://github.com/iancoleman/strcase

## GUI


Cross platform GUI in Go based on Material Design https://fyne.io/
https://github.com/fyne-io/fyne


请注意，默认情况下，Windows应用程序是从命令提示符加载的，这意味着，如果单击图标，则可能会看到命令窗口。 要解决此问题，请在运行或构建命令中添加参数-ldflags -H = windowsgui。



Prerequisites
https://fyne.io/develop/
Windows
1. Download Go from the download page and follow instructions
2. Install one of the available C compilers for windows, the following are tested with Go and Fyne:
    - MSYS2 with MingW-w64 - msys2.org
    - TDM-GCC - tdm-gcc.tdragon.net
    - Cygwin - cygwin.com
3. In Windows your graphics driver will already be installed, but it is recommended to ensure they are up to date.