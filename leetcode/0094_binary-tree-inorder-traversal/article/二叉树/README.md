"+++
title = "0094. Binary Tree Inorder Traversal 二叉树 "
author = ["jiang-jiang-jiang-jiang-a"]
date = 2020-09-02T01:23:08+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# 二叉树

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [二叉树](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/er-cha-shu-by-jiang-jiang-jiang-jiang-a/) by [jiang-jiang-jiang-jiang-a](https://leetcode-cn.com/u/jiang-jiang-jiang-jiang-a/)

层序遍历，然后在奇数时逆序一下
class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int> > res;
        if(!root) return res;
        
        queue<TreeNode *> q;
        q.push(root);
        int flag = 0;
        while(!q.empty())
        {
            vector<int> out;
            int size = q.size(); //取得每一层的长度
            for(int i = 0; i < size; i++)
            {
                auto temp = q.front();
                q.pop();
                out.push_back(temp->val);
                if(temp->left)
                {
                    q.push(temp->left);
                }
                if(temp->right)
                {
                    q.push(temp->right);
                }
            }
            if(flag%2==1)
            {
                reverse(out.begin(),out.end());
            }
            res.push_back(out);
            flag++;
        }
        return res;
    }
};
```
