"+++
title = "0501. Find Mode in Binary Search Tree 写出来真不容易呀 "
author = ["cai-bi-d"]
date = 2020-09-17T12:01:58+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 写出来真不容易呀

> [0501. Find Mode in Binary Search Tree](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)
> [写出来真不容易呀](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/xie-chu-lai-zhen-bu-rong-yi-ya-by-cai-bi-d/) by [cai-bi-d](https://leetcode-cn.com/u/cai-bi-d/)

### 解题思路
此处撰写解题思路

### 代码

```cpp
class Solution {
public:
	void findMore(TreeNode* root, vector<int>& v)//先把所有的节点值都放到一个数组里面
	{
		if(root == NULL)
		{
			return ;
		}
		v.push_back(root->val);
		findMore(root->left, v);
		findMore(root->right,v);
	}
	vector<int> findNumber(vector<int> &v)//从数组里找到众数们
	{
		map<int,int>mp;
		map<int,bool>flag;
		int max = 0;
		vector<int> res;
		for(int i = 0; i < v.size(); i++)
		{
			mp[v[i]]++;
			flag[v[i]] = true;
			if(mp[v[i]] > max)
			{
				max = mp[v[i]];
			}
		}
		for(int i = 0; i < v.size(); i++)
		{
			if(mp[v[i]] == max && flag[v[i]] == true)
			{
				res.push_back(v[i]);
				flag[v[i]] = false;
			}
		}
		return res;
	}
    vector<int> findMode(TreeNode* root) //调用前两个出结果啦
	{
		vector<int> res1;
		if(root == NULL)
		{
			return res1;
		}
		findMore(root, res1);
		return findNumber(res1);
    }
};
```