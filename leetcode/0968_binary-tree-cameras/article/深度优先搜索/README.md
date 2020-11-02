"+++
title = "0968. Binary Tree Cameras 深度优先搜索 "
author = ["hundanLi"]
date = 2019-11-17T08:20:02+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 深度优先搜索

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [深度优先搜索](https://leetcode-cn.com/problems/binary-tree-cameras/solution/shen-du-you-xian-sou-suo-by-hundanli-6/) by [hundanLi](https://leetcode-cn.com/u/hundanli/)

```java
class Solution {
    private int num;
    public int minCameraCover(TreeNode root) {
        if (dfs(root) == 3) {
            num++;
        }
        return num;
    }

    /** 1.安装监视器；2.被监视；3.未被监视 */
    private int dfs(TreeNode root) {
        if (root == null) {
            return 2;
        }
        int left = dfs(root.left);
        int right = dfs(root.right);
        if (left == 3 || right == 3) {
            // 子节点未被监视，因此需要安装监视器
            num++;
            return 1;
        }
        if (left == 1 || right == 1) {
            // 子节点安装了监视器，因此已经被监视
            return 2;
        }
        // 当前未被监视
        return 3;
    }
}
```
