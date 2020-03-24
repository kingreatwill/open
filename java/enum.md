<!--toc-->
[TOC]

# 枚举

## 为什么需要枚举

**常量定义它不香吗？为啥非得用枚举？**

举个栗子，就以B站上传视频为例，视频一般有三个状态：**草稿**、**审核**和**发布**，我们可以将其定义为**静态常量**：

    public class VideoStatus {
        
        public static final int Draft = 1; //草稿
        
        public static final int Review = 2; //审核
        
        public static final int Published = 3; //发布
    }

对于这种**单值类型**的静态常量定义，本身也没错，主要是在使用的地方没有一个明确性的约束而已，比如：

    void judgeVideoStatus( int status ) {
        
        ...
        
    }

比如这里的 `judgeVideoStatus` 函数的本意是传入 `VideoStatus` 的三种静态常量之一，但由于没有类型上的约束，因此传入任意一个`int`值都是可以的，编译器也不会提出任何警告。

但是在枚举类型出现之后，上面这种情况就可以用枚举严谨地去约束，比如用枚举去定义视频状态就非常简洁了：

    public enum VideoStatus {
        Draft, Review, Published
    }

而且主要是在用枚举的地方会有更强的**类型约束**：

    // 入参就有明确类型约束
    void judgeVideoStatus( VideoStatus status ) {
        
        ...
        
    }

这样在使用 `judgeVideoStatus` 函数时，入参类型就会受到明确的类型约束，一旦传入无效值，编译器就会帮我们检查，从而规避潜在问题。

除此之外，枚举在扩展性方面比普常量更方便、也更优雅。

* * *

## 重新系统认识一下枚举

还是拿前文《答应我，别再if/else走天下了可以吗》中的那个例子来说：比如，在后台管理系统中，肯定有用户角色一说，而且角色一般都是固定的，适合定义成一个枚举：

    public enum UserRole {
    
        ROLE_ROOT_ADMIN,  // 系统管理员
    
        ROLE_ORDER_ADMIN, // 订单管理员
    
        ROLE_NORMAL       // 普通用户
    }

接下来我们就用这个`UserRole`为例来说明**枚举的所有基本用法**：

    UserRole role1 = UserRole.ROLE_ROOT_ADMIN;
    UserRole role2 = UserRole.ROLE_ORDER_ADMIN;
    UserRole role3 = UserRole.ROLE_NORMAL;
    
    // values()方法：返回所有枚举常量的数组集合
    for ( UserRole role : UserRole.values() ) {
        System.out.println(role);
    }
    // 打印：
    // ROLE_ROOT_ADMIN
    // ROLE_ORDER_ADMIN
    // ROLE_NORMAL
    
    // ordinal()方法：返回枚举常量的序数，注意从0开始
    System.out.println( role1.ordinal() ); // 打印0
    System.out.println( role2.ordinal() ); // 打印1
    System.out.println( role3.ordinal() ); // 打印2
    
    // compareTo()方法：枚举常量间的比较
    System.out.println( role1.compareTo(role2) ); //打印-1
    System.out.println( role2.compareTo(role3) ); //打印-2
    System.out.println( role1.compareTo(role3) ); //打印-2
    
    // name()方法：获得枚举常量的名称
    System.out.println( role1.name() ); // 打印ROLE_ROOT_ADMIN
    System.out.println( role2.name() ); // 打印ROLE_ORDER_ADMIN
    System.out.println( role3.name() ); // 打印ROLE_NORMAL
    
    // valueOf()方法：返回指定名称的枚举常量
    System.out.println( UserRole.valueOf( "ROLE_ROOT_ADMIN" ) );
    System.out.println( UserRole.valueOf( "ROLE_ORDER_ADMIN" ) );
    System.out.println( UserRole.valueOf( "ROLE_NORMAL" ) );

除此之外，枚举还可以用于`switch`语句中，而且意义更加明确：

    UserRole userRole = UserRole.ROLE_ORDER_ADMIN;
    switch (userRole) {
        case ROLE_ROOT_ADMIN:  // 比如此处的意义就非常清晰了，比1，2，3这种数字好！
            System.out.println("这是系统管理员角色");
            break;
        case ROLE_ORDER_ADMIN:
            System.out.println("这是订单管理员角色");
            break;
        case ROLE_NORMAL:
            System.out.println("这是普通用户角色");
            break;
    }

* * *

## 自定义扩充枚举

上面展示的枚举例子非常简单，仅仅是**单值的情形**，而实际项目中用枚举往往是**多值**用法。

比如，我想扩充一下上面的`UserRole`枚举，在里面加入 **角色名 \-\- 角色编码** 的对应关系，这也是实际项目中常用的用法。

这时候我们可以在枚举里自定义各种属性、构造函数、甚至各种方法：

    public enum UserRole {
    
        ROLE_ROOT_ADMIN( "系统管理员", 000000 ),
        ROLE_ORDER_ADMIN( "订单管理员", 100000 ),
        ROLE_NORMAL( "普通用户", 200000 ),
        ;
    
        // 以下为自定义属性
        
        private final String roleName;  //角色名称
    
        private final Integer roleCode; //角色编码
    
        // 以下为自定义构造函数
        
        UserRole( String roleName, Integer roleCode ) {
            this.roleName = roleName;
            this.roleCode = roleCode;
        }
    
        // 以下为自定义方法
        
        public String getRoleName() {
            return this.roleName;
        }
    
        public Integer getRoleCode() {
            return this.roleCode;
        }
    
        public static Integer getRoleCodeByRoleName( String roleName ) {
            for( UserRole enums : UserRole.values() ) {
                if( enums.getRoleName().equals( roleName ) ) {
                    return enums.getRoleCode();
                }
            }
            return null;
        }
    
    }

