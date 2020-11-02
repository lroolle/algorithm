"+++
title = "0404. Sum of Left Leaves 分治法 and 前序遍历 "
author = ["durden-dong"]
date = 2020-06-28T12:49:20+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3"]
categories = ["leetcode"]
draft = false
+++

# 分治法 and 前序遍历

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [分治法 and 前序遍历](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/fen-zhi-fa-and-qian-xu-bian-li-by-durden-dong/) by [durden-dong](https://leetcode-cn.com/u/durden-dong/)

# 解法一： 分治法递归
```
if root is None:
    return 0
if root.left is None: 
    return self.sumOfLeftLeaves(root.right)

if root.left.left is None and root.left.right is None: 
    return root.left.val + self.sumOfLeftLeaves(root.right)

return self.sumOfLeftLeaves(root.left) + self.sumOfLeftLeaves(root.right)
```
# 解法二： 前序遍历
```
# check 是否为叶子结点：
def helper(root):

    if root is None:
        return False
    if root.left is None and root.right is None:
        return True
    else:
        return False

# 前序遍历，在遍历过程中，发现叶子结点 and 是左树的，则加入到self.res中。
def dfs(root,marker):
    if root is None:
        return 0

    if helper(root) and marker:
        self.res += root.val
    dfs(root.left,marker = 1)
    dfs(root.right,marker = 0)

self.res = 0    
dfs(root,0)
return self.res
```