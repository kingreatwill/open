
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
