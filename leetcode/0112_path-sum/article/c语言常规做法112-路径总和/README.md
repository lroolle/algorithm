"+++
title = "0112. Path Sum C语言常规做法｜112. 路径总和 "
author = ["bevischou"]
date = 2020-09-22T07:10:44+08:00
tags = ["Leetcode", "Algorithms", "C", "Tree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# C语言常规做法｜112. 路径总和

> [0112. Path Sum](https://leetcode-cn.com/problems/path-sum/)
> [C语言常规做法｜112. 路径总和](https://leetcode-cn.com/problems/path-sum/solution/cyu-yan-chang-gui-zuo-fa-112-lu-jing-zong-he-by-be/) by [bevischou](https://leetcode-cn.com/u/bevischou/)

通过递归即可解决。另外也可通过深度优先搜索与回溯法解决，但递归更方便。
```c
bool hasPathSum(struct TreeNode* root, int sum){
    bool ret = false;
    if(root){
        int nextSum = sum - root -> val;
        if(root -> left || root -> right){
            if(root -> left){
                ret = hasPathSum(root -> left, nextSum);
            }
            if(root -> right){
                ret = ret || hasPathSum(root -> right, nextSum);
            }
        }
        else{
            ret = nextSum == 0;
        }
    }
    return ret;
}
```
以下是第一次提交的代码。个人认为代码会好看很多，但可惜无法通过特定测试点，即root为空，sum为0。
```c
bool hasPathSum(struct TreeNode* root, int sum){
    bool ret;
    if(root){
        int nextSum = sum - root -> val;
        ret = hasPathSum(root -> left, nextSum) || hasPathSum(root -> right, nextSum);
    }
    else{
        ret = sum == 0;
    }
    return ret;
}
```