# [译] Coding Interview University 一套完整的学习计划

> * 原文地址：[Coding Interview University](https://github.com/jwasham/coding-interview-university)

--- 

## 目录

- [必备知识](#必备知识)
- [算法复杂度 / Big-O / 渐进分析法](#算法复杂度--big-o--渐进分析法)
- [数据结构](#数据结构)
    - [数组（Arrays）](#数组arrays)
    - [链表（Linked Lists）](#链表linked-lists)
    - [堆栈（Stack）](#堆栈stack)
    - [队列（Queue）](#队列queue)
    - [哈希表（Hash table）](#哈希表hash-table)
- [更多的知识](#更多的知识)
    - [二分查找（Binary search）](#二分查找binary-search)
    - [按位运算（Bitwise operations）](#按位运算bitwise-operations)
- [树（Trees）](#树trees)
    - [树 —— 笔记 & 背景](#树--笔记--背景)
    - [二叉查找树（Binary search trees）：BSTs](#二叉查找树binary-search-treesbsts)
    - [堆（Heap） / 优先级队列（Priority Queue） / 二叉堆（Binary Heap）](#堆heap--优先级队列priority-queue--二叉堆binary-heap)
    - [字典树（Tries）](#字典树tries)
    - [平衡查找树（Balanced search trees）](#平衡查找树balanced-search-trees)
    - [N 叉树（K 叉树、M 叉树）](#n-叉树k-叉树m-叉树)
- [排序](#排序sorting)
- [图（Graphs）](#图graphs)
- [更多知识](#更多知识)
    - [递归](#递归recursion)
    - [动态规划](#动态规划dynamic-programming)
    - [组合 & 概率](#组合combinatorics-n-中选-k-个--概率probability)
    - [NP, NP-完全和近似算法](#np-np-完全和近似算法)
    - [缓存](#缓存cache)
    - [进程和线程](#进程processe和线程thread)
    - [系统设计、可伸缩性、数据处理](#系统设计可伸缩性数据处理)
    - [论文](#论文)
    - [测试](#测试)
    - [调度](#调度)
    - [实现系统例程](#实现系统例程)
    - [字符串搜索和操作](#字符串搜索和操作)
- [终面](#终面)
- [书籍](#书籍)
- [编码练习和挑战](#编码练习和挑战)
- [当你临近面试时](#当你临近面试时)
- [你的简历](#你的简历)
- [当面试来临的时候](#当面试来临的时候)
- [问面试官的问题](#问面试官的问题)
- [当你获得了梦想的职位](#当你获得了梦想的职位)

---------------- 下面的内容是可选的 ----------------

- [附加的学习](#附加的学习)
    - [Unicode](#unicode)
    - [字节顺序](#字节顺序)
    - [Emacs and vi(m)](#emacs-and-vim)
    - [Unix 命令行工具](#unix-命令行工具)
    - [信息资源 (视频)](#信息资源-视频)
    - [奇偶校验位 & 汉明码 (视频)](#奇偶校验位--汉明码-视频)
    - [系统熵值（系统复杂度）](#系统熵值系统复杂度)
    - [密码学](#密码学)
    - [压缩](#压缩)
    - [网络 (视频)](#网络-视频)
    - [计算机安全](#计算机安全)
    - [释放缓存](#释放缓存)
    - [并行/并发编程](#并行并发编程)
    - [设计模式](#设计模式)
    - [信息传输, 序列化, 和队列化的系统](#信息传输-序列化和队列化的系统)
    - [快速傅里叶变换](#快速傅里叶变换)
    - [布隆过滤器](#布隆过滤器)
    - [van Emde Boas 树](#van-emde-boas-树)
    - [更深入的数据结构](#更深入的数据结构)
    - [跳表](#跳表)
    - [网络流](#网络流)
    - [不相交集 & 联合查找](#不相交集--联合查找)
    - [快速处理数学](#math-for-fast-processing)
    - [树堆 (Treap)](#树堆-treap)
    - [线性规划](#线性规划linear-programming视频)
    - [几何：凸包（Geometry, Convex hull）](#几何凸包geometry-convex-hull视频)
    - [离散数学](#离散数学)
    - [机器学习](#机器学习machine-learning)
    - [Go 语言](#go-语言)
- [一些主题的额外内容](#一些主题的额外内容)
- [视频系列](#视频系列)
- [计算机科学课程](#计算机科学课程)

---


## 必备知识

- [x] **计算机是如何处理一段程序：**
    - [x] [CPU 是如何执行代码（视频）](https://www.youtube.com/watch?v=42KTvGYQYnA)
    - [x] [机器码指令（视频）](https://www.youtube.com/watch?v=Mv2XQgpbTNE)

- [x] **编译器**
    - [x] [编译器是如何在 ~1 分钟内工作（视频）](https://www.youtube.com/watch?v=IhC7sdYe-Jg)
    - [x] [Hardvard CS50 —— 编译器（视频）](https://www.youtube.com/watch?v=CSZLNYF4Klo)
    - [x] [C++（视频）](https://www.youtube.com/watch?v=twodd1KFfGk)
    - [x] [掌握编译器的优化（C++）（视频）](https://www.youtube.com/watch?v=FnGCDLhaxKU)

- [x] **浮点数是如何存储的：**
    - [x] 简单的 8-bit：[浮点数的表达形式　—— 1（视频 —— 在计算上有一个错误 —— 详情请查看视频的介绍）](https://www.youtube.com/watch?v=ji3SfClm8TU)
    - [x] 32 bit：[IEEE754 32-bit 浮点二进制（视频）](https://www.youtube.com/watch?v=50ZYcZebIec)
    简单的说，一个小数在计算机中是以二进制存储的，而且存储的位数是有限的（双精度和单精度不一样），比如：9.4-9-0.4 != 0
    而0.4的二进制是0.0110-0110-0110-……无限循环的
    从而9.4的二进制是1001.0110-0110
    9.4比0.4 多了1001 所以最后面肯定没有0.4的位数多，从而导致了!=0
    [参见float](../../golang/float.md)

## 算法复杂度 / Big-O / 渐进分析法
- 并不需要实现
- [ ] [Harvard's CS50](https://cs50.harvard.edu/x/2020/)
- [ ] [Harvard CS50 —— 渐进表示（视频）](https://www.youtube.com/watch?v=iOq5kSKqeR4)
- [ ] [Big O 记号（通用快速教程）（视频）](https://www.youtube.com/watch?v=V6mKVRU1evU)
- [ ] [Big O 记号（以及 Omega 和 Theta）——  最佳数学解释（视频）](https://www.youtube.com/watch?v=ei-A_wy5Yxw&index=2&list=PL1BaGV1cIH4UhkL8a9bJGG356covJ76qN)
- [ ] Skiena 算法：
    - [视频](https://www.youtube.com/watch?v=gSyDMtdPNpU&index=2&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
    - [幻灯片](http://www3.cs.stonybrook.edu/~algorith/video-lectures/2007/lecture2.pdf)
- [ ] [对于算法复杂度分析的一次详细介绍](http://discrete.gr/complexity/)
- [ ] [增长阶数（Orders of Growth）（视频）](https://class.coursera.org/algorithmicthink1-004/lecture/59)
- [ ] [渐进性（Asymptotics）（视频）](https://class.coursera.org/algorithmicthink1-004/lecture/61)
- [ ] [UC Berkeley Big O（视频）](https://youtu.be/VIS4YDpuP98)
- [ ] [UC Berkeley Big Omega（视频）](https://youtu.be/ca3e7UVmeUc)
- [ ] [平摊分析法（Amortized Analysis）（视频）](https://www.youtube.com/watch?v=B3SpQZaAZP4&index=10&list=PL1BaGV1cIH4UhkL8a9bJGG356covJ76qN)
- [ ] [举证“Big O”（视频）](https://class.coursera.org/algorithmicthink1-004/lecture/63)
- [ ] 高级编程（包括递归关系和主定理）：
    - [计算性复杂度：第一部](https://www.topcoder.com/community/data-science/data-science-tutorials/computational-complexity-section-1/)
    - [计算性复杂度：第二部](https://www.topcoder.com/community/data-science/data-science-tutorials/computational-complexity-section-2/)
- [ ] [速查表（Cheat sheet）](http://bigocheatsheet.com/)

    如果部分课程过于学术性，你可直接跳到文章底部，去查看离散数学的视频以获取相关背景知识。

## 数据结构

- ### 数组（Arrays）
    - 实现一个可自动调整大小的动态数组。
    - [ ] 介绍：
        - [数组（视频）](https://www.coursera.org/learn/data-structures/lecture/OsBSF/arrays)
        - [数组的基础知识（视频）](https://archive.org/details/0102WhatYouShouldKnow/02_04-basicArrays.mp4)
        - [多维数组（视频）](https://archive.org/details/0102WhatYouShouldKnow/02_05-multidimensionalArrays.mp4)
        - [动态数组（视频）](https://www.coursera.org/learn/data-structures/lecture/EwbnV/dynamic-arrays)
        - [不规则数组（视频）](https://archive.org/details/0102WhatYouShouldKnow/02_06-jaggedArrays.mp4)
        - [调整数组的大小（视频）](https://archive.org/details/0102WhatYouShouldKnow/03_01-resizableArrays.mp4)
    - [ ] 实现一个动态数组（可自动调整大小的可变数组）：
        - [ ] 练习使用数组和指针去编码，并且指针是通过计算去跳转而不是使用索引
        - [ ] 通过分配内存来新建一个原生数据型数组
            - 可以使用 int 类型的数组，但不能使用其语法特性
            - 从大小为16或更大的数（使用2的倍数 —— 16、32、64、128）开始编写
        - [ ] size() —— 数组元素的个数
        - [ ] capacity() —— 可容纳元素的个数
        - [ ] is_empty()
        - [ ] at(index) —— 返回对应索引的元素，且若索引越界则愤然报错
        - [ ] push(item)
        - [ ] insert(index, item) —— 在指定索引中插入元素，并把后面的元素依次后移
        - [ ] prepend(item) —— 可以使用上面的 insert 函数，传参 index 为 0
        - [ ] pop() —— 删除在数组末端的元素，并返回其值
        - [ ] delete(index) —— 删除指定索引的元素，并把后面的元素依次前移
        - [ ] remove(item) —— 删除指定值的元素，并返回其索引（即使有多个元素）
        - [ ] find(item) —— 寻找指定值的元素并返回其中第一个出现的元素其索引，若未找到则返回 -1
        - [ ] resize(new_capacity) // 私有函数
            - 若数组的大小到达其容积，则变大一倍
            - 获取元素后，若数组大小为其容积的1/4，则缩小一半
    - [ ] 时间复杂度
        - 在数组末端增加/删除、定位、更新元素，只允许占 O(1) 的时间复杂度（平摊（amortized）去分配内存以获取更多空间）
        - 在数组任何地方插入/移除元素，只允许 O(n) 的时间复杂度
    - [ ] 空间复杂度
        - 因为在内存中分配的空间邻近，所以有助于提高性能
        - 空间需求 = （大于或等于 n 的数组容积）* 元素的大小。即便空间需求为 2n，其空间复杂度仍然是 O(n)

- ### 链表（Linked Lists）
    - [ ] 介绍：
        - [ ] [单向链表（视频）](https://www.coursera.org/learn/data-structures/lecture/kHhgK/singly-linked-lists)
        - [ ] [CS 61B —— 链表（视频）](https://www.youtube.com/watch?v=sJtJOtXCW_M&list=PL-XXv-cvA_iAlnI-BQr9hjqADPBtujFJd&index=5)
    - [ ] [C 代码（视频）](https://www.youtube.com/watch?v=QN6FPiD0Gzo)
        - 并非看完整个视频，只需要看关于节点结果和内存分配那一部分即可
    - [ ] 链表 vs 数组：
        - [基本链表 Vs 数组（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/rjBs9/core-linked-lists-vs-arrays)
        - [在现实中，链表 Vs 数组（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/QUaUd/in-the-real-world-lists-vs-arrays)
    - [ ] [为什么你需要避免使用链表（视频）](https://www.youtube.com/watch?v=YQs6IC-vgmo)
    - [ ] 的确：你需要关于“指向指针的指针”的相关知识：（因为当你传递一个指针到一个函数时，该函数可能会改变指针所指向的地址）该页只是为了让你了解“指向指针的指针”这一概念。但我并不推荐这种链式遍历的风格。因为，这种风格的代码，其可读性和可维护性太低。
        - [指向指针的指针](https://www.eskimo.com/~scs/cclass/int/sx8.html)
    - [ ] 实现（我实现了使用尾指针以及没有使用尾指针这两种情况）：
        - [ ] size() —— 返回链表中数据元素的个数
        - [ ] empty() —— 若链表为空则返回一个布尔值 true
        - [ ] value_at(index) —— 返回第 n 个元素的值（从0开始计算）
        - [ ] push_front(value) —— 添加元素到链表的首部
        - [ ] pop_front() —— 删除首部元素并返回其值
        - [ ] push_back(value) —— 添加元素到链表的尾部
        - [ ] pop_back() —— 删除尾部元素并返回其值
        - [ ] front() —— 返回首部元素的值
        - [ ] back() —— 返回尾部元素的值
        - [ ] insert(index, value) —— 插入值到指定的索引，并把当前索引的元素指向到新的元素
        - [ ] erase(index) —— 删除指定索引的节点
        - [ ] value_n_from_end(n) —— 返回倒数第 n 个节点的值
        - [ ] reverse() —— 逆序链表
        - [ ] remove_value(value) —— 删除链表中指定值的第一个元素
    - [ ] 双向链表
        - [介绍（视频）](https://www.coursera.org/learn/data-structures/lecture/jpGKD/doubly-linked-lists)
        - 并不需要实现

- ### 堆栈（Stack）
    - [ ] [堆栈（视频）](https://www.coursera.org/learn/data-structures/lecture/UdKzQ/stacks)
    - [ ] [使用堆栈 —— 后进先出（视频）](https://archive.org/details/0102WhatYouShouldKnow/05_01-usingStacksForLast-inFirst-out.mp4)
    - [ ] 可以不实现，因为使用数组来实现并不重要

- ### 队列（Queue）
    - [ ] [使用队列 —— 先进先出（视频）](https://archive.org/details/0102WhatYouShouldKnow/05_03-usingQueuesForFirst-inFirst-out.mp4)
    - [ ] [队列（视频）](https://www.coursera.org/learn/data-structures/lecture/EShpq/queue)
    - [ ] [原型队列/先进先出（FIFO）](https://en.wikipedia.org/wiki/Circular_buffer)
    - [ ] [优先级队列（视频）](https://archive.org/details/0102WhatYouShouldKnow/05_04-priorityQueuesAndDeques.mp4)
    - [ ] 使用含有尾部指针的链表来实现:
        - enqueue(value) —— 在尾部添加值
        - dequeue() —— 删除最早添加的元素并返回其值（首部元素）
        - empty()
    - [ ] 使用固定大小的数组实现：
        - enqueue(value) —— 在可容的情况下添加元素到尾部
        - dequeue() —— 删除最早添加的元素并返回其值
        - empty()
        - full()
    - [ ] 花销：
        - 在糟糕的实现情况下，使用链表所实现的队列，其入列和出列的时间复杂度将会是 O(n)。因为，你需要找到下一个元素，以致循环整个队列
        - enqueue：O(1)（平摊（amortized）、链表和数组 [探测（probing）]）
        - dequeue：O(1)（链表和数组）
        - empty：O(1)（链表和数组）

- ### 哈希表（Hash table）
    - [ ] 视频：
        - [ ] [链式哈希表（视频）](https://www.youtube.com/watch?v=0M_kIqhwbFo&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=8)
        - [ ] [Table Doubling 和 Karp-Rabin（视频）](https://www.youtube.com/watch?v=BRO7mVIFt08&index=9&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
        - [ ] [Open Addressing 和密码型哈希（Cryptographic Hashing）（视频）](https://www.youtube.com/watch?v=rvdJDijO2Ro&index=10&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
        - [ ] [PyCon 2010：The Mighty Dictionary（视频）](https://www.youtube.com/watch?v=C4Kc8xzcA68)
        - [ ] [（进阶）随机取样（Randomization）：全域哈希（Universal Hashing）& 完美哈希（Perfect Hashing）（视频）](https://www.youtube.com/watch?v=z0lJ2k0sl1g&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp&index=11)
        - [ ] [（进阶）完美哈希（Perfect hashing）（视频）](https://www.youtube.com/watch?v=N0COwN14gt0&list=PL2B4EEwhKD-NbwZ4ezj7gyc_3yNrojKM9&index=4)

    - [ ] 在线课程：
        - [ ] [哈希函数的掌握（视频）](https://archive.org/details/0102WhatYouShouldKnow/06_02-understandingHashFunctions.mp4)
        - [ ] [使用哈希表（视频）](https://archive.org/details/0102WhatYouShouldKnow/06_03-usingHashTables.mp4)
        - [ ] [哈希表的支持（视频）](https://archive.org/details/0102WhatYouShouldKnow/06_04-supportingHashing.mp4)
        - [ ] [哈希表的语言支持（视频）](https://archive.org/details/0102WhatYouShouldKnow/06_05-languageSupportForHashTables.mp4)
        - [ ] [基本哈希表（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/m7UuP/core-hash-tables)
        - [ ] [数据结构（视频）](https://www.coursera.org/learn/data-structures/home/week/3)
        - [ ] [电话薄问题（Phone Book Problem）（视频）](https://www.coursera.org/learn/data-structures/lecture/NYZZP/phone-book-problem)
        - [ ] 分布式哈希表：
            - [Dropbox 中的瞬时上传及存储优化（视频）](https://www.coursera.org/learn/data-structures/lecture/DvaIb/instant-uploads-and-storage-optimization-in-dropbox)
            - [分布式哈希表（视频）](https://www.coursera.org/learn/data-structures/lecture/tvH8H/distributed-hash-tables)

    - [ ] 使用线性探测的数组去实现
        - hash(k, m) —— m 是哈希表的大小
        - add(key, value) —— 如果 key 已存在则更新值
        - exists(key)
        - get(key)
        - remove(key)

## 更多的知识

- ### 二分查找（Binary search）
    - [ ] [二分查找（视频）](https://www.youtube.com/watch?v=D5SrAga1pno)
    - [ ] [二分查找（视频）](https://www.khanacademy.org/computing/computer-science/algorithms/binary-search/a/binary-search)
    - [ ] [详情](https://www.topcoder.com/community/data-science/data-science-tutorials/binary-search/)
    - [ ] 实现：
        - 二分查找（在一个已排序好的整型数组中查找）
        - 迭代式二分查找

- ### 按位运算（Bitwise operations）
    - [ ] [Bits 速查表](https://github.com/jwasham/coding-interview-university/blob/master/extras/cheat%20sheets/bits-cheat-cheet.pdf)
        - 你需要知道大量2的幂数值（从2^1 到 2^16 及 2^32）
    - [ ] 好好理解位操作符的含义：&、|、^、~、>>、<<
        - [ ] [字码（words）](https://en.wikipedia.org/wiki/Word_(computer_architecture))
        - [ ] 好的介绍：
            [位操作（视频）](https://www.youtube.com/watch?v=7jkIUgLC29I)
        - [ ] [C 语言编程教程 2-10：按位运算（视频）](https://www.youtube.com/watch?v=d0AwjSpNXR0)
        - [ ] [位操作](https://en.wikipedia.org/wiki/Bit_manipulation)
        - [ ] [按位运算](https://en.wikipedia.org/wiki/Bitwise_operation)
        - [ ] [Bithacks](https://graphics.stanford.edu/~seander/bithacks.html)
        - [ ] [位元抚弄者（The Bit Twiddler）](http://bits.stephan-brumme.com/)
        - [ ] [交互式位元抚弄者（The Bit Twiddler Interactive）](http://bits.stephan-brumme.com/interactive.html)
    - [ ] 一补数和补码
        - [二进制：利 & 弊（为什么我们要使用补码）（视频）](https://www.youtube.com/watch?v=lKTsv6iVxV4)
        - [一补数（1s Complement）](https://en.wikipedia.org/wiki/Ones%27_complement)
        - [补码（2s Complement）](https://en.wikipedia.org/wiki/Two%27s_complement)
    - [ ] 计算置位（Set Bits）
        - [计算一个字节中置位（Set Bits）的四种方式（视频）](https://youtu.be/Hzuzo9NJrlc)
        - [计算比特位](https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan)
        - [如何在一个 32 位的整型中计算置位（Set Bits）的数量](http://stackoverflow.com/questions/109023/how-to-count-the-number-of-set-bits-in-a-32-bit-integer)
    - [ ] 四舍五入2的幂数：
        - [四舍五入到2的下一幂数](http://bits.stephan-brumme.com/roundUpToNextPowerOfTwo.html)
    - [ ] 交换值：
        - [交换（Swap）](http://bits.stephan-brumme.com/swap.html)
    - [ ] 绝对值：
        - [绝对整型（Absolute Integer）](http://bits.stephan-brumme.com/absInteger.html)

## 树（Trees）

- ### 树 —— 笔记 & 背景
    - [ ] [系列：基本树（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/ovovP/core-trees)
    - [ ] [系列：树（视频）](https://www.coursera.org/learn/data-structures/lecture/95qda/trees)
    - 基本的树形结构
    - 遍历
    - 操作算法
    - BFS（广度优先检索，breadth-first search）
        - [MIT（视频）](https://www.youtube.com/watch?v=s-CYnVz-uh4&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=13)
        - 层序遍历（使用队列的 BFS 算法）
            - 时间复杂度： O(n)
            - 空间复杂度：
                - 最好情况： O(1)
                - 最坏情况：O(n/2)=O(n)
    - DFS（深度优先检索，depth-first search）
        - [MIT（视频）](https://www.youtube.com/watch?v=AfSk24UTFS8&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=14)
        - 笔记：
            - 时间复杂度：O(n)
            - 空间复杂度：
                - 最好情况：O(log n) - 树的平均高度
                - 最坏情况：O(n)
        - 中序遍历（DFS：左、节点本身、右）
        - 后序遍历（DFS：左、右、节点本身）
        - 先序遍历（DFS：节点本身、左、右）

- ### 二叉查找树（Binary search trees）：BSTs
    - [ ] [二叉查找树概览（视频）](https://www.youtube.com/watch?v=x6At0nzX92o&index=1&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
    - [ ] [系列（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/p82sw/core-introduction-to-binary-search-trees)
        - 从符号表开始到 BST 程序
    - [ ] [介绍（视频）](https://www.coursera.org/learn/data-structures/lecture/E7cXP/introduction)
    - [ ] [MIT（视频）](https://www.youtube.com/watch?v=9Jry5-82I68)
    - C/C++:
        - [ ] [二叉查找树 —— 在 C/C++ 中实现（视频）](https://www.youtube.com/watch?v=COZK7NATh4k&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=28)
        - [ ] [BST 的实现 —— 在堆栈和堆中的内存分配（视频）](https://www.youtube.com/watch?v=hWokyBoo0aI&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=29)
        - [ ] [在二叉查找树中找到最小和最大的元素（视频）](https://www.youtube.com/watch?v=Ut90klNN264&index=30&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
        - [ ] [寻找二叉树的高度（视频）](https://www.youtube.com/watch?v=_pnqMz5nrRs&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=31)
        - [ ] [二叉树的遍历 —— 广度优先和深度优先策略（视频）](https://www.youtube.com/watch?v=9RHO6jU--GU&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=32)
        - [ ] [二叉树：层序遍历（视频）](https://www.youtube.com/watch?v=86g8jAQug04&index=33&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
        - [ ] [二叉树的遍历：先序、中序、后序（视频）](https://www.youtube.com/watch?v=gm8DUJJhmY4&index=34&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
        - [ ] [判断一棵二叉树是否为二叉查找树（视频）](https://www.youtube.com/watch?v=yEwSGhSsT0U&index=35&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
        - [ ] [从二叉查找树中删除一个节点（视频）](https://www.youtube.com/watch?v=gcULXE7ViZw&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=36)
        - [ ] [二叉查找树中序遍历的后继者（视频）](https://www.youtube.com/watch?v=5cPbNCrdotA&index=37&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] 实现：
        - [ ] insert    // 往树上插值
        - [ ] get_node_count // 查找树上的节点数
        - [ ] print_values // 从小到大打印树中节点的值
        - [ ] delete_tree
        - [ ] is_in_tree // 如果值存在于树中则返回 true
        - [ ] get_height // 返回节点所在的高度（如果只有一个节点，那么高度则为1）
        - [ ] get_min   // 返回树上的最小值
        - [ ] get_max   // 返回树上的最大值
        - [ ] is_binary_search_tree
        - [ ] delete_value
        - [ ] get_successor // 返回给定值的后继者，若没有则返回-1

- ### 堆（Heap） / 优先级队列（Priority Queue） / 二叉堆（Binary Heap）
    - 可视化是一棵树，但通常是以线性的形式存储（数组、链表）
    - [ ] [堆](https://en.wikipedia.org/wiki/Heap_(data_structure))
    - [ ] [介绍（视频）](https://www.coursera.org/learn/data-structures/lecture/2OpTs/introduction)
    - [ ] [无知的实现（视频）](https://www.coursera.org/learn/data-structures/lecture/z3l9N/naive-implementations)
    - [ ] [二叉树（视频）](https://www.coursera.org/learn/data-structures/lecture/GRV2q/binary-trees)
    - [ ] [关于树高的讨论（视频）](https://www.coursera.org/learn/data-structures/supplement/S5xxz/tree-height-remark)
    - [ ] [基本操作（视频）](https://www.coursera.org/learn/data-structures/lecture/0g1dl/basic-operations)
    - [ ] [完全二叉树（视频）](https://www.coursera.org/learn/data-structures/lecture/gl5Ni/complete-binary-trees)
    - [ ] [伪代码（视频）](https://www.coursera.org/learn/data-structures/lecture/HxQo9/pseudocode)
    - [ ] [堆排序 —— 跳到起点（视频）](https://youtu.be/odNJmw5TOEE?list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&t=3291)
    - [ ] [堆排序（视频）](https://www.coursera.org/learn/data-structures/lecture/hSzMO/heap-sort)
    - [ ] [构建一个堆（视频）](https://www.coursera.org/learn/data-structures/lecture/dwrOS/building-a-heap)
    - [ ] [MIT：堆与堆排序（视频）](https://www.youtube.com/watch?v=B7hVxCmfPtM&index=4&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
    - [ ] [CS 61B Lecture 24：优先级队列（视频）](https://www.youtube.com/watch?v=yIUFT6AKBGE&index=24&list=PL4BBB74C7D2A1049C)
    - [ ] [构建线性时间复杂度的堆（大顶堆）](https://www.youtube.com/watch?v=MiyLo8adrWw)
    - [ ] 实现一个大顶堆：
        - [ ] insert
        - [ ] sift_up —— 用于插入元素
        - [ ] get_max —— 返回最大值但不移除元素
        - [ ] get_size() —— 返回存储的元素数量
        - [ ] is_empty() —— 若堆为空则返回 true
        - [ ] extract_max —— 返回最大值并移除
        - [ ] sift_down —— 用于获取最大值元素
        - [ ] remove(i) —— 删除指定索引的元素
        - [ ] heapify —— 构建堆，用于堆排序
        - [ ] heap_sort() —— 拿到一个未排序的数组，然后使用大顶堆进行就地排序
            - 注意：若用小顶堆可节省操作，但导致空间复杂度加倍。（无法做到就地）

- ### 字典树（Tries）
    - 需要注意的是，字典树各式各样。有些有前缀，而有些则没有。有些使用字符串而不使用比特位来追踪路径。
    - 阅读代码，但不实现。
    - [ ] [数据结构笔记及编程技术](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#Tries)
    - [ ] 短课程视频：
        - [ ] [对字典树的介绍（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/08Xyf/core-introduction-to-tries)
        - [ ] [字典树的性能（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/PvlZW/core-performance-of-tries)
        - [ ] [实现一棵字典树（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/DFvd3/core-implementing-a-trie)
    - [ ] [字典树：一个被忽略的数据结构](https://www.toptal.com/java/the-trie-a-neglected-data-structure)
    - [ ] [高级编程 —— 使用字典树](https://www.topcoder.com/community/data-science/data-science-tutorials/using-tries/)
    - [ ] [标准教程（现实中的用例）（视频）](https://www.youtube.com/watch?v=TJ8SkcUSdbU)
    - [ ] [MIT，高阶数据结构，使用字符串追踪路径（可事半功倍）](https://www.youtube.com/watch?v=NinWEPPrkDQ&index=16&list=PLUl4u3cNGP61hsJNdULdudlRL493b-XZf)

- ### 平衡查找树（Balanced search trees）
    - 掌握至少一种平衡查找树（并懂得如何实现）：
    - “在各种平衡查找树当中，AVL 树和2-3树已经成为了过去，而红黑树（red-black trees）看似变得越来越受人青睐。这种令人特别感兴趣的数据结构，亦称伸展树（splay tree）。它可以自我管理，且会使用轮换来移除任何访问过根节点的 key。” —— Skiena
    - 因此，在各种各样的平衡查找树当中，我选择了伸展树来实现。虽然，通过我的阅读，我发现在 Google 的面试中并不会被要求实现一棵平衡查找树。但是，为了胜人一筹，我们还是应该看看如何去实现。在阅读了大量关于红黑树的代码后，我才发现伸展树的实现确实会使得各方面更为高效。
        - 伸展树：插入、查找、删除函数的实现，而如果你最终实现了红黑树，那么请尝试一下：
        - 跳过删除函数，直接实现搜索和插入功能
    - 我希望能阅读到更多关于 B 树的资料，因为它也被广泛地应用到大型的数据库当中。
    - [ ] [自平衡二叉查找树](https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree)

    - [ ] **AVL 树**
        - 实际中：我能告诉你的是，该种树并无太多的用途，但我能看到有用的地方在哪里：AVL 树是另一种平衡查找树结构。其可支持时间复杂度为 O(log n) 的查询、插入及删除。它比红黑树严格意义上更为平衡，从而导致插入和删除更慢，但遍历却更快。正因如此，才彰显其结构的魅力。只需要构建一次，就可以在不重新构造的情况下读取，适合于实现诸如语言字典（或程序字典，如一个汇编程序或解释程序的操作码）。
        - [ ] [MIT AVL 树 / AVL 树的排序（视频）](https://www.youtube.com/watch?v=FNeL18KsWPc&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=6)
        - [ ] [AVL 树（视频）](https://www.coursera.org/learn/data-structures/lecture/Qq5E0/avl-trees)
        - [ ] [AVL 树的实现（视频）](https://www.coursera.org/learn/data-structures/lecture/PKEBC/avl-tree-implementation)
        - [ ] [分离与合并](https://www.coursera.org/learn/data-structures/lecture/22BgE/split-and-merge)

    - [ ] **伸展树**
        - 实际中：伸展树一般用于缓存、内存分配者、路由器、垃圾回收者、数据压缩、ropes（字符串的一种替代品，用于存储长串的文本字符）、Windows NT（虚拟内存、网络及文件系统）等的实现。
        - [ ] [CS 61B：伸展树（Splay trees）（视频）](https://www.youtube.com/watch?v=Najzh1rYQTo&index=23&list=PL-XXv-cvA_iAlnI-BQr9hjqADPBtujFJd)
        - [ ] MIT 教程：伸展树（Splay trees）：
            - 该教程会过于学术，但请观看到最后的10分钟以确保掌握。
            - [视频](https://www.youtube.com/watch?v=QnPl_Y6EqMo)

    - [ ] **2-3查找树**
        - 实际中：2-3树的元素插入非常快速，但却有着查询慢的代价（因为相比较 AVL 树来说，其高度更高）。
        - 你会很少用到2-3树。这是因为，其实现过程中涉及到不同类型的节点。因此，人们更多地会选择红黑树。
        - [ ] [2-3树的直感与定义（视频）](https://www.youtube.com/watch?v=C3SsdUqasD4&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6&index=2)
        - [ ] [2-3树的二元观点](https://www.youtube.com/watch?v=iYvBtGKsqSg&index=3&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
        - [ ] [2-3树（学生叙述）（视频）](https://www.youtube.com/watch?v=TOb1tuEZ2X4&index=5&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)

    - [ ] **2-3-4树 (亦称2-4树)**
        - 实际中：对于每一棵2-4树，都有着对应的红黑树来存储同样顺序的数据元素。在2-4树上进行插入及删除操作等同于在红黑树上进行颜色翻转及轮换。这使得2-4树成为一种用于掌握红黑树背后逻辑的重要工具。这就是为什么许多算法引导文章都会在介绍红黑树之前，先介绍2-4树，尽管**2-4树在实际中并不经常使用**。
        - [ ] [CS 61B Lecture 26：平衡查找树（视频）](https://www.youtube.com/watch?v=zqrqYXkth6Q&index=26&list=PL4BBB74C7D2A1049C)
        - [ ] [自底向上的2-4树（视频）](https://www.youtube.com/watch?v=DQdMYevEyE4&index=4&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
        - [ ] [自顶向下的2-4树（视频）](https://www.youtube.com/watch?v=2679VQ26Fp4&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6&index=5)

    - [ ] **B 树**
        - 有趣的是：为啥叫 B 仍然是一个神秘。因为 B 可代表波音（Boeing）、平衡（Balanced）或 Bayer（联合创造者）
        - 实际中：B 树会被广泛适用于数据库中，而现代大多数的文件系统都会使用到这种树（或变种)。除了运用在数据库中，B 树也会被用于文件系统以快速访问一个文件的任意块。但存在着一个基本的问题，那就是如何将文件块 i 转换成一个硬盘块（或一个柱面-磁头-扇区）上的地址。
        - [ ] [B 树](https://en.wikipedia.org/wiki/B-tree)
        - [ ] [B 树的介绍（视频）](https://www.youtube.com/watch?v=I22wEC1tTGo&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6&index=6)
        - [ ] [B 树的定义及其插入操作（视频）](https://www.youtube.com/watch?v=s3bCdZGrgpA&index=7&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
        - [ ] [B 树的删除操作（视频）](https://www.youtube.com/watch?v=svfnVhJOfMc&index=8&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
        - [ ] [MIT 6.851 —— 内存层次模块（Memory Hierarchy Models）（视频）](https://www.youtube.com/watch?v=V3omVLzI0WE&index=7&list=PLUl4u3cNGP61hsJNdULdudlRL493b-XZf)
            - 覆盖有高速缓存参数无关型（cache-oblivious）B 树和非常有趣的数据结构
            - 头37分钟讲述的很专业，或许可以跳过（B 指块的大小、即缓存行的大小）

    - [ ] **红黑树**
        - 实际中：红黑树提供了在最坏情况下插入操作、删除操作和查找操作的时间保证。这些时间值的保障不仅对时间敏感型应用有用，例如实时应用，还对在其他数据结构中块的构建非常有用，而这些数据结构都提供了最坏情况下的保障；例如，许多用于计算几何学的数据结构都可以基于红黑树，而目前 Linux 系统所采用的完全公平调度器（the Completely Fair Scheduler）也使用到了该种树。在 Java 8中，红黑树也被用于存储哈希列表集合中相同的数据，而不是使用链表及哈希码。
        - [ ] [Aduni —— 算法 —— 课程4（该链接直接跳到开始部分）（视频）](https://youtu.be/1W3x0f_RmUo?list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&t=3871)
        - [ ] [Aduni —— 算法 —— 课程5（视频）](https://www.youtube.com/watch?v=hm2GHwyKF1o&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=5)
        - [ ] [黑树（Black Tree）](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree)
        - [ ] [二分查找及红黑树的介绍](https://www.topcoder.com/community/data-science/data-science-tutorials/an-introduction-to-binary-search-and-red-black-trees/)

- ### N 叉树（K 叉树、M 叉树）
    - 注意：N 或 K 指的是分支系数（即树的最大分支数）：
        - 二叉树是一种分支系数为2的树
        - 2-3树是一种分支系数为3的树
    - [ ] [K 叉树](https://en.wikipedia.org/wiki/K-ary_tree)

## 排序（Sorting）

- [ ] 笔记:
    - 实现各种排序 & 知道每种排序的最坏、最好和平均的复杂度分别是什么场景:
        - 不要用冒泡排序 - 大多数情况下效率感人 - 时间复杂度 O(n^2), 除非 n <= 16
    - [ ] 排序算法的稳定性 ("快排是稳定的么?")
        - [排序算法的稳定性](https://en.wikipedia.org/wiki/Sorting_algorithm#Stability)
        - [排序算法的稳定性](http://stackoverflow.com/questions/1517793/stability-in-sorting-algorithms)
        - [排序算法的稳定性](http://stackoverflow.com/questions/1517793/stability-in-sorting-algorithms)
        - [排序算法的稳定性](http://www.geeksforgeeks.org/stability-in-sorting-algorithms/)
        - [排序算法 - 稳定性](http://homepages.math.uic.edu/~leon/cs-mcs401-s08/handouts/stability.pdf)
    - [ ] 哪种排序算法可以用链表？哪种用数组？哪种两者都可？
        - 并不推荐对一个链表排序，但归并排序是可行的.
        - [链表的归并排序](http://www.geeksforgeeks.org/merge-sort-for-linked-list/)

- 关于堆排序，请查看前文堆的数据结构部分。堆排序很强大，不过是非稳定排序。

- [ ] [冒泡排序 (video)](https://www.youtube.com/watch?v=P00xJgWzz2c&index=1&list=PL89B61F78B552C1AB)
- [ ] [冒泡排序分析 (video)](https://www.youtube.com/watch?v=ni_zk257Nqo&index=7&list=PL89B61F78B552C1AB)
- [ ] [插入排序 & 归并排序 (video)](https://www.youtube.com/watch?v=Kg4bqzAqRBM&index=3&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
- [ ] [插入排序 (video)](https://www.youtube.com/watch?v=c4BRHC7kTaQ&index=2&list=PL89B61F78B552C1AB)
- [ ] [归并排序 (video)](https://www.youtube.com/watch?v=GCae1WNvnZM&index=3&list=PL89B61F78B552C1AB)
- [ ] [快排 (video)](https://www.youtube.com/watch?v=y_G9BkAm6B8&index=4&list=PL89B61F78B552C1AB)
- [ ] [选择排序 (video)](https://www.youtube.com/watch?v=6nDMgr0-Yyo&index=8&list=PL89B61F78B552C1AB)

- [ ] 斯坦福大学关于排序算法的视频:
    - [ ] [课程 15 | 编程抽象 (video)](https://www.youtube.com/watch?v=ENp00xylP7c&index=15&list=PLFE6E58F856038C69)
    - [ ] [课程 16 | 编程抽象 (video)](https://www.youtube.com/watch?v=y4M9IVgrVKo&index=16&list=PLFE6E58F856038C69)

- [ ] Shai Simonson 视频, [Aduni.org](http://www.aduni.org/):
    - [ ] [算法 - 排序 - 第二讲 (video)](https://www.youtube.com/watch?v=odNJmw5TOEE&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=2)
    - [ ] [算法 - 排序2 - 第三讲 (video)](https://www.youtube.com/watch?v=hj8YKFTFKEE&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=3)

- [ ] Steven Skiena 关于排序的视频:
    - [ ] [课程从 26:46 开始 (video)](https://youtu.be/ute-pmMkyuk?list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&t=1600)
    - [ ] [课程从 27:40 开始 (video)](https://www.youtube.com/watch?v=yLvp-pB8mak&index=8&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
    - [ ] [课程从 35:00 开始 (video)](https://www.youtube.com/watch?v=q7K9otnzlfE&index=9&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
    - [ ] [课程从 23:50 开始 (video)](https://www.youtube.com/watch?v=TvqIGu9Iupw&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=10)

- [ ] 加州大学伯克利分校（UC Berkeley） 大学课程:
    - [ ] [CS 61B 课程 29: 排序 I (video)](https://www.youtube.com/watch?v=EiUvYS2DT6I&list=PL4BBB74C7D2A1049C&index=29)
    - [ ] [CS 61B 课程 30: 排序 II (video)](https://www.youtube.com/watch?v=2hTY3t80Qsk&list=PL4BBB74C7D2A1049C&index=30)
    - [ ] [CS 61B 课程 32: 排序 III (video)](https://www.youtube.com/watch?v=Y6LOLpxg6Dc&index=32&list=PL4BBB74C7D2A1049C)
    - [ ] [CS 61B 课程 33: 排序 V (video)](https://www.youtube.com/watch?v=qNMQ4ly43p4&index=33&list=PL4BBB74C7D2A1049C)

- [ ] - 归并排序:
    - [ ] [使用外部数组](http://www.cs.yale.edu/homes/aspnes/classes/223/examples/sorting/mergesort.c)
    - [ ] [对原数组直接排序](https://github.com/jwasham/practice-cpp/blob/master/merge_sort/merge_sort.cc)
- [ ] - 快速排序:
    - [ ] [实现](http://www.cs.yale.edu/homes/aspnes/classes/223/examples/randomization/quick.c)
    - [ ] [实现](https://github.com/jwasham/practice-c/blob/master/quick_sort/quick_sort.c)

- [ ] 实现:
    - [ ] 归并：平均和最差情况的时间复杂度为 O(n log n)。
    - [ ] 快排：平均时间复杂度为 O(n log n)。
    - 选择排序和插入排序的最坏、平均时间复杂度都是 O(n^2)。
    - 关于堆排序，请查看前文堆的数据结构部分。

- [ ] 有兴趣的话，还有一些补充 - 但并不是必须的:
    - [ ] [基数排序](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#radixSort)
    - [ ] [基数排序 (video)](https://www.youtube.com/watch?v=xhr26ia4k38)
    - [ ] [基数排序, 计数排序 (线性时间内) (video)](https://www.youtube.com/watch?v=Nz1KZXbghj8&index=7&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
    - [ ] [随机算法: 矩阵相乘, 快排, Freivalds' 算法 (video)](https://www.youtube.com/watch?v=cNB2lADK3_s&index=8&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)
    - [ ] [线性时间内的排序 (video)](https://www.youtube.com/watch?v=pOKy3RZbSws&list=PLUl4u3cNGP61hsJNdULdudlRL493b-XZf&index=14)

## 图（Graphs）

图论能解决计算机科学里的很多问题，所以这一节会比较长，像树和排序的部分一样。

- Yegge 的笔记:
    - 有 3 种基本方式在内存里表示一个图:
        - 对象和指针
        - 矩阵
        - 邻接表
    - 熟悉以上每一种图的表示法，并了解各自的优缺点
    - 宽度优先搜索和深度优先搜索 - 知道它们的计算复杂度和设计上的权衡以及如何用代码实现它们
    - 遇到一个问题时，首先尝试基于图的解决方案，如果没有再去尝试其他的。

- [ ] Skiena 教授的课程 - 很不错的介绍:
    - [ ] [CSE373 2012 - 课程 11 - 图的数据结构 (video)](https://www.youtube.com/watch?v=OiXxhDrFruw&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=11)
    - [ ] [CSE373 2012 - 课程 12 - 广度优先搜索 (video)](https://www.youtube.com/watch?v=g5vF8jscteo&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=12)
    - [ ] [CSE373 2012 - 课程 13 - 图的算法 (video)](https://www.youtube.com/watch?v=S23W6eTcqdY&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=13)
    - [ ] [CSE373 2012 - 课程 14 - 图的算法 (1) (video)](https://www.youtube.com/watch?v=WitPBKGV0HY&index=14&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
    - [ ] [CSE373 2012 - 课程 15 - 图的算法 (2) (video)](https://www.youtube.com/watch?v=ia1L30l7OIg&index=15&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
    - [ ] [CSE373 2012 - 课程 16 - 图的算法 (3) (video)](https://www.youtube.com/watch?v=jgDOQq6iWy8&index=16&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)

- [ ] 图 (复习和其他):

    - [ ] [6.006 单源最短路径问题 (video)](https://www.youtube.com/watch?v=Aa2sqUhIn-E&index=15&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
    - [ ] [6.006 Dijkstra 算法 (video)](https://www.youtube.com/watch?v=2E7MmKv0Y24&index=16&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
    - [ ] [6.006 Bellman-Ford 算法(video)](https://www.youtube.com/watch?v=ozsuci5pIso&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=17)
    - [ ] [6.006 Dijkstra 效率优化 (video)](https://www.youtube.com/watch?v=CHvQ3q_gJ7E&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=18)
    - [ ] [Aduni: 图的算法 I - 拓扑排序, 最小生成树, Prim 算法 -  第六课 (video)]( https://www.youtube.com/watch?v=i_AQT_XfvD8&index=6&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm)
    - [ ] [Aduni: 图的算法 II - 深度优先搜索, 广度优先搜索, Kruskal 算法, 并查集数据结构 - 第七课 (video)]( https://www.youtube.com/watch?v=ufj5_bppBsA&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=7)
    - [ ] [Aduni: 图的算法 III: 最短路径 - 第八课 (video)](https://www.youtube.com/watch?v=DiedsPsMKXc&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=8)
    - [ ] [Aduni: 图的算法. IV: 几何算法介绍 - 第九课 (video)](https://www.youtube.com/watch?v=XIAQRlNkJAw&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=9)
    - [ ] [CS 61B 2014 (从 58:09 开始) (video)](https://youtu.be/dgjX4HdMI-Q?list=PL-XXv-cvA_iAlnI-BQr9hjqADPBtujFJd&t=3489)
    - [ ] [CS 61B 2014: 加权图 (video)](https://www.youtube.com/watch?v=aJjlQCFwylA&list=PL-XXv-cvA_iAlnI-BQr9hjqADPBtujFJd&index=19)
    - [ ] [贪心算法: 最小生成树 (video)](https://www.youtube.com/watch?v=tKwnms5iRBU&index=16&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)
    - [ ] [图的算法之强连通分量 Kosaraju 算法 (video)](https://www.youtube.com/watch?v=RpgcYiky7uw)

- 完整的 Coursera 课程:
    - [ ] [图的算法 (video)](https://www.coursera.org/learn/algorithms-on-graphs/home/welcome)

- Yegge: 如果有机会，可以试试研究更酷炫的算法:
    - [ ] Dijkstra 算法 - 上文 - 6.006
    - [ ] A* 算法
        - [ ] [A* 算法](https://en.wikipedia.org/wiki/A*_search_algorithm)
        - [ ] [A* 寻路教程 (video)](https://www.youtube.com/watch?v=KNXfSOx4eEE)
        - [ ] [A* 寻路 (E01: 算法解释) (video)](https://www.youtube.com/watch?v=-L-WgKMFuhE)

- 我会实现:
    - [ ] DFS 邻接表 (递归)
    - [ ] DFS 邻接表 (栈迭代)
    - [ ] DFS 邻接矩阵 (递归)
    - [ ] DFS 邻接矩阵 (栈迭代)
    - [ ] BFS 邻接表
    - [ ] BFS 邻接矩阵
    - [ ] 单源最短路径问题 (Dijkstra)
    - [ ] 最小生成树
    - 基于 DFS 的算法 (根据上文 Aduni 的视频):
        - [ ] 检查环 (我们会先检查是否有环存在以便做拓扑排序)
        - [ ] 拓扑排序
        - [ ] 计算图中的连通分支
        - [ ] 列出强连通分量
        - [ ] 检查双向图

可以从 Skiena 的书（参考下面的书推荐小节）和面试书籍中学习更多关于图的实践。

## 更多知识

- ### 递归（Recursion）
    - [ ] Stanford 大学关于递归 & 回溯的课程:
        - [ ] [课程 8 | 抽象编程 (video)](https://www.youtube.com/watch?v=gl3emqCuueQ&list=PLFE6E58F856038C69&index=8)
        - [ ] [课程 9 | 抽象编程 (video)](https://www.youtube.com/watch?v=uFJhEPrbycQ&list=PLFE6E58F856038C69&index=9)
        - [ ] [课程 10 | 抽象编程 (video)](https://www.youtube.com/watch?v=NdF1QDTRkck&index=10&list=PLFE6E58F856038C69)
        - [ ] [课程 11 | 抽象编程 (video)](https://www.youtube.com/watch?v=p-gpaIGRCQI&list=PLFE6E58F856038C69&index=11)
    - 什么时候适合使用
    - 尾递归会更好么?
        - [ ] [什么是尾递归以及为什么它如此糟糕?](https://www.quora.com/What-is-tail-recursion-Why-is-it-so-bad)
        - [ ] [尾递归 (video)](https://www.youtube.com/watch?v=L1jjXGfxozc)

- ### 动态规划（Dynamic Programming）
    - **注意** ：动态规划是门极为重要的技术，尽管其并未被 Google 提供的准备手册提及，但你可能会对寻求最佳解的方式有点疑问，所以我将其列入这份表单。
    - 这一部分会有点困难，每个可以用动态规划解决的问题都必须先定义出递推关系，要推导出来可能会有点棘手。
    - 我建议先阅读和学习足够多的动态规划的例子，以便对解决 DP 问题的一般模式有个扎实的理解。

    - [ ] 视频:
        - Skiena 的视频可能会有点难跟上，有时候他用白板写的字会比较小，难看清楚。
        - [ ] [Skiena: CSE373 2012 - 课程 19 - 动态规划介绍 (video)](https://youtu.be/Qc2ieXRgR0k?list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&t=1718)
        - [ ] [Skiena: CSE373 2012 - 课程 20 - 编辑距离 (video)](https://youtu.be/IsmMhMdyeGY?list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&t=2749)
        - [ ] [Skiena: CSE373 2012 - 课程 21 - 动态规划举例 (video)](https://youtu.be/o0V9eYF4UI8?list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&t=406)
        - [ ] [Skiena: CSE373 2012 - 课程 22 - 动态规划应用 (video)](https://www.youtube.com/watch?v=dRbMC1Ltl3A&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=22)
        - [ ] [Simonson: 动态规划 0 (starts at 59:18) (video)](https://youtu.be/J5aJEcOr6Eo?list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&t=3558)
        - [ ] [Simonson: 动态规划 I - 课程 11 (video)](https://www.youtube.com/watch?v=0EzHjQ_SOeU&index=11&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm)
        - [ ] [Simonson: 动态规划 II - 课程 12 (video)](https://www.youtube.com/watch?v=v1qiRwuJU7g&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=12)
        - [ ] 单独的 DP 问题 (每一个视频都很短):
            [动态规划 (video)](https://www.youtube.com/playlist?list=PLrmLmBdmIlpsHaNTPP_jHHDx_os9ItYXr)
    - [ ] Yale 课程笔记:
        - [ ] [动态规划](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#dynamicProgramming)
    - [ ] Coursera 课程:
        - [ ] [RNA 二级结构问题 (video)](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/80RrW/the-rna-secondary-structure-problem)
        - [ ] [动态规划算法 (video)](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/PSonq/a-dynamic-programming-algorithm)
        - [ ] [DP 算法描述 (video)](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/oUEK2/illustrating-the-dp-algorithm)
        - [ ] [DP 算法的运行时间 (video)](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/nfK2r/running-time-of-the-dp-algorithm)
        - [ ] [DP vs 递归实现 (video)](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/M999a/dp-vs-recursive-implementation)
        - [ ] [全局成对序列排列 (video)](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/UZ7o6/global-pairwise-sequence-alignment)
        - [ ] [本地成对序列排列 (video)](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/WnNau/local-pairwise-sequence-alignment)

- ### 组合（Combinatorics） (n 中选 k 个) & 概率（Probability）
    - [ ] [数据技巧: 如何找出阶乘、排列和组合(选择) (video)](https://www.youtube.com/watch?v=8RRo6Ti9d0U)
    - [ ] [来点学校的东西: 概率 (video)](https://www.youtube.com/watch?v=sZkAAk9Wwa4)
    - [ ] [来点学校的东西: 概率和马尔可夫链 (video)](https://www.youtube.com/watch?v=dNaJg-mLobQ)
    - [ ] 可汗学院:
        - 课程设置:
            - [ ] [概率理论基础](https://www.khanacademy.org/math/probability/probability-and-combinatorics-topic)
        - 视频 - 41 (每一个都短小精悍):
            - [ ] [概率解释 (video)](https://www.youtube.com/watch?v=uzkc-qNVoOk&list=PLC58778F28211FA19)

- ### NP, NP-完全和近似算法
    - 知道最经典的一些 NP 完全问题，比如旅行商问题和背包问题,
        而且能在面试官试图忽悠你的时候识别出他们。
    - 知道 NP 完全是什么意思.
    - [ ] [计算复杂度 (video)](https://www.youtube.com/watch?v=moPtwq_cVH8&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=23)
    - [ ] Simonson:
        - [ ] [贪心算法. II & 介绍 NP-完全性 (video)](https://youtu.be/qcGnJ47Smlo?list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&t=2939)
        - [ ] [NP-完全性 II & 归约 (video)](https://www.youtube.com/watch?v=e0tGC6ZQdQE&index=16&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm)
        - [ ] [NP-完全性 III (Video)](https://www.youtube.com/watch?v=fCX1BGT3wjE&index=17&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm)
        - [ ] [NP-完全性 IV (video)](https://www.youtube.com/watch?v=NKLDp3Rch3M&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=18)
    - [ ] Skiena:
        - [ ] [CSE373 2012 - 课程 23 - 介绍 NP-完全性 IV (video)](https://youtu.be/KiK5TVgXbFg?list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&t=1508)
        - [ ] [CSE373 2012 - 课程 24 - NP-完全性证明 (video)](https://www.youtube.com/watch?v=27Al52X3hd4&index=24&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
        - [ ] [CSE373 2012 - 课程 25 - NP-完全性挑战 (video)](https://www.youtube.com/watch?v=xCPH4gwIIXM&index=25&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
    - [ ] [复杂度: P, NP, NP-完全性, 规约 (video)](https://www.youtube.com/watch?v=eHZifpgyH_4&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp&index=22)
    - [ ] [复杂度: 近视算法 Algorithms (video)](https://www.youtube.com/watch?v=MEz1J9wY2iM&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp&index=24)
    - [ ] [复杂度: 固定参数算法 (video)](https://www.youtube.com/watch?v=4q-jmGrmxKs&index=25&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)
    - Peter Norvik 讨论旅行商问题的近似最优解:
        - [Jupyter 笔记本](http://nbviewer.jupyter.org/url/norvig.com/ipython/TSP.ipynb)
    - 《算法导论》的第 1048 - 1140 页。

- ### 缓存（Cache）
    - [ ] LRU 缓存:
        - [ ] [LRU 的魔力 (100 Days of Google Dev) (video)](https://www.youtube.com/watch?v=R5ON3iwx78M)
        - [ ] [实现 LRU (video)](https://www.youtube.com/watch?v=bq6N7Ym81iI)
        - [ ] [LeetCode - 146 LRU Cache (C++) (video)](https://www.youtube.com/watch?v=8-FZRAjR7qU)
    - [ ] CPU 缓存:
        - [ ] [MIT 6.004 L15: 存储体系 (video)](https://www.youtube.com/watch?v=vjYF_fAZI5E&list=PLrRW1w6CGAcXbMtDFj205vALOGmiRc82-&index=24)
        - [ ] [MIT 6.004 L16: 缓存的问题 (video)](https://www.youtube.com/watch?v=ajgC3-pyGlk&index=25&list=PLrRW1w6CGAcXbMtDFj205vALOGmiRc82-)

- ### 进程（Processe）和线程（Thread）
    - [ ] 计算机科学 162 - 操作系统 (25 个视频):
        - 视频 1-11 是关于进程和线程
        - [操作系统和系统编程 (video)](https://www.youtube.com/playlist?list=PL-XXv-cvA_iBDyz-ba4yDskqMDY6A1w_c)
    - [进程和线程的区别是什么?](https://www.quora.com/What-is-the-difference-between-a-process-and-a-thread)
    - 涵盖了:
        - 进程、线程、协程
            - 进程和线程的区别
            - 进程
            - 线程
            - 锁
            - 互斥
            - 信号量
            - 监控
            - 他们是如何工作的
            - 死锁
            - 活锁
        - CPU 活动, 中断, 上下文切换
        - 现代多核处理器的并发式结构
        - 进程资源需要（内存：代码、静态存储器、栈、堆、文件描述符、I/O）
        - 线程资源需要（在同一个进程内和其他线程共享以上的资源，但是每个线程都有独立的程序计数器、栈计数器、寄存器和栈）
        - Fork 操作是真正的写时复制（只读），直到新的进程写到内存中，才会生成一份新的拷贝。
        - 上下文切换
            - 操作系统和底层硬件是如何初始化上下文切换的。
    - [ ] [C++ 的线程 (系列 - 10 个视频)](https://www.youtube.com/playlist?list=PL5jc9xFGsL8E12so1wlMS0r0hTQoJL74M)
    - [ ] Python 的协程 (视频):
        - [ ] [线程系列](https://www.youtube.com/playlist?list=PL1H1sBF1VAKVMONJWJkmUh6_p8g4F2oy1)
        - [ ] [Python 线程](https://www.youtube.com/watch?v=Bs7vPNbB9JM)
        - [ ] [理解 Python 的 GIL (2010)](https://www.youtube.com/watch?v=Obt-vMVdM8s)
            - [参考](http://www.dabeaz.com/GIL)
        - [ ] [David Beazley - Python 协程 - PyCon 2015](https://www.youtube.com/watch?v=MCs5OvhV9S4)
        - [ ] [Keynote David Beazley - 兴趣主题 (Python 异步 I/O)](https://www.youtube.com/watch?v=ZzfHjytDceU)
        - [ ] [Python 中的互斥](https://www.youtube.com/watch?v=0zaPs8OtyKY)


    系统设计以及可伸缩性，要把软硬件的伸缩性设计的足够好有很多的东西要考虑，所以这是个包含非常多内容和资源的大主题。需要花费相当多的时间在这个主题上。

- ### 系统设计、可伸缩性、数据处理
    - Yegge 的注意事项:
        - 伸缩性
            - 把大数据集提取为单一值
            - 大数据集转换
            - 处理大量的数据集
        - 系统
            - 特征集
            - 接口
            - 类层次结构
            - 在特定的约束下设计系统
            - 轻量和健壮性
            - 权衡和折衷
            - 性能分析和优化
    - [ ] **从这里开始**: [HiredInTech：系统设计](http://www.hiredintech.com/system-design/)
    - [ ] [该如何为技术面试里设计方面的问题做准备?](https://www.quora.com/How-do-I-prepare-to-answer-design-questions-in-a-technical-interview?redirected_qid=1500023)
    - [ ] [在系统设计面试前必须知道的 8 件事](http://blog.gainlo.co/index.php/2015/10/22/8-things-you-need-to-know-before-system-design-interviews/)
    - [ ] [算法设计](http://www.hiredintech.com/algorithm-design/)
    - [ ] [数据库范式 - 1NF, 2NF, 3NF and 4NF (video)](https://www.youtube.com/watch?v=UrYLYV7WSHM)
    - [ ] [系统设计面试](https://github.com/checkcheckzz/system-design-interview) - 这一部分有很多的资源，浏览一下我放在下面的文章和例子。
    - [ ] [如何在系统设计面试中脱颖而出](http://www.palantir.com/2011/10/how-to-rock-a-systems-design-interview/)
    - [ ] [每个人都该知道的一些数字](http://everythingisdata.wordpress.com/2009/10/17/numbers-everyone-should-know/)
    - [ ] [上下文切换操作会耗费多少时间?](http://blog.tsunanet.net/2010/11/how-long-does-it-take-to-make-context.html)
    - [ ] [跨数据中心的事务 (video)](https://www.youtube.com/watch?v=srOgpXECblk)
    - [ ] [简明 CAP 理论介绍](http://ksat.me/a-plain-english-introduction-to-cap-theorem/)
    - [ ] Paxos 一致性算法:
        - [时间很短](https://www.youtube.com/watch?v=s8JqcZtvnsM)
        - [用例 和 multi-paxos](https://www.youtube.com/watch?v=JEpsBg0AO6o)
        - [论文](http://research.microsoft.com/en-us/um/people/lamport/pubs/paxos-simple.pdf)
    - [ ] [一致性哈希](http://www.tom-e-white.com/2007/11/consistent-hashing.html)
    - [ ] [NoSQL 模式](http://horicky.blogspot.com/2009/11/nosql-patterns.html)
    - [ ] [OOSE: UML 2.0 系列 (video)](https://www.youtube.com/watch?v=OkC7HKtiZC0&list=PLGLfVvz_LVvQ5G-LdJ8RLqe-ndo7QITYc)
    - [ ] OOSE: 使用 UML 和 Java 开发软件 (21 videos):
        - 如果你对 OO 都深刻的理解和实践，可以跳过这部分。
        - [OOSE: 使用 UML 和 Java 开发软件](https://www.youtube.com/playlist?list=PLJ9pm_Rc9HesnkwKlal_buSIHA-jTZMpO)
    - [ ] 面向对象编程的 SOLID 原则:
        - [ ] [Bob Martin 面向对象的 SOLID 原则和敏捷设计 (video)](https://www.youtube.com/watch?v=TMuno5RZNeE)
        - [ ] [C# SOLID 设计模式 (video)](https://www.youtube.com/playlist?list=PL8m4NUhTQU48oiGCSgCP1FiJEcg_xJzyQ)
        - [ ] [SOLID 原则 (video)](https://www.youtube.com/playlist?list=PL4CE9F710017EA77A)
        - [ ] S - [单一职责原则](http://www.oodesign.com/single-responsibility-principle.html) | [每个对象的单一职责](http://www.javacodegeeks.com/2011/11/solid-single-responsibility-principle.html)
            - [更多](https://docs.google.com/open?id=0ByOwmqah_nuGNHEtcU5OekdDMkk)
        - [ ] O - [开闭原则](http://www.oodesign.com/open-close-principle.html)  | [生产环境里的对象应该为扩展做准备而不是为更改](https://en.wikipedia.org/wiki/Open/closed_principle)
            - [更多](http://docs.google.com/a/cleancoder.com/viewer?a=v&pid=explorer&chrome=true&srcid=0BwhCYaYDn8EgN2M5MTkwM2EtNWFkZC00ZTI3LWFjZTUtNTFhZGZiYmUzODc1&hl=en)
        - [ ] L - [里氏代换原则](http://www.oodesign.com/liskov-s-substitution-principle.html) | [基类和继承类遵循 ‘IS A’ 原则](http://stackoverflow.com/questions/56860/what-is-the-liskov-substitution-principle)
            - [更多](http://docs.google.com/a/cleancoder.com/viewer?a=v&pid=explorer&chrome=true&srcid=0BwhCYaYDn8EgNzAzZjA5ZmItNjU3NS00MzQ5LTkwYjMtMDJhNDU5ZTM0MTlh&hl=en)
        - [ ] I - [接口隔离原则](http://www.oodesign.com/interface-segregation-principle.html) | 客户端被迫实现用不到的接口
            - [5 分钟讲解接口隔离原则 (video)](https://www.youtube.com/watch?v=3CtAfl7aXAQ)
            - [更多](http://docs.google.com/a/cleancoder.com/viewer?a=v&pid=explorer&chrome=true&srcid=0BwhCYaYDn8EgOTViYjJhYzMtMzYxMC00MzFjLWJjMzYtOGJiMDc5N2JkYmJi&hl=en)
        - [ ] D -[依赖反转原则](http://www.oodesign.com/dependency-inversion-principle.html) | 减少对象里的依赖。
            - [什么是依赖倒置以及它为什么重要](http://stackoverflow.com/questions/62539/what-is-the-dependency-inversion-principle-and-why-is-it-important)
            - [更多](http://docs.google.com/a/cleancoder.com/viewer?a=v&pid=explorer&chrome=true&srcid=0BwhCYaYDn8EgMjdlMWIzNGUtZTQ0NC00ZjQ5LTkwYzQtZjRhMDRlNTQ3ZGMz&hl=en)
    - [ ] 可伸缩性:
        - [ ] [很棒的概述 (video)](https://www.youtube.com/watch?v=-W9F__D3oY4)
        - [ ] 简短系列:
            - [克隆](http://www.lecloud.net/post/7295452622/scalability-for-dummies-part-1-clones)
            - [数据库](http://www.lecloud.net/post/7994751381/scalability-for-dummies-part-2-database)
            - [缓存](http://www.lecloud.net/post/9246290032/scalability-for-dummies-part-3-cache)
            - [异步](http://www.lecloud.net/post/9699762917/scalability-for-dummies-part-4-asynchronism)
        - [ ] [可伸缩的 Web 架构和分布式系统](http://www.aosabook.org/en/distsys.html)
        - [ ] [错误的分布式系统解释](https://pages.cs.wisc.edu/~zuyu/files/fallacies.pdf)
        - [ ] [实用编程技术](http://horicky.blogspot.com/2010/10/scalable-system-design-patterns.html)
            - [extra: Google Pregel 图形处理](http://horicky.blogspot.com/2010/07/google-pregel-graph-processing.html)
        - [ ] [Jeff Dean - 在 Goolge 构建软件系统 (video)](https://www.youtube.com/watch?v=modXC5IWTJI)
        - [ ] [可伸缩系统架构设计介绍](http://lethain.com/introduction-to-architecting-systems-for-scale/)
        - [ ] [使用 App Engine 和云存储扩展面向全球用户的手机游戏架构实践(video)](https://www.youtube.com/watch?v=9nWyWwY2Onc)
        - [ ] [How Google Does Planet-Scale Engineering for Planet-Scale Infra (video)](https://www.youtube.com/watch?v=H4vMcD7zKM0)
        - [ ] [算法的重要性](https://www.topcoder.com/community/data-science/data-science-tutorials/the-importance-of-algorithms/)
        - [ ] [分片](http://highscalability.com/blog/2009/8/6/an-unorthodox-approach-to-database-design-the-coming-of-the.html)
        - [ ] [Facebook 系统规模扩展实践 (2009)](https://www.infoq.com/presentations/Scale-at-Facebook)
        - [ ] [Facebook 系统规模扩展实践 (2012), "为 10 亿用户构建" (video)](https://www.youtube.com/watch?v=oodS71YtkGU)
        - [ ] [Long Game 工程实践 - Astrid Atkinson Keynote(video)](https://www.youtube.com/watch?v=p0jGmgIrf_M&list=PLRXxvay_m8gqVlExPC5DG3TGWJTaBgqSA&index=4)
        - [ ] [30 分钟看完 YouTuBe 7 年系统扩展经验](http://highscalability.com/blog/2012/3/26/7-years-of-youtube-scalability-lessons-in-30-minutes.html)
            - [video](https://www.youtube.com/watch?v=G-lGCC4KKok)
        - [ ] [PayPal 如何用 8 台虚拟机扛住 10 亿日交易量系统](http://highscalability.com/blog/2016/8/15/how-paypal-scaled-to-billions-of-transactions-daily-using-ju.html)
        - [ ] [如何对大数据集去重](https://blog.clevertap.com/how-to-remove-duplicates-in-large-datasets/)
        - [ ] [Etsy 的扩展和工程文化探究 Jon Cowie (video)](https://www.youtube.com/watch?v=3vV4YiqKm1o)
        - [ ] [是什么造就了 Amazon 自己的微服务架构](http://thenewstack.io/led-amazon-microservices-architecture/)
        - [ ] [压缩还是不压缩，是 Uber 面临的问题](https://eng.uber.com/trip-data-squeeze/)
        - [ ] [异步 I/O Tarantool 队列](http://highscalability.com/blog/2016/3/3/asyncio-tarantool-queue-get-in-the-queue.html)
        - [ ] [什么时候应该用近视查询处理?](http://highscalability.com/blog/2016/2/25/when-should-approximate-query-processing-be-used.html)
        - [ ] [Google 从单数据中心到故障转移, 到本地多宿主架构的演变]( http://highscalability.com/blog/2016/2/23/googles-transition-from-single-datacenter-to-failover-to-a-n.html)
        - [ ] [Spanner](http://highscalability.com/blog/2012/9/24/google-spanners-most-surprising-revelation-nosql-is-out-and.html)
        - [ ] [Egnyte: 构建和扩展 PB 级分布式系统架构的经验教训](http://highscalability.com/blog/2016/2/15/egnyte-architecture-lessons-learned-in-building-and-scaling.html)
        - [ ] [机器学习驱动的编程: 新世界的新编程方式](http://highscalability.com/blog/2016/7/6/machine-learning-driven-programming-a-new-programming-for-a.html)
        - [ ] [日服务数百万请求的图像优化技术](http://highscalability.com/blog/2016/6/15/the-image-optimization-technology-that-serves-millions-of-re.html)
        - [ ] [Patreon 架构](http://highscalability.com/blog/2016/2/1/a-patreon-architecture-short.html)
        - [ ] [Tinder: 推荐引擎是如何决定下一个你将会看到谁的?](http://highscalability.com/blog/2016/1/27/tinder-how-does-one-of-the-largest-recommendation-engines-de.html)
        - [ ] [现代缓存设计](http://highscalability.com/blog/2016/1/25/design-of-a-modern-cache.html)
        - [ ] [Facebook 实时视频流扩展](http://highscalability.com/blog/2016/1/13/live-video-streaming-at-facebook-scale.html)
        - [ ] [在 Amazon AWS 上把服务扩展到 1100 万量级的新手教程](http://highscalability.com/blog/2016/1/11/a-beginners-guide-to-scaling-to-11-million-users-on-amazons.html)
        - [ ] [对延时敏感的应用是否应该使用 Docker?](http://highscalability.com/blog/2015/12/16/how-does-the-use-of-docker-effect-latency.html)
        - [ ] [AMP（Accelerated Mobile Pages）的存在是对 Google 的威胁么?](http://highscalability.com/blog/2015/12/14/does-amp-counter-an-existential-threat-to-google.html)
        - [ ] [360 度解读 Netflix 技术栈](http://highscalability.com/blog/2015/11/9/a-360-degree-view-of-the-entire-netflix-stack.html)
        - [ ] [延迟无处不在 - 如何搞定它？](http://highscalability.com/latency-everywhere-and-it-costs-you-sales-how-crush-it)
        - [ ] [无服务器架构](http://martinfowler.com/articles/serverless.html)
        - [ ] [是什么驱动着 Instagram: 上百个实例、几十种技术](http://instagram-engineering.tumblr.com/post/13649370142/what-powers-instagram-hundreds-of-instances)
        - [ ] [Cinchcast 架构 - 每天处理 1500 小时的音频](http://highscalability.com/blog/2012/7/16/cinchcast-architecture-producing-1500-hours-of-audio-every-d.html)
        - [ ] [Justin.Tv 实时视频播放架构](http://highscalability.com/blog/2010/3/16/justintvs-live-video-broadcasting-architecture.html)
        - [ ] [Playfish's 社交游戏架构 - 每月五千万用户增长](http://highscalability.com/blog/2010/9/21/playfishs-social-gaming-architecture-50-million-monthly-user.html)
        - [ ] [猫途鹰架构 - 40 万访客, 200 万动态页面访问, 30TB 数据](http://highscalability.com/blog/2011/6/27/tripadvisor-architecture-40m-visitors-200m-dynamic-page-view.html)
        - [ ] [PlentyOfFish 架构](http://highscalability.com/plentyoffish-architecture)
        - [ ] [Salesforce 架构 - 如何扛住 13 亿日交易量](http://highscalability.com/blog/2013/9/23/salesforce-architecture-how-they-handle-13-billion-transacti.html)
        - [ ] [ESPN's 架构扩展](http://highscalability.com/blog/2013/11/4/espns-architecture-at-scale-operating-at-100000-duh-nuh-nuhs.html)
        - [ ] 下面 『消息、序列化和消息系统』部分的内容会提到什么样的技术能把各种服务整合到一起
        - [ ] Twitter:
            - [O'Reilly MySQL CE 2011: Jeremy Cole, "Big and Small Data at @Twitter" (video)](https://www.youtube.com/watch?v=5cKTP36HVgI)
            - [时间线的扩展](https://www.infoq.com/presentations/Twitter-Timeline-Scalability)
        - 更多内容可以查看视频部分的『大规模数据挖掘』视频系列。
    - [ ] 系统设计问题练习：下面有一些指导原则，每一个都有相关文档以及在现实中该如何处理。
        - 复习: [HiredInTech 的系统设计](http://www.hiredintech.com/system-design/)
        - [cheat sheet](https://github.com/jwasham/coding-interview-university/blob/master/extras/cheat%20sheets/system-design.pdf)
        - 流程:
            1. 理解问题和范围:
                - 在面试官的帮助下定义用例
                - 提出附加功能的建议
                - 去掉面试官认定范围以外的内容
                - 假定高可用是必须的，而且要作为一个用例
            2. 考虑约束:
                - 问一下每月请求量
                - 问一下每秒请求量 (他们可能会主动提到或者让你算一下)
                - 评估读写所占的百分比
                - 评估的时候牢记 2/8 原则
                - 每秒写多少数据
                - 总的数据存储量要考虑超过 5 年的情况
                - 每秒读多少数据
            3. 抽象设计:
                - 分层 (服务, 数据, 缓存)
                - 基础设施: 负载均衡, 消息
                - 粗略的概括任何驱动整个服务的关键算法
                - 考虑瓶颈并指出解决方案
        - 练习:
            - [设计一个 CDN 网络](http://repository.cmu.edu/cgi/viewcontent.cgi?article=2112&context=compsci)
            - [设计一个随机唯一 ID 生成系统](https://blog.twitter.com/2010/announcing-snowflake)
            - [设计一个在线多人卡牌游戏](http://www.indieflashblog.com/how-to-create-an-asynchronous-multiplayer-game.html)
            - [设计一个 key-value 数据库](http://www.slideshare.net/dvirsky/introduction-to-redis)
            - [设计一个函数获取过去某个时间段内前 K 个最高频访问的请求]( https://icmi.cs.ucsb.edu/research/tech_reports/reports/2005-23.pdf)
            - [设计一个图片分享系统](http://highscalability.com/blog/2011/12/6/instagram-architecture-14-million-users-terabytes-of-photos.html)
            - [设计一个推荐系统](http://ijcai13.org/files/tutorial_slides/td3.pdf)
            - [设计一个短域名生成系统](http://www.hiredintech.com/system-design/the-system-design-process/)
            - [设计一个缓存系统](https://www.adayinthelifeof.nl/2011/02/06/memcache-internals/)

- ### 论文
    - 有 Google 的论文和一些知名的论文.
    - 你很可能实在没时间一篇篇完整的读完他们。我建议可以有选择的读其中一些论文里的核心部分。
    - [ ] [1978: 通信顺序处理](http://spinroot.com/courses/summer/Papers/hoare_1978.pdf)
        - [Go 实现](https://godoc.org/github.com/thomas11/csp)
        - [喜欢经典的论文?](https://www.cs.cmu.edu/~crary/819-f09/)
    - [ ] [2003: The Google 文件系统](http://static.googleusercontent.com/media/research.google.com/en//archive/gfs-sosp2003.pdf)
        - 2012 年被 Colossus 取代了
    - [ ] [2004: MapReduce: Simplified Data Processing on Large Clusters]( http://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf)
        - 大多被云数据流取代了?
    - [ ] [2007: 每个程序员都应该知道的内存知识 (非常长，作者建议跳过某些章节来阅读)](https://www.akkadia.org/drepper/cpumemory.pdf)
    - [ ] [2012: Google 的 Colossus](https://www.wired.com/2012/07/google-colossus/)
        - 没有论文
    - [ ] 2012: AddressSanitizer: 快速的内存访问检查器:
        - [论文](http://static.googleusercontent.com/media/research.google.com/en//pubs/archive/37752.pdf)
        - [视频](https://www.usenix.org/conference/atc12/technical-sessions/presentation/serebryany)
    - [ ] 2013: Spanner: Google 的分布式数据库:
        - [论文](http://static.googleusercontent.com/media/research.google.com/en//archive/spanner-osdi2012.pdf)
        - [视频](https://www.usenix.org/node/170855)
    - [ ] [2014: Machine Learning: The High-Interest Credit Card of Technical Debt](http://static.googleusercontent.com/media/research.google.com/en//pubs/archive/43146.pdf)
    - [ ] [2015: Continuous Pipelines at Google](http://static.googleusercontent.com/media/research.google.com/en//pubs/archive/43790.pdf)
    - [ ] [2015: 大规模高可用: 构建 Google Ads 的数据基础设施](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/44686.pdf)
    - [ ] [2015: TensorFlow: 异构分布式系统上的大规模机器学习](http://download.tensorflow.org/paper/whitepaper2015.pdf )
    - [ ] [2015: 开发者应该如何搜索代码：用例学习](http://static.googleusercontent.com/media/research.google.com/en//pubs/archive/43835.pdf)
    - [ ] [2016: Borg, Omega, and Kubernetes](http://static.googleusercontent.com/media/research.google.com/en//pubs/archive/44843.pdf)

- ### 测试
    - 涵盖了:
        - 单元测试是如何工作的
        - 什么是模拟对象
        - 什么是集成测试
        - 什么是依赖注入
    - [ ] [James Bach 讲敏捷软件测试 (video)](https://www.youtube.com/watch?v=SAhJf36_u5U)
    - [ ] [James Bach 软件测试公开课 (video)](https://www.youtube.com/watch?v=ILkT_HV9DVU)
    - [ ] [Steve Freeman - 测试驱动的开发 (video)](https://vimeo.com/83960706)
        - [slides](http://gotocon.com/dl/goto-berlin-2013/slides/SteveFreeman_TestDrivenDevelopmentThatsNotWhatWeMeant.pdf)
    - [ ] [测试驱动的开发已死。测试不朽。](http://david.heinemeierhansson.com/2014/tdd-is-dead-long-live-testing.html)
    - [ ] [测试驱动的开发已死? (video)](https://www.youtube.com/watch?v=z9quxZsLcfo)
    - [ ] [视频系列 (152 个) - 并不都是必须 (video)](https://www.youtube.com/watch?v=nzJapzxH_rE&list=PLAwxTw4SYaPkWVHeC_8aSIbSxE_NXI76g)
    - [ ] [Python：测试驱动的 Web 开发](http://www.obeythetestinggoat.com/pages/book.html#toc)
    - [ ] 依赖注入:
        - [ ] [视频](https://www.youtube.com/watch?v=IKD2-MAkXyQ)
        - [ ] [测试之道](http://jasonpolites.github.io/tao-of-testing/ch3-1.1.html)
    - [ ] [如何编写测试](http://jasonpolites.github.io/tao-of-testing/ch4-1.1.html)

- ### 调度
    - 在操作系统中是如何运作的
    - 在操作系统部分的视频里有很多资料

- ### 实现系统例程
    - 理解你使用的系统 API 底层有什么
    - 你能自己实现它们么?

- ### 字符串搜索和操作
    - [ ] [文本的搜索模式 (video)](https://www.coursera.org/learn/data-structures/lecture/tAfHI/search-pattern-in-text)
    - [ ] Rabin-Karp (videos):
        - [Rabin Karps 算法](https://www.coursera.org/learn/data-structures/lecture/c0Qkw/rabin-karps-algorithm)
        - [预先计算的优化](https://www.coursera.org/learn/data-structures/lecture/nYrc8/optimization-precomputation)
        - [优化: 实现和分析](https://www.coursera.org/learn/data-structures/lecture/h4ZLc/optimization-implementation-and-analysis)
        - [Table Doubling, Karp-Rabin](https://www.youtube.com/watch?v=BRO7mVIFt08&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=9)
        - [滚动哈希](https://www.youtube.com/watch?v=w6nuXg0BISo&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=32)
    - [ ] Knuth-Morris-Pratt (KMP) 算法:
        - [Pratt 算法](https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm)
        - [教程: Knuth-Morris-Pratt (KMP) 字符串匹配算法](https://www.youtube.com/watch?v=2ogqPWJSftE)
    - [ ] Boyer–Moore 字符串搜索算法
        - [Boyer-Moore字符串搜索算法](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_string_search_algorithm)
        - [Boyer-Moore-Horspool 高级字符串搜索算法 (video)](https://www.youtube.com/watch?v=QDZpzctPf10)
    - [ ] [Coursera: 字符串的算法](https://www.coursera.org/learn/algorithms-on-strings/home/week/1)

---

## 终面

    这一部分有一些短视频，你可以快速的观看和复习大多数重要概念。
    这对经常性的巩固很有帮助。

#### 综述:

- [ ] 2-3 分钟的短视频系列 (23 个)
    - [Videos](https://www.youtube.com/watch?v=r4r1DZcx1cM&list=PLmVb1OknmNJuC5POdcDv5oCS7_OUkDgpj&index=22)
- [ ] 2-5 分钟的短视频系列 - Michael Sambol (18 个):
    - [Videos](https://www.youtube.com/channel/UCzDJwLWoYCUQowF_nG3m5OQ)

#### 排序:

- [ ] 归并排序: https://www.youtube.com/watch?v=GCae1WNvnZM


## 书籍

### Google Coaching 里提到的

**阅读并做练习:**

- [ ] 算法设计手册 (Skiena)
    - 书 (Kindle 上可以租到):
        - [Algorithm Design Manual](http://www.amazon.com/Algorithm-Design-Manual-Steven-Skiena/dp/1849967202)
    - Half.com 是一个资源丰富且性价比很高的在线书店.
    - 答案:
        - [解答](http://www.algorithm.cs.sunysb.edu/algowiki/index.php/The_Algorithms_Design_Manual_(Second_Edition))
        - [解答](http://blog.panictank.net/category/algorithmndesignmanualsolutions/page/2/)
    - [勘误表](http://www3.cs.stonybrook.edu/~skiena/algorist/book/errata)

    read and do exercises from the books below. Then move to coding challenges (further down below)
    一旦你理解了每日计划里的所有内容，就去读上面所列的书并完成练习，然后开始读下面所列的书并做练习，之后就可以开始实战写代码了（本文再往后的部分）

**首先阅读:**
- [ ] [Programming Interviews Exposed: Secrets to Landing Your Next Job, 2nd Edition](http://www.wiley.com/WileyCDA/WileyTitle/productCd-047012167X.html)

**然后阅读 (这本获得了很多推荐， 但是不在 Google coaching 的文档里):**
- [ ] [Cracking the Coding Interview, 6th Edition](http://www.amazon.com/Cracking-Coding-Interview-6th-Programming/dp/0984782850/)
    - 如果你看到有人在看 "The Google Resume", 实际上它和 "Cracking the Coding Interview" 是同一个作者写的，而且后者是升级版。

### 附加书单

这些没有被 Google 推荐阅读，不过我因为需要这些背景知识所以也把它们列在了这里。

- [ ] C Programming Language, Vol 2
    - [练习的答案](https://github.com/lekkas/c-algorithms)

- [ ] C++ Primer Plus, 6th Edition

- [ ] [《Unix 环境高级编程》 The Unix Programming Environment](http://product.half.ebay.com/The-UNIX-Programming-Environment-by-Brian-W-Kernighan-and-Rob-Pike-1983-Other/54385&tg=info)

- [ ] [《编程珠玑》 Programming Pearls](http://www.amazon.com/Programming-Pearls-2nd-Jon-Bentley/dp/0201657880)

- [ ] [Algorithms and Programming: Problems and Solutions](http://www.amazon.com/Algorithms-Programming-Solutions-Alexander-Shen/dp/0817638474)

### 如果你有时间

- [ ] [Introduction to Algorithms](https://www.amazon.com/Introduction-Algorithms-3rd-MIT-Press/dp/0262033844)

- [ ] [Elements of Programming Interviews](https://www.amazon.com/Elements-Programming-Interviews-Insiders-Guide/dp/1479274836)
    - 如果你希望在面试里用 C++ 写代码，这本书的代码全都是 C++ 写的
    - 通常情况下能找到解决方案的好书.

## 编码练习和挑战

一旦你学会了理论基础，就应该把它们拿出来练练。
尽量坚持每天做编码练习，越多越好。

编程问题预备:

- [ ] [不错的介绍 (摘自 System Design 章节): 算法设计:](http://www.hiredintech.com/algorithm-design/)
- [ ] [如何找到解决方案](https://www.topcoder.com/community/data-science/data-science-tutorials/how-to-find-a-solution/)
- [ ] [如何剖析 Topcoder 题目描述](https://www.topcoder.com/community/data-science/data-science-tutorials/how-to-dissect-a-topcoder-problem-statement/)
- [ ] [Topcoders 里用到的数学](https://www.topcoder.com/community/data-science/data-science-tutorials/mathematics-for-topcoders/)
- [ ] [动态规划 – 从入门到精通](https://www.topcoder.com/community/data-science/data-science-tutorials/dynamic-programming-from-novice-to-advanced/)

- [MIT 面试材料](https://courses.csail.mit.edu/iap/interview/materials.php)

- [针对编程语言本身的练习](http://exercism.io/languages)

编码练习平台:

- [LeetCode](https://leetcode.com/)
- [TopCoder](https://www.topcoder.com/)
- [Project Euler (数学方向为主)](https://projecteuler.net/index.php?section=problems)
- [Codewars](http://www.codewars.com)
- [HackerRank](https://www.hackerrank.com/)
- [Codility](https://codility.com/programmers/)
- [InterviewCake](https://www.interviewcake.com/)
- [InterviewBit](https://www.interviewbit.com/invite/icjf)

- [模拟大公司的面试](http://www.gainlo.co/)

## 当你临近面试时

- [ ] 搞定代码面试 (videos):
    - [Cracking The Code Interview](https://www.youtube.com/watch?v=4NIb9l3imAo)
    - [Cracking the Coding Interview - 全栈系列](https://www.youtube.com/watch?v=Eg5-tdAwclo)
    - [Ask Me Anything: Gayle Laakmann McDowell (Cracking the Coding Interview 的作者)](https://www.youtube.com/watch?v=1fqxMuPmGak)

## 你的简历

- [10 条小贴士让你写出一份还算不错的简历](http://steve-yegge.blogspot.co.uk/2007_09_01_archive.html)
- 这是搞定面试的第一个关键步骤


## 当面试来临的时候

    随着下面列举的问题思考下你可能会遇到的 20 个面试问题
    每个问题准备 2-3 种回答
    准备点故事，不要只是摆一些你完成的事情的数据，相信我，人人都喜欢听故事

- 你为什么想得到这份工作？
- 你解决过的最有难度的问题是什么？
- 面对过的最大挑战是什么?
- 见过的最好或者最坏的设计是怎么样的?
- 对某项 Google 产品提出改进建议。
- 你作为一个个体同时也是团队的一员，如何达到最好的工作状态?
- 你的什么技能或者经验是你的角色中不可或缺的?为什么？
- 你在某份工作或某个项目中最享受的是什么?
- 你在某份工作或某个项目中面临过的最大挑战是什么?
- 你在某份工作或某个项目中遇到过的最蛋疼的 Bug 是什么样的？
- 你在某份工作或某个项目中学到了什么？
- 你在某份工作或某个项目中哪些地方还可以做的更好？

## 问面试官的问题

    我会问的一些：(可能我已经知道了答案但我想听听面试官的看法或者了解团队的前景):

- 团队多大规模?
- 开发周期是怎样的? 会使用瀑布流/极限编程/敏捷开发么?
- 经常会为 deadline 加班么? 或者是有弹性的?
- 团队里怎么做技术选型?
- 每周平均开多少次会?
- 你觉得工作环境有助于员工集中精力吗?
- 目前正在做什么工作?
- 喜欢这些事情吗?
- 工作期限是怎么样的?

## 当你获得了梦想的职位

我还能说些什么呢，恭喜你！

- [我希望在 Google 的第一天就知道的 10 件事](https://medium.com/@moonstorming/10-things-i-wish-i-knew-on-my-first-day-at-google-107581d87286#.livxn7clw)

坚持继续学习。

得到这份工作只是一个开始。

---

    *****************************************************************************************************
    *****************************************************************************************************

    下面的内容都是可选的。这些是我的推荐，不是 Google 的。
    通过学习这些内容，你将会得到更多的有关 CS 的概念，并将为所有的软件工程工作做更好的准备。

    *****************************************************************************************************
    *****************************************************************************************************

---

## 附加的学习

- ### Unicode
    - [ ] [每一个软件开发者的绝对最低限度，必须要知道的关于 Unicode 和字符集知识]( http://www.joelonsoftware.com/articles/Unicode.html)
    - [ ] [关于处理文本需要的编码和字符集, 每个程序员绝对需要知道的知识](http://kunststube.net/encoding/)

- ### 字节顺序
    - [ ] [大、小端字节序](https://www.cs.umd.edu/class/sum2003/cmsc311/Notes/Data/endian.html)
    - [ ] [大端字节 Vs 小端字节(视频)](https://www.youtube.com/watch?v=JrNF0KRAlyo)
    - [ ] [大、小端字节序的里里外外(Big And Little Endian Inside/Out) (视频)](https://www.youtube.com/watch?v=oBSuXP-1Tc0)
        - 内核开发者的讨论非常技术性，如果大多数都超出了你的理解范围，不要太担心。
        - 前半段已经足够了。

- ### Emacs and vi(m)
    - Yegge 的建议，从一个很早以前的亚马逊招聘信息中而来：熟悉基于 unix 的代码编辑器
    - vi(m):
        - [使用 vim 进行编辑 01 - 安装, 设置和模式 (视频)](https://www.youtube.com/watch?v=5givLEMcINQ&index=1&list=PL13bz4SHGmRxlZVmWQ9DvXo1fEg4UdGkr)
        - [VIM 的冒险之旅](http://vim-adventures.com/)
        - 4 个视频集:
            - [vi/vim 编辑器 - 课程 1](https://www.youtube.com/watch?v=SI8TeVMX8pk)
            - [vi/vim 编辑器 - 课程 2](https://www.youtube.com/watch?v=F3OO7ZIOaJE)
            - [vi/vim 编辑器 - 课程 4](https://www.youtube.com/watch?v=1lYD5gwgZIA)
            - [vi/vim 编辑器 - 课程 3](https://www.youtube.com/watch?v=ZYEccA_nMaI)
        - [使用 Vi 而不是 Emacs](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#Using_Vi_instead_of_Emacs)
    - emacs:
        - [基础 Emacs 教程 (视频)](https://www.youtube.com/watch?v=hbmV1bnQ-i0)
        - 3 个视频集:
            - [Emacs 教程 (初学者) -第 1 部分- 文件命令, 剪切/复制/粘贴,  自定义命令](https://www.youtube.com/watch?v=ujODL7MD04Q)
            - [Emacs 教程 (初学者 -第 2 部分- Buffer 管理, 搜索, M-x grep 和 rgrep 模式](https://www.youtube.com/watch?v=XWpsRupJ4II)
            - [Emacs 教程 (初学者 -第 3 部分- 表达式, 声明, ~/.emacs 文件和包机制](https://www.youtube.com/watch?v=paSgzPso-yc)
        - [Evil 模式: 或许, 我是怎样对 Emacs 路人转粉的 (视频)](https://www.youtube.com/watch?v=JWD1Fpdd4Pc)
        - [使用 Emacs 开发 C 程序](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#Writing_C_programs_with_Emacs)
        - [(或许) 深度组织模式:管理结构 (视频)](https://www.youtube.com/watch?v=nsGYet02bEk)

- ### Unix 命令行工具
    - 下列内容中的优秀工具由的 Yegge 推荐，Yegge 目前致力于 Amazon 人事招聘处。
    - [ ] bash
    - [ ] cat
    - [ ] grep
    - [ ] sed
    - [ ] awk
    - [ ] curl or wget
    - [ ] sort
    - [ ] tr
    - [ ] uniq
    - [ ] [strace](https://en.wikipedia.org/wiki/Strace)
    - [ ] [tcpdump](https://danielmiessler.com/study/tcpdump/)

- ### 信息资源 (视频)
    - [ ] [Khan Academy 可汗学院](https://www.khanacademy.org/computing/computer-science/informationtheory)
    - [ ] 更多有关马尔可夫的内容:
        - [ ] [Core Markov Text Generation马尔可夫内容生成](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/waxgx/core-markov-text-generation)
        - [ ] [Core Implementing Markov Text Generation马尔可夫内容生成补充](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/gZhiC/core-implementing-markov-text-generation)
        - [ ] [Project = Markov Text Generation Walk Through一个马尔可夫内容生成器的项目](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/EUjrq/project-markov-text-generation-walk-through)
    - 关于更多信息，请参照下方 MIT 6.050J 信息和系统复杂度的内容.

- ### 奇偶校验位 & 汉明码 (视频)
    - [ ] [入门](https://www.youtube.com/watch?v=q-3BctoUpHE)
    - [ ] [奇偶校验位](https://www.youtube.com/watch?v=DdMcAUlxh1M)
    - [ ] 汉明码(Hamming Code):
        - [发现错误](https://www.youtube.com/watch?v=1A_NcXxdoCc)
        - [修正错误](https://www.youtube.com/watch?v=JAMLuxdHH8o)
    - [ ] [检查错误](https://www.youtube.com/watch?v=wbH2VxzmoZk)

- ### 系统熵值（系统复杂度）
    - 请参考下方视频
    - 观看之前，请先确定观看了信息论的视频
    - [ ] [信息理论, 克劳德·香农, 熵值, 系统冗余, 数据比特压缩 (视频)](https://youtu.be/JnJq3Py0dyM?t=176)

- ### 密码学
    - 请参考下方视频
    - 观看之前，请先确定观看了信息论的视频
    - [ ] [可汗学院](https://www.khanacademy.org/computing/computer-science/密码学)
    - [ ] [密码学: 哈希函数](https://www.youtube.com/watch?v=KqqOXndnvic&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp&index=30)
    - [ ] [密码学: 加密](https://www.youtube.com/watch?v=9TNI2wHmaeI&index=31&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)

- ### 压缩
    - 观看之前，请先确定观看了信息论的视频
    - [ ] 压缩 (视频):
        - [ ] [压缩](https://www.youtube.com/watch?v=Lto-ajuqW3w)
        - [ ] [压缩熵值](https://www.youtube.com/watch?v=M5c_RFKVkko)
        - [ ] [由上而下的树 (霍夫曼编码树)](https://www.youtube.com/watch?v=umTbivyJoiI)
        - [ ] [额外比特 - 霍夫曼编码树](https://www.youtube.com/watch?v=DV8efuB3h2g)
        - [ ] [优雅的压缩数据 (无损数据压缩方法)](https://www.youtube.com/watch?v=goOa3DGezUA)
        - [ ] [Text Compression Meets Probabilities](https://www.youtube.com/watch?v=cCDCfoHTsaU)
    - [ ] [数据压缩的艺术](https://www.youtube.com/playlist?list=PLOU2XLYxmsIJGErt5rrCqaSGTMyyqNt2H)
    - [ ] [(可选) 谷歌开发者: GZIP 还差远了呢!](https://www.youtube.com/watch?v=whGwm0Lky2s)

- ### 网络 (视频)
    - [ ] [可汗学院](https://www.khanacademy.org/computing/computer-science/internet-intro)
    - [ ] [网络传输协议中的数据压缩](https://www.youtube.com/watch?v=Vdc8TCESIg8)
    - [ ] [TCP/IP 和 OSI 模型解析!](https://www.youtube.com/watch?v=e5DEVa9eSN0)
    - [ ] [TCP/IP 教程：传输数据包.](https://www.youtube.com/watch?v=nomyRJehhnM)
    - [ ] [HTTP](https://www.youtube.com/watch?v=WGJrLqtX7As)
    - [ ] [SSL 和 HTTPS](https://www.youtube.com/watch?v=S2iBR2ZlZf0)
    - [ ] [SSL/TLS](https://www.youtube.com/watch?v=Rp3iZUvXWlM)
    - [ ] [HTTP 2.0](https://www.youtube.com/watch?v=E9FxNzv1Tr8)
    - [ ] [视频](https://www.youtube.com/playlist?list=PLEbnTDJUr_IegfoqO4iPnPYQui46QqT0j)
    - [ ] [子网络解密 - 第五部分 经典内部域名指向 CIDR 标记](https://www.youtube.com/watch?v=t5xYI0jzOf4)

- ### 计算机安全
    - [MIT](https://www.youtube.com/playlist?list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [威胁模型：入门](https://www.youtube.com/watch?v=GqmQg-cszw4&index=1&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [控制攻击](https://www.youtube.com/watch?v=6bwzNg5qQ0o&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh&index=2)
        - [ ] [缓冲数据注入和防御](https://www.youtube.com/watch?v=drQyrzRoRiA&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh&index=3)
        - [ ] [优先权区分](https://www.youtube.com/watch?v=6SIJmoE9L9g&index=4&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [能力](https://www.youtube.com/watch?v=8VqTSY-11F4&index=5&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [在沙盒中运行原生代码](https://www.youtube.com/watch?v=VEV74hwASeU&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh&index=6)
        - [ ] [网络安全模型](https://www.youtube.com/watch?v=chkFBigodIw&index=7&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [网络安全应用](https://www.youtube.com/watch?v=EBQIGy1ROLY&index=8&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [标志化执行](https://www.youtube.com/watch?v=yRVZPvHYHzw&index=9&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [网络安全](https://www.youtube.com/watch?v=SIEVvk3NVuk&index=11&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [网络协议](https://www.youtube.com/watch?v=QOtA76ga_fY&index=12&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)
        - [ ] [旁路攻击](https://www.youtube.com/watch?v=PuVMkSEcPiI&index=15&list=PLUl4u3cNGP62K2DjQLRxDNRi0z2IRWnNh)

- ### 释放缓存
    - [ ] [Java 释放缓存; 片段化数据 (视频)](https://www.youtube.com/watch?v=StdfeXaKGEc&list=PL-XXv-cvA_iAlnI-BQr9hjqADPBtujFJd&index=25)
    - [ ] [编译器 (视频)](https://www.youtube.com/playlist?list=PLO9y7hOkmmSGTy5z6HZ-W4k2y8WXF7Bff)
    - [ ] [Python 释放缓存 (视频)](https://www.youtube.com/watch?v=iHVs_HkjdmI)
    - [ ] [深度解析：论释放缓存在 JAVA 中的重要性](https://www.infoq.com/presentations/garbage-collection-benefits)
    - [ ] [深度解析：论释放缓存在 Python 中的重要性(视频)](https://www.youtube.com/watch?v=P-8Z0-MhdQs&list=PLdzf4Clw0VbOEWOS_sLhT_9zaiQDrS5AR&index=3)

- ### 并行/并发编程
    - [ ] [Coursera (Scala)](https://www.coursera.org/learn/parprog1/home/week/1)
    - [ ] [论并行/并发编程如何提高 Python 执行效率 (视频)](https://www.youtube.com/watch?v=uY85GkaYzBk)

- ### 设计模式
    - [ ] [UML统一建模语言概览 (视频)](https://www.youtube.com/watch?v=3cmzqZzwNDM&list=PLGLfVvz_LVvQ5G-LdJ8RLqe-ndo7QITYc&index=3)
    - [ ] 主要有如下的设计模式:
        - [ ] s(strategy)
        - [ ] singleton
        - [ ] adapter
        - [ ] prototype
        - [ ] decorator
        - [ ] visitor
        - [ ] factory, abstract factory
        - [ ] facade
        - [ ] observer
        - [ ] proxy
        - [ ] delegate
        - [ ] command
        - [ ] state
        - [ ] memento
        - [ ] 