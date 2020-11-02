"+++
title = "0617. Merge Two Binary Trees 递归逻辑运算（5行代码） "
author = ["mantoufan"]
date = 2020-09-23T09:12:08+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 递归逻辑运算（5行代码）

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [递归逻辑运算（5行代码）](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/di-gui-luo-ji-yun-suan-fu-5xing-dai-ma-by-mantoufa/) by [mantoufan](https://leetcode-cn.com/u/mantoufan/)

### 解题思路
- 三种情况
    1. 都存在：值相加
    2. 一个不存在：存在的替换不存在
    3. 都不存在：返回`null`
- `null` || `null` === `null`，所以：
    `2` 和 `3` 合并为 t1 || t2
### 代码

```javascript
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} t1
 * @param {TreeNode} t2
 * @return {TreeNode}
 */
var mergeTrees = function(t1, t2) {
    if (!t1 || !t2) {
        return t1 || t2
    }
    t1.val += t2.val
    t1.left = mergeTrees(t1.left, t2.left)
    t1.right = mergeTrees(t1.right, t2.right)
    return t1
};
```

### 结果
![QQ拼音截图20200922191744.png](https://pic.leetcode-cn.com/1600852255-MyeKZv-QQ%E6%8B%BC%E9%9F%B3%E6%88%AA%E5%9B%BE20200922191744.png)
