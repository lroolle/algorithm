"+++
title = "0112. Path Sum 112. 路径总和 | 递归 "
author = ["xiao_ben_zhu"]
date = 2020-06-23T07:42:06+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 112. 路径总和 | 递归

> [0112. Path Sum](https://leetcode-cn.com/problems/path-sum/)
> [112. 路径总和 | 递归](https://leetcode-cn.com/problems/path-sum/solution/di-gui-ti-by-hyj8/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)


#### 思路
- sum —— 从根节点到叶子节点的路径上的节点值相加的目标和
- 递归。转为判断：左、右子树中能否找出和为 `sum - root.val` 的路径 
- 每遍历一个节点，sum 就减去当前节点值，当遍历到叶子节点时，已经没有子节点了，如果 sum - 当前叶子节点值 == 0 ，就是找到了从根节点到叶子节点的和为 sum 的路径
- 时间复杂度：O(n)，每个节点被遍历一次
#### 代码
```js
const hasPathSum = (root, sum) => {
  if (root == null) return false;                // 遍历到null节点
  if (root.left == null && root.right == null) { // 遍历到叶子节点
    return sum - root.val == 0;                  // 如果满足这个就返回true
  }
  return hasPathSum(root.left, sum - root.val) ||
    hasPathSum(root.right, sum - root.val);      // 大问题转成两个子树的问题
}
```

#### 如果觉得还可以，点个赞鼓励我继续写下去，如果哪里写得不对或不通顺，指出我我看到就改。
