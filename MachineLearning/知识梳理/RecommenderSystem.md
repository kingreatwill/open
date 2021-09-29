# 推荐系统 Recommender system

[推荐系统](../../BigData/推荐引擎.md)([recommender system](https://en.jinzhao.wiki/wiki/Recommender_system) or recommendation system) 有时也称为推荐平台或推荐引擎(platform or engine)，是[Information filtering system](https://en.jinzhao.wiki/wiki/Information_filtering_system)的一个子类。



## 方法Approaches或算法Recommended Algorithms

[Collaborative filtering](https://en.jinzhao.wiki/wiki/Collaborative_filtering)
Content-based filtering
Session-based recommender systems
Reinforcement learning for recommender systems
Multi-criteria recommender systems
Risk-aware recommender systems
[Mobile recommender systems](https://en.jinzhao.wiki/wiki/Location_based_recommendation)
Hybrid recommender systems


## 开源库介绍
### 基于显式反馈（explicit feedback）的数据集
比如评分相关的推荐算法[Surprise](http://surpriselib.com/)
数据格式：user item rating timestamp（时间戳可以不要）

#### 矩阵分解
由原始数据组成的用户-物品的评分矩阵（User-Item Matrix）

.|I1|I2|I3|I4|I5
---|---|---|---|---|---
A|5|4|4.5|?|3
B|?|4|?|4|?
C|4|?|3.5|4|3
D|?|5|?|?|3
E|4|?|4.5|5|?
我需要找到一个矩阵$\hat{R}$近似未知的矩阵$R$
$$R \approx \hat{R} = Q^TP$$

损失函数
$$\sum_{r_{ui} \in R_{train}} \left(r_{ui} - \hat{r}_{ui} \right)^2 $$

上面只是最简单的原理介绍
详细使用方法以及公式，见[Matrix Factorization-based algorithms](https://surprise.readthedocs.io/en/stable/matrix_factorization.html)

[Matrix Factorization: A Simple Tutorial and Implementation in Python](http://www.quuxlabs.com/blog/2010/09/matrix-factorization-a-simple-tutorial-and-implementation-in-python/)

#### k-NN
> collaborative filtering

基于邻域的协同过滤算法：knns.KNNBasic;knns.KNNWithMeans;
knns.KNNWithZScore;knns.KNNBaseline;
主要参数：
k:设置领域的个数；
user_based：是否为基于用户的协同过滤，True则为UserCF，False为ItemCF；
name：相似度计算方式，默认为MSD，也可设置为cosine，pearson，pearson_baseline

[符号解释](https://surprise.readthedocs.io/en/stable/notation_standards.html)
[sim_options参数](https://surprise.readthedocs.io/en/stable/prediction_algorithms.html#similarity-measures-configuration)

- KNNBasic
`surprise.prediction_algorithms.knns.KNNBasic(k=40, min_k=1, sim_options={}, verbose=True, **kwargs)`
**user_based=True**
$$\hat{r}_{ui} = \frac{
\sum\limits_{v \in N^k_i(u)} \text{sim}(u, v) \cdot r_{vi}}
{\sum\limits_{v \in N^k_i(u)} \text{sim}(u, v)}$$
其中$sim(u,v)$表示user相似度；
$N^k_i(u)$表示评价过item i的user中与当前user u 最相似的前k个user的集合；
**user_based=True**
$$\hat{r}_{ui} = \frac{
\sum\limits_{j \in N^k_u(i)} \text{sim}(i, j) \cdot r_{uj}}
{\sum\limits_{j \in N^k_u(i)} \text{sim}(i, j)}$$
其中$sim(i,j)$表示item相似度；
$N^k_u(i)$表示user u 评价过的item集合中与当前item i最相似的前k个item的集合。

- KNNWithMeans
`surprise.prediction_algorithms.knns.KNNWithMeans(k=40, min_k=1, sim_options={}, verbose=True, **kwargs)`
该方法认为，每个用户的打分跟用户本身的习惯有关，都是基于某一个特定值上下浮动。
$$\hat{r}_{ui} = \mu_u + \frac{ \sum\limits_{v \in N^k_i(u)}
\text{sim}(u, v) \cdot (r_{vi} - \mu_v)} {\sum\limits_{v \in
N^k_i(u)} \text{sim}(u, v)}$$
or
$$\hat{r}_{ui} = \mu_i + \frac{ \sum\limits_{j \in N^k_u(i)}
\text{sim}(i, j) \cdot (r_{uj} - \mu_j)} {\sum\limits_{j \in
N^k_u(i)} \text{sim}(i, j)}$$
其中$\mu_u$表示用户u所有评分的均值；


- KNNWithZScore
`surprise.prediction_algorithms.knns.KNNWithZScore(k=40, min_k=1, sim_options={}, verbose=True, **kwargs)`
该算法通过同时考虑均值和方差来对标准的KNN推荐算法进行改进
$$\hat{r}_{ui} = \mu_u + \sigma_u \frac{ \sum\limits_{v \in N^k_i(u)}
\text{sim}(u, v) \cdot (r_{vi} - \mu_v) / \sigma_v} {\sum\limits_{v
\in N^k_i(u)} \text{sim}(u, v)}$$
or
$$\hat{r}_{ui} = \mu_i + \sigma_i \frac{ \sum\limits_{j \in N^k_u(i)}
\text{sim}(i, j) \cdot (r_{uj} - \mu_j) / \sigma_j} {\sum\limits_{j
\in N^k_u(i)} \text{sim}(i, j)}$$
其中$\sigma_u$表示用户u所有评分的方差；


- KNNBaseline
`surprise.prediction_algorithms.knns.KNNBaseline(k=40, min_k=1, sim_options={}, bsl_options={}, verbose=True, **kwargs)`
该算法不再使用传统的统计指标，而是允许使用自定义的计算方法来为每一个用户计算一个一个基准值。
$$\hat{r}_{ui} = b_{ui} + \frac{ \sum\limits_{v \in N^k_i(u)}
\text{sim}(u, v) \cdot (r_{vi} - b_{vi})} {\sum\limits_{v \in
N^k_i(u)} \text{sim}(u, v)}$$
or
$$\hat{r}_{ui} = b_{ui} + \frac{ \sum\limits_{j \in N^k_u(i)}
\text{sim}(i, j) \cdot (r_{uj} - b_{uj})} {\sum\limits_{j \in
N^k_u(i)} \text{sim}(i, j)}$$

#### SlopeOne
> collaborative filtering

$$\hat{r}_{ui} = \mu_u + \frac{1}{
|R_i(u)|}
\sum\limits_{j \in R_i(u)} \text{dev}(i, j),  $$
$$\text{dev}(i, j) = \frac{1}{
|U_{ij}|}\sum\limits_{u \in U_{ij}} r_{ui} - r_{uj}$$

例如下面表格里有3个用户对4个物品的评分
|       | 101 | 102 | 103 | 104 |
|-------|-----|-----|-----|-----|
| UserX | 5   | 3.5 |     |     |
| UserY | 2   | 5   | 4   | 2   |
| UserZ | 4.5 | 3.5 | 1   | 4   |
求物品两两之间的差值平均分:
物品102和101：{(3.5-5)+(5-2)+(3.5-4.5)}/3=0.5/3
物品103跟101：{(4-2)+(1-4.5)}/2=-1.5/2
物品104跟101：{(2-2)+(4-4.5)}/2=-0.5/2
物品103跟102：{(4-5)+(1-3.5)}/2=-3.5/2
物品104跟102：{(2-5)+(4-3.5)}/2=-2.5/2
物品104跟103：{(2-4)+(4-1)}/2=1/2
能得到下面表格
|     | 101   | 102   | 103 | 104 |
|-----|-------|-------|-----|-----|
| 101 |       |       |     |     |
| 102 | 0.17  |       |     |     |
| 103 | -0.75 | -1.75 |     |     |
| 104 | -0.25 | -1.25 | 0.5 |     |
OK，现在准备工作已经完成了，然后进行推荐，例如对X用户进行推荐，103和104个预测评分根据101、102的评分来的。
X预测103评分={(-0.75+5)+(-1.75+3.5)}/2=(4.25+1.75)/2=3
X预测104评分={(-0.25+5)+(-1.25+3.5)}/2=(4.75+2.25)/2=3.5
那么给X用户推荐的顺序就是：先推荐104在推荐103

[SlopeOne 算法基本原理](https://blog.csdn.net/redhatforyou/article/details/86656356)

#### CoClustering
> collaborative filtering

$$\hat{r}_{ui} = \overline{C_{ui}} + (\mu_u - \overline{C_u}) + (\mu_i- \overline{C_i})$$

协同聚类${C_{ui}}$ 的均值(聚类质心)$\overline{C_{ui}}$ 
聚类${C_u}$（基于user） 的均值(聚类质心)$\overline{C_{u}}$ 
聚类${C_i}$（基于item） 的均值(聚类质心)$\overline{C_{i}}$ 

```
cc = CoClustering(n_cltr_u=3,n_cltr_i =2)
cc.avg_cltr_i.shape # (2,)
cc.avg_cltr_u.shape # (3,)
cc.avg_cocltr.shape # (3,2)
```

### 基于隐式反馈（implicit feedback）的数据集
> Collaborative Filtering

比如根据用户行为相关的推荐算法[implicit](https://implicit.readthedocs.io/en/latest/index.html)

#### 论文
[Collaborative Filtering for Implicit Feedback Datasets](http://yifanhu.net/PUB/cf.pdf)
[Applications of the Conjugate Gradient Method for Implicit Feedback Collaborative Filtering.](http://www.sze.hu/~gtakacs/download/recsys_2011_draft.pdf)