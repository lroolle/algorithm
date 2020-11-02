"+++
title = "0112. Path Sum 112. 路径总和（递归、BFS两种方法） "
author = ["Sunny_SMILE"]
date = 2020-07-08T02:46:24+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 112. 路径总和（递归、BFS两种方法）

> [0112. Path Sum](https://leetcode-cn.com/problems/path-sum/)
> [112. 路径总和（递归、BFS两种方法）](https://leetcode-cn.com/problems/path-sum/solution/112-lu-jing-zong-he-di-gui-bfs-by-sunny_smile/) by [Sunny_SMILE](https://leetcode-cn.com/u/sunny_smile/)

这个题做完后还有这两道类似的题目：
[113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/)，附上我的[题解](https://leetcode-cn.com/problems/path-sum-ii/solution/113-lu-jing-zong-he-iidi-gui-by-sunny_smile/)
[437. 路径总和 III](https://leetcode-cn.com/problems/path-sum-iii/)，附上我的[题解](https://leetcode-cn.com/problems/path-sum-iii/solution/437-lu-jing-zong-he-iiishuang-zhong-dfs-by-sunny_s/)

# 1、递归
**代码**
```
class Solution {
public:
    bool hasPathSum(TreeNode* root, int sum) {
        if(!root)
            return false;
        if(root->val == sum && !root->left && !root->right)
            return true;
        return hasPathSum(root->left, sum - root->val) || hasPathSum(root->right, sum - root->val);
    }
};
```

# 2、广度优先搜索BFS
**代码**
```
class Solution {
public:
    bool hasPathSum(TreeNode* root, int sum) {
        if(!root)
            return false;
        queue<TreeNode *> q_node;
        queue<int> q_val;
        q_node.push(root);
        q_val.push(root->val);
        while(!q_node.empty())
        {
            TreeNode* p = q_node.front();
            int t = q_val.front();
            q_node.pop();
            q_val.pop();
            if(p->left == nullptr && p->right == nullptr)
            {
                if(t == sum)
                    return true;
                continue;
            }
            if(p->left)
            {
                q_node.push(p->left);
                q_val.push(p->left->val + t);
            }
            if(p->right)
            {
                q_node.push(p->right);
                q_val.push(p->right->val + t);
            }
        }
        return false;
    }
};
```