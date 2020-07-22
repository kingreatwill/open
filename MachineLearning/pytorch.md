
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
> 目前from torch.utils.tensorboard import SummaryWriter中的tensorboard是tensorflow开发的，
> 还有这个玩意from tensorboardX import SummaryWriter 


## pytorch安装
```
# https://pytorch.org/
pip install torch===1.5.1 torchvision===0.6.1 -f https://download.pytorch.org/whl/torch_stable.html

or

conda install pytorch torchvision cudatoolkit=10.2 -c pytorch

# 非常慢，我们可以下载安装
https://download.pytorch.org/whl/torch_stable.html

cp37 python 3.7

cpu版本
https://download.pytorch.org/whl/cpu/torch-1.5.1%2Bcpu-cp37-cp37m-win_amd64.whl

gpu版本
https://download.pytorch.org/whl/cu102/torch-1.5.1-cp37-cp37m-win_amd64.whl

还有需要下载torchvision
https://download.pytorch.org/whl/cu102/torchvision-0.6.1-cp37-cp37m-win_amd64.whl

pip install xx.whl
```