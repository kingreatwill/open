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
- 用$\boldsymbol{x} \in \Bbb{R}^{n}$表示具有$n$个元素的向量。按照惯例，$n$维向量。通常被认为是$n$行$1$列的矩阵，称为**列向量**。如果我们想表示一个**行向量**: 具有 $1$ 行和$n$列的矩阵 - 我们通常写$\boldsymbol{x}^T$（这里$\boldsymbol{x}^T$表示$\boldsymbol{x}$的转置，我们稍后将定义它）。
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

给两个向量$\boldsymbol{x},\boldsymbol{y} \in \Bbb{R}^n$, $\boldsymbol{x}^T \boldsymbol{y}$通常称为向量的**内积**或者**点积**，结果是个实数：

$$
\boldsymbol{x}^T \boldsymbol{y} \in \Bbb{R} =
\begin{bmatrix}x_1 & x_2 & \cdots & x_n\end{bmatrix}
\begin{bmatrix}y_1 \\ y_2 \\ \vdots \\ y_n \end{bmatrix}
= \sum_{i=1}^n{x_iy_i}
$$

请注意，内积实际上只是矩阵乘法的特例。$\boldsymbol{x}^T \boldsymbol{y} = \boldsymbol{y}^T \boldsymbol{x}$ 始终成立。
给定向量$\boldsymbol{x} \in \Bbb{R}^m , \boldsymbol{y} \in \Bbb{R}^n$($m$不一定等于$n$), $\boldsymbol{x} \boldsymbol{y}^T \in \Bbb{R}^{m \times n}$叫向量**外积**，它是一个矩阵，由$(\boldsymbol{x} \boldsymbol{y}^T)_{ij} = x_iy_j$组成，也就是(i.e.)：

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

换句话说，$\boldsymbol{y}$是$A$的列的线性组合，其中线性组合的系数由$\boldsymbol{x}$的元素给出。

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

如果您不熟悉这些属性，请花点时间自己验证它们。 例如，为了检查矩阵乘法的结合性，假设$A \in \mathbb{R}^ {m \times n}，$ $B \in \mathbb{R}^ {n \times p} $，$C \in \mathbb{R}^ {p \times q}$。 注意$AB \in \mathbb{R}^ {m \times p}$，所以$(AB)C \in \mathbb{R}^ {m \times q}$。 类似地，$BC \in \mathbb{R}^ {n \times q}$，所以$A(BC) \in \mathbb{R}^ {m \times q}$。 因此，所得矩阵的维度一致。 为了验证矩阵乘法的结合性，检查$(AB)C$的($i,j$)元素是否等于$A(BC)$的($i,j$)元素。 我们可以使用矩阵乘法的定义直接验证这一点：

$$
\begin{aligned} % aligned &= 等号对齐
((A B) C)_{ij} &= \sum_{k=1}^p{(AB)_{ik}C_{kj}} = \sum_{k=1}^p \Bigg( \sum_{l=1}^n{A_{il}B_{lk}} \Bigg) C_{kj} \\
&=\sum_{k=1}^p \Bigg( \sum_{l=1}^n{A_{il}B_{lk}C_{kj}}\Bigg) = \sum_{l=1}^n \Bigg( \sum_{k=1}^p{A_{il}B_{lk}C_{kj}}\Bigg)\\
&=\sum_{l=1}^nA_{il}\Bigg(\sum_{k=1}^p{B_{lk}C_{kj}}\Bigg)  = \sum_{l=1}^n{A_{il}(BC)_{lj}} = (A(BC))_{ij}
\end{aligned}
$$

这里，第一个和最后两个等式简单地使用了矩阵乘法的定义，第三个和第五个等式使用了标量乘法对加法的分配性质，第四个等式使用了标量加法的交换性和结合性。 这种通过简化为简单标量属性来证明矩阵属性的技术会经常出现，因此请确保您熟悉它。

## 3 操作及其性质

在本节中，我们将介绍矩阵和向量的几种操作和性质。 希望其中的大部分内容都可以帮您复习，此笔记可以作为参考。

### 3.1 单位矩阵和对角矩阵

**单位矩阵**用$I \in \Bbb{R}^{n \times n}$表示，它是一个方阵，对角线的元素是 1，其余元素都是 0。可以这样表示：

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

**对角矩阵**的非对角元素全为 0。对角阵通常表示为：$D=diag(d_1, d_2,\cdots, d_n)$，其中：

$$
D_{ij} =
\begin{cases}
   d_i & i=j \\
   0 & i \neq j
\end{cases}.
$$

很明显，单位矩阵$ I= diag(1, 1, \cdots , 1)$。

### 3.2 转置

矩阵的**转置**是指翻转矩阵的行和列。给定一个矩阵$A \in \Bbb{R}^{m \times n}$，它的转置$A^T \in \Bbb{R}^{n \times m}$ ,其中的元素为：

$$
(A^T)_{ij} = A_{ji}.
$$

事实上，我们在描述行向量时已经使用了转置，因为列向量的转置自然是行向量。

转置有以下性质，且很容易验证：

- $(A^T)^T = A$
- $(AB)^T = B^TA^T$
- $(A+B)^T = A^T + B^T$

### 3.3 对称矩阵

如果$A = A^T$，那么**方阵**$A \in \Bbb{R}^{n \times n}$是**对称**的。

- 元素满足$a_{ij} = a_{ji} , \forall i,j$
- $A = A^T$
- 对于任意方阵$A$，$A + A^T$是对称的
- 对角矩阵都是对称矩阵

如果$A = -A^T$，那么它就是**反对称**的。

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

