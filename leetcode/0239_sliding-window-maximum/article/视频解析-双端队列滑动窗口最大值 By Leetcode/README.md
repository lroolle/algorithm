"+++
title = "0239. Sliding Window Maximum 🎦【视频解析】 双端队列滑动窗口最大值 "
author = ["LeetCode"]
date = 2019-08-13T10:26:24+08:00
tags = ["Leetcode", "Algorithms", "SlidingWindow"]
categories = ["leetcode"]
draft = false
+++

# 🎦【视频解析】 双端队列滑动窗口最大值

> [0239. Sliding Window Maximum](https://leetcode-cn.com/problems/sliding-window-maximum/)
> [🎦【视频解析】 双端队列滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/solution/shi-pin-jie-xi-shuang-duan-dui-lie-hua-dong-chuang/) by [LeetCode](https://leetcode-cn.com/u/leetcode/)

双端队列和普通队列最大的不同在于，它允许我们在队列的头尾两端都能在 $O(1)$ 的时间内进行数据的查看、添加和删除。

与队列相似，我们可以利用一个双链表实现双端队列。双端队列最常用的地方就是实现一个长度动态变化的窗口或者连续区间，而动态窗口这种数据结构在很多题目里都有运用。

![...- 常用数据结构和技巧 copy.mov](200d9cd5-4ed7-433a-af37-187b3a73f3d1)

这道题而言，既然每次都要在一个移动的窗口中找到最大值，那很简单，我们就移动这个窗口，然后扫描一遍窗口获得最大值。假设数组里有 $n$ 个元素，这样的算法复杂度就是 $O(n * k)$。

#### 那么我们能不能在移动窗口的过程中，更快地获得最大值呢？

可以利用一个双端队列来表示这个窗口。这个双端队列保存当前窗口中最大那个数的下标，双端队列新的头总是当前窗口中最大的那个数。
同时，有了这个下标，我们可以很快地知道新的窗口是否已经不再包含原来那个最大的数，如果不再包含，我们就把旧的数从双端队列的头删除。按照这样的操作，不管窗口的长度是多长，因为数组里的每个数都分别被压入和弹出双端队列一次，所以我们可以在 $O(n)$ 的时间里完成任务。
