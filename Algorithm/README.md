

[手撕LeetCode题目](https://github.com/labuladong/fucking-algorithm)
[LeetCode 题目用动画的形式演示](https://github.com/MisterBooo/LeetCodeAnimation)

[哈弗大学 排序算法动画](https://www.cs.usfca.edu/~galles/visualization/ComparisonSort.html)
[Sorting Algorithms Animations](https://www.toptal.com/developers/sorting-algorithms)

[算法可视化 Algorithm Visualizer](https://github.com/algorithm-visualizer/algorithm-visualizer)
[算法可视化 demo](https://algorithm-visualizer.org/simple-recursive/cellular-automata)
[可视化排序算法](https://github.com/LucasPilla/Sorting-Algorithms-Visualizer)
[hello-algorithm](https://github.com/geekxh/hello-algorithm)

[动画模拟算法](https://github.com/chefyuan/algorithm-base)

[机器学习算法的动画可视化](https://davpinto.github.io/ml-simulations)

[The Algorithms](https://the-algorithms.com/zh_Hans)

[JavaScript Algorithms and Data Structures](https://github.com/trekhleb/javascript-algorithms)

![https://www.bigocheatsheet.com/](https://www.bigocheatsheet.com/img/big-o-cheat-sheet-poster.png)

[可视化数据结构网站](https://www.cs.usfca.edu/~galles/visualization/Algorithms.html)

[《Hello 算法》：动画图解](https://github.com/krahets/hello-algo)
```
#!/bin/bash

# git clone git@github.com:krahets/hello-algo.git

git -C "/root/code/hello-algo" pull
docker build -f /root/code/hello-algo/Dockerfile -t algo:v0.01 /root/code/hello-algo

# docker run -d -p 10003:8000 --name algo --restart always algo:v0.01
```