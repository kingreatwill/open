<!-- toc -->
[TOC]
# k8s部署策略
https://github.com/ContainerSolutions/k8s-deployment-strategies
https://www.jianshu.com/p/71e14c31cb82
https://blog.container-solutions.com/kubernetes-deployment-strategies
## 重建(recreate)：停止旧版本部署新版本
k8s重新部署 Recreate
先停止旧服务，然后启动新服务，这是最简单的一种部署方式
```
spec:
  strategy:
    type: Recreate
```
### 总结
- 应用状态全部更新
- 停机时间取决于应用程序的关闭和启动消耗的时间

## 蓝绿(blue/green)：新版本与旧版本一起存在，然后切换流量

### 总结
最好用来验证 API 版本问题
- 实时部署/回滚

- 避免版本问题，因为一次更改是整个应用的改变

- 需要两倍的资源

- 在发布到生产之前，应该对整个应用进行适当的测试

## 金丝雀(canary)：将新版本面向一部分用户发布，然后继续全量发布
### 总结
让部分用户参与测试
- 部分用户获取新版本
- 方便错误和性能监控
- 快速回滚
- 发布较慢
- 流量精准控制很浪费（99％A / 1％B = 99 Pod A，1 Pod B）


## A/B测(a/b testing)：以精确的方式（HTTP 头、cookie、权重等）向部分用户发布新版本
A/B测实际上是一种基于数据统计做出业务决策的技术。在 Kubernetes 中并不原生支持，需要额外的一些高级组件来完成改设置（比如Istio、Linkerd、Traefik、或者自定义 Nginx/Haproxy 等）。
### 总结
A/B测试(A/B testing) - 最适合部分用户的功能测试
- 几个版本并行运行
- 完全控制流量分配
- 特定的一个访问错误难以排查，需要分布式跟踪
- Kubernetes 没有直接的支持，需要其他额外的工具

## 滚动更新(rolling-update)：一个接一个地以滚动更新方式发布新版本


为了服务升级过程中提供可持续的不中断的服务，Kubernetes 提供了rolling update机制，具体配置需要修改对应服务的yaml文件

### 参数解析:
```yaml
minReadySeconds: 100 # 容器启动创建多少s后服务可用
strategy:
  # indicate which strategy we want for rolling update
  type: RollingUpdate
  rollingUpdate:
     maxSurge: 1 # 升级过程中最多可以比原先设置多出的POD数量

     maxUnavailable: 1 # 升级过程中最多有多少个POD处于无法提供服务的状态

replicas: 2             # 目的副本集个数
```

### 相关命令：
1. 滚动升级
       kubwx apply -f svc-zipkin.yaml --record

2. 暂停升级
      kubwx rollout pause deployment zipkin-server

3. 继续升级
      kubwx rollout resume deployment zipkin-server

4. 查看升级历史
      kubwx rollout history deployment zipkin-server

5. 回滚操作
     回滚到上一级
       kubwx rollout undo deployment zipkin-server

     回滚制定版本（根据rollout history的查看结果）
       kubwx rollout undo deployment zipkin-server --to-revision=13

### 总结
- 版本在实例之间缓慢替换
- rollout/rollback 可能需要一定时间
- 无法控制流量