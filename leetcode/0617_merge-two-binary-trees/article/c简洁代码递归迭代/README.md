"+++
title = "0617. Merge Two Binary Trees C++简洁代码（递归/迭代）： "
author = ["OrangeMan"]
date = 2020-06-24T02:12:54+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# C++简洁代码（递归/迭代）：

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [C++简洁代码（递归/迭代）：](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/cjian-jie-dai-ma-di-gui-die-dai-by-orangeman/) by [OrangeMan](https://leetcode-cn.com/u/orangeman/)

### 解题思路
1.方法1：递归
2.方法2：栈迭代
3.方法3：队列迭代，方法2和方法3基本相同，只是容器不同，还有方法3加入一些简便的写法
### 代码

```递归
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if(!t1) return t2;
        if(!t2) return t1;
        t1->val += t2->val;
        t1->left = mergeTrees(t1->left, t2->left);
        t1->right = mergeTrees(t1->right, t2->right);
        return t1;
    }
};
```
```栈迭代
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if(!t1) return t2;
        stack<pair<TreeNode*, TreeNode*>> stk;
        if(t1 && t2) stk.push(make_pair(t1, t2));
        while(stk.size()) {
            pair<TreeNode*,TreeNode*> cur = stk.top(); stk.pop();
            cur.first->val = cur.first->val + cur.second->val;
            if(cur.second->left) {
                if(!cur.first->left) cur.first->left = new TreeNode(0); //如果左子树不存在，新建值为0的节点
                stk.push(make_pair(cur.first->left, cur.second->left));
            }
            if(cur.second->right) {
                if(!cur.first->right ) cur.first->right = new TreeNode(0); //如果右子树不存在，新建值为0的节点
                stk.push(make_pair(cur.first->right, cur.second->right));
            }
        }
        return t1;
    }
};
```
```队列迭代
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if(!t1) return t2;
        if(!t2) return t1;
        queue<pair<TreeNode*, TreeNode*>> q;
        q.push(make_pair(t1, t2));
        while(q.size()) {
            pair<TreeNode*,TreeNode*> cur = q.front(); 
            q.pop();
            cur.first->val = cur.first->val + cur.second->val;
            if(cur.second->left) {
                if(!cur.first->left){
                    cur.first->left = new TreeNode(0);
                }
                q.push(make_pair(cur.first->left, cur.second->left));
            }
            if(cur.second->right) {
                if(!cur.first->right ){
                    cur.first->right = new TreeNode(0);
                }
                q.push(make_pair(cur.first->right, cur.second->right));
            }
        }
        return t1;
    }
};
```

