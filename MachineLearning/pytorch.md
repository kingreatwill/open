
## Visdom的介绍
Visdom是Facebook专为PyTorch开发的实时可视化工具包，其作用相当于TensorFlow中的Tensorboard，灵活高效且界面美观
https://github.com/facebookresearch/visdom 7.5k
用于创建、组织和共享实时丰富数据可视化的灵活工具。支持Torch 和Numpy。

```
pip install visdom
#开启监听命令
python -m visdom.server # 或者直接visdom
# 然后在浏览器里输入：http://localhost:8097
```