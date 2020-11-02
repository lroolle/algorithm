"+++
title = "0230. Kth Smallest Element in a BST Go 遍历到数组中存储 "
author = ["Chelsea"]
date = 2020-09-05T05:26:11+08:00
tags = ["Leetcode", "Algorithms", "Go"]
categories = ["leetcode"]
draft = false
+++

# Go 遍历到数组中存储

> [0230. Kth Smallest Element in a BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)
> [Go 遍历到数组中存储](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/solution/go-bian-li-dao-shu-zu-zhong-cun-chu-by-chelsea/) by [Chelsea](https://leetcode-cn.com/u/chelsea/)

```
func kthSmallest(root *TreeNode, k int) int {

	arr := []int{}
	find(root, &arr)

	return arr[k-1]

}
func find(root *TreeNode, arr *[]int) {

	if root == nil {
		return
	}

	find(root.Left, arr)
	*arr = append(*arr, root.Val)
	find(root.Right, arr)
	return

}
```