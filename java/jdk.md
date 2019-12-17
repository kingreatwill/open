[TOC]
https://www.oracle.com/technetwork/java/javase/jdk-relnotes-index-2162236.html

# 名词介绍

## TM
TM是英文Trademark商标的意思。
在我国，商标符号是：® 或 注 ，没有使用TM的规定，采用“先注册原则”，所以如果谁注册了某个东西为商标就可以加个商标符号。
在美国，商标采用“先使用原则”，如果产生冲突，法律（法庭）会保护优先使用者权利即使没有注册。美国一般习惯使用TM来表明：“这是一个我们已经使用的商标”

## JCP
Java Community Process ，Java社区进程
JCP成立于1998年，[官网](https://www.jcp.org/en/home/index)，由社会各界Java组成的社区，规划和领导Java的发展，其[成员](https://jcp.org/en/participation/members)可以在这里看到

## JSR

Java Specification Requests，Java规范请求，由JCP成员向委员会提交的Java发展议案，经过一系列流程后，如果通过最终会体现在未来的Java中


## TCK
Technology Compatibility Kit，技术兼容性测试
如果一个平台型程序想要宣称自己兼容Java，就必须通过TCK测试

## JEP
JDK Enhancement Proposal，a process used by the OpenJDK community for collecting proposals for enhancements to the Java Development Kit.
JDK增强提案，OpenJDK社区使用的一种过程，用于收集有关Java开发工具包的增强提案。
[官网](http://openjdk.java.net/jeps)


## OpenJDK
Sun公司初始设立的开发Java源码组织，是组织也是开源JDK的名字


# JDK

版本号|名称|中文名|发布日期
--|--|--|--
JDK 1.1.4|Sparkler|宝石|1997-09-12
JDK 1.1.5|Pumpkin|南瓜|1997-12-13
JDK 1.1.6|Abigail|阿比盖尔--女子名|1998-04-24
JDK 1.1.7|Brutus|布鲁图--古罗马政治家和将军|1998-09-28
JDK 1.1.8|Chelsea|切尔西--城市名|1999-04-08
J2SE 1.2|Playground|运动场|1998-12-04
J2SE 1.2.1|none|无|1999-03-30
J2SE 1.2.2|Cricket|蟋蟀|1999-07-08
J2SE 1.3|Kestrel|美洲红隼|2000-05-08
J2SE 1.3.1|Ladybird|瓢虫|2001-05-17
J2SE 1.4.0|Merlin|灰背隼|2002-02-13
J2SE 1.4.1|grasshopper|蚱蜢|2002-09-16
J2SE 1.4.2|Mantis|螳螂|2003-06-26
[Java SE 5.0/1.5](https://www.oracle.com/java/technologies/java-archive-javase5-downloads.html)|Tiger|老虎|2004-09-30
[Java SE 6.0/1.6](https://www.oracle.com/java/technologies/javase-java-archive-javase6-downloads.html)|Mustang|野马|2006-04
[Java SE 7.0/1.7](https://www.oracle.com/java/technologies/javase-java-archive-javase7-downloads.html)|Dolphin|海豚|2011-07-28
[Java SE 8.0/1.8 (LTS)](https://www.oracle.com/java/technologies/javase-java-archive-javase8-downloads.html)|Spider|蜘蛛|2014-03-18
Java SE 9.0 | | |2017-09-21
Java SE 10.0	| | |	2018-03-21
Java SE 11.0(LTS)| | |	2018-09-25
Java SE 12.0	| | |	2019/03/19
Java SE 13.0	| | |	2019/09/17
Java SE 14.0	| | |	2020/03/17	
Java SE 15.0	| | |	
Java SE 16.0	| | |	
Java SE 17.0(LTS)	| | |	


时间 | 事迹
--|--
1995年5月23日| Java语言诞生
1996年1月| 第一个JDK-JDK1.0诞生
1996年4月| 10个最主要的操作系统供应商申明将在其产品中嵌入JAVA技术
1996年9月| 约8.3万个网页应用了JAVA技术来制作
1997年2月18日| JDK1.1发布
1997年4月2日| JavaOne会议召开，参与者逾一万人，创当时全球同类会议规模之纪录
1997年9月| JavaDeveloperConnection社区成员超过十万
1998年2月| JDK1.1被下载超过2,000,000次
1998年12月8日| JAVA2企业平台J2EE发布
1999年6月| SUN公司发布Java的三个版本：标准版、企业版和微型版（J2SE、J2EE、J2ME）
2000年5月8日| JDK1.3发布
2000年5月29日| JDK1.4发布
2001年6月5日| NOKIA宣布，到2003年将出售1亿部支持Java的手机
2001年9月24日| J2EE1.3发布
2002年2月26日| J2SE1.4发布，自此Java的计算能力有了大幅提升。
2004年9月30日18:00PM| J2SE1.5发布，是Java语言的发展史上的又一里程碑事件。为了表示这个版本的重要性，J2SE1.5更名为J2SE5.0
2005年6月| JavaOne大会召开，SUN公司公开Java SE 6。此时，Java的各种版本已经更名以取消其中的数字“2”：J2EE更名为Java EE, J2SE更名为Java SE，J2ME更名为Java ME。
2006年11月13日| SUN公司宣布Java全线采纳GNU General Public License Version 2，从而公开了Java的源代码。
2009年4月20日| 甲骨文以现金收购Sun微系统公司，交易价格达74亿美元。
 Oracle| 从2006年12月份Sun发布Java 6后，经过五年多的不懈努力在2011年7月底发布了Java 7正式版！这也是Sun被Oracle收购以来发行的第一个版本。而在三年后的今天，被冠名为“跳票王”的Oracle终于发布了Java 8的正式版，但对于很多开发者来说却比Java 7来的更漫长一些。主要原因还是因为Oracle原本计划在2013年发布正式版Java 8，却因受困于安全性的问题经过了两次跳票，经历9个里程碑版本。当然，我们更不愿意看到Oracle因如期发布而牺牲质量，把原先没有解决的一些缺陷的安全问题带到Java 8当中去。同时也很可能将放弃掉Lambda而导致广泛应用的可能性更小。不管怎么样，Java 8如今来了，全新“革命”而不再是“进化”的功能将会让无数开发者动容。
2014年3月18日| Java8.0发布，这是继Java5.0以来变化最大的版本。一共有10大新特性。最主要的是Lambda表达式和强大的StreamAPI和新版的日期时间API，函数式接口和接口的默认方法和静态方法等。
2017年9月21日| Java9的发布。Java 9的最主要目标是最大限度实现模块化以帮助人们实现积木式的应用编写。目的是帮助人们从JAR的束缚中解脱出来。该特性将贯穿整个Java库，并以单依赖图的方式重新整理依赖。Java 9会把所有三个Java开发平台统一起来，模块化特性会使得Java ME的可复用性得到增强，这将是反击Android和iOS的有力武器。

## JDK1.4
### 正则表达式
### 异常链
### NIO
### 日志类
### XML解析器
### XLST转换器

## JDK1.5

### 自动拆装箱
Java数据类型分两种：基本数据类型 和  引用数据类型(对象)
有时候我们需要将基本数据类型包装为对象进行处理
在JKD5以前我们的处理方式：
```java
//int 转换为 Integer
int i = 10;
Integer integer = new Integer(i);

//Integer 转换为 int
Integer integer1 = new Integer(100);
int i1 = integer1.intValue();
```
自动拆装箱处理方式：
```java
//int 转换为 Integer
Integer integer = 10;
//Integer 转换为 int
int i = integer;
```
将class反编译可以看出自动拆装箱的代码如下：
```java
Integer integer = Integer.valueOf(10);
int i = integer.intValue();
```
以上就是自动拆装箱的效果，同理其余基本类型也可以自动裁装箱对应的对象。详细对应关系如下表：
基本数据类型 |	封装类
--|--
int	|Integer
byte|	Byte
short|	Short
long|	Long
char|	Character
double|	Double
float	|Float
boolean	|Boolean

---

自动装箱的过程：每当需要一种类型的对象时，这种基本类型就自动地封装到与它相同类型的包装中。
自动拆箱的过程：每当需要一个值时，被装箱对象中的值就被自动地提取出来，没必要再去调用intValue()和doubleValue()方法。
自动装箱，只需将该值赋给一个类型包装器引用，java会自动创建一个对象。
自动拆箱，只需将该对象值赋给一个基本类型即可。
java——类的包装器
类型包装器有：Double,Float,Long,Integer,Short,Character和Boolean

### For-Each循环 
增强for循环，新增一种循环语法，格式：for( : )
普通for与增强for循环对比如下：
```java
List<String> list = new ArrayList<String>();
list.add("111");
list.add("222");
list.add("333");

//JDK5 以前循环需要定义下标“index”并初始化，判断是否小于集合长度并自增，循环体还需要赋值
for (int index = 0; index < list.size(); index++ ) {
    String str = list.get(index);
    System.out.println("str: " + str);
}

//foreach 增强for循环只需要下面代码即可完成上面的操作。
for (String str : list){
    System.out.println("str:" + str);
}
```
反编译class文件可以看到增强for循环会被编译器自动处理如下代码：
```java
Iterator var4 = list.iterator();

while(var4.hasNext()) {
    str = (String)var4.next();
    System.out.println("str:" + str);
}
```
具体编译成什么类型还的根据循环数据实际的数据类型，例如：

```java
/int数组 foreach 
int[] ints = new int[]{1, 2, 3, 4, 5};
for(int i : ints){
    System.out.println("i: " + i);
}

//class反编译结果
int[] ints = new int[]{1, 2, 3, 4, 5};
int[] var9 = ints;
int var4 = ints.length;

for(int var5 = 0; var5 < var4; ++var5) {
    int i = var9[var5];
    System.out.println("i: " + i);
}
```
从上面代码可以大致了解foreach，它基本可以替换掉你所有用普通for循环的代码。

### 静态导入
静态导入可以将静态方法和静态变量通过 import static 和导包一样的方式直接导入，使用的时候无需使用类名即可直接访问。
```java
import static java.lang.System.out;
import static java.lang.Math.*;

public class ImportStaticTest {

    public static void main(String[] args) {
        /*
         *  使用静态导入 import static java.lang.System.out;
         *  out.println 可以直接使用调用 而不再需要类System对象去调用
         *  同时也支持*通配符
         */
        out.println("max(3,5): " + max(3,5));
        out.println("random(): " + random());
    }
}
```

### 可变参数（Varargs）

当传入到方法的参数不固定时，就可以使用可变参数 格式：func(数据类型... 参数名)
```java
public class VarArgsTest {

    // Tips: 和以往main方式不一样，一般这样写 public static void main（String[] args）
    public static void main(String... args) {
        varArgs(1);
        varArgs(2,3);

        // ...

        varArgs(7,8,9,10,11);
    }

    // 可变参数的格式： 数据类型... 参数名
    public static void varArgs(int... ints) {
        for (int i : ints){
            System.out.println(i);
        }
    }
}
```
可变参数必须放在方法参数的最后一个并且有且只有一个 

### 枚举
把集合里的对象元素一个一个提取出来。枚举类型使代码更具可读性，理解清晰，易于维护。枚举类型是强类型的，从而保证了系统安全性。而以类的静态字段实现的类似替代模型，不具有枚举的简单性和类型安全性。

简单的用法：JavaEnum简单的用法一般用于代表一组常用常量，可用来代表一类相同类型的常量值。
复杂用法：Java为枚举类型提供了一些内置的方法，同事枚举常量还可以有自己的方法。可以很方便的遍历枚举对象。

注意事项：
- 不能含有public修饰的构造器，即构造器只能是private修饰，如果没有构造器编译器同样也会自动生成一个带private修饰无参默认构造器;
- 所有的枚举值默认都是public static final 修饰的;
- 枚举值与值之间用逗号分割，最后一个用分号，如果枚举值后面没有任何东西该分号可以省略;
- 每一个枚举值就是一个枚举类型的实例;
- 枚举类型中可以定义带任意修饰符的非枚举值变量;
- 枚举值必须位于枚举类型的第一位置，即非枚举值必须位于枚举值之后;
- 枚举类型自带两个方法，values() 和 value(String name) 两个方法；
- 在枚举中支持定义其他结构，如构造方法、普通方法和普通属性等；
- 当枚举类中还包含其他结构时，枚举对象的定义必须放在首行（否则会编译报错）;
- 枚举不能够再继承类，因为它默认继承了Enum类，但枚举可以实现一个或多个接口。
```java
//定义枚举类型
public enum SexEnum {
    MAN,WOMAN
}

// class反编译结果：
public enum SexEnum {
    MAN,
    WOMAN;

    private SexEnum() {
    }
}

```
Enum类常用的三个方法:
- ordinal（）：返回枚举对象下标，默认第一个对象为0；
- name（）：返回枚举对象名称，默认与对象名称保持一致；
- values（）：返回所有枚举对象的数组；

### 格式化输出
JDK5推出了printf-style格式字符串的解释器 java.util.Formatter 工具类，和C语言的printf()有些类似。
```java
//创建对象时指定将结果输出到控制台
Formatter formatter = new Formatter(System.out);
formatter.format("x = %d , y = %s\n",1 , "test");
formatter.close();
```
Formatter类可以将一些特定风格的字符转换为另一种格式进行输出，给出一下常用格式转换。
Format|类型
--|--
d   |	整数型
s	|	String
f   |	浮点数
c	|	Unicode字符
b	|	布尔值
e	|	科学计数
x	|	整数（16进制）
h	|	散列码（16进制）

System.out.printf 和 System.out.foramt 方法的格式化输出就是调用了Formatter工具类。其中System.out.printf 的源码实际就是调用用了System.out.foramt方法。
System.out.printf 源码如下：
```java
public PrintStream printf(String format, Object ... args) {
      return format(format, args);
}
```

### 泛型
定义：类或方法在定义时类型不确定。使用时才确定类型。

JDK5引入泛型是一个很大的功能增强，使用也比较广泛。使用多态进行数据传输时，JDK5之前使用Object传输，然后进行向下转型，这里可能在运行期强转失败抛ClassCastException异常，导致程序异常终止。引入泛型可以将此运行期异常转移到编译异常，在编写代码时就可以检测出问题。
泛型可以避免向下转型的安全隐患
```java
import java.util.ArrayList;
import java.util.List;

public class Test {
    public static void main(String[] args) {

        //此处只能在运行期报ClassCastException异常。
        Object obj = new String();
        Integer i = (Integer) obj;

        //泛型
        List<String> list = new ArrayList<String>();
        list.add("abc");
        list.add("efg");
        //此处编译不通过，类型检测不能通过，在编译期就能解决错误。
//        list.add(1);
//        list.add(false);
//        list.add(0.5);

    }
}
```
泛型关键技术：
1. 通配符问号(?)表示任意类型.如"List<?>"表示可以存放任意对象类型的List，和List<Object>一样。

2. 通配符可以连接限制符：extends 和 super
如:上边界List<? extends Parent> 申明的List只能存放Parent及其子类，而下边界 List<? super Child> 申明的List只能存放Child及其父类。

？extends类 . 表示泛型上限，类与方法均可使用。
？super 类 . 表示泛型下限 ，只能用在方法级别

3. 通用类型，通常使用一个大写字母表示如:T(这里可以使用任意符合命名规则的字符都可以，不过我通常喜欢使用一个大写字母表示)，它能代表任何类型。


如果使用通用类型申明一个变量，那么必须在类申明后面加上<T>,尖括号里面的符号必须是前面申明的通用类型变量，如果有多个可以使用逗号分割如：<T,D>；

如果使用通用类型申明一个方法返回值或者方法参数，要么如上在类申明后加使用<>申明通用类型，要么在方法前申明通用类型。

```java
//在类申明后申明通用类型T，则可以在变量、方法返回值和方法参数使用
public class Test<T> {

    //在变量处使用通用类型，且并需在类申明后申明通用类型
    T t;
    //此处报错因为，变量通用类型必须在类申明后申明
//  E e;

    //在方法返回值处使用通用类型
    public T getT() {
        return t;
    }

    //在方法参数使用通用类型
    public String getType(T t) {
        return t.getClass().getSimpleName();
    }

    //方法返回值通用类型 和 方法参数通用类型 可以在方法前申明
    public <E> E getE(E e) {
        return e;
    }
}
```
- <>内的T被称作是类型参数，用于指代任意类型，不包括基本类型；
- 一般使用单个大写字母
- 引入泛型后，一个泛型类的类型在使用时已经确定好，因此无需向下转型；（解除了向下转型的安全隐患）
- 技巧：IDEA变量自动补全：ctrl+alt+V

**类型擦除（向下兼容）**

泛型只存在于编译阶段，进入JVM之前，与泛型相关的信息会被完全擦除。
在JVM看来，根本就不存在泛型的概念；

泛型类在进行类型擦除时，若未指定泛型的上限，泛型相关信息会被擦除为Object类，否则，擦除为相应的类型上限。


### ProcessBuilder
 ProcessBuilder可以用来创建操作系统进程，它的每一个实例管理着Process集合，start()方法可以创建一个新的Process实例
 主要方法：
- ProcessBuilder的start()方法：执行命令并返回一个Process对象；

- ProcessBuilder的environment()方法：返回运行进程的环境变量Map<String,String>集合；

- ProcessBuilder的directory()方法：返回工作目录；

- Process的getInputStream()方法：获得进程的标准输出流；

- Process的getErrorStream()方法：获得进程的错误输出流。
```java
import java.io.IOException;
import java.io.InputStream;
import java.util.Map;

public class Test {
    public static void main(String[] args) {
        //创建进程
        ProcessBuilder processBuilder = new ProcessBuilder("ipconfig","/all");
        //获取当前进程的环境变量
        Map<String, String> map = processBuilder.environment();
        Process process = null;
        try {
            //执行 ipconfig/all 命令并返回Process对象
            process = processBuilder.start();
        } catch (IOException e) {
            e.printStackTrace();
        }
        //获取进程标准输出流
        InputStream in = process.getInputStream();
        StringBuffer sb = new StringBuffer();
        int readbytes = -1;
        byte[] b = new byte[1024];
        try{
            while((readbytes = in.read(b)) != -1){
                sb.append(new String(b,0,readbytes));
            }
        }catch(IOException e1){
        }finally {
            try{
                in.close();
            }catch (IOException e2){
            }
        }
        System.out.println(sb.toString());

    }
}
```
### 内省  Introspector

内省机制就是基于反射的基础，Java语言对JavaBean类属性、事件的一种缺省处理方法。


JavaBean是一种特殊的类，主要用于传递数据信息，这种类中的方法主要用于访问私有的字段，且方法名符合某种命名规则。

#### JavaBean
- 性都是私有的；
- 有无参的public构造方法；
- 对私有属性根据需要提供 公有的getXxx方法以及setXxx方法；例如属性名称为name,则有getName方法返回属性name值，setName方法设置name值；注意方法的名称通常是get或set加上属性名称，并把属性名称的首字母大写；这些方法称为getters/setters；getters必须有返回值没有方法参数；setter值没有返回值，有方法参数；

符合这些特征的类，被称为JavaBean；

#### 内省与反射的区别

反射可以操作各种类的属性，而内省就是基于反射来操作JavaBean的属性。



#### 内省获得的JavaBean属性

内省中确定属性是根据 set 或 get 方法，如果只要类中有getXXX方法，或者setXXX方法，或者同时有getXXX及setXXX方法，其中getXXX方法没有方法参数，有返回值；setXXX方法没有返回值，有一个方法参数；那么内省机制就认为XXX为一个属性。



如果在两个模块之间传递信息，可以将信息封装进JavaBean中，这种对象称为“值对象”(Value Object)，或“VO”，这些信息储存在类的私有变量中，通过set()、get()获得，如下所示：
```java
public class User {

    private String name;
    private int age;
    private String address;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getAge() {
        return age;
    }

    public void setAge(int age) {
        this.age = age;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }
}
```
在类User中有属性name、age和address。
通过 getName/setName来访问name属性，getAge/setAge来访问age属性，getAddress/setAddress来访问address属性，这是我们默认遵循的规则。
Java JDK中提供了一套 API用来访问某个属性的 getter/setter 方法，这就是内省.

#### 使用内省操作User对象
```java
public static void main(String[] args) throws IntrospectionException {
		//获取所要操作的JavaBean信息
		//BeanInfo是接口不能直接实例化，只能通过Introspector的getBeanInfo获得实例
		BeanInfo beanInfo = Introspector.getBeanInfo(User.class);
		
		//获得JavaBean中的属性描述符，通过PropertyDescriptor可以获得属性一些信息
		//属性是通过set和get方法确定的，与是否在类中定义无关，详情见文章开头
		PropertyDescriptor[] pds = beanInfo.getPropertyDescriptors();
		for (PropertyDescriptor pd : pds) {
			//获取属性名字
			System.out.println(pd.getName());
			//获取属性类型
			System.out.println(pd.getPropertyType());
			//获取get方法，如果不存在则返回null
			System.out.println(pd.getReadMethod());
			//获取set方法，如果不存在则返回null
			System.out.println(pd.getWriteMethod());
		}                    
	}
```
#### 通过内省可以很轻松的将form表单的内容填充进对象里面，比反射轻松省力多了
```java
public void insertUser(HttpServletRequest request) {
		User user = new User();
		
		BeanInfo beanInfo = Introspector.getBeanInfo(User.class);
		
		PropertyDescriptor[] pds = beanInfo.getPropertyDescriptors();
		for (PropertyDescriptor pd : pds) {
            //需要注意参数类型匹配问题，此问题不在本文讨论范围里面
			pd.getWriteMethod().invoke(user, request.getParameter(pd.getName()));
		}
	}
```

内省类库：PropertyDescriptor类库：
PropertyDescriptor类表示JavaBean类通过存储器导出一个属性。主要方法：
- getPropertyType()：获得属性的class对象；
- getReadMeth()：获得读取属性值的方法，返回Method对象；
- getWriteMethod()：获得写入属性值的方法，返回Method对象；
- hashCode()：获得对象的哈希值；
- setReadMethod(Method readMethod)：设置读取属性值；
- setWriteMethod(Method writeMethod)：设置写入属性值。

```java
//创建Person对象，并赋初始值
Person person = new Person();
person.setName("niannianjiuwang");
PropertyDescriptor propertyDescriptor = new PropertyDescriptor("name",Person.class);
//获得属性的Class对象
System.out.println("Class: " + propertyDescriptor.getPropertyType().getSimpleName());
Method method = propertyDescriptor.getReadMethod();
System.out.println("Value: " + method.invoke(person));
System.out.println("HashCode: " + propertyDescriptor.hashCode());
```


### 线程并发库（JUC）
JDK5提供了线程处理的高级功能，在（java.util.concurrent）包下。包括：
#### 线程护斥：Lock 类、ReadWriteLock接口
Lock的方法：

ReadWriteLock的方法：

#### 线程通信：Condition接口
Condition的方法：

#### 线程池：ExecutorService接口
ExecutorService的方法：

#### 同步队列：ArrayBlockingQueue类
ArrayBlockingQueue的方法：

#### 同步集合：ConcurrentHashMap类、CopyOnWriteArrayList类
ConcurrentHashMap相当于一个HashMap集合，但前者是线程安全的，所以性能上比后者略低。

CopyOnWriteArrayList相当于一个ArrayList集合，前者其所有可变操作(add和set等)都是通过对底层的数组进行一次复制来实现，所以代价非常昂贵。

#### 线程同步工具：Semaphore类
Semaphore的方法：

#### 关于JUC的并发库类容非常的多，这里将不一一列举

### 监控和管理虚拟机
在JDK5中使用Bean监控和管理Java虚拟机，java.lang.management.ManagementFactory是管理Bean的工厂类，通过它的get系列方法能够获得不同的管理Bean的实例。

ManagementFactory的方法：

- MemoryMXBean对象：该Bean用于管理Java虚拟机的内存系统，一个Java虚拟机具有一个实例。

- ClassLoadingMXBean对象：该Bean用于管理Java虚拟机的类加载系统，一个Java虚拟机具有一个实例。

- TreadMXBean对象：该Bean用于管理Java虚拟机的线程系统，一个Java虚拟机具有一个实例。

- RuntimeMXBean对象：该Bean用于管理Java虚拟机的线程系统，一个Java虚拟机具有一个实例。

- OperatingSystemMXBean对象：该Bean用于管理操作系统，一个Java虚拟机具有一个实例。


- CompilationMXBean对象：该Bean用于管理Java虚拟机的编译系统，一个Java虚拟机具有一个实例。

- GarbageCollectorMXBean对象：该Bean用于管理Java虚拟机的垃圾回收系统，一个Java虚拟机具有一个或者多个实例。

```java
import java.lang.management.*;
import java.util.List;

public class Test {
    public static void main(String[] args){
        //Java虚拟机的内存系统
        MemoryMXBean memoryMXBean = ManagementFactory.getMemoryMXBean();
        System.out.println("虚拟机的堆内存使用量: " + memoryMXBean.getHeapMemoryUsage());
        System.out.println("虚拟机的非堆内存使用量: " + memoryMXBean.getNonHeapMemoryUsage());
        //Java虚拟机的类加载系统
        ClassLoadingMXBean classLoadingMXBean = ManagementFactory.getClassLoadingMXBean();
        System.out.println("当前加载到Java虚拟机中的类的数量: " + classLoadingMXBean.getLoadedClassCount());
        System.out.println("自Java虚拟机开始执行到目前已经加载的类的总数: " + classLoadingMXBean.getTotalLoadedClassCount());
        //Java虚拟机的线程系统
        ThreadMXBean threadMXBean = ManagementFactory.getThreadMXBean();
        System.out.println("当前线程的总CPU时间: " + threadMXBean.getCurrentThreadCpuTime());
        System.out.println("当前活动线程的数目，包括守护线程和非守护线程: " + threadMXBean.getThreadCount());
        //Java虚拟机的线程系统
        RuntimeMXBean runtimeMXBean = ManagementFactory.getRuntimeMXBean();
        System.out.println("当前Java库路径: " + runtimeMXBean.getLibraryPath());
        System.out.println("当前Java虚拟机实现提供商: " + runtimeMXBean.getVmVendor());
        //操作系统
        OperatingSystemMXBean operatingSystemMXBean = ManagementFactory.getOperatingSystemMXBean();
        System.out.println("当前Java虚拟机可以使用的处理器数目: " + operatingSystemMXBean.getAvailableProcessors());
        System.out.println("当前操作系统名称: " + operatingSystemMXBean.getName());
        //Java虚拟机的编译系统
        CompilationMXBean compilationMXBean = ManagementFactory.getCompilationMXBean();
        System.out.println("当前(JIT)编译器的名称: " + compilationMXBean.getName());
        System.out.println("当前即时(JIT)编译器的名称: " + compilationMXBean.getTotalCompilationTime());
        //Java虚拟机的垃圾回收系统
        List<GarbageCollectorMXBean> garbageCollectorMXBeanList = ManagementFactory.getGarbageCollectorMXBeans();
        for (GarbageCollectorMXBean garbageCollectorMXBean : garbageCollectorMXBeanList) {
            System.out.println("当前垃圾收集器的名字: " + garbageCollectorMXBean.getName());
            System.out.println("当前垃圾收集器累计回收总次数: " + garbageCollectorMXBean.getCollectionCount());
            System.out.println("当前垃圾收集器累计回收总时间: " + garbageCollectorMXBean.getCollectionTime());
        }
    }
}
```

### 元数据(注解)
元数据也可以叫注解，这个名字估计容易理解，格式：@注解名

注解的作用范围，可以通过java.lang.annotation.ElementType查看：

- TYPE：类、接口（包括注释类型）或enum声明

- FIELD：字段声明（包括enum常量）

- METHOD：方法申明

- PARAMETER：参数申明

- CONSTRUCTOR：构造器申明

- LOCAL_VARIABLE：局部变量申明

- ANNOTATION_TYPE：注解类型申明

- PACKAGE：包申明

JDK内置三种标准注解： 

@Override：　注解只能使用在方法上，表示当前的方法定义将覆盖超类中的方法。如果你不小心拼写错误，或者方法签名对不上被覆盖的方法，编译器就会发出错误的提示

@Deprecated：　注解可使用在构造器、字段、局部变量、方法、包、类接口以及枚举上，表示被弃用，不鼓励使用，编译器会发出警告信息。通常是因为它是危险的，或则因为有更好的选择。

@SuppressWarnings：注解可以使用在构造器、字段、局部变量、方法、类接口以及枚举上，必须指定value值，关闭不当的编译器警告信息。告诉编译器不提示某某警告信息。

注意的几个问题：

1.  当注解的元素没有默认值的时候，在使用的时候必须为其指定初始值

2.  如果注解元素有了初始值，那么在使用的时候可以为其赋新的值，否则将使用默认值

3.  一个较为特殊的情况：注解元素当且仅当其只有一个元素且名称为value时，在使用该注解的时候为其赋值时可以不用写属性（元素）名称

元注解：

java内置了4种元注解，元注解负责注解其它的注解，可以理解成java中用来注解Annotation的Annotation

@Retention：　保留策略，表示注解有多长保留，先了解JAVA文件的三个时期：SOURCE 源文件期`（*.java文件） -> CLASS 编译器编译期（*.class文件） -> RUNTIME` jvm运行时期。

@Target：　表示注解使用的上下文，TYPE、FIELD、METHOD、PARAMETER、CONSTRUCTOR、LOCAL_VARIABLE、ANNOTATION_TYPE和PACKAGE。详细说明返回看注解的作用范围。

@Documented：　表示将被javadoc记录。

@Inherited：　表明注释类型是自动继承的。如果一个继承的元注释出现在注释类型上声明，用户在一个类上查询注释类型声明，类声明没有这种类型的注释，然后该类的超类将自动被查询注释类型。这个过程将会重复直到这个注释类型被找到，或者类层次结构的顶部（对象）是达到了。如果没有超类有这种类型的注释，那么查询将表明的类没有这样的注释。



## JDK1.6
提供动态语言支持、提供编译API和卫星HTTP服务器API，改进JVM的锁，同步垃圾回收，类加载
JDK6 改动最大的就是java GUI界面的显示，JDK6支持最新的windows vista系统的Windows Aero视图效果，而JDK5不支持。
### Desktop类和SystemTray类 
JDK6在java.awt包下新增了两个类：Desktop和SystemTray
#### Desktop类：
允许 Java 应用程序启动已在本机桌面上注册的关联应用程序来处理 URI 或文件。
Desktop类的主要方法：
- browse(URI uri)： 使用默认浏览器打开uri资源。

- checkActionSupport(Action actionType)： 检查是否支持的Action。

- open(File file)： 启动关联应用程序来打开文件。

- edit(File file)： 启动关联编辑器应用程序并打开一个文件编辑。

- print(File file)： 使用关联应用程序打印文件。

- mail()： 启动用户默认邮件客户端的邮件组合窗口。

- mail(URI mailtoURI)： 启动用户默认邮件客户端的邮件组合窗口, 填充由 mailto:URI 指定的消息字段。
```java
import java.awt.*;
import java.io.File;
import java.io.IOException;
import java.net.URI;

public class Test {
    public static void main(String[] args) throws IOException {
        //先判断当前平台是否支持Desktop类
        if (Desktop.isDesktopSupported()) {
            //获取Desktop实例
            Desktop desktop = Desktop.getDesktop();
            //使用默认浏览器打开链接
            desktop.browse(URI.create("http://www.cnblogs.com/peter1018"));
            // 打开指定文件/文件夹
            desktop.open(new File("D:\\"));

            //......

            //其他方法可以自行演示...
        }else {
            System.out.println("当前平台不支持 Desktop类！");
        }

    }

}
```
#### SystemTray类：
代表一个系统托盘桌面。在微软的Windows上，它被称为“任务栏”状态区域，在Gnome上，它被称为“通知”在KDE上，它被称为“系统托盘”。该系统托盘由在桌面上运行的所有应用程序共享。
SystemTray类的主要方法：

- isSupported() ： 检测系统是否支持SystemTray类。

- getSystemTray() ： 获得SystemTray实例。

- add(TrayIcon trayIcon)： 在“系统托盘”上添加一个“TrayIcon”。TrayIcon对象代表一个托盘图标，它可以有一个工具提示（文本），一个图像，一个弹出框菜单，以及一组与之相关的侦听器。托盘图标一旦出现，就会在系统托盘中显示。

- remove(TrayIcon trayIcon)： 删除指定的“TrayIcon”。

- getTrayIcons()： 返回一个由这个托盘添加到托盘上的所有图标的数组应用程序。

- getTrayIconSize()：  以像素为单位，返回在系统托盘中占据的一个托盘图标的大小。

- addPropertyChangeListener(String propertyName,PropertyChangeListener listener）： 将 PropertyChangeListener 添加到监听器

- removePropertyChangeListener(String propertyName,PropertyChangeListener listener)： 将PropertyChangeListener 从监听器移除

- getPropertyChangeListeners(String propertyName) ： 返回所有已关联的监听器数组
```java
import javax.swing.*;
import java.awt.*;

public class Test {
    public static void main(String[] args) throws AWTException {
        //先判断当前平台是否支持SystemTray类
        if (SystemTray.isSupported()) {
            //获取SystemTray实例
            SystemTray systemTray = SystemTray.getSystemTray();
            //获得托盘显示图标
            ImageIcon imageIcon=new ImageIcon("D:\\icon.jpg");
            //获得Image对象
            Image icon=imageIcon.getImage();
            //任务栏程序托盘图标
            TrayIcon trayicon = new TrayIcon(icon,"JAVA系统托盘");
            //关键点，设置托盘图标的自适应属性，这样才能在系统显示托盘处正常显示出需要的图片。
            trayicon.setImageAutoSize(true);
            //添加系统托盘
            systemTray.add(trayicon);
            //......
            //其他方法可以自行演示...
            //移除系统托盘
            systemTray.remove(trayicon);
        }else {
            System.out.println("当前平台不支持 SystemTray类！");
        }

    }

}
```

### JAXB2实现对象与XML之间的映射
JAXB是Java Architecture for XML Binding的缩写，可以将一个Java对象转变成为XML格式，反之亦然。我们把对象与关系数据库之间的映射称为ORM, 其实也可以把对象与XML之间的映射称为OXM(Object XML Mapping)。原来JAXB是Java EE的一部分，在JDK6中，SUN将其放到了Java SE中，这也是SUN的一贯做法。JDK6中自带的这个JAXB版本是2.0, 比起1.0(JSR 31)来，JAXB2(JSR 222)用JDK5的新特性Annotation来标识要作绑定的类和属性等，这就极大简化了开发的工作量。实际上，在Java EE 5.0中，EJB和Web Services也通过Annotation来简化开发工作。另外,JAXB2在底层是用StAX(JSR 173)来处理XML文档。

重要的注解和说明

注解 |	说明
--|--
@XmlRootElement | 将类或enum类型映射到XML元素，类名或名作为根节点
@XmlAttribute | 将JavaBean属性映射到XML属性
@XmlElement | 将JavaBean属性映射到源自属性名的XML元素
@XmlAnyAttribute | 将JavaBean属性映射到通配符属性的映射
@XmlAnyElement | 将JavaBean属性映射到XML infoset或JAXB
@XmlElements | 一个用于多XmlElement注解的容器。
@XmlID | 将JavaBean属性映射到XML ID
@XmlIDREF | 将JavaBean属性映射到XML IDREF
@XmlList | 用来将一个属性映射到一个List
@XmlSchema | 将包映射到XML名称空间
@XmlTransient | 该属性无需映射到XML

Address类
```java
import javax.xml.bind.annotation.XmlAttribute;
import javax.xml.bind.annotation.XmlElement;

public class Address {

    @XmlAttribute
    String province;
    @XmlElement
    String area;
    @XmlElement
    String street;

    public Address() {}

    public Address(String province, String area, String street) {
        this.province = province;
        this.area = area;
        this.street = street;
    }

    @Override
    public String toString() {
        return "Address{" +
                "province='" + province + '\'' +
                ", area='" + area + '\'' +
                ", street='" + street + '\'' +
                '}';
    }
}
```
Person类
```java
import javax.xml.bind.annotation.XmlAttribute;
import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Person{

    @XmlAttribute
    private String name;

    @XmlElement
    private int age;

    @XmlElement
    private Address address;

    public Person() {
    }

    public Person(String name, int age, Address address) {
        this.name = name;
        this.age = age;
        this.address = address;
    }

    @Override
    public String toString() {
        return "Person{" +
                "name='" + name + '\'' +
                ", age=" + age +
                ", address=" + address +
                '}';
    }
}
```
测试类：
```java
import javax.xml.bind.JAXBContext;
import javax.xml.bind.JAXBException;
import javax.xml.bind.Marshaller;
import javax.xml.bind.Unmarshaller;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;

public class Test {
    public static void main(String[] args) throws JAXBException, IOException {
        JAXBContext jaxbContext = JAXBContext.newInstance(Person.class);

        //根据Person对象转换为person.xml文件
        Marshaller marshaller = jaxbContext.createMarshaller();
        Address address = new Address("广东省","深圳市","宝安区");
        Person person = new Person("niannianjiuwang", 20, address);
        FileWriter fileWriter = new FileWriter("D:\\person.xml");
        marshaller.marshal(person, fileWriter);

        //根据person.xml文件转换为对象
        FileReader fileReader = new FileReader("D:\\person.xml");
        Unmarshaller unmarshaller = jaxbContext.createUnmarshaller();
        Person personNew = (Person) unmarshaller.unmarshal(fileReader);
        System.out.println(personNew);
    }
}
```
person.xml文件
```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<person name="niannianjiuwang">
    <age>20</age>
    <address province="广东省">
        <area>深圳市</area>
        <street>宝安区</street>
    </address>
</person>
```

### StAX

StAX的来历 ：

在JAXP1.3([JSR 206](https://jcp.org/en/jsr/detail?id=206))有两种处理XML文档的方法:DOM(Document Object Model)和SAX(Simple API for XML).由于JDK6.0中的JAXB2([JSR 222](https://jcp.org/en/jsr/detail?id=222))和JAX-WS 2.0([JSR 224](https://jcp.org/en/jsr/detail?id=224))都会用到StAX所以Sun决定把StAX加入到JAXP家族当中来，并将JAXP的版本升级到1.4(JAXP1.4是JAXP1.3的维护版本). JDK6里面JAXP的版本就是1.4。

StAX简介：

StAX是Streaming API for XML的缩写，它包含在2004年3月的JSR 173 中。StAX是JDK6.0中包含的新特性。
在推式模型中，直到整个XML文件全部被解析，解析器才会产生解析事件。而拉式解析由应用程序来控制，也就是说解析事件由应用程序产生。这就意味着，你可以暂缓解析、解析时跳过某个元素或者同时解析多个文件。用DOM解析式要先将XML文件解析成DOM结构，因此降低了解析效率。使用StAX，解析事件在XML文件解析过程中产生。

下面是各种解析方法之间的比较：
XML Parser API Feature Summary

Feature |	StAX|	SAX|	DOM|	TrAX
--|--|--|--|--
API Type|	Pull, streaming	|Pull, streaming	|In memory tree	|XSLT Rule
Ease of Use|	High|	Medium|	High|	Medium
XPath Capability|	No|	No|	Yes	|Yes
CPU and Memory Efficiency|	Good	|Good	|Varies|	Varies
Forward Only|	Yes|	Yes|	No|	No
Read XML|	Yes|	Yes	|Yes|	Yes
Write XML|	Yes	|No	|Yes|	Yes
Create, Read, Update, Delete|	No|	No|	Yes	|No

StAX API同样也在JWSDP（Java Web Services Developer Pack ）1.6中得到实现，你可以在包javax.xml.stream 中找到它。XMLStreamReader接口用来解析XML文件。XMLStreamWriter接口用来生成XML文件，XMLEventReader用一个对象事件迭代器来解析XML事件。与之相反，XMLStreamReader采用的是游标机制。


```java
import javax.xml.stream.*;
import javax.xml.stream.events.XMLEvent;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;

public class Test {
    public static void main(String[] args) throws FileNotFoundException, XMLStreamException {
        writeXMLByStAX();//用XMLStreamWriter写xml文档
        readXMLByStAX();//用XMLEventReader解析xml文档
    }

    /**
     * 通过StAX读XML
     */
    private static void readXMLByStAX() throws XMLStreamException, FileNotFoundException {
        XMLInputFactory xmlif = XMLInputFactory.newInstance();
        XMLEventReader xmler = xmlif.createXMLEventReader(new FileInputStream("D:\\write.xml"));
        XMLEvent event;
        StringBuffer parsingResult = new StringBuffer();
        while (xmler.hasNext()) {
            event = xmler.nextEvent();
            parsingResult.append(event.toString());
        }
        System.out.println(parsingResult);
    }

    /**
     * 通过StAX写XML
     */
    private static void writeXMLByStAX() throws FileNotFoundException, XMLStreamException {
        XMLOutputFactory xmlOutputFactory = XMLOutputFactory.newFactory();
        XMLStreamWriter xmlStreamWriter = xmlOutputFactory.createXMLStreamWriter(new FileOutputStream("D:\\write.xml"));

        // 写入默认的 XML 声明到xml文档
        xmlStreamWriter.writeStartDocument();
        // 写入注释到xml文档
        xmlStreamWriter.writeComment("testing comment");
        // 写入一个catalogs根元素
        xmlStreamWriter.writeStartElement("catalogs");
        xmlStreamWriter.writeNamespace("myUrl", "http://www.cnblogs.com/peter1018");
        xmlStreamWriter.writeAttribute("name","niannianjiuwang");
        // 写入子元素catalog
        xmlStreamWriter.writeStartElement("http://www.cnblogs.com/peter1018", "catalog");
        xmlStreamWriter.writeAttribute("id","StAX");
        xmlStreamWriter.writeCharacters("Apparel");
        // 写入catalog元素的结束标签
        xmlStreamWriter.writeEndElement();
        // 写入catalogs元素的结束标签
        xmlStreamWriter.writeEndElement();
        // 结束 XML 文档
        xmlStreamWriter.writeEndDocument();
        xmlStreamWriter.close();
    }
}
```
输出结果：

write.xml文件内容：

```xml
<?xml version="1.0" ?><!--testing comment--><catalogs xmlns:myUrl="http://www.cnblogs.com/peter1018" name="niannianjiuwang"><myUrl:catalog id="StAX">Apparel</myUrl:catalog></catalogs>
```
控制台打印结果：
```
<?xml version="1.0" encoding='UTF-8' standalone='no'?><!--testing comment--><catalogs name='niannianjiuwang' xmlns:myUrl='http://www.cnblogs.com/peter1018'><['http://www.cnblogs.com/peter1018']:myUrl:catalog id='StAX'>Apparel</['http://www.cnblogs.com/peter1018']:myUrl:catalog></catalogs>ENDDOCUMENT
```

### Compiler API
Java编程语言编译器javac读取以Java编程语言编写的源文件，并将它们编译为字节码class文件。或者，编译器也可以使用注解找到源文件和类文件并使用comiler API。编译器是一个命令行工具，但也可以使用Java compiler API调用。原话：（[官网介绍](https://docs.oracle.com/javase/6/docs/technotes/guides/javac/)）

我们可以用JDK6 的Compiler API(JSR 199)去动态编译Java源文件，Compiler API结合反射功能就可以实现动态的产生Java代码并编译执行这些代码，有点动态语言的特征。这个特性对于某些需要用到动态编译的应用程序相当有用， 比如JSP Web Server，当我们手动修改JSP后，是不希望需要重启Web Server才可以看到效果的，这时候我们就可以用Compiler API来实现动态编译JSP文件，当然，现在的JSP Web Server也是支持JSP热部署的，现在的JSP Web Server通过在运行期间通过Runtime.exec或ProcessBuilder来调用javac来编译代码，这种方式需要我们产生另一个进程去做编译工作，不够优雅而且容易使代码依赖与特定的操作系统；Compiler API通过一套易用的标准的API提供了更加丰富的方式去做动态编译,而且是跨平台的。


运行代码之前必须在D目录下存放TestObject.java文件
```java
import javax.tools.JavaCompiler;
import javax.tools.JavaFileObject;
import javax.tools.StandardJavaFileManager;
import javax.tools.ToolProvider;
import java.io.IOException;

public class Test {
    public static void main(String[] args) throws IOException {
        JavaCompiler compiler = ToolProvider.getSystemJavaCompiler();
        StandardJavaFileManager standardJavaFileManager = compiler.getStandardFileManager(null,null,null);
        Iterable<? extends JavaFileObject> sourcefiles = standardJavaFileManager.getJavaFileObjects("D:\\TestObject.java");
        compiler.getTask(null, standardJavaFileManager, null, null, null, sourcefiles).call();
        standardJavaFileManager.close();

    }
}
```
输出结果：会在D目录下生成TestObject.class文件。

### 轻量级 Http Server API 

JDK6 提供了一个简单的Http Server API,据此我们可以构建自己的嵌入式Http Server,它支持Http和Https协议,提供了HTTP1.1的部分实现，没有被实现的那部分可以通过扩展已有的Http Server API来实现,必须自己实现HttpHandler接口,HttpServer会调用HttpHandler实现类的回调方法来处理客户端请求,我们把一个Http请求和它的响应称为一个交换,包装成HttpExchange类,HttpServer负责将HttpExchange传给 HttpHandler实现类的回调方法。

```java
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

import java.io.*;

public class TestHandler implements HttpHandler {


    @Override
    public void handle(HttpExchange httpExchange) throws IOException {
        System.out.println("==进入Hadnler方法");
        String responseMsg = "OK";   //响应信息
        InputStream in = httpExchange.getRequestBody(); //获得输入流
        BufferedReader reader = new BufferedReader(new InputStreamReader(in));
        String temp = null;
        while((temp = reader.readLine()) != null) {
            System.out.println("client request:"+temp);
        }
        httpExchange.sendResponseHeaders(200, responseMsg.length()); //设置响应头属性及响应信息的长度
        OutputStream out = httpExchange.getResponseBody();  //获得输出流
        out.write(responseMsg.getBytes());
        out.flush();
        httpExchange.close();
    }
}
```
```java
import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.spi.HttpServerProvider;

import java.io.IOException;
import java.net.InetSocketAddress;

public class Test {
    public static void main(String[] args) throws IOException {
        //启动服务，监听来自客户端的请求
        HttpServerProvider provider = HttpServerProvider.provider();
        //监听8888端口,能同时接受100个请求
        HttpServer httpserver =provider.createHttpServer(new InetSocketAddress(8888), 100);
        //将 /test  请求交给 TestHandler 处理器处理
        httpserver.createContext("/test", new TestHandler());
        httpserver.setExecutor(null);
        httpserver.start();
        System.out.println("server started");
    }
}
```
### 插入式注解处理 API(Pluggable Annotation Processing API)

插入式注解处理API([JSR 269](https://jcp.org/en/jsr/detail?id=269))提供一套标准API来处理Annotations([JSR 175](https://jcp.org/en/jsr/detail?id=175))。实际上JSR 269不仅仅用来处理Annotation，它还建立了Java 语言本身的一个模型，它把method， package， constructor， type, variable， enum,，annotation等Java语言元素映射为Types和Elements， 从而将Java语言的语义映射成为对象，我们可以在javax.lang.model包下面可以看到这些类. 所以我们可以利用JSR 269提供的API来构建一个功能丰富的元编程(metaprogramming)环境。JSR 269用Annotation Processor在编译期间而不是运行期间处理Annotation， Annotation Processor相当于编译器的一个插件，所以称为插入式注解处理。如果Annotation Processor处理Annotation时(执行process方法)产生了新的Java代码，编译器会再调用一次Annotation Processor，如果第二次处理还有新代码产生，就会接着调用Annotation Processor，直到没有新代码产生为止。每执行一次process()方法被称为一个"round"，这样整个Annotation processing过程可以看作是一个round的序列，JSR 269主要被设计成为针对Tools或者容器的API。

举个例子,我们想建立一套基于Annotation的单元测试框架(如TestNG),在测试类里面用Annotation来标识测试期间需要执行的测试方法,如下所示:
```java
@TestMethod
 public void testCheckName(){
       //do something here
 }
```
这时我们就可以用JSR 269提供的API来处理测试类，根据Annotation提取出需要执行的测试方法。

另一个例子是如果我们出于某种原因需要自行开发一个符合Java EE 5.0的Application Server(当然不建议这样做),我们就必须处理Common Annotations([JSR 250](https://jcp.org/en/jsr/detail?id=250)),Web Services Metadata([JSR 181](https://jcp.org/en/jsr/detail?id=181))等规范的Annotations,这时可以用JSR 269提供的API来处理这些Annotations. 在现在的开发工具里面,Eclipse 3.3承诺将支持JSR 269

下面我用代码演示如何来用JSR 269提供的API来处理Annotations和读取Java源文件的元数据(metadata)
```java
@SupportedAnnotationTypes("PluggableAPT.ToBeTested")//可以用"*"表示支持所有Annotations
@SupportedSourceVersion(SourceVersion.RELEASE_6)
public class MyAnnotationProcessor extends AbstractProcessor {
    private void note(String msg) {
        processingEnv.getMessager().printMessage(Diagnostic.Kind.NOTE, msg);
    }
    public boolean process(Set<? extends TypeElement> annotations, RoundEnvironment roundEnv) {
        //annotations的值是通过@SupportedAnnotationTypes声明的且目标源代码拥有的所有Annotations
        for(TypeElement te:annotations){
            note("annotation:"+te.toString());
        }
        Set<? extends Element> elements = roundEnv.getRootElements();//获取源代码的映射对象
        for(Element e:elements){
            //获取源代码对象的成员
            List<? extends Element> enclosedElems = e.getEnclosedElements();
            //留下方法成员,过滤掉其他成员
            List<? extends ExecutableElement> ees = ElementFilter.methodsIn(enclosedElems);
            for(ExecutableElement ee:ees){
                note("--ExecutableElement name is "+ee.getSimpleName());
                List<? extends AnnotationMirror> as = ee.getAnnotationMirrors();//获取方法的Annotations
                note("--as="+as); 
                for(AnnotationMirror am:as){
                    //获取Annotation的值
                    Map<? extends ExecutableElement, ? extends AnnotationValue> map= am.getElementValues();
                    Set<? extends ExecutableElement> ks = map.keySet();
                    for(ExecutableElement k:ks){//打印Annotation的每个值
                        AnnotationValue av = map.get(k);
                        note("----"+ee.getSimpleName()+"."+k.getSimpleName()+"="+av.getValue());
                    }
                }
            }
        }
        return false;
    }
}
```
```java
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
@interface ToBeTested{
    String owner() default "Chinajash";
    String group();
}
```
编译以上代码,然后再创建下面的Testing对象,不要编译Testing对象,我在后面会编译它

```java
public class Testing{    
    @ToBeTested(group="A")
    public void m1(){
    }
    @ToBeTested(group="B",owner="QQ")
    public void m2(){
    }    
    @PostConstruct//Common Annotation里面的一个Annotation
    public void m3(){
    }    
}
```
下面我用以下命令编译Testing对象
```
javac -XprintRounds -processor PluggableAPT.MyAnnotationProcessor Testing.java
```
-XprintRounds表示打印round的次数,运行上面命令后在控制台会看到如下输出:
```
Round 1:
        input files: {PluggableAPT.Testing}
        annotations: [PluggableAPT.ToBeTested, javax.annotation.PostConstruct]
        last round: false
Note: annotation:PluggableAPT.ToBeTested
Note: --ExecutableElement name is m1
Note: ")
Note: ----m1.group=A
Note: --ExecutableElement name is m2
Note: ", owner="QQ")
Note: ----m2.group=B
Note: ----m2.owner=QQ
Note: --ExecutableElement name is m3
Note: --as=@javax.annotation.PostConstruct
Round 2:
        input files: {}
        annotations: []
        last round: true
```
本来想用JDK6.0的Compiler API来执行上面编译命令,可是好像现在Compiler API还不支持-processor参数,运行时总报以下错误
```
Exception in thread "main" java.lang.IllegalArgumentException: invalid flag: -processor PluggableAPT.MyAnnotationProcessor
```
调用Compiler API的代码是这样的
```java
JavaCompiler compiler = ToolProvider.getSystemJavaCompiler();
StandardJavaFileManager fileManager = compiler.getStandardFileManager(null, null, null);
Iterable<? extends JavaFileObject> sourcefiles = fileManager.getJavaFileObjects("Testing.java");
Set<String> options = new HashSet<String>();
options.add("-processor PluggableAPT.MyAnnotationProcessor");
compiler.getTask(null, fileManager, null, options, null, sourcefiles).call();
```
### 用Console开发控制台程序
JDK6中提供 了java.io.Console类专用来访问基于字符的控制台设备。程序如果要与Windows下的cmd或者Linux下的Terminal交互,就可以用Console类代劳。但我们不总是能得到可用的Console，一个JVM是否有可用的Console依赖于底层平台和JVM如何被调用。如果JVM是在交互式命令行(比如Windows的cmd)中启动的，并且输入输出没有重定向到另外的地方，那么就可以得到一个可用的Console实例 。
```java
import java.io.Console;

public class Test {
    public static void main(String[] args){
        Console console = System.console();
        if(console!=null){//判断console是否可用
            String user = new String(console.readLine("Enter user:")); //读取整行字符
            String pwd = new String(console.readPassword("Enter passowrd:")); //读取密码,密码输入时不会显示
            console.printf("User is:"+user+"/n");
            console.printf("Password is:"+pwd+"/n");
        }else{
            System.out.println("JVM无法使用当前的Console");
        }
    }
}
```
在IntelliJ IDEA 中直接运行执行结果：
```
JVM无法使用当前的Console
```
在cmd里面直接 java Test 结果：红色字体是输入部分，密码后面输入默认是隐藏的
```
F:\test\>java Test
Enter user:niannianjiuwang
Enter passowrd:
User is:niannianjiuwang/nPassword is:123456/n
```

### 对脚本语言的支持（如: ruby, groovy, javascript）
JDK6增加了对脚本语言的支持([JSR 223](https://jcp.org/en/jsr/detail?id=223))，原理上是将脚本语言编译成字节码，这样脚本语言也能享用Java平台的诸多优势，包括可移植性，安全等，另外，由于现在是编译成字节码后再执行，所以比原来边解释边执行效率要高很多。加入对脚本语言的支持后，对

Java语言也提供了以下好处。
1. 许多脚本语言都有动态特性，比如，你不需要用一个变量之前先声明它，你可以用一个变量存放完全不同类型的对象，你不需要做强制类型转换，因为转换都是自动的。现在Java语言也可以通过对脚本语言的支持间接获得这种灵活性。
2. 可以用脚本语言快速开发产品原型，因为现在可以Edit-Run，而无需Edit-Compile-Run，当然，因为Java有非常好的IDE支持，我 们完全可以在IDE里面编辑源文件，然后点击运行(隐含编译)，以此达到快速开发原型的目的，所以这点好处基本上可以忽略。
3. 通过引入脚本语言可以轻松实现Java应用程序的扩展和自定义，我们可以把原来分布在在Java应用程序中的配置逻辑，数学表达式和业务规则提取出来，转用JavaScript来处理。

Sun的JDK6实现包含了一个基于[Mozilla Rhino](https://developer.mozilla.org/en-US/docs/Mozilla/Projects/Rhino)的 脚本语言引擎，支持JavaScript，这并不是说明JDK6只支持JavaScript，任何第三方都可以自己实现一个JSR-223兼容的脚本引擎 使得JDK6支持别的脚本语言，比如，你想让JDK6支持Ruby，那你可以自己按照JSR 223的规范实现一个Ruby的脚本引擎类，具体一点，你需要实现 javax.script.ScriptEngine(简单起见，可以继承 javax.script.AbstractScriptEngine) 和 javax.script.ScriptEngineFactory 两个接口。当然，在你实现自己的脚本语引擎之前，先到 [scripting.dev.java.net](http://www.oracle.com/splash/java.net/maintenance/index.html) project 这里看看是不是有人已经帮你做了工作，这样你就可以直接拿来用。
#### Scripting API
Scripting API是用于在Java里面编写脚本语言程序的API， 在Javax.script中可以找到Scripting API，我们就是用这个API来编写JavaScript程序，这个包里面有一个ScriptEngineManager类，它是使用Scriptng API 的入口，ScriptEngineManager可以通过Jar服务发现（service discovery）机制寻找合适的脚本引擎类(ScriptEngine)，使用Scripting API的最简单方式只需下面三步
1. 创建一个ScriptEngineManager对象
2. 通过ScriptEngineManager获得ScriptEngine对象
3. 用ScriptEngine的eval方法执行脚本
```java
import javax.script.Invocable;
import javax.script.ScriptEngine;
import javax.script.ScriptEngineManager;

public class Test {
    public static void main(String[] args){
        ScriptEngineManager factory = new ScriptEngineManager();
        ScriptEngine engine = factory.getEngineByName("JavaScript");
        String script;
        try {
            script = "print('Hello')";
            engine.eval(script);// 执行脚本
            script = "1-23*9/3+77";
            System.out.println(engine.eval(script).toString());// 不用对字符串做解析便可得到算式结果
            engine.put("a", "一个字符串");
            script = "print(a)";
            engine.eval(script);// 脚本调用java对象
            script = "function hello(name) { return 'Hello,' + name;}";
            engine.eval(script);
            Invocable inv = (Invocable) engine;
            System.out.println(inv.invokeFunction("hello", "Scripting"));//java调用脚本方法
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

}
```

###  Common Annotations

Common annotations原本是Java EE 5.0([JSR 244](https://jcp.org/en/jsr/detail?id=244))规范的一部分，现在SUN把它的一部分放到了Java SE 6.0中。随着Annotation元数据功能([JSR 175](https://jcp.org/en/jsr/detail?id=175))加入到Java SE 5.0里面，很多Java 技术(比如EJB，Web Services)都会用Annotation部分代替XML文件来配置运行参数（或者说是支持声明式编程，如EJB的声明式事务）， 如果这些技术为通用目的都单独定义了自己的Annotations，显然有点重复建设， 所以，为其他相关的Java技术定义一套公共的Annotation是有价值的，可以避免重复建设的同时，也保证Java SE和Java EE 各种技术的一致性。

下面列举出Common Annotations 1.0里面的10个Annotations Common Annotations 
注解| 说明
--|--
@Generated |生成的注释用于标记已生成的源代码。
@Resource| 用于标注所依赖的资源，容器据此注入外部资源依赖，有基于字段的注入和基于setter方法的注入两种方式 。
@Resources| 同时标注多个外部依赖，容器会把所有这些外部依赖注入 。
@PostConstruct|标注当容器注入所有依赖之后运行的方法，用来进行依赖注入后的初始化工作，只有一个方法可以标注为PostConstruct 。
@RunAs|执行期间应用程序的角色 。
@PreDestroy|当对象实例将要被从容器当中删掉之前，要执行的回调方法要标注为RunAs用于标注用什么安全角色来执行被标注类的方法，这个安全角色必须和Container 的Security角色一致的。
@RolesAllowed| 用于标注允许执行被标注类或方法的安全角色，这个安全角色必须和Container 的Security角色一致的 。
@PermitAll|允许所有角色执行被标注的类或方法 。
@DenyAll|不允许任何角色执行被标注的类或方法，表明该类或方法不能在Java EE容器里面运行 。
@DeclareRoles|用来定义可以被应用程序检验的安全角色，通常用isUserInRole来检验安全角色 。


注意：

- RolesAllowed,PermitAll,DenyAll不能同时应用到一个类或方法上 
- 标注在方法上的RolesAllowed,PermitAll,DenyAll会覆盖标注在类上的RolesAllowed,PermitAll,DenyAll 
- RunAs,RolesAllowed,PermitAll,DenyAll和DeclareRoles还没有加到Java SE 6.0上来 
- 处理以上Annotations的工作是由Java EE容器来做, Java SE 6.0只是包含了上面表格的前五种Annotations的定义类,并没有包含处理这些Annotations的引擎,这个工作可以由Pluggable Annotation Processing API([JSR 269](https://jcp.org/en/jsr/detail?id=269))来做





## JDK1.7
提供GI收集器、加强对非Java语言的调用支持（JSR-292,升级类加载架构
https://www.oracle.com/technetwork/java/javase/jdk7-relnotes-418459.html

*   [Swing](http://docs.oracle.com/javase/7/docs/technotes/guides/swing/enhancements-7.html)
*   [IO and New IO](http://docs.oracle.com/javase/7/docs/technotes/guides/io/enhancements.html#7)
*   [Networking](http://docs.oracle.com/javase/7/docs/technotes/guides/net/enhancements-7.0.html)
*   [Security](http://docs.oracle.com/javase/7/docs/technotes/guides/security/enhancements-7.html)
*   [Concurrency Utilities](http://docs.oracle.com/javase/7/docs/technotes/guides/concurrency/changes7.html)
*   [Rich Internet Applications (RIA)/Deployment](http://docs.oracle.com/javase/7/docs/technotes/guides/jweb/clientJRECapabilitiesCheatSheet.html)
    *   [Requesting and Customizing Applet Decoration in Dragg able Applets](http://docs.oracle.com/javase/tutorial/deployment/applet/draggableApplet.html#decoration)
    *   [Embedding JNLP File in Applet Tag](http://docs.oracle.com/javase/tutorial/deployment/deploymentInDepth/embeddingJNLPFileInWebPage.html)
    *   [Deploying without Codebase](http://docs.oracle.com/javase/tutorial/deployment/deploymentInDepth/deployingWithoutCodebase.html)
    *   [Handling Applet Initialization Status with Event Handlers](http://docs.oracle.com/javase/tutorial/deployment/applet/appletStatus.html)
*   [Java 2D](http://docs.oracle.com/javase/7/docs/technotes/guides/2d/enhancements70.html)
*   [Java XML](http://docs.oracle.com/javase/7/docs/technotes/guides/xml/enhancements.html) \- JAXP, JAXB, and JAX-WS
*   [Internationalization](http://docs.oracle.com/javase/7/docs/technotes/guides/intl/enhancements.7.html)
*   [java.lang Package](http://docs.oracle.com/javase/7/docs/technotes/guides/lang/enhancements.html#7)
    *   [Multithreaded Custom Class Loaders in Java SE 7](http://docs.oracle.com/javase/7/docs/technotes/guides/lang/cl-mt.html)
*   [Java Programming Language](http://docs.oracle.com/javase/7/docs/technotes/guides/language/enhancements.html#javase7)
    *   [Binary Literals](http://docs.oracle.com/javase/7/docs/technotes/guides/language/binary-literals.html)
    *   [Strings in switch Statements](http://docs.oracle.com/javase/7/docs/technotes/guides/language/strings-switch.html)
    *   [The try-with-resources Statement](http://docs.oracle.com/javase/7/docs/technotes/guides/language/try-with-resources.html)
    *   [Catching Multiple Exception Types and Rethrowing Exceptions with Improved Type Checking](http://docs.oracle.com/javase/7/docs/technotes/guides/language/catch-multiple.html)
    *   [Underscores in Numeric Literals](http://docs.oracle.com/javase/7/docs/technotes/guides/language/underscores-literals.html)
    *   [Type Inference for Generic Instance Creation](http://docs.oracle.com/javase/7/docs/technotes/guides/language/type-inference-generic-instance-creation.html)
    *   [Improved Compiler Warnings and Errors When Using Non-Reifiable Formal Parameters with Varargs Methods](http://docs.oracle.com/javase/7/docs/technotes/guides/language/non-reifiable-varargs.html)
*   [Java Virtual Machine (JVM)](http://docs.oracle.com/javase/7/docs/technotes/guides/vm/)
    *   [Java Virtual Machine Support for Non-Java Languages](http://docs.oracle.com/javase/7/docs/technotes/guides/vm/multiple-language-support.html)
    *   [Garbage-First Collector](http://docs.oracle.com/javase/7/docs/technotes/guides/vm/G1.html)
    *   [Java HotSpot Virtual Machine Performance Enhancements](http://docs.oracle.com/javase/7/docs/technotes/guides/vm/performance-enhancements-7.html)
*   [JDBC](http://docs.oracle.com/javase/7/docs/technotes/guides/jdbc/)


### 二进制字面值 
在Java SE 7，整数类型（byte，short，int 和 long）也可以使用二进制数。要指定二进制，请添加前缀0b或0B编号。以下示例显示了二进制：
```java
// 一个 8-bit 'byte' 值:
byte aByte = (byte)0b00100001;

// 一个 16-bit 'short' 值:
short aShort = (short)0b1010000101000101;

// 一些 32-bit 'int' 值:
int anInt1 = 0b10100001010001011010000101000101;
int anInt2 = 0b101;
int anInt3 = 0B101; // B可以是大写 或 小写.

// 一个 64-bit 'long' 值。 注意 'L' 后缀:
long aLong = 0b1010000101000101101000010100010110100001010001011010000101000101L;
```
二进制可以使数据之间的关系比以十六进制或八进制更明显。例如，以下数组中的每个连续数字都移动一位：
```java
public static final int [] phases = {
        0b00110001,
        0b01100010,
        0b11000100,
        0b10001001,
        0b00010011,
        0b00100110,
        0b01001100,
        0b10011000
}
```
在十六进制中，数字之间的关系并不明显：
```java
public static final int[] phases = {
        0x31, 0x62, 0xC4, 0x89, 0x13, 0x26, 0x4C, 0x98
}
```
您可以在代码中使用二进制常量来验证一个规范文档，（例如假想的8位微处理器的模拟器）进行验证：
```java
public State decodeInstruction(int instruction, State state) {
        if ((instruction & 0b11100000) == 0b00000000) {
            final int register = instruction & 0b00001111;
            switch (instruction & 0b11110000) {
                case 0b00000000: return state.nop();
                case 0b00010000: return state.copyAccumTo(register);
                case 0b00100000: return state.addToAccum(register);
                case 0b00110000: return state.subFromAccum(register);
                case 0b01000000: return state.multiplyAccumBy(register);
                case 0b01010000: return state.divideAccumBy(register);
                case 0b01100000: return state.setAccumFrom(register);
                case 0b01110000: return state.returnFromCall();
                default: throw new IllegalArgumentException();
            }
        } else {
            final int address = instruction & 0b00011111;
            switch (instruction & 0b11100000) {
                case 0b00000000: return state.jumpTo(address);
                case 0b00100000: return state.jumpIfAccumZeroTo(address);
                case 0b01000000: return state.jumpIfAccumNonzeroTo(address);
                case 0b01100000: return state.setAccumFromMemory(address);
                case 0b10100000: return state.writeAccumToMemory(address);
                case 0b11000000: return state.callTo(address);
                default: throw new IllegalArgumentException();
            }
        }
    }
```
您可以使用二进制来使位图更具可读性：
```java
public static final short[] HAPPY_FACE = {
        (short)0b0000011111100000,
        (short)0b0000100000010000,
        (short)0b0001000000001000,
        (short)0b0010000000000100,
        (short)0b0100000000000010,
        (short)0b1000011001100001,
        (short)0b1000011001100001,
        (short)0b1000000000000001,
        (short)0b1000000000000001,
        (short)0b1001000000001001,
        (short)0b1000100000010001,
        (short)0b0100011111100010,
        (short)0b0010000000000100,
        (short)0b0001000000001000,
        (short)0b0000100000010000,
        (short)0b0000011111100000
}
```
### switch 语句支持 String
Java编译器通过switch使用String对象的 if-then-else 语句比链式语句生成通常更高效的字节码。
### try-with-resources
try-with-resources可以自动关闭相关的资源（只要该资源实现了AutoCloseable接口,jdk7为绝大部分资源对象都实现了这个接口）

您可以在try-with-resources语句中声明一个或多个资源。以下示例：
```java
public static void writeToFileZipFileContents(String zipFileName, String outputFileName)
        throws java.io.IOException {

    java.nio.charset.Charset charset = java.nio.charset.Charset.forName("US-ASCII");
    java.nio.file.Path outputFilePath = java.nio.file.Paths.get(outputFileName);

    // Open zip file and create output file with try-with-resources statement

    try (
            java.util.zip.ZipFile zf = new java.util.zip.ZipFile(zipFileName);
            java.io.BufferedWriter writer = java.nio.file.Files.newBufferedWriter(outputFilePath, charset)
    ) {

        // Enumerate each entry

        for (java.util.Enumeration entries = zf.entries(); entries.hasMoreElements();) {

            // Get the entry name and write it to the output file

            String newLine = System.getProperty("line.separator");
            String zipEntryName = ((java.util.zip.ZipEntry)entries.nextElement()).getName() + newLine;
            writer.write(zipEntryName, 0, zipEntryName.length());
        }
    }
}
```
当跟随它的代码块正常结束或由于异常终止时，将按照BufferedWriter再ZipFile对象此顺序自动调用close方法。请注意，close资源的方法按照其创建的相反顺序进行调用。

注意：try-with-resources语句可以像普通try语句一样拥有catch和finally代码块。在try-with-resources语句中，任何catch或finally块在声明的资源关闭后才运行。

#### 抑制异常

可以在try-with-resources语句的代码块中抛出异常。在该示例中writeToFileZipFileContents，可以从该try块中引发异常，并且try在尝试关闭ZipFile和BufferedWriter对象时最多可以从 try-with-resources 语句中抛出两个异常。如果从try块中抛出异常，并且从try-with-resources语句抛出一个或多个异常，那么从try-with-resources语句抛出的异常将被抑制，并且该块引发的异常将被writeToFileZipFileContents方法抛出。您可以通过Throwable.getSuppressed从该try块引发的异常中调用方法来检索这些抑制的异常。

#### 实现AutoCloseable或Closeable接口的类
请参阅Javadoc AutoCloseable和Closeable接口以获取实现这些接口之一的类的列表。该Closeable接口继承了AutoCloseable接口。接口Closeable的close方法抛出IOException异常，而接口AutoCloseable的close方法抛出Exception异常。因此，AutoCloseable接口的子类可以重写此close方法的行为来抛出特殊的异常，例如IOException根本没有异常。

### catch 多个类型异常
#### catch 多种类型异常
```java
catch (IOException|SQLException ex) {
    logger.log(ex);
    throw ex;
}
```
#### 分析异常并重新抛出精准异常
与早期版本的JDK相比，JDK7 编译器对重新产生的异常执行更精确的分析。这使您可以在throws方法声明的子句中指定更具体的异常类型。
```java
static class FirstException extends Exception { }
static class SecondException extends Exception { }
public void rethrowException(String exceptionName) throws Exception {
    try {
        if (exceptionName.equals("First")) {
            throw new FirstException();
        } else {
            throw new SecondException();
        }
    } catch (Exception e) {
        throw e;
    }
}
```
这个例子的try块可能会抛出FirstException或者SecondException。假设你想在rethrowException方法中throws这些异常类型。在JDK 7之前的版本中，您不能这样做。由于catch子句的异常参数e是Exception类型，并且catch块重新抛出异常参数e，所以只能在rethrowException方法抛出Exception异常类型。

在JDK7中的编译器可以通过catch块确定FirstException和SecondException异常。即使该catch子句的异常参数e类型为Exception，编译器也可以确定它是FirstException实例或者是SecondException实例：

```java
public void rethrowException(String exceptionName)
  throws FirstException, SecondException {
    try {
      // ...
    }
    catch (Exception e) {
      throw e;
    }
  }
```

如果catch参数分配给catch块中的另一个值，则此分析将失效。但是，如果catch参数分配给另一个值，则必须在方法声明Exception的throws子句中指定异常类型。

详细地说，在JDK7及更高版本中，当您在catch子句中声明一个或多个异常类型并重新抛出由此catch块处理的异常时，编译器将验证重新抛出异常的类型是否满足以下条件：

- 该try块可以throw它。
- 没有前面其他catch可以处理它。
- 它是catch子句的异常参数之一的子类型或超类型。

JDK7编译器允许在rethrowException方法throws抛出指定Exception类型FirstException和SecondException。因为您可以重新抛出一个throws声明的任何类型的超类型。

在JDK7之前的版本中，您不能抛出该catch子句的异常参数之一的超类型。在JDK7之前的编译器Exception会在语句中生成错误未报告的异常;必须捕获或声明throw,编译器检查抛出异常的类型是否可分配给rethrowException方法声明的throws子句中声明的任何类型。然而，捕捉参数的类型e是Exception，这是一种超类型，而不是一个子类型FirstException和SecondException。

### 字面值中使用下划线
在JDK7中，任意数量的下划线字符（_）可以出现在字面值的任意位置。这个特性使您能够在字面值中分离数字组，这可以提高代码的可读性。

例如，如果您的代码包含具有许多数字的数字，则可以使用下划线字符以三个一组来分隔数字，这与使用逗号或空格等标点符号作为分隔符类似。

以下示例显示了可以在字面值中使用下划线的其他方法：
```java
long creditCardNumber = 1234_5678_9012_3456L;
long socialSecurityNumber = 999_99_9999L;
float pi =     3.14_15F;
long hexBytes = 0xFF_EC_DE_5E;
long hexWords = 0xCAFE_BABE;
long maxLong = 0x7fff_ffff_ffff_ffffL;
byte nybbles = 0b0010_0101;
long bytes = 0b11010010_01101001_10010100_10010010;
```
您只能在数字之间放置下划线; 你不能在下列地方放置下划线：

- 在数字的开头或结尾
- 与浮点数字中的小数点相邻
- 在F或L后缀之前
- 在期望一串数字的位置（In positions where a string of digits is expected）
下面的例子演示了字面值中有效和无效的下划线位置（突出显示）：
```java
float pi1 = 3_.1415F;      // 无效; 不能在小数点附近加下划线
float pi2 = 3._1415F;      // 无效; 不能在小数点附近加下划线
long socialSecurityNumber1
        = 999_99_9999_L;         // 无效; 在L后缀之前不能加下划线

int x1 = _52;              // 这是一个标识符，而不是字面值
int x2 = 5_2;              // OK（十进制文字）
int x3 = 52_;              // 无效; 不能在文字的末尾加上下划线
int x4 = 5_______2;        // OK（十进制文字）

int x5 = 0_x52;            // 无效; 不能在0x基数前缀中加下划线
int x6 = 0x_52;            // 无效; 不能在数字的开头加下划线
int x7 = 0x5_2;            // OK（十六进制文字）
int x8 = 0x52_;            // 无效; 不能在数字的末尾加上下划线

int x9 = 0_52;             // OK（八进制文字）
int x10 = 05_2;            // OK（八进制文字）
int x11 = 052_;            // 无效; 不能在数字的末尾加上下划线
```
### 类型推断
只要编译器可以从上下文中推断出类型参数，你就可以用一组空的类型参数（<>）来代替调用泛型类的构造函数所需的类型参数。这一对尖括号被非正式地称为钻石。

例如，请考虑以下变量声明：
```java
Map<String, List<String>> myMap = new HashMap<String, List<String>>();
```
在JDK7中，可以使用一组空的类型参数（<>）替换构造函数的参数化类型：
```java
Map<String, List<String>> myMap = new HashMap<>();
```

### 改进泛型类型可变参数
#### 堆污染
大多数参数化类型（如ArrayList<Number>和List<String>）都是非具体化的类型。非具体化类型是一种在运行时不确定的类型。在编译时，非具体化类型在此过程中经历了一个称为类型擦除的过程，擦除参数类型和类型参数相关的信息。这确保了在泛型之前创建的Java库和应用程序的二进制兼容性。因为类型擦除在编译时从参数类型中删除信息，所以这些类型是非具体化的。

当参数化类型的变量引用不是该参数化类型的对象时，会发生堆污染。这种情况只有在程序执行一些操作时才会发生，从而在编译时产生不受约束的警告。一个未经检查的警告如果产生，无论是在编译时（在编译时类型检查规则的限制范围内）或运行时。一个涉及参数化类型（例如，一次转换或方法调用）的操作的正确性是无法被验证的。

思考下面的例子：
```java
List l = new ArrayList<Number>();
List<String> ls = l;       // 未经检查的警告
l.add(0, new Integer(42)); // 另一个未经检查的警告
String s = ls.get(0);      // 抛出类型转换异常 ClassCastException
```
在类型擦除时，ArrayList<Number>类型和List<String>类型分别变成ArrayList和List。

该ls变量具有参数化类型List<String>，当List引用l赋值给ls时，编译器会生成未经检查的警告;如果编译器无法在编译时确定，而且JVM也无法在运行时确定l它不是一个List<String>类型;因此，产生堆污染。

因此，在编译时，编译器会在add语句处生成另一个未经检查的警告。编译器无法确定变量l是List<String>类型还是List<Integer>类型（以及另一种发生堆污染的情况）。然而，编译器不会在get语句中产生警告或错误。此声明有效; 它调用List<String>的get方法来获得一个String对象。相反，在运行时，get语句会抛出ClassCastException。

详细地说，当List<Number>对象l被赋值给另一个List<String>对象ls时，就会出现堆污染情况，它有一个不同的静态类型。然而，编译器仍然允许这个赋值。它必须允许这个赋值来保证与不支持泛型的JDK版本兼容性。因为类型擦除，List<Number>和List<String>变成List。因此，编译器允许对象l的赋值个体ls对象，因为ls是一个List类型。

另外，l.add调用该方法时会发生堆污染情况。该方法add第二参数应该是String，但实际参数是Integer。但是，编译器仍然允许这种方法调用。因为类型擦除，add方法第二个参数（List<E>.add(int,E)）变成Object。因此，编译器允许这种方法调用。因为在类型擦除之后，该l.add方法可以添加任何类型，包括Integer类型对象，因为它Object的子类。
#### 可变参数方法和非具体化形式参数
思考下面的ArrayBuilder.addToList方法。它将类型为T的elements可变参数，添加到List listArg参数中：
```java
import java.util.*;

public class ArrayBuilder{

    public static <T> void addToList (List<T> listArg, T... elements) {
        for (T x : elements) {
            listArg.add(x);
        }
    }

    public static void faultyMethod(List<String>... l) {
        Object[] objectArray = l;  // 有效
        objectArray[0] = Arrays.asList(new Integer(42));
        String s = l[0].get(0);    // 抛出ClassCastException异常
    }

}
```
```java
import java.util.*;

public class HeapPollutionExample {

    public static void main(String[] args) {

        List<String> stringListA = new ArrayList<String>();
        List<String> stringListB = new ArrayList<String>();

        ArrayBuilder.addToList(stringListA, "Seven", "Eight", "Nine");
        ArrayBuilder.addToList(stringListA, "Ten", "Eleven", "Twelve");
        List<List<String>> listOfStringLists = new ArrayList<List<String>>();
        ArrayBuilder.addToList(listOfStringLists, stringListA, stringListB);

        ArrayBuilder.faultyMethod(Arrays.asList("Hello!"), Arrays.asList("World!"));
    }
}
```
JDK7编译器为该方法的定义生成以下警告ArrayBuilder.addToList：
```
warning: [varargs] Possible heap pollution from parameterized vararg type T
```
当编译器遇到可变参数方法时，它将可变参数形式参数转换为数组。但是，Java编程语言不允许创建参数化类型的数组。在ArrayBuilder.addToList方法中，编译器将可变参数形式参数T... elements转换为形式参数T[] elements。即数组。但是，由于类型擦除，编译器将可变参数形式参数转换为Object[] 。因此，可能存在堆污染

下面是ArrayBuilder.addToList方法的反编译结果：
```java
public static <T> void addToList(List<T> listArg, T... elements) {
    Object[] var2 = elements;
    int var3 = elements.length;

    for(int var4 = 0; var4 < var3; ++var4) {
        T x = var2[var4];
        listArg.add(x);
    }

}
```
注意：JDK5和6编译器在ArrayBuilder.addToList调用时会生成警告HeapPollutionExample。这些编译器不会在声明上生成警告。但是，JDK7会在声明和调用上生成警告（除非这些警告被注解取代）。


#### 可变参数方法在传递非具体参数的缺点
该方法ArrayBuilder.faultyMethod显示了编译器为什么会警告您这些类型的方法。该方法的第一条语句将可变参数形式参数l赋值给objectArgs的Object数组：
```
Object[] objectArray = l;
```
这里可能产生堆污染。可变参数l可以赋值给objectArray数组。但是，编译器不会在此语句中生成未经检查的警告。当编译器编译List<String>... l 为List[] l 时已经生成警告。List[]是Object[]类型的子类，所以该声明是有效的。

因此，如果您将List任何类型的对象赋值给数组objectArray，则编译器不会发出警告或错误，如此语句所示：
```
objectArray[0] = Arrays.asList(new Integer(42));
```
这个语句将List<Integer>赋值给objectArray数组。

如果你调用以下方法：
```
ArrayBuilder.faultyMethod(Arrays.asList("Hello!"), Arrays.asList("World!"));
```
调用以下方法，在运行时JVM会抛出ClassCastException异常。
```
String s = l[0].get(0);    // 抛出ClassCastException 
```
变量l存储的是List<Integer>类型，但确期待是一个List<String>类型。

#### 抑制可变参数方法在传递不具体形式参数的警告
如果声明了一个具体参数化的可变参数方法。且方法的体不会抛出ClassCastException或其他类似的异常。您可以使用以下几种方式来抑制编译器为这些可变参数方法生成的警告：

- 将@SafeVarargs注解添加到静态非构造函数方法声明中：
- 将@SuppressWarnings({"unchecked", "varargs"})注解添加到方法声明中：
- 使用编译器选项-Xlint:varargs

例如，以下ArrayBuilder类的有两个方法addToList2和addToList3：

```java
import java.util.*;

public class ArrayBuilder {

    public static <T> void addToList (List<T> listArg, T... elements) {
        for (T x : elements) {
            listArg.add(x);
        }
    }

    @SuppressWarnings({"unchecked", "varargs"})
    public static <T> void addToList2 (List<T> listArg, T... elements) {
        for (T x : elements) {
            listArg.add(x);
        }
    }

    @SafeVarargs
    public static <T> void addToList3 (List<T> listArg, T... elements) {
        for (T x : elements) {
            listArg.add(x);
        }
    }

    // ...

}
```
```java
import java.util.*;

public class HeapPollutionExample {

    // ...

    public static void main(String[] args) {

        // ...

        ArrayBuilder.addToList(listOfStringLists, stringListA, stringListB);
        ArrayBuilder.addToList2(listOfStringLists, stringListA, stringListB);
        ArrayBuilder.addToList3(listOfStringLists, stringListA, stringListB);

        // ...

    }
}
```
Java编译器为此示例生成以下警告：

- addToList：　　
  - 在该方法的声明中：[unchecked] Possible heap pollution from parameterized vararg type T
  - 当该方法被调用时：[unchecked] unchecked generic array creation for varargs parameter of type List<String>[]
- addToList2：调用方法时（在方法声明中不生成警告）： [unchecked] unchecked generic array creation for varargs parameter of type List<String>[]
- addToList3：方法声明或调用方法时不会生成警告。
### 增强 Swing

Swing 的增强主要包含如下几点：

- JLayer 类： 是Swing组件的通用装饰器，您能够绘制组件并响应组件事件，而无需直接修改底层组件。
- Nimbus Look&Feel ： Nimbus Look＆Feel（L＆F）已经从com.sun.java.swing标准的API命名空间转移 到了javax.swing。查看javax.swing.plaf.nimbus包有更详细的介绍，虽然它不是默认的L＆F，但您可以轻松使用它。
- 重量级和轻量级组件： 历史上，在同一个容器中混合重量级（AWT）和轻量级（Swing）组件是有问题的。但是，混合重量级和轻量级组件很容易在Java SE 7中完成。
- 窗口的形状和透明度： Java SE 7版本支持具有透明度和非矩形形状的窗口。
- JColorChooser 类HSL选择： JColorChooser该类已添加HSL选项卡，允许用户使用色相饱和度亮度（HSL）颜色模型选择颜色。

自定义WallpaperLayerUI装饰类继承LayerUI重写paint方法
```java
import javax.swing.*;
import javax.swing.plaf.LayerUI;
import java.awt.*;

//继承 LayerUI 实现自定义装饰内容。
public class WallpaperLayerUI extends LayerUI<JComponent> {

    @Override
    public void paint(Graphics g, JComponent c) {
        super.paint(g, c);
        Graphics2D g2 = (Graphics2D) g.create();
        int w = c.getWidth();
        int h = c.getHeight();
        g2.setComposite(AlphaComposite.getInstance(
                AlphaComposite.SRC_OVER, .5f));
        //透明颜色渐变
        g2.setPaint(new GradientPaint(0, 0, Color.yellow, 0, h, Color.red));
        g2.fillRect(0, 0, w, h);
        g2.dispose();
    }
}
```

```java
import javax.swing.*;
import javax.swing.plaf.LayerUI;

public class Test {
    public static void main(String[] args) {
        javax.swing.SwingUtilities.invokeLater(new Runnable() {
            public void run() {
                createUI();
            }
        });
    }

    public static void createUI() {
        //创建JFrame
        JFrame f = new JFrame("Wallpaper");
        JPanel panel = createPanel();

//        这就是使用JDK7新特性：JLayer 类的装饰功能。如果不使用装饰 可以将 直接 f.add(panel)
        LayerUI<JComponent> layerUI = new WallpaperLayerUI();
        JLayer<JComponent> jlayer = new JLayer<JComponent>(panel, layerUI);
        f.add (jlayer);
//        f.add(panel);

        //设置 JFrame 宽高 默认关闭 相对位置 和 可见。
        f.setSize(300, 200);
        f.setDefaultCloseOperation (JFrame.EXIT_ON_CLOSE);
        f.setLocationRelativeTo (null);
        f.setVisible (true);
    }

    private static JPanel createPanel() {
        //创建面板
        JPanel p = new JPanel();

        //创建一个按钮组，并添加 Beef、Chicken 和 Vegetable 三个单选按钮
        ButtonGroup entreeGroup = new ButtonGroup();
        JRadioButton radioButton;
        p.add(radioButton = new JRadioButton("Beef", true));
        entreeGroup.add(radioButton);
        p.add(radioButton = new JRadioButton("Chicken"));
        entreeGroup.add(radioButton);
        p.add(radioButton = new JRadioButton("Vegetable"));
        entreeGroup.add(radioButton);

        //添加 Ketchup 、 Mustard 和 Pickles 复选框
        p.add(new JCheckBox("Ketchup"));
        p.add(new JCheckBox("Mustard"));
        p.add(new JCheckBox("Pickles"));

        //添加 Special requests: 标签 和 文本框
        p.add(new JLabel("Special requests:"));
        p.add(new JTextField(20));

        //添加 Place Order 按钮
        JButton orderButton = new JButton("Place Order");
        p.add(orderButton);

        return p;
    }
}
```

## JDK1.8
https://www.oracle.com/technetwork/cn/java/javase/8-whats-new-2157071-zhs.html

*   [Java 编程语言](http://docs.oracle.com/javase/8/docs/technotes/guides/language/enhancements.html#javase8)
    
    *   Lambda 表达式是一个新的语言特性，已经在此版本中引入。该特性让您可以将功能视为方法参数，或者将代码视为数据。使用 Lambda 表达式，您可以更简洁地表示单方法接口（称为功能接口）的实例。
        
    *   方法引用为已经具有名称的方法提供了易于理解的 lambda 表达式。
        
    *   默认方法允许将新功能添加到库的接口中，并确保与为这些接口的旧版本编写的代码的二进制兼容性。
        
    *   重复批注支持对同一个声明或类型的使用多次应用相同的批注类型。
        
    *   类型批注支持在使用类型的任何地方应用批注，而不仅限于声明。与可插拔类型系统结合使用时，此特性可改进代码的类型检查。
        
    *   改进类型推断。
        
    *   方法参数反射。
        
*   [集合](http://docs.oracle.com/javase/8/docs/technotes/guides/collections/changes8.html)
    
    *   新的 `java.util.stream` 包中的类提供了一个 Stream API，支持对元素流进行函数式操作。Stream API 集成在 Collections API 中，可以对集合进行批量操作，例如顺序或并行的 map-reduce 转换。
        
    *   针对存在键冲突的 HashMap 的性能改进
        
*   [紧凑 profile](http://docs.oracle.com/javase/8/docs/technotes/guides/compactprofiles/)包含 Java SE 平台的预定义子集，并且支持不需要在小型设备上部署和运行整个平台的应用。
    
*   [安全性](http://docs.oracle.com/javase/8/docs/technotes/guides/security/enhancements-8.html)
    
    *   默认启用客户端 TLS 1.2
        
    *   `AccessController.doPrivileged` 的新变体支持代码断言其权限的子集，而不会阻止完全遍历堆栈来检查其他权限
        
    *   更强大的基于密码的加密算法
        
    *   JSSE 服务器端支持 SSL/TLS 服务器名称指示 (SNI) 扩展
        
    *   支持 AEAD 算法：SunJCE 提供程序得到了增强，支持 AES/GCM/NoPadding 密码实现以及 GCM 算法参数。而且 SunJSSE 提供程序也得到了增强，支持基于 AEAD 模式的密码套件。请参阅 Oracle 提供程序文档，JEP 115。
        
    *   密钥库增强，包括新的域密钥库类型 `java.security.DomainLoadStoreParameter` 和为 keytool 实用程序新增的命令选项 `-importpassword`
        
    *   SHA-224 消息摘要
        
    *   增强了对 NSA Suite B 加密的支持
        
    *   更好地支持高熵随机数生成
        
    *   新增了 `java.security.cert.PKIXRevocationChecker` 类，用于配置 X.509 证书的撤销检查
        
    *   适用于 Windows 的 64 位 PKCS11
        
    *   Kerberos 5 重放缓存中新增了 rcache 类型
        
    *   支持 Kerberos 5 协议转换和受限委派
        
    *   默认禁用 Kerberos 5 弱加密类型
        
    *   适用于 GSS-API/Kerberos 5 机制的未绑定 SASL
        
    *   针对多个主机名称的 SASL 服务
        
    *   JNI 桥接至 Mac OS X 上的原生 JGSS
        
    *   SunJSSE 提供程序中支持更强大的临时 DH 密钥
        
    *   JSSE 中支持服务器端加密套件首选项自定义
        
*   [JavaFX](http://docs.oracle.com/javase/8/javase-clienttechnologies.htm)
    
    *   本版本中实施了新的 Modena 主题。有关更多信息，请参阅 [xexperience.com](http://fxexperience.com/2013/03/modena-theme-update/) 上的博客。
        
    *   新的 `SwingNode` 类允许开发人员将 Swing 内容嵌入到 JavaFX 应用中。请参阅 [`SwingNode`](http://docs.oracle.com/javase/8/javafx/api/javafx/embed/swing/SwingNode.html) javadoc 和[将 Swing 内容嵌入 JavaFX 应用中](http://docs.oracle.com/javase/8/javafx/interoperability-tutorial/embed-swing.htm)。
        
    *   新的 UI 控件包括 [`DatePicker`](http://docs.oracle.com/javase/8/javafx/api/javafx/scene/control/DatePicker.html) 和 [`TreeTableView`](http://docs.oracle.com/javase/8/javafx/api/javafx/scene/control/TreeTableView.html) 控件。
        
    *   `javafx.print` 程序包为 JavaFX Printing API 提供了公共类。有关更多信息，请参阅 [javadoc](http://docs.oracle.com/javase/8/javafx/api/javafx/print/package-summary.html)。
        
    *   3D 图形特性现在包括 3D 形状、摄像头、灯光、子场景、材料、挑选和抗锯齿。JavaFX 3D 图形库中新增了 `Shape3D`（`Box`、`Cylinder`、`MeshView` 和 `Sphere` 子类）、`SubScene`、`Material`、`PickResult`、`LightBase`（`AmbientLight` 和 `PointLight` 子类）和 `SceneAntialiasing` API 类。此版本中的 `Camera` API 类也已更新。请参阅 `javafx.scene.shape.Shape3D`、`javafx.scene.SubScene`、`javafx.scene.paint.Material`、`javafx.scene.input.PickResult` 和 `javafx.scene.SceneAntialiasing` 类的相关 javadoc 以及 [JavaFX 3D 图形入门](http://docs.oracle.com/javase/8/javafx/graphics-tutorial/javafx-3d-graphics.htm)文档。
        
    *   `WebView` 类包含新特性和改进。有关其他 HTML5 特性（包括 Web 套接字、Web 辅助进程和 Web 字体）的更多信息，请参阅 [HTML5 支持的特性](http://docs.oracle.com/javase/8/javafx/embedded-browser-tutorial/index.html)。
        
    *   增强了文本支持，包括双向文本、复杂文本脚本（如泰语和印地语控件）以及文本节点中的多行多样式文本。
        
    *   此版本添加了对 Hi-DPI 显示的支持。
        
    *   CSS Styleable\* 类已成为公共 API。有关更多信息，请参阅 [`Javafx.css`](http://docs.oracle.com/javase/8/javafx/api/javafx/css/package-frame.html) javadoc。
        
    *   新的 [`ScheduledService`](http://docs.oracle.com/javase/8/javafx/api/javafx/concurrent/ScheduledService.html) 类允许自动重新启动服务。
        
    *   JavaFX 现在可用于 ARM 平台。适用于 ARM 的 JDK 包含 JavaFX 的基础组件、图形组件和控制组件。
        
*   [工具](http://docs.oracle.com/javase/8/docs/technotes/tools/enhancements-8.html)
    
    *   可通过 `jjs` 命令来调用 Nashorn 引擎。
        
    *   `java` 命令用于启动 JavaFX 应用。
        
    *   重新编写了 `java` 手册页。
        
    *   可通过 `jdeps` 命令行工具来分析类文件。
        
    *   Java Management Extensions (JMX) 支持远程访问诊断命令。
        
    *   `jarsigner` 工具提供了一个选项用于请求获取时间戳机构 (TSA) 的签名时间戳。
        
    *   [Javac 工具](http://docs.oracle.com/javase/8/docs/technotes/guides/javac/index.html)
        
        *   `javac` 命令的 `-parameters` 选项可用于存储正式参数名称，并启用反射 API 来检索正式参数名称。
            
        *   `javac` 命令现已正确实施了 Java 语言规范 (JLS) 第 15.21 节中的相等运算符的类型规则。
            
        *   `javac` 工具现在支持检查 `javadoc` 注释的内容，从而避免在运行 `javadoc` 时生成的文件中产生各种问题，例如无效的 HTML 或可访问性问题。可通过新的 `Xdoclint` 选项来启用此特性。有关更多详细信息，请参阅运行“`javac-X`”时的输出。此特性也可以在 `javadoc` 工具中使用，并且默认启用。
            
        *   `javac` 工具现在支持根据需要生成原生标头。这样便无需在构建管道中单独运行 `javah` 工具。可以使用新的 `-h` 选项在 `javac` 中启用此特性，该选项用于指定写入头文件的目录。将为任何具有原生方法或者使用 `java.lang.annotation.Native` 类型的新批注的类进行批注的常量字段生成头文件。
            
    *   [Javadoc 工具](http://docs.oracle.com/javase/8/docs/technotes/guides/javadoc/whatsnew-8.html)
        
        *   `javadoc` 工具支持新的 `DocTree` API，让您可以将 Javadoc 注释作为抽象语法树来进行遍历。
            
        *   `javadoc` 工具支持新的 Javadoc Access API，让您可以直接从 Java 应用中调用 Javadoc 工具，而无需执行新的进程。有关更多信息，请参阅 [javadoc 新特性](http://docs.oracle.com/javase/8/docs/technotes/guides/javadoc/whatsnew-8.html)页面。
            
        *   `javadoc` 工具现在支持检查 `javadoc` 注释的内容，从而避免在运行 `javadoc` 时生成的文件中产生各种问题，例如无效的 HTML 或可访问性问题。此特性默认为启用状态，可以通过新的 `-Xdoclint` 选项加以控制。有关更多详细信息，请参阅运行“`javadoc -X`”时的输出。`javac` 工具也支持此特性，但默认情况下并未启用它。
            
*   [国际化](http://docs.oracle.com/javase/8/docs/technotes/guides/intl/enhancements.8.html)
    
    *   Unicode 增强，包括对 Unicode 6.2.0 的支持
        
    *   采用 Unicode CLDR 数据和 java.locale.providers 系统属性
        
    *   新增日历和区域设置 API
        
    *   支持将自定义资源包作为扩展进行安装
        
*   [部署](http://docs.oracle.com/javase/8/docs/technotes/guides/jweb/enhancements-8.html)
    
    *   现在可以使用 `URLPermission` 允许沙盒小程序和 Java Web Start 应用连接回启动它们的服务器。不再授予 `SocketPermission`。
        
    *   在所有安全级别，主 JAR 文件的 JAR 文件清单中都需要 Permissions 属性。
        
*   [Date-Time 程序包](http://docs.oracle.com/javase/8/docs/technotes/guides/datetime/index.html) — 一组新程序包，提供全面的日期-时间模型。
    
*   [脚本编写](http://docs.oracle.com/javase/8/docs/technotes/guides/scripting/enhancements.html#jdk8)
    
    *   Rhino Javascript 引擎已被替换为 [Nashorn](http://docs.oracle.com/javase/8/docs/technotes/guides/scripting/nashorn/) JavaScript 引擎
        
*   [Pack200](http://docs.oracle.com/javase/8/docs/technotes/guides/pack200/enhancements.html)
    
    *   Pack200 支持 JSR 292 引入的常量池条目和新字节码
        
    *   JDK8 支持 JSR-292、JSR-308 和 JSR-335 指定的类文件更改
        
*   [IO 和 NIO](http://docs.oracle.com/javase/8/docs/technotes/guides/io/enhancements.html#jdk8)
    
    *   全新的基于 Solaris 事件端口机制的面向 Solaris 的 `SelectorProvider` 实现。要使用它，请将系统属性 `java.nio.channels.spi.Selector` 的值设置为 `sun.nio.ch.EventPortSelectorProvider`。
        
    *   减小 `<JDK_HOME>/jre/lib/charsets.jar` 文件的大小
        
    *   提高了 `java.lang.String(byte[], *)` 构造函数和 `java.lang.String.getBytes()` 方法的性能。
        
*   [java.lang 和 java.util 程序包](http://docs.oracle.com/javase/8/docs/technotes/guides/lang/enhancements.html#jdk8)
    
    *   并行数组排序
        
    *   标准编码和解码 Base64
        
    *   无符号算术支持
        
*   [JDBC](http://docs.oracle.com/javase/8/docs/technotes/guides/jdbc/)
    
    *   删除了 JDBC-ODBC Bridge。
        
    *   JDBC 4.2 引入了新特性。
        
*   Java DB
    
    *   JDK 8 包含 Java DB 10.10。
        
*   [网络](http://docs.oracle.com/javase/8/docs/technotes/guides/net/enhancements-8.0.html)
    
    *   已添加 `java.net.URLPermission` 类。
        
    *   在 `java.net.HttpURLConnection` 类中，如果安装了安全管理器，那么请求打开连接的调用需要权限。
        
*   [并发性](http://docs.oracle.com/javase/8/docs/technotes/guides/concurrency/changes8.html)
    
    *   `java.util.concurrent` 程序包中新增了一些类和接口。
        
    *   `java.util.concurrent.ConcurrentHashMap` 类中新增了一些方法，支持基于新增流工具和 lambda 表达式的聚合操作。
        
    *   `java.util.concurrent.atomic` 程序包中新增了一些类来支持可扩展、可更新的变量。
        
    *   `java.util.concurrent.ForkJoinPool` 类中新增了一些方法来支持通用池。
        
    *   新增的 `java.util.concurrent.locks.StampedLock` 类提供了一个基于能力的锁，可通过三种模式来控制读/写访问。
        
*   [Java XML](http://docs.oracle.com/javase/8/docs/technotes/guides/xml/enhancements.html) - [JAXP](http://docs.oracle.com/javase/8/docs/technotes/guides/xml/jaxp/enhancements-8.html)
    
*   [HotSpot](http://docs.oracle.com/javase/8/docs/technotes/guides/vm/)
    
    *   新增的硬件内部函数以便使用高级加密标准 (AES)。`UseAES` 和 `UseAESIntrinsics` 标志用于为 Intel 硬件启用基于硬件的 AES 内部函数。硬件必须是 2010 年或更新的 Westmere 硬件。例如，要启用硬件 AES，请使用以下标志：
        
        \-XX:+UseAES -XX:+UseAESIntrinsics
        
        要禁用硬件 AES，请使用以下标志：
        
        \-XX:-UseAES -XX:-UseAESIntrinsics
        
    *   删除了 PermGen。
        
    *   方法调用的字节码指令支持 Java 编程语言中的默认方法。
        
*   [Java Mission Control 5.3 版本说明](http://www.oracle.com/technetwork/java/javase/jmc53-release-notes-2157171.html)
    
    *   JDK 8 包含 Java Mission Control 5.3。



https://www.cnblogs.com/peter1018/p/9183548.html

### Lambda 表达式
Lambda 允许把函数作为一个方法的参数（函数作为参数传递到方法中）。
https://www.runoob.com/java/java8-lambda-expressions.html

### 方法引用 
方法引用提供了非常有用的语法，可以直接引用已有Java类或对象（实例）的方法或构造器。与lambda联合使用，方法引用可以使语言的构造更紧凑简洁，减少冗余代码。
https://www.runoob.com/java/java8-method-references.html

### 函数式接口
https://www.runoob.com/java/java8-functional-interfaces.html

### 默认方法
默认方法就是一个在接口里面有了一个实现的方法。
https://www.runoob.com/java/java8-default-methods.html

### 新工具
 新的编译工具，如：Nashorn引擎 jjs、 类依赖分析器jdeps。

### Stream API 
新添加的Stream API（java.util.stream） 把真正的函数式编程风格引入到Java中。
https://www.runoob.com/java/java8-streams.html

### Date Time API
加强对日期与时间的处理。
https://www.runoob.com/java/java8-datetime-api.html

### Optional 类
Optional 类已经成为 Java 8 类库的一部分，用来解决空指针异常。
https://www.runoob.com/java/java8-optional-class.html

### Nashorn, JavaScript 引擎
Java 8提供了一个新的Nashorn javascript引擎，它允许我们在JVM上运行特定的javascript应用。
https://www.runoob.com/java/java8-nashorn-javascript.html

### Base64
在Java 8中，Base64编码已经成为Java类库的标准。

Java 8 内置了 Base64 编码的编码器和解码器。

Base64工具类提供了一套静态方法获取下面三种BASE64编解码器：

- **基本**：输出被映射到一组字符A-Za-z0-9+/，编码不添加任何行标，输出的解码仅支持A-Za-z0-9+/。
- **URL**：输出映射到一组字符A-Za-z0-9+_，输出是URL和文件。
- **MIME**：输出隐射到MIME友好格式。输出每行不超过76字符，并且使用'\r'并跟随'\n'作为分割。编码输出最后没有行分割。

https://www.runoob.com/java/java8-base64.html

## JDK9 
https://docs.oracle.com/javase/9/whatsnew/toc.htm#JSNEW-GUID-C23AFD78-C777-460B-8ACE-58BE5EA681F6

> 102: [Process API Updates](https://openjdk.java.net/jeps/102)  
> 110: [HTTP 2 Client](https://openjdk.java.net/jeps/110)  
> 143: [Improve Contended Locking](https://openjdk.java.net/jeps/143)  
> 158: [Unified JVM Logging](https://openjdk.java.net/jeps/158)  
> 165: [Compiler Control](https://openjdk.java.net/jeps/165)  
> 193: [Variable Handles](https://openjdk.java.net/jeps/193)  
> 197: [Segmented Code Cache](https://openjdk.java.net/jeps/197)  
> 199: [Smart Java Compilation, Phase Two](https://openjdk.java.net/jeps/199)  
> 200: [The Modular JDK](https://openjdk.java.net/jeps/200)  
> 201: [Modular Source Code](https://openjdk.java.net/jeps/201)  
> 211: [Elide Deprecation Warnings on Import Statements](https://openjdk.java.net/jeps/211)  
> 212: [Resolve Lint and Doclint Warnings](https://openjdk.java.net/jeps/212)  
> 213: [Milling Project Coin](https://openjdk.java.net/jeps/213)  
> 214: [Remove GC Combinations Deprecated in JDK 8](https://openjdk.java.net/jeps/214)  
> 215: [Tiered Attribution for javac](https://openjdk.java.net/jeps/215)  
> 216: [Process Import Statements Correctly](https://openjdk.java.net/jeps/216)  
> 217: [Annotations Pipeline 2.0](https://openjdk.java.net/jeps/217)  
> 219: [Datagram Transport Layer Security (DTLS)](https://openjdk.java.net/jeps/219)  
> 220: [Modular Run-Time Images](https://openjdk.java.net/jeps/220)  
> 221: [Simplified Doclet API](https://openjdk.java.net/jeps/221)  
> 222: [jshell: The Java Shell (Read-Eval-Print Loop)](https://openjdk.java.net/jeps/222)  
> 223: [New Version-String Scheme](https://openjdk.java.net/jeps/223)  
> 224: [HTML5 Javadoc](https://openjdk.java.net/jeps/224)  
> 225: [Javadoc Search](https://openjdk.java.net/jeps/225)  
> 226: [UTF-8 Property Files](https://openjdk.java.net/jeps/226)  
> 227: [Unicode 7.0](https://openjdk.java.net/jeps/227)  
> 228: [Add More Diagnostic Commands](https://openjdk.java.net/jeps/228)  
> 229: [Create PKCS12 Keystores by Default](https://openjdk.java.net/jeps/229)  
> 231: [Remove Launch-Time JRE Version Selection](https://openjdk.java.net/jeps/231)  
> 232: [Improve Secure Application Performance](https://openjdk.java.net/jeps/232)  
> 233: [Generate Run-Time Compiler Tests Automatically](https://openjdk.java.net/jeps/233)  
> 235: [Test Class-File Attributes Generated by javac](https://openjdk.java.net/jeps/235)  
> 236: [Parser API for Nashorn](https://openjdk.java.net/jeps/236)  
> 237: [Linux/AArch64 Port](https://openjdk.java.net/jeps/237)  
> 238: [Multi-Release JAR Files](https://openjdk.java.net/jeps/238)  
> 240: [Remove the JVM TI hprof Agent](https://openjdk.java.net/jeps/240)  
> 241: [Remove the jhat Tool](https://openjdk.java.net/jeps/241)  
> 243: [Java-Level JVM Compiler Interface](https://openjdk.java.net/jeps/243)  
> 244: [TLS Application-Layer Protocol Negotiation Extension](https://openjdk.java.net/jeps/244)  
> 245: [Validate JVM Command-Line Flag Arguments](https://openjdk.java.net/jeps/245)  
> 246: [Leverage CPU Instructions for GHASH and RSA](https://openjdk.java.net/jeps/246)  
> 247: [Compile for Older Platform Versions](https://openjdk.java.net/jeps/247)  
> 248: [Make G1 the Default Garbage Collector](https://openjdk.java.net/jeps/248)  
> 249: [OCSP Stapling for TLS](https://openjdk.java.net/jeps/249)  
> 250: [Store Interned Strings in CDS Archives](https://openjdk.java.net/jeps/250)  
> 251: [Multi-Resolution Images](https://openjdk.java.net/jeps/251)  
> 252: [Use CLDR Locale Data by Default](https://openjdk.java.net/jeps/252)  
> 253: [Prepare JavaFX UI Controls & CSS APIs for Modularization](https://openjdk.java.net/jeps/253)  
> 254: [Compact Strings](https://openjdk.java.net/jeps/254)  
> 255: [Merge Selected Xerces 2.11.0 Updates into JAXP](https://openjdk.java.net/jeps/255)  
> 256: [BeanInfo Annotations](https://openjdk.java.net/jeps/256)  
> 257: [Update JavaFX/Media to Newer Version of GStreamer](https://openjdk.java.net/jeps/257)  
> 258: [HarfBuzz Font-Layout Engine](https://openjdk.java.net/jeps/258)  
> 259: [Stack-Walking API](https://openjdk.java.net/jeps/259)  
> 260: [Encapsulate Most Internal APIs](https://openjdk.java.net/jeps/260)  
> 261: [Module System](https://openjdk.java.net/jeps/261)  
> 262: [TIFF Image I/O](https://openjdk.java.net/jeps/262)  
> 263: [HiDPI Graphics on Windows and Linux](https://openjdk.java.net/jeps/263)  
> 264: [Platform Logging API and Service](https://openjdk.java.net/jeps/264)  
> 265: [Marlin Graphics Renderer](https://openjdk.java.net/jeps/265)  
> 266: [More Concurrency Updates](https://openjdk.java.net/jeps/266)  
> 267: [Unicode 8.0](https://openjdk.java.net/jeps/267)  
> 268: [XML Catalogs](https://openjdk.java.net/jeps/268)  
> 269: [Convenience Factory Methods for Collections](https://openjdk.java.net/jeps/269)  
> 270: [Reserved Stack Areas for Critical Sections](https://openjdk.java.net/jeps/270)  
> 271: [Unified GC Logging](https://openjdk.java.net/jeps/271)  
> 272: [Platform-Specific Desktop Features](https://openjdk.java.net/jeps/272)  
> 273: [DRBG-Based SecureRandom Implementations](https://openjdk.java.net/jeps/273)  
> 274: [Enhanced Method Handles](https://openjdk.java.net/jeps/274)  
> 275: [Modular Java Application Packaging](https://openjdk.java.net/jeps/275)  
> 276: [Dynamic Linking of Language-Defined Object Models](https://openjdk.java.net/jeps/276)  
> 277: [Enhanced Deprecation](https://openjdk.java.net/jeps/277)  
> 278: [Additional Tests for Humongous Objects in G1](https://openjdk.java.net/jeps/278)  
> 279: [Improve Test-Failure Troubleshooting](https://openjdk.java.net/jeps/279)  
> 280: [Indify String Concatenation](https://openjdk.java.net/jeps/280)  
> 281: [HotSpot C++ Unit-Test Framework](https://openjdk.java.net/jeps/281)  
> 282: [jlink: The Java Linker](https://openjdk.java.net/jeps/282)  
> 283: [Enable GTK 3 on Linux](https://openjdk.java.net/jeps/283)  
> 284: [New HotSpot Build System](https://openjdk.java.net/jeps/284)  
> 285: [Spin-Wait Hints](https://openjdk.java.net/jeps/285)  
> 287: [SHA-3 Hash Algorithms](https://openjdk.java.net/jeps/287)  
> 288: [Disable SHA-1 Certificates](https://openjdk.java.net/jeps/288)  
> 289: [Deprecate the Applet API](https://openjdk.java.net/jeps/289)  
> 290: [Filter Incoming Serialization Data](https://openjdk.java.net/jeps/290)  
> 291: [Deprecate the Concurrent Mark Sweep (CMS) Garbage Collector](https://openjdk.java.net/jeps/291)  
> 292: [Implement Selected ECMAScript 6 Features in Nashorn](https://openjdk.java.net/jeps/292)  
> 294: [Linux/s390x Port](https://openjdk.java.net/jeps/294)  
> 295: [Ahead-of-Time Compilation](https://openjdk.java.net/jeps/295)  
> 297: [Unified arm32/arm64 Port](https://openjdk.java.net/jeps/297)  
> 298: [Remove Demos and Samples](https://openjdk.java.net/jeps/298)  
> 299: [Reorganize Documentation](https://openjdk.java.net/jeps/299)




### 模块系统：
模块是一个包的容器，Java 9 最大的变化之一是引入了模块系统（Jigsaw 项目）。
https://www.runoob.com/java/java9-module-system.html
https://zhuanlan.zhihu.com/p/31104953

### REPL (JShell)：
交互式编程环境。
https://www.runoob.com/java/java9-repl.html

### HTTP 2 客户端：
HTTP/2标准是HTTP协议的最新版本，新的 HTTPClient API 支持 WebSocket 和 HTTP2 流以及服务器推送特性。

### 改进的 Javadoc：
Javadoc 现在支持在 API 文档中的进行搜索。另外，Javadoc 的输出现在符合兼容 HTML5 标准。支持搜索功能.
JDK8 生成的java帮助文档是在 HTML4 中。而HTML4 已经是很久的标准了。
JDK9 的javadoc，现支持HTML5 标准。
https://www.runoob.com/java/java9-improved-javadocs.html

### 多版本兼容 JAR 包：
多版本兼容 JAR 功能能让你创建仅在特定版本的 Java 环境中运行库程序时选择使用的 class 版本。
https://www.runoob.com/java/java9-multirelease-jar.html

### 集合工厂方法：
List，Set 和 Map 接口中，新的静态工厂方法可以创建这些集合的不可变实例。
https://www.runoob.com/java/java9-collection-factory-methods.html


### 私有接口方法：
在接口中使用private私有方法。我们可以使用 private 访问修饰符在接口中编写私有方法。
https://www.runoob.com/java/java9-private-interface-methods.html

### 进程 API: 
改进的 API 来控制和管理操作系统进程。引进 java.lang.ProcessHandle 及其嵌套接口 Info 来让开发者逃离时常因为要获取一个本地进程的 PID 而不得不使用本地代码的窘境。
https://www.runoob.com/java/java9-process-api-improvements.html

### 改进的 Stream API：
改进的 Stream API 添加了一些便利的方法，使流处理更容易，并使用收集器编写复杂的查询。
https://www.runoob.com/java/java9-stream-api-improvements.html

### 改进 try-with-resources：
如果你已经有一个资源是 final 或等效于 final 变量,您可以在 try-with-resources 语句中使用该变量，而无需在 try-with-resources 语句中声明一个新变量。
https://www.runoob.com/java/java9-try-with-resources-improvement.html

### 改进的弃用注解 @Deprecated：
注解 @Deprecated 可以标记 Java API 状态，可以表示被标记的 API 将会被移除，或者已经破坏。
https://www.runoob.com/java/java9-enhanced-deprecated-annotation.html

### 改进钻石操作符(Diamond Operator) ：
匿名类可以使用钻石操作符(Diamond Operator)。
https://www.runoob.com/java/java9-inner-class-diamond-operator.html

### 改进 Optional 类：
java.util.Optional 添加了很多新的有用方法，Optional 可以直接转为 stream。
https://www.runoob.com/java/java9-optional-class-improvements.html

### 多分辨率图像 API：
定义多分辨率图像API，开发者可以很容易的操作和展示不同分辨率的图像了。
https://www.runoob.com/java/java9-multiresolution-image_api.html

### 改进的 CompletableFuture API ： 
CompletableFuture 类的异步机制可以在 ProcessHandle.onExit 方法退出时执行操作。
https://www.runoob.com/java/java9-completablefuture-api-improvements.html

### 轻量级的 JSON API：
内置了一个轻量级的JSON API
### 响应式流（Reactive Streams) API: 
Java 9中引入了新的响应式流 API 来支持 Java 9 中的响应式编程。


### java 动态编译器 AOT
JIT（Just-in-time）编译器可以在运行时将热点编译成本地代码，速度很快。但是 Java 项目现在变得很大很复杂，因此 JIT 编译器需要花费较长时间才能热身完，而且有些 Java 方法还没法编译，性能方面也会下降。AoT 编译就是为了解决这些问题而生的。
在 JDK 9 中， AOT（JEP 295: Ahead-of-Time Compilation）作为实验特性被引入进来，开发者可以利用新的 jaotc 工具将重点代码转换成类似类库一样的文件。虽然仍处于试验阶段，但这个功能使得 Java应用在被虚拟机启动之前能够先将 Java 类编译为原生代码。此功能旨在改进小型和大型应用程序的启动时间，同时对峰值性能的影响很小。
但是 Java 技术供应商 Excelsior 的营销总监 Dmitry Leskov 担心 AoT 编译技术不够成熟，希望 Oracle 能够等到 Java 10时有个更稳定版本才发布。

### 统一的 JVM 日志系统
日志是解决问题的唯一有效途径：曾经很难知道导致 JVM 性能问题和导致 JVM 崩溃的根本原因。不同的 JVM 日志的碎片化和日志选项（例如：JVM 组件对于日志使用的是不同的机制和规则），这使得 JVM 难以进行调试。
解决该问题最佳方法：对所有的 JVM 组件引入一个单一的系统，这些 JVM 组件支持细粒度的和易配置的 JVM 日志。
### 参考
https://www.cnblogs.com/peter1018/p/9209951.html

## JDK10
http://cr.openjdk.java.net/~iris/se/10/latestSpec/

> 286: [Local-Variable Type Inference](https://openjdk.java.net/jeps/286)  
> 296: [Consolidate the JDK Forest into a Single Repository](https://openjdk.java.net/jeps/296)  
> 304: [Garbage-Collector Interface](https://openjdk.java.net/jeps/304)  
> 307: [Parallel Full GC for G1](https://openjdk.java.net/jeps/307)  
> 310: [Application Class-Data Sharing](https://openjdk.java.net/jeps/310)  
> 312: [Thread-Local Handshakes](https://openjdk.java.net/jeps/312)  
> 313: [Remove the Native-Header Generation Tool (javah)](https://openjdk.java.net/jeps/313)  
> 314: [Additional Unicode Language-Tag Extensions](https://openjdk.java.net/jeps/314)  
> 316: [Heap Allocation on Alternative Memory Devices](https://openjdk.java.net/jeps/316)  
> 317: [Experimental Java-Based JIT Compiler](https://openjdk.java.net/jeps/317)  
> 319: [Root Certificates](https://openjdk.java.net/jeps/319)  
> 322: [Time-Based Release Versioning](https://openjdk.java.net/jeps/322)

### 局部变量类型推断
很多人抱怨Java是一种强类型，需要引入大量的样板代码。甚至在这些情况下，给定好变量名，通常很清楚发生了什么，明显类型声明往往被认为是不必要的。许多流行的编程语言都已经支持某种形式的局部变量类型推断：如C++ (auto), C# (var), Scala (var/val), Go (declaration with :=)等。

JDK10 可以使用var作为局部变量类型推断标识符，此符号仅适用于局部变量，增强for循环的索引，以及传统for循环的本地变量；它不能使用于方法形式参数，构造函数形式参数，方法返回类型，字段，catch形式参数或任何其他类型的变量声明。

标识符var不是关键字；相反，它是一个保留的类型名称。这意味着var用作变量，方法名或则包名称的代码不会受到影响；但var不能作为类或则接口的名字（但这样命名是比较罕见的，因为他违反了通常的命名约定，类和接口首字母应该大写）。

参考一下示例：
```java
var str = "ABC"; //根据推断为 字符串类型
var l = 10L;//根据10L 推断long 类型
var flag = true;//根据 true推断 boolean 类型
var flag1 = 1;//这里会推断boolean类型。0表示false 非0表示true
var list = new ArrayList<String>();  // 推断 ArrayList<String>
var stream = list.stream();          // 推断 Stream<String>
```
反编译class文件：
```java
String str = "ABC";
long l = 10L;
boolean flag = true;
int flag1 = true;
ArrayList<String> list = new ArrayList();
Stream<String> stream = list.stream();
```
从上面示例可以看出，当我们是用复杂的方法时，不需要特意去指定他的具体类型返回，可以使用var推断出正确的数据类型，这在编码中，可以大幅减少我们对方法返回值的探究。
### 将JDK多存储库合并为单存储库
为了简化和简化开发，将JDK多存储库合并到一个存储库中。多年来，JDK的完整代码已经被分解成多个存储库。在JDK9 中有八个仓库：root、corba、hotspot、jaxp、jaxws、jdk、langtools和nashorn。在JDK10中被合并为一个存储库。

虽然这种多存储库模型具有一些有点，但它也有许多缺点，并且在支持各种可取的源代码管理操作方面做得很差。特别是，不可能在相互依赖的变更存储库之间执行原子提交。例如，如果一个bug修复或RFE的代码现在同时跨越了jdk和hotspot 存储库，那么对于两个存储库来说，在托管这两个不同的存储库中，对两个存储库的更改是不可能实现的。跨多个存储库的变更是常见。
### 垃圾回收接口
这不是让开发者用来控制垃圾回收的接口；而是一个在 JVM 源代码中的允许另外的垃圾回收器快速方便的集成的接口。

垃圾回收接口为HotSpot的GC代码提供更好的模块化；在不影响当前代码的基础情况下，将GC添加到HotSpot变的更简单；更容易从JDK构建中排除GC。实际添加或删除GC不是目标，这项工作将使HotSpot中GC算法的构建时间隔离取得进展，但它不是完全构建时间隔离的目标。


### 并行Full GC 的G1
JDK10 通过并行Full GC，改善G1的延迟。G1垃圾收集器在JDK 9中是默认的。以前的默认值并行收集器中有一个并行的Full GC。为了尽量减少对使用GC用户的影响，G1的Full GC也应该并行。

G1垃圾收集器的设计目的是避免Full收集，但是当集合不能足够快地回收内存时，就会出现完全GC。目前对G1的Full GC的实现使用了单线程标记-清除-压缩算法。JDK10 使用并行化标记-清除-压缩算法，并使用Young和Mixed收集器相同的线程数量。线程的数量可以由-XX:ParallelGCThreads选项来控制，但是这也会影响用Young和Mixed收集器的线程数量。

### 应用数据共享
为了提高启动和内存占用，扩展现有的类数据共享（CDS）特性，允许将应用程序类放置在共享档案中。

- 通过在不同的Java进程间共享公共类元数据来减少占用空间。
- 提升启动时间。
- CDS允许将来自JDK的运行时映像文件（$JAVA_HOME/lib/modules）的归档类和应用程序类路径加载到内置平台和系统类加载器中。
- CDS允许将归档类加载到自定义类加载器中。

### 线程局部管控
在不执行全局VM安全点的情况下对线程执行回调的方法。让它停止单个线程而不是全部线程。

### 移除Native-Header Generation Tool （javah）
JDK10 从JDK中移除了javah 工具。该工具已被JDK8 （[JDK-7150368](https://bugs.openjdk.java.net/browse/JDK-7150368)）中添加javac高级功能所取代。此功能提供了在编译java源代码时编写本机头文件的功能，从而无需使用单独的工具。

### Unicode 标签扩展
JDK10 改善 java.util.Locale 类和相关的 API 以实现额外 [BCP 47](http://www.rfc-editor.org/rfc/bcp/bcp47.txt) 语言标签的 Unicode 扩展。尤其以下扩展支持：

- cu：货币类型
- fw：一周的第一天
- rg：区域覆盖
- tz：时区

为支持以上扩展，JDK10对以下API进行更改：

- java.text.DateFormat::get*Instance：将根据扩展ca、rg或tz返回实例。
- java.text.DateFormatSymbols::getInstance：将根据扩展rg返回实例。
- java.text.DecimalFormatSymbols::getInstance：将根据扩展rg返回实例。
- java.text.NumberFormat::get*Instance：将根据nu或rg返回实例。
- java.time.format.DateTimeFormatter::localizedBy：将返回DateTimeFormatter 根据ca，rg或rz的实例。
- java.time.format.DateTimeFormatterBuilder::getLocalizedDateTimePattern：将根据rg返回String。
- java.time.format.DecimalStyle::of：将返回DecimalStyle根据nu或rg的实例。
- java.time.temporal.WeekFields::of：将返回WeekFields根据fw或rg的实例。
- java.util.Calendar::{getFirstDayOfWeek,getMinimalDaysInWeek}：将根据fw或rg返回值。
- java.util.Currency::getInstance：将返回Currency根据cu或rg返回实例。
- java.util.Locale::getDisplayName：将返回一个包含这些U扩展名的显示名称的字符串。
- java.util.spi.LocaleNameProvider：将为这些U扩展的键和类型提供新的SPI。


### 备用内存设备上分配堆内存
启用HotSpot VM以在用户指定的备用内存设备上分配Java对象堆。随着廉价的NV-DIMM内存的可用性，未来的系统可能配备了异构的内存架构。这种技术的一个例子是英特尔的3D XPoint。这样的体系结构，除了DRAM之外，还会有一种或多种类型的非DRAM内存，具有不同的特征。具有与DRAM具有相同语义的可选内存设备，包括原子操作的语义，因此可以在不改变现有应用程序代码的情况下使用DRAM代替DRAM。所有其他的内存结构，如代码堆、metaspace、线程堆栈等等，都将继续驻留在DRAM中。

参考以下使用案例：

- 在多JVM部署中，某些JVM（如守护进程，服务等）的优先级低于其他JVM。与DRAM相比，NV-DIMM可能具有更高的访问延迟。低优先级进程可以为堆使用NV-DIMM内存，允许高优先级进程使用更多DRAM。
- 诸如大数据和内存数据库等应用程序对内存的需求不断增加。这种应用可以将NV-DIMM用于堆，因为与DRAM相比，NV-DIMM可能具有更大的容量，成本更低。

### 基于实验JAVA 的JIT 编译器
启用基于Java的JIT编译器Graal，将其用作Linux / x64平台上的实验性JIT编译器。Graal是一个基于Java的JIT编译器,它是JDK 9中引入的Ahead-of-Time（AOT）编译器的基础。使它成为实验性JIT编译器是Project Metropolis的一项举措，它是下一步是研究JDK的基于Java的JIT的可行性。

使Graal可用作实验JIT编译器，从Linux / x64平台开始。Graal将使用JDK 9中引入的JVM编译器接口（JVMCI）。Graal已经在JDK中，因此将它作为实验JIT将主要用于测试和调试工作。要启用Graal作为JIT编译器，请在java命令行上使用以下选项：
```
-XX:+UnlockExperimentalVMOptions -XX:+UseJVMCICompiler
```
### Root 证书
在JDK中提供一组默认的root 认证权威（CA）证书。在Oracle的Java SE根CA程序中开源root证书，以使OpenJDK构建对开发人员更有吸引力，并减少这些构建和Oracle JDK构建之间的差异。

cacerts密钥存储库是JDK的一部分，它的目的是包含一组root证书，这些root证书可以用来在各种安全协议中使用的证书链中建立信任。然而，JDK源代码中的cacerts密钥库目前是空的。因此，诸如TLS之类的关键安全组件在OpenJDK构建中不会默认工作。为了解决这个问题，用户必须配置和填充cacerts密钥库，并使用一组root证书来记录，例如， [JDK 9 release notes](https://www.oracle.com/technetwork/java/javase/9all-relnotes-3704433.html#JDK-8189131)。

### 基于时间的版本控制
在[JEP 223](http://openjdk.java.net/jeps/223) 引入的版本字符串方案比以往有了显著的改进，但是，该方案并不适合未来，现在Java SE平台和JDK的新版本[严格按照六个月](https://mreinhold.org/blog/forward-faster)的节奏发布。JEP 223方案的主要困难在于发行版的版本号对于其前身的重要性和兼容性进行了编码。然而，在基于时间发布模式中，这些品质并不是事先知道的。在发布的开发周期中，它们可能会发生变化，直到最终的功能被集成为止。因此发布的版本号也是未知的。

使用JEP 223的版本号语义，每个在JDK发布版上工作或者在其上构建或使用组件的人都必须先说明发布的发布日期，然后切换到说版本号，已知。维护库，框架和工具的开发人员必须准备好更改在每个JDK发布周期后期检查版本号的代码。这对所有参与者来说都很尴尬和混乱。

因此，这里提出的主要改变是重新编制版本号来编码，而不是编码的兼容性和重要性，而是按照发布周期的时间推移。这是更适合基于时间的发布模型，因为每个发布周期，因此每个发布的版本号，总是提前知道。

后续的版本格式为：[1-9][0-9]*((\.0)*\.[1-9][0-9]*)*

该格式可以是任意长度，但前四个被赋予特定含义，如：$FEATURE.$INTERIM.$UPDATE.$PATCH

- $FEATURE：功能发布计数器，不管发布内容如何，都会为每个功能发布增加。功能可能会添加到功能发布中; 如果提前通知提前至少发布一次功能发布，它们也可能会被删除。如果合理，可能会做出不兼容的更改。
- $INTERIM：临时版本计数器，对于包含兼容错误修复和增强功能的非功能版本递增，但没有不兼容的更改，没有功能移除，也没有对标准API的更改。
- $UPDATE：更新版本计数器增加了兼容更新版本，可解决新功能中的安全问题，回归和错误。
- $PATCH：紧急修补程序释放计数器只有在需要生成紧急释放以解决关键问题时才会增加。

版本号永远不会有零元素结尾。如果一个元素及其后的所有元素在逻辑上具有零值，那么它们全部被省略。

在严格六个月的发布模式下，版本号如下所示：

- $FEATURE 每六个月增加一次。如：2018年3月发布的是JDK 10，2018年9月发布的是JDK 11，等等。
- $INTERIM 总是为零，因为六个月的模型不包括临时版本。在此为保留其灵活性，以便将来对版本模型的修订可能包括此类版本。
- $UPDATE 在$FEATURE发布后的一个月递增，之后每三个月递增一次：如2018年4月发布JDK 10.0.1。7月发布的是JDK 10.0.2等等。


## JDK11
> 181: [Nest-Based Access Control](https://openjdk.java.net/jeps/181)  
> 309: [Dynamic Class-File Constants](https://openjdk.java.net/jeps/309)  
> 315: [Improve Aarch64 Intrinsics](https://openjdk.java.net/jeps/315)  
> 318: [Epsilon: A No-Op Garbage Collector](https://openjdk.java.net/jeps/318)  
> 320: [Remove the Java EE and CORBA Modules](https://openjdk.java.net/jeps/320)  
> 321: [HTTP Client (Standard)](https://openjdk.java.net/jeps/321)  
> 323: [Local-Variable Syntax for Lambda Parameters](https://openjdk.java.net/jeps/323)  
> 324: [Key Agreement with Curve25519 and Curve448](https://openjdk.java.net/jeps/324)  
> 327: [Unicode 10](https://openjdk.java.net/jeps/327)  
> 328: [Flight Recorder](https://openjdk.java.net/jeps/328)  
> 329: [ChaCha20 and Poly1305 Cryptographic Algorithms](https://openjdk.java.net/jeps/329)  
> 330: [Launch Single-File Source-Code Programs](https://openjdk.java.net/jeps/330)  
> 331: [Low-Overhead Heap Profiling](https://openjdk.java.net/jeps/331)  
> 332: [Transport Layer Security (TLS) 1.3](https://openjdk.java.net/jeps/332)  
> 333: [ZGC: A Scalable Low-Latency Garbage Collector  
>    (Experimental)](https://openjdk.java.net/jeps/333)  
> 335: [Deprecate the Nashorn JavaScript Engine](https://openjdk.java.net/jeps/335)  
> 336: [Deprecate the Pack200 Tools and API](https://openjdk.java.net/jeps/336)

## JDK12

> 189: [Shenandoah: A Low-Pause-Time Garbage Collector (Experimental)](https://openjdk.java.net/jeps/189)
> 230: [Microbenchmark Suite](https://openjdk.java.net/jeps/230)
> 325: [Switch Expressions (Preview)](https://openjdk.java.net/jeps/325)
> 334: [JVM Constants API](https://openjdk.java.net/jeps/334)
> 340: [One AArch64 Port, Not Two](https://openjdk.java.net/jeps/340)
> 341: [Default CDS Archives](https://openjdk.java.net/jeps/341)
> 344: [Abortable Mixed Collections for G1](https://openjdk.java.net/jeps/344)
> 346: [Promptly Return Unused Committed Memory from G1](https://openjdk.java.net/jeps/346)
## JDK13

> 350:[Dynamic CDS Archives](https://openjdk.java.net/jeps/350)
> 351:[ZGC: Uncommit Unused Memory](https://openjdk.java.net/jeps/351)
> 353:[Reimplement the Legacy Socket API](https://openjdk.java.net/jeps/353)
> 354:[Switch Expressions (Preview)](https://openjdk.java.net/jeps/354)
> 355:[Text Blocks (Preview)](https://openjdk.java.net/jeps/355)

## JDK14
http://openjdk.java.net/projects/jdk/14/
## JDK15
http://openjdk.java.net/projects/jdk/15/
## JDK16
http://openjdk.java.net/projects/jdk/16/
## JDK17
http://openjdk.java.net/projects/jdk/17/








