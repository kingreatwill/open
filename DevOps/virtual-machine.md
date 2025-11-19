
虚拟化工具：VMware、VirtualBox、Hyper-V、Xen、KVM、multipass 

## 开源oVirt
https://github.com/oVirt/ovirt-engine

## Windows10 wsl2 开启Hype-v后如何兼容Virtualbox和VMware

### 与Virtualbox兼容
- Virtualbox升级到6.0以上或者下载最新的版本
- 使用管理员的方式打开命令提示符，并切换到Virtualbox的按照命令，执行下面的命令
```
VBoxManage setextradata global "VBoxInternal/NEM/UseRing0Runloop" 0
```

### 与VMware兼容
VMware Workstation 20H1预览版已经兼容了Hype-V，VMware workstation 20H1预览版下载地址：

https://download3.vmware.com/software/wkst/file/VMware-workstation-full-e.x.p-15679048.exe

https://blogs.vmware.com/workstation/2020/05/vmware-workstation-now-supports-hyper-v-mode.html

[VMware Workstation 16 发布](https://sysin.org/article/Download-VMware-Workstation-16/) 支持Hype-V


## VMware
付费, 好用

## VirtualBox
免费
## multipass
https://github.com/canonical/multipass
https://multipass.run/
Multipass 是一个轻量虚拟机管理器，是由 Ubuntu 运营公司 Canonical 所推出的开源项目。运行环境支持 Linux、Windows、macOS。在不同的操作系统上，使用的是不同的虚拟化技术。
在 Linux 上使用的是 KVM、
Window 上使用 Hyper-V、
macOS 中使用 HyperKit 以最小开销运行VM，支持在笔记本模拟小型云。

> 官网上最新描述: Multipass 在 Windows 上使用 Hyper-V，在 macOS 上使用 QEMU 和 HyperKit，在 Linux 上使用 LXD，以实现最小的开销和最快的启动时间。将虚拟机管理程序切换到Virtualbox是一件轻而易举的事。

## QEMU 
https://github.com/qemu/qemu
https://gitlab.com/qemu-project/qemu.git
https://www.qemu.org/contribute/

## KVM、webKVM
- [Virtual Machine for the Web](https://github.com/leaningtech/webvm)
基于内核的虚拟机（KVM）是一种内建于 Linux® 的开源虚拟化技术。具体而言，KVM 可帮助您将 Linux 转变为虚拟机监控程序，使主机计算机能够运行多个隔离的虚拟环境，即虚拟客户机或虚拟机（VM）。

[KVM](https://www.linux-kvm.org/page/Main_Page) 是 Linux 的一部分。Linux 2.6.20 或更新版本包括 KVM。KVM 于 2006 年首次公布，并在一年后合并到主流 Linux 内核版本中。
### Barrier
https://github.com/debauchee/barrier

## lxcfs容器资源视图隔离 for k8s/docker
https://github.com/lxc/lxcfs
比如我们代码获取runtime.CPU() 获取的是主机上的CPU核数, 而不是pod的限制的CPU核数, 需要使用lxcfs进行隔离


## 其它
https://github.com/tiny-pilot/tinypilot

### proxmox
https://www.proxmox.com/en/
https://sysin.org/blog/proxmox-ve-8/

https://git.proxmox.com/

### Android
手机跑操作系统
https://github.com/xoureldeen/Vectras-VM-Android

# "Windows Emulator"，即Windows模拟器

## Wine
Wine其实为"Wine Is Not anEmulator"的递归缩写,Wine不是Windows模拟器，而是运用API转换技术实做出Linux对应到Windows相对应的函数来调用DLL以运行Windows程序。
`wine 软件全名.exe`

