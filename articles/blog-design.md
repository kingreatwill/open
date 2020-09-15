<!--toc-->
[TOC]
#  汪宇杰博客 博客系统知多少：揭秘那些不为人知的学问
```
 博客基本功能设计要点

    4.1 文章（Post）

    4.2 评论（Comment）

    4.3 分类（Category）

    4.4 标签（Tag）

    4.5 归档（Archive）

    4.6 页面（Page）

    4.7 订阅

    4.8 版本控制

    4.9 主题及个性化

    4.10 用户及权限

    4.11 插件

    4.12 图片及附件的处理

    4.13 脏词过滤及评论审查

    4.14 静态化

    4.15 通知系统

5. 博客协议或标准

    5.1 RSS

    5.2 ATOM

    5.3 OPML

    5.4 APML

    5.5 FOAF

    5.6 BlogML

    5.7 Open Search

    5.8 Pingback

    5.9 Trackback

    5.10 MetaWeblog

    5.11 RSD

    5.12 阅读器视图

6. 设计博客系统有哪些知识点

    6.1 时区真的全用UTC？

    6.2 HTML还是Markdown

    6.3 MVC还是SPA

    6.4 安全
SNS 社交服务（相册，朋友圈）
```
## 源码参考
https://github.com/EdiWang/Moonglade
## “博客”的前世今生
博客一开始不叫 Blog，而叫 Weblog，可能让很多人诧异的是，它并不诞生于Web 2.0时代，而是早在1997年已经问世。博客从最早的单用户（单独作者），逐渐发展为多用户（一个团队），即博客平台。而 Web 2.0 时代赋予了博客社交属性，可以让读者进行评论、订阅（RSS/ATOM），博客作者之间可以互相抱团（FOAF）、引用文章（Pingback），才让博客逐渐热门了起来。

博客系统也是各有千秋，PHP 有 WordPress，.NET 有BlogEngine。而最终，WordPress 几乎成为了事实上的博客系统的标准，它同时具备一些 CMS 的功能，微软官方 .NET 团队的博客也是采用 WordPress 搭建。

阅读博客的用户除了使用浏览器，还会使用 RSS/Atom 阅读器。在 iPad 刚出来的年代，阅读器应用曾经风靡一时。订阅的博客一旦有新文章，阅读器就会自动收入，读者无需每天人肉检查是否有新文章发布。Microsoft 365 的 Outlook 至今保留着 RSS 阅读器的功能。

博客至今依然是表达自我、传播信息并与社区互动的最佳途径之一，就算微博（microblogging）出现，也没能使博客变得不再流行。所以说，博客之于互联网，就如同电子邮件一样，“姜还是老的辣”，博客作为一种文化载体，历久弥新，持久散发着光芒。

## 博客基本功能设计要点

### 文章（Post）
其实博客类型的文章，正确的表达是post，英文单词里post和article的区别在于，post只是随心所欲写的文章，而article指的是论文那样的经过精心雕琢，旁征博引，并且有可能在学术期刊上发表的文章。因此设计博客系统的时候，尽量避免使用article这个单词来命名代码。更具体来讲，post中可以出现不严谨、口语化的表达方式，例如本文就算是个post。而article讲究言语上的规范，连 “让我们” ， “我们来看看” 这种字眼都不能出现。

文章需要具备`标题、Slug、创建时间、发布时间、修改时间、摘要和内容`等要素，也会包含所属`分类、标签、阅读量和点赞量`等次要信息。其中Slug是博客的特色，它指的是一篇文章的URL。例如我的文章：《Try the New Azure .NET SDK》，它的URL为 `https://edi.wang/post/2019/12/5/try-the-new-azure-net-sdk`，其中try-the-new-azure-net-sdk即为该文章的Slug。Slug讲究的是“人类可读”，一般情况下均为博客标题对应的英文表达，用中划线分割英文单词，Slug也对博客的SEO起到了关键作用。如果你的博客文章用的是数据库ID、文章标题的HTML Encoding等做URL，请更换为Slug。特别是遇到中文文章，如果标题被URL Encoding了，那么对于SEO和链接分享，都是灾难。一个Slug一旦定下，尽量不要改动，虽然大部分博客系统都支持修改Slug，但是对于被搜索引擎收入的文章，改了Slug就会导致404。比较完备的博客系统（如WordPress）支持采用301重定向方式告诉搜索引擎原文地址已变化。

摘要有两个作用，一是用于在列表视图中显示文章信息预览，二是用于SEO，放在description这个meta标签中，可以帮助搜索引擎精准定收入的内容。对于中文内容，需要注意是否输出的HTML源代码被Encoding过，ASP.NET Core默认的Encoding会对SEO造成灾难（我的博客系统因为面向英语用户，不考虑中文支持，所以并不解决这个问题）。

摘要可以自动抓取文章前几百字，也可以像微信公众号那样要求用户手工填写。我的博客采用的是自动取文章前400字。结合SEO的关系，我的文章通常开头段落就是概要，这样可以让用户在搜索引擎预览页面就能看到准确内容，而不是页面上无关紧要的UI元素。

文章的状态通常包括：草稿、发布、回收。用户仅能看到已发布的文章，管理员可在后台更改文章状态。


### 评论（Comment）
评论是博客中作者和读者互动的主要方式。有些博客要求读者登录后才能发表评论，而有些可以允许游客评论（比如我的博客及WordPress）。登录的好处在于可以识别你的读者，并有效防止垃广告评论。但要求登录也会给用户造成操作上多了一个步骤，嫌麻烦的用户就不会进行评论。

我的博客及WordPress默认都设计为需要管理员在后台审核评论后，才能放出显示。这也能有效避免垃圾广告、骚扰信息甚至是一些恶意的煽动性言论。对于提供Email地址的用户，管理员也能够在后台回复用户的评论，并由博客系统向用户发送Email通知。

