"+++
title = "0404. Sum of Left Leaves 4ms 递归算法 "
author = ["scream419"]
date = 2020-07-13T12:36:00+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 4ms 递归算法

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [4ms 递归算法](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/4ms-di-gui-suan-fa-by-scream419-2/) by [scream419](https://leetcode-cn.com/u/scream419/)

#算法思路
对任意一个节点，它只需要做两件事；
1、判断它的左孩子是不是左叶子；
2、让它的左孩子和右孩子分别向它汇报，以该孩子为根的sumOfLeftLeaves；
最后简单相加即可，很典型的递归算法。
```
class Solution {
public:
    int sumOfLeftLeaves(TreeNode* root) {
        if(root==NULL)return 0;
        int ans=0;
        if(root->left!=NULL){
            if(root->left->left==NULL&&root->left->right==NULL)ans+=root->left->val;
            else ans+=sumOfLeftLeaves(root->left);
        }
        if(root->right!=NULL){
            ans+=sumOfLeftLeaves(root->right);
        }
        return ans;
    }
};
```
