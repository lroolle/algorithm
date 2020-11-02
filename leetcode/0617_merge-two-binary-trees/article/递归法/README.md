"+++
title = "0617. Merge Two Binary Trees 递归法 "
author = ["jason-2"]
date = 2020-08-02T03:19:54+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 递归法

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [递归法](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/di-gui-fa-by-jason-2-9/) by [jason-2](https://leetcode-cn.com/u/jason-2/)

思路： 递归合并左子树，再递归合并右子树，再合并根节点。拼成最终结果。

```
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if(!t1 && !t2) return NULL;
        if(!t1 && t2) {
            return t2;
        }
        if(t1 && !t2){
            return t1;
        }
        t1->val += t2->val;
        t1->left = mergeTrees(t1->left,t2->left);
        t1->right = mergeTrees(t1->right,t2->right);
        return t1;
    }
};
```
