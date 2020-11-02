"+++
title = "0236. Lowest Common Ancestor of a Binary Tree Recusion + Iteration 两种解法 "
author = ["SupermanBlues"]
date = 2020-08-30T05:01:10+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Recursion", "Python"]
categories = ["leetcode"]
draft = false
+++

# Recusion + Iteration 两种解法

> [0236. Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
> [Recusion + Iteration 两种解法](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/recusion-iteration-liang-chong-jie-fa-by-supermanb/) by [SupermanBlues](https://leetcode-cn.com/u/supermanblues/)

### 解题思路
此处撰写解题思路

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    """
        Two Solutions (LCA lowest common ancestor)
        Sol 0: Recursion
        Time Complexity O(N) 
        Space Complexity O(1) without extra space
        
        Sol 1: Iteration 
        Time Complexity O(N)
        Space Complexity O(N)
    """
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        # iteration
        parents, seen = dict(), set()
        def dfs(c = root, father = None): # father == parent == 爸爸
            if not c: return
            parents[c] = father
            dfs(c.left, c)
            dfs(c.right, c)

        dfs()
        while p:
            seen.add(p)
            p = parents[p]

        while q:
            if q in seen: return q
            q = parents[q]

        return # 一定有解， 只是为了编译通过

    def lowestCommonAncestorII(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        # Recursion
        if root in (None, p, q): return root
        l, r =  self.lowestCommonAncestor(root.left,  p, q), \
                self.lowestCommonAncestor(root.right, p, q)
        return r if l is None else l if r is None else root

```