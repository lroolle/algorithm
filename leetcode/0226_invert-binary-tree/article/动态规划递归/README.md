"+++
title = "0226. Invert Binary Tree 动态规划+递归 "
author = ["xu-chen-chen-d"]
date = 2020-08-20T10:52:05+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Tree", "Recursion", "DynamicProgramming"]
categories = ["leetcode"]
draft = false
+++

# 动态规划+递归

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [动态规划+递归](https://leetcode-cn.com/problems/invert-binary-tree/solution/dong-tai-gui-hua-di-gui-by-xu-chen-chen-d/) by [xu-chen-chen-d](https://leetcode-cn.com/u/xu-chen-chen-d/)

### 解题思路
此处撰写解题思路
简单的动态规划+递归。
### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def invertTree(self, root: TreeNode) -> TreeNode:
        while root:
            temp = root.left
            root.left = root.right
            root.right = temp
            self.invertTree(root.left)
            self.invertTree(root.right)
            return root

```