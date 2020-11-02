"+++
title = "0147. Insertion Sort List 双指针暴力排序 "
author = ["yi-wen-statistics"]
date = 2020-08-15T03:24:16+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 双指针暴力排序

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [双指针暴力排序](https://leetcode-cn.com/problems/insertion-sort-list/solution/shuang-zhi-zhen-bao-li-pai-xu-by-yi-wen-statistics/) by [yi-wen-statistics](https://leetcode-cn.com/u/yi-wen-statistics/)

## 解题思路
如代码

### 代码

```python3
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def insertionSortList(self, head: ListNode) -> ListNode:
        new = ListNode(-2**31) # 设定排序链表初始空间
        cur1 = head
        while cur1:
            tmp = cur1.next # 存储后继节点
            cur1.next = None # 切断
            curn = new # 第二指针
            while curn:
                if curn.next is None: # 判空
                    curn.next = cur1
                    break
                elif curn.next.val > cur1.val: # 排序
                    t = curn.next
                    curn.next = cur1
                    curn.next.next = t
                    break
                curn = curn.next
            cur1 = tmp # 这样滑动的好处可以同前继节点断开，成为一个新链表的头节点，而不是在老链表中滑动
        return new.next
```

### 复杂度分析
时间复杂度：O(N^2)
空间复杂度：O(1)