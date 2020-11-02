"+++
title = "0617. Merge Two Binary Trees 「手画图解」合并二叉树 "
author = ["xiao_ben_zhu"]
date = 2020-09-23T02:36:41+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 「手画图解」合并二叉树

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [「手画图解」合并二叉树](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/shou-hua-tu-jie-by-xiao_ben_zhu-3/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)

#### 思路


同步地遍历两棵树上的节点，直接在 t1 上修改。

> **如果把`mergeTrees`函数作为递归函数，它接收的`t1`和`t2`是指：当前遍历的节点（子树）**

递归总是关注当前节点：
1. t1、t2 都存在，将 t2 的节点值加到 t1 上来。
2. t1 为 null、t2 不是 null，t1 换成 t2。
3. t2 为 null、t1 不是 null，t1 依然 t1。
3. t1 和 t2 都为 null，t1 依然 t1。

子树的合并，更新 t1 的子树，交给递归去做，它会对每一个节点做同样的事情。

![image.png](https://pic.leetcode-cn.com/1600828469-FbXAmS-image.png)


#### 
#### 代码
时间复杂度我觉得是：$O(N)$，N 是 t1 和 t2 重合后的节点的个数（并集）。
```js
const mergeTrees = (t1, t2) => {
  if (t1 == null && t2) {
    return t2;
  }
  if (t1 && t2 == null || t1 == null && t2 == null) {
    return t1;
  }
  t1.val += t2.val;

  t1.left = mergeTrees(t1.left, t2.left);
  t1.right = mergeTrees(t1.right, t2.right);

  return t1;
};
```
#### 感谢阅读