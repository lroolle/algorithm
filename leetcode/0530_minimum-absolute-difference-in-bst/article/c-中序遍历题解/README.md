"+++
title = "0530. Minimum Absolute Difference in BST C++ 中序遍历题解 "
author = ["da-li-wang"]
date = 2020-01-03T11:11:39+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 中序遍历题解

> [0530. Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
> [C++ 中序遍历题解](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/solution/c-zhong-xu-bian-li-ti-jie-by-da-li-wang-2/) by [da-li-wang](https://leetcode-cn.com/u/da-li-wang/)

### 解题思路
prev存储前继数
然后与当前数做差

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
    void dfs(TreeNode* root, int& prev, int& res) {
        if (root == NULL) return;
        dfs(root->left, prev, res);
        if (prev >= 0) res = min(res, root->val - prev);
        prev = root->val;
        dfs(root->right, prev, res);
    }
    int getMinimumDifference(TreeNode* root) {
        int prev = -1;
        int res = INT_MAX;
        dfs(root, prev, res);
        return res;
    }
};
```

![image.png](https://pic.leetcode-cn.com/1d78f36e0eb704c2d32efbed956507ba0c4d9c9cd31b35795d2c7800cf1c46f3-image.png)
