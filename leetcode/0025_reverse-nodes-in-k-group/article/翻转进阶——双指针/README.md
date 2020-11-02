"+++
title = "0025. Reverse Nodes in k-Group 翻转进阶——双指针 "
author = ["yi-wen-statistics"]
date = 2020-08-15T13:59:22+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 翻转进阶——双指针

> [0025. Reverse Nodes in k-Group](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
> [翻转进阶——双指针](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/fan-zhuan-jin-jie-shuang-zhi-zhen-by-yi-wen-statis/) by [yi-wen-statistics](https://leetcode-cn.com/u/yi-wen-statistics/)

## 解题思路
第一道hard，主要思路就是计数——逐步截断翻转拼接——补余

### 代码

```python3
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        length = 0
        cur = head
        while cur:
            cur = cur.next
            length += 1

        new = ListNode(0)
        cur2 = new
        cur1, cur3 = head, head
        s = 1
        while length >= k:
            if s == k:
                t = cur3.next
                cur3.next = None
                pre, curp = self.ReverseLinkedList(cur1)
                cur2.next = pre
                cur2 = curp
                cur1 = t
                cur3 = t
                length -= k
                s = 1
            else:
                cur3 = cur3.next
                s += 1
            
        cur2.next = cur1
        return new.next

    def ReverseLinkedList(self, l):
        curl, pre = l, None
        while curl:
            tmp = curl.next
            curl.next = pre
            pre = curl
            curl = tmp

        curp = pre
        while curp.next:
            curp = curp.next

        return pre, curp

```

### 复杂度分析
时间复杂度：O(N)
空间复杂度：O(1)