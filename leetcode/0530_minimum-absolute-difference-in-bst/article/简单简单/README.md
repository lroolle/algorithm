"+++
title = "0530. Minimum Absolute Difference in BST 简单简单 "
author = ["Wuzuobin"]
date = 2020-09-17T12:33:40+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 简单简单

> [0530. Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
> [简单简单](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/solution/jian-dan-jian-dan-by-cai-bi-d/) by [Wuzuobin](https://leetcode-cn.com/u/wuzuobin/)

### 解题思路
此处撰写解题思路

### 代码

```cpp
class Solution 
{
public:
	void inVector(TreeNode* root, vector<int>& v)
	{
		if(root == NULL)
		{
			return ;
		}
		
		inVector(root->left, v);
		v.push_back(root->val);
		inVector(root->right,v);
	}
    int getMinimumDifference(TreeNode* root) 
	{
		
		vector<int> v;
		inVector(root, v);
		//sort(v.begin(), v.end());
		int res = INT_MAX;
		for(int i = 0; i < v.size()-1; i++)
		{
			res = min(res, v[i+1]-v[i]);
		} 
		return res;
    }
};
```