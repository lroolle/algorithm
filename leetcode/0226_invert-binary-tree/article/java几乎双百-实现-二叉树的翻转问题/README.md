"+++
title = "0226. Invert Binary Tree 【Java】几乎双百 实现 ”二叉树的翻转“问题 "
author = ["leetcoder-youzg"]
date = 2020-09-16T02:52:53+08:00
tags = ["Leetcode", "Algorithms", "Java", "Backtracking", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 【Java】几乎双百 实现 ”二叉树的翻转“问题

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [【Java】几乎双百 实现 ”二叉树的翻转“问题](https://leetcode-cn.com/problems/invert-binary-tree/solution/java-ji-hu-shuang-bai-shi-xian-er-cha-shu-de-fan-z/) by [leetcoder-youzg](https://leetcode-cn.com/u/leetcoder-youzg/)

### 解题思路
本题的解题思路也非常简单：
> 递归 交换 每一个节点的 左右子节点

### 运行结果
![image.png](https://pic.leetcode-cn.com/1600223903-vaajTj-image.png)

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
    public TreeNode invertTree(TreeNode root) {
        dfs(root);
        return root;
    }

    private void dfs(TreeNode curRoot) {
        if (curRoot == null) {
            return;
        }

        TreeNode leftNode = curRoot.left;
        TreeNode rightNode = curRoot.right;
        if (leftNode != null || rightNode != null) {
            TreeNode tmpNode = leftNode;
            curRoot.left = curRoot.right;
            curRoot.right = leftNode;
            dfs(leftNode);
            dfs(rightNode);
        }
    }
}
```
打卡第57天
昨天玩了[《Doki Doki Literature Club!》](https://ddlc.moe)，一款 **颠覆了我对galgame认知** 的游戏
(游戏很小，就200多M，Steam上下载的话，是免费的)
(无法点开网页，或者 没找到下载路径的同学，请前往Steam进行下载，当然，汉化包需要自行搜索下载)
![image.png](https://pic.leetcode-cn.com/1600224278-MRBoPf-image.png)
女主最后唱的bgm一出，眼泪差点没忍住
(bgmURL:[https://music.163.com/#/song?id=523658880](https://music.163.com/#/song?id=523658880))
如果感兴趣的同学，玩到一半是致郁，玩到结束是 **究极治愈**
请一定要玩到最后，否则，不会明白这个游戏真正的精髓

(心里承受能力较弱的同学请忽视!!!)
(心里承受能力较弱的同学请忽视!!!)
(心里承受能力较弱的同学请忽视!!!)
