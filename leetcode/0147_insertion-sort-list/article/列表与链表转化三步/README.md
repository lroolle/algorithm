"+++
title = "0147. Insertion Sort List 列表与链表转化（三步） "
author = ["xu-chen-chen-d"]
date = 2020-08-14T07:42:34+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3", "Stack", "Sort", "Queue", "LinkedList"]
categories = ["leetcode"]
draft = false
+++

# 列表与链表转化（三步）

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [列表与链表转化（三步）](https://leetcode-cn.com/problems/insertion-sort-list/solution/lie-biao-yu-lian-biao-zhuan-hua-san-bu-by-xu-chen-/) by [xu-chen-chen-d](https://leetcode-cn.com/u/xu-chen-chen-d/)

### 解题思路
此处撰写解题思路
第一步：直接计算出链表元素个数。
第二步：把链表中val值用列表表示出来，同时按大到小排序。
第三步：按顺序构建链表，即答案值。
### 代码

```python3
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def insertionSortList(self, head: ListNode) -> ListNode:
        count = 0
        temp = head
        while temp:
            temp = temp.next
            count += 1
            #print(count,temp,head)
        
        temp2 = head
        stack = []
        for i in range(count):
            stack.append(temp2.val)
            temp2 = temp2.next
            stack.sort(reverse=True)
            #print(stack)
        

        for i in range(count):
            temp3 = ListNode(stack[i])
            temp3.next = temp
            temp = temp3
            #print(temp)
        
        return temp
            

```