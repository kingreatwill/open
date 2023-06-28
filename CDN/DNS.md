[TOC]
## CoreDNS
使用CoreDNS作为内网DNS服务器
CoreDNS是Golang编写的一个插件式DNS服务器，是Kubernetes 1.13 后所内置的默认DNS服务器
采用的开源协议为Apache License Version 2
CoreDNS也是CNCF孵化项目，目前已经从CNCF毕业。
CoreDNS 的目标是成为 Cloud Native（云原生）环境下的 DNS 服务器和服务发现解决方案。

CoreDNS 插件编写目前有两种方式：
1. 深度耦合 CoreDNS，使用 Go 编写插件，直接编译进 CoreDNS 二进制文件
2. 通过 GRPC 解耦(实际上grpc就是一个插件)，任意语言编写 GRPC 接口实现，CoreDNS 通过 GRPC 与插件交互


> 由于 GRPC 链接实际上借助于 CoreDNS 的 GRPC 插件，同时 GRPC 会有网络开销，TCP 链接不稳定可能造成 DNS 响应过慢等问题。



[基于etcd插件(自带的)的CoreDNS动态域名添加](https://www.cnblogs.com/boshen-hzb/p/7541901.html)
https://coredns.io/plugins/etcd/


### 在docker环境中安装coredns
`docker pull coredns/coredns:1.10.1`
创建配置文件
```
mkdir -p /etc/coredns

vi /etc/coredns/corefile

.:53 {
    hosts {
        10.0.0.1 my.host.com // 你域名和ip
        fallthrough
    }
    forward . 114.114.114.114:53 //你的备用 dns 
    log
}

也可以将 hosts 放在单独一个文件中

.:53 {
    hosts /etc/coredns/hostsfile {
        fallthrough
    }
    forward . 8.8.8.8:53
    log
}

# cat hostsfile
10.0.0.1 example1.org
```
启动服务
`docker run -it -d --name coredns --net=host -v /e/dockerv/coredns/:/etc/coredns/ coredns/coredns:1.10.1 -conf /etc/coredns/corefile`



```
mkdir -p /etc/coredns

cat >/etc/coredns/corefile<<EOF
.:53 {
    forward . 8.8.8.8:53
    log
}
EOF

docker run -d --name coredns \
  --restart=always \
  -v /e/dockerv/coredns/:/etc/coredns/ \
  -p 53:53/udp \
  coredns/coredns:1.10.1 -conf /etc/coredns/corefile
```

测试
`dig @127.0.0.1 -p 53 my.host.com`


全部一起执行的命令
```
rm -rf /etc/coredns && mkdir -p /etc/coredns && echo "
.:53 {
    hosts {
        # ip host
        127.0.0.1 host.com 
        fallthrough
    }
    # forward . dns-1 dns-2
    forward . 114.114.114.114
    log
}
">> /etc/coredns/corefile && docker run -it -d --name dns --net=host -v /etc/coredns:/etc/coredns/ coredns/coredns:latest -conf /etc/coredns/corefile
```

## tool

### DNS client
#### nslookup
#### dig
`dig www.wcoder.com`
#### dog 
https://kgithub.com/ogham/dog/
#### doggo
https://github.com/mr-karan/doggo



```sh
$ cd "$(mktemp -d)"
$ curl -sL "https://github.com/mr-karan/doggo/releases/download/v0.5.5/doggo_0.5.5_linux_amd64.tar.gz" | tar xz
$ mv doggo /usr/local/bin
```

`doggo www.wcoder.com`