"+++
title = "0206. Reverse Linked List [Java] Iterative solution. "
author = ["lincs"]
date = 2020-08-31T01:24:51+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# [Java] Iterative solution.

> [0206. Reverse Linked List](https://leetcode-cn.com/problems/reverse-linked-list/)
> [[Java] Iterative solution.](https://leetcode-cn.com/problems/reverse-linked-list/solution/java-iterative-solution-by-lincs/) by [lincs](https://leetcode-cn.com/u/lincs/)

```java
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode prev = null;
        ListNode cur  = head;
        while (cur != null) {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
        return prev;
    }
}
```
