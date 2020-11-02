"+++
title = "0226. Invert Binary Tree python题解--直接交换左右子树 "
author = ["xiao-xue-66"]
date = 2020-08-26T01:19:41+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# python题解--直接交换左右子树

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [python题解--直接交换左右子树](https://leetcode-cn.com/problems/invert-binary-tree/solution/pythonti-jie-zhi-jie-jiao-huan-zuo-you-zi-shu-by-x/) by [xiao-xue-66](https://leetcode-cn.com/u/xiao-xue-66/)

### 解题思路
![image.png](https://pic.leetcode-cn.com/1598404715-EVyvWg-image.png)
- 仔细看这个，其实就是一个镜像了，只需要将左右子树直接对换就好了
![image.png](https://pic.leetcode-cn.com/1598404731-geTfox-image.png)

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
            return None

        def dfs(root):
            if not root:
                return
            root.left, root.right = root.right, root.left
            dfs(root.left)
            dfs(root.right)

        dfs(root) 

        return root
```