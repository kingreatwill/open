## 第 10 章 隐马尔可夫模型

**隐马尔可夫模型**（[Hidden Markov Model,HMM](https://en.jinzhao.wiki/wiki/Hidden_Markov_model)）是可用于**标注问题**的统计学习模型，描述由隐藏的马尔可夫链随机生成观测序列的过程，属于生成模型。

马尔可夫模型是关于时序的概率模型，描述由一个隐藏的马尔可夫链随机生成**不可观测的状态随机序列**，再由各个状态生成一个观测而产生观测随机序列的过程。
隐藏的马尔可夫链随机生成的状态的序列，称为**状态序列**（state sequence）；每个状态生成一个观测，而由此产生的观测的随机序列，称为**观测序列**（observation sequence）。序列的每一个位置又可以看作是一个时刻。

隐马尔可夫模型由初始概率分布、状态转移概率分布以及观测概率分布确定。
隐马尔可夫模型的形式定义如下：
设 Q 是所有可能的状态的集合，N 是可能的状态数；V 是所有可能的观测的集合，M 是可能的观测数。
$$Q = \{q_1,q_2,...,q_N\} , V= \{v_1,v_2,...,v_M\}$$
长度为 T 的状态序列$I = (i_1,i_2,...,i_T)$以及与状态序列对应的长度为 T 的观测序列$O = (o_1,o_2,...,o_T)$

**状态转移矩阵(状态转移概率分布)**：
$$A=[a_{ij}]_{N\times N}$$
其中$a_{ij} = P(i_{t+1} = q_j | i_t = q_i) ,下标i,j = 1,...,N$表示在时刻$t$处于状态$q_i$的条件下 在时刻$t+1$转移到状态$q_j$的概率

**观测矩阵(观测概率分布)**：
$$B = [b_j(k)]_{N \times M}$$
其中$b_j(k) = P(o_t = v_k | i_t = q_j) ,k = 1,...,M,j = 1,...,N$表示在时刻$t$处于状态$q_j$的条件下生成观测$v_k$的概率

**初始状态概率向量（初始概率分布）**：
$$\pi = (\pi_i)$$
其中$\pi_i = P(i_1 =q_i) ,下标i = 1,...,N$表示时刻$t=1$时 处于状态$q_i$的概率

隐马尔可夫模型由初始状态概率向量$\pi$、状态转移概率矩阵$A$和观测概率矩阵$B$决定。
$\pi$和$A$决定状态序列，$B$决定观测序列。
因此，**一个隐马尔可夫模型**可以用三元符号表示，即
$$\lambda = (A,B,\pi)$$
称为隐马尔可夫模型的三要素。

状态转移概率矩阵$A$与初始状态概率向量$\pi$确定了隐藏的马尔可夫链，生成不可观测的状态序列。观测概率矩阵$B$确定了如何从状态生成观测，与状态序列综合确定了如何产生观测序列。

从定义可知，隐马尔可夫模型作了**两个基本假设**：

1. **齐次马尔可夫性假设**，即假设隐藏的马尔可夫链在任意时刻 t 的状态只依赖于其前一时刻的状态，与其他时刻的状态及观测无关，也与时刻 t 无关。
   $$P(i_{t}|i_{t-1},o_{t-1},...,i_{1},o_{1}) = P(i_{t}|i_{t-1}), t=1,2,...,T$$
1. **观测独立性假设**，即假设任意时刻的观测只依赖于该时刻的马尔可夫链的状态，与其他观测及状态无关。
   $$P(o_{t}|i_{T},o_{T},i_{T-1},o_{T-1},...,i_{t+1},o_{t+1},i_{t},o_{t},i_{t-1},o_{t-1},...,i_{1},o_{1}) = P(o_{t}|i_{t})$$

隐马尔可夫模型的**三个基本问题**：

1. 概率计算问题。给定模型$\lambda = (A,B,\pi)$和观测序列$O = (o_1,o_2,...,o_T)$，计算在模型$\lambda$下观测序列$O$出现的概率$P(O|\lambda)$。

1. 学习问题。已知观测序列$O = (o_1,o_2,...,o_T)$，估计模型$\lambda = (A,B,\pi)$参数，使得在该模型下观测序列概率$P(O|\lambda)$最大。即用极大似然估计的方法估计参数。（$\lambda_{MLE}=\argmax_{\lambda}P(O|\lambda)$，使用 EM 算法求解。）

1. 预测问题，也称为解码（decoding）问题。已知模型$\lambda = (A,B,\pi)$和观测序列$O = (o_1,o_2,...,o_T)$，求对给定观测序列条件概率$P(I|O)$最大的状态序列$I = (i_1,i_2,...,i_T)$。即给定观测序列，求最有可能的对应的状态序列。（Viterbi 算法求$\hat{I}=\argmax_{I}P(I|O,\lambda)$）

**概率计算问题**：
引入隐变量，对完全数据建模(这里还是一样$P(O|\lambda),P(O;\lambda)$是一样的，$\lambda$是参数)
$$P(O|\lambda) = \sum_{I}P(O,I|\lambda)= \sum_{I}P(O|I,\lambda)P(I|\lambda)$$
根据乘法规则（概率论基础教程 51 页，注意$P(i_1|\lambda) = P(i_1)$）以及马尔可夫假设有：
$$P(I|\lambda) = P(i_1,i_2,...,i_T|\lambda)=P(i_1).P(i_2|i_1,\lambda).P(i_3|i_1,i_2,\lambda)...P(i_T|i_1,i_2,...,i_{T-1},\lambda) \\= P(i_1)\prod_{t=2}^T P(i_t|i_1,i_2,...,i_{t-1},\lambda) \\= P(i_1)\prod_{t=2}^T P(i_t|i_{t-1},\lambda) \\= \pi_{i_1}\prod_{t=2}^T a_{i_{t-1}i_{t}}$$
根据乘法规则以及观测独立性假设有：
$$P(O|I,\lambda) = P(o_1,o_2,...,o_T|i_1,i_2,...,i_{T},\lambda) \\= P(o_1|i_1,i_2,...,i_{T},\lambda).P(o_2|o_1,i_1,i_2,...,i_{T},\lambda).P(o_3|o_1,o_2,i_1,i_2,...,i_{T},\lambda)...P(o_T|o_1,o_2,...,o_{T-1},i_1,i_2,...,i_{T},\lambda) \\ = P(o_1|i_1,\lambda).P(o_2|i_2,\lambda)...P(o_T|i_T,\lambda) \\= \prod_{t=1}^Tb_{i_t}(o_t)$$
那么

$$
P(O,I|\lambda) = P(O|I,\lambda)P(I|\lambda) = \pi_{i_1}\prod_{t=2}^T a_{i_{t-1}i_{t}}\prod_{t=1}^Tb_{i_t}(o_t)
\\= \pi_{i_1}b_{i_1}(o_1) .a_{i_1i_2}b_{i_2}(o_2)...a_{i_{T-1}i_T}b_{i_T}(o_T) = \pi_{i_1}b_{i_1}(o_1)\prod_{t=2}^T a_{i_{t-1}i_{t}}b_{i_t}(o_t)
$$

**概率计算问题- 直接由上面计算概率**可得
$$P(O|\lambda) = \sum_{I}P(O,I|\lambda)= \sum_{I}P(O|I,\lambda)P(I|\lambda) \\= \sum_{i_1,i_2,...,i_T} \pi_{i_1}b_{i_1}(o_1)\prod_{t=2}^T a_{i_{t-1}i_{t}}b_{i_t}(o_t) \\= \sum_{i_1 \in N}...\sum_{i_T\in N} \pi_{i_1}b_{i_1}(o_1)\prod_{t=2}^T a_{i_{t-1}i_{t}}b_{i_t}(o_t) $$
时间复杂度$O(TN^{T})$，所以不可行。

上面说过直接求不好求，有以下方法可求得：
**概率计算问题- 前向计算**：
首先我们定义**前向概率**$\alpha_t(i) = P(o_1,o_2,...,o_t,i_t=q_i | \lambda)$，表示时刻$t$部分观测序列为$o_1,o_2,...,o_t$且状态为$q_i$的概率，那么
$$P(O|\lambda) = \sum_{i=1}^N P(O,i_T=q_i|\lambda) = \sum_{i=1}^N P(o_1,...,o_T,i_T=q_i|\lambda) = \sum_{i=1}^N \alpha_T(i)$$

> 其实$P(O|\lambda) = \sum_{j=1}^N P(O,i_1=q_j|\lambda) =...= \sum_{j=1}^N P(O,i_t=q_j|\lambda) = \sum_{i=1}^N\sum_{j=1}^N P(O,i_1=q_i,i_2=q_j|\lambda)$，注意这里是小 t，只不过我们定义了前向概率，并且$O=(o_1,...,o_T)$

所以我们只要求出$\alpha_T(i)$，如何求？依次$\alpha_1(i) ... \alpha_{t+1}(i) ... \alpha_T(i)$

$$\alpha_1(i) = P(o_1,i_1=q_i | \lambda) =P(i_1=q_i | \lambda)P(o_1|i_1=q_i , \lambda) = \pi_ib_i(o_1) \\  \vdots\\ \alpha_{t+1}(i) = P(o_1,o_2,...,o_t,o_{t+1},i_{t+1}=q_i | \lambda)  \\=\sum_{j=1}^N P(o_1,o_2,...,o_t,o_{t+1},i_{t+1}=q_i,i_{t}=q_j | \lambda) \\ =\sum_{j=1}^NP(o_{t+1}|o_1,..,o_t,i_{t+1}=q_i,i_{t}=q_j,\lambda)P(o_1,o_2,...,o_t,i_{t+1}=q_i,i_{t}=q_j | \lambda) \\=\sum_{j=1}^NP(o_{t+1}|i_{t+1}=q_i)P(o_1,o_2,...,o_t,i_{t+1}=q_i,i_{t}=q_j | \lambda)  \\= \sum_{j=1}^NP(o_{t+1}|i_{t+1}=q_i)P(i_{t+1}=q_i | o_1,o_2,...,o_t,i_{t}=q_j,\lambda)P(o_1,o_2,...,o_t,i_{t}=q_j | \lambda)  \\=\sum_{j=1}^NP(o_{t+1}|i_{t+1}=q_i)P(i_{t+1}=q_i | i_{t}=q_j,\lambda)P(o_1,o_2,...,o_t,i_{t}=q_j | \lambda) \\=\sum_{j=1}^NP(o_{t+1}|i_{t+1}=q_i)P(i_{t+1}=q_i | i_{t}=q_j,\lambda)\alpha_t(j) \\= P(o_{t+1}|i_{t+1}=q_i)\sum_{j=1}^NP(i_{t+1}=q_i | i_{t}=q_j,\lambda)\alpha_t(j) \\= \bigg[\sum_{j=1}^N\alpha_t(j)a_{ji} \bigg]  b_i(o_{t+1})$$

**概率计算问题- 后向计算**：
首先我们定义**后向概率**$\beta_t(i) = P(o_{t+1},o_{t+2},...,o_T|i_t=q_i , \lambda)$，表示时刻状态为$q_i$的条件下，从$t+1$到$T$的部分观测序列为$o_{t+1},o_{t+2},...,o_T$概率，那么
$$P(O|\lambda) = \sum_{i=1}^N P(O,i_1=q_i|\lambda) = \sum_{i=1}^N P(o_1,...,o_T,i_1=q_i|\lambda) \\= \sum_{i=1}^N P(o_1,...,o_T|i_1=q_i,\lambda)P(i_1=q_i|\lambda) \\= \sum_{i=1}^N P(o_1|o_2,...,o_T,i_1=q_i,\lambda)P(o_2,...,o_T|i_1=q_i,\lambda)P(i_1=q_i|\lambda) \\ = \sum_{i=1}^N P(o_1|i_1=q_i,\lambda)P(o_2,...,o_T|i_1=q_i,\lambda)P(i_1=q_i|\lambda) \\= \sum_{i=1}^N b_i(o_1)\beta_1(i)\pi_i$$
所以我们只要求出$\beta_1(i)$，如何求？依次$\beta_T(i) ... \beta_1{t-1}(i) ... \beta_1(i)$

$$\beta_T(i) = P(i_T = q_i,\lambda) = 1 \\ \vdots \\ \beta_t(i) = P(o_{t+1},o_{t+2},...,o_T|i_t=q_i , \lambda) \\= \sum_{j=1}^N P(o_{t+1},o_{t+2},...,o_T,i_{t+1}=q_j|i_t=q_i , \lambda) \\= \sum_{j=1}^N P(o_{t+1},o_{t+2},...,o_T|i_{t+1}=q_j,i_t=q_i , \lambda) P(i_{t+1}=q_j|i_t=q_i , \lambda) \\ 条件前面没有o_t(根据概率图也能得出给定i_{t+1}时，i_t与o_{t+1},...,o_T无关) \\= \sum_{j=1}^N P(o_{t+1},o_{t+2},...,o_T|i_{t+1}=q_j, \lambda) P(i_{t+1}=q_j|i_t=q_i , \lambda) \\= \sum_{j=1}^N P(o_{t+1}|o_{t+2},...,o_T,i_{t+1}=q_j, \lambda)P(o_{t+2},...,o_T|i_{t+1}=q_j, \lambda) P(i_{t+1}=q_j|i_t=q_i , \lambda)\\= \sum_{j=1}^N P(o_{t+1}|i_{t+1}=q_j, \lambda)P(o_{t+2},...,o_T|i_{t+1}=q_j, \lambda) P(i_{t+1}=q_j|i_t=q_i , \lambda) \\ =\sum_{j=1}^N b_j(o_{t+1}) \beta_{t+1}(j) a_{ij}$$

- **模型**：
  $$P(O|\lambda) = \sum_{I}P(O,I|\lambda) = \sum_{I}P(O|I,\lambda)P(I|\lambda)$$

- **策略**：
  $$\argmax_{\lambda} P(O|I,\lambda)$$

- **算法**：
  Baum-Welch 算法，其实就是 EM 算法的一个实现

### 附加知识

#### 随机过程

**随机过程**（[Stochastic process](https://en.jinzhao.wiki/wiki/Stochastic_process)）：

设$(\Omega ,{\mathcal {F}},P)$为一个**概率空间**（[Probability space](https://en.jinzhao.wiki/wiki/Probability_space)）,$\Omega$ 为**样本空间**（[sample space](https://en.jinzhao.wiki/wiki/Sample_space)）， $\mathcal {F}$ 是（[Sigma-algebra](https://en.jinzhao.wiki/wiki/%CE%A3-algebra)），$P$ 是（[Probability measure](https://en.jinzhao.wiki/wiki/Probability_measure)）；
设$(S,\Sigma )$为可测量的空间（measurable space），$S$为随机变量的集合
$${\displaystyle \{X(t):t\in T\}} or {\displaystyle \{X(t,\omega ):t\in T\}}$$

其中$X(t)$是一个随机变量，（在自然科学的许多问题中$t$表示时间，那么$X(t)$表示在时刻$t$观察到的值）；$\omega \in \Omega$；$T$是指标集 or 参数集（index set or parameter set），一般表示时间或空间，如：离散$T=\{0,1,2,...\}$一般称为随机序列或时间序列，连续$T=[a,b] ,a可以取0或者 -\infty,b可以取+\infty$

映射$X(t,\omega):T \times \Omega \to R$，即$X(.,.)$是定义在$T \times \Omega$上的二元值函数;
$\forall t \in T$（固定$t \in T$）,$ X(t,.)$是定义在样本空间$\Omega $上的函数，称为**随机变量**; 
$\forall \omega \in \Omega$,映射$X(.,\omega):T \to S$（其实就是固定$\omega \in \Omega $，变成关于T的函数）,被称为**样本函数**（sample function）,特别是当$T$表示时间时，称为随机过程${\displaystyle \{X(t,\omega ):t\in T\}}$的**样本路径**（sample path）。

#### 傻傻分不清楚的马尔可夫

**马尔可夫性质**（[Markov property](https://en.jinzhao.wiki/wiki/Markov_property)）：
如果**随机过程**（[Stochastic process](https://en.jinzhao.wiki/wiki/Stochastic_process)）的未来状态的条件概率分布（以过去和现在值为条件）仅取决于当前状态，则随机过程具有马尔可夫性质；与此属性的过程被认为是**马氏**或**马尔可夫过程**（[Markov process](https://en.jinzhao.wiki/wiki/Markov_chain)）。最著名的马尔可夫过程是**马尔可夫链**（[Markov chain](https://en.jinzhao.wiki/wiki/Markov_chain)）。**布朗运动**（[Brownian motion](https://en.jinzhao.wiki/wiki/Brownian_motion)）是另一个著名的马尔可夫过程。马尔可夫过程是不具备记忆特质的（[Memorylessness](https://en.jinzhao.wiki/wiki/Memorylessness)）

**一阶 离散**
$${\displaystyle P(X_{n}=x_{n}\mid X_{n-1}=x_{n-1},\dots ,X_{0}=x_{0})=P(X_{n}=x_{n}\mid X_{n-1}=x_{n-1}).}$$

**m 阶 离散**
$${\displaystyle {\begin{aligned}&\Pr(X_{n}=x_{n}\mid X_{n-1}=x_{n-1},X_{n-2}=x_{n-2},\dots ,X_{1}=x_{1})\\={}&\Pr(X_{n}=x_{n}\mid X_{n-1}=x_{n-1},X_{n-2}=x_{n-2},\dots ,X_{n-m}=x_{n-m}){\text{ for }}n>m\end{aligned}}}$$

**时间齐次**（Time-homogeneous）
$${\displaystyle \Pr(X_{n+1}=x\mid X_{n}=y)=\Pr(X_{n}=x\mid X_{n-1}=y)}$$

**马尔可夫模型**（[Markov model](https://en.jinzhao.wiki/wiki/Markov_model)）：
.|System state is fully observable |System state is partially observable
---|---|---
System is autonomous |[Markov chain](https://en.jinzhao.wiki/wiki/Markov_chain) | [Hidden Markov model](https://en.jinzhao.wiki/wiki/Hidden_Markov_model)
System is controlled |[Markov decision process](https://en.jinzhao.wiki/wiki/Markov_decision_process) |[Partially observable Markov decision process](https://en.jinzhao.wiki/wiki/Partially_observable_Markov_decision_process)

马尔可夫模型是具有马尔可夫性假设的随机过程，最简单的马尔可夫模型是**马尔可夫链**（[Markov chain](https://en.jinzhao.wiki/wiki/Markov_chain)）

| .                                        | Countable state space（对应随机过程的离散$\Omega$）               | Continuous or general state space（对应随机过程的连续$\Omega$）                              |
| ---------------------------------------- | ----------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| Discrete-time（对应随机过程的离散$T$）   | (discrete-time) Markov chain on a countable or finite state space | Markov chain on a measurable state space (for example, Harris chain)                         |
| Continuous-time（对应随机过程的连续$T$） | Continuous-time Markov process or Markov jump process             | Any continuous stochastic process with the Markov property (for example, the Wiener process) |

令$\{X_n|n=1,2,\cdots\}$是有限个或可数个可能值的随机过程。除非特别提醒，这个随机过程的可能值的集合都将记为非负整数集$\{0,1,2,\cdots\}$ 。
如果 $X_n=i$，那么称该过程在时刻 $t$ 在状态 $i$ ，并假设$P_{i,j}$ 称为(单步(one-step))**转移概率**(transition probability)，表示处在$i$状态的随机变量下一时刻处在$j$ 状态的概率，如果对所有的状态 $i_0,i_1,\cdots,i_{n-1},i,j$及任意$n\ge 0$ ，$P\{X_{n+1}=j|X_n=i,X_{n-1}=i_{n-1},\cdots,X_1=i_1,X_0=i_0\}=P_{i,j}$，这样的过程称为**马尔可夫链**(Markov chain)。

> 一个随机过程$\{X(t):t \geq 0\}$如果$t \in \mathbb{R}_+$则称为连续时间的马尔科夫链，如果$t \in \mathbb{N}_0$则称为离散时间的马尔科夫链

根据 $P_{i,j}$的定义显然有$P_{i,j}\ge0,\;i,j\ge0;\;\;\sum_{j=0}^\infty P_{i,j}=1,\;i=0,1,\cdots$,
用 $P_{i,j}$ 记录从 $i$ 到 $j$ 的(单步)**转移（概率）矩阵**（[transition matrix](https://en.jinzhao.wiki/wiki/Stochastic_matrix)）也称为**随机矩阵、概率矩阵、转移矩阵、替代矩阵或马尔可夫矩阵**：
$$\mathbf{P}_{i,j}=(P_{i_{n},i_{n+1}}) =\begin{bmatrix}P_{0,0}&P_{0,1}&P_{0,2}&\cdots\\P_{1,0}&P_{1,1}&P_{1,2}&\cdots\\\vdots&\vdots&\vdots\\P_{i,0}&P_{i,1}&P_{i,2}&\cdots\\\vdots&\vdots&\vdots\end{bmatrix}$$
现在定义 n 步(n-step)**转移概率**$P_{i,j}^n$ ：$P_{i,j}^n=P\{X_{n+k=j}|X_k=i\},\;n\ge 0,i,j\ge 0$

**右随机矩阵**是一个非负实数的方阵，每个行总和为 1。
**左随机矩阵**是一个非负实数的方阵，每个列的总和为 1。
**双随机矩阵**是一个非负实数的方阵，每个行和每个列的总和为 1。

假设$A$是马尔可夫矩阵，其性质有：

1. 矩阵$A$的 k 次幂$A^k$也是马尔可夫矩阵。
1. 至少有一个特征值为 1，其特征值在[-1,1]区间，特征值为 1 对应的特征向量$\pi$称为**平稳概率向量**（stationary probability vector）。
1. 对于任意**概率向量**（[Probability vector](https://en.jinzhao.wiki/wiki/Probability_vector)）或者**随机向量**$\pi_0$有$\lim_{k \to \infty} A^k \pi_0 = \pi$ （这里是在没有-1 特征值的情况下）。
1. 对于任意**概率向量**$\mu_0$有$\mu_1 = A \mu_0$也是概率向量。

特征值的求解：$\det(A-\lambda I)=0$

由于 $A$ 的每一列相加等于 1，所以 $A−I$ 的每一列相加等于 0，这也就是说 $A−I$ 的行是相关的，其行列式$\det(A-I)=0$为零，所以 $A−I$奇异矩阵，所以 1 是 $A$ 的一个特征值。

[对角化](https://en.jinzhao.wiki/wiki/Diagonalizable_matrix#Diagonalization) $A = P \Lambda P^{-1} $ （参见线性代数及其应用279页，特征值相同特征向量不一定相同）,其中$\Lambda$是由$A$的特征值组成的对角矩阵
$$\mu_k = A^k \mu_0 = (P \Lambda P^{-1})^k \mu_0 = P \Lambda^k P^{-1} \mu_0$$
不妨设 $A$的特征向量和相应的特征值分别为 ${x_1},...,{x_n}$和 $\lambda_1,...,\lambda_n$，可以用特征向量来做一组基，可以把空间中任何向量写成它的线性组合：$\mu_0 = c_1{x_1} + \cdots + c_n{x_n}$
那么：
$$A^k\mathbf{\mu_0} = A^kc_1\mathbf{x_1} + \cdots + A^kc_n\mathbf{x_n}\\ = c_1A^k\mathbf{x_1} + \cdots + c_nA^k\mathbf{x_n} \\= c_1A^{k-1}A\mathbf{x_1} + \cdots + c_nA^{k-1}A\mathbf{x_n} \\= c_1A^{k-1}\lambda_1\mathbf{x_1} + \cdots + c_nA^{k-1}\lambda_n\mathbf{x_n}\\=c_1\lambda_1^k\mathbf{x_1} + \cdots + c_n\lambda_n^k\mathbf{x_n}\\=\sum_{i=1}^n{c_i\lambda_i^k\bm{x_i}}$$

不妨令$\lambda_1=1$, 有$|\lambda_i|\leq 1$,那么：
$$\bm{u_\infty}=\lim_{k\to\infty}{A^k\bm{u_0}}=\lim_{k\to\infty}{\sum_{i=1}^k{c_i\lambda_i^k\bm{x_i}}}=c_1\bm{x_1}$$

因为$u_\infty$是概率向量，而特征值为 1 对应的特征向量也是概率向量，所以$c_1=1$，得到$\bm{u_\infty}=\bm{x_1}$

值得注意的是，除包含 $\lambda=1$ 的情形还包含 $\lambda=-1$ 的情形。
上式如果除$\lambda_1=1$还有$\lambda_2=-1$，那么就有：
$$\bm{u_\infty}=\lim_{k\to\infty}{\sum_{i=1}^k{c_i\lambda_i^k\bm{x_i}}}=c_1\bm{x_1}+(-1)^k c_2\bm{x_2}$$

得$\bm{u_\infty}=\bm{x_1}+(-1)^k\bm{x_2}$,此时$k$为奇数和偶数结果是不同的，造成的结果就是在两种结果之间反复横跳，永远达不到稳态。

如：
$$A=\begin{bmatrix}0&1\\1&0\\\end{bmatrix}$$
其特征值为$\lambda_1=1，\lambda_2=-1$

> 也可以参考第 21 章 PageRank 算法

### 参考文献

[10-1] Rabiner L,Juang B. An introduction to hidden markov Models. IEEE ASSPMagazine,January 1986

[10-2] Rabiner L. A tutorial on hidden Markov models and selected applications in speechrecognition. Proceedings of IEEE,1989

[10-3] Baum L,et al. A maximization technique occuring in the statistical analysis ofprobabilistic functions of Markov chains. Annals of Mathematical Statistics,1970,41: 164–171

[10-4] Bilmes JA. [A gentle tutorial of the EM algorithm and its application to parameterestimation for Gaussian mixture and hidden Markov models](https://people.ece.uw.edu/bilmes/p/mypubs/bilmes1997-em.pdf).

[10-5] Lari K,Young SJ. Applications of stochastic context-free grammars using the Inside-Outside algorithm,Computer Speech & Language,1991,5(3): 237–257

[10-6] Ghahramani Z. Learning Dynamic Bayesian Networks. Lecture Notes in ComputerScience,Vol. 1387,1997,168–197

## 第 11 章 条件随机场

- **模型**：
- **策略**：
- **算法**：

### 附加知识

### 参考文献

[11-1] Bishop M. Pattern Recognition and Machine Learning. Springer-Verlag,2006

[11-2] Koller D,Friedman N. Probabilistic Graphical Models: Principles and Techniques.MIT Press,2009

[11-3] Lafferty J,McCallum A,Pereira F. Conditional random fields: probabilistic models for segmenting and labeling sequence data. In: International Conference on Machine Learning,2001

[11-4] Sha F,Pereira F. Shallow parsing with conditional random fields. In: Proceedings ofthe 2003 Conference of the North American Chapter of Association for ComputationalLinguistics on Human Language Technology,Vol.1,2003

[11-5] McCallum A,Freitag D,Pereira F. Maximum entropy Markov models for informationextraction and segmentation. In: Proc of the International Conference on Machine Learning,2000

[11-6] Taskar B,Guestrin C,Koller D. Max-margin Markov networks. In: Proc of the NIPS2003,2003

[11-7] Tsochantaridis I,Hofmann T,Joachims T. Support vector machine learning forinterdependent and structured output spaces. In: ICML,2004

## 第 12 章 监督学习方法总结
