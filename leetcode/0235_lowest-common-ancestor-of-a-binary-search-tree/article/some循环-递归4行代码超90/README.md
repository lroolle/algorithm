"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree Some循环 + 递归（4行代码，超90%） "
author = ["mantoufan"]
date = 2020-09-27T04:26:04+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# Some循环 + 递归（4行代码，超90%）

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [Some循环 + 递归（4行代码，超90%）](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/somexun-huan-di-gui-4xing-dai-ma-chao-90-by-mantou/) by [mantoufan](https://leetcode-cn.com/u/mantoufan/)

### 解题思路
- 数组 `a` 记录找到 `p` 和 `q` 经历的节点
    - 找到 `p`，结果集 `r` 放入 `a`
    - 找到 `q`，结果集 `r` 放入 `a`   
    - `p` 和 `q` 都找到，终止`递归`
- 先找到的，经历节点少。结果集 `r[0]` < `r[1]`
- `some`找 `r[0]` 和 `r[1]` 中第一个相等的节点

### 代码

```javascript
var lowestCommonAncestor = function(root, p, q) {
    var r = [], dfs = (n, a) => {
        a.unshift(n), n === p && r.push(a), n === q && r.push(a)
        r.length < 2 && (n.left && dfs(n.left, a.slice()), n.right && dfs(n.right, a.slice()))
    }
    return dfs(root, []) || r[0].some(n => r[1].some(m => m === n) && (r = n)) && r
};
```

### 结果
![QQ拼音截图20200927121314.png](https://pic.leetcode-cn.com/1601180285-JDMKts-QQ%E6%8B%BC%E9%9F%B3%E6%88%AA%E5%9B%BE20200927121314.png)