# dockerfile

## dockerfile 文档
https://docs.docker.com/engine/reference/builder/

## docker cli
https://docs.docker.com/engine/reference/run/

## dockerfile 术语表
https://docs.docker.com/glossary

## docker api
https://docs.docker.com/engine/api/latest/

## docker compose file
https://docs.docker.com/compose/compose-file/

## docker compose cli
https://docs.docker.com/compose/reference/overview/

## dockerd - daemon
https://docs.docker.com/engine/reference/commandline/dockerd/
> dockerd是docker后台的真正的进程。docker只是命令行工具。
> docker的后台的所有行为，都是dockerd来实现的。而docker命令，仅仅是交互工具，只是一个客户端。
> 当dockerd未启动时，docker是可以使用，但是执行build等一些命令时就会发现找不到daemon。
> 链接其它远程的docker
> docker -H tcp://192.168.0.83:2376 info
> export DOCKER_HOST="tcp://192.168.0.83:2376"
> docker info
> -H后面就是指定连接的服务端地址 info表示查看服务端daemon的信息

## Dockerfile 中 RUN, CMD, ENTRYPOINT 的区别

RUN 指令：用于指定 docker build 过程中要运行的命令。
CMD 在docker run 时运行，而非docker build;
CMD 指令的首要目的在于为启动的容器指定默认要运行的程序，程序运行结束，容器也就结束；
注意: CMD 指令指定的程序可被 docker run 命令行参数中指定要运行的程序所覆盖。
如果 dockerfile 中如果存在多个CMD指令，仅最后一个生效；

ENTRYPOINT 指令：类似于 CMD 指令，但其不会被 docker run 的命令行参数指定的指令所覆盖，而且这些命令行参数会被当作参数送给 ENTRYPOINT 指令指定的程序；
但是, 如果运行 docker run 时使用了 --entrypoint 选项，此选项的参数可当作要运行的程序覆盖 ENTRYPOINT 指令指定的程序；

两者联合使用技巧
```
ENTRYPOINT ["/usr/sbin/nginx"]
CMD ["-h"] # 为 ENTRYPOINT 指令指定的程序提供默认参数；只要docker run 里指定参数，这个`-h`就会被覆盖
```
当使用docker run --name test -it test_nginx 不传递任何参数时，此时启动容器会使用cmd 指令后的命令作为默认参数，打印nginx的帮助信息。此时cmd 后的内容并不是一个完整的指令，而是参数，如果其内容是一个完整的指令，那么它将覆盖掉ENTRYPOINT 中的内容。

如果使用docker run --name test -it test_nginx -g "daemon off" 启动时，此时给定的运行参数会覆盖掉CMD 指令对应的内容，此时nginx将作为前台进程运行，作为一个web服务器使用，通过browser可以看到hello world


## Here-Documents

https://docs.docker.com/engine/reference/builder/#here-documents

在多行内容中 首行 指定解释器 `#!/bin/bash`， 则所有内容整体被看作一个 Shell 脚本。
```
RUN <<EOT
#!/bin/bash
dist=$(uname -s)
wget -O example.com/app-${dist} app-${dist}
EOT
```
如果bash 直接跟在 EOT 后面。第一行的变量 dist 在第二行 wget 中无法使用。
```
RUN <<EOT bash
dist=$(uname -s)
wget -O example.com/app-${dist} app-${dist}
EOT
```

也可以使用其它解释器
```
FROM python:3.6
RUN <<EOT
#!/usr/bin/env python
print("hello world")
EOT
```

## HEALTHCHECK监控检查
```
FROM nginx:1.23
HEALTHCHECK --interval=5s --timeout=3s  --retries=3 \
    CMD curl -fs http://localhost/ || exit 1

```

docker ps 查看, 在STATUS列会看到健康检查的状态
```
CONTAINER ID   IMAGE                         COMMAND                  CREATED              STATUS                        PORTS                                       NAMES
49f844ec25bc   louislam/uptime-kuma:1.23.6   "/usr/bin/dumb-init …"   About a minute ago   Up About a minute (healthy)   0.0.0.0:3001->3001/tcp, :::3001->3001/tcp   uptime-kuma

```