"+++
title = "0226. Invert Binary Tree 226.翻转二叉树 "
author = ["yi-wen-statistics"]
date = 2020-09-07T03:33:16+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 226.翻转二叉树

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [226.翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/solution/226fan-zhuan-er-cha-shu-by-yi-wen-statistics/) by [yi-wen-statistics](https://leetcode-cn.com/u/yi-wen-statistics/)

## 方法1——深度优先搜索
自顶向下逐步翻转

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
        if not root:
            return root
        tmp = root.left
        root.left = root.right
        root.right = tmp
        self.invertTree(root.left)
        self.invertTree(root.right)
        return root
```

### 复杂度分析
时间复杂度：O(N)
空间复杂度：O(N)