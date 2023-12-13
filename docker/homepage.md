## homepage

### gethomepage
https://github.com/gethomepage/homepage

```
docker run -d --name homepage \
  -p 10005:3000 \
  -v /data/dockerv/homepage/config:/app/config \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --restart unless-stopped \
  ghcr.io/gethomepage/homepage:v0.8.3
```

```
/data/dockerv/homepage/icons:/app/public/icons
```

> -v /var/run/docker.sock:/var/run/docker.sock:ro  `:ro` [read-only volume](https://docs.docker.com/storage/volumes/#use-a-read-only-volume)
> 已内置icons:https://github.com/walkxcode/dashboard-icons
> https://gethomepage.dev/v0.8.3/configs/settings/#card-background-blur

### 其它
https://github.com/Lissy93/dashy

https://github.com/linuxserver/Heimdall
https://github.com/bastienwirtz/homer
