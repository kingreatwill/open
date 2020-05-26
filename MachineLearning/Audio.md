# 几个最新免费开源的中文语音数据集


工欲善其事必先利其器，做机器学习，我们需要有利器，才能完成工作，数据就是我们最重要的利器之一。做中文语音识别，我们需要有对应的中文语音数据集，以帮助我们完成和不断优化改进项目。我们可能很难拿到成千上万小时的语音数据集，但是这里有一些免费开源的语音数据集，大家一定不要错过。文末附数据集下载地址。我们也非常感谢相关单位和团体为国内的开源界做出的贡献。

### 普通语音识别数据集

**THCHS30**

THCHS30是一个很经典的中文语音数据集了，包含了1万余条语音文件，大约40小时的中文语音数据，内容以文章诗句为主，全部为女声。它是由清华大学语音与语言技术中心（CSLT）出版的开放式中文语音数据库。原创录音于2002年由朱晓燕教授在清华大学计算机科学系智能与系统重点实验室监督下进行，原名为“TCMSD”，代表“清华连续”普通话语音数据库’。13年后的出版由王东博士发起，并得到了朱晓燕教授的支持。他们希望为语音识别领域的新入门的研究人员提供玩具级别的数据库，因此，数据库对学术用户完全免费。

license: Apache License v.2.0

**ST-CMDS**

ST-CMDS是由一个AI数据公司发布的中文语音数据集，包含10万余条语音文件，大约100余小时的语音数据。数据内容以平时的网上语音聊天和智能语音控制语句为主，855个不同说话者，同时有男声和女声，适合多种场景下使用。

License: Creative Common BY-NC-ND 4.0 (Attribution-NonCommercial-NoDerivatives 4.0 International)

**AISHELL开源版**

AISHELL是由北京希尔公司发布的一个中文语音数据集，其中包含约178小时的开源版数据。该数据集包含400个来自中国不同地区、具有不同的口音的人的声音。录音是在安静的室内环境中使用高保真麦克风进行录音，并采样降至16kHz。通过专业的语音注释和严格的质量检查，手动转录准确率达到95％以上。该数据免费供学术使用。他们希望为语音识别领域的新研究人员提供适量的数据。

License: Apache License v.2.0

**Primewords Chinese Corpus Set 1**

Primewords包含了大约100小时的中文语音数据，这个免费的中文普通话语料库由上海普力信息技术有限公司发布。语料库由296名母语为英语的智能手机录制。转录准确度大于98％，置信水平为95％，学术用途免费。抄本和话语之间的映射以JSON格式给出。

**aidatatang**

Aidatatang_200zh是由北京数据科技有限公司（数据堂）提供的开放式中文普通话电话语音库。

语料库长达200小时，由Android系统手机（16kHz，16位）和iOS系统手机（16kHz，16位）记录。邀请来自中国不同重点区域的600名演讲者参加录音，录音是在安静的室内环境或环境中进行，其中包含不影响语音识别的背景噪音。参与者的性别和年龄均匀分布。语料库的语言材料是设计为音素均衡的口语句子。每个句子的手动转录准确率大于98％。

**MAGICDATA Mandarin Chinese Read Speech Corpus**

Magic Data技术有限公司的语料库，语料库包含755小时的语音数据，其主要是移动终端的录音数据。邀请来自中国不同重点区域的1080名演讲者参与录制。句子转录准确率高于98％。录音在安静的室内环境中进行。数据库分为训练集，验证集和测试集，比例为51：1：2。诸如语音数据编码和说话者信息的细节信息被保存在元数据文件中。录音文本领域多样化，包括互动问答，音乐搜索，SNS信息，家庭指挥和控制等。还提供了分段的成绩单。该语料库旨在支持语音识别，机器翻译，说话人识别和其他语音相关领域的研究人员。因此，语料库完全免费用于学术用途。

**AISHELL-2 高校学术免费授权版数据集**

