"+++
title = "0968. Binary Tree Cameras 后序遍历简单易懂 "
author = ["shi-di-zi-007"]
date = 2020-06-20T09:05:10+08:00
tags = ["Leetcode", "Algorithms", "C++", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 后序遍历简单易懂

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [后序遍历简单易懂](https://leetcode-cn.com/problems/binary-tree-cameras/solution/hou-xu-bian-li-jian-dan-yi-dong-by-shi-di-zi-007/) by [shi-di-zi-007](https://leetcode-cn.com/u/shi-di-zi-007/)


主要思想：根据每个节点的左右子节点来判断当前节点的状态，因此左右根后续遍历
```
class Solution {
public:
    //记录需要放置摄像头的数量
    int res = 0;
    int minCameraCover(TreeNode* root) 
    {
        //后序遍历，从下自上遍历。
        //若遍历至最上面，root标志为0，则多加一个摄像头
        if(dfs(root) == 0)
        {
            res++;  
        }
        return res;
    }
    int dfs(TreeNode* root)
    {

        //0：未被覆盖(当前节点未被照到)
        //1：已被覆盖(摄像头已经照到这个节点)
        //2：需放置摄像头

        //到根节点，
        if(root == NULL) return 1;
        //遍历左节点
        int left = dfs(root->left);
        //遍历右节点
        int right = dfs(root->right);
        //一个节点左右确定后，判断左右节点情况
        //所有情况00,01,02,11,12,22
        //左右孩子中有一个未被覆盖，则当前节点需要放置摄像头，当前节点标志为2
        if(left ==0 || right==0)
        {
            res++;
            return 2;
        }
        //左右孩子均为已覆盖状态,则当前节点未被覆盖，标志为0
        if(left == 1 && right == 1)
        {
            return 0;
        }
        //若左右孩子为一个覆盖一个放置了摄像头，则当前节点为已被覆盖，标志为1
        if(left+right >= 3)
        {
            return 1;
        }
        //此时已经组合完了根节点所有情况，随便返回一个整数即可
        return 0;
    }
};
```




















原JAVA大佬的：https://leetcode-cn.com/problems/binary-tree-cameras/solution/chao-ji-hao-li-jie-de-da-an-by-levyjeng/