对于技术博客，评论可考虑开放markdown格式。这是一种在程序员之间格外受欢迎的语法，在GitHub得到了广泛应用。

评论需要采用验证码或其他人机验证技术，以防止机器人发广告。但根据经验，验证码并不能100%阻止垃圾信息，因为现代化的垃圾信息还真的是人组团发的，有专门的公司、团队、微信群等，国外也有。于是，你可能需要考虑关键词过滤，购买三方过滤接口等。

评论也得记得做字数限制，不然也有可能会造成部分用户“灌水”、刷屏的现象。

如果不想自己写功能，还可以整合三方的评论服务，即博客系统本身不实现评论功能，通过三方服务加载外部JS，在文章阅读页面“注入”一个评论区，通常这要求文章的URL不变（WordPress里叫做永久性URL）。


### 分类（Category）

像建文件夹一样将文章根据内容进行区分，即为分类。文章分类后，可以帮助读者快速检索同种类的文章。

例如写.NET、PHP、JS的文章都属于 “开发技术（Development）”这个分类。而技术圈新闻和职场经验分享和等文章，则属于 “工作” 分类，分类的划分完全由用户控制。分类可以多对多。例如写一篇文章介绍了用ASP.NET Core开发Angular应用的文章，可以同时属于 “.NET技术” 及 “前端开发” 分类。

分类需要一个标题、一个简介，以及一个路由名称。例如我博客，微软云 Azure 的分类，标题为 Microsoft Azure，简介为  The Best Cloud，路由名称为azure。标题需要同时显示在标题栏，以便于SEO。简介是对于标题的补充说明，便于用户查看。设计路由名称的原因请参考下一段介绍的标签的设计。

分类的另一个功能就是产生 OPML 及 RSS/Atom 订阅源


### 标签（Tag）
一篇文章所提到的话题，即为文章的标签。和分类一样，标签也是多对多关系。标签可以作为检索文章的依据，类似关键词，快速查找相关内容的文章。

标签需要考虑到标签含义重复的情况，例如：VS和Visual Studio是一个意思，VSCode、VSC和Visual Studio Code也是一个意思。那么用户选择标签的时候，最好使用智能提示推荐用户使用已有标签。

对于博客系统设计者来说，还要考虑标签的URL。如果URL用的是标签本身的内容，会导致很多问题。当标签名称为整个英文单词，例如Excel，并不会发生问题，因为URL通常是https://yourblog/tags/excel。但是如果标签内容为 .NET Core、C#、Robots.txt，事情就变的有意思起来。https://yourblog/tags/robots.txt到底是在请求 tags下的robots.txt 文件还是在请求标签？自己作为博客系统设计者，当然可以从程序上限制所有tags接受的路由参数都为标签，好像是解决了问题，但SEO和扫描工具可不这么认为，他们有大量by convention的规则会认为是请求文件。

对于需要URL Encoding的标签内容，更会导致URL缺乏可读性，从而影响SEO。千万不要自作聪明地以为现代的搜索引擎可以处理好URL Encoding，一个URL是否干净对SEO的影响很大。特别是当标签是中文内容的时候，如果全encoding了，URL就会非常冗长，甚至影响到SEO，也影响到博主分享链接。因此，我的博客系统为了处理标签URL，给每个标签都设计了规范化名称（normalized name），由系统根据标签内容自动产生，如 .NET Core 经过 normalize，会变成 dotnet-core，最终产生的URL即 https://edi.wang/tags/list/dotnet-core。


对于用户来说，最容易犯错之一的就是把标签用成搜索关键词。例如用户写了一篇关于 Visual Studio Code的文章，那么标签可能会同时打VSCode、VSC和Visual Studio Code，但其实只要选择一个标签即可。打太多同样含义的标签会导致读者无法完整检索到所有相关文章，对搜索引擎来说，也是如此。所以如何用好标签，是博客设计者和用户需要共同关注的要点。

标签云（Tag Cloud）是博客中用来列出最热门标签的功能。通常使用跟大号字、更明显的颜色来标识出对应文章较多的标签。标签云可以作为博客博主的个性化属性，一眼就能看出博主热衷于什么话题（比如Windows Phone？0.0）。

### 归档（Archive）
以时间（年、月、日）整理的博客文章即为归档，它和分类的区别在于归档只以时间为标准来划分文章。Archive的SEO相对于文章、分类、标签来说，并不那么关键。所以除了URL可以按年月划分以外，并没有额外的讲究。

例如：https://edi.wang/archive/2019/9 表示2019年9月的文章。https://edi.wang/archive/2019 则表示2019年所有的文章。归档功能主要用于给读者按时间查询，看看博主某个时间都在干什么。设计这样的功能可以提高读者对博主的兴趣，也是个人对外形象的一种展示。

### 页面（Page）

页面是博客的可选功能之一，事实上，它更接近于CMS的功能。有些内容并不适合以文章的形式发布，比如“关于”页面。这样的页面通常与发布时候的时间无关，内容也经常更新，排版设计也非常自由，不单纯是文字。

页面通常不需要评论、标签和分类等属性，但可以有发布和编辑时间。和文章一样，页面也需要注意Slug。

在我的博客系统中，页面也选择是否隐藏侧边栏，用户也可以完全编写页面的HTML及CSS代码，并把页面添加为导航菜单。WordPress对于页面的处理更加完备，接近于CMS系统。


### 订阅（Subscription）

读者订阅博客的主要方式有Feed（RSS/ATOM）及Newsletter。Feed方式本质上是被动订阅，需要客户端软件发起请求给服务器，检查是否有新文章发表，才能显示到客户端里。Newsletter一般采用Email形式主动发送给订阅用户，但这要求博客系统的编写者实现Email订阅功能，也要求管理员维护Email服务。订阅一般只推送近期发表的新文章，例如前10、20篇，而不会每次都推送全部文章导致客户端爆炸。

