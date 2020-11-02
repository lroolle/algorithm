"+++
title = "0968. Binary Tree Cameras C++ 监控二叉树 易懂简洁的状态转移 2020.8.13 "
author = ["wsy-is-god"]
date = 2020-08-13T09:39:51+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 监控二叉树 易懂简洁的状态转移 2020.8.13

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [C++ 监控二叉树 易懂简洁的状态转移 2020.8.13](https://leetcode-cn.com/problems/binary-tree-cameras/solution/c-jian-kong-er-cha-shu-yi-dong-jian-ji-de-zhuang-t/) by [wsy-is-god](https://leetcode-cn.com/u/wsy-is-god/)

```
class Solution {
public:
    int ans = 0;
    int dfs(TreeNode* root) {
        if (root == NULL)
            return 2;
        int left_state = dfs(root->left);
        int right_state = dfs(root->right);
        if (left_state == 0 || right_state == 0) {
            ans++;
            return 1;
        }
        else if (left_state == 1 || right_state == 1)
            return 2;
        else return 0;
    }
    int minCameraCover(TreeNode* root) {
        return dfs(root) == 0 ? ans + 1 : ans;
    }
};
```
这道题是看了题解的状态转移法才理解明白的。
要明确好状态才行，每一个结点有3个状态：
我们定义0：该结点没有被摄像机覆盖到，意味着父节点没有摄像机，1：该结点处有一个摄像机，2：该结点没有摄像机，但子节点有摄像机
我们设想一下，如果要求最少的安置摄像头数量，我们肯定会择优选择不在头尾处安插，于是利用后续遍历，即最后LRT的顺序遍历。当我们从底层向上遍历的时候我们肯定一开始左右结点处于0状态，就是父节点没有摄像头，因此我们要在0状态的时候给父节点安插一个摄像头，父节点变为状态1，当我们在父节点安插好摄像头后，对于父节点的上层结点该结点就成了子节点，那么上层结点的状态即为2，因此我们要return 2.

最初使的时候遍历到底层遇到空结点，我们返回2，因为我们不希望在尾结点安置摄像头，在最下层结点安置摄像头不如它的上层结点安置摄像头覆盖的多。
