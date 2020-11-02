"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 俺是fw "
author = ["Tcl_love"]
date = 2020-09-27T04:59:20+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 俺是fw

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [俺是fw](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/an-shi-fw-by-tcl_love/) by [Tcl_love](https://leetcode-cn.com/u/tcl_love/)

### 思路
思路:
可以把遍历想像成多条路寻找节点
A情况: p = 2, q = 8 找到p,q需要两条路[6->2]和[6->8] 公共祖先是6
B情况: p = 2 ,q = 4 找到p,q需要一条路[6->2->4] 公共祖先是2

q和p的值小于当前节点值,说明q,p在左子树,去左子树遍历
q和p的值大于当前节点值,说明q,p在右子树,去右子树遍历
其他情况(q等于当前节点 或 p等于当前节点 或 p大于当前节点值,q小于当前节点值 或 p小于当前节点值,q大于当前节点值)退出循环返回结果

### 代码
```java
public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        boolean flag = true;
        while(flag){
            //q和p的值小于当前节点值,说明q,p在左子树,去左子树遍历
            if(p.val < root.val && q.val < root.val){
                root=root.left;
            //q和p的值大于当前节点值,说明q,p在右子树,去右子树遍历
            }else if(p.val > root.val && q.val > root.val){
                root=root.right;
            }else{
                flag=false;
            }
        }
        return root;
    }
```