"+++
title = "0968. Binary Tree Cameras go "
author = ["huan-le-ma-77"]
date = 2020-04-24T13:45:03+08:00
tags = ["Leetcode", "Algorithms", "Go"]
categories = ["leetcode"]
draft = false
+++

# go

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [go](https://leetcode-cn.com/problems/binary-tree-cameras/solution/go-by-huan-le-ma-77/) by [huan-le-ma-77](https://leetcode-cn.com/u/huan-le-ma-77/)

### 解题思路
此处撰写解题思路

### 代码

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res int

func minCameraCover(root *TreeNode) int {

    res = 0
    if dfs(root) == 1 {
        // 如果最顶端的元素没有被安装，也没有被监视
        // 那么这个节点需要安装一个摄像头
        res ++
    }
    return res;
}

// 1未被监视, 2安装摄像头
func dfs(root *TreeNode) int {

    // 最后一个节点不关心
    if root == nil {
        return 0
    }

    left := dfs(root.Left)
    right := dfs(root.Right)

    // 如果子节点没有被监视
    if left == 1 || right == 1 {
        // 安装摄像头
        res++
        return 2
    }

    // 如果子节点安装摄像头
    if left == 2 || right == 2 {
        // 标记被监视
        return 3
    }
    return 1
}
```