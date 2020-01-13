> * 原文地址：[https://wiki.postgresql.org/wiki/Don%27t_Do_This#Why_not.3F](https://wiki.postgresql.org/wiki/Don%27t_Do_This#Why_not.3F)
> * 译文地址：[https://github.com/watermelo/dailyTrans](https://github.com/watermelo/dailyTrans/blob/master/db/postgresql_dont_do_this.md)
> * 译者：咔叽咔叽  
> * 译者水平有限，如有翻译或理解谬误，烦请帮忙指出

## 1. 数据库编码

### 1.1 不要使用 SQL_ASCII 编码

#### 1.1.1 为什么
虽然这个名称看起来和 `ASCII` 有关系，但并非如此。相反，它只是禁止使用空字节。

更重要的是，`SQL_ASCII` 对所有的编码转换的函数而言意味着“不转换”。也就是说，原始字节编码是什么就是什么。除非特别小心，否则 `SQL_ASCII` 编码的数据库可能最终会存储许多不同编码的混合数据，可能无法可靠地恢复原始字符。

#### 1.1.2 什么时候可以使用
如果你的输入数据已经是编码的混合的数据，例如 IRC 的日志或非 MIME 兼容的电子邮件，那么 `SQL_ASCII` 可能是最后的方法 - 应该首先考虑使用 [`bytea`](http://postgres.cn/docs/9.4/datatype-binary.html#AEN5493) 编码，或者可以检测是否为 UTF8 编码，如果非 UTF8 编码，例如 WIN1252 编码的数据可以假定为 UTF8 编码。

## 2. 工具

### 2.1 不要使用 psql -W 或者 psql --password
不要使用 psql -W 或者 psql --password

#### 2.1.1 为什么
如果使用 --password 或 -W 标志连接服务，`psql` 会提示你需要输入密码 - 因此即使服务器不需要密码，也会提示你输入密码。

这个选项没有必要，它会让人误以为服务器需要密码。如果你登录的用户没有设置密码，或者你在提示时输入了错误的密码，你仍然会成功登录并认为这就是正确的密码 - 但你无法使用这个密码从其他客户端（通过 localhost 连接）或以其他用户身份登录。

#### 2.1.2 什么时候可以使用
不要使用。

### 2.2 不要使用 RULE
不要使用 [RULE](https://www.postgresql.org/docs/current/sql-createrule.html)，（译者注： CREATE RULE 定义一个适用于特定表或者视图的新规则）如果想要用，请使用[触发器](https://www.postgresql.org/docs/current/plpgsql-trigger.html)替代。

#### 2.2.1 为什么
RULE 非常强大，但并不好理解。它看起来像是一些条件逻辑，但实际上它会重写查询或向查询中添加其他查询。

这意味着[所有 non-trivial 的规则都是不正确的](http://blog.rhodiumtoad.org.uk/2010/06/21/the-rule-challenge/)。（译者注：关于 non-trivial 的定义参考引用文章）

Depesz 对此有[更多话要说](https://www.depesz.com/2010/06/15/to-rule-or-not-to-rule-that-is-the-question/)。

#### 2.2.2 什么时候可以使用
不要使用。

### 2.3 不要使用表继承
不要使用[表继承](https://www.postgresql.org/docs/current/tutorial-inheritance.html)，如果真的要用，可以使用外键来替代。

#### 2.3.1 为什么
表继承是一个时髦的概念，其中数据库与面向对象的代码紧耦合。事实证明，这些耦合的东西实际上并没有产生预期的结果。

#### 2.3.2 什么时候可以使用
几乎不使用……差不多。现在表分区是本地完成的，表继承的常见场景已经被一些特性所取代。

## 3. SQL 语句
### 3.1 不要使用 NOT IN
不要使用 NOT IN，或 NOT 和 IN 的任意组合，如 NOT(x IN (select...))。

（如果你认为你想要 NOT IN (select...) 那么你应该使用 NOT EXISTS 替代。）

#### 3.1.1 为什么
两个理由：

1. 如果存在 NULL 值，则 NOT IN 会以意外的方式运行：

```plain
select * from foo where col not in (1,null);
  -- always returns 0 rows

select * from foo where col not in (select x from bar);
  -- returns 0 rows if any value of bar.x is null
```

发生这种情况是因为如果 col = 1 则 col IN(1，null)  返回 TRUE，否则返回 NULL（即它永远不会返回 FALSE）。由于 NOT(TRUE) 为 FALSE，但 NOT(NULL) 仍为 NULL，因此 NOT(col IN(1，null)) （与 col NOT IN(1，null)相同）无法返回 TRUE，也就是说 NOT IN (1, NULL) 这种形式永远不会返回数据。

2. 由于上面的第 1 点，NOT IN (SELECT ...) 不能很好地优化。特别是，规划器（planner 负责生成查询计划）无法将其转换为 anti-join，因此它变为哈希子规划或普通子规划。哈希子规划很快，但规划器只允许该计划用于小结果集;普通的子计划非常慢（事实上是 O(N²)时间复杂度）。这意味着在小规模测试中性能看起来不错，但一旦数据量大，性能就会减慢 5 个或更多个数量级; 我们不希望这种情况发生。

#### 3.1.2 什么时候可以使用
NOT IN (list, of, values, ...)只是在列表中有 null 值（参数或其他方式）时，才会有问题。所以在排除没有 null 值时，是可以用的。

### 3.2 不要使用大写字母命名
不要使用 `NamesLikeThis`，而是使用 `names_like_this` 的命名方式。

#### 3.2.1 为什么
PostgreSQL 将表，列，函数等名称转换为小写，除非它们使用“双引号”扩起来才不会被转换。

所以 `create table Foo()` 将会创建一个表名为 `foo` 的表，执行 `create table "Bar"()` 才会创建表名为 `Bar` 的表。

这些查询语句将会正常执行：`select * from Foo, select * from foo, select * from "Bar"`

这些查询语句会报错 “no such table”：`select * from "Foo", select * from Bar, select * from bar`

这意味着如果在表名或列名中使用大写字母，则查询时必须使用双引号。这很烦人，但是当你使用其他工具访问数据库时，其中一些名称使用双引号，而另一些则不使用，这会让人感到困惑。

坚持使用 a-z，0-9 和下划线来表示名称，就不必再担心了。

#### 3.2.2 什么时候可以使用
如果在输出中显示好看的名称很重要，那么你可能想使用大写字母。但是你也可以使用列别名，也可以在查询中输出好看的名称：`select character_name as "Character Name" from foo`。

### 3.3 不要使用 BETWEEN（尤其是时间戳）

#### 3.3.1 为什么
BETWEEN 使用闭区间比较：范围两端的值会包含在结果中。

这是一个查询的问题

```sql
SELECT * FROM blah WHERE timestampcol BETWEEN '2018-06-01' AND '2018-06-08'
```

这将包括时间戳恰好为 2018-06-08 00:00:00.000000 的结果。查询可以工作，但是由于是闭区间，可能在下一次查询会包含这个时刻值（例如 '2018-06-08' AND '2018-06-09' 就会包含那一刻的值）。

用下面的语句替换

```sql
SELECT * FROM blah WHERE timestampcol >= '2018-06-01' AND timestampcol < '2018-06-08'
```

#### 3.3.2 什么时候可以使用
BETWEEN 对于整数或日期等离散类型的数据是安全的，需要记住 BETWEEN 是闭区间。但使用 BETWEEN 可能是一个坏习惯。

## 4. 日期/时间 的存储
（译者注：[日期/时间 中文文档](http://postgres.cn/docs/9.4/datatype-datetime.html)）

### 4.1 不要使用 timestamp（without time zone）
不要使用 `timestamp` 类型来存储时间戳，而是使用 `timestamptz`（也称为带时区的时间戳）来存储。

#### 4.1.1 为什么
`timestamptz` 记录 UTC 的微秒数，你可以插入任何时区的值。默认情况下，它将显示当前时区的时间，但你可以转换成其他时区。

因为它存储了时间戳信息，可以用算法来转换成不同时区的时间戳。

`timestamp`（也称为没有时区的时间戳）不会执行任何操作，它只会存储你提供的日期和时间。你可以将其视为日历和时钟的图片，而不是时间点，没有时区信息。因此，没有时区的时间戳没法转换时区。

因此，如果你要存储的是时间点，而不是时钟图片，请使用 `timestamptz`。

[更多有关 timestamptz 的信息](https://it.toolbox.com/blogs/josh-berkus/zone-of-misunderstanding-092811)

#### 4.1.2 什么时候可以使用
如果你以抽象方式处理时间戳，或者只是为了 app 的保存和检索，而不对它们进行时间计算，那么 `timestamp` 可能是合适的。

### 4.2 不要使用 timestamp（without time zone）来存储 UTC 时间
将 UTC 值存储在没有时区的 `timestamp` ，通常是从其他缺乏可用时区支持的数据库继承数据的做法。

改为使用 `timestamp with time zone` 即 `timestamptz`。

#### 4.2.1 为什么
因为数据库无法知道是否是 UTC 时区。

这让时间计算变得复杂。例如，“给定时区 u.timezone 的最后一个午夜”的计算语句为：

```sql
date_trunc('day', now() AT TIME ZONE u.timezone) AT TIME ZONE u.timezone AT TIME ZONE 'UTC'
```

并且“u.timezone 中 x.datecol 日期之前的午夜”的计算语句为：

```sql
date_trunc('day', x.datecol AT TIME ZONE 'UTC' AT TIME ZONE u.timezone)
  AT TIME ZONE u.timezone AT TIME ZONE 'UTC'
```

#### 4.2.2 什么时候可以使用
如果与非时区支持数据库的兼容性胜过所有其他考虑因素。

### 4.3 不要使用 timetz
不要使用 `timetz` 类型，可以使用 `timestamptz` 代替。

#### 4.3.1 为什么
甚至手册也告诉你它只是为了遵守 SQL 标准而实现的。

> 带有时区的时间类型由 SQL 标准定义，但定义显示的属性让人怀疑它的可用性。在大多数情况下，日期，时间，没有时区的时间戳和带时区的时间戳的组合应该可以提供任何应用程序所需的日期/时间功能。

#### 4.3.2 什么时候可以使用
从不使用。

### 4.4 不要使用 CURRENT_TIME
不要使用 CURRENT_TIME 函数。使用以下是合适的：

- CURRENT_TIMESTAMP 或者 now() 如果你想要 `timestamp with time zone`
- LOCALTIMESTAMP 如果你想要 `timestamp without time zone`
- CURRENT_DATE 如果你想要 `date`
- LOCALTIME 如果你想要 `time`

#### 4.4.1 为什么
它返回 `timetz` 类型的值，关于 `timetz` 请看上一条解释。

#### 4.4.2 什么时候可以使用
从不使用。

### 4.5 不要使用 timestamp(0) 或者 timestamptz(0)
不要使用 timestamp() 或者 timestamptz() 进行时间戳的转换（尤其是 0）。

使用 `date_trunc('second', blah)` 替换。

#### 4.5.1 为什么
因为它会将小数部分四舍五入截断它。这可能会导致意外问题;考虑到当你将 `now()` 存储到这样一个列中时，你可能会在将来存储一个小数秒的值。

#### 4.5.2 什么时候可以使用
从不使用。

## 5. 文本存储

### 5.1 不要使用 char(n)
不要使用 `char(n)`，使用 `text` 可能更适合。

#### 5.1.1 为什么
使用 `char(n)` 类型的字段，如果长度不够会使用空格填充到声明的长度。这可能不是你想要的。

|名字 | 描述
 --- | ---
character varying(n), varchar(n) | 变长，有长度限制
character(n), char(n) | 定长，不足补空白
text | 变长，无长度限制


手册上说：

> char 类型的数值物理上都用空白填充到指定的长度 n， 并且以这种方式存储和显示。不过，在比较两个 char 类型的值时，尾随的空白是无关紧要的，不需要理会。 在空白比较重要的排序规则中，这个行为会导致意想不到的结果， 比如 `SELECT 'a '::CHAR(2) collate "C" < 'a\n'::CHAR(2)`返回真。 在将 char 值转换成其它字符串类型的时候， 它后面的空白会被删除。请注意， 在 varchar 和 text 数值里， 结尾的空白是有语意的。 并且当使用模式匹配时，如 LIKE，使用正则表达式。

空格填充确实浪费空间，也不会让操作变得更快;事实上，很多情况下我们还要去掉空格。

> 提示: 这三种类型之间没有性能差别，除了当使用填充空白类型时的增加存储空间， 和当存储长度约束的列时一些检查存入时长度的额外的 CPU 周期。 虽然在某些其它的数据库系统里，char(n) 有一定的性能优势，但在 PostgreSQL 里没有。 事实上，char(n) 通常是这三个中最慢的， 因为额外存储成本。在大多数情况下，应该使用 text 或 varchar。

#### 5.1.2 什么时候可以使用
当你移植使用了固定宽度字段的非常非常旧的软件时。或者当你阅读上面手册中的片段并认为“是的，这是完全合理的，并且符合我的要求” 时可以使用。

### 5.2 即使对于固定长度的标识符，也不要使用 char(n)
有时候人们用“我的值刚好是 N 个字符”（例如国家代码，哈希值或来自其他系统的标识符）来回应“为什么不要使用 char(n)”。其实，即使在这些场景下使用 `char(n)` 也不是一个好主意。

#### 5.2.1 为什么
对于太短的值，`char(n)` 会用空格填充它们。因此，带有确定长度的 `char(n)` 比较 `text` 而言没有任何实际好处。

#### 5.2.2 什么时候可以使用
从不使用。

### 5.3 默认情况下不要使用 varchar(n)
默认情况下不要使用 `varchar(n)` 类型。考虑使用 `varchar`（没有长度限制）或 `text` 替代。

#### 5.3.1 为什么
`varchar(n)` 是一个带长度的文本字段，如果你尝试将长度超过 n 个字符（而不是字节）的字符串插入其中，则会引发错误。

varchar（没有 (n) ）或 `text` 是相似的，没有长度限制。如果在三种字段类型中插入相同的字符串，它们将占用完全相同的空间，并且性能基本没有差异。

如果你想要一个长度限制的文本字段，那么 `varchar(n)` 很不错，但是如果你定义姓氏字段为 `varchar(20)`，那么当 Hubert Blaine Wolfeschlegelsteinhausenbergerdorff 注册到你的服务时，将会报错。

有些数据库没有长 `text` 的类型，或者它们没有像 `varchar(n)` 那样被良好支持。那这些数据库的用户通常会使用类似 `varchar(255)` 的表示方法，但他们真正想要的是 `text`。

如果你需要约束字段中的值，比如说约束最大长度 - 或者是最小长度，或者是一组有限制的字符串 - [检查约束](https://www.postgresql.org/docs/current/ddl-constraints.html#DDL-CONSTRAINTS-CHECK-CONSTRAINTS)可以做到这些。

#### 5.3.2 什么时候可以使用
如果你想要一个文本字段，而且插入太长的字符串需要抛出错误，并且不想使用显式检查约束，那么 `varchar(n)` 是一个非常好的类型。只是使用时需要多考虑。

## 6. 其他数据类型

### 6.1 不要使用 money
`money` 数据类型实际上不太适合存储货币值。数字或整数可能更好。

#### 6.1.1 为什么
[大量理由](https://www.postgresql.org/message-id/flat/20130328092819.237c0106@imp#20130328092819.237c0106@imp)

money 类型存储带有固定小数精度的货币金额。 `lc_monetary` 用来设置格式化数字。但它的四舍五入的行为可能不是你想要的。

名字 | 存储容量 | 描述 | 范围
 --- | --- |  --- | ---
money | 8 字节 | 货币金额 | -92233720368547758.08 到 +92233720368547758.07

如果你更改了 `lc_monetary` 设置，则所有 money 列都将包含错误的值。这意味着如果你插入'$ 10.00'而 `lc_monetary` 设置为 `en_US.UTF-8`，你检索的值可能是'10，00 Lei'或'¥1,000'。

#### 6.1.2 什么时候可以使用
如果你只是用单一货币工作，不处理小数美分并且只做加法和减法运算，那么 `money` 类型可能是正确的。

### 6.2 不要使用 `serial`
对于新的应用程序，应使用 `identity`。

#### 6.2.1 为什么
`serial` 类型有一些[奇怪的行为](https://www.2ndquadrant.com/en/blog/postgresql-10-identity-columns/)，使结构，依赖和权限管理更繁琐。

#### 6.2.2 什么时候可以使用
- 如果你需要支持 10.0 版之前的 PostgreSQL。
- 在表继承的某些组合中
- 更一般地说，如果你以某种方式使用来自多个表的相同序列，尽管在这些情况下，显式声明可能优于 `serial` 类型。
