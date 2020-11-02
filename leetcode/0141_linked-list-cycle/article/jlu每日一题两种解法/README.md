"+++
title = "0141. Linked List Cycle [JLU每日一题]两种解法 "
author = ["HdlugJqzc5"]
date = 2020-10-09T04:23:54+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# [JLU每日一题]两种解法

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [[JLU每日一题]两种解法](https://leetcode-cn.com/problems/linked-list-cycle/solution/jlumei-ri-yi-ti-liang-chong-jie-fa-by-hdlugjqzc5/) by [HdlugJqzc5](https://leetcode-cn.com/u/HdlugJqzc5/)

# 解法1
## 思路
快慢指针,一个快一个慢总有一天会相见
## 代码
```
    bool hasCycle(ListNode *head) {
        ListNode *slow, *fast;
        int cnt = 1;
        slow = fast = head;
        while (fast) {
            if (!(cnt % 2)) 
                slow = slow->next;
            fast = fast->next;
            ++cnt;
            if (fast == slow)
                return true;
        }
        return false;
    }
```
# 解法2
## 思路
用完就拆,特殊值标记
## 代码
```
    bool hasCycle(ListNode *head) {
        while (head) {
            if (head->next == (ListNode*)-1) return true;
            ListNode *temp = head->next;
            head->next = (ListNode*)-1;
            head = temp;
        }
        return false;
    }
```

