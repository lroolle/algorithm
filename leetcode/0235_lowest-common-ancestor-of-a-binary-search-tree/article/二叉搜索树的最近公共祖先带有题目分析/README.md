"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 二叉搜索树的最近公共祖先，带有题目分析 "
author = ["zglg"]
date = 2020-08-02T15:19:35+08:00
tags = ["Leetcode", "Algorithms", "Python", "二分搜索"]
categories = ["leetcode"]
draft = false
+++

# 二叉搜索树的最近公共祖先，带有题目分析

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [二叉搜索树的最近公共祖先，带有题目分析](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-da/) by [zglg](https://leetcode-cn.com/u/zglg/)

## Day 59: 二叉搜索树的最近公共祖先

### 1 题目

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

示例 1:

```markdown
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6 
解释: 节点 2 和节点 8 的最近公共祖先是 6。
```
示例 2:

```markdown
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
```
说明:

- 所有节点的值都是唯一的。
- p、q 为不同节点且均存在于给定的二叉搜索树中。

### 2 分析

对于二叉搜索树，已知两个节点，寻找最近祖先节点，根据下面的递推关系式：

若 node.val < min(p.val,q.val)，则p和q的最近祖先节点一定在右子树；

若 max(p.val,q.val) < node.val，则p和q的最近祖先节点一定在左子树；

其他情况 node.val位于p.val和q.val间(可能等于node.val)，则node就是p和q的最近祖先。

### 3 代码

```python
class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        if root.val < min(p.val,q.val):
            return self.lowestCommonAncestor(root.right,p,q)
        elif max(p.val,q.val) < root.val:
            return self.lowestCommonAncestor(root.left,p,q)
        else:
            return root
```

若需学习到更多算法解题思路，欢迎关注我的👉👉👉 [公众号：算法刷题日记](https://static01.imgkr.com/temp/024a2b9488ea4b2597c8b89d0173584f.png)👈👈👈

如果能再点个赞👍👍 就更感激啦💓💓
