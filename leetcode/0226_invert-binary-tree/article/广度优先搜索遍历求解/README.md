"+++
title = "0226. Invert Binary Tree 广度优先搜索遍历求解 "
author = ["jkkk1996"]
date = 2020-09-09T13:17:35+08:00
tags = ["Leetcode", "Algorithms", "Python3", "BreadthfirstSearch", "迭代"]
categories = ["leetcode"]
draft = false
+++

# 广度优先搜索遍历求解

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [广度优先搜索遍历求解](https://leetcode-cn.com/problems/invert-binary-tree/solution/yan-du-you-xian-sou-suo-bian-li-qiu-jie-by-jkkk199/) by [jkkk1996](https://leetcode-cn.com/u/jkkk1996/)

### 解题思路
用广度优先搜索的方法来进行解题，将每一层的字数加入到stack中，然后对每一个子树的左右子树进行交换

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
        if not root : return 
        stack = [root]
        while stack :
            length = len(stack)
            for i in range(length):
                cur = stack.pop(0)
                cur.left,cur.right = cur.right,cur.left
                if cur.left :
                    stack.append(cur.left)
                if cur.right :
                    stack.append(cur.right)
        return root
```