"+++
title = "0501. Find Mode in Binary Search Tree python3广度优先 "
author = ["peter-165"]
date = 2020-04-10T08:12:04+08:00
tags = ["Leetcode", "Algorithms", "Python3", "BreadthfirstSearch", "Python"]
categories = ["leetcode"]
draft = false
+++

# python3广度优先

> [0501. Find Mode in Binary Search Tree](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)
> [python3广度优先](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/python3yan-du-you-xian-by-peter-165/) by [peter-165](https://leetcode-cn.com/u/peter-165/)

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def findMode(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        ans = []
        
        res = []
        queue = collections.deque()
        queue.append(root)

        while queue:
            for _ in range(len(queue)):
                node = queue.popleft()
                res.append(node.val)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

        tmp = collections.Counter(res).most_common()
        for k, v in tmp:
            if v == tmp[0][1]:
                ans.append(k)

        return ans
```