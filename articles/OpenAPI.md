https://github.com/OAI/OpenAPI-Specification
https://swagger.io/specification/
API（Application Programming Interface，应用编程接口）

## OpenAPI规范
OpenAPI规范始于Swagger规范，经过Reverb Technologies和SmartBear等公司多年的发展，OpenAPI计划拥有该规范（捐赠之后），OpenAPI Initiative在GitHub上托管社区驱动的规范。

规范是一种与语言无关的格式，用于描述RESTful Web服务，应用程序可以解释生成的文件，这样才能生成代码、生成文档并根据其描述的服务创建模拟应用。


Apifox = Postman + Swagger + Mock
支持导入 OpenApi格式（原Swagger）、Postman、HAR、RAP2、yapi、Eolinker、DOClever、ApiPost、Apizza 等数据格式。

## swagger

### swagger-codegen-cli
https://mvnrepository.com/artifact/io.swagger/swagger-codegen-cli
https://mvnrepository.com/artifact/io.swagger.codegen.v3/swagger-codegen-cli

java -jar swagger-codegen-cli-3.0.21.jar config-help -l python

-D{optionName}={optionValue}

https://gitee.com/api/v5/doc_json

java -jar swagger-codegen-cli-3.0.21.jar  generate -i doc_json.json -l python -o client -DpackageName=gitee -DprojectName=pygitee -DpackageName=gitee -DprojectName=pygitee

## tool
### postman
### Apifox
### Insomnia
The open-source, cross-platform API client for GraphQL, REST, WebSockets and gRPC.
https://insomnia.rest/
https://github.com/Kong/insomnia