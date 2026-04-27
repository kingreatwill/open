
## 什么是esim
嵌入式SIM（又称 eSIM 或 eUICC）技术，移动用户可以在没有实体 SIM 卡的情况下，下载运营商配置文件并激活运营商服务。该技术是由 GSMA 推动的全球规范，支持在任何移动设备上进行远程 SIM 配置 (RSP)。从 Android9 开始，Android 框架为访问 eSIM 和管理 eSIM 上的订阅配置文件提供了标准 API。借助这些 eUICC API，第三方可以在支持 eSIM 的 Android 设备上开发自己的运营商应用和 Local Profile Assistant (LPA)。


传统的 eSIM 是固化在主板上的芯片。而我们使用的方案是将 eUICC (Embedded Universal Integrated Circuit Card) 的文件系统跑在一张标准的 SIM 卡单片机上。
由于国内外ARA-M证书签发标准不同，导致设备与配置文件的授权验证不互通。这意味着国内运营商的eSIM无法写入该卡，海外设备也无法写入国内eSIM。
购买的小白卡必须确认带ARA-M证书，否则需ROOT。

https://euicc-manual.osmocom.org/



## 开源写卡
> 切卡: STK菜单 -> SIM Tool Kit
### OpenEUICC 和 EasyEUICC
https://gitea.angry.im/PeterCxy/OpenEUICC
https://github.com/estkme-group/openeuicc

https://openeuicc.com/

#### EasyEUICC
https://easyeuicc.org/zh-hans/
https://github.com/BastelPichi/easyeuicc-9esim/releases

最后，你的安卓手机需要支持OMAPI，不是所有的安卓手机都可以用的。
有时OpenEUICC可能不是一个选项。也许你的设备无法解锁。EasyEUICC不需要系统权限，但作为权衡，它只能管理厂商明确允许的可拆卸eSIM芯片。
- ESTKme
- 9eSIM
“EasyEUICC 兼容卡片” 是指符合 GSMA SGP.22 规范，且必须包含此特定 ARA-M SHA-1 值的产品：2A2FA878BC7C3354C2CF82935A5945A3EDAE4AFA。