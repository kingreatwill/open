<!-- toc -->
[TOC]

# 在线文档管理系统
- js gitbook
- go MinDoc
- go BookStack


## 安装gitbook
```
npm install gitbook-cli -g

gitbook -V
gitbook --version

# 其实就是生成 SUMMARY.md和README.md文件
gitbook init

# http://localhost:4000
gitbook serve
#运行该命令后会在书籍的文件夹中生成一个 _book 文件夹, 里面的内容即为生成的 html 文件，我们可以使用下面命令来生成网页而不开启服务器。
gitbook build 不开启服务器
```

## 构建书籍 gitbook build
默认：将生成的静态网站输出到 _book 目录
指定路径：gitbook build [书籍路径] [输出路径]
Debugging: gitbook build ./ --log=debug --debug
指定端口：gitbook serve --port 2333
生成pdf格式：gitbook pdf ./ ./mybook.pdf
生成epub格式：gitbook epub ./ ./mybook.epub
生成 mobi 格式：gitbook mobi ./ ./mybook.mobi


## book.json

```json
{
    "title": "Blankj's Glory",
    "author": "Blankj",
    "description": "select * from learn",
    "language": "zh-hans",
    "gitbook": "3.2.3",
    "styles": {
        "website": "./styles/website.css"
    },
    "structure": {
        "readme": "README.md"
    },
    "links": {
        "sidebar": {
            "我的狗窝": "https://blankj.com"
        }
    },
    "plugins": [
        "-sharing",
        "splitter",
        "expandable-chapters-small",
        "anchors",
        "github",
        "github-buttons",
        "donate",
        "sharing-plus",
        "anchor-navigation-ex",
        "favicon"
    ],
    "pluginsConfig": {
        "github": {
            "url": "https://github.com/Blankj"
        },
        "github-buttons": {
            "buttons": [{
                "user": "Blankj",
                "repo": "glory",
                "type": "star",
                "size": "small",
                "count": true
                }
            ]
        },
        "donate": {
            "alipay": "./source/images/donate.png",
            "title": "",
            "button": "赞赏",
            "alipayText": " "
        },
        "sharing": {
            "douban": false,
            "facebook": false,
            "google": false,
            "hatenaBookmark": false,
            "instapaper": false,
            "line": false,
            "linkedin": false,
            "messenger": false,
            "pocket": false,
            "qq": false,
            "qzone": false,
            "stumbleupon": false,
            "twitter": false,
            "viber": false,
            "vk": false,
            "weibo": false,
            "whatsapp": false,
            "all": [
                "google", "facebook", "weibo", "twitter",
                "qq", "qzone", "linkedin", "pocket"
            ]
        },
        "anchor-navigation-ex": {
            "showLevel": false
        },
        "favicon":{
            "shortcut": "./source/images/favicon.jpg",
            "bookmark": "./source/images/favicon.jpg",
            "appleTouch": "./source/images/apple-touch-icon.jpg",
            "appleTouchMore": {
                "120x120": "./source/images/apple-touch-icon.jpg",
                "180x180": "./source/images/apple-touch-icon.jpg"
            }
        }
    }
}
```


## 插件
"livereload","highlight","search","lunr","sharing","fontsettings","theme-default"
如果要去除自带的插件，可以在插件名称前面加 -，比如：
```
"plugins": [
    "-search"
]
```
如果要指定插件的版本可以使用 plugin@0.3.1，因为一些插件可能不会随着 GitBook 版本的升级而升级。

### 插件安装 
在book.json中添加
```
{
    "plugins": ["pluginName"]
}
```
在book.json所在目录运行gitbook install or gitbook install ./,安装插件

#### 安装全局包
也可以
npm i gitbook-plugin-toc -g

### 不支持markdown的TOC命令自动生成目录
安装插件[gitbook-plugin-toc](https://www.npmjs.com/package/gitbook-plugin-toc)
在book.json中添加
```
{
    "plugins": ["toc"]
}
```
在使用TOC命令的地方使用`<!-- toc -->`代替。即可自动生成文档目录。

### page-footer插件
```
{
    "plugins": [
        "page-footer"
    ],
    "pluginsConfig": {
        "page-footer": {
            "description": "modified at",
            "signature": "Aleen",
            "wisdom": "More than a coder, more than a designer",
            "format": "yyyy-MM-dd hh:mm:ss",
            "copyright": "Copyright &#169; aleen42",
            "timeColor": "#666",
            "copyrightColor": "#666",
            "utcOffset": "8",
            "isShowQRCode": true,
            "baseUri": "https://aleen42.gitbooks.io/personalwiki/content/",
                        "isShowIssues": true,
            "repo": "aleen42/PersonalWiki",
            "issueNum": "8",
            "token": "",
                        "style": "normal"
        }
    }
}
```

### github插件
```
{
    "plugins": [ "github" ],
    "pluginsConfig": {
        "github": {
            "url": "https://github.com/your/repo"
        }
    }
}
```
