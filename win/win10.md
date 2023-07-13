win+↑↓←→ 是分屏快捷键

win+Ctrl+D 是新建虚拟桌面
win+tab 切换建虚拟桌面
CTRL + win+ ←→ 切换建虚拟桌面
win+ctrl+f4快捷键可以关闭当前虚拟桌面



ctrl+alt+tab快捷键：
跟alt+tab快捷键功能一样，但是松开键盘后，界面不会消失

cmd -> 右击窗口栏 -> 属性 -> “选项” -> “编辑选项” ->勾选“启用Ctrl键快捷方式”

win+w  调用便签、屏幕草图和随手画(白板)


ctrl+shift+T 恢复误删除网页

ctrl+backspace 删除光标前的一串 ,ctrl+delete 删除光标后的一串，ctrl+方向键 光标移动到一串字符的最前或者最后

Windows设置>设备>鼠标>"当我悬停在非活动窗口上方时对其进行滚动"改为开；在多窗口叠加的工作中，不用将工作窗口调整到最前面，就可以实现滚轮下拉页面

Alt +1：查看文件/文件夹属性；Alt +2是新建文件夹

Win + Q 快速打开搜索

### 禁用路径长度限制

环境变量PATH的最大长度是有限制的, 最大260个字符
https://learn.microsoft.com/en-us/windows/win32/fileio/maximum-file-path-limitation?tabs=registry

https://learn.microsoft.com/zh-cn/windows/win32/fileio/maximum-file-path-limitation?tabs=registry

将注册表的如下键值，从 0 改成 1 即可（本选项实际上是修改了 NTFS filesystem 的默认显示。在 windows 系统上，该限制是默认开启的。0 代表开启限制，即 false，1 代表关闭限制，即 true）
`Computer\HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\FileSystem\LongPathsEnabled`

```powershell
New-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Control\FileSystem" `
-Name "LongPathsEnabled" -Value 1 -PropertyType DWORD -Force
```