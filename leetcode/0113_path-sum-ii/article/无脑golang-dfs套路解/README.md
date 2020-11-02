"+++
title = "0113. Path Sum II 无脑Golang DFS套路解 "
author = ["ukcuf"]
date = 2020-08-02T11:34:23+08:00
tags = ["Leetcode", "Algorithms", "Go"]
categories = ["leetcode"]
draft = false
+++

# 无脑Golang DFS套路解

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [无脑Golang DFS套路解](https://leetcode-cn.com/problems/path-sum-ii/solution/wu-nao-golang-dfstao-lu-jie-by-ukcuf/) by [ukcuf](https://leetcode-cn.com/u/ukcuf/)

### 解题思路
看着确实很无脑

### 代码

```golang
func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	path := []int{}
	dfs(&ret, root, path, sum)
	return ret
}

func dfs(ret *[][]int, root *TreeNode, path []int, target int) {
	switch {
	case root == nil:
		return
	case root.Left == nil && root.Right == nil && root.Val == target:
		dst := make([]int, len(path)+1)
		copy(dst, append(path, root.Val))
		*ret = append(*ret, dst)
		return
	}
	path = append(path, root.Val)
	dfs(ret, root.Left, path, target-root.Val)
	dfs(ret, root.Right, path, target-root.Val)
}
```