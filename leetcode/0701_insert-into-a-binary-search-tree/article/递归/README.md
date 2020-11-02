"+++
title = "0701. Insert into a Binary Search Tree 递归 "
author = ["linbingyuan"]
date = 2020-04-20T12:59:34+08:00
tags = ["Leetcode", "Algorithms", "Go", "DivideandConquer"]
categories = ["leetcode"]
draft = false
+++

# 递归

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [递归](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/di-gui-by-linbingyuan-3/) by [linbingyuan](https://leetcode-cn.com/u/linbingyuan/)

### 解题思路
1. 对于节点处理的过程，就把递归当作一种分治思想，上层问题分给下层处理，下层结果返回至上层，
2. 例如 [450. 删除二叉搜索树中的节点](https://leetcode-cn.com/problems/delete-node-in-a-bst/) ,[669. 修剪二叉搜索树](https://leetcode-cn.com/problems/trim-a-binary-search-tree/) 都是一样的思路，

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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		if root.Left != nil {
			insertIntoBST(root.Left, val)
		} else {
			root.Left = &TreeNode{Val: val}
		}
	} else {
		if root.Right != nil {
			insertIntoBST(root.Right, val)
		} else {
			root.Right = &TreeNode{Val: val}
		}
	}

	return root
}
```
[![](https://pic.leetcode-cn.com/9251550b5598374757497f02f0edb32d32e2868ef21f63c096dc5b7015da9c01.svg)](https://github.com/bygo/leetcode)

[![](https://pic.leetcode-cn.com/4173da0f2e61f69ed742c363fed66ff7900bb5a2d17994de1f0bafde0155b3e6.svg)](https://www.yuque.com/daizhuansheng/lc/cfg57t)