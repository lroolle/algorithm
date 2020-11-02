"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 【Java】简单易懂地 解决 “最近公共祖先” 问题 "
author = ["leetcoder-youzg"]
date = 2020-09-27T03:00:47+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 【Java】简单易懂地 解决 “最近公共祖先” 问题

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [【Java】简单易懂地 解决 “最近公共祖先” 问题](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/java-jian-dan-yi-dong-di-jie-jue-zui-jin-gong-gong/) by [leetcoder-youzg](https://leetcode-cn.com/u/leetcoder-youzg/)

### 解题思路
本题解 借鉴了 **官方题解** 的思想：
> 利用 二叉搜索树 的性质：
> 左 > 根 > 右
> 我们就能判断 **`p`** 和 **`q`**，相对于 **`当前root`** 的位置
>   1、若p、q 均位于 当前root 的左子树，则让 当前root指向它的**左孩子**
>   2、若p、q 均位于 当前root 的右子树，则让 当前root指向它的**右孩子**
>   3、若p、q 分别位于 当前root 的一左一右，则 **当前root** 就是所求的 **`最近公共祖先`**

### 运行结果
![image.png](https://pic.leetcode-cn.com/1601175215-vxCSRv-image.png)

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
        while (true) {
            if (p.val < root.val && q.val < root.val) { // p、q均在 当前root 的左子树
                root = root.left;
            } else if (p.val > root.val && q.val > root.val) { // p、q均在 当前root 的右子树
                root = root.right;
            } else {    // 当前root 就是 p、q的 共同祖先，结束 迭代
                break;
            }
        }
        return root;
    }
}
```
打卡第68天，加油！！！