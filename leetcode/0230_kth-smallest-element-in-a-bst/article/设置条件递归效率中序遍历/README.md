"+++
title = "0230. Kth Smallest Element in a BST 设置条件递归效率中序遍历 "
author = ["IvTCQF8cRd"]
date = 2020-08-16T13:31:03+08:00
tags = ["Leetcode", "Algorithms", "Java", "BinarySearchTree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 设置条件递归效率中序遍历

> [0230. Kth Smallest Element in a BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)
> [设置条件递归效率中序遍历](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/solution/she-zhi-tiao-jian-di-gui-xiao-lu-zhong-xu-bian-li-/) by [IvTCQF8cRd](https://leetcode-cn.com/u/IvTCQF8cRd/)

设置条件递归效率高点 



    int count = 1;
    int val = -1;
    /**
     * 二叉搜索树中找出第K小的元素
     * @param root
     * @param k
     * @return
     */
    public int kthSmallest(Node root, int k) {
        if(root==null) return 0;
        if(val==-1) kthSmallest(root.left,k);
        if(count==k)  val = root.val;
        count++;
        if(val==-1) kthSmallest(root.right,k);
        return val;
    }