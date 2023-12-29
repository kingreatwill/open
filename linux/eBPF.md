# eBPF（extened Berkeley Packet Filter）

eBPF（extened Berkeley Packet Filter）是一种内核技术，它允许开发人员在不修改内核代码的情况下运行特定的功能。eBPF 的概念源自于 Berkeley Packet Filter（BPF），后者是由贝尔实验室开发的一种网络过滤器，可以捕获和过滤网络数据包。

出于对更好的 Linux 跟踪工具的需求，eBPF 从 [dtrace](https://illumos.org/books/dtrace/chp-intro.html) 中汲取灵感，dtrace 是一种主要用于 Solaris 和 BSD 操作系统的动态跟踪工具。与 dtrace 不同，Linux 无法全面了解正在运行的系统，因为它仅限于系统调用、库调用和函数的特定框架。在[Berkeley Packet Filter  (BPF)](https://www.kernel.org/doc/html/latest/bpf/index.html)（一种使用内核 VM 编写打包过滤代码的工具）的基础上，一小群工程师开始扩展 BPF 后端以提供与 dtrace 类似的功能集。 eBPF 诞生了。2014 年随 Linux 3.18 首次限量发布，充分利用 eBPF 至少需要 Linux 4.4 以上版本。


eBPF 比起传统的 BPF 来说，传统的 BPF 只能用于网络过滤，而 eBPF 则可以用于更多的应用场景，包括网络监控、安全过滤和性能分析等。另外，eBPF 允许常规用户空间应用程序将要在 Linux 内核中执行的逻辑打包为字节码，当某些事件（称为挂钩）发生时，内核会调用 eBPF 程序。此类挂钩的示例包括系统调用、网络事件等。用于编写和调试 eBPF 程序的最流行的工具链称为 [BPF 编译器集合 (BCC)](https://github.com/iovisor/bcc)，它基于 LLVM 和 CLang。

BCC（BPF Compiler Collection）是一套开源的工具集，可以在 Linux 系统中使用 BPF（Berkeley Packet Filter）程序进行系统级性能分析和监测。BCC 包含了许多实用工具，如：

- bcc-tools：一个包含许多常用的 BCC 工具的软件包。
- bpftrace：一个高级语言，用于编写和执行 BPF 程序。
- tcptop：一个实时监控和分析 TCP 流量的工具。
- execsnoop：一个用于监控进程执行情况的工具。
- filetop：一个实时监控和分析文件系统流量的工具。
- trace：一个用于跟踪和分析函数调用的工具。
- funccount：一个用于统计函数调用次数的工具。
- opensnoop：一个用于监控文件打开操作的工具。
- pidstat：一个用于监控进程性能的工具。
- profile：一个用于分析系统 CPU 使用情况的工具。


eBPF 有一些类似的工具。例如，SystemTap 是一种开源工具，可以帮助用户收集 Linux 内核的运行时数据。它通过动态加载内核模块来实现这一功能，类似于 eBPF。另外，DTrace 是一种动态跟踪和分析工具，可以用于收集系统的运行时数据，类似于 eBPF 和 SystemTap。

## eBPF 相关的知名的开源项目
### Katran
Facebook 高性能 4 层负载均衡器 [Katran](https://github.com/facebookincubator/katran) ；
### Cilium
[Cilium](https://github.com/cilium/cilium) 为下一代微服务 ServiceMesh 打造了具备API感知和安全高效的容器网络方案；底层主要使用 eBPF 中 XDP 和 TC；
### IO Visor
IO Visor 项目开源的 [BCC](https://github.com/iovisor/bcc) 、 [BPFTrace](https://github.com/iovisor/bpftrace) 和 [Kubectl-Trace](https://github.com/iovisor/kubectl-trace) ： BCC 提供了更高阶的抽象，可以让用户采用 Python、C++ 和 Lua 等高级语言快速开发 BPF 程序； BPFTrace 采用类似于 awk 语言快速编写 eBPF 程序； Kubectl-Trace 则提供了在 kubernetes 集群中使用 BPF 程序调试的方便操作；

### CloudFlare
CloudFlare 公司开源的 [eBPF Exporter](https://github.com/cloudflare/ebpf_exporter) 和 [bpf-tools](https://github.com/cloudflare/bpftools) ： eBPF Exporter 将 eBPF 技术与监控 Prometheus 紧密结合起来； bpf-tools 可用于网络问题分析和排查；