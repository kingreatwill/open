# 损失函数（loss function）

损失函数（Loss Function ）是定义在单个样本上的，算的是一个样本的误差。

代价函数（Cost Function ）是定义在整个训练集上的，是所有样本误差的平均，也就是损失函数的平均。

目标函数（Object Function）定义为：最终需要优化的函数。等于经验风险+结构风险（也就是Cost Function + 正则化项）。

关于目标函数和代价函数的区别还有一种通俗的区别：

目标函数是最大化或者最小化，而代价函数是最小化

[Pytorch学习之十九种损失函数](https://blog.csdn.net/shanglianlm/article/details/85019768)
