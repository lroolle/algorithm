"+++
title = "0094. Binary Tree Inorder Traversal 写给自己的题解：花式进出栈~~~~~ "
author = ["kknike"]
date = 2020-09-14T03:00:08+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Stack"]
categories = ["leetcode"]
draft = false
+++

# 写给自己的题解：花式进出栈~~~~~

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [写给自己的题解：花式进出栈~~~~~](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/xie-gei-zi-ji-de-ti-jie-hua-shi-jin-chu-zhan-by-kk/) by [kknike](https://leetcode-cn.com/u/kknike/)

### 解题思路
之前想问题太死了，如果逻辑中同时要进栈或者出栈的话，
可以同时写两个判断，对不同情况同时进行不同处理

### 代码

```python3
class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        stack, res, cur = [], [], root
        while cur or stack:  # 混合进出栈的时候，可以使用这样的判断
            while cur: 
                stack.append(cur)
                cur = cur.left
            if stack:
                this = stack.pop()
                res.append(this.val)
                cur = this.right
        return res
```