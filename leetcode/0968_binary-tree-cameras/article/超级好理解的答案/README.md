"+++
title = "0968. Binary Tree Cameras 超级好理解的答案 "
author = ["levyjeng"]
date = 2020-04-11T02:07:54+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 超级好理解的答案

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [超级好理解的答案](https://leetcode-cn.com/problems/binary-tree-cameras/solution/chao-ji-hao-li-jie-de-da-an-by-levyjeng/) by [levyjeng](https://leetcode-cn.com/u/levyjeng/)

![image.png](https://pic.leetcode-cn.com/d932a52a73034b7ac055da7205b1b25bf23ce8768757edc6aee2207916aac59f-image.png)

### 解题思路
有三个状态:
0=>这个结点待覆盖
1=>这个结点已经覆盖
2=>这个结点上安装了相机

### 代码

```java
 class Solution {
   public int minCameraCover(TreeNode root) {
		if (lrd(root) == 0) {
			res++;
		}
		return res;
	}

	int res=0;
	int lrd(TreeNode node) {

		if (node == null) {
			return 1;
		}
		int left=lrd(node.left);
		int right=lrd(node.right);
		if (left == 0 || right == 0) {
			res++;
			return 2;
		}
		if(left==1&&right==1){
			return 0;
		}
		if(left+right>=3){
			return 1;
		}
	
		return -1;
	}
}
```