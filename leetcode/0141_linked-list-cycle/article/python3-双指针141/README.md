"+++
title = "0141. Linked List Cycle Python3, 双指针，141 "
author = ["lionKing_njuer"]
date = 2020-09-09T15:28:47+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# Python3, 双指针，141

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [Python3, 双指针，141](https://leetcode-cn.com/problems/linked-list-cycle/solution/python3-shuang-zhi-zhen-141-by-littlelion_njuer/) by [lionKing_njuer](https://leetcode-cn.com/u/lionking_njuer/)

### 解题思路
双指针，一个走一步，一个走两步。

### 代码

```python3
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head):
        if not head or not head.next:   #节点不存在或下一个节点不存在
            return False
        
        slow = head
        fast = head.next                #起点错开，防止下面循环判断有问题
        
        while fast and fast.next:       #快节点和快节点下一个节点不为空
            if slow == fast:
                return True
            slow = slow.next
            fast = fast.next.next
        return False
```