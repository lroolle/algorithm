"+++
title = "0617. Merge Two Binary Trees 递归方法，新手来 "
author = ["jia-you-jia-you-5"]
date = 2020-07-15T02:22:49+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 递归方法，新手来

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [递归方法，新手来](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/di-gui-fang-fa-xin-shou-lai-by-jia-you-jia-you-5/) by [jia-you-jia-you-5](https://leetcode-cn.com/u/jia-you-jia-you-5/)

### 解题思路
此处撰写解题思路

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
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        //if(t1==nullptr){
        if(!t1){ //t1是空，则t2返回
            return t2;
        }
        if(!t2){//t2是空，则t1返回
            return t1;
        }
        t1->val += t2->val; //精髓，两个相加，都不为空的话
        t1->left = mergeTrees(t1->left, t2->left); //精髓， 左子树 加 左子树 
        t1->right = mergeTrees(t1->right, t2->right); //右子树递归加右子树
        return t1;
    }
};
```