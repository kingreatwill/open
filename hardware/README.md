[TOC]

# IoT相关

[Awesome IoT](https://github.com/phodal/awesome-iot)
[Awesome Open IoT](https://github.com/Agile-IoT/awesome-open-iot)
## Raspberry
Raspberry Pi 内核
https://github.com/raspberrypi/linux c 7.2k

## mbed
https://github.com/ARMmbed/mbed-os  c 3.5k 
mbedOS是ARM自己打造、主打IoT的一整套软件解决方案，是一个针对ARM CortexM系列处理器的嵌入式开源生态。

## Arduino
开放源码电子原型平台，使用户能够创建交互式电子对象。
https://github.com/arduino/Arduino java python 10.8k

### PlatformIO
PlatformIO 是开源的物联网开发生态系统。提供跨平台的代码构建器、集成开发环境（IDE），兼容 Arduino 和 MBED。 PlatformIO 使用纯 Python 开发

## OpenWrt
OpenWrt 可以被描述为一个嵌入式的 Linux 发行版
https://github.com/openwrt/openwrt c 5.8k

> HyperWRT 是一个无线路由器的固件

### wifi-densepose
https://github.com/ruvnet/wifi-densepose
根据WIFI的CSI检测出物体

## RT-Thread
国产嵌入式实时操作系统
RT-Thread是一个集实时操作系统（RTOS）内核、中间件组件和开发者社区于一体的技术平台，具有极小内核、稳定可靠、简单易用、高度可伸缩、组件丰富等特点。RT-Thread拥有一个国内最大的嵌入式开源社区，同时被广泛应用于能源、车载、医疗、消费电子等多个行业，累积装机量达数千万台，成为国人自主开发、国内最成熟稳定和装机量最大的开源RTOS。

https://github.com/RT-Thread/rt-thread c 5.1k

## RTOS
实时操作系统（Real-time operating system, RTOS）
RTOS是一个嵌入式开发套件，包括一个小型但功能强大的操作系统，它为资源受限的设备提供了可靠的、超快的性能。它易于使用，经过市场验证，已在全球超过62亿设备上部署。
> 典型的实时操作系统有VxWorks，RT-Thread，uCOS，QNX，WinCE等。
### FreeRTOS
FreeRTOS是一个迷你的实时操作系统内核。
https://www.freertos.org/

## LiteOS
https://github.com/LiteOS/LiteOS c 4.1k
Huawei LiteOS 是华为面向IoT领域，构建的"统一物联网操作系统和中间件软件平台"，以轻量级（内核小于10k）、低功耗（1节5号电池最多可以工作5年），快速启动，互联互通，安全等关键能力，为开发者提供 "一站式" 完整软件平台，有效降低开发门槛、缩短开发周期。

## RTLinux
RTLinux（AReal-Time Linux,亦称作实时Linux）是Linux中的一种实时操作系统。它由新墨西哥矿业及科技学院的V. Yodaiken开发。目前，RTLinux有一个由社区支持的免费版本，称为RTLinux Free，以及一个来自FSMLabs的商业版本，称作RTLinux Pro。

https://github.com/clrkwllms/rt-linux c

## VxWorks

VxWorks 是美国 Wind River System 公司（ 以下简称风河公司 ，即 WRS 公司）推出的一个实时操作系统。Tornado 是WRS 公司推出的一套实时操作系统开发环境，类似MicrosoftVisual C，但是提供了更丰富的调试、仿真环境和工具。
> NASA的火星车是用的VxWorks
## tinyOS
https://github.com/tinyos

## contiki
## uC/OS

## ThingsBoard 物联网平台
https://github.com/thingsboard/thingsboard



# 其它

## NAS （Network Attached Storage：网络附属存储）
### TrueNAS
https://github.com/truenas
### FreeNAS
https://github.com/freenas
### openmediavault
openmediavault is the next generation network attached storage (NAS) solution based on Debian Linux.
https://github.com/openmediavault/openmediavault

## U盘工具

### 全能维护 U 盘工具：Ventoy
https://github.com/ventoy/Ventoy

# 硬件知识
## 指令集CISC和RISC
CISC：主要是两家，包括IntelCPU（非安腾系列）、AMD CPU。
RISC：服务器领域主要是IBM Power系列、Sun Spark系列，消费级的代表是ARM架构的CPU。

CISC(Complex Instruction Set Computers，复杂指令集计算集)和RISC(Reduced Instruction Set Computers)是两大类主流的CPU指令集类型，其中CISC以Intel，AMD的X86 CPU为代表，而RISC以ARM，IBM Power为代表。RISC的设计初衷针对CISC CPU复杂的弊端，选择一些可以在单个CPU周期完成的指令，以降低CPU的复杂度，将复杂性交给编译器。举一个例子，CISC提供的乘法指令，调用时可完成内存a和内存b中的两个数相乘，结果存入内存a，需要多个CPU周期才可以完成；而RISC不提供“一站式”的乘法指令，需调用四条单CPU周期指令完成两数相乘：内存a加载到寄存器，内存b加载到寄存器，两个寄存器中数相乘，寄存器结果存入内存a。按照此思路，早期的设计出的RISC指令集，指令数是比CISC少些，单后来，很多RISC的指令集中指令数反超了CISC，因此，引用指令的复杂度而非数量来区分两种指令集。

当然，CISC也是要通过操作内存、寄存器、运算器来完成复杂指令的。它在实现时，是将复杂指令转换成了一个微程序，微程序在制造CPU时就已存储于微服务存储器。一个微程序包含若干条微指令（也称微码），执行复杂指令时，实际上是在执行一个微程序。这也带来两种指令集的一个差别，微程序的执行是不可被打断的，而RISC指令之间可以被打断，所以理论上RISC可更快响应中断。

在此，总结一下CISC和RISC的主要区别：
1. CISC的指令能力强，单多数指令使用率低却增加了CPU的复杂度，指令是可变长格式；RISC的指令大部分为单周期指令，指令长度固定，操作寄存器，只有Load/Store操作内存
2. CISC支持多种寻址方式；RISC支持方式少
3. CISC通过微程序控制技术实现；RISC增加了通用寄存器，硬布线逻辑控制为主，是和采用流水线
4. CISC的研制周期长
5. RISC优化编译，有效支持高级语言

## 服务器分类
1. 按照外形
- 塔式服务器
- 机架服务器
- 刀片服务器
- 高密服务器

2. 按照指令
- 复杂指令集CISC（x86架构是典型代表）
- 精简指令集RISC（非x86架构）

3. 按处理器数量
- 单路服务器(1颗CPU)
- 双路服务器(2颗CPU)
- 多路服务器(4颗CPU及以上)

2017年7月，Intel正式发布了代号为Purley的新一代服务器平台，包括代号为Skylake的新一代Xeon CPU，命名为英特尔至强可扩展处理器(Intel Xeon Scalable Processor，SP)，也宣告了延续4代的至强E5/E7系列命名方式的终结。

Xeon至强可扩展处理器不再以E7、E5的方式来划分定位，而代之以铂金(Platinum)、金(Gold)、银(Silver)、铜(Bronze)的方式。Skylake是新命名方式的第一代产品，Cascade Lake是是二代，共用Purley平台。


**大型机**：普通人很少接触，用于大规模计算的计算机系统.大型机通常用于政府、银行、交通、保险公司和大型制造企业。特点是处理数据能力强大、稳定性和安全性又非常高
**小型机**：往往应用于金融、电力、电信等行业，这些用户看重的是Unix操作系统和专用服务器RAS特性、纵向扩展性和高并发访问下的出色处理能力。这些特性是普通的X86服务器很难达到的，所以在数据库等关键应用一般都采用“高大贵”的小型机方案。
**x86服务器**：采用CISC架构处理器。1978年6月8日，Intel发布了一款新型的微处理器8086，意味着x86架构的诞生，而x86作为特定微处理器执行计算机语言的指令集，定义了芯片的基本使用规则。
**ARM服务器**：ARM全称为Advanced RISC Machine，即进阶精简指令集机器。ARM是RISC微处理器的代表作之一，最大的特点在于节能。


四大服务器CPU厂商

RISC/Unix 服务器
- IBM（已成为小型机市场的统治地位）
- SUN/Oracle(最具创新的技术领导者)
- ARM（数据中心领域的创新者）

x86/IA 架构服务器
- Inter/AMD(x86服务器市场的领导者，延续摩尔定律的新神话)


硬盘类型
1. 按接口
- ATA/IDE
- SATA/NL SAS
- SCSI
- SAS
- FC

2. 按盘径
- 5.25英寸
- 3.5
- 2.5
- 1.8

3. 按介质
- 机械硬盘(HDD)
- 固态硬盘(SDD)

4. 按功能
- 桌面级(MTBF 50万小时)
- 企业级(MTBF > 100万小时)

> MTBF，即平均无故障工作时间，英文全称是“Mean Time Between Failure”。

> SATA RPM=5400/7200 , MTBF=120万
> SAS RPM=15000/10000 , MTBF=160万
> NL-SAS RPM=5400/7200 , MTBF=120万
> SSD RPM=NA , MTBF=200万

[RAID](../computer/RAID.md)


服务器网卡

网络接口控制器（英语：network interface controller，NIC），又称网络接口控制器，网络适配器（network adapter），网卡（network interface card），或局域网接收器（LAN adapter），是一块被设计用来允许计算机在计算机网络上进行通讯的计算机硬件。

网卡在TCP/IP的模型中，工作在物理层和数据链路层，用来接收和发送数据。除了数据的收发，网卡还有一些其他功能：

1、代表固定的地址：数据发送出去，发给谁，又从哪里接收。这都是通过IP区分的
2、数据的封装、解封：比如寄一封信，信封里的信纸是data，信封是帧头和帧尾。
3、链路管理：因为以太网是共享链路的，在使用时候可能会有其他人也在发送数据。如果同时发送，就会产生冲突，这就要求在发送的时候，检测链路的状态是否空闲；
4、数据的编码和译码：在物理介质中，传送的是电平或光信号。这时就需要将二进制数据转换成电平信号或光信号。
5、发送和接收数据

总线分类：PCIe、USB、ISA、PCI，ISA/PCI等总线是比较早期的网络总线，现在已很少用了，USB接口的网卡主要用在消费级电子中。
结构形态：集成网卡（LOM）、PCIe标卡网卡、Mezz卡。
应用类型：按网卡所应用的的计算机类型来区分，可以将网卡分为应用于工作站的网卡和应用于服务器的网卡。

网卡接口
电口，PC上常见到的那种网口接口，这种接口叫RJ45，使用的是普通的网线
光口，用于连接光模块，网卡上用于插光模块的接口，我们叫光笼子。