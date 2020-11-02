"+++
title = "0230. Kth Smallest Element in a BST 中序遍历BST "
author = ["elevenxx"]
date = 2020-08-25T01:02:12+08:00
tags = ["Leetcode", "Algorithms", "Python3", "BinarySearchTree", "Python"]
categories = ["leetcode"]
draft = false
+++

# 中序遍历BST

> [0230. Kth Smallest Element in a BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)
> [中序遍历BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/solution/zhong-xu-bian-li-bst-by-elevenxx/) by [elevenxx](https://leetcode-cn.com/u/elevenxx/)


```python3
class Solution:
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        stack = []
        while True:
            while root:
                stack.append(root)
                root = root.left
            root = stack.pop()
            k -= 1
            if k == 0: return root.val
            root = root.right
```

