# 安装JDK
```
yum -y install java-1.8.0-openjdk*
```
配置环境变量 打开 vim /etc/profile添加一下内容
```
export JAVA_HOME=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.161-0.b14.el7_4.x86_64
export PATH=$PATH:$JAVA_HOME/bin
```
修改完成之后，使其生效'
```
source /etc/profile
```
输入java -version 返回版本信息则安装正常。