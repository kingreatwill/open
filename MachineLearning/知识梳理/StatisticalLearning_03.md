[TOC]
## 第 13 章 无监督学习概论
**基本问题**：
**聚类**（clustering）是将样本集合中相似的样本归到相同的类，相似的定义一般用距离度量。
如果一个样本只能属于一个类，则称为硬聚类（hard clustering），如k-means；如果一个样本可以属于多个类，每一个样本以概率属于每一个类$\sum_{i=1}^N p(z_i|x_i) =1$，则称为软聚类（soft clustering），如GMM。
聚类主要用于数据分析，也可用于监督学习的前处理。聚类可以帮助发现数据中的统计规律。


**降维**（dimensionality reduction）是将样本集合中的样本（实例）从高维空间转换到低维空间。
高维空间通常是高维的欧氏空间，而低维空间是低维的欧氏空间或流形（manifold）。低维空间是从数据中自动发现的。降维有线性降维和非线性降维，降维方法有主成分分析。
降维的好处有：节省存储空间、加速计算、解决**维度灾难**（前面章节有讲到）等
降维主要用于数据分析，也可用于监督学习的前处理。降维可以帮助发现高维数据中的统计规律。

**概率模型估计**（probability model estimation），简称概率估计，假设训练数据由一个概率模型生成，由训练数据学习概率模型的结构和参数。
概率模型的结构类型或者概率模型的集合事先给定，而模型的具体结构与参数从数据中自动学习。假设数据由GMM生成（已知结构），学习的目标是估计这个模型的参数。
概率模型包括混合模型、概率图模型等。概率图模型又包括有向图模型和无向图模型（前面章节有讲到）。

**无监督学习方法**
- 聚类
    - 硬聚类：
        - k-means
    - 软聚类：
        - GMM
- 降维
    - 线性：
        - 主成分分析
    - 非线性：
        - 流形学习

- 话题分析
话题分析是文本分析的一种技术。给定一个文本集合，话题分析旨在发现文本集合中每个文本的话题，而话题由单词的集合表示。话题分析方法有：
    - 潜在语义分析
    - 概率潜在语义分析
    - 潜在狄利克雷分配

- 图分析
图分析 的目的是 发掘隐藏在图中的统计规律或潜在结构；
    - 链接分析 是图分析的一种，主要是发现 有向图中的重要结点，包括 PageRank 算法
    - PageRank 算法最初是为互联网搜索而提出。将互联网看作是一个巨大的有向图，网页是结点，网页的超链接是有向边。PageRank 算法可以算出网页的 PageRank 值，表示其重要度，在搜索引擎的排序中网页的重要度起着重要作用

同监督学习一样，无监督学习也有**三要素**：模型、策略、算法
**模型** 就是函数$z=g_\theta(x)$，条件概率分布$P_\theta(z |x)$，或$P_\theta(x|z)$，在聚类、降维、概率模型估计中拥有不同的形式
- 聚类 中模型的输出是 类别
- 降维 中模型的输出是 低维向量
- 概率模型估计 中的模型可以是混合概率模型，也可以是有向概率图模型和无向概率图模型

**策略** 在不同的问题中有不同的形式，但都可以表示为目标函数的优化
- 聚类 中样本与所属类别中心距离的最小化
- 降维 中样本从高维空间转换到低维空间过程中信息损失的最小化
- 概率模型估计 中模型生成数据概率的最大化

**算法** 通常是迭代算法，通过迭代达到目标函数的最优化，比如，梯度下降法。
- 层次聚类法、k均值聚类 是硬聚类方法
- 高斯混合模型 EM算法是软聚类方法
- 主成分分析、潜在语义分析 是降维方法
- 概率潜在语义分析、潜在狄利克雷分配 是概率模型估计方法



### 参考文献
[13-1] Hastie T, Tibshirani R, Friedman J. The elements of statistical learning:data mining, inference, and prediction. Springer. 2001. 

[13-2] Bishop M. Pattern Recognition and Machine Learning. Springer, 2006.

[13-3] Koller D, Friedman N. Probabilistic graphical models: principles and techniques. Cambridge, MA: MIT Press, 2009.

[13-4] Goodfellow I,Bengio Y,Courville A. Deep learning. Cambridge, MA: MIT Press, 2016.

[13-5] Michelle T M. Machine Learning. McGraw-Hill Companies, Inc. 1997.（中译本：机器学习。北京：机械工业出版社，2003.）

[13-6] Barber D. Bayesian reasoning and machine learning, Cambridge, UK:Cambridge University Press, 2012.

[13-7] 周志华. 机器学习. 北京：清华大学出版社，2017.

## 第 14 章 聚类方法

聚类分析（[Cluster analysis or clustering](https://en.jinzhao.wiki/wiki/Cluster_analysis)）是针对给定的样本，根据它们特征点的相似度（距离），将其归并到若干个类（簇）中的分析问题。聚类分析本身不是一种特定的算法，而是要解决的一般任务。

聚类分析算法（[Cluster analysis algorithms](https://en.jinzhao.wiki/wiki/Category:Cluster_analysis_algorithms)）：[sklearn中的聚类算法和介绍](https://scikit-learn.org/stable/modules/clustering.html#clustering)
**基于连接的聚类（层次聚类）** Connectivity-based clustering（[Hierarchical clustering](https://en.jinzhao.wiki/wiki/Hierarchical_clustering)）：
Agglomerative

**基于质心的聚类** Centroid-based clustering：
K均值聚类（[k-means clustering](https://en.jinzhao.wiki/wiki/K-means_clustering)）

**基于分布的聚类** Distribution-based clustering：
高斯混合模型聚类（[Gaussian mixture model](https://en.jinzhao.wiki/wiki/Mixture_model#Gaussian_mixture_model)）

**基于密度的聚类** Density-based clustering：
https://en.jinzhao.wiki/wiki/DBSCAN
https://en.jinzhao.wiki/wiki/OPTICS_algorithm
https://en.jinzhao.wiki/wiki/Mean-shift


**基于网格的聚类** Grid-based clustering：
STING 和 CLIQUE

高维数据的聚类
https://en.jinzhao.wiki/wiki/Clustering_high-dimensional_data

https://en.jinzhao.wiki/wiki/Kernel_density_estimation

- **模型**：
- **策略**：
- **算法**：
### 附加知识
### 参考文献

## 第 15 章 奇异值分解
- **模型**：
- **策略**：
- **算法**：
### 附加知识
### 参考文献

## 第 16 章 主成分分析
- **模型**：
- **策略**：
- **算法**：
### 附加知识
### 参考文献


## 第 17 章 潜在语义分析
- **模型**：
- **策略**：
- **算法**：
### 附加知识
### 参考文献

## 第 18 章 概率潜在语义分析
- **模型**：
- **策略**：
- **算法**：
### 附加知识
### 参考文献

## 第 19 章 马尔科夫链蒙特卡罗法
- **模型**：
- **策略**：
- **算法**：
### 附加知识
### 参考文献

## 第 20 章 潜在狄利克雷分配
- **模型**：
- **策略**：
- **算法**：
### 附加知识
### 参考文献

## 第 21 章 PageRank算法
- **模型**：
- **策略**：
- **算法**：
### 附加知识
### 参考文献

## 第 22 章 无监督学习方法总结
### 附加知识
### 参考文献