订阅一般可按文章分类提供，以便于只对某些分类感兴趣的读者阅读。有些博客系统也提供文章评论的订阅源，以便读者观摩吐槽大会。

关于RSS及ATOM的详细介绍请看5.1、5.2章节。

### 版本控制

更接近CMS的博客系统通常提供版本控制功能，允许用户回滚文章或页面的历史版本。设计版本控制的时候，不能只考虑往前回滚，得还能再滚得回来。通常，用户每次编辑一篇已经写好的文章，就会产生一个新版本，类似于git对于一个文件的commit。博客的版本控制也类似于代码版本控制，你可以选择保存一篇文章的完整内容作为历史版本，也可以选择每次只保存变化量信息（delta）。保存完整内容不容易后续花费大量时间精力 ，但是会占用较多存储空间。保存内容变化量节省数据库空间，但实现代码容易占用大量精力。


### 主题及个性化

好用的博客系统通常支持主题，毕竟个性化是博客本身应有的特点之一。WordPress积累了大量的主题库，也允许自制主题。但是我的博客只支持更改主题色，还有很大上升空间。

### 用户及权限

博客系统分为个人、团队及博客平台。个人博客系统一般为单用户（例如我的博客），不需要设计权限、注册等功能。多用户博客则需要实现不同的角色和权限，比如博客管理员、审核专员、撰稿人、评论管理员等等。无论是单用户还是多用户博客，集成一套成熟的三方RBAC方案可能是最高效的选择，多数三方方案也都支持SSO，例如我博客支持的Azure AD。


### 插件

插件功能可以在不更改博客代码的情况下，按需拓展博客的功能。WordPress以及BlogEngine都支持插件，但Moonglade还不行。

### 图片及附件的处理

- 图片格式


在2020年，图片格式已经非常自由，一般的博客JPG居多，程序员的博客PNG居多（毕竟都是屏幕截图），像微信公众号那样采用WEBP格式现在同样可取，只要读者的设备兼容即可。一般BMP格式由于体积大会导致网络传输慢，所以不推荐。同样道理，GIF也要注意限制尺寸。

博客系统输出图片时，需要采用正确的Mime Type，以保证客户端的兼容性。一般直接输出静态文件本身不需要博客编写者手工处理Mime Type，但有专门图片处理逻辑的博客（例如我的Moonglade）则需要留意保留图片原本的Mime Type。

- 图片水印


给上传的图片自动加水印有助于保护版权，水印内容一般是博客的地址或博主名字。添加水印时要注意图片尺寸调整水印的比例，以免挡住图中重要内容影响阅读。对于过小的图片，可选择性的忽略水印。

另外，考虑到博客有可能会在发展过程中改名，建议添加水印的时候在系统中保留一份原始图片，以便于后期更新水印内容。

具体方法可参考我的文章《ASP.NET Core 给上传的图片加水印》。

- 图片存储


图片存哪里是个值得思考的问题。一般有3个地方存放：文件系统、数据库、云上的Blob存储服务。Moonglade支持文件系统及Azure Blob存储。这三者各有优缺点。

文件系统的优势在于直接serve static file速度最快，但如果图片目录本身位于网站目录底下，会导致目录不只读而引起潜在安全问题。比如初中时候很流行的给DVBBS上传个改了拓展名的ASP web shell，尽管给web服务器上传可执行文件在2020年已经基本绝迹了，但依然存在隐患，就好比就算你家里请了007当保镖也是需要夜间锁好门。

数据库存图安全性最高，并且让博客的数据只位于一个位置，方便管理和备份，十几年前很流行这么做，但其实读写图片对数据库有一定开销，并且再由网站输出，双倍开销，一般不推荐。

而云端Blob存储服务目前来说是最适合这个时代的方案，将图片存储在Blob中不仅能保证服务器目录只读，又能采用云本身的安全特性限制非正常访问，还能通过CDN加速图片输出。要硬说缺点，就是云服务需要额外的金钱，而没钱，是自己的问题，不是云的问题。

- 图片防盗链


作为网站开发者，我们有时候不希望自己网站的图片被其他网站直接引用。这在某些场景下会导致自己数据中心里巨大的带宽消耗，也就意味着别人使用我们的图片，我们要为此付钱。例如，你的网站是a.com，你有一张图片是http://a.com/facepalm.jpg，而b.com在他们的网站上使用一个img标签来引用了你的图片，这导致网络请求是进入你的数据中心，消耗你的资源。因此博客可选择性的启用防盗链功能，具体方法可参考我的文章《ASP.NET / Core 网站图片防盗链》。

- 附 件


通常程序员的技术博客会提供读者下载代码样例等附件。设计附件功能和设计图片存储非常类似，完全可行。但我更建议技术博客将代码示例等附件托管到其他网站（例如GitHub）提供读者下载。

自己博客实现附件下载的坏处有：

- 大文件

不同的Web服务器及防火墙产品对文件尺寸的限制不同，而部署博客的用户很可能无权管理这些限制，就会导致大附件无法提供下载。

- 域及IP黑名单

某些公司或组织（特别是安全规范较高的软件公司）会屏蔽非白名单域的文件下载，尽管你可以用浏览器正常打开该域的网页，但无法下载文件（防火墙只允许HTML/CSS/JS等，而不允许ZIP、EXE等）。而程序员博客的读者很有可能就处在这样的公司里。

- CDN资源耗费

如果你的附件较多，较大，并且你也像设计图片存储一样给附件系统套了个CDN，此时根据CDN服务商计费模式的不同，如果按流量计费，恐怕你的附件下载会导致你的钱包加速瘦身。

而采用三方文件下载（如GitHub、OneDrive）的好处有：

√ 你的文件不仅可以分享到博客文章中，也可以分享到别的位置；

√ 这些三方服务有自己的CDN，而不用担心消耗你自己的钱包；

√ 许多文件托管服务有完整的管理功能，例如文件删除、恢复、版本控制、权限等，要是自己在博客系统里写一个这个，需要花费大量时间……


