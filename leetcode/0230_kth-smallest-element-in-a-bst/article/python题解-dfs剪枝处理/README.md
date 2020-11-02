"+++
title = "0230. Kth Smallest Element in a BST python题解--DFS剪枝处理 "
author = ["xiao-xue-66"]
date = 2020-07-08T00:49:20+08:00
tags = ["Leetcode", "Algorithms", "Python3", "DepthfirstSearch", "Python"]
categories = ["leetcode"]
draft = false
+++

# python题解--DFS剪枝处理

> [0230. Kth Smallest Element in a BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)
> [python题解--DFS剪枝处理](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/solution/pythonti-jie-dfsjian-zhi-chu-li-by-xiao-xue-66/) by [xiao-xue-66](https://leetcode-cn.com/u/xiao-xue-66/)

### 解题思路
![image.png](https://pic.leetcode-cn.com/042b325988a67759c9c979e08605be483576aae6c1aa0e098308894572ea672e-image.png)

- 虽然这是一道中等难度的题目，但我们拿到题目后可以很简单的想到使用DFS，如题需要找到第K个小的数，根据二叉搜索树的特性可以知道中序遍历的结果是从小到大的
- 所以我们采用采用中序遍历，找到第k个节点即可
- 这里我加入了剪枝的处理，就是当左子树种已经找到了我们的第K个节点后，其实没必要在遍历右子树
- 时间复杂度`O(n)`n为二叉树节点个数，空间复杂度为`O(1)`
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
        
        num = 0
        res = 0
        def dfs(root,k):
            nonlocal res
            nonlocal num

            flag = False

            if not root:
                return False

            flag = dfs(root.left,k)

            if flag:#剪枝
                return True

            num += 1
            if num == k:
                res = root.val
                return True

            flag = dfs(root.right,k)
            return flag
        dfs(root,k)
        return res
        
            
```