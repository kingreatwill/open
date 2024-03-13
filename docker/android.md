
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

### 在PC上控制Android设备
https://github.com/pdone/FreeControl

https://github.com/karma9874/AndroRAT