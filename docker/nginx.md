docker run -p 80:80 -v /dockerv/nginx/www:/www -v /dockerv/nginx/conf/nginx.conf:/etc/nginx/nginx.conf -v /dockerv/nginx/logs:/wwwlogs  -d  nginx
--restart always