"+++
title = "0113. Path Sum II python题解--简单的dfs "
author = ["xiao-xue-66"]
date = 2020-08-25T03:35:05+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# python题解--简单的dfs

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [python题解--简单的dfs](https://leetcode-cn.com/problems/path-sum-ii/solution/pythonti-jie-jian-dan-de-dfs-by-xiao-xue-66-5/) by [xiao-xue-66](https://leetcode-cn.com/u/xiao-xue-66/)

### 解题思路
![image.png](https://pic.leetcode-cn.com/1598326498-nikbuD-image.png)
### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> List[List[int]]:
        if not root:
            return []
        res = []

        def dfs(root,stack,sum_):
            sum_ += root.val
            if not root.left and not root.right:
                if sum == sum_:
                    res.append(stack[:]+[root.val])
            elif not root.left:
                dfs(root.right,stack+[root.val],sum_)
            elif not root.right:
                dfs(root.left,stack+[root.val],sum_)
            else:
                dfs(root.left,stack+[root.val],sum_)
                dfs(root.right,stack+[root.val],sum_)
        
        dfs(root,[],0)
        return res

```