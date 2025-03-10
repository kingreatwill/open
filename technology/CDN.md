
[其他CDN知识见](../CDN/README.md)

## CDN基础知识

### CNAME、加速域名、源站地址与回源host之间的关系

#### 1. CNAME
首先需要了解，CNAME 是什么东西。

1.1 A记录
即Address记录，它并不是一个IP或者一个域名，我们可以把它理解为一种指向关系：

域名 www.xx.com → 111.111.111.111

主机名 DD → 222.222.222.222

可以理解为，最终的域名与IP的对应关系这条记录，就是A记录

1.2 CNAME
为什么要区分A记录和CNAME？我们可以把CNAME记录叫做别名记录，就是小名

比如A记录为：
www.baidu.com → 111.111.111.111

那么可能有多个CNAME记录
www.100fen.com → www.baidu.com
www.baifen.com → www.baidu.com

所以大概理解了吧，CNAME就是你主域名A记录的小名

CNAME指向A记录，A记录指向具体的IP地址。

一个网址可以有多个CNAME，可以理解为就是域名转发

#### 2. 加速域名
加速域名是接入CDN的域名，例如使用www.baidu.com域名接入CDN，那么加速域名就是www.baidu.com

#### 3. 源站地址
顾名思义，就是用户的源站主机地址。

当选择自建源时，源站地址也有两种类型：域名（一个）和IP地址（可多个），并且都支持端口。

示例一：
源站类型：自建源
源站地址：115.115.115.115

示例二：
源站类型：自建源
源站地址：115.115.115.115:8080

示例三：
源站类型：自建源
源站地址：source.baidu.com

示例四：
源站类型：自建源
源站地址：source.baidu.com:8080

源站地址填写为域名最大的好处是，当你的源站拥有多个运营商的IP，例如source.baidu.com有三个A记录，分别是100.100.100.100（电信）、101.101.101.101（联通）、102.102.102.102（移动），那么当CDN回源时，会根据用户的来源运营商选择对应的A记录。当访问者为电信用户，回源时则会选择100.100.100.100（电信）来回源，这样就可以避免跨运营商回源，造成回源延迟等问题。

注意：如果源站地址填写为域名时，此域名的作用仅仅是用于DNS解析，例如source.baidu.com对应的A记录为100.100.100.100，那最终源站地址则是100.100.100.100。

建议：能选择源站地址为域名最好是选择为域名。

#### 4. 回源host
回源host也可以说是回源域名，比如源站拥有多个站点：bbs.baidu.com、blog.baidu.com、api.baidu.com，那回源host就是指定到哪个站点上获取资源，具体详情参考以下几个示例。

示例一：
源站类型：自建源
源站地址：115.115.115.115
回源host：blog.baidu.com

当CDN回源时，会到115.115.115.115这台主机上的blog.baidu.com站点拉取资源。

示例二：
源站类型：自建源
源站地址：115.115.115.115
回源host：api.baidu.com

当CDN回源时，会到115.115.115.115这台主机上的api.baidu.com站点拉取资源。
注意：如果填写的回源host在源站上不存在时，则会到源站上的默认站点拉取资源。
建议：当源站拥有多个站点时，一定要选择正确的域名，否则CDN拉取到的资源可能不是你想要的。

#### DNS和CDN整体流程的总结
比如我们请求 www.baidu.com 域名

- 1. 首先，浏览器会从自身的DNS缓存中去查找（chrome://net-internals/#dns），如果没有则进行下一步

- 2. 然后，浏览器会先从操作系统里的DNS缓存中找，windows系统中，命令行 ipconfig/displaydns 查看，linux上的NSCD缓存服务；如果没有则进行下一步

- 3. 从计算机host文件里找；如果没有则进行下一步

- 4. 请求本地域名服务器（可以认为是 阿里云等域名供应商），
发现阿里云里面有进行过配置，CNMAE记录： www.baidu.com → cdn.baidu.com ，所以此时告诉浏览器转为请求 cdn.baidu.com

