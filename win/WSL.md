WSL2的卸载操作如下：
wslconfig /l
# 从列表中选择要卸载的发行版（例如Ubuntu）并键入命令
wslconfig /u Ubuntu

以下命令不行
`lxrun /uninstall /full`