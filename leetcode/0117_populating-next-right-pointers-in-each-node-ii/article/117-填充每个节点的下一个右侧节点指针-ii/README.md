"+++
title = "0117. Populating Next Right Pointers in Each Node II 117: 填充每个节点的下一个右侧节点指针 II "
author = ["wulin-v"]
date = 2020-08-08T06:12:40+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 117: 填充每个节点的下一个右侧节点指针 II

> [0117. Populating Next Right Pointers in Each Node II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)
> [117: 填充每个节点的下一个右侧节点指针 II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/117tian-chong-jie-dian-de-xia-yi-ge-jie-dian-you-c/) by [wulin-v](https://leetcode-cn.com/u/wulin-v/)

### 解题思路
**解法一** 
本题解法一和116题解法一相同，不要要做任何修改
详解过程参见[116：填充每个节点的下一个右侧节点指针](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/solution/116tian-chong-jie-dian-de-xia-yi-ge-jie-dian-you-c/)
**解法二**
从左向右看一个树的每一层就是一个单链表，
在每一层的开始时创建一个虚拟头结点`dumm`

变量说明
- `cur` 是每层的第一个节点
- 对于同一层节点来说 `cur.next` 的作用是向下走一步
- 虚拟头节点`dumm`的`next`就是位于`cur`层的下一层的第一个节点
- `tail`指针是当前要修改next指向的节点，也是实际操作`next`的指针
- `tail`和`tail.next`都是不跨层的

代码如下

### 代码

```java
public Node connect(Node root) {
    if(root==null)
        return root;
    Node cur=root;
    while(cur!=null){           //控制cur到下一层的循环
        Node dumm=new Node();   //创建一个虚拟头结点(每一层都会创建)
        Node tail=dumm;         //维护一个尾节点指针（初始化是虚拟节点）

        while(cur!=null){        //控制cur同一层的循环
            if(cur.left!=null){  //判断cur的左节点是否为空，不为空时就是cur的下一层的第一个节点了
                tail.next=cur.left;
                tail=tail.next;
            }
            if(cur.right!=null){  //判断cur的右节点是否为空，此时不为空时就是cur的下一层的第一个节点了
                tail.next=cur.right;
                tail=tail.next;
            }
            cur=cur.next;         //cur同层移动到下一位置
        }
        cur=dumm.next;            //内循环结束，开始cur的下一层
    }
    return root;
}
```