"+++
title = "0968. Binary Tree Cameras 监控二叉树状态查询法 "
author = ["w1sl1y"]
date = 2019-05-13T09:38:05+08:00
tags = ["Leetcode", "Algorithms", "Recursion", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# 监控二叉树状态查询法

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [监控二叉树状态查询法](https://leetcode-cn.com/problems/binary-tree-cameras/solution/jian-kong-er-cha-shu-zhuang-tai-cha-xun-fa-by-w1sl/) by [w1sl1y](https://leetcode-cn.com/u/w1sl1y/)

当遍历到一个节点时，我们可以定义三种状态：
0 ： 初始状态，如果节点为null可以返回，也就是不影响其他节点，当两个节点都是0时，我们直接设置当前节点为未监控状态
1： 未监控状态，如果子节点含有该状态，则此节点必须添加摄像头，同时返回当前状态为监控态
2： 监控态，表明此节点已经被监控，当子节点为此状态时，父节点不需要添加摄像头，可以返回初始态

```
private int dfs(TreeNode node){
        if (node == null) return 0;

        int l = dfs(node.left);
        int r = dfs(node.right);

        if (l + r == 0)  
            return 1;
        else if (l == 1 || r == 1) {
            cameras ++; return 2;
        } else  
            return 0;
    }
```
当调用时，有一个小技巧，我们需要为传入的根节点添加一个虚拟的头，因为向上遍历时，根节点的监控状态我们无法保证，所以添加一个虚拟头可以简化编程。
```
public int minCameraCover(TreeNode root) {
        TreeNode dummyHead = new TreeNode(0);
        dummyHead.left = root;
        dfs(dummyHead);
        return cameras;
    }
```

时间复杂度O(N)
空间复杂度如果不算递归的隐式调用栈，为O(1),否则为O(h),h为树的高度。