"+++
title = "0236. Lowest Common Ancestor of a Binary Tree Marvel中等的学习笔记-236 "
author = ["tyanyonecancode"]
date = 2020-08-25T02:32:08+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# Marvel中等的学习笔记-236

> [0236. Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
> [Marvel中等的学习笔记-236](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/marvelzhong-deng-de-xue-xi-bi-ji-236-by-tyanyone-2/) by [tyanyonecancode](https://leetcode-cn.com/u/tyanyonecancode/)

### 思路
我们不妨定义一个方法findLCA()，用于检测当前树中是否包含p或q，是则返回true，否则返回false。

在当前树中包含p或q无外乎以下三种情况：
1. 当前根结点就是p或q；
2. 左子树中有p或q；
3. 右子树中有p或q。

第一种情况可直接求得，用变量onCur记录结果。
第二第三种情况则可递归求解，分别用变量onLeft和onRight记录结果。

得到三个变量后，若
x本身就是p或q，然后q或p在x的某一子树中，即`onCur && (onLeft || onRight)`，或者
p、q一个在x的左树中一个在x的右树中，即`onLeft && onRight`
则意为找到LCA，将其赋值到实例变量ans中。

最后，将该树是否包含p或q的结果向上返回，即`return onCur || onLeft || onRight;`。

递归的边界为：
若遇到空树，则返回false，意为空树中不含p或q；
若ans已赋值，说明LCA已找到，直接返回false，意为其他结点都不会是LCA。

这样可以保证得到的公共祖先一定是最近的，即深度最大的。
因为整个递归过程是以自底向上的形式进行的，发现的第一个公共祖先，一定是深度最大的，即最近公共祖先。

代码

```java
class Solution {
    private TreeNode ans;
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        findLCA(root, p, q);
        return ans;
    }
    private boolean findLCA(TreeNode r, TreeNode p, TreeNode q) {
        if(ans != null)    return false; //已找到LCA，无需再检查其他结点
        if(r == null)    return false;
        boolean onLeft = findLCA(r.left, p, q); //pq是否在左子树
        boolean onRight = findLCA(r.right, p, q); //pq是否在右子树
        boolean onCur = r == p || r == q; //pq是否为当前结点
        if((onLeft && onRight) || (onCur && (onLeft || onRight))) //满足最近公共祖先
            ans = r;
        return onCur || onLeft || onRight; //pq是否在当前树中
    }
}
```