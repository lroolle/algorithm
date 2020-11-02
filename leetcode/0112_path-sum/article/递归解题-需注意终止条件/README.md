"+++
title = "0112. Path Sum 递归解题, 需注意终止条件 "
author = ["MuYunyun"]
date = 2020-09-05T06:50:26+08:00
tags = ["Leetcode", "Algorithms", "Recursion", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 递归解题, 需注意终止条件

> [0112. Path Sum](https://leetcode-cn.com/problems/path-sum/)
> [递归解题, 需注意终止条件](https://leetcode-cn.com/problems/path-sum/solution/di-gui-jie-ti-xu-zhu-yi-zhong-zhi-tiao-jian-by-muy/) by [MuYunyun](https://leetcode-cn.com/u/muyunyun/)

### 112. Path Sum

使用递归的思路进行解题:

```js
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} sum
 * @return {boolean}
 */
var hasPathSum = function(root, sum) {
  if (!root) return false
  return ifHasPathSum(root, sum)
}

var ifHasPathSum = function(node, sum) {
  if (!node && sum !== 0) return false
  if (!node && sum === 0) return true
  const remainingVal = sum - node.val
  const leftResult = ifHasPathSum(node.left, remainingVal)
  if (leftResult) return true
  const rightResult = ifHasPathSum(node.right, remainingVal)
  if (rightResult) return true
  return false
}
```

此时 ac, 卡在了以下测试用例中(此题的难点在于对递归终止条件的判断)

```js
    1
  /   \
null   2
```

先对代码进行整理精简:

```js
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} sum
 * @return {boolean}
 */
var hasPathSum = function(root, sum) {
  if (!root) return false
  return ifHasPathSum(root, sum)
}

var ifHasPathSum = function(node, sum) {
  if (!node) return sum === 0
  const remainingVal = sum - node.val
  return ifHasPathSum(node.left, remainingVal) || ifHasPathSum(node.right, remainingVal)
}
```

终止条件应由`判断当前节点为空而且 sum === 0`调整为`判断当前节点为叶子节点且 node.val === sum`, 再次修正代码

```js
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} sum
 * @return {boolean}
 */
var hasPathSum = function(root, sum) {
  if (!root) return false
  if (!root.left && !root.right) return root.val === sum
  const remainingVal = sum - root.val
  return hasPathSum(root.left, remainingVal) || hasPathSum(root.right, remainingVal)
}
```

为确保内容的实时、准确性, 可以查看[JavaScript 题解](https://github.com/MuYunyun/blog/blob/master/LeetCode/README.md)