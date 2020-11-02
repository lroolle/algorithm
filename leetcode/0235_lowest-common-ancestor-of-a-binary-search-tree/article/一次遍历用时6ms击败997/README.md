"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 一次遍历用时6ms，击败99.7% "
author = ["byakuya-2"]
date = 2020-09-27T05:31:37+08:00
tags = ["Leetcode", "Algorithms", "Java", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 一次遍历用时6ms，击败99.7%

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [一次遍历用时6ms，击败99.7%](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/yi-ci-bian-li-yong-shi-6msji-bai-997-by-byakuya-2/) by [byakuya-2](https://leetcode-cn.com/u/byakuya-2/)

利用二叉搜索树的特性，
如果当前结点大于p和q，最近公共祖先一定在左子树；
如果小于p和q，一定在右子树；
如果只大于其中一个，或者出现相等的情况，返回当前结点。

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null) return null;
        if (root.val > p.val && root.val > q.val) {
            //往左找
            return lowestCommonAncestor(root.left, p, q);
        } else if (root.val < p.val && root.val < q.val) {
            //往右找
            return lowestCommonAncestor(root.right, p, q);
        } else {
            return root;
        }
    }
}
```
