"+++
title = "0701. Insert into a Binary Search Tree Java 递归 "
author = ["mmmmmJCY"]
date = 2019-07-09T07:27:51+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# Java 递归

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [Java 递归](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/java-di-gui-by-zxy0917/) by [mmmmmJCY](https://leetcode-cn.com/u/mmmmmjcy/)

递归结束的条件：到达最后一层的目标节点时直接返回新节点。

因为这里只插入一个节点，所以要不就是左子树改变，要不就是右子树数改变，根据BST的性质找到插入点的位置即可。

```
 public TreeNode insertIntoBST(TreeNode root, int val) {
    if(root == null) return new TreeNode(val);
    //根据BST的性质找到需要插入的点
    if(root.val > val){
        root.left =  insertIntoBST(root.left,val);
    }else{
        root.right = insertIntoBST(root.right,val);
    }
    return root;
}
```

