
# WebAssembly
https://github.com/WebAssembly

## WebAssembly Micro Runtime
2019 年 5 月，英特尔公司在 GitHub 上开源了 [WebAssembly Micro Runtime](https://github.com/bytecodealliance/wasm-micro-runtime) 项目（简称 WAMR）。

WASM 未来在嵌入式设备到云端都将具有极其广泛的应用空间

WAMR 项目包括以下三部分功能：

- iwasmVM 内核，是 WASM 字节码的执行引擎，支持解释器、提前编译 （AoT） 和实时编译 （JIT）多种模式；

- WASM 应用程序 API 和多实例应用框架；

- WASM 应用程序的动态管理；

应用场景：

- 可信计算环境（TEE）、Trusted FaaS、联邦计算
- 硬件加速
- 手机、IoT、智能设备的小程序引擎
- 超轻量级 Serverless 容器
- 区块链智能合约
- 工业控制逻辑引擎
- 游戏引擎
- 支持可独立开发、动态装载的固件模块

## Ruffle （Flash 播放器模拟器）
https://github.com/ruffle-rs/ruffle