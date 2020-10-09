<!-- toc -->
[TOC]

# 在线文档管理系统
- js gitbook
- js Gatsby https://github.com/gatsbyjs/gatsby 47.3k
- ts Docusaurus https://github.com/facebook/docusaurus 19.4k
- ruby Jekyll https://github.com/jekyll/jekyll 41.4k
- js VuePress https://github.com/vuejs/vuepress 17.5k
- go MinDoc
- go BookStack
- mdbook https://github.com/rust-lang/mdBook 4.5k
- peach https://github.com/peachdocs/peach 1.1k 
- gohugo+git book theme  https://github.com/gohugoio/hugo 45.1k
https://github.com/alex-shpak/hugo-book

•mdbook是“门槛”最低的，几乎无需配置，就能搭建出一个像模像样的类似gitbook的图书站点；
•peach门槛较高一些，要配置的东西多一些；
•gohugo门槛适中，但却最为灵活和强大。如果对gohugo的模板语法十分熟悉，可以定义出一套满足自己风格的电子书浏览页面风格。


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
            "shortcut": "./img/logo/logo.jpg",
            "bookmark": "./img/logo/logo.jpg",
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
```json
"plugins": [
    "-search"
]
```
如果要指定插件的版本可以使用 plugin@0.3.1，因为一些插件可能不会随着 GitBook 版本的升级而升级。

目前 Gitbook 官方已不再为维护插件网站,只能通过 npmjs 发现 Gitbook 插件.
https://www.npmjs.com/package/
gitbook-plugin


### 插件安装 
在book.json中添加
```json
{
    "plugins": ["pluginName"]
}
```
在book.json所在目录运行gitbook install or gitbook install ./,安装插件

一种是使用npm install pluginName安装，然后写入配置
npm i gitbook-plugin-toc

