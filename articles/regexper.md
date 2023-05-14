# regular expression 正则表达式

## 工具
[图解正则表达式 - ok+](https://regexper.com/)
https://gitlab.com/javallone/regexper-static

[图解正则 - ok++](https://jex.im/regulex/#!flags=&re=%5E(a%7Cb)*%3F%24)
https://github.com/CJex/regulex

[Regex visualizer & editor](https://github.com/Bowen7/regex-vis)

[正则表达式工具 - ok+](http://tool.chinaz.com/regex/)
[Regex Generator正则表达式生成器 - ok+](https://regex-generator.olafneumann.org/)

[正则表达式工具 - good](https://regexr.com/)
https://github.com/gskinner/regexr/

[regex101 - good++](https://regex101.com/)
https://github.com/firasdib/Regex101

[The Regular Expression Visualizer可视化](https://blog.robertelder.org/regular-expression-visualizer/)

[在线正则表达式工具](https://www.regextester.com/)
[在线正则表达式](http://www.regexe.com/)
[在线正则表达式工具](https://www.computerhope.com/cgi-bin/text-tool.pl)
[正则表达式教程 - Regex Tutorial](http://www.regular-expressions.info/wordboundaries.html)


[正则表达式工具- 在线](https://deerchao.cn/tools/wegester/)
[正则表达式工具 Regester](https://deerchao.cn/tools/regester/)

[正则表达式其它工具](http://www.regular-expressions.info/tools.html)
[RegexBuddy 收费 - good](http://www.regexbuddy.com/)
下载4.10：https://www.0daydown.com/01/243180.html
Download 百度网盘
链接: https://pan.baidu.com/s/1VIiYSL2OYosxh_T3pGvHLg 提取码: e3uu

[RegexMagic 正则表达式生成工具 收费 - good](https://www.regexmagic.com/)
下载：2.8：https://www.0daydown.com/01/149682.html
https://pan.baidu.com/s/1mENJB-YFUfkiM1-WC2-ycg

[PowerGREP  Windows grep 收费](https://www.powergrep.com/)

[.NET Regex - 不更新了](http://regexstorm.net/)
https://github.com/lonekorean/regex-storm

[visualRegex](https://wangwl.net/static/projects/visualRegex#)

[rubular](https://rubular.com/)

[正则表达式-不同语言之间的差别](https://www.regular-expressions.info/refcharacters.html)
[工具和教程](http://www.regexlab.com/zh/mtracer/download.htm)
[.NET regular expressions](https://docs.microsoft.com/en-us/dotnet/standard/base-types/regular-expression-language-quick-reference)

## 教程
[JavaScript RegExp](https://learnbyexample.github.io/learn_js_regexp/anchors.html)
[Python re(gex)](https://learnbyexample.github.io/py_regular_expressions/buy.html)
[Regular Expressions](https://github.com/EbookFoundation/free-programming-books/blob/master/books/free-programming-books.md#regular-expressions)



[正则表达式-不同语言之间的差别](https://www.regular-expressions.info/refcharacters.html)

[常用的正则表达式](https://github.com/cdoco/common-regex)
[正则表达式30分钟入门教程](https://deerchao.cn/tutorials/regex/regex.htm)
[学习正则表达式](https://github.com/ziishaned/learn-regex)
[Regex 正则表达式入门](https://www.cnblogs.com/codeshell/p/12825243.html)
[工具和教程](http://www.regexlab.com/zh/mtracer/download.htm)

## 常用的正则表达式
[常用正则表达式汇总](https://www.cnblogs.com/cxsabc/p/10627631.html)
[正则表达式汇总](https://mp.weixin.qq.com/s/h11xHEPxjy1IPE514IjYyw)

## 其它
vs code 插件any-rule   F1



```go
package main

import (
	"fmt"
	"regexp"
)

var myExp = regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)

func main() {
	match := myExp.FindStringSubmatch("1234.5678.9")
	result := make(map[string]string)
	for i, name := range myExp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	fmt.Printf("by name: %s %s\n", result["first"], result["second"])
}
```


# Cron Expression
## 工具
https://www.pppet.net/
## 将cron表达式转换为人类可读描述的
dotnet - https://github.com/bradymholt/cron-expression-descriptor
JavaScript - https://github.com/bradymholt/cRonstrue
Java - https://github.com/RedHogs/cron-parser
Java - https://github.com/voidburn/cron-expression-descriptor
Ruby - https://github.com/alpinweis/cronex
Python - https://github.com/Salamek/cron-descriptor
Go - https://github.com/lnquy/cron

## linux crontab任务(5位)
cron表达式　
表达式： `* * * * *`
含义：分钟(0-59) 小时(0-23) 日期(1-31) 月份(1-12) 星期(0-6)

crontab任务配置基本格式： 
`* *　 *　 *　 *　　command`

星号（*）可以用来代表所有有效的值。譬如，月份值中的星号意味着在满足其它制约条件后每月都执行该命令。
整数间的短线（-）指定一个整数范围。譬如，1-4 意味着整数 1、2、3、4。
用逗号（,）隔开的一系列值指定一个列表。譬如，1，2，3，4 标明这四个指定的整数。
正斜线（/）可以用来指定间隔频率。譬如：*/2用在日期字段中表示每2天执行一次该命令。

> 查看用户下的定时任务:crontab -l 或 cat /var/spool/cron/用户名

## 6位
Seconds Minutes Hours DayofMonth Month DayofWeek [Year]
*： 匹配所有值
?：匹配dayofMonth 和DayofWeek的一个值。
-：表范围。例如在分钟域中写1-10，表示第1到第10分钟每分钟触发一次。
/：前的表示起始，后的表示间隔。例如在Minutes域使用5/20,则意味着5分钟触发一次，而25,45等分别触发一次.
L：表示最后一个。在DayofWeek，4L表示每个月最后一个星期五
W:表示有效工作日(周一到周五)， 只能出现在DayofMonth域，系统将在离指定日期的最近的有效工作日触发事件。例如在DayofMonth使用5W，如果5日是星期六，则将在最近的工作日星期五，即4日触发。另外一点，W的最近寻找不会跨过月份。

LW:这两个字符可以连用，表示在某个月最后一个工作日，即最后一个星期五。

## 7位

cron的表达式被用来配置CronTrigger实例。 cron的表达式是字符串，实际上是由七子表达式，描述个别细节的时间表。这些子表达式是分开的空白，代表：

1. Seconds

2. Minutes

3. Hours

4. Day-of-Month

5. Month

6. Day-of-Week

7. Year (可选字段)

例 "0 0 12 ? * WED" 在每星期三下午12:00 执行,

个别子表达式可以包含范围, 例如，在前面的例子里("WED")可以替换成 "MON-FRI", "MON, WED, FRI"甚至"MON-WED,SAT". “*” 代表整个时间段.
每一个字段都有一套可以指定有效值，如

Seconds (秒) ：可以用数字0－59 表示，

Minutes(分) ：可以用数字0－59 表示，

Hours(时) ：可以用数字0-23表示,

Day-of-Month(天) ：可以用数字1-31 中的任一一个值，但要注意一些特别的月份

Month(月) ：可以用0-11 或用字符串 “JAN, FEB, MAR, APR, MAY, JUN, JUL, AUG, SEP, OCT, NOV and DEC” 表示

Day-of-Week(每周)：可以用数字1-7表示（1 ＝ 星期日）或用字符口串“SUN, MON, TUE, WED, THU, FRI and SAT”表示

“/”：为特别单位，表示为“每”如“0/15”表示每隔15分钟执行一次,“0”表示为从“0”分开始, “3/20”表示表示每隔20分钟执行一次，“3”表示从第3分钟开始执行

“?”：表示每月的某一天，或第周的某一天

“L”：用于每月，或每周，表示为每月的最后一天，或每个月的最后星期几如“6L”表示“每月的最后一个星期五”

“W”：表示为最近工作日，如“15W”放在每月（day-of-month）字段上表示为“到本月15日最近的工作日”

““#”：是用来指定“的”每月第n个工作日,例 在每周（day-of-week）这个字段中内容为"6#3" or "FRI#3" 则表示“每月第三个星期五”

## 常用Cron表达式

0 15 10 * * ? *	每天10点15分触发
0 15 10 * * ? 2017	2017年每天10点15分触发
0 * 14 * * ?	每天下午的 2点到2点59分每分触发
0 0/5 14 * * ?	每天下午的 2点到2点59分(整点开始，每隔5分触发)
0 0/5 14,18 * * ?	每天下午的 2点到2点59分、18点到18点59分(整点开始，每隔5分触发)
0 0-5 14 * * ?	每天下午的 2点到2点05分每分触发
0 15 10 ? * 6L	每月最后一周的星期五的10点15分触发
0 15 10 ? * 6#3	每月的第三周的星期五开始触发