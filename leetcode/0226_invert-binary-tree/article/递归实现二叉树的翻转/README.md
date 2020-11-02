"+++
title = "0226. Invert Binary Tree 递归实现二叉树的翻转 "
author = ["zhangyuanyuan-h"]
date = 2020-09-16T01:35:25+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 递归实现二叉树的翻转

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [递归实现二叉树的翻转](https://leetcode-cn.com/problems/invert-binary-tree/solution/di-gui-shi-xian-er-cha-shu-de-fan-zhuan-by-zhangyu/) by [zhangyuanyuan-h](https://leetcode-cn.com/u/zhangyuanyuan-h/)

### 解题思路
用时100，内存86
1.从root开始遍历二叉树；
2.只要当前节不为空，交换两个子节点；
3.对子节点重复上述步骤
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

    //创建一个全局变量作为交换子节点的中间节点
    TreeNode node;
    public TreeNode invertTree(TreeNode root) {
        node = new TreeNode();

        //执行翻转方法翻转二叉树
        reverseTree(root);

        //返回翻转后的二叉树的根节点
        return root;
    }

    public void reverseTree(TreeNode root){

        //如果当前节点为空直接返回（递归的终止条件）
        if(root==null) return;

        //判断当前节点不为空之后交换该节点的左右两个子节点
        node = root.right;
        root.right = root.left;
        root.left = node;

        //对当前节点的子节点进行翻转
        reverseTree(root.left);
        reverseTree(root.right);

    }
}
```