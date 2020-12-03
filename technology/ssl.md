
## SSL
[11种免费获取SSL证书的方式](https://www.toutiao.com/i6883395048126284292)



### ACME 协议
Let’s Encrypt 使用 ACME 协议来验证您对给定域名的控制权并向您颁发证书。要获得 Let’s Encrypt 证书，您需要选择一个要使用的 ACME 客户端软件。

https://letsencrypt.org/zh-cn/docs/client-options/

https://github.com/letsencrypt/website/

#### certbot - ACME 客户端
[让网站永久拥有HTTPS - 申请免费SSL证书并自动续期](https://blog.csdn.net/xs18952904/article/details/79262646)
https://www.cnblogs.com/dissipate/p/13606006.html
http://www.mntm520.com/post/48
也可以用certbot的docker镜像来生成证书


#### zerossl - ACME 客户端
https://zerossl.com/

## acme.sh

这个脚本acme.sh，实现了 acme 协议, 可以帮你持续自动从Letsencrypt更新CA证书
https://github.com/acmesh-official/acme.sh

安装
https://github.com/acmesh-official/acme.sh/wiki/%E8%AF%B4%E6%98%8E
`curl get.acme.sh | sh`