"+++
title = "0501. Find Mode in Binary Search Tree 二叉搜索树中的众数 "
author = ["junstat"]
date = 2019-05-16T13:02:58+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# 二叉搜索树中的众数

> [0501. Find Mode in Binary Search Tree](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)
> [二叉搜索树中的众数](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/er-cha-sou-suo-shu-zhong-de-zhong-shu-by-junstat/) by [junstat](https://leetcode-cn.com/u/junstat/)

C++实现，不使用额外空间。

思路：二叉搜索树的中序遍历是一个升序序列，逐个比对当前结点(root)值与前驱结点（pre)值。更新当前节点值出现次数(curTimes)及最大出现次数(maxTimes)，更新规则：若curTimes=maxTimes,将root->val添加到结果向量(res)中；若curTimes>maxTimes,清空res,将root->val添加到res,并更新maxTimes为curTimes。

```cpp
void inOrder(TreeNode* root, TreeNode*& pre, int& curTimes, 
             int& maxTimes, vector<int>& res){
    if (!root) return;
    inOrder(root->left, pre, curTimes, maxTimes, res);
    if (pre)
        curTimes = (root->val == pre->val) ? curTimes + 1 : 1;
    if (curTimes == maxTimes)
        res.push_back(root->val);
    else if (curTimes > maxTimes){
        res.clear();
        res.push_back(root->val);
        maxTimes = curTimes;
    }
    pre = root;
    inOrder(root->right, pre, curTimes, maxTimes, res);
}
vector<int> findMode(TreeNode* root) {
    vector<int> res;
    if (!root) return res;
    TreeNode* pre = NULL;
    int curTimes = 1, maxTimes = 0;
    inOrder(root, pre, curTimes, maxTimes, res);
    return res;
}
```