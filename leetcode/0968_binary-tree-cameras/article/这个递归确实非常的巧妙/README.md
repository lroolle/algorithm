"+++
title = "0968. Binary Tree Cameras 这个递归确实非常的巧妙 "
author = ["mike-meng"]
date = 2019-12-25T00:33:00+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# 这个递归确实非常的巧妙

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [这个递归确实非常的巧妙](https://leetcode-cn.com/problems/binary-tree-cameras/solution/zhe-ge-di-gui-que-shi-fei-chang-de-qiao-miao-by-mi/) by [mike-meng](https://leetcode-cn.com/u/mike-meng/)

```c++
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
    int dfs(TreeNode * root,int & res){
        if(!root){
            return 2;
        }
        
        int l = dfs(root->left,res);
        int r = dfs(root->right,res);
        /*0： no camer watch
          1:  one camer on this node
          2:  no camer on this node, but child has camer
        */
        if( l == 0 || r == 0){
            res++;
            return 1;
        }else if(l == 1 || r == 1){
            return 2;
        }else{
            return 0;
        }
    }
    
    int minCameraCover(TreeNode* root) {
        int res = 0;
        if(dfs(root,res) == 0){
            res++;
        }
        
        return res;
    }
};
```