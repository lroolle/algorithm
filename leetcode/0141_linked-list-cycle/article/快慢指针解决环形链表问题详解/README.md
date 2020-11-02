"+++
title = "0141. Linked List Cycle 快慢指针解决环形链表问题详解 "
author = ["1234-52"]
date = 2020-09-26T12:59:08+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 快慢指针解决环形链表问题详解

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [快慢指针解决环形链表问题详解](https://leetcode-cn.com/problems/linked-list-cycle/solution/kuai-man-zhi-zhen-jie-jue-huan-xing-lian-biao-we-2/) by [1234-52](https://leetcode-cn.com/u/1234-52/)

一.解题思路：
1. 判空
2. 设置慢指针l=head;快指针r=head,
3. 相当于两人从同一起跑线开始跑步，l每次跑一步，r每次跑2步，如果是环形跑道，l，r终究会相遇
4. 如果没有环，则r一定先跑到终点,所以无需判断l是否到达终点
5. 只要r!=null(r没到终点)且r.next != null(r跑一步也不会到终点),则两人一直循环跑下去
6. 每循环一次，l跑一步l = l.next，r跑两步r = r.next.next
7. 跑完后判断下，l == r（l跟r是否相遇了），相遇表示有环返回true
8. 如果循环结束，标示r跑到了终点还没跟l相遇，则没有环，返回false

二.时间复杂度：o(N)
三.空间复杂度：o(1))
四.执行用时: 0 ms,超过来100%的java提交记录
五.内存消耗: 38.9 MB,超过了65.13%%的java提交记录
六.java代码：
```
public class Solution {
    public boolean hasCycle(ListNode head) {
        if (head == null) {
            return false;
        }
        ListNode l = head;
        ListNode r = head;
        while (r != null && r.next != null) {
            l = l.next;
            r = r.next.next;
            if (l == r) {
                return true;
            }
        }
        return false;
    }
}
```