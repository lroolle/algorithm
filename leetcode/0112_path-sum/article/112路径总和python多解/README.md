"+++
title = "0112. Path Sum 112.路径总和（Python多解） "
author = ["yi-wen-statistics"]
date = 2020-09-07T02:26:08+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 112.路径总和（Python多解）

> [0112. Path Sum](https://leetcode-cn.com/problems/path-sum/)
> [112.路径总和（Python多解）](https://leetcode-cn.com/problems/path-sum/solution/112lu-jing-zong-he-shen-du-you-xian-sou-suo-by-yi-/) by [yi-wen-statistics](https://leetcode-cn.com/u/yi-wen-statistics/)

## 方法1——深度优先
访问所有完整的路径，将其保存，最后判断。

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        res = []
        def helper(root, sum, tmp):
            if not root:
                return
            tmp += root.val
            if not root.left and not root.right:
                if tmp == sum:
                    res.append(True)
                    return
            helper(root.left, sum, tmp)
            helper(root.right, sum, tmp)

        helper(root, sum, 0)

        return res != []
```

### 复杂度分析
时间复杂度：O(N)
空间复杂度：最坏情况下O(N)

## 方法2——加速剪枝
遇到True的情况全局停止搜索

### 代码

```
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    flag = False
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        def helper(root, sum, tmp):
            if not root or self.flag:
                return
            tmp += root.val
            if not root.left and not root.right:
                if tmp == sum:
                    self.flag = True
                    return
            helper(root.left, sum, tmp)
            helper(root.right, sum, tmp)

        helper(root, sum, 0)

        return self.flag
```

### 复杂度分析
同上，但效率更高

## 方法3——广度优先搜索
设置值队列和节点队列，遇到叶子节点时只要存在value==sum的情况，剪枝

### 代码
```
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        if not root:
            return False
        deq1 = collections.deque()
        deq2 = collections.deque()
        deq1.append(root)
        deq2.append(root.val)
        flag = False
        while deq1:
            n = len(deq1)
            for i in range(n):
                node = deq1.popleft()
                value = deq2.popleft()
                if not node.left and not node.right:
                    if value == sum:
                        flag = True
                        break
                if node.left:
                    deq1.append(node.left)
                    deq2.append(node.left.val+value)
                if node.right:
                    deq1.append(node.right)
                    deq2.append(node.right.val+value)
            if flag:
                break            
        
        return flag
```

### 复杂度分析
时间复杂度：O(N)
空间复杂度：O(N)