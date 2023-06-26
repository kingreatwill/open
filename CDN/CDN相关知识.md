
## Apache Traffic Server(ATS)
ATS最重要的是CDN缓存能力,反向代理是其第二功能.

[ATS CDN 搭建](https://www.taterli.com/8527/)

## CDN基础知识
[CDN基础知识-CNAME、加速域名、源站地址与回源host之间的关系](https://blog.csdn.net/qq_43442524/article/details/110479939)

## 回源SNI
如果您的源站IP绑定了多个域名，且CDN回源协议为HTTPS时，需配置回源SNI，在回源SNI内指明所请求的具体域名，服务器会根据该域名正确地返回对应的SSL证书。
根据 SNI（Server Name Indication）信息来获取域名证书和私钥信息。

[配置回源SNI](https://help.aliyun.com/document_detail/111152.html?spm=a2c4g.125959.0.0.258a2240mTnVex)

## 其它
### 第一公里
是指万维网流量向用户传送的第一个出口,即网站服务器接入互联网所能提供的带宽

## 最后一公里
是指万维网流量向用户传送的最后一段接入链路,即用户接入带宽