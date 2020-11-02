"+++
title = "0530. Minimum Absolute Difference in BST 简单直接打败94%~ "
author = ["yi-li-dan-lu-fen"]
date = 2020-07-24T15:57:59+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 简单直接打败94%~

> [0530. Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
> [简单直接打败94%~](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/solution/jian-dan-zhi-jie-da-bai-94-by-yi-li-dan-lu-fen/) by [yi-li-dan-lu-fen](https://leetcode-cn.com/u/yi-li-dan-lu-fen/)

### 解题思路
大佬们写中序遍历都写烂了。那么针对非搜索树该如何做呢？
楼主的思路如下：首先利用模拟栈来统计所有节点的值，然后对所有值排序，最后只需计算有序数组相邻元素之差的绝对值，取最小值返回即可。
![image.png](https://pic.leetcode-cn.com/4e39e98e35fc10291bf378708b08b30bce664ffcdfafa4e9942d4656ffa6d587-image.png)

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def getMinimumDifference(self, root: TreeNode) -> int:
        stack = [root]
        vals = []
        while stack:
            tmp = stack.pop()
            vals.append(tmp.val)
            if tmp.left:
                stack.append(tmp.left) 
            if tmp.right:
                stack.append(tmp.right)
        vals.sort()
        return min([ abs(vals[i]-vals[i+1]) for i in range(0,len(vals)-1) ])
```