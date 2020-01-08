[TOC]
# 域名解析--什么是A记录、别名记录(CNAME)、MX记录、TXT记录、NS记录

## A (Address) 记录

是用来指定主机名（或域名）对应的IP地址记录。用户可以将该域名下的网站服务器指向到自己的web server上。同时也可以设置您域名的二级域名。

## 别名(CNAME)记录

也被称为规范名字。这种记录允许您将多个名字映射到同一台计算机。 通常用于同时提供WWW和MAIL服务的计算机。例如，有一台计算机名为“host.mydomain.com”（A记录）。 它同时提供WWW和MAIL服务，为了便于用户访问服务。可以为该计算机设置两个别名（CNAME）：WWW和MAIL。 这两个别名的全称就是“www.mydomain.com”和“mail.mydomain.com”。实际上他们都指向“host.mydomain.com”。 同样的方法可以用于当您拥有多个域名需要指向同一服务器IP，此时您就可以将一个域名做A记录指向服务器IP然后将其他的域名做别名到之前做A记录的域名上，那么当您的服务器IP地址变更时您就可以不必麻烦的一个一个域名更改指向了 只需要更改做A记录的那个域名其他做别名的那些域名的指向也将自动更改到新的IP地址上了。

### 如何检测CNAME记录？

1、进入命令状态；（开始菜单 - 运行 - CMD[回车]）；

2、输入命令" nslookup -q=cname 这里填写对应的域名或二级域名"，查看返回的结果与设置的是否一致即可。 



## MX（Mail Exchanger）记录

是邮件交换记录，它指向一个邮件服务器，用于电子邮件系统发邮件时根据 收信人的地址后缀来定位邮件服务器。例如，当Internet上的某用户要发一封信给 user@mydomain.com 时，该用户的邮件系统通过DNS查找mydomain.com这个域名的MX记录，如果MX记录存在， 用户计算机就将邮件发送到MX记录所指定的邮件服务器上。



## 什么是TXT记录？

TXT记录一般指为某个主机名或域名设置的说明，如：

1）admin IN TXT "jack, mobile:13800138000"；

2）mail IN TXT "邮件主机, 存放在xxx ,管理人：AAA"，Jim IN TXT "contact: abc@mailserver.com"

也就是您可以设置 TXT ，以便使别人联系到您。



### 如何检测TXT记录？

1、进入命令状态；（开始菜单 - 运行 - CMD[回车]）；

2、输入命令" nslookup -q=txt 这里填写对应的域名或二级域名"，查看返回的结果与设置的是否一致即可。 



## 什么是NS记录？

ns记录全称为Name Server 是一种域名服务器记录，用来明确当前你的域名是由哪个DNS服务器来进行解析的。

## 代码
```cmd
nslookup www.baidu.com
```
```go
package main

import (
	"fmt"
	"net"
)

func main() {
	main1()
	main2()
	main3()
}

// 查找CNAME
func main1() {
	cname, err := net.LookupCNAME("www.baidu.com")
	if err != nil {
		panic(err)
	} // dig +short www.baidu.com cname
	fmt.Printf("%s\n", cname) //www.a.shifen.com.
}

// 查找MX
func main2() {
	mxs, err := net.LookupMX("cndns.com")
	if err != nil {
		panic(err)
	}
	// dig +short cndns.com mx
	for _, mx := range mxs {
		fmt.Printf("%s %v\n", mx.Host, mx.Pref) // mx1.chengmail.cn. 5
	}
}

// 查找TXT
func main3() {

	txts, err := net.LookupTXT("cndns.com")
	if err != nil {
		panic(err)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		//dig +short cndns.com txt
		fmt.Printf("%s\n", txt) // v=spf1 include:spf.chengmail.cn ~all
	}
}

```