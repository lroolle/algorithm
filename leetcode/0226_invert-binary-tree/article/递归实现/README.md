"+++
title = "0226. Invert Binary Tree 递归实现 "
author = ["eager-2"]
date = 2020-08-15T03:40:05+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 递归实现

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [递归实现](https://leetcode-cn.com/problems/invert-binary-tree/solution/di-gui-shi-xian-by-eager-2/) by [eager-2](https://leetcode-cn.com/u/eager-2/)

### 解题思路
递归实现

### 代码

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public TreeNode invertTree(TreeNode root) {
        if(root == null) return null;
        TreeNode tmp = root.left;
        root.left = root.right;
        root.right = tmp;
        invertTree(root.left);
        invertTree(root.right);
        return root;
    }
}
```