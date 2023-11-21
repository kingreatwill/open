# ExifTool完全入门指南

## ExifTool介绍

  ExifTool由Phil Harvey开发，是一款免费、跨平台的开源软件，用于读写和处理图像（主要）、音视频和PDF等文件的元数据（metadata）。ExifTool可以作为Perl库（Image::ExifTool）使用，也有功能齐全的命令行版本。ExifTool支持很多类型的元数据，包括Exif、IPTC、XMP、JFIF、GeoTIFF、ICC配置文件、Photoshop IRB、FlashPix、AFCP和ID3，以及众多品牌的数码相机的私有格式的元数据。

## 什么是Exif

  Exif是可交换图像文件格式（Exchangeable image file format），是一种标准，定义了与数码相机捕获的图像（或其他媒体）有关的信息，用于存储重要的数据，比如相机的曝光、拍摄日期和时间，甚至GPS定位等。在早期，摄影师需要随身携带笔记本来记录重要信息，如日期、快门速度、光圈等，这非常麻烦而且容易出错。如今，每台数码相机都支持Exif，能够将拍摄时的很多参数通过这种格式（Exif）记录到照片中，这些照片（或其他类型的文件）中的额外数据就叫元数据（metadata），它由一系列参数组成，如快门速度、光圈、白平衡、相机品牌和型号、镜头、焦距等等。Exif信息可能会造成隐私泄露（相机型号、位置等），在社会工程学中，Exif也是获取目标信息的一种手段，所以建议在把照片上传到互联网之前先清理Exif数据。

## Official website

[ExifTool by Phil Harvey](https://exiftool.org)

[ExifTool in Github](https://github.com/exiftool/exiftool)

[ExifTool完全入门指南](https://www.rmnof.com/article/exiftool-introduction/)

[ExifTool 安装](https://exiftool.org/install.html)


## ExifTool命令格式

 读取：exiftool [OPTIONS] [-TAG...] [--TAG...] FILE...
 写入：exiftool [OPTIONS] -TAG[+-<]=[VALUE]... FILE...
 复制：exiftool [OPTIONS] -tagsFromFile SRCFILE [-SRCTAG[>DSTTAG]...] FILE...
 其他：exiftool [ -ver | -list[w|f|r|wf|g[NUM]|d|x] ]

### ExifTool参数一览

```
标签选项
  -TAG or --TAG                    提取或排除指定的标签
  -TAG[+-^]=[VALUE]                向标签写入新值
  -TAG[+-]<=DATFILE                从数据文件读取标签写入文件
  -TAG[+-]<SRCTAG                  复制标签值（见-tagsFromFile）

  -tagsFromFile SRCFILE            复制某文件的标签值
  -x TAG      (-exclude)           排出指定标签
  
输入输出文本格式
  -args       (-argFormat)         将元数据格式化为ExifTool参数
  -b          (-binary)            以二进制输出元数据
  -c FMT      (-coordFormat)       设置GPS坐标格式
  -charset [[TYPE=]CHARSET]        指定字符编码
  -csv[[+]=CSVFILE]                以CSV格式导出/导入标签
  -d FMT      (-dateFormat)        设置日期/时间的格式
  -D          (-decimal)           以十进制显示标签ID号
  -E,-ex,-ec  (-escape(HTML|XML|C))为HTML，XML或C的转义标记值
  -f          (-forcePrint)        强制打印所有指定的标签
  -g[NUM...]  (-groupHeadings)     按标签组输出
  -G[NUM...]  (-groupNames)        打印每个标签的组名
  -h          (-htmlFormat)        输出为HTML格式
  -H          (-hex)               以十六进制显示标签ID号
  -htmlDump[OFFSET]                生成HTML格式的二进制转储
  -j[[+]=JSONFILE] (-json)         以JSON格式导出/导入标签
  -l          (-long)              使用长2行输出格式（标签、值各一行）
  -L          (-latin)             使用Windows Latin1编码
  -lang [LANG]                     设定当前语言
  -listItem INDEX                  从列表中提取特定项目
  -n          (--printConv)        不转换，直接打印
  -p FMTFILE  (-printFormat)       以指定格式输出
  -php                             将标签导出为PHP数组
  -s[NUM]     (-short)             简短输出
  -S          (-veryShort)         非常简短输出
  -sep STR    (-separator)         设置列表项的分隔符字符串
  -sort                            按字母顺序对输出进行排序
  -struct                          启用结构化信息的输出
  -t          (-tab)               以制表符分隔的列表格式输出
  -T          (-table)             以表格格式输出
  -v[NUM]     (-verbose)           打印详细消息
  -w[+|!] EXT (-textOut)           写入（或覆盖！）输出的文本文件
  -W[+|!] FMT (-tagOut)            为每个标签写入输出文本文件
  -Wext EXT   (-tagOutExt)         用-W指定要写入的文件类型
  -X          (-xmlFormat)         使用RDF/XML输出格式

执行选项
  -a          (-duplicates)        允许提取重复的标签
  -e          (--composite)        不生成复合标签
  -ee         (-extractEmbedded)   从嵌入式文件中提取信息
  -ext[+] EXT (-extension)         只处理具有指定扩展名的文件
  -F[OFFSET]  (-fixBase)           修复制造商Makernotes偏移
  -fast[NUM]                       提取元数据时提高速度
  -fileOrder[NUM] [-]TAG           设置文件处理顺序
  -i DIR      (-ignore)            忽略指定的目录名称
  -if[NUM] EXPR                    按条件处理文件
  -m          (-ignoreMinorErrors) 忽略小错误和警告
  -o OUTFILE  (-out)               设置输出文件或目录名称
  -overwrite_original              重命名tmp文件覆盖原始文件
  -overwrite_original_in_place     通过复制tmp文件覆盖原始文件
  -P          (-preserve)          保留文件修改日期/时间
  -password PASSWD                 处理受保护文件的密码
  -progress[:[TITLE]]              显示文件进度计数
  -q          (-quiet)             -q不显示正常消息，-q-q不显示警告
  -r[.]       (-recurse)           递归处理子目录
  -scanForXMP                      扫描所有文件以获取XMP
  -u          (-unknown)           提取未知标签
  -U          (-unknown2)          提取未知的二进制标签
  -wm MODE    (-writeMode)         设置用于写入/创建标签的模式
  -z          (-zip)               读/写压缩信息

其他选项
  -@ ARGFILE                       从文件中读取命令行参数
  -k          (-pause)             结束前先暂停
  -list[w|f|wf|g[NUM]|d|x]         列出各种exiftool功能
  -ver                             打印版本号
  --                               结束选项

特殊功能
  -geotag TRKFILE                  从指定的GPS日志对图像进行地理标记
  -globalTimeShift SHIFT           移位所有格式化的日期/时间值
  -use MODULE                      从插件模块添加功能

实用工具
  -delete_original[!]              删除"_original"备份
  -restore_original                从"_original"备份还原

高级选项
  -api OPT[[^]=[VAL]]              设置ExifTool API选项
  -common_args                     定义通用参数
  -config CFGFILE                  指定配置文件名
  -echo[NUM] TEXT                  将文本回显到stdout或stderr
  -efile[NUM][!] ERRFILE           保存错误的文件名
  -execute[NUM]                    一行执行多个命令
  -srcfile FMT                     处理其他文件来源
  -stay_open FLAG                  继续阅读-@ argfile，即使在EOF之后
  -userParam PARAM[[^]=[VAL]]      设置用户参数 (API UserParam opt)
```