事实证明，对称矩阵在实践中用到很多，它们有很多很好的属性，我们很快就会看到它们。
通常将大小为$n$的所有对称矩阵的集合表示为$\mathbb{S}^n$，因此$A \in \mathbb{S}^n$意味着$A$是对称的$n\times n$矩阵。

### 3.4 矩阵的迹

方矩阵$A \in \mathbb{R}^ {n \times n}$的**迹**，表示为$\operatorname{tr} (A)$（或者$\operatorname{tr} A$，括号显然是隐含的），是矩阵中对角元素的总和：

$$
\operatorname{tr} A=\sum_{i=1}^{n} A_{i i}
$$

如 CS229 讲义中所述，迹具有以下属性（如下所示）：

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

向量的**范数**$\|x\|$是非正式度量的向量的“长度” 。 例如，我们有常用的欧几里德或$\ell_{2}$范数，

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

一个向量集合$\{ x_1,x_2, \cdots x_n \} \subset \mathbb{R}^m$， 如果没有向量可以表示为其余向量的线性组合，则称称该向量是线性无相关的。 相反，如果属于该组的一个向量可以表示为其余向量的线性组合，则称该向量是**线性相关**的。 也就是说，如果：

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

矩阵$A  \in \mathbb{R}^{m \times n}$的**列秩**是构成线性无关集合的$A$的最大列子集的大小。 由于术语的多样性，这通常简称为$A$的线性无关列的数量。同样，行秩是构成线性无关集合的$A$的最大行数。 对于任何矩阵$A  \in \mathbb{R}^{m \times n}$，事实证明$A$的列秩等于$A$的行秩（尽管我们不会证明这一点），因此两个量统称为$A$的**秩**，用 $\text{rank}(A)$表示。 以下是秩的一些基本属性：

- 对于 $A  \in \mathbb{R}^{m \times n}$，$\text{rank}(A) \leq min(m, n)$，如果$ \text(A) = \text{min} (m, n)$，则： $A$ 被称作**满秩**。
- 对于 $A  \in \mathbb{R}^{m \times n}$， $\text{rank}(A) = \text{rank}(A^T)$
- 对于 $A  \in \mathbb{R}^{m \times n}$,$B  \in \mathbb{R}^{n \times p}$ ,$\text{rank}(AB) \leq \text{min} ( \text{rank}(A), \text{rank}(B))$
- 对于 $A,B \in \mathbb{R}^{m \times n}$，$\text{rank}(A + B) \leq \text{rank}(A) + \text{rank}(B)$

### 3.7 方阵的逆

方阵$A  \in \mathbb{R}^{n \times n}$的**逆**表示为$A^{-1}$，并且是这样的唯一矩阵:

$$
A^{-1}A=I=AA^{-1}
$$

请注意，并非所有矩阵都具有逆。 例如，非方形矩阵根据定义没有逆(存在**伪逆**[^4])。 然而，对于一些方形矩阵$A$，$A^{-1}$也可能不存在。 特别是，如果$A^{-1}$存在，我们说$A$是**可逆**的或**非奇异**的，否则就是**不可逆**或**奇异**的[^5]。

为了使方阵 A 具有逆$A^{-1}$，则$A$必须是满秩。 我们很快就会发现，除了满秩之外，还有许多其它的充分必要条件。
以下是逆的性质; 假设$A,B  \in \mathbb{R}^{n \times n}$，而且是非奇异的：

- $(A^{-1})^{-1} = A$
- $(AB)^{-1} = B^{-1}A^{-1}$
- $(A^{-1})^{T} =(A^{T})^{-1} $因此，该矩阵通常表示为$A^{-T}$。

作为如何使用逆的示例，考虑线性方程组，$Ax = b$，其中$A  \in \mathbb{R}^{n \times n}$，$x,b\in \mathbb{R}$， 如果$A$是非奇异的（即可逆的），那么$x = A^{-1}b$。 （如果$A  \in \mathbb{R}^{m \times n}$，不是方阵，这公式还有用吗？ - **伪逆**[^4]）

[^4]: 参考[Moore–Penrose inverse](https://en.jinzhao.wiki/wiki/Moore%E2%80%93Penrose_inverse)
[^5]: 很容易混淆并认为非奇异意味着不可逆。 但实际上，意思正好相反！ 小心！

### 3.8 正交矩阵

如果 $x^Ty=0$，则两个向量$x,y\in \mathbb{R}^{n}$ 是**正交**的。如果$\|x\|_2=1$，则向量$x\in \mathbb{R}^{n}$ 被**归一化**。如果一个方阵$U\in \mathbb{R}^{n \times n}$的所有列彼此正交并被归一化，则方阵$U$是**正交矩阵**（注意在讨论向量与矩阵时的意义不一样，**两个向量正交只需要内积为0，正交矩阵是各列相互正交并且被归一化**）。

它可以从正交性和正态性的定义中得出:

$$
U^ TU = I = U U^T
$$

换句话说，正交矩阵的逆是其转置。 注意，如果$U$不是方阵， 即，$U\in \mathbb{R}^{m \times n}, n < m $ ，但其列仍然是正交的，则$U^TU = I$，但是$UU^T \neq I$。所以**正交矩阵一定是方阵**。

正交矩阵的另一个好的特性是在具有正交矩阵的向量上操作不会改变其欧几里德范数，即(i.e.):

$$
\|U x\|_{2}=\|x\|_{2}
$$

对于任何 $x\in \mathbb{R}^n$ , $U\in \mathbb{R}^{n \times n}$是正交矩阵。

### 3.9 矩阵的值域和零空间
