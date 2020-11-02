"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree C++ 中规中矩的40ms解法（时间O(logN)） "
author = ["Gary_coding"]
date = 2020-04-23T15:46:07+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 中规中矩的40ms解法（时间O(logN)）

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [C++ 中规中矩的40ms解法（时间O(logN)）](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/c-zhong-gui-zhong-ju-de-40msjie-fa-shi-jian-ologn-/) by [Gary_coding](https://leetcode-cn.com/u/gary_coding/)

### 二叉搜索树的公共祖先
```cpp
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == nullptr) return nullptr;
        if (p->val < root->val && q->val < root->val) return lowestCommonAncestor(root->left, p, q);
        if (p->val > root->val && q->val > root->val) return lowestCommonAncestor(root->right, p, q);
        return root;
    }
};
```
### 二叉树的公共祖先
```cpp
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == nullptr || root == p || root == q) return root;
        TreeNode* l = lowestCommonAncestor(root->left, p, q);
        TreeNode* r = lowestCommonAncestor(root->right, p, q);
        if (l && r) return root;
        return l ? l : r;
    }
};
```