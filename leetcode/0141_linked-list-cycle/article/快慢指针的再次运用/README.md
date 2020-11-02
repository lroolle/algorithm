"+++
title = "0141. Linked List Cycle 快慢指针的再次运用 "
author = ["haleeey"]
date = 2020-09-29T13:01:08+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# 快慢指针的再次运用

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [快慢指针的再次运用](https://leetcode-cn.com/problems/linked-list-cycle/solution/kuai-man-zhi-zhen-de-zai-ci-yun-yong-by-haleeey/) by [haleeey](https://leetcode-cn.com/u/haleeey/)

### 解题思路
切记不可写while(p){p = p->next },原因就是如果有圈的话,就会一直循坏,直到超时才会退出程序,这次用快慢指针,如果有环,则必定会相遇.

### 代码

```c
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool hasCycle(struct ListNode *head) {
    struct ListNode *fast = head;
    struct ListNode *slow = head;
    while(fast && fast->next){
        slow = slow->next;
        fast = fast -> next -> next;
        if(slow == fast)
            return true;
    }
    return false; 

}
```