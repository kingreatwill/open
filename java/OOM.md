[Spring Boot引起的“堆外内存泄漏”排查及经验总结](堆外内存泄漏排查.md)

[OutOfMemoryError常见原因及解决方法](https://github.com/StabilityMan/StabilityGuide/blob/master/docs/diagnosis/jvm/exception/%E7%B3%BB%E7%BB%9F%E7%A8%B3%E5%AE%9A%E6%80%A7%E2%80%94%E2%80%94OutOfMemoryError%E5%B8%B8%E8%A7%81%E5%8E%9F%E5%9B%A0%E5%8F%8A%E8%A7%A3%E5%86%B3%E6%96%B9%E6%B3%95.md)

[记一次线上内存报警排查过程](https://github.com/StabilityMan/StabilityGuide/blob/master/docs/diagnosis/system/memory/case/%E8%AE%B0%E4%B8%80%E6%AC%A1%E7%BA%BF%E4%B8%8A%E5%86%85%E5%AD%98%E6%8A%A5%E8%AD%A6%E6%8E%92%E6%9F%A5%E8%BF%87%E7%A8%8B.md)

[JVM常见线上问题 → CPU 100%、内存泄露 问题排查](https://www.toutiao.com/article/6881508437667086856/)

[Java内存区域与内存溢出异常](https://blog.csdn.net/javaymm/article/details/113449674)

### 实战OutOfMemoryError异常
#### Java堆溢出
Java堆用于存储对象实例，只要不断地创建对象，并且保证GC Roots到对象之间有可达路径来避免垃圾回收机制清除这些对象，那么在对象数量到达最大堆的容量限制后就会产生内存溢出异常。
```java
public class HeapOOM {
    /**
     * VM Args: -Xms20m -Xmx20m -XX:+HeapDumpOnOutOfMemoryError
     * -Xms,-Xmx限制堆的大小为20m，不可扩展。
     * -XX:+HeapDumpOnOutOfMemoryError
     * 用于让虚拟机在出现内存溢出异常时Dump出当前的内存堆转储快照以便事后分析
     */
    static class OOMObject{

    }
    public static void main(String[] args) {
	// write your code here
        List<OOMObject> list = new ArrayList<>();
        while (true){
            list.add(new OOMObject());
        }
    }
}

```
#### 虚拟机栈和本地方法栈溢出
虚拟机栈和本地方法栈，在Java虚拟机中描述了两种异常

- 如果线程请求的栈深度大于虚拟机所允许的最大深度，将抛出StackOverflowError。
- 如果虚拟机在扩展栈时无法申请到足够的内存空间，则抛出OutOfMemoryError。

由于实验范围限制于单线程中的操作，下面两种方式都无法让虚拟机产生OutOfMemoryError异常，尝试的结果都是获得StackOverflowError异常。

- 使用-Xss参数减少栈内存容量。结果：抛出StackOverflowError异常，异常出现时输出的堆栈深度相应减少。
- 定义大量本地方法，增大此方法帧中本地变量表的长度。结果抛出StackOverflowError异常时输出的堆栈深度响应缩小。

```java
/**
 * VM Args: -Xss:128k
 */
public class JavaVMStackSOF {
    private int stackLength = 1;

    public void stackLeak() {
        stackLength++;
        stackLeak();
    }

    public static void main(String[] args) throws Throwable{
        JavaVMStackSOF oom = new JavaVMStackSOF();
        try {
            oom.stackLeak();
        } catch (Throwable e) {
            System.out.println("stack length:" + oom.stackLength);
            throw e;
        }
    }
}

```

- 在单个线程下，无论是由于栈帧太大还是虚拟机栈容量太小，当内存无法分配的时候，虚拟机抛出的都是StackOverflowError。
- 如果测试时不仅限于单线程，通过不断建立线程的方式倒是可以产生内存溢出异常，但这样产生的内存溢出异常与栈空间是否足够大并不存在任何联系，在这种情况下，为每个线程的栈分配的内存越大，反而越容易产生内存溢出异常(如下代码)。
- 如果是建立过多线程导致的内存溢出，在不能减少线程或更换64位虚拟机的情况下，只能通过减少最大堆和栈容量来换取更多的线程。

```java
/**
 * VM Args: -Xss2M
 *
 * @author ym
 */
public class JavaVmStackOOM {
    private void dontStop() {
        while (true) {
        }
    }

    public void stackLeadByThread() {
        while (true) {
            Thread thread = new Thread(new Runnable() {
                @Override
                public void run() {
                    dontStop();
                }
            });
        }
    }

    public static void main(String[] args) {
        JavaVmStackOOM oom = new JavaVmStackOOM();
        oom.stackLeadByThread();
    }
}

```

#### 方法区和运行时常量池溢出
String.intern()是一个Native方法，它的作用是：如果字符串常量池中已经包含一个等于此String对象的字符串，则返回代表池中这个字符串的String对象；否则，将此String对象包含的字符串添加到常量池中，并返回此String对象的引用。在JDK 1.6及之前的版本中，由于常量池分配在永久代内，可通过-XX:PermSize和–XX:MaxPermSize限制方法区大小，从而间接限制其中常量池的容量。
```java
/**
 * VM Args: -XX:PermSize=10M -XX:MaxPermSize=10M
 */
public class RuntimeConstantPoolOOM {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        int i = 0;
        while (true) {
            list.add(String.valueOf(i++).intern());
        }
    }
}


```
- 运行时常量池溢出，在OutOfMemoryError后面跟随的提示信息是“PermGen space”,说明运行时常量池属于方法区的一部分。
- 而使用JDK 1.7运行不会得到相同的结果，while循环将一直进行下去。
- 方法区用于存放Class的相关信息。如类名，访问修饰符，常量池，字段描述，方法描述等。对于这些区域的测试，基本的思路是运行时产生大量的类去填满方法区，直到溢出。

#### 本机直接内存溢出
DirectMemory容量可通过-XX:MaxDirectMemorySize指定，如果不指定，则默认与Java最大值（-Xmx指定）一样。

```java
/**
 * VM Args:-Xmx20M -XX:MaxDirectMemorySize=10M
 */
public class DirectMemoryOOM {
    private static final int _1MB = 1024*1024;

    public static void main(String[] args) throws IllegalAccessException {
        Field unsafeField = Unsafe.class.getDeclaredFields()[0];
        unsafeField.setAccessible(true);
        Unsafe unsafe = (Unsafe)unsafeField.get(null);
        while (true) {
            unsafe.allocateMemory(_1MB);
        }
    }
}

```
由DirectMemory导致的内存溢出，一个明显特征是在Heap Dump文件中不会看见明显的异常。