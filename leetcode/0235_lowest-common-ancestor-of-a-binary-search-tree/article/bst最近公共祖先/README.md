"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree BST最近公共祖先 "
author = ["beney-2"]
date = 2020-09-10T01:18:20+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree"]
categories = ["leetcode"]
draft = false
+++

# BST最近公共祖先

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [BST最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/bstzui-jin-gong-gong-zu-xian-by-beney-2/) by [beney-2](https://leetcode-cn.com/u/beney-2/)

### 解题思路
与这题非常类似的还有二叉树的最近公共祖先，对于BST这种特殊的二叉树我们可以充分利用其性质：对于树上的每一个**子树根root**
1. 任意`root`左子树的节点`p`，有`p.val < root.val`
2. 任意`root`右子树的节点`q`，有`q.val > root.val`

那么我们可推出如下结论：
对于要查找LCA的两个节点`p`，`q`：
1. 若`p`、`q`的值都**小于**`root`节点，那么`p`、`q`的LCA必然在`root`的左子树。
2. 若`p`、`q`的值都**大于**`root`节点，那么`p`、`q`的LCA必然在`root`的右子树。
3. 否则，`root`就是`p`、`q`搜索路径的**分岔点**，`root`就是要找的LCA。

这样遍历的长度最多就是树的高度，时间复杂度为$O(lgn)$。
### 代码

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */

class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        TreeNode it = root;
        while (it != null) {
            if (p.val < it.val && q.val < it.val)
                it = it.left;
            else if (p.val > it.val && q.val > it.val)
                it = it.right;
            else
                return it;
        }
        return null; 
    }
}
```
### 对你有帮助就点个棒棒吧！欢迎评论:)