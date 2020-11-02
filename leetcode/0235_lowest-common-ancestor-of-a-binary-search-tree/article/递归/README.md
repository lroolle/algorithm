"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 递归 "
author = ["ling-ling-zhu-qi"]
date = 2020-09-27T05:41:44+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 递归

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [递归](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/di-gui-by-ling-ling-zhu-qi-2/) by [ling-ling-zhu-qi](https://leetcode-cn.com/u/ling-ling-zhu-qi/)

# 思路

首先，二叉搜索树具有左子树<根节点<右子树的特点，最近公共祖先必须满足节点要么在祖先节点的左右两个子树上，要么其中一个指定节点为祖先节点，另一个在其子树上。代码如下：
```
public class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root.val > p.val && root.val > q.val) {
            return lowestCommonAncestor(root.left, p, q);
        } else if (root.val < p.val && root.val < q.val) {
            return lowestCommonAncestor(root.right, p, q);
        } else { //其中一个节点为祖先节点或两个节点分布在左右子树上
            return root;
        }
    }
}
```
