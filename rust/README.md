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

## web框架对比
https://github.com/the-benchmarker/web-frameworks#full-table