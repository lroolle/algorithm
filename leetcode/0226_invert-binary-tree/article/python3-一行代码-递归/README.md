"+++
title = "0226. Invert Binary Tree python3 一行代码 递归 "
author = ["ting-ting-28"]
date = 2020-07-14T04:11:37+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Recursion", "Python"]
categories = ["leetcode"]
draft = false
+++

# python3 一行代码 递归

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [python3 一行代码 递归](https://leetcode-cn.com/problems/invert-binary-tree/solution/python3-yi-xing-dai-ma-di-gui-by-ting-ting-28-2/) by [ting-ting-28](https://leetcode-cn.com/u/ting-ting-28/)

# 解法一：递归——新定义二叉树
1. 停止条件：
    - 如果`root`等于`None`则返回空。
2. 返回的二叉树。
    - `val`: `root.val`；
    - `left`: 反转后的`root.right`；
    - `right`: 反转后的`root.left`；
```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def invertTree(self, root: TreeNode) -> TreeNode:
        return None if not root else TreeNode(root.val, self.invertTree(root.right), self.invertTree(root.left))
```
- - -
# 解法二：递归——修改`root`
1. 停止条件：
    - 如果`root`等于`None`则返回空。
2. 交换：
    - 左右交换；
3. 返回：
    - 返回根。
```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def invertTree(self, root: TreeNode) -> TreeNode:
        if not root:
            return 
        root.left, root.right = root.right, root.left
        self.invertTree(root.left)
        self.invertTree(root.right)
        return root
```