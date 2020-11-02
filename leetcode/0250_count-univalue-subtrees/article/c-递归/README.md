"+++
title = "0250. Count Univalue Subtrees C++ 递归 "
author = ["xing-chen-da-hai-36"]
date = 2020-08-12T07:49:33+08:00
tags = ["Leetcode", "Algorithms", "C++", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# C++ 递归

> [0250. Count Univalue Subtrees](https://leetcode-cn.com/problems/count-univalue-subtrees/)
> [C++ 递归](https://leetcode-cn.com/problems/count-univalue-subtrees/solution/c-di-gui-by-xing-chen-da-hai-36-3/) by [xing-chen-da-hai-36](https://leetcode-cn.com/u/xing-chen-da-hai-36/)

### 解题思路
这里面是依次递归了每个树  有的树被重复计算了，不知道怎么改成每个节点只遍历一次，有好方法的同学请告知一下

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
    int countUnivalSubtrees(TreeNode* root) {
        if(!root)
            return 0;
        return help(root) + countUnivalSubtrees(root->left) + countUnivalSubtrees(root->right);
    }

    bool help(TreeNode* root){
        if(!root)
            return true;
        if(!root->left && !root->right)
            return true;
        if(root->left){
            if(root->val != root->left->val)
                return false;
            if(!help(root->left))
                return false;
        }
        if(root->right){
            if(root->val != root->right->val)
                return false;
            if(!help(root->right))
                return false;
        }
        return true;
    }
};
```