"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree Marvel简单的学习笔记-235 "
author = ["tyanyonecancode"]
date = 2020-03-06T15:17:53+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# Marvel简单的学习笔记-235

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [Marvel简单的学习笔记-235](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/marveljian-dan-de-xue-xi-bi-ji-235-by-marvel_ty/) by [tyanyonecancode](https://leetcode-cn.com/u/tyanyonecancode/)

### 解法一：递归
题目要求从根结点为root的树返回p、q的最近公共祖先，使用递归，很自然地，如果p、q均小于root，则应该由root的左子树返回p、q的最近公共祖先；如果p、q均大于root，则应该由root的右子树返回p、q的最近公共祖先。
由此可以写出递归代码。
直到p、q既不同时大于root，也不同时小于root，则root就是p、q的最近公共祖先，返回root。

代码：
```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if(p.val < root.val && q.val < root.val)
            return lowestCommonAncestor(root.left, p, q);
        if(p.val > root.val && q.val > root.val)
            return lowestCommonAncestor(root.right, p, q);
        return root;
    }
}
```

### 解法二： 迭代
这道题迭代应该比递归容易写。
从根结点开始顺着往下找，如果p、q均小于root，则到左子树中找；如果p、q均大于root，则到右子树中找；
其余情况包括：p、q一个比root大一个比root小，或者p、q中有一个就是root，以上均意味当前root就是p、q的最近公共祖先。

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        while(root != null)
        {
            if(p.val < root.val && q.val < root.val)
                root = root.left;
            else if(p.val > root.val && q.val > root.val)
                root = root.right;
            else    return root;
        }
        return null;
    }
}
```