# 高引用论文

## 机器学习TOP 100论文：
https://github.com/terryum/awesome-deep-learning-papers#understanding--generalization--transfer

## 理解、泛化、迁移学习
1、Distilling the knowledge in a neural network (2015), G. Hinton et al.
http://arxiv.org/pdf/1503.02531

这篇介绍了Hinton大神在15年做的一个黑科技技术，Hinton在一些报告中称之为Dark Knowledge，技术上一般叫做知识蒸馏（Knowledge Distillation）。这篇论文的核心思想是通过迁移知识，从而以训练好的大模型得到更加适合推理的小模型。

2、Deep neural networks are easily fooled: High confidence predictions for unrecognizable images (2015), A. Nguyen et al.
http://arxiv.org/pdf/1412.1897

研究结果揭示了人的视觉和目前DNNs的差异。具体来说，卷积神经网络在ImageNet或MNIST数据集上训练都表现良好，但发现通过进化算法或梯度上升处理的图片，DNNs以很高的置信度贴以标签属于某个数据集类（其实不属于这个数据集类）。

3、How transferable are features in deep neural networks? (2014), J. Yosinski et al.
http://papers.nips.cc/paper/5347-how-transferable-are-features-in-deep-neural-networks.pdf

本文通过实验，量化了深度神经网络每层神经元的通用性与特殊性，并对结果进行了展示。网络第一层的特征并非特定于某一数据集或者某一任务，而是通用的特征，它们适用于许多数据集和普遍的任务。在较深的模型层，特征会从通用的特征逐渐转换为更专业的特征（和任务、数据集紧密相关的特征）。

4、CNN features off-the-Shelf: An astounding baseline for recognition (2014), A. Razavian et al.
http://www.cv-foundation.org//openaccess/content_cvpr_workshops_2014/W15/papers/Razavian_CNN_Features_Off-the-Shelf_2014_CVPR_paper.pdf

本文考虑了一种问题，假设有一个现成的，针对某个具体问题A训练好的CNN，仅仅使用它的前几层来提取图像信息，再配合使用一些经典分类器（SVM等），是否可以在其他的问题B,C上也得到比较好的结果？

5、Learning and transferring mid-Level image representations using convolutional neural networks (2014), M. Oquab et al.
https://www.cv-foundation.org/openaccess/content_cvpr_2014/papers/Oquab_Learning_and_Transferring_2014_CVPR_paper.pdf

CNN的学习需要建立数以百万计的参数，并且需要大量已经标注好的图像。这种特性目前阻止了CNN在有限训练集问题上的应用。本文展示了在大规模标记的数据上、用CNN学习出的图像表示，是如何有效地被迁移到其他视觉识别的任务中的。

6、Visualizing and understanding convolutional networks (2014), M. Zeiler and R. Fergus
http://arxiv.org/pdf/1311.2901

这篇论文的目的，就是通过特征可视化，查看精度变化，从而知道CNN学习到的特征如何。这篇论文阐述了CNN的每一层到底学习到了什么特征，然后作者通过可视化进行调整网络。

7、Decaf: A deep convolutional activation feature for generic visual recognition (2014), J. Donahue et al.
http://arxiv.org/pdf/1310.1531

这篇论文验证了卷积特征在各种场合上的效果，算是transfer learning和一些验证的论文。而且，DeCAF可以算是著名的框架Caffe的前身。

## 优化、技巧方法
8、Training very deep networks (2015), R. Srivastava et al.
http://papers.nips.cc/paper/5850-training-very-deep-networks.pdf

作者提出了一种全新的高速网络结构 (Highway Networks)，用于优化深度神经网络由于梯度爆炸和梯度消失而导致的训练困难的问题。而且，ResNet 的思路和这篇文章所提出的想法有很多相似之处。（小tips，这篇论文发表于 2015 年 05 月份，ResNet 发表于 2015 年 12 月份）

9、Batch normalization: Accelerating deep network training by reducing internal covariate shift (2015), S. Loffe and C. Szegedy
http://arxiv.org/pdf/1502.03167

这篇文章引入了BN层，并介绍了引入原因。引入 BN 后，我们可以不用太在意参数的初始化，同时使用更大的学习率，而且也会有正则化的效果，在一些情况下可以不用再使用 Dropout。

10、Delving deep into rectifiers: Surpassing human-level performance on imagenet classification (2015), K. He et al.
http://www.cv-foundation.org/openaccess/content_iccv_2015/papers/He_Delving_Deep_into_ICCV_2015_paper.pdf

这篇论文是来自MSRA的何恺明的论文，论文首次公开宣布图像的识别率超越人类水平。

11、Dropout: A simple way to prevent neural networks from overfitting (2014), N. Srivastava et al.
http://jmlr.org/papers/volume15/srivastava14a/srivastava14a.pdf

