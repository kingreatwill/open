
### Agent

gastown
https://github.com/gastownhall/gastown


维度 | 语言 | 持久化 | Agent 类型 | 通信 | 规模 | 开源协议
|---|---|---|---|---|---|---|
gastown | Go | git | 多 | 邮件 | 20-30 | MIT
CrewAI | python | 否 | 统一 | 函数 | 10+ | MIT
AutoGen| python | 否 | 统一 | 函数 | 10+ | MIT
LangGraph| python | 否 | 统一 | 图 | 10+ | MIT




### 爬取网站
#### Moli
为 AI Agent 重新设计的浏览器内核
https://github.com/lexmount/moli

#### firecrawl
https://www.firecrawl.dev/
https://github.com/firecrawl/firecrawl

#### Open Lovable
https://lovable.dev/
https://github.com/firecrawl/open-lovable


### OCR

[OmniDocBench](https://github.com/opendatalab/OmniDocBench) 是一个用于评估多样文档解析能力的综合基准和数据集，覆盖多种文档类型和复杂布局，提供多维度评测指标。



[olmOCR-bench](https://github.com/allenai/olmocr/tree/main/olmocr/bench)是由AI2发布的开源OCR基准测试数据集，包含1403个PDF文档和7010个细粒度单元测试用例，专用于评估VLM驱动的PDF到Markdown/HTML转换能力。 其评测覆盖文本存在性、文本缺席性、自然阅读顺序、表格准确性及数学公式准确性五大维度，聚焦真实场景难点如多栏排版、表格结构还原与公式符号布局，显著超越传统字符级准确率评测范式。



[DocParsingBench](https://github.com/SoMarkAI/DocParsingBench) 是一款 专为智能文档解析产品设计的 测评工具。我们开放了测评数据集、完整测评方法和分析工具，帮助开发者用统一、可复现、可诊断的方式比较不同文档解析模型。


[PureDocBench](https://github.com/zhihengli-casia/puredocbench/) 是一个面向 OCR 和文档解析的源可追踪基准测试数据集，用于评估文档解析系统在干净、数字退化及真实退化三种场景下的表现。 数据由 HTML/CSS 源文件渲染生成图像，并从同一源结构中抽取真实标注（GT），有效减少了标注噪声。 图像轨道：包含三条图像轨道——干净（clean）、数字退化（digital-degraded）和真实退化（real-degraded）。


[Wild-OmniDocBench](https://github.com/VirtualLUOUCAS/Wild_OmniDocBench/) 是一个用于评估文档解析模型在真实拍摄条件下鲁棒性的基准测试。 它基于 OmniDocBench，通过受控的物理模拟将扫描/数字文档转换为真实拍摄图像，包括打印、物理变形以及在多种光照条件下拍摄。 与依赖清晰扫描或数字原生页面的标准基准不同，Wild-OmniDocBench 引入了真实场景下的各类干扰：
几何畸变（透视变换、弯曲、褶皱）
光照变化（定向光、不均匀光、低光环境）
屏幕翻拍伪影（摩尔纹、反光）
环境干扰（背景叠加、阴影）

#### OXR
本地启动

操作系统：Linux
GPU：NVIDIA（支持 BF16）
驱动：NVIDIA 驱动（兼容 CUDA 13.0）
Python：3.12（模型下载／源码启动）
起步显存：单卡 24GB（低并发试用）
推荐显存：单卡 ≥48GB（长文档／较高并发）

```
git clone https://github.com/SoMarkAI/OXR.git
cd OXR

# pip install -U huggingface_hub
hf download SoMarkAI/OXR-1.0 --local-dir ./models/OXR-1.0

# or using ModelScope
# pip install -U modelscope-hub
modelscope download SoMark/OXR-1.0 --local_dir ./models/OXR-1.0
```

#### Xiaomi-OCR
Xiaomi-OCR-0 是一个统一的 0.8B 视觉语言模型，面向文档解析与 OCR 中心理解。模型以 Qwen3.5-0.8B 为起点，在约 1.7 亿条 OCR 中心样本上训练，采用 Q-Mask 文本锚定、继续预训练（CPT）和混合任务强化学习（Mix-RL）。

https://github.com/xiaomi-research/Xiaomi-OCR-0/blob/main/README_zh.md