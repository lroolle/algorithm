"+++
title = "0112. Path Sum 递归来做，用sum减去当前的val值 "
author = ["Eng_not_Ape"]
date = 2020-09-04T07:32:40+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# 递归来做，用sum减去当前的val值

> [0112. Path Sum](https://leetcode-cn.com/problems/path-sum/)
> [递归来做，用sum减去当前的val值](https://leetcode-cn.com/problems/path-sum/solution/di-gui-lai-zuo-yong-sumjian-qu-dang-qian-de-valzhi/) by [Eng_not_Ape](https://leetcode-cn.com/u/eng_not_ape/)

### 解题思路
此处撰写解题思路

### 代码

```c
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

bool hasPathSum(struct TreeNode* root, int sum){
    if(root==NULL)  return false;
    sum-=root->val;
    if(root->left==NULL&&root->right==NULL){
        if(sum==0)    return true;
        else    return false;
    }
    return hasPathSum(root->left,sum)||hasPathSum(root->right,sum); 
}
```