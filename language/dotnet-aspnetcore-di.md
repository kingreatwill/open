# ASP.NET Core应用的7种依赖注入方式
## Startup构造器

  

构成HostBuilderContext上下文的两个核心对象（表示配置的IConfiguration对象和表示承载环境的IHostEnvironment对象）可以直接注入Startup构造函数中进行消费。由于ASP.NET Core应用中的承载环境通过IWebHostEnvironment接口表示，IWebHostEnvironment接口派生于IHostEnvironment接口，所以也可以通过注入IWebHostEnvironment对象的方式得到当前承载环境相关的信息。

  

我们可以通过一个简单的实例来验证针对Startup的构造函数注入。如下面的代码片段所示，我们在调用IWebHostBuilder接口的Startup<TStartup>方法时注册了自定义的Startup类型。在定义Startup类型时，我们在其构造函数中注入上述3个对象，提供的调试断言不仅证明了3个对象不为Null，还表明采用IHostEnvironment接口和IWebHostEnvironment接口得到的其实是同一个实例。

  
```csharp
    class Program
    {    
        static void Main()    
        {        
            Host.CreateDefaultBuilder().ConfigureWebHostDefaults(builder => builder.UseStartup<Startup>())        .Build().Run();
        }
    }
    
    public class Startup
    {    
        public Startup(IConfiguration configuration, IHostEnvironment  hostingEnvironment,         IWebHostEnvironment webHostEnvironment)    
        {        
            Debug.Assert(configuration != null);        
            Debug.Assert(hostingEnvironment != null);        
            Debug.Assert(webHostEnvironment != null);        
            Debug.Assert(ReferenceEquals(hostingEnvironment, webHostEnvironment));    
        }    
        public void Configure(IApplicationBuilder app) { }
    }
```

## Startup.Configure方法

  

依赖服务还可以直接注入用于注册中间件的Configure方法中。如果构造函数注入还可以对注入的服务有所选择，那么对于Configure方法来说，通过任意方式注册的服务都可以注入其中，包括通过调用IHostBuilder、IWebHostBuilder和Startup自身的ConfigureServices方法注册的服务，还包括框架自行注册的所有服务。

  

如下面的代码代码片段所示，我们分别调用IWebHostBuilder和Startup的ConfigureServices方法注册了针对IFoo接口和IBar接口的服务，这两个服务直接注入Startup的Configure方法中。另外，Configure方法要求提供一个用来注册中间件的IApplicationBuilder对象作为参数，但是对该参数出现的位置并未做任何限制。

  
```csharp
    class Program
    {    
        static void Main()    
        {        
            Host.CreateDefaultBuilder().ConfigureWebHostDefaults(builder => builder.UseStartup<Startup>()            .ConfigureServices(svcs => svcs.AddSingleton<IFoo, Foo>())).Build().Run();    
        }
    }
    
    public class Startup
    {    
        public void ConfigureServices(IServiceCollection services)  => services.AddSingleton<IBar, Bar>();
            
        public void Configure(IApplicationBuilder app, IFoo foo, IBar bar)   
        {        
            Debug.Assert(foo != null);        
            Debug.Assert(bar != null);    
        }
    }

```

## 中间件构造器

  

ASP.NET Core请求处理管道最重要的对象是用来真正处理请求的中间件。由于ASP.NET Core在创建中间件对象并利用它们构建整个请求处理管道时，所有的服务都已经注册完毕，所以任何一个注册的服务都可以注入中间件类型的构造函数中。如下所示的代码片段体现了针对中间件类型的构造函数注入。
```csharp
    class Program
    {    
        static void Main()    
        {        
            Host.CreateDefaultBuilder().ConfigureWebHostDefaults(builder => builder.ConfigureServices(svcs => svcs.AddSingleton<FoobarMiddleware>().AddSingleton<IFoo, Foo>().AddSingleton<IBar, Bar>()) .Configure(app => app.UseMiddleware<FoobarMiddleware>())).Build().Run();    
        }
    }
    
    public class FoobarMiddleware: IMiddleware
    {    
        public FoobarMiddleware(IFoo foo, IBar bar) {        
            Debug.Assert(foo != null);       
            Debug.Assert(bar != null);    
        }    
        
        public Task InvokeAsync(HttpContext context, RequestDelegate next)    
        {        
            Debug.Assert(next != null);        
            return Task.CompletedTask;    
        }
    }
```

## 中间件Invoke方法

  

