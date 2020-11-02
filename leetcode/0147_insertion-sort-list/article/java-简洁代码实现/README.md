"+++
title = "0147. Insertion Sort List Java 简洁代码实现 "
author = ["gary-liu"]
date = 2020-04-23T16:08:20+08:00
tags = ["Leetcode", "Algorithms", "Java", "LinkedList"]
categories = ["leetcode"]
draft = false
+++

# Java 简洁代码实现

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [Java 简洁代码实现](https://leetcode-cn.com/problems/insertion-sort-list/solution/java-jian-ji-dai-ma-shi-xian-by-serenade0816/) by [gary-liu](https://leetcode-cn.com/u/gary-liu/)

### 解题思路
关键是几个链表节点的作用：
1. `dummy`指向第一个节点，一是用于结果输出，二是寻找插入位置时用于回到起点
2. `cur`用于遍历的节点
3. `pre`一是用于和当前节点作比较，二是做标记，让交换后的cur回到pre后面

### 代码

```java
class Solution {
    public ListNode insertionSortList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode dummy = new ListNode(0);
        ListNode pre = head;
        ListNode cur = head.next;
        dummy.next = head;
        while (cur != null) {
            if (pre.val <= cur.val) {// 本来就有序
                pre = cur;
                cur = cur.next;
            } else {
                ListNode p = dummy;
                // 找到一个位置使得p < cur < p.next
                while (p.next != cur && p.next.val < cur.val) {
                    p = p.next;
                }
                // 将cur插入到p和p.next之间
                // 因为cur被插到前面，pre的指针需要跳过cur
                pre.next = cur.next;
                cur.next = p.next;
                p.next = cur;
                // 完成插入后，cur回到pre后面
                cur = pre.next;
            }
        }
        return dummy.next;
    }
}
```