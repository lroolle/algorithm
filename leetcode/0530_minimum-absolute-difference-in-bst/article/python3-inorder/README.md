"+++
title = "0530. Minimum Absolute Difference in BST Python3, Inorder "
author = ["lionKing_njuer"]
date = 2020-06-17T14:57:30+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# Python3, Inorder

> [0530. Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
> [Python3, Inorder](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/solution/python3-inorder-by-littlelion_njuer/) by [lionKing_njuer](https://leetcode-cn.com/u/lionking_njuer/)

### 解题思路
中序遍历一颗二叉搜索树是很显然的，因为可以获得升序的数组
递归过程中同步更新最小值
注意初始化的数值要正确

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def __init__(self):
        self.res = float('inf')
        self.pre = float('-inf')
    def getMinimumDifference(self, root):
        self.inOrder(root)
        return self.res
    def inOrder(self, root):
        if not root:
            return None
        self.inOrder(root.left)
        current = root.val
        self.res = min(self.res, current - self.pre)
        self.pre = current
        self.inOrder(root.right)
```