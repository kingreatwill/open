
## HTTP

### Accept-Ranges
https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Range_requests

服务器响应输出`Accept-Ranges` 首部（并且它的值不为“none”），那么表示该服务器支持范围请求。

发送一个 HEAD 请求来进行检测。
```sh
curl -I http://i.imgur.com/z4d4kWk.jpg

HTTP/1.1 200 OK
...
Accept-Ranges: bytes
Content-Length: 146515
```
Accept-Ranges: bytes 表示界定范围的单位是 bytes。
如果站点未发送 Accept-Ranges 首部，那么它们有可能不支持范围请求。一些站点会明确将其值设置为 "none"，以此来表明不支持。

### Range
客户端请求内容范围
"-H" 选项可以在请求中追加一个首部行
```
curl -vo /dev/null 'http://i.imgur.com/z4d4kWk.jpg' -r 0-300

curl -vo /dev/null http://i.imgur.com/z4d4kWk.jpg -H "Range: bytes=0-300"
```

`curl http://i.imgur.com/z4d4kWk.jpg -i -H "Range: bytes=0-1023"`
这样生成的请求如下：
```
GET /z4d4kWk.jpg HTTP/1.1
Host: i.imgur.com
Range: bytes=0-1023
```
服务器端会返回状态码为 206 Partial Content 的响应：
```
HTTP/1.1 206 Partial Content
Content-Range: bytes 0-1023/146515
Content-Length: 1024 
...
(binary content)
```
在这里，Content-Length 首部现在用来表示先前请求范围的大小（而不是整张图片的大小）。Content-Range 响应首部则表示这一部分内容在整个资源中所处的位置。

> 在请求成功的情况下，服务器会返回 206 Partial Content 状态码。
> 在请求的范围越界的情况下（范围值超过了资源的大小），服务器会返回 416 Requested Range Not Satisfiable （请求的范围无法满足）状态码。表示客户端错误.
> 在不支持范围请求的情况下，服务器会返回 200 OK 状态码。


Range 头部的格式有以下几种情况：
```
Range: <unit>=<range-start>-
Range: <unit>=<range-start>-<range-end>
Range: <unit>=<range-start>-<range-end>, <range-start>-<range-end>
Range: <unit>=<range-start>-<range-end>, <range-start>-<range-end>, <range-start>-<range-end>
```

对于多重范围`curl http://www.example.com -i -H "Range: bytes=0-50, 100-150"`
服务器返回 206 Partial Content 状态码和 Content-Type：multipart/byteranges; boundary=3d6b6a416f9b5 头部，Content-Type：multipart/byteranges 表示这个响应有多个 byterange。每一部分 byterange 都有他自己的 Content-type 头部和 Content-Range，并且使用 boundary 参数对 body 进行划分。
```
HTTP/1.1 206 Partial Content
Content-Type: multipart/byteranges; boundary=3d6b6a416f9b5
Content-Length: 282

--3d6b6a416f9b5
Content-Type: text/html
Content-Range: bytes 0-50/1270

<!doctype html>
<html>
<head>
    <title>Example Do
--3d6b6a416f9b5
Content-Type: text/html
Content-Range: bytes 100-150/1270

eta http-equiv="Content-type" content="text/html; c
--3d6b6a416f9b5--
```

### Content-Range
服务器响应首部 Content-Range 显示的是一个数据片段在整个文件中的位置。
```
Content-Range: <unit> <range-start>-<range-end>/<size>
Content-Range: <unit> <range-start>-<range-end>/*
Content-Range: <unit> */<size>
```
如果大小未知则用"*"表示

### If-Range
当（中断之后）重新开始请求更多资源片段的时候，必须确保自从上一个片段被接收之后该资源没有进行过修改。

If-Range HTTP 请求头字段用来使得 Range 头字段在一定条件下起作用：假如条件满足的话，条件请求就会生效，服务器会返回状态码为 206 Partial 的响应，以及相应的消息主体。假如条件未能得到满足，那么就会返回状态码为 200 OK 的响应，同时返回整个资源。该首部可以与 Last-Modified 验证器或者 ETag 一起使用，但是二者不能同时使用。
```
Range: bytes=1024-2047
If-Range: Wed, 21 Oct 2015 07:28:00 GMT
```