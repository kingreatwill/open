[Cache-Control](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Cache-Control)

HTTP协议的Cache -Control指定请求和响应遵循的缓存机制。
请求时的缓存指令包括: no-cache、no-store、max-age、 max-stale、min-fresh、only-if-cached等。
响应消息中的指令包括: public、private、no-cache、no- store、no-transform、must-revalidate、proxy-revalidate、max-age。

### 1. Cache-Control:
设置相对过期时间, max-age指明以秒为单位的缓存时间. 若对静态资源只缓存一次, 可以设置max-age的值为315360000000 (一万年). 比如对于提交的订单，为了防止浏览器回退重新提交，可以使用Cache-Control之no-store绝对禁止缓存，即便浏览器回退依然请求的是服务器，进而判断订单的状态给出相应的提示信息！

Http协议的cache-control的常见取值及其组合释义:
no-cache: 数据内容不能被缓存, 每次请求都重新访问服务器, 若有max-age, 则缓存期间不访问服务器.
no-store: 不仅不能缓存, 连暂存也不可以(即: 临时文件夹中不能暂存该资源).
private(默认): 只能在浏览器中缓存, 只有在第一次请求的时候才访问服务器, 若有max-age, 则缓存期间不访问服务器.
public: 可以被任何缓存区缓存, 如: 浏览器、服务器、代理服务器等.
max-age: 相对过期时间, 即以秒为单位的缓存时间(TTL).
no-cache, private: 打开新窗口时候重新访问服务器, 若设置max-age, 则缓存期间不访问服务器.
    - private, 正数的max-age: 后退时候不会访问服务器.
    - no-cache, 正数的max-age: 后退时会访问服务器.


### 2. Expires(绝对到期时间):
设置以分钟为单位的绝对过期时间, 优先级比Cache-Control低, 同时设置Expires和Cache-Control则后者生效. 也就是说要注意一点:  Cache-Control的优先级高于Expires

expires起到控制页面缓存的作用，合理配置expires可以减少很多服务器的请求, expires的配置可以在http段中或者server段中或者location段中.  比如控制图片等过期时间为30天, 可以配置如下:
```
location ~ \.(gif|jpg|jpeg|png|bmp|ico)$ {
           root /var/www/img/;
           expires 30d;
       }
```
再比如:
```
location ~ \.(wma|wmv|asf|mp3|mmf|zip|rar|swf|flv)$ {
              root /var/www/upload/;
              expires max;
      }
```

### 3. Last-Modified:
该资源的最后修改时间, 在浏览器下一次请求资源时, 浏览器将先发送一个请求到服务器上, 并附上If-Unmodified-Since头来说明浏览器所缓存资源的最后修改时间, 如果服务器发现没有修改, 则直接返回304(Not Modified)回应信息给浏览器(内容很少), 如果服务器对比时间发现修改了, 则照常返回所请求的资源. 

需要注意:
1) Last-Modified属性通常和Expires或Cache-Control属性配合使用, 因为即使浏览器设置缓存, 当用户点击”刷新”按钮时, 浏览器会忽略缓存继续向服务器发送请求, 这时Last-Modified将能够很好的减小回应开销.

2) ETag将返回给浏览器一个资源ID, 如果有了新版本则正常发送并附上新ID, 否则返回304， 但是在服务器集群情况下, 每个服务器将返回不同的ID, 因此不建议使用ETag.

以上描述的客户端浏览器缓存是指存储位置在客户端浏览器, 但是对客户端浏览器缓存的实际设置工作是在服务器上的资源中完成的. 虽然上面介绍了有关于客户端浏览器缓存的属性, 但是实际上对这些属性的设置工作都需要在服务器的资源中做设置. 通常有两种操作手段对浏览器缓存进行设置, 一个是通过页面指令声明来设置, 另外一个是通过编程方式来设置.

下面是相关页面设置Cache-Control头信息的几个简单配置:
例一:
```
if ($request_uri ~* "^/$|^/search/.+/|^/company/.+/") {
   add_header    Cache-Control  max-age=3600;
  }
```
个人理解的max-age意思是：客户端本地的缓存，在配置的生存时间内的，客户端可以直接使用，超出生存时间的，到服务器上取新数据。当然这些还要看客户端浏览器的设置。

例二:
```
location ~ .*\.(css|js|swf|php|htm|html )$ {
      add_header Cache-Control no-store;
}
```
例三:

```
location ~ .*\.(js|css)$ {
     expires 10d;
}
```

例四: 将html结尾的请求加上no-cache
```
location / {
    access_log /data/nginx/log/xxx.log api;
    root /home/www/html;
    if ($request_filename ~ .*\.(htm|html)$)
     {
            add_header Cache-Control no-cache;
     }
}
```

### 参考
[Nginx下关于缓存控制字段cache-control的配置说明 - 运维小结](https://www.cnblogs.com/kevingrace/p/10459429.html)