"+++
title = "0094. Binary Tree Inorder Traversal 学习笔记 94 - 1 "
author = ["wangsidi"]
date = 2020-09-14T06:19:22+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 学习笔记 94 - 1

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [学习笔记 94 - 1](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/xue-xi-bi-ji-94-1-by-wangsidi/) by [wangsidi](https://leetcode-cn.com/u/wangsidi/)

### 解题思路
方法一、 递归

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
 std::vector<int> arr;
class Solution {
public:
    void pushVector(TreeNode* root, std::vector<int> &arr){
        //结束点
        if(root == nullptr) return;
        //左
        pushVector(root -> left, arr);
        //opt
        arr.push_back(root -> val);
        //右
        pushVector(root -> right, arr);
    }
    vector<int> inorderTraversal(TreeNode* root) {
        std::vector<int> arr;

        pushVector(root, arr);

        return arr;
    }
};
```