### 敏感词过滤及评论审查

博客难免引来一些抱有敌意的人，也会引来发广告的人，所以通常需要敏感词过滤和评论审查。如果没有审查直接将用户的评论显示在文章下，那么可能会对博主和网站本身带来不良影响。例如，有人发了政治敏感的言论或者不合规的广告，没有经过后台审核就直接显示出来了，而你的博客部署在大陆，那么你的博客很可能会被马上关停整顿，并且自己也会解锁程序员从入门到如入狱的成就。也千万不要以为部署在境外就没事了，一些仇恨性质的言论甚至可以帮你引来黑客，在你的博客里下毒，勒索你或你的读者。

因此我强烈建议个人博客启用敏感词过滤及评论审查功能。WordPress及我的Moonglade博客系统均支持敏感词过滤和评论审查。

### 静态化

早期的新闻系统、博客、CMS为了提高大访问量下的响应速度，都会采用静态化技术，即将服务端渲染完的页面保存为真正的HTML文件于磁盘上，进行static file的输出，Web服务器输出static file的效率非常高，对于不变的内容，用户的后续访问不会hit数据库，因此极大减小了服务器的压力。在2020年的今天，静态化已经不是唯一的方案，Redis Cache也可以帮助我们减少对数据库的频繁访问。对于个人博客来说，如果你的访问量不高，其实并不需要996一个静态化或Redis出来增加开发和维护成本。但如果你设计的是博客平台，那么最好还是用上静态化或Redis吧。

### 通知系统

博客通常通过Email的形式给管理员或用户发送通知。但是没有规范或或约定表示博客是否一定得使用Email进行通知推送，可由博客系统设计者自行决定。

通知通常包括：

向博主发送的通知：新评论、文章被他人博客引用（参见第5.8, 5.9章）。

向用户发送通知：新文章发布（订阅Newsletter）、评论被回复、评论审核通过或被拒。

Email通知系统要注意垃圾邮件及用户隐私保护问题。



垃圾邮件发给博主本身问题并不大，但得注意邮件系统是否会允许未经博主许可的针对读者的邮件发送，其中可能会被人利用发垃圾邮件，从而导致服务器被封禁。有些服务器供应商，例如微软Azure，对于邮件有更加严格的规定，部署在部分PaaS业务上的代码调用SMTP终端会被直接屏蔽。

对于用户隐私问题，在用户向博客系统提供Email地址的同时，需要告知用户该Email地址会被如何使用（可写在隐私协议或页面可见区域），也可以让用户勾选是否允许博主使用该Email进行通知推送。另一个问题是邮件地址暴露，这通常发生在Newsletter的订阅群发，如果把所有订户的Email地址都放在To或CC里，那么每个用户都会知道其余所有人的Email地址，从而互相约炮、欺诈，因此Newsletter请采用BCC或单独发送，并允许用户退订。

Moonglade的通知系统采用Email方式，但设计比较基础。一个完善的通知系统需要采用消息队列及事件设计，并采用三方服务。例如Azure上可以使用Storage Queue + Function App + SendGrid，以免遇到大批量Email发送的时候原地爆炸。

## 博客协议或标准
RSS / ATOM / RSD / OPML / APML / BlogML / Yahoo Media / iTunes

https://github.com/shawnwildermuth/RssSyndication

https://github.com/stimpy77/Argotic

https://github.com/argotic-syndication-framework/Argotic

### RSS
RSS（Really Simple Syndication）是一种基于XML的标准，普遍应用于包括博客在内的内容类网站，由Dave Winer于1999年发明，少年计算机天才Aaron Swartz参与定义规范，可惜后者于2013年1月自杀，年仅26岁。

RSS也是博客系统中最有标志性特性之一，其在博客中的应用广泛度成为了事实上的标准，没有RSS的博客系统就像看到不带摄像头的手机一样有趣。

RSS文件的扩展名可通常是 .rss 或 .xml，也可以不定义拓展名（如Moonglade的RSS）。内容为近期发表的博客文章的XML描述，包括标题、时间、作者、分类、摘要（也可以是全文）等信息。

RSS是写给机器看的，可用于网站之间同步内容，例如当年人人网（前校内网）可通过RSS导入博客文章为日记。而对于普通用户，则需要RSS阅读器应用来订阅博客。通常这样的阅读器里不止订阅一个作者的博客，而是该用户关心的所有博客。阅读器通常也是跨平台、跨设备的，用户可以在电脑、平板、手机，甚至树莓派上订阅RSS源。

部分浏览器（如早期的火狐）也可以自动识别一个博客的RSS地址，并在浏览器中订阅。其自动发现原理是查找网页head中有没有这么一个东西：

`<link rel="alternate" type="application/rss+xml" title="Edi Wang" href="/rss" />`

但是RSS有个缺点，它并不能够由服务器主动向客户端推送，而需要靠客户端自动去服务器拉取。而过去10年中，随着移动端的兴起，消息推送服务弥补了RSS的不足，各大平台也几乎都推出了自己的手机APP，因此RSS已经被许多网站淘汰。但并不意味着RSS没用了，至今仍有大量网站仍然提供RSS订阅。例如微软Channel 9电视台的RSS: `https://channel9.msdn.com/Feeds/RSS/`，国内的博客园的`RSS：http://feed.cnblogs.com/blog/sitehome/rss`，有意思的是博客园网站的logo其实就是个RSS图标。

对于构建博客系统而言，你通常不会再专门做个手机App，用户也不会为每一个博客都单独下载一个App，并且博客系统与其他博客、网站之间依然需要同步，不可能为每个合作伙伴都开发一套同步协议，大家依然都用已经是公认标准的RSS，因此RSS在2020年依然是博客系统推送文章的最佳途径。

参考：https://en.wikipedia.org/wiki/RSS

### ATOM

