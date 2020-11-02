"+++
title = "0404. Sum of Left Leaves 404.左叶子之和 "
author = ["Geralt_U"]
date = 2020-04-16T01:43:13+08:00
tags = ["Leetcode", "Algorithms", "Java", "二叉树", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 404.左叶子之和

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [404.左叶子之和](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/404zuo-xie-zi-zhi-he-by-ming-zhi-shan-you-m9rfkvkd/) by [Geralt_U](https://leetcode-cn.com/u/geralt_u/)

### 解题思路
题解：本题同样需要根据树形结构的基本思路“遍历”，来解决。本题需要计算二叉树所有左叶子之和，只需在遍历的过程中判断是否到达了左叶子，在这里到达左叶子的判断标准为：

- 是否是当前节点的左孩子
- 当前节点的左孩子是不是叶子节点（叶子结点：没有左右孩子）

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
    public int result = 0;
    public int sumOfLeftLeaves(TreeNode root) {
        if(root == null){
            return 0;
        }
        //是否是当前节点的左孩子
        //当前节点的左孩子是不是叶子节点（叶子结点：没有左右孩子）
        if(root.left != null && (root.left.left == null && root.left.right == null)){
            result = result +root.left.val;
        }
        sumOfLeftLeaves(root.left);
        sumOfLeftLeaves(root.right);
        return result;
    }
}
```