"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 二叉搜索树的最近公共祖先 | 两种解法：递归、迭代法 "
author = ["xiao_ben_zhu"]
date = 2020-06-25T20:09:34+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "迭代", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 二叉搜索树的最近公共祖先 | 两种解法：递归、迭代法

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [二叉搜索树的最近公共祖先 | 两种解法：递归、迭代法](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/di-gui-he-die-dai-fa-by-hyj8/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)


#### 思路
由二叉搜索树（BST）的性质，如果 p.val 和 q.val 都比 root.val 小，那么 p 和 q 节点肯定出现在 root 的左子树。

那问题就规模就变小了，不是递归调用整个树，而是递归左子树。

如果 p.val 和 q.val 都比 root.val 大，则递归右子树。

其他情况，root 即为所求，不管是一个大于 root.val 一个小于，还是一个等于一个小于，还是一个等于一个大于。

为什么？

因为 root 为 p 和 q 的最近公共祖先，包含下面的情况：
  -  p 和 q 分居 root 的左右子树
  - p 为 root ，q 在 root 的左或右子树中
  - q 为 root ， p 在 root 的左或右子树中

所以，只要不是 p 和 q 的值都大于、或小于 root.val，即只要不是同处在一个子树中，它们的公共祖先都是 root。

#### 递归 代码
```js
const lowestCommonAncestor = (root, p, q) => {
  if (p.val < root.val && q.val < root.val) {
    return lowestCommonAncestor(root.left, p, q);
  }
  if (p.val > root.val && q.val > root.val) {
    return lowestCommonAncestor(root.right, p, q);
  }
  return root;
};
```
#### 迭代
开启 while 循环，当 root 为 null 时结束循环（可以把 root 看作一个指针）。

如果 p 和 q 的节点值都小于 root.val，即它们都在 root 的左子树中，则`root = root.left`，遍历到 root 的左子节点。

如果 p 和 q 的节点值都大于 root.val，即它们都在 root 的右子树中，则`root = root.right`，遍历到 root 的右子节点。

否则，其他情况下，当前的 root 就是最近公共祖先，找到了，此时 break。

最后返回 root。
```js
const lowestCommonAncestor = (root, p, q) => {
  while (root) {
    if (p.val < root.val && q.val < root.val) {
      root = root.left;
    } else if (p.val > root.val && q.val > root.val) {
      root = root.right;
    } else {
      break;
    }
  }
  return root;
};
```
#### 如果有帮助，点个赞鼓励我继续写下去，写写画画了一百多篇题解和图解，如果哪里不对或哪里不好，指出我我继续修改。