# cargo介绍

作为一门现代语言，rust自然要摒弃石器时代项目代码管理的方法和手段。rust项目组为各位猿提供了超级大杀器cargo，以解决项目代码管理所带来的干扰和困惑。用过node.js的猿们，应该对node.js中的神器npm、grunt、gulp等工具印象深刻。作为新一代静态语言中的翘楚，rust官方参考了现有语言管理工具的优点，于是就产生了cargo。

言而总之，作为rust的代码组织管理工具，cargo提供了一系列的工具，从项目的建立、构建到测试、运行直至部署，为rust项目的管理提供尽可能完整的手段。同时，与rust语言及其编译器rustc本身的各种特性紧密结合，可以说既是语言本身的知心爱人，又是rust猿们的贴心小棉袄，谁用谁知道。

## toml配置文件
https://github.com/toml-lang/toml

## 创建项目 hellorust
 cargo new hellorust --bin
## 编译和运行

cargo build

cargo build --release # 这个属于优化编译

./target/debug/hellorust.exe

./target/release/hellorust.exe # 如果前面是优化编译，则这样运行

cargo run # 编译和运行合在一起

cargo run --release # 同上，区别是是优化编译的

cargo test #运行所有的测试

cargo package
```
build       编译当前包
check       检查当前包并寻出错误，但不进行编译
clean       删除编译结果（即target文件夹）
doc         构建当前包以及依赖项得文档
new         新建一个crate
init        以当前文件夹初始化一个crate
run         编译并执行src/main.rs
test        执行测试项
bench       执行基准测试项
update      更新所需的依赖项并预编译
search      搜索crates
publish     打包发布
install     安装cargo相关可执行文件，默认路径为 $HOME/.cargo/bin
uninstall   卸载相关可执行文件
```
## 建议项目结构
```
.
├── Cargo.toml
├── Cargo.lock
├── tests
├── examples
├── benches
├── target
    ├── debug
    └── release
└── src
    ├── lib.rs
    └── main.rs
```

cargo.toml和cargo.lock文件总是位于项目根目录下。

源代码位于src目录下。

默认的库入口文件是src/lib.rs。

默认的可执行程序入口文件是src/main.rs。

