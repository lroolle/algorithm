"+++
title = "0530. Minimum Absolute Difference in BST 中序遍历，保存前一节点 "
author = ["qi-xi-5"]
date = 2020-05-31T11:59:34+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 中序遍历，保存前一节点

> [0530. Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
> [中序遍历，保存前一节点](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/solution/zhong-xu-bian-li-bao-cun-qian-yi-jie-dian-by-qi-xi/) by [qi-xi-5](https://leetcode-cn.com/u/qi-xi-5/)

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
    // 初始化最小值
    int min = Integer.MAX_VALUE;
    // 前一个节点，初始化为null
    TreeNode pre = null;
    public int getMinimumDifference(TreeNode root) {
        pre(root);
        return min;
    }
    public void pre(TreeNode root){
        if(root == null)
            return ;
        // 中序遍历
        pre(root.left);
        // 判空，寻找最小差值
        if(pre != null)
            min = Math.min(min,root.val - pre.val);
        // 将此节点设置为前一节点
        pre = root;
        pre(root.right);
    }
}
```