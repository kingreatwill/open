[TOC]
# Time Series

时间序列分析（[Time Series Analysis](https://en.jinzhao.wiki/wiki/Time_series)）包括分析时间序列数据以提取有意义的统计数据和数据的其他特征的方法。
时间序列预测（[Time series forecasting](https://en.jinzhao.wiki/wiki/Time_series)）是使用模型根据先前观察到的值来预测未来值。


时间序列依据其特征，有以下几种表现形式，并产生与之相适应的分析方法：
1. 长期趋势变化：受某种基本因素的影响，数据依时间变化时表现为一种确定倾向，它按某种规则稳步地增长或下降。使用的分析方法有：移动平均法、指数平滑法、模型拟和法等。
2. 季节性周期变化：受季节更替等因素影响，序列依一固定周期规则性的变化，又称商业循环。采用的方法：季节指数。
3. 循环变化：周期不固定的波动变化。
4. 随机性变化：由许多不确定因素引起的序列变化。

时间序列有三种基本模式：
- 平稳性 / 随机性(Stationarity)：
    1. 期望为常数
    $E(Z_t) = constant$
    2. 方差为常数
    $Var(Z_t) = constant$
    3. 只要K固定，任意两个相隔K个时间段的数据组的协方差相同的(也就是说协方差只依赖于K这个时间跨度，不依赖于时间点t本身)
    $Cov(Z_t,Z_{t+k}) \text{ depends on k}$
- 趋势性(Trend)：
- 季节性(Seasonarity)：

时间序列分析的基本思路：

1. 平稳性检验（单位根检验、自相关图ACF、偏自相关图PACF）
否：非平稳——差分变换——平稳
是：下一步

2. 白噪声检验
否：计算ACF、PACF（如果没有计算这些）
是：——停止，无分析价值

3. 模型识别（根据ACF、PACF等找到具体使用哪一种模型，也就是最优模型）
4. 参数估计
5. 模型检验
6. 模型优化
7. 模型预测

## 参考资料
https://www.jianshu.com/nb/47529965
https://www.jianshu.com/p/68016de16984

[应用时间序列分析备课笔记](https://www.math.pku.edu.cn/teachers/lidf/course/atsa/atsanotes/html/_atsanotes/index.html)

## 相关模型
[Python 中11 种经典时间序列预测方法](https://zhuanlan.zhihu.com/p/407589768)
[11 Classical Time Series Forecasting Methods in Python (Cheat Sheet)](https://machinelearningmastery.com/time-series-forecasting-methods-in-python-cheat-sheet/)

Seasonal Autoregressive Integrated Moving Average
SARIMA模型（Seasonal Autoregressive Integrated Moving Average），季节性差分自回归滑动平均模型，时间序列预测分析方法之一

ARIMA - Autoregressive Integrated Moving Average model
ARIMA模型（英语：Autoregressive Integrated Moving Average model），差分整合移动平均自回归模型，又称整合移动平均自回归模型（移动也可称作滑动），是时间序列预测分析方法之一
对应实现的库

VARIMA - Autoregressive integrated moving average
Prophet - Modeling Multiple Seasonality With Linear or Non-linear Growth
HWAAS - Exponential Smoothing With Additive Trend and Additive Seasonality
HWAMS - Exponential Smoothing with Additive Trend and Multiplicative Seasonality
Hidden Markov Models (HMM)



## 相关开源库

### statsmodels
[statmodels - 6.5k](https://github.com/statsmodels/statsmodels)：是一个Python包，提供了对scipy统计计算的补充，包括描述性统计和统计模型的估计和推断。

> 注重统计分析，[时间序列分析](https://www.statsmodels.org/stable/tsa.html)只是其中一小部分。
### tsfresh
[tsfresh - 5.9k](https://github.com/blue-yonder/tsfresh)：从时间序列中自动提取相关特征

> 注重特征提取和特征选择

### AutoTS
[AutoTS - 200](https://github.com/winedarksea/AutoTS)：自动时间序列预测（Automated Time Series Forecasting)
> 注重自动机器学习
> 支持: statsmodels,prophet[等](https://github.com/winedarksea/AutoTS/blob/master/extended_tutorial.md#optional-packages)

### Prophet
[Prophet - 13.2k](https://github.com/facebook/prophet)：为具有多重季节性的线性或非线性增长的时间序列数据生成高质量预测的工具。

Prophet是一种基于可加模型预测时间序列数据的程序，其中非线性趋势与年度、每周和每日的季节性以及节假日影响相匹配。它最适用于具有强烈季节性影响的时间序列和几个季节的历史数据。Prophet对于丢失的数据和趋势的变化是稳健的，并且通常能很好地处理异常值。

> 适用于具有强烈季节性影响的数据和多个季节的历史数据
> 遵循sklearn模型 API

### Kats
[Kats - 2.8k](https://github.com/facebookresearch/Kats/tree/master/tutorials)：Kats是一个用于分析时间序列数据的工具包，是一个轻量级的、易于使用的、可概括的和可扩展的框架，用于执行时间序列分析，从理解关键统计数据和特征，检测变化点和异常，到预测未来趋势。

> 时间序列分析
> 模式检测，包括季节性、异常值、趋势变化
> 对时间序列数据建立预测模型，包括 Prophet、ARIMA、Holt Winters 等。

### Darts
[Darts - 2.3k](https://github.com/unit8co/darts)：一个用于简单操作和预测时间序列的python库。

> 语法是“sklearn-friendly”

### AtsPy
[AtsPy - 370](https://github.com/firmai/atspy)：Python中的自动时间序列模型
> 自动时间序列模型

### Sktime
[Sktime - 4.4k](https://github.com/alan-turing-institute/sktime)：时间序列机器学习的统一框架

> 深度学习:[sktime-dl](https://github.com/sktime/sktime-dl)
> 遵循sklearn模型 API
