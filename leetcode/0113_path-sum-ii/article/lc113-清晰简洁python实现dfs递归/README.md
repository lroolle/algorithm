"+++
title = "0113. Path Sum II LC113 清晰简洁Python实现——dfs递归 "
author = ["jhhuang"]
date = 2020-08-14T06:40:08+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3", "DepthfirstSearch", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# LC113 清晰简洁Python实现——dfs递归

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [LC113 清晰简洁Python实现——dfs递归](https://leetcode-cn.com/problems/path-sum-ii/solution/lc113-qing-xi-jian-ji-pythonshi-xian-dfsdi-gui-by-/) by [jhhuang](https://leetcode-cn.com/u/jhhuang/)

### 代码

```python3
class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> List[List[int]]:
        res = []
        if not root: 
            return []
        
        def dfs(track, root, sum):
            if not root:
                return
            if not root.left and not root.right and root.val == sum:
                track += [root.val]
                res.append(track)
            
            dfs(track + [root.val], root.left, sum - root.val)
            dfs(track + [root.val], root.right, sum - root.val)
        
        dfs([], root, sum)
        return res
```