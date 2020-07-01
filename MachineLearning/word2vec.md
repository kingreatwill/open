#  word2vec

https://radimrehurek.com/gensim/models/word2vec.html
https://radimrehurek.com/gensim/models/doc2vec.html

## word2vec

[腾讯开源的word2vec模型](https://ai.tencent.com/ailab/nlp/en/index.html)
https://ai.tencent.com/ailab/nlp/en/embedding.html

搜狗实验室(没有更新)：http://www.sogou.com/labs/resource/ca.php
使用自己的语料训练word2vec模型：https://cloud.tencent.com/developer/article/1327481
利用gensim使用腾讯开源的预训练好的词向量: https://www.jianshu.com/p/bba1bf9518dc
CS224n笔记2 词的向量表示：word2vec:https://www.hankcs.com/nlp/word-vector-representations-word2vec.html

https://cloud.tencent.com/developer/article/1418647
https://www.ctolib.com/mip/cliuxinxin-TX-WORD2VEC-SMALL.html

### 为什么要进行embedding
word2vec就是对word进行embedding

首先,我们知道,在机器学习和深度学习中,对word的最简单的表示就是使用`one-hot([0,0,1,0,0…..]`来表示一个word). 但是用one-hot表示一个word的话,会有一些弊端:从向量中无法看出word之间的关系`((wworda)Twwordb=0(w^{word_a})^Tw^{word_b}=0)`,而且向量也太稀疏. 所以一些人就想着能否用更小的向量来表示一个word,希望这些向量能够承载一些语法和语义上的信息, 这就产生了word2vec

> 智能手机键盘进行下一词预测, 分析语义

## doc2vec