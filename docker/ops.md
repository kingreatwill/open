

## nightingale夜莺
https://github.com/ccfos/nightingale
## uptime-kuma
https://github.com/louislam/uptime-kuma

`docker run -d --restart=always -p 3001:3001 -v /data/dockerv/uptime-kuma/data:/app/data -v /var/run/docker.sock:/var/run/docker.sock --name uptime-kuma louislam/uptime-kuma:1.23.6`

### UptimeFlare
https://github.com/lyc8503/UptimeFlare

[文档](https://github.com/lyc8503/UptimeFlare/wiki/Quickstart)

- 开源且易于部署，用户可以在10分钟内完成设置，无需安装任何本地开发工具。
- 免费使用，利用Cloudflare Workers的免费层，为用户提供预算友好的监控解决方案。
- 支持高达50个每两分钟检查一次的任务。
- 提供全球化监控，确保全面了解服务性能。
- 高度定制化，包括自定义状态页面、请求配置和Webhooks，满足各类需求