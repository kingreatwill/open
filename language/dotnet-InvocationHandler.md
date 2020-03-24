# 动态接口代理
## JAVA代码
```java
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
 
interface IFoo {
	void say();
}
 
class FooHandler implements InvocationHandler {
 
	@Override
	public Object invoke(Object arg0, Method arg1, Object[] arg2)
			throws Throwable {
		if (("say").equals(arg1.getName())) {
			System.out.println("hello world!");
		}
		return null;
	}
 
}
 
public abstract class Program {
 
	public static void main(String[] args) throws InterruptedException {
 
		InvocationHandler handler = new FooHandler();
		Class<?> clazz = IFoo.class;
		IFoo foo = (IFoo) Proxy.newProxyInstance(clazz.getClassLoader(),
				new Class[] { clazz }, handler);
		foo.say();
	}
}
```

## 使用Emit 实现dotnet版本

InvocationHandler.cs
```csharp
namespace Proxy
{
    public interface InvocationHandler
    {
        object InvokeMember(object obj, int rid, string name, params object[] args);
    }
}
```

InterfaceProxy.cs
```csharp
namespace Proxy
{
    using System;
    using System.Collections.Generic;
    using System.Reflection;
    using System.Reflection.Emit;

    /// <summary>
    /// 动态接口代理
    /// </summary>
    public partial class InterfaceProxy
    {
        private class Map
        {
            public Type New
            {
                get;
                set;
            }

            public Type Org
            {
                get;
                set;
            }
        }

        private static IList<Map> maps = null;

        public static T New<T>(InvocationHandler hanlder) where T : class
        {
            object value = New(typeof(T), hanlder);
            if (value == null)
            {
                return null;
            }
            return (T)value;
        }

        public static object New(Type clazz, InvocationHandler hanlder)
        {
            if (clazz == null || !clazz.IsInterface)
            {
                throw new ArgumentException("clazz");
            }
            if (hanlder == null)
            {
                throw new ArgumentException("hanlder");
            }
            lock (maps)
            {
                Type type = GetType(clazz);
                if (type == null)
                {
                    type = CreateType(clazz);
                    maps.Add(new Map() { New = type, Org = clazz });
                }
                return Activator.CreateInstance(type, hanlder);
            }
        }
    }

    public partial class InterfaceProxy
    {
        private const MethodAttributes METHOD_ATTRIBUTES = MethodAttributes.Public | MethodAttributes.NewSlot |
            MethodAttributes.Virtual | MethodAttributes.Final | MethodAttributes.HideBySig;

        private const TypeAttributes TYPE_ATTRIBUTES = TypeAttributes.Public | TypeAttributes.Sealed |
            TypeAttributes.Serializable;

        private const FieldAttributes FIELD_ATTRIBUTES = FieldAttributes.Private;

        private const CallingConventions CALLING_CONVENTIONS = CallingConventions.HasThis;

        private const PropertyAttributes PROPERTY_ATTRIBUTES = PropertyAttributes.SpecialName;

        private static ModuleBuilder MODULE_BUILDER = null;

        static InterfaceProxy()
        {
            maps = new List<Map>();
            AssemblyName an = new AssemblyName("?");
            AssemblyBuilder ab = AppDomain.CurrentDomain.DefineDynamicAssembly(an, AssemblyBuilderAccess.RunAndSave);
            MODULE_BUILDER = ab.DefineDynamicModule(an.Name);
        }

        private static Type GetType(Type clazz)
        {
            for (int i = 0; i < maps.Count; i++)
            {
                Map map = maps[i];
                if (map.Org == clazz)
                {
                    return map.New;
                }
            }
            return null;
        }

        // 1、构造（Constructor）在这里定义构造时提供一个“InvocationHandler”类型的参数输入 这肯定是很重要的地方当然你也可以定义一个公共属性或字段在输入引用但显然不是太好
        private static void CreateConstructor(TypeBuilder tb, FieldBuilder fb)
        {
            Type[] args = new Type[] { typeof(InvocationHandler) };
            ConstructorBuilder ctor = tb.DefineConstructor(MethodAttributes.Public, CallingConventions.Standard, args);
            ILGenerator il = ctor.GetILGenerator();
            //
            il.Emit(OpCodes.Ldarg_0);
            il.Emit(OpCodes.Ldarg_1);
            il.Emit(OpCodes.Stfld, fb);
            il.Emit(OpCodes.Ret);
        }

        // 2、后台字段（Blank field）它主要保存对“InvocationHandler”的一个地址引用，在接口里面调用一个方法或属性时那么就会去使用这个引用 所以这个字段的重要性是很高 当然既然是后台字段肯定是不公开私有的字段
        private static FieldBuilder CreateField(TypeBuilder tb)
        {
            return tb.DefineField("handler", typeof(InvocationHandler), FIELD_ATTRIBUTES);
        }

        private static Type CreateType(Type clazz)
        {
            TypeBuilder tb = MODULE_BUILDER.DefineType(string.Format("{0}+{1}", typeof(InterfaceProxy).FullName, clazz.Name));
            tb.AddInterfaceImplementation(clazz);
            //
            FieldBuilder fb = CreateField(tb);
            //
            CreateConstructor(tb, fb);
            CreateMethods(clazz, tb, fb);
            CreateProperties(clazz, tb, fb);
            //
            return tb.CreateType();
        }

        private static void CreateMethods(Type clazz, TypeBuilder tb, FieldBuilder fb)
        {
            foreach (MethodInfo met in clazz.GetMethods())
            {
                CreateMethod(met, tb, fb);
            }
        }

        private static Type[] GetParameters(ParameterInfo[] pis)
        {
            Type[] buffer = new Type[pis.Length];
            for (int i = 0; i < pis.Length; i++)
            {
                buffer[i] = pis[i].ParameterType;
            }
            return buffer;
        }

        private static MethodBuilder CreateMethod(MethodInfo met, TypeBuilder tb, FieldBuilder fb)
        {
            ParameterInfo[] args = met.GetParameters();
            MethodBuilder mb = tb.DefineMethod(met.Name, InterfaceProxy.METHOD_ATTRIBUTES, met.ReturnType, GetParameters(args));
            ILGenerator il = mb.GetILGenerator();
            il.DeclareLocal(typeof(object[]));

            if (met.ReturnType != typeof(void))
            {
                il.DeclareLocal(met.ReturnType);
            }

            // 定义固态数组
            il.Emit(OpCodes.Nop);
            il.Emit(OpCodes.Ldc_I4, args.Length);
            il.Emit(OpCodes.Newarr, typeof(object));
            il.Emit(OpCodes.Stloc_0);
            // 把函数调用参数放入数组
            for (int i = 0; i < args.Length; i++)
            {
                il.Emit(OpCodes.Ldloc_0);
                il.Emit(OpCodes.Ldc_I4, i);
                il.Emit(OpCodes.Ldarg, (1 + i));
                il.Emit(OpCodes.Box, args[i].ParameterType);
                il.Emit(OpCodes.Stelem_Ref);
            }
            // 桥接到“InvocationHandler”中进行处理
            il.Emit(OpCodes.Ldarg_0);
            il.Emit(OpCodes.Ldfld, fb);
            il.Emit(OpCodes.Ldarg_0);
            il.Emit(OpCodes.Ldc_I4, met.MetadataToken);
            il.Emit(OpCodes.Ldstr, met.Name);
            il.Emit(OpCodes.Ldloc_0);
            il.Emit(OpCodes.Call, typeof(InvocationHandler).GetMethod("InvokeMember", BindingFlags.Instance | BindingFlags.Public));
            // 对“InvocationHandler”的返回值进行处理
            if (met.ReturnType == typeof(void))
            {
                il.Emit(OpCodes.Pop);
            }
            else
            {
                il.Emit(OpCodes.Unbox_Any, met.ReturnType);
                il.Emit(OpCodes.Stloc_1);
                il.Emit(OpCodes.Ldloc_1);
            }
            il.Emit(OpCodes.Ret);
            //
            return mb;
        }

        private static void CreateProperties(Type clazz, TypeBuilder tb, FieldBuilder fb)
        {
            foreach (PropertyInfo prop in clazz.GetProperties())
            {
                PropertyBuilder pb = tb.DefineProperty(prop.Name, PROPERTY_ATTRIBUTES, prop.PropertyType, Type.EmptyTypes);
                MethodInfo met = prop.GetGetMethod();
                if (met != null)
                {
                    MethodBuilder mb = CreateMethod(met, tb, fb);
                    pb.SetGetMethod(mb);
                }
                met = prop.GetSetMethod();
                if (met != null)
                {
                    MethodBuilder mb = CreateMethod(met, tb, fb);
                    pb.SetSetMethod(mb);
                }
            }
        }
    }
}

```

使用
```csharp
namespace Proxy
{
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Reflection;
    using System.Text;

    public static class Program
    {
        public interface IFoo
        {
            void Say();
        }

        public class FooInvocationHandler : InvocationHandler
        {
            object InvocationHandler.InvokeMember(object obj, int rid, string name, params object[] args)
            {
                MethodBase met = typeof(IFoo).Module.ResolveMethod(rid);
                if (met.Name == "Say")
                {
                    Console.WriteLine("hello world!");
                }
                return null;
            }
        }

        static void Main()
        {
            IFoo foo = InterfaceProxy.New<IFoo>(new FooInvocationHandler());
            foo.Say();
            Console.ReadKey(false);
        }
    }
}

```