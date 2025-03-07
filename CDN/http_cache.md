


### 强缓存（Strong Cache）
1. 定义
强缓存直接告诉浏览器：在缓存过期前，无需与服务器通信，直接使用本地缓存。
由服务器通过响应头 Cache-Control 和 Expires 控制。

2. 响应头
- `Cache-Control: max-age=3600`
表示资源在 3600 秒（1小时） 内有效（优先级高于 Expires）。

- `Expires: Thu, 31 Dec 2030 23:59:59 GMT`
指定一个绝对过期时间（依赖于客户端本地时间，可能存在误差）。

3. Nginx 配置示例
```
location /static/ {
    # 设置强缓存：1年内有效
    add_header Cache-Control "public, max-age=31536000";
    expires 1y;
}
```

4. 行为
浏览器首次请求资源时，服务器返回资源并附带缓存头。
后续请求时，浏览器直接读取本地缓存（状态码 200 (from disk cache)），不发送请求到服务器。

### 协商缓存（Weak Cache）
1. 定义
协商缓存要求浏览器 每次向服务器验证缓存是否过期，若未过期则返回 304 Not Modified，继续使用本地缓存。
由服务器通过响应头 Last-Modified 和 ETag 控制。

2. 响应头
- `Last-Modified: Wed, 21 Oct 2023 07:28:00 GMT`
表示资源最后修改时间（精度为秒，可能因时间同步问题失效）。

- `ETag: "5d8c72a5-264"`
资源的唯一标识符（哈希值或版本号），精度更高。

3. Nginx 配置示例
```
location /dynamic/ {
    # 启用协商缓存（默认已支持，无需显式配置）
    add_header Last-Modified "";
    etag on;
}
```

4. 行为
浏览器首次请求资源时，服务器返回资源并附带 Last-Modified 或 ETag。
后续请求时，浏览器通过以下请求头验证缓存：
- `If-Modified-Since: [Last-Modified值]`
向服务器询问资源是否在指定时间后修改过。

- `If-None-Match: [ETag值]`
向服务器验证资源的 ETag 是否变化。
若资源未修改，服务器返回 304 Not Modified，浏览器继续使用缓存；若已修改，返回新资源（状态码 200）。



特性|	强缓存	| 协商缓存
---|---|---
通信成本|	无网络请求（直接读缓存）|	需发送请求验证缓存
响应状态码|	200 (from disk cache)|	304 Not Modified
优先级|	优先于协商缓存|	强缓存过期后触发
适用资源|	长期不变的静态资源	|频繁更新的动态资源


```
# 图片、字体等强缓存
location ~* \.(jpg|png|gif|woff2)$ {
    expires 1y;
    add_header Cache-Control "public, max-age=31536000";
}

# HTML 文件禁用强缓存（总是协商）
location ~* \.html$ {
    add_header Cache-Control "no-cache, must-revalidate";
}
```
