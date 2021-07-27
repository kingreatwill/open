# regular expression 正则表达式

## 工具
[图解正则表达式 - ok+](https://regexper.com/)
https://gitlab.com/javallone/regexper-static

[图解正则 - ok++](https://jex.im/regulex/#!flags=&re=%5E(a%7Cb)*%3F%24)
https://github.com/CJex/regulex


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
