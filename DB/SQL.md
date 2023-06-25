

## sql可视化
[sqlflow](https://sqlflow.gudusoft.com/#/)

## 什么是SQL？

SQL意为结构化查询语言（Structured Query Language）。SQL用于和数据库交流。它是关系数据库管理系统的标准语言。SQL语句用于执行诸如更新数据库中的数据，或从数据库中检索数据等任务。

## 乘风破浪的 SQL
![](../img/db/sql-01.jpg)

实际上，2016 年 SQL 标准就已经增加了 JSON 功能的支持，包括：

- JSON 对象的存储与检索；
- 将 JSON 对象表示成 SQL 数据；
- 将 SQL 数据表示成 JSON 对象。

如今，主流的关系数据库也都增加了原生 JSON 数据类型和相关函数的支持，包括 Oracle、MySQL、SQL Server、PostgreSQL 等。

2019 年 9 月 17 图形查询语言（GQL）成为了继 SQL 之后另一种新的 ISO 标准数据库查询语言。与此同时，SQL 标准将会出现一个新的第 16 部分（SQL/PGQ）（Property Graph Query），在 SQL 中直接提供一些 GQL 功能。

随着互联网和大数据等新技术的发展，SQL 早已不仅仅是当年的关系数据库查询语言了；无论是面向对象特性（例如自定义类型）、文档数据（XML、JSON）的存储和处理、时态数据的存储和处理、复杂事件和流数据处理、数据科学中的多维数组以及图形数据库等各种 NoSQL 功能已经或者即将成为 SQL 标准中的一部分，One SQL to Rule Them All！

[原文](https://blog.csdn.net/horses/article/details/107185387)

## SQL发展的简要历史：
为了在各个数据库厂商之间取得更大的统一性，美国国家标准学会（American NationalStandards Institute,ANSI）于1986年发布了第一个SQL标准。
现在，SQL标准由ANSI和国际标准化组织（International Standards Organization,ISO）作为ISO/IEC 9075标准维护。

1986年，ANSI X3.135-1986，ISO/IEC 9075:1986，SQL-86
1989年，ANSI X3.135-1989，ISO/IEC 9075:1989，SQL-89
1992年，ANSI X3.135-1992，ISO/IEC 9075:1992，SQL-92（SQL2）
1999年，ISO/IEC 9075:1999，SQL:1999（SQL3）
2003年，ISO/IEC 9075:2003，SQL:2003
2008年，ISO/IEC 9075:2008，SQL:2008
2011年，ISO/IEC 9075:2011，SQL:2011
2016年，ISO/IEC 9075:2016，SQL:2016
2023年，ISO/IEC 9075:2023，SQL:2023

从SQL:1999开始，标准简称中的短横线（-）被换成了冒号（:），而且标准制定的年份也改用四位数字了。前一个修改的原因是ISO标准习惯上采用冒号，ANSI标准则一直采用短横线。后一个修改的原因是标准的命名也遇到了2000年问题。

SQL86大概只有几十页，SQL92正文大约有500页，而SQL99则超过了1000页。可以看出，从SQL99开始，SQL标准的个头就非常庞大了，内容包罗万象，已经没有人能够掌握标准的所有内容了。以SQL:2003为例，它包括以下9个部分（中间编号空缺是曾经被占用，之后被废弃的标准造成的）：

ISO/IEC9075-1: Framework (SQL/Framework)
ISO/IEC 9075-2: Foundation (SQL/Foundation)
ISO/IEC 9075-3: Call Level Interface (SQL/CLI)
ISO/IEC 9075-4: Persistent Stored Modules (SQL/PSM)
ISO/IEC 9075-9: Management of External Data (SQL/MED)
ISO/IEC 9075-10: Object Language Bindings (SQL/OLB)
ISO/IEC 9075-11: Information and Definition Schemas (SQL/Schemata)
ISO/IEC 9075-13: Java Routines and Types Using the Java Programming Language(SQL/JRT)
ISO/IEC 9075-14: XML-Related Specifications (SQL/XML)


> 国际标准化组织（ISO）于 2023 年 6 月 1 日正式发布了最新 SQL 标准，也就是 SQL:2023。
[SQL:2023标准正式发布！](https://blog.csdn.net/horses/article/details/131008387)SQL 标准是一个公开资料，但是并不免费。最新标准包含 11 个部分的内容，具体如下：
1. [ISO/IEC 9075-1:2023(en) Information technology — Database languages SQL — Part 1: Framework (SQL/Framework)](https://www.iso.org/standard/76583.html)
[ISO/IEC 9075-1 信息技术 – 数据库语言 – SQL – 第 1 部分：框架（SQL/框架）](https://www.iso.org/standard/76583.html)
1. [ISO/IEC 9075-2 信息技术 – 数据库语言 – SQL – 第 2 部分：基本原则（SQL/基本原则）](https://www.iso.org/standard/76584.html)
1. [ISO/IEC 9075-3 信息技术 – 数据库语言 – SQL – 第 3 部分：调用级接口（SQL/CLI）](https://www.iso.org/standard/84803.html)
1. [ISO/IEC 9075-4 信息技术 – 数据库语言 – SQL – 第 4 部分：持久存储模块（SQL/PSM）](https://www.iso.org/standard/76585.html)
1. [ISO/IEC 9075-9 信息技术 – 数据库语言 – SQL – 第 9 部分：外部数据管理（SQL/MED）](https://www.iso.org/standard/84804.html)
1. [ISO/IEC 9075-10 信息技术 – 数据库语言 – SQL – 第10 部分：对象语言绑定（SQL/OLB）](https://www.iso.org/standard/84805.html)
1. [ISO/IEC 9075-11 信息技术 – 数据库语言 – SQL – 第 11 部分：信息与定义概要（SQL/Schemata）](https://www.iso.org/standard/76586.html)
1. [ISO/IEC 9075-13 信息技术 – 数据库语言 – SQL – 第 13 部分：使用 Java 编程语言的 SQL 程序与类型（SQL/JRT）](https://www.iso.org/standard/84806.html)
1. [ISO/IEC 9075-14 信息技术 – 数据库语言 – SQL – 第 14 部分：XML 相关规范（SQL/XML）](https://www.iso.org/standard/76587.html)
1. [ISO/IEC 9075-15 信息技术 – 数据库语言 – SQL – 第 15 部分：多维数组（SQL/MDA）](https://www.iso.org/standard/84807.html)
1. [ISO/IEC 9075-16 信息技术 – 数据库语言 – SQL – 第 16 部分：属性图查询（SQL/PGQ）](https://www.iso.org/standard/79473.html)


**如果要了解标准的内容，比较推荐的方法是泛读SQL92**

## 参考文献
http://www.ansi.org
http://www.iso.ch
http://sqlstandards.org
http://www.wiscorp.com

[适用于SQL-92，SQL-99和SQL-2003的BNF语法](https://github.com/ronsavage/SQL)

https://ronsavage.github.io/SQL/sql-92.bnf.html

[An Introduction to SQL Standard（SQL标准简介）](https://zedware.github.io/SQL-Standard/)

## SQL统计
### mysql按日、月进行统计，不全的记录补0
1、生成日期表
```sql
select @cdate:=DATE_ADD(@cdate,INTERVAL -1 day) as DAY_TIME
from(select @cdate:=DATE_ADD(STR_TO_DATE('2021-10-31','%Y-%m-%d'),INTERVAL 1 day)
from b2b_order limit 31) aa 
where @cdate>'2021-10-01' and @cdate<='2021-11-01'
```
2、生成日期表后，就是进行左关联
```sql
select STR_TO_DATE(a.DAY_TIME,'%Y-%m-%d'),round(COALESCE(b.totalMoney,0),2)
from 
(select @cdate:=DATE_ADD(@cdate,INTERVAL -1 day) as DAY_TIME
from(select @cdate:=DATE_ADD(STR_TO_DATE('2021-10-31','%Y-%m-%d'),INTERVAL 1 day)
from b2b_order limit 31) aa 
where @cdate>'2021-10-01' and @cdate<='2021-11-01') a 
left join
(select STR_TO_DATE(t.ORDER_DATE,'%Y-%m-%d') as DAY_TIME,sum(t.TOTAL_DEAL_AMOUNT) as totalMoney
from  BASE_ORDER t
where t.dr = 0 
 and t.ORDER_DATE>= STR_TO_DATE('2021-10-01','%Y-%m-%d')
 and t.ORDER_DATE<= STR_TO_DATE('2021-11-01','%Y-%m-%d')
GROUP BY STR_TO_DATE(t.ORDER_DATE,'%Y-%m-%d')
) b
on a.DAY_TIME = b.DAY_TIME
order by a.day_time
```
3、如果按照月份进行统计思路一样的，先构建月份表
```sql
select @cdate:=DATE_ADD(@cdate,INTERVAL -1 month),
DATE_FORMAT(@cdate,'%Y-%m') as DAY_TIME
from
 (select @cdate:=DATE_ADD(STR_TO_DATE('2021-12-01','%Y-%m-%d'),INTERVAL 1 month)
from b2b_order limit 12) aa
where @cdate>'2021-01-01' and @cdate<='2022-01-01'
```
4、再进行关联查询
```sql
select a.DAY_TIME,round(COALESCE(tttt.total_money,0),2)
 from 
 (select @cdate:=DATE_ADD(@cdate,INTERVAL -1 month),DATE_FORMAT(@cdate,'%Y-%m') as DAY_TIME 
    from 
     (select @cdate:=DATE_ADD(STR_TO_DATE('2021-12-01','%Y-%m-%d'),INTERVAL 1 month) 
        from b2b_order limit 12) aa
        where @cdate>'2021-01-01' and @cdate<='2022-01-01') a
left join
  (select ttt.DAY_TIME ,sum(ttt.total_money) as total_money
   from 
     (select DATE_FORMAT(tt.DAY_TIME,'%Y-%m') as DAY_TIME,tt.total_money
      from 
        (select STR_TO_DATE(t.ORDER_DATE,'%Y-%m-%d') as DAY_TIME,sum(t.TOTAL_DEAL_AMOUNT) as total_money
         from  BASE_ORDER t 
          where t.dr = 0
           and t.ORDER_DATE>= STR_TO_DATE('2021-01-01','%Y-%m-%d')
           and t.ORDER_DATE<= STR_TO_DATE('2022-01-01','%Y-%m-%d')
            GROUP BY STR_TO_DATE(t.ORDER_DATE,'%Y-%m-%d')
        ) tt 
      ) ttt
      GROUP BY ttt.DAY_TIME
    ) tttt
on a.DAY_TIME = tttt.DAY_TIME 
order by a.day_time
```
5、这几个sql中涉及到的函数
- STR_TO_DATE（'字符串'，'%Y-%M-%D'）
- DATE_UP()
- DATE_FORMAT()
- ROUND
- COALESCE --没有记录则为0