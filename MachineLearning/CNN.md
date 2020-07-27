

[Tensorflow 2.0 手写全连接MNIST数据集（理论+实战）](https://blog.csdn.net/LQ_qing/article/details/99623008) 
[TensorFlow 2.0 (五) - mnist手写数字识别(CNN卷积神经网络)](https://geektutu.com/post/tensorflow2-mnist-cnn.html) 
[大白话讲解卷积神经网络工作原理 视频](https://www.bilibili.com/video/av35087157/) 
[大白话讲解卷积神经网络工作原理](https://zhuanlan.zhihu.com/p/49184702) 

[什么是CNN？写给小白的机器学习入门贴，Facebook员工打造，47k访问量](https://blog.csdn.net/QbitAI/article/details/107455120)

[CNNs，第1部分:卷积神经网络简介-原文](https://victorzhou.com/blog/intro-to-cnns-part-1/)

[卷积神经网络](https://baike.baidu.com/item/%E5%8D%B7%E7%A7%AF%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C/17541100)

[12种Dropout方法：应用于DNNs，CNNs，RNNs中的数学和可视化解释](https://blog.csdn.net/weixin_42137700/article/details/107248725)

[写给程序员的机器学习入门 (八) - 卷积神经网络 (CNN) - 图片分类和验证码识别](https://www.cnblogs.com/zkweb/p/13354826.html)

## 可视化

[卷积可视化](https://github.com/vdumoulin/conv_arithmetic)

[网站：三维可视化卷积神经网络](http://scs.ryerson.ca/~aharley/vis/conv/)

[3D模型可视化框架](https://tensorspace.org/html/playground/lenet.html)

TensorSpace 是一款3D 模型可视化框架
官网链接：

https://tensorspace.org/

Github链接：

https://github.com/tensorspace-team/tensorspace

## CNN系列模型
1. LeNet
2. AlexNet
3. VGG
4. GoogLeNet
5. ResNet
6. DenseNet
7. Non-Local Networks
8. Deformable Convolutional Networks
9. Dilated Convolutional Networks
10. SENET
[CNN系列模型发展简述（附github代码--已全部跑通）](https://zhuanlan.zhihu.com/p/66215918)

github代码依赖：
python 2.7, Pytorch 0.3.1
https://github.com/liuyuemaicha/cnn_model


## 概念

### 卷积和反卷积
[一文搞懂CNN中的卷积和反卷积](https://www.toutiao.com/i6642655643314422275)
**卷积(Convolutional)**：卷积在图像处理领域被广泛的应用，像滤波、边缘检测、图片锐化等，都是通过不同的卷积核来实现的。在卷积神经网络中通过卷积操作可以提取图片中的特征，低层的卷积层可以提取到图片的一些边缘、线条、角等特征，高层的卷积能够从低层的卷积层中学到更复杂的特征，从而实现到图片的分类和识别。

**反卷积**：反卷积也被称为转置卷积，反卷积其实就是卷积的逆过程。大家可能对于反卷积的认识有一个误区，以为通过反卷积就可以获取到经过卷积之前的图片，实际上通过反卷积操作并不能还原出卷积之前的图片，只能还原出卷积之前图片的尺寸。那么到底反卷积有什么作用呢？通过反卷积可以用来可视化卷积的过程，反卷积在GAN等领域中有着大量的应用。