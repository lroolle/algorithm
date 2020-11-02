"+++
title = "0250. Count Univalue Subtrees &&的顺序问题 "
author = ["hu-xi-yan"]
date = 2020-08-13T01:56:15+08:00
tags = ["Leetcode", "Algorithms", "C++", "DFS"]
categories = ["leetcode"]
draft = false
+++

# &&的顺序问题

> [0250. Count Univalue Subtrees](https://leetcode-cn.com/problems/count-univalue-subtrees/)
> [&&的顺序问题](https://leetcode-cn.com/problems/count-univalue-subtrees/solution/de-shun-xu-wen-ti-by-hu-xi-yan/) by [hu-xi-yan](https://leetcode-cn.com/u/hu-xi-yan/)

A && B && C && D：
注意顺序问题：
1.提前判断可行性，以达到剪枝的目的，需要提前结束递归，应把递归式子放在最后D。
2.需先递归检索所有可能的情况，再由下至上判断，则需把递归式子放在最前面A。

下面递归的两种写法，会导致非常不同的结果，需要特别注意：
```
if(root->left)
    valid = util(root->left, re) && valid && (root->left->val == root->val);
if(root->right)
    valid = util(root->right, re) && valid && (root->right->val == root->val);
```
或者：
```
if(root->left)
    valid = valid && (root->left->val == root->val) && util(root->left, re);
if(root->right)
    valid = valid && (root->right->val == root->val) && util(root->right, re);
```
结合此题，需要找出所有单值子树，所以必须检查所有情况，不能提前剪枝，因此题解如下：
```
class Solution {
public:
    int countUnivalSubtrees(TreeNode* root) {
        int re = 0;
        util(root, re);
        return re;
    }
    bool util(TreeNode* root, int& re){
        if(!root) return false;

        bool valid = true;
        if(root->left)
            valid = util(root->left, re) && valid && (root->left->val == root->val);
        if(root->right)
            valid = util(root->right, re) && valid && (root->right->val == root->val);
        re += valid;
        return valid;
    }
};
```
