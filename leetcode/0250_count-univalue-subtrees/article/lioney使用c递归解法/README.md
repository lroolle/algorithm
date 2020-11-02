"+++
title = "0250. Count Univalue Subtrees lioney:)使用C++递归解法 "
author = ["lioney"]
date = 2019-11-21T07:45:50+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# lioney:)使用C++递归解法

> [0250. Count Univalue Subtrees](https://leetcode-cn.com/problems/count-univalue-subtrees/)
> [lioney:)使用C++递归解法](https://leetcode-cn.com/problems/count-univalue-subtrees/solution/shi-yong-cdi-gui-jie-fa-by-lioney/) by [lioney](https://leetcode-cn.com/u/lioney/)


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
    int countUnivalSubtrees(TreeNode* root) {
        int res = 0;
        helper(root, res);
        return res;
    }
    
    bool helper(TreeNode* root, int& res) {
        if(!root) return true;
        bool left = helper(root->left, res);
        bool right = helper(root->right, res);
        if(root->left && root->left->val != root->val || root->right && root->right->val != root->val)
            return false;
        if(left && right) res += 1;
        return true;
    }
};

