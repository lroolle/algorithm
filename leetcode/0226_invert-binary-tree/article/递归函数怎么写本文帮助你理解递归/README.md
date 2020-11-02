"+++
title = "0226. Invert Binary Tree 递归函数怎么写？本文帮助你理解递归 "
author = ["fuxuemingzhu"]
date = 2020-09-15T17:20:24+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 递归函数怎么写？本文帮助你理解递归

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [递归函数怎么写？本文帮助你理解递归](https://leetcode-cn.com/problems/invert-binary-tree/solution/di-gui-han-shu-zen-yao-xie-ben-wen-bang-zhu-ni-li-/) by [fuxuemingzhu](https://leetcode-cn.com/u/fuxuemingzhu/)

这个题能够很好地帮助我们理解递归。

递归函数本身也是函数，调用递归函数就把它当做普通函数来看待，一定要只思考当前层的处理逻辑，明白该递归函数的输入输出是什么即可，调用的时候不要管函数内部实现。不要用肉脑 debug 递归函数的调用过程，会被绕进去。

首先来分析`invertTree(TreeNode root)`函数的定义：

1. 函数的定义是什么？
该函数可以翻转一棵二叉树，即**将二叉树中的每个节点的左右孩子都进行互换。**
2. 函数的输入是什么？
函数的输入是要被翻转的二叉树。
3. 函数的输出是什么？
返回的结果就是已经翻转后的二叉树。

然后我们来分析函数的写法：

1. 递归终止的条件
当要翻转的节点是空，停止翻转，返回空节点。
2. 返回值
虽然对 `root` 的左右子树都进行了翻转，但是翻转后的二叉树的根节点不变，故返回 `root` 节点。
3. 函数内容
`root` 节点的新的左子树：是翻转了的 `root.right` => 即 `root.left = invert(root.right)`;
`root` 节点的新的右子树：是翻转了的 `root.left` => 即 `root.right = invert(root.left)`;

至此，递归函数就写完了。在『函数内容』编写的时候，是不是把递归函数`invertTree(TreeNode root)`当做了普通函数来用？调用`invertTree(TreeNode root)`函数就是能实现翻转二叉树的目的，不需要理解函数内部怎么实现的。

最后，提醒大家避免踩一个小坑，不能直接写成下面这样的代码：

```python
root.left = invert(root.right)
root.right = invert(root.left)
```

这是因为第一行修改了`root.left`，会影响了第二行。在 Python 中，正确的写法是把两行写在同一行，就能保证 `root.left` 和 `root.right` 的修改是同时进行的。

Python 解法如下：

```python
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def invertTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if not root:
            return
        root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
        return root
```
#### 欢迎关注[负雪明烛的刷题博客](https://blog.csdn.net/fuxuemingzhu)，刷题800多，每道都记录了写法！目前访问量已经100多万次！

#### 送大家刷题模板和套路总结：[【LeetCode】代码模板，刷题必会](https://blog.csdn.net/fuxuemingzhu/article/details/101900729)

#### 力扣每日一题打卡群，大家一起监督力扣刷题：[https://ojeveryday.com/](https://ojeveryday.com/)

