"+++
title = "0701. Insert into a Binary Search Tree Python 116 ms 利用父节点指针达到 O(1)空间复杂度 "
author = ["carrysteve"]
date = 2020-02-22T12:34:56+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3", "BinarySearchTree"]
categories = ["leetcode"]
draft = false
+++

# Python 116 ms 利用父节点指针达到 O(1)空间复杂度

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [Python 116 ms 利用父节点指针达到 O(1)空间复杂度](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/python-116-ms-li-yong-fu-jie-dian-zhi-zhen-da-dao-/) by [carrysteve](https://leetcode-cn.com/u/carrysteve/)

### 解题思路
利用两个指针分别指向分节点与当前节点，减少判断

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def insertIntoBST(self, root: TreeNode, val: int) -> TreeNode:
        
        vnode = TreeNode(val)
        if not root: return vnode
        else:
            pre, node = None, root
            while node:
                pre = node
                node = node.left if val < node.val else node.right

            if pre.val < val: pre.right = vnode
            else: pre.left = vnode
            return root

```