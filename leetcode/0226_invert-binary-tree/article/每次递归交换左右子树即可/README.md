"+++
title = "0226. Invert Binary Tree 每次递归交换左右子树即可。 "
author = ["tea-57"]
date = 2020-08-19T06:25:34+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 每次递归交换左右子树即可。

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [每次递归交换左右子树即可。](https://leetcode-cn.com/problems/invert-binary-tree/solution/mei-ci-di-gui-jiao-huan-zuo-you-zi-shu-ji-ke-by-te/) by [tea-57](https://leetcode-cn.com/u/tea-57/)

### 解题思路
每次递归交换左右子树即可。遇到空节点返回空节点。
![1597818257(1).png](https://pic.leetcode-cn.com/3bc14be8f16de568658ac0c3446a281a8b56355bc8b48f56b14ca08ece6d3da0-1597818257\(1\).png)

### 代码

```cpp
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
    TreeNode* invertTree(TreeNode* root) {
        if(root==NULL)
            return NULL;
        TreeNode* temp=root->left;
        root->left=root->right;
        root->right=temp;
        invertTree(root->left);//递归左子树
        invertTree(root->right);//递归右子树 
        return root;
    }
};
```