大牛集结的论文，Hinton、Bengio都有参与。这篇文章对dropout进行了研究，结果表明，在视觉、语音识别、文档分类和计算生物学等方面，dropout都能提高神经网络在有监督学习任务中的性能，在许多基准数据集上都获得了最新的结果。

12、Adam: A method for stochastic optimization (2014), D. Kingma and J. Ba
http://arxiv.org/pdf/1412.6980

本文展示了如何将优化算法的设计转换为一个学习问题，使算法能够自动地在感兴趣的问题中利用结构。文中的学习算法由LSTMs实现。

13、Improving neural networks by preventing co-adaptation of feature detectors (2012), G. Hinton et al.
http://arxiv.org/pdf/1207.0580.pdf

Hinton的论文，文章对过拟合问题进行了研究。训练网络时，随机忽略一半的feature detectors能够防止因训练集太小带来的过拟合问题。这能够防止一些detectors联合在一起才起作用的情况，每个神经元预测一个特征有利于提高准确率，这种dropout的方法能提高很多benchmark的成绩。

14、Random search for hyper-parameter optimization (2012) J. Bergstra and Y. Bengio
http://www.jmlr.org/papers/volume13/bergstra12a/bergstra12a

Bengio的论文，关于超参数优化的方法。论文指出，Random Search比Gird Search更有效。实际操作的时候，一般也是先用Gird Search的方法，得到所有候选参数，然后每次从中随机选择进行训练。

## 无监督学习、生成模型
15、Pixel recurrent neural networks (2016), A. Oord et al.
http://arxiv.org/pdf/1601.06759v2.pdf

本文提出了一个深度神经网络，它根据顺序沿着两个空间维度来预测图片中的像素。这种模型离散了原始像素值的可能性，同时编码保证了整个图片的完整性。对自然图片的分布进行建模一直以来都是无监督学习中的里程碑式的难题。这要求图片模型易表达、易处理、可拓展。

16、Improved techniques for training GANs (2016), T. Salimans et al.
http://papers.nips.cc/paper/6125-improved-techniques-for-training-gans.pdf

本文提出了可以用到GAN上的一些新的结构特征和训练过程。本文主要应用于半监督学习和生成视觉上真实的图像两个方向。使用这种方法，可以在MNIST，CIFAR10，SVHN上达到很好的半监督效果。

17、Unsupervised representation learning with deep convolutional generative adversarial networks (2015), A. Radford et al.
https://arxiv.org/pdf/1511.06434v2

这篇论文旨在帮助缩小监督学习和非监督学习成功运用于CNN上的差距。论文介绍了CNN的一个类，称为深度卷积生成对抗网络（DCGANs），这个网络有着明确的结构约束，并且表明他们对非监督学习有着强烈的可信度。

18、DRAW: A recurrent neural network for image generation (2015), K. Gregor et al.
http://arxiv.org/pdf/1502.04623

本文介绍了深度递归书写器（DRAW）神经网络用于图像生成。DRAW网络是一种模仿人眼空间注意力机制的、带有视觉偏好性的可变自动编码框架，其主要功能是用于复杂图像的迭代构造。

19、Generative adversarial nets (2014), I. Goodfellow et al.
http://papers.nips.cc/paper/5423-generative-adversarial-nets.pdf

GANs来了。论文提出了一个通过对抗过程估计生成模型的新框架，在新框架中同时训练两个模型：一个用来捕获数据分布的生成模型G，和一个用来估计样本来自训练数据而不是G的概率的判别模型D，G的训练过程是最大化D产生错误的概率。在训练或生成样本期间不需要任何马尔科夫链或展开的近似推理网络。

20、Auto-encoding variational Bayes (2013), D. Kingma and M. Welling
http://arxiv.org/pdf/1312.6114

AEV与GAN是现在生成网络中的两个趋势。文中引入了随机变分推理和学习算法，扩展到大数据集，并且可以在一些温和的差异性条件下、甚至某些棘手的情况下工作。论文表明，变分下界的重新参数化产生了可以使用标准随机梯度法直接优化的下限估计器。

21、Building high-level features using large scale unsupervised learning (2013), Q. Le et al.
http://arxiv.org/pdf/1112.6209

GoogleBrain中特征学习的原理，通过使用未标记的图像学习人脸、猫脸high-level特征，得到检测器。文章使用大数据构建了一个9层的局部连接稀疏自编码网络，使用模型并行化和异步SGD在1000个机器（16000核）上训练了3天，结果显示，可以在未标记图像是否有人脸的情况下训练出一个人脸检测器。

由于文章比较多，此处只介绍前20篇论文，除此之外，还有卷积神经网络模型、目标检测、视频图像处理、NLP算法、RNN模型、强化学习和机器人领域等近年来最经典的论文。