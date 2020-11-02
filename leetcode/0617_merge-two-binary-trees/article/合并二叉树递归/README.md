"+++
title = "0617. Merge Two Binary Trees 【合并二叉树】递归 "
author = ["ikaruga"]
date = 2020-08-10T13:50:12+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 【合并二叉树】递归

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [【合并二叉树】递归](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/merge-two-binary-trees-by-ikaruga/) by [ikaruga](https://leetcode-cn.com/u/ikaruga/)

### 思路
1. 同时递归两颗树，按照结构创建一颗全新的树
2. 注意对于需要 `new TreeNode()` 的新树，从参数传入需要指向指针的指针

### 答题
```C++
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == nullptr && t2 == nullptr) return nullptr;

        TreeNode* root = new TreeNode(0);
        root->val += (t1 != nullptr) ? t1->val : 0;
        root->val += (t2 != nullptr) ? t2->val : 0;

        root->left = mergeTrees((t1 != nullptr) ? t1->left : nullptr, (t2 != nullptr) ? t2->left : nullptr);
        root->right = mergeTrees((t1 != nullptr) ? t1->right : nullptr, (t2 != nullptr) ? t2->right : nullptr);

        return root;
    }
};
```
```C++
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        TreeNode* head = nullptr;
        dfs(&head, t1, t2);
        return head;
    }

    void dfs(TreeNode** t, TreeNode* t1, TreeNode* t2) {
        if (t1 == nullptr && t2 == nullptr) return;
        
        *t = new TreeNode(0);
        (*t)->val += (t1 != nullptr) ? t1->val : 0;
        (*t)->val += (t2 != nullptr) ? t2->val : 0;

        dfs(&((*t)->left), (t1 != nullptr) ? t1->left : nullptr, (t2 != nullptr) ? t2->left : nullptr);
        dfs(&((*t)->right), (t1 != nullptr) ? t1->right : nullptr, (t2 != nullptr) ? t2->right : nullptr);
    }
};
```


### 致谢

感谢您的观看，希望对您有帮助，欢迎热烈的交流！  

**如果感觉还不错就点个赞吧~**

这是 [我的leetcode](https://github.com/AhJo53589/leetcode-cn) ，帮助我收集整理题目，可以方便的 `visual studio` 调试，欢迎关注，star

