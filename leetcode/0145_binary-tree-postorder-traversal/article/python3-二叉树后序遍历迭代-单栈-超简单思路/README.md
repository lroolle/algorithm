"+++
title = "0145. Binary Tree Postorder Traversal python3 二叉树后序遍历迭代 单栈 超简单思路 "
author = ["harris-han"]
date = 2020-07-09T13:10:11+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3"]
categories = ["leetcode"]
draft = false
+++

# python3 二叉树后序遍历迭代 单栈 超简单思路

> [0145. Binary Tree Postorder Traversal](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)
> [python3 二叉树后序遍历迭代 单栈 超简单思路](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/python3-er-cha-shu-hou-xu-bian-li-die-dai-fang-fa-/) by [harris-han](https://leetcode-cn.com/u/harris-han/)


让我们先看后序遍历的顺序
left right root

接着来看前序遍历的顺序
root left right

如果我们把后序遍历翻转，将会得到：
root right left

相比较前序遍历，仅仅改变了left 和 right的顺序：
所以本题思路将会是：在前序遍历中，把left 和 right的顺序调换，然后输出反转的树即可。

上代码：

```
class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        if root==None:
            return []
        stack=[root]
        res=[]
        while stack:
            s=stack.pop()
            res.append(s.val)
            if s.left:#与前序遍历相反
                stack.append(s.left)
            if s.right:
                stack.append(s.right)
        return res[::-1]
```
