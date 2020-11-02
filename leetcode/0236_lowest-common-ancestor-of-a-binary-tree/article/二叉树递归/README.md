"+++
title = "0236. Lowest Common Ancestor of a Binary Tree 二叉树+递归 "
author = ["xu-chen-chen-d"]
date = 2020-08-21T10:12:47+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Tree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 二叉树+递归

> [0236. Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
> [二叉树+递归](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/er-cha-shu-di-gui-by-xu-chen-chen-d/) by [xu-chen-chen-d](https://leetcode-cn.com/u/xu-chen-chen-d/)

### 解题思路
此处撰写解题思路
如果两边各有一个p或q，那么答案应该直接是根节点root。
如果都在左，则答案是left。
如果都在右，则答案是right。
### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if not root:
            return root
        if root == q or root == p:
            return root
        left = self.lowestCommonAncestor(root.left,p,q)
        right = self.lowestCommonAncestor(root.right,p,q)
        
        if left and right:
            return root
        if not left:
            return right
        if not right:
            return left
    
```