"+++
title = "0113. Path Sum II 113. 路径总和 II（递归） "
author = ["Sunny_SMILE"]
date = 2020-07-08T03:10:48+08:00
tags = ["Leetcode", "Algorithms", "C++", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 113. 路径总和 II（递归）

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [113. 路径总和 II（递归）](https://leetcode-cn.com/problems/path-sum-ii/solution/113-lu-jing-zong-he-iidi-gui-by-sunny_smile/) by [Sunny_SMILE](https://leetcode-cn.com/u/sunny_smile/)

可以先做[112. 路径总和](https://leetcode-cn.com/problems/path-sum/)，附上我的[题解](https://leetcode-cn.com/problems/path-sum/solution/112-lu-jing-zong-he-di-gui-bfs-by-sunny_smile/)
还有[437. 路径总和 III](https://leetcode-cn.com/problems/path-sum-iii/)，附上我的[题解](https://leetcode-cn.com/problems/path-sum-iii/solution/437-lu-jing-zong-he-iiishuang-zhong-dfs-by-sunny_s/)

**代码**
```
class Solution {
public:
    vector<vector<int>> ans;

    void DFS(TreeNode* root, vector<int>& t, int sum)
    {
        if(root == nullptr)
            return;
        t.push_back(root->val);
        if(root->val == sum && root->left == nullptr && root->right == nullptr)
        // 是叶子节点且从根节点到叶子节点路径总和=sum -> 符合题目的路径
            ans.push_back(t);
        // if(root->left)
        DFS(root->left, t, sum - root->val);
        // if(root->right)
        DFS(root->right, t, sum - root->val);
        // 弹出最后一个元素
        t.pop_back();
    }

    vector<vector<int>> pathSum(TreeNode* root, int sum) {
        vector<int> t;
        DFS(root, t, sum);
        return ans;
    }
};
```