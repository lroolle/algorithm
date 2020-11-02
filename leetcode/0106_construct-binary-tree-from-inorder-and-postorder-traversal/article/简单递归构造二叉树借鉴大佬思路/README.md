"+++
title = "0106. Construct Binary Tree from Inorder and Postorder Traversal 简单递归构造二叉树，借鉴大佬思路 "
author = ["shaw-r"]
date = 2020-08-31T08:07:04+08:00
tags = ["Leetcode", "Algorithms", "C", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 简单递归构造二叉树，借鉴大佬思路

> [0106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
> [简单递归构造二叉树，借鉴大佬思路](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/jian-dan-di-gui-gou-zao-er-cha-shu-jie-jian-da-lao/) by [shaw-r](https://leetcode-cn.com/u/shaw-r/)

### 解题思路
通过后序遍历确定根结点；
根据中序遍历确定左右子树结点个数left、right；
在后序列表中，前left个值为左子树，left+1到left+right为右子树；
在中序列表中，根结点前为左子树，根结点后为右子树；
确定左右子树后，便可作为参数递归构建左右子树。

### 代码

```c

struct TreeNode* buildTree(int* inorder, int inorderSize, int* postorder, int postorderSize){
    if(postorderSize == 0 || inorderSize == 0)return NULL;      //叶子结点的左右子树为空

    struct TreeNode* root = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    root->val = postorder[postorderSize-1];                     //根结点值为后序遍历最后一位

    int left;
    for(left=0;left<inorderSize;left++){
        if(inorder[left] == root->val)break;                    //找到中序列表中的根结点，其索引为左子树结点个数
    }
    int right = inorderSize - left - 1;                         //计算右子树结点个数

    root->left = buildTree(inorder,left,postorder,left);        //递归构建左、右子树
    root->right = buildTree(inorder+left+1,right,postorder+left,right);

    return root;
}
```