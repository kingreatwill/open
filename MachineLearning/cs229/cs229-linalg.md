> CS229 斯坦福大学机器学习复习材料(数学基础) - [线性代数](http://cs229.stanford.edu/summer2020/cs229-linalg.pdf)

[TOC]

# 线性代数回顾与参考

## 1 基本概念和符号

线性代数提供了一种紧凑地表示和运算“线性方程组”的方法。例如，考虑以下方程组：

$$
4x_{\scriptscriptstyle 1} - 5x_{\scriptscriptstyle 2} = -13\\
-2x_{\scriptscriptstyle 1} + 3x_{\scriptscriptstyle 2} = 9.
$$

这是两个方程和两个变量，正如你从高中代数中所知，你可以找到 $x_1$ 和 $x_2$ 的唯一解（除非方程以某种方式退化，例如，如果第二个方程只是第一个的倍数，但在上面的情况下，实际上只有一个唯一解）。 在矩阵表示法中，我们可以更紧凑地表达：

$$
Ax= b
$$

其中

$$
A=\begin{bmatrix}
   4 & -5 \\
   -2 & 3
\end{bmatrix},
b=\begin{bmatrix}
   -13 \\
   9
\end{bmatrix}.
$$

我们可以看到，以这种形式分析线性方程有许多优点（包括明显的节省空间）。

### 1.1 基本符号

我们使用以下符号：

- 用$A \in \Bbb{R}^{m \times n} $表示一个$m$行$n$列的矩阵，其中$A$的各项都是实数。
- 用$\boldsymbol{x} \in \Bbb{R}^{n}$表示具有$n$个元素的向量。按照惯例，$n$维向量。通常被认为是$n$行$1$列的矩阵，称为<span id="column_vector"></span>**列向量**。如果我们想表示一个<span id="row_vector"></span>**行向量**: 具有 $1$ 行和$n$列的矩阵 - 我们通常写$\boldsymbol{x}^T$（这里$\boldsymbol{x}^T$表示$\boldsymbol{x}$的转置，我们稍后将定义它）。
- 用$x_i$表示向量$\boldsymbol{x}$的第$i$个元素：
  $$
  \boldsymbol{x}=\begin{bmatrix}
     x_1 \\
     x_2 \\
     \vdots \\
     x_n
  \end{bmatrix}.
  $$
- 我们用符号$a_{\scriptscriptstyle ij} $(or $A_{ij}$, $A_{i,j}$)表示$A$的第$i$行第$j$列元素：

  $$
  A=\begin{bmatrix}
     a_{11} & a_{12} & \cdots & a_{1n} \\
     a_{21} & a_{22} & \cdots & a_{2n} \\
     \vdots & \vdots & \ddots & \vdots \\
     a_{m1} & a_{m2} & \cdots & a_{mn}
  \end{bmatrix}.
  $$

- 我们将$A$的第$j$列表示为$a^j$ or $A_{:,j}$ ：

  $$
  A = \begin{bmatrix}
  \text{\textbar} & \text{\textbar} &  & \text{\textbar} \\
  a^1 & a^2 & \cdots & a^n \\
  \text{\textbar} & \text{\textbar} &  & \text{\textbar}
  \end{bmatrix}.
  $$

- 我们将$A$的第$i$行表示为$a_i^T$ or $A_{i,:}$ ：

  $$
  A = \begin{bmatrix}
  \text{\textemdash} & a_1^T & \text{\textemdash} \\
  \text{\textemdash} & a_2^T & \text{\textemdash} \\
   & \vdots &  \\
  \text{\textemdash} & a_m^T & \text{\textemdash} \\
  \end{bmatrix}.
  $$

- 在许多情况下，将矩阵视为列向量或行向量的集合是非常重要和方便的。一般来说，在数学上(和概念上)向量级别上的操作比标量级别上的操作会更简洁。表示矩阵的列或行没有通用约定，因此你可以使用任何符号明确定义它。

## 2 矩阵乘法

矩阵$A \in \Bbb{R}^{m \times n}$和矩阵$B \in \Bbb{R}^{n \times p}$的乘积仍然是一个矩阵$C = AB \in \Bbb{R}^{m \times p}$，其中$C_{ij} = \displaystyle\sum_{k=1}^n {A_{ik}B_{kj}} $.
请注意，为了使矩阵乘积存在，$A$中的列数必须等于$B$中的行数。有很多方法可以查看矩阵乘法，我们将从检查一些特殊情况开始。

### 2.1 向量-向量乘法

给两个向量$\boldsymbol{x},\boldsymbol{y} \in \Bbb{R}^n$, $\boldsymbol{x}^T \boldsymbol{y}$通常称为向量的<span id="inner_product"></span>**内积**或者<span id="dot_product"></span>**点积**，结果是个实数：

$$
\boldsymbol{x}^T \boldsymbol{y} \in \Bbb{R} =
\begin{bmatrix}x_1 & x_2 & \cdots & x_n\end{bmatrix}
\begin{bmatrix}y_1 \\ y_2 \\ \vdots \\ y_n \end{bmatrix}
= \sum_{i=1}^n{x_iy_i}
$$

请注意，内积实际上只是矩阵乘法的特例。$\boldsymbol{x}^T \boldsymbol{y} = \boldsymbol{y}^T \boldsymbol{x}$ 始终成立。
给定向量$\boldsymbol{x} \in \Bbb{R}^m , \boldsymbol{y} \in \Bbb{R}^n$($m$不一定等于$n$), $\boldsymbol{x} \boldsymbol{y}^T \in \Bbb{R}^{m \times n}$叫向量<span id="outer_product"></span>**外积**，它是一个矩阵，由$(\boldsymbol{x} \boldsymbol{y}^T)_{ij} = x_iy_j$组成，也就是(i.e.)：

$$
\boldsymbol{x} \boldsymbol{y}^T \in \Bbb{R}^{m \times n}=
\begin{bmatrix}x_1 \\ x_2 \\ \vdots \\ x_m \end{bmatrix}
\begin{bmatrix}y_1 & y_2 & \cdots & y_n\end{bmatrix}=
\begin{bmatrix}
     x_1y_1 & x_1y_2 & \cdots & x_1y_n \\
     x_2y_1 & x_2y_2 & \cdots & x_2y_n \\
     \vdots & \vdots & \ddots & \vdots \\
     x_my_1 & x_my_2 & \cdots & x_my_n
\end{bmatrix}
$$

举一个外积如何使用的一个例子：让$\boldsymbol{1}\in \Bbb{R}^{n}$表示一个$n$维向量，其元素都等于 1，此外，考虑矩阵$A \in \Bbb{R}^{m \times n}$，其列全部是某个向量 $\boldsymbol{x} \in R^{m}$。 我们可以使用外积紧凑地表示矩阵 $A$:

$$
A=\begin{bmatrix}
  \text{\textbar} & \text{\textbar} &  & \text{\textbar} \\
  x & x & \cdots & x \\
  \text{\textbar} & \text{\textbar} &  & \text{\textbar}
\end{bmatrix}=
\begin{bmatrix}
     x_{1} & x_{1} & \cdots & x_{1} \\
     x_{2} & x_{2} & \cdots & x_{2} \\
     \vdots & \vdots & \ddots & \vdots \\
     x_{m} & x_{m} & \cdots & x_{m}
\end{bmatrix}=
\begin{bmatrix}x_1 \\ x_2 \\ \vdots \\ x_m \end{bmatrix}
\begin{bmatrix}1 & 1 & \cdots & 1\end{bmatrix}=\boldsymbol{x}\boldsymbol{1}^T
$$

### 2.2 矩阵-向量乘法

给定矩阵 $A \in \mathbb{R}^{m \times n}$，向量 $\boldsymbol{x} \in  \mathbb{R}^{n}$ , 它们的积是一个向量 $\boldsymbol{y} = A\boldsymbol{x} \in \mathbb{R}^{m}$。 有几种方法可以查看矩阵向量乘法。

如果我们按行写$A$，那么我们可以表示$A\boldsymbol{x}$为：

$$
\boldsymbol{y} = A\boldsymbol{x} =
\begin{bmatrix}
  \text{\textemdash} & a_1^T & \text{\textemdash} \\
  \text{\textemdash} & a_2^T & \text{\textemdash} \\
   & \vdots &  \\
  \text{\textemdash} & a_m^T & \text{\textemdash} \\
\end{bmatrix}\boldsymbol{x}=
\begin{bmatrix} a_1^T\boldsymbol{x} \\ a_2^T\boldsymbol{x} \\ \vdots \\ a_m^T\boldsymbol{x} \end{bmatrix}
$$

换句话说，第$i$个$y$的元素是$A$的第$i$行和$\boldsymbol{x}$的内积，即：$y_i=a_{i}^{T} \boldsymbol{x}$。

同样的， 可以把 $A$ 写成列的方式，如下：

$$
\boldsymbol{y} = A\boldsymbol{x} =
\begin{bmatrix}
  \text{\textbar} & \text{\textbar} &  & \text{\textbar} \\
  a^1 & a^2 & \cdots & a^n \\
  \text{\textbar} & \text{\textbar} &  & \text{\textbar}
\end{bmatrix}
\begin{bmatrix}x_1 \\ x_2 \\ \vdots \\ x_n \end{bmatrix}=
[a^1]x_1 + [a^2]x_2 + \cdots +[a^n]x_n
\label{1}\tag{1}
$$

换句话说，$\boldsymbol{y}$是$A$的列的<span id="linear_combination"></span>**线性组合**，其中线性组合的系数由$\boldsymbol{x}$的元素给出。

到目前为止，我们一直是矩阵右乘一个列向量，但也可以是矩阵左乘一个行向量。如这样表示：$\boldsymbol{y}^T = \boldsymbol{x}^TA$ 其中$A \in \mathbb{R}^{m \times n}$，$\boldsymbol{x} \in \mathbb{R}^{m}$，$\boldsymbol{y} \in \mathbb{R}^{n}$。 和以前一样，我们可以用两种可行的方式表达$\boldsymbol{y}^T$，这取决于我们是否根据行或列表达$A$.

首先，我们把$A$用列表示：

$$
\boldsymbol{y}^T = \boldsymbol{x}^TA = \boldsymbol{x}^T
\begin{bmatrix}
  \text{\textbar} & \text{\textbar} &  & \text{\textbar} \\
  a^1 & a^2 & \cdots & a^n \\
  \text{\textbar} & \text{\textbar} &  & \text{\textbar}
\end{bmatrix}=
\begin{bmatrix}\boldsymbol{x}^Ta^1 & \boldsymbol{x}^Ta^2 & \cdots & \boldsymbol{x}^Ta^n \end{bmatrix}
$$

这表明$\boldsymbol{y}^T$的第$i$个元素等于$\boldsymbol{x}$和$A$的第$i$列的内积。

最后，根据行表示$A$，我们得到了向量-矩阵乘积的最终表示：

$$
\begin{aligned}
\boldsymbol{y}^T &= \boldsymbol{x}^TA \\&=
\begin{bmatrix} x_1 & x_2 & \cdots & x_n \end{bmatrix}
\begin{bmatrix}
  \text{\textemdash} & a_1^T & \text{\textemdash} \\
  \text{\textemdash} & a_2^T & \text{\textemdash} \\
   & \vdots &  \\
  \text{\textemdash} & a_m^T & \text{\textemdash} \\
\end{bmatrix}\\&=
x_1\begin{bmatrix}\text{\textemdash} & a_1^T & \text{\textemdash}\end{bmatrix}+
x_2\begin{bmatrix}\text{\textemdash} & a_2^T & \text{\textemdash}\end{bmatrix}+ \cdots +
x_n\begin{bmatrix}\text{\textemdash} & a_n^T & \text{\textemdash}\end{bmatrix}
\end{aligned}
$$

所以我们看到$\boldsymbol{y}^T$是$A$的行的线性组合，其中线性组合的系数由$\boldsymbol{x}$的元素给出。

### 2.3 矩阵-矩阵乘法

有了这些知识，我们现在可以看看四种不同的（当然是等价的）查看矩阵与矩阵乘法$ C = AB $的方法。

首先，我们可以将矩阵-矩阵乘法视为一组向量-向量乘积。 从定义中可以得出：最明显的观点是$C$的($i,j$)元素等于$A$的第$i$行和$B$的的$j$列的内积。如下所示：

$$
C = AB =
\begin{bmatrix}
  \text{\textemdash} & a_1^T & \text{\textemdash} \\
  \text{\textemdash} & a_2^T & \text{\textemdash} \\
   & \vdots &  \\
  \text{\textemdash} & a_m^T & \text{\textemdash} \\
\end{bmatrix}
\begin{bmatrix}
  \text{\textbar} & \text{\textbar} &  & \text{\textbar} \\
  b^1 & b^2 & \cdots & b^p \\
  \text{\textbar} & \text{\textbar} &  & \text{\textbar}
\end{bmatrix}=
\begin{bmatrix}
     a_1b_1 & a_1b_2 & \cdots & a_1b_p \\
     a_2b_1 & a_2b_2 & \cdots & a_2b_p \\
     \vdots & \vdots & \ddots & \vdots \\
     a_mb_1 & a_mb_2 & \cdots & a_mb_p
\end{bmatrix}
$$

这里的矩阵$A \in \Bbb{R}^{m \times n} , B \in \Bbb{R}^{n \times p}$ ， 向量$a_i \in \Bbb{R}^n , b^j \in \Bbb{R}^n $ ，所以它们可以计算内积。当我们用行表示 $A$ 和用列表示 $B$ 时，这是最“自然”的表示。
另外，我们可以用列表示 $A$，用行表示 $B$。这种表示导致将 $AB$ 解释为外积之和,这种表示则复杂得多。象征性地，

$$
C = AB =
\begin{bmatrix}
  \text{\textbar} & \text{\textbar} &  & \text{\textbar} \\
  a^1 & a^2 & \cdots & a^n \\
  \text{\textbar} & \text{\textbar} &  & \text{\textbar}
\end{bmatrix}
\begin{bmatrix}
  \text{\textemdash} & b_1^T & \text{\textemdash} \\
  \text{\textemdash} & b_2^T & \text{\textemdash} \\
   & \vdots &  \\
  \text{\textemdash} & b_n^T & \text{\textemdash} \\
\end{bmatrix}=
\sum_{i=1}^n{a^ib_i^T}
$$

换句话说，$AB$等于所有的$A$的第$i$列和$B$第$i$行的外积的和。因此，在这种情况下， $a^i \in \mathbb{R}^ m $和$b_i \in \mathbb{R}^p$， 外积$a^ib_i^T$的维度是$m×p$，与$C$的维度一致。

其次，我们还可以将矩阵-矩阵乘法视为一组矩阵-向量乘法。如果我们把$B$用列表示，我们可以将$C$的列视为$A$和$B$的列(矩阵-向量)的乘积。如下所示：

$$
C = AB = A
\begin{bmatrix}
  \text{\textbar} & \text{\textbar} &  & \text{\textbar} \\
  b^1 & b^2 & \cdots & b^p \\
  \text{\textbar} & \text{\textbar} &  & \text{\textbar}
\end{bmatrix}=
\begin{bmatrix}
  \text{\textbar} & \text{\textbar} &  & \text{\textbar} \\
  Ab^1 & Ab^2 & \cdots & Ab^p \\
  \text{\textbar} & \text{\textbar} &  & \text{\textbar}
\end{bmatrix}
\label{2}\tag{2}
$$

这里$C$的第$i$列由矩阵-向量乘积给出，右边的向量为$c_i = Ab_i$。

最后，我们有类似的观点，我们用行表示$A$，并将$C$的行视为$A$的行和$B$之间的矩阵-向量乘积。如下所示：

$$
C = AB =
\begin{bmatrix}
  \text{\textemdash} & a_1^T & \text{\textemdash} \\
  \text{\textemdash} & a_2^T & \text{\textemdash} \\
   & \vdots &  \\
  \text{\textemdash} & a_m^T & \text{\textemdash} \\
\end{bmatrix}B=
\begin{bmatrix}
  \text{\textemdash} & a_1^TB & \text{\textemdash} \\
  \text{\textemdash} & a_2^TB & \text{\textemdash} \\
   & \vdots &  \\
  \text{\textemdash} & a_m^TB & \text{\textemdash} \\
\end{bmatrix}
$$

这里$C$的第$i$行由矩阵-向量乘积给出：$c_i^T = a_i^T B$。

将矩阵乘法剖析到如此大的程度似乎有点矫枉过正，特别是当所有这些观点都紧跟在我们在本节开头给出的初始定义（$C=AB$）之后。

这些不同方法的直接优势在于它们允许您**在向量的级别/单位而不是标量上进行操作**。 为了完全理解线性代数而不会迷失在复杂的索引操作中，关键是操作尽可能大(向量而不是标量)的概念。[^1]

[^1]: E.g., 如果你可以用矩阵或向量来写出你所有的数学推导，那会比用标量元素来写要好。

实际上所有的线性代数都是在处理某种矩阵乘法，多花一些时间对这里提出的观点进行直观的理解是非常必要的。

除此之外，你还应该了解一些更高级别的矩阵乘法的基本性质：

- 矩阵乘法结合律: $(AB)C = A(BC)$
- 矩阵乘法分配律: $A(B + C) = AB + AC$
- 矩阵乘法一般是不可交换的; 也就是说，通常$AB \ne BA$。 （例如，假设$ A \in \mathbb{R}^ {m \times n}，$ $B \in \mathbb{R}^ {n \times p} $，如果$m$和$q$不相等，矩阵乘积$BA$甚至不存在！）

如果您不熟悉这些性质，请花点时间自己验证它们。 例如，为了检查矩阵乘法的结合性，假设$A \in \mathbb{R}^ {m \times n}，$ $B \in \mathbb{R}^ {n \times p} $，$C \in \mathbb{R}^ {p \times q}$。 注意$AB \in \mathbb{R}^ {m \times p}$，所以$(AB)C \in \mathbb{R}^ {m \times q}$。 类似地，$BC \in \mathbb{R}^ {n \times q}$，所以$A(BC) \in \mathbb{R}^ {m \times q}$。 因此，所得矩阵的维度一致。 为了验证矩阵乘法的结合性，检查$(AB)C$的($i,j$)元素是否等于$A(BC)$的($i,j$)元素。 我们可以使用矩阵乘法的定义直接验证这一点：

$$
\begin{aligned} % aligned &= 等号对齐
((A B) C)_{ij} &= \sum_{k=1}^p{(AB)_{ik}C_{kj}} = \sum_{k=1}^p \Bigg( \sum_{l=1}^n{A_{il}B_{lk}} \Bigg) C_{kj} \\
&=\sum_{k=1}^p \Bigg( \sum_{l=1}^n{A_{il}B_{lk}C_{kj}}\Bigg) = \sum_{l=1}^n \Bigg( \sum_{k=1}^p{A_{il}B_{lk}C_{kj}}\Bigg)\\
&=\sum_{l=1}^nA_{il}\Bigg(\sum_{k=1}^p{B_{lk}C_{kj}}\Bigg)  = \sum_{l=1}^n{A_{il}(BC)_{lj}} = (A(BC))_{ij}
\end{aligned}
$$

这里，第一个和最后两个等式简单地使用了矩阵乘法的定义，第三个和第五个等式使用了标量乘法对加法的分配性质，第四个等式使用了标量加法的交换性和结合性。 这种通过简化为简单标量性质来证明矩阵性质的技术会经常出现，因此请确保您熟悉它。

## 3 操作及其性质

在本节中，我们将介绍矩阵和向量的几种操作和性质。 希望其中的大部分内容都可以帮您复习，此笔记可以作为参考。

### 3.1 单位矩阵和对角矩阵

**单位矩阵**<span id="identity_matrix"></span>用$I \in \Bbb{R}^{n \times n}$表示，它是一个方阵，对角线的元素是 1，其余元素都是 0。可以这样表示：

$$
I_{ij} =
\begin{cases}
   1 & i=j \\
   0 & i \neq j
\end{cases}.
$$

对于所有矩阵$A \in \mathbb{R}^ {m \times n}$，有：
$$AI=A=IA$$
注意，在某种意义上，上面单位矩阵的表示法是不明确的，因为它没有指定$I$的维数。通常，$I$的维数是从上下文推断出来的，以便使矩阵乘法成为可能。 例如，在上面的等式中，$AI = A$中的$I$是$n\times n$矩阵，而$A = IA$中的$I$是$m\times m$矩阵。

**对角矩阵**<span id="diagonal_matrix"></span>的非对角元素全为 0。对角阵通常表示为：$D=diag(d_1, d_2,\cdots, d_n)$，其中：

$$
D_{ij} =
\begin{cases}
   d_i & i=j \\
   0 & i \neq j
\end{cases}.
$$

很明显，单位矩阵$ I= diag(1, 1, \cdots , 1)$。

### 3.2 转置

矩阵的**转置**<span id="transpose"></span>是指翻转矩阵的行和列。给定一个矩阵$A \in \Bbb{R}^{m \times n}$，它的转置$A^T \in \Bbb{R}^{n \times m}$ ,其中的元素为：

$$
(A^T)_{ij} = A_{ji}.
$$

事实上，我们在描述行向量时已经使用了转置，因为列向量的转置自然是行向量。

转置有以下性质，且很容易验证：

- $(A^T)^T = A$
- $(AB)^T = B^TA^T$
- $(A+B)^T = A^T + B^T$

### 3.3 对称矩阵

如果$A = A^T$，那么**方阵**$A \in \Bbb{R}^{n \times n}$是**对称**<span id="symmetric"></span>的。

- 元素满足$a_{ij} = a_{ji} , \forall i,j$
- $A = A^T$
- 对于任意方阵$A$，$A + A^T$是对称的
- 对角矩阵都是对称矩阵

如果$A = -A^T$，那么它就是**反对称**<span id="anti-symmetric"></span>的。

- 元素满足$a_{ij} = -a_{ji} , \forall i,j$,所以当$i=j$时，$a_{ij} = 0$
- $A,B$都为反对称矩阵，则$A \plusmn B$仍为反对称矩阵[^2]
- 若$A$为反对称矩阵，$B$为对称矩阵，则$AB-BA$为对称矩阵[^3]

很容易证明，对于任何矩阵$A \in \mathbb{R}^ {n \times n}$，矩阵$A  +  A^ T$是对称的，矩阵$A -A^T$是反对称的[^2]。

[^2]: 设$A,B$为反对称矩阵，即有$A^T = -A , B^T=-B$则：$(A \plusmn B)^T = A^T \plusmn B^T = (-A) \plusmn (-B) = -(A \plusmn B)$
[^3]: 设$A$为反对称矩阵，$B$为对称矩阵，即有$A^T = -A , B^T=B$则：

    $$
    (AB - BA)^T = (AB)^T - (BA)^T = B^TA^T - A^TB^T = -BA + AB =(AB - BA)
    $$

由此得出，任意方矩阵$A \in \mathbb{R}^ {n \times n}$可以表示为对称矩阵和反对称矩阵的和，所以：

$$
A=\frac{1}{2}(A+A^T)+\frac{1}{2}(A-A^T)
$$

事实证明，对称矩阵在实践中用到很多，它们有很多很好的性质，我们很快就会看到它们。
通常将大小为$n$的所有对称矩阵的集合表示为$\mathbb{S}^n$，因此$A \in \mathbb{S}^n$意味着$A$是对称的$n\times n$矩阵。

### 3.4 矩阵的迹

方矩阵$A \in \mathbb{R}^ {n \times n}$的**迹**<span id="trace"></span>，表示为$\operatorname{tr} (A)$（或者$\operatorname{tr} A$，括号显然是隐含的），是矩阵中对角元素的总和：

$$
\operatorname{tr} A=\sum_{i=1}^{n} A_{i i}
$$

如 CS229 讲义中所述，迹具有以下性质（如下所示）：

- 对于矩阵$A \in \mathbb{R}^ {n \times n}$，则：$\operatorname{tr}A =\operatorname{tr}A^T$
- 对于矩阵$A,B \in \mathbb{R}^ {n \times n}$，则：$\operatorname{tr}(A + B) = \operatorname{tr}A + \operatorname{tr}B$
- 对于矩阵$A \in \mathbb{R}^ {n \times n}$，$ t \in \mathbb{R}$，则：$\operatorname{tr}(tA) = t\operatorname{tr}A$.
- 对于矩阵 $A$, $B$，$AB$ 为方阵, 则：$\operatorname{tr}AB = \operatorname{tr}BA$
- 对于矩阵 $A$, $B$, $C$, $ABC$为方阵(包含 1\*1 的矩阵-标量), 则：$\operatorname{tr}ABC = \operatorname{tr}BCA=\operatorname{tr}CAB$, 同理，更多矩阵的积也是有这个性质。

我们给出第四个性质的证明。假设$A \in \mathbb{R}^ {m \times n}$和$B \in \mathbb{R}^ {n \times m}$（因此$AB \in \mathbb{R}^ {m \times m}$是方阵）。 观察到$BA \in \mathbb{R}^ {n \times n}$也是一个方阵，因此对它们进行迹的运算是有意义的。 要证明$\operatorname{tr}AB = \operatorname{tr}BA$，注意：

$$
\begin{aligned}
\operatorname{tr} A B &=\sum_{i=1}^{m}(A B)_{i i}=\sum_{i=1}^{m}\left(\sum_{j=1}^{n} A_{i j} B_{j i}\right) \\
&=\sum_{i=1}^{m} \sum_{j=1}^{n} A_{i j} B_{j i}=\sum_{j=1}^{n} \sum_{i=1}^{m} B_{j i} A_{i j} \\
&=\sum_{j=1}^{n}\left(\sum_{i=1}^{m} B_{j i} A_{i j}\right)=\sum_{j=1}^{n}(B A)_{j j}=\operatorname{tr} B A
\end{aligned}
$$

这里，第一个和最后两个等式使用了迹运算符和矩阵乘法的定义。 重点在第四个等式,使用标量乘法的交换性来反转每个乘积中的项的顺序，以及标量加法的交换性和结合性来重新排列求和的顺序。

### 3.5 范数

向量的**范数**<span id="norm"></span>$\|x\|$是非正式度量的向量的“长度” 。 例如，我们有常用的欧几里德或$\ell_{2}$范数，

$$
\|x\|_{2}=\sqrt{\sum_{i=1}^{n} x_{i}^{2}}
$$

注意：$\|x\|_{2}^{2}=x^{T} x$

更正式地，范数是满足 4 个性质的函数（$f : \mathbb{R}^{n} \rightarrow \mathbb{R}$）：

1. 对于所有的 $x \in \mathbb{R}^ {n}$, $f(x) \geq 0 $(非负性).
2. 当且仅当$x = 0$ 时，$f(x) = 0$ (确定性).
3. 对于所有$x \in \mathbb{R}^ {n}$,$t\in \mathbb{R}$，则 $f(tx) = \left| t \right|f(x)$ (正齐次性).
4. 对于所有 $x,y \in \mathbb{R}^ {n}$, $f(x + y) \leq f(x) + f(y)$ (三角不等式)

其他范数的例子，如：$\ell_1$范数：

$$
\|x\|_{1}=\sum_{i=1}^{n}|x_{i}|
$$

和$\ell_{\infty }$范数：

$$
\|x\|_{\infty}=\max_{i}\left|x_{i}\right|
$$

事实上，到目前为止所提出的所有三个范数都是$\ell_p$范数族的例子，它们由实数$p \geq 1$参数化，并定义为：

$$
\|x\|_{p}=\left(\sum_{i=1}^{n}\left|x_{i}\right|^{p}\right)^{1 / p}
$$

也可以为矩阵定义范数，例如**Frobenius**范数:

$$
\|A\|_{F}=\sqrt{\sum_{i=1}^{m} \sum_{j=1}^{n} A_{i j}^{2}}=\sqrt{\operatorname{tr}\left(A^{T} A\right)}
$$

还有很多其他范数，但它们超出了这个复习材料的范围。

### 3.6 线性相关性和秩

一个向量集合$\{ x_1,x_2, \cdots x_n \} \subset \mathbb{R}^m$， 如果没有向量可以表示为其余向量的线性组合，则称称该向量是**线性无关**<span id="linearly_independent"></span>的。 相反，如果属于该组的一个向量可以表示为其余向量的线性组合，则称该向量是**线性相关**<span id="linearly_dependent"></span>的。 也就是说，如果：

$$
x_{j}=\sum_{i=1,i \neq j}^{n} \alpha_{i} x_{i}
$$

存在$\alpha_1,\cdots \alpha_{n} \in \mathbb{R}$，那么向量$x_1,x_2, \cdots x_n$是线性相关的; 否则，向量是线性无关的。
另一种**线性相关**的描述（存在不全为零的数$\alpha_{i}$，使得等式成立）：

$$
\sum_{i=1}^{n} \alpha_{i} x_{i} = 0,\exists \alpha_i \neq 0
$$

例如，向量：

$$
x_{1}=
\begin{bmatrix}
  1 \\
  2 \\
  3
\end{bmatrix} \quad
x_{2}=
\begin{bmatrix}
  4 \\
  1 \\
  5
\end{bmatrix} \quad
x_{3}=
\begin{bmatrix}
  2 \\
  -3 \\
  -1
\end{bmatrix}
$$

是线性相关的，因为：$x_3=-2x_1+x_2$。

矩阵$A  \in \mathbb{R}^{m \times n}$的**列秩**<span id="column_rank"></span>是构成线性无关集合的$A$的最大列子集的大小。 由于术语的多样性，这通常简称为$A$的线性无关列的数量。同样，**行秩**<span id="row_rank"></span>是构成线性无关集合的$A$的最大行数。
对于任何矩阵$A  \in \mathbb{R}^{m \times n}$，事实证明$A$的列秩等于$A$的行秩（尽管我们不会证明这一点），因此两个量统称为$A$的**秩**<span id="rank"></span>，用 $\text{rank}(A)$表示。 以下是秩的一些基本性质：

- 对于 $A  \in \mathbb{R}^{m \times n}$，$\text{rank}(A) \leq min(m, n)$，如果$ \text(A) = \text{min} (m, n)$，则： $A$ 被称作**满秩**<span id="full_rank"></span>。
- 对于 $A  \in \mathbb{R}^{m \times n}$， $\text{rank}(A) = \text{rank}(A^T)$
- 对于 $A  \in \mathbb{R}^{m \times n}$,$B  \in \mathbb{R}^{n \times p}$ ,$\text{rank}(AB) \leq \text{min} ( \text{rank}(A), \text{rank}(B))$
- 对于 $A,B \in \mathbb{R}^{m \times n}$，$\text{rank}(A + B) \leq \text{rank}(A) + \text{rank}(B)$

### 3.7 方阵的逆

方阵$A  \in \mathbb{R}^{n \times n}$的**逆**<span id="inverse"></span>表示为$A^{-1}$，并且是这样的唯一矩阵:

$$
A^{-1}A=I=AA^{-1}
$$

请注意，并非所有矩阵都具有逆。 例如，非方形矩阵根据定义没有逆(存在**伪逆**[^4])。 然而，对于一些方形矩阵$A$，$A^{-1}$也可能不存在。 特别是，如果$A^{-1}$存在，我们说$A$是**可逆**<span id="invertible"></span>的或**非奇异**<span id="non-singular"></span>的，否则就是**不可逆**<span id="non-invertible"></span>或**奇异**<span id="singular"></span>的[^5]。

为了使方阵 A 具有逆$A^{-1}$，则$A$必须是满秩。 我们很快就会发现，除了满秩之外，还有许多其它的充分必要条件。
以下是逆的性质; 假设$A,B  \in \mathbb{R}^{n \times n}$，而且是非奇异的：

- $(A^{-1})^{-1} = A$
- $(AB)^{-1} = B^{-1}A^{-1}$
- $(A^{-1})^{T} =(A^{T})^{-1} $因此，该矩阵通常表示为$A^{-T}$。

作为如何使用逆的示例，考虑线性方程组，$Ax = b$，其中$A  \in \mathbb{R}^{n \times n}$，$x,b\in \mathbb{R}$， 如果$A$是非奇异的（即可逆的），那么$x = A^{-1}b$。 （如果$A  \in \mathbb{R}^{m \times n}$，不是方阵，这公式还有用吗？ - **伪逆**[^4]）

[^4]: 参考[Moore–Penrose inverse](https://en.jinzhao.wiki/wiki/Moore%E2%80%93Penrose_inverse)
[^5]: 很容易混淆并认为非奇异意味着不可逆。 但实际上，意思正好相反！ 小心！

### 3.8 正交矩阵

如果 $x^Ty=0$，则两个向量$x,y\in \mathbb{R}^{n}$ 是**正交**<span id="orthogonal"></span>的。如果$\|x\|_2=1$，则向量$x\in \mathbb{R}^{n}$ 被**归一化**<span id="normalized"></span>。如果一个方阵$U\in \mathbb{R}^{n \times n}$的所有列彼此正交并被归一化，则方阵$U$是**正交矩阵**（注意在讨论向量与矩阵时的意义不一样，**两个向量正交只需要内积为 0，正交矩阵是各列相互正交并且被归一化**）。

它可以从正交性和正态性的定义中得出:

$$
U^ TU = I = U U^T
$$

换句话说，正交矩阵的逆是其转置。 注意，如果$U$不是方阵， 即，$U\in \mathbb{R}^{m \times n}, n < m $ ，但其列仍然是正交的，则$U^TU = I$，但是$UU^T \neq I$。所以**正交矩阵一定是方阵**。

正交矩阵的另一个好的特性是在具有正交矩阵的向量上操作不会改变其欧几里德范数，即(i.e.):

$$
\|U x\|_{2}=\|x\|_{2}
\label{3}\tag{3}
$$

对于任何 $x\in \mathbb{R}^n$ , $U\in \mathbb{R}^{n \times n}$是正交矩阵。

### 3.9 矩阵的值域和零空间

**张成**<span id="span"></span>一个向量集合$\{ x_1,x_2, \cdots x_n \} $可以表示为一个向量集合$\{ x_1, \cdots x_n \} $的所以线性组合：

$$
\operatorname{span}(\{x_1, \cdots x_n \}) = \Bigg\{v:v=\sum_{i=1}^n{\alpha_i x_i}, \alpha_i \in \Bbb{R} \Bigg\}
$$

可以看到，如果$\{x_{1}, \cdots x_{n}\}$是一组$n$个线性无关的向量，其中每个$x_i \in \mathbb{R}^{n}$，则$\text{span}(\{x_{1}, \ldots x_{n}\})=\mathbb{R}^{n}$。 换句话说，任何向量$v\in \mathbb{R}^{n}$都可以写成$x_1$到$x_n$的线性组合。
向量$y\in \mathbb{R}^{m}$**投影**<span id="projection"></span>到$\{x_{1}, \ldots x_{n}\}$所张成的空间（这里我们假设$x_i \in \mathbb{R}^{m}$）得到向量$v \in \operatorname{span}(\{x_{1}, \ldots, x_{n}\})$，由欧几里德范数$\|v  -  y\|_2$可以得知，这样$v$尽可能接近$y$。

我们将投影表示为$\operatorname{Proj}\left(y ;\left\{x_{1}, \ldots x_{n}\right\}\right)$，并且可以将其正式定义为:

$$
\operatorname{Proj}\left(y ;\left\{x_{1}, \ldots x_{n}\right\}\right)=\operatorname{argmin}_{v \in \operatorname{span}\left(\left\{x_{1}, \ldots, x_{n}\right\}\right)}\|y-v\|_{2}
$$

矩阵$A\in \mathbb{R}^{m \times n}$的**值域**<span id="range"></span>（有时也称为**列空间**<span id="columnspace"></span>），表示为$\mathcal{R}(A)$，是$A$的列所张成的空间。换句话说，

$$
\mathcal{R}(A)=\left\{v \in \mathbb{R}^{m} : v=A x, x \in \mathbb{R}^{n}\right\}
$$

做一些技术性的假设（即$A$是满秩且$n <m$），向量$y \in \mathbb{R}^{m}$到$A$的值域的投影由下式给出:

$$
\operatorname{Proj}(y ; A)=\operatorname{argmin}_{v \in \mathcal{R}(A)}\|v-y\|_{2}=A\left(A^{T} A\right)^{-1} A^{T} y
$$

这个最后的方程应该看起来非常熟悉，因为它几乎与我们在课程中（我们将很快再次得出）得到的公式：与参数的最小二乘估计一样。
看一下投影的定义，显而易见，这实际上是我们在最小二乘问题中最小化的目标（除了范数的平方这里有点不一样，这不会影响找到最优解），所以这些问题自然是非常相关的。

当$A$只包含一列时，$a \in \mathbb{R}^{m}$，这给出了向量投影到一条线上的特殊情况：

$$
\operatorname{Proj}(y ; a)=\frac{a a^{T}}{a^{T} a} y
$$

一个矩阵$A\in \mathbb{R}^{m \times n}$的**零空间**<span id="nullspace"></span> $\mathcal{N}(A)$ 是所有乘以$A$时等于 0 向量的集合，即：

$$
\mathcal{N}(A)=\left\{x \in \mathbb{R}^{n} : A x=0\right\}
$$

注意，$\mathcal{R}(A)$中的向量的大小为$m$，而 $\mathcal{N}(A)$ 中的向量的大小为$n$，因此$\mathcal{R}(A^T)$和 $\mathcal{N}(A)$ 中的向量的大小均为$\mathbb{R}^{n}$。 事实上，还有很多例子。 证明：

$$
\left\{w : w=u+v, u \in \mathcal{R}\left(A^{T}\right), v \in \mathcal{N}(A)\right\}=\mathbb{R}^{n} \text { and } \mathcal{R}\left(A^{T}\right) \cap \mathcal{N}(A)=\{\mathbf{0}\}
$$

换句话说，$\mathcal{R}(A^T)$和 $\mathcal{N}(A)$ 是不相交的子集，它们一起跨越$\mathbb{R}^{n}$的整个空间。 这种类型的集合称为**正交补**<span id="orthogonal_complements"></span>，我们用$\mathcal{R}(A^T)= \mathcal{N}(A)^{\perp}$表示。

### 3.10 行列式

一个方阵$A  \in \mathbb{R}^{n \times n}$的**行列式**<span id="determinant"></span>是函数$\text {det}$：$\mathbb{R}^{n \times n} \rightarrow \mathbb{R}^{n} $，并且表示为$\left| A \right|$或者$\text{det} A$（有点像迹运算符，我们通常省略括号）。在代数上，我们可以写出 A 的行列式的明确公式，但不幸的是，这并不能直观地理解它的含义。 相反，我们将从提供行列式的几何解释开始，然后访问其一些特定的代数性质。

给定一个矩阵：

$$
\begin{bmatrix}
  \text{\textemdash} & a_1^T & \text{\textemdash} \\
  \text{\textemdash} & a_2^T & \text{\textemdash} \\
   & \vdots &  \\
  \text{\textemdash} & a_n^T & \text{\textemdash} \\
\end{bmatrix}
$$

考虑通过采用$A$行向量$a_{1}, \ldots a_{n}\in  \mathbb{R}^{n}$的所有可能线性组合形成的点$S \subset \mathbb{R}^{n}$的集合，其中线性组合的系数都在 0 和 1 之间; 也就是说，集合$S$是$\text{span}(\{a_{1}, \ldots a_{n}\})$受到系数$\alpha_{1}, \ldots \alpha_{n}$的限制的线性组合，$\alpha_1, \cdots ,\alpha_n$满足$0 \leq \alpha_{i} \leq 1, i=1, \ldots, n$。从形式上看，

$$
S=\left\{v \in \mathbb{R}^{n} : v=\sum_{i=1}^{n} \alpha_{i} a_{i} \text { where } 0 \leq \alpha_{i} \leq 1, i=1, \ldots, n\right\}
$$

事实证明，$A$的行列式的绝对值是对集合$S$的“体积”的度量[^6]。

[^6]: 诚然，我们实际上并没有定义我们在这里所说的“体积”是什么意思，但希望直觉应该足够清楚。 当 $n = 2$ 时，我们的“体积”概念对应于笛卡尔平面中 $S$ 的面积。 当 $n = 3$ 时，“体积”对应于我们通常的三维物体体积概念。

比方说：一个$2 \times2$的矩阵(4)：

$$
A=
\begin{bmatrix}
  1 & 3 \\
  3 & 2
\end{bmatrix}
\label{4}\tag{4}
$$

它的矩阵的行是：

$$
a_{1}=\left[\begin{array}{l}{1} \\ {3}\end{array}\right]
\quad
a_{2}=\left[\begin{array}{l}{3} \\ {2}\end{array}\right]
$$

对应于这些行对应的集合$S$如图 1 所示。对于二维矩阵，$S$通常具有平行四边形的形状。 在我们的例子中，行列式的值是$\left| A \right| = -7$（可以使用本节后面显示的公式计算），因此平行四边形的面积为 7。（请自己验证！）

在三维中，集合$S$对应于一个称为平行六面体的对象（一个有倾斜边的三维框，这样每个面都有一个平行四边形）。行定义$S$的$3×3$矩阵 S 的行列式的绝对值给出了平行六面体的三维体积。在更高的维度中，集合$S$是一个称为$n$维平行体的对象。

<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" width="241px" height="226px" viewBox="-0.5 -0.5 241 226" content="&lt;mxfile&gt;&lt;diagram id=&quot;sW9IoM3PunykceuHWi1C&quot; name=&quot;Page-1&quot;&gt;7VjLcpswFP0aZpyFOwjxcJbBSdtFO9OZLLpWQAZNAblCfuXrKywJkMGxk0Azde2FjQ/SFbrnnoMkC87z7ReGlul3GuPMcux4a8F7y3FcMBPfFbCTwMy2JZAwEksINMAjecYK1M1WJMal0ZBTmnGyNMGIFgWOuIEhxujGbLagmTnqEiW4AzxGKOuiP0nMUzULJ2jwr5gkqR4Z+LfyTo50YzWTMkUx3bQg+GDBOaOUy6t8O8dZlTudF9nv85G79YMxXPBzOkDZYY2ylZqbei6+05PFRXxX5Uz8izJUliSyYJjyPBMAEJclZ/QXntOMsn0HaO8/4o4MheNOMpunA/WcRa1gmmPOdqLJpsmqpzKVthKqMYYzxMnaDI8UuUkdrh7hByViYMdWdejoslNl6Lq2GaKkKxZh1audxYNA7qlAHLEE804gcdGadgPtSeonzL0SNgBhh4Gc8Qi77SHMz0RewgXdP1JUE+H/XlW6DxtKWpCfVL8TgYkhbftGRxHjy0CyQbcaBJvf0JMwYKMIUEaSoqoQwSwWw4drzDgRFnenbuQkjqsYIcMleUZP+3hVkSyrpOzT5IWWd1+PWQXA2z7/VZ0b12sXlNtfUCrQ1P6kXVVRNQWDlNJUd1FhA2hGoItFid/LvTeUWPVbBl7la8hVBYLOaPIFoMMh0tIrV09adqAlxxb8EWqs3OB4dbxRpd6LKhUi9YZRJfBMVY4gStB9hU6AdFV4898QFthmpp2B+HNMsx6FQf9qq6PYam2j2lb98WzV7uHw7cuiPkt2jlnyZayb/BMCF3vOd2paVQEYxYS7G88JlCbsXI4Jn+LId/1DvxzEEaA9AmXBq1y3oAV+yXL/8Y3oqQXo2Zbr/7WNKOjuRiaulJx3OZILTkgucPWZl6ZuEJOc6iPNQSU3u0ru/MXJx0tO/G1ObmXz5vgbPvwB&lt;/diagram&gt;&lt;/mxfile&gt;"><defs/><g><path d="M 32 216 L 225.63 216" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 230.88 216 L 223.88 219.5 L 225.63 216 L 223.88 212.5 Z" fill="#000000" stroke="#000000" stroke-miterlimit="10" pointer-events="all"/><path d="M 32 216 L 32 22.37" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 32 17.12 L 35.5 24.12 L 32 22.37 L 28.5 24.12 Z" fill="#000000" stroke="#000000" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject style="overflow: visible; text-align: left;" pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 217px; margin-left: 13px;"><div style="box-sizing: border-box; font-size: 0; text-align: center; "><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: #000000; line-height: 1.2; pointer-events: all; background-color: #ffffff; white-space: nowrap; "><font color="#000000">(0, 0)</font></div></div></div></foreignObject><text x="13" y="220" fill="#000000" font-family="Helvetica" font-size="11px" text-anchor="middle">(0, 0)</text></switch></g><path d="M 32 216 L 68.8 105.59" fill="none" stroke="#000000" stroke-width="3" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 70.94 99.18 L 72.36 109.14 L 68.8 105.59 L 63.82 106.3 Z" fill="#000000" stroke="#000000" stroke-width="3" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject style="overflow: visible; text-align: left;" pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 149px; margin-left: 42px;"><div style="box-sizing: border-box; font-size: 0; text-align: center; "><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: #000000; line-height: 1.2; pointer-events: all; background-color: #ffffff; white-space: nowrap; ">a<sub>1</sub></div></div></div></foreignObject><text x="42" y="152" fill="#000000" font-family="Helvetica" font-size="11px" text-anchor="middle">a1</text></switch></g><g transform="translate(-0.5 -0.5)"><switch><foreignObject style="overflow: visible; text-align: left;" pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 86px; margin-left: 53px;"><div style="box-sizing: border-box; font-size: 0; text-align: center; "><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: #000000; line-height: 1.2; pointer-events: all; background-color: #ffffff; white-space: nowrap; ">(1, 3)</div></div></div></foreignObject><text x="53" y="90" fill="#000000" font-family="Helvetica" font-size="11px" text-anchor="middle">(1, 3)</text></switch></g><path d="M 32 216 L 143.59 141.6" fill="none" stroke="#000000" stroke-width="3" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 149.21 137.86 L 144.22 146.6 L 143.59 141.6 L 139.22 139.11 Z" fill="#000000" stroke="#000000" stroke-width="3" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject style="overflow: visible; text-align: left;" pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 187px; margin-left: 102px;"><div style="box-sizing: border-box; font-size: 0; text-align: center; "><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: #000000; line-height: 1.2; pointer-events: all; background-color: #ffffff; white-space: nowrap; "><font color="#000000">a<sub>2</sub></font></div></div></div></foreignObject><text x="102" y="190" fill="#000000" font-family="Helvetica" font-size="11px" text-anchor="middle">a2</text></switch></g><g transform="translate(-0.5 -0.5)"><switch><foreignObject style="overflow: visible; text-align: left;" pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 152px; margin-left: 163px;"><div style="box-sizing: border-box; font-size: 0; text-align: center; "><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: #000000; line-height: 1.2; pointer-events: all; background-color: #ffffff; white-space: nowrap; ">(3, 2)</div></div></div></foreignObject><text x="163" y="155" fill="#000000" font-family="Helvetica" font-size="11px" text-anchor="middle">(3, 2)</text></switch></g><path d="M 72 96 L 192 16" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="stroke"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject style="overflow: visible; text-align: left;" pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 6px; margin-left: 176px;"><div style="box-sizing: border-box; font-size: 0; text-align: center; "><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: #000000; line-height: 1.2; pointer-events: all; background-color: #ffffff; white-space: nowrap; ">(4, 5)</div></div></div></foreignObject><text x="176" y="9" fill="#000000" font-family="Helvetica" font-size="11px" text-anchor="middle">(4, 5)</text></switch></g><path d="M 152 136 L 192 16" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="stroke"/></g><switch><g requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"/><a transform="translate(0,-5)" xlink:href="https://www.diagrams.net/doc/faq/svg-export-text-problems" target="_blank"><text text-anchor="middle" font-size="10px" x="50%" y="100%">Viewer does not support full SVG 1.1</text></a></switch></svg>

> 图 1：（4）中给出的$2×2$矩阵$A$的行列式的图示。 这里，$a_1$和$a_2$是对应于$A$行的向量，并且集合$S$对应于平行四边形区域。 这个行列式的绝对值，$\left| \text{det} A \right| = 7$，即平行四边形的面积。

在代数上，行列式满足以下三个性质（所有其他性质都遵循这些性质，包括通用公式）：

1. 单位阵的行列式为 1, $\left| I \right|= 1$（几何上，单位超立方体的体积为 1）。

2. 给定一个矩阵 $A  \in \mathbb{R}^{n \times n}$, 如果我们将$A$中的一行乘上一个标量$t  \in \mathbb{R}$，那么新矩阵的行列式是$t\left| A \right|$

   $$
   \left|\begin{bmatrix}
     \text{\textemdash} & t a_1^T & \text{\textemdash} \\
     \text{\textemdash} & a_2^T & \text{\textemdash} \\
      & \vdots &  \\
     \text{\textemdash} & a_m^T & \text{\textemdash} \\
   \end{bmatrix}\right| = t|A|
   $$

   几何上，将集合$S$的一个边乘以系数$t$，体积也会增加一个系数$t$。

3. 如果我们交换任意两行在$a_i^T$和$a_j^T$，那么新矩阵的行列式是$-\left| A \right|$，例如：
   $$
   \left|\begin{bmatrix}
     \text{\textemdash} & a_2^T & \text{\textemdash} \\
     \text{\textemdash} & a_1^T & \text{\textemdash} \\
      & \vdots &  \\
     \text{\textemdash} & a_m^T & \text{\textemdash} \\
   \end{bmatrix}\right| = -|A|
   $$
   你一定很奇怪，满足上述三个性质的函数的存在并不多。事实上，这样的函数确实存在，而且是唯一的（我们在这里不再证明了）。

从上述三个性质中得出的几个性质包括：

- 对于 $A  \in \mathbb{R}^{n \times n}$, $\left| A \right| = \left| A^T \right|$
- 对于 $A,B \in \mathbb{R}^{n \times n}$, $\left| AB \right|= \left| A \right|\left| B \right|$
- 对于 $A  \in \mathbb{R}^{n \times n}$,如果$\left| A \right|= 0$ 有且只有当$A$是奇异的（即不可逆）（如果 $A$ 是奇异的，那么它没有满秩，因此它的列是线性相关的。在这种情况下，集合 $S$ 对应于 $n$ 维空间中的“平面”，因此体积为零。）
- 对于 $A  \in \mathbb{R}^{n \times n}$ 同时，$A$为非奇异的，则：$\left| A ^{−1}\right| = 1/\left| A \right|$

在给出行列式的一般定义之前，我们定义，对于$A  \in \mathbb{R}^{n \times n}$，$A_{\bcancel i, \bcancel j}\in \mathbb{R}^{(n-1) \times (n-1)}$是由于删除第$i$行和第$j$列而产生的矩阵。 行列式的一般（递归）公式是：

$$
\begin{aligned}|A| &=\sum_{i=1}^{n}(-1)^{i+j} a_{i j}\left|A_{\backslash i, \backslash j}\right| \quad(\text { for any } j \in 1, \ldots, n) \\ &=\sum_{j=1}^{n}(-1)^{i+j} a_{i j}\left|A_{\backslash i, \backslash j}\right| \quad(\text { for any } i \in 1, \ldots, n) \end{aligned}
$$

对于 $A  \in \mathbb{R}^{1 \times 1}$，初始情况为$\left| A \right|= a_{11}$。如果我们把这个公式完全展开为 $A  \in \mathbb{R}^{n \times n}$，就等于$n!$（$n$阶乘）不同的项。因此，对于大于$3×3$的矩阵，我们几乎没有明确地写出完整的行列式方程。然而，$3×3$大小的矩阵的行列式方程是相当常见的，建议好好地了解它们：

$$
\begin{aligned}
\left|\left[a_{11}\right]\right| &=a_{11}
\\
\left|\left[\begin{array}{ll}{a_{11}} & {a_{12}} \\ {a_{21}} & {a_{22}}\end{array}\right]\right|&=a_{11} a_{22}-a_{12} a_{21}
\\
\left|\left[\begin{array}{l}{a_{11}} & {a_{12}} & {a_{13}} \\ {a_{21}} & {a_{22}} & {a_{23}} \\ {a_{31}} & {a_{32}} & {a_{33}}\end{array}\right]\right| &=
\begin{array}{c}{a_{11} a_{22} a_{33}+a_{12} a_{23} a_{31}+a_{13} a_{21} a_{32}} \\ {-a_{11} a_{23} a_{32}-a_{12} a_{21} a_{33}-a_{13} a_{22} a_{31}} \end{array}
\end{aligned}
$$

矩阵$A  \in \mathbb{R}^{n \times n}$的**经典伴随矩阵**<span id="classical_adjoint"></span>（通常称为<span id="adjoint"></span>**伴随矩阵**[^7]）表示为$\operatorname{adj}(A)$，并定义为：

$$
\operatorname{adj}(A) \in \mathbb{R}^{n \times n}, \quad(\operatorname{adj}(A))_{i j}=(-1)^{i+j}\left|A_{\backslash j, \backslash i}\right|
$$

（注意索引$A_{\backslash j, \backslash i}$中的变化）。可以看出，对于任何非奇异$A  \in \mathbb{R}^{n \times n}$，

$$
A^{-1}=\frac{1}{|A|} \operatorname{adj}(A)
$$

虽然这是一个很好的“显式”的逆矩阵公式，但我们应该注意，从数字上讲，有很多更有效的方法来计算逆矩阵。

[^7]:
    $A_{ij}$的**余子式**<span id="minor"></span>（余子式其实是一个数）表示为$M_{ij} = \left|A_{\backslash i, \backslash j}\right|$,就是删除第 i 行和第 j 列而产生矩阵的行列式；$A_{ij}$的**代数余子式**<span id="cofactor"></span>（代数余子式也是一个数）表示为$C_{ij} = (-1)^{i+j}M_{ij}$；$A$的**余子矩阵(代数余子式矩阵,记为 cof)**<span id="cofactor_matrix"></span>是一个$n$阶矩阵$C$，使得其第$i$行第$j$列的元素是$A$关于第$i$行第$j$列的代数余子式。则伴随矩阵的定义如下：

    $$
    A^* = \operatorname{adj}(A) = C^T \in \mathbb{R}^{n \times n}, \quad(\operatorname{adj}(A))_{i j}=(-1)^{i+j}\left|A_{\backslash j, \backslash i}\right| =C_{ji}
    $$

    伴随矩阵的一些性质(这里用$A^*$表示$\operatorname{adj}(A) $)：

    - $A$可逆当且仅当$A^*$可逆
    - $A$可逆，则$A^* = |A|A^{-1}$
    - $|A^*|=|A|^{n-1}$
    - $(kA)^*=k^{n-1}A^*$
    - $A$可逆，则$(A^{-1})^* = (A^*)^{-1}$
    - $(A^T)^* = (A^*)^T$
    - $(AB)^* = B^*A^*$
    - $\operatorname{rank}(A^*) = n, \operatorname{rank}(A) = n\\\operatorname{rank}(A^*) = 1, \operatorname{rank}(A) = n-1\\\operatorname{rank}(A^*) = 0, \operatorname{rank}(A) < n-1$

### 3.11 二次型和半正定矩阵

给定方矩阵$A  \in \mathbb{R}^{n \times n}$和向量$x \in \mathbb{R}^{n}$，标量$x^T Ax$被称为**二次型**<span id="quadratic_form"></span>。 写得清楚些，我们可以看到：

$$
x^{T} A x=\sum_{i=1}^{n} x_{i}(A x)_{i}=\sum_{i=1}^{n} x_{i}\left(\sum_{j=1}^{n} A_{i j} x_{j}\right)=\sum_{i=1}^{n} \sum_{j=1}^{n} A_{i j} x_{i} x_{j}
$$

注意：

$$
x^{T} A x=\left(x^{T} A x\right)^{T}=x^{T} A^{T} x=x^{T}\left(\frac{1}{2} A+\frac{1}{2} A^{T}\right) x
$$

小技巧：

$$
A = \frac{A+A^T}{2} + \frac{A-A^T}{2}\\,
\\
x^{T} A x = x^{T}\frac{A+A^T}{2}x + x^{T}\frac{A-A^T}{2}x = x^{T}\frac{A+A^T}{2}x +0
$$

第一个等号的是因为是标量的转置与自身相等，而第二个等号是因为是我们平均两个本身相等的量。 由此，我们可以得出结论，只有$A$的对称部分有助于形成二次型($A+A^T$是[对称的](#symmetric))。 出于这个原因，我们经常隐含地假设以二次型出现的矩阵是对称阵。
我们给出以下定义：

- 对于所有非零向量$x \in \mathbb{R}^n$，$x^TAx>0$，对称阵$A \in \mathbb{S}^n$为**正定(PD)**<span id="positive_definite"></span>。这通常表示为$A\succ0$（或$A>0$），并且通常将所有正定矩阵的集合表示为$\mathbb{S}_{++}^n$。

- 对于所有向量$x^TAx\geq 0$，对称矩阵$A \in \mathbb{S}^n$是**半正定(PSD)**<span id="positive_semidefinite"></span>。 这写为$A \succeq 0$（或$A≥0$），并且所有半正定矩阵的集合通常表示为$\mathbb{S}_+^n$。

- 同样，对称矩阵$A \in \mathbb{S}^n$是**负定(ND)**<span id="negative_definite"></span>，如果对于所有非零$x \in \mathbb{R}^n$，则$x^TAx <0$表示为$A\prec0$（或$A <0$）。

- 类似地，对称矩阵$A \in \mathbb{S}^n$是**半负定(NSD)**<span id="negative_semidefinite"></span>，如果对于所有$x \in \mathbb{R}^n$，则$x^TAx \leq 0$表示为$A\preceq 0$（或$A≤0$）。

- 最后，对称矩阵$A \in \mathbb{S}^n$是**不定**<span id="indefinite"></span>的，如果它既不是正半定也不是负半定，即，如果存在$x_1,x_2 \in \mathbb{R}^n$，那么$x_1^TAx_1>0$且$x_2^TAx_2<0$。

很明显，如果$A$是正定的，那么$−A$是负定的，反之亦然。同样，如果$A$是半正定的，那么$−A$是是半负定的，反之亦然。如果果$A$是不定的，那么$−A$是也是不定的。

正定矩阵和负定矩阵的一个重要性质是它们总是满秩，因此是可逆的。为了了解这是为什么，假设某个矩阵$A \in \mathbb{S}^n$不是满秩。然后，假设$A$的第$j$列可以表示为其他$n-1$列的线性组合：

$$
a_{j}=\sum_{i \neq j} x_{i} a_{i}
$$

对于某些$x_1,\cdots x_{j-1},x_{j + 1} ,\cdots ,x_n\in \mathbb{R}$。设$x_j = -1$，则：

$$
Ax=\sum_{i =1}^n x_{i} a_{i}=0
$$

但这意味着对于某些非零向量$x$，$x^T Ax = 0$，因此$A$必须既不是正定也不是负定。如果$A$是正定或负定，则必须是满秩。
最后，有一种类型的正定矩阵经常出现，因此值得特别提及。 给定矩阵$A  \in \mathbb{R}^{m \times n}$（不一定是对称或偶数平方），矩阵$G = A^T A \in \Bbb{R}^{n \times n}$（有时称为**Gram 矩阵**<span id="gram_matrix"></span>）总是半正定的。 此外，如果$m\geq n$（同时为了方便起见，我们假设$A$是满秩,即$\operatorname{rank}(A) = n$，则$G = A^T A$是正定的。

> $AA^T$(即 Gram 矩阵)是半正定矩阵；协方差矩阵要是半正定矩阵
> 正定矩阵的所有特征值为正实数
> 半正定矩阵的所有特征值为非负实数
> 负定矩阵的所有特征值为负实数
> 半负定矩阵的所有特征值为非正实数
> 不定矩阵的特征值既有正数也有负数

### 3.12 特征值和特征向量

给定一个方阵$A \in\mathbb{R}^{n\times n}$，我们认为在以下条件下，$\lambda \in\mathbb{C}$是$A$的**特征值**<span id="eigenvalue"></span>，$x\in\mathbb{C}^n$是相应的**特征向量**[^8]<span id="eigenvector"></span>：

[^8]: 请注意，$\lambda$ 和 $x$ 的元素实际上在 $C $中，即复数集，而不仅仅是实数； 我们很快就会看到为什么这是必要的。 现在不要担心这个问题，你可以像实向量一样思考复向量。

$$
Ax=\lambda x,x \ne 0
$$

直观地说，这个定义意味着将$A$乘以向量$x$会得到一个新的向量，该向量指向与$x$相同的方向，但按系数$\lambda$缩放。
值得注意的是，对于任何特征向量$x\in\mathbb{C}^n$和标量$c\in\mathbb{C}$，$A(cx)=cAx=c\lambda x=\lambda(cx)$，$cx$也是一个特征向量。因此，当我们讨论与$\lambda$相关的特征向量时，我们通常假设特征向量被标准化为长度为 1（这仍然会造成一些歧义，因为$x$和$−x$都是特征向量，但我们必须接受这一点）。

我们可以重写上面的等式来说明$(\lambda,x)$是$A$的特征值-特征向量对：

$$
(\lambda I-A)x=0,x \ne 0
$$

但是$(\lambda I-A)x=0$只有当$(\lambda I-A)$有一个非空零空间时，同时$(\lambda I-A)$是奇异的，$x$才具有非零解，即：

$$
|(\lambda I-A)|=0
$$

我们现在可以使用之前的行列式定义来展开这个表达式式$|(\lambda I-A)|$为$\lambda$的（非常大的）多项式，其中，$\lambda$的次数为$n$。它通常被称为矩阵$A$的特征多项式。

然后我们找到这个特征多项式的$n$个根（可能是复数），并用$\lambda_1,\cdots,\lambda_n$表示。这些都是矩阵$A$的特征值，但我们注意到它们可能不明显。为了找到特征值$\lambda_i$对应的特征向量，我们只需解线性方程$(\lambda I-A)x=0$，因为$(\lambda I-A)$是奇异的，所以保证有一个非零解（但也可能有多个或无穷多个解）。

应该注意的是，这不是实际用于数值计算特征值和特征向量的方法（记住行列式的完全展开式有$n!$项），这是一个数学论证。

以下是特征值和特征向量的性质（所有假设在$A \in\mathbb{R}^{n\times n}$具有特征值$\lambda_1,\cdots,\lambda_n$的前提下）：

- $A$的迹等于其特征值之和,也等于对角元素之和(迹的定义)

  $$
  \operatorname{tr} A=\sum_{i=1}^{n} \lambda_{i} =\sum_{i=1}^{n} A_{ii}
  $$

- $A$的行列式等于其特征值的乘积

  $$
  |A|=\prod_{i=1}^{n} \lambda_{i}
  $$

- $A$的秩等于$A$的非零特征值的个数
- 假设$A$非奇异，其特征值为$\lambda$和特征向量为$x$。那么$1/\lambda$是具有相关特征向量$x$的$A^{-1}$的特征值，即$A^{-1}x=(1/\lambda)x$。（要证明这一点，取特征向量方程，$Ax=\lambda x$，两边都左乘$A^{-1}$）
- 对角阵的特征值$D=\operatorname{diag}(d_1,\cdots d_n)$实际上就是对角元素$d_1,\cdots d_n$

### 3.13 对称矩阵的特征值和特征向量

一般而言，一般方阵的特征值和特征向量的结构很难表征。 幸运的是，在机器学习的大多数情况下，处理对称实矩阵就足够了，其特征值和特征向量具有显着的性质。

在本节中，我们假设$A$是实对称矩阵, 具有以下性质：

1. $A$的所有特征值都是实数。 我们用$\lambda_1,\cdots,\lambda_n$表示。

2. 存在一组特征向量$u_1,\cdots,u_n$，对于所有$i$，$u_i$是特征值$\lambda_{i}$对应的特征向量。以及$u_1,\cdots,u_n$是单位向量并且彼此正交[^9]。

[^9]: 在数学上，我们有$\forall{i},Au_i = \lambda_iu_i, \|u_i\|_2 = 1, \text{and } \forall{j} \neq i, u_i^Tu_j = 0$。此外，我们注意到任意矩阵 A（而这里我们主要描述对称矩阵）的特征向量，并不是都满足彼此正交，因为特征值可以是重复的，特征向量也是如此。

设$U$是包含$u_i$作为列的正交矩阵[^10]：

[^10]: 这里为了符号的简单性，我们偏离了前几节中矩阵列的符号约定(本来是应该用$u^i$表示的，这里我们用$u_i$来表示)。

$$
U=\left[\begin{array}{cccc}{ |} & { |} & {} & { |} \\ {u_{1}} & {u_{2}} & {\cdots} & {u_{n}} \\ { |} & { |} & {} & { |}\end{array}\right]
\label{5}\tag{5}
$$

设$\Lambda= \operatorname{diag}(\lambda_1,\cdots,\lambda_n)$是包含$\lambda_1,\cdots,\lambda_n$作为对角线上的元素的对角矩阵。 使用 2.3 节的方程$\eqref{2}$中的矩阵 - 矩阵向量乘法的方法，我们可以验证：

$$
A U=\left[\begin{array}{cccc}
{ |} & { |} & {} & { |} \\
{A u_{1}} & {A u_{2}} & {\cdots} & {A u_{n}} \\
{ |} & { |} & {} & { |}\end{array}\right]=
\left[\begin{array}{ccc}
{ |} & { |} & { } & { |}\\
{\lambda_{1} u_{1}} & {\lambda_{2} u_{2}} & {\cdots} & {\lambda_{n} u_{n}} \\
{ |} & { |} & {} & { |}
\end{array}\right]=
U \operatorname{diag}\left(\lambda_{1}, \ldots, \lambda_{n}\right)=U \Lambda
$$

考虑到正交矩阵$U$满足$UU^T=I$，利用上面的方程，我们得到：

$$
A=AUU^T=U\Lambda U^T
\label{6}\tag{6}
$$

这种$A$的新的表示形式为$U\Lambda U^T$，通常称为矩阵$A$的**对角化**<span id="diagonalizing"></span>。术语对角化是这样来的：通过这种表示，我们通常可以有效地将对称矩阵$A$视为对角矩阵--这更容易理解--关于由特征向量$U$定义的基础， 我们将通过几个例子详细说明。

**背景知识：关于另一个基的向量**。

任何正交矩阵$U=\left[\begin{array}{cccc}{ |} & { |} & {} & { |} \\ {u_{1}} & {u_{2}} & {\cdots} & {u_{n}} \\ { |} & { |} & {} & { |}\end{array}\right]$定义了一个新的属于$\mathbb {R}^{n}$的基（坐标系），意义如下：对于任何向量$x \in\mathbb{R}^{n}$都可以表示为$u_1,\cdots,u_n$的线性组合，其系数为$\hat x_1,\cdots,\hat x_n$：

$$
x=\hat x_1u_1+\cdots + \hat x_nu_n=U\hat x
$$

在第二个等式中，我们使用矩阵和向量相乘的方法,查看式$\eqref{1}$。 实际上，这种$\hat x$是唯一存在的:

$$
x=U \hat{x} \Leftrightarrow U^{T} x=\hat{x}
$$

换句话说，向量$\hat x=U^Tx$可以作为向量$x$的另一种表示，与$U$定义的基有关。

**“对角化”矩阵向量乘法**。 通过上面的设置，我们将看到左乘矩阵$A$可以被视为左乘对角矩阵，也就是特征向量组成的基。 假设$x$是一个向量，$\hat x$是以$U$为基$x$的表示。设$z=Ax$为矩阵向量积。现在让我们计算以$U$为基来表示$z$：
然后，再利用$UU^T=U^TU=I$和$A=AUU^T=U\Lambda U^T$，也就是式$\eqref{6}$，我们得到：

$$
\hat{z}=U^{T} z=U^{T} A x=U^{T} U \Lambda U^{T} x=\Lambda \hat{x}=\left[\begin{array}{c}{\lambda_{1} \hat{x}_{1}} \\ {\lambda_{2} \hat{x}_{2}} \\ {\vdots} \\ {\lambda_{n} \hat{x}_{n}}\end{array}\right]
$$

我们可以看到，原始空间中的左乘矩阵$A$等于左乘对角矩阵$\Lambda$相对于新的基，即仅将每个坐标缩放相应的特征值。
在新的基上，矩阵多次相乘也变得简单多了。例如，假设$q=AAAx$。根据$A$的元素导出$q$的分析形式，使用原始的基可能是一场噩梦，但使用新的基就容易多了：

$$
\hat{q}=
U^{T} q=
U^{T} AAA x=
U^{T} U \Lambda U^{T} U \Lambda U^{T} U \Lambda U^{T} x=
\Lambda^{3} \hat{x}=
\left[\begin{array}{c}
{\lambda_{1}^{3} \hat{x}_{1}} \\ {\lambda_{2}^{3} \hat{x}_{2}} \\ {\vdots} \\ {\lambda_{n}^{3} \hat{x}_{n}}
\end{array}\right]
\label{7}\tag{7}
$$

**“对角化”二次型**。作为直接推论，二次型$x^TAx$也可以在新的基上简化。

$$
x^{T} A x=x^{T} U \Lambda U^{T} x=\hat{x}^T \Lambda \hat{x}=\sum_{i=1}^{n} \lambda_{i} \hat{x}_{i}^{2}
\label{8}\tag{8}
$$

(回想一下，在旧的表示法中，$x^{T} A x=\sum_{i=1, j=1}^{n} x_{i} x_{j} A_{i j}$涉及一个$n^2$项的和，而不是上面等式中的$n$项。)利用这个观点，我们还可以证明矩阵$A$的正定性完全取决于其特征值的符号：

1. 如果所有的$\lambda_i>0$，则矩阵$A$正定的，因为对于任意的$\hat x \ne 0$,$x^{T} A x=\sum_{i=1}^{n} \lambda_{i} \hat{x}_{i}^{2}>0$[^11]

2. 如果所有的$\lambda_i\geq 0$，则矩阵$A$是为正半定，因为对于任意的$\hat x $,$x^{T} A x=\sum*{i=1}^{n} \lambda*{i} \hat{x}\_{i}^{2} \geq 0$

3. 同样，如果所有$\lambda_i<0$或$\lambda_i\leq 0$，则矩阵$A$分别为负定或半负定。

4. 最后，如果$A$同时具有正特征值和负特征值，比如 $\lambda_i>0$和$\lambda_j<0$，那么它是不定的。这是因为如果我们让$\hat x$满足$\hat x_i=1 \text{ and } \hat x_k=0, \forall k\ne i$，那么$x^{T} A x=\sum_{i=1}^{n} \lambda_{i} \hat{x}_{i}^{2}>0$ ,同样的我们让$\hat x$满足$\hat x_j=1 \text{ and } \hat x_k=0,\forall k\ne j$，那么$x^{T} A x=\sum_{i=1}^{n} \lambda_{i} \hat{x}_{i}^{2}<0$[^12]

[^11]: 注意$\hat x \ne 0  \hArr x \ne 0$
[^12]: 注意$x=U \hat x$,因此构造$\hat x$给出来$x$的隐式构造

特征值和特征向量经常出现的应用是最大化矩阵的某些函数。特别是对于矩阵$A \in \mathbb{S}^{n}$，考虑以下最大化问题：

$$
\max _{x \in \mathbb{R}^{n}} \ x^{T} A x=\sum_{i=1}^{n} \lambda_{i} \hat{x}_{i}^{2} \quad \text { subject to }\|x\|_{2}^{2}=1
\label{9}\tag{9}
$$

也就是说，我们要找到（范数 1）的向量，它使二次型最大化。假设特征值的阶数为$\lambda_1 \geq \lambda _2 \geq \cdots \lambda_n$，此优化问题的最优值为$\lambda_1$，且与$\lambda_1$对应的任何特征向量$u_1$都是最大值之一。（如果$\lambda_1 > \lambda_2$，那么有一个与特征值$\lambda_1$对应的唯一特征向量，它是上面那个优化问题$\eqref{9}$的唯一最大值。）

我们可以通过使用对角化技术来证明这一点：注意，通过公式$\|U x\|_{2}\overset{\eqref{3}}{=}\|x\|_{2} $推出$\|x\|_{2}=\|\hat{x}\|_{2}$，并利用公式$x^{T} A x=x^{T} U \Lambda U^{T} x=\hat{x}^T \Lambda \hat{x}=\sum*{i=1}^{n} \lambda*{i} \hat{x}\_{i}^{2} ,\text{也就是式}\eqref{8}$，我们可以将上面那个优化问题改写为：

$$
\max _{\hat{x} \in \mathbb{R}^{n}}\ \hat{x}^{T} \Lambda \hat{x}=\sum_{i=1}^{n} \lambda_{i} \hat{x}_{i}^{2} \quad \text { subject to }\|\hat{x}\|_{2}^{2}=1
$$

然后，我们得到目标的上界为$\lambda_1$：

$$
\hat{x}^{T} \Lambda \hat{x}=\sum_{i=1}^{n} \lambda_{i} \hat{x}_{i}^{2} \leq \sum_{i=1}^{n} \lambda_{1} \hat{x}_{i}^{2}=\lambda_{1}
$$

此外，设置$\hat{x}=\left[\begin{array}{c}{1} \\ {0} \\ {\vdots} \\ {0}\end{array}\right]$可让上述等式成立，这与设置$x=u_1$相对应。

## 4.矩阵微积分

虽然前几节中的主题通常在线性代数的标准课程中涵盖，但一个似乎不经常涉及（我们将广泛使用）的主题是微积分对向量设置的扩展。 尽管我们使用的所有实际微积分都相对微不足道，但符号通常会使事情看起来比实际困难得多。 在本节中，我们将介绍矩阵微积分的一些基本定义并提供一些示例。

### 4.1 梯度

假设$f: \mathbb{R}^{m \times n} \rightarrow \mathbb{R}$是将维度为$m \times n$的矩阵$A\in \mathbb{R}^{m \times n}$作为输入并返回实数值的函数。 然后$f$的**梯度**<span id="gradient"></span>（相对于$A\in \mathbb{R}^{m \times n}$）是偏导数矩阵，定义如下：

$$
\nabla_{A} f(A) \in \mathbb{R}^{m \times n}=\left[\begin{array}{cccc}{\frac{\partial f(A)}{\partial A_{11}}} & {\frac{\partial f(A)}{\partial A_{12}}} & {\cdots} & {\frac{\partial f(A)}{\partial A_{1n}}} \\ {\frac{\partial f(A)}{\partial A_{21}}} & {\frac{\partial f(A)}{\partial A_{22}}} & {\cdots} & {\frac{\partial f(A)}{\partial A_{2 n}}} \\ {\vdots} & {\vdots} & {\ddots} & {\vdots} \\ {\frac{\partial f(A)}{\partial A_{m 1}}} & {\frac{\partial f(A)}{\partial A_{m 2}}} & {\cdots} & {\frac{\partial f(A)}{\partial A_{m n}}}\end{array}\right]
$$

即，$m \times n$矩阵:

$$
\left(\nabla_{A} f(A)\right)_{i j}=\frac{\partial f(A)}{\partial A_{i j}}
$$

**请注意**，$\nabla_{A} f(A) $的维度始终与$A$的维度相同。特殊情况，如果$A$只是向量$A\in \mathbb{R}^{n}$，则

$$
\nabla_{x} f(x)=\left[\begin{array}{c}{\frac{\partial f(x)}{\partial x_{1}}} \\ {\frac{\partial f(x)}{\partial x_{2}}} \\ {\vdots} \\ {\frac{\partial f(x)}{\partial x_{n}}}\end{array}\right]
$$

重要的是要记住，只有当函数是实值时，即如果函数返回标量值，才定义函数的梯度。例如，$A\in \mathbb{R}^{m \times n}$相对于$x$，我们不能取$Ax$的梯度，因为这个量(输出)是向量值。

直接从偏导数的等价性质得出：

- $\nabla_{x}(f(x)+g(x))=\nabla_{x} f(x)+\nabla_{x} g(x)$

- $\text{For }t \in \mathbb{R},\nabla_{x}(t f(x))=t \nabla_{x} f(x)$

原则上，梯度是偏导数对多维变量函数的自然延伸。然而，在实践中，由于符号的原因，使用梯度有时是很棘手的。例如，假设$A\in \mathbb{R}^{m \times n}$是一个固定系数矩阵，假设$b\in \mathbb{R}^{m}$是一个固定系数向量。设$f: \mathbb{R}^{m} \rightarrow \mathbb{R}$为$f(z)=z^Tz$定义的函数，因此$\nabla_{z}f(z)=2z$。但现在考虑表达式，

$$
\nabla f(Ax)
$$

该表达式应该如何解释？ 至少有两种可能性：

1. 在第一个解释中，回想起$\nabla_{z}f(z)=2z$。 在这里，我们将$\nabla f(Ax)$解释为评估点$Ax$处的梯度，因此:

$$
\nabla f(A x)=2(A x)=2 A x \in \mathbb{R}^{m}
$$

2. 在第二种解释中，我们将数量$f(Ax)$视为输入变量$x$的函数。 更正式地说，设$g(x) =f(Ax)$。 然后在这个解释中:

$$
\nabla f(A x)=\nabla_{x} g(x) \in \mathbb{R}^{n}
$$

在这里，我们可以看到这两种解释确实不同。 一种解释产生$m$维向量作为结果，而另一种解释产生$n$维向量作为结果($x$的维度是$n$，所以$\nabla_{x} g(x)$也是$n$，上面有讲到)！ 我们怎么解决这个问题？

这里，关键是要明确我们要区分的变量。
在第一种情况下，我们将函数$f$与其参数$z$进行区分，然后替换参数$Ax$。
在第二种情况下，我们将复合函数$g(x)=f(Ax)$直接与$x$进行微分。

我们将第一种情况表示为$\nabla zf(Ax)$，第二种情况表示为$\nabla xf(Ax)$[^13]。

保持符号清晰是非常重要的，以后完成课程作业时候你就会发现。

[^13]: 我们必须接受这种符号的一个缺点是，在第一种情况下，$\nabla zf(Ax)$ 似乎我们正在对一个变量进行微分，而这个变量甚至没有出现在被微分的表达式中！ 出于这个原因，第一种情况通常写为 $\nabla f(Ax)$，并且可以理解我们对 $f$ 的参数进行微分这一事实。 然而，第二种情况总是写成 $\nabla xf(Ax)$。

### 4.2 黑塞矩阵

假设$f: \mathbb{R}^{n} \rightarrow \mathbb{R}$是一个函数，它接受$\mathbb{R}^{n}$中的向量并返回实数。那么关于$x$的**黑塞矩阵**<span id="hessian"></span>（也有翻译作海森矩阵），写做：$\nabla_x ^2 f(A x)$，或者简单地说，$H$是$n \times n$的偏导数矩阵：

$$
\nabla_{x}^{2} f(x) \in \mathbb{R}^{n \times n}=\left[\begin{array}{cccc}{\frac{\partial^{2} f(x)}{\partial x_{1}^{2}}} & {\frac{\partial^{2} f(x)}{\partial x_{1} \partial x_{2}}} & {\cdots} & {\frac{\partial^{2} f(x)}{\partial x_{1} \partial x_{n}}} \\ {\frac{\partial^{2} f(x)}{\partial x_{2} \partial x_{1}}} & {\frac{\partial^{2} f(x)}{\partial x_{2}^{2}}} & {\cdots} & {\frac{\partial^{2} f(x)}{\partial x_{2} \partial x_{n}}} \\ {\vdots} & {\vdots} & {\ddots} & {\vdots} \\ {\frac{\partial^{2} f(x)}{\partial x_{n} \partial x_{1}}} & {\frac{\partial^{2} f(x)}{\partial x_{n} \partial x_{2}}} & {\cdots} & {\frac{\partial^{2} f(x)}{\partial x_{n}^{2}}}\end{array}\right]
$$

换句话说，$\nabla_{x}^{2} f(x) \in \mathbb{R}^{n \times n}$，其：

$$
\left(\nabla_{x}^{2} f(x)\right)_{i j}=\frac{\partial^{2} f(x)}{\partial x_{i} \partial x_{j}}
$$

注意：黑塞矩阵通常是对称阵：

$$
\frac{\partial^{2} f(x)}{\partial x_{i} \partial x_{j}}=\frac{\partial^{2} f(x)}{\partial x_{j} \partial x_{i}}
$$

与梯度相似，只有当$f(x)$为实值时黑塞矩阵才有定义。

很自然地认为梯度与向量函数的一阶导数的相似，而黑塞矩阵与二阶导数的相似（我们使用的符号也暗示了这种关系）。 这种直觉通常是正确的，但需要记住以下几个注意事项。
首先，对于一个变量$f: \mathbb{R} \rightarrow \mathbb{R}$的实值函数，它的基本定义：二阶导数是一阶导数的导数，即：

$$
\frac{\partial^{2} f(x)}{\partial x^{2}}=\frac{\partial}{\partial x} \frac{\partial}{\partial x} f(x)
$$

然而，对于向量的函数，函数的梯度是一个向量，我们不能取向量的梯度，即:

$$
\nabla_{x} \nabla_{x} f(x)=\nabla_{x}\left[\begin{array}{c}{\frac{\partial f(x)}{\partial x_{1}}} \\ {\frac{\partial f(x)}{\partial x_{2}}} \\ {\vdots} \\ {\frac{\partial f(x)}{\partial x_{n}}}\end{array}\right]
$$

上面这个表达式没有意义。 因此，黑塞矩阵不是梯度的梯度。 然而，下面这种情况却这几乎是正确的：如果我们看一下梯度$\left(\nabla_{x} f(x)\right)_{i}=\partial f(x) / \partial x_{i}$的第$i$个元素，并取关于于$x$的梯度我们得到：

$$
\nabla_{x} \frac{\partial f(x)}{\partial x_{i}}=\left[\begin{array}{c}{\frac{\partial^{2} f(x)}{\partial x_{i} \partial x_{1}}} \\ {\frac{\partial^{2} f(x)}{\partial x_{2} \partial x_{2}}} \\ {\vdots} \\ {\frac{\partial f(x)}{\partial x_{i} \partial x_{n}}}\end{array}\right]
$$

这是黑塞矩阵第$i$行（列）,所以：

$$
\nabla_{x}^{2} f(x)=\left[\nabla_{x}\left(\nabla_{x} f(x)\right)_{1} \quad \nabla_{x}\left(\nabla_{x} f(x)\right)_{2} \quad \cdots \quad \nabla_{x}\left(\nabla_{x} f(x)\right)_{n}\right]
$$

简单地说：我们可以说由于：$\nabla_{x}^{2} f(x)=\nabla_{x}\left(\nabla_{x} f(x)\right)^{T}$，只要我们理解，这实际上是取$\nabla_{x} f(x)$的每个元素的梯度，而不是整个向量的梯度。

最后，请注意，虽然我们可以对矩阵$A\in \mathbb{R}^{n}$取梯度，但对于这门课，我们只考虑对向量$x \in \mathbb{R}^{n}$取黑塞矩阵。
这会方便很多（事实上，我们所做的任何计算都不要求我们找到关于矩阵的黑森方程），因为关于矩阵的黑塞方程就必须对矩阵所有元素求偏导数$\partial^{2} f(A) /\left(\partial A_{i j} \partial A_{k \ell}\right)$，将其表示为矩阵相当麻烦。

### 4.3 二次函数和线性函数的梯度和黑塞矩阵

现在让我们尝试确定几个简单函数的梯度和黑塞矩阵。 应该注意的是，这里给出的所有梯度都是**CS229**讲义中给出的梯度的特殊情况。

对于$x \in \mathbb{R}^{n}$, 设$f(x)=b^Tx$ 的某些已知向量$b \in \mathbb{R}^{n}$ ，则：

$$
f(x)=\sum_{i=1}^{n} b_{i} x_{i}
$$

所以：

$$
\frac{\partial f(x)}{\partial x_{k}}=\frac{\partial}{\partial x_{k}} \sum_{i=1}^{n} b_{i} x_{i}=b_{k}
$$

由此我们可以很容易地看出$\nabla_{x} b^{T} x=b$。 这应该与单变量微积分中的类似情况进行比较，其中$\partial /(\partial x) a x=a$。
现在考虑$A\in \mathbb{S}^{n}$的二次函数$f(x)=x^TAx$。 记住这一点：

$$
f(x)=\sum_{i=1}^{n} \sum_{j=1}^{n} A_{i j} x_{i} x_{j}
$$

为了取偏导数，我们将分别考虑包括$x_k$和$x_2^k$因子的项：

$$
\begin{aligned} \frac{\partial f(x)}{\partial x_{k}} &=\frac{\partial}{\partial x_{k}} \sum_{i=1}^{n} \sum_{j=1}^{n} A_{i j} x_{i} x_{j} \\ &=\frac{\partial}{\partial x_{k}}\left[\sum_{i \neq k} \sum_{j \neq k} A_{i j} x_{i} x_{j}+\sum_{i \neq k} A_{i k} x_{i} x_{k}+\sum_{j \neq k} A_{k j} x_{k} x_{j}+A_{k k} x_{k}^{2}\right] \\ &=\sum_{i \neq k} A_{i k} x_{i}+\sum_{j \neq k} A_{k j} x_{j}+2 A_{k k} x_{k} \\ &=\sum_{i=1}^{n} A_{i k} x_{i}+\sum_{j=1}^{n} A_{k j} x_{j}=2 \sum_{i=1}^{n} A_{k i} x_{i} \end{aligned}
$$

最后一个等式，是因为$A$是对称的（我们可以安全地假设，因为它以二次形式出现）。 注意，$\nabla_{x} f(x)$的第$k$个元素是$A$和$x$的第$k$行的内积。 因此，$\nabla_{x} x^{T} A x=2 A x$。 同样，这应该提醒你单变量微积分中的类似事实，即$\partial /(\partial x) a x^{2}=2 a x$。

最后，让我们来看看二次函数$f(x)=x^TAx$黑塞矩阵（显然，线性函数$b^Tx$的黑塞矩阵为零）。在这种情况下:

$$
\frac{\partial^{2} f(x)}{\partial x_{k} \partial x_{\ell}}=\frac{\partial}{\partial x_{k}}\left[\frac{\partial f(x)}{\partial x_{\ell}}\right]=\frac{\partial}{\partial x_{k}}\left[2 \sum_{i=1}^{n} A_{\ell i} x_{i}\right]=2 A_{\ell k}=2 A_{k \ell}
$$

因此，应该很清楚$\nabla_{x}^2 x^{T} A x=2 A$，这应该是完全可以理解的（同样类似于$\partial^2 /(\partial x^2) a x^{2}=2a$的单变量事实）。

简要概括起来：

- $\nabla_{x} b^{T} x=b$

- $\nabla_{x} x^{T} A x=2 A x$ (如果$A$是对称阵)

- $\nabla_{x}^2 x^{T} A x=2 A $  (如果$A$是对称阵)

### 4.4 最小二乘法

让我们应用上一节中得到的方程来推导最小二乘方程。假设我们得到矩阵$A\in \mathbb{R}^{m \times n}$（为了简单起见，我们假设$A$是满秩）和向量$b\in \mathbb{R}^{m}$，从而使$b \notin \mathcal{R}(A)$。在这种情况下，我们将无法找到向量$x\in \mathbb{R}^{n}$，由于$Ax = b$，因此我们想要找到一个向量$x$，使得$Ax$尽可能接近 $b$，用欧几里德范数的平方$\|A x-b\|\_{2}^{2} $来衡量。

使用公式$\|x\|^{2}=x^Tx$，我们可以得到：

$$
\begin{aligned}\|A x-b\|_{2}^{2} &=(A x-b)^{T}(A x-b) \\ &=x^{T} A^{T} A x-2 b^{T} A x+b^{T} b \end{aligned}
$$

根据$x$的梯度，并利用上一节中推导的性质：

$$
\begin{aligned} \nabla_{x}\left(x^{T} A^{T} A x-2 b^{T} A x+b^{T} b\right) &=\nabla_{x} x^{T} A^{T} A x-\nabla_{x} 2 b^{T} A x+\nabla_{x} b^{T} b \\ &=2 A^{T} A x-2 A^{T} b \end{aligned}
$$

将最后一个表达式设置为零，然后解出$x$，得到了正规方程：

$$
x = (A^TA)^{-1}A^Tb
$$

这和我们在课堂上得到的相同。

### 4.5 行列式的梯度

现在让我们考虑一种情况，我们找到一个函数相对于矩阵的梯度，也就是说，对于$A\in \mathbb{R}^{n \times n}$，我们要找到$\nabla_{A}|A|$。回想一下我们对行列式的讨论：

$$
|A|=\sum_{i=1}^{n}(-1)^{i+j} A_{i j}\left|A_{\backslash i, \backslash j}\right| \quad(\text { for any } j \in 1, \ldots, n)
$$

所以：

$$
\frac{\partial}{\partial A_{k \ell}}|A|=\frac{\partial}{\partial A_{k \ell}} \sum_{i=1}^{n}(-1)^{i+j} A_{i j}\left|A_{\backslash i, \backslash j}\right|=(-1)^{k+\ell}\left|A_{\backslash k,\backslash \ell}\right|=(\operatorname{adj}(A))_{\ell k}
$$

从这里可以知道，它直接从伴随矩阵的性质得出：

$$
\nabla_{A}|A|=(\operatorname{adj}(A))^{T}=|A| A^{-T}
$$

现在我们来考虑函数$f : \mathbb{S}_{++}^{n} \rightarrow \mathbb{R}$，$f(A)=\log |A|$。注意，我们必须将$f$的域限制为正定矩阵，因为这确保了$|A|>0$，因此$|A|$的对数是实数。在这种情况下，我们可以使用链式法则（没什么奇怪的，只是单变量演算中的普通链式法则）来看看：

$$
\frac{\partial \log |A|}{\partial A_{i j}}=\frac{\partial \log |A|}{\partial|A|} \frac{\partial|A|}{\partial A_{i j}}=\frac{1}{|A|} \frac{\partial|A|}{\partial A_{i j}}
$$

从这一点可以明显看出：

$$
\nabla_{A} \log |A|=\frac{1}{|A|} \nabla_{A}|A|=A^{-1}
$$

我们可以在最后一个表达式中删除转置，因为$A$是对称的。注意与单值情况的相似性，其中$\partial /(\partial x) \log x=1 / x$。

### 4.6 特征值优化

最后，我们使用矩阵演算以直接导致特征值/特征向量分析的方式求解优化问题。 考虑以下等式约束优化问题：

$$
\max _{x \in \mathbb{R}^{n}} x^{T} A x \quad \text { subject to }\|x\|_{2}^{2}=1
$$

对于对称矩阵$A\in \mathbb{S}^{n}$。求解等式约束优化问题的标准方法是采用**拉格朗日**<span id="lagrangian"></span>形式，一种包含等式约束的目标函数[^14]，在这种情况下，拉格朗日函数可由以下公式给出：

$$
\mathcal{L}(x, \lambda)=x^{T} A x-\lambda x^{T} x
$$

其中，$\lambda $被称为与等式约束关联的拉格朗日乘子。可以确定，要使$x^{\star}$成为问题的最佳点，拉格朗日的梯度必须在$x^\star$处为零（这不是唯一的条件，但它是必需的）。也就是说，

$$
\nabla_{x} \mathcal{L}(x, \lambda)=\nabla_{x}\left(x^{T} A x-\lambda x^{T} x\right)=2 A^{T} x-2 \lambda x=0
$$

请注意，这只是线性方程$Ax =\lambda x$。 这表明假设$x^T x = 1$，可能最大化（或最小化）$x^T Ax$的唯一点是$A$的特征向量。

[^14]: 如果您以前没有见过拉格朗日，请不要担心，因为我们将在后面的 CS229 中更详细地介绍它们。

## 名词索引

[column vector 列向量](#column_vector)
[row vector 行向量](#row_vector)
[inner product 内积](#inner_product)
[dot product 点积](#dot_product)
[outer product 外积](#outer_product)
[linear combination 线性组合](#linear_combination)
[identity matrix 单位矩阵](#identity_matrix)
[diagonal matrix 对角矩阵](#diagonal_matrix)
[transpose 转置](#transpose)
[symmetric matrix 对称矩阵](#symmetric)
[anti-symmetric matrix 反对称矩阵](#anti-symmetric)
[trace 迹](#trace)
[norm 范数](#norm)
[(linearly) independent 线性无关](#linearly_independent)
[(linearly) dependent 线性相关](#linearly_dependent)
[column rank 列秩](#column_rank)
[row rank 行秩](#row_rank)
[rank 秩](#rank)
[full rank 满秩](#full_rank)
[inverse 逆](#inverse)
[invertible 可逆的](#invertible)
[non-singular 非奇异](#non-singular)
[non-invertible 不可逆](#non-invertible)
[singular 奇异](#singular)
[orthogonal 正交](#orthogonal)
[normalized 归一化](#normalized)
[span 张成](#span)
[projection 投影](#projection)
[range 值域](#range)
[columnspace 列空间](#columnspace)
[nullspace 零空间](#nullspace)
[orthogonal complements 正交补](#orthogonal_complements)
[determinant 行列式](#determinant)
[classical adjoint(adjugate) matrix 经典伴随矩阵](#classical_adjoint)
[adjoint(adjugate) matrix 伴随矩阵](#adjoint)
[minor 余子式](#minor)
[cofactor 代数余子式](#cofactor)
[cofactor matrix 代数余子式矩阵](#cofactor_matrix)
[quadratic form 二次型](#quadratic_form)
[positive definite(PD) 正定](#positive_definite)
[positive semidefinitee (PSD) 半正定](#positive_semidefinite)
[negative definite (ND) 负定](#negative_definite)
[negative semidefinite(NSD) 半负定](#negative_semidefinite)
[indefinite 不定](#indefinite)
[Gram matrix 格拉姆矩阵](#gram_matrix)
[eigenvalue 特征值](#eigenvalue)
[eigenvector 特征向量](#eigenvector)
[Diagonalizing 对角化](#diagonalizing)
[gradient 梯度](#gradient)
[Hessian 黑塞矩阵](#hessian)
[Lagrangian 拉格朗日](#lagrangian)
