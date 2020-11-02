"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 二叉搜索树（BST）递归 "
author = ["xu-chen-chen-d"]
date = 2020-08-21T09:41:29+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Tree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 二叉搜索树（BST）递归

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [二叉搜索树（BST）递归](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/er-cha-sou-suo-shu-bstdi-gui-by-xu-chen-chen-d/) by [xu-chen-chen-d](https://leetcode-cn.com/u/xu-chen-chen-d/)

### 解题思路
此处撰写解题思路
首先要懂什么事二叉搜索树：左子树val小于根节点val，根节点值小于右子树val。
下面就开始做题：
1.p与q,将小的赋值给p，大的赋值给q。
2.只有程序中的三种情况了，分别是：
p.val <= root.val <= q.val（结束，返回root值即结果）、
root.val < p.val < q.val（root.right递归）、
root.val > q.val > p.val（root.left递归）。
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
        temp = None
        if p.val > q.val:
            temp = p
            p = q
            q = temp

        if p.val <= root.val <= q.val:
            return root
        if root.val < p.val < q.val:
            return self.lowestCommonAncestor(root.right,p,q)
        if root.val > q.val > p.val:
            return self.lowestCommonAncestor(root.left,p,q)

```