"+++
title = "0250. Count Univalue Subtrees 深度搜索 "
author = ["ren-can-r"]
date = 2020-08-14T06:39:22+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 深度搜索

> [0250. Count Univalue Subtrees](https://leetcode-cn.com/problems/count-univalue-subtrees/)
> [深度搜索](https://leetcode-cn.com/problems/count-univalue-subtrees/solution/shen-du-sou-suo-by-ren-can-r/) by [ren-can-r](https://leetcode-cn.com/u/ren-can-r/)

   class Solution {
        int count = 0;
        public int countUnivalSubtrees(TreeNode root) {
            if (root == null) return 0;
            isUnival(root);
            return count;
        }
        private boolean isUnival(TreeNode root) {
            boolean flag = true;
            if (root.left != null)
                if (!isUnival(root.left) || root.val != root.left.val) flag = false;
            if (root.right != null)
                if (!isUnival(root.right) || root.val != root.right.val) flag = false;
            if (flag) count++;
            return flag;
        }
    }