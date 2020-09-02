学习视频
https://www.ixigua.com/home/4375653785021037/video/

JAVA高级架构师
https://profile.zjurl.cn/rogue/ugc/profile/?version_code=769&version_name=70609&user_id=800046079287044

Java Tutorial
https://www.javatpoint.com/spring-boot-tutorial


## 微服务框架

[Micronaut 2.0 M2 vs Quarkus 1.3.1 vs Spring Boot 2.3 M2 on JDK 14](https://github.com/graemerocher/framework-comparison-2020)

https://github.com/lizzyTheLizard/medium-java-framework-compare/tree/master/compare
### spring-boot
https://github.com/spring-projects/spring-framework 39k
https://github.com/spring-projects/spring-boot 50.2k

https://spring.io/projects/spring-cloud
### helidon
Oracle公司在2018年发起，并提供了对于[MicroProfile](https://microprofile.io/)规范的实现(MicroProfile项目始于2016年)，支持 GraalVM
https://github.com/oracle/helidon 2.1k

### micronaut
由Grails框架的作者在2018年创造，支持 GraalVM
https://micronaut.io/

https://github.com/micronaut-projects/micronaut-core 4.1k

> 易用，推荐

### Quarkus
在2019年由红帽(Red Hat)创造的Kubernetes原生Java框架，支持GraalVM
它基于[MicroProfile](https://microprofile.io/)，[Vert.x](https://vertx.io/)，[resteasy](https://github.com/resteasy/Resteasy),[Netty](https://netty.io/)和[Hibernate](https://hibernate.org/)等标准构建。

https://quarkus.io/

https://github.com/quarkusio/quarkus 5.7k

Quarkus 1.4.1（当前1.7.1）不推荐使用Java 8。 建议用户从现在开始使用Java 11。

> 很活跃，推荐


#### 微服务12要素
https://www.12facor.net/zh_cn/
https://blog.csdn.net/zeb_perfect/article/details/52536411
```
I. 基准代码
一份基准代码，多份部署
II. 依赖
显式声明依赖关系
III. 配置
在环境中存储配置
IV. 后端服务
把后端服务当作附加资源
V. 构建，发布，运行
严格分离构建和运行
VI. 进程
以一个或多个无状态进程运行应用
VII. 端口绑定
通过端口绑定提供服务
VIII. 并发
通过进程模型进行扩展
IX. 易处理
快速启动和优雅终止可最大化健壮性
X. 开发环境与线上环境等价
尽可能的保持开发，预发布，线上环境相同
XI. 日志
把日志当作事件流
XII. 管理进程
后台管理任务当作一次性进程运行
```