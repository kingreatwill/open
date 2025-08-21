
在生产环境中，我们用 systemd slices 自动化管理 cgroup，并将 cgroup 相关指标接入监控系统，确保配置持续生效。

cpu.cfs_quota_us - CPU 配额限制
cpu.cfs_period_us - CPU 周期时间
cpu.stat - CPU 使用统计
cpu.usage - CPU 使用时间
cpu.shares - CPU 分配权重

### 创建和配置 cgroup

```
#!/bin/bash

# 创建 cgroup 目录
mkdir -p /sys/fs/cgroup/cpu/myapp

# 设置不同的 CPU 权重,# 权重值越大，分配的 CPU 时间越多
echo 1024 > /sys/fs/cgroup/cpu/app1/cpu.shares    # 基础权重
echo 2048 > /sys/fs/cgroup/cpu/myapp/cpu.shares    # 2倍权重
# 设置 CPU 限制为 1.5 个核心
echo 150000 > /sys/fs/cgroup/cpu/myapp/cpu.cfs_quota_us
echo 100000 > /sys/fs/cgroup/cpu/myapp/cpu.cfs_period_us

# 验证设置
echo "CPU 配额: $(cat /sys/fs/cgroup/cpu/myapp/cpu.cfs_quota_us)"
echo "CPU 周期: $(cat /sys/fs/cgroup/cpu/myapp/cpu.cfs_period_us)"
echo "CPU 限制: $(echo "scale=2; $(cat /sys/fs/cgroup/cpu/myapp/cpu.cfs_quota_us) / $(cat /sys/fs/cgroup/cpu/myapp/cpu.cfs_period_us)" | bc) 核心"
```

### 将进程分配到 cgroup
```
#!/bin/bash

# 创建 cgroup
mkdir -p /sys/fs/cgroup/cpu/myapp

# 设置 CPU 限制为 0.5 核心
echo 50000 > /sys/fs/cgroup/cpu/myapp/cpu.cfs_quota_us
echo 100000 > /sys/fs/cgroup/cpu/myapp/cpu.cfs_period_us

# 启动测试进程
stress-ng --cpu 4 --timeout 60 &

# 获取进程 PID
PID=$!

# 将进程分配到 cgroup
echo $PID > /sys/fs/cgroup/cpu/myapp/tasks

echo "进程 $PID 已分配到 cgroup，CPU 限制为 0.5 核心"

# 监控 CPU 使用情况
echo "监控 CPU 使用情况..."
while kill -0 $PID 2>/dev/null; do
    echo "CPU 使用: $(cat /sys/fs/cgroup/cpu/myapp/cpu.stat | grep nr_throttled)"
    sleep 1
done
```
### 动态调整 CPU 限制
```
#!/bin/bash

CGROUP_PATH="/sys/fs/cgroup/cpu/myapp"

# 创建 cgroup
mkdir -p $CGROUP_PATH

# 初始设置：1 个 CPU 核心
echo 100000 > $CGROUP_PATH/cpu.cfs_period_us
echo 100000 > $CGROUP_PATH/cpu.cfs_quota_us

echo "初始 CPU 限制: 1.0 核心"

# 启动测试进程
stress-ng --cpu 4 --timeout 120 &
PID=$!
echo $PID > $CGROUP_PATH/tasks

echo "进程 $PID 已启动"

# 动态调整 CPU 限制
sleep 10
echo "调整 CPU 限制为 0.5 核心"
echo 50000 > $CGROUP_PATH/cpu.cfs_quota_us

sleep 10
echo "调整 CPU 限制为 2.0 核心"
echo 200000 > $CGROUP_PATH/cpu.cfs_quota_us

sleep 10
echo "调整 CPU 限制为 0.25 核心"
echo 25000 > $CGROUP_PATH/cpu.cfs_quota_us

# 等待进程结束
wait $PID
echo "测试完成"
```