"+++
title = "0701. Insert into a Binary Search Tree C++ 中规中矩的96ms解法（dfs + bfs） "
author = ["Gary_coding"]
date = 2020-06-08T15:40:26+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 中规中矩的96ms解法（dfs + bfs）

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [C++ 中规中矩的96ms解法（dfs + bfs）](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/c-zhong-gui-zhong-ju-de-96msjie-fa-dfs-by-gary_cod/) by [Gary_coding](https://leetcode-cn.com/u/gary_coding/)

### dfs
```cpp
class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if (root == nullptr) {
            TreeNode* ans = new TreeNode(val);
            return ans;
        }
        if (val < root->val) root->left = insertIntoBST(root->left, val);
        else root->right = insertIntoBST(root->right, val);
        return root;
    }
};
```
### bfs
```cpp
class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if (root == nullptr) return new TreeNode(val);
        TreeNode* head = root;
        while (root) {
            if (val < root->val) {
                if (root->left == nullptr) {
                    root->left = new TreeNode(val);
                    return head;
                }
                else root = root->left;
            }
            else {
                if (root->right == nullptr) {
                    root->right = new TreeNode(val);
                    return head;
                }
                else root = root->right;
            }
        }
        return head;
    }
};
```