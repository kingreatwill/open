# windows 命令
## 参考
https://www.thewindowsclub.com/
https://docs.microsoft.com/zh-cn/windows-server/administration/windows-commands/windows-commands
https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/windows-commands

## 网络相关
### ipconfig
[英文](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/ipconfig)
[中文](https://docs.microsoft.com/zh-cn/windows-server/administration/windows-commands/ipconfig)
显示所有当前 TCP/IP 网络配置值并刷新动态主机配置协议 (DHCP) 和域名系统 (DNS) 设置。 在没有参数的情况下使用， ipconfig 显示 Internet 协议版本 4 (IPv4) 以及所有适配器的 IPv6 地址、子网掩码和默认网关。

```
ipconfig /?
ipconfig
ipconfig /flushdns
```

### ping
通过向另一台 TCP/IP 计算机发送 Internet 控制消息协议 (ICMP) 回响请求消息来验证 IP 级连接。 将显示相应的回响回复消息以及往返时间。 ping 是用于排查连接性、可访问性和名称解析的主要 TCP/IP 命令。 使用没有参数的情况下，此命令显示帮助内容。
你还可以使用此命令测试计算机的计算机名和 IP 地址。 如果成功完成了 IP 地址的 ping 操作，但不对计算机名称进行 ping 操作，则可能存在名称解析问题。 在这种情况下，请确保你指定的计算机名可以通过本地主机文件进行解析，方法是使用域名系统 (DNS) 查询，或通过 NetBIOS 名称解析技术。

> 仅当在 "网络连接" 中网络适配器的属性中将 "Internet 协议" (TCP/IP) 协议安装为组件时，此命令才可用。

```
ping /? 
# 以下 -参数也是可以的 如：ping -t 127.0.0.1 
ping /a 127.0.0.1 # 显示主机名
ping /n 10 127.0.0.1  # ping10次
ping /l 1000 127.0.0.1 # ping发送1000字节，默认32字节
ping /t 127.0.0.1  # 一直ping
```
### netstat
显示处于活动状态的 TCP 连接、计算机正在侦听的端口、以太网统计信息、IP 路由表、用于 IP、ICMP、TCP 和 UDP 协议的 IPv4 统计信息 () 和 ipv6 统计信息 (ipv6、ICMPv6、TCP over IPv6 和 UDP over IPv6 协议) 。 使用没有参数的情况下，此命令显示活动 TCP 连接。
> 仅当在 "网络连接" 中网络适配器的属性中将 "Internet 协议" (TCP/IP) 协议安装为组件时，此命令才可用。

```
netstat /?
-o 显示PID
-n 显示TCP
-e 以太网
-s 按协议显示统计信息
netstat -s -p tcp # 若要仅显示 TCP 协议的统计信息
```

状态	指示 TCP 连接的状态，包括：
- CLOSE_WAIT
- CLOSED
- 端建立
- FIN_WAIT_1
- FIN_WAIT_2
- LAST_ACK
- 侦听
- SYN_RECEIVED
- SYN_SEND
- TIMED_WAIT


### telnet
与运行 telnet 服务器服务的计算机通信。
```
telnet /?
telnet /f telnetlog.txt telnet.microsoft.com 44 # /f 记录日志文件
```
### tracert
通过向目标发送 Internet 控制消息协议) (（ (ICMP) 回响请求或 ICMPv6 消息），确定到达目标的路径。 显示的路径为源主机与目标之间路径中的路由器的 near/端路由器接口列表。 近/侧接口是最接近路径中发送主机的路由器接口。 在没有参数的情况下使用，tracert 显示帮助。

```
tracert /?
tracert 127.0.0.1
tracert /d 127.0.0.1 # 并防止将每个 IP 地址解析为其名称
```
## Windows
### wmic
显示交互式命令行界面中的 WMI 信息。
```
wmic /?

WMIC 已弃用。
[全局开关] <命令>

可以使用以下全局开关:
/NAMESPACE           别名在其上操作的命名空间的路径。
/ROLE                包含别名定义的角色的路径。
/NODE                别名在其上操作的服务器。
/IMPLEVEL            客户端模拟级别。
/AUTHLEVEL           客户端身份验证级别。
/LOCALE              客户端应使用的语言 ID。
/PRIVILEGES          启用或禁用所有权限。
/TRACE               将调试信息输出到 stderr。
/RECORD              记录所有输入命令和输出内容。
/INTERACTIVE         设置或重置交互模式。
/FAILFAST            设置或重置 FailFast 模式。
/USER                会话期间要使用的用户。
/PASSWORD            登录会话时要使用的密码。
/OUTPUT              指定输出重定向模式。
/APPEND              指定输出重定向模式。
/AGGREGATE           设置或重置聚合模式。
/AUTHORITY           指定连接的 <授权类型>。
/?[:<BRIEF|FULL>]    用法信息。

有关特定全局开关的详细信息，请键入: switch-name /?


子命令	说明
class	从 WMIC 的默认别名模式转义以直接访问 WMI 架构中的类。
path	从 WMIC 的默认别名模式进行转义，以直接访问 WMI 架构中的实例。
context	显示所有全局开关的当前值。

如：wmic context

NAMESPACE    : root\cimv2
ROLE         : root\cli
NODE(S)      : BOBENTERPRISE
IMPLEVEL     : IMPERSONATE
[AUTHORITY   : N/A]
AUTHLEVEL    : PKTPRIVACY
LOCALE       : ms_409
PRIVILEGES   : ENABLE
TRACE        : OFF
RECORD       : N/A
INTERACTIVE  : OFF
FAILFAST     : OFF
OUTPUT       : STDOUT
APPEND       : STDOUT
USER         : N/A
AGGREGATE    : ON


若要将命令行使用的语言 ID 更改为英语 (区域设置 ID 409) ，请键入：
wmic /locale:ms_409
```
### rundll32
加载并运行32位动态链接库 (Dll) 。 没有适用于 Rundll32.exe 的可配置设置。 为使用 rundll32.exe 命令运行的特定 DLL 提供了帮助信息。
必须从提升的命令提示符运行 rundll32.exe 命令。 若要打开提升的命令提示符，请单击 " 开始"，右键单击 " 命令提示符"，然后单击 "以 管理员身份运行"。
```
Rundll32 <DLLname>

Rundll32.exe printui.dll,PrintUIEntry # 显示打印机用户界面
or
rundll32 printui.dll PrintUIEntry # 显示打印机用户界面
or
rundll32 printui PrintUIEntry
or
rundll32 printui,PrintUIEntry

rundll32.exe printui.dll,PrintUIEntry/?

# https://www.thewindowsclub.com/rundll32-shortcut-commands-windows

rundll32.exe keymgr.dll,KRShowKeyMgr # 凭证，查看存储的用户名和密码

RunDll32.exe shell32.dll,Control_RunDLL appwiz.cpl,,0 # 打开程序和功能

# Content Advisor（好像要有Internet explorer ）
RunDll32.exe msrating.dll,RatingSetupUI

RunDll32.exe shell32.dll,Control_RunDLL  # 控制面板

Delete Temporary Internet Files
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 8

Delete Cookies
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 2

Delete History
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 1

Delete Form Data
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 16

Delete Passwords
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 32

Delete All
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 255

Delete All + files and settings stored by Add-ons:
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 4351

Date and Time Properties
RunDll32.exe shell32.dll,Control_RunDLL timedate.cpl

Display Settings
RunDll32.exe shell32.dll,Control_RunDLL access.cpl,,3

Device Manager
RunDll32.exe devmgr.dll DeviceManager_Execute

Folder Options – General
RunDll32.exe shell32.dll,Options_RunDLL 0

Folder Options – File Types
RunDll32.exe shell32.dll,Control_Options 2

Folder Options – Search
RunDll32.exe shell32.dll,Options_RunDLL 2

Folder Options – View
RunDll32.exe shell32.dll,Options_RunDLL 7

Forgotten Password Wizard
RunDll32.exe keymgr.dll,PRShowSaveWizardExW

Hibernate
RunDll32.exe powrprof.dll,SetSuspendState

Internet Explorer’s Internet Properties dialog box.
RunDll32.exe Shell32.dll,Control_RunDLL Inetcpl.cpl,,4

Keyboard Properties
RunDll32.exe shell32.dll,Control_RunDLL main.cpl @1

Lock Screen
RunDll32.exe user32.dll,LockWorkStation

Mouse Button – Swap left button to function as right
Rundll32 User32.dll,SwapMouseButton

Mouse Properties Dialog Box
Rundll32 Shell32.dll,Control_RunDLL main.cpl @0,0

Map Network Drive Wizard
Rundll32 Shell32.dll,SHHelpShortcuts_RunDLL Connect

Network Connections
RunDll32.exe shell32.dll,Control_RunDLL ncpa.cpl

Organize IE Favourites
Rundll32.exe shdocvw.dll,DoOrganizeFavDlg

Open With Dialog Box
Rundll32 Shell32.dll,OpenAs_RunDLL Any_File-name.ext

Printer User Interface
Rundll32 Printui.dll,PrintUIEntry /?

Printer Management Folder.
Rundll32 Shell32.dll,SHHelpShortcuts_RunDLL PrintersFolder

Power Options
RunDll32.exe Shell32.dll,Control_RunDLL powercfg.cpl

Process Idle Tasks
RunDll32.exe advapi32.dll,ProcessIdleTasks

Regional and Language Options
Rundll32 Shell32.dll,Control_RunDLL Intl.cpl,,0

Stored Usernames and Passwords
RunDll32.exe keymgr.dll,KRShowKeyMgr

Safely Remove Hardware Dialog Box
Rundll32 Shell32.dll,Control_RunDLL HotPlug.dll

Sound Properties Dialog Box
Rundll32 Shell32.dll,Control_RunDLL Mmsys.cpl,,0

System Properties Box
Rundll32 Shell32.dll,Control_RunDLL Sysdm.cpl,,3

System Properties – Advanced
RunDll32.exe shell32.dll,Control_RunDLL sysdm.cpl,,4

System Properties: Automatic Updates
RunDll32.exe shell32.dll,Control_RunDLL sysdm.cpl,,5

Taskbar Properties
RunDll32.exe shell32.dll,Options_RunDLL 1

User Accounts
RunDll32.exe shell32.dll,Control_RunDLL nusrmgr.cpl

Unplug/Eject Hardware
RunDll32.exe shell32.dll,Control_RunDLL hotplug.dll

Windows Security Center
RunDll32.exe shell32.dll,Control_RunDLL wscui.cpl

Windows – About
RunDll32.exe SHELL32.DLL,ShellAboutW

Windows Fonts Installation Folder
Rundll32 Shell32.dll,SHHelpShortcuts_RunDLL FontsFolder

Windows Firewall
RunDll32.exe shell32.dll,Control_RunDLL firewall.cpl

Wireless Network Setup
RunDll32.exe shell32.dll,Control_RunDLL NetSetup.cpl,@0,WNSW
```

## 文件目录相关命令
### attrib

help attrib

- 查看文件属性
attrib d:\2.txt

- 为文件添加一个隐藏属性和一个只读属性。
attrib +r +h d:\2.txt