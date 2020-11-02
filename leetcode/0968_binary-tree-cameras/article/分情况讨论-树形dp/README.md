"+++
title = "0968. Binary Tree Cameras 分情况讨论 树形dp "
author = ["chwma"]
date = 2020-09-22T07:03:21+08:00
tags = ["Leetcode", "Algorithms", "DynamicProgramming", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 分情况讨论 树形dp

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [分情况讨论 树形dp](https://leetcode-cn.com/problems/binary-tree-cameras/solution/fen-qing-kuang-tao-lun-shu-xing-dp-by-chwma/) by [chwma](https://leetcode-cn.com/u/chwma/)

### 解题思路

对于每个节点 $node$，如果被监控，有三种状态：
- 状态0：做为儿子，它的父节点装有摄像头，此时它被监控
- 状态1：做为自己，它本身装有摄像头，此时它被监控
- 状态2：做为父亲，它的儿子节点装有摄像头，此时它被监控

### 状态表示
- $dp[node][0]$：表示当前节点 $node$为 $root$ 的子树，且做为儿子且被父亲监控，所需要的最小摄像头数量
- $dp[node][1]$：表示当前节点 $node$为 $root$ 的子树，且做为自己且被自己监控，所需要的最小摄像头数量
- $dp[node][2]$：表示当前节点 $node$为 $root$ 的子树，且做为父亲且被儿子监控，所需要的最小摄像头数量
- 由于第二维只有3个状态，也可以写成3个一维 $dp0$，$dp1$，$dp2$ 来表示

### 状态转移
假设当前节点为$node$，左儿子为 $left$，右儿子为 $right$，且假设左右儿子均存在，不存在的情况做特殊处理即可
- 当前节点作为儿子被监控，此时两个儿子节点只需要保证，以各自为根的子树分别所需最小摄像头数量，然后求和即为node所有最小摄像头数量，即：
    - 左儿子所需最小摄像头数量有两种情况，即左儿子作为父亲或者自己：$l0 = min(dp1[left], dp2[left])$
    - 右儿子所需最小摄像头数量有两种情况，即右儿子作为父亲或者自己：$r0 = min(dp1[right], dp2[right])$
    - $dp0[node] = l0 + r0$
- 当前节点作为自己被监控，此时每个节点都可能有三种选择，每种选择取最小，然后两个节点相加后，再加上根节点，即：
    - $dp1[node] = 1 + min(dp0[lef], dp1[left], dp2[left]) + min(dp0[right], dp1[right], dp2[right])$
- 当前节点作为父亲被监控，此时必须存在一个儿子节点有摄像头，即：
    - 左边必选，右边可选可不选：$l2 = dp1[left] + min(dp1[right], dp2[right])$
    - 右边必选，左边可选可不选：$r2 = dp1[right] + min(dp1[left], dp2[left])$
    - $dp2[node] = min(l2, r2)$

### 最终答案
$min(dp1[root], dp2[root])$

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
    unordered_map<TreeNode*, int> dp0; // son
    unordered_map<TreeNode*, int> dp1; // self
    unordered_map<TreeNode*, int> dp2; // parent

    void dfs(TreeNode* root) {
      if (!root) {
        return;
      }

      if (!root->left && !root->right) {
        dp2[root] = 1;
        dp1[root] = 1;
        dp0[root] = 0;
      } else if (!root->left) {
        dfs(root->right);
        dp2[root] = dp1[root->right];
        dp1[root] = 1 + min(dp2[root->right], min(dp1[root->right], dp0[root->right]));
        dp0[root] = min(dp1[root->right], dp2[root->right]);
      } else if (!root->right) {
        dfs(root->left);
        dp2[root] = dp1[root->left];
        dp1[root] = 1 + min(dp2[root->left], min(dp1[root->left], dp0[root->left]));
        dp0[root] = min(dp1[root->left], dp2[root->left]);
      } else {
        dfs(root->left);
        dfs(root->right);
        int l0 = min(dp1[root->left], dp2[root->left]);
        int l1 = min(dp2[root->left], min(dp1[root->left], dp0[root->left]));
        int l2 = dp1[root->left] + min(dp1[root->right], dp2[root->right]);
        
        int r0 = min(dp1[root->right], dp2[root->right]);
        int r1 = min(dp2[root->right], min(dp1[root->right], dp0[root->right]));
        int r2 = dp1[root->right] + min(dp1[root->left], dp2[root->left]);

        dp0[root] = l0 + r0;
        dp1[root] = 1 + l1 + r1;
        dp2[root] = min(l2, r2);
      }
    }

    int minCameraCover(TreeNode* root) {
      dfs(root);

      return min(dp1[root], dp2[root]);
    }
};
```