"+++
title = "0617. Merge Two Binary Trees 617. 合并二叉树（递归 | 先序遍历） "
author = ["yiluolion"]
date = 2020-09-23T09:29:41+08:00
tags = ["Leetcode", "Algorithms", "Python", "Recursion", "前序遍历"]
categories = ["leetcode"]
draft = false
+++

# 617. 合并二叉树（递归 | 先序遍历）

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [617. 合并二叉树（递归 | 先序遍历）](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/617-he-bing-er-cha-shu-di-gui-xian-xu-bian-li-by-y/) by [yiluolion](https://leetcode-cn.com/u/yiluolion/)

## 思路：递归

这道题中，其实题目中给出的提示还是比较多的。我们先审题，题目中给定两个二叉树，想象其中一个二叉树覆盖在另外一个上面，将重叠的节点进行合并。这里合并的规则如下：

- 若两个节点重叠（这里其实更明确地说是两个不为 None 的节点），将他们的值相加作为节点合并后的新值；
- 如果两个节点重叠（但节点存在为 None 的情况），这个时候，将不为 None 的节点作为合并后二叉树的节点。

> 【否则不为 NULL 的节点将直接作为新二叉树的节点】，题目中这一句最后， “新二叉树的节点”，这里存疑。个人理解是合并后的二叉树就是这里所述的新二叉树，而不是一定要创建新的节点。所以本篇后面的代码直接将 t2 合并到 t1。

题目最后的 **注意** 部分，说明合并必须从两个树的根节点开始。那么，这里我们用先序遍历来处理，在进入子树递归前，先处理值。

#### 递归

这里整理下，使用递归的做法：

- 首先，根据题目提供的规则，设置终止条件：
  - 当 t1 为空时，返回 t2 作为新二叉树的节点；
  - 当 t2 为空时，返回 t1 作为新二叉树的节点。
- 使用先序遍历，递归逐层逻辑的处理（这里将 t2 合并到 t1）：
  - 将两个二叉树重叠的节点值相加，合并到 t1；
  - 递归左子树；
  - 递归右子树。

这里用简单图示来进一步解析上述做法：

![图示](https://pic.leetcode-cn.com/1600853227-USsVUW-617_%E5%90%88%E5%B9%B6%E4%BA%8C%E5%8F%89%E6%A0%91.gif)



具体的代码如下。

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:

        def dfs(t1, t2):
            # t1 为空，返回 t2
            if not t1:
                return t2
            # t2 为空，返回 t1
            if not t2:
                return t1
            
            # 先序遍历，先处理值
            t1.val += t2.val

            t1.left = dfs(t1.left, t2.left)
            t1.right = dfs(t1.right, t2.right)
        
            return t1

        return dfs(t1, t2)
```

## 欢迎关注

---

公众号 【[书所集录](https://i.loli.net/2020/07/09/sNEGeV8g6fmW5Ub.jpg)】

---

如有错误，烦请指出，欢迎指点交流。