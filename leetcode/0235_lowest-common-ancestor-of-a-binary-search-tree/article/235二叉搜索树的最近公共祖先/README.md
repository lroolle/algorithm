"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 235.二叉搜索树的最近公共祖先 "
author = ["jue-qiang-zha-zha"]
date = 2020-09-18T08:26:10+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3"]
categories = ["leetcode"]
draft = false
+++

# 235.二叉搜索树的最近公共祖先

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [235.二叉搜索树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/235er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xi-5/) by [jue-qiang-zha-zha](https://leetcode-cn.com/u/jue-qiang-zha-zha/)

## 解题思路

两种解法，递归和遍历。

解体关键：二叉搜索树（BST）的性质。所有比父节点小的节点都位于左子树，所有比父节点大的节点都位于右子树，且任意节点的子树也都是二叉搜索树。

因此，关键是找到划分节点。如果两个节点p、q的值都比某一结点node的大，则p、q一定在node节点的右子树；如果p、q的值都比node的值小，则p、q一定在node的左子树；如果不满足上面的两种情况，则node为我们要找到的节点。

### 解法1：递归

判断：
（1）如果两个节点p、q的值都比某一结点node的大，则对node节点的右子树进行调用自身；
（2）如果p、q的值都比node的值小，则对node节点的左子树进行调用自身；
（3）如果不满足上面的两种情况，则return node。

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
        if p.val > root.val and q.val > root.val:
            return self.lowestCommonAncestor(root.right, p, q)
        elif p.val < root.val and q.val < root.val:
            return self.lowestCommonAncestor(root.left, p, q)
        else:
            return root
        
```
时间复杂度：O(n)， 每个节点访问一次。
空间复杂度：O(n)，取决于递归栈的深度，最大不超过n。
### 解法2：遍历

问题的关键是找到两个节点p、q的划分节点，因此不需要利用递归，直接遍历树节点并进行值大小的比较即可。

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
        node = root     # 扫描指针
        while node:
            if p.val > node.val and q.val > node.val:
                node = node.right
            elif p.val < node.val and q.val < node.val:
                node = node.left
            else:
                return node
        
```
时间复杂度：O(n)，每个节点只访问一次。
空间复杂度：O(1)，不需要辅助栈和队列。