"+++
title = "0025. Reverse Nodes in k-Group 递归java "
author = ["reedfan"]
date = 2020-03-15T13:16:54+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 递归java

> [0025. Reverse Nodes in k-Group](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
> [递归java](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/di-gui-java-by-reedfan-2/) by [reedfan](https://leetcode-cn.com/u/reedfan/)

#### 分析
大致过程可以分解为
1、找到待翻转的k个节点（注意：若剩余数量小于 k 的话，则不需要反转，因此直接返回待翻转部分的头结点即可）。
2、对其进行翻转。并返回翻转后的头结点（注意：翻转为左闭又开区间，所以本轮操作的尾结点其实就是下一轮操作的头结点）。
3、对下一轮 k 个节点也进行翻转操作。
4、将上一轮翻转后的尾结点指向下一轮翻转后的头节点，即将每一轮翻转的k的节点连接起来。
####  具体过程看下图。
![](https://pic.leetcode-cn.com/f63d5ca4d3f055ce8e4591c8bc51c288791f88da9ccec9617bc8bb51c26163a2.png)
#### 代码
```Java
    public ListNode reverseKGroup(ListNode head, int k) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode tail = head;
        for (int i = 0; i < k; i++) {
            //剩余数量小于k的话，则不需要反转。
            if (tail == null) {
                return head;
            }
            tail = tail.next;
        }
        // 反转前 k 个元素
        ListNode newHead = reverse(head, tail);
        //下一轮的开始的地方就是tail
        head.next = reverseKGroup(tail, k);

        return newHead;
    }

    /*
    左闭又开区间
     */
    private ListNode reverse(ListNode head, ListNode tail) {
        ListNode pre = null;
        ListNode next = null;
        while (head != tail) {
            next = head.next;
            head.next = pre;
            pre = head;
            head = next;
        }
        return pre;

    }
```