此时，浏览器转为请求cdn.baidu.com ，上面的1-3步还得再来一遍。

1-3步骤重复

- 4. 请求本地域名服务器（可以认为是 阿里云等域名供应商），发现阿里云里面有进行过配置，A记录：cdn.baidu.com → 222.222.222.222 ，然后就把 IP 222.222.222.222 返回给浏览器。

- 5. 浏览器得到了IP地址，注意这个IP地址，实际上是CDN负载均衡服务器的地址。继续请求这个地址

- 6. 请求进入到了CDN负载均衡服务器后，服务器会根据算法策略等，返回一个最合适的文件缓存服务器IP地址，至于怎么选择合适的，看下面的优化

- 7. 浏览器访问文件缓存服务器IP地址，最后得到文件资源

#### 附加题目
题目：加速域名为：www.baidu.com，源站地址填写为域名：source.baidu.com，并且此域名拥有三个IP：100.100.100.100（电信）、101.101.101.101（联通）、102.102.102.102（移动），回源host为blog.baidu.com。

问题一：请将回源步骤描述下。

问题二：当访问者是联通用户时，请将回源步骤描述下。

问题三：当源站不存在blog.baidu.com站点时，请将回源步骤描述下。

问题四：当访问者是教育网用户时，请将回源步骤描述下。

#### 附加答案
问题一：
先将source.baidu.com域名做DNS解析，得到具体的IP地址（100.100.100.100）。
CDN使用回源host（blog.baidu.com）到100.100.100.100这台主机上拉取blog.baidu.com站点的资源。
blog.baidu.com站点返回资源给CDN。
回源流程：访问者（www.baidu.com） -> CDN（回源） -> blog.baidu.com（100.100.100.100） -> CDN（返回资源） -> 访问者（获取资源）

问题二：
先将source.baidu.com域名做DNS解析，得到101.101.101.101（CDN会帮其选择最优的IP，因为访问者是联通用户，当然是到联通IP回源更好，这样就不会出现跨运营商问题）。
CDN使用回源host（blog.baidu.com）到101.101.101.101这台主机上拉取blog.baidu.com站点的资源。
blog.baidu.com站点返回资源给CDN。
回源流程：访问者（www.baidu.com） -> CDN（回源） -> blog.baidu.com（101.101.101.101） -> CDN（返回资源） -> 访问者（获取资源）

问题三：
先将source.baidu.com域名做DNS解析，得到具体的IP地址（100.100.100.100）。
CDN使用回源host（blog.baidu.com）到100.100.100.100这台主机上拉取默认站点的资源。
100.100.100.100这台主机的默认站点返回资源给CDN。
回源流程：访问者（www.baidu.com） -> CDN（回源） -> blog.baidu.com（100.100.100.100） -> CDN（返回资源） -> 访问者（获取资源）

问题四：
先将source.baidu.com域名做DNS解析，得到具体的IP地址102.102.102.102（因为源站没有合适对应的教育网运营商IP，所以CDN会自动帮其选择最优的IP，至于哪个是最优的就无法直接断定了，交由CDN判断）。
CDN使用回源host（blog.baidu.com）到102.102.102.102这台主机上拉取默认站点的资源。
blog.baidu.com站点返回资源给CDN。
回源流程：访问者（www.baidu.com） -> CDN（回源） -> blog.baidu.com（102.102.102.102） -> CDN（返回资源） -> 访问者（获取资源）

#### 结论总结

加速域名为用户接入的CDN域名。
源站地址为用户的源站，当源站地址为域名时，此域名仅用于做DNS解析。
当源站拥有多个站点，回源时则选择回源host指定站点拉取资源。
当源站不存在回源host域名时，则到源站的默认站点拉取资源。

#### 参考链接
https://juejin.im/post/6854573212425814030
https://cloud.tencent.com/developer/article/1195058