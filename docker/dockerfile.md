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