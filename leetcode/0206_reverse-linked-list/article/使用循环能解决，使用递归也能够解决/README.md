"+++
title = "0206. Reverse Linked List 使用循环能解决，使用递归也能够解决 "
author = ["wu-xian-sen-2"]
date = 2020-09-04T08:16:55+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 使用循环能解决，使用递归也能够解决

> [0206. Reverse Linked List](https://leetcode-cn.com/problems/reverse-linked-list/)
> [使用循环能解决，使用递归也能够解决](https://leetcode-cn.com/problems/reverse-linked-list/solution/shi-yong-xun-huan-neng-jie-jue-shi-yong-di-gui-ye-/) by [wu-xian-sen-2](https://leetcode-cn.com/u/wu-xian-sen-2/)

### 解题思路
1:  使用递归解决单链表反转问题
```
def reverseList(head: ListNode) -> ListNode:
    """  循环迭代 """
    pre, cur = None, head
    while cur is not None:
        tmp = cur.next
        cur.next = pre
        pre = cur
        cur = tmp
    return pre

def reverseList_01(head: ListNode) -> ListNode:
    """
    - 循环迭代2
    - 使用同时赋值
    """
    pre, cur = None, head
    while cur is not None:
        cur.next, pre, cur = pre, cur, cur.next
    return pre
```
2： 使用递归解决单链表反转问题 
- 递归层操作：循环修改相邻两个元素指针指向，使得后一个节点指针指向前一个节点
- 递归退出条件：递归推出条件其实也是使用循环时退出条件，cur is None，也就是后一个节点为空节点时退出递归。

### 代码

```python3
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:  
        """
        - 使用循环能解决，使用递归也能够解决
        - 递归层操作：循环修改相邻两个元素指针指向，使得后一个节点指针指向前一个节点
        - 递归退出条件：递归推出条件其实也是使用循环时退出条件，cur is None，也就是后一个节点为空节点时退出递归。
        """
        def helper(pre, cur):
            if cur is None:
                return pre
            cur_next = cur.next
            cur.next = pre
            return helper(cur, cur_next)
        return helper(None, head)

```