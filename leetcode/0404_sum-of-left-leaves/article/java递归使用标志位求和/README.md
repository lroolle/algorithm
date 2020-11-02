"+++
title = "0404. Sum of Left Leaves Java递归，使用标志位求和 "
author = ["ha-ha-ha-192"]
date = 2020-07-01T07:57:43+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree"]
categories = ["leetcode"]
draft = false
+++

# Java递归，使用标志位求和

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [Java递归，使用标志位求和](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/javadi-gui-shi-yong-biao-zhi-wei-qiu-he-by-ha-ha-h/) by [ha-ha-ha-192](https://leetcode-cn.com/u/ha-ha-ha-192/)

本题要求，求所有树的左叶子节点，这里抓关键字，一定要是左孩子并且是叶子节点
如何判断叶子节点很简单`root.left == null && root.right == null`
那么如何判断是左孩子呢？
这里我使用一个**标志位**判断叶子节点

具体是这样的写的
当你写递归函数的时候是不是递归左孩子`sumOfLeftLeaves(root.left);`
递归右孩子`sumOfLeftLeaves(root.right);`
这里就非常明显了，递归左孩子是不是说明这个节点一定是左孩子
我们就可以将代码这样改写`sumOfLeftLeaves(root.left, true);`,`sumOfLeftLeaves(root.right, false);`

参考代码
```java
class Solution {
    public int sumOfLeftLeaves(TreeNode root) {
        return sumOfLeftLeaves(root, false);
    }
    public int sumOfLeftLeaves(TreeNode root, boolean flag){
        if(root == null){
            return 0;
        }
        if(flag && root.left == null && root.right == null){
            return root.val;
        }
        int leftSum = sumOfLeftLeaves(root.left, true);
        int rightSum = sumOfLeftLeaves(root.right, false);
        return leftSum + rightSum;
    }
}
```