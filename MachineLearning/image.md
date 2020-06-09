<!--toc-->
[toc]
# 图像处理需要掌握的基本知识

给图片打标签的工具（不要有中文路径）
https://github.com/tzutalin/labelImg

标注工具labelimg和labelme

矩形标注工具：labelimg
多边形标准工具：labelme
 

前者官网发布了可执行文件，后者只有python源码，如果需要编译windows exe，可以这样：
pip install labelme

然后运行labelme确保程序可以正常执行

下载源码：

cd  D:\github\wkentaro\labelme-3.16.7

pip install .

pip install pyinstaller

pyinstaller labelme.spec



## 静态图像类型
### 矢量（Vector）图
用数学公式描述的图像，用一系列绘图指令表示图像：图像中每个形状都用一个完整的公式描述，称为一个对象。
- 优点：
    - 文件数据量很小：
    - 图像质量与分辨率无关：无论图像放大或缩小多少倍，总是以显示设备允许的最大清晰度显示。
    计算机计算与显示图像时，往往能看到画图的过程。

- 缺点：
    - 不易制作色调丰富或色彩变化太多的图像；
    - 绘出来的图像不是很逼真；
    - 不易在不同的软件间交换文件。

### 位图（Bitmap）
通过像素点表示图像，每个像素具有颜色属性和位置属性。
- 优点：
    - 显示速度快；
    - 真实世界的图像可以通过扫描仪、数码相机、摄像机等设备方便的转化为点位图。
- 缺点：
    - 存储和传输时数据量比较大；
    - 缩放、旋转时算法复杂且容易失真。

1. 线画稿（LineArt）
只有黑白两种颜色。适合于由黑白两色构成而没有灰度阴影的图像
2. 灰度图像（GrayScale）
从技术上说，就是具有从黑到白的若干种灰度的单色图像。若灰度图像像素的灰度级用8bit表示，则每个像素都是介于黑色和白色之间的256（2*=256）种灰度中的一种。通常所说的黑白照片，其实包含了黑白之间的所有灰度色调。

灰度图像是除了黑白之外，还添加了第三种颜色：灰色，灰色把灰度划分为 256 个不同的亮度，例如纯白色，它的亮度级别是255。

图像转化为灰度图像有以下几种算法：

浮点算法：Gray = R * 0.3 + G * 0.59 + B * 0.11
整数方法：Gray = ( R * 30 + G * 59 + B * 11 ) / 100
移位方法：Gray = ( R * 76 + G * 151 + B * 28 ) >> 8
平均值法：Gray = ( R + G + B ) / 3
仅取绿色：Gray = G
加权平均值算法：R = G = B = R * 0.299 + G * 0.587 + B * 0.144


3. 二值图像 (Binary Image)
二值图像是取值只有0和1的逻辑数组

4. 索引颜色图像（Index Color）
索引颜色通常也称为映射颜色。在这种模式下，颜色是一组预先定义的、有限的颜色。索引颜色的图像最多只能显示256种颜色。索引颜色图像在图像文件里定义索引颜色。打开该文件时，构成该图像具体颜色的索引值就被读入程序里，然后根据索引值找到最终的颜色。

5. 伪彩色图像(Pseudo-color)
我们知道可以观察出颜色的光的波长范围是有限的，仅仅有那么一小段，换句话说也就是说有一大段光，仅仅有一小段有颜色。
其它都是灰度的，但人类视觉有一个特点就是，仅仅能分辨出二十几种灰度，也就是说採集到的灰度图像分辨率超级高。有一千个灰度级，但非常遗憾。
人们仅仅能看出二十几个，也就是说信息损失了五十倍，但人类视觉对彩色的分辨能力相当强，可以分辨出几千种色度。
在从採集的角度说下伪彩和真彩色，伪彩色原始图像是灰度图像

