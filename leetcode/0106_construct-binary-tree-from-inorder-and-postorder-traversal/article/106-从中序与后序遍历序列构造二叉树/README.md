"+++
title = "0106. Construct Binary Tree from Inorder and Postorder Traversal 106. 从中序与后序遍历序列构造二叉树 "
author = ["akian"]
date = 2020-08-28T17:19:45+08:00
tags = ["Leetcode", "Algorithms", "Java", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 106. 从中序与后序遍历序列构造二叉树

> [0106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
> [106. 从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/106-cong-zhong-xu-yu-hou-xu-bian-li-xu-lie-gou-13/) by [akian](https://leetcode-cn.com/u/akian/)

**简单介绍：**
题目：从中序与后序遍历序列构造二叉树
题目难度：中等
使用语言：JAVA
这道题来自leetcode题库的深度优先搜索标签。
**解题思路：**
**主要解法：递归。**

**利用后序遍历来寻找root结点，再使用root结点把中序遍历，分成左子树，根结点，右子树。不断递归，直到子树为叶子结点。
重点：前序遍历、中序遍历左右子树边界位置需要确定好。可以根据下图来确定边界。
小技巧：使用hashmap保存中序遍历元素的索引位置。**

![image.png](https://pic.leetcode-cn.com/1598635143-eCSPLz-image.png)
**数据结构：**
要实现对数据的操作，我们要先明确存储数据的数据结构。
该题的数据结构的作用：
**1.HashMap保存inorder数组元素的索引。**
**算法：**
既然明确了我们的数据结构，我们就可以开始我们的算法分析了。
**1.第一步，初始化工作。
2.第二步，递归建立树
3.buildTreeTool：建立树的工具。后序遍历寻找到root结点，在使用root结点结合中序遍历建立root结点的左右子树。**
**代码部分：**

```java

public class Solution {
    public TreeNode buildTree(int[] inorder, int[] postorder) {
        HashMap<Integer,Integer> map=new HashMap<>();
        for(int i=0;i<inorder.length;i++){
            map.put(inorder[i],i);
        }
        return buildTreeTool(inorder,0,inorder.length,postorder,0,postorder.length,map);
    }

    public TreeNode buildTreeTool(int[] inorder,int i_start,int i_end,int[] postorder,int p_start,int p_end,HashMap<Integer,Integer> map){
        if(p_start==p_end) return null;
        int root_val=postorder[p_end-1];
        TreeNode root=new TreeNode(root_val);
        int i_root_index=map.get(root_val);
        int right_num=i_end-i_root_index-1;

        root.left=buildTreeTool(inorder,i_start,i_root_index,postorder,p_start,p_end-right_num-1,map);
        root.right=buildTreeTool(inorder,i_root_index+1,i_end,postorder,p_end-right_num-1,p_end-1,map);

        return root;

    }
}
```

**时间、空间复杂度：**

![image.png](https://pic.leetcode-cn.com/1598635112-YouMtu-image.png)
**结语：**
晚安！晚安！晚安！晚安！晚安！晚安！晚安！晚安！晚安！晚安！晚安！
