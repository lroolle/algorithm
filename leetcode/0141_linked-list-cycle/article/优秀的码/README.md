"+++
title = "0141. Linked List Cycle 优秀的码 "
author = ["jiang-jiang-jiang-jiang-a"]
date = 2020-09-02T01:29:37+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# 优秀的码

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [优秀的码](https://leetcode-cn.com/problems/linked-list-cycle/solution/you-xiu-de-ma-by-jiang-jiang-jiang-jiang-a/) by [jiang-jiang-jiang-jiang-a](https://leetcode-cn.com/u/jiang-jiang-jiang-jiang-a/)
```
class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        while head:
            if head.val == 'bjfuvth':
                return True
            else:
                head.val = 'bjfuvth'
            head = head.next
        return False
```