其他可选的可执行文件位于src/bin/*.rs(这里每一个rs文件均对应一个可执行文件)。

外部测试源代码文件位于tests目录下。

示例程序源代码文件位于examples。

基准测试源代码文件位于benches目录下。

```
.
├── Cargo.lock
├── Cargo.toml
├── src/
│   ├── lib.rs
│   ├── main.rs
│   └── bin/
│       ├── named-executable.rs
│       ├── another-executable.rs
│       └── multi-file-executable/
│           ├── main.rs
│           └── some_module.rs
├── benches/
│   ├── large-input.rs
│   └── multi-file-bench/
│       ├── main.rs
│       └── bench_module.rs
├── examples/
│   ├── simple.rs
│   └── multi-file-example/
│       ├── main.rs
│       └── ex_module.rs
└── tests/
    ├── some-integration-tests.rs
    └── multi-file-test/
        ├── main.rs
        └── test_module.rs
```
- Cargo.toml and Cargo.lock are stored in the root of your package (package root).
- Source code goes in the src directory.
- The default library file is src/lib.rs.
- The default executable file is src/main.rs.
    - Other executables can be placed in src/bin/.
- Benchmarks go in the benches directory.
- Examples go in the examples directory.
- Integration tests go in the tests directory.


## 定义项目依赖

基于rust官方仓库crates.io，通过版本说明来描述：

基于项目源代码的git仓库地址，通过URL来描述：

基于本地项目的绝对路径或者相对路径，通过类Unix模式的路径来描述： 这三种形式具体写法如下：
````
[dependencies]
typemap = "0.3"
plugin = "0.2*"
hammer = { version = "0.5.0"}
color = { git = "https://github.com/bjz/color-rs" }
geometry = { path = "crates/geometry" }
````

## cargo镜像
在用户目录.cargo文件夹或在与Cargo.toml同级目录.cargo文件夹下创建config文件

```
[source.crates-io]
registry = "https://github.com/rust-lang/crates.io-index"
# 指定镜像
replace-with = 'sjtu'

# 清华大学
[source.tuna]
registry = "https://mirrors.tuna.tsinghua.edu.cn/git/crates.io-index.git"

# 中国科学技术大学
[source.ustc]
registry = "git://mirrors.ustc.edu.cn/crates.io-index"

# 上海交通大学
[source.sjtu]
registry = "https://mirrors.sjtug.sjtu.edu.cn/git/crates.io-index"

# rustcc社区
[source.rustcc]
registry = "https://code.aliyun.com/rustcc/crates.io-index.git"
```

Cargo 支持用另一个来源更换一个来源的能力，可根据镜像或 vendoring 依赖关系来表达倾向。要配置这些，目前通过.cargo/config配置机制完成，像这样:
```
# `source` 表下，就是存储有关要更换的来源名称
[source]

# 在`source` 表格之下的，可为一定数量的有关来源名称. 示例下面就# 定义了一个新源， 叫 `my-awesome-source`， 其内容来自本地 # `vendor`目录 ，其相对于包含`.cargo/config`文件的目录
[source.my-awesome-source]
directory = "vendor"

# Git sources 也指定一个 branch/tag/rev
git = "https://example.com/path/to/repo"
# branch = "master"
# tag = "v1.0.1"
# rev = "313f44e8"

# The crates.io 默认源 在"crates-io"名称下， 且在这里我们使用 `replace-with` 字段指明 默认源更换成"my-awesome-source"源
[source.crates-io]
replace-with = "my-awesome-source"
```

使用此配置，Cargo 会尝试在"vendor"目录中，查找所有包，而不是 查询在线注册表 crates.io 。Cargo 有两种来源更换的表达 :

- 供应(Vendoring) - 可以定义自定义源，它们表示本地文件系统上的包。这些源是它们正在更换的源的子集，并在需要时可以检入包中。

- 镜像(Mirroring) - 可以更换为等效版本的源，行为表现为 `crates.io` 本身的缓存。

Cargo 有一个关于来源更换的核心假设，源代码从两个完全相同的源而来。在上面的例子中，Cargo 假设所有的箱子都来自my-awesome-source，与crates-io副本完全相同。请注意，这也意味着my-awesome-source，不允许有crates-io源不存在的箱。

因此，来源更换不适用于依赖项补丁(fix bug)，或私有注册表等情况。Cargo 是通过使用[replace]字段支持依赖项补丁，计划为未来版本的 Cargo 提供私人注册表的支持。

## 发布到 crates.io

进入：https://crates.io/me

点击：New Token
cmd: cargo login (token)
查看：~/.cargo/credentials

我们将使用cargo package子命令。这将把我们的整个包装箱全部打包成一个*.crate文件，其在target/package目录中。
```
$ cargo package
```
作为一个额外的功能，`*.crate`将独立于当前源树进行验证。在`*.crate`创建之后，会解压到target/package目录，然后从头开始构建，以确保构建成功的所有必要文件。可以使用`--no-verify`参数禁用此行为。

现在是时候看看`*.crate`文件了，为了确保您不会意外地打包 2GB 视频资源，或用于代码生成，集成测试或基准测试的大型数据文件。目前存在 10MB 的`*.crate`文件上传大小限制。所以，如果tests和benches目录及其依赖项大小，最多只达 几 MB，您仍可以将它们保存在包; 不然的话，最好排除它们。

在打包时，Cargo 会自动忽略版本控制系统的忽略文件，但是如果要指定要额外的忽略文件集，则可以使用清单中的exclude字段:

```
[package]
# ...
exclude = [
    "public/assets/*",
    "videos/*",
]
```

这个数组中每个元素接受的语法是[rust-lang/glob](https://github.com/rust-lang/glob)。如果您宁愿使用白名单，而不是黑名单,Cargo 也支持include字段，如果设置,则会覆盖exclude字段:
```
[package]
# ...
include = [
    "**/*.rs",
    "Cargo.toml",
]
```
上传(也可以跳过cargo package)
```
$ cargo publish
```
> 注意语义版本控制 [Semantic Versioning 2.0.0](https://semver.org/)
## config
https://doc.rust-lang.org/cargo/reference/config.html

## Cargo.toml
[The Manifest Format 清单格式](https://cargo.budshome.com/reference/manifest.html)

## 交叉编译
查看支持的目标平台
rustup target list
```
aarch64-apple-ios
aarch64-fuchsia
aarch64-linux-android
aarch64-pc-windows-msvc
aarch64-unknown-linux-gnu
aarch64-unknown-linux-musl
aarch64-unknown-none
aarch64-unknown-none-softfloat
arm-linux-androideabi
arm-unknown-linux-gnueabi
arm-unknown-linux-gnueabihf
arm-unknown-linux-musleabi
arm-unknown-linux-musleabihf
armebv7r-none-eabi
armebv7r-none-eabihf
armv5te-unknown-linux-gnueabi
armv5te-unknown-linux-musleabi
armv7-linux-androideabi
armv7-unknown-linux-gnueabi
armv7-unknown-linux-gnueabihf
armv7-unknown-linux-musleabi
armv7-unknown-linux-musleabihf
armv7a-none-eabi
armv7r-none-eabi
armv7r-none-eabihf
asmjs-unknown-emscripten
i586-pc-windows-msvc
i586-unknown-linux-gnu
i586-unknown-linux-musl
i686-linux-android
i686-pc-windows-gnu
i686-pc-windows-msvc
i686-unknown-freebsd
i686-unknown-linux-gnu
i686-unknown-linux-musl
mips-unknown-linux-gnu
mips-unknown-linux-musl
mips64-unknown-linux-gnuabi64
mips64-unknown-linux-muslabi64
mips64el-unknown-linux-gnuabi64
mips64el-unknown-linux-muslabi64
mipsel-unknown-linux-gnu
mipsel-unknown-linux-musl
nvptx64-nvidia-cuda
powerpc-unknown-linux-gnu
powerpc64-unknown-linux-gnu
powerpc64le-unknown-linux-gnu
riscv32i-unknown-none-elf
riscv32imac-unknown-none-elf
riscv32imc-unknown-none-elf
riscv64gc-unknown-linux-gnu
riscv64gc-unknown-none-elf
riscv64imac-unknown-none-elf
s390x-unknown-linux-gnu
sparc64-unknown-linux-gnu
sparcv9-sun-solaris
thumbv6m-none-eabi
thumbv7em-none-eabi
thumbv7em-none-eabihf
thumbv7m-none-eabi
thumbv7neon-linux-androideabi
thumbv7neon-unknown-linux-gnueabihf
thumbv8m.base-none-eabi
thumbv8m.main-none-eabi
thumbv8m.main-none-eabihf
wasm32-unknown-emscripten
wasm32-unknown-unknown
wasm32-wasi
x86_64-apple-darwin
x86_64-apple-ios
x86_64-fortanix-unknown-sgx
x86_64-fuchsia
x86_64-linux-android
x86_64-pc-windows-gnu
x86_64-pc-windows-msvc (installed)
x86_64-rumprun-netbsd
x86_64-sun-solaris
x86_64-unknown-cloudabi
x86_64-unknown-freebsd
x86_64-unknown-linux-gnu
x86_64-unknown-linux-gnux32
x86_64-unknown-linux-musl
x86_64-unknown-netbsd
x86_64-unknown-redox
```
在~/.cargo/config配置参数
在该文件的末尾加上下面这条交叉编译工具(linker里面的内容可能不对，这点不熟悉)
```
[target.x86_64-unknown-linux-musl]
linker = "x86_64-linux-musl-gcc"
```
安装在config配置的target.x86_64-unknown-linux-musl工具
```
$ rustup target add x86_64-unknown-linux-musl
```
准备源代码进行交叉编译
```
$ cargo build --target=x86_64-unknown-linux-musl
```
若是想要省略该参数则需要对config作如下的修改，以改变默认的构建目标
```
[build]
target = "x86_64-unknown-linux-musl"
```
另外，也可以给build --target x86_64-unknown-linux-musl 命令设置别名从而缩短构建命令。比如按下面的方式修改config文件后，就可以使用cargo build_linux来构建程序了
```
[alias]
build_linux = "build --target x86_64-unknown-linux-musl"
```

## 常见问题
- Blocking waiting for file lock on package cache
    - 删~/.cargo/.package-cache文件(rm -rf ~/.cargo/.package-cache)