# 安装MAVEN
下载：http://mirrors.shu.edu.cn/apache/maven/maven-3/3.5.2/binaries/apache-maven-3.5.2-bin.tar.gz
```
## 解压
tar vxf apache-maven-3.5.2-bin.tar.gz
## 移动
mv apache-maven-3.5.2 /usr/local/maven3
```
修改环境变量， 在/etc/profile中添加以下几行
```
MAVEN_HOME=/usr/local/maven3
export MAVEN_HOME
export PATH=${PATH}:${MAVEN_HOME}/bin
```
记得执行source /etc/profile使环境变量生效。
输入mvn -version 返回版本信息则安装正常。

## docker maven
```
#maven:3.6.3-jdk-8 
#maven:3.6.3-jdk-8-alpine
FROM maven:3.6.3-jdk-8-alpine
COPY settings.xml /usr/share/maven/ref/

# -v $HOME/.m2:/root/.m2 # Maven都会在下载依赖项之前查看本地目录.
```

## 打包
```
# 功能：  打包
# 参数说明：
#		-Dmaven.test.skip=true：跳过单元测试
#		-U：每次构建检查依赖更新，可避免缓存中快照版本依赖不更新问题，但会牺牲部分性能
#		-e -X ：打印调试信息，定位疑难构建问题时建议使用此参数构建
#		-B：以batch模式运行，可避免日志打印时出现ArrayIndexOutOfBoundsException异常
#		-P：profile
# 使用场景： 打包项目且不需要执行单元测试时使用
mvn package -Dmaven.test.skip=true -U -e -X -B -P dev

#功能：打包;执行单元测试，但忽略单元测试用例失败，每次构建检查依赖更新
#使用场景： 需要执行单元测试，且使用构建提供的单元测试报告服务统计执行情况
# 使用条件：在”单元测试“中选择处理单元测试结果，并正确填写测试结果文件路径
#mvn package -Dmaven.test.failure.ignore=true -U -e -X -B

#功能：打包并发布依赖包到私有依赖库
#使用场景： 需要将当前项目构建结果发布到私有依赖仓库以供其他maven项目引用时使用
#mvn deploy -Dmaven.test.skip=true -U -e -X -B
```

## java dockerfile
```
FROM openjdk:17-slim-bullseye as base
WORKDIR /app
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
ENV JAVA_OPTS="-XX:InitialRAMPercentage=75 -XX:MaxRAMPercentage=75"

FROM maven:3.8-openjdk-17-slim as build
WORKDIR /src
# 这里可以修改settings.xml
# 方案一:修改镜像中的settings.xml ;
# 方案二: 把settings.xml复制到源代码中,copy到镜像中后 使用mvn命令指定settings.xml
COPY . ./
# -P dev
RUN mvn package --settings settings.xml -Dmaven.test.skip=true -U -e -X -B


FROM base as final
WORKDIR /app
COPY --from=build /src/target/springboot-0.0.1-SNAPSHOT.jar .
ENTRYPOINT ["java", "-jar", "springboot-0.0.1-SNAPSHOT.jar"]



#FROM openjdk:17-slim-bullseye
#WORKDIR /app
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#RUN echo 'Asia/Shanghai' >/etc/timezone
#ENV JAVA_OPTS="-XX:InitialRAMPercentage=75 -XX:MaxRAMPercentage=75"
#
#WORKDIR /app
#EXPOSE 80
#ADD ./target/springboot-0.0.1-SNAPSHOT.jar .

# docker build -t springbootdemo:v1 .
# docker run -it --rm -p 9808:80 springbootdemo:v1
# 访问http://localhost:9808/test/get
```
