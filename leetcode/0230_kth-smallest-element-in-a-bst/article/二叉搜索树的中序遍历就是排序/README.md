"+++
title = "0230. Kth Smallest Element in a BST 二叉搜索树的中序遍历就是排序 "
author = ["aff0aoHk97"]
date = 2020-07-26T02:09:29+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 二叉搜索树的中序遍历就是排序

> [0230. Kth Smallest Element in a BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)
> [二叉搜索树的中序遍历就是排序](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/solution/er-cha-sou-suo-shu-de-zhong-xu-bian-li-jiu-shi-pai/) by [aff0aoHk97](https://leetcode-cn.com/u/aff0aoHk97/)

```
class Solution {
public:
    vector<int>order;
    void dfs(TreeNode* root)
    {
        if(root==nullptr)
            return;
        dfs(root->left);
        order.push_back(root->val);
        dfs(root->right);
    }
    int kthSmallest(TreeNode* root, int k) {
        dfs(root);
        return order[k-1];
        
    }
};
```
