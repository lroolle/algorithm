"+++
title = "0094. Binary Tree Inorder Traversal 递归，耗时100%；简单易懂 "
author = ["ddd-12f"]
date = 2020-09-14T06:23:02+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 递归，耗时100%；简单易懂

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [递归，耗时100%；简单易懂](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/di-gui-hao-shi-100jian-dan-yi-dong-by-ddd-12f/) by [ddd-12f](https://leetcode-cn.com/u/ddd-12f/)

### 解题思路
不知各位大佬的内存使用如何。
首先必须明白什么是中序遍历（可BaiDu，打钱打钱）
其次就是按照中序遍历定义的步骤，依次递归执行。

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
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        if (root != NULL) {
            auto temp = (inorderTraversal(root->left));
            res.insert(res.end(),temp.begin(),temp.end());
            res.push_back(root->val);
            temp = (inorderTraversal(root->right));
            res.insert(res.end(),temp.begin(),temp.end());
        }
        return res;
    }
};
```