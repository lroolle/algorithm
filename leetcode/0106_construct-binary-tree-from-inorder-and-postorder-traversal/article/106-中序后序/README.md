"+++
title = "0106. Construct Binary Tree from Inorder and Postorder Traversal 106: 中序+后序 "
author = ["wulin-v"]
date = 2020-08-05T16:19:20+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 106: 中序+后序

> [0106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
> [106: 中序+后序](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/106-zhong-xu-hou-xu-by-wulin-v/) by [wulin-v](https://leetcode-cn.com/u/wulin-v/)

### 解题思路
Leetcode的105、106、889这三题的解题思路都是一样的，主要是注意一下两点：
1、三个题目迭代构建子树的区间变化
2、三个题目创建树的根节点是不样的

- Leetcode106题 [中序+后序](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/106-zhong-xu-hou-xu-by-wulin-v/)
- Leetcode105题 [前序+后序](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/105qian-xu-zhong-xu-by-wulin-v/)
- Leetcode889题 [前序+中序](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/solution/889qian-xu-hou-xu-zhu-shi-xiang-jie-by-wulin-v/)

变量说明：
- in 中序遍历的数组
- post 后序遍历的数组
- inL 中序遍历的第一个节点
- inR 中序遍历的最后一个节点
- postL 后序遍历的第一个节点
- postR 后序遍历的最后一个节点

图片说明：
![106.png](https://pic.leetcode-cn.com/45c11cc528868a97c9b28a2a1361e1c9d56605b7ab43350964e2561549b0170d-106.png)
详细过程参考，如有错误请留言
### 代码

```java
    class Solution{
    Map<Integer,Integer> map = new HashMap<>();
        public TreeNode buildTree(int[] inorder, int[] postorder) {
        //为了和105、889题目结构一致都利用HashMap方便查找节点位置（索引）
        //我们利用中序来区分左右子数
        for(int i=0;i<inorder.length;i++) 
            map.put(inorder[i],i);
        return helper(inorder,postorder,0,inorder.length-1,0,postorder.length-1);
    }
    private TreeNode helper(int[] in, int[] post, int inL, int inR, int postL, int postR){
        if (inL>inR)  //说明数组无效
            return null;
        
        TreeNode root=new TreeNode(post[postR]);//利用后序的最后一个节点创建新的树的根节点
        int postRootVal=post[postR];    //后序遍历的最后一个节点post_last_node
        int inMid=map.get(postRootVal); //后序遍历节点post_last_node在中序遍历中的位置（利用map直接以O(1)的复杂度取出索引）
        int inLeftNum=inMid-inL;        //中序遍历的左子树的个数
        /*
            中序左子树的范围说明:
                从中序第一个节点开始，到后序遍历的最后一个节点在中序遍历位置的前一个节点结束 inL,inMid-1
            中序右子树的范围说明:
                从后序遍历的最后一个节点在中序遍历位置的后一个节点开始，到中序的最后一个位置结束 inMid+1,inR
            后序左子树的范围说明:
                从后序的第一个的节点开始，到第一个节点+左子树的个数-1结束 postL,postL+inLeftNum-1
            后序右子树的范围说明:
                从后序的第一个节点+左子树的个数开始，到到后序的最后一个位置-1结束 postL+inLeftNum,postR-1

            中序左子树的范围[inL,inMid-1] 后序左子树的范围[postL,postL+inLeftNum-1]
            中序右子树的范围[inMid+1,inR] 后序右子树的范围[postL+inLeftNum,postR-1]
        */

        root.left=helper(in,post,inL,inMid-1,postL,postL+inLeftNum-1);//递归构建左子树
        root.right=helper(in,post,inMid+1,inR,postL+inLeftNum,postR-1);//递归构建右子树
        return root;
    }
}

```
**注意：**
- 同一棵树中序和后序遍历左子树的节点数是相同的，所以**从后序遍历的第一个节点加上中序确定的左节点个数，再减去一**就可以确定出后序遍历中左子树的节点范围
- 后序的右节点是从去除左节点开始的(后序中除去左节点个数就剩右节点和根了)所以是**从后序的第一个节点+左子树的个数开始，到后序的最后一个节点减一(减的是根)**

千万不要被后序遍历的顺序给骗了，利用后序遍历的顺序 (右 左 根)做题 (你会被坑的连你爸都不认识你。。信不信由你)后序遍历顺序是  右 左 根是对的。但是，但是，她也是**从最左边子树的最左边的子节点开始遍历**(请认真思考，这意味着什么)，这意味着后序遍历的开始打印出来的都是树的左边的值。