"+++
title = "0968. Binary Tree Cameras 【监控二叉树】后序遍历 "
author = ["ikaruga"]
date = 2020-09-22T06:18:14+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 【监控二叉树】后序遍历

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [【监控二叉树】后序遍历](https://leetcode-cn.com/problems/binary-tree-cameras/solution/binary-tree-cameras-by-ikaruga/) by [ikaruga](https://leetcode-cn.com/u/ikaruga/)

### 思路
1. 想要最大化摄像头范围
    1. 如果从上往下捋，上面最大化了，下面因为分叉，为了补齐，需要多个摄像头
    2. 如果从下往上捋，下面最大化，上面由于分支合并，就算要补齐，补的个数也少
    3. 所以应该后序遍历

2. 将后序遍历的返回值定义为与摄像头的距离
    1. 如果这里安装了摄像头，就返回 0 
    2. 其父节点的距离就是 0 + 1 = 1
    3. 祖父节点的距离就是 0 + 1 + 1 = 2
    4. 当距离为 2 时，这里就覆盖不到了，再上一层就该安装新摄像头了
    5. 安装时，记录安装个数

3. 由于可能有左右两个节点
    1. 对于自身与摄像头的距离为：`ret = min(l, r) + 1`
    2. 而对于是否要安装摄像头，要考虑其子节点是否已经无法接受照射了
    3. 所以要判断：`if (max(l, r) == 2)`

4. 注意根节点的距离是 2 时，没有上一层给安摄像头，记得在外边判断一下

### 答题
```C++
class Solution {
public:
    int minCameraCover(TreeNode* root) {
        ans += postOrder(root) == 2;
        return ans;
    }

    // 返回值表示与摄像头的距离
    int postOrder(TreeNode* root) {
        if (root == nullptr) return 1;

        int l = postOrder(root->left);
        int r = postOrder(root->right);

        int ret = min(l, r) + 1;
        if (max(l, r) == 2) {
            ret = 0;
            ans++;
        }
        return ret;
    }

private:
    int ans = 0;
};
```



### 致谢
感谢您的观看，希望对您有帮助，欢迎热烈的交流！  

**如果感觉还不错就点个赞吧~**

这是 [我的leetcode](https://github.com/AhJo53589/leetcode-cn) ，帮助我收集整理题目，可以方便的 `visual studio` 调试，欢迎关注，star

