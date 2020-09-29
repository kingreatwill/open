# podman
Podman 是一个开源的容器运行时项目，可在大多数 Linux 平台上使用。Podman 提供与 Docker 非常相似的功能。正如前面提到的那样，它不需要在你的系统上运行任何守护进程，并且它也可以在没有 root 权限的情况下运行。
Podman 可以管理和运行任何符合 OCI（Open Container Initiative）规范的容器和容器镜像。Podman 提供了一个与 Docker 兼容的命令行前端来管理 Docker 镜像。

Podman 官网地址：https://podman.io/
Podman 项目地址：https://github.com/containers/libpod

## Podman 和docker不同之处？

- docker 需要在我们的系统上运行一个守护进程(docker daemon)，而podman 不需要
- 启动容器的方式不同：
    - docker cli 命令通过API跟 Docker Engine(引擎)交互告诉它我想创建一个container，然后docker Engine才会调用OCI container runtime(runc)来启动一个container。这代表container的process(进程)不会是Docker CLI的child process(子进程)，而是Docker Engine的child process。
    - Podman是直接给OCI containner runtime(runc)进行交互来创建container的，所以container process直接是podman的child process。

- 因为docke有docker daemon，所以docker启动的容器支持--restart策略，但是podman不支持，如果在k8s中就不存在这个问题，我们可以设置pod的重启策略，在系统中我们可以采用编写systemd服务来完成自启动
- docker需要使用root用户来创建容器，但是podman不需要

### 容器中的loginuid
- podman run alpine cat /proc/self/loginuid
    输出1001
- docker run alpine cat /proc/self/loginuid
  输出4294967295

Podman使用传统的fork/exec模型创建容器，因此容器进程是Podman进程的子进程。而Docker使用C/S(客户端/服务器)模型。执行docker命令的是Docker客户端工具，它通过C/S操作与Docker守护进程通信。Docker守护进程创建容器并处理stdin/stdout和Docker客户端工具的通信。

init默认loginuid是4294967295。由于容器是Docker守护程序的子进程，而Docker守护进程是init系统的子进程，所以 Docker守护进程和容器进程全部具有相同的loginuid为4294967295,在审计日志中都为unset，包括crond等

cat /proc/1/loginuid
4294967295

### 安全问题
我们来看看如果通过Docker启动的容器进程修改/etc/shadow文件会发生什么。

sudo docker run --privileged -v /:/host alpine touch /host/etc/shadow

sudo ausearch -f / etc / shadow –i

type=PROCTITLE msg=audit(07/31/2019 14:25:37.654:7990) : proctitle=sudo docker run --privileged -v /:/host alpine touch /host/etc/shadow

type=PATH msg=audit(07/31/2019 14:25:37.654:7990) : item=0 name=/etc/shadow inode=68873511 dev=fd:00 mode=file,000 ouid=root ogid=root rdev=00:00 objtype=NORMAL cap_fp=none cap_fi=none cap_fe=0 cap_fver=0

type=CWD msg=audit(07/31/2019 14:25:37.654:7990) : cwd=/home/test

type=SYSCALL msg=audit(07/31/2019 14:25:37.654:7990) : arch=x86_64 syscall=open success=yes exit=7 a0=0x7f442f2354f3 a1=O_RDONLY|O_CLOEXEC a2=0x1b6 a3=0x24 items=1 ppid=32003 pid=1773 auid=unset uid=test gid=test euid=root suid=root fsuid=root egid=test sgid=test fsgid=test tty=pts0 ses=998 comm=sudo exe=/usr/bin/sudo key=(null)

上面Docker审计日志中auid为unset（4294967295）。这样审计时候可虽然知道知道进程修改了/etc/shadow文件，但是谁就无法确定了。

如果该攻击者随后删除了Docker容器，那么修改/etc/shadow文件的系统上将没有任何跟踪。

现在让我们看看与Podman完全相同的场景。

sudo podman run --privileged -v /:/host alpine touch /host/etc/shadow

sudo ausearch -f /etc/shadow –i

type=PROCTITLE msg=audit(07/31/2019 14:30:01.061:8007) : proctitle=sudo podman run --privileged -v /:/host alpine touch /host/etc/shadow

type=PATH msg=audit(07/31/2019 14:30:01.061:8007) : item=0 name=/etc/shadow inode=68873511 dev=fd:00 mode=file,000 ouid=root ogid=root rdev=00:00 objtype=NORMAL cap_fp=none cap_fi=none cap_fe=0 cap_fver=0

type=CWD msg=audit(07/31/2019 14:30:01.061:8007) : cwd=/home/test

type=SYSCALL msg=audit(07/31/2019 14:30:01.061:8007) : arch=x86_64 syscall=open success=yes exit=7 a0=0x7fb84435e4f3 a1=O_RDONLY|O_CLOEXEC a2=0x1b6 a3=0x24 items=1 ppid=32003 pid=1916 auid=test uid=test gid=test euid=root suid=root fsuid=root egid=test sgid=test fsgid=test tty=pts0 ses=998 comm=sudo exe=/usr/bin/sudo key=(null)

由于它使用传统的fork/exec，因此Podman正确记录了所有内容。

通过审计追踪/etc/shadow文件的一个简单示例，但审计系统对于审计系统上的进程非常有用。使用fork /exec容器运行时启动容器（而不是客户端/服务器容器运行时）允许安全人员通过审计日志记录保持更好的安全性。

在启动容器时，fork /exec模型与C/S架构先比较还有许多其他不错的功能。例如，systemd有关功能：

SD_NOTIFY：如果将Podman命令放入systemd单元文件中，容器进程可以通过Podman向堆栈返回通知，表明服务已准备好接收任务。这是基于C/S架构的docker无法完成的事情。

套接字激活：可以将连接的套接字通过systemd传递到Podman并传递到容器进程中使用它们。这在C/S架构的docker中也是不可能的。

## CLI
https://podman.readthedocs.io/en/latest/Commands.html

Podman CLI 里面87%的指令都和Docker CLI 相同，官方给出了这么个例子alias docker=podman,所以说经常使用Docker CLI的人使用podman上手非常快

