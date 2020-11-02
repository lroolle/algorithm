"+++
title = "0501. Find Mode in Binary Search Tree 迭代写法。 中序遍历。不使用额外空间。符合要求。 "
author = ["maserhe"]
date = 2020-09-21T10:13:49+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 迭代写法。 中序遍历。不使用额外空间。符合要求。

> [0501. Find Mode in Binary Search Tree](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)
> [迭代写法。 中序遍历。不使用额外空间。符合要求。](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/die-dai-xie-fa-zhong-xu-bian-li-bu-shi-yong-e-wai-/) by [maserhe](https://leetcode-cn.com/u/maserhe/)

为什么程序里面有个栈，还说没有使用额外空间？**用递归写的话，使用的的系统栈。我这里手动写的栈，没有使用系统栈。**
![QQ截图20200921180602.png](https://pic.leetcode-cn.com/1600682770-NGduSg-QQ%E6%88%AA%E5%9B%BE20200921180602.png)

![QQ截图20200921181250.png](https://pic.leetcode-cn.com/1600683179-vRBohb-QQ%E6%88%AA%E5%9B%BE20200921181250.png)


为什么不使用递归？
- **能用一个函数解决问题的话，就别写两个函数。**
- **手写栈能更好的理解递归**
**中序遍历** 对于 BST 刚好是 从大到小排列。然后边执行边 。纪录结果就行了。
```
vector<int> findMode(TreeNode* root) {
        
        if (!root) return {};
        vector<int> ans;
        // BST 的中序遍历 是 顺序排列的。

        stack<TreeNode *> st; 
        TreeNode * tmp = root;

        // ma 出现最多的 次数， last 上一次出现的值。 now 当前数已经出现过的次数。
        int ma = 0, last = -1, now = 0;

        // 迭代写法。 没有使用递归。手写栈模拟系统栈 ，实现递归程序。
        while (!st.empty() || tmp)
        {
            while (tmp)
            {
                st.push(tmp);
                tmp = tmp->left;
            }

            TreeNode * t = st.top();
            st.pop();

            if (last == t->val) now ++ ;
            else  now = 1;
            
            if (now > ma){
                ans.clear();
                ans.push_back(t->val);
                ma = now;
            }
            else if (now == ma){
                ans.push_back(t->val);
            }

            last = t->val;
            tmp = t->right;
        }

        return ans;
    }
```