## 第 15 章 奇异值分解
奇异值分解([Singular Value Decomposition, SVD](https://en.jinzhao.wiki/wiki/Singular_value_decomposition))是在机器学习领域广泛应用的算法，它不光可以用于降维算法中的特征分解，还可以用于推荐系统，以及自然语言处理等领域。是很多机器学习算法的基石。也是矩阵分解（[Matrix decomposition](https://en.jinzhao.wiki/wiki/Category:Matrix_decompositions)）的一种。

先了解下特征值分解（[Eigenvalues and eigenvectors](https://en.jinzhao.wiki/wiki/Eigenvalues_and_eigenvectors)以及[Eigendecomposition of a matrix](https://en.jinzhao.wiki/wiki/Eigendecomposition_of_a_matrix)）以及对角化（[Diagonalizable matrix](https://en.jinzhao.wiki/wiki/Diagonalizable_matrix)）
特征值（有些方阵是没有特征值的）：
$${\displaystyle A\mathbf {u} =\lambda \mathbf {u} \implies (A-I\lambda)\mathbf {u} = 0}$$
特征值分解：
设$A_{n\times n}$，是具有n个线性无关的特征向量$q_i$（不一定是不同特征值,可以有相同的特征值），那么A可以分解为：
$${\displaystyle A=Q\Lambda Q^{-1}}$$
其中A是方阵，$\Lambda$是由特征值组成的对角矩阵（[diagonal matrix](https://en.jinzhao.wiki/wiki/Diagonal_matrix)）；
$q_i$通常是标准化的，但不是必须的;
因为Q的列是线性无关的，所以 Q 是可逆的;
如果A的特征值都不为0那么A是可逆的（[invertible](https://en.jinzhao.wiki/wiki/Invertible_matrix)）${\displaystyle \mathbf {A} ^{-1}=\mathbf {Q} \mathbf {\Lambda } ^{-1}\mathbf {Q} ^{-1}}$

注意只有[可对角化](https://en.jinzhao.wiki/wiki/Diagonalizable_matrix)的矩阵分解才能用这种方式：如以下矩阵不可被对角化
$${\displaystyle A = \left[{\begin{matrix}1&1\\0&1\end{matrix}}\right]}$$
其特征值为$[1,1]$,特征向量为$[1,0]^T , [-1,0]^T$

如果 $\mathbf {A}$  是对称矩阵（[symmetric matrix](https://en.jinzhao.wiki/wiki/Symmetric_matrix)），因为$\mathbf {Q}$  由特征向量构成 $\mathbf {A}$ 它保证是一个正交矩阵（[orthogonal matrix](https://en.jinzhao.wiki/wiki/Orthogonal_matrix)）,有${\displaystyle \mathbf {Q} ^{-1}=\mathbf {Q} ^{\mathrm {T} }}$

> Q其实也是酉矩阵（[Unitary Matrix](https://en.jinzhao.wiki/wiki/Unitary_matrix)），它是 正交矩阵 在复数上的推广。

> 不可对角化的矩阵称为有缺陷的（[defective](https://en.jinzhao.wiki/wiki/Defective_matrix)）。对于有缺陷的矩阵，特征向量的概念推广到广义特征向量（[generalized eigenvectors](https://en.jinzhao.wiki/wiki/Generalized_eigenvector)），特征值的对角矩阵推广到Jordan 范式（[Jordan normal form](https://en.jinzhao.wiki/wiki/Jordan_normal_form)）。在代数闭域上，任何矩阵A都具有Jordan 范式，因此允许广义特征向量的基和分解为广义特征空间（[generalized eigenspaces](https://en.jinzhao.wiki/wiki/Generalized_eigenspace)）。



- **模型**：
对于复矩阵 $M_{m \times n} = {\displaystyle \mathbf {U\Sigma V^{H}} }$
对于实矩阵 $M_{m \times n} = {\displaystyle \mathbf {U\Sigma V^{T}} }$
其中$U$是${\displaystyle m\times m}$复酉矩阵(正交矩阵)
$\Sigma = diag(\sigma_1,\sigma_2,...,\sigma_p)$是矩形对角矩阵（rectangular diagonal matrix），对角元素是非负的实数并且降序排列,$p=\min(m,n), \sigma_1 \ge \sigma_2 \ge ... \ge \sigma_p \ge 0$
$V$是一个${\displaystyle n\times n}$复酉矩阵(正交矩阵)
$\sigma_i$称为矩阵M的**奇异值**
$U$的列向量称为左奇异向量left-singular vector
$V$的列向量称为右奇异向量right-singular vector
如果矩阵$M$的秩为$r = rank(M),r \le \min(m,n)$
$M$的紧SVD（compact SVD）为$M_{m \times n} = U_{m \times r} \Sigma_{r \times r} V_{n \times r}^T$，$rank(\Sigma_{r \times r}) = rank(M) = r$（可以无损压缩）
$M$的截断SVD（truncated SVD）为$M_{m \times n} \approx U_{m \times k} \Sigma_{k \times k} V_{n \times k}^T , 0 \lt k \lt r$（有损压缩）
矩阵$M$的伪逆（[pseudoinverse](https://en.jinzhao.wiki/wiki/Moore%E2%80%93Penrose_inverse)）$M^{+} = V\Sigma^{+}U^{H}$
$M$的奇异值$\sigma_i$是$M^TM$的特征值的平方根$\sqrt{\lambda_i}$,$V$的列向量是$M^TM$的特征向量，$U$的列向量是$MM^T$的特征向量
$M^TM$很明显还是一个对称矩阵，其特征值为非负，证明：假设$\lambda$是$M^TM$的一个特征值
$\|Mx\|^2 = x^TA^TAx = x^T \lambda x= \lambda x^Tx = \lambda\|x\|^2 \Rightarrow \lambda = \frac{\|Mx\|^2}{\|x\|^2} \ge 0$
$${\begin{aligned}\mathbf {M} ^{T}\mathbf {M} &=\mathbf {V} {\boldsymbol {\Sigma }}^{T}\mathbf {U} ^{T}\,\mathbf {U} {\boldsymbol {\Sigma }}\mathbf {V} ^{T}=\mathbf {V} ({\boldsymbol {\Sigma }}^{T}{\boldsymbol {\Sigma }})\mathbf {V} ^{T}\\\mathbf {M} \mathbf {M} ^{T}&=\mathbf {U} {\boldsymbol {\Sigma }}\mathbf {V} ^{T}\,\mathbf {V} {\boldsymbol {\Sigma }}^{T}\mathbf {U} ^{T}=\mathbf {U} ({\boldsymbol {\Sigma }}{\boldsymbol {\Sigma }}^{T})\mathbf {U} ^{T}\end{aligned}}$$
这不就是矩阵的特征值分解吗【上面有将特征值分解有讲到对称矩阵的对角化（特征值分解）】

- **策略**：
- **算法**：
### 附加知识


#### 矩阵性质
这里介绍一些参见的[矩阵](https://en.jinzhao.wiki/wiki/Category:Matrices)，以及其性质。

##### 共轭转置（[Conjugate transpose](https://en.jinzhao.wiki/wiki/Conjugate_transpose)）

共轭（[Complex conjugate](https://en.jinzhao.wiki/wiki/Complex_conjugate)）是复数上的概念
对于一个复数$z =a + bi$，其共轭为$\bar{z} = a-bi$，所以有$z\bar{z} = a^2 + b^2$
共轭转置也有其它叫法，如：Hermitian conjugate, bedaggered matrix, adjoint matrix or transjugate。值得注意的是adjoint matrix而不是 这个[Adjugate matrix](https://en.jinzhao.wiki/wiki/Adjugate_matrix)，虽然有时候他们都用$A^*$表示。这里为了统一我用$A^H$表示A的共轭转置矩阵。可以参考[共轭转置矩阵与伴随矩阵都用A*表示合理吗？](https://zhuanlan.zhihu.com/p/87330558)
> 有个神奇的公式：欧拉公式 $e^{\pi i}+1=0$

$${\displaystyle \left({\boldsymbol {A}}^{\mathrm {H} }\right)_{ij}={\overline {{\boldsymbol {A}}_{ji}}}}$$

$${\displaystyle {\boldsymbol {A}}^{\mathrm {H} }=\left({\overline {\boldsymbol {A}}}\right)^{\mathsf {T}}={\overline {{\boldsymbol {A}}^{\mathsf {T}}}}}$$

例如：
$${\displaystyle {\boldsymbol {A}}={\begin{bmatrix}1&-2-i&5\\1+i&i&4-2i\end{bmatrix}}} , {\displaystyle {\boldsymbol {A}}^{\mathsf {T}}={\begin{bmatrix}1&1+i\\-2-i&i\\5&4-2i\end{bmatrix}}} , {\displaystyle {\boldsymbol {A}}^{\mathrm {H} }={\begin{bmatrix}1&1-i\\-2+i&-i\\5&4+2i\end{bmatrix}}}$$

性质：
- ${\displaystyle ({\boldsymbol {A}}+{\boldsymbol {B}})^{\mathrm {H} }={\boldsymbol {A}}^{\mathrm {H} }+{\boldsymbol {B }}^{\mathrm {H} }}$ 
- ${\displaystyle (z{\boldsymbol {A}})^{\mathrm {H} }={\overline {z}}{\boldsymbol {A}}^{\mathrm {H} }}$
- ${\displaystyle ({\boldsymbol {A}}{\boldsymbol {B}})^{\mathrm {H} }={\boldsymbol {B}}^{\mathrm {H} }{\boldsymbol {A}} ^{\mathrm {H} }}$
- ${\displaystyle \left({\boldsymbol {A}}^{\mathrm {H} }\right)^{\mathrm {H} }={\boldsymbol {A}}}$
- 如果$A$可逆，当且仅当$A^H$可逆，有${\displaystyle \left({\boldsymbol {A}}^{\mathrm {H} }\right)^{-1}=\left({\boldsymbol {A}}^{-1}\right)^{ \mathrm {H} }}$
- $A^H$的特征值是$A$特征值的复共轭
- 内积性质${\displaystyle \left\langle {\boldsymbol {A}}x,y\right\rangle _{m}=\left\langle x,{\boldsymbol {A}}^{\mathrm {H} }y\right\rangle _{n}}$，`A是m*n,x是n*1,y是m*1`，下标m表示是m维向量作内积。
- （A是方阵）行列式性质${\displaystyle \det \left({\boldsymbol {A}}^{\mathrm {H} }\right)={\overline {\det \left({\boldsymbol {A}}\right)}}}$
- （A是方阵）迹的性质${\displaystyle \operatorname {tr} \left({\boldsymbol {A}}^{\mathrm {H} }\right)={\overline {\operatorname {tr} ({\boldsymbol {A}})}}}$

##### 埃尔米特矩阵（[Hermitian matrix](https://en.jinzhao.wiki/wiki/Hermitian_matrix)）
Hermitian matrix (or self-adjoint matrix)
A是复**方阵**
$${\displaystyle A{\text{ Hermitian}}\quad \iff \quad A=A^{\mathsf {H}}}$$
例子：
$$A = {\displaystyle {\begin{bmatrix}0&a-ib&c-id\\a+ib&1&m-in\\c+id&m+in&2\end{bmatrix}}}$$
**埃尔米特矩阵是对称矩阵在复数上的推广**。

其矩阵有很多性质，具体见维基百科。

Skew-Hermitian matrix：${\displaystyle A{\text{ skew-Hermitian}}\quad \iff \quad A^{\mathsf {H}}=-A}$

##### [Normal matrix](https://en.jinzhao.wiki/wiki/Normal_matrix)
A是复**方阵**
$${\displaystyle A{\text{ normal}}\quad \iff \quad A^{H}A=AA^{H}}$$

例子：
$${\displaystyle A={\begin{bmatrix}1&1&0\\0&1&1\\1&0&1\end{bmatrix}}} , {\displaystyle AA^{H}={\begin{bmatrix}2&1&1\\1&2&1\\1&1&2\end{bmatrix}}=A^{H}A.}$$

Normal matrix一定是可对角化的$A = U\Lambda U^H$，$U$是酉矩阵，$\Lambda = diag(\lambda_1,...)$是$A$的特征值组成的对角矩阵

> 对于复矩阵，所有的unitary, Hermitian, and skew-Hermitian 矩阵都是normal矩阵
> 对应的对于实矩阵，所有的 orthogonal, symmetric, and skew-symmetric 矩阵也都是normal矩阵

##### 酉矩阵（[Unitary matrix](https://en.jinzhao.wiki/wiki/Unitary_matrix)）
U是复**方阵**
$$U^{H} = U^{-1}$$
性质：
- ${\displaystyle U^{H}U=UU^{H}=I,}$
- $\left\langle Ux,Uy\right\rangle = \left\langle x,y\right\rangle$
- U是可对角化的${\displaystyle U=VDV^{H},}$where V is unitary, and D is diagonal and unitary.
- ${\displaystyle \left|\det(U)\right|=1}$
- 其特征向量是相互正交的（废话，正交矩阵的推广）

**酉矩阵它是正交矩阵在复数上的推广**。

#### 矩阵分解(因子分解)

> sympy.Matrix除了分解还有diagonalize对角化（也是一种矩阵分解），eig特征值（其实也可以特征值分解），rref行简化阶梯型，det行列式，inv逆矩阵，广义逆矩阵pinv；更多[参考](https://docs.sympy.org/latest/modules/matrices/matrices.html#linear-algebra)
> scipy.linalg中也有很多关于线性代数的方法：scipy.linalg.svd，以及各种矩阵分解的方法；更多[参考](http://scipy.github.io/devdocs/reference/linalg.html)
> numpy.linalg中也有很多关于线性代数的方法：np.linalg.svd；更多[参考](https://docs.scipy.org/doc/numpy-1.15.0/reference/routines.linalg.html)


除了SVD和PCA，NMF以外，还有很多矩阵分解（[Matrix decomposition](https://en.jinzhao.wiki/wiki/Matrix_decomposition)）的方法。不过有很多分解是有要求的，比如必须是方阵（特征值分解就要求必须是方阵）等。
- **LU分解**（[LU decomposition](https://en.jinzhao.wiki/wiki/LU_decomposition)）
$${\displaystyle A=LU.}$$
L是下三角矩阵（[lower triangular matrix](https://en.jinzhao.wiki/wiki/Triangular_matrix)）
U是上三角矩阵（[upper triangular matrix](https://en.jinzhao.wiki/wiki/Triangular_matrix)）
有时还会包含一个置换矩阵（[permutation matrix](https://en.jinzhao.wiki/wiki/Permutation_matrix)），它在每行和每列中只有一个1，而在其他地方则为0。
$${\displaystyle A=PLU}$$

- **QR分解**（[QR decomposition](https://en.jinzhao.wiki/wiki/QR_decomposition)）
$$A = QR$$
Q是正交矩阵（[Orthogonal Matrix](https://en.jinzhao.wiki/wiki/Orthogonal_matrix)）；
R是上三角矩阵（[right(upper) triangular matrix](https://en.jinzhao.wiki/wiki/Triangular_matrix)）
> 类似的可以定义QL、RQ 和 LQ，L是下三角矩阵（[left(lower) triangular matrix](https://en.jinzhao.wiki/wiki/Triangular_matrix)）

### 参考文献

[15-1] Leon S J. Linear algebra with applications. Pearson, 2009(中译本：线性代数。张文博，张丽静 译. 北京：机械工业出版社)

[15-2] Strang G. Introduction to linear algebra. Fourth Edition. Wellesley-Cambridge Press, 2009.

[15-3] Cline A K. Dhillon I S. Computation of the singular value decomposition, Handbook of linear algebra. CRC Press, 2006.

[15-4] 徐树方. 矩阵计算的理论与方法。北京：北京大学出版社, 1995.

[15-5] Kolda T G,Bader B W. [Tensor decompositions and applications](https://old-www.sandia.gov/~tgkolda/pubs/pubfiles/SAND2007-6702.pdf). SIAM Review, 2009, 51(3):455-500.

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