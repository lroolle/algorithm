"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 【5月第4篇题解】分类讨论 递归 "
author = ["Qtaro"]
date = 2020-05-13T12:22:52+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 【5月第4篇题解】分类讨论 递归

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [【5月第4篇题解】分类讨论 递归](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/5yue-di-4pian-ti-jie-fen-lei-tao-lun-di-gui-by-qta/) by [Qtaro](https://leetcode-cn.com/u/qtaro/)

### 解题思路
![235.jpg](https://pic.leetcode-cn.com/5248f41cda0047c0723f8bdd49dd34a036056ee216e11da195fa4a623824cc21-235.jpg)
### 代码

```cpp
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
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (!root) return nullptr;
        if (p->val < root->val && q->val < root->val)
            return lowestCommonAncestor(root->left, p, q);
        if (p->val > root->val && q->val > root->val)
            return lowestCommonAncestor(root->right, p, q);
        return root;
    }
};
```