"+++
title = "0113. Path Sum II 优化DFS，简洁代码 （100） "
author = ["dxian-sen-2"]
date = 2020-05-17T09:51:33+08:00
tags = ["Leetcode", "Algorithms", "C++", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 优化DFS，简洁代码 （100）

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [优化DFS，简洁代码 （100）](https://leetcode-cn.com/problems/path-sum-ii/solution/you-hua-dfsjian-ji-dai-ma-100-by-dxian-sen-2/) by [dxian-sen-2](https://leetcode-cn.com/u/dxian-sen-2/)

一个简单的自顶向下DFS，但是有些细节需要注意
```
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> res;
    vector<int> temp;//防止反复初始化数组 
    void dfs(TreeNode* root,int sum){
        int resum=sum-root->val;
        temp.push_back(root->val);
        if(resum==0&&!root->left&&!root->right)
			res.push_back(temp);//找到答案 
        if(root->left)
			dfs(root->left,resum);
        if(root->right)
			dfs(root->right,resum);
        temp.pop_back();//回溯 
    }
    vector<vector<int>> pathSum(TreeNode* root, int sum) {
        if(root)dfs(root,sum);
        return res;
    }
};
```
建议不要把temp当做参数传入