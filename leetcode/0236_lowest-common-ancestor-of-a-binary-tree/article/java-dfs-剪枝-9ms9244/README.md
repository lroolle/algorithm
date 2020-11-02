"+++
title = "0236. Lowest Common Ancestor of a Binary Tree JAVA DFS + 剪枝   9ms，92.44% "
author = ["lava-4"]
date = 2020-01-14T13:45:57+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# JAVA DFS + 剪枝   9ms，92.44%

> [0236. Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
> [JAVA DFS + 剪枝   9ms，92.44%](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/java-dfs-jian-zhi-9ms9244-by-lava-4/) by [lava-4](https://leetcode-cn.com/u/lava-4/)

我们使用DFS搜索每一个节点的左右子树：
1、若子树上存在p和q的公共节点，返回此公共节点
2、若不存在公共节点，但是存在p或q任意一个节点，返回此节点
3、若不存在公共、p、q节点，则返回null。

那么，有以下几个结论：
0、若当前节点为null、p、q之一，直接返回当前节点
1、若左子树上存在公共节点（返回值非p、q），则函数返回值为左子树返回值，不需再遍历右子树
2、若左子树返回null，则函数返回值为右子树返回值
3、若左子树、右子树返回值均为非null，则肯定为一个p，一个q，则公共节点为当前节点
4、最后一种情况，即左子树返回非null，右子树返回null，则函数返回值为左子树返回值

*针对p是q的子节点这种特殊情况，上述方案依然可行（但就没有办法剪枝了，可以考虑针对此情况加一个标记，不再遍历右子树）。


```
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
        //结论0
        if (root == null || root == p || root == q)
            return root;
        TreeNode left = lowestCommonAncestor(root.left, p, q);

        //结论1
        if (left != null && left != q && left != p)
            return left;
        TreeNode right = lowestCommonAncestor(root.right, p, q);

        //结论3
        if (left != null && right != null)
            return root;

        //结论2 4
        else
            return left == null ? right : left;
    }
}
```
