"+++
title = "0094. Binary Tree Inorder Traversal 「手画图解」用栈模拟中序遍历，怎么做以及为什么 | 递归与迭代 "
author = ["xiao_ben_zhu"]
date = 2020-09-14T00:57:16+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 「手画图解」用栈模拟中序遍历，怎么做以及为什么 | 递归与迭代

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [「手画图解」用栈模拟中序遍历，怎么做以及为什么 | 递归与迭代](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/shou-hua-tu-jie-yong-zhan-mo-ni-zhong-xu-bian-li-z/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)


#### 中序遍历的递归实现
这个大家都比较容易写出来，但我还是想回到它的最原始定义。

```js
const inorderTraversal = (root) => {
  const res = [];
  const inorder = (root) => {
    if (root == null) {
      return;
    }
    inorder(root.left);
    res.push(root.val);
    inorder(root.right);
  };
  inorder(root);
  return res;
};
```
正如我之前提过，大家别含糊地记说：“中序遍历是先访问左子树，再访问根节点，再访问右子树。”

这么描述不准确的，容易产生误导。

因为无论是前、中、后序遍历，都是先访问根节点，再访问它的左子树，再访问它的右子树。

那区别在哪里呢？比如中序遍历，它是将 do something with root（处理当前节点的val）放在了访问完它的左子树之后。如果是打印节点值，就会产生「左根右」的打印顺序。

![image.png](https://pic.leetcode-cn.com/1600022563-FCcWIJ-image.png)

前、中、后序遍历都是基于DFS，节点的访问顺序如上图所示，可以看到，每个节点有三个不同的驻留阶段，即每个节点会被经过 三次：
1. 在递归它的左子树之前。
2. 在递归完它的左子树之后，在递归它的右子树之前。
3. 在递归完它的右子树之后。

我们将 do something with root 放在这三个时间点之一，就分别对应了：前、中、后序遍历。

所以，它们的唯一区别是：在什么时间点去处理一个节点的数据部分。

下面伪代码就是『中序遍历』，拿 root.val 做一些事情，放在刚结束了它的左子树的递归的时候。
```
Inorder (root) {
  call Inorder(root.left) 
  access the content of root
  call Inorder(root.right)
}
```
#### 中序遍历的迭代实现
搞清楚概念后，我们想一下，怎么用栈去模拟递归遍历，并且是中序遍历呢？


对一棵树进行 dfs 先做什么？如下图，先递归节点A，再递归左子节点B，再递归左子节点D，递归调用是压入递归栈的操作。

![image.png](https://pic.leetcode-cn.com/1600027366-jHBJMK-image.png)


先是不断地将左节点压入栈，我们可以确定出这部分的代码：
```js
while (root) {
    stack.push(root);
    root = root.left;
}
```
DFS的访问顺序：根节点、左子树、右子树，现在轮到访问栈中的节点的右子树了。

并且是先访问位于树的底部、栈的顶部的节点的右子树。

于是，让栈中的节点逐个出栈，出栈的同时，把它的右子节点压入栈，即相当于对右子节点进行递归。

因为中序遍历，在这之前，我们要访问处理当前出栈节点的节点值：将出栈节点的节点值，输出。

![image.png](https://pic.leetcode-cn.com/1600046829-vQeqLn-image.png)

我们知道，新入栈的右子节点，或者说右子树，就是对他进行递归（递归压栈）。和节点A、B、D的递归压栈一样，它们都是子树，只是不同的子树。

所以，也要做一样的事情，同样也要先不断将它的左子节点压入栈，然后再出栈，带出右子节点入栈。

![image.png](https://pic.leetcode-cn.com/1600046603-eBtEoP-image.png)

即栈顶出栈的过程，也要包含下面这部分代码，你可以看到这部分代码出现了两次：
```js
while (root) {
    stack.push(root);
    root = root.left;
}
```
其实这两次出现，分别对应了下面的两次 inorder 调用：
```js
inorder(root.left);
res.push(root.val);
inorder(root.right);
```

#### 迭代实现 代码
```js
const inorderTraversal = (root) => {
  const res = [];
  const stack = [];

  while (root) {        // 先把能压入栈的左子节点都压进来
    stack.push(root);
    root = root.left;
  }
  while (stack.length) {
    let root = stack.pop(); // 栈顶的节点出栈
    res.push(root.val);     // 在压入右子树之前，处理它的数值部分（中序遍历的定义）
    root = root.right;      // 获取它的右子树
    while (root) {          // 右子树存在，执行while循环    
      stack.push(root);     // 压入当前root
      root = root.left;     // 不断压入左子节点
    }
  }
  return res;
};
```
大家可以和递归写法对比一下，感受一下，区别在哪里，相同点在哪里。
#### 全流程大致图示
![image.png](https://pic.leetcode-cn.com/1600046351-mjZerJ-image.png)

#### 后记
时间有限，暂时分析中序遍历，前序遍历和后序遍历的迭代的分析之后再补充上来。

明确这三种遍历都是基于DFS递归，清楚了概念，再明白递归其实是压栈出栈的操作，比照着，自己用一个栈，去模拟递归栈，用循环迭代，去类比，去模拟递归的逻辑，就不难写出迭代法的代码。
#### 如果觉得有帮助，点个赞支持我继续总结分享写字画图，如果哪里不对或可以更好，指出我我继续修改。