6. 真彩色图像（True Color）
自然界中几乎所有颜色都可以由红、绿、蓝（R，G，B）组合而成。真彩色图像中，每一个像素由红、绿和蓝三个字节组成，每个字节为8bit，表示0到255之间的不同的亮度值。256×256×256，能表示约1670万种颜色。颜色深度为每像素24位的数字图像是目前所能获取、浏览和保存的颜色信息最丰富的彩色图像，由于它所表达的颜色远远超出了人眼所能辨别的范围，故将其称为“真彩色”。

- **常用颜色的RGB值**
白色：rgb(255,255,255)
黑色：rgb(0,0,0)
红色：rgb(255,0,0)
绿色：rgb(0,255,0)
蓝色：rgb(0,0,255)
青色：rgb(0,255,255)
紫色：rgb(255,0,255)
调整相关数字，便可以得到深浅不一的各种颜色。
ARGB：
A = Alpha表示透明度
FF - 不透明
00 - 全透明 
<table border="1" cellspacing="3" cellpadding="0" align="center" bgcolor="#F6ECE6">
<tbody>
<tr>
<td align="center" width="100" height="10">
<p align="center"><span style="color: #b12c2c;"><strong>颜色样式</strong></span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: #b12c2c;"><strong><strong>RGB</strong>数值</strong></span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: #b12c2c;"><strong>颜色代码</strong></span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: #b12c2c;"><strong>颜色样式</strong></span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: #b12c2c;"><strong><strong>RGB</strong>数值</strong></span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: #b12c2c;"><strong>颜色代码</strong></span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#000000" width="100" height="10">
<p align="center"><span style="color: white;">黑色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">0,0,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#000000</span></p>
</td>
<td align="center" bgcolor="#FFFFFF" width="100" height="10">
<p align="center"><span style="color: navy;">白色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,255,255</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FFFFFF</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#666666" width="100" height="10">
<p align="center"><span style="color: white;">象牙黑</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">88,87,86</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#666666</span></p>
</td>
<td align="center" bgcolor="#F0FFFF" width="100" height="10">
<p align="center"><span style="color: navy;">天蓝灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">202,235,216</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#F0FFFF</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#808A87" width="100" height="10">
<p align="center"><span style="color: white;">冷灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">128,138,135</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#808A87</span></p>
</td>
<td align="center" bgcolor="#CCCCCC" width="100" height="10">
<p align="center"><span style="color: navy;">灰色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">192,192,192</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#CCCCCC</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#808069" width="100" height="10">
<p align="center"><span style="color: white;">暖灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">128,118,105</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#808069</span></p>
</td>
<td align="center" bgcolor="#FAFFF0" width="100" height="10">
<p align="center"><span style="color: navy;">象牙灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">251,255,242</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FAFFF0</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#708069" width="100" height="10">
<p align="center"><span style="color: white;">石板灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">118,128,105</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#E6E6E6</span></p>
</td>
<td align="center" bgcolor="#FAF0E6" width="100" height="10">
<p align="center"><span style="color: navy;">亚麻灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">250,240,230</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FAF0E6</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#F5F5F5" width="100" height="10">
<p align="center"><span style="color: navy;">白烟灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">245,245,245</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#F5F5F5</span></p>
</td>
<td align="center" bgcolor="#FFFFCD" width="100" height="10">
<p align="center"><span style="color: navy;">杏仁灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,235,205</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FFFFCD</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#FCE6C9" width="100" height="10">
<p align="center"><span style="color: navy;">蛋壳灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">252,230,202</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FCE6C9</span></p>
</td>
<td align="center" bgcolor="#FFF5EE" width="100" height="10">
<p align="center"><span style="color: navy;">贝壳灰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,245,238</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FFF5EE</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#FF0000" width="100" height="10">
<p align="center"><span style="color: white;">红色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,0,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FF0000</span></p>
</td>
<td align="center" bgcolor="#FFFF00" width="100" height="10">
<p align="center"><span style="color: navy;">黄色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,255,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FFFF00</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#E3170D" width="100" height="10">
<p align="center"><span style="color: white;">镉红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">227,23,13</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#E3170D</span></p>
</td>
<td align="center" bgcolor="#FF9912" width="100" height="10">
<p align="center"><span style="color: navy;">镉黄</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,153,18</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FF9912</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#9C661F" width="100" height="10">
<p align="center"><span style="color: white;">砖红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">156,102,31</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#9C661F</span></p>
</td>
<td align="center" bgcolor="#E3CF57" width="100" height="10">
<p align="center"><span style="color: navy;">香蕉黄</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">227,207,87</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#E3CF57</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#FF7F50" width="100" height="10">
<p align="center"><span style="color: white;">珊瑚红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,127,80</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FF7F50</span></p>
</td>
<td align="center" bgcolor="#FFD700" width="100" height="10">
<p align="center"><span style="color: navy;">金黄</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,215,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FFD700</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#FF6347" width="100" height="10">
<p align="center"><span style="color: white;">番茄红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,99,71</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FF6347</span></p>
</td>
<td align="center" bgcolor="#FF7D40" width="100" height="10">
<p align="center"><span style="color: navy;">肉黄</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,125,64</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FF7D40</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#FFC0CB" width="100" height="10">
<p align="center"><span style="color: white;">粉红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,192,203</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FFC0CB</span></p>
</td>
<td align="center" bgcolor="#FFE384" width="100" height="10">
<p align="center"><span style="color: navy;">粉黄</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,227,132</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FFE384</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#B0171F" width="100" height="10">
<p align="center"><span style="color: white;">印度红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">176,23,31</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#B0171F</span></p>
</td>
<td align="center" bgcolor="#FF8000" width="100" height="10">
<p align="center"><span style="color: white;">橘黄</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,128,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FF8000</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#FF00FF" width="100" height="10">
<p align="center"><span style="color: white;">深红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">255,0,255</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#FF00FF</span></p>
</td>
<td align="center" bgcolor="#ED9121" width="100" height="10">
<p align="center"><span style="color: white;">萝卜黄</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">237,145,33</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#ED9121</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#990033" width="100" height="10">
<p align="center"><span style="color: white;">黑红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">116,0,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#990033</span></p>
</td>
<td align="center" bgcolor="#8B864E" width="100" height="10">
<p align="center"><span style="color: white;">黑黄</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">85,102,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#8B864E</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#00FF00" width="100" height="10">
<p align="center"><span style="color: navy;">绿色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">0,255,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#00FF00</span></p>
</td>
<td align="center" bgcolor="#802A2A" width="100" height="10">
<p align="center"><span style="color: white;">棕色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">128,42,42</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#802A2A</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#00FFFF" width="100" height="10">
<p align="center"><span style="color: navy;"><strong>青色</strong></span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">0,255,255</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#00FFFF</span></p>
</td>
<td align="center" bgcolor="#C76114" width="100" height="10">
<p align="center"><span style="color: white;">土色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">199,97,20</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#C76114</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#7FFF00" width="100" height="10">
<p align="center"><span style="color: navy;">黄绿色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">127,255,0</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#7FFF00</span></p>
</td>
<td align="center" bgcolor="#F4A460" width="100" height="10">
<p align="center"><span style="color: white;">沙棕色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">244,164,95</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#F4A460</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#40E0D0" width="100" height="10">
<p align="center"><span style="color: navy;">青绿色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">64,224,205</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#40E0D0</span></p>
</td>
<td align="center" bgcolor="#D2B48C" width="100" height="10">
<p align="center"><span style="color: navy;">棕褐色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">210,180,140</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#D2B48C</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#082E54" width="100" height="10">
<p align="center"><span style="color: white;">靛<strong>青色</strong></span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">8,46,84</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#082E54</span></p>
</td>
<td align="center" bgcolor="#BC8F8F" width="100" height="10">
<p align="center"><span style="color: white;">玫瑰红</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">188,143,143</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#BC8F8F</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#228B22" width="100" height="10">
<p align="center"><span style="color: white;">森林绿</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">34,139,34</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#228B22</span></p>
</td>
<td align="center" bgcolor="#A0522D" width="100" height="10">
<p align="center"><span style="color: white;">赫色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">160,82,45</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#A0522D</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#6B8E23" width="100" height="10">
<p align="center"><span style="color: white;">草绿色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">107,142,35</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#6B8E23</span></p>
</td>
<td align="center" bgcolor="#C76114" width="100" height="10">
<p align="center"><span style="color: white;">肖贡土色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">199,97,20</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#C76114</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#0000FF" width="100" height="10">
<p align="center"><span style="color: white;">蓝色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">0,0,255</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#0000FF</span></p>
</td>
<td align="center" bgcolor="#A020F0" width="100" height="10">
<p align="center"><span style="color: white;">肖贡土色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">160,32,240</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#A020F0</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#03A89E" width="100" height="10">
<p align="center"><span style="color: white;">锰蓝</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">3,168,158</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#03A89E</span></p>
</td>
<td align="center" bgcolor="#DA70D6" width="100" height="10">
<p align="center"><span style="color: white;">淡紫色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">218,112,214</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#DA70D6</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#191970" width="100" height="10">
<p align="center"><span style="color: white;">深蓝</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">25,25,112</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#191970</span></p>
</td>
<td align="center" bgcolor="#8A2BE2" width="100" height="10">
<p align="center"><span style="color: white;">紫罗兰</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">138,43,226</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#8A2BE2</span></p>
</td>
</tr>
<tr>
<td align="center" bgcolor="#00C78C" width="100" height="10">
<p align="center"><span style="color: white;">土耳其蓝</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">0,199,140</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#00C78C</span></p>
</td>
<td align="center" bgcolor="#9933FA" width="100" height="10">
<p align="center"><span style="color: white;">胡紫色</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">153,51,250</span></p>
</td>
<td align="center" width="100" height="10">
<p align="center"><span style="color: navy;">#9933FA</span></p>
</td>
</tr>
</tbody>
</table>

