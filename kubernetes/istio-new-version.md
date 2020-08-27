## istio 1.5
- 改变的第一个版本
- 推翻了之前的架构设计，提出了“回归单体”的架构设计（将控制平面的所有组件组合并成一个单体结构 Istiod。）
- Envoy 与 WebAssembly 集成

## istio 1.6
Istio 1.6 版本其实是对 1.5 版本未完成工作的收尾，其中最大的简化工作是将原有组件的功能完全整合入 Istiod ，让 Istiod 更加完整，也彻底移除了 Citadel、Sidecar Injector 和 Galley。另外，添加了 istioctl install 命令来替代 manifest apply 的安装过程，用更直观、更精简的命令改善安装过程的体验。


## istio 1.7
- 专访 Christian Posta：Istio 1.7 将成为生产可用的最稳定版本
- Require Kubernetes 1.16+