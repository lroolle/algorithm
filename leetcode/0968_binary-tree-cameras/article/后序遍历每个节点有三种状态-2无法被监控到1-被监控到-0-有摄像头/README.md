"+++
title = "0968. Binary Tree Cameras 后序遍历，每个节点有三种状态: 2无法被监控到，1 被监控到,   0 有摄像头 "
author = ["jane-73"]
date = 2020-02-15T08:10:46+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# 后序遍历，每个节点有三种状态: 2无法被监控到，1 被监控到,   0 有摄像头

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [后序遍历，每个节点有三种状态: 2无法被监控到，1 被监控到,   0 有摄像头](https://leetcode-cn.com/problems/binary-tree-cameras/solution/hou-xu-bian-li-mei-ge-jie-dian-you-san-chong-zhuan/) by [jane-73](https://leetcode-cn.com/u/jane-73/)

### 解题思路
执行用时 :0 ms, 在所有 Java 提交中击败了100.00%的用户
内存消耗 :47.5 MB, 在所有 Java 提交中击败了5.26%的用户

如果左右子节点都无法被监控到 则当前放置一个摄像头。
如果左右子节点能够被监控到，则当前返回无法被监控
如果左右子节点有一个有摄像头，则当前返回可以被监控到。
叶子节点始终返回无法被监控到。  叶子节点不能放摄像头

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
    int ret = 0;
    public int minCameraCover(TreeNode root) {
        if (root == null) {
            return 0;
        }
        if (root.left == null && root.right == null) {
            return 1;
        }
        ret = 0;
        int status = see(root);
        if (status == 2) {
            ret++;
        }
        return ret;
    }

    public int see(TreeNode node) {

        if (node.left == null && node.right == null) {
            return 2;
        }

        int left = -1;
        if (node.left != null) {
            left = see(node.left);
        }

        int rigth = -1;
        if (node.right != null) {
            rigth = see(node.right);
        }

        if (left == 2 || rigth == 2) {
            ret++;
            return 0;
        }

        if ((left == 1 && rigth == 1) || (left == 1 && rigth == -1) || (rigth == 1 && left == -1)) {
            return 2;
        } else if (left == 0 || rigth == 0) {
            return 1;
        }
        return 0;
    }
}
```