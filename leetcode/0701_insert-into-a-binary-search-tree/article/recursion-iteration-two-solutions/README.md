"+++
title = "0701. Insert into a Binary Search Tree Recursion + Iteration Two Solutions... "
author = ["SupermanBlues"]
date = 2020-08-03T06:22:26+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Recursion", "迭代", "Python"]
categories = ["leetcode"]
draft = false
+++

# Recursion + Iteration Two Solutions...

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [Recursion + Iteration Two Solutions...](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/recursion-iteration-two-solutions-by-supermanblues/) by [SupermanBlues](https://leetcode-cn.com/u/supermanblues/)

### 解题思路
此处撰写解题思路

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def insertIntoBST(self, root: TreeNode, val: int) -> TreeNode:
        # recursion solution
        if root is None: return TreeNode(val)
        if val < root.val:
            root.left = self.insertIntoBST(root.left, val)
        else:
            root.right = self.insertIntoBST(root.right, val)
        return root

    def insertIntoBST2(self, root: TreeNode, val: int) -> TreeNode:
        # iteration solution
        if root is None: return TreeNode(val)
        p = root
        while True:
            if val < p.val:
                if p.left is None:
                    p.left = TreeNode(val)
                    break
                else:
                    p = p.left
            else:
                if p.right is None:
                    p.right = TreeNode(val)
                    break
                else:
                    p = p.right
        
        return root

```