"+++
title = "0226. Invert Binary Tree 递归解法 "
author = ["LHN_Leetcode"]
date = 2020-09-16T01:55:25+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 递归解法

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [递归解法](https://leetcode-cn.com/problems/invert-binary-tree/solution/di-gui-jie-fa-by-lhn_leetcode/) by [LHN_Leetcode](https://leetcode-cn.com/u/lhn_leetcode/)

### 解题思路
不断交换，递归时根节点的左右子树

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
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var invertTree = function (root) {
    const helper = (root) => {
        if (root) {
            [root.left, root.right] = [root.right, root.left];// 交换左右子树
            helper(root.left);
            helper(root.right);
        }
    }
    helper(root);
    return root;
};
```