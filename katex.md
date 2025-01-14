

### mathjax

https://www.mathjax.org/
https://github.com/mathjax/MathJax
### KaTeX
https://katex.org/

https://github.com/KaTeX/KaTeX

### Latex数学公式(能正确复制到公众号等平台）:

https://khan.github.io/KaTeX/docs/supported.html
“复制”时会自动把Latex数学公式转换为图片，并自动上传到云图床（如果在“图片”设置了“...,自动上传到云图床”）。
[请参考：Md2All,让公众号完美显示Latex数学公式](https://www.cnblogs.com/garyyan/p/9228994.html)

#### 行内公式：`$...$`

是的，我就是行内公式：$e^{x^2}\neq{e^x}^2$，排得OK吗？

#### 块公式：

$$ e^{x^2}\neq{e^x}^2 $$

来个复杂点的:

$$H(D_2) = -(\frac{2}{4}\ log_2 \frac{2}{4} + \frac{2}{4}\ log_2 \frac{2}{4}) = 1$$

矩阵：

$$\begin{pmatrix}
    1 & a_1 & a_1^2 & \cdots & a_1^n \\
    1 & a_2 & a_2^2 & \cdots & a_2^n \\
    \vdots & \vdots & \vdots & \ddots & \vdots \\
    1 & a_m & a_m^2 & \cdots & a_m^n \\
    \end{pmatrix}$$

对应“一键排版”的css样式关键字为：`.katex`

#### Latex复制到公众号等各平台的特别说明
##### 复杂的行内公式（顶部和底部突出很多那种），转换后，如果显示不完整，请改为块公式
有些比较复杂的行内公式,转换后，可能会出现顶部和底部很突出的部分看不见的情况，把它改成块公式就OK。

##### 公众号报”图片粘贴失败“时，配合云图床完美解决
如果你发现复制到公众号时报**”图片粘贴失败“**，那是因为公众号有个很奇怪的问题：当复制很小很小的本地图片时，就可能会报”图片粘贴失败“，而其它的平台暂时没遇到。
解决的办法是点“图片”图标，设置好图床信息，并选“...,自动上传到云图床”。
[请参考云图床教程:https://www.cnblogs.com/garyyan/p/9181809.html](https://www.cnblogs.com/garyyan/p/9181809.html)

##### 针对“知乎”的解决方法
知乎是一个比较神奇的网站，会把你的所有的样式恢复为默认的，并图片一定是居中的，也不能直接复制本地的图片。
所以你如果想要在知乎上正常显示：
1:只用块公式，或你可以接受行内公式在知乎上显示变成了块公式;
2:设置云图床，参考上面公众号那样设置“图片”->“...,自动上传到云图床”。