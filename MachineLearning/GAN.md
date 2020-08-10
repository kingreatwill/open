# GAN
生成式对抗网络（GAN, Generative Adversarial Networks ）是一种深度学习模型，是近年来复杂分布上无监督学习最具前景的方法之一。
GAN是机器学习近十年来最有意思的想法，GAN，即生成对抗网络，是一个生成模型，也是半监督和无监督学习模型，它可以在不需要大量标注数据的情况下学习深度表征。最大的特点就是提出了一种让两个深度网络对抗训练的方法。
## xx
### Generative（Model）：生成模型

机器学习的模型可大体分为两类: 
- 生成模型（Generative Model）
生成模型是给定某种隐含信息，来随机产生观测数据。
- 判别模型（Discriminative Model）
判别模型需要输入变量x ，通过某种模型来预测p(y|x)

判别模型：给定一张图，判断这张图里的动物是猫还是狗
生成模型：给一系列猫的图片，生成一张新的猫咪（不在数据集里）

对于判别模型，损失函数是容易定义的，因为输出的目标相对简单。但对于生成模型，损失函数的定义就不是那么容易。
例如对于NLP方面的生成语句，虽然有BLEU这一优秀的衡量指标，但由于难以求导，以至于无法放进模型训练；对于生成猫咪图片的任务，如果简单地将损失函数定义为“和已有图片的欧式距离”，那么结果将是数据库里图片的诡异混合，效果惨不忍睹。
当我们希望神经网络画一只猫的时候，显然是希望这张图有一个动物的轮廓、带质感的毛发、和一个霸气的眼神，而不是冷冰冰的欧式距离最优解。
如何将我们对于猫的期望放到模型中训练呢？
这就是GAN的Adversarial部分解决的问题。

### Adversarial：对抗（互怼 ）
生成模型的回馈部分，交给判别模型，Generative和Discrimitive给紧密地联合在了一起。
### Networks：（深度）神经网络
创作艺术品（Gatys 的 Neural Alorightm for Artistic Style），AlphaGo（CNN估值 + 蒙特卡洛剪枝），高质量的机器翻译（Attention + seq2seq）

## 其它
### gan-playground

https://github.com/reiinakano/gan-playground
```
$ git clone https://github.com/reiinakano/gan-playground.git
$ cd gan-playground
$ npm install && bower install # Install node modules and bower components
```

### GAN Image-to-Image
https://affinelayer.com/pixsrv/index.html

### gans-awesome
https://github.com/nashory/gans-awesome-applications

### gan paper
https://github.com/hindupuravinash/the-gan-zoo

### GAN图像生成器—BigGAN

https://github.com/sxhxliang/BigGAN-pytorch
https://github.com/ajbrock/BigGAN-PyTorch

https://github.com/ajbrock/Neural-Photo-Editor

## 参考
[生成对抗网络（GAN）的半监督学习](https://www.toutiao.com/i6853711638252651022/)