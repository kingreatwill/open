package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"github.com/PuerkitoBio/goquery"
	"github.com/mxschmitt/playwright-go"
)

type suport struct {
	url               string
	file              string
	selector          string
	useHeadlessChrome bool
	mdConverter       *md.Converter
}

// .gitignore忽略文件把我害惨了~~
func New(url_str, file, selector string) *suport {
	var converter = md.NewConverter("", true, nil)
	converter.Use(plugin.GitHubFlavored())
	converter.Use(plugin.TaskListItems())
	converter.Use(plugin.Table())
	suport := &suport{
		url:               url_str,
		file:              file,
		selector:          selector,
		useHeadlessChrome: true,
		mdConverter:       converter,
	}
	suport.parse()
	return suport
}

func (s *suport) Do() {
	doc := s.request()
	html_str := s.saveImg(doc)
	md_str := s.convert2md(html_str)
	s.saveFile(fmt.Sprintf("%v.md", s.file), []byte(md_str))
}

func (s *suport) parse() {
	u, err := url.Parse(s.url)
	if err != nil {
		log.Fatal(err)
	}
	switch strings.ToLower(u.Host) {
	case "www.cnblogs.com", "cnblogs.com":
		s.selector = "#cnblogs_post_body"
		s.useHeadlessChrome = false
	}

}

// 1.
// go get github.com/mxschmitt/playwright-go

// 2.
//go run github.com/mxschmitt/playwright-go/cmd/playwright install --with-deps
//# Or
//go install github.com/mxschmitt/playwright-go/cmd/playwright
//playwright install --with-deps

//  Set-ExecutionPolicy RemoteSigned

func (s *suport) request() *goquery.Document {
	if !s.useHeadlessChrome {
		res, err := http.Get(s.url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatalf("goquery NewDocument: %v", err)
		}
		if s.file == "" {
			s.file = doc.Find("title").First().Text()
		}
		return doc
	}
	// 使用浏览器;
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	response, err := page.Goto(s.url)
	if err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	bodys, err := response.Body()
	//htmlContent, err := page.InnerHTML("body") // page.TextContent("body")
	if err != nil {
		log.Fatalf("could not get text content: %v", err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodys))
	if err != nil {
		log.Fatalf("goquery NewDocument: %v", err)
	}
	if s.file == "" {
		s.file = doc.Find("title").First().Text()
	}
	return doc
	// entries, err := page.QuerySelectorAll(".athing")
	// if err != nil {
	// 	log.Fatalf("could not get entries: %v", err)
	// }
	// for i, entry := range entries {
	// 	titleElement, err := entry.QuerySelector("td.title > a")
	// 	if err != nil {
	// 		log.Fatalf("could not get title element: %v", err)
	// 	}
	// 	title, err := titleElement.TextContent()
	// 	if err != nil {
	// 		log.Fatalf("could not get text content: %v", err)
	// 	}
	// 	fmt.Printf("%d: %s\n", i+1, title)
	// }
	// if err = browser.Close(); err != nil {
	// 	log.Fatalf("could not close browser: %v", err)
	// }
	// if err = pw.Stop(); err != nil {
	// 	log.Fatalf("could not stop Playwright: %v", err)
	// }
}

func (s *suport) convert2md(html_str string) string {
	md_str, _ := s.mdConverter.ConvertString(html_str)
	return md_str
}

func (s *suport) saveImg(doc *goquery.Document) string {
	reg := regexp.MustCompile(`[\s\[\]|/&?？\\]`)
	s.file = reg.ReplaceAllString(s.file, "")

	content_selection := doc.Find(s.selector).First()
	content_selection.Find("img").Each(func(i int, img_selection *goquery.Selection) {
		if img_src, exists := img_selection.Attr("src"); exists {
			ss := strings.Split(strings.Split(img_src, "?")[0], ".")
			// 下载图片;
			resp, _ := http.Get(img_src)
			body, _ := ioutil.ReadAll(resp.Body)
			imgName := fmt.Sprintf("img/%v-%v.%v", s.file, i+1, ss[len(ss)-1])
			s.saveFile(imgName, body)
			// 更改属性;
			img_selection.SetAttr("src", imgName)
		}
	})

	html_str, _ := content_selection.Html()
	return html_str
}

func (s *suport) saveFile(fileName string, fileContent []byte) {
	os.MkdirAll("img", os.ModePerm)
	out, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Create: %v", err)
	}
	_, err = io.Copy(out, bytes.NewReader(fileContent))
	if err != nil {
		log.Fatalf("Create: %v", err)
	}
}
