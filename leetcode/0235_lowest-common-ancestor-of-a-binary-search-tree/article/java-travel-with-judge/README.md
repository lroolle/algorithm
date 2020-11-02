"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree [Java] Travel with Judge. "
author = ["lincs"]
date = 2020-09-27T02:45:05+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# [Java] Travel with Judge.

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [[Java] Travel with Judge.](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/java-travel-with-judge-by-lincs/) by [lincs](https://leetcode-cn.com/u/lincs/)

- 二叉搜索树的最近公共祖先
```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (p.val < root.val && q.val < root.val) {
            return lowestCommonAncestor(root.left, p, q); 
        }
        if (root.val < p.val && root.val < q.val) {
            return lowestCommonAncestor(root.right, p, q);
        }
        return root;
    }
}
```
- 普通树的最近公共祖先
```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null) return null;
        if (p == root) return p;
        if (q == root) return q;
        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);
        if (left == null) return right;
        if (right == null) return left;
        return root;
    }
}
```