ATOM和RSS的作用几乎一样，但ATOM的出现是为了弥补RSS的一些设计缺陷。例如对于文章发表日期，ATOM采用RFC 3339的时间戳，而RSS采用的是RFC 822标准。ATOM也可以标识文章的语言、允许payload中出现RSS不允许的XHTML、XML和Base64编码内容等。

许多博客系统（包括我的Moonglade）同时提供RSS及ATOM源。

参考链接：https://en.wikipedia.org/wiki/Atom_(Web_standard) 

### OPML
“OPML（概述处理器标记语言）是用于轮廓的XML格式（定义为“一棵树，其中每个节点包含一组具有字符串值的命名属性” ）。它最初由UserLand在其Radio UserLand产品中作为大纲应用程序的本机文件格式开发，此后已被用于其他用途，最常见的是在Web Feed聚合器之间交换Web Feed列表。

OPML规范将大纲定义为任意元素的层次结构，有序列表。该规范相当开放，因此适用于多种类型的列表数据。

Mozilla Thunderbird 和许多其他RSS阅读器网站和应用程序都支持以OPML格式导入和导出RSS feed列表。”

参考：https://en.wikipedia.org/wiki/OPML

通俗易懂的说，OPML对于博客来说，就是告诉阅读器，这个博客一共有哪些订阅源以及他们各自的订阅地址，通常就是每个文章分类是一个订阅源，全部文章又是一个订阅源。

#### OPML概念

http://dev.opml.org/spec1.html

