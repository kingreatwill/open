[TOC]
# k8s之Pod健康检测

## Pod健康检测机制
对于Pod的健康状态检测，kubernetes提供了两类探针(Probe)来执行对Pod的健康状态检测：
- **LivenessProbe探针**：
用于判断容器是否存活，即Pod是否为running状态，如果LivenessProbe探针探测到容器不健康，则kubelet将kill掉容器，并根据容器的重启策略是否重启，如果一个容器不包含LivenessProbe探针，则Kubelet认为容器的LivenessProbe探针的返回值永远成功。
- **ReadinessProbe探针**：
用于判断容器是否启动完成，即容器的Ready是否为True，可以接收请求，如果ReadinessProbe探测失败，则容器的Ready将为False，控制器将此Pod的Endpoint从对应的service的Endpoint列表中移除，从此不再将任何请求调度此Pod上，直到下次探测成功。

每类探针都支持三种探测方法：

- **exec**：通过执行命令来检查服务是否正常，针对复杂检测或无HTTP接口的服务，命令返回值为0则表示容器健康。
- **httpGet**：通过发送http请求检查服务是否正常，返回200-399状态码则表明容器健康。
- **tcpSocket**：通过容器的IP和Port执行TCP检查，如果能够建立TCP连接，则表明容器健康。

探针探测的结果有以下三者之一：
- **Success**：Container通过了检查。
- **Failure**：Container未通过检查。
- **Unknown**：未能执行检查，因此不采取任何措施。

## LivenessProbe探针配置

> #### httpGet探测方式有如下可选的控制字段
> - host：要连接的主机名，默认为Pod IP，可以在http request head中设置host头部。
> - scheme: 用于连接host的协议，默认为HTTP。
> - path：http服务器上的访问URI。
> - httpHeaders：自定义HTTP请求headers，HTTP允许重复headers。
> - port： 容器上要访问端口号或名称。


## 配置探针(Probe)相关属性
探针(Probe)有许多可选字段，可以用来更加精确的控制Liveness和Readiness两种探针的行为(Probe)：

- initialDelaySeconds：Pod启动后延迟多久才进行检查，单位：秒。
- periodSeconds：检查的间隔时间，默认为10，单位：秒。
- timeoutSeconds：探测的超时时间，默认为1，单位：秒。
- successThreshold：探测失败后认为成功的最小连接成功次数，默认为1，在Liveness探针中必须为1，最小值为1。
- failureThreshold：探测失败的重试次数，重试一定次数后将认为失败，在readiness探针中，Pod会被标记为未就绪，默认为3，最小值为1。

[原文](https://blog.51cto.com/newfly/2137136)