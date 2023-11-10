
## openresty

### 镜像内部信息
OpenResty 默认安装位置：`/usr/local/openresty/`
安装目录中 Nginx 相关文件：`/usr/local/openresty/nginx/`
默认服务指向 Web 文件夹：`/usr/local/openresty/nginx/html/`
映射关系：`/bin/openresty -> /usr/local/openresty/nginx/sbin/nginx` 和 `/bin/opm -> /usr/local/openresty/bin/opm`
默认配置文件位置（后续的配置会覆盖这里的内容）：`/etc/nginx/conf.d/`
`/etc/nginx/conf.d/default.conf`
```conf
# nginx.vh.default.conf  --  docker-openresty
#
# This file is installed to:
#   `/etc/nginx/conf.d/default.conf`
#
# It tracks the `server` section of the upstream OpenResty's `nginx.conf`.
#
# This config (and any other configs in `etc/nginx/conf.d/`) is loaded by
# default by the `include` directive in `/usr/local/openresty/nginx/conf/nginx.conf`.
#
# See https://github.com/openresty/docker-openresty/blob/master/README.md#nginx-config-files
#


server {
    listen       80;
    server_name  localhost;

    #charset koi8-r;
    #access_log  /var/log/nginx/host.access.log  main;

    location / {
        root   /usr/local/openresty/nginx/html;
        index  index.html index.htm;
    }

    #error_page  404              /404.html;

    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/local/openresty/nginx/html;
    }

    # proxy the PHP scripts to Apache listening on 127.0.0.1:80
    #
    #location ~ \.php$ {
    #    proxy_pass   http://127.0.0.1;
    #}

    # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
    #
    #location ~ \.php$ {
    #    root           /usr/local/openresty/nginx/html;
    #    fastcgi_pass   127.0.0.1:9000;
    #    fastcgi_index  index.php;
    #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
    #    include        fastcgi_params;
    #}

    # deny access to .htaccess files, if Apache's document root
    # concurs with nginx's one
    #
    #location ~ /\.ht {
    #    deny  all;
    #}
}

```

