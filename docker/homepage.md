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
> 已内置icons: https://github.com/walkxcode/dashboard-icons
https://materialdesignicons.com/ `mdi-XX`
https://simpleicons.org/ `si-XX`
颜色: `mdi-XX-#f0d453` or `si-XX-#a712a2`; 当然url也可以

> 配置: https://gethomepage.dev/v0.8.3/configs/settings/#card-background-blur
> 小部件: https://github.com/gethomepage/homepage/blob/main/docs/widgets/services/caddy.md
> 小部件: https://github.com/gethomepage/homepage/blob/main/docs/widgets/services/portainer.md

services.yaml
```yaml
---
# For configuration options and examples, please see:
# https://gethomepage.dev/latest/configs/services

- Docker:
    - homepage:
        href: https://homepage.wcoder.com/
        description: homepage
        server: my-docker
        container: homepage
    - DocuSeal:
        href: https://docuseal.wcoder.com/
        description: docuseal
        server: my-docker
        container: docuseal
    - algo:
        href: https://algo.wcoder.com/
        description: algo
        server: my-docker
        container: algo
    - shiori:
        href: https://record.ren/
        description: shiori
        server: my-docker
        container: shiori
    - frps:
        href: https://frp.wcoder.com/
        description: frp
        server: my-docker
        container: frps
    - uptime-kuma:
        href: https://kuma.wcoder.com/
        description: uptime-kuma
        server: my-docker
        container: uptime-kuma


- Portainer:
    - Portainer:
        icon: docker.svg
        href: https://portainer.wcoder.com/
        description: portainer
        server: my-docker
        container: portainer
        widget:
          type: portainer
          url: http://43.155.152.66:10006
          env: 2
          key: ptr_KOF45eXZ663qG5PSB6FPSAmWNXmDMquZpELZnmTyZVs=
    - Caddy:
        icon: nginx.svg
        href: https://wcoder.com/
        description: Caddy
        server: my-docker
        container: caddy
        widget:
          type: caddy
          url: http://10.0.0.11:2019


```

### 其它
https://github.com/Lissy93/dashy

https://github.com/linuxserver/Heimdall
https://github.com/bastienwirtz/homer
