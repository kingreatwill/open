
## 群辉
黑群辉安装引导:https://github.com/fbelavenuto/arpl

## 威联通

## 开源

### freeNAS

### openmediavault
简称 OMV

### U-NAS
国产免费

### XigmaNAS
### TrueNAS
### openfiler


## ONVIF

rtsp://192.168.31.183:554/cam/realmonitor?channel=1&subtype=0&unicast=true&proto=Onvif

rtsp://username:password@ip:port/cam/realmonitor?channel=1&subtype=0
说明:
username：用户名；
password：密码；
ip：为设备IP；
port：端口号默认为554，若为默认可不填写。
channel：通道号，起始为1。例如通道2，则为channel=2。
subtype：码流类型，主码流为0（即subtype=0），辅码流为1（即subtype=1）。
例如，请求某设备的通道1的主码流，Url如下：
rtsp://admin:lierdalux@192.168.1.108:554/cam/realmonitor?channel=1&subtype=1

可以用VLC直接打开此视频流！

乐橙云开放平台 https://open.imou.com/book/readme/model.html


- ONVIF Device Manager
- VLC
- shinobi