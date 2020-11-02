"+++
title = "0701. Insert into a Binary Search Tree 701. 二叉搜索树中的插入操作 C++ "
author = ["eric-345"]
date = 2020-02-21T07:05:00+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 701. 二叉搜索树中的插入操作 C++

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [701. 二叉搜索树中的插入操作 C++](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/c-by-eric-345-48/) by [eric-345](https://leetcode-cn.com/u/eric-345/)

### 解题思路
### 代码

```cpp

class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        
        insert(root, val);
        return root;
    }

    // 传引用的方式
    void insert(TreeNode*& node, int val)
    {
        if(node == nullptr)
            node = new TreeNode(val);
        
        if(val < node->val)
            insert(node->left, val);

        else if (val > node->val)
            insert(node->right, val);
    }
    // 返回值的方式
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        
        if(root == nullptr)
            root = new TreeNode(val);
        
        if(val < root->val)
            root->left = insertIntoBST(root->left, val);
        
        else if(val > root->val)
            root->right = insertIntoBST(root->right, val);
            
        return root;
    }
};
```