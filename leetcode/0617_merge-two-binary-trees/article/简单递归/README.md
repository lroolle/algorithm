"+++
title = "0617. Merge Two Binary Trees 简单递归 "
author = ["li-zhi-chao-4"]
date = 2020-09-22T23:54:55+08:00
tags = ["Leetcode", "Algorithms", "cpp", "Tree", "Recursion", "DFS", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# 简单递归

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [简单递归](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/jian-dan-di-gui-by-li-zhi-chao-4/) by [li-zhi-chao-4](https://leetcode-cn.com/u/li-zhi-chao-4/)

### 解题思路
![QQ截图20200923075210.png](https://pic.leetcode-cn.com/1600818821-BJHWMJ-QQ%E6%88%AA%E5%9B%BE20200923075210.png)
简单递归即可，依次合并根据节点、合并左右子节点
微信搜索“编程猿来如此”关注公众号
![TIM截图20200602154558.png](https://pic.leetcode-cn.com/1600818885-VGBGqp-TIM%E6%88%AA%E5%9B%BE20200602154558.png)
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
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if( t1 == NULL) return t2;
        if( t2 == NULL) return t1;
        t1->val += t2->val;//t2 合并到 t1 上
        //分别合并左右子节点
        t1->left = mergeTrees(t1->left,t2->left);
        t1->right = mergeTrees(t1->right,t2->right);
        return t1;
    }
    
};
```