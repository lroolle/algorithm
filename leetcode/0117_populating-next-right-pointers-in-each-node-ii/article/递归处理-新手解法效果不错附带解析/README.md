"+++
title = "0117. Populating Next Right Pointers in Each Node II 递归处理 | 新手解法（效果不错，附带解析） "
author = ["finlu"]
date = 2020-06-30T01:05:11+08:00
tags = ["Leetcode", "Algorithms", "Java", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 递归处理 | 新手解法（效果不错，附带解析）

> [0117. Populating Next Right Pointers in Each Node II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)
> [递归处理 | 新手解法（效果不错，附带解析）](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/di-gui-chu-li-xin-shou-jie-fa-xiao-guo-bu-cuo-fu-d/) by [finlu](https://leetcode-cn.com/u/finlu/)

![image.png](https://pic.leetcode-cn.com/6ccb17260c5ebbc48a18892b40ecffde8bdeb4fd5c72c36465ebc9a90a9fae3a-image.png)
#### 分析
题目中所给的是一颗二叉树，单并没有说是一棵完全二叉树，所以节点对应的子树可能没有左子树或者没有右子树或者二者都没有。这时候我的思路就是使用递归并结合==自顶向下==的思想来解决。

自顶向下的分析思想：根据当前节点已有的信息寻找答案，如果需要传递给子节点参数信息，则将参数信息进行传递。

好了，接下来看一下题的解题思路吧！

### 解题思路

1. 首先确定递归条件：递归函数传入的是叶节点的孩子节点的时候结束递归
2. 确定每个节点需要做的处理：
3. (1) 当前节点的左节点不为空时，将左节点的next指针指向当前节点的右节点（不需要考虑右节点是否为空）；
4. (2) 如果当前节点的next值不为null，则需要先找到当前节点的最右侧节点R（右节点不为null的时候为右节点，否则为左节点），之后再找到第一个左右节点不全为null的节点，找出其子节点中最左的不为null的节点L，当节点R和和节点L都不为空的时候，将R的next设置为L
5. 递归处理右子树
6. 递归处理左子树

为什么先处理右子树？
答：因为再递归过程中，需要一直找到最右侧节点，所以必须保证右侧的next指针指向是正确的！

### 代码

```java
/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}
    
    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/

class Solution {
    private void dfs(Node root) {
        if (root == null) return;
        if (root.left != null) {
            root.left.next = root.right;
        }
        if (root.next != null) {
            Node leftNode = root.right != null ? root.right : root.left;
            Node rightRootNode = root.next;
            while(rightRootNode != null && rightRootNode.left == null && rightRootNode.right == null) {
                rightRootNode = rightRootNode.next;
            }
            if (leftNode != null && rightRootNode != null) {
                Node rightNode = rightRootNode.left != null ? rightRootNode.left : rightRootNode.right;
                leftNode.next = rightNode;
            }
        }
        dfs(root.right);
        dfs(root.left);
    }
    
    public Node connect(Node root) {
        dfs(root);
        return root;  
    }
}
```