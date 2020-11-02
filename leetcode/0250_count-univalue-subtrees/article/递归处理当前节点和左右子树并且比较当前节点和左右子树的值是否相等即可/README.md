"+++
title = "0250. Count Univalue Subtrees 递归处理当前节点和左右子树并且比较当前节点和左右子树的值是否相等即可 "
author = ["yand-3"]
date = 2020-08-15T06:32:59+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 递归处理当前节点和左右子树并且比较当前节点和左右子树的值是否相等即可

> [0250. Count Univalue Subtrees](https://leetcode-cn.com/problems/count-univalue-subtrees/)
> [递归处理当前节点和左右子树并且比较当前节点和左右子树的值是否相等即可](https://leetcode-cn.com/problems/count-univalue-subtrees/solution/di-gui-chu-li-dang-qian-jie-dian-he-zuo-you-zi-shu/) by [yand-3](https://leetcode-cn.com/u/yand-3/)

### 解题思路
此处撰写解题思路

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
    public  int countUnivalSubtrees(TreeNode root) {
        helper(root);
        return result;
    }

     int result = 0;

    private  boolean helper(TreeNode root) {
        if (root == null) {
            return true;
        }
        if (root.left == null && root.right == null) {
            // 叶子节点肯定满足条件
            result++;
            return true;
        }
        // 不是叶子节点的话，判断当前节点是否和左右节点的值相等
        boolean left = helper(root.left);
        boolean right = helper(root.right);
        // 左右子树均满足条件，判断当前节点是否和左右孩子节点值相等
        if (left && ((root.left == null || root.val == root.left.val))
                && right && (root.right == null || root.val == root.right.val)) {
            result++;
            return true;
        }
        return false;
    }
}
```