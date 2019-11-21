https://github.com/BuiltCloud/grpc-web-demo

![](../img/grpc-web-demo.png)

1. 创建grpc-service
https://github.com/KnowitSolutions/Grpc.Helpers
2. 创建grpc-web
https://github.com/grpc/grpc-web/tree/master/net/grpc/gateway/examples/helloworld
grpc-web
```
protoc -I=. helloworld.proto \
  --js_out=import_style=commonjs:. \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
```
```
  npm install
```
```
npx webpack client.js
```
```
dotnet serve
```