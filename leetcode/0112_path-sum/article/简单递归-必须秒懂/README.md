"+++
title = "0112. Path Sum 简单递归，🤷‍♀️ 必须秒懂！ "
author = ["sweetiee"]
date = 2020-07-07T04:53:33+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree", "DepthfirstSearch", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 简单递归，🤷‍♀️ 必须秒懂！

> [0112. Path Sum](https://leetcode-cn.com/problems/path-sum/)
> [简单递归，🤷‍♀️ 必须秒懂！](https://leetcode-cn.com/problems/path-sum/solution/jian-dan-di-gui-bi-xu-miao-dong-by-sweetiee-2/) by [sweetiee](https://leetcode-cn.com/u/sweetiee/)

🙋 今日打卡！

## 一、题目分析

求解从 `root` 到叶子节点是否存在路径和为 `sum` 的路径 `hasPathSum(root, sum)`。

可以转换成求解从 `root.left` 或者 `root.right` 到叶子节点是否存在路径和为 `sum - root.val` 的路径，即 `hasPathSum(root.left, sum - root.val) || hasPathSum(root.right, sum - root.val) `。
## 二、具体实现

``` Java
class Solution {
    public boolean hasPathSum(TreeNode root, int sum) {
        if (root == null) {
            return false;
        }
        // 到达叶子节点时，递归终止，判断 sum 是否符合条件。
        if (root.left == null && root.right == null) {
            return root.val == sum;
        }
        // 递归地判断root节点的左孩子和右孩子。
        return hasPathSum(root.left, sum - root.val) || hasPathSum(root.right, sum - root.val);
    }
}
```
