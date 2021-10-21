package main

import (
	"fmt"
	"regexp"
	"testing"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)

// go get -u all

func conv(html_str string) string {

	//opt := &md.Options{
	//	StrongDelimiter: "__", // default: **
	//}
	var converter = md.NewConverter("", true, nil)
	// newline := md.Rule{
	// 	Filter: []string{"br"}, // register <br>
	// 	Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
	// 		return md.String("\n") // return markdown
	// 	},
	// }

	converter.Use(plugin.GitHubFlavored())
	converter.Use(plugin.TaskListItems())
	converter.Use(plugin.Table())
	//converter.AddRules(newline)
	md_str, _ := converter.ConvertString(html_str)
	return md_str
}

func Test_md(t *testing.T) {
	//html := `<p>1. xxx <br/>2. xxxx<br/>3. xxx</p><p><span class="img-wrap"><img src="xxx"></span><br>4. golang<br>a. xx<br>b. xx</p> jmap –histo[:live]`
	html := `<code>last_30_days</code>`
	md_str := conv(html)
	t.Log(md_str)
}

func Test_request(t *testing.T) {
	doc := New("https://segmentfault.com/a/1190000014395186", "file", "selector").request()
	html_str, _ := doc.Html()
	md_str := conv(html_str)
	t.Log(md_str)
}

func Test_default(t *testing.T) {
	do("https://www.xx.com/a6746056891136213515", "", "")
	// html2md.exe https://www.xx.com/a6746056891136213515
}

func Test_regexp(t *testing.T) {
	// 特殊字符的查找
	reg := regexp.MustCompile(`[\[\]|/&\\]`)
	s := fmt.Sprintf("%v", reg.ReplaceAllString("[]|/&\\", "-"))
	t.Log(s)
}

// cnblogs_post_body

func Test_segmentfault(t *testing.T) {
	do("https://segmentfault.com/a/1190000014395186", "", "")
	// html2md.exe https://segmentfault.com/a/1190000014395186
}

func Test_cnblogs(t *testing.T) {
	do("https://www.cnblogs.com/coder-programming/p/12604799.html", "", "")
	// html2md.exe https://www.cnblogs.com/chenyishi/p/12586512.html
}

func Test_blog_csdn(t *testing.T) {
	do("https://blog.csdn.net/qq_36894974/article/details/103778956", "", "")
	// html2md.exe https://blog.csdn.net/qq_36894974/article/details/103778956
}

func Test_jianshu(t *testing.T) {
	do("https://www.jianshu.com/p/16719baa1713", "", "")
	// html2md.exe https://www.jianshu.com/p/16719baa1713
}

func Test_mpweixin(t *testing.T) {
	do("https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651437583&idx=2&sn=ddb41d563ce289aea4ef50e57fa58acd&chksm=80bb64fdb7ccedeb5ee8210371ae88ff871173160b75b55bb8878c3a30cf0935cbc1fad32ebb&scene=21#wechat_redirect", "", "")
	// html2md.exe https://mp.weixin.qq.com/s/UMv7C2KkEdKfEgIZxOksbg
}

func Test_oschina(t *testing.T) {
	do("https://www.oschina.net/news/114446/benchmark-openjdk-corretto-graalvm", "", "")
	// html2md.exe https://www.oschina.net/news/114446/benchmark-openjdk-corretto-graalvm
}

func Test_cloud_tencent(t *testing.T) {
	do("https://cloud.tencent.com/developer/article/1548200", "", "")
	// html2md.exe https://cloud.tencent.com/developer/article/1548200
}

func Test_zhuanlan_zhihu(t *testing.T) {
	do("https://zhuanlan.zhihu.com/p/100511637", "", "")
	// html2md.exe https://zhuanlan.zhihu.com/p/100511637
}

func Test_toutiao(t *testing.T) {
	do("https://www.toutiao.com/a6746056891136213515", "", "")
	// html2md.exe https://www.toutiao.com/a6746056891136213515
}
