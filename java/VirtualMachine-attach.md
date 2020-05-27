
在JDK中com.sun.tools.attach.VirtualMachine提供了一些从外部进程attach到jvm上，并执行一些操作的功能。

它提供了以下功能：
1. public void loadAgentLibrary(String agentLibrary, String options)
载入一个dll的agent
2. public void loadAgent(String agent, String options)
载入java的agent，官方名字叫做 JPLIS agent （ Java Programming Language Instrumentation Services），并且调用agentmain方法
3. public InputStream remoteDataDump(Object … args)
用于dumptreahd，jstack会用到
4. public InputStream dumpHeap(Object … args)
用于堆的dump，jmap会用到
5. public InputStream heapHisto(Object … args)
堆的柱状图的dump ，jmap会用到它
6. public InputStream setFlag(String name, String value)
设置一些标记的值，一些简单的可以，复杂的不能生效的，jinfo会用到它
7. public InputStream printFlag(String name)
打印标记的值，jinfo需要的。
## 创建一个简单的 Java Agent
```java
// https://github.com/xl1605368195/SimpleAgent
package com.demo;

import java.lang.instrument.Instrumentation;

public class Agent {
    public static void premain(String param, Instrumentation inst) {
        main(param, inst);
    }

    public static void agentmain(String param, Instrumentation inst) {
        main(param, inst);
    }

    private synchronized static void main(String args, final Instrumentation inst) {
        // 打印 attach时传入的参数
        System.out.println("args: " + args);
    }
}
```

## 主动发起 attach 代码

```java
import com.sun.tools.attach.*;
import java.io.IOException;
public class Main {
    public static void main(String[] args) {
        VirtualMachine vm = null;
        try {
            vm = VirtualMachine.attach("50447");
            // 指定Java Agent的jar包路径
            String agentPath = "../agent-1.0-SNAPSHOT-jar-with-dependencies.jar";
            vm.loadAgent(agentPath, "agent test");
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            try {
                vm.detach();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}

```