RGB按1：1：1取值的时候是一个灰度图像

### 图像文件格式
图像文件的格式，即图像文件的数据构成。一般每种图像文件均有一个文件头，在文件头之后是图像数据。
- 文件头：
一般包括文件类型、文件制作者、制作时间、版本号、文件大小等内容。
内容由制作该图像文件的公司决定。
- 图像数据：
各种图像文件的制作还涉及到图像文件的压缩方式和存储效率等。
数字图像有多种存储格式，每种格式一般由不同的开发商支持。随着信息技术的发展和图像应用领域的不断拓宽，还会出现新的图像格式。
- 图像文件格式体系
    - 互联网用：GIF、JPG、PNG
    - 印刷用：TIF、JPG、TAG、PCX
    - 国际标准：TIF、JPG
    - bmp(bitmap)文件:Windows 位图可以用任何颜色深度（从黑白到 24 位颜色）存储单个光栅图像。Windows 位图文件格式与其他 Microsoft Windows 程序兼容。它不支持文件压缩，也不适用于 Web 页。优点：BMP 支持 1 位到 24 位颜色深度。

### 色彩深度（Color Depth，简称色深）/位分辨率（Bit Resolution）
色彩深度（Color Depth，简称色深），是用 位（bit）数来表示数码影像色彩数目的单位。
色彩深度越高，可用的颜色就越多，颜色之间的过渡更自然和平滑。
像素深度越深，所占用的存储空间越大。
相反，如果像素深度太浅，那也影响图像的质量，图像看起来让人觉得很粗糙和很不自然。

