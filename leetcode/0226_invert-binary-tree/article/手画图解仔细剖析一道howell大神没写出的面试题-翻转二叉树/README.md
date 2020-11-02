"+++
title = "0226. Invert Binary Tree 「手画图解」仔细剖析一道Howell大神没写出的面试题 | 翻转二叉树 "
author = ["xiao_ben_zhu"]
date = 2020-09-16T00:49:38+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 「手画图解」仔细剖析一道Howell大神没写出的面试题 | 翻转二叉树

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [「手画图解」仔细剖析一道Howell大神没写出的面试题 | 翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/solution/shou-hua-tu-jie-san-chong-xie-fa-di-gui-liang-chon/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)

这道题的故事背景是：Howell 大神在推上吐槽，自己面试谷歌，因为没写出翻转二叉树被拒了。Leetcode 问讯立马收了这道题，并把难度标为 easy。

![image.png](https://pic.leetcode-cn.com/1600217921-cKePTW-image.png)

#### 思路
![image.png](https://pic.leetcode-cn.com/1600223748-vxiJaf-image.png)


每个子树的根节点都说：“先翻转我的左右孩子吧！”

结果就是，位于树的底部的、左右孩子都是 null 的子树，先被翻转。

向上返回的过程中，以这个子树为孩子的子树被翻转……

递归全部出栈时，整个树翻转好了。

一个子树的翻转怎么实现呢？——交换一个子树的左右子节点（左右子树）

于是，问题在递归出栈时，被解决。

#### 代码

```js
const invertTree = (root) => {
  if (root == null) {
    return root;
  }
  // 总是先 invert subtree
  invertTree(root.left);
  invertTree(root.right);
  // swap
  const temp = root.left;
  root.left = root.right;
  root.right = temp;
  
  return root;
};
```
#### 递归写法 2
![image.png](https://pic.leetcode-cn.com/1600220552-rZdRfp-image.png)


先交换整个左右子树，它们内部还没翻转，没事，交给递归去翻转。

也就是，把交换的操作，放在递归调用之前。

于是，问题在递归压栈时，被解决。
#### 递归代码2
```js
const invertTree = (root) => {
  if (root == null) {
    return root;
  }
  const temp = root.left;
  root.left = root.right;
  root.right = temp;

  invertTree(root.left);
  invertTree(root.right);

  return root;
};
```

#### 复盘总结

第一种其实就是后序遍历，第二种是前序遍历。

前序、后序遍历，都是基于DFS，都是先访问根节点、再访问左子树、再访问右子树（如果有人反驳你，那是他错了）。

**唯一的区别是：**

前序遍历：将「对节点的处理操作」放到了「递归左子树」之前！
后序遍历：将「对节点的处理操作」放到了「递归右子树」之后！

这个「对节点的处理操作」，就是「do something with root」，就是解决问题的代码，就是下面这段：
```js
const temp = root.left;
root.left = root.right;
root.right = temp;
```

递归只是帮你 walk through 这棵树，是一个辅助角色，核心还是「解决问题的代码」，递归只是把这段代码，应用到每一个子树上，解决每一个子问题，最后解决整个问题。

你可以选择将这部分代码放到，DFS遍历过程中的一个合适的时间点，殊途同归。

#### 迭代写法（BFS写法）
用层序遍历的方式去扫描二叉树。

根节点先入列，然后出列，出列就先交换它的左右子节点（左右子树）。

并且让它的左右子节点（左右子树）入列，往后，这些子节点（子树）会出列，也会被翻转。

直到队列为空，就遍历完所有的节点（子树），也就翻转了整个二叉树。
```js
const invertTree = (root) => {
  if (root == null) {
    return root;
  }
  const queue = [];   // 维护一个队列
  queue.push(root);   // 初始推入第一层的root
  
  while (queue.length) {
    const cur = queue.shift(); // 出列的节点

    const temp = cur.left;     // 交换出列节点左右子树
    cur.left = cur.right;
    cur.right = temp;

    if (cur.left) {            // 左右子节点继续入列考察，因为要继续翻转它们的孩子  
      queue.push(cur.left);
    }
    if (cur.right) {
      queue.push(cur.right);
    }
  }

  return root;
};
```

#### 复盘BFS写法

解决问题的代码放在哪里？

放在节点出列的时候，出列就考察它。
#### 后记
树应该是最重要的数据结构了，只要吃编程这碗饭，就要和它打交道。

数据库离不开它，JSON对象也是树，浏览器的DOM结构也是树，很多设计也是树形的。

只要往深一点思考，你会发现一道题目的涵盖面其实是不小的，可以感受出很多东西。

这样一总结，就很晴朗，可以关注我，我尽量在每道题都想深一丢丢，并很注重阅读体验。


#### 如果有帮助，点个赞鼓励我继续写下去画下去，坚持不易，如果哪里写得不对或不够好，指出我我会继续修改调整。