如果采用基于约定的中间件类型定义方式，注册的服务还可以直接注入真正用于处理请求的InvokeAsync方法或者Invoke方法中。另外，将方法命名为InvokeAsync更符合TAP（Task-based Asynchronous Pattern）编程模式，之所以保留Invoke方法命名，主要是出于版本兼容的目的。如下所示的代码片段展示了针对InvokeAsync方法的服务注入。
```csharp
    class Program
    {    
        static void Main()    
        {        
            Host.CreateDefaultBuilder().ConfigureWebHostDefaults(builder => builder.ConfigureServices(svcs => svcs.AddSingleton<IFoo, Foo>().AddSingleton<IBar, Bar>()).Configure(app => app.UseMiddleware<FoobarMiddleware>())).Build().Run();    
        }
    }
    
    public class FoobarMiddleware
    {    
        private readonly RequestDelegate _next;    
        public FoobarMiddleware(RequestDelegate next) => _next = next;    
        public Task InvokeAsync(HttpContext context, IFoo foo, IBar bar)    
        {        
            Debug.Assert(context != null);        
            Debug.Assert(foo != null);        
            Debug.Assert(bar != null);        
            return _next(context);    
        }
    }
```
  

虽然约定定义的中间件类型和Startup类型采用了类似的服务注入方式，它们都支持构造函数注入和方法注入，但是它们之间有一些差别。中间件类型的构造函数、Startup类型的Configure方法和中间件类型的Invoke方法或者InvokeAsync方法都具有一个必需的参数，其类型分别为RequestDelegate、IApplicationBuilder和HttpContext，对于该参数在整个参数列表的位置，前两者都未做任何限制，只有后者要求表示当前请求上下文的参数HttpContext必须作为方法的第一个参数。

  

按照上述约定，如下这个中间件类型FoobarMiddleware的定义是不合法的，但是Starup类型的定义则是合法的。对于这一点，笔者认为可以将这个限制放开，这样不仅使中间件类型的定义更加灵活，还能保证注入方式的一致性。
```csharp
    public class FoobarMiddleware
    {    
        public FoobarMiddleware(RequestDelegate next);    
        public Task InvokeAsync(IFoo foo, IBar bar, HttpContext context);
    }
    public class Startup
    {    
        public void Configure(IFoo foo, IBar bar, IApplicationBuilder app);
    }

```
## Controller构造器

  

在一个ASP.NET Core MVC应用中，我们可以在定义的Controller中以构造函数注入的方式注入所需的服务。在如下所示的代码片段中，我们将IFoobar服务注入到HomeController的构造函数中。
```csharp
    class Program
    {    
        static void Main()    
        {        
            Host.CreateDefaultBuilder().ConfigureWebHostDefaults(builder => builder.ConfigureServices(svcs => svcs.AddSingleton<IFoobar, Foobar>().AddSingleton<IBar, Bar>().AddControllersWithViews())            .Configure(app => app.UseRouting().UseEndpoints(endpoints => endpoints.MapControllers()))).Build().Run();    
        }
    }
    public class HomeController : Controller
    {    
        public HomeController(IFoobar foobar) => Debug.Assert(foobar != null);
    }
```


## Action方法

  

借助于ASP.NET Core MVC基于模型绑定的参数绑定机制，我们可以将注册的服务绑定到目标Action方法的参数上，进而实现针对Action方法的依赖注入。在采用这种类型的注入方式时，我们需要在注入参数上按照如下的方式标注FromServicesAttribute特性，用以确定参数绑定的来源是注册的服务。
```csharp
    class Program
    {    
        static void Main()    
        {        
            Host.CreateDefaultBuilder().ConfigureWebHostDefaults(builder => builder.ConfigureServices(svcs => svcs.AddSingleton<IFoobar, Foobar>().AddControllersWithViews()).Configure(app => app.UseRouting().UseEndpoints(endpoints => endpoints.MapControllers()))).Build().Run();    
        }
    }
    public class HomeController: Controller
    {    
        [HttpGet("/")]    
        public void Index([FromServices]IFoobar foobar)    
        {        
            Debug.Assert(foobar != null);    
        }
   }
```


## 视图

  

在ASP.NET Core MVC应用中，我们还可以将服务注册到现的View中。假设我们定义了如下这个简单的MVC程序，并定义了一个简单的HomeController。
```csharp
    class Program
    {    
        static void Main()    
        {        
            Host.CreateDefaultBuilder().ConfigureWebHostDefaults(
                builder => builder.ConfigureServices(svcs => svcs.AddSingleton<IFoobar, Foobar>().AddControllersWithViews())
                .Configure(app => app.UseRouting().UseEndpoints(endpoints => endpoints.MapControllers()))
                ).Build().Run();    
        }
    }public class HomeController: Controller
    {        
        [HttpGet("/")]        
        public IActionResult Index() => View();
    }
```
  

我们为HomeController定义了一个路由指向根路径（“/”）的Action方法Index，该方法在调用View方法呈现默认的View之前，将注入的IFoo服务以ViewBag的形式传递到View中。如下所示的代码片段是这个Action方法对应View（/Views/Home/Index.cshtml）的定义，我们通过@inject指令注入了IFoobar服务，并将属性名设置为Foobar，这意味着当前View对象将添加一个Foobar属性来引用注入的服务。  
```csharp
    @inject IFoobar Foobar@{    Debug.Assert(Foobar!= null);}
```
  
