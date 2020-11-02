"+++
title = "0113. Path Sum II DFS非递归和递归python3 "
author = ["baiyizhe"]
date = 2019-11-04T02:54:18+08:00
tags = ["Leetcode", "Algorithms", "Python3"]
categories = ["leetcode"]
draft = false
+++

# DFS非递归和递归python3

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [DFS非递归和递归python3](https://leetcode-cn.com/problems/path-sum-ii/solution/dfsfei-di-gui-python3-by-baiyizhe/) by [baiyizhe](https://leetcode-cn.com/u/baiyizhe/)

非递归
```
class Solution:
    def pathSum(self, root: TreeNode, sum_: int) -> List[List[int]]:
        if not root: return []
        stack = [([root.val], root)]
        res = []
        while stack:
            tmp, node = stack.pop()
            if not node.right and not node.left and sum(tmp) == sum_:
                res.append(tmp)
            if node.right:
                stack.append((tmp + [node.right.val], node.right))
            if node.left:
                stack.append((tmp + [node.left.val], node.left))
        return res
```


递归
```
class Solution:
    def pathSum(self, root: TreeNode, sum_: int) -> List[List[int]]:
        def helper(root, tmp, sum_):
            if not root:
                return 
            if not root.left and not root.right and sum_ - root.val == 0:
                tmp += [root.val]
                res.append(tmp)
            helper(root.left, tmp + [root.val], sum_ - root.val)
            helper(root.right, tmp + [root.val], sum_ - root.val)
        res = []
        helper(root, [], sum_)
        return res
```
