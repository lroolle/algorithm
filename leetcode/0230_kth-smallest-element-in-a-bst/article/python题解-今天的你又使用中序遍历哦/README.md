"+++
title = "0230. Kth Smallest Element in a BST python题解--今天的你又使用中序遍历哦！！！ "
author = ["xiao-xue-66"]
date = 2020-08-27T01:15:49+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# python题解--今天的你又使用中序遍历哦！！！

> [0230. Kth Smallest Element in a BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)
> [python题解--今天的你又使用中序遍历哦！！！](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/solution/pythonti-jie-jin-tian-de-ni-you-shi-yong-zhong-xu-/) by [xiao-xue-66](https://leetcode-cn.com/u/xiao-xue-66/)

### 解题思路
![image.png](https://pic.leetcode-cn.com/1598490893-StZdHR-image.png)
- 二叉搜索树的中序遍历是递增序列哦

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        self.num = 0
        self.res = 0
        def dfs(root):
            if not root:
                return
            dfs(root.left)
            self.num += 1
            if self.num == k:
                self.res = root.val
                return 
            dfs(root.right)
        
        dfs(root)
        return self.res
            
```