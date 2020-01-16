[编译出dll](https://www.cnblogs.com/timeddd/p/11731160.html)
```
go build --buildmode=c-shared -o Test.dll
dotnet 
[DllImport(DLL_NAME, EntryPoint = "Test")]
```

### 让出CPU时间片
#### GO
用于让出CPU时间片
runtime.Gosched()

#### .net
Thread.Sleep()方法：是强制放弃CPU的时间片，然后重新和其他线程一起参与CPU的竞争。
用Sleep()方法是会让线程放弃CPU的使用权。
用SpinWait()方法是不会放弃CPU的使用权。

#### .net
await Task.Yield();// 让出时间片
Thread.SpinWait();// 不让出时间片

AsyncLocal 4.6+

4.5
using System.Runtime.Remoting;
using System.Runtime.Remoting.Messaging;
CallContext



#### java
Thread.yield(); //让出当前剩余的CPU时间片
Thread.onSpinWait();// 不让出时间片

TransmittableThreadLocal

## grpc Balance
```
conn, err := grpc.Dial(
    "",
    grpc.WithInsecure(),
    grpc.WithBalancer(grpc.RoundRobin(resolver.NewPseudoResolver([]string{
        "10.0.0.1:10000",
        "10.0.0.2:10000",
    }))),
)
```