指每个像素所用的位数（bit），像素位决定了彩色图像的每个像素可能的颜色数，或者确定灰度图像每个像素可能有的灰度级数。例如，一个彩色图像的每个像素用R,G，B,三个分量来表示，如每个分量用8位，那么一个像素用24位表示，就说这个像素的深度是24

位深度计算是以2为底数的指数的幂。常见的有：
- 1位：2种颜色，单色光，黑白二色，用于compact Macintoshes。
- 2位：4种颜色，CGA，用于gray-scale早期的NeXTstation及color Macintoshes。
- 3位：8种颜色，用于大部分早期的电脑显示器。
- 4位：16种颜色，用于EGA及不常见及在更高的分辨率的VGA标准，color Macintoshes。
- 5位：32种颜色，用于Original Amiga chipset。
- 6位：64种颜色，用于Original Amiga chipset。
- 7位：128种颜色
- 8位：256种颜色，用于最早期的彩色Unix工作站，低分辨率的VGA，Super VGA，AGA，color Macintoshes。
灰阶，有256种灰色（包括黑白）。若以24位模式来表示，则RGB的数值均一样，例如(200,200,200)。
彩色图像，若以24位模式来表示，则RGB的数值均一样，例如(200,200,200)。就是常说的24位真彩，约为1670万色。
- 9位：512种颜色
- 10位:1024种颜色，
- 12位：用于部分硅谷图形系统，Neo Geo，彩色NeXTstation及Amiga系统于HAM mode。
- 16位：用于部分color Macintoshes( 红色占5 个位、蓝色占 5 个位、绿色占 6 个位，所以红色、蓝色、绿色各有 32、32、64 种明暗度的变化总共可以组合出 64K 种颜色 )。
- 24位：有16,777,216色，真彩色，能提供比肉眼能识别更多的颜色，用于拍摄照片。
- 32位：基于24位而生，增加8个位的Alpha通道。



