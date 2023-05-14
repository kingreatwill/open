[为什么SpringBoot可以直接运行 Jar 包？](https://www.toutiao.com/article/7064428929943257608/)

[7min到40s：SpringBoot启动优化实践](https://www.toutiao.com/article/7181632167833846306/)

[SpringBoot优雅停机](https://www.toutiao.com/article/7194248840076083712/)

[Spring容器获取Bean的9种方式，你能get几种？](https://www.toutiao.com/article/7187938000448635424/)

[事务 @Transactional](https://mp.weixin.qq.com/s/AZr5i6FR9xwj_1nFQvX20g)
[SpringBoot 多数据源及事务解决方案](https://mp.weixin.qq.com/s/e7ICGVN9jcGupNtYV-Labg)

[SpringBoot教程(7) @ConditionalOnProperty 详细讲解和示例](https://blog.csdn.net/winterking3/article/details/114822929)
[通过自定义starter，优雅实现接口统一返回和统一异常处理](https://www.toutiao.com/article/7223019919686943289/)

[使用ImportSelector接口动态加载Bean](https://www.toutiao.com/article/7222132283204010508/)

## profile
application.properties

`spring.profiles.active=@profiles.active@`

项目级别pom.xml
```xml
<profiles>
		<profile>
			<id>dev</id>
			<properties>
				<profiles.active>dev</profiles.active>
			</properties>
			<activation>
				<activeByDefault>true</activeByDefault>
			</activation>
		</profile>
		<profile>
			<id>uat</id>
			<properties>
				<profiles.active>uat</profiles.active>
			</properties>
		</profile>
	</profiles>
```
全局设置源码下载
settings.xml
```xml
<profiles>
<profile>
        <id>downloadSources</id>
        <properties>
            <downloadSources>true</downloadSources>
            <downloadJavadocs>true</downloadJavadocs>
        </properties>
    </profile>
</profiles>
<activeProfiles>
      <activeProfile>downloadSources</activeProfile>
  </activeProfiles>
```
