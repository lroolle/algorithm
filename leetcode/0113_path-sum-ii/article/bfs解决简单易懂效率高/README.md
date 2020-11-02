"+++
title = "0113. Path Sum II BFS解决，简单易懂，效率高。 "
author = ["ybzdqhl"]
date = 2020-08-21T04:59:36+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Tree", "BreadthfirstSearch", "Python"]
categories = ["leetcode"]
draft = false
+++

# BFS解决，简单易懂，效率高。

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [BFS解决，简单易懂，效率高。](https://leetcode-cn.com/problems/path-sum-ii/solution/bfsjie-jue-jian-dan-yi-dong-xiao-lu-gao-by-bold-ki/) by [ybzdqhl](https://leetcode-cn.com/u/ybzdqhl/)

### 解题思路
采用广度优先遍历，每遍历一个节点，就构建一个从根节点到当前节点的路径数组，如果当前节点为叶子节点，那么该路径数组便包含了从根节点到该叶子节点的路径(经过的所有节点)，此时对该路径数组进行求和，如果求和后的值等于目标值sum,那么就将该数组添加到结果列表中。
遍历结束后输出结果列表即可。

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def pathSum(self, root: TreeNode, sum_: int) -> List[List[int]]:
        if not root:
            return []
        
        stack = [(root, [root.val])]
        res = list()
        while stack:
            node, tmp = stack.pop(0)
            if not node.left and not node.right and sum(tmp) == sum_:
                res.append(tmp)
            if node.left:
                stack.append((node.left,tmp+[node.left.val]))
            if node.right:
                stack.append((node.right,tmp+[node.right.val]))
        
        return res
```