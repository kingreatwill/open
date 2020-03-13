[awesome-go](https://github.com/avelino/awesome-go)

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

