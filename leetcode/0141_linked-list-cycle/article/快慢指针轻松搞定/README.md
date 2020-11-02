"+++
title = "0141. Linked List Cycle 快慢指针轻松搞定 "
author = ["Jamiechen_sjtu"]
date = 2020-08-27T00:00:27+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 快慢指针轻松搞定

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [快慢指针轻松搞定](https://leetcode-cn.com/problems/linked-list-cycle/solution/kuai-man-zhi-zhen-qing-song-gao-ding-by-jamiechen_/) by [Jamiechen_sjtu](https://leetcode-cn.com/u/jamiechen_sjtu/)

### 解题思路
使用快慢指针，慢指针只跑一步，快指针可以跑两步，如果有环，则两者能相遇，如果无环，那快指针先到达终点

### 代码

```python3
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if not head or not head.next:
            return False
        slow=head
        fast=head.next
        while slow!=fast:
            if fast is None or fast.next is None:
                return False
            slow=slow.next
            fast=fast.next.next
        return True

```