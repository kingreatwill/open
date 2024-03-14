
## 云手机
https://github.com/budtmo/docker-android

> x86架构，部分软件不能运行


首先要保证你的机器支持KVM虚拟化; 浏览器中访问你的http://ip:6080
```
# 创建三星S10手机
docker run -d -p 6008:6008 -e EMULATOR_DEVICE="Samsung Galaxy S10" -e WEB_VNC=true -e WEB_VNC_PORT=6008 --device /dev/kvm --name android budtmo/docker-android:emulator_11.

# 创建三星S6手机
docker run --privileged -d -p 6080:6080 -p 5554:5554 -p 5555:5555 -e DEVICE="Samsung Galaxy S6" --name android-container budtmo/docker-android-x86-8.1
```


## 远程手机/远程控制

https://github.com/bk138/droidVNC-NG

https://github.com/bk138/droidVNC-NG/releases/download/v2.2.0/droidvnc-ng-2.2.0.apk

Windows远程安卓,输入手机安装droidVNC-NG后的IP+端口: https://www.realvnc.com/en/connect/download/viewer/windows/

Mac远程安卓手机: 屏幕共享->输入手机安装droidVNC-NG后的IP+端口

### 在PC上控制Android设备/Android实时投屏软件
https://github.com/pdone/FreeControl

https://github.com/karma9874/AndroRAT


可以使用todesk远程连接电脑,进行操作
Android实时投屏软件，此应用程序提供USB(或通过TCP/IP)连接的Android设备的显示和控制。它不需要任何root访问权限
https://gitee.com/Barryda/QtScrcpy

1. 进入手机系统设置，找到手机版本号页面（一般在“关于手机”里面），连续点击5~6下，开启开发者模式
2. 找到并进入“开发人员选项”（一般在“系统和更新”里面，如果找不到的话直接搜索），打开“USB调试”功能
3. 在电脑的QtScrcpy上点击“一键USB连接”
4. 在手机上允许USB调试
5. 连接成功后会自动添加设备，双击即可打开屏幕共享
6. 电脑设置从不进入睡眠状态，并且关闭手机的锁屏密码