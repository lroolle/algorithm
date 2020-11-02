"+++
title = "0250. Count Univalue Subtrees C语言 双百 "
author = ["cheng-xian-sheng-3"]
date = 2019-11-13T13:45:24+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# C语言 双百

> [0250. Count Univalue Subtrees](https://leetcode-cn.com/problems/count-univalue-subtrees/)
> [C语言 双百](https://leetcode-cn.com/problems/count-univalue-subtrees/solution/cyu-yan-shuang-bai-by-cheng-xian-sheng-3/) by [cheng-xian-sheng-3](https://leetcode-cn.com/u/cheng-xian-sheng-3/)

简单递归
```
int g_sum = 0;
int IsUnivaltree(struct TreeNode* root){
    if (root == NULL) {
        return 1;
    }
    int left = IsUnivaltree(root->left);
    int right = IsUnivaltree(root->right);

    if (root->left != NULL && (left == 0 || root->val != root->left->val)) {
        return 0;
    } 
    if (root->right != NULL && (right == 0 || root->val != root->right->val)) {
        return 0;
    }
    g_sum++;
    return 1;
}
int countUnivalSubtrees(struct TreeNode* root){
    if (root == NULL) {
        return 0;
    }
    g_sum = 0;
    IsUnivaltree(root);
    return g_sum;
}
```