希尔贝壳中文普通话语音数据库AISHELL-2的语音时长为**1000**小时，其中718小时来自AISHELL-ASR0009-\[ZH-CN\]，282小时来自AISHELL-ASR0010-\[ZH-CN\]。录音文本涉及唤醒词、语音控制词、智能家居、无人驾驶、工业生产等12个领域。录制过程在安静室内环境中， 同时使用3种不同设备： 高保真麦克风（44.1kHz，16bit）；Android系统手机（16kHz，16bit）；iOS系统手机（16kHz，16bit）。AISHELL-2采用iOS系统手机录制的语音数据。1991名来自中国不同口音区域的发言人参与录制。经过专业语音校对人员转写标注，并通过严格质量检验，此数据库文本正确率在96%以上。（支持学术研究，未经允许禁止商用。）

[AISHELL-2 中文语音数据库申请链接](http://www.aishelltech.com/aishell_2)

**数据堂1505小时中文语音数据集（高校学术免费授权版）**

数据有效时长达1505小时，。录音内容超过3万条口语化句子，由6408名来自中国不同地区的录音人参与录制。经过专业语音校对及人员转写标注，通过严格质量检验，句准确率达98%以上，是行业内句准确率的最高标准。

[数据堂1050小时数据集申请获取链接](https://datatang.com/opensource)

### 说话人验证数据集

**cn-celeb**

此数据是“在野外”收集的大规模说话人识别数据集。该数据集包含来自1000位中国名人的13万种语音，涵盖了现实世界中的11种不同流派。所有音频文件都编码为单通道，并以16位精度以16kHz采样。数据收集过程由清华大学语音与语言技术中心组织。它也由国家自然科学基金61633013和博士后科学基金2018M640133资助。

**附下载链接：**

*   THCHS30

[国内镜像](http://cn-mirror.openslr.org/resources/18/data_thchs30.tgz) | [国外镜像](http://www.openslr.org/resources/18/data_thchs30.tgz)

*   ST-CMDS

[国内镜像](http://cn-mirror.openslr.org/resources/38/ST-CMDS-20170001_1-OS.tar.gz) | [国外镜像](http://www.openslr.org/resources/38/ST-CMDS-20170001_1-OS.tar.gz)

*   AISHELL开源版

[国内镜像](http://cn-mirror.openslr.org/resources/33/data_aishell.tgz) | [国外镜像](http://www.openslr.org/resources/33/data_aishell.tgz)

*   Primewords Chinese Corpus Set 1

[国内镜像](http://cn-mirror.openslr.org/resources/47/primewords_md_2018_set1.tar.gz) | [国外镜像](http://www.openslr.org/resources/47/primewords_md_2018_set1.tar.gz)

*   aidatatang

[国内镜像](http://cn-mirror.openslr.org/resources/62/aidatatang_200zh.tgz) | [国外镜像](http://www.openslr.org/resources/62/aidatatang_200zh.tgz)

*   Magic Data Mandarin Chinese Read Speech Corpus

[国内镜像(train)](http://cn-mirror.openslr.org/resources/68/train_set.tar.gz)  |  [国内镜像(dev)](http://cn-mirror.openslr.org/resources/68/dev_set.tar.gz)  |  [国内镜像(test)](http://cn-mirror.openslr.org/resources/68/test_set.tar.gz)  |  [国内镜像(metadata)](http://cn-mirror.openslr.org/resources/68/metadata.tar.gz)

[国外镜像(train)](http://www.openslr.org/resources/68/train_set.tar.gz)  |  [国外镜像(dev)](http://www.openslr.org/resources/68/dev_set.tar.gz) |  [国外镜像(test)](http://www.openslr.org/resources/68/test_set.tar.gz) | [国外镜像(metadata)](http://www.openslr.org/resources/68/metadata.tar.gz)

*   cn-celeb

[国内镜像](http://cn-mirror.openslr.org/resources/82/cn-celeb.tgz) | [国外镜像](http://www.openslr.org/resources/82/cn-celeb.tgz)

更多语音数据集，请访问：

[OpenSLR国内镜像](http://cn-mirror.openslr.org) | [OpenSLR国外镜像](http://www.openslr.org)

[AI柠檬OpenSLR镜像站(部分镜像)](https://dev.ailemon.me/openslr)


