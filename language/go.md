https://github.com/hwholiday/learning_tools


go runtime
深入浅出Golang Runtime
https://mp.weixin.qq.com/s?__biz=MjM5OTcxMzE0MQ==&mid=2653373399&idx=1&sn=041f3efab73730e7c5e3ef3e1d10f20d

练识课堂——Go内存详解.pdf



go内存模型
https://cloud.tencent.com/developer/article/1359184
https://www.cnblogs.com/shijingxiang/articles/11466957.html

pprof
https://www.cnblogs.com/qcrao-2018/p/11832732.html


编译出dll
https://www.cnblogs.com/timeddd/p/11731160.html

go build --buildmode=c-shared -o Test.dll
dotnet 
[DllImport(DLL_NAME, EntryPoint = "Test")]




GO
用于让出CPU时间片
runtime.Gosched()


.net
Thread.Sleep()方法：是强制放弃CPU的时间片，然后重新和其他线程一起参与CPU的竞争。
用Sleep()方法是会让线程放弃CPU的使用权。
用SpinWait()方法是不会放弃CPU的使用权。

.net
await Task.Yield();// 让出时间片
Thread.SpinWait();// 不让出时间片

AsyncLocal 4.6+

4.5
using System.Runtime.Remoting;
using System.Runtime.Remoting.Messaging;
CallContext



java
Thread.yield(); //让出当前剩余的CPU时间片
Thread.onSpinWait();// 不让出时间片

TransmittableThreadLocal