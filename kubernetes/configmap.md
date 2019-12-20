configmap 热更新后自动重启 pods

https://github.com/stakater/Reloader

原理：部署一个服务监控configmap的变化，重启pod