## 图像像素操作
### 几何运算-加减乘除
#### 图像加法
图像加法有两种方式，一种是通过 Numpy 直接对两个图像进行相加，另一种是通过 OpenCV 的 add() 函数进行相加。

不管使用哪种方法，相加的两个图像必须具有相同的深度和类型，简单理解就是图像的大小和类型必须一致。

- Numpy 加法

Numpy 的运算方法是： img = img1 + img2 ，然后再对最终的运算结果取模。

当最终的像素值 <= 255 时，则运算结果直接为 img1 + img2 。
当最终的像素值 > 255 时，则运算的结果需对 255 进行取模运算(% 256)。

- OpenCV 加法

OpenCV 的运算方式是直接调用 add() 函数进行的，这时的运算方式是饱和运算。

当最终的像素值 <= 255 时，则运算结果直接为 img1 + img2 。
当最终的像素值 > 255 时，这时则是饱和运算，结果固定为 255 。

### 逻辑运算-与或非取反

### 像素读写

### 通道混合与调整

### 对比度与亮度调整

## 图像变换
### 插值(zoom in或out)

### 旋转(rotate)

### 透视变换

### 错切变换

### 翻转

## 像素统计
### 计算均值与方差

### 计算直方图

### 计算最大最小

### 计算像素内方差

## 色彩空间
### RGB

### HSL

### YUV

### YCrCb

### 色彩空间转换

### 灰度转换

### 调整饱和度与亮度

### 主色彩提取与分析

## 卷积图像处理
### 空间域卷积

### 频率域卷积

### FFT空域到时域转换

### 模糊

### 边缘提取
Cancy理论和方法，数学形态学理论，多尺度分辨与分析理论，线性理论，模糊数学理论，分类理论，梯度理论，基于区域的方法，自适应方法，小波理论，混沌免疫模糊聚类法，Hilbeit理论，模板分解和图像积分理论
### 去噪

### 增强

### 直方图均衡化

### 直方图反向投影

