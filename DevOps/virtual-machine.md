


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