从上述代码可知，在`enum`枚举类中完全可以像在普通`Class`里一样声明属性、构造函数以及成员方法。

***

## 枚举 \+ 接口 = ?

比如在我的前文《答应我，别再if/else走天下了可以吗》中讲烦人的`if/else`消除时，就讲过如何通过**让枚举去实现接口**来方便的完成。

这地方不妨再回顾一遍：

什么角色能干什么事，这很明显有一个对应关系，所以我们首先定义一个公用的接口`RoleOperation`，表示不同角色所能做的操作：

    public interface RoleOperation {
        String op();  // 表示某个角色可以做哪些op操作
    }

接下来我们将不同角色的情况全部交由枚举类来做，定义一个枚举类`RoleEnum`，并让它去实现`RoleOperation`接口：

    public enum RoleEnum implements RoleOperation {
    
        // 系统管理员(有A操作权限)
        ROLE_ROOT_ADMIN {
            @Override
            public String op() {
                return "ROLE_ROOT_ADMIN:" + " has AAA permission";
            }
        },
    
        // 订单管理员(有B操作权限)
        ROLE_ORDER_ADMIN {
            @Override
            public String op() {
                return "ROLE_ORDER_ADMIN:" + " has BBB permission";
            }
        },
    
        // 普通用户(有C操作权限)
        ROLE_NORMAL {
            @Override
            public String op() {
                return "ROLE_NORMAL:" + " has CCC permission";
            }
        };
    }

这样，在调用处就变得异常简单了，一行代码就行了，根本不需要什么`if/else`：

    public class JudgeRole {
        public String judge( String roleName ) {
            // 一行代码搞定！之前的if/else灰飞烟灭
            return RoleEnum.valueOf(roleName).op();
        }
    }

而且这样一来，以后假如我想扩充条件，只需要去枚举类中**加代码**即可，而不用改任何老代码，非常符合**开闭原则**！

* * *

## 枚举与设计模式

什么？枚举还能实现设计模式？

是的！不仅能而且还能实现好几种！

### 1、单例模式

    public class Singleton {
    
        // 构造函数私有化，避免外部创建实例
        private Singleton() {
    
        }
    
        //定义一个内部枚举
        public enum SingletonEnum{
    
            SEED;  // 唯一一个枚举对象，我们称它为“种子选手”！
    
            private Singleton singleton;
    
            SingletonEnum(){
                singleton = new Singleton(); //真正的对象创建隐蔽在此！
            }
    
            public Singleton getInstnce(){
                return singleton;
            }
        }
    
        // 故意外露的对象获取方法，也是外面获取实例的唯一入口
        public static Singleton getInstance(){
            return SingletonEnum.SEED.getInstnce(); // 通过枚举的种子选手来完成
        }
    }

### 2、策略模式

这个也比较好举例，比如用枚举就可以写出一个基于策略模式的加减乘除计算器

    public class Test {
    
        public enum Calculator {
    
            ADDITION {
                public Double execute( Double x, Double y ) {
                    return x + y; // 加法
                }
            },
    
            SUBTRACTION {
                public Double execute( Double x, Double y ) {
                    return x - y; // 减法
                }
            },
    
            MULTIPLICATION {
                public Double execute( Double x, Double y ) {
                    return x * y; // 乘法
                }
            },
    
    
            DIVISION {
                public Double execute( Double x, Double y ) {
                    return x/y;  // 除法
                }
            };
    
            public abstract Double execute(Double x, Double y);
        }
        
        public static void main(String[] args) {
            System.out.println( Calculator.ADDITION.execute( 4.0, 2.0 ) );       
            // 打印 6.0
            System.out.println( Calculator.SUBTRACTION.execute( 4.0, 2.0 ) );    
            // 打印 2.0
            System.out.println( Calculator.MULTIPLICATION.execute( 4.0, 2.0 ) ); 
            // 打印 8.0
            System.out.println( Calculator.DIVISION.execute( 4.0, 2.0 ) );       
            // 打印 2.0
        }
    }

* * *

## 专门用于枚举的集合类

我们平常一般习惯于使用诸如：`HashMap` 和 `HashSet`等集合来盛放元素，而对于枚举，有它专门的集合类：`EnumSet`和`EnumMap`

### 1、EnumSet

`EnumSet` 是专门为盛放枚举类型所设计的 `Set` 类型。

还是举例来说，就以文中开头定义的角色枚举为例：

    public enum UserRole {
    
        ROLE_ROOT_ADMIN,  // 系统管理员
    
        ROLE_ORDER_ADMIN, // 订单管理员
    
        ROLE_NORMAL       // 普通用户
    }

比如系统里来了一批人，我们需要查看他是不是某个角色中的一个：

    // 定义一个管理员角色的专属集合
    EnumSet<UserRole> userRolesForAdmin 
        = EnumSet.of( 
            UserRole.ROLE_ROOT_ADMIN,
            UserRole.ROLE_ORDER_ADMIN 
        );
    
    // 判断某个进来的用户是不是管理员
    Boolean isAdmin( User user ) {
        
        if( userRoles.contains( user.getUserRole() ) )
            return true;
        
        return false;
    }

### 2、EnumMap

同样，`EnumMap` 则是用来专门盛放枚举类型为`key`的 `Map` 类型。

比如，系统里来了一批人，我们需要统计不同的角色到底有多少人这种的话：

    Map<UserRole,Integer> userStatisticMap = new EnumMap<>(UserRole.class);
    
    for ( User user : userList ) {
        Integer num = userStatisticMap.get( user.getUserRole() );
        if( null != num ) {
            userStatisticMap.put( user.getUserRole(), num+1 );
        } else {
            userStatisticMap.put( user.getUserRole(), 1 );
        }
    }

用`EnumMap`可以说非常方便了。

