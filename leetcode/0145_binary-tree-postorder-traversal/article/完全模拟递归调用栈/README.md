"+++
title = "0145. Binary Tree Postorder Traversal 完全模拟递归调用栈 "
author = ["arsenal591"]
date = 2020-08-26T13:58:09+08:00
tags = ["Leetcode", "Algorithms", "C++", "Stack", "Recursion", "DepthfirstSearch", "二叉树"]
categories = ["leetcode"]
draft = false
+++

# 完全模拟递归调用栈

> [0145. Binary Tree Postorder Traversal](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)
> [完全模拟递归调用栈](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/wan-quan-mo-ni-di-gui-diao-yong-zhan-by-arsenal591/) by [arsenal591](https://leetcode-cn.com/u/arsenal591/)

#### 思路

首先写出递归求解的程序（伪代码）：
```C++
void postorderRecursive(TreeNode* pos) {
    if (pos == nullptr) {
        return;
    }
    postorderRecursive(pos->left);
    postorderRecursive(pos->right);
    visit(pos); // 泛指一切的遍历操作，在本题中，即为将 pos->val 插入到输出队列当中
}
```

对于每个递归调用子程序，有且只有一个参数：当前位置 `pos`。但每次调用栈来到 `postorderRecursive(pos)` 时（设 `pos != nullptr`），有四种可能的阶段：
- 即将访问 `pos->left`
- 即将访问 `pos->right`
- 即将调用 `visit(pos)`

观察到，我们用一对 `(TreeNode* pos, int n)` 就能够描述一个子程序目前的「状态」：其中 `TreeNode* pos` 即为子程序的输入参数， `int n` 为当前正在执行的阶段，取值为 0, 1, 2，分别对应上面描述的三个阶段。

我们使用一个 `stack<pair<TreeNode*, int>>` 来模拟递归调用栈，调用栈的栈顶代表了当前正在被模拟执行的子程序。每次迭代，需要先取出栈顶元素，再将栈顶元素弹出。根据栈顶元素 `(pos, n)` 的取值，有以下几种情况：
- 如果 `pos == nullptr`，直接跳过，对应递归调用中的 `return`。
- 否则，如果 `n == 0`，意味着正在执行子程序 `postorderRecursive(pos)` 的阶段 1。此时首先将 `(pos, 1)` 重新压回栈中，意味着该子程序（`postorderRecursive(pos)`）此后还会继续执行；随后将 `(pos->left, 0)` 压入栈中，意味着开启一个新的子程序，即`postorderRecursive(pos->left)`。
- 如果  `n == 1`，与 `n == 0` 同理
- 如果 `n == 2`，意味着此时需要访问节点 `pos`，故直接将 `pos->val` 插入到输出队列当中，然后继续迭代过程，意味着子程序 `postorderRecursive(pos)` 执行完毕。

翻译成代码，即为：


```C++
vector<int> postorderTraversal(TreeNode* root) {
    stack<pair<TreeNode*, int>> s;
    vector<int> ret;
    s.push(make_pair(root, 0));

    while (!s.empty()) {
        auto [pos, n] = s.top();
        s.pop();

        if (pos == nullptr) {
            continue;
        }
        if (n == 0) {
            s.push(make_pair(pos, 1));
            s.push(make_pair(pos->left, 0));
        } else if (n == 1) {
            s.push(make_pair(pos, 2));
            s.push(make_pair(pos->right, 0));
        } else {
            ret.push_back(pos->val);
        }
    }
    return ret;
}
```

#### 总结

本文提供了一套将递归调用「无脑」转化成迭代的一种通用方法。具体而言，这套方法需要以下步骤：
- 确定子程序的输入参数 `params`
- 确定子程序的所有「阶段」，其中以每个递归调用作为阶段的划分。如果有 $n$ 个递归调用（本题中 $n = 2$），则对应 $n + 1$ 个阶段
- 确定子程序在每个「阶段」要做的事情，即两个递归调用之间的代码部分。

具体而言，如果递归调用程序是
``` C++
ReturnType f(Input params) {
    // something in stage 1
    f (params_1);
    // something in stage 2
    f (params_2);

    // ...

    // something in stage n
    f (params_n)
    // something in stage n + 1
}
```

则对应的迭代写法就为
```
ReturnType g(Input initial_params) {
    stack<pair<Input, int>> s;
    s.push({initial_params, 0});

    while (!s.empty()) {
        auto [params, i] = s.top();
        s.pop();

        if (i == 0) {
            // something in stage 1
            s.push(params, 1);
            s.push(params_1, 0);
        } else if (n == 1) {
            // something in stage 2
            s.push(params, 2);
            s.push(params_2, 0);
        } 
        // ...
        else if (i == n) {
            // something in stage n
            s.push(params, n + 1);
            s.push(params_n, 0);
        } else { // i == n + 1
            // something in stage n + 1
        }

    }
}
```
