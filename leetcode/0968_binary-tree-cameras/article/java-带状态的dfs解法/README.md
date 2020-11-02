"+++
title = "0968. Binary Tree Cameras Java  带状态的dfs解法 "
author = ["don-vito-corleone"]
date = 2020-04-12T12:50:52+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# Java  带状态的dfs解法

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [Java  带状态的dfs解法](https://leetcode-cn.com/problems/binary-tree-cameras/solution/java-dai-zhuang-tai-de-dfsjie-fa-by-don-vito-corle/) by [don-vito-corleone](https://leetcode-cn.com/u/don-vito-corleone/)

### 解题思路
节点一共有三种状态，MONITOR 代表这个节点是一个监视器

NEED_PARENT代表这个节点需要一个监视器来监视

NO_NEED 代表这个节点不需要监视器来监视

然后从叶子节点开始递归， 当某个节点的左节点或者右节点需要监控器，那就令这个节点为MONITOR状态

如果左节点或右节点有一个是MONITOR状态，那这个节点就不需要监控器，为NO_NEED状态

否则返回NEED_PARMENT状态
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
    int NEED_PARENT = 0;
    int NO_NEED = 1;
    int MONITOR = 2;
    int res;
    public int minCameraCover(TreeNode root) {
        if (dfs(root) == NEED_PARENT) {
            res ++;
        }
        return res;
    }

    int dfs(TreeNode root) {
        if (root == null)  {
            return NO_NEED;
        }
        int left = dfs(root.left);
        int right = dfs(root.right);
        if (left == NEED_PARENT || right == NEED_PARENT) {
            res ++;
            return MONITOR;
        }
        if (left == MONITOR || right == MONITOR) {
            return NO_NEED;
        }
        return NEED_PARENT;
    }
}
```