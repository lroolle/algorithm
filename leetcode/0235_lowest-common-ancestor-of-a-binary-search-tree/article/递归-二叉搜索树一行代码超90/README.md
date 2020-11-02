"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 递归 + 二叉搜索树（一行代码，超90%） "
author = ["mantoufan"]
date = 2020-09-27T05:10:14+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "二叉树", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 递归 + 二叉搜索树（一行代码，超90%）

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [递归 + 二叉搜索树（一行代码，超90%）](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/di-gui-er-cha-sou-suo-shu-yi-xing-dai-ma-chao-90-b/) by [mantoufan](https://leetcode-cn.com/u/mantoufan/)

### 解题思路
1. `二叉搜索树`：左子节点 < 根节点 < 右子节点，`root` `一定`是 `p`  `q` 根节点的可能：
    - `p`左 `q`右：`p` < `root` < `q`
    - `p`右 `q`左：`q` < `root` < `p`
    - `q`是`p`根节点：`p` < `q` = `root`
    - `p`是`q`根节点：`q` < `p` = `root`
    - `p`和`q`互为根节点（同节点）：`p` =  `q` = `root`
2. 上面5条件合二：
    - `q` <= `root` <= `p`
    - `p` <= `root` <= `q`
3. `递归`：遍历每个节点，只要满足上面两条件之一，则该节点一定是 `p` 和 `q` 的 `最近根节点`

### 代码

```javascript
var lowestCommonAncestor = function(root, p, q) {
    return (p.val <= root.val && root.val <= q.val || q.val <= root.val && root.val <= p.val) && root || lowestCommonAncestor(p.val < root.val && root.left || root.right, p, q)
};
```

### 结果
![2.png](https://pic.leetcode-cn.com/1601182093-SXrzpO-2.png)
