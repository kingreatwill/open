https://openwrt.org/
https://github.com/openwrt/openwrt

### openwrt
#### fanchmwrt
https://github.com/fanchmwrt/fanchmwrt

#### iStoreOS
iStoreOS 目标是提供一个人人会用的路由兼轻 NAS 系统，不管是作为路由还是 NAS，你都有相似的操作体验。
https://github.com/istoreos



### 小米路由器

#### AX6000手动开启ssh权限
Redmi AX6000固件版本：1.0.60
默认网关：192.168.31.1

https://cdn.cnbj1.fds.api.mi-img.com/xiaoqiang/rom/rb06/miwifi_rb06_firmware_7ddeb_1.0.60.bin

https://cdn.cnbj1.fds.api.mi-img.com/xiaoqiang/rom/rb06/miwifi_rb06_firmware_847e9_1.0.48.bin


路由器初始密码的计算：找到Redmi AX6000的SN码，并在此[计算初始密码](https://miwifi.dev/ssh)。得到初始密码，将其永久保存下来，路由器恢复出厂设置后会用到；

打开路由器的web端管理界面，输入管理员密码，从上方地址栏即可获得stok
> 每次路由器重启，stok值都会改变。

1. 固件降级从web端的管理界面将Redmi AX6000的固件版本降级到1.0.60。


2. 开启Redmi AX6000的开发者模式
```
1. 更改路由器的crash分区，使其进入到开发者模式

    在浏览器地址栏输入：

http://192.168.31.1/cgi-bin/luci/;stok=token/api/misystem/set_sys_time?timezone=%20%27%20%3B%20zz%3D%24%28dd%20if%3D%2Fdev%2Fzero%20bs%3D1%20count%3D2%202%3E%2Fdev%2Fnull%29%20%3B%20printf%20%27%A5%5A%25c%25c%27%20%24zz%20%24zz%20%7C%20mtd%20write%20-%20crash%20%3B%20
    其中，将stok=token中的token替换成路由器的stok。

    例如，我的stok为030b24d39b1a4a549aa12dac23c52313，则我应输入到浏览器地址栏的链接为：

http://192.168.31.1/cgi-bin/luci/;stok=030b24d39b1a4a549aa12dac23c52313/api/misystem/set_sys_time?timezone=%20%27%20%3B%20zz%3D%24%28dd%20if%3D%2Fdev%2Fzero%20bs%3D1%20count%3D2%202%3E%2Fdev%2Fnull%29%20%3B%20printf%20%27%A5%5A%25c%25c%27%20%24zz%20%24zz%20%7C%20mtd%20write%20-%20crash%20%3B%20
2. 特殊方法重启路由器

    在浏览器地址栏输入：

http://192.168.31.1/cgi-bin/luci/;stok=token/api/misystem/set_sys_time?timezone=%20%27%20%3b%20reboot%20%3b%20
    其中，将stok=token中的token替换成路由器的stok即可。
```

3. 设置Redmi AX6000的Bdata参数
```
1. 设置参数telnet_en、 ssh_en、uart_en

    在浏览器地址栏输入：

http://192.168.31.1/cgi-bin/luci/;stok=token/api/misystem/set_sys_time?timezone=%20%27%20%3B%20bdata%20set%20telnet_en%3D1%20%3B%20bdata%20set%20ssh_en%3D1%20%3B%20bdata%20set%20uart_en%3D1%20%3B%20bdata%20commit%20%3B%20
    其中，将stok=token中的token替换成路由器新的stok即可。

2. 特殊方法重启路由器

    在浏览器地址栏输入：

http://192.168.31.1/cgi-bin/luci/;stok=token/api/misystem/set_sys_time?timezone=%20%27%20%3b%20reboot%20%3b%20
    其中，将stok=token中的token替换成路由器新的stok即可。
```
4. 通过telnet开启ssh

```
telnet 192.168.31.1


进入telnet后，输入指令开启ssh并修改密码为admin

指令:
bdata set boot_wait=on bdata commit nvram set ssh_en=1 nvram set telnet_en=1 nvram set uart_en=1 nvram set boot_wait=on nvram commit sed -i 's/channel=.*/channel="debug"/g' /etc/init.d/dropbear /etc/init.d/dropbear restart echo -e 'admin\nadmin' | passwd root
```

5. 永久开启 ssh
目前从telnet开启的ssh，路由器重启会失效，因此添加一个开启自动运行的脚本，来实现自动开启 ssh。缺点就是恢复出厂设置或重新刷机后需要重新添加。
注：此处需要路由器联网才可以执行成功。
```
ssh root@192.168.31.1

mkdir /data/auto_ssh && cd /data/auto_ssh curl -O https://cdn.jsdelivr.net/gh/lemoeo/AX6S@main/auto_ssh.sh chmod +x auto_ssh.sh uci set firewall.auto_ssh=include uci set firewall.auto_ssh.type='script' uci set firewall.auto_ssh.path='/data/auto_ssh/auto_ssh.sh' uci set firewall.auto_ssh.enabled='1' uci commit firewall

```

6. 关闭Redmi AX6000的开发者模式
```
新开一个Windows Terminal窗口，ssh连接进去：

ssh root@192.168.31.1
输入密码admin，即可登入ssh。

输入指令清除crash分区并重启：

mtd erase crash reboot

```
7. 修复时区异常
```
用上述方法开启ssh后，可能会导致路由器时区异常，在ssh中执行以下指令即可：

uci set system.@system[0].timezone='CST-8' uci set system.@system[0].webtimezone='CST-8' uci set system.@system[0].timezoneindex='2.84' uci commit

```
8. 注意事项
```
每次升级固件或重置固件都需要执行以下指令：

sed -i 's/channel=.*/channel="debug"/g' /etc/init.d/dropbear /etc/init.d/dropbear restart echo -e 'admin\nadmin' | passwd root
新开一个Windows Terminal窗口，ssh连接进去：

注：此处是新开一个窗口，不要从telnet界面进行ssh连接。

ssh root@192.168.31.1
输入密码admin，即可登入ssh。

登入后输入以下指令：

mkdir /data/auto_ssh && cd /data/auto_ssh curl -O https://cdn.jsdelivr.net/gh/lemoeo/AX6S@main/auto_ssh.sh chmod +x auto_ssh.sh uci set firewall.auto_ssh=include uci set firewall.auto_ssh.type='script' uci set firewall.auto_ssh.path='/data/auto_ssh/auto_ssh.sh' uci set firewall.auto_ssh.enabled='1' uci commit firewall
```

#### xmir-patcher脚本开启ssh权限
使用 xmir-patcher 工具开启SSH
https://github.com/openwrt-xiaomi/xmir-patcher

https://gitee.com/longzheng0315/Xiaomi-Router-patcher

RA72 - 小米路由器 AX6000
安装教程
Windows 用户
```
下载工具包
下载并解压项目文件到任意目录
运行工具
双击运行 run.bat 文件
自动环境配置
工具会自动使用内置的 Python 环境
无需手动安装 Python 或依赖包
```

Linux / macOS 用户
```
系统要求
# 确保已安装 Python 3.8+ 和 OpenSSL
python3 --version  # 应显示 3.8 或更高版本
openssl version    # 确认 OpenSSL 已安装
安装依赖

# Ubuntu/Debian
sudo apt update
sudo apt install python3 python3-pip python3-venv openssl

# CentOS/RHEL
sudo yum install python3 python3-pip openssl

# macOS (使用 Homebrew)
brew install python3 openssl
下载并运行

# 下载项目
git clone https://github.com/longzheng268/Xiaomi-Router-patcher.git
cd Xiaomi-Router-patcher

# 运行工具
chmod +x run.sh
./run.sh
```

运行工具后会看到主菜单：
```
==========================================================

Xiaomi MiR Patcher  

 1 - Set IP-address (current value: 192.168.31.1)
 2 - Connect to device (install exploit)
 3 - Read full device info
 4 - Create full backup
 5 - Install EN/RU languages
 6 - Install permanent SSH
 7 - Install firmware (from directory "firmware")
 8 - {{{ Other functions }}}
 9 - [[ Reboot device ]]
 0 - Exit

Select:
```

执行操作
```

必要的前置步骤
设置IP地址 (选项1)
选择 1 → 输入路由器IP地址192.168.31.1 → 按回车确认

连接设备并安装漏洞利用 (选项2)
选择 2 → 工具会自动检测设备型号并安装相应的漏洞利用程序
⚠️ 重要: 这是所有后续操作的前提，必须先成功执行此步骤！

常用操作流程
完整备份 (推荐首先执行)
选择 4 → 工具会创建完整的固件备份，保存在 backup 目录中

启用SSH访问
选择 6 → 工具会永久启用SSH访问，默认用户名/密码: root/root

安装语言包
选择 5 → 安装英文/俄文/中文语言包，让路由器界面支持更多语言
```

```
# 检查SSH服务状态
ssh root@192.168.31.1

修改默认密码: SSH启用后立即修改默认密码

passwd root

# 如不需要Telnet，建议禁用
/etc/init.d/telnet disable


# 限制SSH访问的IP范围
iptables -A INPUT -p tcp --dport 22 -s 192.168.31.0/24 -j ACCEPT
iptables -A INPUT -p tcp --dport 22 -j DROP
```


固件地址 https://opt.cn2qq.com/
KoolCenter梅林固件发布页： https://www.koolcenter.com/fw
在线编译OpenWrt固件：https://openwrt.ai/
https://mao.fan/
路由器查询： https://mao.fan/select 这个页面，点击支持刷机，会筛选出来支持刷机的路由器，还可以筛选无线协议、品牌等。