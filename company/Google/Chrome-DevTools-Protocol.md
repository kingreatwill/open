# Chrome DevTools Protocol
https://chromedevtools.github.io/devtools-protocol/

chrome.exe --remote-debugging-port=9222
Chrome DevTools Protocol 是基于 WebScoket 协议的，当使用上面代码启动 chrome 后，我们可以在另一个页面输入 localhost:9222 来查看所有可以被 inspect 的页面。
Chrome 的 F12 想必大家已经再熟悉不过了，平常 debug 页面全靠它，然而它却是通过一个叫 Chrome DevTools Protocol 的协议来和目标页面通讯的。

Chrome DevTools Protocol (CDP)
Chrome DevTools Protocol (CDP) 的主页在：
https://chromedevtools.github.io/devtools-protocol/。它提供一系列的接口来查看，检查，调整并且检查 Chromium 的性能。Chrome 的开发者工具就是使用这一系列的接口，并且 Chrome 开发者工具来维护这些接口。

所谓 CDP 的协议，本质上是什么呢？本质上是基于 websocket 的一种协议

## 开发
https://github.com/ChromeDevTools/awesome-chrome-devtools#chrome-devtools-protocol