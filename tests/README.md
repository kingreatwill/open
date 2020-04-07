### selenium
https://github.com/SeleniumHQ/selenium
https://www.selenium.dev/documentation/en/

selenium 是一个 web 的自动化测试工具，不少学习功能自动化的同学开始首选 selenium ，因为它相比 QTP 有诸多有点：

- 免费，也不用再为破解 QTP 而大伤脑筋
- 小巧，对于不同的语言它只是一个包而已，而 QTP 需要下载安装1个多 G 的程序。
- 这也是最重要的一点，不管你以前更熟悉 C、 java、ruby、python、或都是 C# ，你都可以通过 selenium 完成自动化测试，而 QTP 只支持 VBS
- 支持多平台：windows、linux、MAC ，支持多浏览器：ie、ff、safari、opera、chrome
- 支持分布式测试用例的执行，可以把测试用例分布到不同的测试机器的执行，相当于分发机的功能。

### Pytest
pytest是一个非常成熟的全功能的Python测试框架，主要特点有以下几点：

1、简单灵活，容易上手，文档丰富；
2、支持参数化，可以细粒度地控制要测试的测试用例；
3、能够支持简单的单元测试和复杂的功能测试，还可以用来做selenium/appnium等自动化测试、接口自动化测试（pytest+requests）;
4、pytest具有很多第三方插件，并且可以自定义扩展，比较好用的如pytest-selenium（集成selenium）、pytest-html（完美html测试报告生成）、pytest-rerunfailures（失败case重复执行）、pytest-xdist（多CPU分发）等；
5、测试用例的skip和xfail处理；
6、可以很好的和CI工具结合，例如jenkins


### SeleniumBase
https://github.com/seleniumbase/SeleniumBase

Selenium & pytest