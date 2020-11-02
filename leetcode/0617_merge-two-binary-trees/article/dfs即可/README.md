"+++
title = "0617. Merge Two Binary Trees DFS即可 "
author = ["flyhigher139"]
date = 2020-09-23T02:13:55+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Java"]
categories = ["leetcode"]
draft = false
+++

# DFS即可

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [DFS即可](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/dfsji-ke-by-flyhigher139/) by [flyhigher139](https://leetcode-cn.com/u/flyhigher139/)

本题比较简单，不用解释了，直接看代码即可：
```python3
class Solution:
    def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:
        if not t1 or not t2:
            return t1 or t2
        t1.val = t1.val + t2.val
        t1.left = self.mergeTrees(t1.left, t2.left)
        t1.right = self.mergeTrees(t1.right, t2.right)
        return t1
```
```java
class Solution {
    public TreeNode mergeTrees(TreeNode t1, TreeNode t2) {
        if (t1 == null) {
            return t2;
        }
        if (t2 == null) {
            return t1;
        }
        t1.val = t1.val + t2.val;
        t1.left = mergeTrees(t1.left, t2.left);
        t1.right = mergeTrees(t1.right, t2.right);
        return t1;
    }
}
```