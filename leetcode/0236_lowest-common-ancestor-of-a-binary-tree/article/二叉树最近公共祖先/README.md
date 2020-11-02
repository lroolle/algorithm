"+++
title = "0236. Lowest Common Ancestor of a Binary Tree 二叉树最近公共祖先 "
author = ["beney-2"]
date = 2020-08-26T13:26:00+08:00
tags = ["Leetcode", "Algorithms", "Java", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# 二叉树最近公共祖先

> [0236. Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
> [二叉树最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/er-cha-shu-zui-jin-gong-gong-zu-xian-by-beney-2/) by [beney-2](https://leetcode-cn.com/u/beney-2/)

### 解题思路
这个是我最初始的比较暴力的解法：

要找到p、q的公共祖先，我们必须要知道**从根root搜索分别搜索到p和q的路径**，这是最基本的判断信息。获取搜索路径的方式使用`DFS`对树进行遍历,同时利用栈`Stack`来保存路径。

1. 首先DFS获取`p`的搜索路径`pPath`。
2. 若`q`在搜索路径`pPath`上，则`q`就是最近公共祖先。
3. 否则，再DFS获取`q`的搜索路径`qPath`。接着从其中一条路径自下而上遍历元素`e`，若`e`在**另一个结点的搜索路径**上则`e`就是解。

### 复杂度分析
至少需要进行一次DFS，时间复杂度$O(n)$。

最坏情况下，搜索到p和搜索到q的路径都要获取，然后匹配路径上最低的相同结点，由于我使用的是**栈**来存储路径，所以匹配操作的复杂度为$O(n^2)O$。

总体的时间复杂度最坏为$O(n^2)$

### 代码

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        Stack<TreeNode> pPath = new Stack<>();
        Stack<TreeNode> qPath = new Stack<>();
        findNodeDFS(root, p, pPath);
        if (pPath.contains(q)) {
            return q;
        } else {
            findNodeDFS(root, q, qPath);
            TreeNode res;
            while (!pPath.contains(res = qPath.pop())) {}
            return res;
        }
    }

    private boolean findNodeDFS(TreeNode root, TreeNode tar, Stack<TreeNode> path) {
        if (root == null)
            return false;
        else if (root == tar) { 
            path.push(root);
            return true; 
        }
        path.push(root);
        if (findNodeDFS(root.left, tar, path)) {
            return true;
        } else if (findNodeDFS(root.right, tar, path)) {
            return true;
        } else {
            path.pop();
            return false;
        }
    }
}
```

#### 对你有帮助就点个棒棒吧!欢迎评论:)