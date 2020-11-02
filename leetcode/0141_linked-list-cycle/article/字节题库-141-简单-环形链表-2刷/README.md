"+++
title = "0141. Linked List Cycle 字节题库 - #141 - 简单 - 环形链表 - 2刷 "
author = ["superkakayong"]
date = 2020-08-02T03:32:28+08:00
tags = ["Leetcode", "Algorithms", "C++", "LinkedList", "TwoPointers"]
categories = ["leetcode"]
draft = false
+++

# 字节题库 - #141 - 简单 - 环形链表 - 2刷

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [字节题库 - #141 - 简单 - 环形链表 - 2刷](https://leetcode-cn.com/problems/linked-list-cycle/solution/zi-jie-ti-ku-141-jian-dan-huan-xing-lian-biao-1shu/) by [superkakayong](https://leetcode-cn.com/u/superkakayong/)

### 先放C++代码，非常简单易懂，就不在代码中写注释了哈。
- 用了快慢指针方法，快指针走两步慢指针走一步。如果有闭环，那么快慢指针一定会相遇；如果没有闭环，那么快指针一定会先走到头，直接return false就好了。

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        if(!head || !head->next) return nullptr;
        ListNode *fast = head, *slow = head;
        while(fast && fast->next)
        {
            fast = fast -> next -> next;
            slow = slow -> next;
            if(fast == slow) return true;
        }
        return false;
    }
};
```
### 执行结果截图
![image.png](https://pic.leetcode-cn.com/f6a3ef379f9eb64a91da8d7a79381dcdeff514c1d9a9352179c1c389f2e2c710-image.png)
