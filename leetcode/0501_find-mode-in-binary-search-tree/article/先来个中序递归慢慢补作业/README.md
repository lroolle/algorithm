"+++
title = "0501. Find Mode in Binary Search Tree 先来个中序递归，慢慢补作业 "
author = ["linbingyuan"]
date = 2020-04-14T14:35:10+08:00
tags = ["Leetcode", "Algorithms", "Go", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# 先来个中序递归，慢慢补作业

> [0501. Find Mode in Binary Search Tree](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)
> [先来个中序递归，慢慢补作业](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/xian-lai-ge-zhong-xu-di-gui-man-man-bu-zuo-ye-by-l/) by [linbingyuan](https://leetcode-cn.com/u/linbingyuan/)

# dfs 递归

```golang
var max int     //最大值
var res []int   //结果
var cur int     //当前
var counter int //当前计数

func findMode(root *TreeNode) []int {
	res, max, cur, counter = []int{}, 1, 0, 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		if root.Val != cur {
			counter = 0
		}
		counter++
		if max < counter {
			max = counter
			res = []int{root.Val}
		} else if max == counter {
			res = append(res, root.Val)
		}
		cur = root.Val
		dfs(root.Right)
	}
}
```
[![](https://pic.leetcode-cn.com/9251550b5598374757497f02f0edb32d32e2868ef21f63c096dc5b7015da9c01.svg)](https://github.com/temporaries/leetcode)

[![](https://pic.leetcode-cn.com/4173da0f2e61f69ed742c363fed66ff7900bb5a2d17994de1f0bafde0155b3e6.svg)](https://www.yuque.com/daizhuansheng/lc/cfg57t)