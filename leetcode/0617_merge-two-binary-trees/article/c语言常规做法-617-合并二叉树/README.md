"+++
title = "0617. Merge Two Binary Trees C语言常规做法 | 617. 合并二叉树 "
author = ["bevischou"]
date = 2020-09-23T02:16:24+08:00
tags = ["Leetcode", "Algorithms", "C", "Tree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# C语言常规做法 | 617. 合并二叉树

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [C语言常规做法 | 617. 合并二叉树](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/cyu-yan-chang-gui-zuo-fa-617-he-bing-er-cha-shu-by/) by [bevischou](https://leetcode-cn.com/u/bevischou/)

通过递归的方法，在t1上原地修改（当然在t2上也是一样的）。
```c
struct TreeNode* mergeTrees(struct TreeNode* t1, struct TreeNode* t2){
    if(t1 && t2){
        t1 -> val = t1 -> val + t2 -> val;
        t1 -> left = mergeTrees(t1 -> left, t2 -> left);
        t1 -> right = mergeTrees(t1 -> right, t2 -> right);
    }
    else if(t2){
        t1 = t2;
    }
    return t1;
}
```