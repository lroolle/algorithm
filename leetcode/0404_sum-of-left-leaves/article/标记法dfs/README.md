"+++
title = "0404. Sum of Left Leaves 标记法dfs "
author = ["yi-wen-statistics"]
date = 2020-09-08T08:02:40+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 标记法dfs

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [标记法dfs](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/biao-ji-fa-dfs-by-yi-wen-statistics/) by [yi-wen-statistics](https://leetcode-cn.com/u/yi-wen-statistics/)

## 方法1——递归式
设置left和right两个标记参数，把往左子树递归和往右子树递归区分开来

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def sumOfLeftLeaves(self, root: TreeNode) -> int:
        res = []
        def dfs(root, left, right, res):
            if not root:
                return
            if not root.left and not root.right:
                if left and not right:
                    res.append(root.val)
                    return
            dfs(root.left, True, False, res)
            dfs(root.right, False, True, res)
        dfs(root, False, False, res)
        return sum(res)
```

### 复杂度分析
时间复杂度：O(N)
空间复杂度：O(N)