原始的起点配置文件: `/usr/local/openresty/nginx/conf/nginx.conf`
```conf
# nginx.conf  --  docker-openresty
#
# This file is installed to:
#   `/usr/local/openresty/nginx/conf/nginx.conf`
# and is the file loaded by nginx at startup,
# unless the user specifies otherwise.
#
# It tracks the upstream OpenResty's `nginx.conf`, but removes the `server`
# section and adds this directive:
#     `include /etc/nginx/conf.d/*.conf;`
#
# The `docker-openresty` file `nginx.vh.default.conf` is copied to
# `/etc/nginx/conf.d/default.conf`.  It contains the `server section
# of the upstream `nginx.conf`.
#
# See https://github.com/openresty/docker-openresty/blob/master/README.md#nginx-config-files
#

#user  nobody;
#worker_processes 1;

# Enables the use of JIT for regular expressions to speed-up their processing.
pcre_jit on;



#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;

#pid        logs/nginx.pid;


events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;

    # Enables or disables the use of underscores in client request header fields.
    # When the use of underscores is disabled, request header fields whose names contain underscores are marked as invalid and become subject to the ignore_invalid_headers directive.
    # underscores_in_headers off;

    #log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
    #                  '$status $body_bytes_sent "$http_referer" '
    #                  '"$http_user_agent" "$http_x_forwarded_for"';

    #access_log  logs/access.log  main;

        # Log in JSON Format
        # log_format nginxlog_json escape=json '{ "timestamp": "$time_iso8601", '
        # '"remote_addr": "$remote_addr", '
        #  '"body_bytes_sent": $body_bytes_sent, '
        #  '"request_time": $request_time, '
        #  '"response_status": $status, '
        #  '"request": "$request", '
        #  '"request_method": "$request_method", '
        #  '"host": "$host",'
        #  '"upstream_addr": "$upstream_addr",'
        #  '"http_x_forwarded_for": "$http_x_forwarded_for",'
        #  '"http_referrer": "$http_referer", '
        #  '"http_user_agent": "$http_user_agent", '
        #  '"http_version": "$server_protocol", '
        #  '"nginx_access": true }';
        # access_log /dev/stdout nginxlog_json;

    # See Move default writable paths to a dedicated directory (#119)
    # https://github.com/openresty/docker-openresty/issues/119
    client_body_temp_path /var/run/openresty/nginx-client-body;
    proxy_temp_path       /var/run/openresty/nginx-proxy;
    fastcgi_temp_path     /var/run/openresty/nginx-fastcgi;
    uwsgi_temp_path       /var/run/openresty/nginx-uwsgi;
    scgi_temp_path        /var/run/openresty/nginx-scgi;

    sendfile        on;
    #tcp_nopush     on;

    #keepalive_timeout  0;
    keepalive_timeout  65;

    #gzip  on;

    include /etc/nginx/conf.d/*.conf;

    # Don't reveal OpenResty version to clients.
    # server_tokens off;
}

```
> https://github.com/openresty/docker-openresty/blob/master/README.md#nginx-config-files

### docker
proxy.conf
```
server {
    listen       8080;
    access_log  /var/log/nginx/access.log;
    location / {
	resolver 119.29.29.29;
        proxy_pass http://$http_host$uri$is_args$args;
        proxy_set_header Host $host;
    }
}
```
```sh
docker run -d --name openresty -v /data/dockerv/openresty/nginx.conf:/usr/local/openresty/nginx/conf/nginx.conf -v /data/dockerv/openresty/conf:/etc/nginx/conf.d -v /data/dockerv/openresty/www:/usr/local/openresty/nginx/html  -v /data/dockerv/openresty/logs:/var/log/nginx --restart always --network host openresty:v0.0.1
```

> 镜像在下面有命令,  `openresty/openresty:1.21.4.1-0-bullseye`

### 支持https
目前为止, 可以实现http的正向代理,不支持代理 Https 网站
作为 web_server Nginx 当然是可以处理 ssl 的，但作为proxy则是不行的。因为nginx不支持CONNECT，收到“CONNECT /:443 HTTP/1.1”后会报一个包含“client sent invalid request while reading client request line,” 的错误。因为 CONNECT 是正向代理的特性。

例：
访问：# curl -I -x 192.168.10.154:9999 'https://www.baidu.com/?tn=93380420_hao_pg'
日志：192.168.10.X - [04/Nov/2017:10:23:46 +0800] "CONNECT www.baidu.com:443 HTTP/1.1" 400 173 "-" "-" - - - - "-"


那么，如何让nginx的正向代理，既支持http又支持https的代理访问呢？





需要安装模块：[ngx_http_proxy_connect_module](https://github.com/chobits/ngx_http_proxy_connect_module)

```
server {

    listen 8080;

    # dns resolver used by forward proxying
    resolver 8.8.8.8;

    # forward proxy for CONNECT request
    proxy_connect;
    proxy_connect_allow 443 80;
    proxy_connect_connect_timeout 10s;
    proxy_connect_read_timeout 10s;
    proxy_connect_send_timeout 10s;

    # forward proxy for non-CONNECT request
    location / {
        auth_basic           "authentication";
        auth_basic_user_file htpasswd;
        proxy_pass http://$http_host$uri$is_args$args;
        proxy_set_header Host $host;
    }
}
```

密码文件`conf.d/htpasswd`
```
# comment
name1:password1
name2:password2:comment
name3:password3
```
#### docker-openresty携带参数安装
https://github.com/openresty/docker-openresty#building-from-source

```
git clone git@github.com:openresty/docker-openresty.git

cd docker-openresty

docker build --build-arg RESTY_EVAL_PRE_CONFIGURE="wget -O ngx_http_proxy_connect_module-0.0.5.tar.gz  https://github.com/chobits/ngx_http_proxy_connect_module/archive/refs/tags/v0.0.5.tar.gz && tar zxf ngx_http_proxy_connect_module-0.0.5.tar.gz" --build-arg RESTY_CONFIG_OPTIONS_MORE="--add-module=/tmp/ngx_http_proxy_connect_module-0.0.5" --build-arg RESTY_PCRE_OPTIONS="--with-pcre-jit && patch -d build/nginx-1.21.4/ -p1 < /tmp/ngx_http_proxy_connect_module-0.0.5/patch/proxy_connect_rewrite_102101.patch " -f jammy/Dockerfile -t openresty:v0.0.1 . > build.log 2>&1
```

Dockerfile
```
# Dockerfile - Ubuntu Jammy
# https://github.com/openresty/docker-openresty

ARG RESTY_IMAGE_BASE="ubuntu"
ARG RESTY_IMAGE_TAG="jammy"

FROM ${RESTY_IMAGE_BASE}:${RESTY_IMAGE_TAG}

LABEL maintainer="Evan Wies <evan@neomantra.net>"

# Docker Build Arguments
ARG RESTY_IMAGE_BASE="ubuntu"
ARG RESTY_IMAGE_TAG="jammy"
ARG RESTY_VERSION="1.21.4.2"
ARG RESTY_LUAROCKS_VERSION="3.9.2"
ARG RESTY_OPENSSL_VERSION="1.1.1w"
ARG RESTY_OPENSSL_PATCH_VERSION="1.1.1f"
ARG RESTY_OPENSSL_URL_BASE="https://www.openssl.org/source"
ARG RESTY_PCRE_VERSION="8.45"
ARG RESTY_PCRE_BUILD_OPTIONS="--enable-jit"
ARG RESTY_PCRE_SHA256="4e6ce03e0336e8b4a3d6c2b70b1c5e18590a5673a98186da90d4f33c23defc09"
ARG RESTY_J="1"
ARG RESTY_CONFIG_OPTIONS="\
    --with-compat \
    --with-file-aio \
    --with-http_addition_module \
    --with-http_auth_request_module \
    --with-http_dav_module \
    --with-http_flv_module \
    --with-http_geoip_module=dynamic \
    --with-http_gunzip_module \
    --with-http_gzip_static_module \
    --with-http_image_filter_module=dynamic \
    --with-http_mp4_module \
    --with-http_random_index_module \
    --with-http_realip_module \
    --with-http_secure_link_module \
    --with-http_slice_module \
    --with-http_ssl_module \
    --with-http_stub_status_module \
    --with-http_sub_module \
    --with-http_v2_module \
    --with-http_xslt_module=dynamic \
    --with-ipv6 \
    --with-mail \
    --with-mail_ssl_module \
    --with-md5-asm \
    --with-sha1-asm \
    --with-stream \
    --with-stream_ssl_module \
    --with-threads \
    "
ARG RESTY_CONFIG_OPTIONS_MORE=""
ARG RESTY_LUAJIT_OPTIONS="--with-luajit-xcflags='-DLUAJIT_NUMMODE=2 -DLUAJIT_ENABLE_LUA52COMPAT'"
ARG RESTY_PCRE_OPTIONS="--with-pcre-jit"

ARG RESTY_ADD_PACKAGE_BUILDDEPS=""
ARG RESTY_ADD_PACKAGE_RUNDEPS=""
ARG RESTY_EVAL_PRE_CONFIGURE=""
ARG RESTY_EVAL_POST_DOWNLOAD_PRE_CONFIGURE=""
ARG RESTY_EVAL_POST_MAKE=""

# These are not intended to be user-specified
ARG _RESTY_CONFIG_DEPS="--with-pcre \
    --with-cc-opt='-DNGX_LUA_ABORT_AT_PANIC -I/usr/local/openresty/pcre/include -I/usr/local/openresty/openssl/include' \
    --with-ld-opt='-L/usr/local/openresty/pcre/lib -L/usr/local/openresty/openssl/lib -Wl,-rpath,/usr/local/openresty/pcre/lib:/usr/local/openresty/openssl/lib' \
    "

LABEL resty_image_base="${RESTY_IMAGE_BASE}"
LABEL resty_image_tag="${RESTY_IMAGE_TAG}"
LABEL resty_version="${RESTY_VERSION}"
LABEL resty_luarocks_version="${RESTY_LUAROCKS_VERSION}"
LABEL resty_openssl_version="${RESTY_OPENSSL_VERSION}"
LABEL resty_openssl_patch_version="${RESTY_OPENSSL_PATCH_VERSION}"
LABEL resty_openssl_url_base="${RESTY_OPENSSL_URL_BASE}"
LABEL resty_pcre_version="${RESTY_PCRE_VERSION}"
LABEL resty_pcre_build_options="${RESTY_PCRE_BUILD_OPTIONS}"
LABEL resty_pcre_sha256="${RESTY_PCRE_SHA256}"
LABEL resty_config_options="${RESTY_CONFIG_OPTIONS}"
LABEL resty_config_options_more="${RESTY_CONFIG_OPTIONS_MORE}"
LABEL resty_config_deps="${_RESTY_CONFIG_DEPS}"
LABEL resty_add_package_builddeps="${RESTY_ADD_PACKAGE_BUILDDEPS}"
LABEL resty_add_package_rundeps="${RESTY_ADD_PACKAGE_RUNDEPS}"
LABEL resty_eval_pre_configure="${RESTY_EVAL_PRE_CONFIGURE}"
LABEL resty_eval_post_download_pre_configure="${RESTY_EVAL_POST_DOWNLOAD_PRE_CONFIGURE}"
LABEL resty_eval_post_make="${RESTY_EVAL_POST_MAKE}"
LABEL resty_luajit_options="${RESTY_LUAJIT_OPTIONS}"
LABEL resty_pcre_options="${RESTY_PCRE_OPTIONS}"


RUN DEBIAN_FRONTEND=noninteractive apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        build-essential \
        ca-certificates \
        curl \
        gettext-base \
        libgd-dev \
        libgeoip-dev \
        libncurses5-dev \
        libperl-dev \
        libreadline-dev \
        libxslt1-dev \
        make \
        perl \
        unzip \
        wget \
        zlib1g-dev \
        ${RESTY_ADD_PACKAGE_BUILDDEPS} \
        ${RESTY_ADD_PACKAGE_RUNDEPS} \
    && cd /tmp \
    && if [ -n "${RESTY_EVAL_PRE_CONFIGURE}" ]; then eval $(echo ${RESTY_EVAL_PRE_CONFIGURE}); fi \
    && curl -fSL "${RESTY_OPENSSL_URL_BASE}/openssl-${RESTY_OPENSSL_VERSION}.tar.gz" -o openssl-${RESTY_OPENSSL_VERSION}.tar.gz \
    && tar xzf openssl-${RESTY_OPENSSL_VERSION}.tar.gz \
    && cd openssl-${RESTY_OPENSSL_VERSION} \
    && if [ $(echo ${RESTY_OPENSSL_VERSION} | cut -c 1-5) = "1.1.1" ] ; then \
        echo 'patching OpenSSL 1.1.1 for OpenResty' \
        && curl -s https://raw.githubusercontent.com/openresty/openresty/master/patches/openssl-${RESTY_OPENSSL_PATCH_VERSION}-sess_set_get_cb_yield.patch | patch -p1 ; \
    fi \
    && if [ $(echo ${RESTY_OPENSSL_VERSION} | cut -c 1-5) = "1.1.0" ] ; then \
        echo 'patching OpenSSL 1.1.0 for OpenResty' \
        && curl -s https://raw.githubusercontent.com/openresty/openresty/ed328977028c3ec3033bc25873ee360056e247cd/patches/openssl-1.1.0j-parallel_build_fix.patch | patch -p1 \
        && curl -s https://raw.githubusercontent.com/openresty/openresty/master/patches/openssl-${RESTY_OPENSSL_PATCH_VERSION}-sess_set_get_cb_yield.patch | patch -p1 ; \
    fi \
    && ./config \
      no-threads shared zlib -g \
      enable-ssl3 enable-ssl3-method \
      --prefix=/usr/local/openresty/openssl \
      --libdir=lib \
      -Wl,-rpath,/usr/local/openresty/openssl/lib \
    && make -j${RESTY_J} \
    && make -j${RESTY_J} install_sw \
    && cd /tmp \
    && curl -fSL https://downloads.sourceforge.net/project/pcre/pcre/${RESTY_PCRE_VERSION}/pcre-${RESTY_PCRE_VERSION}.tar.gz -o pcre-${RESTY_PCRE_VERSION}.tar.gz \
    && echo "${RESTY_PCRE_SHA256}  pcre-${RESTY_PCRE_VERSION}.tar.gz" | shasum -a 256 --check \
    && tar xzf pcre-${RESTY_PCRE_VERSION}.tar.gz \
    && cd /tmp/pcre-${RESTY_PCRE_VERSION} \
    && ./configure \
        --prefix=/usr/local/openresty/pcre \
        --disable-cpp \
        --enable-utf \
        --enable-unicode-properties \
        ${RESTY_PCRE_BUILD_OPTIONS} \
    && make -j${RESTY_J} \
    && make -j${RESTY_J} install \
    && cd /tmp \
    && curl -fSL https://openresty.org/download/openresty-${RESTY_VERSION}.tar.gz -o openresty-${RESTY_VERSION}.tar.gz \
    && tar xzf openresty-${RESTY_VERSION}.tar.gz \
    && cd /tmp/openresty-${RESTY_VERSION} \
    && if [ -n "${RESTY_EVAL_POST_DOWNLOAD_PRE_CONFIGURE}" ]; then eval $(echo ${RESTY_EVAL_POST_DOWNLOAD_PRE_CONFIGURE}); fi \
    && eval ./configure -j${RESTY_J} ${_RESTY_CONFIG_DEPS} ${RESTY_CONFIG_OPTIONS} ${RESTY_CONFIG_OPTIONS_MORE} ${RESTY_LUAJIT_OPTIONS} ${RESTY_PCRE_OPTIONS} \
    && make -j${RESTY_J} \
    && make -j${RESTY_J} install \
    && cd /tmp \
    && rm -rf \
        openssl-${RESTY_OPENSSL_VERSION}.tar.gz openssl-${RESTY_OPENSSL_VERSION} \
        pcre-${RESTY_PCRE_VERSION}.tar.gz pcre-${RESTY_PCRE_VERSION} \
        openresty-${RESTY_VERSION}.tar.gz openresty-${RESTY_VERSION} \
    && curl -fSL https://luarocks.github.io/luarocks/releases/luarocks-${RESTY_LUAROCKS_VERSION}.tar.gz -o luarocks-${RESTY_LUAROCKS_VERSION}.tar.gz \
    && tar xzf luarocks-${RESTY_LUAROCKS_VERSION}.tar.gz \
    && cd luarocks-${RESTY_LUAROCKS_VERSION} \
    && ./configure \
        --prefix=/usr/local/openresty/luajit \
        --with-lua=/usr/local/openresty/luajit \
        --lua-suffix=jit-2.1.0-beta3 \
        --with-lua-include=/usr/local/openresty/luajit/include/luajit-2.1 \
    && make build \
    && make install \
    && cd /tmp \
    && if [ -n "${RESTY_EVAL_POST_MAKE}" ]; then eval $(echo ${RESTY_EVAL_POST_MAKE}); fi \
    && rm -rf luarocks-${RESTY_LUAROCKS_VERSION} luarocks-${RESTY_LUAROCKS_VERSION}.tar.gz \
    && if [ -n "${RESTY_ADD_PACKAGE_BUILDDEPS}" ]; then DEBIAN_FRONTEND=noninteractive apt-get remove -y --purge ${RESTY_ADD_PACKAGE_BUILDDEPS} ; fi \
    && DEBIAN_FRONTEND=noninteractive apt-get autoremove -y \
    && mkdir -p /var/run/openresty \
    && ln -sf /dev/stdout /usr/local/openresty/nginx/logs/access.log \
    && ln -sf /dev/stderr /usr/local/openresty/nginx/logs/error.log

# Add additional binaries into PATH for convenience
ENV PATH=$PATH:/usr/local/openresty/luajit/bin:/usr/local/openresty/nginx/sbin:/usr/local/openresty/bin

# Add LuaRocks paths
# If OpenResty changes, these may need updating:
#    /usr/local/openresty/bin/resty -e 'print(package.path)'
#    /usr/local/openresty/bin/resty -e 'print(package.cpath)'
ENV LUA_PATH="/usr/local/openresty/site/lualib/?.ljbc;/usr/local/openresty/site/lualib/?/init.ljbc;/usr/local/openresty/lualib/?.ljbc;/usr/local/openresty/lualib/?/init.ljbc;/usr/local/openresty/site/lualib/?.lua;/usr/local/openresty/site/lualib/?/init.lua;/usr/local/openresty/lualib/?.lua;/usr/local/openresty/lualib/?/init.lua;./?.lua;/usr/local/openresty/luajit/share/luajit-2.1.0-beta3/?.lua;/usr/local/share/lua/5.1/?.lua;/usr/local/share/lua/5.1/?/init.lua;/usr/local/openresty/luajit/share/lua/5.1/?.lua;/usr/local/openresty/luajit/share/lua/5.1/?/init.lua"

ENV LUA_CPATH="/usr/local/openresty/site/lualib/?.so;/usr/local/openresty/lualib/?.so;./?.so;/usr/local/lib/lua/5.1/?.so;/usr/local/openresty/luajit/lib/lua/5.1/?.so;/usr/local/lib/lua/5.1/loadall.so;/usr/local/openresty/luajit/lib/lua/5.1/?.so"

# Copy nginx configuration files
COPY nginx.conf /usr/local/openresty/nginx/conf/nginx.conf
COPY nginx.vh.default.conf /etc/nginx/conf.d/default.conf

CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]

# Use SIGQUIT instead of default SIGTERM to cleanly drain requests
# See https://github.com/openresty/docker-openresty/blob/master/README.md#tips--pitfalls
STOPSIGNAL SIGQUIT
```


#### 网上参考安装
[安装](https://www.cnblogs.com/yanjieli/p/15229907.html)

```
./configure --add-module=/path/to/ngx_http_proxy_connect_module
make && make install
```
```
wget http://dlsw.91donkey.com/software/source/nginx/ngx_http_proxy_connect_module.tgz && tar zxf ngx_http_proxy_connect_module.tgz

./configure \
--user=www \
--group=www \
--prefix=/usr/local/nginx \
--with-http_ssl_module \
--with-http_stub_status_module \
--with-http_realip_module \
--with-threads \
--add-module=/root/src/ngx_http_proxy_connect_module
```



```
server {

    listen 8080;

    # dns resolver used by forward proxying
    resolver 8.8.8.8;

    # forward proxy for CONNECT request
    proxy_connect;
    proxy_connect_allow 443 80;
    proxy_connect_connect_timeout 10s;
    proxy_connect_read_timeout 10s;
    proxy_connect_send_timeout 10s;

    # forward proxy for non-CONNECT request
    location / {
        auth_basic           "authentication";
        auth_basic_user_file htpasswd;
        proxy_pass http://$host;
        proxy_set_header Host $host;
    }
}
```

### 测试
设置代理
```
export http_proxy=http://user123:hahapwd@47.113.67.125:18888
export https_proxy=http://user123:hahapwd@47.113.67.125:18888
```
测试
```
curl -I --proxy https://xx:xx@proxy.wcoder.com  https://www.baidu.com

curl -I --proxy http://43.155.152.66:8080  https://www.baidu.com

curl -x https://proxy.wcoder.com --proxy-user x:x -L https://www.baidu.com


curl -x 43.155.152.66:8888 --proxy-user x:xx -I https://hub.docker.com/
```



## nginx正向代理配置

https://github.com/97994598069/97994598069.github.io/blob/10b7af1bff8b885de2e97a744f1e2bab87b55070/Nginx/nginx/nginx%E6%AD%A3%E5%90%91%E4%BB%A3%E7%90%86%E9%85%8D%E7%BD%AE.txt

```
说明：
1.正向代理的nginx安装正常安装就可以，没有特别的要求，
2.nginx当正向代理的时候，通过代理访问https的网站会失败，而失败的原因是客户端同nginx代理服务器之间建立连接失败，并非nginx不能将https的请求转发出去。因此要解决的问题就是客户端如何同nginx代理服务器之间建立起连接。有了这个思路之后，就可以很简单的解决问题。我们可以配置两个SERVER节点，一个处理HTTP转发，另一个处理HTTPS转发，而客户端都通过HTTP来访问代理，通过访问代理不同的端口，来区分HTTP和HTTPS请求。


下面看nginx的配置文件如下：
# cat nginx.conf
# For more information on configuration, see:
#   * Official English Documentation: http://nginx.org/en/docs/
#   * Official Russian Documentation: http://nginx.org/ru/docs/

user nginx;
worker_processes auto;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;

include /usr/share/nginx/modules/*.conf;

events {
    worker_connections 1024;
}

http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile            on;
    tcp_nopush          on;
    tcp_nodelay         on;
    keepalive_timeout   65;
    types_hash_max_size 2048;

    include             /etc/nginx/mime.types;
    default_type        application/octet-stream;

    include /etc/nginx/conf.d/*.conf;

#HTTP proxy       #这里位http的正向代理配置
    server{
        resolver 8.8.8.8;
        access_log /var/log/nginx/access_proxy-80.log main;
    listen 80;
    location / {
    root html;
    index index.html index.htm;
    proxy_pass $scheme://$host$request_uri;
    proxy_set_header HOST $http_host;
    proxy_buffers 256 4k;
    proxy_max_temp_file_size 0k;
    proxy_connect_timeout 30;
    proxy_send_timeout 60;
    proxy_read_timeout 60;
    proxy_next_upstream error timeout invalid_header http_502;
    }
    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
    root html;
        }
    }

#HTTPS proxy        #这里为：https的正向代理配置      
    server{
    resolver 8.8.8.8;
    access_log /var/log/nginx/access_proxy-443.log main;
    listen 443;
    location / {
    root html;
    index index.html index.htm;
    proxy_pass https://$host$request_uri;
    proxy_buffers 256 4k;
    proxy_max_temp_file_size 0k;
    proxy_connect_timeout 30;
    proxy_send_timeout 60;
    proxy_read_timeout 60;
    proxy_next_upstream error timeout invalid_header http_502;
    }
    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
    root html;
    }
    }
}


配置后重启nginx，
然后我们来访问测试下：
1、如果访问HTTP网站，可以直接这样的方式: curl --proxy proxy_server-ip:80 http://www.hm.net/

2、如果访问HTTPS网站，例如https://www.alipay.com，那么可以使用nginx的HTTPS转发的server：
curl --proxy proxy_server:443 http://www.alipay.com

3、使用浏览器访问
这里使用的是firefox浏览器
选项-->网络代理-->手动配置代理-->http代理


++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++




正向代理功能比较简单，但是原生nginx不支持https代理，如果访问https网站，会报错。
# nginx代理不支持http CONNECT方法：
curl: (56) Received HTTP code 400 from proxy after CONNECT

为了实现对https代理的支持，需要对原有nginx源码打补丁，就可以让nginx支持CONNECT模式了。
https://github.com/chobits/ngx_http_proxy_connect_module#build-openresty

二、下载并安装openresty
shell> yum -y install lua-devel
shell> wget http://dlsw.91donkey.com/software/source/nginx/openresty-1.15.8.1.tar.gz
shell> wget http://dlsw.91donkey.com/software/source/nginx/ngx_http_proxy_connect_module.tgz && tar zxf ngx_http_proxy_connect_module.tgz
shell> tar zxf openresty-1.15.8.1.tar.gz
shell> vim bundle/nginx-1.15.8/auto/cc/gcc
# 将下列代码注释掉，能够减少编译后nginx二级制文件的大小，提高程序执行效率。
# debug
CFLAGS="$CFLAGS -g"

shell> cd openresty-1.15.8.1
shell> ./configure --prefix=/opt/openresty --with-http_stub_status_module --with-http_sub_module \
    --with-http_auth_request_module --with-http_addition_module \
    --add-module=/usr/local/src/ngx_http_proxy_connect_module
shell> patch -d build/nginx-1.15.8/ -p 1 < /usr/local/src/ngx_http_proxy_connect_module/patch/proxy_connect_rewrite_101504.patch
shell> gmake -j 8 && gmake install
shell> echo "/opt/openresty/nginx/sbin/nginx" >> /etc/rc.d/rc.local
shell> /opt/openresty/nginx/sbin/nginx


三、配置服务器端nginx正向代理
# 在nginx.conf中增加server{}块，具体如下：
  server {
      listen 8080;
      resolver 8.8.8.8;
      resolver_timeout 5s;
      proxy_connect;
      proxy_connect_allow 443 563;
      proxy_connect_connect_timeout 10s;
      proxy_connect_read_timeout 10s;
      proxy_connect_send_timeout 10s;
      location / {
          proxy_pass $scheme://$host$request_uri;
          proxy_set_header Host $http_host;
          proxy_buffers 256 4k;
          proxy_max_temp_file_size 0;
          proxy_connect_timeout 30;
      }
      access_log /export/home/logs/proxy/access.log main;
      error_log /export/home/logs/proxy/error.log;
  }
  

四、配置终端代理
# 在 /etc/profile 文件中增加如下三项。
export proxy="http://{proxy_server_ip}:8080"
export http_proxy=$proxy
export https_proxy=$proxy

# 使配置生效
shell> source /etc/profile
```


### nginx ngx_http_proxy_connect_module
```
FROM alpine:3.9

LABEL maintainer="NGINX Docker Maintainers <docker-maint@nginx.com>"

ENV NGINX_VERSION 1.15.12

# https://github.com/chobits/ngx_http_proxy_connect_module下载的主分支包
ADD ngx_http_proxy_connect_module.tar.gz /opt/

RUN GPG_KEYS=B0F4253373F8F6F510D42178520A9993A1C052F8 \
	&& CONFIG="\
		--prefix=/etc/nginx \
		--sbin-path=/usr/sbin/nginx \
		--modules-path=/usr/lib/nginx/modules \
		--conf-path=/etc/nginx/nginx.conf \
		--error-log-path=/var/log/nginx/error.log \
		--http-log-path=/var/log/nginx/access.log \
		--pid-path=/var/run/nginx.pid \
		--lock-path=/var/run/nginx.lock \
		--http-client-body-temp-path=/var/cache/nginx/client_temp \
		--http-proxy-temp-path=/var/cache/nginx/proxy_temp \
		--http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp \
		--http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp \
		--http-scgi-temp-path=/var/cache/nginx/scgi_temp \
		--user=nginx \
		--group=nginx \
		--with-http_ssl_module \
		--with-http_realip_module \
		--with-http_addition_module \
		--with-http_sub_module \
		--with-http_dav_module \
		--with-http_flv_module \
		--with-http_mp4_module \
		--with-http_gunzip_module \
		--with-http_gzip_static_module \
		--with-http_random_index_module \
		--with-http_secure_link_module \
		--with-http_stub_status_module \
		--with-http_auth_request_module \
		--with-http_xslt_module=dynamic \
		--with-http_image_filter_module=dynamic \
		--with-http_geoip_module=dynamic \
		--with-threads \
		--with-stream \
		--with-stream_ssl_module \
		--with-stream_ssl_preread_module \
		--with-stream_realip_module \
		--with-stream_geoip_module=dynamic \
		--with-http_slice_module \
		--with-mail \
		--with-mail_ssl_module \
		--with-compat \
		--with-file-aio \
		--with-http_v2_module \
		# 对应上面添加的模块目录
                --add-module=/opt/ngx_http_proxy_connect_module \
	" \
	&& addgroup -S nginx \
	&& adduser -D -S -h /var/cache/nginx -s /sbin/nologin -G nginx nginx \
	&& apk add --no-cache --virtual .build-deps \
		gcc \
		libc-dev \
		make \
		openssl-dev \
		pcre-dev \
		zlib-dev \
		linux-headers \
		curl \
		gnupg1 \
		libxslt-dev \
		gd-dev \
		geoip-dev \
		# 编译ngx_http_proxy_connect_module依赖的
                patch \
                pcre \
                zlib \
	&& curl -fSL https://nginx.org/download/nginx-$NGINX_VERSION.tar.gz -o nginx.tar.gz \
	&& curl -fSL https://nginx.org/download/nginx-$NGINX_VERSION.tar.gz.asc  -o nginx.tar.gz.asc \
	&& export GNUPGHOME="$(mktemp -d)" \
	&& found=''; \
        for server in \
		ha.pool.sks-keyservers.net \
		hkp://keyserver.ubuntu.com:80 \
		hkp://p80.pool.sks-keyservers.net:80 \
		pgp.mit.edu \
	; do \
		echo "Fetching GPG key $GPG_KEYS from $server"; \
		gpg --keyserver "$server" --keyserver-options timeout=10 --recv-keys "$GPG_KEYS" && found=yes && break; \
	done; \
	test -z "$found" && echo >&2 "error: failed to fetch GPG key $GPG_KEYS" && exit 1; \
	gpg --batch --verify nginx.tar.gz.asc nginx.tar.gz \
	&& rm -rf "$GNUPGHOME" nginx.tar.gz.asc \
	&& mkdir -p /usr/src \
	&& tar -zxC /usr/src -f nginx.tar.gz \
	&& rm nginx.tar.gz \
	&& cd /usr/src/nginx-$NGINX_VERSION \
        # 对应版本的patch文件
        && patch -p1 < /opt/ngx_http_proxy_connect_module/patch/proxy_connect_rewrite_101504.patch \
	&& ./configure $CONFIG --with-debug \
	&& make -j$(getconf _NPROCESSORS_ONLN) \
	&& mv objs/nginx objs/nginx-debug \
	&& mv objs/ngx_http_xslt_filter_module.so objs/ngx_http_xslt_filter_module-debug.so \
	&& mv objs/ngx_http_image_filter_module.so objs/ngx_http_image_filter_module-debug.so \
	&& mv objs/ngx_http_geoip_module.so objs/ngx_http_geoip_module-debug.so \
	&& mv objs/ngx_stream_geoip_module.so objs/ngx_stream_geoip_module-debug.so \
	&& ./configure $CONFIG \
	&& make -j$(getconf _NPROCESSORS_ONLN) \
	&& make install \
	&& rm -rf /etc/nginx/html/ \
	&& mkdir /etc/nginx/conf.d/ \
	&& mkdir -p /usr/share/nginx/html/ \
	&& install -m644 html/index.html /usr/share/nginx/html/ \
	&& install -m644 html/50x.html /usr/share/nginx/html/ \
	&& install -m755 objs/nginx-debug /usr/sbin/nginx-debug \
	&& install -m755 objs/ngx_http_xslt_filter_module-debug.so /usr/lib/nginx/modules/ngx_http_xslt_filter_module-debug.so \
	&& install -m755 objs/ngx_http_image_filter_module-debug.so /usr/lib/nginx/modules/ngx_http_image_filter_module-debug.so \
	&& install -m755 objs/ngx_http_geoip_module-debug.so /usr/lib/nginx/modules/ngx_http_geoip_module-debug.so \
	&& install -m755 objs/ngx_stream_geoip_module-debug.so /usr/lib/nginx/modules/ngx_stream_geoip_module-debug.so \
	&& ln -s ../../usr/lib/nginx/modules /etc/nginx/modules \
	&& strip /usr/sbin/nginx* \
	&& strip /usr/lib/nginx/modules/*.so \
	&& rm -rf /usr/src/nginx-$NGINX_VERSION \
	\
	# Bring in gettext so we can get `envsubst`, then throw
	# the rest away. To do this, we need to install `gettext`
	# then move `envsubst` out of the way so `gettext` can
	# be deleted completely, then move `envsubst` back.
	&& apk add --no-cache --virtual .gettext gettext \
	&& mv /usr/bin/envsubst /tmp/ \
	\
	&& runDeps="$( \
		scanelf --needed --nobanner --format '%n#p' /usr/sbin/nginx /usr/lib/nginx/modules/*.so /tmp/envsubst \
			| tr ',' '\n' \
			| sort -u \
			| awk 'system("[ -e /usr/local/lib/" $1 " ]") == 0 { next } { print "so:" $1 }' \
	)" \
	&& apk add --no-cache --virtual .nginx-rundeps $runDeps \
	&& apk del .build-deps \
	&& apk del .gettext \
	&& mv /tmp/envsubst /usr/local/bin/ \
	\
	# Bring in tzdata so users could set the timezones through the environment
	# variables
	&& apk add --no-cache tzdata \
	\
	# forward request and error logs to docker log collector
	&& ln -sf /dev/stdout /var/log/nginx/access.log \
	&& ln -sf /dev/stderr /var/log/nginx/error.log

EXPOSE 80

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]
```