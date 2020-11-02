"+++
title = "0250. Count Univalue Subtrees 250. 统计同值子树 "
author = ["klb"]
date = 2020-09-13T06:34:55+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 250. 统计同值子树

> [0250. Count Univalue Subtrees](https://leetcode-cn.com/problems/count-univalue-subtrees/)
> [250. 统计同值子树](https://leetcode-cn.com/problems/count-univalue-subtrees/solution/250-tong-ji-tong-zhi-zi-shu-by-klb/) by [klb](https://leetcode-cn.com/u/klb/)

### 解题思路

使用深度优先搜索遍历。

首先，叶子节点是同值子树；

对于非叶子节点：如果存在左右子树，且左右子树都是同值子树，且当前节点的值和左右子树的值一样，则当前节点为根节点的树也是同值子树。

### 代码

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    int count = 0;

    public int countUnivalSubtrees(TreeNode root) {
        if (root == null) return 0;
        is_uni(root);
        return count;
    }

    /**
     * 判断以 node 为根节点的树是否为同值子树
     *
     * @param node
     * @return
     */
    boolean is_uni(TreeNode node) {
        // 叶子节点是同值子树
        if (node.left == null && node.right == null) {
            count++;
            return true;
        }
        // 非叶子节点的情况
        boolean is_unival = true;
        if (node.left != null) {
            // 左子节点不为空，且左子节点是同值子树，且左子节点的值和 node 一样
            is_unival = is_uni(node.left) && node.left.val == node.val;
        }
        if (node.right != null) {
            // 右子节点不为空，右子节点是同值子树，node 和右子节点的值相同
            is_unival = is_uni(node.right) && is_unival && node.right.val == node.val;
        }
        if (!is_unival) return false;
        count++;
        return true;
    }
}
```