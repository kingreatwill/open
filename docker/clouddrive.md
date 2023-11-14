# WebDAV 

WebDAV 基于 HTTP 协议的通信协议，在GET、POST、HEAD等几个HTTP标准方法以外添加了一些新的方法，使应用程序可对Web Server直接读写，并支持写文件锁定(Locking)及解锁(Unlock)，还可以支持文件的版本控制。

常用的文件共享有三种：FTP、Samba、WebDAV，它们各有优缺点，了解后才能更好地根据自己的需求选择方案。

- FTP属于古老的文件共享方式了，因为安全性，现代浏览器最新已默认不能打开FTP协议。SFTP在FTP基础上增加了加密，在Linux上安装OpenSSH后可以直接用SFTP协议传输。使用SFTP临时传送文件还可以，但做文件共享，性能不高，速度较慢。
- Samba是Linux下CIFS协议的实现，优势在于对于小白使用简章，和Windows系统文件共享访问一样，不需要安装第三方软件，而且移动端也有大量APP支持。苹果手机文件APP中添加网络存储用的就是这种方式。Windows下文件共享使用445端口，且不能更改。445端口常常受黑客关照，在广域网上大多运营封掉了访端口，所以这种文件共享只适合在内网使用。
- WebDAV 基于 HTTP 协议的通信协议，在GET、POST、HEAD等几个HTTP标准方法以外添加了一些新的方法，使应用程序可对Web Server直接读写，并支持写文件锁定(Locking)及解锁(Unlock)，还可以支持文件的版本控制。因为基于HTTP，在广域网上共享文件有天然的优势，移动端文件管理APP也大多支持WebDAV协议。使用HTTPS还能保安全性。Apache和Nginx支持WebDAV，可作为WebDAV文件共享服务器软件。也可以使用专门的WebDAV软件部署。


clouddrive和alist两者最重要的功能都是挂载云盘为本地磁盘，clouddrive出来时间早，挂载出来是smb协议，兼容性高但支持网盘有限，仅支持阿里云盘、天翼云盘、115等几个；
alist出来时间较晚，挂在出来是webdav协议，支持网盘数量丰富，且支持直接挂载阿里云盘共享内容，不需要转存占用自己网盘空间，缺点是webdav协议在电视端支持有限，kodi支持但操作反人类，其余当贝播放器和专业的播放机z9x支持较差（Z9X连接webdav速度较慢，无法流畅播放4k原盘）。因此，对于webdav支持不好的终端，需要将alist挂载的webdav协议内容通过clouddrive二次挂载为smb协议使用。

## clouddrive(非开源)
https://hub.docker.com/r/cloudnas/clouddrive
```
docker run -d --name clouddrive --restart unless-stopped -v /share/Public/dockerv/clouddrive/CloudNAS:/CloudNAS:shared  -v /share/Public/dockerv/clouddrive/Config:/Config -v /share/Public/dockerv/clouddrive/media:/media:shared --privileged --device /dev/fuse:/dev/fuse -p 9798:9798 cloudnas/clouddrive
```

## RaiDrive
RaiDrive为国外软件，对国内网盘支持并不好。如果有挂载阿里云盘、天翼网盘的需要，推荐使用替代品CloudDrive。
RaiDrive 基于开源项目开发，承诺不保存用户的账号密码。

## NetDrive
NetDrive 支持将网络文件夹映射到本地，体验相当不错。唯一遗憾的是这款软件价格不菲，并从 NetDrive 3 开始转为订阅制

## alist
一个支持多种存储，支持网页浏览和 WebDAV 的文件列表程序，由 gin 和 Solidjs 驱动。
https://github.com/alist-org/alist

https://alist.nn.ci/zh/guide/install/script.html


## WebDAV Server
https://github.com/hacdias/webdav

## Nginx 开启 WebDAV
在Nginx中实现WebDAV需要安装 libnginx-mod-http-dav-ext 模块，以下是Nginx的配置：
```
server {
        listen 80;
        listen [::]:80;

        server_name dav.engr-z.com;
        auth_basic "Authorized Users Only";
        auth_basic_user_file /etc/.htpasswd;

        location / {
                root /data/webdav;
                client_body_temp_path /var/temp;
                dav_methods PUT DELETE MKCOL COPY MOVE;
                dav_ext_methods PROPFIND OPTIONS;
                create_full_put_path on;
                client_max_body_size 10G;
        }
}

server {
        listen 443;
        listen [::]:443;
        server_name dav.engr-z.com;

        ssl on;
        ssl_certificate /data/www/cert/dav.engr-z.com_nginx/cert.pem;
        ssl_certificate_key /data/www/cert/dav.engr-z.com_nginx/cert.key;
        ssl_session_timeout 5m;
        ssl_protocols SSLv2 SSLv3 TLSv1;
        ssl_ciphers ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;
        ssl_prefer_server_ciphers on;

        location / {
                root /data/webdav;
                client_body_temp_path /var/temp;
                dav_methods PUT DELETE MKCOL COPY MOVE;
                dav_ext_methods PROPFIND OPTIONS;
                create_full_put_path on;
                client_max_body_size 10G;
        }

}
```
.htpasswd 用户密码文件的创建方式和 Apache 一样，htpasswd是apache的工具，如果使用nginx，可以单独安装该工具而不安装整个apache。在Ubuntu中使用 sudo apt install apache2-utils 安装。

Nginx 对 WebDAV 支持不是太好，建议使用 Apache 或专用于 WebDAV 服务软件架设。

## WebDAV挂载/映射
- Windows
打开 “计算机” ，点右键添加一个网络位置，按向导填入地址，用户名，密码。

挂载指定盘符：`net use Y: https://dav.engr-z.com/ /user:engrz /persistent:YES 密码`

[RaiDrive](https://www.raidrive.com) 是一款能够将一些网盘映射为本地网络磁盘的工具，支持 Google Drive、Google Photos、Dropbox、OneDrive、FTP、SFTP、WebDAV。

- Linux
使用命令挂载，需要安装 davfs2 ：
apt install davfs2
执行命令后系统会自动安装，出现提示，选是。

挂载：
sudo mount -t davfs http://dav.engr-z/ ./webdav/

## 可以用来挂载WebDav的软件
- Windows
Potplayer，kmplayer，RaiDrive，kodi，OneCommander，Mountain Duck，netdrive ❌，rclone
- Android
Nplayer，kmplayer，ES文件管理器，kodi，nova魔改，reex，cx 文件管理器，Solid Expore
- IOS
Nplayer，kmplayer，infuse，Fileball文件管理器
- 电视TV
Nplayer，kodi，nova魔改
- Mac
IINA，Mountain Duck，infuse，netdrive，rclone
- Linux
davfs2，rclone

## 参考
[网络存储文件共享之WebDAV](https://zhuanlan.zhihu.com/p/352216119)