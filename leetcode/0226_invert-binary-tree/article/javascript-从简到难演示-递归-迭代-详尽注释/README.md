"+++
title = "0226. Invert Binary Tree JavaScript 从简到难演示 递归 + 迭代 详尽注释 "
author = ["jsliang"]
date = 2020-09-16T02:32:05+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Recursion", "迭代"]
categories = ["leetcode"]
draft = false
+++

# JavaScript 从简到难演示 递归 + 迭代 详尽注释

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [JavaScript 从简到难演示 递归 + 迭代 详尽注释](https://leetcode-cn.com/problems/invert-binary-tree/solution/javascript-cong-jian-dao-nan-yan-shi-di-gui-die-da/) by [jsliang](https://leetcode-cn.com/u/jsliang/)


这道题其实可以由简到难，逐步推演，且看下文！

首先，假设这棵树长这样：

```
  2
 / \
1   3
```

那么怎么交换它左右节点，这个很简单吧？！

```js
const invertTree = (root) => {
  // 1. 获取左节点，存入 tempRoot 中
  const tempRoot = root.left;

  // 2. 获取右节点，存入左节点
  root.left = root.right;

  // 3. 将 tempRoot 存入进右节点，交换完成
  root.right = tempRoot;

  // 4. 返回交换后节点
  return root;
};

const root = {
  val: 2,
  left: { val: 1, left: null, right: null },
  right: { val: 3, left: null, right: null },
};

console.log(invertTree(root));
```

既然交换一个简单的树可以，那么我们延伸一下，交换的树稍微复杂一点点，这也不是难事吧！

* 递归：5 行代码解决问题

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
 * @return {TreeNode}
 */
const invertTree = (root) => {
  // 1. 如果树已经空了，那么返回 null
  if (!root) {
    return null;
  }
  
  // 2. 使用 tempRoot 获取左子树
  const tempRoot = invertTree(root.left);

  // 3. 使用 左子树 获取 右子树
  root.left = invertTree(root.right);

  // 4. 使用 右子树 获取 tempRoot（左子树），实现简单互换
  root.right = tempRoot;

  // 5. 关键：返回交换后的树
  return root;
};

const root = {
  val: 4,
  left: {
    val: 2,
    left: { val: 1, left: null, right: null },
    right: { val: 3, left: null, right: null },
  },
  right: {
    val: 7,
    left: { val: 6, left: null, right: null },
    right: { val: 9, left: null, right: null },
  },
};

console.log(invertTree(root));
```

既然有递归，那么迭代也不能放过啊，套用广度优先搜索的那个套路，直接上手。

* 迭代：

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
 * @return {TreeNode}
 */
const invertTree = (root) => {
  // 狗不理
  if (!root) {
    return null;
  }

  // 1. 设置当前层的节点
  let nowRoot = [root];

  // 2. 当我们迭代后，还可以遍历当前层时，那就继续迭代
  while (nowRoot.length) {

    // 2.1 设置下一层节点
    const nextRoot = [];

    // 2.2 遍历当前层节点
    for (let i = 0; i < nowRoot.length; i++) {

      // 2.3 如果当前层存在左子树或者右子树
      if (nowRoot[i].left || nowRoot[i].right) {

        // 2.3 SWOP 交换
        [nowRoot[i].left, nowRoot[i].right] = [nowRoot[i].right, nowRoot[i].left];

        // 2.4 如果存在左子树，那么下一层节点添加左子树，方便将它里面内容进行交换
        if (nowRoot[i].left) {
          nextRoot.push(nowRoot[i].left);
        }

        // 2.5 如果存在右子树，那么同上
        if (nowRoot[i].right) {
          nextRoot.push(nowRoot[i].right);
        }
      }
    }

    // 2.6 将当前层节点替换为下一层节点
    nowRoot = nextRoot;
  }

  // 3. 搞定收工
  return root;
};

const root = {
  val: 4,
  left: {
    val: 2,
    left: { val: 1, left: null, right: null },
    right: { val: 3, left: null, right: null },
  },
  right: {
    val: 7,
    left: { val: 6, left: null, right: null },
    right: { val: 9, left: null, right: null },
  },
};

console.log(invertTree(root));
```

搞定收工，你懂了没？
