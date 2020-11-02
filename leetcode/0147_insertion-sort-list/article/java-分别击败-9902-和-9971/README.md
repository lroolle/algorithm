"+++
title = "0147. Insertion Sort List java 分别击败 99.02% 和 99.71%  "
author = ["ning-meng-lll"]
date = 2020-09-21T13:50:38+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# java 分别击败 99.02% 和 99.71% 

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [java 分别击败 99.02% 和 99.71% ](https://leetcode-cn.com/problems/insertion-sort-list/solution/zhi-xing-yong-shi-3-ms-zai-suo-you-java-ti-jiao-56/) by [ning-meng-lll](https://leetcode-cn.com/u/ning-meng-lll/)

### 解题思路
插入排序，两层循环

### 代码

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode insertionSortList(ListNode head) {
        if (head==null||head.next==null)
            return head;

        ListNode dummy = new ListNode(Integer.MIN_VALUE);
        dummy.next = head;
        ListNode cur = head.next;
        ListNode pre = head;
        while (cur!=null){
            ListNode temp = cur.next;
            if (cur.val<pre.val){
                ListNode start = dummy;
                while (cur.val>start.next.val){
                    start = start.next;
                }
                pre.next = cur.next;
                cur.next = start.next;
                start.next = cur;

                cur = temp;
            }
            else {
                cur = cur.next;
                pre = pre.next;
            }

        }
        return dummy.next;
    }
}
```