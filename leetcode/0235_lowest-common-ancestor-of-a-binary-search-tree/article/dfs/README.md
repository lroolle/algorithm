"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree DFS "
author = ["7QQQQQQQ"]
date = 2020-09-10T15:44:45+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# DFS

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [DFS](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/dfs-by-7qqqqqqq-2/) by [7QQQQQQQ](https://leetcode-cn.com/u/7qqqqqqq/)

### 解题思路
挺简单的。
直接DFS即可。每次DFS记录出来当前点的左右子树是不是有这两个点，有的话个数+1，因为dfs是栈的过程。
从顶层往下跑，回溯是从下往上，碰到有一个结点的cnt=2，即他的子树中存在这两个结点，那么当前结点就是答案。

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
    TreeNode *ans;
    int flag=0;
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        dfs(root,p,q);
        return ans;
    }   
    int dfs(TreeNode* root, TreeNode* p, TreeNode* q){
            int cnt=0;
            if(root==p || root==q) cnt=1;
            if(root->left!=NULL) cnt+=dfs(root->left,p,q);
            if(root->right!=NULL) cnt+=dfs(root->right,p,q);
            if(flag) return 0;
            if(cnt==2) ans=root,flag=1;
            return cnt;
    }
};
```