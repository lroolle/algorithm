"+++
title = "0617. Merge Two Binary Trees python简单递归 "
author = ["liuxiaolong"]
date = 2020-07-13T13:34:20+08:00
tags = ["Leetcode", "Algorithms", "Python3"]
categories = ["leetcode"]
draft = false
+++

# python简单递归

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [python简单递归](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/pythonjian-dan-di-gui-by-liuxiaolong/) by [liuxiaolong](https://leetcode-cn.com/u/liuxiaolong/)

```
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:
        if not (t1 and t2):
            return t1 if t1 else t2
        t1.val += t2.val
        t1.left = self.mergeTrees(t1.left, t2.left)
        t1.right = self.mergeTrees(t1.right, t2.right)
        return t1
```
