# Python程序分析

## Py-Spy
https://github.com/benfred/py-spy
Py-Spy 是一款 Python 应用采样分析器，允许在不重启和修改代码的情况下，可视化你的 Python 项目在哪些地方耗时较久。

Py-Spy 采用 Rust 编写，速度快，不会与要配置的 Python 项目运行相同的进程，也不会以任何方式中断正在运行的应用。

### 安装
pip install py-spy
or
cargo install py-spy

### 分析
#### top
py-spy top --pid 12345
py-spy top -- python myprogram.py

#### dump
py-spy dump --pid 12345

#### record
还支持从运行过程生成火焰图：
py-spy record -o profile.svg --pid 12345
py-spy record -o profile.svg -- python myprogram.py