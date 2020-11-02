"+++
title = "0530. Minimum Absolute Difference in BST 中序遍历 "
author = ["powcai"]
date = 2020-06-14T14:53:56+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3"]
categories = ["leetcode"]
draft = false
+++

# 中序遍历

> [0530. Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
> [中序遍历](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/solution/zhong-xu-bian-li-by-powcai-2/) by [powcai](https://leetcode-cn.com/u/powcai/)

## 思路:

中序遍历

## 代码:

```python
class Solution:
    def getMinimumDifference(self, root: TreeNode) -> int:
        res, prev  = float("inf"), float("-inf") 
        
        def dfs(root):
            nonlocal res, prev
            if not root:
                return 
            dfs(root.left)
            res = min(res, root.val - prev)
            prev = root.val
            dfs(root.right) 
            
        dfs(root)
        return res
```

