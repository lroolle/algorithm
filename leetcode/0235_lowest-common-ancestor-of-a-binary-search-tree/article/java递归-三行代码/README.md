"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree Java(递归) 三行代码 "
author = ["get996"]
date = 2019-12-12T06:40:41+08:00
tags = ["Leetcode", "Algorithms", "Java", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# Java(递归) 三行代码

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [Java(递归) 三行代码](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/javadi-gui-san-xing-dai-ma-by-get996/) by [get996](https://leetcode-cn.com/u/get996/)

### 解题思路
三行代码
根结点比两个结点都大 就在左子树找
根结点比两个结点都小 就在右子树找
否则就返回根结点

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
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if(root.val>p.val && root.val>q.val) return lowestCommonAncestor(root.left, p, q);
        if(root.val<p.val && root.val<q.val) return lowestCommonAncestor(root.right,p,q);
        return root;
    }
    
}
```