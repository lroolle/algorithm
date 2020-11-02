"+++
title = "0141. Linked List Cycle Python 【全国最菜转圈圈】 "
author = ["flying_du"]
date = 2020-10-09T01:09:17+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# Python 【全国最菜转圈圈】

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [Python 【全国最菜转圈圈】](https://leetcode-cn.com/problems/linked-list-cycle/solution/python-quan-guo-zui-cai-zhuan-quan-quan-by-flying_/) by [flying_du](https://leetcode-cn.com/u/flying_du/)

### 解题思路
首先就是暴力转圈法。
我们就跑10000步，如果后面还能跑，说明有环。
如果突然跑到头了，就说明没有环。
### 暴力解法

```python3
class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        i = 0
        
        if head == None:
            return False

        while head.next != None:
            head = head.next
            i += 1

            if i == 10000 and head.next != None:
                return True

        return False

```
### 快慢指针解法

找两个小朋友跑，
一个一次跑一步，
一个一次跑两步。
跑两步的小朋友如果有一天突然和跑一步的小朋友跑到一个位置了，说明有环。
否则没环。

```python3
class Solution:
    def hasCycle(self, head: ListNode) -> bool:

        if head == None:
            return False

        slow = fast = head

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow is fast:
                return True

        return False
```