package main

import (
	"testing"
)

func Test_segmentfault(t *testing.T) {
	do("https://segmentfault.com/a/1190000014395186", "", "")
	// html2md.exe https://segmentfault.com/a/1190000014395186
}

func Test_cnblogs(t *testing.T) {
	do("https://www.cnblogs.com/chenyishi/p/12586512.html", "", "")
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