### 不支持markdown的TOC命令自动生成目录
安装插件[gitbook-plugin-toc](https://www.npmjs.com/package/gitbook-plugin-toc)
在book.json中添加
```json
{
    "plugins": ["toc"]
}
```
在使用TOC命令的地方使用`<!-- toc -->`代替。即可自动生成文档目录。

### page-footer插件
```json
{
    "plugins": [
        "page-footer"
    ],
    "pluginsConfig": {
        "page-footer": {
            "description": "modified at",
            "signature": "Kingreatwill",
            "wisdom": "More than a coder, more than a designer",
            "format": "yyyy-MM-dd hh:mm:ss",
            "copyright": "Copyright &#169; kingreatwill",
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
```json
{
    "plugins": [ "github" ],
    "pluginsConfig": {
        "github": {
            "url": "https://github.com/your/repo"
        }
    }
}
```

### hide-element 隐藏元素
如：默认的gitbook左侧提示：Published with GitBook
```json
{
    "plugins": [
        "hide-element"
    ],
    "pluginsConfig": {
        "hide-element": {
            "elements": [".gitbook-link"]
        }
    }
}
```

### back-to-top-button 回到顶部(anchor-navigation-ex 也可以)
```json
{
    "plugins": [
         "back-to-top-button"
    ]
}
```

### chapter-fold 导航目录折叠
```json
{
    "plugins": ["chapter-fold"]
}
```
### code 复制代码
```json
{
    "plugins" : [ "code" ]
}
```

### splitter 侧边栏宽度可调节
左侧目录和右侧文章可以拖动调节宽度。
```json
{
    "plugins": [
        "splitter"
    ]
}
```

### search-pro 高级搜索
支持中英文，准确率更高一些。
```json
{
    "plugins": [
          "-lunr", 
          "-search", 
          "search-pro"
    ]
}
```

### insert-logo 插入logo
在左侧导航栏上方插入logo。
```json
{
    "plugins": [ "insert-logo" ]
    "pluginsConfig": {
      "insert-logo": {
        "url": "images/logo.png",
        "style": "background: none; max-height: 30px; min-height: 30px"
      }
    }
}
```

### custom-favicon 修改标题栏图标
```json
{
    "plugins" : ["custom-favicon"],
    "pluginsConfig" : {
        "favicon": "icon/favicon.ico"
    }
}
```

### pageview-count 阅读量计数

```json
{
  "plugins": [ "pageview-count"]
}
```

### tbfed-pagefooter 页面添加页脚
在每个文章下面标注版权信息和文章时间。
```json
{
    "plugins": [
       "tbfed-pagefooter"
    ],
    "pluginsConfig": {
        "tbfed-pagefooter": {
            "copyright":"Copyright &copy dsx2016.com 2019",
            "modify_label": "该文章修订时间：",
            "modify_format": "YYYY-MM-DD HH:mm:ss"
        }
    }
}
```

### popup 弹出大图
```json
{
  "plugins": [ "popup" ]
}
```

### sharing-plus 分享当前页面
gitbook默认只有Facebook、Google+、Twiter、Weibo、Instapaper
```json
{
    "plugins": ["-sharing", "sharing-plus"],
    "pluginsConfig": {
        "sharing": {
             "douban": true,
             "facebook": true,
             "google": true,
             "pocket": true,
             "qq": true,
             "qzone": true,
             "twitter": true,
             "weibo": true,
          "all": [
               "douban", "facebook", "google", "instapaper", "linkedin","twitter", "weibo", 
               "messenger","qq", "qzone","viber","whatsapp"
           ]
       }
    }
}
```

### book.json
```json
{
    "plugins": [
        "back-to-top-button",
        "chapter-fold",
        "code",
        "splitter",
        "-lunr",
        "-search",
        "search-pro",
        "insert-logo",
        "custom-favicon",
        "pageview-count",
        "tbfed-pagefooter",
        "popup",
        "-sharing",
        "sharing-plus"
    ],
    "pluginsConfig": {
        "insert-logo": {
            "url": "https://file.smallzhiyun.com/%E4%B9%A6.png",
            "style": "background: none; max-height: 30px; min-height: 30px"
        },
        "favicon": "./icon/book.ico",
        "tbfed-pagefooter": {
            "copyright": "Copyright &copy dsx2016.com 2019",
            "modify_label": "该文章修订时间：",
            "modify_format": "YYYY-MM-DD HH:mm:ss"
        },
        "sharing": {
            "douban": true,
            "facebook": true,
            "google": true,
            "pocket": true,
            "qq": true,
            "qzone": true,
            "twitter": true,
            "weibo": true,
            "all": [
                "douban",
                "facebook",
                "google",
                "instapaper",
                "linkedin",
                "twitter",
                "weibo",
                "messenger",
                "qq",
                "qzone",
                "viber",
                "whatsapp"
            ]
        }
    }
}
```

### gitalk 评论插件
### mermaid 使用流程图。


## 插件开发
GitBook 插件是在 npm 上发布的遵循传统定义的 node 包,除了标准的 node 规范外还有一些 Gitbook 自身定义的相关规范.

### 目录结构
Gitbook 插件最基本的项目结构至少包括配置文件 package.json 和入口文件 index.js ,其他目录文件根据插件用途自行增减.
```
.
├── index.js
└── package.json
```
> 实际插件项目略有不同,可能还会有 _layouts 布局目录, asset 资源目录以及自定义 example 示例目录和 docs 文档目录等等.

### package.json
package.json 是**nodejs**的配置文件,Gitbook 插件同样遵循该规范,配置文件声明了插件的版本描述性信息,除此之外还有 Gitbook 相关字段,遵循schema准则,基本示例如下:
```json
{
    "name": "gitbook-plugin-mytest",
    "version": "0.0.1",
    "description": "This is my first GitBook plugin",
    "engines": {
        "gitbook": ">1.x.x"
    },
    "gitbook": {
        "properties": {
            "myConfigKey": {
                "type": "string",
                "default": "it's the default value",
                "description": "It defines my awesome config!"
            }
        }
    }
}

```
> 值得注意的是,包名称必须以 gitbook-plugin-开头，包引擎应该包含gitbook.如需了解 package.json 的规范,可参考[官方文档](https://docs.npmjs.com/files/package.json)

### index.js
`index.js` 是插件运行时的入口,基本示例如下:
```
module.exports = {
    // 钩子函数
    hooks: {},

    // 代码块
    blocks: {},

    // 过滤器
    filters: {}
};
```
### 发布插件
GitBook 插件可以在[npmjs官网](https://www.npmjs.com/)上发布.

如需发布插件,首先需要在npmjs官网上注册帐户,然后通过命令行发布.
```
$ npm publish
```
### 专用插件
专用插件可以托管在 GitHub 上,并使用 git urls:
```
{
    "plugins": [
        "myplugin@git+https://github.com/MyCompany/mygitbookplugin.git#1.0.0"
    ]
}
```

### 本地测试插件
使用 npm link 可以在发布之前测试你的插件,命令详情参考[官方文档](https://docs.npmjs.com/cli/link)

在插件的文件夹中,运行：
```
$ npm link
```
然后在您的书或者文档的文件夹中执行:
```
$ npm link gitbook-plugin-<name>
```
### 单元测试插件
[gitbook-tester](https://github.com/todvora/gitbook-tester)可以方便地为你的插件编写`Node.js/Mocha`单元测试.

使用Travis.可以对每个提交/标签运行测试.

