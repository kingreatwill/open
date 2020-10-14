# regular expression 正则表达式

[正则表达式30分钟入门教程](https://deerchao.cn/tutorials/regex/regex.htm)
[图解正则表达式 - ok+](https://regexper.com/)
https://gitlab.com/javallone/regexper-static

[图解正则 - ok++](https://jex.im/regulex/#!flags=&re=%5E(a%7Cb)*%3F%24)
https://github.com/CJex/regulex


[正则表达式工具 - ok+](http://tool.chinaz.com/regex/)


[正则表达式工具 - good](https://regexr.com/)
https://github.com/gskinner/regexr/

[regex101 - good++](https://regex101.com/)
https://github.com/firasdib/Regex101


[在线正则表达式](http://www.regexe.com/)

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

[不同语言之间的差别](https://www.regular-expressions.info/refcharacters.html)

[.NET Regex - 不更新了](http://regexstorm.net/)
https://github.com/lonekorean/regex-storm


[学习正则表达式](https://github.com/ziishaned/learn-regex)

[工具和教程](http://www.regexlab.com/zh/mtracer/download.htm)

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


[Regex 正则表达式入门](https://www.cnblogs.com/codeshell/p/12825243.html)

vs code 插件any-rule   F1
