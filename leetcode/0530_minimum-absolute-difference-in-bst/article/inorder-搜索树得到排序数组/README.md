"+++
title = "0530. Minimum Absolute Difference in BST InOrder 搜索树得到排序数组 "
author = ["Eng_not_Ape"]
date = 2020-08-15T10:28:21+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# InOrder 搜索树得到排序数组

> [0530. Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
> [InOrder 搜索树得到排序数组](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/solution/inorder-sou-suo-shu-de-dao-pai-xu-shu-zu-by-eng_no/) by [Eng_not_Ape](https://leetcode-cn.com/u/eng_not_ape/)

### 解题思路
此处撰写解题思路

### 代码

```c

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
#define MAXSIZE 101
int len=0;

void get_array(struct TreeNode* root,int *array){/*in order*/
    if(root==NULL)  return;
    get_array(root->left,array);
    array[len++]=root->val;
    get_array(root->right,array);
    //return ans;
}
int get_min(int *array){
    int delta=100000,temp;
    for(int i=1;i<len;i++){
        temp=(array[i]>array[i-1]?(array[i]-array[i-1]):(array[i-1]-array[i]));
        if(delta>temp){
            delta=temp;
        }
    }
    return delta;
}
int getMinimumDifference(struct TreeNode* root){
    if(root==NULL)  return false;
    int *array[MAXSIZE]={0};
    get_array(root,array);
    return get_min(array);
}

```