## 图像增强
在图像增强领域，运用的主要知识有：模糊数学理论，直方图法，小波理论，多尺度分析理论，非线性理论，HSI彩色模型饱和度分量法，Retinex理论，Contourlet变换等。
在空域中，图像处理的方法主要内容有灰度变换，空域滤波，直方图处理等，基于变换域的图像处理方法是将图像信息特征从时域中转换到频域下，利用一定的指数修正方法修正变换域内的系数，从而获得增强后的图像，与空域方法相比更好的是频域，包括傅里叶变换的算法，小波变换算法，同时这些方法也存在着一定的缺点，其中傅里叶变换容易出现振铃现象，小波分析算法在多分辨率分析中可以有效抑制图像噪声，易于控制图像增强区域，小波分析具有有限的方向，对于提取方向信息不能很好表示。


## 图像滤波
三角形滤波法，Retinex理论，统计理论，小波理论分析，非线性理论，中值理论，偏微分方程法，马尔科夫随机场理论，多尺度分析理论，贝叶斯理论，模拟退火理论
常用的滤波算法有中值滤波、均值滤波、高斯滤波等，中值滤波虽然有避免模糊的优点，但是在图像角点和线条等细节上出现丢失现象。

### 中值滤波(cv.medianBlur)
中值滤波是基于排序统计理论的一种能有效抑制噪声的非线性信号处理技术。它也是一种邻域运算，类似于卷积，但是计算的不是加权求和，而是把数字图像或数字序列中一点的值用该点的一个邻域中各点值的中值代替，让周围像素灰度值的差比较大的像素改取与周围的像素值接近的值，从而可以消除孤立的噪声点。它能减弱或消除傅立叶空间的高频分量，但影响低频分量。因为高频分量对应图像中的区域边缘的灰度值具有较大较快变化的部分，该滤波可将这些分量滤除，使图像平滑。值滤波技术在衰减噪声的同时能较好的保护图像的边缘。



## 形态学处理
### 腐蚀

### 膨胀

### 开闭操作

### 形态学梯度

### 顶帽

### 黑帽

### 内梯度与外梯度

## 图像分割
### K-Means

### Mean-Sift

### 分水岭

### Fuzzy-C Means

### GMM

### Graphic Cut

### 区域生长

## 特征提取
### SIFT

### SURF

### LBP

### HOG

### Haars

### Blob

### DOG或者LOG

### 金字塔

### Haars Corner

### Shi-Tomasi Corner

### Hessian

## 二值图像
### 全局阈值二值化

### 局部阈值二值化

### 轮廓提取

### 区域测量

### 几何矩特性

### 连通区域计算

### 泛洪填充

### 霍夫变换

### 距离变换

### 分水岭分割

### 链式编码

### 骨架提取

### 欧拉数计算

## 对象识别与匹配
### 直方图匹配

### 相关性匹配

### 模板匹配

### KNN

### SVM


## 其它
均值：反映了图像的亮度，均值越大说明图像亮度越大，反之越小；

标准差：反映了图像像素值与均值的离散程度，标准差越大说明图像的质量越好；

平均梯度：反映了图像的清晰度和纹理变化，平均梯度越大说明图像越清晰；

通道：

单通道图，俗称灰度图，每个像素点只能有有一个值表示颜色，它的像素值在0到255之间，0是黑色，255是白色，中间值是一些不同等级的灰色。（也有3通道的灰度图，3通道灰度图只有一个通道有值，其他两个通道的值都是零）。

三通道图，每个像素点都有3个值表示 ，所以就是3通道。也有4通道的图。例如RGB图片即为三通道图片，RGB色彩模式是工业界的一种颜色标准，是通过对红(R)、绿(G)、蓝(B)三个颜色通道的变化以及它们相互之间的叠加来得到各式各样的颜色的，RGB即是代表红、绿、蓝三个通道的颜色，这个标准几乎包括了人类视力所能感知的所有颜色，是目前运用最广的颜色系统之一。总之，每一个点由三个值表示。
