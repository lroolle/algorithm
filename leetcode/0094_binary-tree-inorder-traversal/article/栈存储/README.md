"+++
title = "0094. Binary Tree Inorder Traversal 栈存储 "
author = ["wu-ming-shi-7"]
date = 2020-08-30T09:41:33+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 栈存储

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [栈存储](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/zhan-cun-chu-by-wu-ming-shi-7/) by [wu-ming-shi-7](https://leetcode-cn.com/u/wu-ming-shi-7/)

### 解题思路
采用栈保存中间节点，以便回溯
### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        # 左  ->  中  ->  右
        stack = []
        res = []
        while root or len(stack) != 0:
            if root:
                stack.append(root)
                root = root.left
            else:
                root = stack.pop(-1)
                res.append(root.val)
                root = root.right
        
        return res
        
```