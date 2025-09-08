# Rust

## 标准库
https://doc.rust-lang.org/std/index.html

## cargo
[详情](cargo.md)

https://doc.rust-lang.org/cargo/index.html

## rustdoc

https://doc.rust-lang.org/rustdoc/index.html
参数：https://doc.rust-lang.org/rustdoc/command-line-arguments.html
```
$ rustdoc -h
$ rustdoc --help
$ rustdoc -V
$ rustdoc --version

更多的输出
$ rustdoc -v src/lib.rs
$ rustdoc --verbose src/lib.rs

rustdoc src/lib.rs # doc/lib/index.html
rustdoc src/lib.rs --crate-name docs #  doc/docs/index.html

rustdoc README.md  # md文件生成README.html

cargo doc
or
rustdoc --crate-name docs src/lib.rs -o <path>/docs/target/doc -L
dependency=<path>/docs/target/debug/deps
```
## rustc
https://doc.rust-lang.org/rustc/index.html
https://doc.rust-lang.org/rustc/platform-support.html
参数：https://doc.rust-lang.org/rustc/command-line-arguments.html
```
如果main.rs中引用了foo.rs
只用rustc main.rs 编译就可以了

```
大多数Rust程序员并不直接调用rustc，而是通过Cargo进行调用。不过这都是为rustc服务的!如果你想知道货物如何呼叫rustc，你可以`cargo build --verbose`

## rustup
[详情](rustup.md)

## 编译错误指南
https://doc.rust-lang.org/error-index.html

## Live Reload / Hot Reload 热更新

1. 通过WASM运行时进行热加载。比如 Substrate框架，自己实现了一个WASM运行时。也可以通过wasmer这样的 wasm vm 来支持热加载。
2. 利用 Rust 为宿主的脚本语言。比如社区里有一个 mun-runtime，是[mun 编程语言](https://github.com/mun-lang/mun) （Rust实现），可以和 Rust 无缝集成，并先天支持Hot Reload。

3. 通过动态库实现。比如，[live-reloading-rs](https://github.com/porglezomp-misc/live-reloading-rs) 。这里也有篇[相关的文章](https://silverweed.github.io/Rust_game_programming_code_hotloading/)。

4. 平时开发的话，可以用[cargo-watch](https://github.com/passcod/cargo-watch)，它会监控代码文件的变化，并自动重新编译、运行代码。
还有另外一种热更新是资源热更新，比如配置文件、游戏素材什么的，可参考：
https://github.com/phaazon/warmy
https://github.com/amethyst/amethyst/blob/master/amethyst_assets/src/reload.rs

## 参考
Rust语言示例：https://rustwiki.org/zh-CN//rust-by-example/hello.html

Rust项目列表：https://github.com/rust-unofficial/awesome-rust

Rust入门书籍：https://rustcc.gitbooks.io/rustprimer/content/

Rust程序设计（第一版）：https://kaisery.gitbooks.io/rust-book-chinese/content/

Rust程序设计（第二版）：https://kaisery.github.io/trpl-zh-cn/

## rust基础
http://wiki.jikexueyuan.com/project/rust-primer/cargo-projects-manager/cargo-projects-manager.html
https://rustcc.gitbooks.io/rustprimer/content/quickstart/rust-travel.html

## 每日rust
https://zhuanlan.zhihu.com/p/50737367

## Awesome Rust
[一个可以比较各种库的网站](https://rust.libhunt.com/categories)

[awesome-rust](https://github.com/rust-unofficial/awesome-rust)

### web框架对比
https://github.com/the-benchmarker/web-frameworks#full-table

https://rust.libhunt.com/categories/5464-web