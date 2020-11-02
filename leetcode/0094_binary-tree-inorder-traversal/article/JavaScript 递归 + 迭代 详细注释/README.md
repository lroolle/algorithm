"+++
title = "0094. Binary Tree Inorder Traversal JavaScript 递归 + 迭代 详细注释 "
author = ["jsliang"]
date = 2020-09-14T02:42:42+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Recursion", "迭代"]
categories = ["leetcode"]
draft = false
+++

# JavaScript 递归 + 迭代 详细注释

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [JavaScript 递归 + 迭代 详细注释](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/javascript-di-gui-die-dai-xiang-xi-zhu-shi-by-jsli/) by [jsliang](https://leetcode-cn.com/u/jsliang/)

递归：

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
 * @return {number[]}
 */
const inorderTraversal = (root) => {
  // 1. 设置结果集
  const result = [];

  // 2. 递归
  const recursion = (root) => {
    // 2.1 如果见底，则返回
    if (!root) {
      return;
    }
    // 2.2 中序遍历：左 -> 中 -> 右
    recursion(root.left);
    result.push(root.val);
    recursion(root.right);
  };
  recursion(root);

  // 3. 返回结果集
  return result;
};

const root = {
  val: 1,
  left: null,
  right: {
    val: 2,
    left: { val: 3, left: null, right: null },
    right: null,
  },
};

console.log(inorderTraversal(root));
```

迭代：

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
 * @return {number[]}
 */
const inorderTraversal = (root) => {
  // 1. 设置结果集
  const result = [];

  // 2. 设置栈队列
  const stack = [];

  // 3. 遍历
  while (root || stack.length) {
    // 3.1 栈添加左子树
    while (root) {
      stack.push(root);
      root = root.left;
    }

    // 3.2 推出
    root = stack.pop();

    // 3.3 获取值
    result.push(root.val);

    // 3.4 转换右子树
    root = root.right;
  }

  // 4. 返回结果集
  return result;
};

const root = {
  val: 1,
  left: null,
  right: {
    val: 2,
    left: { val: 3, left: null, right: null },
    right: null,
  },
};

console.log(inorderTraversal(root));
```