http://dev.opml.org/spec2.html
[使用并解析 OPML 格式的订阅列表来转移自己的 RSS 订阅（解析篇）](https://blog.walterlv.com/post/deserialize-opml-using-dotnet)
- 典型的 OPML 文件
https://edi.wang/opml
http://home.opml.org/?format=opml
```xml
<?xml version="1.0" encoding="UTF-8"?>
<opml version="1.0">
  <head>
    <title>walterlv</title>
  </head>
  <body>
    <outline text="walterlv" title="walterlv" type="rss" xmlUrl="https://blog.walterlv.com/feed.xml" htmlUrl="https://blog.walterlv.com/" />

    <outline title="Team" text="Team">
      <outline text="林德熙" title="林德熙" type="rss" xmlUrl="https://blog.lindexi.com/feed.xml" htmlUrl="https://blog.lindexi.com/" />
    </outline>

    <outline title="Microsoft" text="Microsoft">
      <outline text="Microsoft .NET Blog" title="Microsoft .NET Blog" type="rss" xmlUrl="https://blogs.msdn.microsoft.com/dotnet/feed/"/>
      <outline text="Microsoft The Visual Studio Blog" title="Microsoft The Visual Studio Blog" type="rss" xmlUrl="https://blogs.msdn.microsoft.com/visualstudio/feed/"/>
    </outline>
  </body>
</opml>
```
你可以很容易地看出它的一些特征。比如以 opml 为根，head 中包含 title，body 中包含分组的 outline。每一个 outline 中包含 text, type, xmlUrl 等属性。接下来我们详细描述这个格式。

- OPML 文件中的节点解释
**opml 根节点**
`<opml>` 是 OPML 格式文件的根节点，其 version 属性是必要的。它的值可能为 1.0 或 2.0；如果是 1.0，则视为符合 [OPML 1.0 规范](http://dev.opml.org/spec1.html)；如果是 2.0，则视为符合 [OPML 2.0 规范](http://dev.opml.org/spec2.html)。额外的，值也可能是 1.1，那么也视为符合 1.0 规范。
opml 根节点中包含 head 和 body 节点。
**head 节点**
head 节点可包含 0 个或多个元素：
    - title 这就是 OPML 文档标题
    - dateCreated 文档创建时间
    - dateModified 文档修改时间
    - ownerName 文档作者
    - ownerEmail 文档作者的邮箱
    - ownerId 文档作者的 url，要求不存在相同 Id 的两个作者
    - docs 描述此文档的文档的 url
    当然，这些都是可选的。
    额外的，还有 expansionState, vertScrollState, windowTop, windowLeft, windowBottom, windowRight。

**body 节点**
body 节点包含一个或多个 outline 元素。
**outline（普通）**
outline 元素组成一个树状结构。也就是说，如果我们使用 OPML 储存 RSS 订阅列表，那么可以存为树状结构。在前面的例子中，我把自己的 RSS 订阅独立开来，把朋友和微软的 RSS 订阅分成了单独的组。

outline 必须有 text 属性，其他都是可选的。而 text 属性就是 RSS 订阅的显示文字，如果没有这个属性，那么 RSS 的订阅列表中将会是空白一片。

于是，我们解析 text 属性便可以得到可以显示出来的 RSS 订阅列表。对于前面的例子对应的 RSS 订阅列表就可以显示成下面这样：
```
- walterlv
- Team
    - 林德熙
- Microsoft
    - Microsoft .NET Blog
    - Microsoft The Visual Studio Blog
```
outline 还有其他可选属性：
type 指示此 outline 节点应该如何解析
isComment 布尔值，为 true 或 false；如果为 true，那么次 outline 就只是注释而已
isBreakpoint 适用于脚本，执行时可下断点
created 一个时间，表示此节点的创建时间
category 逗号分隔的类别：如果表示分类，则要用 / 分隔子类别；如果表示标签，则不加 /
例如：/Boston/Weather, /Harvard/Berkman,/Politics（例子来源于[官方规范](http://dev.opml.org/spec2.html)）

**outline（RSS 专属）**
当 type 是 rss 时，还有一些 RSS 专属属性。这时，必要属性就有三个了：
type
text
xmlUrl
其中，xmlUrl 就指的是订阅源的 url 地址了。在官方规范中，规定解析器不应该总认为 text 存在，相比之下，xmlUrl 显得更加重要。
还有一些可选属性：
description
htmlUrl
language
title
version



### APML
APML即Attention Profiling Mark-up Language，它比OPML更鲜为人知。APML目前在互联网上已经非常少见了，比WP还惨。作为博客行业的历史遗迹之一，抱着情怀简短介绍一下。

与OPML类似，它也是一种XML格式的声明文件，用来描述个人感兴趣的事物或话题，并分享给其他读者或博主，以帮助阅读器或者博客系统本身针对用户感兴趣的内容提供服务或更有针对性的广告。

参考链接：https://en.wikipedia.org/wiki/Attention_Profiling_Mark-up_Language

WordPress可以通过插件实现APML，BlogEngine则自带APML，我的Moonglade不支持APML。

### FOAF
FOAF即Friend of a Friend，也是个写给机器看的文件，描述了一个人类的社交关系，通常在博客中可以用FOAF表示博主和其他博客之间的 “友情链接” ，只不过这个友情链接是写给机器看的。好让机器明白，谁才是你的基友，从而给读者推荐基友博客里的内容。

WordPress可以通过插件实现FOAF，BlogEngine自带FOAF，我的Moonglade不支持FOAF。FOAF和APML的现状差不多，已快绝迹。

参考链接：https://en.wikipedia.org/wiki/FOAF_(ontology)


### BlogML
BlogML是一套跨博客系统的数据标准，凡是实现了BlogML的博客系统，就算语言、平台不一样，也都可以互相导入、导出文章等数据。就好比HTML5是个标准，Edge、Chrome、Firefox是浏览器，只要针对HTML5写的网页都能跨这些浏览器运行。

BlogML也诞生于.NET社区之中，随后发展成了标准。除了本身就是.NET的BlogEngine等系统以外，PHP写的WordPress都支持BlogML。当年支持BlogML的还有Windows Live Spaces，Subtext，DasBlog等。我的Moonglade不支持BlogML。

当前BlogML的标准schema是2.0，更新于2006年11月25日。看起来这个标准也已经……

参考：https://en.wikipedia.org/wiki/BlogML

### Open Search
如果博客实现了Open Search规范，那么博客的搜索功能就能够自动整合到用户的浏览器里，从而便于用户直接在浏览器地址栏使用你博客的搜索服务作为搜索引擎（就像必应、谷歌那样）。

实现Open Search只需两部，首先在网页的head里加入指向opensearch定义文件的link

`<link type="application/opensearchdescription+xml" rel="search" title="Edi Wang" href="/opensearch" />`

然后输出opensearch文件即可
```
<OpenSearchDescription xmlns="http://a9.com/-/spec/opensearch/1.1/">

<ShortName>Edi Wang</ShortName>

<Description>Latest posts from Edi Wang</Description>

<Image height="16" width="16" type="image/vnd.microsoft.icon">https://edi.wang/favicon.ico</Image>

<Url type="text/html" template="https://edi.wang/search/{searchTerms}"/>

</OpenSearchDescription>
```
文件描述了博客的名称、简介、图标以及搜索内容的URL pattern。浏览器一旦识别这个文件，会自动将你的博客注册到搜索引擎列表里去。然后读者就可以直接在浏览器地址栏里搜索关键词，并显示博客自己的搜索结果页面。

Open Search的具体规范和标准可参考：https://en.wikipedia.org/wiki/OpenSearch

### Pingback

Pingback用于博客系统之间通讯，一旦自己的文章被他人引用就会收到pingback请求，而自己引用了他人的文章就会向对方博客发送一个pingback请求，因此完成一次Pingback需要己方和对方的博客共同支持pingback协议。由于是标准协议，所以pingback并不要求双方的博客使用同一款博客产品，例如我用.NET Core写的Moonglade可以完美和PHP写的WordPress互相ping。Pingback也并不限制网站类型一定得是博客，任何CMS或内容网站想要支持Pingback都没问题。

Pingback的技术原理也不复杂。

发送Pingback请求：

得到自己文章的URL A、对面被引用文章的URL B，请求B，看看它有没有pingback终端，如果有，构建一个HTTP Request，内容是一段XML：
```
<methodCall>
       <methodName>pingback.ping</methodName>
       <param>
              <param><value><string>A</string></value></param>
              <param><value><string>B</string></value></param>
       </param>
</methodCall>
```
这样B所在的网站就知道A文章引用了B文章，处理pingback后，会给A所在的网站一个成功与否的响应。

接受Pingback请求：

自己的文章URL A被他人文章B引用，并收到了一个pingback XML。首先自己要验证别人的pingback请求长得是否奇怪，以保证安全性，例如有没有正常的methodName、有没有合法的双方URL、URL是否能正常访问、是否有奇怪的URL（例如localhost或有潜在攻击行为的特殊构造）。保证pingback请求没问题后，请求B的页面，抓取B网页的title内容、B的IP地址，记录到自己的数据库中，并和A文章关联。

收到的Pingback通常以系统身份自动在文章下加评论，但这个设计不是规范之一，你可以自由发挥，例如Moonglade把pingback集中起来在后台给博客管理员查看。

参考：https://en.wikipedia.org/wiki/Pingback

### Trackback
Trackback允许一个网站将更新通知给另一个网站。这是网站作者在有人链接到其文档之一时请求通知的四种类型的链接方法之一。这使作者可以跟踪谁链接到他们的文章。

参考：https://en.wikipedia.org/wiki/Trackback

尽管功能和Pingback类似，但Trackback通常需要手工发送，并需要给对方提供一篇文章的摘要。而Pingback的过程是又双方博客系统共同完成的全自动操作。

### MetaWeblog

MetaWeblog是一套基于XML-RPC 的Web Service，这套API定义了几个标准接口，用于文章、分类、标签等博客常规内容的CRUD。只要实现了这些接口的博客系统，就可以让博主不用通过浏览器登录博客后台写文章，而使用计算机上安装的客户端去写博客。主流的客户端包括 Windows Live Writer、Microsoft Word。在客户端里可以完整的编辑文章、插入图片、设置分类，甚至可以将博客的主题同步到客户端中。

可能它看起来也像是过时的博客协议之一，但直到2020年的今天，最新版的Microsoft 365套件依然完整支持实现了MetaWeblog API的博客系统。

类似MetaWeblog的博客API还有Blogger API, Atom Publishing Protocol, Micropub。

参考：https://en.wikipedia.org/wiki/MetaWeblog

我的博客在2012年曾经996 007完整实现了MetaWeblog + RSD，但如今30岁了，在.NET Core里暂时不打算实现这个了，毕竟有多少人还在用Live Writer和Word写博客（哭。

### RSD

Really Simple Discovery（RSD）是XML格式和一种发布约定，用于使博客或其他Web软件公开的服务可由客户端软件发现。这是一种将设置编辑/博客软件所需的信息减少到三个众所周知的元素的方法：用户名，密码和主页URL。任何其他关键设置都应该在与网站相关的RSD文件中定义，或者可以使用提供的信息来发现。

为了使用RSD，网站的所有者在首页的head里放置了一个链接标记，用于指示RSD文件的位置。MediaWiki使用的一个示例是：
```
<link rel="EditURI" type="application/rsd+xml" href="https://en.wikipedia.org/w/api.php?action=rsd" />
```
然后用RSD文件去表示各种API的接口
```
<?xml version="1.0"?>
<rsd version="1.0" xmlns="http://archipelago.phrasewise.com/rsd">
    <service>
        <apis>
            <api name="MediaWiki" preferred="true" apiLink="http://en.wikipedia.org/w/api.php" blogID="">
                <settings>
                    <docs xml:space="preserve">http://mediawiki.org/wiki/API</docs>
                    <setting name="OAuth" xml:space="preserve">false</setting>
                </settings>
            </api>
        </apis>
        <engineName xml:space="preserve">MediaWiki</engineName>
        <engineLink xml:space="preserve">http://www.mediawiki.org/</engineLink>
    </service>
</rsd>
```
参考：https://en.wikipedia.org/wiki/Really_Simple_Discovery

RSD也几乎和上面的MetaWeblog接口一起使用。这样Windows Live Writer、Microsoft Word等工具才可以自动发现博客的MetaWeblog服务，而不需要手工去输URL。

### 阅读器视图
大部分浏览器和客户端都有阅读器视图，可以让读者在与博客网站页面风格完全不一样的视图中阅读文章。

浏览器识别到我的博客支持阅读器视图，就会亮起沉浸式阅读按钮

进入沉浸式阅读界面后，浏览器会自动提取文章的内容，识别文章的标题、章节、图片，去掉导航栏、侧边栏等与文章无关的元素，并可让用户控制文本大小、背景色，甚至朗读文章内容。

不仅我的博客有阅读器视图，设计良好的博客、新闻内容站都有，例如Azure的：
另外，支持阅读器视图的网站，SEO一定不会差。因此设计博客系统时，请考虑支持阅读器视图。

## 设计博客系统有哪些知识点
### 时区真的全用UTC
存储时间使用UTC在2020年应该已经是猿尽皆知的实践了，博客系统其实也是如此，我的博客所有时间数据最终保存都采用UTC时间。但博客有个特殊的地方，即它不应该按读者的时区去转换UTC时间进行显示，而应该按照博客作者的时区去显示时间。

这并不是技术上的原因，就算你按读者时区去显示时间也不会有代码爆炸，原因在于博客的诞生初衷，就是为了彰显个性，让博主在互联网上有自己的展示空间，因此突出博主本人的属性非常重要，博主所在时区也是个让读者了解博主的属性之一，因此，正宗的博客系统都会给一个时区设置选项，并以此转换UTC时间作为显示，WordPress和我的Moonglade博客系统均是如此。博客系统不自动转换读者所在时区的时间，纯粹就是个鲜为人知的情怀设计，但必须得尊重。

那么有意思的事情来了，搜索引擎要怎么理解博客文章的时间？最好将UTC时间仅告诉搜索引擎，不要给用户显示，方法也很简单，用HTML5的time标签的datetime属性即可。在HTML5标准推广以后，搜索引擎更喜欢看标签类型来判断内容的含义，而不是根据标签里的内容来猜意思。

在C#里，ToString(“u”)指的是Universal sortable date/time patter。
```
<time datetime="@Model.PostModel.PubDateUtc.ToString("u")" title="GMT @Model.PostModel.PubDateUtc">@DateTimeResolver.GetDateTimeWithUserTZone(Model.PostModel.PubDateUtc).ToString("MM/dd/yyyy")</time>
```
对于刚才截图里的文章，时间的HTML为：
`<time datetime="2020-04-29 11:41:02Z" title="GMT 4/29/2020 11:41:02 AM">04/29/2020</time>`

### HTML还是Markdown

许多技术人士编写博客系统的时候喜欢选用Markdown作为编辑器，如果单纯只是个技术博客，自己使用并没有什么问题。但如果你在给他人编写博客系统，请记住，不是每个人，都是程序员，不是每个人，都喜欢Markdown。

在这种情况下，一个WSIWYG的HTML编辑器（如TinyMCE）是不错的选择，HTML编辑器相对Markdown也支持更高级的排版方式。Moonglade 同时支持HTML和Markdown编辑器。

保存文章内容到数据库时，Markdown格式需要选择原始内容，而非生成的HTML，因为还需要支持后续编辑。HTML格式现在也不建议encoding存储，毕竟都已经2020年了，市面上的主流数据库都可以正确支持各种神奇的Unicode，比如文章中突然出现个emoji😂，而如果你使用了encoding，就会像我的博客一样面临一些福报：https://github.com/EdiWang/Moonglade/issues/280。并且encoding和decoding的过程会影响性能。我的Moonglade博客系统也刚刚完成了去除encoding的改造。

### MVC还是SPA

许多社区里写博客系统的程序员都偏向于使用SPA架构建博客，而鄙视用MVC，觉得落后，真的是这样吗？这个问题就像是飞机为什么不飞直线，是航空公司不会规划吗？关于这一点，我曾经在以前的博客文章《我的 .NET Core 博客性能优化经验总结》中写过：

2014年以后，随着SPA的兴起，Angular等框架逐渐成为了前端开发的主流。它们解决的问题正是提升前端的响应度，让Web应用尽量接近本地原生应用的体验。我也面临过不少朋友的质疑：为什么你的博客不用angular写？是你不擅长吗？

其实并不是那么简单。实际上我任职的岗位的目前主要工作内容也是写angular，博客曾经的.NET Framework版的后台也用过angularjs以及angular2，经过一系列的实践表明，我博客这样的内容站用angular收益并不大。

其实这并不奇怪，在盲目选择框架之前，我们得注意一个前提条件：SPA框架所针对的，其实是Web应用。而应用的意思是重交互，即像Azure Portal或Outlook邮箱那样，目的是把网页当应用程来开发，这时候SPA不仅能提升用户体验，也能降低开发成本，何乐而不为？但是博客属于内容为主的网站，不是应用，要说应用也勉强只能说博客的后台管理可以是应用。博客前台唯一的交互就是评论、搜索，因此SPA并不适合这样的工作。这就像你要去菜场买菜，骑自行车反而比你开个坦克过去方便。

在微软官方文档里也有同样的关于何时选择SPA，何时选择传统网站的参考：

https://docs.microsoft.com/en-us/dotnet/architecture/modern-web-apps-azure/choose-between-traditional-web-and-single-page-apps 

博客前台仍然选用MVC的另一个原因，请回顾一下本文开头“博客的读者是谁”，我运营博客十余年，统计的数据表明，几乎所有的用户都来源于搜索引擎，都只点进来看一篇文章，然后关闭网页。现在仔细想想，SPA解决的最大的问题之一是什么？是不是通过只刷新局部来提高前端性能（可响应度）？而用户从搜索引擎过来，只看一篇文章就关闭网页，真的用得到SPA只刷新局部的优势吗？用户只看一篇文章，你用个SPA框架，用户得加载一堆框架本身的文件，其中包括导航、交互等功能，而99%的用户根本就不会点到别的地方去，于是你只为了1%的用户，去加载硕大的一个框架，值得吗？这性能到底是提高了，还是降低了？

MVC框架虽然每次都会输出服务器端渲染的完整HTML，但由于99%的用户只看一篇文章就关闭网页，所以对于99%的用户来说，他们所需要加载的资源，远小于加载一套SPA，速度更快，还更SEO友好。SPA适合用在博客的后台管理portal，而不是前台。

### 安全

根据运营博客多年的后台监控数据，最常见的攻击行为是全自动的漏洞扫描工具。他们会请求例如 data.zip, wp-admin.php, git目录等常见的安全疏忽，或是想要通过某些博客系统的已知漏洞进行攻击。目的是为了控制服务器，在你的博客网页里加入对用户的恶意代码（例如勒索病毒、挖矿）等，有些也会将服务器本身变成矿机。

设计博客系统时，常用的安全对策可参考OWASP（https://owasp.org/），但同时保留灵活性。例如，加入JavaScript的CSP时，请考虑正常博客用户可能需要添加三方统计插件（如Azure Application Insights，国内的CNZZ等），请设计一定的黑、白名单或功能开关。

大部分设计者都知道要防范用户的输入，即博客的读者，输入的入口通常只有评论和搜索功能。但不要忘了，博主在博客后台管理中的输入也需要防范，因为不一定是博主本人在操作。举个例子，博主的账号被盗，黑客在后台将导航栏的链接指向黑客的服务器或localhost上早已准备好的奇妙的机关（是的，不要以为localhost在正常人的电脑上不起作用），那么读者就会受到严重影响。

关于后台登录的身份认证，能采用成熟的SSO的就优先采用SSO，例如Moonglade支持Azure Active Directory验证，这样能够利用微软这样的专业服务管理授权认证，尽可能小的避免账户上产生安全问题。如果用户没有SSO的环境，才fallback到本地账号认证。千万不要认为用三方服务没自己写安全，觉得自己写的逻辑没人知道就不会被黑了，除非你是世界顶级大牛，不然自己写的系统易黑程度远高于三方服务。

另有一些攻击通常由一些敌对阵营的无聊程序员发起，例如使用脚本或工具持续不断的请求博客系统的某个URL，企图像DDOS那样击爆服务器，对于这种无聊刷刷党，博客系统设计者只要加入有关URL endpoint的rate limit即可。对于真实的DDOS攻击，只有云端抗DDOS服务或硬件DDOS防火墙才能解决。

最后别忘了OWASP里没有的东西，博客的协议也会有设计缺陷，例如pingback可以用来DDOS（https://www.imperva.com/blog/wordpress-security-alert-pingback-ddos/），也能扫描服务器端口（https://www.avsecurity.in/wordpress-xml-rpc-pingback-vulnerability/）

## 结束语

设计一个优秀的博客系统，每一处细节都值得斟酌。这些设计绝对不可能一开始就能做对，而是得靠长期运营博客的数据去发现并思考。并且，市场会变化，用户行为会变化，标准会被淘汰，也会被发明，因此你的系统需要跟着进化。

任何看似简单的系统，就算普通到烂大街，也有背后看不见的一套完整体系。博客如此，电子商城、外卖、金融清算系统更是复杂，不要光凭自己表面看到的就开始做。就如同造飞机，造个纸飞机和真飞机，绝对不是一回事。

技术人员也不要觉得什么流行就得用什么，优秀的产品并不是堆砌时髦技术做出来的，而先得分析你的用户到底是怎么用你的产品，才能做最合适的选择。要记住，想要一件事情做成功，思路不要只局限于技术本身，学会分析市场，用户行为，才能更准确的选择和应用技术。